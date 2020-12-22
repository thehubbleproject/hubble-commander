package core

import (
	"errors"
)

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
