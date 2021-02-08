package bidder

import (
	"context"
	"sync"
	"time"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/db"
	"github.com/BOPR/log"
)

var (
	BiddingService = "bidder"
)

type Bidder struct {
	// Base service
	core.BaseService

	// DB instance
	DB db.DB

	Bazooka bazooka.Bazooka

	cfg config.Configuration

	// header listener subscription
	cancelBidding context.CancelFunc

	// wait group
	wg sync.WaitGroup
}

// NewBidder returns new aggregator object
func NewBidder(cfg config.Configuration) *Bidder {
	// create logger
	logger := log.Logger.With("module", BiddingService)
	bi := &Bidder{}
	bi.BaseService = *core.NewBaseService(logger, BiddingService, bi)
	DB, err := db.NewDB(cfg)
	if err != nil {
		panic(err)
	}
	bz, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		panic(err)
	}
	bi.Bazooka = bz
	bi.DB = DB

	bi.cfg = cfg

	return bi
}

// OnStart starts new block subscription
func (bi *Bidder) OnStart() error {
	err := bi.BaseService.OnStart() // Always call the overridden method.
	if err != nil {
		return err
	}

	ctx, cancelBidding := context.WithCancel(context.Background())
	bi.cancelBidding = cancelBidding

	// start bidding for the next slots
	go bi.startBidding(ctx, bi.cfg.PollingInterval)
	return nil
}

// OnStop stops all necessary go routines
func (bi *Bidder) OnStop() {
	bi.BaseService.OnStop() // Always call the overridden method.
	bi.DB.Close()
	// cancel ack process
	bi.cancelBidding()
}

func (bi *Bidder) startBidding(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	// stop ticker when everything done
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// a.wg.Wait()
			// a.wg.Add(1)
			// go a.pickBatch()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// ShouldPropose checks whether we should propose or not
func (bi *Bidder) ShouldPropose() {

	// check the current bidable slot

	// check if the bid is below our bid

}
