package wallet

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	blswallet "github.com/kilic/bn254/bls"
	"github.com/stretchr/testify/require"
)

func TestPrivateKey(t *testing.T) {
	domainBytes, _ := hex.DecodeString("1234ABCD")
	// keyMaterial, _ := hex.DecodeString("5566")
	// key := crypto.Keccak256Hash(keyMaterial).Bytes()
	key, _ := hex.DecodeString("0bbb61c6ddb95fa018b3cbd40fe7abe2861c54040646dcc76011013e8352ce1a")
	fmt.Printf("%x", key)
	wallet, err := SecretToWallet(key, crypto.Keccak256Hash(domainBytes))
	require.NoError(t, err)
	pubkeyBytes, _ := hex.DecodeString("1d66806e962cbb8c9ef907c2b17996b5d908ee93d86cddc3c3f0401f8ba1dfba24030326130d84e3d2019b0bbc5bf724b0ec84ec4f341586d891cc49f0424fa3061ac5817571d009b58a7b3f2e05ed9d0c4b8f4f741e3fed47217561785c23f11118f4af592b7f48163427ca489eee044594ee1148b0b8e8eb7e35e8a7c13137")
	require.Equal(t, fmt.Sprintf("%x", pubkeyBytes), fmt.Sprintf("%x", wallet.signer.Account.Public.ToBytes()))
}

func TestSignAndVerify(t *testing.T) {
	wallet, err := NewWallet(DefaultDomain)
	require.Equal(t, err, nil, "error creating wallet")
	signBytes := []byte("0x123222")
	signature, err := wallet.Sign(signBytes)
	require.Equal(t, err, nil, "error signing transaction")
	fmt.Println(hex.EncodeToString(signature.ToBytes()))
	valid, err := wallet.VerifySignature(signBytes, signature, *wallet.signer.Account.Public)
	require.Equal(t, err, nil, "error verifying signature")
	require.Equal(t, valid, true, "error verifying signature")
}

func TestVerifyAggregated(t *testing.T) {
	signBytes := []byte("0x123222")
	signerSize := 2
	publicKeys := make([]*blswallet.PublicKey, signerSize)
	messages := make([]blswallet.Message, signerSize)
	signatures := make([]*blswallet.Signature, signerSize)
	for i := 0; i < signerSize; i++ {
		wallet, err := NewWallet(DefaultDomain)
		if err != nil {
			t.Fatal(err)
		}
		accountSignature, err := wallet.Sign(signBytes)
		if err != nil {
			t.Fatal(err)
		}
		messages[i] = signBytes
		publicKeys[i] = wallet.signer.Account.Public
		signatures[i] = &accountSignature
	}

	aggregatedSignature := blswallet.AggregateSignatures(signatures)
	aggregatedSignatureWallet, err := NewAggregateSignature(signatures)
	if err != nil {
		panic(err)
	}
	fmt.Println("aggregated sig", hex.EncodeToString(aggregatedSignature.ToBytes()), hex.EncodeToString(aggregatedSignatureWallet.ToBytes()))
	verified, err := VerifyAggregatedSignature(messages, publicKeys, aggregatedSignatureWallet, DefaultDomain)
	if err != nil {
		t.Fatal(err)
	}
	if !verified {
		t.Fatalf("signature is not verified")
	}
}
