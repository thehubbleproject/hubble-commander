package core

import (
	"testing"

	"github.com/BOPR/config"
	"github.com/BOPR/wallet"
	blswallet "github.com/kilic/bn254/bls"
	"github.com/stretchr/testify/require"
)

func TestFetchBatchInput(t *testing.T) {
	// data ::= "5f5b95b8000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000001edb3e000fa2bce948fe20edec89567eb469f406088938db809b7f5c0c57403fa00000000000000000000000000000000000000000000000000000000000000010d187c89a3a3d52a5415de61538149796f205f93b1c9477c6c5acd633d97dec21fdb9d57d407fad614c2857588f31d8422664b25abd87638df01d8f501581ff8000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000180000000200000003000100000000000200000003000100000000000000000000"
	// inputMap := "map[feeReceivers:[0] signatures:[[5923330888430140098499636240370570309915734290469917058249406549173531565762 14409723755289018913401382336605519721181704378557932192077219658986738425848]] stateRoots:[[237 179 224 0 250 43 206 148 143 226 14 222 200 149 103 235 70 159 64 96 136 147 141 184 9 183 245 192 197 116 3 250]] txss:[[0 0 0 2 0 0 0 3 0 1 0 0 0 0 0 2 0 0 0 3 0 1 0 0]]]"
}

func TestVerifySingle(t *testing.T) {
	err := config.ParseAndInitGlobalConfig("../.")
	require.Nil(t, err, "error should be nil")

	bazooka, err := NewPreLoadedBazooka()
	require.Nil(t, err, "error should be nil")

	newWallet, err := wallet.NewWallet()
	require.Nil(t, err)

	secret, pubkey := newWallet.Bytes()
	txData, err := bazooka.EncodeTransferTx(1, 2, 0, 0, 0, TX_TRANSFER_TYPE)
	require.Nil(t, err)

	tx, err := NewPendingTx(1, 2, TX_TRANSFER_TYPE, []byte(""), txData)
	require.Nil(t, err)

	txBytes, err := tx.GetSignBytes(bazooka)
	require.Nil(t, err)

	err = tx.SignTx(secret, pubkey, txBytes)
	require.Nil(t, err)

	sig, err := blswallet.SignatureKeyFromBytes(tx.Signature)
	require.Nil(t, err)

	pubkeyObj, err := blswallet.PublicKeyFromBytes(pubkey)
	require.Nil(t, err)

	// verify signature using go-lib
	valid, err := newWallet.VerifySignature(txBytes, *sig, *pubkeyObj)
	// this passes
	require.True(t, valid)
	require.Nil(t, err)

	pubkeyInt, err := Pubkey(pubkey).ToSol()
	require.Nil(t, err)

	solSignature, err := BytesToSolSignature(tx.Signature)
	require.Nil(t, err)

	// verify single
	err = bazooka.SC.Transfer.Validate(nil, txData, solSignature, pubkeyInt, wallet.DefaultDomain)
	// this doesnt
	require.Nil(t, err)
}
