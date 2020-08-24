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
	"github.com/tendermint/tendermint/libs/log"

	"github.com/BOPR/contracts/logger"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/contracts/rollupcaller"
	"github.com/BOPR/contracts/rolluputils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Broadcast details is stored to the DB with each broadcast
// It is useful because now we can broadcast txs simulatneously without caring about confirmation times
type BroadcastDetails struct {
	DBModel
	TxHash     string `json:"tx"`
	Nonce      uint64 `json:"nonce"`
	BatchType  uint64 `json:"batchType"`
	UpdateRoot string `json:"updateRoot"`
	Txs        string `gorm:"size:1000000"`
	Status     uint64 `json:"status"`
}

func (db *DB) LogBatch(nonce uint64, batchType uint64, root string, txs []byte) error {
	var details BroadcastDetails
	// details.TxHash = txHash
	details.Nonce = nonce
	details.BatchType = batchType
	details.UpdateRoot = root
	details.Txs = hex.EncodeToString(txs)
	details.Status = 0
	return db.Instance.Create(&details).Error
}

func (db *DB) GetLastTransaction() (BroadcastDetails, error) {
	var details BroadcastDetails
	if err := db.Instance.Order("nonce desc").First(&details).Error; err != nil {
		return details, err
	}
	return details, nil
}

// GetPendingBatches get the pending batches
func (db *DB) GetPendingBatches() ([]BroadcastDetails, error) {
	var details []BroadcastDetails
	if err := db.Instance.Order("created_at").Find(&details).Error; err != nil {
		db.Logger.Error("error while fetching pending batches", err)
		return details, err
	}
	return details, nil
}

// MarkBroadcastDone updates the status of the broadcast
func (db *DB) MarkBroadcastDone(id string, txHash string) error {
	return db.Instance.Where("id = ? AND status = ?", id, 0).Updates(BroadcastDetails{TxHash: txHash, Status: 1}).Error
}

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
	log       log.Logger
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

func (b *Bazooka) DecompressTransferTx(compressedTx []byte) (from, to, amount *big.Int, sig []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTx, err := b.RollupUtils.DecompressTx(&opts, compressedTx)
	if err != nil {
		return from, to, amount, sig, err
	}
	return decompressedTx.From, decompressedTx.To, decompressedTx.Amount, sig, nil
}

func (b *Bazooka) DecompressCreateAccountTx(compressedTx []byte) (to, stateID, tokenType *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTx, err := b.RollupUtils.DecompressCreateAccount(&opts, compressedTx)
	if err != nil {
		return to, stateID, tokenType, err
	}
	return decompressedTx.AccountID, decompressedTx.StateID, decompressedTx.TokenType, nil
}

func (b *Bazooka) DecompressBurnConsentTx(compressedTx []byte) (from, amount, nonce *big.Int, sig []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTx, err := b.RollupUtils.DecompressBurnConsent(&opts, compressedTx)
	if err != nil {
		return from, amount, nonce, sig, err
	}
	return decompressedTx.FromIndex, decompressedTx.Amount, decompressedTx.Nonce, sig, nil
}

func (b *Bazooka) DecompressBurnExecTx(compressedTx []byte) (from *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTxFromIndex, err := b.RollupUtils.DecompressBurnExecution(&opts, compressedTx)
	if err != nil {
		return from, err
	}
	return decompressedTxFromIndex, nil
}

func (b *Bazooka) DecompressAirdropTx(compressedTx []byte) (from, to, amount *big.Int, sig []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	decompressedTx, err := b.RollupUtils.DecompressAirdrop(&opts, compressedTx)
	if err != nil {
		return
	}
	return big.NewInt(1), decompressedTx.ToIndex, decompressedTx.Amount, decompressedTx.Signature, nil
}

func (b *Bazooka) SignBytesForTransfer(txType, fromIndex, toIndex, nonce, amount int64) ([32]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.GetTxSignBytes(&opts, big.NewInt(txType), big.NewInt(fromIndex), big.NewInt(toIndex), big.NewInt(nonce), big.NewInt(amount))
}

func (b *Bazooka) SignBytesForAirdrop(txType, fromIndex, toIndex, nonce, amount int64) ([32]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.AirdropSignBytes(&opts, big.NewInt(txType), big.NewInt(fromIndex), big.NewInt(toIndex), big.NewInt(nonce), big.NewInt(amount))
}
func (b *Bazooka) SignBytesForBurnConsent(txType, from, nonce, amount int64) ([32]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BurnConsentSignBytes(&opts, big.NewInt(txType), big.NewInt(from), big.NewInt(nonce), big.NewInt(amount))
}

//
// Encoders and Decoders for transactions
//

func (b *Bazooka) EncodeTransferTx(from, to, token, nonce, amount, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromTxDeconstructed(&opts, big.NewInt(txType), big.NewInt(from), big.NewInt(to), big.NewInt(token), big.NewInt(nonce), big.NewInt(amount))
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
	return b.RollupUtils.BytesFromAirdropNoStruct(&opts, big.NewInt(txType), big.NewInt(from), big.NewInt(to), big.NewInt(token), big.NewInt(nonce), big.NewInt(amount))
}

func (b *Bazooka) DecodeAirdropTx(txBytes []byte) (from, to, token, nonce, txType, amount *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.RollupUtils.AirdropFromBytesNoStruct(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.From, tx.To, tx.TokenType, tx.Nonce, tx.TxType, tx.Amount, nil
}

func (b *Bazooka) EncodeCreateAccountTx(to, stateID, token int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromCreateAccountNoStruct(&opts, big.NewInt(TX_CREATE_ACCOUNT), big.NewInt(to), big.NewInt(stateID), big.NewInt(token))
}

func (b *Bazooka) DecodeCreateAccountTx(txBytes []byte) (toAccID, toStateID, token *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.RollupUtils.CreateAccountFromBytes(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.AccountID, tx.StateID, tx.TokenType, nil
}

func (b *Bazooka) EncodeBurnConsentTx(from, amount, nonce, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromBurnConsentNoStruct(&opts, big.NewInt(TX_BURN_CONSENT), big.NewInt(from), big.NewInt(amount), big.NewInt(nonce))
}

func (b *Bazooka) DecodeBurnConsentTx(txBytes []byte) (from, amount, nonce, txType *big.Int, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	tx, err := b.RollupUtils.BurnConsentFromBytes(&opts, txBytes)
	if err != nil {
		return
	}
	return tx.FromIndex, tx.Amount, tx.Nonce, big.NewInt(TX_BURN_CONSENT), nil
}

func (b *Bazooka) EncodeBurnExecTx(from, txType int64) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.RollupUtils.BytesFromBurnExecutionNoStruct(&opts, big.NewInt(TX_BURN_EXEC), big.NewInt(from))
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

	b.log.Debug("Decoded account", "ID", account.ID, "balance", account.Balance, "token", account.TokenType, "nonce", account.Nonce, "burn", account.Burn, "lastBurn", account.LastBurn)
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
		genesisAccount = append(genesisAccount, *NewUserAccount(ID.Uint64(), STATUS_ACTIVE, UintToString(ID.Uint64()), account))
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
			b.log.Error("unable to convert HexToByteArray", err)
			return err
		}
		siblingData = append(siblingData, data)
	}

	accountProof := rollup.TypesAccountMerkleProof{}
	accountProof.AccountIP.PathToAccount = StringToBigInt(TBreplaced.Path)
	userAccount, err := TBreplaced.ToABIAccount()
	if err != nil {
		b.log.Error("unable to convert", "error", err)
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
	lastTxBroadcasted, err := DBInstance.GetLastTransaction()
	if err != nil {
		return err
	}
	if lastTxBroadcasted.Nonce+1 != auth.Nonce.Uint64() {
		b.log.Info("Replacing nonce", "nonceEstimated", auth.Nonce.String(), "replacedBy", lastTxBroadcasted.Nonce+1)
		auth.Nonce = big.NewInt(int64(lastTxBroadcasted.Nonce + 1))
	}
	b.log.Info("Broadcasting deposit finalisation transaction")

	tx, err := b.RollupContract.FinaliseDepositsAndSubmitBatch(auth, depositSubTreeHeight, accountProof)
	if err != nil {
		return err
	}
	// TODO change this to deposit type
	err = DBInstance.LogBatch(auth.Nonce.Uint64(), 100, "", []byte(""))
	if err != nil {
		return err
	}
	b.log.Info("Deposits successfully finalized!", "TxHash", tx.Hash())
	return nil
}

// SubmitBatch submits the batch on chain with updated root and compressed transactions
func (b *Bazooka) SubmitBatch(commitments []Commitment) error {
	b.log.Info(
		"Attempting to submit a new batch",
		"NumOfCommitments",
		len(commitments),
	)

	if len(commitments) == 0 {
		b.log.Info("No transactions to submit, waiting....")
		return nil
	}

	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTransactionsBeingCommitted int
	for _, commitment := range commitments {
		compressedTxs, err := b.CompressTxs(commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, commitment.UpdatedRoot)
		totalTransactionsBeingCommitted += len(commitment.Txs)
		sig1 := commitment.AggregatedSignature[0:32]
		sig2 := commitment.AggregatedSignature[32:64]
		sig1bigInt := big.NewInt(0)
		sig1bigInt.SetBytes(sig1)
		sig2bigInt := big.NewInt(0)
		sig2bigInt.SetBytes(sig2)
		aggregatedSigBigInt := [2]*big.Int{sig1bigInt, sig2bigInt}
		fmt.Println("creeated aggregated sig", aggregatedSigBigInt)
		aggregatedSig = append(aggregatedSig, aggregatedSigBigInt)
	}

	b.log.Info("Batch prepared", "totalTransactions", totalTransactionsBeingCommitted)
	data, err := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Pack("submitBatch", txs, updatedRoots, uint8(commitments[0].BatchType), aggregatedSig)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return err
	}

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	stakeAmount := big.NewInt(0)
	stakeAmount.SetString("3200000000000000000", 10)

	// generate call msg
	callMsg := ethereum.CallMsg{
		To:    &rollupAddress,
		Data:  data,
		Value: stakeAmount,
	}

	auth, err := b.GenerateAuthObj(b.EthClient, callMsg)
	if err != nil {
		b.log.Error("Error creating auth object", "error", err)
		return err
	}

	// lastTxBroadcasted, err := DBInstance.GetLastTransaction()
	// if err != nil {
	// 	return err
	// }

	// if lastTxBroadcasted.Nonce+1 != auth.Nonce.Uint64() {
	// 	b.log.Info("Replacing nonce", "nonceEstimated", auth.Nonce.String(), "replacedBy", lastTxBroadcasted.Nonce+1)
	// 	auth.Nonce = big.NewInt(int64(lastTxBroadcasted.Nonce + 1))
	// }

	// latestBatch, err := DBInstance.GetLatestBatch()
	// if err != nil {
	// 	return err
	// }

	// newBatch := Batch{
	// 	BatchID:   latestBatch.BatchID + 1,
	// 	StateRoot: updatedRoot.String(),
	// 	Committer: config.OperatorAddress.String(),
	// 	Status:    BATCH_BROADCASTED,
	// }

	// b.log.Info("Broadcasting a new batch", "newBatch", newBatch)
	// err = DBInstance.AddNewBatch(newBatch)
	// if err != nil {
	// 	return err
	// }

	tx, err := b.RollupContract.SubmitBatch(auth, txs, updatedRoots, uint8(commitments[0].BatchType), aggregatedSig)
	if err != nil {
		b.log.Error("Error submitting batch", "err", err)
		return err
	}
	b.log.Info("Sent a new batch!", "TxHash", tx.Hash().String())

	// err = DBInstance.LogBatch(0, txs[0].Type, updatedRoot.String(), compressedTxs)
	// if err != nil {
	// 	return err
	// }

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
