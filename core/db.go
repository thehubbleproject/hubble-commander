package core

import (
	"github.com/tendermint/tendermint/libs/log"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// global DB instance created while doing init
var DBInstance DB

type DB struct {
	Instance *gorm.DB
	Logger   log.Logger
}

// NewDB creates a new DB instance
// NOTE: it uses the configrations present in the config.toml file
// returns error if not able to open the DB
func NewDB() (DB, error) {
	if err := config.ParseAndInitGlobalConfig(""); err != nil {
		return DB{}, err
	}
	db, err := gorm.Open(config.GlobalCfg.DB, config.GlobalCfg.FormattedDBURL())
	if err != nil {
		return DB{}, err
	}
	db.LogMode(config.GlobalCfg.DBLogMode)
	// create logger
	logger := common.Logger.With("module", "DB")
	return DB{Instance: db, Logger: logger}, nil
}

func (db *DB) Close() {
	db.Instance.Close()
}
