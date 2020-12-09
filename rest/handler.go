package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BOPR/core"
)

type TxReceiver struct {
	Message   []byte `json:"message"`
	Signature []byte `json:"sig"`
	Type      uint64 `json:"type"`
}

// TxReceiverHandler handles user txs
func TxReceiverHandler(w http.ResponseWriter, r *http.Request) {
	// receive the payload and read
	var tx TxReceiver
	if !ReadRESTReq(w, r, &tx) {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}

	to, from, err := getToFromFromBytes(tx.Message, tx.Type)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Error decoding message")
	}

	// create a new pending transaction
	userTx, err := core.NewPendingTx(from, to, core.TX_TRANSFER_TYPE, tx.Signature, tx.Message)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	// add the transaction to pool
	err = db.InsertTx(&userTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}

	output, err := json.Marshal(userTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Unable to write output:%v", err))
	}
}

func getToFromFromBytes(tx []byte, txType uint64) (to, from uint64, err error) {
	switch txType {
	case core.TX_TRANSFER_TYPE:
		fromInt, toInt, _, _, _, _, err := bazooka.DecodeTransferTx(tx)
		if err != nil {
			return 0, 0, err
		}
		return toInt.Uint64(), fromInt.Uint64(), nil
	case core.TX_CREATE_2_TRANSFER:
		fromInt, _, _, _, _, _, err := bazooka.DecodeCreate2TransferWithPub(tx)
		if err != nil {
			return 0, 0, err
		}
		return 0, fromInt.Uint64(), nil
	case core.TX_MASS_MIGRATIONS:
		fromInt, _, _, _, _, _, err := bazooka.DecodeMassMigrationTx(tx)
		if err != nil {
			return 0, 0, err
		}
		return 0, fromInt.Uint64(), nil
	}

	return to, from, nil
}
