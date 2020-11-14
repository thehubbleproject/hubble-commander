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
	// ID is the path of the user state in the state Tree
	// Cannot be changed once created
	StateID uint64 `gorm:"not null;index:StateID"`

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
		StateID: id,
		Path:    path,
		Status:  status,
		Type:    TYPE_TERMINAL,
		Data:    data,
	}
	newState.UpdatePath(newState.Path)
	newState.CreateAccountHash()
	return newState
}

// newStateNode creates a new non-terminal user account, the only this useful in this is
// Path, Status, Hash, PubkeyHash
func newStateNode(path, hash string) *UserStateNode {
	newUserState := &UserStateNode{
		StateID: ZERO,
		Path:    path,
		Status:  STATUS_ACTIVE,
		Type:    TYPE_NON_TERMINAL,
	}
	newUserState.UpdatePath(newUserState.Path)
	newUserState.Hash = hash
	return newUserState
}

// NewPendingUserState creates a new terminal user state but in pending state
// It is to be used while adding new deposits while they are not finalised
func NewPendingUserState(id uint64, data []byte) *UserStateNode {
	state := &UserStateNode{
		StateID: id,
		Path:    UNINITIALIZED_PATH,
		Status:  STATUS_PENDING,
		Type:    TYPE_TERMINAL,
		Data:    data,
	}
	state.UpdatePath(state.Path)
	state.CreateAccountHash()
	return state
}

func (state *UserStateNode) UpdatePath(path string) {
	state.Path = path
	state.Level = uint64(len(path))
}

func (state *UserStateNode) String() string {
	id, balance, nonce, token, _ := LoadedBazooka.DecodeState(state.Data)
	return fmt.Sprintf("ID: %d Bal: %d Nonce: %d Token: %v", id, balance, nonce, token)
}

func (state *UserStateNode) ToABIAccount() (solState rollupclient.TypesUserState, err error) {
	solState.PubkeyIndex, solState.Balance, solState.Nonce, solState.TokenType, err = LoadedBazooka.DecodeState(state.Data)
	if err != nil {
		return
	}
	return
}

func (state *UserStateNode) HashToByteArray() ByteArray {
	ba, err := HexToByteArray(state.Hash)
	if err != nil {
		panic(err)
	}
	return ba
}

func (state *UserStateNode) IsActive() bool {
	return state.Status == STATUS_ACTIVE
}

func (state *UserStateNode) IsCoordinator() bool {
	if state.Path != "" {
		return false
	}

	if state.Status != 1 {
		return false
	}

	if state.Type != 0 {
		return false
	}

	return true
}

func (state *UserStateNode) CreateAccountHash() {
	state.Hash = common.Keccak256(state.Data).String()
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
func (db *DB) InitStateTree(depth uint64, genesisStates []UserStateNode) error {
	// calculate total number of leaves
	totalLeaves := math.Exp2(float64(depth))
	if int(totalLeaves) != len(genesisStates) {
		return errors.New("Depth and number of leaves do not match")
	}
	db.Logger.Debug("Attempting to init balance tree", "totalStates", totalLeaves)

	var err error

	// insert coodinator leaf
	err = db.InsertCoordinatorAccounts(&genesisStates[0], depth)
	if err != nil {
		db.Logger.Error("Unable to insert coodinator account", "err", err)
		return err
	}

	var insertRecords []interface{}
	prevNodePath := genesisStates[0].Path

	for i := 1; i < len(genesisStates); i++ {
		pathToAdjacentNode, err := GetAdjacentNodePath(prevNodePath)
		if err != nil {
			return err
		}
		genesisStates[i].UpdatePath(pathToAdjacentNode)
		insertRecords = append(insertRecords, genesisStates[i])
		prevNodePath = genesisStates[i].Path
	}

	db.Logger.Info("Creating user states, might take a minute or two, sit back.....", "count", len(insertRecords))
	err = gormbulk.BulkInsert(db.Instance, insertRecords, CHUNK_SIZE)
	if err != nil {
		db.Logger.Error("Unable to insert states to DB", "err", err)
		return errors.New("Unable to insert states")
	}

	// merkelise
	// 1. Pick all leaves at level depth
	// 2. Iterate 2 of them and create parents and store
	// 3. Persist all parents to database
	// 4. Start with next round
	for i := depth; i > 0; i-- {
		// get all leaves at depth N
		states, err := db.GetStatesAtDepth(i)
		if err != nil {
			return err
		}
		var nextLevelStates []interface{}

		// iterate over 2 at a time and create next level
		for i := 0; i < len(states); i += 2 {
			left, err := HexToByteArray(states[i].Hash)
			if err != nil {
				return err
			}
			right, err := HexToByteArray(states[i+1].Hash)
			if err != nil {
				return err
			}
			parentHash, err := GetParent(left, right)
			if err != nil {
				return err
			}
			parentPath := GetParentPath(states[i].Path)
			newAccNode := *newStateNode(parentPath, parentHash.String())
			nextLevelStates = append(nextLevelStates, newAccNode)
		}
		err = gormbulk.BulkInsert(db.Instance, nextLevelStates, CHUNK_SIZE)
		if err != nil {
			db.Logger.Error("Unable to insert states to DB", "err", err)
			return errors.New("Unable to insert states")
		}
	}

	// mark the root node type correctly
	return nil
}

func (db *DB) GetStatesAtDepth(depth uint64) ([]UserStateNode, error) {
	var states []UserStateNode
	err := db.Instance.Where("level = ?", depth).Find(&states).Error
	if err != nil {
		return states, err
	}
	return states, nil
}

func (db *DB) UpdateState(state UserStateNode) error {
	db.Logger.Info("Updated state", "PATH", state.Path)
	state.CreateAccountHash()
	siblings, err := db.GetSiblings(state.Path)
	if err != nil {
		return err
	}

	db.Logger.Debug("Updating state", "Hash", state.Hash, "Path", state.Path, "countOfSiblings", len(siblings))
	return db.StoreLeaf(state, state.Path, siblings)
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
		parentState, err := db.GetStateByPath(GetParentPath(computedNode.Path))
		if err != nil {
			return err
		}
		computedNode = parentState
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
		// update left state
		err = db.updateState(leftNode, leftNode.Path)
		if err != nil {
			return err
		}
	} else {
		// update right state
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
	var state UserStateNode
	state.Path = pathToParent
	state.Hash = newHash.String()
	return db.updateState(state, pathToParent)
}

func (db *DB) UpdateRootNodeHashes(newRoot ByteArray) error {
	var state UserStateNode
	state.Path = ""
	state.Hash = newRoot.String()
	return db.updateState(state, state.Path)
}

func (db *DB) AddNewPendingAccount(state UserStateNode) error {
	return db.Instance.Create(&state).Error
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
	var state UserStateNode
	err := db.Instance.Where("path = ?", path).Find(&state).GetErrors()
	if len(err) != 0 {
		return state, ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return state, nil
}

func (db *DB) GetStateByIndex(index uint64) (state UserStateNode, err error) {
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
	var state UserStateNode
	if db.Instance.First(&state, hash).RecordNotFound() {
		return state, ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return state, nil
}

func (db *DB) GetDepositSubTreeRoot(hash string, level uint64) (UserStateNode, error) {
	var state UserStateNode
	err := db.Instance.Where("level = ? AND hash = ?", level, hash).First(&state).Error
	if gorm.IsRecordNotFoundError(err) {
		return state, ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return state, nil
}

func (db *DB) GetRoot() (UserStateNode, error) {
	var state UserStateNode
	err := db.Instance.Where("level = ?", 0).Find(&state).GetErrors()
	if len(err) != 0 {
		return state, ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v", err))
	}
	return state, nil
}

func (db *DB) InsertCoordinatorAccounts(state *UserStateNode, depth uint64) error {
	state.UpdatePath(GenCoordinatorPath(depth))
	state.CreateAccountHash()
	state.Type = TYPE_TERMINAL
	return db.Instance.Create(&state).Error
}

// updateState will simply replace all the changed fields
func (db *DB) updateState(state UserStateNode, path string) error {
	return db.Instance.Model(&state).Where("path = ?", path).Updates(UserStateNode{StateID: state.StateID, Status: state.Status, Data: state.Data, Hash: state.Hash}).Error
}

func (db *DB) GetAccountCount() (int, error) {
	var count int
	db.Instance.Table("user_accounts").Count(&count)
	return count, nil
}

// GetFirstEmptyAccount fetches the first empty state
func (db *DB) GetFirstEmptyAccount() (state UserStateNode, err error) {
	params, err := db.GetParams()
	if err != nil {
		return state, err
	}
	expectedHash := defaultHashes[params.MaxDepositSubTreeHeight]
	return db.GetAccountByHash(expectedHash.String())
}

func (db *DB) DeletePendingAccount(ID uint64) error {
	var state UserStateNode
	if err := db.Instance.Where("state_id = ? AND status = ?", ID, STATUS_PENDING).Delete(&state).Error; err != nil {
		return ErrRecordNotFound(fmt.Sprintf("unable to delete record for ID: %v", ID))
	}
	return nil
}

//
// Deposit Account Handling
//

func (db *DB) AttachDepositInfo(root ByteArray) error {
	// find all pending accounts
	var state UserStateNode
	state.CreatedByDepositSubTree = root.String()
	if err := db.Instance.Model(&state).Where("status = ?", STATUS_PENDING).Update(&state).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetPendingAccByDepositRoot(root ByteArray) ([]UserStateNode, error) {
	// find all states with CreatedByDepositSubTree as `root`
	var pendingStates []UserStateNode
	if err := db.Instance.Where("created_by_deposit_sub_tree = ? AND status = ?", root.String(), STATUS_PENDING).Find(&pendingStates).Error; err != nil {
		return pendingStates, err
	}

	return pendingStates, nil
}
