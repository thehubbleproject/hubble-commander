package main

import (
	"fmt"

	"github.com/BOPR/core"
	"github.com/spf13/cobra"
)

//  viewState
func viewState() *cobra.Command {
	cmd := cobra.Command{
		Use:   "state-info",
		Short: "returns decoded state info",
		RunE: func(cmd *cobra.Command, args []string) error {
			stateID, err := cmd.Flags().GetUint64(FlagStateID)
			if err != nil {
				return err
			}
			db, err := core.NewDB()
			if err != nil {
				return err
			}
			defer db.Close()

			bazooka, err := core.NewPreLoadedBazooka()
			if err != nil {
				return err
			}

			state, err := db.GetStateByIndex(stateID)
			if err != nil {
				return err
			}

			fmt.Println("State", state, stateID)

			ID, bal, nonce, token, err := bazooka.DecodeState(state.Data)
			if err != nil {
				return err
			}

			fmt.Println("State info", "AccountID", ID, "Balance", bal, "Nonce", nonce, "Token", token)
			return nil
		},
	}

	cmd.Flags().Uint64(FlagStateID, 0, "--id=<state-id>")
	err := cmd.MarkFlagRequired(FlagStateID)
	if err != nil {
		panic(err)
	}
	return &cmd
}
