package core

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Deposit is the user data stored on the node per user
type Deposit struct {
	ID string `json:"-" gorm:"primary_key;size:100;default:'6ba7b810-9dad-11d1-80b4-00c04fd430c8'"`

	// AccountID is the path of the pubkey in the account Tree
	AccountID uint64 `gorm:"not null;index:AccountID"`

	Data []byte `gorm:"type:varbinary(255)"`

	Hash string `gorm:"not null;index:Hash"`

	// Add the deposit root for the state
	DepositRoot string
}

// BeforeCreate sets id before creating
func (s *Deposit) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("id", uuid.NewV4().String())
	if err != nil {
		return err
	}
	return nil
}

// NewDeposit creates a new deposit
func NewDeposit(accID uint64, data []byte) *Deposit {
	hash := Keccak256(data)
	return &Deposit{
		AccountID: accID,
		Data:      data,
		Hash:      hash.String(),
	}
}
