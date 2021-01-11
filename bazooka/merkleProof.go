package bazooka

import (
	"github.com/BOPR/core"
)

type StateMerkleProof struct {
	State    core.UserState
	Siblings []core.UserState
}

func NewStateMerkleProof(state core.UserState, siblings []core.UserState) StateMerkleProof {
	return StateMerkleProof{State: state, Siblings: siblings}
}

func (m *StateMerkleProof) ToABIVersion(b Bazooka) (stateMP TypesStateMerkleProof, err error) {
	var witnesses [][32]byte
	for _, s := range m.Siblings {
		witnesses = append(witnesses, s.HashToByteArray())
	}

	state, err := ToABIAccount(b, m.State)
	if err != nil {
		return
	}

	stateMP = TypesStateMerkleProof{
		State:   state,
		Witness: witnesses,
	}

	return stateMP, nil
}

func ToABIAccount(b Bazooka, s core.UserState) (solState TypesUserState, err error) {
	if len(s.Data) == 0 {
		return *(NewEmptyTypesUserState()), nil
	}
	solState.PubkeyID, solState.Balance, solState.Nonce, solState.TokenID, err = b.DecodeState(s.Data)
	return
}

type AccountMerkleProof struct {
	Path      string
	PublicKey string
	Siblings  []core.Account
}

func NewAccountMerkleProof(path string, publicKey string, siblings []core.Account) AccountMerkleProof {
	return AccountMerkleProof{PublicKey: publicKey, Siblings: siblings, Path: path}
}
