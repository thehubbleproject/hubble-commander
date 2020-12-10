package rest

import (
	"github.com/BOPR/core"
	"github.com/gorilla/mux"
)

var db core.DB
var bazooka core.Bazooka

// LoadRouters loads router
func LoadRouters() (r *mux.Router, err error) {
	tempDB, err := core.NewDB()
	if err != nil {
		return
	}
	db = tempDB

	bz, err := core.NewPreLoadedBazooka()
	if err != nil {
		return
	}
	bazooka = bz

	r = mux.NewRouter()
	r.HandleFunc("/tx", TxHandler).Methods("POST")
	r.HandleFunc("/account", GetAccountHandler).Methods("GET")
	r.HandleFunc("/state/{id}", stateDecoderHandler).Methods("GET")
	return r, nil
}
