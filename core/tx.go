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
	// row ID
	ID string `json:"-" gorm:"primary_key;size:100;default:'6ba7b810-9dad-11d1-80b4-00c04fd430c8'"`

	// encoded tx data
	Data []byte `json:"data"`

	// tranasction meta data
	Type uint64 `json:"type" gorm:"not null"`

	// derived fields essentail for faster querying
	From      uint64 `json:"from"`
	Nonce     uint64 `json:"nonce"`
	Fee       uint64 `json:"fee"`
	TokenType uint64 `json:"token"`

	// transaction signature
	Signature []byte `json:"sig" gorm:"null"`

	TxHash string `json:"hash" gorm:"not null"`
	Status uint64 `json:"status"`
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
func NewTx(data, sig []byte, from, nonce, fee, token, txType uint64) Tx {
	return Tx{
		Data:      data,
		From:      from,
		Nonce:     nonce,
		Fee:       fee,
		TokenType: token,
		Signature: sig,
		Type:      txType,
	}
}

// NewPendingTx creates a new pending transaction
func NewPendingTx(data, sig []byte, from, nonce, fee, token, txType uint64) (tx Tx, err error) {
	tx = Tx{
		Data:      data,
		Nonce:     nonce,
		From:      from,
		Fee:       fee,
		TokenType: token,
		Signature: sig,
		Type:      txType,
		Status:    TX_STATUS_PENDING,
	}

	if err = tx.AssignHash(); err != nil {
		return
	}
	return
}

// SignTx returns the transaction data that has to be signed
// it populates the signature field of the transaction or returns an error
func (tx *Tx) SignTx(wallet wallet.Wallet, txBytes []byte) (err error) {
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
	return fmt.Sprintf("Status:%v Hash: %v Data: %v", tx.Status, tx.TxHash, hex.EncodeToString(tx.Data))
}
