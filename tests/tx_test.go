package tests_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/migrations"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func setupDB() (db core.DB, cleanup func(), err error) {
	tmpfile, err := ioutil.TempFile("", "test.*.db")
	if err != nil {
		return
	}

	sqliteDb, err := gorm.Open("sqlite3", tmpfile.Name())
	if err != nil {
		return
	}
	logger := common.Logger.With("module", "tests")
	db = core.DB{Instance: sqliteDb, Logger: logger}

	allMigrations := migrations.GetMigrations()
	m := migrations.NewGormigrate(db.Instance, migrations.DefaultOptions, allMigrations)
	err = m.Migrate()
	if err != nil {
		return
	}
	cleanup = func() {
		db.Close()
		os.Remove(tmpfile.Name())
	}

	return db, cleanup, nil

}

func TestPopTx(t *testing.T) {
	db, cleanup, err := setupDB()
	if err != nil {
		t.Errorf("setupDB error %s", err)
	}
	defer cleanup()

	var txType uint64 = 1
	config.GlobalCfg.TxsPerBatch = 2

	tx1 := core.NewTx(1, 2, txType, []byte{00}, []byte{00})
	tx2, err := core.NewPendingTx(1, 2, txType, []byte{00}, []byte{01})
	if err != nil {
		t.Errorf("PopTxs error %s", err)
	}
	tx3, err := core.NewPendingTx(1, 2, txType, []byte{00}, []byte{02})
	if err != nil {
		t.Errorf("PopTxs error %s", err)
	}

	if err = db.InsertTx(&tx1); err != nil {
		t.Errorf("PopTxs error %s", err)
	}
	if err = db.InsertTx(&tx2); err != nil {
		t.Errorf("PopTxs error %s", err)
	}
	if err = db.InsertTx(&tx3); err != nil {
		t.Errorf("PopTxs error %s", err)
	}

	fetchedTxType, err := db.FetchTxType()
	assert.Equal(t, txType, fetchedTxType)

	txs, err := db.PopTxs()
	if err != nil {
		t.Errorf("PopTxs error %s", err)
	}
	for i, tx := range []core.Tx{tx2, tx3} {
		assert.Equal(t, tx.TxHash, txs[i].TxHash)
	}
}
