package core

import (
	"errors"
)

var (
	ErrUnableToInsertLeaves = errors.New("Unable to insert leaves")
)

// Account is the copy of the accounts tree
type Account struct {
	// ID is the path of the user account in the Account Tree
	// Cannot be changed once created
	ID uint64 `gorm:"column:account_id"`

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
	node := &Account{
		ID:        id,
		PublicKey: pubkey,
		Path:      path,
		Type:      TYPE_TERMINAL,
	}
	err := node.PopulateHash()
	return node, err
}

func NewAccountNode(path, hash string) *Account {
	node := &Account{
		ID:   ZERO,
		Path: path,
		Type: TYPE_NON_TERMINAL,
	}
	node.UpdatePath(path)
	node.Hash = hash
	return node
}

// NewEmptyAccount creates new empty account which generates zero hash
func NewEmptyAccount() *Account {
	return &Account{ID: ZERO, PublicKey: []byte{}, Type: TYPE_TERMINAL}
}

func (node *Account) UpdatePath(path string) {
	node.Path = path
	node.Level = uint64(len(path))
}

func (node *Account) HashToByteArray() ByteArray {
	ba, err := HexToByteArray(node.Hash)
	if err != nil {
		panic(err)
	}
	return ba
}

func (node *Account) PopulateHash() error {
	if len(node.PublicKey) == 0 {
		node.Hash = ZeroLeaf.String()
		return nil
	}
	hash, err := Pubkey(node.PublicKey).ToHash()
	if err != nil {
		return err
	}
	node.Hash = hash
	return nil
}
