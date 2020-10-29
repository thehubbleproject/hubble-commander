package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenDefaultHashes(t *testing.T) {
	defaultHashes, err := GenDefaultHashes(1)
	require.NoError(t, err, "error generating default hashes")
	expected := "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563"
	require.Equal(t, defaultHashes[0].String(), expected)
	defaultHashes, err = GenDefaultHashes(2)
	require.NoError(t, err, "error generating default hashes")
	fmt.Println("defaultHashes", defaultHashes[0].String(), defaultHashes[1].String())
	require.Equal(t, defaultHashes[0].String(), "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563")
	require.Equal(t, defaultHashes[1].String(), "0x633dc4d7da7256660a892f8f1604a44b5432649cc8ec5cb3ced4c4e6ac94dd1d")
}
