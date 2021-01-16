package db

import (
	"errors"
	"fmt"

	"github.com/BOPR/core"
)

func (db *DB) AddNewDeposit(deposit core.Deposit) error {
	return db.Instance.Create(&deposit).Error
}

// GetDepositNodeAndSiblings fetches the right intermediate node that has to be replaced for deposits
func (db *DB) GetDepositNodeAndSiblings() (nodeToBeReplaced core.UserState, siblings []core.UserState, err error) {
	// get params
	params, err := db.GetParams()
	if err != nil {
		return
	}

	// find out the number of leves in the level
	// 2^depth == number of leaves in the depth
	totalLeaves := core.TotalLeavesForDepth(int(params.MaxDepth))
	expectedHash := core.DefaultHashes[params.MaxDepositSubTreeHeight]

	for i := 0; i < totalLeaves; i++ {
		path, errr := core.SolidityPathToNodePath(uint64(i), params.MaxDepth-params.MaxDepositSubTreeHeight)
		if errr != nil {
			return
		}

		nodeToBeReplaced, errr = db.GetStateByPath(path)
		if errr != nil {
			return
		}

		if nodeToBeReplaced.Hash == expectedHash.String() {
			break
		}
	}

	// get siblings for the path to node
	siblings, err = db.GetSiblings(nodeToBeReplaced.Path)
	if err != nil {
		return
	}

	return
}

// FinaliseDepositsAndAddBatch finalises deposits and a
func (db *DB) FinaliseDepositsAndAddBatch(depositRoot core.ByteArray, pathToDepositSubTree uint64) error {
	db.Logger.Info("Finalising accounts", "depositRoot", depositRoot, "pathToDepositSubTree", pathToDepositSubTree)
	// update the empty leaves with new accounts
	err := db.FinaliseDeposits(pathToDepositSubTree, depositRoot)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) FinaliseDeposits(pathToDepositSubTree uint64, depositRoot core.ByteArray) error {
	params, err := db.GetParams()
	if err != nil {
		return err
	}
	// find out the accounts that are finalised
	deposits, err := db.GetPendingDepositsByDepositRoot(depositRoot)
	if err != nil {
		return err
	}
	db.Logger.Info("Got pending deposits", "depositRoot", depositRoot, "PendingDepositCount", len(deposits))

	// find out where the insertion was made
	height := params.MaxDepth - params.MaxDepositSubTreeHeight
	getTerminalNodesOf, err := core.SolidityPathToNodePath(pathToDepositSubTree, height)
	if err != nil {
		return err
	}

	terminalNodes, err := core.GetAllChildren(getTerminalNodesOf, int(params.MaxDepth))
	if err != nil {
		return err
	}
	if len(terminalNodes) != len(deposits) {
		return errors.New("deposit subtree cannot be empty")
	}

	for i, deposit := range deposits {
		// convery deposit to user state
		newUserState := core.NewUserState(deposit.AccountID, terminalNodes[i], deposit.Data)
		err := db.UpdateState(*newUserState)
		if err != nil {
			return err
		}
		// delete pending account
		err = db.DeletePendingDeposit(deposit.AccountID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) AttachDepositInfo(root core.ByteArray) error {
	var deposit core.Deposit
	result := db.Instance.Model(&deposit).Update("deposit_root", root.String())
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// GetPendingDepositsByDepositRoot fetches all deposits created by a specific deposit subtree
func (db *DB) GetPendingDepositsByDepositRoot(root core.ByteArray) ([]core.Deposit, error) {
	var pendingDeposits []core.Deposit
	query := db.Instance.Scopes(QueryByDepositRoot(root.String())).Find(&pendingDeposits)
	if err := query.Error; err != nil {
		return pendingDeposits, err
	}

	return pendingDeposits, nil
}

func (db *DB) DeletePendingDeposit(ID uint64) error {
	var deposit core.Deposit
	err := db.Instance.Scopes(QueryByAccountID(ID)).Delete(&deposit).Error
	if err != nil {
		return core.ErrRecordNotFound(fmt.Sprintf("unable to delete record for ID: %v", ID))
	}
	return nil
}

func (db *DB) GetPendingDeposits(numberOfAccs uint64) ([]core.Deposit, error) {
	var deposits []core.Deposit
	err := db.Instance.Limit(numberOfAccs).Find(&deposits).Error
	if err != nil {
		return deposits, err
	}
	return deposits, nil
}
