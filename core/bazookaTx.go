package core

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/BOPR/config"
	"github.com/BOPR/contracts/accountregistry"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SubmitTransferMethod        = "submitTransfer"
	SubmitCreate2TransferMethod = "submitCreate2Transfer"
	SubmitMassMigrationMethod   = "submitMassMigration"
)

type Calldata interface {
	Pack(b Bazooka, args interface{}) (data []byte, err error)
	Unpack(b Bazooka, data []byte) (err error)
	Method() string
}

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

// SubmitBatch submits the batch on chain with updated root and compressed transactions
func (b *Bazooka) SubmitBatch(commitments []Commitment) (txHash string, err error) {
	b.log.Info(
		"Attempting to submit a new batch",
		"NumOfCommitments",
		len(commitments),
	)

	if len(commitments) == 0 {
		b.log.Info("No transactions to submit, waiting....")
		return "", nil
	}

	switch txType := commitments[0].BatchType; txType {
	case TX_TRANSFER_TYPE:
		txHash, err := b.submitTransferBatch(commitments)
		if err != nil {
			return "", err
		}
		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", TX_TRANSFER_TYPE)
	case TX_CREATE_2_TRANSFER:
		txHash, err := b.submitCreate2TransferBatch(commitments)
		if err != nil {
			return "", err
		}
		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", TX_TRANSFER_TYPE)
	case TX_MASS_MIGRATIONS:
		txHash, err := b.submitMassMigrationBatch(commitments)
		if err != nil {
			return "", err
		}
		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", TX_TRANSFER_TYPE)
	default:
		b.log.Error("Tx not indentified", "txType", commitments[0].BatchType)
	}

	return txHash, nil
}

func (b *Bazooka) submitTransferBatch(commitments []Commitment) (string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int
	var feeReceivers []*big.Int

	dummyReceivers := big.NewInt(0)
	for i := 0; i <= len(commitments); i++ {
		feeReceivers = append(feeReceivers, dummyReceivers)
	}

	for _, commitment := range commitments {
		compressedTxs, err := b.CompressTxs(commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, commitment.UpdatedRoot)
		totalTxs += len(commitment.Txs)

		sig, err := BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return "", err
		}
		aggregatedSig = append(aggregatedSig, sig)
	}
	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)

	// TODO https://github.com/thehubbleproject/hubble-commander/issues/68
	stakeAmount := big.NewInt(1000000000000000000)

	data, err := b.RollupABI.Pack("submitTransfer", updatedRoots, aggregatedSig, feeReceivers, txs)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return "", err
	}

	auth, err := b.generateAuthObj(b.EthClient, rollupAddress, stakeAmount, data)
	if err != nil {
		b.log.Error("Estimate gas failed, tx reverting", "error", err)
		return "", err
	}

	tx, err := b.SC.RollupContract.SubmitTransfer(auth, updatedRoots, aggregatedSig, feeReceivers, txs)
	if err != nil {
		b.log.Error("Error submitting batch", "err", err)
		return "", err
	}
	receipt, err := bind.WaitMined(context.Background(), b.EthClient, tx)
	if err != nil {
		return "", err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return "", errors.New("Error tx reverted")
	}

	return tx.Hash().String(), nil
}

func (b *Bazooka) submitCreate2TransferBatch(commitments []Commitment) (string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int
	var feeReceivers []*big.Int

	dummyReceivers := big.NewInt(0)
	for i := 0; i <= len(commitments); i++ {
		feeReceivers = append(feeReceivers, dummyReceivers)
	}

	for _, commitment := range commitments {
		compressedTxs, err := b.CompressTxs(commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, commitment.UpdatedRoot)
		totalTxs += len(commitment.Txs)

		sig, err := BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return "", err
		}
		aggregatedSig = append(aggregatedSig, sig)
	}
	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)

	// TODO https://github.com/thehubbleproject/hubble-commander/issues/68
	stakeAmount := big.NewInt(1000000000000000000)

	data, err := b.RollupABI.Pack("submitCreate2Transfer", updatedRoots, aggregatedSig, feeReceivers, txs)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return "", nil
	}

	auth, err := b.generateAuthObj(b.EthClient, rollupAddress, stakeAmount, data)
	if err != nil {
		b.log.Error("Estimate gas failed, tx reverting", "error", err)
		return "", nil
	}

	tx, err := b.SC.RollupContract.SubmitCreate2Transfer(auth, updatedRoots, aggregatedSig, feeReceivers, txs)
	if err != nil {
		b.log.Error("Error submitting batch", "err", err)
		return "", nil
	}

	return tx.Hash().String(), nil
}

func (b *Bazooka) submitMassMigrationBatch(commitments []Commitment) (string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int

	var meta [][4]*big.Int
	var withdrawRoots [][32]byte

	dummyReceiver := big.NewInt(0)

	for _, commitment := range commitments {
		compressedTxs, err := b.CompressTxs(commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, commitment.UpdatedRoot)
		totalTxs += len(commitment.Txs)

		sig, err := BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return "", err
		}

		aggregatedSig = append(aggregatedSig, sig)

		var spokeID = big.NewInt(0)
		var tokenID = big.NewInt(0)
		var totalAmount = big.NewInt(0)

		for i, tx := range commitment.Txs {
			from, spoke, _, _, amount, _, err := b.DecodeMassMigrationTx(tx.Data)
			if err != nil {
				return "", err
			}

			if i == 0 {
				spokeID = spoke
				state, err := DBInstance.GetStateByIndex(from.Uint64())
				if err != nil {
					return "", err
				}
				_, _, _, token, err := b.DecodeState(state.Data)
				if err != nil {
					return "", err
				}
				tokenID = token
			}

			totalAmount.Add(amount, totalAmount)
		}

		var metaValues [4]*big.Int
		metaValues[0] = spokeID
		metaValues[1] = tokenID
		metaValues[2] = totalAmount
		metaValues[3] = dummyReceiver

		meta = append(meta, metaValues)
	}

	withdrawRoots = updatedRoots

	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)

	// TODO https://github.com/thehubbleproject/hubble-commander/issues/68
	stakeAmount := big.NewInt(100000000000000000)

	data, err := b.RollupABI.Pack("submitMassMigration", updatedRoots, aggregatedSig, meta, withdrawRoots, txs)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return "", nil
	}

	auth, err := b.generateAuthObj(b.EthClient, rollupAddress, stakeAmount, data)
	if err != nil {
		b.log.Error("Estimate gas failed, tx reverting", "error", err)
		return "", nil
	}

	fmt.Println("data here", hex.EncodeToString(data))

	tx, err := b.SC.RollupContract.SubmitMassMigration(auth, updatedRoots, aggregatedSig, meta, withdrawRoots, txs)
	if err != nil {
		b.log.Error("Error submitting batch", "err", err)
		return "", nil
	}

	return tx.Hash().String(), nil
}

func (b *Bazooka) FireDepositFinalisation(TBreplaced UserState, siblings []UserState, subTreeHeight uint64) (err error) {
	b.log.Info(
		"Attempting to finalise deposits",
		"NodeToBeReplaced",
		TBreplaced.String(),
		"NumberOfSiblings",
		len(siblings),
		"atDepth",
		subTreeHeight,
	)

	// depositSubTreeHeight := big.NewInt(0)
	// depositSubTreeHeight.SetUint64(subTreeHeight)
	// var siblingData [][32]byte
	// for _, sibling := range siblings {
	// 	data, err := HexToByteArray(sibling.Hash)
	// 	if err != nil {
	// 		b.log.Error("unable to convert HexToByteArray", err)
	// 		return err
	// 	}
	// 	siblingData = append(siblingData, data)
	// }

	// accountProof := rollup.TypesAccountMerkleProof{}
	// accountProof.AccountIP.PathToAccount = StringToBigInt(TBreplaced.Path)
	// userAccount, err := TBreplaced.ToABIAccount()
	// if err != nil {
	// 	b.log.Error("unable to convert", "error", err)
	// 	return
	// }
	// accountProof.AccountIP.Account = rollup.TypesUserState(userAccount)

	// accountProof.Siblings = siblingData
	// data, err := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Pack("finaliseDepositsAndSubmitBatch", depositSubTreeHeight, accountProof)
	// if err != nil {
	// 	fmt.Println("Unable to craete data", err)
	// 	return
	// }

	// rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	// stakeAmount := big.NewInt(0)
	// stakeAmount.SetString("32000000000000000000", 10)

	// // generate call msg
	// callMsg := ethereum.CallMsg{
	// 	To:    &rollupAddress,
	// 	Data:  data,
	// 	Value: stakeAmount,
	// }

	// auth, err := b.generateAuthObj(b.EthClient, callMsg)
	// if err != nil {
	// 	return err
	// }
	// b.log.Info("Broadcasting deposit finalisation transaction")

	// tx, err := b.RollupContract.FinaliseDepositsAndSubmitBatch(auth, depositSubTreeHeight, accountProof)
	// if err != nil {
	// 	return err
	// }
	// b.log.Info("Deposits successfully finalized!", "TxHash", tx.Hash())
	return nil
}

func (b *Bazooka) generateAuthObj(client *ethclient.Client, toAddr ethCmn.Address, value *big.Int, data []byte) (auth *bind.TransactOpts, err error) {
	// from address
	fromAddress := config.OperatorAddress

	// fetch gas price
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	// fetch nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}

	callMsg := ethereum.CallMsg{
		To:    &toAddr,
		Data:  data,
		Value: value,
	}

	// fetch gas limit
	callMsg.From = fromAddress
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return
	}
	// create auth
	auth = bind.NewKeyedTransactor(config.OperatorKey)
	auth.GasPrice = gasprice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(gasLimit)
	return
}

// RegisterPubkeys registers pubkeys in a batch
func (b *Bazooka) RegisterPubkeys(pubkeys [16][4]*big.Int) (txHash string, err error) {
	registryABI, err := abi.JSON(strings.NewReader(accountregistry.AccountregistryABI))
	if err != nil {
		return
	}
	data, err := registryABI.Pack("registerBatch", pubkeys)
	if err != nil {
		b.log.Error("Error packing data for register batch", "err", err)
		return
	}

	auth, err := b.generateAuthObj(b.EthClient, ethCmn.HexToAddress(config.GlobalCfg.AccountRegistry), big.NewInt(0), data)
	if err != nil {
		b.log.Error("Estimate gas failed, tx reverting", "error", err)
		return
	}

	tx, err := b.SC.AccountRegistry.RegisterBatch(auth, pubkeys)
	if err != nil {
		return
	}

	b.log.Info("Registered pubkeys", "count", len(pubkeys), "txHash", tx.Hash().String())

	return tx.Hash().String(), nil
}

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
	case TX_GENESIS:
		return calldata, ErrNoTxs
	case TX_DEPOSIT:
		return calldata, ErrNoTxs
	case TX_TRANSFER_TYPE:
		method = b.RollupABI.Methods[SubmitTransferMethod]
	case TX_CREATE_2_TRANSFER:
		method = b.RollupABI.Methods[SubmitCreate2TransferMethod]
	case TX_MASS_MIGRATIONS:
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
	case TX_GENESIS:
		return inputDataMap, ErrNoTxs
	case TX_DEPOSIT:
		return inputDataMap, ErrNoTxs
	case TX_TRANSFER_TYPE:
		method = b.RollupABI.Methods["submitTransfer"]
	case TX_CREATE_2_TRANSFER:
		method = b.RollupABI.Methods["submitCreate2Transfer"]
	case TX_MASS_MIGRATIONS:
		method = b.RollupABI.Methods["submitMassMigration"]
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
