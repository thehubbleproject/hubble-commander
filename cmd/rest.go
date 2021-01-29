package main

import (
	"fmt"
	"net/http"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/rest"
	"github.com/gorilla/handlers"
	"github.com/spf13/cobra"
)

// startRestServerCmd starts the daemon
func startRestServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Starts hubble rest-server",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.ParseConfig()
			common.PanicIfError(err)
			r, err := rest.LoadRouters(cfg)

			headersOk := handlers.AllowedHeaders([]string{"*"})
			originsOk := handlers.AllowedOrigins([]string{"*"})
			methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

			common.PanicIfError(err)
			http.Handle("/", r)

			fmt.Println("Server started on port 3000 🎉")

			// TODO replace this with port from config
			err = http.ListenAndServe("0.0.0.0:3000", handlers.CORS(originsOk, headersOk, methodsOk)(r))
			common.PanicIfError(err)
		},
	}
}
