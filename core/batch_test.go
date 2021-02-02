package core

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

// Thanks @miguelmotah <3
func TestTransferBody(t *testing.T) {
	expectedOutput := "3a53dc4890241dbe03e486e785761577d1c369548f6b09aa38017828dcdf5c2707857e73108d077c5b7ef89540d6493f70d940f1763a9d34c9d98418a39d28ac02bb0e4743a7d0586711ee3dd6311256579ab7abcd53c9c76f040bfde4d6d6e90000000000000000000000000000000000000000000000000000000000000000000000000000000100010000"

	// bytes32 accountRoot
	accountRoot := "0x3a53dc4890241dbe03e486e785761577d1c369548f6b09aa38017828dcdf5c27"
	// uint256[2] calldata signatures
	signatures := []string{
		"3402053321874964899321528271743396700217057178612185975187363512030360053932",
		"1235124644010117237054094970590473241953434069965207718920579820322861537001",
	}
	// bytes calldata txss
	txss := "000000000000000100010000"

	accountRootBz, err := HexToByteArray(accountRoot)
	require.Nil(t, err)

	var signature [2]*big.Int
	signature[0] = StringToDecimalBigInt(signatures[0])
	signature[1] = StringToDecimalBigInt(signatures[1])

	txssBz, err := hex.DecodeString(txss)
	require.Nil(t, err)

	tb := TransferBody{
		AccountRoot: accountRootBz,
		Signature:   signature,
		FeeReceiver: big.NewInt(0),
		Txs:         txssBz,
	}

	var tc TransferCommitment
	tc.TransferBody = tb

	data, err := tc.Bytes()
	require.Nil(t, err)

	require.Equal(t, expectedOutput, hex.EncodeToString(data), "error output doesnt match")
}

// StringToDecimalBigInt takes in a string and returns the corresponding big int
func StringToDecimalBigInt(s string) *big.Int {
	t := big.NewInt(0)
	t.SetString(s, 10)
	return t
}
