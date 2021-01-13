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
func (db *DB) InitStateTree(depth uint64, genesisStates []core.UserState) error {
	// calculate total number of leaves
	totalLeaves := math.Exp2(float64(depth))
	if int(totalLeaves) != len(genesisStates) {
		return errors.New("Depth and number of leaves do not match")
	}
	db.Logger.Debug("Attempting to init balance tree", "totalStates", totalLeaves)
	var err error

	var insertRecords []interface{}
	prevNodePath := genesisStates[0].Path
	db.Logger.Info("Num of genesis states", len(genesisStates))

	for i := 0; i < len(genesisStates); i++ {
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
		genesisStates[i].UpdatePath(path)
		insertRecords = append(insertRecords, genesisStates[i])
		prevNodePath = genesisStates[i].Path
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
		nodes, err := db.GetStatesAtDepth(i)
		if err != nil {
			return err
		}

		var nextLevelNodes []interface{}

		// iterate over 2 at a time and create next level
		for i := 0; i < len(nodes); i += 2 {
			left, err := core.HexToByteArray(nodes[i].Hash)
			if err != nil {
				return err
			}
			right, err := core.HexToByteArray(nodes[i+1].Hash)
			if err != nil {
				return err
			}
			parentHash, err := core.GetParent(left, right)
			if err != nil {
				return err
			}
			parentPath := core.GetParentPath(nodes[i].Path)
			newNode := *core.NewStateNode(parentPath, parentHash.String())
			nextLevelNodes = append(nextLevelNodes, newNode)
		}
		err = gormbulk.BulkInsert(db.Instance, nextLevelNodes, core.CHUNK_SIZE)
		if err != nil {
			db.Logger.Error("Unable to insert states to DB", "err", err)
			return errors.New("Unable to insert states")
		}
	}

	// mark the root node type correctly
	return nil
}

func (db *DB) GetStatesAtDepth(depth uint64) ([]core.UserState, error) {
	var nodes []core.UserState
	err := db.Instance.Where("level = ?", depth).Find(&nodes).Error
	return nodes, err
}

func (db *DB) UpdateState(state core.UserState) error {
	state.UpdateHash()
	siblings, err := db.GetSiblings(state.Path)
	if err != nil {
		return err
	}

	db.Logger.Debug("Updating state", "Hash", state.Hash, "Path", state.Path, "countOfSiblings", len(siblings))
	return db.storeLeaf(state, state.Path, siblings)
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

func (db *DB) storeLeaf(state core.UserState, path string, siblings []core.UserState) error {
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
		parent, err := db.GetStateByPath(core.GetParentPath(computedNode.Path))
		if err != nil {
			return err
		}
		computedNode = parent
	}
	// Store the new root
	err = db.UpdateRootNodeHashes(computedNode.HashToByteArray())
	return err
}

// StoreNode updates the nodes given the parent hash
func (db *DB) StoreNode(parentHash core.ByteArray, leftNode core.UserState, rightNode core.UserState, isLeft bool) (err error) {
	if isLeft {
		err = db.updateState(leftNode, leftNode.Path)
	} else {
		err = db.updateState(rightNode, rightNode.Path)
	}
	if err != nil {
		return err
	}
	// update the parent with the new hashes
	return db.UpdateParentWithHash(core.GetParentPath(leftNode.Path), parentHash)
}

func (db *DB) UpdateParentWithHash(pathToParent string, newHash core.ByteArray) error {
	var node core.UserState
	node.Path = pathToParent
	node.Hash = newHash.String()
	return db.updateState(node, pathToParent)
}

func (db *DB) UpdateRootNodeHashes(newRoot core.ByteArray) error {
	var root core.UserState
	root.Path = ""
	root.Hash = newRoot.String()
	return db.updateState(root, root.Path)
}

func (db *DB) AddNewPendingUserState(state core.UserState) error {
	return db.Instance.Create(&state).Error
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
	var node core.UserState
	err := db.Instance.Where("path = ?", path).Find(&node).GetErrors()
	if len(err) != 0 {
		return node, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return node, nil
}

func (db *DB) GetStateByIndex(index uint64) (state core.UserState, err error) {
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

func (db *DB) GetStateByHash(hash string) (core.UserState, error) {
	var state core.UserState
	if db.Instance.First(&state, hash).RecordNotFound() {
		return state, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return state, nil
}

func (db *DB) GetDepositSubTreeRoot(hash string, level uint64) (core.UserState, error) {
	var node core.UserState
	err := db.Instance.Where("level = ? AND hash = ?", level, hash).First(&node).Error
	if gorm.IsRecordNotFoundError(err) {
		return node, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return node, nil
}

func (db *DB) GetRoot() (core.UserState, error) {
	var node core.UserState
	err := db.Instance.Where("level = ? AND status = ?", 0, core.STATUS_INACTIVE).Find(&node).GetErrors()
	if len(err) != 0 {
		return node, core.ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v", err))
	}
	return node, nil
}

// updateState will simply replace all the changed fields
func (db *DB) updateState(node core.UserState, path string) error {
	return db.Instance.Model(&node).Where("path = ?", path).Updates(node).Error
}

func (db *DB) GetAccountCount() (int, error) {
	var count int
	err := db.Instance.Table("user_accounts").Count(&count).Error
	return count, err
}

// GetFirstEmptyAccount fetches the first empty account
func (db *DB) GetFirstEmptyAccount() (node core.UserState, err error) {
	params, err := db.GetParams()
	if err != nil {
		return node, err
	}
	expectedHash := core.DefaultHashes[params.MaxDepositSubTreeHeight]
	return db.GetStateByHash(expectedHash.String())
}

func (db *DB) DeletePendingState(accountID uint64) error {
	var state core.UserState
	if err := db.Instance.Where("account_id = ? AND status = ?", accountID, core.STATUS_PENDING).Delete(&state).Error; err != nil {
		return core.ErrRecordNotFound(fmt.Sprintf("unable to delete record for ID: %v", accountID))
	}
	return nil
}

//
// Deposit State Handling
//

func (db *DB) AttachDepositInfo(root core.ByteArray) error {
	// find all pending states
	var node core.UserState
	node.CreatedByDepositSubTree = root.String()
	return db.Instance.Model(&node).Where("status = ?", core.STATUS_PENDING).Update(&node).Error
}

func (db *DB) GetPendingStateByDepositRoot(root core.ByteArray) ([]core.UserState, error) {
	// find all states with CreatedByDepositSubTree as `root`
	var pendingStates []core.UserState
	err := db.Instance.Where("created_by_deposit_sub_tree = ? AND status = ?", root.String(), core.STATUS_PENDING).Find(&pendingStates).Error
	return pendingStates, err
}
