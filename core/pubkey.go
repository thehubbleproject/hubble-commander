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

// NewPubkeyFromBytes creates a pubkey from bytes
// func NewPubkeyFromBytes(bz []byte) (pubkey Pubkey, err error) {
// 	if len(bz) != pubkeyLength {
// 		return pubkey, ErrInvalidPubkeyLen
// 	}
// 	chunkSize := 32
// 	for i := 0; i < pubkeyLength; i += chunkSize {
// 		end := i + chunkSize
// 		// necessary check to avoid slicing beyond
// 		// slice capacity
// 		if end > len(bz) {
// 			end = len(bz)
// 		}
// 		pubkeyPart := bz[i:end]
// 		tempPubkeyPart := big.NewInt(0)
// 		tempPubkeyPart = tempPubkeyPart.SetBytes(pubkeyPart)
// 		pubkey[i/chunkSize] = tempPubkeyPart
// 	}
// 	return pubkey, nil
// }

// func (p Pubkey) String() (string, error) {
// 	pubBytes, err := p.serialize()
// 	if err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(pubBytes), nil
// }

// // Serialize seralises a public key to bytes
// func (p Pubkey) serialize() ([]byte, error) {
// 	var pubkey []uint64
// 	for i := range p {
// 		pub := p[i].Uint64()
// 		pubkey = append(pubkey, pub)
// 	}
// 	return json.Marshal(pubkey)
// }

// // bytesToPubkey converts bytes to Pubkey
// func bytesToPubkey(b []byte) (pubkey Pubkey, err error) {
// 	var p []uint64
// 	err = json.Unmarshal(b, &p)
// 	if err != nil {
// 		return pubkey, err
// 	}
// 	var pubkeyBigInt [4]*big.Int
// 	for i := range p {
// 		temp := big.NewInt(0)
// 		temp.SetUint64(p[i])
// 		pubkeyBigInt[i] = temp
// 	}
// 	return pubkeyBigInt, nil
// }

// func StrToPubkey(s string) (pubkey Pubkey, err error) {
// 	pubBytes, err := hex.DecodeString(s)
// 	if err != nil {
// 		return
// 	}

// 	return bytesToPubkey(pubBytes)
// }
