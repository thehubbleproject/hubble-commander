package core

// UserState is the user data stored on the node per user
type UserState struct {
	// ID is the path of the user account in the account Tree
	// Cannot be changed once created
	AccountID uint64 `gorm:"not null;index:AccountID"`

	Data []byte `gorm:"type:varbinary(255)"`

	// Path from root to leaf
	// NOTE: not a part of the leaf
	// Path is a string to that we can run LIKE queries
	Path string `gorm:"index:Path"`

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

	// Add the deposit hash for the state
	CreatedByDepositSubTree string
}

// NewUserState creates a new user state
func NewUserState(id, status uint64, path string, data []byte) *UserState {
	node := &UserState{
		AccountID: id,
		Path:      path,
		Status:    status,
		Type:      TYPE_TERMINAL,
		Data:      data,
	}
	node.UpdatePath(node.Path)
	node.UpdateHash()
	return node
}

// NewStateNode creates a new non-terminal user state, the only this useful in this is
// Path, Status, Hash, PubkeyHash
func NewStateNode(path, hash string) *UserState {
	node := &UserState{
		AccountID: ZERO,
		Path:      path,
		Status:    STATUS_INACTIVE,
		Type:      TYPE_NON_TERMINAL,
	}
	node.UpdatePath(node.Path)
	node.Hash = hash
	return node
}

// NewPendingUserState creates a new terminal user state but in pending state
// It is to be used while adding new deposits while they are not finalised
func NewPendingUserState(id uint64, data []byte) *UserState {
	return NewUserState(id, STATUS_PENDING, UNINITIALIZED_PATH, data)
}

func (node *UserState) UpdatePath(path string) {
	node.Path = path
	node.Level = uint64(len(path))
}

func (node *UserState) UpdateHash() {
	node.Hash = Keccak256(node.Data).String()
}

func (node *UserState) HashToByteArray() ByteArray {
	ba, err := HexToByteArray(node.Hash)
	if err != nil {
		panic(err)
	}
	return ba
}

func (node *UserState) IsActive() bool {
	return node.Status == STATUS_ACTIVE
}

//
// Utils
//

// EmptyUserState creates a new state which has the same hash as ZeroLeaf
func EmptyUserState() UserState {
	return *NewUserState(ZERO, STATUS_INACTIVE, "", nil)
}
