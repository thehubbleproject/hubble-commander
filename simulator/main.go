package simulator

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
)

const (
	SimulatorService = "simulator"
	BATCH_SIZE       = 32
	AIRDROP_AMOUNT   = 10
	TRANSFER_AMOUNT  = 1
	TOKEN            = 1
	BURN_AMOUNT      = 1
	REDDIT_ACCOUNT   = 2
	REDDIT_KEY       = "5f9b54d6b94235bc56b537b96bf19746ab2ff29007a353847108b494e124fda9"
)

type Simulator struct {
	// Base service
	core.BaseService

	// DB instance
	DB core.DB

	// contract caller to interact with contracts
	LoadedBazooka core.Bazooka

	// header listener subscription
	cancelSimulator context.CancelFunc

	toSwap bool

	accounts map[uint64]User
}

// NewSimulator returns new simulator object
func NewSimulator() *Simulator {
	logger := common.Logger.With("module", SimulatorService)
	sim := &Simulator{}
	sim.BaseService = *core.NewBaseService(logger, SimulatorService, sim)
	db, err := core.NewDB()
	if err != nil {
		panic(err)
	}
	sim.LoadedBazooka, err = core.NewPreLoadedBazooka()
	if err != nil {
		panic(err)
	}
	sim.accounts = make(map[uint64]User)

	sim.DB = db
	return sim
}

type UserList struct {
	Users []User `json:"users"`
}

type User struct {
	PublicKey string `json:"pubkey"`
	PrivKey   string `json:"privkey"`
}

func (s *Simulator) ReadUsers() error {
	var userListInstance UserList
	users, err := os.Open("users.json")
	if err != nil {
		return err
	}
	defer users.Close()

	genBytes, err := ioutil.ReadAll(users)
	if err != nil {
		return err
	}

	err = json.Unmarshal(genBytes, &userListInstance)
	for i, user := range userListInstance.Users {
		s.accounts[uint64(i+2)] = user
	}

	return err
}

// OnStart starts new block subscription
func (s *Simulator) OnStart() error {
	s.BaseService.OnStart() // Always call the overridden method.
	ctx, cancelSimulator := context.WithCancel(context.Background())
	s.cancelSimulator = cancelSimulator
	totalCycles, err := s.DB.CycleCount()
	if err != nil {
		panic(err)
	}
	if totalCycles == 0 {
		startIndex := uint64(4)
		s.DB.LogCycle(core.STAGE_TRANSFER, startIndex, startIndex+BATCH_SIZE)
	}
	err = s.ReadUsers()
	if err != nil {
		panic(err)
	}

	go s.SimulationStart(ctx, 10*time.Second)
	s.toSwap = false
	return nil
}

// OnStop stops all necessary go routines
func (s *Simulator) OnStop() {
	s.BaseService.OnStop() // Always call the overridden method.
	s.DB.Close()
	s.cancelSimulator()
}

// SimulationStart starts the simulator
func (s *Simulator) SimulationStart(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	// stop ticker when everything done
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			s.simulateRedditTransfers()
			// pick batch from DB
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (s *Simulator) simulateRedditTransfers() {
	fmt.Println("Starting to simulate")
	// Reddit key
	BATCH_SIZE := 1024
	latestFromAcc, err := s.DB.GetAccountByIndex(REDDIT_ACCOUNT)
	if err != nil {
		s.Logger.Error("unable to fetch latest account", "error", err)
		return
	}

	_, _, nonce, _, _, _, err := s.LoadedBazooka.DecodeAccount(latestFromAcc.Data)
	if err != nil {
		s.Logger.Error("unable to decode account", "error", err)
		return
	}
	var txs []core.Tx
	for i := int64(0); i < int64(BATCH_SIZE); i++ {
		txBytes, err := s.LoadedBazooka.EncodeTransferTx(2, 3, TOKEN, nonce.Int64()+i, TRANSFER_AMOUNT, core.TX_TRANSFER_TYPE)
		if err != nil {
			s.Logger.Error("unable to encode tx", "error", err)
			return
		}
		s.Logger.Debug("Encoded tx", "index", i)
		txCore := core.NewPendingTx(2, 3, core.TX_TRANSFER_TYPE, []byte{}, txBytes)
		signBytes, err := s.LoadedBazooka.SignBytesForTransfer(int64(txCore.Type), int64(txCore.From), int64(txCore.To), nonce.Int64()+i, TRANSFER_AMOUNT)
		if err != nil {
			s.Logger.Error("unable to encode tx", "error", err)
			return
		}
		err = txCore.SignTx(s.accounts[REDDIT_ACCOUNT].PrivKey, s.accounts[REDDIT_ACCOUNT].PublicKey, signBytes)
		if err != nil {
			s.Logger.Error("unable to sign transaction", "error", err)
			return
		}
		txs = append(txs, txCore)
	}
	for i := 0; i < len(txs); i++ {
		err = s.DB.InsertTx(&txs[i])
		if err != nil {
			s.Logger.Error("unable to insert tx", "error", err)
			return
		}
		s.Logger.Info("Sent a tx!", "TxHash", txs[i].TxHash, "From", txs[i].From, "To", txs[i].To)
	}
}

// // tries sending transactins to and fro accounts to the rollup node
// func (s *Simulator) startCycle() {
// 	lastCycle, err := s.DB.GetLastCycle()
// 	if err != nil {
// 		s.Logger.Error("Error getting last cycle info", "error", err)
// 		return
// 	}
// 	fmt.Println("Last cycle info:", lastCycle)
// 	switch lastCycle.Stage {
// 	case core.STAGE_TRANSFER:
// 		fmt.Println("Starting create account cycle")
// 		// do something
// 		s.simulateCreateAccounts(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
// 		s.DB.LogCycle(core.STAGE_ACCOUNT_CREATE, lastCycle.StartIndex, lastCycle.EndIndex)
// 	case core.STAGE_ACCOUNT_CREATE:
// 		fmt.Println("Starting burn consent cycle")
// 		// do something
// 		s.simulateBurnConsent(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
// 		s.DB.LogCycle(core.STAGE_BURN_CONSENT, lastCycle.StartIndex, lastCycle.EndIndex)
// 	case core.STAGE_BURN_CONSENT:
// 		fmt.Println("Starting airdrop cycle")
// 		// do something
// 		s.simulateAirdrop(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
// 		s.DB.LogCycle(core.STAGE_AIRDROP, lastCycle.StartIndex, lastCycle.EndIndex)
// 	case core.STAGE_AIRDROP:
// 		fmt.Println("Starting burn exec cycle")
// 		// do something
// 		s.simulateBurnExec(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
// 		s.DB.LogCycle(core.STAGE_BURN_EXEC, lastCycle.StartIndex, lastCycle.EndIndex)
// 	case core.STAGE_BURN_EXEC:
// 		fmt.Println("Starting transfer cycle")
// 		// s.simulateTransfer(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
// 		// Mark cycle complete, update the indexes
// 		s.Logger.Info("Simulation cycle ending", "startIndex", lastCycle.StartIndex, "end", lastCycle.EndIndex)
// 		s.DB.LogCycle(core.STAGE_TRANSFER, lastCycle.EndIndex+1, lastCycle.EndIndex+BATCH_SIZE)
// 	}
// }

// func (s *Simulator) simulateCreateAccounts(startIndex, endIndex int64) {
// 	for i := int64(0); i < BATCH_SIZE; i++ {
// 		txBytes, err := s.LoadedBazooka.EncodeCreateAccountTx(startIndex+i, startIndex+i, TOKEN)
// 		if err != nil {
// 			s.Logger.Error("unable to encode tx", "error", err)
// 			return
// 		}
// 		txCore := core.NewPendingTx(0, uint64(startIndex+i), core.TX_CREATE_ACCOUNT, []byte{}, txBytes)

// 		err = s.DB.InsertTx(&txCore)
// 		if err != nil {
// 			s.Logger.Error("unable to insert tx", "error", err)
// 			return
// 		}
// 		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
// 	}
// }

// func (s *Simulator) simulateBurnConsent(startIndex, endIndex int64) {
// 	for i := int64(0); i < BATCH_SIZE; i++ {
// 		latestFromAcc, err := s.DB.GetAccountByIndex(uint64(startIndex + i))
// 		if err != nil {
// 			s.Logger.Error("unable to fetch latest account", "error", err)
// 			return
// 		}
// 		_, _, nonce, _, _, _, err := s.LoadedBazooka.DecodeAccount(latestFromAcc.Data)
// 		if err != nil {
// 			s.Logger.Error("unable to decode account", "error", err)
// 			return
// 		}
// 		txBytes, err := s.LoadedBazooka.EncodeBurnConsentTx(startIndex+i, BURN_AMOUNT, nonce.Int64()+1, core.TX_BURN_CONSENT)
// 		if err != nil {
// 			s.Logger.Error("unable to encode tx", "error", err)
// 			return
// 		}
// 		txCore := core.NewPendingTx(uint64(startIndex+i), 0, core.TX_BURN_CONSENT, []byte{}, txBytes)
// 		signBytes, err := s.LoadedBazooka.SignBytesForBurnConsent(int64(txCore.Type), int64(txCore.From), nonce.Int64()+1, BURN_AMOUNT)
// 		if err != nil {
// 			s.Logger.Error("unable to encode tx", "error", err)
// 			return
// 		}
// 		// sign it
// 		err = txCore.SignTx(s.accounts[uint64(startIndex+i)], signBytes)
// 		if err != nil {
// 			s.Logger.Error("unable to sign tx", "error", err)
// 			return
// 		}
// 		err = s.DB.InsertTx(&txCore)
// 		if err != nil {
// 			s.Logger.Error("unable to insert tx", "error", err)
// 			return
// 		}
// 		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
// 	}
// }

// func (s *Simulator) simulateTransfer(startIndex, endIndex int64) {
// 	for i := int64(0); i < BATCH_SIZE; i++ {
// 		latestFromAcc, err := s.DB.GetAccountByIndex(uint64(startIndex + i))
// 		if err != nil {
// 			s.Logger.Error("unable to fetch latest account", "error", err)
// 			return
// 		}
// 		_, _, nonce, _, _, _, err := s.LoadedBazooka.DecodeAccount(latestFromAcc.Data)
// 		if err != nil {
// 			s.Logger.Error("unable to decode account", "error", err)
// 			return
// 		}
// 		txBytes, err := s.LoadedBazooka.EncodeTransferTx(startIndex+i, REDDIT_ACCOUNT, TOKEN, nonce.Int64()+1, TRANSFER_AMOUNT, core.TX_TRANSFER_TYPE)
// 		if err != nil {
// 			s.Logger.Error("unable to encode tx", "error", err)
// 			return
// 		}
// 		txCore := core.NewPendingTx(uint64(startIndex+i), REDDIT_ACCOUNT, core.TX_TRANSFER_TYPE, []byte{}, txBytes)
// 		signBytes, err := s.LoadedBazooka.SignBytesForTransfer(int64(txCore.Type), int64(txCore.From), int64(txCore.To), nonce.Int64()+1, TRANSFER_AMOUNT)
// 		if err != nil {
// 			s.Logger.Error("unable to encode tx", "error", err)
// 			return
// 		}
// 		// sign it
// 		err = txCore.SignTx(s.accounts[uint64(startIndex+i)], signBytes)
// 		if err != nil {
// 			s.Logger.Error("unable to sign tx", "error", err)
// 			return
// 		}
// 		err = s.DB.InsertTx(&txCore)
// 		if err != nil {
// 			s.Logger.Error("unable to insert tx", "error", err)
// 			return
// 		}
// 		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
// 	}
// }

// func (s *Simulator) simulateAirdrop(startIndex, endIndex int64) {
// 	var redditNonce int64
// 	latestFromAcc, err := s.DB.GetAccountByIndex(REDDIT_ACCOUNT)
// 	if err != nil {
// 		s.Logger.Error("unable to fetch latest account", "error", err)
// 		return
// 	}
// 	_, _, nonce, _, _, _, err := s.LoadedBazooka.DecodeAccount(latestFromAcc.Data)
// 	if err != nil {
// 		s.Logger.Error("unable to decode account", "error", err)
// 		return
// 	}
// 	redditNonce = nonce.Int64() + 1
// 	for i := int64(0); i < BATCH_SIZE; i++ {
// 		txBytes, err := s.LoadedBazooka.EncodeAirdropTx(REDDIT_ACCOUNT, startIndex+i, TOKEN, redditNonce, TRANSFER_AMOUNT, core.TX_AIRDROP_TYPE)
// 		if err != nil {
// 			s.Logger.Error("unable to encode tx", "error", err)
// 			return
// 		}
// 		txCore := core.NewPendingTx(REDDIT_ACCOUNT, uint64(startIndex+i), core.TX_AIRDROP_TYPE, []byte{}, txBytes)
// 		signBytes, err := s.LoadedBazooka.SignBytesForAirdrop(int64(txCore.Type), int64(txCore.From), int64(txCore.To), redditNonce, TRANSFER_AMOUNT)
// 		if err != nil {
// 			s.Logger.Error("unable to encode tx", "error", err)
// 			return
// 		}
// 		// sign it
// 		err = txCore.SignTx(REDDIT_KEY, signBytes)
// 		if err != nil {
// 			s.Logger.Error("unable to sign tx", "error", err)
// 			return
// 		}
// 		fmt.Println("signature details", txCore.Signature, redditNonce)
// 		err = s.DB.InsertTx(&txCore)
// 		if err != nil {
// 			s.Logger.Error("unable to insert tx", "error", err)
// 			return
// 		}
// 		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
// 		redditNonce++
// 	}
// }

// func (s *Simulator) simulateBurnExec(startIndex, endIndex int64) {
// 	for i := int64(0); i < BATCH_SIZE; i++ {
// 		txBytes, err := s.LoadedBazooka.EncodeBurnExecTx(startIndex+i, core.TX_BURN_EXEC)
// 		if err != nil {
// 			s.Logger.Error("unable to encode tx", "error", err)
// 			return
// 		}
// 		txCore := core.NewPendingTx(uint64(startIndex+i), REDDIT_ACCOUNT, core.TX_BURN_EXEC, []byte{}, txBytes)
// 		err = s.DB.InsertTx(&txCore)
// 		if err != nil {
// 			s.Logger.Error("unable to insert tx", "error", err)
// 			return
// 		}
// 		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
// 	}
// }
