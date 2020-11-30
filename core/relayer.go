package core

import (
	"math/big"

	"github.com/jinzhu/gorm"
)

const (
	statusPackedReceived   = 1
	statusPackedProcessing = 2
	minPacketCount         = 1024
)

// RelayPacket is the relay packet for some specific actions
type RelayPacket struct {
	DBModel
	Data      []byte `json:"data"`
	Signature []byte `json:"sig" gorm:"null"`
	Pubkey    string `json:"pubkey" gorm:"null"`
	TxHash    string `json:"txHash"`
	Status    uint64 `json:"status"`
}

func (rp *RelayPacket) AfterCreate(tx *gorm.DB) (err error) {
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

	pubkeys, err := decodePackets(packets)
	if err != nil {
		return err
	}

	var pubkeyArr [minPacketCount][4]*big.Int
	copy(pubkeyArr[:], pubkeys)

	txHash, err := LoadedBazooka.RegisterPubkeys(pubkeyArr)
	if err != nil {
		return err
	}

	var ids []string
	for _, packet := range packets {
		ids = append(ids, packet.ID)
	}

	if err := tx.Model(rp).Where("id IN = ?", ids).Update(RelayPacket{Status: statusPackedProcessing, TxHash: txHash}).Error; err != nil {
		return err
	}

	return nil
}

// NewRelayPacket creates a new relay packet to be transmitted on-chain
func NewRelayPacket(data, signature []byte, pubkey string, status uint64) *RelayPacket {
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
	_, toPub, _, _, _, _, err := LoadedBazooka.DecodeCreate2TransferWithPub(data)
	if err != nil {
		return err
	}

	// TODO check if to account exists

	// create a new relay packet with status received
	pubkey, err := toPub.String()
	if err != nil {
		return err
	}
	rp := NewRelayPacket(data, sig, pubkey, statusPackedReceived)
	// check if the ToPubkey exists already, send error if true
	return db.Instance.Create(rp).Error
}

func (db *DB) GetPacketByPubkey(pubkey string) (rp RelayPacket, err error) {
	if err := db.Instance.Model(rp).Where("pubkey = ?", pubkey).Find(&rp).Error; err != nil {
		return rp, err
	}

	return rp, nil
}

func (db *DB) MarkPacketDone(pubkey string) error {
	rp, err := db.GetPacketByPubkey(pubkey)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
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

	fromIndex, _, nonce, txType, amount, fee, err := LoadedBazooka.DecodeCreate2TransferWithPub(rp.Data)
	if err != nil {
		return err
	}

	txData, err := LoadedBazooka.EncodeCreate2TransferTx(
		fromIndex.Int64(),
		int64(toStateID),
		int64(toAcc.ID),
		fee.Int64(),
		nonce.Int64(),
		amount.Int64(),
		txType.Int64(),
	)
	if err != nil {
		return err
	}

	tx, err := NewPendingTx(fromIndex.Uint64(), toStateID, txType.Uint64(), rp.Signature, txData)
	if err != nil {
		return err
	}

	if err := db.InsertTx(&tx); err != nil {
		return err
	}

	return nil
}

//
// utils
//

// decodes data in call packets and returns the pubkeys
func decodePackets(packets []RelayPacket) ([][4]*big.Int, error) {
	var tos [][4]*big.Int
	for _, packet := range packets {
		_, to, _, _, _, _, err := LoadedBazooka.DecodeCreate2TransferWithPub(packet.Data)
		if err != nil {
			return tos, err
		}
		tos = append(tos, to)
	}

	return tos, nil
}
