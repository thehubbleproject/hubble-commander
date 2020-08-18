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

func BytesToSignature(b []byte) (blswallet.Signature, error) {
	sig, err := blswallet.SignatureKeyFromBytes(b)
	return *sig, err
}

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

func (w *Wallet) Bytes() (secretKey []byte, pubkey []byte) {
	accountBytes := w.Account.ToBytes()
	secretBytes := accountBytes[128:]
	pubkeyBytes := accountBytes[:128]
	return secretBytes, pubkeyBytes
}

func SecretToWallet(secretKey []byte, pubkey []byte) (wallet Wallet, err error) {
	in := append(pubkey, secretKey...)
	keyPair, err := blswallet.NewKeyPairFromBytes(in)
	if err != nil {
		return
	}
	return Wallet{Account: keyPair}, nil
}

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
