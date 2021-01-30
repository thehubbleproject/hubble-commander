package rest

import (
	"fmt"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/db"
	"github.com/common-nighthawk/go-figure"
	"github.com/gorilla/mux"
)

var dbI db.DB
var bazookaI bazooka.Bazooka

// LoadRouters loads router
func LoadRouters(cfg config.Configuration) (r *mux.Router, err error) {
	myFigure := figure.NewColorFigure("Hubble", "", "red", true)
	myFigure.Print()
	tempDB, err := db.NewDB(cfg)
	if err != nil {
		return
	}
	dbI = tempDB

	bz, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		return
	}
	bazookaI = bz

	r = mux.NewRouter()
	r.HandleFunc("/user/state/{id}", stateDecoderHandler).Methods("GET")
	r.HandleFunc("/user/account/{id}", accountDecoderHandler).Methods("GET")
	r.HandleFunc("/tx/{hash}", txStatusHandler).Methods("GET")

	r.HandleFunc("/tx", TxHandler).Methods("POST")
	r.HandleFunc("/user/{pubkey}", userStateHandler).Methods("GET")

	r.HandleFunc("/transfer", transferTx).Methods("POST")
	r.HandleFunc("/massmigration", massMigrationTx).Methods("POST")
	r.HandleFunc("/create2transfer", create2transferTx).Methods("POST")

	fmt.Println("Here are the available routes...")

	var routes []string
	err = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		routes = append(routes, t)
		return nil
	})
	if err != nil {
		return
	}
	return r, nil
}
