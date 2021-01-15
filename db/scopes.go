package db

import "github.com/jinzhu/gorm"

func QueryByPath(path string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("path = ?", path)
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
func QueryByLevel(level uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("level = ?", level)
	}
}
func QueryByStatus(status uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}
func QueryByHash(hash string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("hash = ?", hash)
	}
}
