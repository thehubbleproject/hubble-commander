package common

import (
	"encoding/hex"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFloat16(t *testing.T) {
	tests := []struct {
		uncompressed string
		compressed   string
	}{
		{"3290000000000000", "d149"},
		{"2207", "089f"},
		{"27020000", "4a8e"},
		{"2816000", "3b00"},
		{"226000000000", "90e2"},
		{"1758000000000", "96de"},
		{"2096000000000", "9830"},
		{"3023000000", "6bcf"},
		{"30460000000", "7be6"},
		{"3045000000000000000", "fbe5"},
		{"36740000000", "7e5a"},
		{"1501000", "35dd"},
		{"3771000", "3ebb"},
		{"1715000", "36b3"},
		{"95600", "23bc"},
		{"14310000000000", "a597"},
		{"3335000000", "6d07"},
		{"99100000", "53df"},
		{"207200000000000", "b818"},
		{"11490000", "447d"},
	}
	for _, tt := range tests {
		ttInt, _ := strconv.ParseUint(tt.uncompressed, 10, 64)
		ttBytes, _ := hex.DecodeString(tt.compressed)
		ttByte2 := [2]byte{ttBytes[0], ttBytes[1]}

		compressed, err := Compress(ttInt)
		require.NoError(t, err)
		require.Equal(t, ttByte2, compressed, "Compression failed")

		decompressed := Decompress(ttByte2)
		require.Equal(t, ttInt, decompressed, "Decompression failed")

	}

}

func TestRounding(t *testing.T) {
	tests := []struct {
		input  uint64
		expect uint64
	}{
		{12345, 12340},
		{56789, 56700},
		{123123, 123100},
		{186950000000, 186900000000},
		{4096, 4090},
		{4095, 4095},
	}
	for _, tt := range tests {
		rounded, err := Round(tt.input)
		require.NoError(t, err)
		require.Equal(t, tt.expect, rounded)
	}

}
