package bazooka

import (
	"math/big"

	"github.com/BOPR/core"
)

type TypesUserState struct {
	PubkeyID *big.Int
	TokenID  *big.Int
	Balance  *big.Int
	Nonce    *big.Int
}

type TypesStateMerkleProof struct {
	State   TypesUserState
	Witness [][32]byte
}

// TypesCommitmentInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesCommitmentInclusionProof struct {
	Commitment core.CommitmentData
	Path       *big.Int
	Witness    [][32]byte
}
