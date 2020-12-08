package core

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPubkey(t *testing.T) {
	pubkeyStr := "1047972a2f18ce99324b762b6e10838bd4f6ed2560fb6ef89f813731e0b1e52105e897c3350ebe44fea7db66c63faa9de9885937ff0cc799abaa9d5eaeaea5be2803e3489934f070d7ed4f1da3e44d9944e62fb1d6aa66280672899630229de52e12403fc86c5d2210f502689506489fbae1671d7cfb9dab38a7752427215169"
	bz, err := hex.DecodeString(pubkeyStr)
	if err != nil {
		panic(err)
	}
	pubkey := Pubkey(bz)
	pubInt, err := pubkey.ToSol()
	if err != nil {
		panic(err)
	}
	require.Equal(t, hex.EncodeToString(NewPubkey(pubInt)), pubkeyStr)
}
