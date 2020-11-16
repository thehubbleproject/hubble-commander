package wallet

import (
	"crypto/rand"

	blswallet "github.com/kilic/bn254/bls"
)

type Wallet struct {
	signer blswallet.BLSSigner
}

var defaultDomain = []byte{0x00, 0x00, 0x00, 0x00}

func BytesToSignature(b []byte) (blswallet.Signature, error) {
	sig, err := blswallet.SignatureFromBytes(b)
	return *sig, err
}

func NewWallet() (wallet Wallet, err error) {
	newAccount, err := blswallet.NewKeyPair(rand.Reader)
	if err != nil {
		return
	}
	signer := blswallet.BLSSigner{Account: newAccount, Domain: defaultDomain}
	return Wallet{signer: signer}, nil
}

func (w *Wallet) Bytes() (secretKey []byte, pubkey []byte) {
	accountBytes := w.signer.Account.ToBytes()
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
	signer := blswallet.BLSSigner{Account: keyPair, Domain: defaultDomain}
	return Wallet{signer: signer}, nil
}

func (w *Wallet) Sign(data []byte) (blswallet.Signature, error) {
	signature, err := w.signer.Sign(data)
	if err != nil {
		return blswallet.Signature{}, err
	}
	return *signature, nil
}

func (w *Wallet) VerifySignature(data []byte, signature blswallet.Signature, pubkey blswallet.PublicKey) (valid bool, err error) {
	verifier := blswallet.NewBLSVerifier(w.signer.Domain)
	valid, err = verifier.Verify(data, &signature, &pubkey)
	return valid, err
}

func VerifyAggregatedSignature(data []blswallet.Message, pubkeys []*blswallet.PublicKey, aggregateSignature blswallet.Signature) (valid bool, err error) {
	verifier := blswallet.NewBLSVerifier(defaultDomain)
	return verifier.VerifyAggregate(data, pubkeys, &aggregateSignature)
}

// NewAggregateSignature creates a new aggregated signature
func NewAggregateSignature(signatures []*blswallet.Signature) (aggregatedSignature blswallet.Signature, err error) {
	aggregatedSig := blswallet.AggregateSignatures(signatures)
	return *aggregatedSig, nil
}
