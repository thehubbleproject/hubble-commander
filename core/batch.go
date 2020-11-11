package core

import "errors"

var (
	ErrBatchAlreadyCommitted = errors.New("Batch Already Committed")
)

// Batch is the batches that need to be submitted on-chain periodically
type Batch struct {
	BatchID        uint64 `json:"batch_id,omitempty"`
	StateRoot      string `json:"state_root,omitempty"`
	Committer      string `json:"committer,omitempty"`
	SubmissionHash string `json:"submission_hash,omitempty"`
	BatchType      uint64 `json:"batch_type,omitempty"`
	Status         uint64 `json:"status,omitempty"`
}

func NewBatch(stateRoot, committer, submissionHash string, batchType, status uint64) Batch {
	return Batch{StateRoot: stateRoot, Committer: committer, SubmissionHash: submissionHash, BatchType: batchType, Status: status}
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
	batchCount, err := db.GetBatchCount()
	if err != nil {
		return err
	}

	// this is because batch ID starts from 0
	batch.BatchID = uint64(batchCount)

	return db.Instance.Create(batch).Error
}

func (db *DB) GetBatchByIndex(index uint64) (batch Batch, err error) {
	if err := db.Instance.Where("batch_id = ?", index).Find(&batch).Error; err != nil {
		return batch, err
	}
	return batch, nil
}

func (db *DB) CommitBatch(ID uint64) error {
	batch, err := db.GetBatchByIndex(ID)
	if err != nil {
		return err
	}

	if batch.Status == BATCH_COMMITTED {
		return ErrBatchAlreadyCommitted
	}

	batch.Status = BATCH_COMMITTED
	return db.Instance.Model(&batch).Where("batch_id = ?", batch.BatchID).Update(batch).Error
}

// IsCatchingUp returns true/false according to the sync status of the node
func IsCatchingUp() (bool, error) {
	totalBatches, err := LoadedBazooka.TotalBatches()
	if err != nil {
		return false, err
	}

	totalBatchedStored, err := DBInstance.GetBatchCount()
	if err != nil {
		return false, err
	}

	// if total batchse are greater than what we recorded we are still catching up
	if totalBatches > uint64(totalBatchedStored) {
		return true, err
	}

	return false, nil
}
