package bidder

import (
	"context"
	"sync"
	"time"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/log"
	ethCmn "github.com/ethereum/go-ethereum/common"
)

var (
	BiddingService = "bidder"
)

type CoordinatorInfo struct {
	DepositAmount uint64
	Address       ethCmn.Address
}

type SlotInfo struct {
	CurrentSlot uint64
}

type Bidder struct {
	// Base service
	core.BaseService

	bz bazooka.Bazooka

	cfg config.Configuration

	bidderInfo CoordinatorInfo

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
	bz, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		panic(err)
	}
	bi.bz = bz
	bi.cfg = cfg

	bi.bidderInfo.Address = ethCmn.HexToAddress(cfg.OperatorAddress)
	bi.bidderInfo.DepositAmount = cfg.MinDeposit

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
			shouldPropose, err := bi.ShouldPropose()
			if err != nil {
				bi.Logger.Error("Could not determine bidding action", "err", err)
				return
			}

			bi.Logger.Info("Bidding action decided", "shouldPropose", shouldPropose)
			if !shouldPropose {
				return
			}

			bi.wg.Wait()
			bi.wg.Add(1)
			go bi.Bid()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// ShouldPropose checks whether we should propose or not
func (bi *Bidder) ShouldPropose() (ok bool, err error) {
	bi.Logger.Info("Deciding if we should propose")

	depositSize, err := bi.bz.GetDeposit(bi.bidderInfo.Address)
	if err != nil {
		bi.Logger.Error("Error getting deposit_size", "error", err)
		return
	}

	bi.bidderInfo.DepositAmount = depositSize
	bi.Logger.Debug("Current amount deposited on-chain", "amount", depositSize)

	// check the current bidable slot
	slotOnAuction, err := bi.bz.GetBidableSlot()
	if err != nil {
		bi.Logger.Error("Error getting bidable slot", "error", err)
		return
	}

	bi.Logger.Debug("Current slot on auction", "slotNum", slotOnAuction)

	proposerAddr, bidAmount, isInit, err := bi.bz.GetCurrentBidForSlot(slotOnAuction)
	if err != nil {
		bi.Logger.Error("Error getting current bid for slot", "error", err)
		return
	}

	// if no one has bid yet, we can be the first ones
	if !isInit {
		bi.Logger.Info("No one has bid yet, bidding...", "proposer", bi.bidderInfo.Address)
		return true, nil
	}

	// we are the highest bidders, no need to bid
	if proposerAddr == bi.bidderInfo.Address {
		bi.Logger.Info("We are the highest bidder", "amount", bidAmount)
		return false, nil
	}

	// we dont bid if the current bid amount is higher or equal than the config set
	if bidAmount >= bi.cfg.BidAmount {
		bi.Logger.Info("We are outbidded", "currentWinner", proposerAddr.String(), "amount", bidAmount, "ourBid", bi.cfg.BidAmount)
		return false, nil
	}

	// this means all the conditions are set for us to bid, lets bid!
	return true, nil
}

// Bid bids the bid amount for the coordinator on-chain
func (bi *Bidder) Bid() {
	bi.Logger.Info("Attempting to bid on slot!")
	defer bi.wg.Done()

	// check if we have enough minimmum amount deposited on-chain
	if bi.bidderInfo.DepositAmount <= bi.cfg.MinDeposit {
		txHash, errI := bi.bz.DepositForAuction(int64(bi.cfg.MinDeposit), bi.bidderInfo.Address)
		if errI != nil {
			return
		}

		bi.Logger.Info("We did not have enough ether deposited, sent new deposit", "txHash", txHash)

		// return wtih no error, we bid on next poll
		return
	}

	// bid for the slot
	txHash, err := bi.bz.Bid(bi.cfg.BidAmount)
	if err != nil {
		return
	}

	bi.Logger.Info("Sent new bid!", "txHash", txHash, "amount", bi.cfg.BidAmount, "proposer", bi.bidderInfo.Address)
}
