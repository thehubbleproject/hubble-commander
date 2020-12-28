package core

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPubkey(t *testing.T) {

	/*
		> const ethers = require("ethers")
		> ethers.utils.solidityKeccak256(["uint256", "uint256", "uint256", "uint256"], ["0x01", "0x02", "0x03", "0x04"])
		'0x392791df626408017a264f53fde61065d5a93a32b60171df9d8a46afdf82992d'
	*/
	pubkey := NewPubkey([4]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4)})
	hash, err := pubkey.ToHash()
	require.NoError(t, err)
	require.Equal(t, "0x392791df626408017a264f53fde61065d5a93a32b60171df9d8a46afdf82992d", hash)

	/*
		> ethers.utils.solidityKeccak256(["uint256", "uint256", "uint256", "uint256"],  [
		  '0x2cbb4859ce0dead5cfc5a8bd444166ad6d913fd3b0cae8146b21a09cfe4ec9f5',
		  '0x2d07390b2b52a872e4c375a55e8fcb1389358b93484b5964eb2d32f5f9b9193c',
		  '0x0ba33185eae981b55323520a4068b5ca1a00a6f26713579b1ac6784474e777fd',
		  '0x067c53e7b0b8b9110534f733a7eb72c86959b37ad7521b35ebd21646bcb19912'
		])
		'0x440f7602722941c9563b2751bb8d69be261a08d1ea65be0780d53410ec705b12'
	*/
	pubkey = FromString([4]string{
		"2cbb4859ce0dead5cfc5a8bd444166ad6d913fd3b0cae8146b21a09cfe4ec9f5",
		"2d07390b2b52a872e4c375a55e8fcb1389358b93484b5964eb2d32f5f9b9193c",
		"0ba33185eae981b55323520a4068b5ca1a00a6f26713579b1ac6784474e777fd",
		"067c53e7b0b8b9110534f733a7eb72c86959b37ad7521b35ebd21646bcb19912",
	})
	hash, err = pubkey.ToHash()
	require.NoError(t, err)
	require.Equal(t, "0x440f7602722941c9563b2751bb8d69be261a08d1ea65be0780d53410ec705b12", hash)

}
