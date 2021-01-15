package db

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/BOPR/core"
)

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
	accounts, err := db.GetPendingAccByDepositRoot(depositRoot)
	if err != nil {
		return err
	}
	fmt.Println("pending accounts by depositRoot", len(accounts))

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

	for i, acc := range accounts {
		acc.Status = core.STATUS_ACTIVE
		acc.Type = core.TYPE_TERMINAL
		acc.UpdatePath(terminalNodes[i])
		acc.CreateAccountHash()
		err := db.UpdateState(acc)
		if err != nil {
			return err
		}

		// delete pending account
		err = db.DeletePendingAccount(acc.AccountID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) GetPendingDeposits(numberOfAccs uint64) ([]core.UserState, error) {
	var accounts []core.UserState
	err := db.Instance.Limit(numberOfAccs).Where("status = ?", core.STATUS_PENDING).Find(&accounts).Error
	if err != nil {
		return accounts, err
	}
	return accounts, nil
}

func (db *DB) GetAllTerminalNodes(pathToDepositSubTree string) (terminalNodes []string, err error) {
	buf := bytes.Buffer{}
	buf.WriteString(pathToDepositSubTree)
	buf.WriteString("%")
	var accounts []core.UserState

	// LIKE query with search for terminal nodes to DB
	if err = db.Instance.Where("path LIKE ? AND type = ?", buf.String(), core.TYPE_TERMINAL).Find(&accounts).Error; err != nil {
		return
	}

	// get all accounts while making sure they are empty and append to paths array
	for _, account := range accounts {
		if account.Hash != core.ZeroLeaf.String() {
			return terminalNodes, errors.New("Account not zero, aborting operation")
		}
		terminalNodes = append(terminalNodes, account.Path)
	}
	return
}
