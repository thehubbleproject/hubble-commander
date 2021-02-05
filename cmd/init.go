package main

import (
	"fmt"
	"strconv"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/spf13/cobra"
)

const (
	defaultGenesisPath = ".contracts/genesis.json"
)

// initCmd generated init command to initialise the config file
func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialises Configration for BOPR",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := config.GetDefaultConfig()
			config.WriteConfigFile("./config.toml", &defaultConfig)
		},
	}
}

func configureGenesisCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "load",
		Short: "Load Genesis from a genesis.json",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.ParseConfig()
			common.PanicIfError(err)
			path := ""
			if len(args) == 0 {
				path = defaultGenesisPath
			} else {
				path = args[0]
			}
			fmt.Printf("Loading genesis from %s\n", path)

			genesis, err := config.ReadGenesisFile(path)
			common.PanicIfError(err)

			common.PanicIfError(genesis.Validate())

			// contracts
			cfg.RollupAddress = genesis.Addresses.Rollup
			cfg.BurnAuction = genesis.Addresses.Chooser
			cfg.TokenRegistry = genesis.Addresses.TokenRegistry
			cfg.AccountRegistry = genesis.Addresses.BlsAccountRegistry
			cfg.DepositManager = genesis.Addresses.DepositManager
			cfg.State = genesis.Addresses.FrontendGeneric
			cfg.Transfer = genesis.Addresses.FrontendTransfer
			cfg.MassMigration = genesis.Addresses.FrontendMassMigration
			cfg.Create2Transfer = genesis.Addresses.FrontendCreate2Transfer

			cfg.MaxTreeDepth = genesis.Parameters.MAXDEPTH
			cfg.MaxDepositSubtree = genesis.Parameters.MAXDEPOSITSUBTREEDEPTH
			cfg.GenesisEth1Block = genesis.Auxiliary.GenesisEth1Block
			stakeAmount, err := strconv.Atoi(genesis.Parameters.STAKEAMOUNT)
			common.PanicIfError(err)
			cfg.StakeAmount = uint64(stakeAmount)
			cfg.AppID = (genesis.Auxiliary.Domain)
			config.WriteConfigFile("./config.toml", &cfg)
		},
	}
}
