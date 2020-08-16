package wallet

import (
	"crypto/rand"

	blswallet "github.com/kilic/bn254/bls"
)

var gHasher blswallet.Hasher

func init() {
	gHasher = &blswallet.HasherKeccak256{}
}

type Wallet struct {
	Account *blswallet.KeyPair
}

// type Signature struct {
// 	Signature blswallet.Signature
// }

// func NewSignature(sig blswallet.Signature) Signature {
// 	return Signature{Signature: sig}
// }

func BytesToSignature(b []byte) (blswallet.Signature, error) {
	sig, err := blswallet.SignatureKeyFromBytes(b)
	return *sig, err
}

// func (s *Signature) Bytes() []byte {
// 	return s.Signature.ToBytes()
// }

func getBLSSignatures(sigs []blswallet.Signature) (blsSigs []*blswallet.Signature) {
	for _, sig := range sigs {
		blsSigs = append(blsSigs, &sig)
	}
	return
}

func NewWallet() (wallet Wallet, err error) {
	newAccount, err := blswallet.NewKeyPair(rand.Reader)
	if err != nil {
		return
	}
	return Wallet{Account: newAccount}, nil
}

// func SecretToPublicKey(secretKey []byte) {
// 	if len(secretKey) != 32 {
// 		// error invalid priv key
// 	}
// 	// var se [32]byte
// 	// copy(s[:], secretKey)
// 	s := big.NewInt(0)
// 	s.SetBytes(secretKey)
// 	g2 := bn254.NewG2()
// 	public := g2.New()
// 	g2.MulScalar(public, g2.One(), s)
// 	secret := &blswallet.SecretKey{}
// 	copy(secret[32-len(s.Bytes()):], s.Bytes()[:])
// 	keyPair := blswallet.KeyPair{secret: secret, Public: public}
// }

func createMessage(data []byte) *blswallet.Message {
	return &blswallet.Message{Message: data, Domain: []byte{0x00, 0x00, 0x00, 0x00}}
}

func (w *Wallet) Sign(data []byte) (blswallet.Signature, error) {
	signer := blswallet.NewBLSSigner(gHasher, w.Account)
	signature, err := signer.Sign(createMessage(data))
	if err != nil {
		return blswallet.Signature{}, err
	}
	return *signature, nil
}

func (w *Wallet) VerifySignature(data []byte, signature blswallet.Signature, pubkey blswallet.PublicKey) (valid bool, err error) {
	verifier := blswallet.NewBLSVerifier(gHasher)
	return verifier.Verify(createMessage(data), &signature, w.Account.Public)
}

func VerifyAggregatedSignature(data []*blswallet.Message, pubkeys []*blswallet.PublicKey, aggregateSignature blswallet.Signature) (valid bool, err error) {
	verifier := blswallet.NewBLSVerifier(gHasher)
	return verifier.VerifyAggregate(data, pubkeys, &aggregateSignature)
}

// NewAggregateSignature creates a new aggregated signature
func NewAggregateSignature(signatures []*blswallet.Signature) (aggregatedSignature blswallet.Signature, err error) {
	verifier := blswallet.NewBLSVerifier(gHasher)
	aggregatedSig := verifier.AggregateSignatures(signatures)
	return *aggregatedSig, nil
}
