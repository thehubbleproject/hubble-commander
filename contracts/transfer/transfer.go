// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transfer

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

// OffchainTransfer is an auto generated low-level Go binding around an user-defined struct.
type OffchainTransfer struct {
	TxType    *big.Int
	FromIndex *big.Int
	ToIndex   *big.Int
	Amount    *big.Int
	Fee       *big.Int
	Nonce     *big.Int
}

// TxTransfer is an auto generated low-level Go binding around an user-defined struct.
type TxTransfer struct {
	FromIndex *big.Int
	ToIndex   *big.Int
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
	PubkeyIndex *big.Int
	TokenType   *big.Int
	Balance     *big.Int
	Nonce       *big.Int
}

// TransferABI is the input ABI used to generate the binding from.
const TransferABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnesses\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"checkSignature\",\"outputs\":[{\"internalType\":\"enumTypes.Result\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"encodedTxs\",\"type\":\"bytes[]\"}],\"name\":\"compress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"decode\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"decompress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structTx.Transfer[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOffchain.Transfer\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"encode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"name\":\"process\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"signBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"validate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"senderEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiverEncoded\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"validateAndApply\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newSender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newReceiver\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.Result\",\"name\":\"result\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txMsg\",\"type\":\"bytes\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32\",\"name\":\"domain\",\"type\":\"bytes32\"}],\"name\":\"verifySingle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Transfer is an auto generated Go binding around an Ethereum contract.
type Transfer struct {
	TransferCaller     // Read-only binding to the contract
	TransferTransactor // Write-only binding to the contract
	TransferFilterer   // Log filterer for contract events
}

// TransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferSession struct {
	Contract     *Transfer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferCallerSession struct {
	Contract *TransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferTransactorSession struct {
	Contract     *TransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferRaw struct {
	Contract *Transfer // Generic contract binding to access the raw methods on
}

// TransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferCallerRaw struct {
	Contract *TransferCaller // Generic read-only contract binding to access the raw methods on
}

// TransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferTransactorRaw struct {
	Contract *TransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfer creates a new instance of Transfer, bound to a specific deployed contract.
func NewTransfer(address common.Address, backend bind.ContractBackend) (*Transfer, error) {
	contract, err := bindTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfer{TransferCaller: TransferCaller{contract: contract}, TransferTransactor: TransferTransactor{contract: contract}, TransferFilterer: TransferFilterer{contract: contract}}, nil
}

// NewTransferCaller creates a new read-only instance of Transfer, bound to a specific deployed contract.
func NewTransferCaller(address common.Address, caller bind.ContractCaller) (*TransferCaller, error) {
	contract, err := bindTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferCaller{contract: contract}, nil
}

// NewTransferTransactor creates a new write-only instance of Transfer, bound to a specific deployed contract.
func NewTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferTransactor, error) {
	contract, err := bindTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferTransactor{contract: contract}, nil
}

// NewTransferFilterer creates a new log filterer instance of Transfer, bound to a specific deployed contract.
func NewTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferFilterer, error) {
	contract, err := bindTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferFilterer{contract: contract}, nil
}

// bindTransfer binds a generic wrapper to an already deployed contract.
func bindTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.TransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.TransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfer *TransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfer *TransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfer *TransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfer.Contract.contract.Transact(opts, method, params...)
}

// CheckSignature is a free data retrieval call binding the contract method 0x729ba7f0.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Transfer *TransferCaller) CheckSignature(opts *bind.CallOpts, signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "checkSignature", signature, proof, stateRoot, accountRoot, domain, txs)
	return *ret0, err
}

// CheckSignature is a free data retrieval call binding the contract method 0x729ba7f0.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Transfer *TransferSession) CheckSignature(signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	return _Transfer.Contract.CheckSignature(&_Transfer.CallOpts, signature, proof, stateRoot, accountRoot, domain, txs)
}

// CheckSignature is a free data retrieval call binding the contract method 0x729ba7f0.
//
// Solidity: function checkSignature(uint256[2] signature, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) proof, bytes32 stateRoot, bytes32 accountRoot, bytes32 domain, bytes txs) view returns(uint8)
func (_Transfer *TransferCallerSession) CheckSignature(signature [2]*big.Int, proof TypesSignatureProof, stateRoot [32]byte, accountRoot [32]byte, domain [32]byte, txs []byte) (uint8, error) {
	return _Transfer.Contract.CheckSignature(&_Transfer.CallOpts, signature, proof, stateRoot, accountRoot, domain, txs)
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Transfer *TransferCaller) Compress(opts *bind.CallOpts, encodedTxs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "compress", encodedTxs)
	return *ret0, err
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Transfer *TransferSession) Compress(encodedTxs [][]byte) ([]byte, error) {
	return _Transfer.Contract.Compress(&_Transfer.CallOpts, encodedTxs)
}

// Compress is a free data retrieval call binding the contract method 0xfa320471.
//
// Solidity: function compress(bytes[] encodedTxs) pure returns(bytes)
func (_Transfer *TransferCallerSession) Compress(encodedTxs [][]byte) ([]byte, error) {
	return _Transfer.Contract.Compress(&_Transfer.CallOpts, encodedTxs)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Transfer *TransferCaller) Decode(opts *bind.CallOpts, encodedTx []byte) (OffchainTransfer, error) {
	var (
		ret0 = new(OffchainTransfer)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "decode", encodedTx)
	return *ret0, err
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Transfer *TransferSession) Decode(encodedTx []byte) (OffchainTransfer, error) {
	return _Transfer.Contract.Decode(&_Transfer.CallOpts, encodedTx)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes encodedTx) pure returns((uint256,uint256,uint256,uint256,uint256,uint256) _tx)
func (_Transfer *TransferCallerSession) Decode(encodedTx []byte) (OffchainTransfer, error) {
	return _Transfer.Contract.Decode(&_Transfer.CallOpts, encodedTx)
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256,uint256)[])
func (_Transfer *TransferCaller) Decompress(opts *bind.CallOpts, txs []byte) ([]TxTransfer, error) {
	var (
		ret0 = new([]TxTransfer)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "decompress", txs)
	return *ret0, err
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256,uint256)[])
func (_Transfer *TransferSession) Decompress(txs []byte) ([]TxTransfer, error) {
	return _Transfer.Contract.Decompress(&_Transfer.CallOpts, txs)
}

// Decompress is a free data retrieval call binding the contract method 0x5cd3f3a1.
//
// Solidity: function decompress(bytes txs) pure returns((uint256,uint256,uint256,uint256)[])
func (_Transfer *TransferCallerSession) Decompress(txs []byte) ([]TxTransfer, error) {
	return _Transfer.Contract.Decompress(&_Transfer.CallOpts, txs)
}

// Encode is a free data retrieval call binding the contract method 0x08909a83.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Transfer *TransferCaller) Encode(opts *bind.CallOpts, _tx OffchainTransfer) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "encode", _tx)
	return *ret0, err
}

// Encode is a free data retrieval call binding the contract method 0x08909a83.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Transfer *TransferSession) Encode(_tx OffchainTransfer) ([]byte, error) {
	return _Transfer.Contract.Encode(&_Transfer.CallOpts, _tx)
}

// Encode is a free data retrieval call binding the contract method 0x08909a83.
//
// Solidity: function encode((uint256,uint256,uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Transfer *TransferCallerSession) Encode(_tx OffchainTransfer) ([]byte, error) {
	return _Transfer.Contract.Encode(&_Transfer.CallOpts, _tx)
}

// Process is a free data retrieval call binding the contract method 0x89840761.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Transfer *TransferCaller) Process(opts *bind.CallOpts, stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	ret := new(struct {
		NewRoot [32]byte
		Result  uint8
	})
	out := ret
	err := _Transfer.contract.Call(opts, out, "process", stateRoot, encodedTx, tokenType, from, to)
	return *ret, err
}

// Process is a free data retrieval call binding the contract method 0x89840761.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Transfer *TransferSession) Process(stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	return _Transfer.Contract.Process(&_Transfer.CallOpts, stateRoot, encodedTx, tokenType, from, to)
}

// Process is a free data retrieval call binding the contract method 0x89840761.
//
// Solidity: function process(bytes32 stateRoot, bytes encodedTx, uint256 tokenType, ((uint256,uint256,uint256,uint256),bytes32[]) from, ((uint256,uint256,uint256,uint256),bytes32[]) to) pure returns(bytes32 newRoot, uint8 result)
func (_Transfer *TransferCallerSession) Process(stateRoot [32]byte, encodedTx []byte, tokenType *big.Int, from TypesStateMerkleProof, to TypesStateMerkleProof) (struct {
	NewRoot [32]byte
	Result  uint8
}, error) {
	return _Transfer.Contract.Process(&_Transfer.CallOpts, stateRoot, encodedTx, tokenType, from, to)
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTx) pure returns(bytes)
func (_Transfer *TransferCaller) SignBytes(opts *bind.CallOpts, encodedTx []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Transfer.contract.Call(opts, out, "signBytes", encodedTx)
	return *ret0, err
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTx) pure returns(bytes)
func (_Transfer *TransferSession) SignBytes(encodedTx []byte) ([]byte, error) {
	return _Transfer.Contract.SignBytes(&_Transfer.CallOpts, encodedTx)
}

// SignBytes is a free data retrieval call binding the contract method 0xd8b1728e.
//
// Solidity: function signBytes(bytes encodedTx) pure returns(bytes)
func (_Transfer *TransferCallerSession) SignBytes(encodedTx []byte) ([]byte, error) {
	return _Transfer.Contract.SignBytes(&_Transfer.CallOpts, encodedTx)
}

// Validate is a free data retrieval call binding the contract method 0x5ef63f4f.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Transfer *TransferCaller) Validate(opts *bind.CallOpts, encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Transfer.contract.Call(opts, out, "validate", encodedTx, signature, pubkey, domain)
	return err
}

// Validate is a free data retrieval call binding the contract method 0x5ef63f4f.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Transfer *TransferSession) Validate(encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Transfer.Contract.Validate(&_Transfer.CallOpts, encodedTx, signature, pubkey, domain)
}

// Validate is a free data retrieval call binding the contract method 0x5ef63f4f.
//
// Solidity: function validate(bytes encodedTx, uint256[2] signature, uint256[4] pubkey, bytes32 domain) view returns()
func (_Transfer *TransferCallerSession) Validate(encodedTx []byte, signature [2]*big.Int, pubkey [4]*big.Int, domain [32]byte) error {
	return _Transfer.Contract.Validate(&_Transfer.CallOpts, encodedTx, signature, pubkey, domain)
}

// ValidateAndApply is a free data retrieval call binding the contract method 0xdd21c2ed.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes receiverEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Transfer *TransferCaller) ValidateAndApply(opts *bind.CallOpts, senderEncoded []byte, receiverEncoded []byte, encodedTx []byte) (struct {
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
	err := _Transfer.contract.Call(opts, out, "validateAndApply", senderEncoded, receiverEncoded, encodedTx)
	return *ret, err
}

// ValidateAndApply is a free data retrieval call binding the contract method 0xdd21c2ed.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes receiverEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Transfer *TransferSession) ValidateAndApply(senderEncoded []byte, receiverEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Transfer.Contract.ValidateAndApply(&_Transfer.CallOpts, senderEncoded, receiverEncoded, encodedTx)
}

// ValidateAndApply is a free data retrieval call binding the contract method 0xdd21c2ed.
//
// Solidity: function validateAndApply(bytes senderEncoded, bytes receiverEncoded, bytes encodedTx) pure returns(bytes newSender, bytes newReceiver, uint8 result)
func (_Transfer *TransferCallerSession) ValidateAndApply(senderEncoded []byte, receiverEncoded []byte, encodedTx []byte) (struct {
	NewSender   []byte
	NewReceiver []byte
	Result      uint8
}, error) {
	return _Transfer.Contract.ValidateAndApply(&_Transfer.CallOpts, senderEncoded, receiverEncoded, encodedTx)
}

// VerifySingle is a free data retrieval call binding the contract method 0x494ac1dd.
//
// Solidity: function verifySingle(bytes txMsg, uint256[4] pubkey, uint256[2] signature, bytes32 domain) view returns()
func (_Transfer *TransferCaller) VerifySingle(opts *bind.CallOpts, txMsg []byte, pubkey [4]*big.Int, signature [2]*big.Int, domain [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _Transfer.contract.Call(opts, out, "verifySingle", txMsg, pubkey, signature, domain)
	return err
}

// VerifySingle is a free data retrieval call binding the contract method 0x494ac1dd.
//
// Solidity: function verifySingle(bytes txMsg, uint256[4] pubkey, uint256[2] signature, bytes32 domain) view returns()
func (_Transfer *TransferSession) VerifySingle(txMsg []byte, pubkey [4]*big.Int, signature [2]*big.Int, domain [32]byte) error {
	return _Transfer.Contract.VerifySingle(&_Transfer.CallOpts, txMsg, pubkey, signature, domain)
}

// VerifySingle is a free data retrieval call binding the contract method 0x494ac1dd.
//
// Solidity: function verifySingle(bytes txMsg, uint256[4] pubkey, uint256[2] signature, bytes32 domain) view returns()
func (_Transfer *TransferCallerSession) VerifySingle(txMsg []byte, pubkey [4]*big.Int, signature [2]*big.Int, domain [32]byte) error {
	return _Transfer.Contract.VerifySingle(&_Transfer.CallOpts, txMsg, pubkey, signature, domain)
}
