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
func QueryBySubtreeID(subtreeID uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("subtree_id = ?", subtreeID)
	}
}
func QueryByAccountID(accountID uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("account_id = ?", accountID)
	}
}
func QueryByID(id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
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
func QueryByTxHash(hash string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("tx_hash = ?", hash)
	}
}
func QueryByBatchID(id uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("batch_id = ?", id)
	}
}
