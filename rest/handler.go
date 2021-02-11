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

func handleTx(r *http.Request) (resp ResponseTx, err error) {
	var tx TxReceiver
	err = json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		return
	}
	txMessageBytes, err := hex.DecodeString(tx.Message)
	if err != nil {
		return
	}
	from, _, nonceInTx, _, fee, err := decodeTx(txMessageBytes, tx.Type)
	if err != nil {
		return
	}
	fromState, err := dbI.GetStateByIndex(from)
	if err != nil {
		return
	}
	_, _, nonce, token, err := bazookaI.DecodeState(fromState.Data)
	if err != nil {
		return
	}
	if nonceInTx != nonce.Uint64()+1 {
		return
	}
	txSignatureBytes, err := hex.DecodeString(tx.Signature)
	if err != nil {
		return
	}

	// create a new pending transaction
	userTx, err := core.NewPendingTx(txMessageBytes, txSignatureBytes, from, nonceInTx, fee, token.Uint64(), tx.Type)
	if err != nil {
		return
	}

	// add the transaction to pool
	err = dbI.InsertTx(&userTx)
	if err != nil {
		return
	}

	resp.From = userTx.From
	resp.Data = hex.EncodeToString(userTx.Data)
	resp.Signature = hex.EncodeToString(userTx.Signature)
	resp.TxHash = userTx.TxHash
	resp.Status = userTx.Status
	resp.Type = userTx.Type
	return
}

// TxHandler handles user txs
func TxHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := handleTx(r)
	WriteRESTResp(w, resp, err)
}

type stateExporter struct {
	Balance   uint64 `json:"balance"`
	AccountID uint64 `json:"account_id"`
	StateID   uint64 `json:"state_id"`
	Token     uint64 `json:"token_id"`
	Nonce     uint64 `json:"nonce"`
}

func handleStateDecode(r *http.Request) (stateData stateExporter, err error) {
	params := mux.Vars(r)
	id := params[KeyID]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	state, err := dbI.GetStateByIndex(uint64(idInt))
	if err != nil {
		return
	}
	if state.Type != core.TYPE_TERMINAL {
		err = errors.New("Invalid state")
		return
	}

	ID, balance, nonce, token, err := bazookaI.DecodeState(state.Data)
	if err != nil {
		return
	}
	stateData.AccountID = ID.Uint64()
	stateData.Balance = balance.Uint64()
	stateData.Nonce = nonce.Uint64()
	stateData.StateID = uint64(idInt)
	stateData.Token = token.Uint64()
	return
}

func stateDecoderHandler(w http.ResponseWriter, r *http.Request) {
	stateData, err := handleStateDecode(r)
	WriteRESTResp(w, stateData, err)
}

type accountExporter struct {
	AccountID uint64 `json:"account_id"`
	Pubkey    string `json:"pubkey"`
}

func handleAccountDecode(r *http.Request) (accountData accountExporter, err error) {
	params := mux.Vars(r)
	id := params[KeyID]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	account, err := dbI.GetAccountLeafByID(uint64(idInt))
	if err != nil {
		return
	}

	if account.Type != core.TYPE_TERMINAL || len(account.PublicKey) == 0 {
		err = errors.New("Invalid account or account not present")
		return
	}

	accountData.AccountID = account.AccountID
	accountData.Pubkey = core.Pubkey(account.PublicKey).String()
	return
}

func accountDecoderHandler(w http.ResponseWriter, r *http.Request) {
	accountData, err := handleAccountDecode(r)
	WriteRESTResp(w, accountData, err)
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

func handlerTransfer(r *http.Request) (response SignBytesResponse, err error) {
	var transferTx TransferTx
	err = json.NewDecoder(r.Body).Decode(&transferTx)
	if err != nil {
		return
	}
	txData, err := bazookaI.EncodeTransferTx(int64(transferTx.From), int64(transferTx.To), int64(transferTx.Fee), int64(transferTx.Nonce), int64(transferTx.Amount), core.TX_TRANSFER_TYPE)
	if err != nil {
		return
	}
	tx := core.Tx{Data: txData}
	signBytes, err := bazookaI.TransferSignBytes(tx)
	if err != nil {
		return
	}
	response.Message = hex.EncodeToString(signBytes)
	response.Type = core.TX_TRANSFER_TYPE
	return
}

func transferTxHandler(w http.ResponseWriter, r *http.Request) {
	response, err := handlerTransfer(r)

	WriteRESTResp(w, response, err)
}

// MassMigrationsTx is the body
type MassMigrationsTx struct {
	From      uint64 `json:"from"`
	ToSpokeID uint64 `json:"to_spoke_id"`
	Nonce     uint64 `json:"nonce"`
	Amount    uint64 `json:"amount"`
	Fee       uint64 `json:"fee"`
}

func handleMassMigration(r *http.Request) (response SignBytesResponse, err error) {
	var mmTx MassMigrationsTx
	err = json.NewDecoder(r.Body).Decode(&mmTx)
	if err != nil {
		return
	}
	txData, err := bazookaI.EncodeMassMigrationTx(int64(mmTx.From), int64(mmTx.ToSpokeID), int64(mmTx.Fee), int64(mmTx.Nonce), int64(mmTx.Amount), core.TX_MASS_MIGRATIONS)
	if err != nil {
		return
	}
	tx := core.Tx{Data: txData}
	signBytes, err := bazookaI.MassMigrationSignBytes(tx)
	if err != nil {
		return
	}
	response.Message = hex.EncodeToString(signBytes)
	response.Type = core.TX_MASS_MIGRATIONS
	return
}

func massMigrationTxHandler(w http.ResponseWriter, r *http.Request) {
	response, err := handleMassMigration(r)

	WriteRESTResp(w, response, err)
}

// Create2transfer is the body
type Create2transfer struct {
	From     uint64 `json:"from"`
	ToPubkey []byte `json:"to_pubkey"`
	Nonce    uint64 `json:"nonce"`
	Amount   uint64 `json:"amount"`
	Fee      uint64 `json:"fee"`
}

func handleCreate2Transfer(r *http.Request) (response SignBytesResponse, err error) {
	var cTx Create2transfer
	err = json.NewDecoder(r.Body).Decode(&cTx)
	if err != nil {
		return
	}

	toPubkey, err := core.Pubkey(cTx.ToPubkey).ToSol()
	if err != nil {
		return
	}

	txData, err := bazookaI.EncodeCreate2TransferTxWithPub(int64(cTx.From), toPubkey, int64(cTx.Fee), int64(cTx.Nonce), int64(cTx.Amount), core.TX_CREATE_2_TRANSFER)
	if err != nil {
		return
	}
	tx := core.Tx{Data: txData}
	signBytes, err := bazookaI.Create2TransferSignBytesWithPub(tx)
	if err != nil {
		return
	}
	response.Message = hex.EncodeToString(signBytes)
	response.Type = core.TX_CREATE_2_TRANSFER
	return
}

func create2transferTxHandler(w http.ResponseWriter, r *http.Request) {
	response, err := handleCreate2Transfer(r)
	WriteRESTResp(w, response, err)
}

func handleTxStatus(r *http.Request) (tx *core.Tx, err error) {
	params := mux.Vars(r)
	txHash := params[KeyHash]
	if len(txHash) == 0 {
		err = errors.New("TxHash not present")
		return
	}

	tx, err = dbI.GetTxByHash(txHash)
	return

}

func txStatusHandler(w http.ResponseWriter, r *http.Request) {
	tx, err := handleTxStatus(r)
	WriteRESTResp(w, tx, err)
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

func handleUserState(r *http.Request) (response UserDetails, err error) {
	params := mux.Vars(r)
	pubkeyStr := params[KeyPubkey]
	pubkeybz, err := hex.DecodeString(pubkeyStr)
	if err != nil {
		return
	}

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
		_, balance, nonce, token, _err := bazookaI.DecodeState(state.Data)
		if _err != nil {
			return response, _err
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
	return
}

func userStateHandler(w http.ResponseWriter, r *http.Request) {
	response, err := handleUserState(r)
	WriteRESTResp(w, response, err)
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

func handleEstimateNonce(r *http.Request) (resp estimateNonceResp, err error) {
	params := mux.Vars(r)
	id := params[KeyID]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	pendingNonce, err := dbI.GetPendingNonce(uint64(idInt))
	if err != nil {
		return
	}

	resp.StateID = uint64(idInt)
	resp.Nonce = pendingNonce
	return
}

func estimateNonceHandler(w http.ResponseWriter, r *http.Request) {
	response, err := handleEstimateNonce(r)
	WriteRESTResp(w, response, err)
}
