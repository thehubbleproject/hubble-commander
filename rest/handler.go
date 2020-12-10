package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/BOPR/core"
	"github.com/gorilla/mux"
)

var (
	ErrInvalidTxType = errors.New("Invalid transaction type")
)

const (
	KeyID = "id"
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

type stateExporter struct {
	Balance   uint64 `json:"balance"`
	AccountID uint64 `json:"account_id"`
	StateID   uint64 `json:"state_id"`
	Token     uint64 `json:"token_id"`
	Nonce     uint64 `json:"nonce"`
}

func stateDecoderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params[KeyID]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to convert ID")
	}

	parameters, err := db.GetParams()
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to convert ID")
	}

	path, err := core.SolidityPathToNodePath(uint64(idInt), parameters.MaxDepth)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to convert ID")
	}

	state, err := db.GetStateByPath(path)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable toID")
	}
	ID, balance, nonce, token, err := bazooka.DecodeState(state.Data)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable toID")
	}
	var stateData stateExporter
	stateData.AccountID = ID.Uint64()
	stateData.Balance = balance.Uint64()
	stateData.Nonce = nonce.Uint64()
	stateData.StateID = uint64(idInt)
	stateData.Token = token.Uint64()

	output, err := json.Marshal(stateData)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
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
