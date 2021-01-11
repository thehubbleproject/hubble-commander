// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountregistry

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

// AccountregistryABI is the input ABI used to generate the binding from.
const AccountregistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"}],\"name\":\"PubkeyRegistered\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATCH_DEPTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATCH_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITNESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32[9]\",\"name\":\"witness\",\"type\":\"bytes32[9]\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filledSubtreesLeft\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filledSubtreesRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leafIndexLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leafIndexRight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[4][4]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][4]\"}],\"name\":\"registerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootLeft\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Accountregistry is an auto generated Go binding around an Ethereum contract.
type Accountregistry struct {
	AccountregistryCaller     // Read-only binding to the contract
	AccountregistryTransactor // Write-only binding to the contract
	AccountregistryFilterer   // Log filterer for contract events
}

// AccountregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountregistrySession struct {
	Contract     *Accountregistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountregistryCallerSession struct {
	Contract *AccountregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AccountregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountregistryTransactorSession struct {
	Contract     *AccountregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AccountregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountregistryRaw struct {
	Contract *Accountregistry // Generic contract binding to access the raw methods on
}

// AccountregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountregistryCallerRaw struct {
	Contract *AccountregistryCaller // Generic read-only contract binding to access the raw methods on
}

// AccountregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountregistryTransactorRaw struct {
	Contract *AccountregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountregistry creates a new instance of Accountregistry, bound to a specific deployed contract.
func NewAccountregistry(address common.Address, backend bind.ContractBackend) (*Accountregistry, error) {
	contract, err := bindAccountregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accountregistry{AccountregistryCaller: AccountregistryCaller{contract: contract}, AccountregistryTransactor: AccountregistryTransactor{contract: contract}, AccountregistryFilterer: AccountregistryFilterer{contract: contract}}, nil
}

// NewAccountregistryCaller creates a new read-only instance of Accountregistry, bound to a specific deployed contract.
func NewAccountregistryCaller(address common.Address, caller bind.ContractCaller) (*AccountregistryCaller, error) {
	contract, err := bindAccountregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountregistryCaller{contract: contract}, nil
}

// NewAccountregistryTransactor creates a new write-only instance of Accountregistry, bound to a specific deployed contract.
func NewAccountregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountregistryTransactor, error) {
	contract, err := bindAccountregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountregistryTransactor{contract: contract}, nil
}

// NewAccountregistryFilterer creates a new log filterer instance of Accountregistry, bound to a specific deployed contract.
func NewAccountregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountregistryFilterer, error) {
	contract, err := bindAccountregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountregistryFilterer{contract: contract}, nil
}

// bindAccountregistry binds a generic wrapper to an already deployed contract.
func bindAccountregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accountregistry *AccountregistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accountregistry.Contract.AccountregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accountregistry *AccountregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountregistry.Contract.AccountregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accountregistry *AccountregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accountregistry.Contract.AccountregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accountregistry *AccountregistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accountregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accountregistry *AccountregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accountregistry *AccountregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accountregistry.Contract.contract.Transact(opts, method, params...)
}

// BATCHDEPTH is a free data retrieval call binding the contract method 0x1c76e77e.
//
// Solidity: function BATCH_DEPTH() view returns(uint256)
func (_Accountregistry *AccountregistryCaller) BATCHDEPTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "BATCH_DEPTH")
	return *ret0, err
}

// BATCHDEPTH is a free data retrieval call binding the contract method 0x1c76e77e.
//
// Solidity: function BATCH_DEPTH() view returns(uint256)
func (_Accountregistry *AccountregistrySession) BATCHDEPTH() (*big.Int, error) {
	return _Accountregistry.Contract.BATCHDEPTH(&_Accountregistry.CallOpts)
}

// BATCHDEPTH is a free data retrieval call binding the contract method 0x1c76e77e.
//
// Solidity: function BATCH_DEPTH() view returns(uint256)
func (_Accountregistry *AccountregistryCallerSession) BATCHDEPTH() (*big.Int, error) {
	return _Accountregistry.Contract.BATCHDEPTH(&_Accountregistry.CallOpts)
}

// BATCHSIZE is a free data retrieval call binding the contract method 0x49faa4d4.
//
// Solidity: function BATCH_SIZE() view returns(uint256)
func (_Accountregistry *AccountregistryCaller) BATCHSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "BATCH_SIZE")
	return *ret0, err
}

// BATCHSIZE is a free data retrieval call binding the contract method 0x49faa4d4.
//
// Solidity: function BATCH_SIZE() view returns(uint256)
func (_Accountregistry *AccountregistrySession) BATCHSIZE() (*big.Int, error) {
	return _Accountregistry.Contract.BATCHSIZE(&_Accountregistry.CallOpts)
}

// BATCHSIZE is a free data retrieval call binding the contract method 0x49faa4d4.
//
// Solidity: function BATCH_SIZE() view returns(uint256)
func (_Accountregistry *AccountregistryCallerSession) BATCHSIZE() (*big.Int, error) {
	return _Accountregistry.Contract.BATCHSIZE(&_Accountregistry.CallOpts)
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Accountregistry *AccountregistryCaller) DEPTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "DEPTH")
	return *ret0, err
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Accountregistry *AccountregistrySession) DEPTH() (*big.Int, error) {
	return _Accountregistry.Contract.DEPTH(&_Accountregistry.CallOpts)
}

// DEPTH is a free data retrieval call binding the contract method 0x98366e35.
//
// Solidity: function DEPTH() view returns(uint256)
func (_Accountregistry *AccountregistryCallerSession) DEPTH() (*big.Int, error) {
	return _Accountregistry.Contract.DEPTH(&_Accountregistry.CallOpts)
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Accountregistry *AccountregistryCaller) SETSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "SET_SIZE")
	return *ret0, err
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Accountregistry *AccountregistrySession) SETSIZE() (*big.Int, error) {
	return _Accountregistry.Contract.SETSIZE(&_Accountregistry.CallOpts)
}

// SETSIZE is a free data retrieval call binding the contract method 0xd0383d68.
//
// Solidity: function SET_SIZE() view returns(uint256)
func (_Accountregistry *AccountregistryCallerSession) SETSIZE() (*big.Int, error) {
	return _Accountregistry.Contract.SETSIZE(&_Accountregistry.CallOpts)
}

// WITNESSLENGTH is a free data retrieval call binding the contract method 0x5e71468b.
//
// Solidity: function WITNESS_LENGTH() view returns(uint256)
func (_Accountregistry *AccountregistryCaller) WITNESSLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "WITNESS_LENGTH")
	return *ret0, err
}

// WITNESSLENGTH is a free data retrieval call binding the contract method 0x5e71468b.
//
// Solidity: function WITNESS_LENGTH() view returns(uint256)
func (_Accountregistry *AccountregistrySession) WITNESSLENGTH() (*big.Int, error) {
	return _Accountregistry.Contract.WITNESSLENGTH(&_Accountregistry.CallOpts)
}

// WITNESSLENGTH is a free data retrieval call binding the contract method 0x5e71468b.
//
// Solidity: function WITNESS_LENGTH() view returns(uint256)
func (_Accountregistry *AccountregistryCallerSession) WITNESSLENGTH() (*big.Int, error) {
	return _Accountregistry.Contract.WITNESSLENGTH(&_Accountregistry.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0xa61936eb.
//
// Solidity: function exists(uint256 pubkeyID, uint256[4] pubkey, bytes32[9] witness) view returns(bool)
func (_Accountregistry *AccountregistryCaller) Exists(opts *bind.CallOpts, pubkeyID *big.Int, pubkey [4]*big.Int, witness [9][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "exists", pubkeyID, pubkey, witness)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0xa61936eb.
//
// Solidity: function exists(uint256 pubkeyID, uint256[4] pubkey, bytes32[9] witness) view returns(bool)
func (_Accountregistry *AccountregistrySession) Exists(pubkeyID *big.Int, pubkey [4]*big.Int, witness [9][32]byte) (bool, error) {
	return _Accountregistry.Contract.Exists(&_Accountregistry.CallOpts, pubkeyID, pubkey, witness)
}

// Exists is a free data retrieval call binding the contract method 0xa61936eb.
//
// Solidity: function exists(uint256 pubkeyID, uint256[4] pubkey, bytes32[9] witness) view returns(bool)
func (_Accountregistry *AccountregistryCallerSession) Exists(pubkeyID *big.Int, pubkey [4]*big.Int, witness [9][32]byte) (bool, error) {
	return _Accountregistry.Contract.Exists(&_Accountregistry.CallOpts, pubkeyID, pubkey, witness)
}

// FilledSubtreesLeft is a free data retrieval call binding the contract method 0x034a29ae.
//
// Solidity: function filledSubtreesLeft(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistryCaller) FilledSubtreesLeft(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "filledSubtreesLeft", arg0)
	return *ret0, err
}

// FilledSubtreesLeft is a free data retrieval call binding the contract method 0x034a29ae.
//
// Solidity: function filledSubtreesLeft(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistrySession) FilledSubtreesLeft(arg0 *big.Int) ([32]byte, error) {
	return _Accountregistry.Contract.FilledSubtreesLeft(&_Accountregistry.CallOpts, arg0)
}

// FilledSubtreesLeft is a free data retrieval call binding the contract method 0x034a29ae.
//
// Solidity: function filledSubtreesLeft(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistryCallerSession) FilledSubtreesLeft(arg0 *big.Int) ([32]byte, error) {
	return _Accountregistry.Contract.FilledSubtreesLeft(&_Accountregistry.CallOpts, arg0)
}

// FilledSubtreesRight is a free data retrieval call binding the contract method 0x8d037962.
//
// Solidity: function filledSubtreesRight(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistryCaller) FilledSubtreesRight(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "filledSubtreesRight", arg0)
	return *ret0, err
}

// FilledSubtreesRight is a free data retrieval call binding the contract method 0x8d037962.
//
// Solidity: function filledSubtreesRight(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistrySession) FilledSubtreesRight(arg0 *big.Int) ([32]byte, error) {
	return _Accountregistry.Contract.FilledSubtreesRight(&_Accountregistry.CallOpts, arg0)
}

// FilledSubtreesRight is a free data retrieval call binding the contract method 0x8d037962.
//
// Solidity: function filledSubtreesRight(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistryCallerSession) FilledSubtreesRight(arg0 *big.Int) ([32]byte, error) {
	return _Accountregistry.Contract.FilledSubtreesRight(&_Accountregistry.CallOpts, arg0)
}

// LeafIndexLeft is a free data retrieval call binding the contract method 0x693c1db7.
//
// Solidity: function leafIndexLeft() view returns(uint256)
func (_Accountregistry *AccountregistryCaller) LeafIndexLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "leafIndexLeft")
	return *ret0, err
}

// LeafIndexLeft is a free data retrieval call binding the contract method 0x693c1db7.
//
// Solidity: function leafIndexLeft() view returns(uint256)
func (_Accountregistry *AccountregistrySession) LeafIndexLeft() (*big.Int, error) {
	return _Accountregistry.Contract.LeafIndexLeft(&_Accountregistry.CallOpts)
}

// LeafIndexLeft is a free data retrieval call binding the contract method 0x693c1db7.
//
// Solidity: function leafIndexLeft() view returns(uint256)
func (_Accountregistry *AccountregistryCallerSession) LeafIndexLeft() (*big.Int, error) {
	return _Accountregistry.Contract.LeafIndexLeft(&_Accountregistry.CallOpts)
}

// LeafIndexRight is a free data retrieval call binding the contract method 0xd7c53ea7.
//
// Solidity: function leafIndexRight() view returns(uint256)
func (_Accountregistry *AccountregistryCaller) LeafIndexRight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "leafIndexRight")
	return *ret0, err
}

// LeafIndexRight is a free data retrieval call binding the contract method 0xd7c53ea7.
//
// Solidity: function leafIndexRight() view returns(uint256)
func (_Accountregistry *AccountregistrySession) LeafIndexRight() (*big.Int, error) {
	return _Accountregistry.Contract.LeafIndexRight(&_Accountregistry.CallOpts)
}

// LeafIndexRight is a free data retrieval call binding the contract method 0xd7c53ea7.
//
// Solidity: function leafIndexRight() view returns(uint256)
func (_Accountregistry *AccountregistryCallerSession) LeafIndexRight() (*big.Int, error) {
	return _Accountregistry.Contract.LeafIndexRight(&_Accountregistry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Accountregistry *AccountregistryCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Accountregistry *AccountregistrySession) Root() ([32]byte, error) {
	return _Accountregistry.Contract.Root(&_Accountregistry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Accountregistry *AccountregistryCallerSession) Root() ([32]byte, error) {
	return _Accountregistry.Contract.Root(&_Accountregistry.CallOpts)
}

// RootLeft is a free data retrieval call binding the contract method 0xd8289463.
//
// Solidity: function rootLeft() view returns(bytes32)
func (_Accountregistry *AccountregistryCaller) RootLeft(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "rootLeft")
	return *ret0, err
}

// RootLeft is a free data retrieval call binding the contract method 0xd8289463.
//
// Solidity: function rootLeft() view returns(bytes32)
func (_Accountregistry *AccountregistrySession) RootLeft() ([32]byte, error) {
	return _Accountregistry.Contract.RootLeft(&_Accountregistry.CallOpts)
}

// RootLeft is a free data retrieval call binding the contract method 0xd8289463.
//
// Solidity: function rootLeft() view returns(bytes32)
func (_Accountregistry *AccountregistryCallerSession) RootLeft() ([32]byte, error) {
	return _Accountregistry.Contract.RootLeft(&_Accountregistry.CallOpts)
}

// RootRight is a free data retrieval call binding the contract method 0xcab2da9b.
//
// Solidity: function rootRight() view returns(bytes32)
func (_Accountregistry *AccountregistryCaller) RootRight(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "rootRight")
	return *ret0, err
}

// RootRight is a free data retrieval call binding the contract method 0xcab2da9b.
//
// Solidity: function rootRight() view returns(bytes32)
func (_Accountregistry *AccountregistrySession) RootRight() ([32]byte, error) {
	return _Accountregistry.Contract.RootRight(&_Accountregistry.CallOpts)
}

// RootRight is a free data retrieval call binding the contract method 0xcab2da9b.
//
// Solidity: function rootRight() view returns(bytes32)
func (_Accountregistry *AccountregistryCallerSession) RootRight() ([32]byte, error) {
	return _Accountregistry.Contract.RootRight(&_Accountregistry.CallOpts)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistryCaller) Zeros(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Accountregistry.contract.Call(opts, out, "zeros", arg0)
	return *ret0, err
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistrySession) Zeros(arg0 *big.Int) ([32]byte, error) {
	return _Accountregistry.Contract.Zeros(&_Accountregistry.CallOpts, arg0)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 ) view returns(bytes32)
func (_Accountregistry *AccountregistryCallerSession) Zeros(arg0 *big.Int) ([32]byte, error) {
	return _Accountregistry.Contract.Zeros(&_Accountregistry.CallOpts, arg0)
}

// Register is a paid mutator transaction binding the contract method 0x95e4bf03.
//
// Solidity: function register(uint256[4] pubkey) returns(uint256)
func (_Accountregistry *AccountregistryTransactor) Register(opts *bind.TransactOpts, pubkey [4]*big.Int) (*types.Transaction, error) {
	return _Accountregistry.contract.Transact(opts, "register", pubkey)
}

// Register is a paid mutator transaction binding the contract method 0x95e4bf03.
//
// Solidity: function register(uint256[4] pubkey) returns(uint256)
func (_Accountregistry *AccountregistrySession) Register(pubkey [4]*big.Int) (*types.Transaction, error) {
	return _Accountregistry.Contract.Register(&_Accountregistry.TransactOpts, pubkey)
}

// Register is a paid mutator transaction binding the contract method 0x95e4bf03.
//
// Solidity: function register(uint256[4] pubkey) returns(uint256)
func (_Accountregistry *AccountregistryTransactorSession) Register(pubkey [4]*big.Int) (*types.Transaction, error) {
	return _Accountregistry.Contract.Register(&_Accountregistry.TransactOpts, pubkey)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x70c462ac.
//
// Solidity: function registerBatch(uint256[4][4] pubkeys) returns(uint256)
func (_Accountregistry *AccountregistryTransactor) RegisterBatch(opts *bind.TransactOpts, pubkeys [4][4]*big.Int) (*types.Transaction, error) {
	return _Accountregistry.contract.Transact(opts, "registerBatch", pubkeys)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x70c462ac.
//
// Solidity: function registerBatch(uint256[4][4] pubkeys) returns(uint256)
func (_Accountregistry *AccountregistrySession) RegisterBatch(pubkeys [4][4]*big.Int) (*types.Transaction, error) {
	return _Accountregistry.Contract.RegisterBatch(&_Accountregistry.TransactOpts, pubkeys)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x70c462ac.
//
// Solidity: function registerBatch(uint256[4][4] pubkeys) returns(uint256)
func (_Accountregistry *AccountregistryTransactorSession) RegisterBatch(pubkeys [4][4]*big.Int) (*types.Transaction, error) {
	return _Accountregistry.Contract.RegisterBatch(&_Accountregistry.TransactOpts, pubkeys)
}

// AccountregistryPubkeyRegisteredIterator is returned from FilterPubkeyRegistered and is used to iterate over the raw logs and unpacked data for PubkeyRegistered events raised by the Accountregistry contract.
type AccountregistryPubkeyRegisteredIterator struct {
	Event *AccountregistryPubkeyRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountregistryPubkeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountregistryPubkeyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountregistryPubkeyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountregistryPubkeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountregistryPubkeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountregistryPubkeyRegistered represents a PubkeyRegistered event raised by the Accountregistry contract.
type AccountregistryPubkeyRegistered struct {
	Pubkey   [4]*big.Int
	PubkeyID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPubkeyRegistered is a free log retrieval operation binding the contract event 0xf0777e5cea47492e18df87dcc844efabdfad315d1a2b4883d87cc2b964eddff0.
//
// Solidity: event PubkeyRegistered(uint256[4] pubkey, uint256 pubkeyID)
func (_Accountregistry *AccountregistryFilterer) FilterPubkeyRegistered(opts *bind.FilterOpts) (*AccountregistryPubkeyRegisteredIterator, error) {

	logs, sub, err := _Accountregistry.contract.FilterLogs(opts, "PubkeyRegistered")
	if err != nil {
		return nil, err
	}
	return &AccountregistryPubkeyRegisteredIterator{contract: _Accountregistry.contract, event: "PubkeyRegistered", logs: logs, sub: sub}, nil
}

// WatchPubkeyRegistered is a free log subscription operation binding the contract event 0xf0777e5cea47492e18df87dcc844efabdfad315d1a2b4883d87cc2b964eddff0.
//
// Solidity: event PubkeyRegistered(uint256[4] pubkey, uint256 pubkeyID)
func (_Accountregistry *AccountregistryFilterer) WatchPubkeyRegistered(opts *bind.WatchOpts, sink chan<- *AccountregistryPubkeyRegistered) (event.Subscription, error) {

	logs, sub, err := _Accountregistry.contract.WatchLogs(opts, "PubkeyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountregistryPubkeyRegistered)
				if err := _Accountregistry.contract.UnpackLog(event, "PubkeyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePubkeyRegistered is a log parse operation binding the contract event 0xf0777e5cea47492e18df87dcc844efabdfad315d1a2b4883d87cc2b964eddff0.
//
// Solidity: event PubkeyRegistered(uint256[4] pubkey, uint256 pubkeyID)
func (_Accountregistry *AccountregistryFilterer) ParsePubkeyRegistered(log types.Log) (*AccountregistryPubkeyRegistered, error) {
	event := new(AccountregistryPubkeyRegistered)
	if err := _Accountregistry.contract.UnpackLog(event, "PubkeyRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
