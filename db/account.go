package db

import (
	"fmt"

	"github.com/BOPR/core"
	"github.com/jinzhu/gorm"
)

// InitAccountTree init account tree with all leaves
func (db *DB) InitAccountTree(depth int) error {
	account := core.NewAccountRoot(depth)
	return db.Instance.Create(&account).Error
}

// GetAccountLeafByPath gets the account of the given path from the DB
func (db *DB) GetAccountLeafByPath(path string) (core.Account, error) {
	var leaf core.Account
	err := db.Instance.Scopes(QueryByPath(path)).Find(&leaf).Error
	if err == gorm.ErrRecordNotFound {
		nodeType, err := db.FindNodeType(path)
		if err != nil {
			return leaf, err
		}
		height, err := db.DepthToHeight(len(path))
		if err != nil {
			return leaf, err
		}
		node := core.NewAccountNode(path, core.DefaultHashes[height].String(), nodeType)
		return *node, nil
	}
	if err != nil {
		return leaf, err
	}
	return leaf, nil
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
	var account *core.Account
	path, err := db.IDToPath(ID)
	if err != nil {
		return *account, err
	}
	return db.GetAccountLeafByPath(path)
}

func (db *DB) GetAccountRoot() (core.Account, error) {
	var account core.Account
	err := db.Instance.Scopes(QueryByType(core.TYPE_ROOT)).Find(&account).Error
	if err != nil {
		return account, core.ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v in core.Account tree", err))
	}
	return account, nil
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
	newAccount.Path = path
	newAccount.UpdatePath(path)

	var temp core.Account
	err := db.Instance.Model(&newAccount).Where("path = ?", path).Find(&temp).Error
	if gorm.IsRecordNotFoundError(err) {
		return db.Instance.Create(&newAccount).Error
	}
	if err != nil {
		return err
	}
	err = db.Instance.Model(&newAccount).Where("path = ?", path).Update(newAccount).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) updateParentAccountWithHash(pathToParent string, newHash core.ByteArray) error {
	// Update the root hash
	var tempAccount core.Account
	tempAccount.Path = pathToParent
	nodeType, err := db.FindNodeType(pathToParent)
	if err != nil {
		return err
	}
	tempAccount.Type = nodeType
	tempAccount.Hash = newHash.String()
	return db.updateSingleAccount(tempAccount, pathToParent)
}

func (db *DB) updateAccountRootNodes(newRoot core.ByteArray) error {
	var tempAccountLeaf core.Account
	tempAccountLeaf.Type = core.TYPE_ROOT
	tempAccountLeaf.Path = ""
	tempAccountLeaf.Hash = newRoot.String()
	return db.updateSingleAccount(tempAccountLeaf, tempAccountLeaf.Path)
}
