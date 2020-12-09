package core

import (
	"math/big"
)

// Pubkey is an alias for public key
type Pubkey []byte

const pubkeyLength = 128

func NewPubkey(p [4]*big.Int) Pubkey {
	var pubkey []byte
	for _, part := range p {
		pubkey = append(pubkey, part.Bytes()...)
	}
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
