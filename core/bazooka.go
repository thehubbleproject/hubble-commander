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
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/BOPR/contracts/logger"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/contracts/rollupcaller"
	"github.com/BOPR/contracts/rolluputils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

// IContractCaller is the common interface using which we will interact with the contracts
// and the ethereum chain
type IBazooka interface {
	FetchBatchInputData(txHash ethCmn.Hash) (txs [][]byte, err error)
}

// Global Contract Caller Object
var LoadedBazooka Bazooka

// ContractCaller satisfies the IContractCaller interface and contains all the variables required to interact
// With the ethereum chain along with contract addresses and ABI's
type Bazooka struct {
	log log.Logger

	EthClient *ethclient.Client

	ContractABI map[string]abi.ABI

	// Rollup contract
	RollupContract *rollup.Rollup
	EventLogger    *logger.Logger
	RollupUtils    *rolluputils.Rolluputils
	RollupCaller   *rollupcaller.Rollupcaller
}

// NewContractCaller contract caller
// NOTE: Reads configration from the config.toml file
func NewPreLoadedBazooka() (bazooka Bazooka, err error) {
	// TODO remove

	// err = config.ParseAndInitGlobalConfig()
	// if err != nil {
	// 	return
	// }
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

	// initialise all variables for rollup utils contract
	rollupUtilsAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupUtilsAddress)
	if bazooka.RollupUtils, err = rolluputils.NewRolluputils(rollupUtilsAddress, bazooka.EthClient); err != nil {
		return bazooka, err
	}
	if bazooka.ContractABI[common.ROLLUP_UTILS], err = abi.JSON(strings.NewReader(rolluputils.RolluputilsABI)); err != nil {
		return bazooka, err
	}

	// initialise all variables for event logger contract
	rollupcallerAddr := ethCmn.HexToAddress(config.GlobalCfg.RollupCallerAddress)
	if bazooka.RollupCaller, err = rollupcaller.NewRollupcaller(rollupcallerAddr, bazooka.EthClient); err != nil {
		return bazooka, err
	}
	if bazooka.ContractABI[common.ROLLUP_CALLER], err = abi.JSON(strings.NewReader(rollupcaller.RollupcallerABI)); err != nil {
		return bazooka, err
	}

	bazooka.log = common.Logger.With("module", "bazooka")

	return bazooka, nil
}

// get main chain block header
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

func (b *Bazooka) FetchBalanceTreeRoot() (ByteArray, error) {
	root, err := b.RollupContract.GetLatestBalanceTreeRoot(nil)
	if err != nil {
		return ByteArray{}, err
	}
	return root, nil
}

func (b *Bazooka) FetchBatchInputData(txHash ethCmn.Hash) (txs [][]byte, err error) {
	tx, isPending, err := b.EthClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		b.log.Error("Cannot fetch transaction from hash", "Error", err)
		return
	}

	if isPending {
		err := errors.New("Transaction is pending")
		b.log.Error("Transaction is still pending, cannot process", "Error", err)
		return txs, err
	}

	payload := tx.Data()
	decodedPayload := payload[4:]
	inputDataMap := make(map[string]interface{})
	method := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Methods["submitBatch"]
	err = method.Inputs.UnpackIntoMap(inputDataMap, decodedPayload)
	if err != nil {
		b.log.Error("Error unpacking payload", "Error", err)
		return
	}

	return GetTxsFromInput(inputDataMap), nil
}

// ProcessTx calls the ProcessTx function on the contract to verify the tx
// returns the updated accounts and the new balance root
func (b *Bazooka) ProcessTx(balanceTreeRoot, accountTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof AccountMerkleProof, pdaProof PDAMerkleProof) (newBalanceRoot ByteArray, from, to []byte, err error) {
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return b.processTransferTx(balanceTreeRoot, accountTreeRoot, tx, fromMerkleProof, toMerkleProof, pdaProof)
	case TX_AIRDROP_TYPE:
		return b.processAirdropTx(balanceTreeRoot, accountTreeRoot, tx, fromMerkleProof, toMerkleProof, pdaProof)
	case TX_CREATE_ACCOUNT:
		return b.processCreateAccountTx(balanceTreeRoot, accountTreeRoot, tx, fromMerkleProof, toMerkleProof, pdaProof)
	case TX_BURN_CONSENT:
		return b.processBurnConsentTx(balanceTreeRoot, accountTreeRoot, tx, fromMerkleProof, toMerkleProof, pdaProof)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return
	}
}

func (b *Bazooka) ApplyTx(accountMP AccountMerkleProof, tx Tx) (updatedAccount []byte, updatedRoot ByteArray, err error) {
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return b.applyTransferTx(accountMP, tx)
	case TX_AIRDROP_TYPE:
		return b.applyAirdropTx(accountMP, tx)
	case TX_CREATE_ACCOUNT:
		return b.applyCreateAccountTx(accountMP, tx)
	case TX_BURN_CONSENT:
		return b.applyBurnConsentTx(accountMP, tx)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return updatedAccount, updatedRoot, errors.New("Didn't match any options")
	}
}

func (b *Bazooka) CompressTx(tx *Tx) ([]byte, error) {
	switch txType := tx.Type; txType {
	case TX_TRANSFER_TYPE:
		return LoadedBazooka.CompressTransferTx(*tx)
	case TX_AIRDROP_TYPE:
		return LoadedBazooka.CompressAirdropTx(*tx)
	case TX_CREATE_ACCOUNT:
		return LoadedBazooka.CompressCreateAccount(*tx)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return []byte(""), nil
	}
}

func (b *Bazooka) processCreateAccountTx(balanceTreeRoot, accountTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof AccountMerkleProof, pdaProof PDAMerkleProof) (newBalanceRoot ByteArray, from, to []byte, err error) {
	decodedSignature, _ := hex.DecodeString(tx.Signature)
	opts := bind.CallOpts{From: config.OperatorAddress}
	toMP, err := toMerkleProof.ToABIVersion()
	if err != nil {
		return
	}
	updatedRoot, newToAccount, errCode, IsValidTx, err := b.RollupCaller.ProcessCreateAccountTx(&opts,
		balanceTreeRoot,
		accountTreeRoot,
		decodedSignature,
		tx.Data,
		pdaProof.ToABIVersion(),
		toMP,
	)
	if err != nil {
		return
	}

	b.log.Info("Processed transaction", "IsSuccess", IsValidTx, "newRoot", updatedRoot)

	if !IsValidTx {
		b.log.Error("Invalid transaction", "error_code", errCode)
		return newBalanceRoot, from, to, errors.New("Tx is invalid")
	}
	newBalanceRoot = BytesToByteArray(updatedRoot[:])
	return newBalanceRoot, fromMerkleProof.Account.Data, newToAccount, nil
}

func (b *Bazooka) processTransferTx(balanceTreeRoot, accountTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof AccountMerkleProof, pdaProof PDAMerkleProof) (newBalanceRoot ByteArray, from, to []byte, err error) {
	decodedSignature, _ := hex.DecodeString(tx.Signature)
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion()
	if err != nil {
		return
	}
	toMP, err := toMerkleProof.ToABIVersion()
	if err != nil {
		return
	}
	typesAccountProofs := rollupcaller.TypesAccountProofs{From: fromMP, To: toMP}
	updatedRoot, newFromAccount, newToAccount, errCode, IsValidTx, err := b.RollupCaller.ProcessTransferTx(&opts,
		balanceTreeRoot,
		accountTreeRoot,
		decodedSignature,
		tx.Data,
		pdaProof.ToABIVersion(),
		typesAccountProofs,
	)
	if err != nil {
		return
	}

	b.log.Info("Processed transaction", "IsSuccess", IsValidTx, "newRoot", updatedRoot)

	if !IsValidTx {
		b.log.Error("Invalid transaction", "error_code", errCode)
		return newBalanceRoot, from, to, errors.New("Tx is invalid")
	}
	newBalanceRoot = BytesToByteArray(updatedRoot[:])
	return newBalanceRoot, newFromAccount, newToAccount, nil
}

func (b *Bazooka) processAirdropTx(balanceTreeRoot, accountTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof AccountMerkleProof, pdaProof PDAMerkleProof) (newBalanceRoot ByteArray, from, to []byte, err error) {
	decodedSignature, _ := hex.DecodeString(tx.Signature)
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion()
	if err != nil {
		return
	}
	toMP, err := toMerkleProof.ToABIVersion()
	if err != nil {
		return
	}
	typesAccountProofs := rollupcaller.TypesAccountProofs{From: fromMP, To: toMP}
	updatedRoot, newFromAccount, newToAccount, errCode, IsValidTx, err := b.RollupCaller.ProcessAirdropTx(&opts,
		balanceTreeRoot,
		accountTreeRoot,
		decodedSignature,
		tx.Data,
		pdaProof.ToABIVersion(),
		typesAccountProofs,
	)
	if err != nil {
		return
	}

	b.log.Info("Processed transaction", "IsSuccess", IsValidTx, "newRoot", updatedRoot)

	if !IsValidTx {
		b.log.Error("Invalid transaction", "error_code", errCode)
		return newBalanceRoot, from, to, errors.New("Tx is invalid")
	}
	newBalanceRoot = BytesToByteArray(updatedRoot[:])
	return newBalanceRoot, newFromAccount, newToAccount, nil
}

func (b *Bazooka) processBurnConsentTx(balanceTreeRoot, accountTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof AccountMerkleProof, pdaProof PDAMerkleProof) (newBalanceRoot ByteArray, from, to []byte, err error) {
	decodedSignature, _ := hex.DecodeString(tx.Signature)
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion()
	if err != nil {
		return
	}
	updatedRoot, newFromAccount, errCode, IsValidTx, err := b.RollupCaller.ProcessBurnConsentTx(&opts,
		balanceTreeRoot,
		accountTreeRoot,
		decodedSignature,
		tx.Data,
		pdaProof.ToABIVersion(),
		fromMP,
	)
	if err != nil {
		return
	}

	b.log.Info("Processed transaction", "IsSuccess", IsValidTx, "newRoot", updatedRoot)

	if !IsValidTx {
		b.log.Error("Invalid transaction", "error_code", errCode)
		return newBalanceRoot, from, to, errors.New("Tx is invalid")
	}
	newBalanceRoot = BytesToByteArray(updatedRoot[:])
	return newBalanceRoot, newFromAccount, toMerkleProof.Account.Data, nil
}

func (b *Bazooka) applyTransferTx(accountMP AccountMerkleProof, tx Tx) ([]byte, ByteArray, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	accMP, err := accountMP.ToABIVersion()
	if err != nil {
		return nil, ByteArray{}, err
	}
	updatedAccountBytes, updatedRoot, err := b.RollupCaller.ApplyTransferTx(&opts, accMP, tx.Data)
	if err != nil {
		return updatedAccountBytes, updatedRoot, err
	}
	return updatedAccountBytes, updatedRoot, nil
}

func (b *Bazooka) applyAirdropTx(accountMP AccountMerkleProof, tx Tx) ([]byte, ByteArray, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	accMP, err := accountMP.ToABIVersion()
	if err != nil {
		return nil, ByteArray{}, err
	}
	updatedAccountBytes, updatedRoot, err := b.RollupCaller.ApplyAirdropTx(&opts, accMP, tx.Data)
	if err != nil {
		return updatedAccountBytes, updatedRoot, err
	}
	return updatedAccountBytes, updatedRoot, nil
}

func (b *Bazooka) applyCreateAccountTx(accountMP AccountMerkleProof, tx Tx) ([]byte, ByteArray, error) {
	// if the account is a from account, return without changes
	if accountMP.Account.AccountID == tx.From {
		return accountMP.Account.Data, ByteArray{}, nil
	}

	opts := bind.CallOpts{From: config.OperatorAddress}
	accMP, err := accountMP.ToABIVersion()
	if err != nil {
		return nil, ByteArray{}, err
	}
	updatedAccountBytes, updatedRoot, err := b.RollupCaller.ApplyAirdropTx(&opts, accMP, tx.Data)
	if err != nil {
		return updatedAccountBytes, updatedRoot, err
	}
	return updatedAccountBytes, updatedRoot, nil
}

func (b *Bazooka) applyBurnConsentTx(accountMP AccountMerkleProof, tx Tx) ([]byte, ByteArray, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	accMP, err := accountMP.ToABIVersion()
	if err != nil {
		return nil, ByteArray{}, err
	}
	postApplyState, err := b.RollupCaller.ApplyBurnConsentTx(&opts, accMP, tx.Data)
	if err != nil {
		return postApplyState.UpdatedAccount, postApplyState.NewRoot, err
	}
	return postApplyState.UpdatedAccount, postApplyState.NewRoot, nil
}

func (b *Bazooka) CompressTransferTx(tx Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	sigBytes, err := hex.DecodeString(tx.Signature[2:])
	if err != nil {
		return nil, err
	}
	return b.RollupUtils.CompressTxWithMessage(&opts, tx.Data, sigBytes)
}

func (b *Bazooka) CompressAirdropTx(tx Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	sigBytes, err := hex.DecodeString(tx.Signature[2:])
	if err != nil {
		return nil, err
	}
	return b.RollupUtils.CompressAirdropTxWithMessage(&opts, tx.Data, sigBytes)
}

func (b *Bazooka) CompressCreateAccount(tx Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.CompressCreateAccountWithMessage(&opts, tx.Data)
}
func (b *Bazooka) CompressBurnConsent(tx Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	sigBytes, err := hex.DecodeString(tx.Signature[2:])
	if err != nil {
		return nil, err
	}
	return b.RollupUtils.CompressBurnConsentTxWithMessage(&opts, tx.Data, sigBytes)
}

func (b *Bazooka) CompressBurnExecution(tx Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.CompressBurnExecutionWithMessage(&opts, tx.Data)
}

func (b *Bazooka) DecompressTransferTxs(compressedTxs [][]byte) (from, to, amount []*big.Int, sig [][]byte, err error) {
	for _, compressedTx := range compressedTxs {
		opts := bind.CallOpts{From: config.OperatorAddress}
		decompressedTx, err := b.RollupUtils.DecompressTx(&opts, compressedTx)
		if err != nil {
			return from, to, amount, sig, err
		}
		from = append(from, decompressedTx.From)
		to = append(to, decompressedTx.To)
		amount = append(amount, decompressedTx.Amount)
		sig = append(sig, decompressedTx.Sig)
	}
	return from, to, amount, sig, nil
}

func (b *Bazooka) DecompressCreateAccountTxs(compressedTxs [][]byte) (to, tokenType []*big.Int, err error) {
	for _, compressedTx := range compressedTxs {
		opts := bind.CallOpts{From: config.OperatorAddress}
		decompressedTx, err := b.RollupUtils.DecompressCreateAccount(&opts, compressedTx)
		if err != nil {
			return to, tokenType, err
		}
		to = append(to, decompressedTx.ToIndex)
		tokenType = append(tokenType, decompressedTx.TokenType)
	}
	return to, tokenType, nil
}

func (b *Bazooka) DecompressBurnConsentTxs(compressedTxs [][]byte) (from, amount, nonce []*big.Int, cancel []bool, sig [][]byte, err error) {
	for _, compressedTx := range compressedTxs {
		opts := bind.CallOpts{From: config.OperatorAddress}
		decompressedTx, err := b.RollupUtils.DecompressBurnConsent(&opts, compressedTx)
		if err != nil {
			return from, amount, nonce, cancel, sig, err
		}

		from = append(from, decompressedTx.FromIndex)
		amount = append(amount, decompressedTx.Amount)
		nonce = append(nonce, decompressedTx.Nonce)
		cancel = append(cancel, decompressedTx.Cancel)
		sig = append(sig, decompressedTx.Signature)
	}
	return from, amount, nonce, cancel, sig, nil
}

func (b *Bazooka) DecompressBurnExecTxs(compressedTxs [][]byte) (from []*big.Int, err error) {
	for _, compressedTx := range compressedTxs {
		opts := bind.CallOpts{From: config.OperatorAddress}
		decompressedTxFromIndex, err := b.RollupUtils.DecompressBurnExecution(&opts, compressedTx)
		if err != nil {
			return from, err
		}

		from = append(from, decompressedTxFromIndex)
	}
	return from, nil
}

//
// Encoders and Decoders for transactions
//

func (b *Bazooka) EncodeTransferTx(from, to, token, nonce, amount, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromTxDeconstructed(&opts, big.NewInt(from), big.NewInt(to), big.NewInt(token), big.NewInt(nonce), big.NewInt(txType), big.NewInt(amount))
}

func (b *Bazooka) DecodeTransferTx(txBytes []byte) (from, to, token, nonce, txType, amount *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.RollupUtils.TxFromBytesDeconstructed(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.From, tx.To, tx.TokenType, tx.Nonce, tx.TxType, tx.Amount, nil
}

func (b *Bazooka) EncodeAirdropTx(from, to, token, nonce, amount, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromAirdropNoStruct(&opts, big.NewInt(from), big.NewInt(to), big.NewInt(token), big.NewInt(nonce), big.NewInt(txType), big.NewInt(amount))
}

func (b *Bazooka) DecodeAirdropTx(txBytes []byte) (from, to, token, nonce, txType, amount *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.RollupUtils.AirdropFromBytesNoStruct(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.From, tx.To, tx.TokenType, tx.Nonce, tx.TxType, tx.Amount, nil
}

func (b *Bazooka) EncodeCreateAccountTx(to, token int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromCreateAccountNoStruct(&opts, big.NewInt(to), big.NewInt(token))
}

func (b *Bazooka) DecodeCreateAccountTx(txBytes []byte) (to, token *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.RollupUtils.CreateAccountFromBytes(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.ToIndex, tx.TokenType, nil
}

func (b *Bazooka) EncodeBurnConsentTx(from, amount, nonce, txType int64, cancel bool) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromBurnConsentNoStruct(&opts, big.NewInt(from), big.NewInt(amount), big.NewInt(nonce), cancel)
}

func (b *Bazooka) DecodeBurnConsentTx(txBytes []byte) (from, amount, nonce, txType *big.Int, cancel bool, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.RollupUtils.BurnConsentTxFromBytes(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.FromIndex, tx.Amount, tx.Nonce, big.NewInt(TX_BURN_CONSENT), tx.Cancel, nil
}

func (b *Bazooka) EncodeBurnExecTx(from, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromBurnExecutionNoStruct(&opts, big.NewInt(from))
}

func (b *Bazooka) DecodeBurnExecTx(txBytes []byte) (from, txType *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.RollupUtils.BurnExecutionFromBytes(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.FromIndex, big.NewInt(TX_BURN_EXEC), nil
}

//
// Encoders and Decoders for accounts
//

func (b *Bazooka) EncodeAccount(id, balance, nonce, token, burn, lastBurn int64) (accountBytes []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	accountBytes, err = b.RollupUtils.BytesFromAccountDeconstructed(&opts, big.NewInt(id), big.NewInt(balance), big.NewInt(nonce), big.NewInt(token), big.NewInt(burn), big.NewInt(lastBurn))
	if err != nil {
		return
	}
	return accountBytes, nil
}

func (b *Bazooka) DecodeAccount(accountBytes []byte) (ID, balance, nonce, token, burn, lastBurn *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	account, err := b.RollupUtils.AccountFromBytes(&opts, accountBytes)
	if err != nil {
		return
	}

	return account.ID, account.Balance, account.Nonce, account.TokenType, account.Burn, account.LastBurn, nil
}

// ----------------------------------------------------------------

func (b *Bazooka) GetGenesisAccounts() (genesisAccount []UserAccount, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	// get genesis accounts
	accounts, err := b.RollupUtils.GetGenesisDataBlocks(&opts)
	if err != nil {
		return
	}

	for _, account := range accounts {
		ID, _, _, _, _, _, _ := b.DecodeAccount(account)
		genesisAccount = append(genesisAccount, *NewUserAccount(ID.Uint64(), STATUS_ACTIVE, "", UintToString(ID.Uint64()), account))
	}
	return
}

//
// Transactions
//

func (b *Bazooka) FireDepositFinalisation(TBreplaced UserAccount, siblings []UserAccount, subTreeHeight uint64) (err error) {
	b.log.Info(
		"Attempting to finalise deposits",
		"NodeToBeReplaced",
		TBreplaced.String(),
		"NumberOfSiblings",
		len(siblings),
		"atDepth",
		subTreeHeight,
	)

	// TODO check latest batch on-chain nd if we need to push new batch
	depositSubTreeHeight := big.NewInt(0)
	depositSubTreeHeight.SetUint64(subTreeHeight)
	var siblingData [][32]byte
	for _, sibling := range siblings {
		data, err := HexToByteArray(sibling.Hash)
		if err != nil {
			fmt.Println("unable to convert HexToByteArray", err)
			return err
		}
		siblingData = append(siblingData, data)
	}

	accountProof := rollup.TypesAccountMerkleProof{}
	accountProof.AccountIP.PathToAccount = StringToBigInt(TBreplaced.Path)
	userAccount, err := TBreplaced.ToABIAccount()
	if err != nil {
		fmt.Println("unable to convert", "error", err)
		return
	}
	accountProof.AccountIP.Account = rollup.TypesUserAccount(userAccount)

	accountProof.Siblings = siblingData
	data, err := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Pack("finaliseDepositsAndSubmitBatch", depositSubTreeHeight, accountProof)
	if err != nil {
		fmt.Println("Unable to craete data", err)
		return
	}

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	stakeAmount := big.NewInt(0)
	stakeAmount.SetString("32000000000000000000", 10)
	fmt.Println("Stake committed", stakeAmount)

	// generate call msg
	callMsg := ethereum.CallMsg{
		To:    &rollupAddress,
		Data:  data,
		Value: stakeAmount,
	}

	auth, err := b.GenerateAuthObj(b.EthClient, callMsg)
	if err != nil {
		return err
	}
	b.log.Info("Broadcasting deposit finalisation transaction")

	tx, err := b.RollupContract.FinaliseDepositsAndSubmitBatch(auth, depositSubTreeHeight, accountProof)
	if err != nil {
		return err
	}
	b.log.Info("Deposits successfully finalized!", "TxHash", tx.Hash())
	return nil
}

// SubmitBatch submits the batch on chain with updated root and compressed transactions
func (b *Bazooka) SubmitBatch(updatedRoot ByteArray, txs []Tx) error {
	b.log.Info(
		"Attempting to submit a new batch",
		"UpdatedRoot",
		updatedRoot.String(),
		"txs",
		len(txs),
	)
	if len(txs) == 0 {
		b.log.Info("No transactions to submit, waiting....")
		return nil
	}
	var compressedTxs [][]byte
	for _, tx := range txs {
		compressedTx, err := b.CompressTx(&tx)
		if err != nil {
			return err
		}
		compressedTxs = append(compressedTxs, compressedTx)
	}

	data, err := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Pack("submitBatch", compressedTxs, updatedRoot, uint8(txs[0].Type))
	if err != nil {
		return err
	}

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	stakeAmount := big.NewInt(0)
	stakeAmount.SetString("32000000000000000000", 10)

	// generate call msg
	callMsg := ethereum.CallMsg{
		To:    &rollupAddress,
		Data:  data,
		Value: stakeAmount,
	}

	auth, err := b.GenerateAuthObj(b.EthClient, callMsg)
	if err != nil {
		return err
	}

	latestBatch, err := DBInstance.GetLatestBatch()
	if err != nil {
		return err
	}

	newBatch := Batch{
		BatchID:   latestBatch.BatchID + 1,
		StateRoot: updatedRoot.String(),
		Committer: config.OperatorAddress.String(),
		Status:    BATCH_BROADCASTED,
	}
	b.log.Info("Broadcasting a new batch", "newBatch", newBatch)
	err = DBInstance.AddNewBatch(newBatch)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	tx, err := b.RollupContract.SubmitBatch(auth, compressedTxs, updatedRoot, uint8(txs[0].Type))
=======
<<<<<<< HEAD
	tx, err := b.RollupContract.SubmitBatch(auth, compressedTxs, updatedRoot,uint8(txs[0].Type))
=======
	tx, err := b.RollupContract.SubmitBatch(auth, compressedTxs, updatedRoot, uint8(txs[0].Type))
>>>>>>> dc2015b843fd4d5dc29d39759213738d8c28bfd6
>>>>>>> bb7321667f07d386de963ac1ff64b0af4b439129
	if err != nil {
		return err
	}

	b.log.Info("Sent a new batch!", "txHash", tx.Hash().String())
	return nil
}

func GetTxsFromInput(input map[string]interface{}) (txs [][]byte) {
	data := input["_txs"].([][]byte)
	return data
}

func (b *Bazooka) GenerateAuthObj(client *ethclient.Client, callMsg ethereum.CallMsg) (auth *bind.TransactOpts, err error) {
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

	// create auth
	auth = bind.NewKeyedTransactor(config.OperatorKey)
	auth.GasPrice = gasprice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(gasLimit) // uint64(gasLimit)
	return
}
