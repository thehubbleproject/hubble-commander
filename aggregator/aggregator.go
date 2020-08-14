package aggregator

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
)

const (
	AggregatingService = "aggregator"
)

// Aggregator is the service which is supposed to create batches
// It has the following tasks:
// 1. Pick txs from the mempool
// 2. Validate these trnsactions
// 3. Update the DB post running each tx
// 4. Finally create a batch of all the transactions and post on-chain
type Aggregator struct {
	// Base service
	core.BaseService

	// contract caller to interact with contracts
	LoadedBazooka core.Bazooka

	// DB instance
	DB core.DB

	// header listener subscription
	cancelAggregating context.CancelFunc

	// wait group
	wg sync.WaitGroup
}

// NewAggregator returns new aggregator object
func NewAggregator() *Aggregator {
	// create logger
	logger := common.Logger.With("module", AggregatingService)
	LoadedBazooka, err := core.NewPreLoadedBazooka()
	if err != nil {
		panic(err)
	}
	aggregator := &Aggregator{}
	aggregator.BaseService = *core.NewBaseService(logger, AggregatingService, aggregator)
	DB, err := core.NewDB()
	if err != nil {
		panic(err)
	}
	aggregator.DB = DB
	aggregator.LoadedBazooka = LoadedBazooka
	return aggregator
}

// OnStart starts new block subscription
func (a *Aggregator) OnStart() error {
	a.BaseService.OnStart() // Always call the overridden method.

	ctx, cancelAggregating := context.WithCancel(context.Background())
	a.cancelAggregating = cancelAggregating

	// start polling for checkpoint in buffer
	go a.startAggregating(ctx, config.GlobalCfg.PollingInterval)
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
			a.wg.Wait()
			a.wg.Add(1)
			go a.pickBatch()
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
	a.Logger.Info("Processing new batch", "numberOfTxs", len(txs))

	// Step-2
	err = a.ProcessTx(txs)
	if err != nil {
		fmt.Println("Error while processing tx", "error", err)
		return
	}

	// Step-3
	// Finally create a merkel root of all updated leafs and push batch on-chain
	rootAcc, err := a.DB.GetRoot()
	if err != nil {
		fmt.Println("Error while getting root", "error", err)
		return
	}

	a.LoadedBazooka.SubmitBatch(rootAcc.HashToByteArray(), txs)
	if err != nil {
		fmt.Println("Error while submitting batch", "error", err)
		return
	}
}

// ProcessTx fetches all the data required to validate tx from smart contact
// and calls the proccess tx function to return the updated balance root and accounts
func (a *Aggregator) ProcessTx(txs []core.Tx) error {
	if len(txs) == 0 {
		return errors.New("no tx to process,aborting!")
	}

	var redditPDAProof core.PDAMerkleProof
	if (txs[0]).Type == core.TX_AIRDROP_TYPE {
		core.VerifierWaitGroup.Add(1)
		err := core.DBInstance.FetchPDAProofWithID(txs[0].From, &redditPDAProof)
		if err != nil {
			return err
		}
	}

	start := time.Now()
	for _, tx := range txs {
		rootAcc, err := a.DB.GetRoot()
		if err != nil {
			return err
		}
		a.Logger.Debug("Latest root", "root", rootAcc.Hash)
		currentRoot, err := core.HexToByteArray(rootAcc.Hash)
		if err != nil {
			return err
		}
		pdaRoot, err := a.DB.GetPDARoot()
		if err != nil {
			return err
		}
		currentAccountTreeRoot := pdaRoot.HashToByteArray()
		fromAccProof, toAccProof, PDAproof, txDBConn, err := tx.GetVerificationData()
		if err != nil {
			a.Logger.Error("Unable to create verification data", "error", err)
			return err
		}
		if (txs[0]).Type == core.TX_AIRDROP_TYPE {
			core.VerifierWaitGroup.Wait()
			PDAproof = redditPDAProof
		}
		updatedRoot, _, updatedTo, err := a.LoadedBazooka.ProcessTx(currentRoot, currentAccountTreeRoot, tx, fromAccProof, toAccProof, PDAproof)
		if err != nil {
			a.Logger.Error("Error processing tx", "tx", tx.String(), "error", err)
			if txDBConn.Instance != nil {
				txDBConn.Instance.Rollback()
				txDBConn.Close()
			}
			return err
		} else {
			if txDBConn.Instance != nil {
				txDBConn.Instance.Commit()
				txDBConn.Close()
			}
		}
		switch txType := tx.Type; txType {
		case core.TX_TRANSFER_TYPE:
			tx.ApplySingleTx(toAccProof.Account, updatedTo)
		case core.TX_AIRDROP_TYPE:
			tx.ApplySingleTx(toAccProof.Account, updatedTo)
		case core.TX_CREATE_ACCOUNT:
			tx.ApplySingleTx(toAccProof.Account, updatedTo)
		case core.TX_BURN_CONSENT:
			fmt.Println("burnconsent")
		case core.TX_BURN_EXEC:
			fmt.Println("burn exec")
		}
		currentRoot = updatedRoot
	}
	elapsed := time.Since(start)
	log.Printf("Process batch took %s", elapsed)
	return nil
}
