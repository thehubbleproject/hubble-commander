package core

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/BOPR/wallet"
)

func TestPubkey(t *testing.T) {
	// create new account
	user, err := wallet.NewWallet()
	if err != nil {
		panic(err)
	}

	_, pubkeyBytes := user.Bytes()
	fmt.Println("data", hex.EncodeToString(pubkeyBytes))
	publicKey, err := NewPubkeyFromBytes(pubkeyBytes)
	if err != nil {
		panic(err)
	}
	pubkeyStr, err := publicKey.String()
	if err != nil {
		panic(err)
	}

	fmt.Println("pubkey string", pubkeyStr)

}
