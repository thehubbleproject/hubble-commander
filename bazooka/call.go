package bazooka

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethCmn "github.com/ethereum/go-ethereum/common"
)

// CompressTxs compresses all transactions
func CompressTxs(b *Bazooka, txs []core.Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	var data [][]byte
	for _, tx := range txs {
		data = append(data, tx.Data)
	}
	switch txType := txs[0].Type; txType {
	case core.TX_TRANSFER_TYPE:
		return b.CompressTransferTxs(opts, data)
	case core.TX_CREATE_2_TRANSFER:
		return b.CompressCreate2TransferTxs(opts, data)
	case core.TX_MASS_MIGRATIONS:
		return b.CompressMassMigrationTxs(opts, data)
	default:
		fmt.Println("TxType didnt match any options", txs[0].Type)
		return []byte(""), errors.New("Did not match any options")
	}
}

// ApplyTx applies the transaction and returns the udpates
func ApplyTx(b Bazooka, sender, receiver []byte, tx core.Tx) (updatedSender, updatedReceiver []byte, err error) {
	switch txType := tx.Type; txType {
	case core.TX_TRANSFER_TYPE:
		return b.ApplyTransferTx(sender, receiver, tx)
	case core.TX_CREATE_2_TRANSFER:
		return b.ApplyCreate2TransferTx(sender, receiver, tx)
	case core.TX_MASS_MIGRATIONS:
		return b.ApplyMassMigrationTx(sender, receiver, tx)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return updatedSender, updatedReceiver, errors.New("Didn't match any options")
	}
}

// ProcessTx calls the ProcessTx function on the contract to verify the tx
// returns the updated accounts and the new balance root
func ProcessTx(b Bazooka, balanceTreeRoot core.ByteArray, tx core.Tx, fromMerkleProof, toMerkleProof StateMerkleProof) (newBalanceRoot core.ByteArray, err error) {
	switch txType := tx.Type; txType {
	case core.TX_TRANSFER_TYPE:
		return b.ProcessTransferTx(balanceTreeRoot, tx, fromMerkleProof, toMerkleProof)
	case core.TX_CREATE_2_TRANSFER:
		return b.ProcessCreate2TransferTx(balanceTreeRoot, tx, fromMerkleProof, toMerkleProof)
	case core.TX_MASS_MIGRATIONS:
		return b.ProcessMassMigrationTx(balanceTreeRoot, tx, fromMerkleProof, toMerkleProof)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return newBalanceRoot, errors.New("Did not match any options")
	}
}

func (b *Bazooka) CompressTransferTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
	return b.SC.Transfer.Compress(&opts, data)
}

func (b *Bazooka) CompressCreate2TransferTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
	return b.SC.Create2Transfer.Compress(&opts, data)
}

func (b *Bazooka) CompressMassMigrationTxs(opts bind.CallOpts, data [][]byte) ([]byte, error) {
	return b.SC.MassMigration.Compress(&opts, data)
}

func (b *Bazooka) TransferSignBytes(tx core.Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.SC.Transfer.SignBytes(&opts, tx.Data)
}

func (b *Bazooka) Create2TransferSignBytesWithPub(tx core.Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	return b.SC.Create2Transfer.SignBytes(&opts, tx.Data)
}

func (b *Bazooka) MassMigrationSignBytes(tx core.Tx) ([]byte, error) {
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

func (b *Bazooka) ApplyTransferTx(sender, receiver []byte, tx core.Tx) (updatedSender, updatedReceiver []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	updates, err := b.SC.Transfer.ValidateAndApply(&opts, sender, receiver, tx.Data)
	if err != nil {
		return
	}

	if err = core.ParseResult(updates.Result); err != nil {
		return
	}

	return updates.NewSender, updates.NewReceiver, nil
}

func (b *Bazooka) ApplyCreate2TransferTx(sender, receiver []byte, tx core.Tx) (updatedSender, updatedReceiver []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	updates, err := b.SC.Create2Transfer.ValidateAndApply(&opts, sender, tx.Data)
	if err != nil {
		return
	}

	if err = core.ParseResult(updates.Result); err != nil {
		return
	}

	return updates.NewSender, updates.NewReceiver, nil
}

func (b *Bazooka) ApplyMassMigrationTx(sender, receiver []byte, tx core.Tx) (updatedSender, updatedReceiver []byte, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	updates, err := b.SC.MassMigration.ValidateAndApply(&opts, sender, tx.Data)
	if err != nil {
		return
	}

	if err = core.ParseResult(updates.Result); err != nil {
		return
	}

	return updates.NewSender, receiver, nil
}

func GetSignBytes(b Bazooka, tx *core.Tx) (signBytes []byte, err error) {
	switch txType := tx.Type; txType {
	case core.TX_TRANSFER_TYPE:
		return b.TransferSignBytes(*tx)
	case core.TX_CREATE_2_TRANSFER:
		return b.Create2TransferSignBytesWithPub(*tx)
	case core.TX_MASS_MIGRATIONS:
		return b.MassMigrationSignBytes(*tx)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return []byte(""), errors.New("Did not match any options")
	}
}

// GetTokenAddress fetches the token address from token registry
func (b *Bazooka) GetTokenAddress(tokenID uint64) (tokenAddr common.Address, err error) {
	opts := bind.CallOpts{From: b.operator}
	tokenAddr, err = b.SC.TokenRegistry.SafeGetAddress(&opts, big.NewInt(int64(tokenID)))
	if err != nil {
		return tokenAddr, err
	}
	return tokenAddr, nil
}

func (b *Bazooka) FetchFromAndToStateIDs(tx core.Tx) (from, to uint64, err error) {
	switch txType := tx.Type; txType {
	case core.TX_TRANSFER_TYPE:
		from, to, _, _, _, _, err := b.DecodeTransferTx(tx.Data)
		return from.Uint64(), to.Uint64(), err
	case core.TX_CREATE_2_TRANSFER:
		from, to, _, _, _, _, _, err := b.DecodeCreate2Transfer(tx.Data)
		return from.Uint64(), to.Uint64(), err
	case core.TX_MASS_MIGRATIONS:
		from, _, _, _, _, _, err := b.DecodeMassMigrationTx(tx.Data)
		return from.Uint64(), 0, err
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return 0, 0, errors.New("Did not match any options")
	}
}

//
// burn auction
//

// GetCurrentProposer fetches the current proposer from on-chain contract
func (b *Bazooka) GetCurrentProposer() (common.Address, error) {
	opts := bind.CallOpts{From: b.operator}
	proposer, err := b.SC.BurnAuction.GetProposer(&opts)
	if err != nil {
		return proposer, err
	}

	return proposer, nil
}

// GetCurrentSlot fetches the current slot
func (b *Bazooka) GetCurrentSlot() (currentSlot uint32, err error) {
	opts := bind.CallOpts{From: b.operator}
	slot, err := b.SC.BurnAuction.CurrentSlot(&opts)
	if err != nil {
		return 0, err
	}
	return slot, nil
}

// GetBidableSlot fetches the current slot
func (b *Bazooka) GetBidableSlot() (bidableSlot uint32, err error) {
	currSlot, err := b.GetCurrentSlot()
	if err != nil {
		return 0, err
	}
	return currSlot + 2, nil
}

// GetCurrentBidForSlot fetches the current bid for the slot
func (b *Bazooka) GetCurrentBidForSlot(slot uint32) (proposer ethCmn.Address, amount uint64, isInit bool, err error) {
	opts := bind.CallOpts{From: b.operator}
	resp, err := b.SC.BurnAuction.Auction(&opts, slot)
	if err != nil {
		return
	}
	return resp.Coordinator, resp.Amount.Uint64(), resp.Initialized, nil
}

// GetDeposit fetches the deposit amount for the proposer
func (b *Bazooka) GetDeposit(proposer ethCmn.Address) (uint64, error) {
	opts := bind.CallOpts{From: b.operator}
	amount, err := b.SC.BurnAuction.Deposits(&opts, proposer)
	if err != nil {
		return 0, err
	}
	return amount.Uint64(), nil
}
