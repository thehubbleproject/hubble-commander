package core

import (
	"fmt"
	"testing"
)

func TestGetAdjacentNodePath(t *testing.T) {
	res, err := GetAdjacentNodePath("00000001")
	fmt.Println("data", res, err)
}

func TestGetParentPath(t *testing.T) {
	leftChildPath := "000"
	// rightChildPath := "111"
	// expectedParentPath := "11"
	fmt.Println(GetParentPath(leftChildPath))
}

func TestBasicPathMutations(t *testing.T) {
	// index 1024 in binary is the path below
	index := uint64(1024)
	depth := uint64(10)

	newPath, err := SolidityPathToNodePath(index, depth)
	if err != nil {
		panic(err)
	}

	data, err := StringToUint(newPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("path generated", newPath, "data", data)
}

func TestStringToBigInt(t *testing.T) {
	path := "001"
	fmt.Println(StringToBigInt(path).String())
}
