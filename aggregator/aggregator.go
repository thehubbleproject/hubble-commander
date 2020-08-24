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
	"github.com/BOPR/router"
	"github.com/BOPR/wallet"
	"github.com/kilic/bn254/bls"
)

const (
	AggregatingService = "aggregator"
	COMMITMENT_SIZE    = 32
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

	// Router for all transactions
	router *router.Router

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
	a.ProcessAndSubmitBatch(txs)
}

func (a *Aggregator) ProcessAndSubmitBatch(txs []core.Tx) {
	a.Logger.Info("Processing new batch", "numberOfTxs", len(txs))

	// Step-2
	commitments, err := a.ProcessTx(txs)
	if err != nil {
		fmt.Println("Error while processing tx", "error", err)
		return
	}

	// Step-3
	// Submit all commitments on-chain
	a.LoadedBazooka.SubmitBatch(commitments)
	if err != nil {
		fmt.Println("Error while submitting batch", "error", err)
		return
	}
}

// ProcessTx fetches all the data required to validate tx from smart contact
// and calls the proccess tx function to return the updated balance root and accounts
func (a *Aggregator) ProcessTx(txs []core.Tx) (commitments []core.Commitment, err error) {
	if len(txs) == 0 {
		return commitments, errors.New("no tx to process,aborting")
	}
	var redditPDAProof core.PDAMerkleProof
	if (txs[0]).Type == core.TX_AIRDROP_TYPE {
		core.VerifierWaitGroup.Add(1)
		err = core.DBInstance.FetchPDAProofWithID(txs[0].Accounts[0], &redditPDAProof)
		if err != nil {
			return
		}
	}
	start := time.Now()
	for i, tx := range txs {
		a.Logger.Info("Processing transaction", "txNumber", i, "of", len(txs))
		updatedRoot, _, errWhileProcessing := a.router.ProcessTx(tx)
		if errWhileProcessing != nil {
			return
		}
		if i%32 == 0 {
			txInCommitment := txs[i : i+32]
			a.Logger.Info("Preparing a commitment", "NumOfTxs", len(txInCommitment), "type", txs[0].Type, "totalCommitmentsYet", len(commitments))
			aggregatedSig, err := aggregateSignatures(txInCommitment)
			if err != nil {
				return commitments, err
			}
			commitment := core.Commitment{Txs: txInCommitment, UpdatedRoot: updatedRoot, BatchType: tx.Type, AggregatedSignature: aggregatedSig.ToBytes()}
			commitments = append(commitments, commitment)
		}
		a.router.ApplyTx(tx)
	}
	elapsed := time.Since(start)
	log.Printf("Process batch took %s", elapsed)
	return commitments, nil
}

// generates aggregated signature for commitment
func aggregateSignatures(txs []core.Tx) (aggregatedSig bls.Signature, err error) {
	var signatures []*bls.Signature
	for _, tx := range txs {
		sig, err := wallet.BytesToSignature(tx.Signature)
		if err != nil {
			return aggregatedSig, err
		}
		signatures = append(signatures, &sig)
	}
	return wallet.NewAggregateSignature(signatures)
}
