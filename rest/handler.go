package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/BOPR/core"
	"github.com/gorilla/mux"
)

var (
	ErrInvalidTxType = errors.New("Invalid transaction type")
)

type TxReceiver struct {
	Message   []byte `json:"message"`
	Signature []byte `json:"sig"`
	Type      uint64 `json:"type"`
}

// TxReceiverHandler handles user txs
func TxHandler(w http.ResponseWriter, r *http.Request) {
	// receive the payload and read
	var tx TxReceiver
	if !ReadRESTReq(w, r, &tx) {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}

	to, from, err := decodeTx(tx.Message, tx.Type)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}

	// create a new pending transaction
	userTx, err := core.NewPendingTx(from, to, core.TX_TRANSFER_TYPE, tx.Signature, tx.Message)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	// add the transaction to pool
	err = core.DBInstance.InsertTx(&userTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}

	output, err := json.Marshal(userTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

func stateDecoderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("params", params)
}

func decodeTx(tx []byte, txType uint64) (to, from uint64, err error) {
	switch txType {
	case core.TX_TRANSFER_TYPE:
		fromInt, toInt, _, _, _, _, err := bazooka.DecodeTransferTx(tx)
		if err != nil {
			return to, from, err
		}
		return toInt.Uint64(), fromInt.Uint64(), nil
	case core.TX_CREATE_2_TRANSFER:
		fromInt, _, _, _, _, _, err := bazooka.DecodeCreate2TransferWithPub(tx)
		if err != nil {
			return to, from, err
		}
		return fromInt.Uint64(), 0, nil
	case core.TX_MASS_MIGRATIONS:
		fromInt, _, _, _, _, _, err := bazooka.DecodeMassMigrationTx(tx)
		if err != nil {
			return to, from, err
		}
		return fromInt.Uint64(), 0, nil
	default:
		return 0, 0, ErrInvalidTxType
	}
}
