package db

import (
	tmLog "github.com/tendermint/tendermint/libs/log"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/log"
	"github.com/globalsign/mgo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type IDB interface {
	// Account related DB functions
	// FetchSiblings(accID uint64) (accs []UserState, err error)
	GetAllAccounts() (accs []core.UserState, err error)
	GetAccount(accID uint64) (core.UserState, error)
	InsertBulkAccounts(accounts []core.UserState) error
	GetAccountCount() (int, error)

	// Tx related functions
	InsertTx(t *core.Tx) error
	PopTxs() (txs []core.Tx, err error)

	// Batch related functions
	InsertBatchInfo(root core.ByteArray, index uint64) error
	GetAllBatches() (batches []core.Batch, err error)
	GetLatestBatch() (core.Batch, error)
	GetBatchCount() (int, error)

	// common functions
	GetBatchCollection() *mgo.Collection
	GetTransactionCollection() *mgo.Collection
	GetAccountCollection() *mgo.Collection
}

// DB is the struct implementing IDB
type DB struct {
	Instance *gorm.DB
	Bazooka  bazooka.Bazooka
	Logger   tmLog.Logger
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
	logger := log.Logger.With("module", "DB")

	bz, err := bazooka.NewPreLoadedBazooka()
	if err != nil {
		return DB{}, err
	}
	return DB{Instance: db, Logger: logger, Bazooka: bz}, nil
}

func (db *DB) Close() {
	db.Instance.Close()
}
