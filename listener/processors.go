package listener

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"

	"github.com/BOPR/contracts/logger"
)

func (s *Syncer) processNewPubkeyAddition(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	// unpack event
	event := new(logger.LoggerPubkeyRegistered)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"accountID", event.PubkeyID.String(),
		"pubkey", event.Pubkey,
	)
	params, err := s.DBInstance.GetParams()
	if err != nil {
		return
	}

	// add new account in pending state to DB and
	pathToNode, err := core.SolidityPathToNodePath(event.PubkeyID.Uint64(), params.MaxDepth)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to convert path", err)
		panic(err)
	}

	pubkey := core.Pubkey(event.Pubkey)
	pubKeyStr, err := pubkey.String()
	if err != nil {
		return
	}
	newAcc, err := core.NewAccount(event.PubkeyID.Uint64(), pubKeyStr, pathToNode)
	if err != nil {
		fmt.Println("unable to create new account")
		panic(err)
	}

	if err := s.DBInstance.AddNewAccount(*newAcc); err != nil {
		panic(err)
	}
}

func (s *Syncer) processDepositQueued(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit found")
	// unpack event
	event := new(logger.LoggerDepositQueued)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}
	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"pubkeyID", event.PubkeyID.String(),
		"Amount", hex.EncodeToString(event.Data),
	)
	// add new account in pending state to DB and
	newAccount := core.NewPendingUserState(event.PubkeyID.Uint64(), event.Data)
	if err := s.DBInstance.AddNewPendingAccount(*newAccount); err != nil {
		panic(err)
	}
}

func (s *Syncer) processDepositSubtreeCreated(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit subtree created")
	// unpack event
	event := new(logger.LoggerDepositSubTreeReady)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}
	err = s.DBInstance.AttachDepositInfo(event.Root)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to attack deposit information:", err)
		panic(err)
	}

	// send deposit finalisation transction to ethereum chain
	catchingup, err := core.IsCatchingUp()
	if err != nil {
		panic(err)
	}

	if !catchingup {
		s.SendDepositFinalisationTx()
	} else {
		s.Logger.Info("Still cathing up, aborting deposit finalisation")
	}

}

func (s *Syncer) processDepositFinalised(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("Deposit batch finalised!")

	// unpack event
	event := new(logger.LoggerDepositsFinalised)

	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}

	depositRoot := core.ByteArray(event.DepositSubTreeRoot)
	pathToDepositSubTree := event.PathToSubTree

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"DepositSubTreeRoot", depositRoot.String(),
		"PathToDepositSubTreeInserted", pathToDepositSubTree.String(),
	)

	// TODO handle error
	newRoot, err := s.DBInstance.FinaliseDepositsAndAddBatch(depositRoot, pathToDepositSubTree.Uint64())
	if err != nil {
		fmt.Println("Error while finalising deposits", err)
	}

	fmt.Println("new root", newRoot)
}

func (s *Syncer) processNewBatch(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New batch found!")

	event := new(logger.LoggerNewBatch)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		s.Logger.Error("Unable to unpack log:", "error", err)
		panic(err)
	}

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"BatchNumber", event.Index.String(),
		"Type", event.BatchType,
		"Committer", event.Committer.String(),
		"TxHash", vLog.TxHash.String(),
	)

	// if the batch has some txs, parse them
	var txs []byte

	// pick the calldata for the batch
	txs, err = s.loadedBazooka.FetchBatchInputData(vLog.TxHash, event.BatchType)
	if err != nil {
		s.Logger.Error("Error fetching input data from tx", "error", err)
		return
	}

	// if we havent seen the batch, apply txs and store batch
	batch, err := s.DBInstance.GetBatchByIndex(event.Index.Uint64())
	if err != nil && gorm.IsRecordNotFoundError(err) {
		s.Logger.Info("Found a new batch, applying transactions and adding new batch", "index", event.Index.Uint64)
		newRoot, err := s.applyTxsFromBatch(txs, uint64(event.BatchType), true)
		if err != nil {
			s.Logger.Error("Error applying transactions from batch", "index", event.Index.String(), "error", err)
			return
		}

		newBatch := core.NewBatch(newRoot.String(), event.Committer.String(), vLog.TxHash.String(), uint64(event.BatchType), core.BATCH_COMMITTED)
		err = s.DBInstance.AddNewBatch(newBatch)
		if err != nil {
			s.Logger.Error("Error adding new batch to DB", "error", err)
			return
		}
		return
	} else if err != nil {
		s.Logger.Error("Unable to fetch batch", "index", event.Index, "err", err)
		return
	}

	// Mark seen batch as committed if we havent already
	if batch.Status != core.BATCH_COMMITTED {
		s.Logger.Info("Found a non committed batch")
		err = s.DBInstance.CommitBatch(event.Index.Uint64())
		if err != nil {
			s.Logger.Error("Unable to commit batch", "index", event.Index.String(), "err", err)
			return
		}
	}
}

func (s *Syncer) processRegisteredToken(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New token registered")
	event := new(logger.LoggerRegisteredToken)

	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}
	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"TokenAddress", event.TokenContract.String(),
		"TokenID", event.TokenType,
	)
	newToken := core.Token{TokenID: event.TokenType.Uint64(), Address: event.TokenContract.String()}
	if err := s.DBInstance.AddToken(newToken); err != nil {
		panic(err)
	}
}

func (s *Syncer) SendDepositFinalisationTx() {
	params, err := s.DBInstance.GetParams()
	if err != nil {
		return
	}
	nodeToBeReplaced, siblings, err := s.DBInstance.GetDepositNodeAndSiblings()
	if err != nil {
		return
	}

	err = s.loadedBazooka.FireDepositFinalisation(nodeToBeReplaced, siblings, params.MaxDepositSubTreeHeight)
	if err != nil {
		return
	}
}

func (s *Syncer) applyTxsFromBatch(txsBytes []byte, txType uint64, isSyncing bool) (newRoot core.ByteArray, err error) {
	// check if the batch has any txs
	if len(txsBytes) == 0 {
		s.Logger.Info("No txs to apply")
		return newRoot, nil
	}
	var transactions []core.Tx

	switch txType {
	case core.TX_TRANSFER_TYPE:
		transactions, err = s.decompressTransfers(txsBytes)
		if err != nil {
			return newRoot, err
		}
	default:
		fmt.Println("TxType didnt match any options", txType)
		return newRoot, errors.New("Didn't match any options")
	}

	commitments, err := core.ProcessTxs(s.DBInstance, s.loadedBazooka, transactions, isSyncing)
	if err != nil {
		return newRoot, err
	}

	return commitments[len(commitments)-1].UpdatedRoot, nil
}

// decompressTransfers decompresses transfer bytes to TX
func (s *Syncer) decompressTransfers(decompressedTxs []byte) (txs []core.Tx, err error) {
	froms, tos, amounts, fees, err := s.loadedBazooka.DecompressTransferTxs(decompressedTxs)
	if err != nil {
		return
	}
	var transactions []core.Tx

	for i := 0; i < len(froms); i++ {
		fromState, err := s.DBInstance.GetStateByIndex(froms[i].Uint64())
		if err != nil {
			return transactions, err
		}
		_, _, nonce, _, err := s.loadedBazooka.DecodeState(fromState.Data)
		if err != nil {
			return transactions, err
		}

		txData, err := s.loadedBazooka.EncodeTransferTx(froms[i].Int64(), tos[i].Int64(), fees[i].Int64(), nonce.Int64(), amounts[i].Int64(), int64(core.TX_TRANSFER_TYPE))
		if err != nil {
			return transactions, err
		}

		newTx := core.NewTx(froms[i].Uint64(), tos[i].Uint64(), core.TX_TRANSFER_TYPE, nil, txData)
		transactions = append(transactions, newTx)
	}

	return transactions, nil
}
