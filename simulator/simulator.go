package simulator

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
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
	DefaultAmount       = 1
	DefaultFee          = 1
)

type UserList struct {
	Users []User `json:"users"`
}

type User struct {
	PublicKey string `json:"pubkey"`
	PrivKey   string `json:"privkey"`
}

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

	users UserList

	states map[uint64]User
}

// NewSimulator returns new simulator object
func NewSimulator(cfg config.Configuration, users UserList) *Simulator {
	logger := log.Logger.With("module", SimulatorService)
	simulator := &Simulator{}
	simulator.BaseService = *core.NewBaseService(logger, SimulatorService, simulator)

	// create DB obj
	DB, err := db.NewDB(cfg)
	if err != nil {
		panic(err)
	}

	// create bazooka obj
	bz, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		panic(err)
	}

	if len(users.Users) < 2 {
		panic(errors.New("simulator needs 2 or more than 2 state to function"))
	}

	simulator.Bazooka = bz
	simulator.DB = DB
	simulator.cfg = cfg
	simulator.users = users
	simulator.states = make(map[uint64]User)
	err = simulator.ParseUserList()
	if err != nil {
		panic(err)
	}
	return simulator
}

// ParseUserList parse user list to states
func (s *Simulator) ParseUserList() error {
	states := make(map[uint64]User)
	// var token uint64
	for _, u := range s.users.Users {
		pubkeyBz, err := hex.DecodeString(u.PublicKey)
		if err != nil {
			return err
		}
		account, err := s.DB.GetAccountByPubkey(pubkeyBz)
		if err != nil {
			return err
		}
		statesList, err := s.DB.GetStateByAccID(account.AccountID)
		if err != nil {
			return err
		}
		for _, statesInList := range statesList {
			if len(statesInList.Path) == 0 {
				continue
			}
			stateID, err := strconv.Atoi(statesInList.Path)
			if err != nil {
				return err
			}
			states[uint64(stateID)] = u
		}
	}
	s.states = states
	return nil
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

// AttemptTransfer attempts to make a transfer using one of the states
func (s *Simulator) AttemptTransfer() {
	defer s.wg.Done()
	// pick the first state in the states list
	for stateID, user := range s.states {
		state, err := s.DB.GetStateByIndex(stateID)
		if err != nil {
			return
		}

		_, bal, _, token, err := s.Bazooka.DecodeState(state.Data)
		if err != nil {
			return
		}

		// if balance is >1 use this account to make a transfer tx, else move on
		if bal.Int64() > 1 {
			// fetch pending nonce
			pendingNonce, err := s.DB.GetPendingNonce(stateID)
			if err != nil {
				return
			}

			receiverStateID, err := s.findReceiver(stateID)
			if err != nil {
				return
			}

			txData, err := s.Bazooka.EncodeTransferTx(int64(stateID), int64(receiverStateID), int64(DefaultFee), int64(pendingNonce+1), int64(DefaultAmount), core.TX_TRANSFER_TYPE)
			if err != nil {
				return
			}

			tx, err := core.NewPendingTx(txData, nil, stateID, pendingNonce+1, uint64(DefaultFee), token.Uint64(), core.TX_TRANSFER_TYPE)
			if err != nil {
				return
			}
			privBz, err := hex.DecodeString(user.PrivKey)
			if err != nil {
				return
			}
			pubkeyBz, err := hex.DecodeString(user.PublicKey)
			if err != nil {
				return
			}

			if err = signAndBroadcast(&s.Bazooka, &s.DB, tx, privBz, pubkeyBz); err != nil {
				return
			}

			return
		}
	}
}

func (s *Simulator) findReceiver(senderStateID uint64) (receiverStateID uint64, err error) {
	for stateID := range s.states {
		if stateID != senderStateID {
			return stateID, nil
		}
	}
	return 0, errors.New("no receiver state found")
}

func signAndBroadcast(b *bazooka.Bazooka, DBI *db.DB, tx core.Tx, priv, pub []byte) (err error) {
	txBytes, err := bazooka.GetSignBytes(*b, &tx)
	if err != nil {
		return
	}

	err = tx.SignTx(priv, pub, txBytes)
	if err != nil {
		return
	}
	err = tx.AssignHash()
	if err != nil {
		return
	}

	fmt.Println("Sending new tx", tx.String())
	err = DBI.InsertTx(&tx)
	if err != nil {
		return err
	}
	return nil
}
