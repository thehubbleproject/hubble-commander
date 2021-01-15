package core

import (
	"encoding/hex"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/willf/pad"
	"golang.org/x/crypto/sha3"
)

// StringToBigInt takes in a string and returns the corresponding big int
func StringToBigInt(s string) *big.Int {
	t := big.NewInt(0)
	t.SetString(s, 2)
	return t
}

func UintToBigInt(a uint64) *big.Int {
	t := big.NewInt(0)
	t.SetUint64(a)
	return t
}

func StringToUint(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func UintToString(a uint64) string {
	return strconv.FormatUint(a, 2)
}

func GetNthBitFromRight(path string, index int) int {
	dataRune := []rune(path)
	// if the bit is 0
	if dataRune[len(dataRune)-index-1] == 48 {
		return 0
	} else {
		return 1
	}
}

func FlipBitInString(s string, i int) string {
	dataRune := []rune(s)
	if dataRune[i] == 48 {
		dataRune[i] = 49
	} else {
		dataRune[i] = 48
	}
	return string(dataRune)
}

func GetOtherChild(path string) string {
	return FlipBitInString(path, len(path)-1)
}

func trimPathToParentPath(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func GetAdjacentNodePath(path string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("path is empty")
	}
	nodePath, err := StringToUint(path)
	if err != nil {
		return "", err
	}
	adjacentNodePath := nodePath + 1

	return pad.Left(UintToString(adjacentNodePath), len(path), "0"), nil
}

// SolidityPathToNodePath goes from 3 to 000000000011
func SolidityPathToNodePath(path uint64, depth uint64) (string, error) {
	pathWithoutPrefix := UintToString(path)
	// pad path with 0's to make it fit depth
	var pathToNode []rune

	if depth < uint64(len(pathWithoutPrefix)) {
		return "", errors.New("Path should be greater or equal to depth")
	}

	for i := uint64(0); i < (depth)-uint64(len(pathWithoutPrefix)); i++ {
		pathToNode = append(pathToNode, 48)
	}

	generatedPath := strings.Join([]string{string(pathToNode), pathWithoutPrefix}, "")

	return generatedPath, nil
}

func BytesToSolSignature(in []byte) (out [2]*big.Int, err error) {
	if len(in) != 64 {
		return out, errors.New("Invalid signature length")
	}
	out[0] = new(big.Int).SetBytes(in[:32])
	out[1] = new(big.Int).SetBytes(in[32:64])
	return out, nil
}

func Keccak256(data []byte) ethCmn.Hash {
	return crypto.Keccak256Hash(data)
}

func KeccakFromString(data string) (hash ethCmn.Hash, err error) {
	bz, err := hex.DecodeString(data)
	if err != nil {
		return
	}
	return Keccak256(bz), nil

}

func RlpHash(x interface{}) (h ethCmn.Hash, err error) {
	hw := sha3.NewLegacyKeccak256()
	if err = rlp.Encode(hw, x); err != nil {
		return
	}
	hw.Sum(h[:0])
	return h, nil
}

// GetAllChildren fetches paths for all children for a particular node
func GetAllChildren(path string, treeDepth int) (childrenPaths []string, err error) {
	if len(path) == treeDepth {
		return
	}
	heightDiff := treeDepth - len(path)
	var suffix []rune
	for i := 0; i < heightDiff; i++ {
		suffix = append(suffix, 48)
	}
	path = strings.Join([]string{path, string(suffix)}, "")
	totalChildren := TotalLeavesForDepth(heightDiff)
	firstIndex, err := StringToUint(path)
	if err != nil {
		return
	}
	for i := uint64(0); i < uint64(totalChildren); i++ {
		idx := firstIndex + i
		// TODO replace with normal uint to string conversion with depth check
		// this is slow
		childPath, errr := SolidityPathToNodePath(idx, uint64(treeDepth))
		if errr != nil {
			return
		}
		childrenPaths = append(childrenPaths, childPath)
	}
	return
}

// TotalLeavesForDepth calculate total leaves for depth
func TotalLeavesForDepth(depth int) (totalLeaves int) {
	return int(math.Exp2(float64(depth)))
}

func encodeChildren(left, right ByteArray) (result []byte, err error) {
	bytes32Type, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		return
	}

	arguments := abi.Arguments{
		{
			Type: bytes32Type,
		},
		{
			Type: bytes32Type,
		},
	}
	bz, err := arguments.Pack(
		[32]byte(left),
		[32]byte(right),
	)
	if err != nil {
		return
	}

	return bz, nil
}
