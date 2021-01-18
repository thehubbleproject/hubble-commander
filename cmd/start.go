package main

import (
	"log"
	"math"
	"os"
	"os/signal"
	"runtime"

	agg "github.com/BOPR/aggregator"
	"github.com/BOPR/bazooka"
	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/db"
	"github.com/BOPR/listener"
	hlog "github.com/BOPR/log"
	"github.com/common-nighthawk/go-figure"

	"github.com/BOPR/core"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

// startCmd starts the daemon
func startCmd() *cobra.Command {
	myFigure := figure.NewColorFigure("COMMANDER", "", "green", true)
	myFigure.Print()
	return &cobra.Command{
		Use:   "start",
		Short: "Starts hubble daemon",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			cfg, err := config.ParseConfig()
			common.PanicIfError(err)

			bz, err := bazooka.NewPreLoadedBazooka(cfg)
			common.PanicIfError(err)

			DBI, err := db.NewDB(cfg)
			common.PanicIfError(err)

			logger := hlog.Logger.With("module", "start")
			// create aggregator service
			aggregator := agg.NewAggregator(cfg)
			// create the syncer service
			syncer := listener.NewSyncer(cfg)

			// if no row is found then we are starting the node for the first time
			syncStatus, err := DBI.GetSyncStatus()
			if err != nil && gorm.IsRecordNotFoundError(err) {
				// read genesis file
				genesis, err := config.ReadGenesisFile()
				common.PanicIfError(err)
				// loads genesis data to the database
				loadGenesisData(&bz, &DBI, genesis)
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
					DBI.Close()
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

			runtime.Goexit()

		},
	}
}

func loadGenesisData(bz *bazooka.Bazooka, DBI *db.DB, genesis config.Genesis) {
	err := genesis.Validate()
	if err != nil {
		common.PanicIfError(err)
	}

	var states []core.UserState
	var accounts []core.Account

	for i := 0; i < int(math.Exp2(float64(genesis.MaxTreeDepth))); i++ {
		// create empty state
		newEmptyState := core.EmptyUserState()
		newEmptyState.Hash = core.ZeroLeaf.String()
		states = append(states, newEmptyState)

		// create empty account
		newAccount := core.NewEmptyAccount()
		newAccount.Hash = core.ZeroLeaf.String()
		accounts = append(accounts, *newAccount)
	}

	err = DBI.InitStateTree(genesis.MaxTreeDepth, states)
	common.PanicIfError(err)

	err = DBI.InitAccountTree(genesis.MaxTreeDepth, accounts)
	common.PanicIfError(err)

	// load params
	newParams := core.Params{StakeAmount: genesis.StakeAmount, MaxDepth: genesis.MaxTreeDepth, MaxDepositSubTreeHeight: genesis.MaxDepositSubTreeHeight}
	err = DBI.UpdateStakeAmount(newParams.StakeAmount)
	common.PanicIfError(err)
	err = DBI.UpdateMaxDepth(newParams.MaxDepth)
	common.PanicIfError(err)
	err = DBI.UpdateDepositSubTreeHeight(newParams.MaxDepositSubTreeHeight)
	common.PanicIfError(err)
	err = DBI.UpdateFinalisationTimePerBatch(40320)
	common.PanicIfError(err)

	// load sync status
	err = DBI.InitSyncStatus(genesis.StartEthBlock)
	common.PanicIfError(err)
}
