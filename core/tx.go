package core

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/BOPR/config"
	"github.com/BOPR/wallet"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

var VerifierWaitGroup sync.WaitGroup

// Tx represets the transaction on hubble
type Tx struct {
	DBModel
	To        uint64 `json:"to"`
	From      uint64 `json:"from"`
	Data      []byte `json:"data"`
	Signature []byte `json:"sig" gorm:"null"`
	TxHash    string `json:"hash" gorm:"not null"`
	Status    uint64 `json:"status"`
	Type      uint64 `json:"type" gorm:"not null"`
}

// NewTx creates a new transaction
func NewTx(from, to, txType uint64, sig, message []byte) Tx {
	return Tx{
		From:      from,
		To:        to,
		Data:      message,
		Signature: sig,
		Type:      txType,
	}
}

// NewPendingTx creates a new transaction
func NewPendingTx(from, to, txType uint64, sig, message []byte) Tx {
	tx := Tx{
		To:        to,
		From:      from,
		Data:      message,
		Signature: sig,
		Status:    TX_STATUS_PENDING,
		Type:      txType,
	}
	tx.AssignHash()
	return tx
}

// GetSignBytes returns the transaction data that has to be signed
func (tx Tx) GetSignBytes() (signBytes []byte) {
	return tx.Data
}

// SignTx returns the transaction data that has to be signed
func (tx *Tx) SignTx(key string, pubkey string, txBytes [32]byte) (err error) {
	privKeyBytes, err := hex.DecodeString(key)
	if err != nil {
		fmt.Println("unable to decode string", err)
		return
	}
	pubkeyBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		fmt.Println("unable to decode string", err)
		return
	}
	wallet, err := wallet.SecretToWallet(privKeyBytes, pubkeyBytes)
	if err != nil {
		return err
	}
	sig, err := wallet.Sign(txBytes[:])
	if err != nil {
		return err
	}
	tx.Signature = sig.ToBytes()
	return nil
}

// AssignHash creates a tx hash and add it to the tx
func (t *Tx) AssignHash() {
	if t.TxHash != "" {
		return
	}
	hash := rlpHash(t)
	t.TxHash = hash.String()
}

func (tx *Tx) Apply(updatedFrom, updatedTo []byte) error {
	// get fresh copy of database
	db, err := NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// begin a transaction
	mysqlTx := db.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			mysqlTx.Rollback()
		}
	}()

	// post this perform all ops on transaction
	db.Instance = mysqlTx

	// apply transaction on from account
	fromAcc, err := db.GetAccountByIndex(tx.From)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	fromAcc.Data = updatedFrom

	err = db.UpdateAccount(fromAcc)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}
	toAcc, err := DBInstance.GetAccountByIndex(tx.To)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	toAcc.Data = updatedTo

	err = db.UpdateAccount(toAcc)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	tx.UpdateStatus(TX_STATUS_PROCESSED)

	// Or commit the transaction
	mysqlTx.Commit()
	return nil
}

func (tx *Tx) ApplySingleTx(account UserAccount, updatedData []byte) error {
	account.Data = updatedData
	err := DBInstance.UpdateAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tx) String() string {
	return fmt.Sprintf("To: %v From: %v Status:%v Hash: %v Data: %v", t.To, t.From, t.Status, t.TxHash, hex.EncodeToString(t.Data))
}

// Insert tx into the DB
func (db *DB) InsertTx(t *Tx) error {
	return db.Instance.Create(t).Error
}

func (db *DB) PopTxs() (txs []Tx, err error) {
	txType, err := db.FetchTxType()
	tx := db.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return txs, err
	}
	var pendingTxs []Tx
	// select N number of transactions which are pending in mempool and
	if err := tx.Limit(config.GlobalCfg.TxsPerBatch).Order("created_at").Where(&Tx{Status: TX_STATUS_PENDING, Type: txType}).Find(&pendingTxs).Error; err != nil {
		db.Logger.Error("error while fetching pending transactions", err)
		return txs, err
	}
	db.Logger.Info("Found txs", "pendingTxs", len(pendingTxs))
	var ids []string
	for _, tx := range pendingTxs {
		ids = append(ids, tx.ID)
	}
	// update the transactions from pending to processing
	errs := tx.Table("txes").Where("id IN (?)", ids).Updates(map[string]interface{}{"status": TX_STATUS_PROCESSING}).GetErrors()
	if err != nil {
		db.Logger.Error("errors while processing transactions", errs)
		return
	}
	return pendingTxs, tx.Commit().Error
}

func (db *DB) FetchTxType() (txType uint64, err error) {
	// find out which txType has the highest count
	var maxTxType uint64
	var maxCount uint64
	txTypes := []uint64{TX_TRANSFER_TYPE}
	for _, txType := range txTypes {
		count, err := db.GetCountPerTxType(txType)
		if err != nil {
			return 0, err
		}
		if count > maxCount {
			maxTxType = txType
			maxCount = count
		}
	}
	return maxTxType, nil
}

func (db *DB) GetCountPerTxType(txType uint64) (uint64, error) {
	var count uint64
	err := db.Instance.Model(&Tx{}).Where("type = ? AND status = ?", txType, TX_STATUS_PENDING).Count(&count).Error
	return count, err
}

func (db *DB) GetTx() (tx []Tx, err error) {
	err = db.Instance.First(&tx).Error
	if err != nil {
		return tx, err
	}
	return
}

func (tx *Tx) UpdateStatus(status uint64) error {
	return DBInstance.Instance.Model(&tx).Update("status", status).Error
}

// GetVerificationData fetches all the data required to prove validity fo transaction
func (tx *Tx) GetVerificationData() (fromMerkleProof, toMerkleProof AccountMerkleProof, PDAProof PDAMerkleProof, txDBConn DB, err error) {
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return tx.CreateVerificationDataForTransfer()
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return
	}
}

func (tx *Tx) CreateVerificationDataForTransfer() (fromMerkleProof, toMerkleProof AccountMerkleProof, PDAProof PDAMerkleProof, txDBConn DB, err error) {
	VerifierWaitGroup.Add(2)
	go DBInstance.FetchPDAProofWithID(tx.From, &PDAProof)
	go DBInstance.FetchMPWithID(tx.From, &fromMerkleProof)
	toAcc, err := DBInstance.GetAccountByIndex(tx.To)
	if err != nil {
		return
	}
	var toSiblings []UserAccount
	dbCopy, _ := NewDB()
	mysqlTx := dbCopy.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			mysqlTx.Rollback()
		}
	}()
	dbCopy.Instance = mysqlTx
	VerifierWaitGroup.Wait()
	updatedFromAccountBytes, _, err := LoadedBazooka.ApplyTx(fromMerkleProof, *tx)
	if err != nil {
		return
	}

	fromMerkleProof.Account.Data = updatedFromAccountBytes
	err = dbCopy.UpdateAccount(fromMerkleProof.Account)
	if err != nil {
		return
	}
	// TODO add a check to ensure that DB copy of state matches the one returned by ApplyTransferTx
	toSiblings, err = dbCopy.GetSiblings(toAcc.Path)
	if err != nil {
		return
	}
	toMerkleProof = NewAccountMerkleProof(toAcc, toSiblings)
	return fromMerkleProof, toMerkleProof, PDAProof, dbCopy, nil
}

func rlpHash(x interface{}) (h ethCmn.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func ConcatTxs(txs [][]byte) []byte {
	var concatenatedTxs []byte
	for _, tx := range txs {
		concatenatedTxs = append(concatenatedTxs, tx[:]...)
	}
	return concatenatedTxs
}

func (db *DB) FetchPDAProofWithID(id uint64, pdaProof *PDAMerkleProof) (err error) {
	leaf, err := DBInstance.GetPDALeafByID(id)
	if err != nil {
		fmt.Println("error while getting pda leaf", err)
		return
	}
	siblings, err := DBInstance.GetPDASiblings(leaf.Path)
	if err != nil {
		fmt.Println("error while getting pda siblings", err)
		return
	}
	pdaMP := NewPDAProof(leaf.Path, leaf.PublicKey, siblings)
	*pdaProof = pdaMP
	VerifierWaitGroup.Done()
	return nil
}

func (db *DB) FetchMPWithID(id uint64, accountMP *AccountMerkleProof) (err error) {
	leaf, err := DBInstance.GetAccountByIndex(id)
	if err != nil {
		fmt.Println("error while getting leaf", err)
		return
	}
	siblings, err := DBInstance.GetSiblings(leaf.Path)
	if err != nil {
		fmt.Println("error while getting siblings", err)
		return
	}
	accMP := NewAccountMerkleProof(leaf, siblings)
	*accountMP = accMP
	VerifierWaitGroup.Done()
	return nil
}
