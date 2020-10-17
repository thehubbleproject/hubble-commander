// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rolluputils

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

// TxTransfer is an auto generated low-level Go binding around an user-defined struct.
type TxTransfer struct {
	FromIndex *big.Int
	ToIndex   *big.Int
	Amount    *big.Int
	Fee       *big.Int
}

// TypesCreate2Transfer is an auto generated low-level Go binding around an user-defined struct.
type TypesCreate2Transfer struct {
	TxType    *big.Int
	FromIndex *big.Int
	ToIndex   *big.Int
	ToAccID   *big.Int
	Nonce     *big.Int
	Amount    *big.Int
	Fee       *big.Int
}

// TypesMassMigrationBody is an auto generated low-level Go binding around an user-defined struct.
type TypesMassMigrationBody struct {
	AccountRoot   [32]byte
	Signature     [2]*big.Int
	TargetSpokeID *big.Int
	WithdrawRoot  [32]byte
	TokenID       *big.Int
	Amount        *big.Int
	Txs           []byte
}

// TypesMassMigrationCommitment is an auto generated low-level Go binding around an user-defined struct.
type TypesMassMigrationCommitment struct {
	StateRoot [32]byte
	Body      TypesMassMigrationBody
}

// TypesTransfer is an auto generated low-level Go binding around an user-defined struct.
type TypesTransfer struct {
	TxType    *big.Int
	FromIndex *big.Int
	ToIndex   *big.Int
	Nonce     *big.Int
	Amount    *big.Int
	Fee       *big.Int
}

// TypesTransferBody is an auto generated low-level Go binding around an user-defined struct.
type TypesTransferBody struct {
	AccountRoot [32]byte
	Signature   [2]*big.Int
	TokenType   *big.Int
	FeeReceiver *big.Int
	Txs         []byte
}

// TypesTransferCommitment is an auto generated low-level Go binding around an user-defined struct.
type TypesTransferCommitment struct {
	StateRoot [32]byte
	Body      TypesTransferBody
}

// TypesUserState is an auto generated low-level Go binding around an user-defined struct.
type TypesUserState struct {
	PubkeyIndex *big.Int
	TokenType   *big.Int
	Balance     *big.Int
	Nonce       *big.Int
}

// RolluputilsABI is the input ABI used to generate the binding from.
const RolluputilsABI = "[{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"BytesFromState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"from\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"to\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BytesFromTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"txBytes\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"CompressManyTransferFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressTransferFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256[4]\",\"name\":\"from\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"to\",\"type\":\"uint256[4]\"}],\"name\":\"Create2IndexToPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"}],\"name\":\"Create2PubkeyToIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"Create2TransferFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAccID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Create2Transfer\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"DecompressManyTransfer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTx.Transfer[]\",\"name\":\"structTxs\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"DecompressTransfers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTx.Transfer[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetGenesisLeaves\",\"outputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"leaves\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"HashFromState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"targetSpokeID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.MassMigrationBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.MassMigrationCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"MMCommitmentToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"stateBytes\",\"type\":\"bytes\"}],\"name\":\"StateFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.TransferBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransferCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"TransferCommitmentToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"TxFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Transfer\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"from\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"to\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"getTxSignBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"getTxSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Rolluputils is an auto generated Go binding around an Ethereum contract.
type Rolluputils struct {
	RolluputilsCaller     // Read-only binding to the contract
	RolluputilsTransactor // Write-only binding to the contract
	RolluputilsFilterer   // Log filterer for contract events
}

// RolluputilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolluputilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolluputilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolluputilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolluputilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolluputilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolluputilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolluputilsSession struct {
	Contract     *Rolluputils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolluputilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolluputilsCallerSession struct {
	Contract *RolluputilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RolluputilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolluputilsTransactorSession struct {
	Contract     *RolluputilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RolluputilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolluputilsRaw struct {
	Contract *Rolluputils // Generic contract binding to access the raw methods on
}

// RolluputilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolluputilsCallerRaw struct {
	Contract *RolluputilsCaller // Generic read-only contract binding to access the raw methods on
}

// RolluputilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolluputilsTransactorRaw struct {
	Contract *RolluputilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRolluputils creates a new instance of Rolluputils, bound to a specific deployed contract.
func NewRolluputils(address common.Address, backend bind.ContractBackend) (*Rolluputils, error) {
	contract, err := bindRolluputils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rolluputils{RolluputilsCaller: RolluputilsCaller{contract: contract}, RolluputilsTransactor: RolluputilsTransactor{contract: contract}, RolluputilsFilterer: RolluputilsFilterer{contract: contract}}, nil
}

// NewRolluputilsCaller creates a new read-only instance of Rolluputils, bound to a specific deployed contract.
func NewRolluputilsCaller(address common.Address, caller bind.ContractCaller) (*RolluputilsCaller, error) {
	contract, err := bindRolluputils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolluputilsCaller{contract: contract}, nil
}

// NewRolluputilsTransactor creates a new write-only instance of Rolluputils, bound to a specific deployed contract.
func NewRolluputilsTransactor(address common.Address, transactor bind.ContractTransactor) (*RolluputilsTransactor, error) {
	contract, err := bindRolluputils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolluputilsTransactor{contract: contract}, nil
}

// NewRolluputilsFilterer creates a new log filterer instance of Rolluputils, bound to a specific deployed contract.
func NewRolluputilsFilterer(address common.Address, filterer bind.ContractFilterer) (*RolluputilsFilterer, error) {
	contract, err := bindRolluputils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolluputilsFilterer{contract: contract}, nil
}

// bindRolluputils binds a generic wrapper to an already deployed contract.
func bindRolluputils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolluputilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rolluputils *RolluputilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rolluputils.Contract.RolluputilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rolluputils *RolluputilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rolluputils.Contract.RolluputilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rolluputils *RolluputilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rolluputils.Contract.RolluputilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rolluputils *RolluputilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rolluputils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rolluputils *RolluputilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rolluputils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rolluputils *RolluputilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rolluputils.Contract.contract.Transact(opts, method, params...)
}

// BytesFromState is a free data retrieval call binding the contract method 0xc2e9bee3.
//
// Solidity: function BytesFromState((uint256,uint256,uint256,uint256) state) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromState(opts *bind.CallOpts, state TypesUserState) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromState", state)
	return *ret0, err
}

// BytesFromState is a free data retrieval call binding the contract method 0xc2e9bee3.
//
// Solidity: function BytesFromState((uint256,uint256,uint256,uint256) state) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromState(state TypesUserState) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromState(&_Rolluputils.CallOpts, state)
}

// BytesFromState is a free data retrieval call binding the contract method 0xc2e9bee3.
//
// Solidity: function BytesFromState((uint256,uint256,uint256,uint256) state) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromState(state TypesUserState) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromState(&_Rolluputils.CallOpts, state)
}

// BytesFromTx is a free data retrieval call binding the contract method 0x760c66ee.
//
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTx(opts *bind.CallOpts, _tx TypesTransfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTx", _tx)
	return *ret0, err
}

// BytesFromTx is a free data retrieval call binding the contract method 0x760c66ee.
//
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTx(_tx TypesTransfer) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTx is a free data retrieval call binding the contract method 0x760c66ee.
//
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTx(_tx TypesTransfer) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTx0 is a free data retrieval call binding the contract method 0x7d11ae08.
//
// Solidity: function BytesFromTx(uint256 txType, uint256[4] from, uint256[4] to, uint256 toAccID, uint256 nonce, uint256 amount, uint256 fee) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTx0(opts *bind.CallOpts, txType *big.Int, from [4]*big.Int, to [4]*big.Int, toAccID *big.Int, nonce *big.Int, amount *big.Int, fee *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTx0", txType, from, to, toAccID, nonce, amount, fee)
	return *ret0, err
}

// BytesFromTx0 is a free data retrieval call binding the contract method 0x7d11ae08.
//
// Solidity: function BytesFromTx(uint256 txType, uint256[4] from, uint256[4] to, uint256 toAccID, uint256 nonce, uint256 amount, uint256 fee) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTx0(txType *big.Int, from [4]*big.Int, to [4]*big.Int, toAccID *big.Int, nonce *big.Int, amount *big.Int, fee *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx0(&_Rolluputils.CallOpts, txType, from, to, toAccID, nonce, amount, fee)
}

// BytesFromTx0 is a free data retrieval call binding the contract method 0x7d11ae08.
//
// Solidity: function BytesFromTx(uint256 txType, uint256[4] from, uint256[4] to, uint256 toAccID, uint256 nonce, uint256 amount, uint256 fee) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTx0(txType *big.Int, from [4]*big.Int, to [4]*big.Int, toAccID *big.Int, nonce *big.Int, amount *big.Int, fee *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx0(&_Rolluputils.CallOpts, txType, from, to, toAccID, nonce, amount, fee)
}

// BytesFromTx1 is a free data retrieval call binding the contract method 0xbe6c2af2.
//
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTx1(opts *bind.CallOpts, _tx TypesCreate2Transfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTx1", _tx)
	return *ret0, err
}

// BytesFromTx1 is a free data retrieval call binding the contract method 0xbe6c2af2.
//
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTx1(_tx TypesCreate2Transfer) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx1(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTx1 is a free data retrieval call binding the contract method 0xbe6c2af2.
//
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTx1(_tx TypesCreate2Transfer) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx1(&_Rolluputils.CallOpts, _tx)
}

// CompressManyTransferFromEncoded is a free data retrieval call binding the contract method 0x8ee4ac06.
//
// Solidity: function CompressManyTransferFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressManyTransferFromEncoded(opts *bind.CallOpts, txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressManyTransferFromEncoded", txBytes, sigs)
	return *ret0, err
}

// CompressManyTransferFromEncoded is a free data retrieval call binding the contract method 0x8ee4ac06.
//
// Solidity: function CompressManyTransferFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressManyTransferFromEncoded(txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyTransferFromEncoded(&_Rolluputils.CallOpts, txBytes, sigs)
}

// CompressManyTransferFromEncoded is a free data retrieval call binding the contract method 0x8ee4ac06.
//
// Solidity: function CompressManyTransferFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressManyTransferFromEncoded(txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyTransferFromEncoded(&_Rolluputils.CallOpts, txBytes, sigs)
}

// CompressTransferFromEncoded is a free data retrieval call binding the contract method 0x57ed1f1f.
//
// Solidity: function CompressTransferFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressTransferFromEncoded(opts *bind.CallOpts, txBytes []byte, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressTransferFromEncoded", txBytes, sig)
	return *ret0, err
}

// CompressTransferFromEncoded is a free data retrieval call binding the contract method 0x57ed1f1f.
//
// Solidity: function CompressTransferFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressTransferFromEncoded(txBytes []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressTransferFromEncoded(&_Rolluputils.CallOpts, txBytes, sig)
}

// CompressTransferFromEncoded is a free data retrieval call binding the contract method 0x57ed1f1f.
//
// Solidity: function CompressTransferFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressTransferFromEncoded(txBytes []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressTransferFromEncoded(&_Rolluputils.CallOpts, txBytes, sig)
}

// Create2IndexToPubkey is a free data retrieval call binding the contract method 0xf31b8bd2.
//
// Solidity: function Create2IndexToPubkey(bytes txBytes, uint256[4] from, uint256[4] to) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) Create2IndexToPubkey(opts *bind.CallOpts, txBytes []byte, from [4]*big.Int, to [4]*big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "Create2IndexToPubkey", txBytes, from, to)
	return *ret0, err
}

// Create2IndexToPubkey is a free data retrieval call binding the contract method 0xf31b8bd2.
//
// Solidity: function Create2IndexToPubkey(bytes txBytes, uint256[4] from, uint256[4] to) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) Create2IndexToPubkey(txBytes []byte, from [4]*big.Int, to [4]*big.Int) ([]byte, error) {
	return _Rolluputils.Contract.Create2IndexToPubkey(&_Rolluputils.CallOpts, txBytes, from, to)
}

// Create2IndexToPubkey is a free data retrieval call binding the contract method 0xf31b8bd2.
//
// Solidity: function Create2IndexToPubkey(bytes txBytes, uint256[4] from, uint256[4] to) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) Create2IndexToPubkey(txBytes []byte, from [4]*big.Int, to [4]*big.Int) ([]byte, error) {
	return _Rolluputils.Contract.Create2IndexToPubkey(&_Rolluputils.CallOpts, txBytes, from, to)
}

// Create2PubkeyToIndex is a free data retrieval call binding the contract method 0xe8791d76.
//
// Solidity: function Create2PubkeyToIndex(bytes txBytes, uint256 from, uint256 to, uint256 toAccID) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) Create2PubkeyToIndex(opts *bind.CallOpts, txBytes []byte, from *big.Int, to *big.Int, toAccID *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "Create2PubkeyToIndex", txBytes, from, to, toAccID)
	return *ret0, err
}

// Create2PubkeyToIndex is a free data retrieval call binding the contract method 0xe8791d76.
//
// Solidity: function Create2PubkeyToIndex(bytes txBytes, uint256 from, uint256 to, uint256 toAccID) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) Create2PubkeyToIndex(txBytes []byte, from *big.Int, to *big.Int, toAccID *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.Create2PubkeyToIndex(&_Rolluputils.CallOpts, txBytes, from, to, toAccID)
}

// Create2PubkeyToIndex is a free data retrieval call binding the contract method 0xe8791d76.
//
// Solidity: function Create2PubkeyToIndex(bytes txBytes, uint256 from, uint256 to, uint256 toAccID) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) Create2PubkeyToIndex(txBytes []byte, from *big.Int, to *big.Int, toAccID *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.Create2PubkeyToIndex(&_Rolluputils.CallOpts, txBytes, from, to, toAccID)
}

// Create2TransferFromBytes is a free data retrieval call binding the contract method 0x352dab36.
//
// Solidity: function Create2TransferFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsCaller) Create2TransferFromBytes(opts *bind.CallOpts, txBytes []byte) (TypesCreate2Transfer, error) {
	var (
		ret0 = new(TypesCreate2Transfer)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "Create2TransferFromBytes", txBytes)
	return *ret0, err
}

// Create2TransferFromBytes is a free data retrieval call binding the contract method 0x352dab36.
//
// Solidity: function Create2TransferFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsSession) Create2TransferFromBytes(txBytes []byte) (TypesCreate2Transfer, error) {
	return _Rolluputils.Contract.Create2TransferFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// Create2TransferFromBytes is a free data retrieval call binding the contract method 0x352dab36.
//
// Solidity: function Create2TransferFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsCallerSession) Create2TransferFromBytes(txBytes []byte) (TypesCreate2Transfer, error) {
	return _Rolluputils.Contract.Create2TransferFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// DecompressManyTransfer is a free data retrieval call binding the contract method 0xd1b4bf51.
//
// Solidity: function DecompressManyTransfer(bytes txs) pure returns((uint256,uint256,uint256,uint256)[] structTxs)
func (_Rolluputils *RolluputilsCaller) DecompressManyTransfer(opts *bind.CallOpts, txs []byte) ([]TxTransfer, error) {
	var (
		ret0 = new([]TxTransfer)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressManyTransfer", txs)
	return *ret0, err
}

// DecompressManyTransfer is a free data retrieval call binding the contract method 0xd1b4bf51.
//
// Solidity: function DecompressManyTransfer(bytes txs) pure returns((uint256,uint256,uint256,uint256)[] structTxs)
func (_Rolluputils *RolluputilsSession) DecompressManyTransfer(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressManyTransfer(&_Rolluputils.CallOpts, txs)
}

// DecompressManyTransfer is a free data retrieval call binding the contract method 0xd1b4bf51.
//
// Solidity: function DecompressManyTransfer(bytes txs) pure returns((uint256,uint256,uint256,uint256)[] structTxs)
func (_Rolluputils *RolluputilsCallerSession) DecompressManyTransfer(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressManyTransfer(&_Rolluputils.CallOpts, txs)
}

// DecompressTransfers is a free data retrieval call binding the contract method 0x67347755.
//
// Solidity: function DecompressTransfers(bytes txs) pure returns((uint256,uint256,uint256,uint256)[])
func (_Rolluputils *RolluputilsCaller) DecompressTransfers(opts *bind.CallOpts, txs []byte) ([]TxTransfer, error) {
	var (
		ret0 = new([]TxTransfer)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressTransfers", txs)
	return *ret0, err
}

// DecompressTransfers is a free data retrieval call binding the contract method 0x67347755.
//
// Solidity: function DecompressTransfers(bytes txs) pure returns((uint256,uint256,uint256,uint256)[])
func (_Rolluputils *RolluputilsSession) DecompressTransfers(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressTransfers(&_Rolluputils.CallOpts, txs)
}

// DecompressTransfers is a free data retrieval call binding the contract method 0x67347755.
//
// Solidity: function DecompressTransfers(bytes txs) pure returns((uint256,uint256,uint256,uint256)[])
func (_Rolluputils *RolluputilsCallerSession) DecompressTransfers(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressTransfers(&_Rolluputils.CallOpts, txs)
}

// GetGenesisLeaves is a free data retrieval call binding the contract method 0x3043be91.
//
// Solidity: function GetGenesisLeaves() pure returns(bytes32[2] leaves)
func (_Rolluputils *RolluputilsCaller) GetGenesisLeaves(opts *bind.CallOpts) ([2][32]byte, error) {
	var (
		ret0 = new([2][32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "GetGenesisLeaves")
	return *ret0, err
}

// GetGenesisLeaves is a free data retrieval call binding the contract method 0x3043be91.
//
// Solidity: function GetGenesisLeaves() pure returns(bytes32[2] leaves)
func (_Rolluputils *RolluputilsSession) GetGenesisLeaves() ([2][32]byte, error) {
	return _Rolluputils.Contract.GetGenesisLeaves(&_Rolluputils.CallOpts)
}

// GetGenesisLeaves is a free data retrieval call binding the contract method 0x3043be91.
//
// Solidity: function GetGenesisLeaves() pure returns(bytes32[2] leaves)
func (_Rolluputils *RolluputilsCallerSession) GetGenesisLeaves() ([2][32]byte, error) {
	return _Rolluputils.Contract.GetGenesisLeaves(&_Rolluputils.CallOpts)
}

// HashFromState is a free data retrieval call binding the contract method 0xa7f5dc9a.
//
// Solidity: function HashFromState((uint256,uint256,uint256,uint256) state) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromState(opts *bind.CallOpts, state TypesUserState) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromState", state)
	return *ret0, err
}

// HashFromState is a free data retrieval call binding the contract method 0xa7f5dc9a.
//
// Solidity: function HashFromState((uint256,uint256,uint256,uint256) state) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromState(state TypesUserState) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromState(&_Rolluputils.CallOpts, state)
}

// HashFromState is a free data retrieval call binding the contract method 0xa7f5dc9a.
//
// Solidity: function HashFromState((uint256,uint256,uint256,uint256) state) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromState(state TypesUserState) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromState(&_Rolluputils.CallOpts, state)
}

// HashFromTx is a free data retrieval call binding the contract method 0x84c1ac63.
//
// Solidity: function HashFromTx((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromTx(opts *bind.CallOpts, _tx TypesTransfer) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromTx", _tx)
	return *ret0, err
}

// HashFromTx is a free data retrieval call binding the contract method 0x84c1ac63.
//
// Solidity: function HashFromTx((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromTx(_tx TypesTransfer) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromTx(&_Rolluputils.CallOpts, _tx)
}

// HashFromTx is a free data retrieval call binding the contract method 0x84c1ac63.
//
// Solidity: function HashFromTx((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromTx(_tx TypesTransfer) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromTx(&_Rolluputils.CallOpts, _tx)
}

// MMCommitmentToHash is a free data retrieval call binding the contract method 0xdc280c1b.
//
// Solidity: function MMCommitmentToHash((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,bytes)) commitment) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) MMCommitmentToHash(opts *bind.CallOpts, commitment TypesMassMigrationCommitment) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "MMCommitmentToHash", commitment)
	return *ret0, err
}

// MMCommitmentToHash is a free data retrieval call binding the contract method 0xdc280c1b.
//
// Solidity: function MMCommitmentToHash((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,bytes)) commitment) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) MMCommitmentToHash(commitment TypesMassMigrationCommitment) ([32]byte, error) {
	return _Rolluputils.Contract.MMCommitmentToHash(&_Rolluputils.CallOpts, commitment)
}

// MMCommitmentToHash is a free data retrieval call binding the contract method 0xdc280c1b.
//
// Solidity: function MMCommitmentToHash((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,bytes)) commitment) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) MMCommitmentToHash(commitment TypesMassMigrationCommitment) ([32]byte, error) {
	return _Rolluputils.Contract.MMCommitmentToHash(&_Rolluputils.CallOpts, commitment)
}

// StateFromBytes is a free data retrieval call binding the contract method 0x717ebc5c.
//
// Solidity: function StateFromBytes(bytes stateBytes) pure returns((uint256,uint256,uint256,uint256) state)
func (_Rolluputils *RolluputilsCaller) StateFromBytes(opts *bind.CallOpts, stateBytes []byte) (TypesUserState, error) {
	var (
		ret0 = new(TypesUserState)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "StateFromBytes", stateBytes)
	return *ret0, err
}

// StateFromBytes is a free data retrieval call binding the contract method 0x717ebc5c.
//
// Solidity: function StateFromBytes(bytes stateBytes) pure returns((uint256,uint256,uint256,uint256) state)
func (_Rolluputils *RolluputilsSession) StateFromBytes(stateBytes []byte) (TypesUserState, error) {
	return _Rolluputils.Contract.StateFromBytes(&_Rolluputils.CallOpts, stateBytes)
}

// StateFromBytes is a free data retrieval call binding the contract method 0x717ebc5c.
//
// Solidity: function StateFromBytes(bytes stateBytes) pure returns((uint256,uint256,uint256,uint256) state)
func (_Rolluputils *RolluputilsCallerSession) StateFromBytes(stateBytes []byte) (TypesUserState, error) {
	return _Rolluputils.Contract.StateFromBytes(&_Rolluputils.CallOpts, stateBytes)
}

// TransferCommitmentToHash is a free data retrieval call binding the contract method 0xadd9019b.
//
// Solidity: function TransferCommitmentToHash((bytes32,(bytes32,uint256[2],uint256,uint256,bytes)) commitment) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) TransferCommitmentToHash(opts *bind.CallOpts, commitment TypesTransferCommitment) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "TransferCommitmentToHash", commitment)
	return *ret0, err
}

// TransferCommitmentToHash is a free data retrieval call binding the contract method 0xadd9019b.
//
// Solidity: function TransferCommitmentToHash((bytes32,(bytes32,uint256[2],uint256,uint256,bytes)) commitment) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) TransferCommitmentToHash(commitment TypesTransferCommitment) ([32]byte, error) {
	return _Rolluputils.Contract.TransferCommitmentToHash(&_Rolluputils.CallOpts, commitment)
}

// TransferCommitmentToHash is a free data retrieval call binding the contract method 0xadd9019b.
//
// Solidity: function TransferCommitmentToHash((bytes32,(bytes32,uint256[2],uint256,uint256,bytes)) commitment) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) TransferCommitmentToHash(commitment TypesTransferCommitment) ([32]byte, error) {
	return _Rolluputils.Contract.TransferCommitmentToHash(&_Rolluputils.CallOpts, commitment)
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsCaller) TxFromBytes(opts *bind.CallOpts, txBytes []byte) (TypesTransfer, error) {
	var (
		ret0 = new(TypesTransfer)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "TxFromBytes", txBytes)
	return *ret0, err
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsSession) TxFromBytes(txBytes []byte) (TypesTransfer, error) {
	return _Rolluputils.Contract.TxFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsCallerSession) TxFromBytes(txBytes []byte) (TypesTransfer, error) {
	return _Rolluputils.Contract.TxFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x6836802f.
//
// Solidity: function getTxSignBytes(uint256 txType, uint256[4] from, uint256[4] to, uint256 nonce, uint256 amount, uint256 fee) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) GetTxSignBytes(opts *bind.CallOpts, txType *big.Int, from [4]*big.Int, to [4]*big.Int, nonce *big.Int, amount *big.Int, fee *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getTxSignBytes", txType, from, to, nonce, amount, fee)
	return *ret0, err
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x6836802f.
//
// Solidity: function getTxSignBytes(uint256 txType, uint256[4] from, uint256[4] to, uint256 nonce, uint256 amount, uint256 fee) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) GetTxSignBytes(txType *big.Int, from [4]*big.Int, to [4]*big.Int, nonce *big.Int, amount *big.Int, fee *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes(&_Rolluputils.CallOpts, txType, from, to, nonce, amount, fee)
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x6836802f.
//
// Solidity: function getTxSignBytes(uint256 txType, uint256[4] from, uint256[4] to, uint256 nonce, uint256 amount, uint256 fee) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) GetTxSignBytes(txType *big.Int, from [4]*big.Int, to [4]*big.Int, nonce *big.Int, amount *big.Int, fee *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes(&_Rolluputils.CallOpts, txType, from, to, nonce, amount, fee)
}

// GetTxSignBytes0 is a free data retrieval call binding the contract method 0x902c3a06.
//
// Solidity: function getTxSignBytes((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) GetTxSignBytes0(opts *bind.CallOpts, _tx TypesTransfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getTxSignBytes0", _tx)
	return *ret0, err
}

// GetTxSignBytes0 is a free data retrieval call binding the contract method 0x902c3a06.
//
// Solidity: function getTxSignBytes((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) GetTxSignBytes0(_tx TypesTransfer) ([]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes0(&_Rolluputils.CallOpts, _tx)
}

// GetTxSignBytes0 is a free data retrieval call binding the contract method 0x902c3a06.
//
// Solidity: function getTxSignBytes((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) GetTxSignBytes0(_tx TypesTransfer) ([]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes0(&_Rolluputils.CallOpts, _tx)
}
