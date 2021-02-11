package db

import (
	"errors"
	"fmt"

	"github.com/BOPR/core"
)

// AddNewDeposit adds a new deposit to DB
func (db *DB) AddNewDeposit(deposit core.Deposit) error {
	return db.Instance.Create(&deposit).Error
}

// GetDepositNodeAndSiblings fetches the right intermediate node and siblings that has to be replaced for incoming deposits
func (db *DB) GetDepositNodeAndSiblings() (nodeToBeReplaced core.UserState, siblings []core.UserState, err error) {
	params, err := db.GetParams()
	if err != nil {
		return
	}

	// find an empty node to accomodate new deposits
	nodeToBeReplaced, err = db.FindEmptyState(int(params.MaxDepth) - int(params.MaxDepositSubTreeHeight))
	if err != nil {
		return
	}

	// get siblings for the path to node
	siblings, err = db.GetSiblings(nodeToBeReplaced.Path)
	if err != nil {
		return
	}

	return
}

// FinaliseDeposits finalises the deposits for a deposit subtree
func (db *DB) FinaliseDeposits(pathToDepositSubTree uint64, depositRoot core.ByteArray) error {
	db.Logger.Info("Finalising deposits", "pathToDepositSubTree", pathToDepositSubTree)

	params, err := db.GetParams()
	if err != nil {
		return err
	}

	// find out the deposits that are to be finalised
	deposits, err := db.GetPendingDeposits(depositRoot)
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
		fmt.Println("here", len(terminalNodes), len(deposits))
		return errors.New("deposit subtree cannot be empty")
	}

	for i, deposit := range deposits {
		// convert deposit to user state
		newUserState := core.NewUserState(deposit.AccountID, terminalNodes[i], deposit.Data)
		err := db.UpdateState(*newUserState)
		if err != nil {
			return err
		}
	}

	return db.ClearPendingDeposits()
}

// GetPendingDeposits fetches all deposits created by a specific deposit subtree
func (db *DB) GetPendingDeposits(root core.ByteArray) ([]core.Deposit, error) {
	var pendingDeposits []core.Deposit
	query := db.Instance.Order("accounc_id asc").Find(&pendingDeposits)
	if err := query.Error; err != nil {
		return pendingDeposits, err
	}
	return pendingDeposits, nil
}

// ClearPendingDeposits empties the pending deposit table
func (db *DB) ClearPendingDeposits() error {
	var deposit core.Deposit
	query := db.Instance.Delete(&deposit)
	err := query.Error
	if err != nil {
		return core.ErrRecordNotFound(fmt.Sprintf("unable to clear pendingDeposits"))
	}
	return nil
}
