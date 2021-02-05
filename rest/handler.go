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
	Message   string `json:"message"`
	Signature string `json:"sig"`
}

type ResponseTx struct {
	From      uint64 `json:"from"`
	Data      string `json:"data"`
	Signature string `json:"sig" gorm:"null"`
	TxHash    string `json:"hash" gorm:"not null"`
	Status    uint64 `json:"status"`
	Type      uint64 `json:"type" gorm:"not null"`
}

func coreTxToResponseTx(_tx core.Tx) ResponseTx {
	var resp ResponseTx
	resp.From = _tx.From
	resp.Data = hex.EncodeToString(_tx.Data)
	resp.Signature = hex.EncodeToString(_tx.Signature)
	resp.TxHash = _tx.TxHash
	resp.Status = _tx.Status
	resp.Type = _tx.Type
	return resp
}

// TxReceiverHandler handles user txs
func TxHandler(w http.ResponseWriter, r *http.Request) {
	// receive the payload and read
	var tx TxReceiver
	if !ReadRESTReq(w, r, &tx) {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
		return
	}
	txMessageBytes, err := hex.DecodeString(tx.Message)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	from, _, nonceInTx, _, fee, err := decodeTx(txMessageBytes, tx.Type)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	fromState, err := dbI.GetStateByIndex(from)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	_, _, _, token, err := bazookaI.DecodeState(fromState.Data)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	// if nonceInTx != nonce.Uint64()+1 {
	// 	WriteErrorResponse(w, http.StatusBadRequest, "Nonce invalid")
	// 	return
	// }
	txSignatureBytes := []byte(tx.Signature)
	fmt.Println("signature length", len(txSignatureBytes))

	// create a new pending transaction
	userTx, err := core.NewPendingTx(txMessageBytes, txSignatureBytes, from, nonceInTx, fee, token.Uint64(), tx.Type)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// add the transaction to pool
	err = dbI.InsertTx(&userTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	output, err := json.Marshal(coreTxToResponseTx(userTx))
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
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
	Message string `json:"message"`
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
	response.Message = hex.EncodeToString(signBytes)
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
	response.Message = hex.EncodeToString(signBytes)
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
	response.Message = hex.EncodeToString(signBytes)
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
	pubkeybz, err := hex.DecodeString(pubkeyStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Could not decode string")
		return
	}

	var response UserDetails
	acc, err := dbI.GetAccountByPubkey(pubkeybz)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Could not get account by pubkey")
		return
	}

	response.AccountID = acc.AccountID
	response.Pubkey = hex.EncodeToString(acc.PublicKey)

	states, err := dbI.GetStateByAccID(acc.AccountID)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Could not get state by accID")
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

func decodeTx(tx []byte, txType uint64) (from, to, nonce, amount, fee uint64, err error) {
	switch txType {
	case core.TX_TRANSFER_TYPE:
		fromInt, toInt, nonceInt, _, amountInt, feeInt, err := bazookaI.DecodeTransferTx(tx)
		if err != nil {
			return from, to, nonce, amount, fee, err
		}
		return fromInt.Uint64(), toInt.Uint64(), nonceInt.Uint64(), amountInt.Uint64(), feeInt.Uint64(), nil
	case core.TX_CREATE_2_TRANSFER:
		fromInt, _, nonceInt, _, amountInt, feeInt, err := bazookaI.DecodeCreate2TransferWithPub(tx)
		if err != nil {
			return from, to, nonce, amount, fee, err
		}
		return fromInt.Uint64(), 0, nonceInt.Uint64(), amountInt.Uint64(), feeInt.Uint64(), nil
	case core.TX_MASS_MIGRATIONS:
		fromInt, _, nonceInt, _, amountInt, feeInt, err := bazookaI.DecodeMassMigrationTx(tx)
		if err != nil {
			return from, to, nonce, amount, fee, err
		}
		return fromInt.Uint64(), 0, nonceInt.Uint64(), amountInt.Uint64(), feeInt.Uint64(), nil
	default:
		return 0, 0, 0, 0, 0, ErrInvalidTxType
	}
}

type estimateNonceResp struct {
	StateID uint64
	Nonce   uint64
}

func estimateNonceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params[KeyID]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to convert ID")
		return
	}

	pendingNonce, err := dbI.GetPendingNonce(uint64(idInt))
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to estimate nonce")
		return
	}

	var resp estimateNonceResp
	resp.StateID = uint64(idInt)
	resp.Nonce = pendingNonce

	output, err := json.Marshal(resp)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
		return
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}
