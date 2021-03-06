package listener

import (
	"context"
	"encoding/hex"
	"strings"
	"sync"
	"time"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/contracts/accountregistry"
	"github.com/BOPR/contracts/burnauction"
	"github.com/BOPR/contracts/depositmanager"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/contracts/tokenregistry"
	"github.com/BOPR/db"
	"github.com/BOPR/log"

	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	syncerServiceName = "syncer"
)

// Syncer to sync events from ethereum chain
type Syncer struct {
	// Base service
	core.BaseService

	// ABIs
	abis []abi.ABI

	// storage client
	DBInstance db.DB

	// contract caller to interact with contracts
	loadedBazooka bazooka.Bazooka

	// local config
	cfg config.Configuration

	// header channel
	HeaderChannel chan *ethTypes.Header
	// cancel function for poll/subscription
	cancelSubscription context.CancelFunc

	// header listener subscription
	cancelHeaderProcess context.CancelFunc

	// wait group
	wg sync.WaitGroup
}

// NewSyncer creates a new syncer object
func NewSyncer(cfg config.Configuration) *Syncer {
	// create logger
	logger := log.Logger.With("module", syncerServiceName)

	// create syncer obj
	syncerService := &Syncer{}

	// create new base service
	syncerService.BaseService = *core.NewBaseService(logger, syncerServiceName, syncerService)
	loadedBazooka, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		panic(err)
	}

	var abis []abi.ABI

	rollupABI, err := abi.JSON(strings.NewReader(rollup.RollupABI))
	common.PanicIfError(err)
	depositManagerABI, err := abi.JSON(strings.NewReader(depositmanager.DepositmanagerABI))
	common.PanicIfError(err)
	trABI, err := abi.JSON(strings.NewReader(tokenregistry.TokenregistryABI))
	common.PanicIfError(err)
	baABI, err := abi.JSON(strings.NewReader(burnauction.BurnauctionABI))
	common.PanicIfError(err)
	arABI, err := abi.JSON(strings.NewReader(accountregistry.AccountregistryABI))
	common.PanicIfError(err)

	abis = append(abis, rollupABI)
	abis = append(abis, depositManagerABI)
	abis = append(abis, trABI)
	abis = append(abis, baABI)
	abis = append(abis, arABI)

	// abis for all the events
	syncerService.abis = abis
	syncerService.loadedBazooka = loadedBazooka
	syncerService.HeaderChannel = make(chan *ethTypes.Header)
	syncerService.DBInstance, err = db.NewDB(cfg)
	if err != nil {
		panic(err)
	}

	syncerService.cfg = cfg

	//nolint:govet // will fix later in #76
	return syncerService
}

// OnStart starts new block subscription
func (s *Syncer) OnStart() error {
	// Always call the overridden method.
	err := s.BaseService.OnStart()
	if err != nil {
		return err
	}

	// create cancellable context
	ctx, cancelSubscription := context.WithCancel(context.Background())
	s.cancelSubscription = cancelSubscription

	// create cancellable context
	headerCtx, cancelHeaderProcess := context.WithCancel(context.Background())
	s.cancelHeaderProcess = cancelHeaderProcess

	// start header process
	go s.startHeaderProcess(headerCtx)

	// subscribe to new head
	subscription, err := s.loadedBazooka.EthClient.SubscribeNewHead(ctx, s.HeaderChannel)
	if err != nil {
		// start go routine to poll for new header using client object
		go s.startPolling(ctx, s.cfg.PollingInterval)
	} else {
		// start go routine to listen new header using subscription
		go s.startSubscription(ctx, subscription)
	}

	s.Logger.Info("Starting syncer")

	return nil
}

// OnStop stops all necessary go routines
func (s *Syncer) OnStop() {

	s.BaseService.OnStop() // Always call the overridden method.

	// cancel subscription if any
	s.cancelSubscription()

	// cancel header process
	s.cancelHeaderProcess()

	s.DBInstance.Close()
}

// startHeaderProcess starts header process when they get new header
func (s *Syncer) startHeaderProcess(ctx context.Context) {
	for {
		select {
		case newHeader := <-s.HeaderChannel:
			s.wg.Wait()
			s.processHeader(*newHeader)
		case <-ctx.Done():
			return
		}
	}
}

// startPolling starts polling
func (s *Syncer) startPolling(ctx context.Context, pollInterval time.Duration) {
	// How often to fire the passed in function in second
	interval := pollInterval

	// Setup the ticker and the channel to signal
	// the ending of the interval
	ticker := time.NewTicker(interval)

	// start listening
	for {
		select {
		case <-ticker.C:
			s.Logger.Info("Searching for new logs...")
			syncStatus, err := s.DBInstance.GetSyncStatus()
			if err != nil {
				s.Logger.Error("Unable to fetch listener log", "error", err)
				return
			}

			header, err := s.loadedBazooka.GetEthBlock(nil)
			if err != nil {
				s.Logger.Error("Error fetching latest blocks", "error", err)
				return
			}

			if header.Number.Uint64()-syncStatus.LastEthBlockRecorded >= s.cfg.ConfirmationBlocks {
				s.HeaderChannel <- header
			}

		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (s *Syncer) startSubscription(ctx context.Context, subscription ethereum.Subscription) {
	for {
		select {
		case err := <-subscription.Err():
			s.Logger.Error("Error while subscribing new blocks", "error", err)
			err = s.Stop()
			s.Logger.Error("Error while stopping", "error", err)
			// cancel subscription
			s.cancelSubscription()
			return
		case <-ctx.Done():
			return
		}
	}
}

func (s *Syncer) processHeader(header ethTypes.Header) {
	syncStatus, err := s.DBInstance.GetSyncStatus()
	if err != nil {
		s.Logger.Error("Unable to fetch listener log", "error", err)
		return
	}

	// to make sure we dont run the same events multiple times
	lastBlockSynced := syncStatus.LastEthBlockBigInt()
	if header.Number.Uint64() <= lastBlockSynced.Uint64() {
		s.Logger.Info("Everything synced, watching for more events", "lastSyncedBlock", lastBlockSynced, "currentHeader", header)
		return
	}

	// we need to filter only by logger contracts
	// since all events are emitted by it
	query := ethereum.FilterQuery{
		FromBlock: syncStatus.LastEthBlockBigInt(),
		ToBlock:   header.Number,
		Addresses: []ethCmn.Address{
			ethCmn.HexToAddress(s.cfg.RollupAddress),
			ethCmn.HexToAddress(s.cfg.TokenRegistry),
			ethCmn.HexToAddress(s.cfg.AccountRegistry),
			ethCmn.HexToAddress(s.cfg.DepositManager),
			ethCmn.HexToAddress(s.cfg.BurnAuction),
		},
		Topics: [][]ethCmn.Hash{},
	}

	// get all logs
	logs, err := s.loadedBazooka.EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		s.Logger.Error("Error while filtering logs from syncer", "error", err)
		return
	} else if len(logs) > 0 {
		s.Logger.Debug("New logs found", "numberOfLogs", len(logs))
	}

	s.wg.Add(1)
	go s.processEvents(logs, header)
}

func (s *Syncer) processEvents(logs []ethTypes.Log, header ethTypes.Header) {
	defer s.wg.Done()
	var err error

	lastBlockNum := header.Number.Uint64()

	for _, vLog := range logs {
		topic := vLog.Topics[0].Bytes()
		for i := 0; i < len(s.abis); i++ {
			abiObject := s.abis[i]
			selectedEvent := EventByID(&abiObject, topic)
			if selectedEvent != nil {
				s.Logger.Debug("Found an event", "name", selectedEvent.Name, "topic", hex.EncodeToString(topic), "blockNum", vLog.BlockNumber)
				switch selectedEvent.Name {
				case "NewBatch":
					err = s.processNewBatch(selectedEvent.Name, &abiObject, &vLog)
					if err != nil {
						s.Logger.Error("Error processign new bathc", err)
						break
					}
				case "PubkeyRegistered":
					err = s.processNewPubkeyAddition(selectedEvent.Name, &abiObject, &vLog)
					if err != nil {
						s.Logger.Error("Error processing pubkey reqis", err)
						break
					}
				case "DepositQueued":
					err = s.processDepositQueued(selectedEvent.Name, &abiObject, &vLog)
					if err != nil {
						s.Logger.Error("Error processing deposit queue", err)
						break
					}
				case "DepositSubTreeReady":
					err = s.processDepositSubtreeCreated(selectedEvent.Name, &abiObject, &vLog)
					if err != nil {
						s.Logger.Error("Error processing subtree event", err)
						break
					}
				case "DepositsFinalised":
					err = s.processDepositFinalised(selectedEvent.Name, &abiObject, &vLog)
					if err != nil {
						s.Logger.Error("Error finalising deposit", err)
						break
					}
				default:
					s.Logger.Debug("Unable to match with any event", "event", selectedEvent.Name)
				}
			} else {
				s.Logger.Debug("Unable to match with any event", "topics", hex.EncodeToString(topic))
			}
		}
	}
	if err != nil {
		s.Logger.Info("Error processing event", "err", err)
		return
	}

	// Update sync status with block num
	err = s.DBInstance.UpdateSyncStatusWithBlockNumber(lastBlockNum)
	if err != nil {
		s.Logger.Error("Unable to update listener log", "error", err)
		return
	}
}
