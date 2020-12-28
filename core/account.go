package core

import (
	"errors"
)

var (
	ErrAccAlreadyExists     = errors.New("Account already exists")
	ErrUnableToInsertLeaves = errors.New("Unable to insert leaves")
	EmptyByteSlice          = []byte{}
)

// Account is the copy of the accounts tree
type Account struct {
	// ID is the path of the user account in the Account Tree
	// Cannot be changed once created
	ID uint64 `gorm:"not null;column:account_id"`

	// Public key for the user
	PublicKey []byte `gorm:"type:varbinary(255)"`

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
func NewAccount(id uint64, pubkey []byte, path string) (*Account, error) {
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

func NewAccountNode(path, hash string) *Account {
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
	return &Account{ID: ZERO, PublicKey: EmptyByteSlice, Type: TYPE_TERMINAL}
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
		p.Hash = ZERO_VALUE_LEAF.String()
		return nil
	}
	hash, err := FromBytes(p.PublicKey).ToHash()
	if err != nil {
		return err
	}
	p.Hash = hash
	return nil
}
