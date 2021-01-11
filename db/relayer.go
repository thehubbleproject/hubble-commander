package db

import (
	"fmt"
	"math/big"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

const (
	statusPackedReceived   = 1
	statusPackedProcessing = 2
	minPacketCount         = 4
)

// RelayPacket is the relay packet for some specific actions
type RelayPacket struct {
	ID        string `json:"-" gorm:"primary_key;size:100;default:'6ba7b810-9dad-11d1-80b4-00c04fd430c8'"`
	Data      []byte `gorm:"type:varbinary(1000)" json:"data"`
	Signature []byte `json:"sig" gorm:"null"`
	Pubkey    []byte `json:"pubkey" gorm:"null"`
	TxHash    string `json:"txHash"`
	Status    uint64 `json:"status"`
}

// BeforeCreate sets id
func (rp *RelayPacket) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("id", uuid.NewV4().String())
	if err != nil {
		return err
	}
	return nil
}

func (rp *RelayPacket) AfterCreate(tx *gorm.DB, cfg config.Configuration) (err error) {
	bz, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		return err
	}
	query := tx.Model(&RelayPacket{}).Where("status = ?", statusPackedReceived)

	var count int
	if err := query.Count(&count).Error; err != nil {
		return err
	}
	if count < minPacketCount {
		return
	}

	var packets []RelayPacket
	if err := query.Find(&packets).Error; err != nil {
		return err
	}

	pubkeys, err := decodePackets(packets, bz)
	if err != nil {
		return err
	}

	var pubkeyArr [minPacketCount][4]*big.Int
	copy(pubkeyArr[:], pubkeys)

	txHash, err := bz.RegisterPubkeys(pubkeyArr)
	if err != nil {
		return err
	}

	for _, packet := range packets {
		if err := tx.Model(&packet).Where("id = ?", packet.ID).Updates(RelayPacket{Status: statusPackedProcessing, TxHash: txHash}).Error; err != nil {
			return err
		}
	}

	return nil
}

// NewRelayPacket creates a new relay packet to be transmitted on-chain
func NewRelayPacket(data, signature []byte, pubkey []byte, status uint64) *RelayPacket {
	return &RelayPacket{
		Data:      data,
		Pubkey:    pubkey,
		Signature: signature,
		Status:    status,
	}
}

// DB

// InsertRelayPacket inserts new relay packet
func (db *DB) InsertRelayPacket(data, sig []byte) error {
	// decode data to fetch pubkey
	_, toPub, _, _, _, _, err := db.Bazooka.DecodeCreate2TransferWithPub(data)
	if err != nil {
		return err
	}

	// TODO check if to account exists

	// create a new relay packet with status received
	pubkey := core.NewPubkey(toPub)
	rp := NewRelayPacket(data, sig, pubkey, statusPackedReceived)

	if err := db.Instance.Create(rp).Error; err != nil {
		return err
	}

	return nil
}

func (db *DB) GetPacketByPubkey(pubkey []byte) (rp RelayPacket, err error) {
	if err := db.Instance.Model(&rp).Where("pubkey = ?", pubkey).Find(&rp).Error; err != nil {
		return rp, err
	}

	return rp, nil
}

func (db *DB) MarkPacketDone(pubkey []byte) error {
	rp, err := db.GetPacketByPubkey(pubkey)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		if err == gorm.ErrRecordNotFound {
			return nil
		}
	}

	// update all records with pubkey as status done
	toAcc, err := db.GetAccountByPubkey(pubkey)
	if err != nil {
		return err
	}

	toStateID, err := db.ReserveEmptyLeaf()
	if err != nil {
		return err
	}

	fromIndex, _, nonce, txType, amount, fee, err := db.Bazooka.DecodeCreate2TransferWithPub(rp.Data)
	if err != nil {
		return err
	}

	txData, err := db.Bazooka.EncodeCreate2TransferTx(
		fromIndex.Int64(),
		int64(toStateID),
		int64(toAcc.ID),
		fee.Int64(),
		nonce.Int64(),
		amount.Int64(),
		txType.Int64(),
	)
	if err != nil {
		fmt.Println("error encoding", err)
		return err
	}

	tx, err := core.NewPendingTx(fromIndex.Uint64(), toStateID, txType.Uint64(), rp.Signature, txData)
	if err != nil {
		return err
	}

	return db.Instance.Create(&tx).Error
}

//
// utils
//

// decodes data in call packets and returns the pubkeys
func decodePackets(packets []RelayPacket, bz bazooka.Bazooka) ([][4]*big.Int, error) {
	var tos [][4]*big.Int
	for _, packet := range packets {
		_, to, _, _, _, _, err := bz.DecodeCreate2TransferWithPub(packet.Data)
		if err != nil {
			return tos, err
		}
		tos = append(tos, to)
	}

	return tos, nil
}
