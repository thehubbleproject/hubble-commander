package core

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Pubkey is an alias for public key
type Pubkey []byte

const pubkeyLength = 128

func NewPubkey(p [4]*big.Int) Pubkey {
	pubkey := make([]byte, pubkeyLength)
	copy(pubkey[:32], common.LeftPadBytes(p[1].Bytes(), 32))
	copy(pubkey[32:64], common.LeftPadBytes(p[0].Bytes(), 32))
	copy(pubkey[64:96], common.LeftPadBytes(p[3].Bytes(), 32))
	copy(pubkey[96:128], common.LeftPadBytes(p[2].Bytes(), 32))
	return pubkey
}

func NewPubkeyFromString(p [4]string) Pubkey {
	pubkey := make([]byte, pubkeyLength)
	copy(pubkey[:32], common.Hex2BytesFixed(p[1], 32))
	copy(pubkey[32:64], common.Hex2BytesFixed(p[0], 32))
	copy(pubkey[64:96], common.Hex2BytesFixed(p[3], 32))
	copy(pubkey[96:128], common.Hex2BytesFixed(p[2], 32))
	return pubkey
}

func (p Pubkey) ToSol() (pubkey [4]*big.Int, err error) {
	if len(p) != pubkeyLength {
		return pubkey, ErrInvalidPubkeyLen
	}
	pubkey[1] = new(big.Int).SetBytes(p[:32])
	pubkey[0] = new(big.Int).SetBytes(p[32:64])
	pubkey[3] = new(big.Int).SetBytes(p[64:96])
	pubkey[2] = new(big.Int).SetBytes(p[96:128])

	return pubkey, nil
}

func (p Pubkey) String() string {
	return hex.EncodeToString(p)
}

func (p Pubkey) ToHash() (str string, err error) {
	uint256Arr4, err := abi.NewType("uint256[4]", "", nil)
	if err != nil {
		return "", err
	}

	arguments := abi.Arguments{{Type: uint256Arr4}}
	ints, err := p.ToSol()
	if err != nil {
		return "", err
	}
	bytes, err := arguments.Pack(ints)
	if err != nil {
		return "", err
	}

	return Keccak256(bytes).String(), nil
}
