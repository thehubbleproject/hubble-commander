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

	// Add the deposit hash for the account
	CreatedByDepositSubTree string
}

// NewUserState creates a new user account
func NewUserState(id, status uint64, path string, data []byte) *UserState {
	newState := &UserState{
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

// NewStateNode creates a new non-terminal user account, the only this useful in this is
// Path, Status, Hash, PubkeyHash
func NewStateNode(path, hash string) *UserState {
	newUserState := &UserState{
		AccountID: ZERO,
		Path:      path,
		Status:    STATUS_INACTIVE,
		Type:      TYPE_NON_TERMINAL,
	}
	newUserState.UpdatePath(newUserState.Path)
	newUserState.Hash = hash
	return newUserState
}

// NewPendingUserState creates a new terminal user account but in pending state
// It is to be used while adding new deposits while they are not finalised
func NewPendingUserState(id uint64, data []byte) *UserState {
	newAcccount := &UserState{
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

func (acc *UserState) UpdatePath(path string) {
	acc.Path = path
	acc.Level = uint64(len(path))
}

func (acc *UserState) HashToByteArray() ByteArray {
	ba, err := HexToByteArray(acc.Hash)
	if err != nil {
		panic(err)
	}
	return ba
}

func (acc *UserState) IsActive() bool {
	return acc.Status == STATUS_ACTIVE
}

func (acc *UserState) IsCoordinator() bool {
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

func (acc *UserState) CreateAccountHash() {
	accountHash := Keccak256(acc.Data)
	acc.Hash = accountHash.String()
}

//
// Utils
//

// EmptyUserState creates a new account which has the same hash as ZeroLeaf
func EmptyUserState() UserState {
	return *NewUserState(ZERO, STATUS_INACTIVE, "", nil)
}
