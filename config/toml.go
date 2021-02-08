package config

import (
	"text/template"
)

// Note: any changes to the comments/variables/mapstructure
// must be reflected in the appropriate struct in config/config.go
const defaultConfigTemplate = `# This is a TOML config file.

##### RPC configrations #####
# RPC endpoint for ethereum chain
eth_RPC_URL = "{{ .EthRPC }}"

##### DB configrations #####
db_type = "{{ .DB }}"
db_url = "{{ .DBURL }}"
trace = "{{ .Trace }}"
db_log_mode = "{{ .DBLogMode }}"

##### Server configrations #####
server_port = "{{ .ServerPort }}"
polling_interval = "{{ .PollingInterval }}"
txs_per_commitment = "{{ .TxsPerCommitment }}"
max_commitments_per_batch = "{{ .MaxCommitmentsPerBatch }}"

#### Keystore #####
operator_key = "{{ .OperatorKey }}"
operator_address = "{{ .OperatorAddress }}"
keystore_passphrase = "{{ .KeystorePassphrase }}"

#### Syncer settings #####
confirmation_blocks = "{{ .ConfirmationBlocks }}"


#################################### NON CONFIGRABLE FIELDS BELOW #################################

##### Contract Addresses #####

### Main contracts ###
rollup_address = "{{ .RollupAddress }}"
token_registry_address = "{{ .TokenRegistry }}"
account_registry_address = "{{ .AccountRegistry }}"
deposit_manager_address = "{{ .DepositManager }}"
burn_auction_address = "{{ .BurnAuction }}"

### Client contracts ###
frontend_generic_address = "{{ .State }}"
transfer_address = "{{ .Transfer }}"
mass_migration_address = "{{ .MassMigration }}"
create2transfer_address = "{{ .Create2Transfer }}"

### Protocol Parameters
max_tree_depth = "{{ .MaxTreeDepth }}"
max_deposit_subtree = "{{ .MaxDepositSubtree }}"
stake_amount = "{{ .StakeAmount }}"
app_id = "{{ .AppID }}"

### Protocol Auxiliary Values
genesis_eth1_block = "{{ .GenesisEth1Block }}"
`

var configTemplate *template.Template

func init() {
	var err error
	tmpl := template.New("appConfigFileTemplate")
	if configTemplate, err = tmpl.Parse(defaultConfigTemplate); err != nil {
		panic(err)
	}
}
