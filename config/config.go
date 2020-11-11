package config

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"regexp"
	"strings"
	"time"

	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	cmn "github.com/tendermint/tendermint/libs/common"
)

const (
	DATABASENAME                = "hubble"
	DefaultMongoDB              = "mongodb://localhost:27017"
	DefaultDB                   = "mysql"
	DefaultDbUrlPrefix          = "mysql://root:root@(localhost:3306)"
	DefaultEthRPC               = "http://localhost:8545"
	DefaultPollingInterval      = 5 * time.Second
	DefaultSeverPort            = "8080"
	DefaultConfirmationBlocks   = 5
	DefaultDepositSubTreeHeight = 4
	DefaultMaxDepth             = 2
)

var GlobalCfg Configuration
var OperatorKey *ecdsa.PrivateKey
var OperatorPubKey *ecdsa.PublicKey
var OperatorAddress ethCmn.Address

// Configuration represents heimdall config
type Configuration struct {
	// DB related configs
	DB        string `mapstructure:"db_type"`
	DBURL     string `mapstructure:"db_url"`
	Trace     bool   `mapstructure:"trace"`
	DBLogMode bool   `mapstructure:"db_log_mode"`

	EthRPC             string        `mapstructure:"eth_RPC_URL"`
	PollingInterval    time.Duration `mapstructure:"polling_interval"`
	TxsPerBatch        uint64        `mapstructure:"txs_per_batch"`
	ServerPort         string        `mapstructure:"server_port"`
	ConfirmationBlocks uint64        `mapstructure:"confirmation_blocks"` // Number of blocks for confirmation

	RollupAddress   string `mapstructure:"rollup_address"`
	LoggerAddress   string `mapstructure:"logger_address"`
	FrontendAddress string `mapstructure:"frontend_address"`

	OperatorKey       string `mapstructure:"operator_key"`
	OperatorAddress   string `mapstructure:"operator_address"`
	LastRecordedBlock string `mapstructure:"last_recorded_block"`
}

// GetDefaultConfig returns the default configration options
func GetDefaultConfig() Configuration {
	return Configuration{
		DB:                 DefaultDB,
		DBURL:              GetDBURL(),
		Trace:              false,
		DBLogMode:          true,
		EthRPC:             DefaultEthRPC,
		TxsPerBatch:        2,
		PollingInterval:    DefaultPollingInterval,
		ServerPort:         DefaultSeverPort,
		ConfirmationBlocks: DefaultConfirmationBlocks,
		RollupAddress:      ethCmn.Address{}.String(),
		LoggerAddress:      ethCmn.Address{}.String(),
		FrontendAddress:    ethCmn.Address{}.String(),
		OperatorKey:        "",
		OperatorAddress:    "",
		LastRecordedBlock:  "0",
	}
}

// ParseConfig retrieves the default environment configuration for the
// application.
func ParseConfig(path string) (*Configuration, error) {
	conf := new(Configuration)
	v := viper.New()
	v.SetConfigName("config")
	if path == "" {
		v.AddConfigPath(".")
	} else {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return conf, err
	}
	err := v.Unmarshal(conf)

	return conf, err
}

func ParseAndInitGlobalConfig(path string) error {
	conf, err := ParseConfig(path)
	if err != nil {
		return err
	}
	GlobalCfg = *conf
	return nil
}

// FormattedDBURL returns formatted db url
func (c *Configuration) FormattedDBURL() string {
	re := regexp.MustCompile(`[a-z0-9]+://`)
	tokens := re.Split(c.DBURL, 2)
	return strings.Join(tokens, "")
}

// WriteConfigFile renders config using the template and writes it to
// configFilePath.
func WriteConfigFile(configFilePath string, config *Configuration) {
	var buffer bytes.Buffer

	if err := configTemplate.Execute(&buffer, config); err != nil {
		panic(err)
	}
	cmn.MustWriteFile(configFilePath, buffer.Bytes(), 0644)
}

// SetOperatorKey sets the operatorKeys which include
// the private key and pubkey globally
func SetOperatorKeys(privKeyStr string) error {
	privKeyBytes, err := hex.DecodeString(privKeyStr)
	if err != nil {
		return err
	}
	OperatorKey = crypto.ToECDSAUnsafe(privKeyBytes)
	publicKey := OperatorKey.Public()
	ecsdaPubKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	OperatorPubKey = ecsdaPubKey
	OperatorAddress = crypto.PubkeyToAddress(*OperatorPubKey)
	return nil
}

func GenOperatorKey() ([]byte, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return crypto.FromECDSA(privKey), nil
}

// PrivKeyToPubKey convert private key to public key
func PrivKeyStringToAddress(privKey string) (ethCmn.Address, error) {
	privKeyBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return ethCmn.Address{}, err
	}

	OperatorKey := crypto.ToECDSAUnsafe(privKeyBytes)
	publicKey := OperatorKey.Public()
	ecsdaPubKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return ethCmn.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return crypto.PubkeyToAddress(*ecsdaPubKey), nil
}

func GetDBURL() string {
	values := []string{DefaultDbUrlPrefix, "/", DATABASENAME, "?parseTime=true"}
	return strings.Join(values, "")
}
