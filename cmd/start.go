package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"

	agg "github.com/BOPR/aggregator"
	"github.com/BOPR/common"
	"github.com/BOPR/config"

	"github.com/BOPR/core"
	"github.com/BOPR/rest"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// StartCmd starts the daemon
func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Starts hubble daemon",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			// populate global config object
			ReadAndInitGlobalConfig()

			InitGlobalDBInstance()

			InitGlobalBazooka()

			logger := common.Logger.With("module", "hubble")

			//
			// Create all the required services
			//

			// create aggregator service
			aggregator := agg.NewAggregator()

			// create the syncer service
			// syncer := listener.NewSyncer()

			// if no row is found then we are starting the node for the first time
			syncStatus, err := core.DBInstance.GetSyncStatus()
			if err != nil && gorm.IsRecordNotFoundError(err) {
				// read genesis file
				genesis, err := config.ReadGenesisFile()
				common.PanicIfError(err)

				// loads genesis data to the database
				LoadGenesisData(genesis)
			} else if err != nil && !gorm.IsRecordNotFoundError(err) {
				logger.Error("Error connecting to database", "error", err)
				common.PanicIfError(err)
			}

			logger.Info("Starting coordinator with sync and aggregator enabled", "lastSyncedEthBlock",
				syncStatus.LastEthBlockBigInt().String(),
				"lastSyncedBatch", syncStatus.LastBatchRecorded)

			// go routine to catch signal
			catchSignal := make(chan os.Signal, 1)
			signal.Notify(catchSignal, os.Interrupt)
			go func() {
				// sig is a ^C, handle it
				for range catchSignal {
					aggregator.Stop()
					// syncer.Stop()
					core.DBInstance.Close()

					// exit
					os.Exit(1)
				}
			}()

			r := mux.NewRouter()
			r.HandleFunc("/tx", rest.TxReceiverHandler).Methods("POST")
			r.HandleFunc("/account", rest.GetAccountHandler).Methods("GET")
			http.Handle("/", r)

			// if err := syncer.Start(); err != nil {
			// 	log.Fatalln("Unable to start syncer", "error")
			// }

			if err := aggregator.Start(); err != nil {
				log.Fatalln("Unable to start aggregator", "error", err)
			}
			// TODO replace this with port from config
			err = http.ListenAndServe(":3000", r)
			if err != nil {
				panic(err)
			}
			fmt.Println("Server started on port 3000 🎉")
		},
	}
}

func ReadAndInitGlobalConfig() {
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

func InitGlobalDBInstance() {
	// create db Instance
	tempDB, err := core.NewDB()
	common.PanicIfError(err)

	// init global DB instance
	core.DBInstance = tempDB
}

func InitGlobalBazooka() {
	var err error
	// create and init global config object
	core.LoadedBazooka, err = core.NewPreLoadedBazooka()
	common.PanicIfError(err)
}

// LoadGenesisData helps load the genesis data into the DB
func LoadGenesisData(genesis config.Genesis) {
	err := genesis.Validate()
	if err != nil {
		common.PanicIfError(err)
	}

	genesisAccounts, err := core.LoadedBazooka.GetGenesisAccounts()
	common.PanicIfError(err)
	zeroAccount := genesisAccounts[0]
	diff := int(math.Exp2(float64(genesis.MaxTreeDepth))) - len(genesisAccounts)
	var allAccounts []core.UserState
	var allAccountLeaf []core.Account

	// convert genesis accounts to user accounts
	for _, account := range genesisAccounts {
		pubkeyHash := core.ZERO_VALUE_LEAF.String()
		allAccountLeaf = append(allAccountLeaf, core.Account{Hash: pubkeyHash})
		allAccounts = append(
			allAccounts,
			account,
		)
	}

	// fill the tree with zero leaves
	for diff > 0 {
		newAcc := core.EmptyAccount()
		newAcc.Data = zeroAccount.Data
		newAcc.Hash = core.ZERO_VALUE_LEAF.String()
		allAccounts = append(allAccounts, newAcc)
		newAccount := core.NewEmptyAccount()
		newAccount.Hash = core.ZERO_VALUE_LEAF.String()
		allAccountLeaf = append(allAccountLeaf, *newAccount)
		diff--
	}
	allAccounts = AlterRedditAccounts(allAccounts)
	allAccountLeaf = AlterRedditPubkeys(allAccountLeaf)

	// load accounts
	err = core.DBInstance.InitBalancesTree(genesis.MaxTreeDepth, allAccounts)
	common.PanicIfError(err)

	err = core.DBInstance.InitAccountTree(genesis.MaxTreeDepth, allAccountLeaf)
	common.PanicIfError(err)

	// load params
	newParams := core.Params{StakeAmount: genesis.StakeAmount, MaxDepth: genesis.MaxTreeDepth, MaxDepositSubTreeHeight: genesis.MaxDepositSubTreeHeight}
	core.DBInstance.UpdateStakeAmount(newParams.StakeAmount)
	core.DBInstance.UpdateMaxDepth(newParams.MaxDepth)
	core.DBInstance.UpdateDepositSubTreeHeight(newParams.MaxDepositSubTreeHeight)
	core.DBInstance.UpdateFinalisationTimePerBatch(40320)

	// load sync status
	err = core.DBInstance.InitSyncStatus(genesis.StartEthBlock)
	if err != nil {
		panic(err)
	}
}

type userList struct {
	Users []user `json:"users"`
}

type user struct {
	Address   string `json:"address"`
	PublicKey string `json:"pubkey"`
	PrivKey   string `json:"privkey"`
}

func ReadUsers() (reddit []user, err error) {
	var userListInstance userList
	users, err := os.Open("users.json")
	if err != nil {
		return
	}
	defer users.Close()

	genBytes, err := ioutil.ReadAll(users)
	if err != nil {
		return
	}

	err = json.Unmarshal(genBytes, &userListInstance)
	if err != nil {
		panic(err)
	}
	return userListInstance.Users, nil
}

func AlterRedditAccounts(allAccounts []core.UserState) (updatedAccounts []core.UserState) {
	account2 := allAccounts[2]
	account3 := allAccounts[3]
	bazooka, err := core.NewPreLoadedBazooka()
	if err != nil {
		panic(err)
	}

	account2Data, err := bazooka.EncodeAccount(int64(account2.AccountID), 2000000, 0, 1, 0, 0)
	if err != nil {
		panic(err)
	}
	account2.Data = account2Data
	account2.AccountID = 2
	account2.CreateAccountHash()

	account3Data, err := bazooka.EncodeAccount(int64(account3.AccountID), 1000000, 0, 1, 0, 0)
	if err != nil {
		panic(err)
	}
	account3.Data = account3Data
	account3.AccountID = 3
	account3.CreateAccountHash()

	allAccounts[2] = account2
	allAccounts[3] = account3

	return allAccounts
}

func AlterRedditPubkeys(allAccounts []core.Account) (updatedAccounts []core.Account) {
	users, err := ReadUsers()
	if err != nil {
		panic(err)
	}
	account2 := allAccounts[2]
	account3 := allAccounts[3]
	account2.PublicKey = users[0].PublicKey
	account2.ID = 2
	account2.PopulateHash()
	account3.PublicKey = users[1].PublicKey
	account3.ID = 3
	account3.PopulateHash()
	allAccounts[2] = account2
	allAccounts[3] = account3
	return allAccounts
}
