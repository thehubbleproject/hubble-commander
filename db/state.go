package db

import (
	"errors"
	"fmt"

	"github.com/BOPR/core"
	"github.com/jinzhu/gorm"
)

// InitStateTree initialises the states tree
func (db *DB) InitStateTree(depth int) error {
	rootNode := core.NewStateRoot(depth)
	return db.Instance.Create(&rootNode).Error
}

func (db *DB) UpdateState(state core.UserState) error {
	state.CreateAccountHash()
	siblings, err := db.GetSiblings(state.Path)
	if err != nil {
		return err
	}

	db.Logger.Debug("Updating state", "Hash", state.Hash, "Path", state.Path, "countOfSiblings", len(siblings))
	return db.storeLeaf(state, state.Path, siblings)
}

// ReserveEmptyLeaf reserves an empty state leaf and maps the accID to state leaf
func (db *DB) ReserveEmptyLeaf(accID uint64) (stateID uint64, err error) {
	params, err := db.GetParams()
	if err != nil {
		return
	}
	state, err := db.FindEmptyAndReserve(int(params.MaxDepth), accID)
	if err != nil {
		return
	}
	return core.StringToUint(state.Path)
}

func (db *DB) storeLeaf(state core.UserState, path string, siblings []core.UserState) error {
	var err error
	var isLeft bool
	computedNode := state
	for i := 0; i < len(siblings); i++ {
		var parentHash core.ByteArray
		sibling := siblings[i]
		isComputedRightSibling := core.GetNthBitFromRight(
			path,
			i,
		)
		if isComputedRightSibling == 0 {
			parentHash, err = core.GetParent(computedNode.HashToByteArray(), sibling.HashToByteArray())
			if err != nil {
				return err
			}
			isLeft = true
			// Store the node!
			err = db.StoreNode(parentHash, computedNode, sibling, isLeft)
			if err != nil {
				return err
			}
		} else {
			parentHash, err = core.GetParent(sibling.HashToByteArray(), computedNode.HashToByteArray())
			if err != nil {
				return err
			}
			isLeft = false
			// Store the node!
			err = db.StoreNode(parentHash, sibling, computedNode, isLeft)
			if err != nil {
				return err
			}
		}
		parentAccount, err := db.GetStateByPath(core.GetParentPath(computedNode.Path))
		if err != nil {
			return err
		}
		computedNode = parentAccount
	}
	// Store the new root
	err = db.UpdateRootNodeHashes(computedNode.HashToByteArray())
	if err != nil {
		return err
	}
	return nil
}

// StoreNode updates the nodes given the parent hash
func (db *DB) StoreNode(parentHash core.ByteArray, leftNode core.UserState, rightNode core.UserState, isLeft bool) (err error) {
	if isLeft {
		// update left account
		err = db.updateState(leftNode, leftNode.Path)
		if err != nil {
			return err
		}
	} else {
		// update right account
		err = db.updateState(rightNode, rightNode.Path)
		if err != nil {
			return err
		}
	}
	// update the parent with the new hashes
	return db.UpdateParentWithHash(core.GetParentPath(leftNode.Path), parentHash)
}

func (db *DB) UpdateParentWithHash(pathToParent string, newHash core.ByteArray) error {
	var tempState core.UserState
	tempState.Path = pathToParent
	nodeType, err := db.FindNodeType(pathToParent)
	if err != nil {
		return err
	}
	tempState.Type = nodeType
	tempState.Hash = newHash.String()
	return db.updateState(tempState, pathToParent)
}

func (db *DB) UpdateRootNodeHashes(newRoot core.ByteArray) error {
	var tempState core.UserState
	tempState.Type = core.TYPE_ROOT
	tempState.Path = ""
	tempState.Hash = newRoot.String()
	return db.updateState(tempState, tempState.Path)
}

func (db *DB) GetSiblings(path string) ([]core.UserState, error) {
	var relativePath = path
	var siblings []core.UserState
	for i := len(path); i > 0; i-- {
		otherChild := core.GetOtherChild(relativePath)
		otherNode, err := db.GetStateByPath(otherChild)
		if err != nil {
			return siblings, err
		}
		siblings = append(siblings, otherNode)
		relativePath = core.GetParentPath(relativePath)
	}

	return siblings, nil
}

func (db *DB) GetStateByIndex(index uint64) (state core.UserState, err error) {
	path, err := db.IDToPath(index)
	if err != nil {
		return
	}
	return db.GetStateByPath(path)
}

func (db *DB) GetStateByPath(path string) (core.UserState, error) {
	var userState core.UserState
	err := db.Instance.Scopes(QueryByPath(path)).Find(&userState).Error
	if err == gorm.ErrRecordNotFound {
		nodeType, err := db.FindNodeType(path)
		if err != nil {
			return userState, err
		}
		height, err := db.DepthToHeight(len(path))
		if err != nil {
			return userState, err
		}
		node := core.NewStateNode(path, core.DefaultHashes[height].String(), nodeType)
		return *node, nil
	}
	if err != nil {
		return userState, err
	}

	return userState, nil
}

// GetRoot fetches the root of the state tree
func (db *DB) GetRoot() (core.UserState, error) {
	var state core.UserState
	err := db.Instance.Scopes(QueryByType(core.TYPE_ROOT)).Find(&state).Error
	if err != nil {
		return state, core.ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v", err))
	}
	return state, nil
}

// updateState will simply replace all the changed fields
func (db *DB) updateState(newState core.UserState, path string) error {
	var state core.UserState
	err := db.Instance.Model(&state).Scopes(QueryByPath(path)).Find(&state).Error
	if gorm.IsRecordNotFoundError(err) {
		return db.Instance.Create(&newState).Error
	}
	if err != nil {
		return err
	}
	err = db.Instance.Model(&state).Scopes(QueryByPath(path)).Update(&newState).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) FindEmptyState(depth int) (state core.UserState, err error) {
	params, err := db.GetParams()
	if err != nil {
		return
	}
	if depth > int(params.MaxDepth) {
		return state, errors.New("depth cannot be greater than max depth")
	}
	totalLeaves := core.TotalLeavesForDepth(depth)
	expectedHash := core.DefaultHashes[int(params.MaxDepth)-depth]
	for i := 0; i < totalLeaves; i++ {
		path, errr := core.SolidityPathToNodePath(uint64(i), uint64(depth))
		if errr != nil {
			return
		}
		state, errr = db.GetStateByPath(path)
		if errr != nil {
			return
		}

		if state.Hash == expectedHash.String() && !state.IsReserved {
			break
		}
	}
	return state, nil
}

func (db *DB) GetStateByAccID(accID uint64) (states []core.UserState, err error) {
	err = db.Instance.Scopes(QueryByAccountID(accID)).Find(&states).Error
	if err != nil {
		return
	}
	return states, nil
}

func (db *DB) FindEmptyAndReserve(depth int, accID uint64) (state core.UserState, err error) {
	state, err = db.FindEmptyState(depth)
	if err != nil {
		return
	}
	state.IsReserved = true
	state.AccountID = accID
	err = db.updateState(state, state.Path)
	if err != nil {
		return
	}
	return
}
