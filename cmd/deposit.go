package main

import (
	"encoding/hex"
	"fmt"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	db "github.com/BOPR/db"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// deposit generated init command to initialise the config file
func deposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit",
		Short: "Registers accounts and makes a deposit",
		RunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()
			// get the pubkey
			pubKey, err := flags.GetString(FlagPubKey)
			if err != nil {
				return err
			}

			// get the tokenID
			token, err := flags.GetUint64(FlagTokenID)
			if err != nil {
				return err
			}

			// get the amount
			amount, err := flags.GetUint64(FlagAmount)
			if err != nil {
				return err
			}

			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}

			DBI, err := db.NewDB(cfg)
			if err != nil {
				return err
			}
			defer DBI.Close()

			bazooka, err := bazooka.NewPreLoadedBazooka(cfg)
			if err != nil {
				return err
			}

			tokenAddr, err := bazooka.GetTokenAddress(token)
			if err != nil {
				return err
			}

			approveTx, err := bazooka.ApproveToken(tokenAddr, common.HexToAddress(bazooka.Cfg.DepositManager), amount)
			if err != nil {
				return err
			}
			fmt.Println("Approved tokens", approveTx)
			toPubHexBz, err := hex.DecodeString(pubKey)
			if err != nil {
				return err
			}
			toPubkey, err := core.Pubkey(toPubHexBz).ToSol()
			if err != nil {
				return err
			}

			txHash, err := bazooka.Deposit(toPubkey, token, amount)
			if err != nil {
				return err
			}

			fmt.Println("hash and err", txHash, err)
			return nil
		},
	}
	cmd.Flags().StringP(FlagPubKey, "", "", "--pubkey=<pubkey>")
	cmd.Flags().Uint64P(FlagTokenID, "", 0, "--token-id=<tokenID>")
	cmd.Flags().Uint64P(FlagAmount, "", 0, "--amount=<amount>")
	return cmd
}
