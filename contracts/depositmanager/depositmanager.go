// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositmanager

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

// DepositmanagerABI is the input ABI used to generate the binding from.
const DepositmanagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"DepositSubTreeReady\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"dequeueToSubmit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"subtreeRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subtreeRoot\",\"type\":\"bytes32\"}],\"name\":\"reenqueue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Depositmanager is an auto generated Go binding around an Ethereum contract.
type Depositmanager struct {
	DepositmanagerCaller     // Read-only binding to the contract
	DepositmanagerTransactor // Write-only binding to the contract
	DepositmanagerFilterer   // Log filterer for contract events
}

// DepositmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositmanagerSession struct {
	Contract     *Depositmanager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositmanagerCallerSession struct {
	Contract *DepositmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DepositmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositmanagerTransactorSession struct {
	Contract     *DepositmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DepositmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositmanagerRaw struct {
	Contract *Depositmanager // Generic contract binding to access the raw methods on
}

// DepositmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositmanagerCallerRaw struct {
	Contract *DepositmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// DepositmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositmanagerTransactorRaw struct {
	Contract *DepositmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositmanager creates a new instance of Depositmanager, bound to a specific deployed contract.
func NewDepositmanager(address common.Address, backend bind.ContractBackend) (*Depositmanager, error) {
	contract, err := bindDepositmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositmanager{DepositmanagerCaller: DepositmanagerCaller{contract: contract}, DepositmanagerTransactor: DepositmanagerTransactor{contract: contract}, DepositmanagerFilterer: DepositmanagerFilterer{contract: contract}}, nil
}

// NewDepositmanagerCaller creates a new read-only instance of Depositmanager, bound to a specific deployed contract.
func NewDepositmanagerCaller(address common.Address, caller bind.ContractCaller) (*DepositmanagerCaller, error) {
	contract, err := bindDepositmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerCaller{contract: contract}, nil
}

// NewDepositmanagerTransactor creates a new write-only instance of Depositmanager, bound to a specific deployed contract.
func NewDepositmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositmanagerTransactor, error) {
	contract, err := bindDepositmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerTransactor{contract: contract}, nil
}

// NewDepositmanagerFilterer creates a new log filterer instance of Depositmanager, bound to a specific deployed contract.
func NewDepositmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositmanagerFilterer, error) {
	contract, err := bindDepositmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerFilterer{contract: contract}, nil
}

// bindDepositmanager binds a generic wrapper to an already deployed contract.
func bindDepositmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositmanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositmanager *DepositmanagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Depositmanager.Contract.DepositmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositmanager *DepositmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositmanager *DepositmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositmanager *DepositmanagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Depositmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositmanager *DepositmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositmanager *DepositmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositmanager.Contract.contract.Transact(opts, method, params...)
}

// DequeueToSubmit is a paid mutator transaction binding the contract method 0xd86ee48d.
//
// Solidity: function dequeueToSubmit() returns(bytes32 subtreeRoot)
func (_Depositmanager *DepositmanagerTransactor) DequeueToSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "dequeueToSubmit")
}

// DequeueToSubmit is a paid mutator transaction binding the contract method 0xd86ee48d.
//
// Solidity: function dequeueToSubmit() returns(bytes32 subtreeRoot)
func (_Depositmanager *DepositmanagerSession) DequeueToSubmit() (*types.Transaction, error) {
	return _Depositmanager.Contract.DequeueToSubmit(&_Depositmanager.TransactOpts)
}

// DequeueToSubmit is a paid mutator transaction binding the contract method 0xd86ee48d.
//
// Solidity: function dequeueToSubmit() returns(bytes32 subtreeRoot)
func (_Depositmanager *DepositmanagerTransactorSession) DequeueToSubmit() (*types.Transaction, error) {
	return _Depositmanager.Contract.DequeueToSubmit(&_Depositmanager.TransactOpts)
}

// Reenqueue is a paid mutator transaction binding the contract method 0xade00026.
//
// Solidity: function reenqueue(bytes32 subtreeRoot) returns()
func (_Depositmanager *DepositmanagerTransactor) Reenqueue(opts *bind.TransactOpts, subtreeRoot [32]byte) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "reenqueue", subtreeRoot)
}

// Reenqueue is a paid mutator transaction binding the contract method 0xade00026.
//
// Solidity: function reenqueue(bytes32 subtreeRoot) returns()
func (_Depositmanager *DepositmanagerSession) Reenqueue(subtreeRoot [32]byte) (*types.Transaction, error) {
	return _Depositmanager.Contract.Reenqueue(&_Depositmanager.TransactOpts, subtreeRoot)
}

// Reenqueue is a paid mutator transaction binding the contract method 0xade00026.
//
// Solidity: function reenqueue(bytes32 subtreeRoot) returns()
func (_Depositmanager *DepositmanagerTransactorSession) Reenqueue(subtreeRoot [32]byte) (*types.Transaction, error) {
	return _Depositmanager.Contract.Reenqueue(&_Depositmanager.TransactOpts, subtreeRoot)
}

// DepositmanagerDepositQueuedIterator is returned from FilterDepositQueued and is used to iterate over the raw logs and unpacked data for DepositQueued events raised by the Depositmanager contract.
type DepositmanagerDepositQueuedIterator struct {
	Event *DepositmanagerDepositQueued // Event containing the contract specifics and raw log

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
func (it *DepositmanagerDepositQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositmanagerDepositQueued)
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
		it.Event = new(DepositmanagerDepositQueued)
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
func (it *DepositmanagerDepositQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositmanagerDepositQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositmanagerDepositQueued represents a DepositQueued event raised by the Depositmanager contract.
type DepositmanagerDepositQueued struct {
	PubkeyID *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositQueued is a free log retrieval operation binding the contract event 0x5a1922090c12e28c132a961f6bb4d74350598a62e8981b5eb9bb5ccbbce9df94.
//
// Solidity: event DepositQueued(uint256 pubkeyID, bytes data)
func (_Depositmanager *DepositmanagerFilterer) FilterDepositQueued(opts *bind.FilterOpts) (*DepositmanagerDepositQueuedIterator, error) {

	logs, sub, err := _Depositmanager.contract.FilterLogs(opts, "DepositQueued")
	if err != nil {
		return nil, err
	}
	return &DepositmanagerDepositQueuedIterator{contract: _Depositmanager.contract, event: "DepositQueued", logs: logs, sub: sub}, nil
}

// WatchDepositQueued is a free log subscription operation binding the contract event 0x5a1922090c12e28c132a961f6bb4d74350598a62e8981b5eb9bb5ccbbce9df94.
//
// Solidity: event DepositQueued(uint256 pubkeyID, bytes data)
func (_Depositmanager *DepositmanagerFilterer) WatchDepositQueued(opts *bind.WatchOpts, sink chan<- *DepositmanagerDepositQueued) (event.Subscription, error) {

	logs, sub, err := _Depositmanager.contract.WatchLogs(opts, "DepositQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositmanagerDepositQueued)
				if err := _Depositmanager.contract.UnpackLog(event, "DepositQueued", log); err != nil {
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
func (_Depositmanager *DepositmanagerFilterer) ParseDepositQueued(log types.Log) (*DepositmanagerDepositQueued, error) {
	event := new(DepositmanagerDepositQueued)
	if err := _Depositmanager.contract.UnpackLog(event, "DepositQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositmanagerDepositSubTreeReadyIterator is returned from FilterDepositSubTreeReady and is used to iterate over the raw logs and unpacked data for DepositSubTreeReady events raised by the Depositmanager contract.
type DepositmanagerDepositSubTreeReadyIterator struct {
	Event *DepositmanagerDepositSubTreeReady // Event containing the contract specifics and raw log

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
func (it *DepositmanagerDepositSubTreeReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositmanagerDepositSubTreeReady)
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
		it.Event = new(DepositmanagerDepositSubTreeReady)
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
func (it *DepositmanagerDepositSubTreeReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositmanagerDepositSubTreeReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositmanagerDepositSubTreeReady represents a DepositSubTreeReady event raised by the Depositmanager contract.
type DepositmanagerDepositSubTreeReady struct {
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDepositSubTreeReady is a free log retrieval operation binding the contract event 0x4744f3a44c5716c9fa423a71cdaa806771a8bf469f4c007ca338b8e8e202a8b5.
//
// Solidity: event DepositSubTreeReady(bytes32 root)
func (_Depositmanager *DepositmanagerFilterer) FilterDepositSubTreeReady(opts *bind.FilterOpts) (*DepositmanagerDepositSubTreeReadyIterator, error) {

	logs, sub, err := _Depositmanager.contract.FilterLogs(opts, "DepositSubTreeReady")
	if err != nil {
		return nil, err
	}
	return &DepositmanagerDepositSubTreeReadyIterator{contract: _Depositmanager.contract, event: "DepositSubTreeReady", logs: logs, sub: sub}, nil
}

// WatchDepositSubTreeReady is a free log subscription operation binding the contract event 0x4744f3a44c5716c9fa423a71cdaa806771a8bf469f4c007ca338b8e8e202a8b5.
//
// Solidity: event DepositSubTreeReady(bytes32 root)
func (_Depositmanager *DepositmanagerFilterer) WatchDepositSubTreeReady(opts *bind.WatchOpts, sink chan<- *DepositmanagerDepositSubTreeReady) (event.Subscription, error) {

	logs, sub, err := _Depositmanager.contract.WatchLogs(opts, "DepositSubTreeReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositmanagerDepositSubTreeReady)
				if err := _Depositmanager.contract.UnpackLog(event, "DepositSubTreeReady", log); err != nil {
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
func (_Depositmanager *DepositmanagerFilterer) ParseDepositSubTreeReady(log types.Log) (*DepositmanagerDepositSubTreeReady, error) {
	event := new(DepositmanagerDepositSubTreeReady)
	if err := _Depositmanager.contract.UnpackLog(event, "DepositSubTreeReady", log); err != nil {
		return nil, err
	}
	return event, nil
}
