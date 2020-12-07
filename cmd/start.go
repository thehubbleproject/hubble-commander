package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"

	agg "github.com/BOPR/aggregator"
	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/listener"
	"github.com/BOPR/rest"

	"github.com/BOPR/core"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd starts the daemon
func startCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Starts hubble daemon",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			// populate global config objects
			initConfigAndGlobals()

			logger := common.Logger.With("module", "start")
			// create aggregator service
			aggregator := agg.NewAggregator()
			// create the syncer service
			syncer := listener.NewSyncer()

			// if no row is found then we are starting the node for the first time
			syncStatus, err := core.DBInstance.GetSyncStatus()
			if err != nil && gorm.IsRecordNotFoundError(err) {
				// read genesis file
				genesis, err := config.ReadGenesisFile()
				common.PanicIfError(err)
				// loads genesis data to the database
				loadGenesisData(genesis)
			} else if err != nil && !gorm.IsRecordNotFoundError(err) {
				logger.Error("Error connecting to database", "error", err)
				common.PanicIfError(err)
			}

			logger.Info("Starting coordinator with sync and aggregator enabled", "lastSyncedEthBlock",
				syncStatus.LastEthBlockBigInt().String())

			// go routine to catch signal
			catchSignal := make(chan os.Signal, 1)
			signal.Notify(catchSignal, os.Interrupt)

			go func() {
				// sig is a ^C, handle it
				for range catchSignal {
					if err := aggregator.Stop(); err != nil {
						log.Fatalln("Unable to stop aggregator", "error", err)
					}
					if err := syncer.Stop(); err != nil {
						log.Fatalln("Unable to stop syncer", "error", err)
					}

					core.DBInstance.Close()
					// exit
					os.Exit(1)
				}
			}()

			if err := syncer.Start(); err != nil {
				log.Fatalln("Unable to start syncer", "error", err)
			}

			if err := aggregator.Start(); err != nil {
				log.Fatalln("Unable to start aggregator", "error", err)
			}
			r := rest.LoadRouters()
			http.Handle("/", &r)

			// TODO replace this with port from config
			err = http.ListenAndServe(":3000", &r)
			if err != nil {
				panic(err)
			}
			fmt.Println("Server started on port 3000 🎉")
		},
	}
}

func loadGenesisData(genesis config.Genesis) {
	err := genesis.Validate()
	if err != nil {
		common.PanicIfError(err)
	}
	genesisAccounts := genesis.GenesisAccounts.Accounts

	diff := int(math.Exp2(float64(genesis.MaxTreeDepth))) - len(genesisAccounts)

	var states []core.UserState
	var accounts []core.Account

	var zeroData []byte

	// convert genesis accounts to user accounts
	for i, acc := range genesisAccounts {
		path, err := core.SolidityPathToNodePath(acc.AccountID, genesis.MaxTreeDepth)
		common.PanicIfError(err)

		// use contracts to get coordinator state bytes
		stateBytes, err := core.LoadedBazooka.EncodeState(acc.AccountID, acc.Balance, acc.Nonce, acc.TokenType)
		common.PanicIfError(err)

		if i == 0 {
			zeroData = stateBytes
		}

		newState := core.NewUserState(acc.AccountID, core.STATUS_ACTIVE, path, stateBytes)
		newAcc, err := core.NewAccount(acc.AccountID, acc.PublicKey, path)
		common.PanicIfError(err)

		states = append(states, *newState)
		accounts = append(accounts, *newAcc)
	}

	for ; diff > 0; diff-- {
		// create empty state
		newEmptyState := core.EmptyUserState()
		newEmptyState.Data = zeroData
		newEmptyState.Hash = core.ZERO_VALUE_LEAF.String()
		states = append(states, newEmptyState)

		// create empty account
		newAccount := core.NewEmptyAccount()
		newAccount.Hash = core.ZERO_VALUE_LEAF.String()
		accounts = append(accounts, *newAccount)
	}

	err = core.DBInstance.InitStateTree(genesis.MaxTreeDepth, states)
	common.PanicIfError(err)

	err = core.DBInstance.InitAccountTree(genesis.MaxTreeDepth, accounts)
	common.PanicIfError(err)

	// load params
	newParams := core.Params{StakeAmount: genesis.StakeAmount, MaxDepth: genesis.MaxTreeDepth, MaxDepositSubTreeHeight: genesis.MaxDepositSubTreeHeight}
	err = core.DBInstance.UpdateStakeAmount(newParams.StakeAmount)
	common.PanicIfError(err)
	err = core.DBInstance.UpdateMaxDepth(newParams.MaxDepth)
	common.PanicIfError(err)
	err = core.DBInstance.UpdateDepositSubTreeHeight(newParams.MaxDepositSubTreeHeight)
	common.PanicIfError(err)
	err = core.DBInstance.UpdateFinalisationTimePerBatch(40320)
	common.PanicIfError(err)

	// load sync status
	err = core.DBInstance.InitSyncStatus(genesis.StartEthBlock)
	common.PanicIfError(err)
}

func initConfigAndGlobals() {
	readAndInitGlobalConfig()
	initGlobalDBInstance()
	initGlobalBazooka()
}

func readAndInitGlobalConfig() {
	// create viper object
	viperObj := viper.New()

	// get current directory
	dir, err := os.Getwd()
	common.PanicIfError(err)

	// set config paths
	viperObj.SetConfigName(ConfigFileName) // name of config file (without extension)
	viperObj.AddConfigPath(dir)

	// finally! read config
	err = viperObj.ReadInConfig()
	common.PanicIfError(err)

	// unmarshall to the configration object
	var cfg config.Configuration
	if err = viperObj.UnmarshalExact(&cfg); err != nil {
		common.PanicIfError(err)
	}

	// init global config
	config.GlobalCfg = cfg
	// TODO use a better way to handle priv keys post testnet
	common.PanicIfError(config.SetOperatorKeys(config.GlobalCfg.OperatorKey))
}

func initGlobalDBInstance() {
	// create db Instance
	tempDB, err := core.NewDB()
	common.PanicIfError(err)

	// init global DB instance
	core.DBInstance = tempDB
}

func initGlobalBazooka() {
	var err error
	// create and init global config object
	core.LoadedBazooka, err = core.NewPreLoadedBazooka()
	common.PanicIfError(err)
}
