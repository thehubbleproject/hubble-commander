// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenregistry

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

// TokenregistryABI is the input ABI used to generate the binding from.
const TokenregistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"RegisteredToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"RegistrationRequest\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"finaliseRegistration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextTokenID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingRegistrations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"requestRegistration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"safeGetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Tokenregistry is an auto generated Go binding around an Ethereum contract.
type Tokenregistry struct {
	TokenregistryCaller     // Read-only binding to the contract
	TokenregistryTransactor // Write-only binding to the contract
	TokenregistryFilterer   // Log filterer for contract events
}

// TokenregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenregistrySession struct {
	Contract     *Tokenregistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenregistryCallerSession struct {
	Contract *TokenregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenregistryTransactorSession struct {
	Contract     *TokenregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenregistryRaw struct {
	Contract *Tokenregistry // Generic contract binding to access the raw methods on
}

// TokenregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenregistryCallerRaw struct {
	Contract *TokenregistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenregistryTransactorRaw struct {
	Contract *TokenregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenregistry creates a new instance of Tokenregistry, bound to a specific deployed contract.
func NewTokenregistry(address common.Address, backend bind.ContractBackend) (*Tokenregistry, error) {
	contract, err := bindTokenregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenregistry{TokenregistryCaller: TokenregistryCaller{contract: contract}, TokenregistryTransactor: TokenregistryTransactor{contract: contract}, TokenregistryFilterer: TokenregistryFilterer{contract: contract}}, nil
}

// NewTokenregistryCaller creates a new read-only instance of Tokenregistry, bound to a specific deployed contract.
func NewTokenregistryCaller(address common.Address, caller bind.ContractCaller) (*TokenregistryCaller, error) {
	contract, err := bindTokenregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenregistryCaller{contract: contract}, nil
}

// NewTokenregistryTransactor creates a new write-only instance of Tokenregistry, bound to a specific deployed contract.
func NewTokenregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenregistryTransactor, error) {
	contract, err := bindTokenregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenregistryTransactor{contract: contract}, nil
}

// NewTokenregistryFilterer creates a new log filterer instance of Tokenregistry, bound to a specific deployed contract.
func NewTokenregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenregistryFilterer, error) {
	contract, err := bindTokenregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenregistryFilterer{contract: contract}, nil
}

// bindTokenregistry binds a generic wrapper to an already deployed contract.
func bindTokenregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenregistry *TokenregistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenregistry.Contract.TokenregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenregistry *TokenregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenregistry.Contract.TokenregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenregistry *TokenregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenregistry.Contract.TokenregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenregistry *TokenregistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenregistry *TokenregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenregistry *TokenregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenregistry.Contract.contract.Transact(opts, method, params...)
}

// NextTokenID is a free data retrieval call binding the contract method 0xf101e481.
//
// Solidity: function nextTokenID() view returns(uint256)
func (_Tokenregistry *TokenregistryCaller) NextTokenID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenregistry.contract.Call(opts, out, "nextTokenID")
	return *ret0, err
}

// NextTokenID is a free data retrieval call binding the contract method 0xf101e481.
//
// Solidity: function nextTokenID() view returns(uint256)
func (_Tokenregistry *TokenregistrySession) NextTokenID() (*big.Int, error) {
	return _Tokenregistry.Contract.NextTokenID(&_Tokenregistry.CallOpts)
}

// NextTokenID is a free data retrieval call binding the contract method 0xf101e481.
//
// Solidity: function nextTokenID() view returns(uint256)
func (_Tokenregistry *TokenregistryCallerSession) NextTokenID() (*big.Int, error) {
	return _Tokenregistry.Contract.NextTokenID(&_Tokenregistry.CallOpts)
}

// PendingRegistrations is a free data retrieval call binding the contract method 0xcd4852e5.
//
// Solidity: function pendingRegistrations(address ) view returns(bool)
func (_Tokenregistry *TokenregistryCaller) PendingRegistrations(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tokenregistry.contract.Call(opts, out, "pendingRegistrations", arg0)
	return *ret0, err
}

// PendingRegistrations is a free data retrieval call binding the contract method 0xcd4852e5.
//
// Solidity: function pendingRegistrations(address ) view returns(bool)
func (_Tokenregistry *TokenregistrySession) PendingRegistrations(arg0 common.Address) (bool, error) {
	return _Tokenregistry.Contract.PendingRegistrations(&_Tokenregistry.CallOpts, arg0)
}

// PendingRegistrations is a free data retrieval call binding the contract method 0xcd4852e5.
//
// Solidity: function pendingRegistrations(address ) view returns(bool)
func (_Tokenregistry *TokenregistryCallerSession) PendingRegistrations(arg0 common.Address) (bool, error) {
	return _Tokenregistry.Contract.PendingRegistrations(&_Tokenregistry.CallOpts, arg0)
}

// SafeGetAddress is a free data retrieval call binding the contract method 0x41eb569e.
//
// Solidity: function safeGetAddress(uint256 tokenID) view returns(address)
func (_Tokenregistry *TokenregistryCaller) SafeGetAddress(opts *bind.CallOpts, tokenID *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tokenregistry.contract.Call(opts, out, "safeGetAddress", tokenID)
	return *ret0, err
}

// SafeGetAddress is a free data retrieval call binding the contract method 0x41eb569e.
//
// Solidity: function safeGetAddress(uint256 tokenID) view returns(address)
func (_Tokenregistry *TokenregistrySession) SafeGetAddress(tokenID *big.Int) (common.Address, error) {
	return _Tokenregistry.Contract.SafeGetAddress(&_Tokenregistry.CallOpts, tokenID)
}

// SafeGetAddress is a free data retrieval call binding the contract method 0x41eb569e.
//
// Solidity: function safeGetAddress(uint256 tokenID) view returns(address)
func (_Tokenregistry *TokenregistryCallerSession) SafeGetAddress(tokenID *big.Int) (common.Address, error) {
	return _Tokenregistry.Contract.SafeGetAddress(&_Tokenregistry.CallOpts, tokenID)
}

// FinaliseRegistration is a paid mutator transaction binding the contract method 0x0b72ccbc.
//
// Solidity: function finaliseRegistration(address tokenContract) returns()
func (_Tokenregistry *TokenregistryTransactor) FinaliseRegistration(opts *bind.TransactOpts, tokenContract common.Address) (*types.Transaction, error) {
	return _Tokenregistry.contract.Transact(opts, "finaliseRegistration", tokenContract)
}

// FinaliseRegistration is a paid mutator transaction binding the contract method 0x0b72ccbc.
//
// Solidity: function finaliseRegistration(address tokenContract) returns()
func (_Tokenregistry *TokenregistrySession) FinaliseRegistration(tokenContract common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.FinaliseRegistration(&_Tokenregistry.TransactOpts, tokenContract)
}

// FinaliseRegistration is a paid mutator transaction binding the contract method 0x0b72ccbc.
//
// Solidity: function finaliseRegistration(address tokenContract) returns()
func (_Tokenregistry *TokenregistryTransactorSession) FinaliseRegistration(tokenContract common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.FinaliseRegistration(&_Tokenregistry.TransactOpts, tokenContract)
}

// RequestRegistration is a paid mutator transaction binding the contract method 0xe9b6bfe3.
//
// Solidity: function requestRegistration(address tokenContract) returns()
func (_Tokenregistry *TokenregistryTransactor) RequestRegistration(opts *bind.TransactOpts, tokenContract common.Address) (*types.Transaction, error) {
	return _Tokenregistry.contract.Transact(opts, "requestRegistration", tokenContract)
}

// RequestRegistration is a paid mutator transaction binding the contract method 0xe9b6bfe3.
//
// Solidity: function requestRegistration(address tokenContract) returns()
func (_Tokenregistry *TokenregistrySession) RequestRegistration(tokenContract common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.RequestRegistration(&_Tokenregistry.TransactOpts, tokenContract)
}

// RequestRegistration is a paid mutator transaction binding the contract method 0xe9b6bfe3.
//
// Solidity: function requestRegistration(address tokenContract) returns()
func (_Tokenregistry *TokenregistryTransactorSession) RequestRegistration(tokenContract common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.RequestRegistration(&_Tokenregistry.TransactOpts, tokenContract)
}

// TokenregistryRegisteredTokenIterator is returned from FilterRegisteredToken and is used to iterate over the raw logs and unpacked data for RegisteredToken events raised by the Tokenregistry contract.
type TokenregistryRegisteredTokenIterator struct {
	Event *TokenregistryRegisteredToken // Event containing the contract specifics and raw log

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
func (it *TokenregistryRegisteredTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenregistryRegisteredToken)
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
		it.Event = new(TokenregistryRegisteredToken)
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
func (it *TokenregistryRegisteredTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenregistryRegisteredTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenregistryRegisteredToken represents a RegisteredToken event raised by the Tokenregistry contract.
type TokenregistryRegisteredToken struct {
	TokenID       *big.Int
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegisteredToken is a free log retrieval operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenID, address tokenContract)
func (_Tokenregistry *TokenregistryFilterer) FilterRegisteredToken(opts *bind.FilterOpts) (*TokenregistryRegisteredTokenIterator, error) {

	logs, sub, err := _Tokenregistry.contract.FilterLogs(opts, "RegisteredToken")
	if err != nil {
		return nil, err
	}
	return &TokenregistryRegisteredTokenIterator{contract: _Tokenregistry.contract, event: "RegisteredToken", logs: logs, sub: sub}, nil
}

// WatchRegisteredToken is a free log subscription operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenID, address tokenContract)
func (_Tokenregistry *TokenregistryFilterer) WatchRegisteredToken(opts *bind.WatchOpts, sink chan<- *TokenregistryRegisteredToken) (event.Subscription, error) {

	logs, sub, err := _Tokenregistry.contract.WatchLogs(opts, "RegisteredToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenregistryRegisteredToken)
				if err := _Tokenregistry.contract.UnpackLog(event, "RegisteredToken", log); err != nil {
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
// Solidity: event RegisteredToken(uint256 tokenID, address tokenContract)
func (_Tokenregistry *TokenregistryFilterer) ParseRegisteredToken(log types.Log) (*TokenregistryRegisteredToken, error) {
	event := new(TokenregistryRegisteredToken)
	if err := _Tokenregistry.contract.UnpackLog(event, "RegisteredToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenregistryRegistrationRequestIterator is returned from FilterRegistrationRequest and is used to iterate over the raw logs and unpacked data for RegistrationRequest events raised by the Tokenregistry contract.
type TokenregistryRegistrationRequestIterator struct {
	Event *TokenregistryRegistrationRequest // Event containing the contract specifics and raw log

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
func (it *TokenregistryRegistrationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenregistryRegistrationRequest)
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
		it.Event = new(TokenregistryRegistrationRequest)
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
func (it *TokenregistryRegistrationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenregistryRegistrationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenregistryRegistrationRequest represents a RegistrationRequest event raised by the Tokenregistry contract.
type TokenregistryRegistrationRequest struct {
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegistrationRequest is a free log retrieval operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Tokenregistry *TokenregistryFilterer) FilterRegistrationRequest(opts *bind.FilterOpts) (*TokenregistryRegistrationRequestIterator, error) {

	logs, sub, err := _Tokenregistry.contract.FilterLogs(opts, "RegistrationRequest")
	if err != nil {
		return nil, err
	}
	return &TokenregistryRegistrationRequestIterator{contract: _Tokenregistry.contract, event: "RegistrationRequest", logs: logs, sub: sub}, nil
}

// WatchRegistrationRequest is a free log subscription operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Tokenregistry *TokenregistryFilterer) WatchRegistrationRequest(opts *bind.WatchOpts, sink chan<- *TokenregistryRegistrationRequest) (event.Subscription, error) {

	logs, sub, err := _Tokenregistry.contract.WatchLogs(opts, "RegistrationRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenregistryRegistrationRequest)
				if err := _Tokenregistry.contract.UnpackLog(event, "RegistrationRequest", log); err != nil {
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
func (_Tokenregistry *TokenregistryFilterer) ParseRegistrationRequest(log types.Log) (*TokenregistryRegistrationRequest, error) {
	event := new(TokenregistryRegistrationRequest)
	if err := _Tokenregistry.contract.UnpackLog(event, "RegistrationRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}
