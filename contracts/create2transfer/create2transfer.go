// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package create2transfer

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
	TxType     *big.Int
	FromIndex  *big.Int
	ToIndex    *big.Int
	ToPubkeyID *big.Int
	Amount     *big.Int
	Fee        *big.Int
	Nonce      *big.Int
}

// OffchainCreate2TransferWithPub is an auto generated low-level Go binding around an user-defined struct.
type OffchainCreate2TransferWithPub struct {
	TxType    *big.Int
	FromIndex *big.Int
	ToPubkey  [4]*big.Int
	Amount    *big.Int
	Fee       *big.Int
	Nonce     *big.Int
}

// TxCreate2Transfer is an auto generated low-level Go binding around an user-defined struct.
type TxCreate2Transfer struct {
	FromIndex  *big.Int
	ToIndex    *big.Int
	ToPubkeyID *big.Int
	Amount     *big.Int
	Fee        *big.Int
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
	PubkeyID *big.Int
	TokenID  *big.Int
	Balance  *big.Int
	Nonce    *big.Int
}

// Create2transferABI is the input ABI used to generate the binding from.
const Create2transferABI = "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeysSender\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnessesSender\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeysReceiver\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnessesReceiver\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProofWithReceiver\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"checkSignature\",\"outputs\":[{\"internalType\":\"enumTypes.Result\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"encodedTxs\",\"type\":\"bytes[]\"}],\"name\":\"compress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decode\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toPubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decodeWithPub\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"toPubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2TransferWithPub\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"decompress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toPubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTx.Create2Transfer[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toPubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"encode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"toPubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Create2TransferWithPub\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"encodeWithPub\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"name\":\"process\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTxWithPub\",\"type\":\"bytes\"}],\"name\":\"signBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkeySender\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkeyReceiver\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"validateAndApply\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newReceiver\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Create2transfer is an auto generated Go binding around an Ethereum contract.
type Create2transfer struct {
	Create2transferCaller     // Read-only binding to the contract
	Create2transferTransactor // Write-only binding to the contract
	Create2transferFilterer   // Log filterer for contract events
}

// Create2transferCaller is an auto generated read-only Go binding around an Ethereum contract.
type Create2transferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create2transferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Create2transferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create2transferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Create2transferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create2transferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Create2transferSession struct {
	Contract     *Create2transfer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Create2transferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Create2transferCallerSession struct {
	Contract *Create2transferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Create2transferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Create2transferTransactorSession struct {
	Contract     *Create2transferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Create2transferRaw is an auto generated low-level Go binding around an Ethereum contract.
type Create2transferRaw struct {
	Contract *Create2transfer // Generic contract binding to access the raw methods on
}

// Create2transferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Create2transferCallerRaw struct {
	Contract *Create2transferCaller // Generic read-only contract binding to access the raw methods on
}

// Create2transferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Create2transferTransactorRaw struct {
	Contract *Create2transferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreate2transfer creates a new instance of Create2transfer, bound to a specific deployed contract.
func NewCreate2transfer(address common.Address, backend bind.ContractBackend) (*Create2transfer, error) {
	contract, err := bindCreate2transfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Create2transfer{Create2transferCaller: Create2transferCaller{contract: contract}, Create2transferTransactor: Create2transferTransactor{contract: contract}, Create2transferFilterer: Create2transferFilterer{contract: contract}}, nil
}

// NewCreate2transferCaller creates a new read-only instance of Create2transfer, bound to a specific deployed contract.
func NewCreate2transferCaller(address common.Address, caller bind.ContractCaller) (*Create2transferCaller, error) {
	contract, err := bindCreate2transfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Create2transferCaller{contract: contract}, nil
}

// NewCreate2transferTransactor creates a new write-only instance of Create2transfer, bound to a specific deployed contract.
func NewCreate2transferTransactor(address common.Address, transactor bind.ContractTransactor) (*Create2transferTransactor, error) {
	contract, err := bindCreate2transfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Create2transferTransactor{contract: contract}, nil
}

// NewCreate2transferFilterer creates a new log filterer instance of Create2transfer, bound to a specific deployed contract.
func NewCreate2transferFilterer(address common.Address, filterer bind.ContractFilterer) (*Create2transferFilterer, error) {
	contract, err := bindCreate2transfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Create2transferFilterer{contract: contract}, nil
}

// bindCreate2transfer binds a generic wrapper to an already deployed contract.
func bindCreate2transfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Create2transferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Create2transfer *Create2transferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Create2transfer.Contract.Create2transferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Create2transfer *Create2transferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Create2transfer.Contract.Create2transferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Create2transfer *Create2transferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Create2transfer.Contract.Create2transferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Create2transfer *Create2transferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Create2transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Create2transfer *Create2transferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Create2transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Create2transfer *Create2transferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Create2transfer.Contract.contract.Transact(opts, method, params...)
}

// CheckSignature is a free data retrieval call binding the contract method 0xb9c7dfcc.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Create2transfer *Create2transferCaller) CheckSignature(opts *bind.CallOpts, signature [2]*big.Int, proof TypesSignatureProofWithReceiver, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "checkSignature", signature, proof, stateRoot, accountRoot, domain, txs)
	return *ret0, err
}

// CheckSignature is a free data retrieval call binding the contract method 0xb9c7dfcc.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Create2transfer *Create2transferSession) CheckSignature(signature [2]*big.Int, proof TypesSignatureProofWithReceiver, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	return _Create2transfer.Contract.CheckSignature(&_Create2transfer.CallOpts, signature, proof, stateRoot, accountRoot, domain, txs)
}

// CheckSignature is a free data retrieval call binding the contract method 0xb9c7dfcc.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Create2transfer *Create2transferCallerSession) CheckSignature(signature [2]*big.Int, proof TypesSignatureProofWithReceiver, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	return _Create2transfer.Contract.CheckSignature(&_Create2transfer.CallOpts, signature, proof, stateRoot, accountRoot, domain, txs)
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Create2transfer *Create2transferCaller) Compress(opts *bind.CallOpts, encodedTxs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "compress", encodedTxs)
	return *ret0, err
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Create2transfer *Create2transferSession) Compress(encodedTxs [][]byte) ([]byte, error) {
	return _Create2transfer.Contract.Compress(&_Create2transfer.CallOpts, encodedTxs)
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Create2transfer *Create2transferCallerSession) Compress(encodedTxs [][]byte) ([]byte, error) {
	return _Create2transfer.Contract.Compress(&_Create2transfer.CallOpts, encodedTxs)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Create2transfer *Create2transferCaller) Decode(opts *bind.CallOpts, encodedTx []byte) (OffchainCreate2Transfer, error) {
	var (
		ret0 = new(OffchainCreate2Transfer)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "decode", encodedTx)
	return *ret0, err
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Create2transfer *Create2transferSession) Decode(encodedTx []byte) (OffchainCreate2Transfer, error) {
	return _Create2transfer.Contract.Decode(&_Create2transfer.CallOpts, encodedTx)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Create2transfer *Create2transferCallerSession) Decode(encodedTx []byte) (OffchainCreate2Transfer, error) {
	return _Create2transfer.Contract.Decode(&_Create2transfer.CallOpts, encodedTx)
}

// DecodeWithPub is a free data retrieval call binding the contract method 0xe48411d3.
//
// Solidity: function decodeWithPub(bytes encodedTx) pure returns((uint256,uint256,uint256[4],uint256,uint256,uint256) _tx)
func (_Create2transfer *Create2transferCaller) DecodeWithPub(opts *bind.CallOpts, encodedTx []byte) (OffchainCreate2TransferWithPub, error) {
	var (
		ret0 = new(OffchainCreate2TransferWithPub)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "decodeWithPub", encodedTx)
	return *ret0, err
}

// DecodeWithPub is a free data retrieval call binding the contract method 0xe48411d3.
//
// Solidity: function decodeWithPub(bytes encodedTx) pure returns((uint256,uint256,uint256[4],uint256,uint256,uint256) _tx)
func (_Create2transfer *Create2transferSession) DecodeWithPub(encodedTx []byte) (OffchainCreate2TransferWithPub, error) {
	return _Create2transfer.Contract.DecodeWithPub(&_Create2transfer.CallOpts, encodedTx)
}

// DecodeWithPub is a free data retrieval call binding the contract method 0xe48411d3.
//
// Solidity: function decodeWithPub(bytes encodedTx) pure returns((uint256,uint256,uint256[4],uint256,uint256,uint256) _tx)
func (_Create2transfer *Create2transferCallerSession) DecodeWithPub(encodedTx []byte) (OffchainCreate2TransferWithPub, error) {
	return _Create2transfer.Contract.DecodeWithPub(&_Create2transfer.CallOpts, encodedTx)
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Create2transfer *Create2transferCaller) Decompress(opts *bind.CallOpts, txs []byte) ([]TxCreate2Transfer, error) {
	var (
		ret0 = new([]TxCreate2Transfer)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "decompress", txs)
	return *ret0, err
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Create2transfer *Create2transferSession) Decompress(txs []byte) ([]TxCreate2Transfer, error) {
	return _Create2transfer.Contract.Decompress(&_Create2transfer.CallOpts, txs)
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256,uint256,uint256)[])
func (_Create2transfer *Create2transferCallerSession) Decompress(txs []byte) ([]TxCreate2Transfer, error) {
	return _Create2transfer.Contract.Decompress(&_Create2transfer.CallOpts, txs)
}

// Encode is a free data retrieval call binding the contract method 0x6cc2d754.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Create2transfer *Create2transferCaller) Encode(opts *bind.CallOpts, _tx OffchainCreate2Transfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "encode", _tx)
	return *ret0, err
}

// Encode is a free data retrieval call binding the contract method 0x6cc2d754.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Create2transfer *Create2transferSession) Encode(_tx OffchainCreate2Transfer) ([]byte, error) {
	return _Create2transfer.Contract.Encode(&_Create2transfer.CallOpts, _tx)
}

// Encode is a free data retrieval call binding the contract method 0x6cc2d754.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Create2transfer *Create2transferCallerSession) Encode(_tx OffchainCreate2Transfer) ([]byte, error) {
	return _Create2transfer.Contract.Encode(&_Create2transfer.CallOpts, _tx)
}

// EncodeWithPub is a free data retrieval call binding the contract method 0xbba9dddc.
//
// Solidity: function encodeWithPub((uint256,uint256,uint256[4],uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Create2transfer *Create2transferCaller) EncodeWithPub(opts *bind.CallOpts, _tx OffchainCreate2TransferWithPub) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "encodeWithPub", _tx)
	return *ret0, err
}

// EncodeWithPub is a free data retrieval call binding the contract method 0xbba9dddc.
//
// Solidity: function encodeWithPub((uint256,uint256,uint256[4],uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Create2transfer *Create2transferSession) EncodeWithPub(_tx OffchainCreate2TransferWithPub) ([]byte, error) {
	return _Create2transfer.Contract.EncodeWithPub(&_Create2transfer.CallOpts, _tx)
}

// EncodeWithPub is a free data retrieval call binding the contract method 0xbba9dddc.
//
// Solidity: function encodeWithPub((uint256,uint256,uint256[4],uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Create2transfer *Create2transferCallerSession) EncodeWithPub(_tx OffchainCreate2TransferWithPub) ([]byte, error) {
	return _Create2transfer.Contract.EncodeWithPub(&_Create2transfer.CallOpts, _tx)
}

// Process is a free data retrieval call binding the contract method 0x89840761.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenID, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Create2transfer *Create2transferCaller) Process(opts *bind.CallOpts, stateRoot [32]byte, encodedTx []byte, tokenID *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	ret := new(struct {
		NewRoot [32]byte
		Result  uint8
	})
	out := ret
	err := _Create2transfer.contract.Call(opts, out, "process", stateRoot, encodedTx, tokenID, from, to)
	return *ret, err
}

// Process is a free data retrieval call binding the contract method 0x89840761.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenID, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Create2transfer *Create2transferSession) Process(stateRoot [32]byte, encodedTx []byte, tokenID *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	return _Create2transfer.Contract.Process(&_Create2transfer.CallOpts, stateRoot, encodedTx, tokenID, from, to)
}

// Process is a free data retrieval call binding the contract method 0x89840761.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenID, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Create2transfer *Create2transferCallerSession) Process(stateRoot [32]byte, encodedTx []byte, tokenID *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	return _Create2transfer.Contract.Process(&_Create2transfer.CallOpts, stateRoot, encodedTx, tokenID, from, to)
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTxWithPub) pure returns(bytes)
func (_Create2transfer *Create2transferCaller) SignBytes(opts *bind.CallOpts, encodedTxWithPub []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "signBytes", encodedTxWithPub)
	return *ret0, err
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTxWithPub) pure returns(bytes)
func (_Create2transfer *Create2transferSession) SignBytes(encodedTxWithPub []byte) ([]byte, error) {
	return _Create2transfer.Contract.SignBytes(&_Create2transfer.CallOpts, encodedTxWithPub)
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTxWithPub) pure returns(bytes)
func (_Create2transfer *Create2transferCallerSession) SignBytes(encodedTxWithPub []byte) ([]byte, error) {
	return _Create2transfer.Contract.SignBytes(&_Create2transfer.CallOpts, encodedTxWithPub)
}

// Validate is a free data retrieval call binding the contract method 0x074aa417.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkeySender, uint256[4] pubkeyReceiver, bytes32 domain) view returns(bool)
func (_Create2transfer *Create2transferCaller) Validate(opts *bind.CallOpts, encodedTx []byte, signature [2]*big.Int, pubkeySender [4]*big.Int, pubkeyReceiver [4]*big.Int, domain [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Create2transfer.contract.Call(opts, out, "validate", encodedTx, signature, pubkeySender, pubkeyReceiver, domain)
	return *ret0, err
}

// Validate is a free data retrieval call binding the contract method 0x074aa417.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkeySender, uint256[4] pubkeyReceiver, bytes32 domain) view returns(bool)
func (_Create2transfer *Create2transferSession) Validate(encodedTx []byte, signature [2]*big.Int, pubkeySender [4]*big.Int, pubkeyReceiver [4]*big.Int, domain [32]byte) (bool, error) {
	return _Create2transfer.Contract.Validate(&_Create2transfer.CallOpts, encodedTx, signature, pubkeySender, pubkeyReceiver, domain)
}

// Validate is a free data retrieval call binding the contract method 0x074aa417.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkeySender, uint256[4] pubkeyReceiver, bytes32 domain) view returns(bool)
func (_Create2transfer *Create2transferCallerSession) Validate(encodedTx []byte, signature [2]*big.Int, pubkeySender [4]*big.Int, pubkeyReceiver [4]*big.Int, domain [32]byte) (bool, error) {
	return _Create2transfer.Contract.Validate(&_Create2transfer.CallOpts, encodedTx, signature, pubkeySender, pubkeyReceiver, domain)
}

// ValidateAndApply is a free data retrieval call binding the contract method 0x5f7dde22.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Create2transfer *Create2transferCaller) ValidateAndApply(opts *bind.CallOpts, senderEncoded []byte, encodedTx []byte) (struct {
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
	err := _Create2transfer.contract.Call(opts, out, "validateAndApply", senderEncoded, encodedTx)
	return *ret, err
}

// ValidateAndApply is a free data retrieval call binding the contract method 0x5f7dde22.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Create2transfer *Create2transferSession) ValidateAndApply(senderEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Create2transfer.Contract.ValidateAndApply(&_Create2transfer.CallOpts, senderEncoded, encodedTx)
}

// ValidateAndApply is a free data retrieval call binding the contract method 0x5f7dde22.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Create2transfer *Create2transferCallerSession) ValidateAndApply(senderEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Create2transfer.Contract.ValidateAndApply(&_Create2transfer.CallOpts, senderEncoded, encodedTx)
}
