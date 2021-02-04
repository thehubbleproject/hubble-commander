package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/BOPR/config"
	"github.com/BOPR/simulator"
	"github.com/spf13/cobra"
)

// startSimulator starts the daemon
func startSimulator() *cobra.Command {
	return &cobra.Command{
		Use:   "simulate",
		Short: "Starts hubble simulator",
		RunE: func(cmd *cobra.Command, args []string) error {
			jsonFile, err := os.Open("users.json")
			if err != nil {
				return err
			}
			// defer the closing of our jsonFile so that we can parse it later on
			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			// we initialize our Users array
			var users simulator.UserList

			// we unmarshal our byteArray which contains our
			// jsonFile's content into 'users' which we defined above
			err = json.Unmarshal(byteValue, &users)
			if err != nil {
				return err
			}

			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}

			sim := simulator.NewSimulator(cfg, users)

			// go routine to catch signal
			catchSignal := make(chan os.Signal, 1)
			signal.Notify(catchSignal, os.Interrupt)

			go func() {
				// sig is a ^C, handle it
				for range catchSignal {
					if err := sim.Stop(); err != nil {
						log.Fatalln("Unable to stop simulator", "error", err)
					}
					// exit
					os.Exit(1)
				}
			}()

			if err := sim.Start(); err != nil {
				return err
			}

			runtime.Goexit()
			return nil
		},
	}
}
