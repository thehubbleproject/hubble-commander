package core

import (
	"fmt"
	"testing"

	"github.com/BOPR/common"
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

	txData, err := bazooka.EncodeTransferTx(1, 2, 0, 0, 0, TX_TRANSFER_TYPE)
	if err != nil {
		return
	}

	tx, err := NewPendingTx(1, 2, TX_TRANSFER_TYPE, []byte(""), txData)
	if err != nil {
		return
	}

	txBytes, err := tx.GetSignBytes(bazooka)
	if err != nil {
		return
	}

	newWallet, err := wallet.NewWallet()
	if err != nil {
		return
	}

	secret, pubkey := newWallet.Bytes()
	err = tx.SignTx(secret, pubkey, common.Keccak256(txBytes))
	if err != nil {
		return
	}

	sig, err := blswallet.SignatureKeyFromBytes(tx.Signature)
	if err != nil {
		fmt.Println("error while getting signature", err)
		return
	}

	pubkeyObj, err := blswallet.PublicKeyFromBytes(pubkey)
	if err != nil {
		fmt.Println("error while getting public key", err)
		return
	}
	pubkeyInt, err := Pubkey(pubkey).ToSol()
	if err != nil {
		return
	}

	solSignature, err := BytesToSolSignature(tx.Signature)
	if err != nil {
		return
	}

	valid, err := newWallet.VerifySignature(common.Keccak256(txBytes).Bytes(), *sig, *pubkeyObj)
	fmt.Println(valid, err)

	err = bazooka.SC.Transfer.VerifySingle(nil, txBytes, pubkeyInt, solSignature, wallet.DefaultDomain)
	if err != nil {
		fmt.Println("error on validate", err)
		return
	}
}
