package wallet

import (
	"crypto/rand"

	blswallet "github.com/kilic/bn254/bls"
)

type Wallet struct {
	Account *blswallet.KeyPair
	Hasher  blswallet.Hasher
}

type Signature struct {
	Signature blswallet.Signature
}

func getBLSSignatures(sigs []Signature) (blsSigs []*blswallet.Signature) {
	for _, sig := range sigs {
		blsSigs = append(blsSigs, &sig.Signature)
	}
	return
}

func NewWallet() (wallet Wallet, err error) {
	newAccount, err := blswallet.NewKeyPair(rand.Reader)
	if err != nil {
		return
	}
	hasher := &blswallet.HasherKeccak256{}
	return Wallet{Account: newAccount, Hasher: hasher}, nil
}

func SecretToPublicKey(secretKey []byte) {
	// if len(secretKey) != 32 {
	// 	// error invalid priv key
	// }
	// // var se [32]byte
	// // copy(s[:], secretKey)
	// s := big.NewInt(0)
	// s.SetBytes(secretKey)
	// g2 := bn254.NewG2()
	// public := g2.New()
	// g2.MulScalar(public, g2.One(), s)
	// secret := &blswallet.SecretKey{}
	// copy(secret[32-len(s.Bytes()):], s.Bytes()[:])
	// keyPair := blswallet.KeyPair{secret: secret, Public: public}
}

func createMessage(data []byte) *blswallet.Message {
	return &blswallet.Message{Message: data, Domain: []byte{}}
}

func (w *Wallet) Sign(data []byte) ([]byte, error) {
	signer := blswallet.NewBLSSigner(w.Hasher, w.Account)
	signature, err := signer.Sign(createMessage(data))
	if err != nil {
		return []byte(""), err
	}
	return signature.ToBytes(), nil
}

func VerifySignature(data []byte, signature Signature, pubkey blswallet.PublicKey) (valid bool, err error) {
	hasher := &blswallet.HasherSHA256{}
	verifier := blswallet.NewBLSVerifier(hasher)
	return verifier.Verify(createMessage(data), &signature.Signature, &pubkey)
}

func VerifyAggregatedSignature(data [][]byte, aggregateSignature Signature, pubkeys []*blswallet.PublicKey) (valid bool, err error) {
	hasher := &blswallet.HasherSHA256{}
	verifier := blswallet.NewBLSVerifier(hasher)
	var messages []*blswallet.Message
	for _, txData := range data {
		messages = append(messages, createMessage(txData))
	}
	return verifier.VerifyAggregate(messages, pubkeys, &aggregateSignature.Signature)
}

// NewAggregateSignature creates a new aggregated signature
func NewAggregateSignature(signatures []Signature) (aggregatedSignature Signature) {
	hasher := &blswallet.HasherSHA256{}
	verifier := blswallet.NewBLSVerifier(hasher)
	blsSigs := getBLSSignatures(signatures)
	aggregatedSig := verifier.AggregateSignatures(blsSigs)
	return Signature{Signature: *aggregatedSig}
}
