// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollupclient

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OffchainCreate2Transfer is an auto generated low-level Go binding around an user-defined struct.
type OffchainCreate2Transfer struct {
	TxType    *big.Int
	FromIndex *big.Int
	ToIndex   *big.Int
	ToAccID   *big.Int
	Amount    *big.Int
	Fee       *big.Int
	Nonce     *big.Int
}

// OffchainMassMigration is an auto generated low-level Go binding around an user-defined struct.
type OffchainMassMigration struct {
	TxType    *big.Int
	FromIndex *big.Int
	Amount    *big.Int
	Fee       *big.Int
	SpokeID   *big.Int
	Nonce     *big.Int
}

// OffchainTransfer is an auto generated low-level Go binding around an user-defined struct.
type OffchainTransfer struct {
	TxType    *big.Int
	FromIndex *big.Int
	ToIndex   *big.Int
	Amount    *big.Int
	Fee       *big.Int
	Nonce     *big.Int
}

// TypesSignatureProof is an auto generated low-level Go binding around an user-defined struct.
type TypesSignatureProof struct {
	States          []TypesUserState
	StateWitnesses  [][][32]byte
	Pubkeys         [][4]*big.Int
	PubkeyWitnesses [][][32]byte
}

// TypesSignatureProofWithReceiver is an auto generated low-level Go binding around an user-defined struct.
type TypesSignatureProofWithReceiver struct {
	States                  []TypesUserState
	StateWitnesses          [][][32]byte
	PubkeysSender           [][4]*big.Int
	PubkeyWitnessesSender   [][][32]byte
	PubkeysReceiver         [][4]*big.Int
	PubkeyWitnessesReceiver [][][32]byte
}

// TypesStateMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type TypesStateMerkleProof struct {
	State   TypesUserState
	Witness [][32]byte
}

// TypesUserState is an auto generated low-level Go binding around an user-defined struct.
type TypesUserState struct {
	PubkeyIndex *big.Int
	TokenType   *big.Int
	Balance     *big.Int
	Nonce       *big.Int
}

// RollupclientABI is the input ABI used to generate the binding from.
const RollupclientABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeysSender\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnessesSender\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeysReceiver\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnessesReceiver\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProofWithReceiver\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"checkSignatureCreate2Transfer\",\"outputs\":[{\"internalType\":\"enumTypes.Result\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnesses\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"checkSignatureMassMigration\",\"outputs\":[{\"internalType\":\"enumTypes.Result\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnesses\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"checkSignatureTransfer\",\"outputs\":[{\"internalType\":\"enumTypes.Result\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"encodedTxs\",\"type\":\"bytes[]\"}],\"name\":\"compressCreate2Transfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"encodedTxs\",\"type\":\"bytes[]\"}],\"name\":\"compressMassMigration\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"encodedTxs\",\"type\":\"bytes[]\"}],\"name\":\"compressTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeCreate2Transfer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeMassMigration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"stateBytes\",\"type\":\"bytes\"}],\"name\":\"decodeState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeTransfer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"encode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"encodeCreate2Transfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"encodeMassMigration\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"encodeTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"name\":\"processCreate2Transfer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"}],\"name\":\"processMassMigration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"freshState\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"name\":\"processTransfer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"fromPubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"toPubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"valiateCreate2Transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"valiateMassMigration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"valiateTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"validateAndApplyCreate2Transfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newReceiver\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"validateAndApplyMassMigration\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"withdrawState\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiverEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"validateAndApplyTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newReceiver\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Rollupclient is an auto generated Go binding around an Ethereum contract.
type Rollupclient struct {
	RollupclientCaller     // Read-only binding to the contract
	RollupclientTransactor // Write-only binding to the contract
	RollupclientFilterer   // Log filterer for contract events
}

// RollupclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupclientSession struct {
	Contract     *Rollupclient     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupclientCallerSession struct {
	Contract *RollupclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RollupclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupclientTransactorSession struct {
	Contract     *RollupclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RollupclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupclientRaw struct {
	Contract *Rollupclient // Generic contract binding to access the raw methods on
}

// RollupclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupclientCallerRaw struct {
	Contract *RollupclientCaller // Generic read-only contract binding to access the raw methods on
}

// RollupclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupclientTransactorRaw struct {
	Contract *RollupclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupclient creates a new instance of Rollupclient, bound to a specific deployed contract.
func NewRollupclient(address common.Address, backend bind.ContractBackend) (*Rollupclient, error) {
	contract, err := bindRollupclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollupclient{RollupclientCaller: RollupclientCaller{contract: contract}, RollupclientTransactor: RollupclientTransactor{contract: contract}, RollupclientFilterer: RollupclientFilterer{contract: contract}}, nil
}

// NewRollupclientCaller creates a new read-only instance of Rollupclient, bound to a specific deployed contract.
func NewRollupclientCaller(address common.Address, caller bind.ContractCaller) (*RollupclientCaller, error) {
	contract, err := bindRollupclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupclientCaller{contract: contract}, nil
}

// NewRollupclientTransactor creates a new write-only instance of Rollupclient, bound to a specific deployed contract.
func NewRollupclientTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupclientTransactor, error) {
	contract, err := bindRollupclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupclientTransactor{contract: contract}, nil
}

// NewRollupclientFilterer creates a new log filterer instance of Rollupclient, bound to a specific deployed contract.
func NewRollupclientFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupclientFilterer, error) {
	contract, err := bindRollupclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupclientFilterer{contract: contract}, nil
}

// bindRollupclient binds a generic wrapper to an already deployed contract.
func bindRollupclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupclientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollupclient *RollupclientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollupclient.Contract.RollupclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollupclient *RollupclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollupclient.Contract.RollupclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollupclient *RollupclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollupclient.Contract.RollupclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollupclient *RollupclientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollupclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollupclient *RollupclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollupclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollupclient *RollupclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollupclient.Contract.contract.Transact(opts, method, params...)
}

// CheckSignatureCreate2Transfer is a free data retrieval call binding the contract method 0x249c2fbd.
//
// Solidity: function checkSignatureCreate2Transfer(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientCaller) CheckSignatureCreate2Transfer(opts *bind.CallOpts, signature [2]*big.Int, proof TypesSignatureProofWithReceiver, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "checkSignatureCreate2Transfer", signature, proof, stateRoot, accountRoot, domain, txs)
	return *ret0, err
}

// CheckSignatureCreate2Transfer is a free data retrieval call binding the contract method 0x249c2fbd.
//
// Solidity: function checkSignatureCreate2Transfer(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientSession) CheckSignatureCreate2Transfer(signature [2]*big.Int, proof TypesSignatureProofWithReceiver, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	return _Rollupclient.Contract.CheckSignatureCreate2Transfer(&_Rollupclient.CallOpts, signature, proof, stateRoot, accountRoot, domain, txs)
}

// CheckSignatureCreate2Transfer is a free data retrieval call binding the contract method 0x249c2fbd.
//
// Solidity: function checkSignatureCreate2Transfer(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientCallerSession) CheckSignatureCreate2Transfer(signature [2]*big.Int, proof TypesSignatureProofWithReceiver, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	return _Rollupclient.Contract.CheckSignatureCreate2Transfer(&_Rollupclient.CallOpts, signature, proof, stateRoot, accountRoot, domain, txs)
}

// CheckSignatureMassMigration is a free data retrieval call binding the contract method 0xe52e3c60.
//
// Solidity: function checkSignatureMassMigration(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, uint256 spokeID, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientCaller) CheckSignatureMassMigration(opts *bind.CallOpts, signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, spokeID *big.Int, txs []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "checkSignatureMassMigration", signature, proof, stateRoot, accountRoot, domain, spokeID, txs)
	return *ret0, err
}

// CheckSignatureMassMigration is a free data retrieval call binding the contract method 0xe52e3c60.
//
// Solidity: function checkSignatureMassMigration(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, uint256 spokeID, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientSession) CheckSignatureMassMigration(signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, spokeID *big.Int, txs []byte) (uint8, error) {
	return _Rollupclient.Contract.CheckSignatureMassMigration(&_Rollupclient.CallOpts, signature, proof, stateRoot, accountRoot, domain, spokeID, txs)
}

// CheckSignatureMassMigration is a free data retrieval call binding the contract method 0xe52e3c60.
//
// Solidity: function checkSignatureMassMigration(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, uint256 spokeID, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientCallerSession) CheckSignatureMassMigration(signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, spokeID *big.Int, txs []byte) (uint8, error) {
	return _Rollupclient.Contract.CheckSignatureMassMigration(&_Rollupclient.CallOpts, signature, proof, stateRoot, accountRoot, domain, spokeID, txs)
}

// CheckSignatureTransfer is a free data retrieval call binding the contract method 0xf6de1ce2.
//
// Solidity: function checkSignatureTransfer(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientCaller) CheckSignatureTransfer(opts *bind.CallOpts, signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "checkSignatureTransfer", signature, proof, stateRoot, accountRoot, domain, txs)
	return *ret0, err
}

// CheckSignatureTransfer is a free data retrieval call binding the contract method 0xf6de1ce2.
//
// Solidity: function checkSignatureTransfer(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientSession) CheckSignatureTransfer(signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	return _Rollupclient.Contract.CheckSignatureTransfer(&_Rollupclient.CallOpts, signature, proof, stateRoot, accountRoot, domain, txs)
}

// CheckSignatureTransfer is a free data retrieval call binding the contract method 0xf6de1ce2.
//
// Solidity: function checkSignatureTransfer(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Rollupclient *RollupclientCallerSession) CheckSignatureTransfer(signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	return _Rollupclient.Contract.CheckSignatureTransfer(&_Rollupclient.CallOpts, signature, proof, stateRoot, accountRoot, domain, txs)
}

// CompressCreate2Transfer is a free data retrieval call binding the contract method 0x647a6636.
//
// Solidity: function compressCreate2Transfer(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) CompressCreate2Transfer(opts *bind.CallOpts, encodedTxs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "compressCreate2Transfer", encodedTxs)
	return *ret0, err
}

// CompressCreate2Transfer is a free data retrieval call binding the contract method 0x647a6636.
//
// Solidity: function compressCreate2Transfer(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientSession) CompressCreate2Transfer(encodedTxs [][]byte) ([]byte, error) {
	return _Rollupclient.Contract.CompressCreate2Transfer(&_Rollupclient.CallOpts, encodedTxs)
}

// CompressCreate2Transfer is a free data retrieval call binding the contract method 0x647a6636.
//
// Solidity: function compressCreate2Transfer(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) CompressCreate2Transfer(encodedTxs [][]byte) ([]byte, error) {
	return _Rollupclient.Contract.CompressCreate2Transfer(&_Rollupclient.CallOpts, encodedTxs)
}

// CompressMassMigration is a free data retrieval call binding the contract method 0x79374249.
//
// Solidity: function compressMassMigration(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) CompressMassMigration(opts *bind.CallOpts, encodedTxs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "compressMassMigration", encodedTxs)
	return *ret0, err
}

// CompressMassMigration is a free data retrieval call binding the contract method 0x79374249.
//
// Solidity: function compressMassMigration(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientSession) CompressMassMigration(encodedTxs [][]byte) ([]byte, error) {
	return _Rollupclient.Contract.CompressMassMigration(&_Rollupclient.CallOpts, encodedTxs)
}

// CompressMassMigration is a free data retrieval call binding the contract method 0x79374249.
//
// Solidity: function compressMassMigration(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) CompressMassMigration(encodedTxs [][]byte) ([]byte, error) {
	return _Rollupclient.Contract.CompressMassMigration(&_Rollupclient.CallOpts, encodedTxs)
}

// CompressTransfer is a free data retrieval call binding the contract method 0xa98d6265.
//
// Solidity: function compressTransfer(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) CompressTransfer(opts *bind.CallOpts, encodedTxs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "compressTransfer", encodedTxs)
	return *ret0, err
}

// CompressTransfer is a free data retrieval call binding the contract method 0xa98d6265.
//
// Solidity: function compressTransfer(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientSession) CompressTransfer(encodedTxs [][]byte) ([]byte, error) {
	return _Rollupclient.Contract.CompressTransfer(&_Rollupclient.CallOpts, encodedTxs)
}

// CompressTransfer is a free data retrieval call binding the contract method 0xa98d6265.
//
// Solidity: function compressTransfer(bytes[] encodedTxs) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) CompressTransfer(encodedTxs [][]byte) ([]byte, error) {
	return _Rollupclient.Contract.CompressTransfer(&_Rollupclient.CallOpts, encodedTxs)
}

// DecodeCreate2Transfer is a free data retrieval call binding the contract method 0xdd9a65a6.
//
// Solidity: function decodeCreate2Transfer(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientCaller) DecodeCreate2Transfer(opts *bind.CallOpts, encodedTx []byte) (OffchainCreate2Transfer, error) {
	var (
		ret0 = new(OffchainCreate2Transfer)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "decodeCreate2Transfer", encodedTx)
	return *ret0, err
}

// DecodeCreate2Transfer is a free data retrieval call binding the contract method 0xdd9a65a6.
//
// Solidity: function decodeCreate2Transfer(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientSession) DecodeCreate2Transfer(encodedTx []byte) (OffchainCreate2Transfer, error) {
	return _Rollupclient.Contract.DecodeCreate2Transfer(&_Rollupclient.CallOpts, encodedTx)
}

// DecodeCreate2Transfer is a free data retrieval call binding the contract method 0xdd9a65a6.
//
// Solidity: function decodeCreate2Transfer(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientCallerSession) DecodeCreate2Transfer(encodedTx []byte) (OffchainCreate2Transfer, error) {
	return _Rollupclient.Contract.DecodeCreate2Transfer(&_Rollupclient.CallOpts, encodedTx)
}

// DecodeMassMigration is a free data retrieval call binding the contract method 0x7da53539.
//
// Solidity: function decodeMassMigration(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientCaller) DecodeMassMigration(opts *bind.CallOpts, encodedTx []byte) (OffchainMassMigration, error) {
	var (
		ret0 = new(OffchainMassMigration)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "decodeMassMigration", encodedTx)
	return *ret0, err
}

// DecodeMassMigration is a free data retrieval call binding the contract method 0x7da53539.
//
// Solidity: function decodeMassMigration(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientSession) DecodeMassMigration(encodedTx []byte) (OffchainMassMigration, error) {
	return _Rollupclient.Contract.DecodeMassMigration(&_Rollupclient.CallOpts, encodedTx)
}

// DecodeMassMigration is a free data retrieval call binding the contract method 0x7da53539.
//
// Solidity: function decodeMassMigration(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientCallerSession) DecodeMassMigration(encodedTx []byte) (OffchainMassMigration, error) {
	return _Rollupclient.Contract.DecodeMassMigration(&_Rollupclient.CallOpts, encodedTx)
}

// DecodeState is a free data retrieval call binding the contract method 0xb3b83621.
//
// Solidity: function decodeState(bytes stateBytes) pure returns((uint256,uint256,uint256,uint256) state)
func (_Rollupclient *RollupclientCaller) DecodeState(opts *bind.CallOpts, stateBytes []byte) (TypesUserState, error) {
	var (
		ret0 = new(TypesUserState)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "decodeState", stateBytes)
	return *ret0, err
}

// DecodeState is a free data retrieval call binding the contract method 0xb3b83621.
//
// Solidity: function decodeState(bytes stateBytes) pure returns((uint256,uint256,uint256,uint256) state)
func (_Rollupclient *RollupclientSession) DecodeState(stateBytes []byte) (TypesUserState, error) {
	return _Rollupclient.Contract.DecodeState(&_Rollupclient.CallOpts, stateBytes)
}

// DecodeState is a free data retrieval call binding the contract method 0xb3b83621.
//
// Solidity: function decodeState(bytes stateBytes) pure returns((uint256,uint256,uint256,uint256) state)
func (_Rollupclient *RollupclientCallerSession) DecodeState(stateBytes []byte) (TypesUserState, error) {
	return _Rollupclient.Contract.DecodeState(&_Rollupclient.CallOpts, stateBytes)
}

// DecodeTransfer is a free data retrieval call binding the contract method 0xdc02f456.
//
// Solidity: function decodeTransfer(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientCaller) DecodeTransfer(opts *bind.CallOpts, encodedTx []byte) (OffchainTransfer, error) {
	var (
		ret0 = new(OffchainTransfer)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "decodeTransfer", encodedTx)
	return *ret0, err
}

// DecodeTransfer is a free data retrieval call binding the contract method 0xdc02f456.
//
// Solidity: function decodeTransfer(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientSession) DecodeTransfer(encodedTx []byte) (OffchainTransfer, error) {
	return _Rollupclient.Contract.DecodeTransfer(&_Rollupclient.CallOpts, encodedTx)
}

// DecodeTransfer is a free data retrieval call binding the contract method 0xdc02f456.
//
// Solidity: function decodeTransfer(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Rollupclient *RollupclientCallerSession) DecodeTransfer(encodedTx []byte) (OffchainTransfer, error) {
	return _Rollupclient.Contract.DecodeTransfer(&_Rollupclient.CallOpts, encodedTx)
}

// Encode is a free data retrieval call binding the contract method 0x17412b8a.
//
// Solidity: function encode((uint256,uint256,uint256,uint256) state) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) Encode(opts *bind.CallOpts, state TypesUserState) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "encode", state)
	return *ret0, err
}

// Encode is a free data retrieval call binding the contract method 0x17412b8a.
//
// Solidity: function encode((uint256,uint256,uint256,uint256) state) pure returns(bytes)
func (_Rollupclient *RollupclientSession) Encode(state TypesUserState) ([]byte, error) {
	return _Rollupclient.Contract.Encode(&_Rollupclient.CallOpts, state)
}

// Encode is a free data retrieval call binding the contract method 0x17412b8a.
//
// Solidity: function encode((uint256,uint256,uint256,uint256) state) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) Encode(state TypesUserState) ([]byte, error) {
	return _Rollupclient.Contract.Encode(&_Rollupclient.CallOpts, state)
}

// EncodeCreate2Transfer is a free data retrieval call binding the contract method 0xbadcaf16.
//
// Solidity: function encodeCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) EncodeCreate2Transfer(opts *bind.CallOpts, _tx OffchainCreate2Transfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "encodeCreate2Transfer", _tx)
	return *ret0, err
}

// EncodeCreate2Transfer is a free data retrieval call binding the contract method 0xbadcaf16.
//
// Solidity: function encodeCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientSession) EncodeCreate2Transfer(_tx OffchainCreate2Transfer) ([]byte, error) {
	return _Rollupclient.Contract.EncodeCreate2Transfer(&_Rollupclient.CallOpts, _tx)
}

// EncodeCreate2Transfer is a free data retrieval call binding the contract method 0xbadcaf16.
//
// Solidity: function encodeCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) EncodeCreate2Transfer(_tx OffchainCreate2Transfer) ([]byte, error) {
	return _Rollupclient.Contract.EncodeCreate2Transfer(&_Rollupclient.CallOpts, _tx)
}

// EncodeMassMigration is a free data retrieval call binding the contract method 0xbb2420f4.
//
// Solidity: function encodeMassMigration((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) EncodeMassMigration(opts *bind.CallOpts, _tx OffchainMassMigration) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "encodeMassMigration", _tx)
	return *ret0, err
}

// EncodeMassMigration is a free data retrieval call binding the contract method 0xbb2420f4.
//
// Solidity: function encodeMassMigration((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientSession) EncodeMassMigration(_tx OffchainMassMigration) ([]byte, error) {
	return _Rollupclient.Contract.EncodeMassMigration(&_Rollupclient.CallOpts, _tx)
}

// EncodeMassMigration is a free data retrieval call binding the contract method 0xbb2420f4.
//
// Solidity: function encodeMassMigration((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) EncodeMassMigration(_tx OffchainMassMigration) ([]byte, error) {
	return _Rollupclient.Contract.EncodeMassMigration(&_Rollupclient.CallOpts, _tx)
}

// EncodeTransfer is a free data retrieval call binding the contract method 0x6712104b.
//
// Solidity: function encodeTransfer((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) EncodeTransfer(opts *bind.CallOpts, _tx OffchainTransfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "encodeTransfer", _tx)
	return *ret0, err
}

// EncodeTransfer is a free data retrieval call binding the contract method 0x6712104b.
//
// Solidity: function encodeTransfer((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientSession) EncodeTransfer(_tx OffchainTransfer) ([]byte, error) {
	return _Rollupclient.Contract.EncodeTransfer(&_Rollupclient.CallOpts, _tx)
}

// EncodeTransfer is a free data retrieval call binding the contract method 0x6712104b.
//
// Solidity: function encodeTransfer((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) EncodeTransfer(_tx OffchainTransfer) ([]byte, error) {
	return _Rollupclient.Contract.EncodeTransfer(&_Rollupclient.CallOpts, _tx)
}

// ProcessCreate2Transfer is a free data retrieval call binding the contract method 0x45b42c83.
//
// Solidity: function processCreate2Transfer(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Rollupclient *RollupclientCaller) ProcessCreate2Transfer(opts *bind.CallOpts, stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	ret := new(struct {
		NewRoot [32]byte
		Result  uint8
	})
	out := ret
	err := _Rollupclient.contract.Call(opts, out, "processCreate2Transfer", stateRoot, encodedTx, tokenType, from, to)
	return *ret, err
}

// ProcessCreate2Transfer is a free data retrieval call binding the contract method 0x45b42c83.
//
// Solidity: function processCreate2Transfer(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Rollupclient *RollupclientSession) ProcessCreate2Transfer(stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	return _Rollupclient.Contract.ProcessCreate2Transfer(&_Rollupclient.CallOpts, stateRoot, encodedTx, tokenType, from, to)
}

// ProcessCreate2Transfer is a free data retrieval call binding the contract method 0x45b42c83.
//
// Solidity: function processCreate2Transfer(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ProcessCreate2Transfer(stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	return _Rollupclient.Contract.ProcessCreate2Transfer(&_Rollupclient.CallOpts, stateRoot, encodedTx, tokenType, from, to)
}

// ProcessMassMigration is a free data retrieval call binding the contract method 0x38621a76.
//
// Solidity: function processMassMigration(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from) pure returns(bytes32 newRoot, bytes freshState, uint8 result)
func (_Rollupclient *RollupclientCaller) ProcessMassMigration(opts *bind.CallOpts, stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof) (struct {
	NewRoot    [32]byte
	FreshState []byte
	Result     uint8
}, error) {
	ret := new(struct {
		NewRoot    [32]byte
		FreshState []byte
		Result     uint8
	})
	out := ret
	err := _Rollupclient.contract.Call(opts, out, "processMassMigration", stateRoot, encodedTx, tokenType, from)
	return *ret, err
}

// ProcessMassMigration is a free data retrieval call binding the contract method 0x38621a76.
//
// Solidity: function processMassMigration(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from) pure returns(bytes32 newRoot, bytes freshState, uint8 result)
func (_Rollupclient *RollupclientSession) ProcessMassMigration(stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof) (struct {
	NewRoot    [32]byte
	FreshState []byte
	Result     uint8
}, error) {
	return _Rollupclient.Contract.ProcessMassMigration(&_Rollupclient.CallOpts, stateRoot, encodedTx, tokenType, from)
}

// ProcessMassMigration is a free data retrieval call binding the contract method 0x38621a76.
//
// Solidity: function processMassMigration(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from) pure returns(bytes32 newRoot, bytes freshState, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ProcessMassMigration(stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof) (struct {
	NewRoot    [32]byte
	FreshState []byte
	Result     uint8
}, error) {
	return _Rollupclient.Contract.ProcessMassMigration(&_Rollupclient.CallOpts, stateRoot, encodedTx, tokenType, from)
}

// ProcessTransfer is a free data retrieval call binding the contract method 0x5340fda8.
//
// Solidity: function processTransfer(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Rollupclient *RollupclientCaller) ProcessTransfer(opts *bind.CallOpts, stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	ret := new(struct {
		NewRoot [32]byte
		Result  uint8
	})
	out := ret
	err := _Rollupclient.contract.Call(opts, out, "processTransfer", stateRoot, encodedTx, tokenType, from, to)
	return *ret, err
}

// ProcessTransfer is a free data retrieval call binding the contract method 0x5340fda8.
//
// Solidity: function processTransfer(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Rollupclient *RollupclientSession) ProcessTransfer(stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	return _Rollupclient.Contract.ProcessTransfer(&_Rollupclient.CallOpts, stateRoot, encodedTx, tokenType, from, to)
}

// ProcessTransfer is a free data retrieval call binding the contract method 0x5340fda8.
//
// Solidity: function processTransfer(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ProcessTransfer(stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	return _Rollupclient.Contract.ProcessTransfer(&_Rollupclient.CallOpts, stateRoot, encodedTx, tokenType, from, to)
}

// ValiateCreate2Transfer is a free data retrieval call binding the contract method 0x984ca4bb.
//
// Solidity: function valiateCreate2Transfer(bytes encodedTx, uint256[2] signature, uint256[4] fromPubkey, uint256[4] toPubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCaller) ValiateCreate2Transfer(opts *bind.CallOpts, encodedTx []byte, signature [2]*big.Int, fromPubkey [4]*big.Int, toPubkey [4]*big.Int, domain [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Rollupclient.contract.Call(opts, out, "valiateCreate2Transfer", encodedTx, signature, fromPubkey, toPubkey, domain)
	return err
}

// ValiateCreate2Transfer is a free data retrieval call binding the contract method 0x984ca4bb.
//
// Solidity: function valiateCreate2Transfer(bytes encodedTx, uint256[2] signature, uint256[4] fromPubkey, uint256[4] toPubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientSession) ValiateCreate2Transfer(encodedTx []byte, signature [2]*big.Int, fromPubkey [4]*big.Int, toPubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateCreate2Transfer(&_Rollupclient.CallOpts, encodedTx, signature, fromPubkey, toPubkey, domain)
}

// ValiateCreate2Transfer is a free data retrieval call binding the contract method 0x984ca4bb.
//
// Solidity: function valiateCreate2Transfer(bytes encodedTx, uint256[2] signature, uint256[4] fromPubkey, uint256[4] toPubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCallerSession) ValiateCreate2Transfer(encodedTx []byte, signature [2]*big.Int, fromPubkey [4]*big.Int, toPubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateCreate2Transfer(&_Rollupclient.CallOpts, encodedTx, signature, fromPubkey, toPubkey, domain)
}

// ValiateMassMigration is a free data retrieval call binding the contract method 0x3b017bf4.
//
// Solidity: function valiateMassMigration(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCaller) ValiateMassMigration(opts *bind.CallOpts, encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Rollupclient.contract.Call(opts, out, "valiateMassMigration", encodedTx, signature, pubkey, domain)
	return err
}

// ValiateMassMigration is a free data retrieval call binding the contract method 0x3b017bf4.
//
// Solidity: function valiateMassMigration(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientSession) ValiateMassMigration(encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateMassMigration(&_Rollupclient.CallOpts, encodedTx, signature, pubkey, domain)
}

// ValiateMassMigration is a free data retrieval call binding the contract method 0x3b017bf4.
//
// Solidity: function valiateMassMigration(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCallerSession) ValiateMassMigration(encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateMassMigration(&_Rollupclient.CallOpts, encodedTx, signature, pubkey, domain)
}

// ValiateTransfer is a free data retrieval call binding the contract method 0x5aba4cf3.
//
// Solidity: function valiateTransfer(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCaller) ValiateTransfer(opts *bind.CallOpts, encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Rollupclient.contract.Call(opts, out, "valiateTransfer", encodedTx, signature, pubkey, domain)
	return err
}

// ValiateTransfer is a free data retrieval call binding the contract method 0x5aba4cf3.
//
// Solidity: function valiateTransfer(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientSession) ValiateTransfer(encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateTransfer(&_Rollupclient.CallOpts, encodedTx, signature, pubkey, domain)
}

// ValiateTransfer is a free data retrieval call binding the contract method 0x5aba4cf3.
//
// Solidity: function valiateTransfer(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCallerSession) ValiateTransfer(encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateTransfer(&_Rollupclient.CallOpts, encodedTx, signature, pubkey, domain)
}

// ValidateAndApplyCreate2Transfer is a free data retrieval call binding the contract method 0x323f6b7b.
//
// Solidity: function validateAndApplyCreate2Transfer(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientCaller) ValidateAndApplyCreate2Transfer(opts *bind.CallOpts, senderEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	ret := new(struct {
		NewSender   []byte
		NewReceiver []byte
		Result      uint8
	})
	out := ret
	err := _Rollupclient.contract.Call(opts, out, "validateAndApplyCreate2Transfer", senderEncoded, encodedTx)
	return *ret, err
}

// ValidateAndApplyCreate2Transfer is a free data retrieval call binding the contract method 0x323f6b7b.
//
// Solidity: function validateAndApplyCreate2Transfer(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientSession) ValidateAndApplyCreate2Transfer(senderEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyCreate2Transfer(&_Rollupclient.CallOpts, senderEncoded, encodedTx)
}

// ValidateAndApplyCreate2Transfer is a free data retrieval call binding the contract method 0x323f6b7b.
//
// Solidity: function validateAndApplyCreate2Transfer(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ValidateAndApplyCreate2Transfer(senderEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyCreate2Transfer(&_Rollupclient.CallOpts, senderEncoded, encodedTx)
}

// ValidateAndApplyMassMigration is a free data retrieval call binding the contract method 0x5d4ac153.
//
// Solidity: function validateAndApplyMassMigration(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes withdrawState, uint8 result)
func (_Rollupclient *RollupclientCaller) ValidateAndApplyMassMigration(opts *bind.CallOpts, senderEncoded []byte, encodedTx []byte) (struct {
	NewSender     []byte
	WithdrawState []byte
	Result        uint8
}, error) {
	ret := new(struct {
		NewSender     []byte
		WithdrawState []byte
		Result        uint8
	})
	out := ret
	err := _Rollupclient.contract.Call(opts, out, "validateAndApplyMassMigration", senderEncoded, encodedTx)
	return *ret, err
}

// ValidateAndApplyMassMigration is a free data retrieval call binding the contract method 0x5d4ac153.
//
// Solidity: function validateAndApplyMassMigration(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes withdrawState, uint8 result)
func (_Rollupclient *RollupclientSession) ValidateAndApplyMassMigration(senderEncoded []byte, encodedTx []byte) (struct {
	NewSender     []byte
	WithdrawState []byte
	Result        uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyMassMigration(&_Rollupclient.CallOpts, senderEncoded, encodedTx)
}

// ValidateAndApplyMassMigration is a free data retrieval call binding the contract method 0x5d4ac153.
//
// Solidity: function validateAndApplyMassMigration(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes withdrawState, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ValidateAndApplyMassMigration(senderEncoded []byte, encodedTx []byte) (struct {
	NewSender     []byte
	WithdrawState []byte
	Result        uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyMassMigration(&_Rollupclient.CallOpts, senderEncoded, encodedTx)
}

// ValidateAndApplyTransfer is a free data retrieval call binding the contract method 0xc5e411be.
//
// Solidity: function validateAndApplyTransfer(bytes senderEncoded, bytes receiverEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientCaller) ValidateAndApplyTransfer(opts *bind.CallOpts, senderEncoded []byte, receiverEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	ret := new(struct {
		NewSender   []byte
		NewReceiver []byte
		Result      uint8
	})
	out := ret
	err := _Rollupclient.contract.Call(opts, out, "validateAndApplyTransfer", senderEncoded, receiverEncoded, encodedTx)
	return *ret, err
}

// ValidateAndApplyTransfer is a free data retrieval call binding the contract method 0xc5e411be.
//
// Solidity: function validateAndApplyTransfer(bytes senderEncoded, bytes receiverEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientSession) ValidateAndApplyTransfer(senderEncoded []byte, receiverEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyTransfer(&_Rollupclient.CallOpts, senderEncoded, receiverEncoded, encodedTx)
}

// ValidateAndApplyTransfer is a free data retrieval call binding the contract method 0xc5e411be.
//
// Solidity: function validateAndApplyTransfer(bytes senderEncoded, bytes receiverEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ValidateAndApplyTransfer(senderEncoded []byte, receiverEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyTransfer(&_Rollupclient.CallOpts, senderEncoded, receiverEncoded, encodedTx)
}
