package core

import (
	"math/big"
)

// Batch is the batches that need to be submitted on-chain periodically
type Batch struct {
	BatchID              uint64
	StateRoot            string
	Committer            string
	StakeAmount          uint64
	FinalisesOn          big.Int
	SubmissionHash       string
	TransactionsIncluded []byte `gorm:"size:1000000"`
	BatchType            uint64
	Status               uint64
}

type Commitment struct {
	Txs                 []Tx
	UpdatedRoot         ByteArray
	BatchType           uint64
	AggregatedSignature []byte
}

func (db *DB) GetLatestBatch() (batch Batch, err error) {
	if err := db.Instance.Order("batch_id desc").First(&batch).Error; err != nil {
		return batch, err
	}
	return batch, nil
}

func (db *DB) GetBatchCount() (int, error) {
	var count int
	db.Instance.Table("batches").Count(&count)
	return count, nil
}

func (db *DB) AddNewBatch(batch Batch) error {
	return db.Instance.Create(batch).Error
}

func (db *DB) GetBatchByIndex(index uint64) (batch Batch, err error) {
	if err := db.Instance.Where("batch_id = ?", index).Find(&batch).Error; err != nil {
		return batch, err
	}
	return batch, nil
}

func (db *DB) CommitBatch(batch Batch) error {
	return db.Instance.Update(batch).Error
}
