// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logger\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATCH_DEPTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATCH_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITNESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accountIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32[31]\",\"name\":\"witness\",\"type\":\"bytes32[31]\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filledSubtreesLeft\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filledSubtreesRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leafIndexLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leafIndexRight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logger\",\"outputs\":[{\"internalType\":\"contractLogger\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[4][1024]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][1024]\"}],\"name\":\"registerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootLeft\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// BATCHDEPTH is a free data retrieval call binding the contract method 0x1c76e77e.
//
// Solidity: function BATCH_DEPTH() view returns(uint256)
func (_Registry *RegistryCaller) BATCHDEPTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "BATCH_DEPTH")
	return *ret0, err
}

// BATCHDEPTH is a free data retrieval call binding the contract method 0x1c76e77e.
//
// Solidity: function BATCH_DEPTH() view returns(uint256)
func (_Registry *RegistrySession) BATCHDEPTH() (*big.Int, error) {
	return _Registry.Contract.BATCHDEPTH(&_Registry.CallOpts)
}

// BATCHDEPTH is a free data retrieval call binding the contract method 0x1c76e77e.
//
// Solidity: function BATCH_DEPTH() view returns(uint256)
func (_Registry *RegistryCallerSession) BATCHDEPTH() (*big.Int, error) {
	return _Registry.Contract.BATCHDEPTH(&_Registry.CallOpts)
}

// BATCHSIZE is a free data retrieval call binding the contract method 0x49faa4d4.
//
// Solidity: function BATCH_SIZE() view returns(uint256)
func (_Registry *RegistryCaller) BATCHSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "BATCH_SIZE")
	return *ret0, err
}

// BATCHSIZE is a free data retrieval call binding the contract method 0x49faa4d4.
//
// Solidity: function BATCH_SIZE() view returns(uint256)
func (_Registry *RegistrySession) BATCHSIZE() (*big.Int, error) {
	return _Registry.Contract.BATCHSIZE(&_Registry.CallOpts)
}

// BATCHSIZE is a free data retrieval call binding the contract method 0x49faa4d4.
//
// Solidity: function BATCH_SIZE() view returns(uint256)
func (_Registry *RegistryCallerSession) BATCHSIZE() (*big.Int, error) {
	return _Registry.Contract.BATCHSIZE(&_Registry.CallOpts)
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Registry *RegistryCaller) DEPTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "DEPTH")
	return *ret0, err
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Registry *RegistrySession) DEPTH() (*big.Int, error) {
	return _Registry.Contract.DEPTH(&_Registry.CallOpts)
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Registry *RegistryCallerSession) DEPTH() (*big.Int, error) {
	return _Registry.Contract.DEPTH(&_Registry.CallOpts)
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Registry *RegistryCaller) SETSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "SET_SIZE")
	return *ret0, err
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Registry *RegistrySession) SETSIZE() (*big.Int, error) {
	return _Registry.Contract.SETSIZE(&_Registry.CallOpts)
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Registry *RegistryCallerSession) SETSIZE() (*big.Int, error) {
	return _Registry.Contract.SETSIZE(&_Registry.CallOpts)
}

// WITNESSLENGTH is a free data retrieval call binding the contract method 0x5e71468b.
//
// Solidity: function WITNESS_LENGTH() view returns(uint256)
func (_Registry *RegistryCaller) WITNESSLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "WITNESS_LENGTH")
	return *ret0, err
}

// WITNESSLENGTH is a free data retrieval call binding the contract method 0x5e71468b.
//
// Solidity: function WITNESS_LENGTH() view returns(uint256)
func (_Registry *RegistrySession) WITNESSLENGTH() (*big.Int, error) {
	return _Registry.Contract.WITNESSLENGTH(&_Registry.CallOpts)
}

// WITNESSLENGTH is a free data retrieval call binding the contract method 0x5e71468b.
//
// Solidity: function WITNESS_LENGTH() view returns(uint256)
func (_Registry *RegistryCallerSession) WITNESSLENGTH() (*big.Int, error) {
	return _Registry.Contract.WITNESSLENGTH(&_Registry.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x709a8b2a.
//
// Solidity: function exists(uint256 accountIndex, uint256[4] pubkey, bytes32[31] witness) view returns(bool)
func (_Registry *RegistryCaller) Exists(opts *bind.CallOpts, accountIndex *big.Int, pubkey [4]*big.Int, witness [31][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "exists", accountIndex, pubkey, witness)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x709a8b2a.
//
// Solidity: function exists(uint256 accountIndex, uint256[4] pubkey, bytes32[31] witness) view returns(bool)
func (_Registry *RegistrySession) Exists(accountIndex *big.Int, pubkey [4]*big.Int, witness [31][32]byte) (bool, error) {
	return _Registry.Contract.Exists(&_Registry.CallOpts, accountIndex, pubkey, witness)
}

// Exists is a free data retrieval call binding the contract method 0x709a8b2a.
//
// Solidity: function exists(uint256 accountIndex, uint256[4] pubkey, bytes32[31] witness) view returns(bool)
func (_Registry *RegistryCallerSession) Exists(accountIndex *big.Int, pubkey [4]*big.Int, witness [31][32]byte) (bool, error) {
	return _Registry.Contract.Exists(&_Registry.CallOpts, accountIndex, pubkey, witness)
}

// FilledSubtreesLeft is a free data retrieval call binding the contract method 0x034a29ae.
//
// Solidity: function filledSubtreesLeft(uint256 ) view returns(bytes32)
func (_Registry *RegistryCaller) FilledSubtreesLeft(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "filledSubtreesLeft", arg0)
	return *ret0, err
}

// FilledSubtreesLeft is a free data retrieval call binding the contract method 0x034a29ae.
//
// Solidity: function filledSubtreesLeft(uint256 ) view returns(bytes32)
func (_Registry *RegistrySession) FilledSubtreesLeft(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.FilledSubtreesLeft(&_Registry.CallOpts, arg0)
}

// FilledSubtreesLeft is a free data retrieval call binding the contract method 0x034a29ae.
//
// Solidity: function filledSubtreesLeft(uint256 ) view returns(bytes32)
func (_Registry *RegistryCallerSession) FilledSubtreesLeft(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.FilledSubtreesLeft(&_Registry.CallOpts, arg0)
}

// FilledSubtreesRight is a free data retrieval call binding the contract method 0x8d037962.
//
// Solidity: function filledSubtreesRight(uint256 ) view returns(bytes32)
func (_Registry *RegistryCaller) FilledSubtreesRight(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "filledSubtreesRight", arg0)
	return *ret0, err
}

// FilledSubtreesRight is a free data retrieval call binding the contract method 0x8d037962.
//
// Solidity: function filledSubtreesRight(uint256 ) view returns(bytes32)
func (_Registry *RegistrySession) FilledSubtreesRight(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.FilledSubtreesRight(&_Registry.CallOpts, arg0)
}

// FilledSubtreesRight is a free data retrieval call binding the contract method 0x8d037962.
//
// Solidity: function filledSubtreesRight(uint256 ) view returns(bytes32)
func (_Registry *RegistryCallerSession) FilledSubtreesRight(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.FilledSubtreesRight(&_Registry.CallOpts, arg0)
}

// LeafIndexLeft is a free data retrieval call binding the contract method 0x693c1db7.
//
// Solidity: function leafIndexLeft() view returns(uint256)
func (_Registry *RegistryCaller) LeafIndexLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "leafIndexLeft")
	return *ret0, err
}

// LeafIndexLeft is a free data retrieval call binding the contract method 0x693c1db7.
//
// Solidity: function leafIndexLeft() view returns(uint256)
func (_Registry *RegistrySession) LeafIndexLeft() (*big.Int, error) {
	return _Registry.Contract.LeafIndexLeft(&_Registry.CallOpts)
}

// LeafIndexLeft is a free data retrieval call binding the contract method 0x693c1db7.
//
// Solidity: function leafIndexLeft() view returns(uint256)
func (_Registry *RegistryCallerSession) LeafIndexLeft() (*big.Int, error) {
	return _Registry.Contract.LeafIndexLeft(&_Registry.CallOpts)
}

// LeafIndexRight is a free data retrieval call binding the contract method 0xd7c53ea7.
//
// Solidity: function leafIndexRight() view returns(uint256)
func (_Registry *RegistryCaller) LeafIndexRight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "leafIndexRight")
	return *ret0, err
}

// LeafIndexRight is a free data retrieval call binding the contract method 0xd7c53ea7.
//
// Solidity: function leafIndexRight() view returns(uint256)
func (_Registry *RegistrySession) LeafIndexRight() (*big.Int, error) {
	return _Registry.Contract.LeafIndexRight(&_Registry.CallOpts)
}

// LeafIndexRight is a free data retrieval call binding the contract method 0xd7c53ea7.
//
// Solidity: function leafIndexRight() view returns(uint256)
func (_Registry *RegistryCallerSession) LeafIndexRight() (*big.Int, error) {
	return _Registry.Contract.LeafIndexRight(&_Registry.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Registry *RegistryCaller) Logger(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "logger")
	return *ret0, err
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Registry *RegistrySession) Logger() (common.Address, error) {
	return _Registry.Contract.Logger(&_Registry.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Registry *RegistryCallerSession) Logger() (common.Address, error) {
	return _Registry.Contract.Logger(&_Registry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Registry *RegistryCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Registry *RegistrySession) Root() ([32]byte, error) {
	return _Registry.Contract.Root(&_Registry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Registry *RegistryCallerSession) Root() ([32]byte, error) {
	return _Registry.Contract.Root(&_Registry.CallOpts)
}

// RootLeft is a free data retrieval call binding the contract method 0xd8289463.
//
// Solidity: function rootLeft() view returns(bytes32)
func (_Registry *RegistryCaller) RootLeft(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "rootLeft")
	return *ret0, err
}

// RootLeft is a free data retrieval call binding the contract method 0xd8289463.
//
// Solidity: function rootLeft() view returns(bytes32)
func (_Registry *RegistrySession) RootLeft() ([32]byte, error) {
	return _Registry.Contract.RootLeft(&_Registry.CallOpts)
}

// RootLeft is a free data retrieval call binding the contract method 0xd8289463.
//
// Solidity: function rootLeft() view returns(bytes32)
func (_Registry *RegistryCallerSession) RootLeft() ([32]byte, error) {
	return _Registry.Contract.RootLeft(&_Registry.CallOpts)
}

// RootRight is a free data retrieval call binding the contract method 0xcab2da9b.
//
// Solidity: function rootRight() view returns(bytes32)
func (_Registry *RegistryCaller) RootRight(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "rootRight")
	return *ret0, err
}

// RootRight is a free data retrieval call binding the contract method 0xcab2da9b.
//
// Solidity: function rootRight() view returns(bytes32)
func (_Registry *RegistrySession) RootRight() ([32]byte, error) {
	return _Registry.Contract.RootRight(&_Registry.CallOpts)
}

// RootRight is a free data retrieval call binding the contract method 0xcab2da9b.
//
// Solidity: function rootRight() view returns(bytes32)
func (_Registry *RegistryCallerSession) RootRight() ([32]byte, error) {
	return _Registry.Contract.RootRight(&_Registry.CallOpts)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_Registry *RegistryCaller) Zeros(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "zeros", arg0)
	return *ret0, err
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_Registry *RegistrySession) Zeros(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.Zeros(&_Registry.CallOpts, arg0)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_Registry *RegistryCallerSession) Zeros(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.Zeros(&_Registry.CallOpts, arg0)
}

// Register is a paid mutator transaction binding the contract method 0x95e4bf03.
//
// Solidity: function register(uint256[4] pubkey) returns(uint256)
func (_Registry *RegistryTransactor) Register(opts *bind.TransactOpts, pubkey [4]*big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register", pubkey)
}

// Register is a paid mutator transaction binding the contract method 0x95e4bf03.
//
// Solidity: function register(uint256[4] pubkey) returns(uint256)
func (_Registry *RegistrySession) Register(pubkey [4]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, pubkey)
}

// Register is a paid mutator transaction binding the contract method 0x95e4bf03.
//
// Solidity: function register(uint256[4] pubkey) returns(uint256)
func (_Registry *RegistryTransactorSession) Register(pubkey [4]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, pubkey)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x4d18e08e.
//
// Solidity: function registerBatch(uint256[4][1024] pubkeys) returns(uint256)
func (_Registry *RegistryTransactor) RegisterBatch(opts *bind.TransactOpts, pubkeys [1024][4]*big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerBatch", pubkeys)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x4d18e08e.
//
// Solidity: function registerBatch(uint256[4][1024] pubkeys) returns(uint256)
func (_Registry *RegistrySession) RegisterBatch(pubkeys [1024][4]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RegisterBatch(&_Registry.TransactOpts, pubkeys)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x4d18e08e.
//
// Solidity: function registerBatch(uint256[4][1024] pubkeys) returns(uint256)
func (_Registry *RegistryTransactorSession) RegisterBatch(pubkeys [1024][4]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RegisterBatch(&_Registry.TransactOpts, pubkeys)
}
