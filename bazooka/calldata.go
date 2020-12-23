package bazooka

import (
	"errors"
	"math/big"
)

// Calldata interface defines different batches and how their call data is packed/unpacked
type Calldata interface {
	Pack(b Bazooka, args interface{}) (data []byte, err error)
	Unpack(b Bazooka, data []byte) (err error)
	Method() string
}

// TransferCalldata implements Calldata and is used by TxType Transfer
type TransferCalldata struct {
	Txss         [][]byte
	StateRoots   [][32]byte
	feeReceivers []*big.Int
	Signatures   [][2]*big.Int
}

func (c *TransferCalldata) Pack(b Bazooka, input interface{}) (data []byte, err error) {
	body, ok := input.(TransferCalldata)
	if !ok {
		return data, errors.New("Input not of type transfer")
	}
	data, err = b.RollupABI.Pack(SubmitTransferMethod, body.StateRoots, body.Signatures, body.feeReceivers, body.Txss)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return data, err
	}
	return data, nil
}

func (c *TransferCalldata) Unpack(b Bazooka, data []byte) (err error) {
	err = b.RollupABI.Unpack(c, SubmitTransferMethod, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *TransferCalldata) Method() string {
	return SubmitTransferMethod
}

type Create2TransferCalldata struct {
	Txss         [][]byte
	StateRoots   [][32]byte
	feeReceivers []*big.Int
	Signatures   [][2]*big.Int
}

func (c *Create2TransferCalldata) Pack(b Bazooka, args ...interface{}) (data []byte, err error) {
	return nil, nil
}

func (c *Create2TransferCalldata) Unpack(b Bazooka, data []byte) (err error) {
	return nil
}
func (c *Create2TransferCalldata) Method() string {
	return SubmitCreate2TransferMethod
}

type MassMigrationCalldata struct {
	Txss          [][]byte
	StateRoots    [][32]byte
	WithdrawRoots [][32]byte
	Meta          [][4]*big.Int
	Signatures    [][2]*big.Int
}

func (c *MassMigrationCalldata) Pack(b Bazooka, args ...interface{}) (data []byte, err error) {
	return nil, nil
}

func (c *MassMigrationCalldata) Unpack(b Bazooka, data []byte) (err error) {
	return nil
}
func (c *MassMigrationCalldata) Method() string {
	return SubmitMassMigrationMethod
}
