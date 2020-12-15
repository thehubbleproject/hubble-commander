package core

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/wallet"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/BOPR/contracts/accountregistry"
	"github.com/BOPR/contracts/create2transfer"
	"github.com/BOPR/contracts/massmigration"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/contracts/state"
	"github.com/BOPR/contracts/transfer"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// IBazooka is the common interface using which we will interact with the contracts
// and the ethereum chain
type IBazooka interface {
	FetchBatchInputData(txHash ethCmn.Hash) (txs [][]byte, err error)
}

var (
	ErrTxPending           = errors.New("Tx Pending, cannot read calldata")
	ErrConvertingTxPayload = errors.New("Error converting tx payload")
	ErrTxParamDoesntExist  = errors.New("Tx param does not exist")
	ErrNoTxs               = errors.New("Error no transactions")
)

const (
	TXS_PARAM      = "txss"
	META_PARAM     = "meta"
	WITHDRAW_ROOTS = "withdrawRoots"
)

var (
	DOMAIN = [32]byte{}
)

// Global bazooka Object
var LoadedBazooka Bazooka

// Bazooka contains everything needed to interact with smart contracts
type Bazooka struct {
	log       log.Logger
	EthClient *ethclient.Client

	RollupABI abi.ABI
	SC        Contracts
}

type Contracts struct {
	RollupContract  *rollup.Rollup
	State           *state.State
	Transfer        *transfer.Transfer
	Create2Transfer *create2transfer.Create2transfer
	MassMigration   *massmigration.Massmigration
	AccountRegistry *accountregistry.Accountregistry
}

// NewPreLoadedBazooka creates
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

	bazooka.RollupABI, err = abi.JSON(strings.NewReader(rollup.RollupABI))
	if err != nil {
		return
	}

	bazooka.SC, err = getContractInstances(bazooka.EthClient)
	if err != nil {
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
	totalBatches, err := b.SC.RollupContract.NextBatchID(nil)
	if err != nil {
		return 0, err
	}
	return totalBatches.Uint64(), nil
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

// ProcessTx calls the ProcessTx function on the contract to verify the tx
// returns the updated accounts and the new balance root
func (b *Bazooka) ProcessTx(balanceTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot ByteArray, err error) {
	b.log.Info("Processing new tx", "type", tx.Type)
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return b.processTransferTx(balanceTreeRoot, tx, fromMerkleProof, toMerkleProof)
	case TX_CREATE_2_TRANSFER:
		return b.processCreate2TransferTx(balanceTreeRoot, tx, fromMerkleProof, toMerkleProof)
	case TX_MASS_MIGRATIONS:
		return b.processMassMigrationTx(balanceTreeRoot, tx, fromMerkleProof, toMerkleProof)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return newBalanceRoot, errors.New("Did not match any options")
	}
}

// ApplyTx applies the transaction and returns the udpates
func (b *Bazooka) ApplyTx(sender, receiver []byte, tx Tx) (updatedSender, updatedReceiver []byte, err error) {
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return b.applyTransferTx(sender, receiver, tx)
	case TX_CREATE_2_TRANSFER:
		return b.applyCreate2TransferTx(sender, receiver, tx)
	case TX_MASS_MIGRATIONS:
		return b.applyMassMigrationTx(sender, receiver, tx)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return updatedSender, updatedReceiver, errors.New("Didn't match any options")
	}
}

// CompressTxs compresses all transactions
func (b *Bazooka) CompressTxs(txs []Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	var data [][]byte
	for _, tx := range txs {
		data = append(data, tx.Data)
	}
	switch txType := txs[0].Type; txType {
	case TX_TRANSFER_TYPE:
		return b.compressTransferTxs(opts, data)
	case TX_CREATE_2_TRANSFER:
		return b.compressCreate2TransferTxs(opts, data)
	case TX_MASS_MIGRATIONS:
		return b.compressMassMigrationTxs(opts, data)
	default:
		fmt.Println("TxType didnt match any options", txs[0].Type)
		return []byte(""), errors.New("Did not match any options")
	}
}

func (b *Bazooka) authenticateTx(db DB, tx Tx, pubkeySender []byte) error {
	opts := bind.CallOpts{From: config.OperatorAddress}
	solPubkeySender, err := Pubkey(pubkeySender).ToSol()
	if err != nil {
		return err
	}
	signature, err := BytesToSolSignature(tx.Signature)
	if err != nil {
		return err
	}

	switch tx.Type {
	case TX_TRANSFER_TYPE:
		err = b.SC.Transfer.Validate(&opts, tx.Data, signature, solPubkeySender, wallet.DefaultDomain)
		if err != nil {
			return err
		}
	case TX_CREATE_2_TRANSFER:
		_, _, toAccID, _, _, _, _, err := b.DecodeCreate2Transfer(tx.Data)
		if err != nil {
			return err
		}
		acc, err := db.GetAccountLeafByID(toAccID.Uint64())
		if err != nil {
			return err
		}
		solPubkeyReceiver, err := Pubkey(acc.PublicKey).ToSol()
		if err != nil {
			return err
		}
		err = b.SC.Create2Transfer.Validate(&opts, tx.Data, signature, solPubkeySender, solPubkeyReceiver, wallet.DefaultDomain)
		if err != nil {
			return err
		}
	case TX_MASS_MIGRATIONS:
		err = b.SC.MassMigration.Validate(&opts, tx.Data, signature, solPubkeySender, wallet.DefaultDomain)
		if err != nil {
			return err
		}
	}

	return nil
}
func (b *Bazooka) processCreate2TransferTx(balanceTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot ByteArray, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion()
	if err != nil {
		return
	}
	toMP, err := toMerkleProof.ToABIVersion()
	if err != nil {
		return
	}

	result, err := b.SC.Create2Transfer.Process(
		&opts,
		balanceTreeRoot,
		tx.Data,
		fromMP.State.TokenID,
		create2transfer.TypesStateMerkleProof{State: create2transfer.TypesUserState(fromMP.State), Witness: fromMP.Witness},
		create2transfer.TypesStateMerkleProof{State: create2transfer.TypesUserState(toMP.State), Witness: toMP.Witness},
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

	result, err := b.SC.Transfer.Process(
		&opts,
		balanceTreeRoot,
		tx.Data,
		fromMP.State.TokenID,
		transfer.TypesStateMerkleProof{State: transfer.TypesUserState(fromMP.State), Witness: fromMP.Witness},
		transfer.TypesStateMerkleProof{State: transfer.TypesUserState(toMP.State), Witness: toMP.Witness},
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

func (b *Bazooka) processMassMigrationTx(balanceTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot ByteArray, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion()
	if err != nil {
		return
	}

	result, err := b.SC.MassMigration.Process(
		&opts,
		balanceTreeRoot,
		tx.Data,
		fromMP.State.TokenID,
		massmigration.TypesStateMerkleProof{State: massmigration.TypesUserState(fromMP.State), Witness: fromMP.Witness},
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

func (b *Bazooka) applyTransferTx(sender, receiver []byte, tx Tx) (updatedSender, updatedReceiver []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	updates, err := b.SC.Transfer.ValidateAndApply(&opts, sender, receiver, tx.Data)
	if err != nil {
		return
	}

	if err = ParseResult(updates.Result); err != nil {
		return
	}

	return updates.NewSender, updates.NewReceiver, nil
}

func (b *Bazooka) applyCreate2TransferTx(sender, receiver []byte, tx Tx) (updatedSender, updatedReceiver []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	updates, err := b.SC.Create2Transfer.ValidateAndApply(&opts, sender, tx.Data)
	if err != nil {
		return
	}

	if err = ParseResult(updates.Result); err != nil {
		return
	}

	return updates.NewSender, updates.NewReceiver, nil
}

func (b *Bazooka) applyMassMigrationTx(sender, receiver []byte, tx Tx) (updatedSender, updatedReceiver []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	updates, err := b.SC.MassMigration.ValidateAndApply(&opts, sender, tx.Data)
	if err != nil {
		return
	}

	if err = ParseResult(updates.Result); err != nil {
		return
	}

	return updates.NewSender, receiver, nil
}

func (b *Bazooka) compressTransferTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
	return b.SC.Transfer.Compress(&opts, data)
}

func (b *Bazooka) compressCreate2TransferTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
	return b.SC.Create2Transfer.Compress(&opts, data)
}

func (b *Bazooka) compressMassMigrationTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
	return b.SC.MassMigration.Compress(&opts, data)
}

func (b *Bazooka) TransferSignBytes(tx Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.SC.Transfer.SignBytes(&opts, tx.Data)
}

func (b *Bazooka) Create2TransferSignBytesWithPub(tx Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.SC.Create2Transfer.SignBytes(&opts, tx.Data)
}

func (b *Bazooka) MassMigrationSignBytes(tx Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.SC.MassMigration.SignBytes(&opts, tx.Data)
}

func (b *Bazooka) DecompressTransferTxs(txs []byte) (froms, tos, amounts, fees []big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTxs, err := b.SC.Transfer.Decompress(&opts, txs)
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

func (b *Bazooka) DecompressCreate2TransferTxs(txs []byte) (froms, tos, toAccIDs, amounts, fees []big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTxs, err := b.SC.Create2Transfer.Decompress(&opts, txs)
	if err != nil {
		return
	}

	for _, decompressedTx := range decompressedTxs {
		froms = append(froms, *decompressedTx.FromIndex)
		tos = append(tos, *decompressedTx.ToIndex)
		toAccIDs = append(tos, *decompressedTx.ToPubkeyID)
		amounts = append(amounts, *decompressedTx.Amount)
		fees = append(fees, *decompressedTx.Fee)
	}

	return
}

func (b *Bazooka) DecompressMassMigrationTxs(txs []byte) (froms, amounts, fees []big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTxs, err := b.SC.MassMigration.Decompress(&opts, txs)
	if err != nil {
		return
	}

	for _, decompressedTx := range decompressedTxs {
		froms = append(froms, *decompressedTx.FromIndex)
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
	return b.SC.Transfer.Encode(&opts, tx)
}

func (b *Bazooka) DecodeTransferTx(txBytes []byte) (from, to, nonce, txType, amount, fee *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.SC.Transfer.Decode(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.FromIndex, tx.ToIndex, tx.Nonce, tx.TxType, tx.Amount, tx.Fee, nil
}

func (b *Bazooka) EncodeCreate2TransferTx(from, to, toAccID, fee, nonce, amount, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx := struct {
		TxType     *big.Int
		FromIndex  *big.Int
		ToIndex    *big.Int
		ToPubkeyID *big.Int
		Amount     *big.Int
		Fee        *big.Int
		Nonce      *big.Int
	}{big.NewInt(txType), big.NewInt(from), big.NewInt(to), big.NewInt(toAccID), big.NewInt(amount), big.NewInt(fee), big.NewInt(nonce)}
	return b.SC.Create2Transfer.Encode(&opts, tx)
}

func (b *Bazooka) DecodeCreate2Transfer(txBytes []byte) (from, to, toAccID, nonce, txType, amount, fee *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.SC.Create2Transfer.Decode(&opts, txBytes)
	if err != nil {
		return
	}

	return tx.FromIndex, tx.ToIndex, tx.ToPubkeyID, tx.Nonce, tx.TxType, tx.Amount, tx.Fee, nil
}

func (b *Bazooka) EncodeCreate2TransferTxWithPub(from int64, toPub [4]*big.Int, fee, nonce, amount, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx := struct {
		TxType    *big.Int
		FromIndex *big.Int
		ToPubkey  [4]*big.Int
		Amount    *big.Int
		Fee       *big.Int
		Nonce     *big.Int
	}{big.NewInt(txType), big.NewInt(from), toPub, big.NewInt(amount), big.NewInt(fee), big.NewInt(nonce)}
	return b.SC.Create2Transfer.EncodeWithPub(&opts, tx)
}

func (b *Bazooka) DecodeCreate2TransferWithPub(txBytes []byte) (fromIndex *big.Int, toPub [4]*big.Int, nonce, txType, amount, fee *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.SC.Create2Transfer.DecodeWithPub(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.FromIndex, tx.ToPubkey, tx.Nonce, tx.TxType, tx.Amount, tx.Fee, nil
}

func (b *Bazooka) EncodeMassMigrationTx(from, toSpoke, fee, nonce, amount, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx := struct {
		TxType    *big.Int
		FromIndex *big.Int
		Amount    *big.Int
		Fee       *big.Int
		SpokeID   *big.Int
		Nonce     *big.Int
	}{big.NewInt(txType), big.NewInt(from), big.NewInt(amount), big.NewInt(fee), big.NewInt(toSpoke), big.NewInt(nonce)}
	return b.SC.MassMigration.Encode(&opts, tx)
}

func (b *Bazooka) DecodeMassMigrationTx(txBytes []byte) (from, toSpoke, nonce, txType, amount, fee *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.SC.MassMigration.Decode(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.FromIndex, tx.SpokeID, tx.Nonce, tx.TxType, tx.Amount, tx.Fee, nil
}

//
// Encoders and Decoders for state
//

func (b *Bazooka) EncodeState(id, balance, nonce, token uint64) (accountBytes []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	accountBytes, err = b.SC.State.Encode(&opts, state.TypesUserState{
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

	state, err := b.SC.State.DecodeState(&opts, stateBytes)
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
		return "", nil
	}

	auth, err := b.generateAuthObj(b.EthClient, rollupAddress, stakeAmount, data)
	if err != nil {
		b.log.Error("Estimate gas failed, tx reverting", "error", err)
		return "", nil
	}

	tx, err := b.SC.RollupContract.SubmitTransfer(auth, updatedRoots, aggregatedSig, feeReceivers, txs)
	if err != nil {
		b.log.Error("Error submitting batch", "err", err)
		return "", nil
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

func getContractInstances(client *ethclient.Client) (contracts Contracts, err error) {
	if contracts.RollupContract, err = rollup.NewRollup(ethCmn.HexToAddress(config.GlobalCfg.RollupAddress), client); err != nil {
		return contracts, err
	}
	if contracts.AccountRegistry, err = accountregistry.NewAccountregistry(ethCmn.HexToAddress(config.GlobalCfg.AccountRegistry), client); err != nil {
		return contracts, err
	}
	if contracts.State, err = state.NewState(ethCmn.HexToAddress(config.GlobalCfg.State), client); err != nil {
		return contracts, err
	}
	if contracts.Transfer, err = transfer.NewTransfer(ethCmn.HexToAddress(config.GlobalCfg.Transfer), client); err != nil {
		return contracts, err
	}
	if contracts.Create2Transfer, err = create2transfer.NewCreate2transfer(ethCmn.HexToAddress(config.GlobalCfg.Create2Transfer), client); err != nil {
		return contracts, err
	}
	if contracts.MassMigration, err = massmigration.NewMassmigration(ethCmn.HexToAddress(config.GlobalCfg.MassMigration), client); err != nil {
		return contracts, err
	}
	return contracts, nil
}
