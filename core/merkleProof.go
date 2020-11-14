package core

import (
	"github.com/BOPR/contracts/rollupclient"
)

type StateMerkleProof struct {
	State    UserStateNode
	Siblings []UserStateNode
}

func NewStateMerkleProof(account UserStateNode, siblings []UserStateNode) StateMerkleProof {
	return StateMerkleProof{State: account, Siblings: siblings}
}

func (m *StateMerkleProof) ToABIVersion() (stateMP rollupclient.TypesStateMerkleProof, err error) {
	var witnesses [][32]byte
	for _, s := range m.Siblings {
		witnesses = append(witnesses, s.HashToByteArray())
	}

	state, err := m.State.ToABIState()
	if err != nil {
		return
	}

	stateMP = rollupclient.TypesStateMerkleProof{
		State:   state,
		Witness: witnesses,
	}

	return stateMP, nil
}

type AccountMerkleProof struct {
	Path      string
	PublicKey string
	Siblings  []Account
}

func NewAccountMerkleProof(path string, publicKey string, siblings []Account) AccountMerkleProof {
	return AccountMerkleProof{PublicKey: publicKey, Siblings: siblings, Path: path}
}
