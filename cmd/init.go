package main

import (
	"encoding/hex"
	"strconv"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/spf13/cobra"
)

// initCmd generated init command to initialise the config file
func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialises Configration for BOPR",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := config.GetDefaultConfig()
			operatorKey, err := config.GenOperatorKey()
			common.PanicIfError(err)
			defaultConfig.OperatorKey = hex.EncodeToString(operatorKey)
			address, err := config.PrivKeyStringToAddress(hex.EncodeToString(operatorKey))
			common.PanicIfError(err)
			defaultConfig.OperatorAddress = address.String()
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
			genesis, err := config.ReadGenesisFile(args[0])
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
			config.WriteConfigFile("./config.toml", &cfg)
		},
	}
}
