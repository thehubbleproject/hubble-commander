package bazooka

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
)

func (b *Bazooka) getTxDataByHash(txHash ethCmn.Hash) (data []byte, err error) {
	tx, isPending, err := b.EthClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		b.log.Error("Cannot fetch transaction from hash", "Error", err)
		return
	}

	if isPending {
		b.log.Error("Transaction is still pending, cannot process", "Error", ErrTxPending)
		return data, ErrTxPending
	}

	payload := tx.Data()[4:]
	return payload, nil
}

func (b *Bazooka) ParseCalldata(txHash ethCmn.Hash, batchType uint8) (calldata Calldata, err error) {
	// inputDataMap := make(map[string]interface{})
	var method abi.Method
	var data []byte

	switch batchType {
	case core.TX_GENESIS:
		return calldata, ErrNoTxs
	case core.TX_DEPOSIT:
		return calldata, ErrNoTxs
	case core.TX_TRANSFER_TYPE:
		method = b.RollupABI.Methods[SubmitTransferMethod]
	case core.TX_CREATE_2_TRANSFER:
		method = b.RollupABI.Methods[SubmitCreate2TransferMethod]
	case core.TX_MASS_MIGRATIONS:
		method = b.RollupABI.Methods[SubmitMassMigrationMethod]
	}
	fmt.Println("data and method", data, method)

	data, err = b.getTxDataByHash(txHash)
	if err != nil {
		return nil, err
	}

	fmt.Println("data here", hex.EncodeToString(data))

	// calldata.Unpack()
	return calldata, nil
}

func (b *Bazooka) getInputData(txHash ethCmn.Hash, batchType uint8) (map[string]interface{}, error) {
	inputDataMap := make(map[string]interface{})
	var method abi.Method
	var data []byte

	switch batchType {
	case core.TX_GENESIS:
		return inputDataMap, ErrNoTxs
	case core.TX_DEPOSIT:
		return inputDataMap, ErrNoTxs
	case core.TX_TRANSFER_TYPE:
		method = b.RollupABI.Methods[SubmitTransferMethod]
	case core.TX_CREATE_2_TRANSFER:
		method = b.RollupABI.Methods[SubmitCreate2TransferMethod]
	case core.TX_MASS_MIGRATIONS:
		method = b.RollupABI.Methods[SubmitMassMigrationMethod]
	}

	data, err := b.getTxDataByHash(txHash)
	if err != nil {
		return nil, err
	}

	err = method.Inputs.UnpackIntoMap(inputDataMap, data)
	if err != nil {
		b.log.Error("Error unpacking payload", "Error", err, "data", data)
		return inputDataMap, err
	}

	return inputDataMap, nil
}

// FetchTxsFromBatch parses the calldata for transactions
func (b *Bazooka) FetchTxsFromBatch(txHash ethCmn.Hash, batchType uint8) (txs []byte, err error) {
	inputs, err := b.getInputData(txHash, batchType)
	if err != nil {
		if err == ErrNoTxs {
			return txs, nil
		}
		return txs, err
	}

	return getTxsFromInput(inputs)
}

// FetchMetaInfoFromBatch parses the calldata for transactions
func (b *Bazooka) FetchMetaInfoFromBatch(txHash ethCmn.Hash, batchType uint8) (withdrawRoots [][32]byte, toSpokeIDs, tokenIDs, amounts []*big.Int, err error) {
	inputs, err := b.getInputData(txHash, batchType)
	if err != nil {
		return
	}

	return getMassMigrationInfo(inputs)
}

func getTxsFromInput(input map[string]interface{}) (txs []byte, err error) {
	if txPayload, ok := input[TXS_PARAM]; ok {
		txList, ok := txPayload.([][]byte)
		if !ok {
			return nil, ErrConvertingTxPayload
		}

		txs = txList[0]
	} else {
		return nil, ErrTxParamDoesntExist
	}
	return txs, nil
}

func getMassMigrationInfo(input map[string]interface{}) (withdrawRoots [][32]byte, toSpokeIDs, tokenIDs, amounts []*big.Int, err error) {
	fmt.Println("print input data", input)
	if txPayload, ok := input[META_PARAM]; ok {
		metaList, ok := txPayload.([][4]*big.Int)
		if !ok {
			return withdrawRoots, toSpokeIDs, tokenIDs, amounts, ErrConvertingTxPayload
		}
		fmt.Println("meta list", metaList)

		// txs = metaList[0]
	} else {
		return withdrawRoots, toSpokeIDs, tokenIDs, amounts, ErrTxParamDoesntExist
	}
	return
}
