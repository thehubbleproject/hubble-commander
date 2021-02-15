package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/bidder"
	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/db"
	"github.com/BOPR/listener"
	hlog "github.com/BOPR/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

// startCmd starts the daemon
func startCmd() *cobra.Command {
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
			// aggregator := agg.NewAggregator(cfg)
			// create the syncer service
			syncer := listener.NewSyncer(cfg)
			bidderInstance := bidder.NewBidder(cfg)

			// if no row is found then we are starting the node for the first time
			syncStatus, err := DBI.GetSyncStatus()
			if err != nil && gorm.IsRecordNotFoundError(err) {
				storeGenesisData(&bz, &DBI, cfg)
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
					// if err := aggregator.Stop(); err != nil {
					// 	log.Fatalln("Unable to stop aggregator", "error", err)
					// }
					if err := syncer.Stop(); err != nil {
						log.Fatalln("Unable to stop syncer", "error", err)
					}
					if err := bidderInstance.Stop(); err != nil {
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
			if err := bidderInstance.Start(); err != nil {
				log.Fatalln("Unable to start syncer", "error", err)
			}

			// if err := aggregator.Start(); err != nil {
			// 	log.Fatalln("Unable to start aggregator", "error", err)
			// }

			runtime.Goexit()
		},
	}
}

func storeGenesisData(bz *bazooka.Bazooka, DBI *db.DB, cfg config.Configuration) {
	err := DBI.InitStateTree(int(cfg.MaxTreeDepth))
	common.PanicIfError(err)

	opts := bind.CallOpts{From: ethCmn.HexToAddress(cfg.OperatorAddress)}
	accountTreeDepth, err := bz.SC.AccountRegistry.DEPTH(&opts)
	common.PanicIfError(err)

	err = DBI.InitAccountTree(int(accountTreeDepth.Uint64()) + 1)
	common.PanicIfError(err)

	newParams := core.Params{StakeAmount: cfg.StakeAmount, MaxDepth: cfg.MaxTreeDepth, MaxDepositSubTreeHeight: cfg.MaxDepositSubtree}
	err = DBI.UpdateStakeAmount(newParams.StakeAmount)
	common.PanicIfError(err)
	err = DBI.UpdateMaxDepth(newParams.MaxDepth)
	common.PanicIfError(err)
	err = DBI.UpdateDepositSubTreeHeight(newParams.MaxDepositSubTreeHeight)
	common.PanicIfError(err)
	err = DBI.UpdateFinalisationTimePerBatch(40320)
	common.PanicIfError(err)
	err = DBI.InitSyncStatus(cfg.GenesisEth1Block)
	common.PanicIfError(err)
}
