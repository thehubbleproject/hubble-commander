package migrations

import (
	types "github.com/BOPR/core"
	dbI "github.com/BOPR/db"
	"github.com/jinzhu/gorm"
)

// NOTE: The order is really important here in both up and down migrations
func init() {
	m := &Migration{
		ID: "1585168847",
		Up: func(db *gorm.DB) error {
			if !db.HasTable(&types.Tx{}) {
				db.CreateTable(&types.Tx{})
			}
			if !db.HasTable(&types.Batch{}) {
				db.CreateTable(&types.Batch{})
			}
			if !db.HasTable(&types.Commitment{}) {
				db.CreateTable(&types.Commitment{})
			}
			if !db.HasTable(&types.SyncStatus{}) {
				db.CreateTable(&types.SyncStatus{})
			}
			if !db.HasTable(&types.Params{}) {
				db.CreateTable(&types.Params{})
			}
			if !db.HasTable(&dbI.Token{}) {
				db.CreateTable(&dbI.Token{})
			}
			if !db.HasTable(&types.Account{}) {
				db.CreateTable(&types.Account{})
			}
			if !db.HasTable(&types.Deposit{}) {
				db.CreateTable(&types.Deposit{})
			}
			if !db.HasTable(&types.UserState{}) {
				db.CreateTable(&types.UserState{})
			}
			if !db.HasTable(&dbI.RelayPacket{}) {
				db.CreateTable(&dbI.RelayPacket{})
			}
			return nil
		},
		Down: func(db *gorm.DB) error {
			db.DropTableIfExists(&types.Tx{})
			db.DropTableIfExists(&types.Batch{})
			db.DropTableIfExists(&types.Commitment{})
			db.DropTableIfExists(&types.Params{})
			db.DropTableIfExists(&types.SyncStatus{})
			db.DropTableIfExists(&dbI.Token{})
			db.DropTableIfExists(&types.Account{})
			db.DropTableIfExists(&types.Deposit{})
			db.DropTableIfExists(&types.UserState{})
			db.DropTableIfExists(&dbI.RelayPacket{})
			return nil
		},
	}

	// add migration to list
	addMigration(m)
}
