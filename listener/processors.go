package listener

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/common"
	"github.com/BOPR/contracts/accountregistry"
	"github.com/BOPR/contracts/depositmanager"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/core"
	"github.com/BOPR/db"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
)

var (
	ErrBadBatch = errors.New("STATE ROOT MISMATCH, BAD BATCH FOUND")
)

func (s *Syncer) processNewPubkeyAddition(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) error {
	// unpack event
	event := new(accountregistry.AccountregistryPubkeyRegistered)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		s.Logger.Error("Unable to unpack log", "error", err)
		return err
	}

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"accountID", event.PubkeyID.String(),
		"pubkey", event.Pubkey,
	)
	params, err := s.DBInstance.GetParams()
	if err != nil {
		s.Logger.Error("Error getting params", "error", err)
		return err
	}

	pathToNode, err := core.SolidityPathToNodePath(event.PubkeyID.Uint64(), params.MaxDepth)
	if err != nil {
		s.Logger.Error("Error getting solidity path", "error", err)
		return err
	}

	pubkey := core.NewPubkey(event.Pubkey)
	nodeType, err := s.DBInstance.FindNodeType(pathToNode)
	if err != nil {
		s.Logger.Error("Error fetching nodetype", "error", err)
		return err
	}

	newAcc, err := core.NewAccount(event.PubkeyID.Uint64(), pubkey, pathToNode, nodeType)
	if err != nil {
		s.Logger.Error("Error creating account", "error", err)
		return err
	}

	if err := s.DBInstance.AddNewAccount(*newAcc); err != nil {
		s.Logger.Error("Error adding new account", "error", err)
		return err
	}

	// if pubkey was added by relayer mark the packet processed
	if err := s.DBInstance.MarkPacketDone(pubkey); err != nil {
		s.Logger.Error("Error marking packet done", "error", err)
		return err
	}

	return nil
}

func (s *Syncer) processDepositQueued(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) error {
	s.Logger.Info("New deposit found")
	event := new(depositmanager.DepositmanagerDepositQueued)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		s.Logger.Error("Error unpacking log", "error", err)
		return err
	}

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"pubkeyID", event.PubkeyID.String(),
		"Data", hex.EncodeToString(event.Data),
	)

	newDeposit := core.NewDeposit(event.PubkeyID.Uint64(), event.Data)
	err = s.DBInstance.AddNewDeposit(*newDeposit)
	if err != nil {
		s.Logger.Error("Error adding new deposit", "error", err)
		return err
	}
	return nil
}

func (s *Syncer) processDepositSubtreeCreated(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) error {
	s.Logger.Info("New deposit subtree created")

	// unpack event
	event := new(depositmanager.DepositmanagerDepositSubTreeReady)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		s.Logger.Error("Error unpacking new log", "error", err)
		return err
	}

	// send deposit finalisation transction to ethereum chain
	catchingup, err := IsCatchingUp(s.loadedBazooka, s.DBInstance)
	if err != nil {
		s.Logger.Error("Error catching up", "error", err)
		return err
	}

	if !catchingup {
		err = s.sendDepositFinalisationTx()
		if err != nil {
			s.Logger.Error("Error sending deposit finalisation", "error", err)
			return err
		}
	} else {
		s.Logger.Info("Still catching up, aborting deposit finalisation")
	}

	// attach deposit information
	err = s.DBInstance.AttachDepositInfo(event.Root)
	if err != nil {
		s.Logger.Error("Unable to attach deposit info", "error", err)
		return err
	}

	return nil
}

func (s *Syncer) processDepositFinalised(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) error {
	s.Logger.Info("Deposit batch finalised!")
	event := new(rollup.RollupDepositsFinalised)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		s.Logger.Error("Error unpacking log", "error", err)
		return err
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
	err = s.DBInstance.FinaliseDepositsAndAddBatch(depositRoot, pathToDepositSubTree.Uint64())
	if err != nil {
		s.Logger.Error("Error finalized deposit and adding new batch", "error", err)
		return err
	}
	return nil
}

func (s *Syncer) processNewBatch(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) error {
	s.Logger.Info("New batch found!")

	event := new(rollup.RollupNewBatch)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		return err
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
		_, commitments, err := s.parseAndApplyBatch(vLog.TxHash, event.BatchType)
		if err == ErrBadBatch {
			s.Logger.Error("Fraud detected, please challenge batch on-chain", "error", err)
			panic(err)
		}
		if err != nil {
			s.Logger.Error("Error applying batch", "error", err)
			return err
		}
		newBatch := core.NewBatch(event.Committer.String(), vLog.TxHash.String(), uint64(event.BatchType), core.BATCH_COMMITTED)
		batchID, err := s.DBInstance.AddNewBatch(newBatch, commitments)
		if err != nil {
			s.Logger.Error("Error adding new batch to DB", "error", err)
			return err
		}

		s.Logger.Info("Added new batch", "ID", batchID)
		return nil
	}

	if err != nil {
		s.Logger.Error("Unable to fetch batch", "index", event.Index, "err", err)
		return err
	}

	// Mark seen batch as committed if we havent already
	if batch.Status != core.BATCH_COMMITTED {
		s.Logger.Info("Found a non committed batch")
		err = s.DBInstance.CommitBatch(event.Index.Uint64())
		if err != nil {
			s.Logger.Error("Unable to commit batch", "index", event.Index.String(), "err", err)
			return err
		}
	}
	return nil
}

func (s *Syncer) sendDepositFinalisationTx() error {
	params, err := s.DBInstance.GetParams()
	if err != nil {
		s.Logger.Error("Error fetching params", "error", err)
		return err
	}
	nodeToBeReplaced, siblings, err := s.DBInstance.GetDepositNodeAndSiblings()
	if err != nil {
		s.Logger.Error("Error fetching deposit node and siblings", "error", err)
		return err
	}
	totalBatches, err := s.DBInstance.GetBatchCount()
	if err != nil {
		s.Logger.Error("Error fetching totalBatches", "error", err)
		return err
	}
	commitmentMP, err := s.DBInstance.GetLastCommitmentMP(uint64(totalBatches) - 1)
	if err != nil {
		s.Logger.Error("Error creating commitment merkle proof", "error", err)
		return err
	}
	err = s.loadedBazooka.FireDepositFinalisation(nodeToBeReplaced, siblings, commitmentMP, params.MaxDepositSubTreeHeight)
	if err != nil {
		s.Logger.Error("Error sending deposit finalisation", "error", err)
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

	s.Logger.Info("Calldata parsed, applying batch", "batchType", batchType)

	newRoot, commitments, err = s.applyBatch(calldata, txHash, uint64(batchType), true)
	if err != nil {
		return
	}

	// make sure the new root post tx apply matches the one we have in calldata
	// if not we have detected a fraud that needs to be challenged on-chain
	if newRoot != calldata.LastStateRoot() {
		return newRoot, commitments, ErrBadBatch

	}

	s.Logger.Info("Batch applied successfully!", "newRoot", newRoot)
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
	var txsInCommitment []int

	switch txType {
	case core.TX_TRANSFER_TYPE:
		batchInfo, ok := calldata.(bazooka.TransferCalldata)
		if !ok {
			return newRoot, commitments, errors.New("Error converting calldata to batchinfo")
		}
		transactions, txsInCommitment, err = decompressTransfers(s.loadedBazooka, s.DBInstance, batchInfo.Txss)
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
		transactions, txsInCommitment, err = decompressMassMigrations(s.loadedBazooka, s.DBInstance, batchInfo.Txss, batchInfo.Meta)
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
		transactions, txsInCommitment, err = decompressCreate2Transfers(s.loadedBazooka, s.DBInstance, batchInfo.Txss)
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

	commitments, err = db.ProcessTxs(&s.loadedBazooka, &s.DBInstance, transactions, txsInCommitment, isSyncing)
	if err != nil {
		return newRoot, commitments, err
	}

	for i := range commitments {
		commitments[i].BodyRoot = commitmentData[i].BodyRoot
		commitments[i].StateRoot = commitmentData[i].StateRoot
	}

	// it is assumed that we will have atleast 1 commitment here because we revery before if there are no transactions
	return core.BytesToByteArray(commitments[len(commitments)-1].StateRoot), commitments, nil
}
