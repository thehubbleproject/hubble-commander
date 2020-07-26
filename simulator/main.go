package simulator

import (
	"context"
	"fmt"
	"time"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
)

const (
	SimulatorService = "simulator"
	BATCH_SIZE       = 32
	AIRDROP_AMOUNT   = 10
	TRANSFER_AMOUNT  = 1
	BURN_AMOUNT      = 1
	REDDIT_ACCOUNT   = 1
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

	sim.DB = db
	return sim
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
		// firstEmptyAccount, err := s.DB.GetFirstEmptyAccount()
		// if err != nil {
		// 	panic(err)
		// }
		// startIndex, err := core.StringToUint(firstEmptyAccount.Path)
		// if err != nil {
		// 	panic(err)
		// }
		startIndex := uint64(2)
		s.DB.LogCycle(core.STAGE_TRANSFER, startIndex, startIndex+BATCH_SIZE)
	}
	go s.SimulationStart(ctx, 30*time.Second)

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
			s.startCycle()
			// pick batch from DB
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// tries sending transactins to and fro accounts to the rollup node
func (s *Simulator) startCycle() {
	lastCycle, err := s.DB.GetLastCycle()
	if err != nil {
		s.Logger.Error("Error getting last cycle info", "error", err)
		return
	}
	fmt.Println("Last cycle info:", lastCycle)
	switch lastCycle.Stage {
	case core.STAGE_TRANSFER:
		fmt.Println("Starting create account cycle")
		// do something
		s.simulateCreateAccounts(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
		s.DB.LogCycle(core.STAGE_ACCOUNT_CREATE, lastCycle.StartIndex, lastCycle.EndIndex)
	case core.STAGE_ACCOUNT_CREATE:
		fmt.Println("Starting burn consent cycle")
		// do something
		s.simulateBurnConsent(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
		s.DB.LogCycle(core.STAGE_BURN_CONSENT, lastCycle.StartIndex, lastCycle.EndIndex)
	case core.STAGE_BURN_CONSENT:
		fmt.Println("Starting airdrop cycle")
		// do something
		s.simulateAirdrop(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
		s.DB.LogCycle(core.STAGE_AIRDROP, lastCycle.StartIndex, lastCycle.EndIndex)
	case core.STAGE_AIRDROP:
		fmt.Println("Starting burn exec cycle")
		// do something
		s.simulateBurnExec(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
		s.DB.LogCycle(core.STAGE_BURN_EXEC, lastCycle.StartIndex, lastCycle.EndIndex)
	case core.STAGE_BURN_EXEC:
		fmt.Println("Starting transfer cycle")
		s.simulateTransfer(int64(lastCycle.StartIndex), int64(lastCycle.EndIndex))
		// Mark cycle complete, update the indexes
		s.Logger.Info("Simulation cycle ending", "startIndex", lastCycle.StartIndex, "end", lastCycle.EndIndex)
		firstEmptyAccount, err := s.DB.GetFirstEmptyAccount()
		if err != nil {
			s.Logger.Error("Error fetching the first empty account", "error", err)
			return
		}
		startIndex, err := core.StringToUint(firstEmptyAccount.Path)
		if err != nil {
			s.Logger.Error("Error converting path to index", "err", err)
			return
		}
		s.DB.LogCycle(core.STAGE_TRANSFER, startIndex, startIndex+BATCH_SIZE)
	}
}

func (s *Simulator) simulateCreateAccounts(startIndex, endIndex int64) {
	for i := int64(0); i < BATCH_SIZE; i++ {
		txBytes, err := s.LoadedBazooka.EncodeCreateAccountTx(startIndex+i, 1)
		if err != nil {
			s.Logger.Error("unable to encode tx", "error", err)
			return
		}
		txCore := core.NewPendingTx(0, uint64(startIndex+i), core.TX_CREATE_ACCOUNT, "0x1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)
		err = s.DB.InsertTx(&txCore)
		if err != nil {
			s.Logger.Error("unable to insert tx", "error", err)
			return
		}
		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
	}
}

func (s *Simulator) simulateBurnConsent(startIndex, endIndex int64) {
	for i := int64(0); i < BATCH_SIZE; i++ {
		latestFromAcc, err := s.DB.GetAccountByIndex(uint64(startIndex + i))
		if err != nil {
			s.Logger.Error("unable to fetch latest account", "error", err)
			return
		}
		_, _, nonce, _, _, _, err := s.LoadedBazooka.DecodeAccount(latestFromAcc.Data)
		if err != nil {
			s.Logger.Error("unable to decode account", "error", err)
			return
		}
		txBytes, err := s.LoadedBazooka.EncodeBurnConsentTx(startIndex+i, BURN_AMOUNT, nonce.Int64()+1, core.TX_BURN_CONSENT)
		if err != nil {
			s.Logger.Error("unable to encode tx", "error", err)
			return
		}
		txCore := core.NewPendingTx(0, uint64(startIndex+i), core.TX_CREATE_ACCOUNT, "0x1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)
		err = s.DB.InsertTx(&txCore)
		if err != nil {
			s.Logger.Error("unable to insert tx", "error", err)
			return
		}
		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
	}
}

func (s *Simulator) simulateTransfer(startIndex, endIndex int64) {
	for i := int64(0); i < BATCH_SIZE; i++ {
		latestFromAcc, err := s.DB.GetAccountByIndex(uint64(startIndex + i))
		if err != nil {
			s.Logger.Error("unable to fetch latest account", "error", err)
			return
		}
		_, _, nonce, token, _, _, err := s.LoadedBazooka.DecodeAccount(latestFromAcc.Data)
		if err != nil {
			s.Logger.Error("unable to decode account", "error", err)
			return
		}
		txBytes, err := s.LoadedBazooka.EncodeTransferTx(startIndex+i, REDDIT_ACCOUNT, token.Int64(), nonce.Int64()+1, TRANSFER_AMOUNT, core.TX_TRANSFER_TYPE)
		if err != nil {
			s.Logger.Error("unable to encode tx", "error", err)
			return
		}
		txCore := core.NewPendingTx(uint64(startIndex+i), REDDIT_ACCOUNT, core.TX_TRANSFER_TYPE, "0x1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)
		err = s.DB.InsertTx(&txCore)
		if err != nil {
			s.Logger.Error("unable to insert tx", "error", err)
			return
		}
		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
	}
}

func (s *Simulator) simulateAirdrop(startIndex, endIndex int64) {
	var redditNonce int64
	latestFromAcc, err := s.DB.GetAccountByIndex(REDDIT_ACCOUNT)
	if err != nil {
		s.Logger.Error("unable to fetch latest account", "error", err)
		return
	}
	_, _, nonce, token, _, _, err := s.LoadedBazooka.DecodeAccount(latestFromAcc.Data)
	if err != nil {
		s.Logger.Error("unable to decode account", "error", err)
		return
	}
	redditNonce = nonce.Int64() + 1
	for i := int64(0); i < BATCH_SIZE; i++ {
		txBytes, err := s.LoadedBazooka.EncodeAirdropTx(REDDIT_ACCOUNT, startIndex+i, token.Int64(), redditNonce, TRANSFER_AMOUNT, core.TX_AIRDROP_TYPE)
		if err != nil {
			s.Logger.Error("unable to encode tx", "error", err)
			return
		}
		txCore := core.NewPendingTx(REDDIT_ACCOUNT, uint64(startIndex+i), core.TX_AIRDROP_TYPE, "0x1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)
		err = s.DB.InsertTx(&txCore)
		if err != nil {
			s.Logger.Error("unable to insert tx", "error", err)
			return
		}
		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
		redditNonce++
	}
}

func (s *Simulator) simulateBurnExec(startIndex, endIndex int64) {
	for i := int64(0); i < BATCH_SIZE; i++ {
		txBytes, err := s.LoadedBazooka.EncodeBurnExecTx(startIndex+i, core.TX_BURN_EXEC)
		if err != nil {
			s.Logger.Error("unable to encode tx", "error", err)
			return
		}
		txCore := core.NewPendingTx(uint64(startIndex+i), REDDIT_ACCOUNT, core.TX_BURN_EXEC, "0x1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)
		err = s.DB.InsertTx(&txCore)
		if err != nil {
			s.Logger.Error("unable to insert tx", "error", err)
			return
		}
		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)
	}
}
