package main

import (
	"fmt"
	"net/http"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/rest"
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
			common.PanicIfError(err)
			http.Handle("/", r)

			fmt.Println("Server started on port 3000 🎉")

			// TODO replace this with port from config
			err = http.ListenAndServe(":3000", r)
			common.PanicIfError(err)
		},
	}
}
