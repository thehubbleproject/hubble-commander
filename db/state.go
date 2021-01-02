package db

import (
	"errors"
	"fmt"
	"math"

	"github.com/BOPR/core"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// InitStateTree initialises the states tree
func (db *DB) InitStateTree(depth uint64, genesisAccounts []core.UserState) error {
	// calculate total number of leaves
	totalLeaves := math.Exp2(float64(depth))
	if int(totalLeaves) != len(genesisAccounts) {
		return errors.New("Depth and number of leaves do not match")
	}
	db.Logger.Debug("Attempting to init balance tree", "totalAccounts", totalLeaves)
	var err error

	var insertRecords []interface{}
	prevNodePath := genesisAccounts[0].Path

	for i := 0; i < len(genesisAccounts); i++ {
		var path string
		if i == 0 {
			path, err = core.SolidityPathToNodePath(0, depth)
			if err != nil {
				return err
			}
		} else {
			path, err = core.GetAdjacentNodePath(prevNodePath)
			if err != nil {
				return err
			}
		}
		genesisAccounts[i].UpdatePath(path)
		insertRecords = append(insertRecords, genesisAccounts[i])
		prevNodePath = genesisAccounts[i].Path
	}

	db.Logger.Info("Creating user accounts, might take a minute or two, sit back.....", "count", len(insertRecords))
	err = gormbulk.BulkInsert(db.Instance, insertRecords, core.CHUNK_SIZE)
	if err != nil {
		db.Logger.Error("Unable to insert accounts to DB", "err", err)
		return errors.New("Unable to insert accounts")
	}

	// merkelise
	// 1. Pick all leaves at level depth
	// 2. Iterate 2 of them and create parents and store
	// 3. Persist all parents to database
	// 4. Start with next round
	for i := depth; i > 0; i-- {
		// get all leaves at depth N
		accs, err := db.GetStatesAtDepth(i)
		if err != nil {
			return err
		}

		var nextLevelAccounts []interface{}

		// iterate over 2 at a time and create next level
		for i := 0; i < len(accs); i += 2 {
			left, err := core.HexToByteArray(accs[i].Hash)
			if err != nil {
				return err
			}
			right, err := core.HexToByteArray(accs[i+1].Hash)
			if err != nil {
				return err
			}
			parentHash, err := core.GetParent(left, right)
			if err != nil {
				return err
			}
			parentPath := core.GetParentPath(accs[i].Path)
			newAccNode := *core.NewStateNode(parentPath, parentHash.String())
			nextLevelAccounts = append(nextLevelAccounts, newAccNode)
		}
		err = gormbulk.BulkInsert(db.Instance, nextLevelAccounts, core.CHUNK_SIZE)
		if err != nil {
			db.Logger.Error("Unable to insert states to DB", "err", err)
			return errors.New("Unable to insert states")
		}
	}

	// mark the root node type correctly
	return nil
}

func (db *DB) GetStatesAtDepth(depth uint64) ([]core.UserState, error) {
	var accs []core.UserState
	err := db.Instance.Where("level = ?", depth).Find(&accs).Error
	if err != nil {
		return accs, err
	}
	return accs, nil
}

func (db *DB) UpdateState(state core.UserState) error {
	state.CreateAccountHash()
	siblings, err := db.GetSiblings(state.Path)
	if err != nil {
		return err
	}

	db.Logger.Debug("Updating state", "Hash", state.Hash, "Path", state.Path, "countOfSiblings", len(siblings))
	return db.StoreLeaf(state, state.Path, siblings)
}

// ReserveEmptyLeaf reserve an empty leaf
func (db *DB) ReserveEmptyLeaf() (id uint64, err error) {
	var states []core.UserState
	// find empty state leaf
	if err := db.Instance.Where("type = ? AND status = ?", core.TYPE_TERMINAL, core.STATUS_INACTIVE).Find(&states).Error; err != nil {
		return 0, err
	}
	// update status to status_active
	states[1].Status = core.STATUS_ACTIVE
	if err := db.updateState(states[1], states[1].Path); err != nil {
		return 0, err
	}
	return core.StringToUint(states[1].Path)
}

func (db *DB) StoreLeaf(state core.UserState, path string, siblings []core.UserState) error {
	var err error
	var isLeft bool
	computedNode := state
	for i := 0; i < len(siblings); i++ {
		var parentHash core.ByteArray
		sibling := siblings[i]
		isComputedRightSibling := core.GetNthBitFromRight(
			path,
			i,
		)
		if isComputedRightSibling == 0 {
			parentHash, err = core.GetParent(computedNode.HashToByteArray(), sibling.HashToByteArray())
			if err != nil {
				return err
			}
			isLeft = true
			// Store the node!
			err = db.StoreNode(parentHash, computedNode, sibling, isLeft)
			if err != nil {
				return err
			}
		} else {
			parentHash, err = core.GetParent(sibling.HashToByteArray(), computedNode.HashToByteArray())
			if err != nil {
				return err
			}
			isLeft = false
			// Store the node!
			err = db.StoreNode(parentHash, sibling, computedNode, isLeft)
			if err != nil {
				return err
			}
		}
		parentAccount, err := db.GetStateByPath(core.GetParentPath(computedNode.Path))
		if err != nil {
			return err
		}
		computedNode = parentAccount
	}
	// Store the new root
	err = db.UpdateRootNodeHashes(computedNode.HashToByteArray())
	if err != nil {
		return err
	}
	return nil
}

// StoreNode updates the nodes given the parent hash
func (db *DB) StoreNode(parentHash core.ByteArray, leftNode core.UserState, rightNode core.UserState, isLeft bool) (err error) {
	if isLeft {
		// update left account
		err = db.updateState(leftNode, leftNode.Path)
		if err != nil {
			return err
		}
	} else {
		// update right account
		err = db.updateState(rightNode, rightNode.Path)
		if err != nil {
			return err
		}
	}
	// update the parent with the new hashes
	return db.UpdateParentWithHash(core.GetParentPath(leftNode.Path), parentHash)
}

func (db *DB) UpdateParentWithHash(pathToParent string, newHash core.ByteArray) error {
	// Update the root hash
	var tempAccount core.UserState
	tempAccount.Path = pathToParent
	tempAccount.Hash = newHash.String()
	return db.updateState(tempAccount, pathToParent)
}

func (db *DB) UpdateRootNodeHashes(newRoot core.ByteArray) error {
	var tempAccount core.UserState
	tempAccount.Path = ""
	tempAccount.Hash = newRoot.String()
	return db.updateState(tempAccount, tempAccount.Path)
}

func (db *DB) AddNewPendingAccount(acc core.UserState) error {
	return db.Instance.Create(&acc).Error
}

func (db *DB) GetSiblings(path string) ([]core.UserState, error) {
	var relativePath = path
	var siblings []core.UserState
	for i := len(path); i > 0; i-- {
		otherChild := core.GetOtherChild(relativePath)
		otherNode, err := db.GetStateByPath(otherChild)
		if err != nil {
			return siblings, err
		}
		siblings = append(siblings, otherNode)
		relativePath = core.GetParentPath(relativePath)
	}

	return siblings, nil
}

// GetStateByDepth gets the state leaf of the given path from the DB
func (db *DB) GetStateByDepth(path string) (core.UserState, error) {
	var account core.UserState
	err := db.Instance.Where("path = ?", path).Find(&account).GetErrors()
	if len(err) != 0 {
		return account, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return account, nil
}

func (db *DB) GetStateByIndex(index uint64) (acc core.UserState, err error) {
	params, err := db.GetParams()
	if err != nil {
		return
	}
	path, err := core.SolidityPathToNodePath(index, params.MaxDepth)
	if err != nil {
		return
	}
	return db.GetStateByPath(path)
}

func (db *DB) GetStateByPath(path string) (core.UserState, error) {
	var userState core.UserState
	err := db.Instance.Where("path = ?", path).Find(&userState).GetErrors()
	if len(err) != 0 {
		return userState, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return userState, nil
}

func (db *DB) GetAccountByHash(hash string) (core.UserState, error) {
	var account core.UserState
	if db.Instance.First(&account, hash).RecordNotFound() {
		return account, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return account, nil
}

func (db *DB) GetDepositSubTreeRoot(hash string, level uint64) (core.UserState, error) {
	var account core.UserState
	err := db.Instance.Where("level = ? AND hash = ?", level, hash).First(&account).Error
	if gorm.IsRecordNotFoundError(err) {
		return account, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return account, nil
}

func (db *DB) GetRoot() (core.UserState, error) {
	var account core.UserState
	err := db.Instance.Where("level = ?", 0).Find(&account).GetErrors()
	if len(err) != 0 {
		return account, core.ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v", err))
	}
	return account, nil
}

// updateState will simply replace all the changed fields
func (db *DB) updateState(newAcc core.UserState, path string) error {
	return db.Instance.Model(&newAcc).Where("path = ?", path).Updates(newAcc).Error
}

func (db *DB) GetAccountCount() (int, error) {
	var count int
	db.Instance.Table("user_accounts").Count(&count)
	return count, nil
}

// GetFirstEmptyAccount fetches the first empty account
func (db *DB) GetFirstEmptyAccount() (acc core.UserState, err error) {
	params, err := db.GetParams()
	if err != nil {
		return acc, err
	}
	expectedHash := core.DefaultHashes[params.MaxDepositSubTreeHeight]
	return db.GetAccountByHash(expectedHash.String())
}

func (db *DB) DeletePendingAccount(ID uint64) error {
	var account core.UserState
	if err := db.Instance.Where("account_id = ? AND status = ?", ID, core.STATUS_PENDING).Delete(&account).Error; err != nil {
		return core.ErrRecordNotFound(fmt.Sprintf("unable to delete record for ID: %v", ID))
	}
	return nil
}

//
// Deposit Account Handling
//

func (db *DB) AttachDepositInfo(root core.ByteArray) error {
	// find all pending accounts
	var account core.UserState
	account.CreatedByDepositSubTree = root.String()
	if err := db.Instance.Model(&account).Where("status = ?", core.STATUS_PENDING).Update(&account).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetPendingAccByDepositRoot(root core.ByteArray) ([]core.UserState, error) {
	// find all accounts with CreatedByDepositSubTree as `root`
	var pendingAccounts []core.UserState
	if err := db.Instance.Where("created_by_deposit_sub_tree = ? AND status = ?", root.String(), core.STATUS_PENDING).Find(&pendingAccounts).Error; err != nil {
		return pendingAccounts, err
	}

	return pendingAccounts, nil
}
