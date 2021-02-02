// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package massmigration

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

// OffchainMassMigration is an auto generated low-level Go binding around an user-defined struct.
type OffchainMassMigration struct {
	TxType    *big.Int
	FromIndex *big.Int
	Amount    *big.Int
	Fee       *big.Int
	SpokeID   *big.Int
	Nonce     *big.Int
}

// TxMassMigration is an auto generated low-level Go binding around an user-defined struct.
type TxMassMigration struct {
	FromIndex *big.Int
	Amount    *big.Int
	Fee       *big.Int
}

// TypesSignatureProof is an auto generated low-level Go binding around an user-defined struct.
type TypesSignatureProof struct {
	States          []TypesUserState
	StateWitnesses  [][][32]byte
	Pubkeys         [][4]*big.Int
	PubkeyWitnesses [][][32]byte
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

// MassmigrationABI is the input ABI used to generate the binding from.
const MassmigrationABI = "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnesses\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"checkSignature\",\"outputs\":[{\"internalType\":\"enumTypes.Result\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"encodedTxs\",\"type\":\"bytes[]\"}],\"name\":\"compress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decode\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"decompress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTx.MassMigration[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.MassMigration\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"encode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"}],\"name\":\"process\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"freshState\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"signBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"validateAndApply\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newReceiver\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Massmigration is an auto generated Go binding around an Ethereum contract.
type Massmigration struct {
	MassmigrationCaller     // Read-only binding to the contract
	MassmigrationTransactor // Write-only binding to the contract
	MassmigrationFilterer   // Log filterer for contract events
}

// MassmigrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type MassmigrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MassmigrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MassmigrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MassmigrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MassmigrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MassmigrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MassmigrationSession struct {
	Contract     *Massmigration    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MassmigrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MassmigrationCallerSession struct {
	Contract *MassmigrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MassmigrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MassmigrationTransactorSession struct {
	Contract     *MassmigrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MassmigrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type MassmigrationRaw struct {
	Contract *Massmigration // Generic contract binding to access the raw methods on
}

// MassmigrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MassmigrationCallerRaw struct {
	Contract *MassmigrationCaller // Generic read-only contract binding to access the raw methods on
}

// MassmigrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MassmigrationTransactorRaw struct {
	Contract *MassmigrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMassmigration creates a new instance of Massmigration, bound to a specific deployed contract.
func NewMassmigration(address common.Address, backend bind.ContractBackend) (*Massmigration, error) {
	contract, err := bindMassmigration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Massmigration{MassmigrationCaller: MassmigrationCaller{contract: contract}, MassmigrationTransactor: MassmigrationTransactor{contract: contract}, MassmigrationFilterer: MassmigrationFilterer{contract: contract}}, nil
}

// NewMassmigrationCaller creates a new read-only instance of Massmigration, bound to a specific deployed contract.
func NewMassmigrationCaller(address common.Address, caller bind.ContractCaller) (*MassmigrationCaller, error) {
	contract, err := bindMassmigration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MassmigrationCaller{contract: contract}, nil
}

// NewMassmigrationTransactor creates a new write-only instance of Massmigration, bound to a specific deployed contract.
func NewMassmigrationTransactor(address common.Address, transactor bind.ContractTransactor) (*MassmigrationTransactor, error) {
	contract, err := bindMassmigration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MassmigrationTransactor{contract: contract}, nil
}

// NewMassmigrationFilterer creates a new log filterer instance of Massmigration, bound to a specific deployed contract.
func NewMassmigrationFilterer(address common.Address, filterer bind.ContractFilterer) (*MassmigrationFilterer, error) {
	contract, err := bindMassmigration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MassmigrationFilterer{contract: contract}, nil
}

// bindMassmigration binds a generic wrapper to an already deployed contract.
func bindMassmigration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MassmigrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Massmigration *MassmigrationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Massmigration.Contract.MassmigrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Massmigration *MassmigrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Massmigration.Contract.MassmigrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Massmigration *MassmigrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Massmigration.Contract.MassmigrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Massmigration *MassmigrationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Massmigration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Massmigration *MassmigrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Massmigration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Massmigration *MassmigrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Massmigration.Contract.contract.Transact(opts, method, params...)
}

// CheckSignature is a free data retrieval call binding the contract method 0x997ee238.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, uint256 spokeID, bytes txs) view returns(uint8)
func (_Massmigration *MassmigrationCaller) CheckSignature(opts *bind.CallOpts, signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, spokeID *big.Int, txs []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Massmigration.contract.Call(opts, out, "checkSignature", signature, proof, stateRoot, accountRoot, domain, spokeID, txs)
	return *ret0, err
}

// CheckSignature is a free data retrieval call binding the contract method 0x997ee238.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, uint256 spokeID, bytes txs) view returns(uint8)
func (_Massmigration *MassmigrationSession) CheckSignature(signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, spokeID *big.Int, txs []byte) (uint8, error) {
	return _Massmigration.Contract.CheckSignature(&_Massmigration.CallOpts, signature, proof, stateRoot, accountRoot, domain, spokeID, txs)
}

// CheckSignature is a free data retrieval call binding the contract method 0x997ee238.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, uint256 spokeID, bytes txs) view returns(uint8)
func (_Massmigration *MassmigrationCallerSession) CheckSignature(signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, spokeID *big.Int, txs []byte) (uint8, error) {
	return _Massmigration.Contract.CheckSignature(&_Massmigration.CallOpts, signature, proof, stateRoot, accountRoot, domain, spokeID, txs)
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Massmigration *MassmigrationCaller) Compress(opts *bind.CallOpts, encodedTxs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Massmigration.contract.Call(opts, out, "compress", encodedTxs)
	return *ret0, err
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Massmigration *MassmigrationSession) Compress(encodedTxs [][]byte) ([]byte, error) {
	return _Massmigration.Contract.Compress(&_Massmigration.CallOpts, encodedTxs)
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Massmigration *MassmigrationCallerSession) Compress(encodedTxs [][]byte) ([]byte, error) {
	return _Massmigration.Contract.Compress(&_Massmigration.CallOpts, encodedTxs)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Massmigration *MassmigrationCaller) Decode(opts *bind.CallOpts, encodedTx []byte) (OffchainMassMigration, error) {
	var (
		ret0 = new(OffchainMassMigration)
	)
	out := ret0
	err := _Massmigration.contract.Call(opts, out, "decode", encodedTx)
	return *ret0, err
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Massmigration *MassmigrationSession) Decode(encodedTx []byte) (OffchainMassMigration, error) {
	return _Massmigration.Contract.Decode(&_Massmigration.CallOpts, encodedTx)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Massmigration *MassmigrationCallerSession) Decode(encodedTx []byte) (OffchainMassMigration, error) {
	return _Massmigration.Contract.Decode(&_Massmigration.CallOpts, encodedTx)
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256)[])
func (_Massmigration *MassmigrationCaller) Decompress(opts *bind.CallOpts, txs []byte) ([]TxMassMigration, error) {
	var (
		ret0 = new([]TxMassMigration)
	)
	out := ret0
	err := _Massmigration.contract.Call(opts, out, "decompress", txs)
	return *ret0, err
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256)[])
func (_Massmigration *MassmigrationSession) Decompress(txs []byte) ([]TxMassMigration, error) {
	return _Massmigration.Contract.Decompress(&_Massmigration.CallOpts, txs)
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256)[])
func (_Massmigration *MassmigrationCallerSession) Decompress(txs []byte) ([]TxMassMigration, error) {
	return _Massmigration.Contract.Decompress(&_Massmigration.CallOpts, txs)
}

// Encode is a free data retrieval call binding the contract method 0x08909a83.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Massmigration *MassmigrationCaller) Encode(opts *bind.CallOpts, _tx OffchainMassMigration) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Massmigration.contract.Call(opts, out, "encode", _tx)
	return *ret0, err
}

// Encode is a free data retrieval call binding the contract method 0x08909a83.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Massmigration *MassmigrationSession) Encode(_tx OffchainMassMigration) ([]byte, error) {
	return _Massmigration.Contract.Encode(&_Massmigration.CallOpts, _tx)
}

// Encode is a free data retrieval call binding the contract method 0x08909a83.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Massmigration *MassmigrationCallerSession) Encode(_tx OffchainMassMigration) ([]byte, error) {
	return _Massmigration.Contract.Encode(&_Massmigration.CallOpts, _tx)
}

// Process is a free data retrieval call binding the contract method 0x73b3ba35.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenID, ((uint256,uint256,uint256,uint256),bytes32[]) from) pure returns(bytes32 newRoot, bytes freshState, uint8 result)
func (_Massmigration *MassmigrationCaller) Process(opts *bind.CallOpts, stateRoot [32]byte, encodedTx []byte, tokenID *big.Int, from TypesStateMerkleProof) (struct {
	NewRoot    [32]byte
	FreshState []byte
	Result     uint8
}, error) {
	ret := new(struct {
		NewRoot    [32]byte
		FreshState []byte
		Result     uint8
	})
	out := ret
	err := _Massmigration.contract.Call(opts, out, "process", stateRoot, encodedTx, tokenID, from)
	return *ret, err
}

// Process is a free data retrieval call binding the contract method 0x73b3ba35.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenID, ((uint256,uint256,uint256,uint256),bytes32[]) from) pure returns(bytes32 newRoot, bytes freshState, uint8 result)
func (_Massmigration *MassmigrationSession) Process(stateRoot [32]byte, encodedTx []byte, tokenID *big.Int, from TypesStateMerkleProof) (struct {
	NewRoot    [32]byte
	FreshState []byte
	Result     uint8
}, error) {
	return _Massmigration.Contract.Process(&_Massmigration.CallOpts, stateRoot, encodedTx, tokenID, from)
}

// Process is a free data retrieval call binding the contract method 0x73b3ba35.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenID, ((uint256,uint256,uint256,uint256),bytes32[]) from) pure returns(bytes32 newRoot, bytes freshState, uint8 result)
func (_Massmigration *MassmigrationCallerSession) Process(stateRoot [32]byte, encodedTx []byte, tokenID *big.Int, from TypesStateMerkleProof) (struct {
	NewRoot    [32]byte
	FreshState []byte
	Result     uint8
}, error) {
	return _Massmigration.Contract.Process(&_Massmigration.CallOpts, stateRoot, encodedTx, tokenID, from)
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTx) pure returns(bytes)
func (_Massmigration *MassmigrationCaller) SignBytes(opts *bind.CallOpts, encodedTx []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Massmigration.contract.Call(opts, out, "signBytes", encodedTx)
	return *ret0, err
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTx) pure returns(bytes)
func (_Massmigration *MassmigrationSession) SignBytes(encodedTx []byte) ([]byte, error) {
	return _Massmigration.Contract.SignBytes(&_Massmigration.CallOpts, encodedTx)
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTx) pure returns(bytes)
func (_Massmigration *MassmigrationCallerSession) SignBytes(encodedTx []byte) ([]byte, error) {
	return _Massmigration.Contract.SignBytes(&_Massmigration.CallOpts, encodedTx)
}

// Validate is a free data retrieval call binding the contract method 0x5ef63f4f.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns(bool)
func (_Massmigration *MassmigrationCaller) Validate(opts *bind.CallOpts, encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Massmigration.contract.Call(opts, out, "validate", encodedTx, signature, pubkey, domain)
	return *ret0, err
}

// Validate is a free data retrieval call binding the contract method 0x5ef63f4f.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns(bool)
func (_Massmigration *MassmigrationSession) Validate(encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) (bool, error) {
	return _Massmigration.Contract.Validate(&_Massmigration.CallOpts, encodedTx, signature, pubkey, domain)
}

// Validate is a free data retrieval call binding the contract method 0x5ef63f4f.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns(bool)
func (_Massmigration *MassmigrationCallerSession) Validate(encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) (bool, error) {
	return _Massmigration.Contract.Validate(&_Massmigration.CallOpts, encodedTx, signature, pubkey, domain)
}

// ValidateAndApply is a free data retrieval call binding the contract method 0x5f7dde22.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Massmigration *MassmigrationCaller) ValidateAndApply(opts *bind.CallOpts, senderEncoded []byte, encodedTx []byte) (struct {
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
	err := _Massmigration.contract.Call(opts, out, "validateAndApply", senderEncoded, encodedTx)
	return *ret, err
}

// ValidateAndApply is a free data retrieval call binding the contract method 0x5f7dde22.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Massmigration *MassmigrationSession) ValidateAndApply(senderEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Massmigration.Contract.ValidateAndApply(&_Massmigration.CallOpts, senderEncoded, encodedTx)
}

// ValidateAndApply is a free data retrieval call binding the contract method 0x5f7dde22.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Massmigration *MassmigrationCallerSession) ValidateAndApply(senderEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Massmigration.Contract.ValidateAndApply(&_Massmigration.CallOpts, senderEncoded, encodedTx)
}
