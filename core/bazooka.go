package core

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/BOPR/contracts/logger"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/contracts/rollupclient"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// IContractCaller is the common interface using which we will interact with the contracts
// and the ethereum chain
type IBazooka interface {
	FetchBatchInputData(txHash ethCmn.Hash) (txs [][]byte, err error)
}

var (
	ErrTxPending           = errors.New("Tx Pending, cannot read calldata")
	ErrConvertingTxPayload = errors.New("Error converting tx payload")
	ErrTxParamDoesntExist  = errors.New("Tx param does not exist")
)

const (
	TXS_PARAM = "txss"
)

// Global Contract Caller Object
var LoadedBazooka Bazooka

// ContractCaller satisfies the IContractCaller interface and contains all the variables required to interact
// With the ethereum chain along with contract addresses and ABI's
type Bazooka struct {
	log       log.Logger
	EthClient *ethclient.Client

	ContractABI map[string]abi.ABI

	RollupContract *rollup.Rollup
	EventLogger    *logger.Logger
	Frontend       *rollupclient.Rollupclient
}

// NewContractCaller contract caller
// NOTE: Reads configration from the config.toml file
func NewPreLoadedBazooka() (bazooka Bazooka, err error) {
	err = config.SetOperatorKeys(config.GlobalCfg.OperatorKey)
	if err != nil {
		return
	}

	if RPCClient, err := rpc.Dial(config.GlobalCfg.EthRPC); err != nil {
		return bazooka, err
	} else {
		bazooka.EthClient = ethclient.NewClient(RPCClient)
	}

	bazooka.ContractABI = make(map[string]abi.ABI)

	// initialise all variables for rollup contract
	rollupContractAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	if bazooka.RollupContract, err = rollup.NewRollup(rollupContractAddress, bazooka.EthClient); err != nil {
		return bazooka, err
	}
	if bazooka.ContractABI[common.ROLLUP_CONTRACT_KEY], err = abi.JSON(strings.NewReader(rollup.RollupABI)); err != nil {
		return bazooka, err
	}

	// initialise all variables for event logger contract
	loggerAddress := ethCmn.HexToAddress(config.GlobalCfg.LoggerAddress)
	if bazooka.EventLogger, err = logger.NewLogger(loggerAddress, bazooka.EthClient); err != nil {
		return bazooka, err
	}
	if bazooka.ContractABI[common.LOGGER_KEY], err = abi.JSON(strings.NewReader(logger.LoggerABI)); err != nil {
		return bazooka, err
	}

	clientAddr := ethCmn.HexToAddress(config.GlobalCfg.FrontendAddress)
	if bazooka.Frontend, err = rollupclient.NewRollupclient(clientAddr, bazooka.EthClient); err != nil {
		return bazooka, err
	}
	if bazooka.ContractABI[common.ROLLUP_CLIENT], err = abi.JSON(strings.NewReader(rollupclient.RollupclientABI)); err != nil {
		return bazooka, err
	}

	bazooka.log = common.Logger.With("module", "bazooka")
	return bazooka, nil
}

// GetMainChainBlock fetches the eth chain block for a block num
func (b *Bazooka) GetMainChainBlock(blockNum *big.Int) (header *ethTypes.Header, err error) {
	latestBlock, err := b.EthClient.HeaderByNumber(context.Background(), blockNum)
	if err != nil {
		return
	}
	return latestBlock, nil
}

// TotalBatches returns the total number of batches that have been submitted on chain
func (b *Bazooka) TotalBatches() (uint64, error) {
	totalBatches, err := b.RollupContract.NumOfBatchesSubmitted(nil)
	if err != nil {
		return 0, err
	}
	return totalBatches.Uint64(), nil
}

<<<<<<< HEAD
func (b *Bazooka) GetTxDataByHash(txHash ethCmn.Hash) (data []byte, err error) {
=======
// FetchBatchInputData parses the calldata for transactions
func (b *Bazooka) FetchBatchInputData(txHash ethCmn.Hash) (txs []byte, err error) {
>>>>>>> 19b9212... revive sync
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

// FetchBatchInputData parses the calldata for transactions
func (b *Bazooka) FetchBatchInputData(txHash ethCmn.Hash, batchType uint8) (txs []byte, err error) {
	inputDataMap := make(map[string]interface{})
	var method abi.Method
	var data []byte

	switch batchType {
	case TX_GENESIS:
		return []byte{}, nil
	case TX_DEPOSIT:
		return []byte{}, nil
	case TX_TRANSFER_TYPE:
		method = b.ContractABI[common.ROLLUP_CONTRACT_KEY].Methods["submitTransfer"]
	case TX_CREATE_2_TRANSFER:
		return []byte{}, nil
	}

	data, err = b.GetTxDataByHash(txHash)
	if err != nil {
		return nil, err
	}
	err = method.Inputs.UnpackIntoMap(inputDataMap, data)
	if err != nil {
		b.log.Error("Error unpacking payload", "Error", err)
		return
	}

<<<<<<< HEAD
	return getTxsFromInput(inputDataMap)
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
=======
	return getTxsFromInput(inputDataMap), nil
}

func getTxsFromInput(input map[string]interface{}) (txs []byte) {
	data := input["_txs"].([]byte)
	return data
>>>>>>> 19b9212... revive sync
}

// ProcessTx calls the ProcessTx function on the contract to verify the tx
// returns the updated accounts and the new balance root
func (b *Bazooka) ProcessTx(balanceTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot ByteArray, err error) {
	b.log.Info("Processing new tx", "type", tx.Type)
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return b.processTransferTx(balanceTreeRoot, tx, fromMerkleProof, toMerkleProof)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return newBalanceRoot, errors.New("Did not match any options")
	}
}

func (b *Bazooka) ApplyTx(sender, receiver []byte, tx Tx) (updatedSender, updatedReceiver []byte, err error) {
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return b.applyTransferTx(sender, receiver, tx)
	// case TX_CREATE_2_TRANSFER:
	// 	return b.applyCreate2TransferTx(sender, receiver, tx)
	// case TX_MASS_MIGRATIONS:
	// 	return b.applyMassMigrationTx(sender, receiver, tx)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return updatedSender, updatedReceiver, errors.New("Didn't match any options")
	}
}

func (b *Bazooka) CompressTxs(txs []Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	var data [][]byte
	for _, tx := range txs {
		data = append(data, tx.Data)
	}
	switch txType := txs[0].Type; txType {
	case TX_TRANSFER_TYPE:
		return LoadedBazooka.compressTransferTxs(opts, data)
	// case TX_CREATE_2_TRANSFER:
	// 	return LoadedBazooka.compressCreate2TransferTxs(opts, data)
	// case TX_MASS_MIGRATIONS:
	// 	return LoadedBazooka.compressMassMigrationTxs(opts, data)
	default:
		fmt.Println("TxType didnt match any options", txs[0].Type)
		return []byte(""), errors.New("Did not match any options")
	}
}

func (b *Bazooka) processTransferTx(balanceTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot ByteArray, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion()
	if err != nil {
		return
	}
	toMP, err := toMerkleProof.ToABIVersion()
	if err != nil {
		return
	}

	result, err := b.Frontend.ProcessTransfer(
		&opts,
		balanceTreeRoot,
		tx.Data,
		fromMP.State.TokenType,
		fromMP,
		toMP,
	)
	if err != nil {
		return
	}
	if err = ParseResult(result.Result); err != nil {
		return
	}

	b.log.Info("Processed transaction", "postTxRoot", ByteArray(result.NewRoot).String(), "resultCode", result.Result)

	// TOOD read result code and bubble up error messages
	return result.NewRoot, nil
}

// TOOD add processCreate2TransferTx
// TOOD add processMassMigrationTx

// func (b *Bazooka) applyCreate2TransferTx(sender, receiver []byte, tx Tx) (updatedSender, updatedReceiver []byte, err error) {
// 	opts := bind.CallOpts{From: config.OperatorAddress}
// 	updates, err := b.Frontend.ValidateAndApplyCreate2Transfer(&opts, sender, tx.Data)
// 	if err != nil {
// 		return
// 	}

// 	if err = ParseResult(updates.Result); err != nil {
// 		return
// 	}

// 	return updates.NewSender, updates.NewReceiver, nil
// }

func (b *Bazooka) applyTransferTx(sender, receiver []byte, tx Tx) (updatedSender, updatedReceiver []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	updates, err := b.Frontend.ValidateAndApplyTransfer(&opts, sender, receiver, tx.Data)
	if err != nil {
		return
	}

	if err = ParseResult(updates.Result); err != nil {
		return
	}

	return updates.NewSender, updates.NewReceiver, nil
}

// func (b *Bazooka) applyMassMigrationTx(sender, receiver []byte, tx Tx) (updatedSender, updatedReceiver []byte, err error) {
// 	opts := bind.CallOpts{From: config.OperatorAddress}
// 	updates, err := b.Frontend.ValidateAndApplyMassMigration(&opts, sender, tx.Data)
// 	if err != nil {
// 		return
// 	}

// 	if err = ParseResult(updates.Result); err != nil {
// 		return
// 	}

// 	return updates.NewSender, updatedReceiver, nil
// }

func (b *Bazooka) compressTransferTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
	return b.Frontend.CompressTransfer(&opts, data)
}

// func (b *Bazooka) compressCreate2TransferTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
// 	return b.Frontend.CompressCreate2Transfer(&opts, data)
// }

// func (b *Bazooka) compressMassMigrationTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
// 	return b.Frontend.CompressMassMigration(&opts, data)
// }

func (b *Bazooka) TransferSignBytes() {
	// TODO hook into tx.GetSignData
}

func (b *Bazooka) DecompressTransferTxs(txs []byte) (froms, tos, amounts, fees []big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTxs, err := b.Frontend.DecompressTransfer(&opts, txs)
	if err != nil {
		return
	}

	for _, decompressedTx := range decompressedTxs {
		froms = append(froms, *decompressedTx.FromIndex)
		tos = append(tos, *decompressedTx.ToIndex)
		amounts = append(amounts, *decompressedTx.Amount)
		fees = append(fees, *decompressedTx.Fee)
	}

	return
}

//
// Encoders and Decoders for transactions
//

func (b *Bazooka) EncodeTransferTx(from, to, fee, nonce, amount, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx := struct {
		TxType    *big.Int
		FromIndex *big.Int
		ToIndex   *big.Int
		Amount    *big.Int
		Fee       *big.Int
		Nonce     *big.Int
	}{big.NewInt(txType), big.NewInt(from), big.NewInt(to), big.NewInt(amount), big.NewInt(fee), big.NewInt(nonce)}
	return b.Frontend.EncodeTransfer(&opts, tx)
}

func (b *Bazooka) DecodeTransferTx(txBytes []byte) (from, to, nonce, txType, amount, fee *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.Frontend.DecodeTransfer(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.FromIndex, tx.ToIndex, tx.Nonce, tx.TxType, tx.Amount, tx.Fee, nil
}

// TODO add encoders decoders to create2transfer
// TODO add encoders decoders for mass migrations txs

//
// Encoders and Decoders for state
//

func (b *Bazooka) EncodeState(id, balance, nonce, token uint64) (accountBytes []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	accountBytes, err = b.Frontend.Encode(&opts, rollupclient.TypesUserState{
		PubkeyIndex: big.NewInt(int64(id)),
		TokenType:   big.NewInt(int64(token)),
		Balance:     big.NewInt(int64(balance)),
		Nonce:       big.NewInt(int64(nonce)),
	})
	if err != nil {
		return
	}
	return accountBytes, nil
}

func (b *Bazooka) DecodeState(stateBytes []byte) (ID, balance, nonce, token *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}

	state, err := b.Frontend.DecodeState(&opts, stateBytes)
	if err != nil {
		return
	}

	b.log.Debug("Decoded state", "ID", state.PubkeyIndex, "balance", state.Balance, "token", state.TokenType, "nonce", state.Nonce)
	return state.PubkeyIndex, state.Balance, state.Nonce, state.TokenType, nil
}

// ----------------------------------------------------------------

//
// Transactions
//

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

	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int

	for _, commitment := range commitments {
		compressedTxs, err := b.CompressTxs(commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, commitment.UpdatedRoot)
		totalTxs += len(commitment.Txs)

		// TODO cleanup
		sig1 := commitment.AggregatedSignature[0:32]
		sig2 := commitment.AggregatedSignature[32:64]
		sig1bigInt := big.NewInt(0)
		sig1bigInt.SetBytes(sig1)
		sig2bigInt := big.NewInt(0)
		sig2bigInt.SetBytes(sig2)
		aggregatedSigBigInt := [2]*big.Int{sig1bigInt, sig2bigInt}
		aggregatedSig = append(aggregatedSig, aggregatedSigBigInt)
	}

	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	// TODO https://github.com/thehubbleproject/hubble-commander/issues/68
	stakeAmount := big.NewInt(100000000000000000)

	// TODO fix
	var feeReceivers []*big.Int
	dummyReceivers := big.NewInt(0)
	feeReceivers = append(feeReceivers, dummyReceivers)

	switch txType := commitments[0].BatchType; txType {
	case TX_TRANSFER_TYPE:
		data, err := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Pack("submitTransfer", updatedRoots, aggregatedSig, feeReceivers, txs)
		if err != nil {
			b.log.Error("Error packing data for submitBatch", "err", err)
			return "", nil
		}

		// generate call msg
		callMsg := ethereum.CallMsg{
			To:    &rollupAddress,
			Data:  data,
			Value: stakeAmount,
		}

		// generate auth
		auth, err := b.generateAuthObj(b.EthClient, callMsg)
		if err != nil {
			b.log.Error("Estimate gas failed, tx reverting", "error", err)
			return "", nil
		}

		tx, err := b.RollupContract.SubmitTransfer(auth, updatedRoots, aggregatedSig, feeReceivers, txs)
		if err != nil {
			b.log.Error("Error submitting batch", "err", err)
			return "", nil
		}

		txHash = tx.Hash().String()
		b.log.Info("Sent a new batch!", "TxHash", txHash)
	default:
		b.log.Error("Tx not indentified", "txType", commitments[0].BatchType)
	}

	return txHash, nil
}

func (b *Bazooka) FireDepositFinalisation(TBreplaced UserState, siblings []UserState, subTreeHeight uint64) (err error) {
	// b.log.Info(
	// 	"Attempting to finalise deposits",
	// 	"NodeToBeReplaced",
	// 	TBreplaced.String(),
	// 	"NumberOfSiblings",
	// 	len(siblings),
	// 	"atDepth",
	// 	subTreeHeight,
	// )

	// // TODO check latest batch on-chain nd if we need to push new batch
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

	// auth, err := b.GenerateAuthObj(b.EthClient, callMsg)
	// if err != nil {
	// 	return err
	// }
	// lastTxBroadcasted, err := DBInstance.GetLastTransaction()
	// if err != nil {
	// 	return err
	// }
	// if lastTxBroadcasted.Nonce+1 != auth.Nonce.Uint64() {
	// 	b.log.Info("Replacing nonce", "nonceEstimated", auth.Nonce.String(), "replacedBy", lastTxBroadcasted.Nonce+1)
	// 	auth.Nonce = big.NewInt(int64(lastTxBroadcasted.Nonce + 1))
	// }
	// b.log.Info("Broadcasting deposit finalisation transaction")

	// tx, err := b.RollupContract.FinaliseDepositsAndSubmitBatch(auth, depositSubTreeHeight, accountProof)
	// if err != nil {
	// 	return err
	// }
	// // TODO change this to deposit type
	// err = DBInstance.LogBatch(auth.Nonce.Uint64(), 100, "", []byte(""))
	// if err != nil {
	// 	return err
	// }
	// b.log.Info("Deposits successfully finalized!", "TxHash", tx.Hash())
	return nil
}

func (b *Bazooka) generateAuthObj(client *ethclient.Client, callMsg ethereum.CallMsg) (auth *bind.TransactOpts, err error) {
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
