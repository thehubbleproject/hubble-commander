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
const RollupclientABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"name\":\"applyTransferTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"senderState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiverState\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeCreate2Transfer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeMassMigration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeTransfer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"name\":\"processTransferTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"senderState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiverState\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

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

// ApplyTransferTx is a free data retrieval call binding the contract method 0x3a449293.
//
// Solidity: function applyTransferTx(bytes txBytes, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes senderState, bytes receiverState)
func (_Rollupclient *RollupclientCaller) ApplyTransferTx(opts *bind.CallOpts, txBytes []byte, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	SenderState   []byte
	ReceiverState []byte
}, error) {
	ret := new(struct {
		SenderState   []byte
		ReceiverState []byte
	})
	out := ret
	err := _Rollupclient.contract.Call(opts, out, "applyTransferTx", txBytes, from, to)
	return *ret, err
}

// ApplyTransferTx is a free data retrieval call binding the contract method 0x3a449293.
//
// Solidity: function applyTransferTx(bytes txBytes, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes senderState, bytes receiverState)
func (_Rollupclient *RollupclientSession) ApplyTransferTx(txBytes []byte, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	SenderState   []byte
	ReceiverState []byte
}, error) {
	return _Rollupclient.Contract.ApplyTransferTx(&_Rollupclient.CallOpts, txBytes, from, to)
}

// ApplyTransferTx is a free data retrieval call binding the contract method 0x3a449293.
//
// Solidity: function applyTransferTx(bytes txBytes, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes senderState, bytes receiverState)
func (_Rollupclient *RollupclientCallerSession) ApplyTransferTx(txBytes []byte, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	SenderState   []byte
	ReceiverState []byte
}, error) {
	return _Rollupclient.Contract.ApplyTransferTx(&_Rollupclient.CallOpts, txBytes, from, to)
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

// ProcessTransferTx is a free data retrieval call binding the contract method 0x8e62209e.
//
// Solidity: function processTransferTx(bytes32 stateRoot, bytes txBytes, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, bytes senderState, bytes receiverState, uint8 result)
func (_Rollupclient *RollupclientCaller) ProcessTransferTx(opts *bind.CallOpts, stateRoot [32]byte, txBytes []byte, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot       [32]byte
	SenderState   []byte
	ReceiverState []byte
	Result        uint8
}, error) {
	ret := new(struct {
		NewRoot       [32]byte
		SenderState   []byte
		ReceiverState []byte
		Result        uint8
	})
	out := ret
	err := _Rollupclient.contract.Call(opts, out, "processTransferTx", stateRoot, txBytes, from, to)
	return *ret, err
}

// ProcessTransferTx is a free data retrieval call binding the contract method 0x8e62209e.
//
// Solidity: function processTransferTx(bytes32 stateRoot, bytes txBytes, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, bytes senderState, bytes receiverState, uint8 result)
func (_Rollupclient *RollupclientSession) ProcessTransferTx(stateRoot [32]byte, txBytes []byte, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot       [32]byte
	SenderState   []byte
	ReceiverState []byte
	Result        uint8
}, error) {
	return _Rollupclient.Contract.ProcessTransferTx(&_Rollupclient.CallOpts, stateRoot, txBytes, from, to)
}

// ProcessTransferTx is a free data retrieval call binding the contract method 0x8e62209e.
//
// Solidity: function processTransferTx(bytes32 stateRoot, bytes txBytes, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, bytes senderState, bytes receiverState, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ProcessTransferTx(stateRoot [32]byte, txBytes []byte, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot       [32]byte
	SenderState   []byte
	ReceiverState []byte
	Result        uint8
}, error) {
	return _Rollupclient.Contract.ProcessTransferTx(&_Rollupclient.CallOpts, stateRoot, txBytes, from, to)
}
