package core

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/BOPR/config"
	"github.com/BOPR/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

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
