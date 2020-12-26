package db

import (
	"errors"
	"fmt"
	"math"

	"github.com/BOPR/core"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// GetAccount gets the account of the given path from the DB
func (db *DB) GetAccountLeafByPath(path string) (core.Account, error) {
	var pdaLeaf core.Account
	err := db.Instance.Where("path = ?", path).Find(&pdaLeaf).GetErrors()
	if len(err) != 0 {
		return pdaLeaf, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return pdaLeaf, nil
}

// GetAccountByPubkey gets the account of the given pubkey
func (db *DB) GetAccountByPubkey(pubkey []byte) (core.Account, error) {
	var leaf core.Account
	err := db.Instance.Where("public_key = ?", pubkey).Find(&leaf).GetErrors()
	if len(err) != 0 {
		return leaf, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for pubkey: %v err:%v", leaf, err))
	}
	return leaf, nil
}

func (db *DB) GetAccountLeafByID(ID uint64) (core.Account, error) {
	var account core.Account
	if err := db.Instance.Where("account_id = ?", ID).Find(&account).Error; err != nil {
		return account, core.ErrRecordNotFound(fmt.Sprintf("unable to find record for ID: %v in core.Account tree", ID))
	}
	return account, nil
}

func (db *DB) GetAccountRoot() (core.Account, error) {
	var account core.Account
	err := db.Instance.Where("level = ?", 0).Find(&account).GetErrors()
	if len(err) != 0 {
		return account, core.ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v in core.Account tree", err))
	}
	return account, nil
}

// GetAccountByDepth fetches all accounts at a level
func (db *DB) GetAccountByDepth(depth uint64) ([]core.Account, error) {
	var accs []core.Account
	err := db.Instance.Where("level = ?", depth).Find(&accs).Error
	if err != nil {
		return accs, err
	}
	return accs, nil
}

func (db *DB) AddNewAccount(acc core.Account) error {
	return db.UpdateAccount(acc)
}

// UpdateAccount updates the account
func (db *DB) UpdateAccount(leaf core.Account) error {
	err := leaf.PopulateHash()
	if err != nil {
		return err
	}
	siblings, err := db.GetAccountSiblings(leaf.Path)
	if err != nil {
		return err
	}

	db.Logger.Debug("Updating account", "Hash", leaf.Hash, "Path", leaf.Path, "countOfSiblings", len(siblings))
	return db.storeAccountLeaf(leaf, leaf.Path, siblings)
}

// GetAccountSiblings fetches siblings for a node
func (db *DB) GetAccountSiblings(path string) ([]core.Account, error) {
	var relativePath = path
	var siblings []core.Account
	for i := len(path); i > 0; i-- {
		otherChild := core.GetOtherChild(relativePath)
		otherNode, err := db.GetAccountLeafByPath(otherChild)
		if err != nil {
			return siblings, err
		}
		siblings = append(siblings, otherNode)
		relativePath = core.GetParentPath(relativePath)
	}
	return siblings, nil
}

func (db *DB) storeAccountLeaf(pdaLeaf core.Account, path string, siblings []core.Account) error {
	var err error
	computedNode := pdaLeaf
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

			err = db.storeAccountNode(parentHash, computedNode, sibling)
			if err != nil {
				return err
			}

		} else {

			parentHash, err = core.GetParent(sibling.HashToByteArray(), computedNode.HashToByteArray())
			if err != nil {
				return err
			}

			err = db.storeAccountNode(parentHash, sibling, computedNode)
			if err != nil {
				return err
			}

		}

		parentAccount, err := db.GetAccountLeafByPath(core.GetParentPath(computedNode.Path))
		if err != nil {
			return err
		}

		computedNode = parentAccount
	}

	// Store the new root
	err = db.updateAccountRootNodes(computedNode.HashToByteArray())
	if err != nil {
		return err
	}

	return nil
}

// InitAccountTree init account tree with all leaves
func (db *DB) InitAccountTree(depth uint64, genesisAccount []core.Account) error {
	// calculate total number of leaves
	totalLeaves := math.Exp2(float64(depth))
	if int(totalLeaves) != len(genesisAccount) {
		return errors.New("Depth and number of leaves do not match")
	}

	db.Logger.Debug("Attempting to init core.Account tree", "totalAccounts", totalLeaves)
	var err error

	var insertRecords []interface{}
	prevNodePath := genesisAccount[0].Path
	for i := 0; i < len(genesisAccount); i++ {
		var path string
		if i == 0 {
			path, err = core.SolidityPathToNodePath(0, depth)
			if err != nil {
				return err
			}
			prevNodePath = path
		} else {
			path, err = core.GetAdjacentNodePath(prevNodePath)
			if err != nil {
				return err
			}
		}
		genesisAccount[i].UpdatePath(path)
		insertRecords = append(insertRecords, genesisAccount[i])
		prevNodePath = genesisAccount[i].Path
	}

	db.Logger.Info("Creating core.Account tree, might take a minute or two, sit back.....", "count", len(insertRecords))

	err = gormbulk.BulkInsert(db.Instance, insertRecords, core.CHUNK_SIZE)
	if err != nil {
		db.Logger.Error("Unable to insert leaves to DB", "err", err)
		return core.ErrUnableToInsertLeaves
	}

	// merkelise
	// 1. Pick all leaves at level depth
	// 2. Iterate 2 of them and create parents and store
	// 3. Persist all parents to database
	// 4. Start with next round
	for j := depth; j > 0; j-- {
		// get all leaves at depth N
		accs, err := db.GetAccountByDepth(j)
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
			newAccNode := core.NewAccountNode(parentPath, parentHash.String())
			nextLevelAccounts = append(nextLevelAccounts, newAccNode)
		}
		err = gormbulk.BulkInsert(db.Instance, nextLevelAccounts, core.CHUNK_SIZE)
		if err != nil {
			db.Logger.Error("Unable to insert core.Account leaves to DB", "err", err)
			return errors.New("Unable to insert core.Account leaves")
		}
	}
	// mark the root node type correctly
	return nil
}

// storeAccountNode updates the nodes given the parent hash
func (db *DB) storeAccountNode(parentHash core.ByteArray, leftNode, rightNode core.Account) (err error) {
	// update left account
	err = db.updateSingleAccount(leftNode, leftNode.Path)
	if err != nil {
		return err
	}
	// update right account
	err = db.updateSingleAccount(rightNode, rightNode.Path)
	if err != nil {
		return err
	}
	// update the parent with the new hashes
	return db.updateParentAccountWithHash(core.GetParentPath(leftNode.Path), parentHash)
}

// updateSingleAccount will simply replace all the changed fields
func (db *DB) updateSingleAccount(newAccount core.Account, path string) error {
	return db.Instance.Model(&newAccount).Where("path = ?", path).Update(newAccount).Error
}

func (db *DB) updateParentAccountWithHash(pathToParent string, newHash core.ByteArray) error {
	// Update the root hash
	var tempAccount core.Account
	tempAccount.Path = pathToParent
	tempAccount.Hash = newHash.String()
	return db.updateSingleAccount(tempAccount, pathToParent)
}

func (db *DB) updateAccountRootNodes(newRoot core.ByteArray) error {
	var tempAccountLeaf core.Account
	tempAccountLeaf.Path = ""
	tempAccountLeaf.Hash = newRoot.String()
	return db.updateSingleAccount(tempAccountLeaf, tempAccountLeaf.Path)
}
