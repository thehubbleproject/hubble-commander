package core

import (
	"encoding/hex"
	"encoding/json"
	"math/big"

	ethCmn "github.com/ethereum/go-ethereum/common"
)

type Hash ethCmn.Hash
type Address ethCmn.Address

type ByteArray [32]byte

func (b ByteArray) String() string {
	bz := b[:]
	enc := make([]byte, len(bz)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], bz)
	return string(enc)
}

// String has to be prefixed with 0x
func HexToByteArray(a string) (b ByteArray, err error) {
	bz, err := hex.DecodeString(a[2:])
	if err != nil {
		return b, err
	}
	return BytesToByteArray(bz), nil
}

func BytesToByteArray(bz []byte) ByteArray {
	var temp [32]byte
	copy(temp[:], bz)
	return temp
}

// Pubkey is an alias for public key
type Pubkey [4]*big.Int

func (p Pubkey) String() (string, error) {
	pubBytes, err := p.Serialize()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(pubBytes), nil
}

// Serialize seralises a public key to bytes
func (p Pubkey) Serialize() ([]byte, error) {
	var pubkey []uint64
	for i := range p {
		pub := p[i].Uint64()
		pubkey = append(pubkey, pub)
	}
	return json.Marshal(pubkey)
}

// BytesToPubkey converts bytes to Pubkey
func BytesToPubkey(b []byte) (pubkey Pubkey, err error) {
	var p []uint64
	err = json.Unmarshal(b, &p)
	if err != nil {
		return pubkey, err
	}
	var pubkeyBigInt [4]*big.Int
	for i := range p {
		temp := big.NewInt(0)
		temp.SetUint64(p[i])
		pubkeyBigInt[i] = temp
	}
	return pubkeyBigInt, nil
}

func StrToPubkey(s string) (pubkey Pubkey, err error) {
	pubBytes, err := hex.DecodeString(s)
	if err != nil {
		return
	}
	return BytesToPubkey(pubBytes)
}
