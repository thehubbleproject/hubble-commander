package core

import (
	"errors"
	"math"
)

var (
	ZeroLeaf      ByteArray
	DefaultHashes []ByteArray
)

func init() {
	var err error
	ZeroLeaf, err = HexToByteArray("0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563")
	if err != nil {
		panic(err)
	}

	DefaultHashes, err = GenDefaultHashes(100)
	if err != nil {
		panic(err)
	}
}

// Tree is a merkle tree that stores a whole tree indexed by `level` and `index`
// All nodes at particular level can be fetched by Nodes[level]
// Specific node can be fetched by level and depth both, Nodes[level][index]
type Tree struct {
	Nodes       [][]ByteArray
	TotalLeaves uint64
	Depth       uint64
	TotalLevels int
}

// NewTree creates a merkle tree from given leaves
// the order of leaves will be preserved
func NewTree(leaves []ByteArray) (newTree Tree, err error) {
	totalLeaves := len(leaves)
	treeDepth := minTreeDepth(len(leaves))
	totalLevels := treeDepth + 1

	if totalLeaves%2 != 0 {
		return newTree, errors.New("Even number of leaves expected")
	}

	newTree.TotalLeaves = uint64(totalLeaves)
	newTree.Depth = treeDepth
	newTree.TotalLevels = int(totalLevels)

	if totalLeaves == 0 {
		newTree.Nodes = make([][]ByteArray, 1)
		newTree.Nodes[0] = append(newTree.Nodes[0], ZeroLeaf)
		return newTree, nil
	}

	// initialize the array
	newTree.Nodes = make([][]ByteArray, int(totalLevels))

	// fill in the leaves
	for i := 0; i < totalLeaves; i++ {
		newTree.Nodes[totalLevels-1] = append(newTree.Nodes[totalLevels-1], leaves[i])
	}

	err = ascend(&newTree)
	if err != nil {
		return newTree, err
	}

	return newTree, nil
}

// GetWitnessForLeaf creates a witness for a given leaf
func (t *Tree) GetWitnessForLeaf(leafIndex uint64) (leaf ByteArray, witness []ByteArray, err error) {
	if leafIndex == 0 {
		return t.Nodes[0][0], t.Nodes[0], nil
	}
	if leafIndex > t.TotalLeaves {
		return leaf, witness, errors.New("Leaf index out of range")
	}
	leaf = t.Nodes[t.TotalLevels-1][leafIndex]
	depth := int(t.Depth)
	witness = make([]ByteArray, depth)
	for i := 0; i < depth; i++ {
		leafIndex ^= 1
		witness[i] = t.Nodes[depth-i][leafIndex]
		leafIndex >>= 1
	}
	return
}

// NodeCount returns the number of nodes on a level
func (t *Tree) NodeCount(level int) int {
	return int(math.Exp2(float64(level)))
}

// ascends from the leaf level towards root creating all intermediate nodes
func ascend(t *Tree) error {
	for level := t.TotalLevels - 1; level > 0; level-- {
		totalNodesAtLevel := t.NodeCount(level)
		for i := 0; i+1 < totalNodesAtLevel; i += 2 {
			parent, err := GetParent(t.Nodes[level][i], t.Nodes[level][i+1])
			if err != nil {
				return err
			}
			t.Nodes[level-1] = append(t.Nodes[level-1], parent)
		}
	}
	return nil
}

func minTreeDepth(totalLeaves int) uint64 {
	if totalLeaves == 1 {
		return 1
	}
	return uint64(math.Ceil(math.Log2(float64(totalLeaves))))
}

// GetParent takes in left and right children and returns the parent hash
func GetParent(left, right ByteArray) (parent ByteArray, err error) {
	data, err := encodeChildren(left, right)
	if err != nil {
		return parent, err
	}
	leaf := Keccak256(data)
	return BytesToByteArray(leaf.Bytes()), nil
}

// GetParentPath given the path to any of the children returns the path to the parent
func GetParentPath(path string) (parentNodePath string) {
	return trimPathToParentPath(path)
}

// GenDefaultHashes generates default hashes
func GenDefaultHashes(depth int) ([]ByteArray, error) {
	hashes := make([]ByteArray, depth)
	hashes[0] = ZeroLeaf
	for i := 1; i < depth; i++ {
		parent, err := GetParent(hashes[i-1], hashes[i-1])
		if err != nil {
			return hashes, err
		}
		hashes[i] = parent
	}
	return hashes, nil
}
