package core

import (
	"errors"
	"fmt"
	"math"

	"github.com/BOPR/common"
	"github.com/BOPR/contracts/rollupclient"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// UserStateNode is the user data stored on the node per user
type UserStateNode struct {
	// ID is the path of the user account in the account Tree
	// Cannot be changed once created
	AccountID uint64 `gorm:"not null;index:AccountID"`

	Data []byte `gorm:"type:varbinary(255)" sql:"DEFAULT:0"`

	// Path from root to leaf
	// NOTE: not a part of the leaf
	// Path is a string to that we can run LIKE queries
	Path string `gorm:"not null;index:Path"`

	// Pending = 0 means has deposit but not merged to balance tree
	// Active = 1
	// InActive = 2 => non leaf node
	// NonInitialised = 100
	Status uint64 `gorm:"not null;index:Status"`

	// Type of nodes
	// 1 => terminal
	// 0 => root
	// 2 => non terminal
	Type uint64 `gorm:"not null;index:Type"`

	// keccak hash of the node
	Hash string `gorm:"not null;index:Hash"`

	Level uint64 `gorm:"not null;index:Level"`

	// Add the deposit hash for the account
	CreatedByDepositSubTree string
}

// NewUserState creates a new user account
func NewUserState(id, status uint64, path string, data []byte) *UserStateNode {
	newState := &UserStateNode{
		AccountID: id,
		Path:      path,
		Status:    status,
		Type:      TYPE_TERMINAL,
		Data:      data,
	}
	newState.UpdatePath(newState.Path)
	newState.CreateAccountHash()
	return newState
}

// newStateNode creates a new non-terminal user account, the only this useful in this is
// Path, Status, Hash, PubkeyHash
func newStateNode(path, hash string) *UserStateNode {
	newUserState := &UserStateNode{
		AccountID: ZERO,
		Path:      path,
		Status:    STATUS_ACTIVE,
		Type:      TYPE_NON_TERMINAL,
	}
	newUserState.UpdatePath(newUserState.Path)
	newUserState.Hash = hash
	return newUserState
}

// NewPendingUserState creates a new terminal user account but in pending state
// It is to be used while adding new deposits while they are not finalised
func NewPendingUserState(id uint64, data []byte) *UserStateNode {
	newAcccount := &UserStateNode{
		AccountID: id,
		Path:      UNINITIALIZED_PATH,
		Status:    STATUS_PENDING,
		Type:      TYPE_TERMINAL,
		Data:      data,
	}
	newAcccount.UpdatePath(newAcccount.Path)
	newAcccount.CreateAccountHash()
	return newAcccount
}

func (acc *UserStateNode) UpdatePath(path string) {
	acc.Path = path
	acc.Level = uint64(len(path))
}

func (s *UserStateNode) String() string {
	id, balance, nonce, token, _ := LoadedBazooka.DecodeState(s.Data)
	return fmt.Sprintf("ID: %d Bal: %d Nonce: %d Token: %v", id, balance, nonce, token)
}

func (s *UserStateNode) ToABIAccount() (solState rollupclient.TypesUserState, err error) {
	solState.PubkeyIndex, solState.Balance, solState.Nonce, solState.TokenType, err = LoadedBazooka.DecodeState(s.Data)
	if err != nil {
		return
	}
	return
}

func (acc *UserStateNode) HashToByteArray() ByteArray {
	ba, err := HexToByteArray(acc.Hash)
	if err != nil {
		panic(err)
	}
	return ba
}

func (acc *UserStateNode) IsActive() bool {
	return acc.Status == STATUS_ACTIVE
}

func (acc *UserStateNode) IsCoordinator() bool {
	if acc.Path != "" {
		return false
	}

	if acc.Status != 1 {
		return false
	}

	if acc.Type != 0 {
		return false
	}

	return true
}

func (acc *UserStateNode) CreateAccountHash() {
	accountHash := common.Keccak256(acc.Data)
	acc.Hash = accountHash.String()
}

//
// Utils
//

// EmptyUserState creates a new account which has the same hash as ZERO_VALUE_LEAF
func EmptyUserState() UserStateNode {
	return *NewUserState(ZERO, STATUS_INACTIVE, "", []byte(""))
}

//
// DB interactions for user state
//

// InitStateTree initialises the states tree
func (db *DB) InitStateTree(depth uint64, genesisAccounts []UserStateNode) error {
	// calculate total number of leaves
	totalLeaves := math.Exp2(float64(depth))
	if int(totalLeaves) != len(genesisAccounts) {
		return errors.New("Depth and number of leaves do not match")
	}
	db.Logger.Debug("Attempting to init balance tree", "totalAccounts", totalLeaves)

	var err error

	// insert coodinator leaf
	err = db.InsertCoordinatorAccounts(&genesisAccounts[0], depth)
	if err != nil {
		db.Logger.Error("Unable to insert coodinator account", "err", err)
		return err
	}

	var insertRecords []interface{}
	prevNodePath := genesisAccounts[0].Path

	for i := 1; i < len(genesisAccounts); i++ {
		pathToAdjacentNode, err := GetAdjacentNodePath(prevNodePath)
		if err != nil {
			return err
		}
		genesisAccounts[i].UpdatePath(pathToAdjacentNode)
		insertRecords = append(insertRecords, genesisAccounts[i])
		prevNodePath = genesisAccounts[i].Path
	}

	db.Logger.Info("Creating user accounts, might take a minute or two, sit back.....", "count", len(insertRecords))
	err = gormbulk.BulkInsert(db.Instance, insertRecords, CHUNK_SIZE)
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
			newAccNode := *newStateNode(parentPath, parentHash.String())
			nextLevelAccounts = append(nextLevelAccounts, newAccNode)
		}
		err = gormbulk.BulkInsert(db.Instance, nextLevelAccounts, CHUNK_SIZE)
		if err != nil {
			db.Logger.Error("Unable to insert states to DB", "err", err)
			return errors.New("Unable to insert states")
		}
	}

	// mark the root node type correctly
	return nil
}

func (db *DB) GetStatesAtDepth(depth uint64) ([]UserStateNode, error) {
	var accs []UserStateNode
	err := db.Instance.Where("level = ?", depth).Find(&accs).Error
	if err != nil {
		return accs, err
	}
	return accs, nil
}

func (db *DB) UpdateState(account UserStateNode) error {
	db.Logger.Info("Updated account", "PATH", account.Path)
	account.CreateAccountHash()
	siblings, err := db.GetSiblings(account.Path)
	if err != nil {
		return err
	}

	db.Logger.Debug("Updating account", "Hash", account.Hash, "Path", account.Path, "countOfSiblings", len(siblings))
	return db.StoreLeaf(account, account.Path, siblings)
}

func (db *DB) StoreLeaf(state UserStateNode, path string, siblings []UserStateNode) error {
	var err error
	var isLeft bool
	computedNode := state
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
			isLeft = true
			// Store the node!
			err = db.StoreNode(parentHash, computedNode, sibling, isLeft)
			if err != nil {
				return err
			}
		} else {
			parentHash, err = GetParent(sibling.HashToByteArray(), computedNode.HashToByteArray())
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
		parentAccount, err := db.GetStateByPath(GetParentPath(computedNode.Path))
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
func (db *DB) StoreNode(parentHash ByteArray, leftNode UserStateNode, rightNode UserStateNode, isLeft bool) (err error) {
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
	return db.UpdateParentWithHash(GetParentPath(leftNode.Path), parentHash)
}

func (db *DB) UpdateParentWithHash(pathToParent string, newHash ByteArray) error {
	// Update the root hash
	var tempAccount UserStateNode
	tempAccount.Path = pathToParent
	tempAccount.Hash = newHash.String()
	return db.updateState(tempAccount, pathToParent)
}

func (db *DB) UpdateRootNodeHashes(newRoot ByteArray) error {
	var tempAccount UserStateNode
	tempAccount.Path = ""
	tempAccount.Hash = newRoot.String()
	return db.updateState(tempAccount, tempAccount.Path)
}

func (db *DB) AddNewPendingAccount(acc UserStateNode) error {
	return db.Instance.Create(&acc).Error
}

func (db *DB) GetSiblings(path string) ([]UserStateNode, error) {
	var relativePath = path
	var siblings []UserStateNode
	for i := len(path); i > 0; i-- {
		otherChild := GetOtherChild(relativePath)
		otherNode, err := db.GetStateByPath(otherChild)
		if err != nil {
			return siblings, err
		}
		siblings = append(siblings, otherNode)
		relativePath = GetParentPath(relativePath)
	}

	return siblings, nil
}

// GetStateByDepth gets the state leaf of the given path from the DB
func (db *DB) GetStateByDepth(path string) (UserStateNode, error) {
	var account UserStateNode
	err := db.Instance.Where("path = ?", path).Find(&account).GetErrors()
	if len(err) != 0 {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return account, nil
}

func (db *DB) GetStateByIndex(index uint64) (acc UserStateNode, err error) {
	params, err := db.GetParams()
	if err != nil {
		return
	}
	path, err := SolidityPathToNodePath(index, params.MaxDepth)
	if err != nil {
		return
	}
	return db.GetStateByPath(path)
}

func (db *DB) GetStateByPath(path string) (UserStateNode, error) {
	var userState UserStateNode
	err := db.Instance.Where("path = ?", path).Find(&userState).GetErrors()
	if len(err) != 0 {
		return userState, ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return userState, nil
}

func (db *DB) GetAccountByHash(hash string) (UserStateNode, error) {
	var account UserStateNode
	if db.Instance.First(&account, hash).RecordNotFound() {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return account, nil
}

func (db *DB) GetDepositSubTreeRoot(hash string, level uint64) (UserStateNode, error) {
	var account UserStateNode
	err := db.Instance.Where("level = ? AND hash = ?", level, hash).First(&account).Error
	if gorm.IsRecordNotFoundError(err) {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return account, nil
}

func (db *DB) GetRoot() (UserStateNode, error) {
	var account UserStateNode
	err := db.Instance.Where("level = ?", 0).Find(&account).GetErrors()
	if len(err) != 0 {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v", err))
	}
	return account, nil
}

func (db *DB) InsertCoordinatorAccounts(acc *UserStateNode, depth uint64) error {
	acc.UpdatePath(GenCoordinatorPath(depth))
	acc.CreateAccountHash()
	acc.Type = TYPE_TERMINAL
	return db.Instance.Create(&acc).Error
}

// updateState will simply replace all the changed fields
func (db *DB) updateState(newAcc UserStateNode, path string) error {
	return db.Instance.Model(&newAcc).Where("path = ?", path).Updates(UserStateNode{AccountID: newAcc.AccountID, Status: newAcc.Status, Data: newAcc.Data, Hash: newAcc.Hash}).Error
}

func (db *DB) GetAccountCount() (int, error) {
	var count int
	db.Instance.Table("user_accounts").Count(&count)
	return count, nil
}

// GetFirstEmptyAccount fetches the first empty account
func (db *DB) GetFirstEmptyAccount() (acc UserStateNode, err error) {
	params, err := db.GetParams()
	if err != nil {
		return acc, err
	}
	expectedHash := defaultHashes[params.MaxDepositSubTreeHeight]
	return db.GetAccountByHash(expectedHash.String())
}

func (db *DB) DeletePendingAccount(ID uint64) error {
	var account UserStateNode
	if err := db.Instance.Where("account_id = ? AND status = ?", ID, STATUS_PENDING).Delete(&account).Error; err != nil {
		return ErrRecordNotFound(fmt.Sprintf("unable to delete record for ID: %v", ID))
	}
	return nil
}

//
// Deposit Account Handling
//

func (db *DB) AttachDepositInfo(root ByteArray) error {
	// find all pending accounts
	var account UserStateNode
	account.CreatedByDepositSubTree = root.String()
	if err := db.Instance.Model(&account).Where("status = ?", STATUS_PENDING).Update(&account).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetPendingAccByDepositRoot(root ByteArray) ([]UserStateNode, error) {
	// find all accounts with CreatedByDepositSubTree as `root`
	var pendingAccounts []UserStateNode
	if err := db.Instance.Where("created_by_deposit_sub_tree = ? AND status = ?", root.String(), STATUS_PENDING).Find(&pendingAccounts).Error; err != nil {
		return pendingAccounts, err
	}

	return pendingAccounts, nil
}
