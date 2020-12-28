package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Genesis describes the fields in genesis.json file
type Genesis struct {
	Parameters struct {
		MAXDEPTH               uint64 `json:"MAX_DEPTH"`
		MAXDEPOSITSUBTREEDEPTH uint64 `json:"MAX_DEPOSIT_SUBTREE_DEPTH"`
		STAKEAMOUNT            string `json:"STAKE_AMOUNT"` // In Wei
		BLOCKSTOFINALISE       uint64 `json:"BLOCKS_TO_FINALISE"`
		MINGASLEFT             uint64 `json:"MIN_GAS_LEFT"`
		MAXTXSPERCOMMIT        uint64 `json:"MAX_TXS_PER_COMMIT"`
		USEBURNAUCTION         bool   `json:"USE_BURN_AUCTION"`
		DONATIONADDRESS        string `json:"DONATION_ADDRESS"`
		DONATIONNUMERATOR      uint64 `json:"DONATION_NUMERATOR"`
		GENESISSTATEROOT       string `json:"GENESIS_STATE_ROOT"`
	} `json:"parameters"`
	Addresses struct {
		ParamManager            string `json:"paramManager"`
		FrontendGeneric         string `json:"frontendGeneric"`
		FrontendTransfer        string `json:"frontendTransfer"`
		FrontendMassMigration   string `json:"frontendMassMigration"`
		FrontendCreate2Transfer string `json:"frontendCreate2Transfer"`
		NameRegistry            string `json:"nameRegistry"`
		BlsAccountRegistry      string `json:"blsAccountRegistry"`
		TokenRegistry           string `json:"tokenRegistry"`
		Transfer                string `json:"transfer"`
		MassMigration           string `json:"massMigration"`
		Create2Transfer         string `json:"create2Transfer"`
		Chooser                 string `json:"chooser"`
		ExampleToken            string `json:"exampleToken"`
		SpokeRegistry           string `json:"spokeRegistry"`
		Vault                   string `json:"vault"`
		DepositManager          string `json:"depositManager"`
		Rollup                  string `json:"rollup"`
		WithdrawManager         string `json:"withdrawManager"`
	} `json:"addresses"`
}

func (g Genesis) Validate() error {
	if g.Parameters.MAXDEPTH <= 0 {
		return fmt.Errorf("Bad Max tree size %d", g.Parameters.MAXDEPTH)
	}
	return nil
}

func ReadGenesisFile(path string) (Genesis, error) {
	var genesis Genesis

	genesisFile, err := os.Open(path)
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
