package main

import (
	"encoding/hex"
	"fmt"

	"github.com/BOPR/core"
	"github.com/BOPR/wallet"
	blswallet "github.com/kilic/bn254/bls"
)

func main() {
	privkeystr := "837bc93f489f59d01e66d9451dbf2c2c2441e28092a197ce3a3ffe12902289a4"
	privKeyBytes, err := hex.DecodeString(privkeystr)
	if err != nil {
		panic(err)
	}
	userWallet, err := wallet.SecretToWallet(privKeyBytes, wallet.DefaultDomain)
	if err != nil {
		panic(err)
	}
	// pubkeystr := "093010d2d971651901219f61dbb7cec96fe77b7d690a6924f5a17ca563772e751bf9007952179170230ecba116ab235d381b9456c5cafa88797b7ed4b7369f261b347f8d4cddeb987d32dd7032b05d5f5bde2d22a58cdd669013cccbfb1da5162a409027360fbaeb96797595d20e0bb3a7faf388696d532e0acf07040dac04c0"
	// pubkeyBz, err := hex.DecodeString(pubkeystr)
	// if err != nil {
	// 	panic(err)
	// }
	// pubkeyStringArr := [4]string{"1bf9007952179170230ecba116ab235d381b9456c5cafa88797b7ed4b7369f26",
	// 	"093010d2d971651901219f61dbb7cec96fe77b7d690a6924f5a17ca563772e75", "2a409027360fbaeb96797595d20e0bb3a7faf388696d532e0acf07040dac04c0", "1b347f8d4cddeb987d32dd7032b05d5f5bde2d22a58cdd669013cccbfb1da516"}
	// pubkey := core.NewPubkeyFromString(pubkeyStringArr)
	tx, err := core.NewPendingTx(nil, nil, 1, 1, 0, 1, core.TX_TRANSFER_TYPE)
	if err != nil {
		panic(err)
	}
	txBytesStr := "000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000"
	txBytes, err := hex.DecodeString(txBytesStr)
	if err != nil {
		panic(err)
	}
	err = tx.SignTx(userWallet, txBytes)
	if err != nil {
		panic(err)
	}

	signature, err := blswallet.SignatureFromBytes(tx.Signature)
	if err != nil {
		panic(err)
	}
	fmt.Println("signature to bytes", hex.EncodeToString(tx.Signature))
	secretBytes, pubkeyBz := userWallet.Bytes()
	if err != nil {
		panic(err)
	}
	corePubkey := core.Pubkey(pubkeyBz)

	fmt.Println("match with secret key striung", hex.EncodeToString(secretBytes))
	fmt.Println("match with pub key striung", corePubkey.String())
	tempPub, err := blswallet.PublicKeyFromBytes(pubkeyBz)
	if err != nil {
		panic(err)
	}

	valid, err := userWallet.VerifySignature(txBytes, *signature, *tempPub)
	if err != nil {
		panic(err)
	}
	fmt.Println("valiud", valid, err)
	// fmt.Println("sigtx", hex.EncodeToString(tx.Signature))

}
