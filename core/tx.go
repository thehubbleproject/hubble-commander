package core

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/BOPR/wallet"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrNoTxsFound          = errors.New("no tx found")
	ErrSignatureNotPresent = errors.New("signature not present")
)

// Tx represets the transaction on hubble
type Tx struct {
	ID        string `json:"-" gorm:"primary_key;size:100;default:'6ba7b810-9dad-11d1-80b4-00c04fd430c8'"`
	To        uint64 `json:"to"`
	From      uint64 `json:"from"`
	Data      []byte `json:"data"`
	Signature []byte `json:"sig" gorm:"null"`
	TxHash    string `json:"hash" gorm:"not null"`
	Status    uint64 `json:"status"`
	Type      uint64 `json:"type" gorm:"not null"`
}

// BeforeCreate sets id
func (tx *Tx) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("id", uuid.NewV4().String())
	if err != nil {
		return err
	}
	return nil
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
func NewPendingTx(from, to, txType uint64, sig, data []byte) (tx Tx, err error) {
	tx = Tx{
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

// SignTx returns the transaction data that has to be signed
func (tx *Tx) SignTx(secret, pubkey []byte, txBytes []byte) (err error) {
	wallet, err := wallet.SecretToWallet(secret, pubkey)
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
func (tx *Tx) AssignHash() (err error) {
	if tx.TxHash != "" {
		return nil
	}
	hash, err := RlpHash(tx)
	if err != nil {
		return
	}
	tx.TxHash = hash.String()
	return nil
}

func (tx *Tx) String() string {
	return fmt.Sprintf("To: %v From: %v Status:%v Hash: %v Data: %v", tx.To, tx.From, tx.Status, tx.TxHash, hex.EncodeToString(tx.Data))
}
