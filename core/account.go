package core

import (
	"errors"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrAccAlreadyExists     = errors.New("Account already exists")
	ErrUnableToInsertLeaves = errors.New("Unable to insert leaves")
	EmptyByteSlice          = []byte{}
)

// Account is the copy of the accounts tree
type Account struct {
	ID string `json:"-" gorm:"primary_key;size:100;default:'6ba7b810-9dad-11d1-80b4-00c04fd430c8'"`
	// ID is the path of the user account in the Account Tree
	// Cannot be changed once created
	AccountID uint64 `gorm:"column:account_id"`

	// Public key for the user
	PublicKey []byte `gorm:"type:varbinary(255)"`

	// Path from root to leaf
	Path string `gorm:"not null;unique;index:Path"`

	// Type of node
	Type uint64 `gorm:"not null"`

	// keccak hash of the node
	Hash string `gorm:"not null"`

	// Level is the level of node in the tree
	Level uint64 `gorm:"not null"`
}

// BeforeCreate sets id before creating
func (acc *Account) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("id", uuid.NewV4().String())
	if err != nil {
		return err
	}
	return nil
}

// NewAccount creates a new account
func NewAccount(accID uint64, pubkey []byte, path string, nodeType uint64) (*Account, error) {
	newAccount := &Account{
		AccountID: accID,
		PublicKey: pubkey,
		Path:      path,
		Type:      nodeType,
	}
	newAccount.UpdatePath(newAccount.Path)
	err := newAccount.PopulateHash()
	if err != nil {
		return nil, err
	}
	return newAccount, nil
}

func NewAccountRoot(depth int) *Account {
	newAccount := &Account{
		AccountID: 0,
		PublicKey: []byte(""),
		Path:      "",
		Type:      TYPE_ROOT,
	}
	newAccount.Hash = DefaultHashes[depth].String()
	return newAccount

}

func NewAccountNode(path, hash string, nodeType uint64) *Account {
	newAccount := &Account{
		AccountID: ZERO,
		Path:      path,
		Type:      nodeType,
	}
	newAccount.UpdatePath(path)
	newAccount.Hash = hash
	return newAccount
}

// NewEmptyAccount creates new empty account which generates zero hash
func NewEmptyAccount() *Account {
	return &Account{AccountID: ZERO, PublicKey: EmptyByteSlice, Type: TYPE_TERMINAL}
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
	if len(p.PublicKey) == 0 {
		p.Hash = ZeroLeaf.String()
		return nil
	}
	hash, err := Pubkey(p.PublicKey).ToHash()
	if err != nil {
		return err
	}
	p.Hash = hash
	return nil
}
