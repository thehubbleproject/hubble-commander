package core

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

type Commitments []Commitment

func NewCommitments(commitments []Commitment, batchID uint64) Commitments {
	for i, commitment := range commitments {
		commitment.Offset = uint64(i)
		commitment.BatchID = batchID
	}
	return commitments
}

type Commitment struct {
	CommitmentData

	// commitment meta information
	BatchID   uint64
	Offset    uint64
	BatchType uint64

	Txs                 []Tx   `gorm:"-"`
	AggregatedSignature []byte `gorm:"-"`
}

// CommitmentData is the crutial information per commitment that needs to be stored
type CommitmentData struct {
	StateRoot ByteArray
	BodyRoot  ByteArray
}

func (c CommitmentData) Leaf() (leaf ByteArray, err error) {
	data, err := c.Bytes()
	if err != nil {
		return
	}
	return ByteArray(Keccak256(data)), nil
}

func (c CommitmentData) Bytes() ([]byte, error) {
	bytes32Ty, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		return []byte(""), err
	}
	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: bytes32Ty,
		},
	}
	data, err := arguments.Pack(
		c.StateRoot,
		c.BodyRoot,
	)
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}

func NewCommitmentData(stateRoot, bodyRoot ByteArray) *CommitmentData {
	return &CommitmentData{StateRoot: stateRoot, BodyRoot: bodyRoot}
}

func NewCommitment(txs []Tx, batchType uint64, newRoot ByteArray, signature []byte) Commitment {
	return Commitment{
		BatchType: batchType,
	}
}

type TxCommitment interface {
	Hash() (ByteArray, error)
}

type TransferCommitment struct {
	StateRoot ByteArray
	TransferBody
}
type TransferBody struct {
	AccountRoot ByteArray
	Signature   [2]*big.Int
	FeeReceiver *big.Int
	Txs         []byte
}

func (t *TransferCommitment) Hash() (ByteArray, error) { return ByteArray(t.StateRoot), nil }

type Create2TransferCommitment struct {
	StateRoot ByteArray
	Create2TransferBody
}
type Create2TransferBody struct {
	AccountRoot ByteArray
	Signature   [2]*big.Int
	FeeReceiver *big.Int
	Txs         []byte
}

func (c *Create2TransferCommitment) Hash() (ByteArray, error) {
	return ByteArray(c.StateRoot), nil
}

type MassMigrationCommitment struct {
	StateRoot ByteArray
	MassMigrationBody
}

type MassMigrationBody struct {
	AccountRoot  ByteArray
	Signature    [2]*big.Int
	SpokeID      *big.Int
	WithdrawRoot ByteArray
	TokenID      *big.Int
	Amount       *big.Int
	FeeReceiver  *big.Int
	Txs          []byte
}

func (m *MassMigrationCommitment) Hash() (ByteArray, error) { return ByteArray(m.StateRoot), nil }
