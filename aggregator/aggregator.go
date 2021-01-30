package aggregator

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/db"
	"github.com/BOPR/log"
)

var (
	AggregatingService       = "aggregator"
	ErrIncorrectTxCount      = errors.New("inaccurate number of transactions")
	ErrNotEnoughTransactions = errors.New("not enough transactions")
)

// Aggregator is the service which is supposed to create batches
// It has the following tasks:
// 1. Pick txs from the mempoolinaccurate number of transactione
// 2. Validate these trnsactions
// 3. Update the DB post running each tx
// 4. Finally create a batch of all the transactions and post on-chain
type Aggregator struct {
	// Base service
	core.BaseService

	// DB instance
	DB db.DB

	Bazooka bazooka.Bazooka

	cfg config.Configuration

	// header listener subscription
	cancelAggregating context.CancelFunc

	// wait group
	wg sync.WaitGroup
}

// NewAggregator returns new aggregator object
func NewAggregator(cfg config.Configuration) *Aggregator {
	// create logger
	logger := log.Logger.With("module", AggregatingService)
	aggregator := &Aggregator{}
	aggregator.BaseService = *core.NewBaseService(logger, AggregatingService, aggregator)
	DB, err := db.NewDB(cfg)
	if err != nil {
		panic(err)
	}
	bz, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		panic(err)
	}
	aggregator.Bazooka = bz
	aggregator.DB = DB

	aggregator.cfg = cfg

	return aggregator
}

// OnStart starts new block subscription
func (a *Aggregator) OnStart() error {
	err := a.BaseService.OnStart() // Always call the overridden method.
	if err != nil {
		return err
	}

	ctx, cancelAggregating := context.WithCancel(context.Background())
	a.cancelAggregating = cancelAggregating

	// start polling for checkpoint in buffer
	go a.startAggregating(ctx, a.cfg.PollingInterval)
	return nil
}

// OnStop stops all necessary go routines
func (a *Aggregator) OnStop() {
	a.BaseService.OnStop() // Always call the overridden method.
	a.DB.Close()
	// cancel ack process
	a.cancelAggregating()
}

func (a *Aggregator) startAggregating(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	// stop ticker when everything done
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if isCatchingUp, err := IsCatchingUp(a.Bazooka, a.DB); err != nil {
				return
			} else if isCatchingUp {
				a.Logger.Info("Commander catching up, aborting aggregation till next poll")
			} else {
				a.wg.Wait()
				a.wg.Add(1)
				go a.pickBatch()
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (a *Aggregator) pickBatch() {
	defer a.wg.Done()
	txs, err := a.DB.PopTxs()
	if err != nil {
		fmt.Println("Error while popping txs from mempool", "Error", err)
	}
	a.processAndSubmitBatch(txs)
}

func (a *Aggregator) processAndSubmitBatch(txs []core.Tx) {
	a.Logger.Info("Processing new batch", "numberOfTxs", len(txs))

	// process all transactions
	commitments, err := a.processTxs(txs)
	if err == ErrNotEnoughTransactions {
		a.Logger.Info("Not enough transactions for a commitment")
		return
	}
	if err != nil {
		a.Logger.Error("Error processing tx", "error", err)
		return
	}

	// fetch account tree root
	accountTreeRoot, err := a.DB.GetAccountRootHash()
	if err != nil {
		a.Logger.Info("Error fetching account root", "error", err)
		return
	}

	// Submit all commitments on-chain
	txHash, commitments, err := a.Bazooka.SubmitBatch(commitments, accountTreeRoot.String())
	if err != nil {
		a.Logger.Error("Error while submitting batch", "error", err)
		return
	}

	// last commitment
	lastCommitment := commitments[len(commitments)-1]

	// record batch locally
	newBatch := core.NewBatch(a.cfg.OperatorAddress, txHash, lastCommitment.BatchType, core.BATCH_BROADCASTED)
	batchID, err := a.DB.AddNewBatch(newBatch, commitments)
	if err != nil {
		a.Logger.Error("Error adding new batch", "error", err)
		return
	}

	a.Logger.Info("Added new batch to DB", "ID", batchID, "numOfCommitments", len(commitments))
}

func (a *Aggregator) processTxs(txs []core.Tx) (commitments []core.Commitment, err error) {
	var txsInCommitment []int

	if len(txs) < int(a.cfg.TxsPerCommitment) {
		return commitments, ErrNotEnoughTransactions
	}

	// if number of transactions divisible into TxsPerCommitment batches
	// allows sending smaller number of transactions
	if len(txs)%int(a.cfg.TxsPerCommitment) != 0 {
		return commitments, ErrIncorrectTxCount
	}

	// calculates number of transactions per commitment
	for i := range txs {
		if i%int(a.cfg.TxsPerCommitment) == 0 {
			txsInCommitment = append(txsInCommitment, int(a.cfg.TxsPerCommitment))
		}
	}

	return db.ProcessTxs(&a.Bazooka, &a.DB, txs, txsInCommitment, false)
}

// IsCatchingUp returns true/false according to the sync status of the node
func IsCatchingUp(b bazooka.Bazooka, db db.DB) (bool, error) {
	totalBatches, err := b.TotalBatches()
	if err != nil {
		return false, err
	}

	totalBatchedStored, err := db.GetBatchCount()
	if err != nil {
		return false, err
	}

	// if total batchse are greater than what we recorded we are still catching up
	if totalBatches > uint64(totalBatchedStored) {
		return true, err
	}

	return false, nil
}
