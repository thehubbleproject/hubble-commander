package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	ZERO_UINT       = 0
	DEFAULT_DEPTH   = 2
	DEFAULT_BALANCE = 0
	DEFAULT_NONCE   = 0
)

// Genesis describes the fields in genesis.json file
type Genesis struct {
	StartEthBlock           uint64 `json:"start_eth_block,omitempty"`     // should be set to the eth block num the contracts were deployed at
	MaxTreeDepth            uint64 `json:"max_tree_depth,omitempty"`      // height for all trees initially
	MaxDepositSubTreeHeight uint64 `json:"max_deposit_subtree,omitempty"` // max height for deposit subtrees initially
	StakeAmount             uint64 `json:"stake_amount,omitempty"`        // initial stake amount set on contracts
}

// Validate validates the genesis file and checks for basic things
func (g Genesis) Validate() error {
	return nil
}

func DefaultGenesis() Genesis {
	return Genesis{
		StartEthBlock:           0,
		MaxTreeDepth:            DEFAULT_DEPTH,
		MaxDepositSubTreeHeight: DEFAULT_DEPTH,
		StakeAmount:             32,
	}
}

func ReadGenesisFile() (Genesis, error) {
	var genesis Genesis

	genesisFile, err := os.Open("genesis.json")
	if err != nil {
		return genesis, err
	}
	defer genesisFile.Close()

	genBytes, err := ioutil.ReadAll(genesisFile)
	if err != nil {
		return genesis, err
	}

	err = json.Unmarshal(genBytes, &genesis)
	return genesis, err
}

func WriteGenesisFile(genesis Genesis) error {
	bz, err := json.MarshalIndent(genesis, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("genesis.json", bz, 0644)
}
