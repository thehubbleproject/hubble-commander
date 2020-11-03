package core

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/BOPR/common"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func TestPopTx(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	db, err := gorm.Open("sqlite", sqlite.Open("gorm.db"))
	logger := common.Logger.With("module", "DB")
	txDB := DB{Instance: db, Logger: logger}

	tx1 := NewTx(1, 2, 1, []byte{00}, []byte{00})
	tx2 := NewPendingTx(1, 2, 1, []byte{00}, []byte{00})
	tx3 := NewPendingTx(1, 2, 1, []byte{00}, []byte{00})

	txDB.InsertTx(&tx1)
	txDB.InsertTx(&tx2)
	txDB.InsertTx(&tx3)

	txs, err := txDB.PopTxs()
	if err != nil {
		t.Errorf("PopTxs error %s", err)
	}
	assert.Equal(t, txs, []Tx{tx2, tx3})
}
