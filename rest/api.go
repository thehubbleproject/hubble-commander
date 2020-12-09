package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BOPR/core"
)

type (
	TransferTx struct {
		From      uint64 `json:"from"`
		To        uint64 `json:"to"`
		Amount    uint64 `json:"amount"`
		Nonce     uint64 `json:"nonce"`
		TokenType uint64 `json:"token"`
		TxType    uint64 `json:"txType"`
	}
)

// GetAccountHandler fetches the user account data like balance, token type and nonce
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	IDstr := vars.Get("ID")
	ID, err := strconv.ParseUint(IDstr, 0, 64)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid ID")
	}
	fmt.Println(ID)
	var account core.UserState
	// account, err := core.DBInstance.GetAccount(ID)
	// if err != nil {
	// 	WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Account with ID %v not found", ID))
	// }
	output, err := json.Marshal(account)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}
