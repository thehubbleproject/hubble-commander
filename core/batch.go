package core

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/math"
)

var (
	ErrBatchAlreadyCommitted = errors.New("Batch Already Committed")
)

// Batch is the batches that need to be submitted on-chain periodically
type Batch struct {
	BatchID        uint64 `json:"batch_id,omitempty"`
	SubmissionHash string `json:"submission_hash,omitempty"`
	BatchType      uint64 `json:"batch_type,omitempty"`
	Status         uint64 `json:"status,omitempty"`
}

func NewBatch(submissionHash string, batchType, status uint64) Batch {
	return Batch{SubmissionHash: submissionHash, BatchType: batchType, Status: status}
}

type Commitments []Commitment

type Commitment struct {
	CommitmentData

	// commitment meta information
	BatchID   uint64
	Offset    uint64
	BatchType uint64

	Txs                 []Tx   `gorm:"-"`
	AggregatedSignature []byte `gorm:"-"`
}

func NewCommitment(txs []Tx, batchType uint64, stateRoot ByteArray, signature []byte) Commitment {
	return Commitment{
		CommitmentData: CommitmentData{
			StateRoot: stateRoot[:],
		},
		BatchType:           batchType,
		Txs:                 txs,
		AggregatedSignature: signature,
	}
}

// Commit adds state root and body root to the commitment
func (c *Commitment) Commit(stateRoot, bodyRoot []byte) {
	c.StateRoot = stateRoot
	c.BodyRoot = bodyRoot
}

// AttachPostSubmissionInfo attaches batch ID and offset information to commitment
// before its pushed to DB
func (c *Commitment) AttachPostSubmissionInfo(batchID, offset uint64) {
	c.BatchID = batchID
	c.Offset = offset
}

// CommitmentData is the crutial information per commitment that needs to be stored
type CommitmentData struct {
	StateRoot []byte
	BodyRoot  []byte
}

func NewCommitmentData(stateRoot, bodyRoot ByteArray) *CommitmentData {
	return &CommitmentData{StateRoot: stateRoot[:], BodyRoot: bodyRoot[:]}
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
		BytesToByteArray(c.StateRoot),
		BytesToByteArray(c.BodyRoot),
	)
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}

type TxCommitment interface {
	Hash() (ByteArray, error)
	Bytes() ([]byte, error)
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

func (t *TransferCommitment) Hash() (ByteArray, error) {
	data, err := t.Bytes()
	if err != nil {
		return ByteArray{}, err
	}
	return ByteArray(Keccak256(data)), nil
}

func (t *TransferCommitment) Bytes() ([]byte, error) {
	accountRootBytes := t.AccountRoot[:]
	signatureBytes, err := decodeUint256Array(t.Signature)
	if err != nil {
		return nil, err
	}
	feeReceiverBytes := math.U256Bytes(t.FeeReceiver)
	return encodePacked(accountRootBytes, signatureBytes, feeReceiverBytes, t.Txs), nil
}

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
	data, err := c.Bytes()
	if err != nil {
		return ByteArray{}, err
	}
	return ByteArray(Keccak256(data)), nil
}

func (c *Create2TransferCommitment) Bytes() ([]byte, error) {
	accountRootBytes := c.AccountRoot[:]
	signatureBytes, err := decodeUint256Array(c.Signature)
	if err != nil {
		return nil, err
	}
	feeReceiverBytes := math.U256Bytes(c.FeeReceiver)
	return encodePacked(accountRootBytes, signatureBytes, feeReceiverBytes, c.Txs), nil
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

func (m *MassMigrationCommitment) Hash() (ByteArray, error) {
	data, err := m.Bytes()
	if err != nil {
		return ByteArray{}, err
	}
	return ByteArray(Keccak256(data)), nil
}

// Bytes converts mass migration to bytes
func (m *MassMigrationCommitment) Bytes() ([]byte, error) {
	accountRootBytes := m.AccountRoot[:]
	signatureBytes, err := decodeUint256Array(m.Signature)
	if err != nil {
		return nil, err
	}
	spokeIDBytes := math.U256Bytes(m.SpokeID)
	withdrawRootBytes := m.WithdrawRoot[:]
	tokenIDBytes := math.U256Bytes(m.TokenID)
	amountBytes := math.U256Bytes(m.Amount)
	feeReceiverBytes := math.U256Bytes(m.FeeReceiver)
	return encodePacked(accountRootBytes, signatureBytes, spokeIDBytes, withdrawRootBytes, tokenIDBytes, amountBytes, feeReceiverBytes, m.Txs), nil
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func decodeUint256Array(arr [2]*big.Int) ([]byte, error) {
	var output [][]byte
	for _, elem := range arr {
		output = append(output, math.U256Bytes(elem))
	}
	return bytes.Join(output, nil), nil
}
