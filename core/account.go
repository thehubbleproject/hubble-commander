package core

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"

	"github.com/BOPR/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

var (
	ErrAccAlreadyExists     = errors.New("Account already exists")
	ErrUnableToInsertLeaves = errors.New("Unable to insert leaves")
)

// Account is the copy of the accounts tree
type Account struct {
	// ID is the path of the user account in the Account Tree
	// Cannot be changed once created
	ID uint64 `gorm:"not null;column:account_id"`

	// Public key for the user
	PublicKey string `gorm:"type:varchar(1000)"`

	// Path from root to leaf
	// Path is a string to that we can run LIKE queries
	Path string `gorm:"not null;index:Path"`

	// Type of node
	Type uint64 `gorm:"not null"`

	// keccak hash of the node
	Hash string `gorm:"not null"`

	// Level is the level of node in the tree
	Level uint64 `gorm:"not null"`
}

// NewAccount creates a new account
func NewAccount(id uint64, pubkey, path string) (*Account, error) {
	newAccount := &Account{
		ID:        id,
		PublicKey: pubkey,
		Path:      path,
		Type:      TYPE_TERMINAL,
	}
	err := newAccount.PopulateHash()
	if err != nil {
		return nil, err
	}
	return newAccount, nil
}

func newAccountNode(path, hash string) *Account {
	newAccount := &Account{
		ID:   ZERO,
		Path: path,
		Type: TYPE_NON_TERMINAL,
	}
	newAccount.UpdatePath(path)
	newAccount.Hash = hash
	return newAccount
}

// NewEmptyAccount creates new empty account which generates zero hash
func NewEmptyAccount() *Account {
	return &Account{ID: ZERO, PublicKey: "", Type: TYPE_TERMINAL}
}

func (p *Account) UpdatePath(path string) {
	p.Path = path
	p.Level = uint64(len(path))
}

func (p *Account) HashToByteArray() ByteArray {
	ba, err := HexToByteArray(p.Hash)
	if err != nil {
		panic(err)
	}
	return ba
}

func (p *Account) PopulateHash() error {
	if p.PublicKey == "" {
		p.Hash = ZERO_VALUE_LEAF.String()
		return nil
	}
	bz, err := encodePubkey(p.PublicKey)
	if err != nil {
		return err
	}
	p.Hash = common.Keccak256(bz).String()
	return nil
}

// ------------------------------------------------------------------------------------- DB -------------------------------------------------------------------------------------

// GetAccount gets the account of the given path from the DB
func (db *DB) GetAccountLeafByPath(path string) (Account, error) {
	var pdaLeaf Account
	err := db.Instance.Where("path = ?", path).Find(&pdaLeaf).GetErrors()
	if len(err) != 0 {
		return pdaLeaf, ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return pdaLeaf, nil
}

// GetAccountByPubkey gets the account of the given pubkey
func (db *DB) GetAccountByPubkey(pubkey string) (Account, error) {
	var leaf Account
	err := db.Instance.Where("public_key = ?", pubkey).Find(&leaf).GetErrors()
	if len(err) != 0 {
		return leaf, ErrRecordNotFound(fmt.Sprintf("unable to find record for pubkey: %v err:%v", leaf, err))
	}
	return leaf, nil
}

func (db *DB) GetAccountLeafByID(ID uint64) (Account, error) {
	var account Account
	if err := db.Instance.Where("account_id = ?", ID).Find(&account).Error; err != nil {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for ID: %v in Account tree", ID))
	}
	return account, nil
}

func (db *DB) GetAccountRoot() (Account, error) {
	var account Account
	err := db.Instance.Where("level = ?", 0).Find(&account).GetErrors()
	if len(err) != 0 {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v in Account tree", err))
	}
	return account, nil
}

// GetAccountByDepth fetches all accounts at a level
func (db *DB) GetAccountByDepth(depth uint64) ([]Account, error) {
	var accs []Account
	err := db.Instance.Where("level = ?", depth).Find(&accs).Error
	if err != nil {
		return accs, err
	}
	return accs, nil
}

func (db *DB) AddNewAccount(acc Account) error {
	return db.UpdateAccount(acc)
}

// UpdateAccount updates the account
func (db *DB) UpdateAccount(leaf Account) error {
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
func (db *DB) GetAccountSiblings(path string) ([]Account, error) {
	var relativePath = path
	var siblings []Account
	for i := len(path); i > 0; i-- {
		otherChild := GetOtherChild(relativePath)
		otherNode, err := db.GetAccountLeafByPath(otherChild)
		if err != nil {
			return siblings, err
		}
		siblings = append(siblings, otherNode)
		relativePath = GetParentPath(relativePath)
	}
	return siblings, nil
}

func (db *DB) storeAccountLeaf(pdaLeaf Account, path string, siblings []Account) error {
	var err error
	computedNode := pdaLeaf
	for i := 0; i < len(siblings); i++ {
		var parentHash ByteArray
		sibling := siblings[i]
		isComputedRightSibling := GetNthBitFromRight(
			path,
			i,
		)
		if isComputedRightSibling == 0 {
			parentHash, err = GetParent(computedNode.HashToByteArray(), sibling.HashToByteArray())
			if err != nil {
				return err
			}

			err = db.storeAccountNode(parentHash, computedNode, sibling)
			if err != nil {
				return err
			}

		} else {

			parentHash, err = GetParent(sibling.HashToByteArray(), computedNode.HashToByteArray())
			if err != nil {
				return err
			}

			err = db.storeAccountNode(parentHash, sibling, computedNode)
			if err != nil {
				return err
			}

		}

		parentAccount, err := db.GetAccountLeafByPath(GetParentPath(computedNode.Path))
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

// InsertCoordinatorPubkeyAccounts inserts the coordinator accounts
func (db *DB) InsertCoordinatorPubkeyAccounts(coordinatorAccount *Account, depth uint64) error {
	coordinatorAccount.UpdatePath(GenCoordinatorPath(depth))
	err := coordinatorAccount.PopulateHash()
	if err != nil {
		return err
	}
	coordinatorAccount.Type = TYPE_TERMINAL
	return db.Instance.Create(&coordinatorAccount).Error
}

// InitAccountTree init account tree with all leaves
func (db *DB) InitAccountTree(depth uint64, genesisAccount []Account) error {
	// calculate total number of leaves
	totalLeaves := math.Exp2(float64(depth))
	if int(totalLeaves) != len(genesisAccount) {
		return errors.New("Depth and number of leaves do not match")
	}

	db.Logger.Debug("Attempting to init Account tree", "totalAccounts", totalLeaves)
	var err error

	// insert coodinator leaf
	err = db.InsertCoordinatorPubkeyAccounts(&genesisAccount[0], depth)
	if err != nil {
		db.Logger.Error("Unable to insert coodinator account", "err", err)
		return err
	}

	var insertRecords []interface{}
	prevNodePath := genesisAccount[0].Path
	for i := 1; i < len(genesisAccount); i++ {
		pathToAdjacentNode, err := GetAdjacentNodePath(prevNodePath)
		if err != nil {
			return err
		}
		genesisAccount[i].UpdatePath(pathToAdjacentNode)
		insertRecords = append(insertRecords, genesisAccount[i])
		prevNodePath = genesisAccount[i].Path
	}

	db.Logger.Info("Creating Account tree, might take a minute or two, sit back.....", "count", len(insertRecords))

	err = gormbulk.BulkInsert(db.Instance, insertRecords, CHUNK_SIZE)
	if err != nil {
		db.Logger.Error("Unable to insert leaves to DB", "err", err)
		return ErrUnableToInsertLeaves
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
			left, err := HexToByteArray(accs[i].Hash)
			if err != nil {
				return err
			}
			right, err := HexToByteArray(accs[i+1].Hash)
			if err != nil {
				return err
			}
			parentHash, err := GetParent(left, right)
			if err != nil {
				return err
			}
			parentPath := GetParentPath(accs[i].Path)
			newAccNode := newAccountNode(parentPath, parentHash.String())
			nextLevelAccounts = append(nextLevelAccounts, newAccNode)
		}
		err = gormbulk.BulkInsert(db.Instance, nextLevelAccounts, CHUNK_SIZE)
		if err != nil {
			db.Logger.Error("Unable to insert Account leaves to DB", "err", err)
			return errors.New("Unable to insert Account leaves")
		}
	}
	// mark the root node type correctly
	return nil
}

// storeAccountNode updates the nodes given the parent hash
func (db *DB) storeAccountNode(parentHash ByteArray, leftNode, rightNode Account) (err error) {
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
	return db.updateParentAccountWithHash(GetParentPath(leftNode.Path), parentHash)
}

// updateSingleAccount will simply replace all the changed fields
func (db *DB) updateSingleAccount(newAccount Account, path string) error {
	return db.Instance.Model(&newAccount).Where("path = ?", path).Update(newAccount).Error
}

func (db *DB) updateParentAccountWithHash(pathToParent string, newHash ByteArray) error {
	// Update the root hash
	var tempAccount Account
	tempAccount.Path = pathToParent
	tempAccount.Hash = newHash.String()
	return db.updateSingleAccount(tempAccount, pathToParent)
}

func (db *DB) updateAccountRootNodes(newRoot ByteArray) error {
	var tempAccountLeaf Account
	tempAccountLeaf.Path = ""
	tempAccountLeaf.Hash = newRoot.String()
	return db.updateSingleAccount(tempAccountLeaf, tempAccountLeaf.Path)
}
func encodePubkey(pubkey string) ([]byte, error) {
	pubkeyBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		panic(err)
	}
	uint256Ty, err := abi.NewType("bytes", "bytes", nil)
	if err != nil {
		return []byte(""), err
	}

	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
	}

	bytes, err := arguments.Pack(
		pubkeyBytes,
	)

	if err != nil {
		return []byte(""), err
	}

	return bytes, nil
}
