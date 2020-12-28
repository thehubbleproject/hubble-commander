package bazooka

import (
	"errors"
	"math/big"

	"github.com/BOPR/core"
)

// Calldata interface defines different batches and how their call data is packed/unpacked
type Calldata interface {
	Pack(b Bazooka) (data []byte, err error)
	Commitments(accountRoot string) ([]core.CommitmentData, error)
	Method() string
}

// TransferCalldata implements Calldata and is used by TxType Transfer
type TransferCalldata struct {
	StateRoots   [][32]byte
	Signatures   [][2]*big.Int
	FeeReceivers []*big.Int
	Txss         [][]byte
}

func (c TransferCalldata) Pack(b Bazooka) (data []byte, err error) {
	data, err = b.RollupABI.Pack(SubmitTransferMethod, c.StateRoots, c.Signatures, c.FeeReceivers, c.Txss)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return data, err
	}

	return data, nil
}

func (c TransferCalldata) Commitments(accountRoot string) ([]core.CommitmentData, error) {
	// loop through all stateroots

	// create core.TranferCommitments

	// call .hash() and create body root

	// return core.CommitmentData

	return []core.CommitmentData{}, nil
}

func (c TransferCalldata) Method() string {
	return SubmitTransferMethod
}

type Create2TransferCalldata struct {
	Txss         [][]byte
	StateRoots   [][32]byte
	FeeReceivers []*big.Int
	Signatures   [][2]*big.Int
}

func (c Create2TransferCalldata) Pack(b Bazooka) (data []byte, err error) {
	data, err = b.RollupABI.Pack(SubmitTransferMethod, c.StateRoots, c.Signatures, c.FeeReceivers, c.Txss)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return data, err
	}
	return data, nil
}

func (c Create2TransferCalldata) Commitments(accountRoot string) ([]core.CommitmentData, error) {
	// loop through all stateroots

	// create core.TranferCommitments

	// call .hash() and create body root

	// return core.CommitmentData

	return []core.CommitmentData{}, nil
}

func (c Create2TransferCalldata) Method() string {
	return SubmitCreate2TransferMethod
}

type MassMigrationCalldata struct {
	Txss          [][]byte
	StateRoots    [][32]byte
	WithdrawRoots [][32]byte
	Meta          [][4]*big.Int
	Signatures    [][2]*big.Int
}

func (c MassMigrationCalldata) Pack(b Bazooka) (data []byte, err error) {
	data, err = b.RollupABI.Pack(SubmitMassMigrationMethod, c.StateRoots, c.Signatures, c.Meta, c.WithdrawRoots, c.Txss)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return data, err
	}
	return data, nil
}

func (c MassMigrationCalldata) Commitments(accountRoot string) ([]core.CommitmentData, error) {
	// loop through all stateroots
	for i := 0; i < len(c.StateRoots); i++ {

	}

	// create core.TranferCommitments

	// call .hash() and create body root

	// return core.CommitmentData

	return []core.CommitmentData{}, nil
}

func (c MassMigrationCalldata) Method() string {
	return SubmitMassMigrationMethod
}

func (b *Bazooka) UnpackBatchCalldata(method string, data []byte) (calldata Calldata, err error) {
	switch method {
	case SubmitTransferMethod:
		var transferCalldata TransferCalldata
		method := b.RollupABI.Methods[SubmitTransferMethod]
		err = method.Inputs.Unpack(&transferCalldata, data)
		if err != nil {
			return transferCalldata, err
		}
		return transferCalldata, nil
	case SubmitCreate2TransferMethod:
		var create2transfer Create2TransferCalldata
		method := b.RollupABI.Methods[SubmitCreate2TransferMethod]
		err = method.Inputs.Unpack(&create2transfer, data)
		if err != nil {
			return create2transfer, err
		}

		return create2transfer, nil
	case SubmitMassMigrationMethod:
		var massMigration MassMigrationCalldata
		method := b.RollupABI.Methods[SubmitMassMigrationMethod]
		err = method.Inputs.Unpack(&massMigration, data)
		if err != nil {
			return massMigration, err
		}

		return massMigration, nil
	}

	return nil, errors.New("Unable to match")
}
