package rest

import "github.com/gorilla/mux"

// LoadRouters loads router
func LoadRouters() mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tx", TxReceiverHandler).Methods("POST")
	r.HandleFunc("/account", GetAccountHandler).Methods("GET")
	return *r
}
