package listener

import (
	"bytes"

	"github.com/BOPR/bazooka"
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

func concatTxs(txss [][]byte) (txs []byte) {
	for _, tx := range txss {
		txs = append(txs, tx[:]...)
	}
	return txs
}
