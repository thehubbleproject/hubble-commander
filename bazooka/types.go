package bazooka

import "math/big"

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
