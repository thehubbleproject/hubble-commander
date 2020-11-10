package core

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/wallet"
	"github.com/kilic/bn254/bls"
)

var (
<<<<<<< HEAD
	ErrNoTxsFound          = errors.New("no tx found")
	ErrSignatureNotPresent = errors.New("signature not present")
=======
	ErrNoTxsFound = errors.New("no tx found")
>>>>>>> 19b9212... revive sync
)

const (
	COMMITMENT_SIZE = 1
)

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
func NewTx(from, to, txType uint64, sig, data []byte) Tx {
	return Tx{
		From:      from,
		To:        to,
		Data:      data,
		Signature: sig,
		Type:      txType,
	}
}

// NewPendingTx creates a new transaction
<<<<<<< HEAD
func NewPendingTx(from, to, txType uint64, sig, data []byte) (tx Tx, err error) {
	tx = Tx{
=======
func NewPendingTx(from, to, txType uint64, sig, data []byte) Tx {
	tx := Tx{
>>>>>>> 19b9212... revive sync
		To:        to,
		From:      from,
		Data:      data,
		Signature: sig,
		Status:    TX_STATUS_PENDING,
		Type:      txType,
	}

	if err = tx.AssignHash(); err != nil {
		return
	}
	return
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
func (t *Tx) AssignHash() (err error) {
	if t.TxHash != "" {
		return nil
	}
	hash, err := common.RlpHash(t)
	if err != nil {
		return
	}
	t.TxHash = hash.String()
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
func (tx *Tx) GetVerificationData() (fromMerkleProof, toMerkleProof StateMerkleProof, txDBConn DB, err error) {
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return tx.GetWitnessTranfer()
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return
	}
}

func (tx *Tx) GetWitnessTranfer() (fromMerkleProof, toMerkleProof StateMerkleProof, txDBConn DB, err error) {
	dbCopy, _ := NewDB()

	// fetch from state MP
	err = DBInstance.FetchMPWithID(tx.From, &fromMerkleProof)
	if err != nil {
		return
	}
	toState, err := DBInstance.GetStateByIndex(tx.To)
	if err != nil {
		return
	}
	var toSiblings []UserState

	newFrom, newTo, err := LoadedBazooka.ApplyTx(fromMerkleProof.State.Data, toState.Data, *tx)
	if err != nil {
		return
	}

	mysqlTx := dbCopy.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			mysqlTx.Rollback()
		}
	}()
	dbCopy.Instance = mysqlTx

	// apply the new from leaf
	currentFromStateCopy := fromMerkleProof.State
	currentFromStateCopy.Data = newFrom
	err = dbCopy.UpdateState(currentFromStateCopy)
	if err != nil {
		return
	}

	// create witness for to leaf
	toSiblings, err = dbCopy.GetSiblings(toState.Path)
	if err != nil {
		return
	}
	toMerkleProof = NewStateMerkleProof(toState, toSiblings)

	// apply the new to leaf
	toState.Data = newTo
	err = dbCopy.UpdateState(toState)
	if err != nil {
		return
	}

	return fromMerkleProof, toMerkleProof, dbCopy, nil
}

func ConcatTxs(txs [][]byte) []byte {
	var concatenatedTxs []byte
	for _, tx := range txs {
		concatenatedTxs = append(concatenatedTxs, tx[:]...)
	}
	return concatenatedTxs
}

func (db *DB) FetchAccountProofWithID(id uint64, pdaProof *AccountMerkleProof) (err error) {
	leaf, err := DBInstance.GetAccountLeafByID(id)
	if err != nil {
		fmt.Println("error while getting pda leaf", err)
		return
	}
	siblings, err := DBInstance.GetAccountSiblings(leaf.Path)
	if err != nil {
		fmt.Println("error while getting pda siblings", err)
		return
	}
	pdaMP := NewAccountMerkleProof(leaf.Path, leaf.PublicKey, siblings)
	*pdaProof = pdaMP
	return nil
}

func (db *DB) FetchMPWithID(id uint64, accountMP *StateMerkleProof) (err error) {
	leaf, err := DBInstance.GetStateByIndex(id)
	if err != nil {
		fmt.Println("error while getting leaf", err)
		return
	}
	siblings, err := DBInstance.GetSiblings(leaf.Path)
	if err != nil {
		fmt.Println("error while getting siblings", err)
		return
	}
	accMP := NewStateMerkleProof(leaf, siblings)
	*accountMP = accMP
	return nil
}

<<<<<<< HEAD
// Validate creates proofs for validating txs and returns new root post validation
=======
// ValidateTx creates proofs for validating txs and returns new root post validation
>>>>>>> 19b9212... revive sync
func (tx *Tx) Validate(bz Bazooka, currentRoot ByteArray) (newRoot ByteArray, err error) {
	fromStateProof, toStateProof, txDBConn, err := tx.GetVerificationData()
	if err != nil {
		return
	}

	newRoot, err = bz.ProcessTx(currentRoot, *tx, fromStateProof, toStateProof)
	if err != nil {
		if txDBConn.Instance != nil {
			txDBConn.Instance.Rollback()
			txDBConn.Close()
		}
		return
	}

	err = tx.authenticate(bz)
	if err != nil {
		txDBConn.Instance.Rollback()
		txDBConn.Close()
		return
	}

	if txDBConn.Instance != nil {
		txDBConn.Instance.Commit()
		txDBConn.Close()
	}

	return
}

func (tx *Tx) authenticate(bz Bazooka) error {
	accID, _, _, _, err := bz.DecodeState(tx.Data)
	if err != nil {
		return err
	}

	acc, err := DBInstance.GetAccountLeafByID(accID.Uint64())
	if err != nil {
		return err
	}

	err = bz.authenticateTx(*tx, acc.PublicKey)
	if err != nil {
		return err
	}

	return nil
}

// ProcessTxs processes all trasnactions and returns the commitment list
<<<<<<< HEAD
func ProcessTxs(db DB, bz Bazooka, txs []Tx, isSyncing bool) (commitments []Commitment, err error) {
=======
func ProcessTxs(db DB, bz Bazooka, txs []Tx) (commitments []Commitment, err error) {
>>>>>>> 19b9212... revive sync
	if len(txs) == 0 {
		return commitments, ErrNoTxsFound
	}
	for i, tx := range txs {
		rootAcc, err := db.GetRoot()
		if err != nil {
			return commitments, err
		}
		currentRoot, err := HexToByteArray(rootAcc.Hash)
		if err != nil {
			return commitments, err
		}
		newRoot, err := tx.Validate(LoadedBazooka, currentRoot)
		if err != nil {
			return commitments, err
		}
		if i%COMMITMENT_SIZE == 0 {
			txInCommitment := txs[i : i+COMMITMENT_SIZE]
			aggregatedSig, err := aggregateSignatures(txInCommitment)
			if err != nil {
<<<<<<< HEAD
				if isSyncing && err == ErrSignatureNotPresent {
					continue
				} else {
					return commitments, err
				}
=======
				return commitments, err
>>>>>>> 19b9212... revive sync
			}
			commitment := Commitment{Txs: txInCommitment, UpdatedRoot: newRoot, BatchType: tx.Type, AggregatedSignature: aggregatedSig.ToBytes()}
			commitments = append(commitments, commitment)
		}
<<<<<<< HEAD
=======
		currentRoot = newRoot
>>>>>>> 19b9212... revive sync
	}

	return commitments, nil
}

// generates aggregated signature for commitment
func aggregateSignatures(txs []Tx) (aggregatedSig bls.Signature, err error) {
	var signatures []*bls.Signature
	for _, tx := range txs {
<<<<<<< HEAD
		if tx.Signature == nil {
			return aggregatedSig, ErrSignatureNotPresent
		}
=======
>>>>>>> 19b9212... revive sync
		sig, err := wallet.BytesToSignature(tx.Signature)
		if err != nil {
			return aggregatedSig, err
		}
		signatures = append(signatures, &sig)
	}
	return wallet.NewAggregateSignature(signatures)
}
