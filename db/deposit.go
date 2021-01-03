package db

import (
	"bytes"
	"errors"

	"github.com/BOPR/core"
)

func (db *DB) GetDepositNodeAndSiblings() (NodeToBeReplaced core.UserState, siblings []core.UserState, err error) {
	// get params
	params, err := db.GetParams()
	if err != nil {
		return
	}

	// get the deposit node
	// it fetches empty node according to the deposit tree height and its hash
	expectedHash := core.DefaultHashes[params.MaxDepositSubTreeHeight]

	// getNode with the expectedHash
	NodeToBeReplaced, err = db.GetDepositSubTreeRoot(expectedHash.String(), params.MaxDepth-params.MaxDepositSubTreeHeight)
	if err != nil {
		return
	}

	// get siblings for the path to node
	siblings, err = db.GetSiblings(NodeToBeReplaced.Path)
	if err != nil {
		return
	}

	return
}

func (db *DB) FinaliseDepositsAndAddBatch(depositRoot core.ByteArray, pathToDepositSubTree uint64) (string, error) {
	var root string
	db.Logger.Info("Finalising accounts", "depositRoot", depositRoot, "pathToDepositSubTree", pathToDepositSubTree)

	// update the empty leaves with new accounts
	err := db.FinaliseDeposits(pathToDepositSubTree, depositRoot)
	if err != nil {
		return root, err
	}

	rootAccount, err := db.GetRoot()
	if err != nil {
		return root, err
	}

	return rootAccount.Hash, nil
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

	// find out where the insertion was made
	height := params.MaxDepth - 1
	getTerminalNodesOf, err := core.SolidityPathToNodePath(pathToDepositSubTree, height)
	if err != nil {
		return err
	}

	// TODO add error for if no account found
	terminalNodes, err := db.GetAllTerminalNodes(getTerminalNodesOf)
	if err != nil {
		return err
	}

	for i, acc := range accounts {
		acc.Status = core.STATUS_ACTIVE
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
	err := db.Instance.Limit(numberOfAccs).Where("status = ?", 0).Find(&accounts).Error
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
