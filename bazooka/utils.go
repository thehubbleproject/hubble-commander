package bazooka

import (
	"context"
	"errors"
	"math/big"

	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var (
	ErrNoEnoughConfirmations = errors.New("Not enough confirmations")
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
	data, err := b.getTxDataByHash(txHash)
	if err != nil {
		return nil, err
	}
	switch batchType {
	case core.TX_GENESIS:
		return calldata, ErrNoTxs
	case core.TX_DEPOSIT:
		return calldata, ErrNoTxs
	case core.TX_TRANSFER_TYPE:
		return b.UnpackBatchCalldata(SubmitTransferMethod, data)
	case core.TX_CREATE_2_TRANSFER:
		return b.UnpackBatchCalldata(SubmitCreate2TransferMethod, data)
	case core.TX_MASS_MIGRATIONS:
		return b.UnpackBatchCalldata(SubmitMassMigrationMethod, data)
	default:
		return nil, errors.New("Unable to match batch type")
	}
}

func (b *Bazooka) SignAndBroadcastBatch(client *ethclient.Client, toAddr ethCmn.Address, value *big.Int, data Calldata) (tx *types.Transaction, err error) {
	inputData, err := data.Pack(*b)
	if err != nil {
		return nil, err
	}
	return b.SignAndBroadcast(client, toAddr, value, inputData)
}

func (b *Bazooka) SignAndBroadcast(client *ethclient.Client, toAddr ethCmn.Address, value *big.Int, data []byte) (tx *types.Transaction, err error) {
	opts, err := b.generateAuthObj(b.EthClient, toAddr, value, data)
	if err != nil {
		b.log.Error("Estimate gas failed, tx reverting", "error", err)
		return
	}
	tx = types.NewTransaction(opts.Nonce.Uint64(), toAddr, opts.Value, opts.GasLimit, opts.GasPrice, data)
	sigTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, tx)
	if err != nil {
		b.log.Error("Error creating signer", "error", err)
		return
	}
	err = b.EthClient.SendTransaction(context.Background(), sigTx)
	if err != nil {
		b.log.Error("error unable to send transaction", err)
		return
	}
	return tx, nil
}

func (b *Bazooka) generateAuthObj(client *ethclient.Client, toAddr ethCmn.Address, value *big.Int, data []byte) (auth *bind.TransactOpts, err error) {
	// from address
	fromAddress := config.OperatorAddress
	callMsg := ethereum.CallMsg{
		From:  fromAddress,
		To:    &toAddr,
		Data:  data,
		Value: value,
	}

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
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return
	}
	// create auth
	auth = bind.NewKeyedTransactor(config.OperatorKey)
	auth.GasPrice = gasprice
	auth.Value = value
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = gasLimit
	return
}

// GetConfirmedTxReceipt returns confirmed tx receipt
func (b *Bazooka) GetConfirmedTxReceipt(tx common.Hash, requiredConfirmations uint64) (txReceipt *types.Receipt, err error) {
	// get main tx receipt
	receipt, err := b.GetReceipt(tx)
	if err != nil {
		return nil, err
	}

	// get main chain block
	latestBlk, err := b.GetEthBlock(nil)
	if err != nil {
		return nil, err
	}

	diff := latestBlk.Number.Uint64() - receipt.BlockNumber.Uint64()
	if diff < requiredConfirmations {
		return nil, ErrNoEnoughConfirmations
	}

	return receipt, nil
}

// GetReceipt returns main tx receipt
func (b *Bazooka) GetReceipt(txHash common.Hash) (*types.Receipt, error) {
	return b.getTxReceipt(b.EthClient, txHash)
}

func (b *Bazooka) getTxReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	return client.TransactionReceipt(context.Background(), txHash)
}

func etherToWei(val *big.Int) *big.Int {
	return new(big.Int).Mul(val, big.NewInt(params.Ether))
}

func weiToEther(val *big.Int) *big.Int {
	return new(big.Int).Div(val, big.NewInt(params.Ether))
}
