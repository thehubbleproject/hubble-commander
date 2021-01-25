package rest

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BOPR/core"
	"github.com/gorilla/mux"
)

var (
	ErrInvalidTxType = errors.New("Invalid transaction type")
)

const (
	KeyID     = "id"
	KeyHash   = "hash"
	KeyPubkey = "pubkey"
)

type TxReceiver struct {
	Type      uint64 `json:"type"`
	Message   []byte `json:"message"`
	Signature []byte `json:"sig"`
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
	err = dbI.InsertTx(&userTx)
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
		return
	}

	state, err := dbI.GetStateByIndex(uint64(idInt))
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to get state by path")
		return
	}
	if state.Type != core.TYPE_TERMINAL {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid state")
		return
	}

	ID, balance, nonce, token, err := bazookaI.DecodeState(state.Data)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to decode state")
		return
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

type accountExporter struct {
	AccountID uint64 `json:"account_id"`
	Pubkey    string `json:"pubkey"`
}

func accountDecoderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params[KeyID]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to convert ID")
		return
	}

	account, err := dbI.GetAccountLeafByID(uint64(idInt))
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to fetch account by ID")
		return
	}

	if account.Type != core.TYPE_TERMINAL || len(account.PublicKey) == 0 {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid account or account not present")
		return
	}

	var accountData accountExporter
	accountData.AccountID = account.AccountID
	accountData.Pubkey = core.Pubkey(account.PublicKey).String()

	output, err := json.Marshal(accountData)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
		return
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

type SignBytesResponse struct {
	Type    uint64 `json:"tx_type"`
	Message []byte `json:"message"`
}
type TransferTx struct {
	From   uint64 `json:"from"`
	To     uint64 `json:"to"`
	Nonce  uint64 `json:"nonce"`
	Amount uint64 `json:"amount"`
	Fee    uint64 `json:"fee"`
}

func transferTx(w http.ResponseWriter, r *http.Request) {
	var transferTx TransferTx
	err := json.NewDecoder(r.Body).Decode(&transferTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to json decode transfer")
		return
	}
	txData, err := bazookaI.EncodeTransferTx(int64(transferTx.From), int64(transferTx.To), int64(transferTx.Fee), int64(transferTx.Nonce), int64(transferTx.Amount), core.TX_TRANSFER_TYPE)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to json decode transfer")
		return
	}
	tx := core.Tx{Data: txData}
	signBytes, err := bazookaI.TransferSignBytes(tx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to json decode transfer")
		return
	}
	var response SignBytesResponse
	response.Message = signBytes
	response.Type = core.TX_TRANSFER_TYPE

	output, err := json.Marshal(response)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to json decode transfer")
		return
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// MassMigrationsTx is the body
type MassMigrationsTx struct {
	From      uint64 `json:"from"`
	ToSpokeID uint64 `json:"to_spoke_id"`
	Nonce     uint64 `json:"nonce"`
	Amount    uint64 `json:"amount"`
	Fee       uint64 `json:"fee"`
}

func massMigrationTx(w http.ResponseWriter, r *http.Request) {
	var mmTx MassMigrationsTx
	err := json.NewDecoder(r.Body).Decode(&mmTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to json decode mass migrations")
		return
	}
	txData, err := bazookaI.EncodeMassMigrationTx(int64(mmTx.From), int64(mmTx.ToSpokeID), int64(mmTx.Fee), int64(mmTx.Nonce), int64(mmTx.Amount), core.TX_MASS_MIGRATIONS)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to encode mass migration transaction")
		return
	}
	tx := core.Tx{Data: txData}
	signBytes, err := bazookaI.MassMigrationSignBytes(tx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to generate sign bytes")
		return
	}
	var response SignBytesResponse
	response.Message = signBytes
	response.Type = core.TX_MASS_MIGRATIONS

	output, err := json.Marshal(response)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall")
		return
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// Create2transfer is the body
type Create2transfer struct {
	From     uint64 `json:"from"`
	ToPubkey []byte `json:"to_pubkey"`
	Nonce    uint64 `json:"nonce"`
	Amount   uint64 `json:"amount"`
	Fee      uint64 `json:"fee"`
}

func create2transferTx(w http.ResponseWriter, r *http.Request) {
	var cTx Create2transfer
	err := json.NewDecoder(r.Body).Decode(&cTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to json decode mass migrations")
		return
	}

	toPubkey, err := core.Pubkey(cTx.ToPubkey).ToSol()
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to json decode mass migrations")
		return
	}

	txData, err := bazookaI.EncodeCreate2TransferTxWithPub(int64(cTx.From), toPubkey, int64(cTx.Fee), int64(cTx.Nonce), int64(cTx.Amount), core.TX_CREATE_2_TRANSFER)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to encode mass migration transaction")
		return
	}
	tx := core.Tx{Data: txData}
	signBytes, err := bazookaI.Create2TransferSignBytesWithPub(tx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to generate sign bytes")
		return
	}
	var response SignBytesResponse
	response.Message = signBytes
	response.Type = core.TX_CREATE_2_TRANSFER
	output, err := json.Marshal(response)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall")
		return
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

func txStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txHash := params[KeyHash]
	if len(txHash) == 0 {
		WriteErrorResponse(w, http.StatusBadRequest, "TxHash not present")
		return
	}

	tx, err := dbI.GetTxByHash(txHash)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to find transaction by hash")
		return
	}

	output, err := json.Marshal(tx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
		return
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// UserDetails is the body
type UserDetails struct {
	States    []UserDetailsState `json:"states"`
	AccountID uint64             `json:"account_id"`
	Pubkey    string             `json:"pubkey"`
}

type UserDetailsState struct {
	Balance uint64 `json:"balance"`
	StateID uint64 `json:"state_id"`
	Token   uint64 `json:"token_id"`
	Nonce   uint64 `json:"nonce"`
}

func userStateHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pubkeyStr := params[KeyPubkey]
	if len(pubkeyStr) == 0 {
		WriteErrorResponse(w, http.StatusBadRequest, "Pubkey not present")
		return
	}
	pubkeybz, err := hex.DecodeString(pubkeyStr)
	if err != nil {
		return
	}

	var response UserDetails

	acc, err := dbI.GetAccountByPubkey(pubkeybz)
	if err != nil {
		return
	}
	response.AccountID = acc.AccountID
	response.Pubkey = hex.EncodeToString(acc.PublicKey)

	states, err := dbI.GetStateByAccID(acc.AccountID)
	if err != nil {
		return
	}

	var userStates []UserDetailsState
	for _, state := range states {
		if state.Type != core.TYPE_TERMINAL {
			continue
		}

		_, balance, nonce, token, err := bazookaI.DecodeState(state.Data)
		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, "Unable to decode state")
			return
		}
		var stateData UserDetailsState
		stateData.Balance = balance.Uint64()
		stateData.Nonce = nonce.Uint64()
		stateID, err := core.StringToUint(state.Path)
		if err != nil {
			fmt.Println("error converting path to ID")
			continue
		}
		stateData.StateID = stateID
		stateData.Token = token.Uint64()

		userStates = append(userStates, stateData)
	}
	response.States = userStates
	output, err := json.Marshal(response)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
		return
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

func decodeTx(tx []byte, txType uint64) (to, from uint64, err error) {
	switch txType {
	case core.TX_TRANSFER_TYPE:
		fromInt, toInt, _, _, _, _, err := bazookaI.DecodeTransferTx(tx)
		if err != nil {
			return to, from, err
		}
		return toInt.Uint64(), fromInt.Uint64(), nil
	case core.TX_CREATE_2_TRANSFER:
		fromInt, _, _, _, _, _, err := bazookaI.DecodeCreate2TransferWithPub(tx)
		if err != nil {
			return to, from, err
		}
		return fromInt.Uint64(), 0, nil
	case core.TX_MASS_MIGRATIONS:
		fromInt, _, _, _, _, _, err := bazookaI.DecodeMassMigrationTx(tx)
		if err != nil {
			return to, from, err
		}
		return fromInt.Uint64(), 0, nil
	default:
		return 0, 0, ErrInvalidTxType
	}
}
