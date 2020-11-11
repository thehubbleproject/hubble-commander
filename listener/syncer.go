package listener

import (
	"context"
	"encoding/hex"
	"sync"
	"time"

	"github.com/BOPR/common"
	"github.com/BOPR/config"

	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// Syncer to sync events from ethereum chain
type Syncer struct {
	// Base service
	core.BaseService

	// ABIs
	abis []abi.ABI

	// storage client
	DBInstance core.DB

	// contract caller to interact with contracts
	loadedBazooka core.Bazooka

	// header channel
	HeaderChannel chan *ethTypes.Header
	// cancel function for poll/subscription
	cancelSubscription context.CancelFunc

	// header listener subscription
	cancelHeaderProcess context.CancelFunc

	// wait group
	wg sync.WaitGroup
}

func NewSyncer() Syncer {
	// create logger
	logger := common.Logger.With("module", SyncerServiceName)

	// create syncer obj
	syncerService := &Syncer{}

	// create new base service
	syncerService.BaseService = *core.NewBaseService(logger, SyncerServiceName, syncerService)
	loadedBazooka, err := core.NewPreLoadedBazooka()
	if err != nil {
		panic(err)
	}
	var abis []abi.ABI
	abis = append(abis, loadedBazooka.ContractABI[common.LOGGER_KEY])

	// abis for all the events
	syncerService.abis = abis
	syncerService.loadedBazooka = loadedBazooka
	syncerService.HeaderChannel = make(chan *ethTypes.Header)
	syncerService.DBInstance, err = core.NewDB()
	if err != nil {
		panic(err)
	}
	//nolint:govet // will fix later in #76
	return *syncerService
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
		go s.startPolling(ctx, config.GlobalCfg.PollingInterval)
	} else {
		// start go routine to listen new header using subscription
		go s.startSubscription(ctx, subscription)
	}
	s.Logger.Info("Starting syncer", "LoggingContract", config.GlobalCfg.LoggerAddress)
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

	// Setup the ticket and the channel to signal
	// the ending of the interval
	ticker := time.NewTicker(interval)

	// start listening
	for {
		select {
		case <-ticker.C:
			s.Logger.Info("Searching for new logs...")
			header, err := s.loadedBazooka.EthClient.HeaderByNumber(ctx, nil)
			if err == nil && header != nil {
				// send data to channel
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
			// stop service
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
	s.Logger.Info("Sync status", "LastLogIndexed", syncStatus.LastEthBlockBigInt().String())
	if header.Number.Uint64() <= syncStatus.LastEthBlockBigInt().Uint64() {
		s.Logger.Error("No need to sync more events", "currentEthBlock", header.Number.String(), "lastSyncedBlock", syncStatus.LastEthBlockBigInt().String())
		return
	}

	// we need to filter only by logger contracts
	// since all events are emitted by it
	query := ethereum.FilterQuery{
		FromBlock: syncStatus.LastEthBlockBigInt(),
		ToBlock:   header.Number,
		Addresses: []ethCmn.Address{
			ethCmn.HexToAddress(config.GlobalCfg.LoggerAddress),
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
	for _, vLog := range logs {
		topic := vLog.Topics[0].Bytes()
		for _, abiObject := range s.abis {
			selectedEvent := EventByID(&abiObject, topic)
			if selectedEvent != nil {
				s.Logger.Debug("Found an event", "name", selectedEvent.Name)
				switch selectedEvent.Name {
				case "RegisteredToken":
					s.processRegisteredToken(selectedEvent.Name, &abiObject, &vLog)
				case "NewBatch":
					s.processNewBatch(selectedEvent.Name, &abiObject, &vLog)
				case "NewPubkeyAdded":
					s.processNewPubkeyAddition(selectedEvent.Name, &abiObject, &vLog)
				case "DepositQueued":
					s.processDepositQueued(selectedEvent.Name, &abiObject, &vLog)
				case "DepositSubTreeReady":
					s.processDepositSubtreeCreated(selectedEvent.Name, &abiObject, &vLog)
				case "DepositsFinalised":
					s.processDepositFinalised(selectedEvent.Name, &abiObject, &vLog)
				default:
					s.Logger.Debug("Unable to match with any event", "event", selectedEvent.Name)
				}
			} else {
				s.Logger.Debug("Unable to match with any event", "topics", hex.EncodeToString(topic))
			}
		}
	}
	err := s.DBInstance.UpdateSyncStatusWithBlockNumber(header.Number.Uint64())
	if err != nil {
		s.Logger.Error("Unable to update listener log", "error", err)
	}
}
