package core

import "math/big"

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
	chunkSize := 32
	for i := 0; i < pubkeyLength; i += chunkSize {
		end := i + chunkSize
		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(p) {
			end = len(p)
		}
		pubkeyPart := p[i:end]
		tempPubkeyPart := big.NewInt(0)
		tempPubkeyPart = tempPubkeyPart.SetBytes(pubkeyPart)
		pubkey[i/chunkSize] = tempPubkeyPart
	}
	return pubkey, nil
}
