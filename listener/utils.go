package listener

import (
	"bytes"
	"math/big"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/core"
	"github.com/BOPR/db"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// EventByID searches for all the loaded ABI's for matching events
func EventByID(abiObject *abi.ABI, sigdata []byte) *abi.Event {
	for _, event := range abiObject.Events {
		if bytes.Equal(event.ID.Bytes(), sigdata) {
			return &event
		}
	}
	return nil
}

// IsCatchingUp returns true/false according to the sync status of the node
func IsCatchingUp(b bazooka.Bazooka, db db.DB) (bool, error) {
	totalBatches, err := b.TotalBatches()
	if err != nil {
		return false, err
	}

	totalBatchedStored, err := db.GetBatchCount()
	if err != nil {
		return false, err
	}

	// if total batchse are greater than what we recorded we are still catching up
	if totalBatches > uint64(totalBatchedStored) {
		return true, err
	}

	return false, nil
}

// decompressTransfers decompresses transfer bytes to TX
func decompressTransfers(b bazooka.Bazooka, DBInstance db.DB, compressedTxs [][]byte) (txs []core.Tx, txsPerCommitment []int, err error) {
	var transactions []core.Tx
	for _, txSet := range compressedTxs {
		froms, tos, amounts, fees, errr := b.DecompressTransferTxs(txSet)
		if errr != nil {
			return
		}
		txsPerCommitment = append(txsPerCommitment, len(froms))
		for i := 0; i < len(froms); i++ {
			fromState, err := DBInstance.GetStateByIndex(froms[i].Uint64())
			if err != nil {
				return transactions, txsPerCommitment, err
			}

			_, _, nonce, _, err := b.DecodeState(fromState.Data)
			if err != nil {
				return transactions, txsPerCommitment, err
			}

			txData, err := b.EncodeTransferTx(froms[i].Int64(), tos[i].Int64(), fees[i].Int64(), nonce.Int64(), amounts[i].Int64(), int64(core.TX_TRANSFER_TYPE))
			if err != nil {
				return transactions, txsPerCommitment, err
			}
			newTx := core.NewTx(froms[i].Uint64(), tos[i].Uint64(), core.TX_TRANSFER_TYPE, nil, txData)
			transactions = append(transactions, newTx)
		}
	}
	return transactions, txsPerCommitment, nil
}

func decompressCreate2Transfers(b bazooka.Bazooka, DBInstance db.DB, compressedTxs [][]byte) (txs []core.Tx, txsPerCommitment []int, err error) {
	var transactions []core.Tx
	for _, txSet := range compressedTxs {
		froms, tos, toAccIDs, amounts, fees, errr := b.DecompressCreate2TransferTxs(txSet)
		if errr != nil {
			return
		}
		txsPerCommitment = append(txsPerCommitment, len(froms))
		for i := 0; i < len(froms); i++ {
			fromState, err := DBInstance.GetStateByIndex(froms[i].Uint64())
			if err != nil {
				return transactions, txsPerCommitment, err
			}
			_, _, nonce, _, err := b.DecodeState(fromState.Data)
			if err != nil {
				return transactions, txsPerCommitment, err
			}
			txData, err := b.EncodeCreate2TransferTx(froms[i].Int64(), tos[i].Int64(), toAccIDs[i].Int64(), fees[i].Int64(), nonce.Int64(), amounts[i].Int64(), int64(core.TX_CREATE_2_TRANSFER))
			if err != nil {
				return transactions, txsPerCommitment, err
			}

			newTx := core.NewTx(froms[i].Uint64(), tos[i].Uint64(), core.TX_CREATE_2_TRANSFER, nil, txData)
			transactions = append(transactions, newTx)
		}
	}
	return transactions, txsPerCommitment, nil
}

// decompressTransfers decompresses transfer bytes to TX
func decompressMassMigrations(b bazooka.Bazooka, DBInstance db.DB, compressedTxs [][]byte, meta [][4]*big.Int) (txs []core.Tx, txsPerCommitment []int, err error) {
	var transactions []core.Tx
	for _, txSet := range compressedTxs {
		froms, amounts, fees, errr := b.DecompressMassMigrationTxs(txSet)
		if errr != nil {
			return
		}
		for i := 0; i < len(froms); i++ {
			fromState, err := DBInstance.GetStateByIndex(froms[i].Uint64())
			if err != nil {
				return transactions, txsPerCommitment, err
			}
			_, _, nonce, _, err := b.DecodeState(fromState.Data)
			if err != nil {
				return transactions, txsPerCommitment, err
			}
			txData, err := b.EncodeMassMigrationTx(froms[i].Int64(), meta[i][0].Int64(), fees[i].Int64(), nonce.Int64(), amounts[i].Int64(), int64(core.TX_MASS_MIGRATIONS))
			if err != nil {
				return transactions, txsPerCommitment, err
			}
			newTx := core.NewTx(froms[i].Uint64(), 0, core.TX_TRANSFER_TYPE, nil, txData)
			transactions = append(transactions, newTx)
		}
	}
	return transactions, txsPerCommitment, nil
}
