package bazooka

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	tmLog "github.com/tendermint/tendermint/libs/log"

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
	DOMAIN                 = [32]byte{}
)

const (
	TXS_PARAM      = "txss"
	META_PARAM     = "meta"
	WITHDRAW_ROOTS = "withdrawRoots"
)

type (

	// Bazooka contains everything needed to interact with smart contracts
	Bazooka struct {
		log       tmLog.Logger
		EthClient *ethclient.Client

		RollupABI abi.ABI
		SC        Contracts
	}

	Contracts struct {
		RollupContract  *rollup.Rollup
		State           *state.State
		Transfer        *transfer.Transfer
		Create2Transfer *create2transfer.Create2transfer
		MassMigration   *massmigration.Massmigration
		AccountRegistry *accountregistry.Accountregistry
	}
)

// NewPreLoadedBazooka loads all contract and creates a ready to go client
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

	bazooka.log = log.Logger.With("module", "bazooka")
	return bazooka, nil
}

// GetEthBlock fetches the eth chain block for a block num
func (b *Bazooka) GetEthBlock(blockNum *big.Int) (header *ethTypes.Header, err error) {
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

func (b *Bazooka) ProcessCreate2TransferTx(balanceTreeRoot core.ByteArray, tx core.Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot core.ByteArray, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion(*b)
	if err != nil {
		return
	}
	toMP, err := toMerkleProof.ToABIVersion(*b)
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
	if err = core.ParseResult(result.Result); err != nil {
		return
	}

	b.log.Info("Processed transaction", "postTxRoot", core.ByteArray(result.NewRoot).String(), "resultCode", result.Result)

	// TOOD read result code and bubble up error messages
	return result.NewRoot, nil
}

func (b *Bazooka) ProcessTransferTx(balanceTreeRoot core.ByteArray, tx core.Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot core.ByteArray, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion(*b)
	if err != nil {
		return
	}
	toMP, err := toMerkleProof.ToABIVersion(*b)
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
	if err = core.ParseResult(result.Result); err != nil {
		return
	}

	b.log.Info("Processed transaction", "postTxRoot", core.ByteArray(result.NewRoot).String(), "resultCode", result.Result)

	// TOOD read result code and bubble up error messages
	return result.NewRoot, nil
}

func (b *Bazooka) ProcessMassMigrationTx(balanceTreeRoot core.ByteArray, tx core.Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot core.ByteArray, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := fromMerkleProof.ToABIVersion(*b)
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
	if err = core.ParseResult(result.Result); err != nil {
		return
	}

	b.log.Info("Processed transaction", "postTxRoot", core.ByteArray(result.NewRoot).String(), "resultCode", result.Result)

	// TOOD read result code and bubble up error messages
	return result.NewRoot, nil
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
