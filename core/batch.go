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

func NewCommitment(batchID uint64, offset uint64, txs []Tx, batchType uint64, stateRoot, bodyRoot ByteArray, signature []byte) Commitment {
	return Commitment{
		CommitmentData: CommitmentData{
			StateRoot: stateRoot,
			BodyRoot:  bodyRoot,
		},

		BatchID:   batchID,
		Offset:    offset,
		BatchType: batchType,

		Txs:                 txs,
		AggregatedSignature: signature,
	}
}

// CommitmentData is the crutial information per commitment that needs to be stored
type CommitmentData struct {
	StateRoot ByteArray
	BodyRoot  ByteArray
}

func NewCommitmentData(stateRoot, bodyRoot ByteArray) *CommitmentData {
	return &CommitmentData{StateRoot: stateRoot, BodyRoot: bodyRoot}
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

func (t *TransferCommitment) Hash() (ByteArray, error) {
	uint2Ty, err := abi.NewType("uint256[2]", "uint256[2]", nil)
	if err != nil {
		return ByteArray{}, err
	}
	bytesTy, err := abi.NewType("bytes", "bytes", nil)
	if err != nil {
		return ByteArray{}, err
	}
	bytes32Ty, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		return ByteArray{}, err
	}
	uintTy, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return ByteArray{}, err
	}
	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: uint2Ty,
		},
		{
			Type: uintTy,
		},
		{
			Type: bytesTy,
		},
	}
	data, err := arguments.Pack(
		t.AccountRoot,
		t.Signature,
		t.FeeReceiver,
		t.Txs,
	)
	if err != nil {
		return ByteArray{}, err
	}

	return ByteArray(Keccak256(data)), nil
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
	uint2Ty, err := abi.NewType("uint256[2]", "uint256[2]", nil)
	if err != nil {
		return ByteArray{}, err
	}
	bytesTy, err := abi.NewType("bytes", "bytes", nil)
	if err != nil {
		return ByteArray{}, err
	}
	bytes32Ty, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		return ByteArray{}, err
	}
	uintTy, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return ByteArray{}, err
	}
	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: uint2Ty,
		},
		{
			Type: uintTy,
		},
		{
			Type: bytesTy,
		},
	}
	data, err := arguments.Pack(
		c.AccountRoot,
		c.Signature,
		c.FeeReceiver,
		c.Txs,
	)
	if err != nil {
		return ByteArray{}, err
	}

	return ByteArray(Keccak256(data)), nil
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
	uint2Ty, err := abi.NewType("uint256[2]", "uint256[2]", nil)
	if err != nil {
		return ByteArray{}, err
	}
	bytesTy, err := abi.NewType("bytes", "bytes", nil)
	if err != nil {
		return ByteArray{}, err
	}
	bytes32Ty, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		return ByteArray{}, err
	}
	uintTy, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return ByteArray{}, err
	}
	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: uint2Ty,
		},
		{
			Type: uintTy,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: uintTy,
		},
		{
			Type: uintTy,
		},
		{
			Type: uintTy,
		},
		{
			Type: bytesTy,
		},
	}
	data, err := arguments.Pack(
		m.AccountRoot,
		m.Signature,
		m.SpokeID,
		m.WithdrawRoot,
		m.TokenID,
		m.Amount,
		m.FeeReceiver,
		m.Txs,
	)
	if err != nil {
		return ByteArray{}, err
	}

	return ByteArray(Keccak256(data)), nil
}
