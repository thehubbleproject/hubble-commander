package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
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
func TestStringToBigInt(t *testing.T) {
	path := "001"
	fmt.Println(StringToBigInt(path).String())
}

func TestGetAllChildren(t *testing.T) {
	subTreePath := "0001"
	depth := 7
	expectedOutput := []string{"0001000", "0001001", "0001010", "0001011", "0001100", "0001101", "0001110", "0001111"}
	childrentPath, err := GetAllChildren(subTreePath, depth)
	require.Nil(t, err, "error should not be nil")
	require.Equal(t, expectedOutput, childrentPath, "children path should match")
}
