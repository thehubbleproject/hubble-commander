// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burnauction

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

// BurnauctionABI is the input ABI used to generate the binding from.
const BurnauctionABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"slot\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"NewBestBid\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"coordinator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numBlock\",\"type\":\"uint256\"}],\"name\":\"block2slot\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSlot\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesisBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Burnauction is an auto generated Go binding around an Ethereum contract.
type Burnauction struct {
	BurnauctionCaller     // Read-only binding to the contract
	BurnauctionTransactor // Write-only binding to the contract
	BurnauctionFilterer   // Log filterer for contract events
}

// BurnauctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnauctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnauctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnauctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnauctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnauctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnauctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnauctionSession struct {
	Contract     *Burnauction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnauctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnauctionCallerSession struct {
	Contract *BurnauctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BurnauctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnauctionTransactorSession struct {
	Contract     *BurnauctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BurnauctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnauctionRaw struct {
	Contract *Burnauction // Generic contract binding to access the raw methods on
}

// BurnauctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnauctionCallerRaw struct {
	Contract *BurnauctionCaller // Generic read-only contract binding to access the raw methods on
}

// BurnauctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnauctionTransactorRaw struct {
	Contract *BurnauctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnauction creates a new instance of Burnauction, bound to a specific deployed contract.
func NewBurnauction(address common.Address, backend bind.ContractBackend) (*Burnauction, error) {
	contract, err := bindBurnauction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Burnauction{BurnauctionCaller: BurnauctionCaller{contract: contract}, BurnauctionTransactor: BurnauctionTransactor{contract: contract}, BurnauctionFilterer: BurnauctionFilterer{contract: contract}}, nil
}

// NewBurnauctionCaller creates a new read-only instance of Burnauction, bound to a specific deployed contract.
func NewBurnauctionCaller(address common.Address, caller bind.ContractCaller) (*BurnauctionCaller, error) {
	contract, err := bindBurnauction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnauctionCaller{contract: contract}, nil
}

// NewBurnauctionTransactor creates a new write-only instance of Burnauction, bound to a specific deployed contract.
func NewBurnauctionTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnauctionTransactor, error) {
	contract, err := bindBurnauction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnauctionTransactor{contract: contract}, nil
}

// NewBurnauctionFilterer creates a new log filterer instance of Burnauction, bound to a specific deployed contract.
func NewBurnauctionFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnauctionFilterer, error) {
	contract, err := bindBurnauction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnauctionFilterer{contract: contract}, nil
}

// bindBurnauction binds a generic wrapper to an already deployed contract.
func bindBurnauction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnauctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Burnauction *BurnauctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Burnauction.Contract.BurnauctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Burnauction *BurnauctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burnauction.Contract.BurnauctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Burnauction *BurnauctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Burnauction.Contract.BurnauctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Burnauction *BurnauctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Burnauction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Burnauction *BurnauctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burnauction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Burnauction *BurnauctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Burnauction.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0xd3dd08e2.
//
// Solidity: function auction(uint32 ) view returns(address coordinator, uint128 amount, bool initialized)
func (_Burnauction *BurnauctionCaller) Auction(opts *bind.CallOpts, arg0 uint32) (struct {
	Coordinator common.Address
	Amount      *big.Int
	Initialized bool
}, error) {
	ret := new(struct {
		Coordinator common.Address
		Amount      *big.Int
		Initialized bool
	})
	out := ret
	err := _Burnauction.contract.Call(opts, out, "auction", arg0)
	return *ret, err
}

// Auction is a free data retrieval call binding the contract method 0xd3dd08e2.
//
// Solidity: function auction(uint32 ) view returns(address coordinator, uint128 amount, bool initialized)
func (_Burnauction *BurnauctionSession) Auction(arg0 uint32) (struct {
	Coordinator common.Address
	Amount      *big.Int
	Initialized bool
}, error) {
	return _Burnauction.Contract.Auction(&_Burnauction.CallOpts, arg0)
}

// Auction is a free data retrieval call binding the contract method 0xd3dd08e2.
//
// Solidity: function auction(uint32 ) view returns(address coordinator, uint128 amount, bool initialized)
func (_Burnauction *BurnauctionCallerSession) Auction(arg0 uint32) (struct {
	Coordinator common.Address
	Amount      *big.Int
	Initialized bool
}, error) {
	return _Burnauction.Contract.Auction(&_Burnauction.CallOpts, arg0)
}

// Block2slot is a free data retrieval call binding the contract method 0xa87a2ead.
//
// Solidity: function block2slot(uint256 numBlock) view returns(uint32)
func (_Burnauction *BurnauctionCaller) Block2slot(opts *bind.CallOpts, numBlock *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Burnauction.contract.Call(opts, out, "block2slot", numBlock)
	return *ret0, err
}

// Block2slot is a free data retrieval call binding the contract method 0xa87a2ead.
//
// Solidity: function block2slot(uint256 numBlock) view returns(uint32)
func (_Burnauction *BurnauctionSession) Block2slot(numBlock *big.Int) (uint32, error) {
	return _Burnauction.Contract.Block2slot(&_Burnauction.CallOpts, numBlock)
}

// Block2slot is a free data retrieval call binding the contract method 0xa87a2ead.
//
// Solidity: function block2slot(uint256 numBlock) view returns(uint32)
func (_Burnauction *BurnauctionCallerSession) Block2slot(numBlock *big.Int) (uint32, error) {
	return _Burnauction.Contract.Block2slot(&_Burnauction.CallOpts, numBlock)
}

// CurrentSlot is a free data retrieval call binding the contract method 0x3359632e.
//
// Solidity: function currentSlot() view returns(uint32)
func (_Burnauction *BurnauctionCaller) CurrentSlot(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Burnauction.contract.Call(opts, out, "currentSlot")
	return *ret0, err
}

// CurrentSlot is a free data retrieval call binding the contract method 0x3359632e.
//
// Solidity: function currentSlot() view returns(uint32)
func (_Burnauction *BurnauctionSession) CurrentSlot() (uint32, error) {
	return _Burnauction.Contract.CurrentSlot(&_Burnauction.CallOpts)
}

// CurrentSlot is a free data retrieval call binding the contract method 0x3359632e.
//
// Solidity: function currentSlot() view returns(uint32)
func (_Burnauction *BurnauctionCallerSession) CurrentSlot() (uint32, error) {
	return _Burnauction.Contract.CurrentSlot(&_Burnauction.CallOpts)
}

// GenesisBlock is a free data retrieval call binding the contract method 0x4cdc9c63.
//
// Solidity: function genesisBlock() view returns(uint256)
func (_Burnauction *BurnauctionCaller) GenesisBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Burnauction.contract.Call(opts, out, "genesisBlock")
	return *ret0, err
}

// GenesisBlock is a free data retrieval call binding the contract method 0x4cdc9c63.
//
// Solidity: function genesisBlock() view returns(uint256)
func (_Burnauction *BurnauctionSession) GenesisBlock() (*big.Int, error) {
	return _Burnauction.Contract.GenesisBlock(&_Burnauction.CallOpts)
}

// GenesisBlock is a free data retrieval call binding the contract method 0x4cdc9c63.
//
// Solidity: function genesisBlock() view returns(uint256)
func (_Burnauction *BurnauctionCallerSession) GenesisBlock() (*big.Int, error) {
	return _Burnauction.Contract.GenesisBlock(&_Burnauction.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Burnauction *BurnauctionCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Burnauction.contract.Call(opts, out, "getBlockNumber")
	return *ret0, err
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Burnauction *BurnauctionSession) GetBlockNumber() (*big.Int, error) {
	return _Burnauction.Contract.GetBlockNumber(&_Burnauction.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Burnauction *BurnauctionCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Burnauction.Contract.GetBlockNumber(&_Burnauction.CallOpts)
}

// GetProposer is a free data retrieval call binding the contract method 0xe9790d02.
//
// Solidity: function getProposer() view returns(address)
func (_Burnauction *BurnauctionCaller) GetProposer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Burnauction.contract.Call(opts, out, "getProposer")
	return *ret0, err
}

// GetProposer is a free data retrieval call binding the contract method 0xe9790d02.
//
// Solidity: function getProposer() view returns(address)
func (_Burnauction *BurnauctionSession) GetProposer() (common.Address, error) {
	return _Burnauction.Contract.GetProposer(&_Burnauction.CallOpts)
}

// GetProposer is a free data retrieval call binding the contract method 0xe9790d02.
//
// Solidity: function getProposer() view returns(address)
func (_Burnauction *BurnauctionCallerSession) GetProposer() (common.Address, error) {
	return _Burnauction.Contract.GetProposer(&_Burnauction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Burnauction *BurnauctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burnauction.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Burnauction *BurnauctionSession) Bid() (*types.Transaction, error) {
	return _Burnauction.Contract.Bid(&_Burnauction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Burnauction *BurnauctionTransactorSession) Bid() (*types.Transaction, error) {
	return _Burnauction.Contract.Bid(&_Burnauction.TransactOpts)
}

// BurnauctionNewBestBidIterator is returned from FilterNewBestBid and is used to iterate over the raw logs and unpacked data for NewBestBid events raised by the Burnauction contract.
type BurnauctionNewBestBidIterator struct {
	Event *BurnauctionNewBestBid // Event containing the contract specifics and raw log

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
func (it *BurnauctionNewBestBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnauctionNewBestBid)
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
		it.Event = new(BurnauctionNewBestBid)
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
func (it *BurnauctionNewBestBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnauctionNewBestBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnauctionNewBestBid represents a NewBestBid event raised by the Burnauction contract.
type BurnauctionNewBestBid struct {
	Slot        uint32
	Coordinator common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBestBid is a free log retrieval operation binding the contract event 0xc2ba2c9fb0ca6acb820a11398dccf241b9126b36ddf994d1baddd51dbf22bc71.
//
// Solidity: event NewBestBid(uint32 slot, address coordinator, uint128 amount)
func (_Burnauction *BurnauctionFilterer) FilterNewBestBid(opts *bind.FilterOpts) (*BurnauctionNewBestBidIterator, error) {

	logs, sub, err := _Burnauction.contract.FilterLogs(opts, "NewBestBid")
	if err != nil {
		return nil, err
	}
	return &BurnauctionNewBestBidIterator{contract: _Burnauction.contract, event: "NewBestBid", logs: logs, sub: sub}, nil
}

// WatchNewBestBid is a free log subscription operation binding the contract event 0xc2ba2c9fb0ca6acb820a11398dccf241b9126b36ddf994d1baddd51dbf22bc71.
//
// Solidity: event NewBestBid(uint32 slot, address coordinator, uint128 amount)
func (_Burnauction *BurnauctionFilterer) WatchNewBestBid(opts *bind.WatchOpts, sink chan<- *BurnauctionNewBestBid) (event.Subscription, error) {

	logs, sub, err := _Burnauction.contract.WatchLogs(opts, "NewBestBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnauctionNewBestBid)
				if err := _Burnauction.contract.UnpackLog(event, "NewBestBid", log); err != nil {
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

// ParseNewBestBid is a log parse operation binding the contract event 0xc2ba2c9fb0ca6acb820a11398dccf241b9126b36ddf994d1baddd51dbf22bc71.
//
// Solidity: event NewBestBid(uint32 slot, address coordinator, uint128 amount)
func (_Burnauction *BurnauctionFilterer) ParseNewBestBid(log types.Log) (*BurnauctionNewBestBid, error) {
	event := new(BurnauctionNewBestBid)
	if err := _Burnauction.contract.UnpackLog(event, "NewBestBid", log); err != nil {
		return nil, err
	}
	return event, nil
}
