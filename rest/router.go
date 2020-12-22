package rest

import (
	"github.com/BOPR/bazooka"
	"github.com/BOPR/db"
	"github.com/gorilla/mux"
)

var dbI db.DB
var bazookaI bazooka.Bazooka

// LoadRouters loads router
func LoadRouters() (r *mux.Router, err error) {
	tempDB, err := db.NewDB()
	if err != nil {
		return
	}
	dbI = tempDB

	bz, err := bazooka.NewPreLoadedBazooka()
	if err != nil {
		return
	}
	bazookaI = bz

	r = mux.NewRouter()
	r.HandleFunc("/tx", TxHandler).Methods("POST")
	r.HandleFunc("/account", GetAccountHandler).Methods("GET")
	r.HandleFunc("/state/{id}", stateDecoderHandler).Methods("GET")
	return r, nil
}
