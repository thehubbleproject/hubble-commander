package listener

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/common"
	"github.com/BOPR/contracts/accountregistry"
	"github.com/BOPR/contracts/depositmanager"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/contracts/tokenregistry"
	"github.com/BOPR/core"
	"github.com/BOPR/db"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
)

func (s *Syncer) processNewPubkeyAddition(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	// unpack event
	event := new(accountregistry.AccountregistryPubkeyRegistered)
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

	pubkey := core.NewPubkey(event.Pubkey)
	newAcc, err := core.NewAccount(event.PubkeyID.Uint64(), pubkey, pathToNode)
	if err != nil {
		fmt.Println("unable to create new account")
		panic(err)
	}

	if err := s.DBInstance.AddNewAccount(*newAcc); err != nil {
		panic(err)
	}

	// if pubkey was added by relayer mark the packet processed
	if err := s.DBInstance.MarkPacketDone(pubkey); err != nil {
		panic(err)
	}
}

func (s *Syncer) processDepositQueued(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit found")
	// unpack event
	event := new(depositmanager.DepositmanagerDepositQueued)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}
	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"pubkeyID", event.PubkeyID.String(),
		"Data", hex.EncodeToString(event.Data),
	)
	// add new account in pending state to DB and
	newAccount := core.NewPendingUserState(event.PubkeyID.Uint64(), event.Data)
	err = s.DBInstance.AddNewPendingUserState(*newAccount)
	if err != nil {
		panic(err)
	}
}

func (s *Syncer) processDepositSubtreeCreated(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit subtree created")
	// unpack event
	event := new(depositmanager.DepositmanagerDepositSubTreeReady)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}

	// send deposit finalisation transction to ethereum chain
	catchingup, err := IsCatchingUp(s.loadedBazooka, s.DBInstance)
	if err != nil {
		panic(err)
	}
	if !catchingup {
		s.sendDepositFinalisationTx()
	} else {
		s.Logger.Info("Still cathing up, aborting deposit finalisation")
	}

	err = s.DBInstance.AttachDepositInfo(event.Root)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to attack deposit information:", err)
		panic(err)
	}
}

func (s *Syncer) processDepositFinalised(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("Deposit batch finalised!")

	// unpack event
	event := new(rollup.RollupDepositsFinalised)

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

	event := new(rollup.RollupNewBatch)
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

	// if we havent seen the batch, apply txs and store batch
	batch, err := s.DBInstance.GetBatchByIndex(event.Index.Uint64())
	if err != nil && gorm.IsRecordNotFoundError(err) {
		s.Logger.Info("Found a new batch, applying transactions and adding new batch", "index", event.Index.Uint64)
		newRoot, commitments, err := s.parseAndApplyBatch(vLog.TxHash, event.BatchType)
		if err != nil {
			fmt.Println("error while applying batch", "error", err)
			return
		}
		newBatch := core.NewBatch(newRoot.String(), event.Committer.String(), vLog.TxHash.String(), uint64(event.BatchType), core.BATCH_COMMITTED)
		batchID, err := s.DBInstance.AddNewBatch(newBatch, commitments)
		if err != nil {
			s.Logger.Error("Error adding new batch to DB", "error", err)
			return
		}

		s.Logger.Info("Added new batch", "ID", batchID)
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
	event := new(tokenregistry.TokenregistryRegisteredToken)

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
	newToken := db.Token{TokenID: event.TokenType.Uint64(), Address: event.TokenContract.String()}
	if err := s.DBInstance.AddToken(newToken); err != nil {
		panic(err)
	}
}

func (s *Syncer) sendDepositFinalisationTx() error {
	params, err := s.DBInstance.GetParams()
	if err != nil {
		fmt.Println("error while getting params")
		return err
	}
	nodeToBeReplaced, siblings, err := s.DBInstance.GetDepositNodeAndSiblings()
	if err != nil {
		fmt.Println("error finding replaced nodes", err)
		return err
	}
	totalBatches, err := s.DBInstance.GetBatchCount()
	if err != nil {
		fmt.Println("error find total batches", err)
		return err
	}
	commitmentMP, err := s.DBInstance.GetLastCommitmentMP(uint64(totalBatches))
	if err != nil {
		fmt.Println("error creating commitmentMP", err)
		return err
	}

	err = s.loadedBazooka.FireDepositFinalisation(nodeToBeReplaced, siblings, commitmentMP, params.MaxDepositSubTreeHeight)
	if err != nil {
		fmt.Println("error sending tx", err)
		return err
	}

	return nil
}

func (s *Syncer) parseAndApplyBatch(txHash ethCmn.Hash, batchType uint8) (newRoot core.ByteArray, commitments []core.Commitment, err error) {
	calldata, err := s.loadedBazooka.ParseCalldata(txHash, batchType)
	if err == bazooka.ErrNoTxs {
		return newRoot, commitments, nil
	}
	if err != nil {
		return newRoot, commitments, fmt.Errorf("unable to parse calldata %s", err)
	}
	newRoot, commitments, err = s.applyBatch(calldata, txHash, uint64(batchType), true)
	if err != nil {
		return
	}

	// apply transactions
	return newRoot, commitments, nil
}

func (s *Syncer) applyBatch(calldata bazooka.Calldata, txHash ethCmn.Hash, txType uint64, isSyncing bool) (newRoot core.ByteArray, commitments []core.Commitment, err error) {
	rootNode, err := s.DBInstance.GetAccountRoot()
	if err != nil {
		return newRoot, nil, err
	}

	accountRoot := rootNode.Hash
	var transactions []core.Tx
	var commitmentData []core.CommitmentData

	switch txType {
	case core.TX_TRANSFER_TYPE:
		batchInfo, ok := calldata.(bazooka.TransferCalldata)
		if !ok {
			return newRoot, commitments, errors.New("Error converting calldata to batchinfo")
		}
		transactions, err = s.decompressTransfers(batchInfo.Txss)
		if err != nil {
			return newRoot, commitments, err
		}

		commitmentData, err = batchInfo.Commitments(accountRoot)
		if err != nil {
			return newRoot, commitments, err
		}
	case core.TX_MASS_MIGRATIONS:
		batchInfo, ok := calldata.(bazooka.MassMigrationCalldata)
		if !ok {
			return newRoot, commitments, errors.New("Error converting calldata to batchinfo")
		}
		transactions, err = s.decompressMassMigrations(batchInfo.Txss, batchInfo.Meta)
		if err != nil {
			return newRoot, commitments, err
		}
		commitmentData, err = batchInfo.Commitments(accountRoot)
		if err != nil {
			return newRoot, commitments, err
		}
	case core.TX_CREATE_2_TRANSFER:
		batchInfo, ok := calldata.(bazooka.Create2TransferCalldata)
		if !ok {
			return newRoot, commitments, errors.New("Error converting calldata to batchinfo")
		}
		transactions, err = s.decompressCreate2Transfers(batchInfo.Txss)
		if err != nil {
			return newRoot, commitments, err
		}
		commitmentData, err = batchInfo.Commitments(accountRoot)
		if err != nil {
			return newRoot, commitments, err
		}
	default:
		fmt.Println("TxType didnt match any options", txType)
		return newRoot, commitments, errors.New("Didn't match any options")
	}

	s.Logger.Info("Parsed calldata", "totalTxs", len(transactions))

	commitments, err = db.ProcessTxs(&s.loadedBazooka, &s.DBInstance, transactions, isSyncing)
	if err != nil {
		return newRoot, commitments, err
	}

	for i := range commitments {
		commitments[i].BodyRoot = commitmentData[i].BodyRoot
		commitments[i].StateRoot = commitmentData[i].StateRoot
	}

	return core.BytesToByteArray(commitments[len(commitments)-1].StateRoot), commitments, nil
}

// decompressTransfers decompresses transfer bytes to TX
func (s *Syncer) decompressTransfers(decompressedTxs [][]byte) (txs []core.Tx, err error) {
	froms, tos, amounts, fees, err := s.loadedBazooka.DecompressTransferTxs(concatTxs(decompressedTxs))
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

func (s *Syncer) decompressCreate2Transfers(decompressedTxs [][]byte) (txs []core.Tx, err error) {
	froms, tos, toAccIDs, amounts, fees, err := s.loadedBazooka.DecompressCreate2TransferTxs(concatTxs(decompressedTxs))
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

		txData, err := s.loadedBazooka.EncodeCreate2TransferTx(froms[i].Int64(), tos[i].Int64(), toAccIDs[i].Int64(), fees[i].Int64(), nonce.Int64(), amounts[i].Int64(), int64(core.TX_CREATE_2_TRANSFER))
		if err != nil {
			return transactions, err
		}

		newTx := core.NewTx(froms[i].Uint64(), tos[i].Uint64(), core.TX_CREATE_2_TRANSFER, nil, txData)
		transactions = append(transactions, newTx)
	}

	return transactions, nil
}

// decompressTransfers decompresses transfer bytes to TX
func (s *Syncer) decompressMassMigrations(decompressedTxs [][]byte, meta [][4]*big.Int) (txs []core.Tx, err error) {
	froms, amounts, fees, err := s.loadedBazooka.DecompressMassMigrationTxs(concatTxs(decompressedTxs))
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

		txData, err := s.loadedBazooka.EncodeMassMigrationTx(froms[i].Int64(), meta[i][0].Int64(), fees[i].Int64(), nonce.Int64(), amounts[i].Int64(), int64(core.TX_MASS_MIGRATIONS))
		if err != nil {
			return transactions, err
		}

		newTx := core.NewTx(froms[i].Uint64(), 0, core.TX_TRANSFER_TYPE, nil, txData)
		transactions = append(transactions, newTx)
	}

	return transactions, nil
}
