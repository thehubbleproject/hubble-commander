package db

import "github.com/jinzhu/gorm"

func QueryByPath(path string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("path = ?", path)
	}
}
func QueryByDepositRoot(depositRoot string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deposit_root = ?", depositRoot)
	}
}
func QueryByAccountID(accountID uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("account_id = ?", accountID)
	}
}
func QueryByType(nodeType uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("type = ?", nodeType)
	}
}
func QueryByDepth(depth uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("depth = ?", depth)
	}
}
func QueryByHash(hash string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("hash = ?", hash)
	}
}
