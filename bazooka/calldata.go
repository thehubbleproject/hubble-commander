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

func (c TransferCalldata) Commitments(accountRoot string) (commitmentDatas []core.CommitmentData, err error) {
	for i := range c.StateRoots {
		var transferCommitment core.TransferCommitment
		transferCommitment.AccountRoot, err = core.HexToByteArray(accountRoot)
		if err != nil {
			return
		}

		transferCommitment.StateRoot = c.StateRoots[i]
		transferCommitment.Signature = c.Signatures[i]
		transferCommitment.FeeReceiver = c.FeeReceivers[i]
		transferCommitment.Txs = c.Txss[i]

		bodyRoot, inErr := transferCommitment.Hash()
		if inErr != nil {
			return
		}

		commitmentDatas = append(commitmentDatas, *core.NewCommitmentData(c.StateRoots[i], bodyRoot))
	}
	return commitmentDatas, nil
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

func (c Create2TransferCalldata) Commitments(accountRoot string) (commitmentDatas []core.CommitmentData, err error) {
	for i := range c.StateRoots {
		var c2tCommitment core.Create2TransferCommitment
		c2tCommitment.AccountRoot, err = core.HexToByteArray(accountRoot)
		if err != nil {
			return
		}

		c2tCommitment.StateRoot = c.StateRoots[i]
		c2tCommitment.Signature = c.Signatures[i]
		c2tCommitment.FeeReceiver = c.FeeReceivers[i]
		c2tCommitment.Txs = c.Txss[i]

		bodyRoot, inErr := c2tCommitment.Hash()
		if inErr != nil {
			return
		}

		commitmentDatas = append(commitmentDatas, *core.NewCommitmentData(c.StateRoots[i], bodyRoot))
	}
	return commitmentDatas, nil
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

func (c MassMigrationCalldata) Commitments(accountRoot string) (commitmentDatas []core.CommitmentData, err error) {
	for i := range c.StateRoots {
		var mmCommitment core.MassMigrationCommitment
		mmCommitment.AccountRoot, err = core.HexToByteArray(accountRoot)
		if err != nil {
			return
		}

		mmCommitment.StateRoot = c.StateRoots[i]
		mmCommitment.Signature = c.Signatures[i]
		mmCommitment.WithdrawRoot = c.WithdrawRoots[i]
		mmCommitment.Txs = c.Txss[i]

		mmCommitment.FeeReceiver = c.Meta[i][3]
		mmCommitment.Amount = c.Meta[i][2]
		mmCommitment.SpokeID = c.Meta[i][0]
		mmCommitment.TokenID = c.Meta[i][1]

		bodyRoot, inErr := mmCommitment.Hash()
		if inErr != nil {
			return
		}

		commitmentDatas = append(commitmentDatas, *core.NewCommitmentData(c.StateRoots[i], bodyRoot))
	}
	return commitmentDatas, nil
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
