package simulator

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/db"
	"github.com/BOPR/log"
)

var (
	SimulatorService    = "simulator"
	ErrIncorrectTxCount = errors.New("inaccurate number of transactions")
)

type Simulator struct {
	// Base service
	core.BaseService

	// DB instance
	DB db.DB

	// Bazooka instance for simlator
	Bazooka bazooka.Bazooka

	// configration
	cfg config.Configuration

	// header listener subscription
	cancelSimulator context.CancelFunc

	// wait group
	wg sync.WaitGroup
}

// NewSimulator returns new simulator object
func NewSimulator(cfg config.Configuration) *Simulator {
	// create logger
	logger := log.Logger.With("module", SimulatorService)
	simulator := &Simulator{}
	simulator.BaseService = *core.NewBaseService(logger, SimulatorService, simulator)
	DB, err := db.NewDB(cfg)
	if err != nil {
		panic(err)
	}
	bz, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		panic(err)
	}

	simulator.Bazooka = bz
	simulator.DB = DB
	simulator.cfg = cfg

	return simulator
}

// OnStart starts new block subscription
func (s *Simulator) OnStart() error {
	err := s.BaseService.OnStart() // Always call the overridden method.
	if err != nil {
		return err
	}
	ctx, cancelAggregating := context.WithCancel(context.Background())
	s.cancelSimulator = cancelAggregating
	// start polling for checkpoint in buffer
	go s.startAggregating(ctx, s.cfg.PollingInterval)
	return nil
}

// OnStop stops all necessary go routines
func (s *Simulator) OnStop() {
	s.BaseService.OnStop() // Always call the overridden method.
	s.DB.Close()
	s.cancelSimulator()
}

func (s *Simulator) startAggregating(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	// stop ticker when everything done
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			s.wg.Wait()
			s.wg.Add(1)
			go s.AttemptTransfer()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (s *Simulator) AttemptTransfer() {

}
