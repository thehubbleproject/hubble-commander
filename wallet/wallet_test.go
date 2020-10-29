package wallet

import (
	"encoding/hex"
	"fmt"
	"testing"

	blswallet "github.com/kilic/bn254/bls"
	"github.com/stretchr/testify/require"
)

func TestSignAndVerify(t *testing.T) {
	wallet, err := NewWallet()
	require.Equal(t, err, nil, "error creating wallet")
	signBytes := []byte("0x123222")
	signature, err := wallet.Sign(signBytes)
	require.Equal(t, err, nil, "error signing transaction")
	fmt.Println(hex.EncodeToString(signature.ToBytes()))
	valid, err := wallet.VerifySignature(signBytes, signature, *wallet.Account.Public)
	require.Equal(t, err, nil, "error verifying signature")
	require.Equal(t, valid, true, "error verifying signature")
}

func TestVerifyAggregated(t *testing.T) {
	hasher := gHasher
	signBytes := []byte("0x123222")
	signerSize := 2
	publicKeys := make([]*blswallet.PublicKey, signerSize)
	messages := make([]*blswallet.Message, signerSize)
	signatures := make([]*blswallet.Signature, signerSize)
	for i := 0; i < signerSize; i++ {
		account, err := NewWallet()
		if err != nil {
			t.Fatal(err)
		}
		accountSignature, err := account.Sign(signBytes)
		messages[i] = createMessage(signBytes)
		publicKeys[i] = account.Account.Public
		signatures[i] = &accountSignature
	}

	verifier := blswallet.NewBLSVerifier(hasher)
	aggregatedSignature := verifier.AggregateSignatures(signatures)
	aggregatedSignatureWallet, err := NewAggregateSignature(signatures)
	if err != nil {
		panic(err)
	}
	fmt.Println("aggregated sig", hex.EncodeToString(aggregatedSignature.ToBytes()), hex.EncodeToString(aggregatedSignatureWallet.ToBytes()))
	verified, err := VerifyAggregatedSignature(messages, publicKeys, aggregatedSignatureWallet)
	if err != nil {
		t.Fatal(err)
	}
	if !verified {
		t.Fatalf("signature is not verified")
	}
}
