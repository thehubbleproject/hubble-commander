package core

import (
	"bytes"
	"errors"
)

func (db *DB) GetDepositNodeAndSiblings() (NodeToBeReplaced UserStateNode, siblings []UserStateNode, err error) {
	// get params
	params, err := db.GetParams()
	if err != nil {
		return
	}

	// get the deposit node
	expectedHash := defaultHashes[params.MaxDepositSubTreeHeight]

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

func (db *DB) FinaliseDepositsAndAddBatch(depositRoot ByteArray, pathToDepositSubTree uint64) (string, error) {
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

func (db *DB) FinaliseDeposits(pathToDepositSubTree uint64, depositRoot ByteArray) error {
	params, err := db.GetParams()
	if err != nil {
		return err
	}

	// find out the states that are finalised
	states, err := db.GetPendingStateByDepositRoot(depositRoot)
	if err != nil {
		return err
	}

	// find out where the insertion was made
	height := params.MaxDepth - 1
	getTerminalNodesOf, err := SolidityPathToNodePath(pathToDepositSubTree, height)
	if err != nil {
		return err
	}

	// TODO add error for if no account found
	terminalNodes, err := db.GetAllTerminalNodes(getTerminalNodesOf)
	if err != nil {
		return err
	}

	for i, state := range states {
		state.Status = STATUS_ACTIVE
		state.UpdatePath(terminalNodes[i])
		state.CreateStateHash()
		err := db.UpdateState(state)
		if err != nil {
			return err
		}

		// delete pending account
		err = db.DeletePendingState(state.StateID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) GetPendingDeposits(numberOfStates uint64) ([]UserStateNode, error) {
	var states []UserStateNode
	err := db.Instance.Limit(numberOfStates).Where("status = ?", 0).Find(&states).Error
	if err != nil {
		return states, err
	}
	return states, nil
}

func (db *DB) GetAllTerminalNodes(pathToDepositSubTree string) (terminalNodes []string, err error) {
	buf := bytes.Buffer{}
	buf.WriteString(pathToDepositSubTree)
	buf.WriteString("%")
	var states []UserStateNode

	// LIKE query with search for terminal nodes to DB
	if err = db.Instance.Where("path LIKE ? AND type = ?", buf.String(), 1).Find(&states).Error; err != nil {
		return
	}

	// get all states while making sure they are empty and append to paths array
	for _, state := range states {
		if state.Hash != ZERO_VALUE_LEAF.String() {
			return terminalNodes, errors.New("State not zero, aborting operation")
		}
		terminalNodes = append(terminalNodes, state.Path)
	}
	return
}
