package tests_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
	"github.com/BOPR/migrations"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func setupDB() (db core.DB, cleanup func(), err error) {
	tmpfile, err := ioutil.TempFile("", "test.*.db")
	println(tmpfile.Name())
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
	m.Migrate()
	cleanup = func() {
		println("Closing DB and removing file")
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

	println("Testing start")

	tx1 := core.NewTx(1, 2, 1, []byte{00}, []byte{00})
	tx2 := core.NewPendingTx(1, 2, 1, []byte{00}, []byte{00})
	tx3 := core.NewPendingTx(1, 2, 1, []byte{00}, []byte{00})

	err = db.InsertTx(&tx1)
	if err != nil {
		t.Errorf("PopTxs error %s", err)
	}
	err = db.InsertTx(&tx2)
	if err != nil {
		t.Errorf("PopTxs error %s", err)
	}
	err = db.InsertTx(&tx3)
	if err != nil {
		t.Errorf("PopTxs error %s", err)
	}

	txs, err := db.PopTxs()
	if err != nil {
		t.Errorf("PopTxs error %s", err)
	}
	assert.Equal(t, []core.Tx{tx2, tx3}, txs)
}
