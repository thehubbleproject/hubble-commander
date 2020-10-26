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

// RollupclientABI is the input ABI used to generate the binding from.
const RollupclientABI = "[{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer[]\",\"name\":\"txs\",\"type\":\"tuple[]\"}],\"name\":\"compressCreate2Transfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration[]\",\"name\":\"txs\",\"type\":\"tuple[]\"}],\"name\":\"compressMassMigration\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer[]\",\"name\":\"txs\",\"type\":\"tuple[]\"}],\"name\":\"compressTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeCreate2Transfer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeMassMigration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeTransfer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"fromPubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"toPubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"valiateCreate2Transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"valiateMassMigration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"valiateTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"validateAndApplyCreate2Transfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newReceiver\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"validateAndApplyMassMigration\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"withdrawState\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiverEncoded\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"validateAndApplyTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newReceiver\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

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

// CompressCreate2Transfer is a free data retrieval call binding the contract method 0x5c3e8167.
//
// Solidity: function compressCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) CompressCreate2Transfer(opts *bind.CallOpts, txs []OffchainCreate2Transfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "compressCreate2Transfer", txs)
	return *ret0, err
}

// CompressCreate2Transfer is a free data retrieval call binding the contract method 0x5c3e8167.
//
// Solidity: function compressCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientSession) CompressCreate2Transfer(txs []OffchainCreate2Transfer) ([]byte, error) {
	return _Rollupclient.Contract.CompressCreate2Transfer(&_Rollupclient.CallOpts, txs)
}

// CompressCreate2Transfer is a free data retrieval call binding the contract method 0x5c3e8167.
//
// Solidity: function compressCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) CompressCreate2Transfer(txs []OffchainCreate2Transfer) ([]byte, error) {
	return _Rollupclient.Contract.CompressCreate2Transfer(&_Rollupclient.CallOpts, txs)
}

// CompressMassMigration is a free data retrieval call binding the contract method 0x2302f766.
//
// Solidity: function compressMassMigration((uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) CompressMassMigration(opts *bind.CallOpts, txs []OffchainMassMigration) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "compressMassMigration", txs)
	return *ret0, err
}

// CompressMassMigration is a free data retrieval call binding the contract method 0x2302f766.
//
// Solidity: function compressMassMigration((uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientSession) CompressMassMigration(txs []OffchainMassMigration) ([]byte, error) {
	return _Rollupclient.Contract.CompressMassMigration(&_Rollupclient.CallOpts, txs)
}

// CompressMassMigration is a free data retrieval call binding the contract method 0x2302f766.
//
// Solidity: function compressMassMigration((uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) CompressMassMigration(txs []OffchainMassMigration) ([]byte, error) {
	return _Rollupclient.Contract.CompressMassMigration(&_Rollupclient.CallOpts, txs)
}

// CompressTransfer is a free data retrieval call binding the contract method 0xf11ba323.
//
// Solidity: function compressTransfer((uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientCaller) CompressTransfer(opts *bind.CallOpts, txs []OffchainTransfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollupclient.contract.Call(opts, out, "compressTransfer", txs)
	return *ret0, err
}

// CompressTransfer is a free data retrieval call binding the contract method 0xf11ba323.
//
// Solidity: function compressTransfer((uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientSession) CompressTransfer(txs []OffchainTransfer) ([]byte, error) {
	return _Rollupclient.Contract.CompressTransfer(&_Rollupclient.CallOpts, txs)
}

// CompressTransfer is a free data retrieval call binding the contract method 0xf11ba323.
//
// Solidity: function compressTransfer((uint256,uint256,uint256,uint256,uint256,uint256)[] txs) pure returns(bytes)
func (_Rollupclient *RollupclientCallerSession) CompressTransfer(txs []OffchainTransfer) ([]byte, error) {
	return _Rollupclient.Contract.CompressTransfer(&_Rollupclient.CallOpts, txs)
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

// ValiateCreate2Transfer is a free data retrieval call binding the contract method 0xf17bafd5.
//
// Solidity: function valiateCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] fromPubkey, uint256[4] toPubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCaller) ValiateCreate2Transfer(opts *bind.CallOpts, _tx OffchainCreate2Transfer, signature [2]*big.Int, fromPubkey [4]*big.Int, toPubkey [4]*big.Int, domain [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Rollupclient.contract.Call(opts, out, "valiateCreate2Transfer", _tx, signature, fromPubkey, toPubkey, domain)
	return err
}

// ValiateCreate2Transfer is a free data retrieval call binding the contract method 0xf17bafd5.
//
// Solidity: function valiateCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] fromPubkey, uint256[4] toPubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientSession) ValiateCreate2Transfer(_tx OffchainCreate2Transfer, signature [2]*big.Int, fromPubkey [4]*big.Int, toPubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateCreate2Transfer(&_Rollupclient.CallOpts, _tx, signature, fromPubkey, toPubkey, domain)
}

// ValiateCreate2Transfer is a free data retrieval call binding the contract method 0xf17bafd5.
//
// Solidity: function valiateCreate2Transfer((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] fromPubkey, uint256[4] toPubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCallerSession) ValiateCreate2Transfer(_tx OffchainCreate2Transfer, signature [2]*big.Int, fromPubkey [4]*big.Int, toPubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateCreate2Transfer(&_Rollupclient.CallOpts, _tx, signature, fromPubkey, toPubkey, domain)
}

// ValiateMassMigration is a free data retrieval call binding the contract method 0x036654c5.
//
// Solidity: function valiateMassMigration((uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCaller) ValiateMassMigration(opts *bind.CallOpts, _tx OffchainMassMigration, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Rollupclient.contract.Call(opts, out, "valiateMassMigration", _tx, signature, pubkey, domain)
	return err
}

// ValiateMassMigration is a free data retrieval call binding the contract method 0x036654c5.
//
// Solidity: function valiateMassMigration((uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientSession) ValiateMassMigration(_tx OffchainMassMigration, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateMassMigration(&_Rollupclient.CallOpts, _tx, signature, pubkey, domain)
}

// ValiateMassMigration is a free data retrieval call binding the contract method 0x036654c5.
//
// Solidity: function valiateMassMigration((uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCallerSession) ValiateMassMigration(_tx OffchainMassMigration, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateMassMigration(&_Rollupclient.CallOpts, _tx, signature, pubkey, domain)
}

// ValiateTransfer is a free data retrieval call binding the contract method 0xe0fa9946.
//
// Solidity: function valiateTransfer((uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCaller) ValiateTransfer(opts *bind.CallOpts, _tx OffchainTransfer, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Rollupclient.contract.Call(opts, out, "valiateTransfer", _tx, signature, pubkey, domain)
	return err
}

// ValiateTransfer is a free data retrieval call binding the contract method 0xe0fa9946.
//
// Solidity: function valiateTransfer((uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientSession) ValiateTransfer(_tx OffchainTransfer, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateTransfer(&_Rollupclient.CallOpts, _tx, signature, pubkey, domain)
}

// ValiateTransfer is a free data retrieval call binding the contract method 0xe0fa9946.
//
// Solidity: function valiateTransfer((uint256,uint256,uint256,uint256,uint256,uint256) _tx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Rollupclient *RollupclientCallerSession) ValiateTransfer(_tx OffchainTransfer, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Rollupclient.Contract.ValiateTransfer(&_Rollupclient.CallOpts, _tx, signature, pubkey, domain)
}

// ValidateAndApplyCreate2Transfer is a free data retrieval call binding the contract method 0xccb302d2.
//
// Solidity: function validateAndApplyCreate2Transfer(bytes senderEncoded, (uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientCaller) ValidateAndApplyCreate2Transfer(opts *bind.CallOpts, senderEncoded []byte, _tx OffchainCreate2Transfer) (struct {
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
	err := _Rollupclient.contract.Call(opts, out, "validateAndApplyCreate2Transfer", senderEncoded, _tx)
	return *ret, err
}

// ValidateAndApplyCreate2Transfer is a free data retrieval call binding the contract method 0xccb302d2.
//
// Solidity: function validateAndApplyCreate2Transfer(bytes senderEncoded, (uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientSession) ValidateAndApplyCreate2Transfer(senderEncoded []byte, _tx OffchainCreate2Transfer) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyCreate2Transfer(&_Rollupclient.CallOpts, senderEncoded, _tx)
}

// ValidateAndApplyCreate2Transfer is a free data retrieval call binding the contract method 0xccb302d2.
//
// Solidity: function validateAndApplyCreate2Transfer(bytes senderEncoded, (uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ValidateAndApplyCreate2Transfer(senderEncoded []byte, _tx OffchainCreate2Transfer) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyCreate2Transfer(&_Rollupclient.CallOpts, senderEncoded, _tx)
}

// ValidateAndApplyMassMigration is a free data retrieval call binding the contract method 0x01c5909d.
//
// Solidity: function validateAndApplyMassMigration(bytes senderEncoded, (uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes withdrawState, uint8 result)
func (_Rollupclient *RollupclientCaller) ValidateAndApplyMassMigration(opts *bind.CallOpts, senderEncoded []byte, _tx OffchainMassMigration) (struct {
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
	err := _Rollupclient.contract.Call(opts, out, "validateAndApplyMassMigration", senderEncoded, _tx)
	return *ret, err
}

// ValidateAndApplyMassMigration is a free data retrieval call binding the contract method 0x01c5909d.
//
// Solidity: function validateAndApplyMassMigration(bytes senderEncoded, (uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes withdrawState, uint8 result)
func (_Rollupclient *RollupclientSession) ValidateAndApplyMassMigration(senderEncoded []byte, _tx OffchainMassMigration) (struct {
	NewSender     []byte
	WithdrawState []byte
	Result        uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyMassMigration(&_Rollupclient.CallOpts, senderEncoded, _tx)
}

// ValidateAndApplyMassMigration is a free data retrieval call binding the contract method 0x01c5909d.
//
// Solidity: function validateAndApplyMassMigration(bytes senderEncoded, (uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes withdrawState, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ValidateAndApplyMassMigration(senderEncoded []byte, _tx OffchainMassMigration) (struct {
	NewSender     []byte
	WithdrawState []byte
	Result        uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyMassMigration(&_Rollupclient.CallOpts, senderEncoded, _tx)
}

// ValidateAndApplyTransfer is a free data retrieval call binding the contract method 0x870e5856.
//
// Solidity: function validateAndApplyTransfer(bytes senderEncoded, bytes receiverEncoded, (uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientCaller) ValidateAndApplyTransfer(opts *bind.CallOpts, senderEncoded []byte, receiverEncoded []byte, _tx OffchainTransfer) (struct {
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
	err := _Rollupclient.contract.Call(opts, out, "validateAndApplyTransfer", senderEncoded, receiverEncoded, _tx)
	return *ret, err
}

// ValidateAndApplyTransfer is a free data retrieval call binding the contract method 0x870e5856.
//
// Solidity: function validateAndApplyTransfer(bytes senderEncoded, bytes receiverEncoded, (uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientSession) ValidateAndApplyTransfer(senderEncoded []byte, receiverEncoded []byte, _tx OffchainTransfer) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyTransfer(&_Rollupclient.CallOpts, senderEncoded, receiverEncoded, _tx)
}

// ValidateAndApplyTransfer is a free data retrieval call binding the contract method 0x870e5856.
//
// Solidity: function validateAndApplyTransfer(bytes senderEncoded, bytes receiverEncoded, (uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Rollupclient *RollupclientCallerSession) ValidateAndApplyTransfer(senderEncoded []byte, receiverEncoded []byte, _tx OffchainTransfer) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Rollupclient.Contract.ValidateAndApplyTransfer(&_Rollupclient.CallOpts, senderEncoded, receiverEncoded, _tx)
}
