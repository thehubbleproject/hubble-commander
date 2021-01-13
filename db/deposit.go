package db

import (
	"bytes"
	"errors"

	"github.com/BOPR/core"
)

// GetDepositNodeAndSiblings fetches the right intermediate node that has to be replaced for deposits
func (db *DB) GetDepositNodeAndSiblings() (nodeToBeReplaced core.UserState, siblings []core.UserState, err error) {
	// get params
	params, err := db.GetParams()
	if err != nil {
		return
	}

	// get the deposit node
	// it fetches empty node according to the deposit tree height and its hash
	expectedHash := core.DefaultHashes[params.MaxDepositSubTreeHeight]

	// getNode with the expectedHash
	nodeToBeReplaced, err = db.GetDepositSubTreeRoot(expectedHash.String(), params.MaxDepth-params.MaxDepositSubTreeHeight)
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

// FinaliseDepositsAndAddBatch finalises deposits and a
func (db *DB) FinaliseDepositsAndAddBatch(depositRoot core.ByteArray, pathToDepositSubTree uint64) (string, error) {
	var root string
	db.Logger.Info("Finalising states", "depositRoot", depositRoot, "pathToDepositSubTree", pathToDepositSubTree)

	// update the empty leaves with new states
	err := db.FinaliseDeposits(pathToDepositSubTree, depositRoot)
	if err != nil {
		return root, err
	}

	rootStateNode, err := db.GetRoot()
	if err != nil {
		return root, err
	}

	return rootStateNode.Hash, nil
}

func (db *DB) FinaliseDeposits(pathToDepositSubTree uint64, depositRoot core.ByteArray) error {
	params, err := db.GetParams()
	if err != nil {
		return err
	}

	// find out the states that are finalised
	userStates, err := db.GetPendingStateByDepositRoot(depositRoot)
	if err != nil {
		return err
	}

	// find out where the insertion was made
	height := params.MaxDepth - 1
	getTerminalNodesOf, err := core.SolidityPathToNodePath(pathToDepositSubTree, height)
	if err != nil {
		return err
	}

	terminalNodes, err := db.GetAllTerminalNodes(getTerminalNodesOf)
	if err != nil {
		return err
	}

	for i, node := range userStates {
		node.Status = core.STATUS_ACTIVE
		node.UpdatePath(terminalNodes[i])
		node.UpdateHash()
		err := db.UpdateState(node)
		if err != nil {
			return err
		}

		// delete pending states
		err = db.DeletePendingState(node.AccountID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) GetPendingDeposits(numberOfStates uint64) ([]core.UserState, error) {
	var states []core.UserState
	err := db.Instance.Limit(numberOfStates).Where("status = ?", 0).Find(&states).Error
	return states, err
}

func (db *DB) GetAllTerminalNodes(pathToDepositSubTree string) (terminalNodes []string, err error) {
	buf := bytes.Buffer{}
	buf.WriteString(pathToDepositSubTree)
	buf.WriteString("%")
	var states []core.UserState

	// LIKE query with search for terminal nodes to DB
	if err = db.Instance.Where("path LIKE ? AND type = ?", buf.String(), core.TYPE_TERMINAL).Find(&states).Error; err != nil {
		return
	}

	// get all states while making sure they are empty and append to paths array
	for _, node := range states {
		if node.Hash != core.ZeroLeaf.String() {
			return terminalNodes, errors.New("State not zero, aborting operation")
		}
		terminalNodes = append(terminalNodes, node.Path)
	}
	return
}
