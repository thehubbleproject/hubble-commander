// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package logger

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

// LoggerABI is the input ABI used to generate the binding from.
const LoggerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"BatchRollback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"DepositSubTreeReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"depositSubTreeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pathToSubTree\",\"type\":\"uint256\"}],\"name\":\"DepositsFinalised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTypes.Usage\",\"name\":\"batchType\",\"type\":\"uint8\"}],\"name\":\"NewBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"}],\"name\":\"PubkeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"RegisteredToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"RegistrationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesSlashed\",\"type\":\"uint256\"}],\"name\":\"RollbackFinalisation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"StakeWithdraw\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"logBatchRollback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"depositSubTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pathToSubTree\",\"type\":\"uint256\"}],\"name\":\"logDepositFinalised\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accountID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"logDepositQueued\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"logDepositSubTreeReady\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"enumTypes.Usage\",\"name\":\"batchType\",\"type\":\"uint8\"}],\"name\":\"logNewBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"}],\"name\":\"logPubkeyRegistered\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"logRegisteredToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"logRegistrationRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalBatchesSlashed\",\"type\":\"uint256\"}],\"name\":\"logRollbackFinalisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"committed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"logStakeWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Logger is an auto generated Go binding around an Ethereum contract.
type Logger struct {
	LoggerCaller     // Read-only binding to the contract
	LoggerTransactor // Write-only binding to the contract
	LoggerFilterer   // Log filterer for contract events
}

// LoggerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoggerSession struct {
	Contract     *Logger           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoggerCallerSession struct {
	Contract *LoggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LoggerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoggerTransactorSession struct {
	Contract     *LoggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoggerRaw struct {
	Contract *Logger // Generic contract binding to access the raw methods on
}

// LoggerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoggerCallerRaw struct {
	Contract *LoggerCaller // Generic read-only contract binding to access the raw methods on
}

// LoggerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoggerTransactorRaw struct {
	Contract *LoggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogger creates a new instance of Logger, bound to a specific deployed contract.
func NewLogger(address common.Address, backend bind.ContractBackend) (*Logger, error) {
	contract, err := bindLogger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logger{LoggerCaller: LoggerCaller{contract: contract}, LoggerTransactor: LoggerTransactor{contract: contract}, LoggerFilterer: LoggerFilterer{contract: contract}}, nil
}

// NewLoggerCaller creates a new read-only instance of Logger, bound to a specific deployed contract.
func NewLoggerCaller(address common.Address, caller bind.ContractCaller) (*LoggerCaller, error) {
	contract, err := bindLogger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerCaller{contract: contract}, nil
}

// NewLoggerTransactor creates a new write-only instance of Logger, bound to a specific deployed contract.
func NewLoggerTransactor(address common.Address, transactor bind.ContractTransactor) (*LoggerTransactor, error) {
	contract, err := bindLogger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerTransactor{contract: contract}, nil
}

// NewLoggerFilterer creates a new log filterer instance of Logger, bound to a specific deployed contract.
func NewLoggerFilterer(address common.Address, filterer bind.ContractFilterer) (*LoggerFilterer, error) {
	contract, err := bindLogger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoggerFilterer{contract: contract}, nil
}

// bindLogger binds a generic wrapper to an already deployed contract.
func bindLogger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoggerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logger *LoggerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Logger.Contract.LoggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logger *LoggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logger.Contract.LoggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logger *LoggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logger.Contract.LoggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logger *LoggerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Logger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logger *LoggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logger *LoggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logger.Contract.contract.Transact(opts, method, params...)
}

// LogBatchRollback is a paid mutator transaction binding the contract method 0x1cb0fb2b.
//
// Solidity: function logBatchRollback(uint256 batchID) returns()
func (_Logger *LoggerTransactor) LogBatchRollback(opts *bind.TransactOpts, batchID *big.Int) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logBatchRollback", batchID)
}

// LogBatchRollback is a paid mutator transaction binding the contract method 0x1cb0fb2b.
//
// Solidity: function logBatchRollback(uint256 batchID) returns()
func (_Logger *LoggerSession) LogBatchRollback(batchID *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogBatchRollback(&_Logger.TransactOpts, batchID)
}

// LogBatchRollback is a paid mutator transaction binding the contract method 0x1cb0fb2b.
//
// Solidity: function logBatchRollback(uint256 batchID) returns()
func (_Logger *LoggerTransactorSession) LogBatchRollback(batchID *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogBatchRollback(&_Logger.TransactOpts, batchID)
}

// LogDepositFinalised is a paid mutator transaction binding the contract method 0xa0eccfa1.
//
// Solidity: function logDepositFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree) returns()
func (_Logger *LoggerTransactor) LogDepositFinalised(opts *bind.TransactOpts, depositSubTreeRoot [32]byte, pathToSubTree *big.Int) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logDepositFinalised", depositSubTreeRoot, pathToSubTree)
}

// LogDepositFinalised is a paid mutator transaction binding the contract method 0xa0eccfa1.
//
// Solidity: function logDepositFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree) returns()
func (_Logger *LoggerSession) LogDepositFinalised(depositSubTreeRoot [32]byte, pathToSubTree *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogDepositFinalised(&_Logger.TransactOpts, depositSubTreeRoot, pathToSubTree)
}

// LogDepositFinalised is a paid mutator transaction binding the contract method 0xa0eccfa1.
//
// Solidity: function logDepositFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree) returns()
func (_Logger *LoggerTransactorSession) LogDepositFinalised(depositSubTreeRoot [32]byte, pathToSubTree *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogDepositFinalised(&_Logger.TransactOpts, depositSubTreeRoot, pathToSubTree)
}

// LogDepositQueued is a paid mutator transaction binding the contract method 0xa557f4d0.
//
// Solidity: function logDepositQueued(uint256 accountID, bytes data) returns()
func (_Logger *LoggerTransactor) LogDepositQueued(opts *bind.TransactOpts, accountID *big.Int, data []byte) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logDepositQueued", accountID, data)
}

// LogDepositQueued is a paid mutator transaction binding the contract method 0xa557f4d0.
//
// Solidity: function logDepositQueued(uint256 accountID, bytes data) returns()
func (_Logger *LoggerSession) LogDepositQueued(accountID *big.Int, data []byte) (*types.Transaction, error) {
	return _Logger.Contract.LogDepositQueued(&_Logger.TransactOpts, accountID, data)
}

// LogDepositQueued is a paid mutator transaction binding the contract method 0xa557f4d0.
//
// Solidity: function logDepositQueued(uint256 accountID, bytes data) returns()
func (_Logger *LoggerTransactorSession) LogDepositQueued(accountID *big.Int, data []byte) (*types.Transaction, error) {
	return _Logger.Contract.LogDepositQueued(&_Logger.TransactOpts, accountID, data)
}

// LogDepositSubTreeReady is a paid mutator transaction binding the contract method 0x346300bc.
//
// Solidity: function logDepositSubTreeReady(bytes32 root) returns()
func (_Logger *LoggerTransactor) LogDepositSubTreeReady(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logDepositSubTreeReady", root)
}

// LogDepositSubTreeReady is a paid mutator transaction binding the contract method 0x346300bc.
//
// Solidity: function logDepositSubTreeReady(bytes32 root) returns()
func (_Logger *LoggerSession) LogDepositSubTreeReady(root [32]byte) (*types.Transaction, error) {
	return _Logger.Contract.LogDepositSubTreeReady(&_Logger.TransactOpts, root)
}

// LogDepositSubTreeReady is a paid mutator transaction binding the contract method 0x346300bc.
//
// Solidity: function logDepositSubTreeReady(bytes32 root) returns()
func (_Logger *LoggerTransactorSession) LogDepositSubTreeReady(root [32]byte) (*types.Transaction, error) {
	return _Logger.Contract.LogDepositSubTreeReady(&_Logger.TransactOpts, root)
}

// LogNewBatch is a paid mutator transaction binding the contract method 0xdb64014c.
//
// Solidity: function logNewBatch(address committer, uint256 index, uint8 batchType) returns()
func (_Logger *LoggerTransactor) LogNewBatch(opts *bind.TransactOpts, committer common.Address, index *big.Int, batchType uint8) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logNewBatch", committer, index, batchType)
}

// LogNewBatch is a paid mutator transaction binding the contract method 0xdb64014c.
//
// Solidity: function logNewBatch(address committer, uint256 index, uint8 batchType) returns()
func (_Logger *LoggerSession) LogNewBatch(committer common.Address, index *big.Int, batchType uint8) (*types.Transaction, error) {
	return _Logger.Contract.LogNewBatch(&_Logger.TransactOpts, committer, index, batchType)
}

// LogNewBatch is a paid mutator transaction binding the contract method 0xdb64014c.
//
// Solidity: function logNewBatch(address committer, uint256 index, uint8 batchType) returns()
func (_Logger *LoggerTransactorSession) LogNewBatch(committer common.Address, index *big.Int, batchType uint8) (*types.Transaction, error) {
	return _Logger.Contract.LogNewBatch(&_Logger.TransactOpts, committer, index, batchType)
}

// LogPubkeyRegistered is a paid mutator transaction binding the contract method 0x1d290ccf.
//
// Solidity: function logPubkeyRegistered(uint256[4] pubkey, uint256 pubkeyID) returns()
func (_Logger *LoggerTransactor) LogPubkeyRegistered(opts *bind.TransactOpts, pubkey [4]*big.Int, pubkeyID *big.Int) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logPubkeyRegistered", pubkey, pubkeyID)
}

// LogPubkeyRegistered is a paid mutator transaction binding the contract method 0x1d290ccf.
//
// Solidity: function logPubkeyRegistered(uint256[4] pubkey, uint256 pubkeyID) returns()
func (_Logger *LoggerSession) LogPubkeyRegistered(pubkey [4]*big.Int, pubkeyID *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogPubkeyRegistered(&_Logger.TransactOpts, pubkey, pubkeyID)
}

// LogPubkeyRegistered is a paid mutator transaction binding the contract method 0x1d290ccf.
//
// Solidity: function logPubkeyRegistered(uint256[4] pubkey, uint256 pubkeyID) returns()
func (_Logger *LoggerTransactorSession) LogPubkeyRegistered(pubkey [4]*big.Int, pubkeyID *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogPubkeyRegistered(&_Logger.TransactOpts, pubkey, pubkeyID)
}

// LogRegisteredToken is a paid mutator transaction binding the contract method 0x88870639.
//
// Solidity: function logRegisteredToken(uint256 tokenType, address tokenContract) returns()
func (_Logger *LoggerTransactor) LogRegisteredToken(opts *bind.TransactOpts, tokenType *big.Int, tokenContract common.Address) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logRegisteredToken", tokenType, tokenContract)
}

// LogRegisteredToken is a paid mutator transaction binding the contract method 0x88870639.
//
// Solidity: function logRegisteredToken(uint256 tokenType, address tokenContract) returns()
func (_Logger *LoggerSession) LogRegisteredToken(tokenType *big.Int, tokenContract common.Address) (*types.Transaction, error) {
	return _Logger.Contract.LogRegisteredToken(&_Logger.TransactOpts, tokenType, tokenContract)
}

// LogRegisteredToken is a paid mutator transaction binding the contract method 0x88870639.
//
// Solidity: function logRegisteredToken(uint256 tokenType, address tokenContract) returns()
func (_Logger *LoggerTransactorSession) LogRegisteredToken(tokenType *big.Int, tokenContract common.Address) (*types.Transaction, error) {
	return _Logger.Contract.LogRegisteredToken(&_Logger.TransactOpts, tokenType, tokenContract)
}

// LogRegistrationRequest is a paid mutator transaction binding the contract method 0x4c7637d2.
//
// Solidity: function logRegistrationRequest(address tokenContract) returns()
func (_Logger *LoggerTransactor) LogRegistrationRequest(opts *bind.TransactOpts, tokenContract common.Address) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logRegistrationRequest", tokenContract)
}

// LogRegistrationRequest is a paid mutator transaction binding the contract method 0x4c7637d2.
//
// Solidity: function logRegistrationRequest(address tokenContract) returns()
func (_Logger *LoggerSession) LogRegistrationRequest(tokenContract common.Address) (*types.Transaction, error) {
	return _Logger.Contract.LogRegistrationRequest(&_Logger.TransactOpts, tokenContract)
}

// LogRegistrationRequest is a paid mutator transaction binding the contract method 0x4c7637d2.
//
// Solidity: function logRegistrationRequest(address tokenContract) returns()
func (_Logger *LoggerTransactorSession) LogRegistrationRequest(tokenContract common.Address) (*types.Transaction, error) {
	return _Logger.Contract.LogRegistrationRequest(&_Logger.TransactOpts, tokenContract)
}

// LogRollbackFinalisation is a paid mutator transaction binding the contract method 0xaeedf0c7.
//
// Solidity: function logRollbackFinalisation(uint256 totalBatchesSlashed) returns()
func (_Logger *LoggerTransactor) LogRollbackFinalisation(opts *bind.TransactOpts, totalBatchesSlashed *big.Int) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logRollbackFinalisation", totalBatchesSlashed)
}

// LogRollbackFinalisation is a paid mutator transaction binding the contract method 0xaeedf0c7.
//
// Solidity: function logRollbackFinalisation(uint256 totalBatchesSlashed) returns()
func (_Logger *LoggerSession) LogRollbackFinalisation(totalBatchesSlashed *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogRollbackFinalisation(&_Logger.TransactOpts, totalBatchesSlashed)
}

// LogRollbackFinalisation is a paid mutator transaction binding the contract method 0xaeedf0c7.
//
// Solidity: function logRollbackFinalisation(uint256 totalBatchesSlashed) returns()
func (_Logger *LoggerTransactorSession) LogRollbackFinalisation(totalBatchesSlashed *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogRollbackFinalisation(&_Logger.TransactOpts, totalBatchesSlashed)
}

// LogStakeWithdraw is a paid mutator transaction binding the contract method 0x1f09cbd4.
//
// Solidity: function logStakeWithdraw(address committed, uint256 batchID) returns()
func (_Logger *LoggerTransactor) LogStakeWithdraw(opts *bind.TransactOpts, committed common.Address, batchID *big.Int) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "logStakeWithdraw", committed, batchID)
}

// LogStakeWithdraw is a paid mutator transaction binding the contract method 0x1f09cbd4.
//
// Solidity: function logStakeWithdraw(address committed, uint256 batchID) returns()
func (_Logger *LoggerSession) LogStakeWithdraw(committed common.Address, batchID *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogStakeWithdraw(&_Logger.TransactOpts, committed, batchID)
}

// LogStakeWithdraw is a paid mutator transaction binding the contract method 0x1f09cbd4.
//
// Solidity: function logStakeWithdraw(address committed, uint256 batchID) returns()
func (_Logger *LoggerTransactorSession) LogStakeWithdraw(committed common.Address, batchID *big.Int) (*types.Transaction, error) {
	return _Logger.Contract.LogStakeWithdraw(&_Logger.TransactOpts, committed, batchID)
}

// LoggerBatchRollbackIterator is returned from FilterBatchRollback and is used to iterate over the raw logs and unpacked data for BatchRollback events raised by the Logger contract.
type LoggerBatchRollbackIterator struct {
	Event *LoggerBatchRollback // Event containing the contract specifics and raw log

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
func (it *LoggerBatchRollbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerBatchRollback)
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
		it.Event = new(LoggerBatchRollback)
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
func (it *LoggerBatchRollbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerBatchRollbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerBatchRollback represents a BatchRollback event raised by the Logger contract.
type LoggerBatchRollback struct {
	BatchID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBatchRollback is a free log retrieval operation binding the contract event 0xd1273cb450780431aa88c2a774163d0be3a1e505ff453810334d2151097c6dbe.
//
// Solidity: event BatchRollback(uint256 batchID)
func (_Logger *LoggerFilterer) FilterBatchRollback(opts *bind.FilterOpts) (*LoggerBatchRollbackIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "BatchRollback")
	if err != nil {
		return nil, err
	}
	return &LoggerBatchRollbackIterator{contract: _Logger.contract, event: "BatchRollback", logs: logs, sub: sub}, nil
}

// WatchBatchRollback is a free log subscription operation binding the contract event 0xd1273cb450780431aa88c2a774163d0be3a1e505ff453810334d2151097c6dbe.
//
// Solidity: event BatchRollback(uint256 batchID)
func (_Logger *LoggerFilterer) WatchBatchRollback(opts *bind.WatchOpts, sink chan<- *LoggerBatchRollback) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "BatchRollback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerBatchRollback)
				if err := _Logger.contract.UnpackLog(event, "BatchRollback", log); err != nil {
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

// ParseBatchRollback is a log parse operation binding the contract event 0xd1273cb450780431aa88c2a774163d0be3a1e505ff453810334d2151097c6dbe.
//
// Solidity: event BatchRollback(uint256 batchID)
func (_Logger *LoggerFilterer) ParseBatchRollback(log types.Log) (*LoggerBatchRollback, error) {
	event := new(LoggerBatchRollback)
	if err := _Logger.contract.UnpackLog(event, "BatchRollback", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerDepositQueuedIterator is returned from FilterDepositQueued and is used to iterate over the raw logs and unpacked data for DepositQueued events raised by the Logger contract.
type LoggerDepositQueuedIterator struct {
	Event *LoggerDepositQueued // Event containing the contract specifics and raw log

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
func (it *LoggerDepositQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerDepositQueued)
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
		it.Event = new(LoggerDepositQueued)
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
func (it *LoggerDepositQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerDepositQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerDepositQueued represents a DepositQueued event raised by the Logger contract.
type LoggerDepositQueued struct {
	PubkeyID *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositQueued is a free log retrieval operation binding the contract event 0x5a1922090c12e28c132a961f6bb4d74350598a62e8981b5eb9bb5ccbbce9df94.
//
// Solidity: event DepositQueued(uint256 pubkeyID, bytes data)
func (_Logger *LoggerFilterer) FilterDepositQueued(opts *bind.FilterOpts) (*LoggerDepositQueuedIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "DepositQueued")
	if err != nil {
		return nil, err
	}
	return &LoggerDepositQueuedIterator{contract: _Logger.contract, event: "DepositQueued", logs: logs, sub: sub}, nil
}

// WatchDepositQueued is a free log subscription operation binding the contract event 0x5a1922090c12e28c132a961f6bb4d74350598a62e8981b5eb9bb5ccbbce9df94.
//
// Solidity: event DepositQueued(uint256 pubkeyID, bytes data)
func (_Logger *LoggerFilterer) WatchDepositQueued(opts *bind.WatchOpts, sink chan<- *LoggerDepositQueued) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "DepositQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerDepositQueued)
				if err := _Logger.contract.UnpackLog(event, "DepositQueued", log); err != nil {
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

// ParseDepositQueued is a log parse operation binding the contract event 0x5a1922090c12e28c132a961f6bb4d74350598a62e8981b5eb9bb5ccbbce9df94.
//
// Solidity: event DepositQueued(uint256 pubkeyID, bytes data)
func (_Logger *LoggerFilterer) ParseDepositQueued(log types.Log) (*LoggerDepositQueued, error) {
	event := new(LoggerDepositQueued)
	if err := _Logger.contract.UnpackLog(event, "DepositQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerDepositSubTreeReadyIterator is returned from FilterDepositSubTreeReady and is used to iterate over the raw logs and unpacked data for DepositSubTreeReady events raised by the Logger contract.
type LoggerDepositSubTreeReadyIterator struct {
	Event *LoggerDepositSubTreeReady // Event containing the contract specifics and raw log

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
func (it *LoggerDepositSubTreeReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerDepositSubTreeReady)
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
		it.Event = new(LoggerDepositSubTreeReady)
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
func (it *LoggerDepositSubTreeReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerDepositSubTreeReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerDepositSubTreeReady represents a DepositSubTreeReady event raised by the Logger contract.
type LoggerDepositSubTreeReady struct {
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDepositSubTreeReady is a free log retrieval operation binding the contract event 0x4744f3a44c5716c9fa423a71cdaa806771a8bf469f4c007ca338b8e8e202a8b5.
//
// Solidity: event DepositSubTreeReady(bytes32 root)
func (_Logger *LoggerFilterer) FilterDepositSubTreeReady(opts *bind.FilterOpts) (*LoggerDepositSubTreeReadyIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "DepositSubTreeReady")
	if err != nil {
		return nil, err
	}
	return &LoggerDepositSubTreeReadyIterator{contract: _Logger.contract, event: "DepositSubTreeReady", logs: logs, sub: sub}, nil
}

// WatchDepositSubTreeReady is a free log subscription operation binding the contract event 0x4744f3a44c5716c9fa423a71cdaa806771a8bf469f4c007ca338b8e8e202a8b5.
//
// Solidity: event DepositSubTreeReady(bytes32 root)
func (_Logger *LoggerFilterer) WatchDepositSubTreeReady(opts *bind.WatchOpts, sink chan<- *LoggerDepositSubTreeReady) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "DepositSubTreeReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerDepositSubTreeReady)
				if err := _Logger.contract.UnpackLog(event, "DepositSubTreeReady", log); err != nil {
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

// ParseDepositSubTreeReady is a log parse operation binding the contract event 0x4744f3a44c5716c9fa423a71cdaa806771a8bf469f4c007ca338b8e8e202a8b5.
//
// Solidity: event DepositSubTreeReady(bytes32 root)
func (_Logger *LoggerFilterer) ParseDepositSubTreeReady(log types.Log) (*LoggerDepositSubTreeReady, error) {
	event := new(LoggerDepositSubTreeReady)
	if err := _Logger.contract.UnpackLog(event, "DepositSubTreeReady", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerDepositsFinalisedIterator is returned from FilterDepositsFinalised and is used to iterate over the raw logs and unpacked data for DepositsFinalised events raised by the Logger contract.
type LoggerDepositsFinalisedIterator struct {
	Event *LoggerDepositsFinalised // Event containing the contract specifics and raw log

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
func (it *LoggerDepositsFinalisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerDepositsFinalised)
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
		it.Event = new(LoggerDepositsFinalised)
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
func (it *LoggerDepositsFinalisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerDepositsFinalisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerDepositsFinalised represents a DepositsFinalised event raised by the Logger contract.
type LoggerDepositsFinalised struct {
	DepositSubTreeRoot [32]byte
	PathToSubTree      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositsFinalised is a free log retrieval operation binding the contract event 0xd1d49921e88d39febcc645126f95d5eb4fab4d9e436d650983b7221eb7bf5c28.
//
// Solidity: event DepositsFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree)
func (_Logger *LoggerFilterer) FilterDepositsFinalised(opts *bind.FilterOpts) (*LoggerDepositsFinalisedIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "DepositsFinalised")
	if err != nil {
		return nil, err
	}
	return &LoggerDepositsFinalisedIterator{contract: _Logger.contract, event: "DepositsFinalised", logs: logs, sub: sub}, nil
}

// WatchDepositsFinalised is a free log subscription operation binding the contract event 0xd1d49921e88d39febcc645126f95d5eb4fab4d9e436d650983b7221eb7bf5c28.
//
// Solidity: event DepositsFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree)
func (_Logger *LoggerFilterer) WatchDepositsFinalised(opts *bind.WatchOpts, sink chan<- *LoggerDepositsFinalised) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "DepositsFinalised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerDepositsFinalised)
				if err := _Logger.contract.UnpackLog(event, "DepositsFinalised", log); err != nil {
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

// ParseDepositsFinalised is a log parse operation binding the contract event 0xd1d49921e88d39febcc645126f95d5eb4fab4d9e436d650983b7221eb7bf5c28.
//
// Solidity: event DepositsFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree)
func (_Logger *LoggerFilterer) ParseDepositsFinalised(log types.Log) (*LoggerDepositsFinalised, error) {
	event := new(LoggerDepositsFinalised)
	if err := _Logger.contract.UnpackLog(event, "DepositsFinalised", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerNewBatchIterator is returned from FilterNewBatch and is used to iterate over the raw logs and unpacked data for NewBatch events raised by the Logger contract.
type LoggerNewBatchIterator struct {
	Event *LoggerNewBatch // Event containing the contract specifics and raw log

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
func (it *LoggerNewBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerNewBatch)
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
		it.Event = new(LoggerNewBatch)
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
func (it *LoggerNewBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerNewBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerNewBatch represents a NewBatch event raised by the Logger contract.
type LoggerNewBatch struct {
	Committer common.Address
	Index     *big.Int
	BatchType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBatch is a free log retrieval operation binding the contract event 0xb34aa33e0e9ecb485e5d2fe496d4135cbedda94674af8795691fe0877da3916f.
//
// Solidity: event NewBatch(address committer, uint256 index, uint8 batchType)
func (_Logger *LoggerFilterer) FilterNewBatch(opts *bind.FilterOpts) (*LoggerNewBatchIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return &LoggerNewBatchIterator{contract: _Logger.contract, event: "NewBatch", logs: logs, sub: sub}, nil
}

// WatchNewBatch is a free log subscription operation binding the contract event 0xb34aa33e0e9ecb485e5d2fe496d4135cbedda94674af8795691fe0877da3916f.
//
// Solidity: event NewBatch(address committer, uint256 index, uint8 batchType)
func (_Logger *LoggerFilterer) WatchNewBatch(opts *bind.WatchOpts, sink chan<- *LoggerNewBatch) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerNewBatch)
				if err := _Logger.contract.UnpackLog(event, "NewBatch", log); err != nil {
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

// ParseNewBatch is a log parse operation binding the contract event 0xb34aa33e0e9ecb485e5d2fe496d4135cbedda94674af8795691fe0877da3916f.
//
// Solidity: event NewBatch(address committer, uint256 index, uint8 batchType)
func (_Logger *LoggerFilterer) ParseNewBatch(log types.Log) (*LoggerNewBatch, error) {
	event := new(LoggerNewBatch)
	if err := _Logger.contract.UnpackLog(event, "NewBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerPubkeyRegisteredIterator is returned from FilterPubkeyRegistered and is used to iterate over the raw logs and unpacked data for PubkeyRegistered events raised by the Logger contract.
type LoggerPubkeyRegisteredIterator struct {
	Event *LoggerPubkeyRegistered // Event containing the contract specifics and raw log

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
func (it *LoggerPubkeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerPubkeyRegistered)
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
		it.Event = new(LoggerPubkeyRegistered)
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
func (it *LoggerPubkeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerPubkeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerPubkeyRegistered represents a PubkeyRegistered event raised by the Logger contract.
type LoggerPubkeyRegistered struct {
	Pubkey   [4]*big.Int
	PubkeyID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPubkeyRegistered is a free log retrieval operation binding the contract event 0xf0777e5cea47492e18df87dcc844efabdfad315d1a2b4883d87cc2b964eddff0.
//
// Solidity: event PubkeyRegistered(uint256[4] pubkey, uint256 pubkeyID)
func (_Logger *LoggerFilterer) FilterPubkeyRegistered(opts *bind.FilterOpts) (*LoggerPubkeyRegisteredIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "PubkeyRegistered")
	if err != nil {
		return nil, err
	}
	return &LoggerPubkeyRegisteredIterator{contract: _Logger.contract, event: "PubkeyRegistered", logs: logs, sub: sub}, nil
}

// WatchPubkeyRegistered is a free log subscription operation binding the contract event 0xf0777e5cea47492e18df87dcc844efabdfad315d1a2b4883d87cc2b964eddff0.
//
// Solidity: event PubkeyRegistered(uint256[4] pubkey, uint256 pubkeyID)
func (_Logger *LoggerFilterer) WatchPubkeyRegistered(opts *bind.WatchOpts, sink chan<- *LoggerPubkeyRegistered) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "PubkeyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerPubkeyRegistered)
				if err := _Logger.contract.UnpackLog(event, "PubkeyRegistered", log); err != nil {
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
func (_Logger *LoggerFilterer) ParsePubkeyRegistered(log types.Log) (*LoggerPubkeyRegistered, error) {
	event := new(LoggerPubkeyRegistered)
	if err := _Logger.contract.UnpackLog(event, "PubkeyRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerRegisteredTokenIterator is returned from FilterRegisteredToken and is used to iterate over the raw logs and unpacked data for RegisteredToken events raised by the Logger contract.
type LoggerRegisteredTokenIterator struct {
	Event *LoggerRegisteredToken // Event containing the contract specifics and raw log

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
func (it *LoggerRegisteredTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerRegisteredToken)
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
		it.Event = new(LoggerRegisteredToken)
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
func (it *LoggerRegisteredTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerRegisteredTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerRegisteredToken represents a RegisteredToken event raised by the Logger contract.
type LoggerRegisteredToken struct {
	TokenType     *big.Int
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegisteredToken is a free log retrieval operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenType, address tokenContract)
func (_Logger *LoggerFilterer) FilterRegisteredToken(opts *bind.FilterOpts) (*LoggerRegisteredTokenIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "RegisteredToken")
	if err != nil {
		return nil, err
	}
	return &LoggerRegisteredTokenIterator{contract: _Logger.contract, event: "RegisteredToken", logs: logs, sub: sub}, nil
}

// WatchRegisteredToken is a free log subscription operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenType, address tokenContract)
func (_Logger *LoggerFilterer) WatchRegisteredToken(opts *bind.WatchOpts, sink chan<- *LoggerRegisteredToken) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "RegisteredToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerRegisteredToken)
				if err := _Logger.contract.UnpackLog(event, "RegisteredToken", log); err != nil {
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

// ParseRegisteredToken is a log parse operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenType, address tokenContract)
func (_Logger *LoggerFilterer) ParseRegisteredToken(log types.Log) (*LoggerRegisteredToken, error) {
	event := new(LoggerRegisteredToken)
	if err := _Logger.contract.UnpackLog(event, "RegisteredToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerRegistrationRequestIterator is returned from FilterRegistrationRequest and is used to iterate over the raw logs and unpacked data for RegistrationRequest events raised by the Logger contract.
type LoggerRegistrationRequestIterator struct {
	Event *LoggerRegistrationRequest // Event containing the contract specifics and raw log

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
func (it *LoggerRegistrationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerRegistrationRequest)
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
		it.Event = new(LoggerRegistrationRequest)
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
func (it *LoggerRegistrationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerRegistrationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerRegistrationRequest represents a RegistrationRequest event raised by the Logger contract.
type LoggerRegistrationRequest struct {
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegistrationRequest is a free log retrieval operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Logger *LoggerFilterer) FilterRegistrationRequest(opts *bind.FilterOpts) (*LoggerRegistrationRequestIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "RegistrationRequest")
	if err != nil {
		return nil, err
	}
	return &LoggerRegistrationRequestIterator{contract: _Logger.contract, event: "RegistrationRequest", logs: logs, sub: sub}, nil
}

// WatchRegistrationRequest is a free log subscription operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Logger *LoggerFilterer) WatchRegistrationRequest(opts *bind.WatchOpts, sink chan<- *LoggerRegistrationRequest) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "RegistrationRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerRegistrationRequest)
				if err := _Logger.contract.UnpackLog(event, "RegistrationRequest", log); err != nil {
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

// ParseRegistrationRequest is a log parse operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Logger *LoggerFilterer) ParseRegistrationRequest(log types.Log) (*LoggerRegistrationRequest, error) {
	event := new(LoggerRegistrationRequest)
	if err := _Logger.contract.UnpackLog(event, "RegistrationRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerRollbackFinalisationIterator is returned from FilterRollbackFinalisation and is used to iterate over the raw logs and unpacked data for RollbackFinalisation events raised by the Logger contract.
type LoggerRollbackFinalisationIterator struct {
	Event *LoggerRollbackFinalisation // Event containing the contract specifics and raw log

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
func (it *LoggerRollbackFinalisationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerRollbackFinalisation)
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
		it.Event = new(LoggerRollbackFinalisation)
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
func (it *LoggerRollbackFinalisationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerRollbackFinalisationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerRollbackFinalisation represents a RollbackFinalisation event raised by the Logger contract.
type LoggerRollbackFinalisation struct {
	TotalBatchesSlashed *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRollbackFinalisation is a free log retrieval operation binding the contract event 0x8efd02dfe309f172ea08425236aa43cec05abc25ebbb9cd4ed1de0d5048fc91a.
//
// Solidity: event RollbackFinalisation(uint256 totalBatchesSlashed)
func (_Logger *LoggerFilterer) FilterRollbackFinalisation(opts *bind.FilterOpts) (*LoggerRollbackFinalisationIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "RollbackFinalisation")
	if err != nil {
		return nil, err
	}
	return &LoggerRollbackFinalisationIterator{contract: _Logger.contract, event: "RollbackFinalisation", logs: logs, sub: sub}, nil
}

// WatchRollbackFinalisation is a free log subscription operation binding the contract event 0x8efd02dfe309f172ea08425236aa43cec05abc25ebbb9cd4ed1de0d5048fc91a.
//
// Solidity: event RollbackFinalisation(uint256 totalBatchesSlashed)
func (_Logger *LoggerFilterer) WatchRollbackFinalisation(opts *bind.WatchOpts, sink chan<- *LoggerRollbackFinalisation) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "RollbackFinalisation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerRollbackFinalisation)
				if err := _Logger.contract.UnpackLog(event, "RollbackFinalisation", log); err != nil {
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

// ParseRollbackFinalisation is a log parse operation binding the contract event 0x8efd02dfe309f172ea08425236aa43cec05abc25ebbb9cd4ed1de0d5048fc91a.
//
// Solidity: event RollbackFinalisation(uint256 totalBatchesSlashed)
func (_Logger *LoggerFilterer) ParseRollbackFinalisation(log types.Log) (*LoggerRollbackFinalisation, error) {
	event := new(LoggerRollbackFinalisation)
	if err := _Logger.contract.UnpackLog(event, "RollbackFinalisation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoggerStakeWithdrawIterator is returned from FilterStakeWithdraw and is used to iterate over the raw logs and unpacked data for StakeWithdraw events raised by the Logger contract.
type LoggerStakeWithdrawIterator struct {
	Event *LoggerStakeWithdraw // Event containing the contract specifics and raw log

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
func (it *LoggerStakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerStakeWithdraw)
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
		it.Event = new(LoggerStakeWithdraw)
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
func (it *LoggerStakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggerStakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggerStakeWithdraw represents a StakeWithdraw event raised by the Logger contract.
type LoggerStakeWithdraw struct {
	Committed common.Address
	BatchID   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdraw is a free log retrieval operation binding the contract event 0x1248d48e2de900a1010c7fce73506969ecec243600bfc08b641b158f26d857cd.
//
// Solidity: event StakeWithdraw(address committed, uint256 batchID)
func (_Logger *LoggerFilterer) FilterStakeWithdraw(opts *bind.FilterOpts) (*LoggerStakeWithdrawIterator, error) {

	logs, sub, err := _Logger.contract.FilterLogs(opts, "StakeWithdraw")
	if err != nil {
		return nil, err
	}
	return &LoggerStakeWithdrawIterator{contract: _Logger.contract, event: "StakeWithdraw", logs: logs, sub: sub}, nil
}

// WatchStakeWithdraw is a free log subscription operation binding the contract event 0x1248d48e2de900a1010c7fce73506969ecec243600bfc08b641b158f26d857cd.
//
// Solidity: event StakeWithdraw(address committed, uint256 batchID)
func (_Logger *LoggerFilterer) WatchStakeWithdraw(opts *bind.WatchOpts, sink chan<- *LoggerStakeWithdraw) (event.Subscription, error) {

	logs, sub, err := _Logger.contract.WatchLogs(opts, "StakeWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggerStakeWithdraw)
				if err := _Logger.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
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

// ParseStakeWithdraw is a log parse operation binding the contract event 0x1248d48e2de900a1010c7fce73506969ecec243600bfc08b641b158f26d857cd.
//
// Solidity: event StakeWithdraw(address committed, uint256 batchID)
func (_Logger *LoggerFilterer) ParseStakeWithdraw(log types.Log) (*LoggerStakeWithdraw, error) {
	event := new(LoggerStakeWithdraw)
	if err := _Logger.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
