package core

import (
	"github.com/BOPR/contracts/rollupclient"
)

type AccountMerkleProof struct {
	Account  UserAccount
	Siblings []UserAccount
}

func NewAccountMerkleProof(account UserAccount, siblings []UserAccount) AccountMerkleProof {
	return AccountMerkleProof{Account: account, Siblings: siblings}
}

func (m *AccountMerkleProof) ToABIVersion() (AccMP rollupclient.TypesStateMerkleProof, err error) {
	// // create siblings
	// var siblingNodes [][32]byte
	// for _, s := range m.Siblings {
	// 	siblingNodes = append(siblingNodes, s.HashToByteArray())
	// }

	// account, err := m.Account.ToABIAccount()
	// if err != nil {
	// 	return
	// }

	// AccMP = rollupcaller.TypesAccountMerkleProof{
	// 	AccountIP: rollupcaller.TypesAccountInclusionProof{
	// 		PathToAccount: StringToBigInt(m.Account.Path),
	// 		Account:       account,
	// 	},
	// 	Siblings: siblingNodes,
	// }

	return AccMP, nil
}

type PDAMerkleProof struct {
	Path      string
	PublicKey string
	Siblings  []PDA
}

func NewPDAProof(path string, publicKey string, siblings []PDA) PDAMerkleProof {
	return PDAMerkleProof{PublicKey: publicKey, Siblings: siblings, Path: path}
}

// func (m *PDAMerkleProof) ToABIVersion() rollupcaller.TypesPDAMerkleProof {
// create siblings
// var siblingNodes [][32]byte
// for _, s := range m.Siblings {
// 	siblingNodes = append(siblingNodes, s.HashToByteArray())
// }
// pubkey, err := hex.DecodeString(m.PublicKey)
// if err != nil {
// 	panic(err)
// }
// pub1 := pubkey[0:32]
// pub2 := pubkey[32:64]
// pub3 := pubkey[64 : 64+32]
// pub4 := pubkey[64+32 : 64+32+32]
// sig1bigInt := big.NewInt(0)
// sig1bigInt.SetBytes(pub1)
// sig2bigInt := big.NewInt(0)
// sig2bigInt.SetBytes(pub2)
// sig3bigInt := big.NewInt(0)
// sig3bigInt.SetBytes(pub3)
// sig4bigInt := big.NewInt(0)
// sig4bigInt.SetBytes(pub4)
// aggregatedSigBigInt := [4]*big.Int{sig1bigInt, sig2bigInt, sig3bigInt, sig4bigInt}
// return rollupcaller.TypesPDAMerkleProof{
// 	Pda: rollupcaller.TypesPDAInclusionProof{
// 		PathToPubkey: StringToBigInt(m.Path),
// 		PubkeyLeaf:   rollupcaller.TypesPDALeaf{Pubkey: aggregatedSigBigInt},
// 	},
// 	Siblings: siblingNodes,
// }
// }
