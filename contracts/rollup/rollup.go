// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollup

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

// TypesAccountInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesAccountInclusionProof struct {
	PathToAccount *big.Int
	Account       TypesUserAccount
}

// TypesAccountMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type TypesAccountMerkleProof struct {
	AccountIP TypesAccountInclusionProof
	Siblings  [][32]byte
}

// TypesAccountProofs is an auto generated low-level Go binding around an user-defined struct.
type TypesAccountProofs struct {
	From TypesAccountMerkleProof
	To   TypesAccountMerkleProof
}

// TypesBatch is an auto generated low-level Go binding around an user-defined struct.
type TypesBatch struct {
	CommitmentRoot [32]byte
	Committer      common.Address
	FinalisesOn    *big.Int
	DepositRoot    [32]byte
	Withdrawn      bool
}

// TypesBatchValidationProofs is an auto generated low-level Go binding around an user-defined struct.
type TypesBatchValidationProofs struct {
	AccountProofs []TypesAccountProofs
	PdaProof      []TypesPDAMerkleProof
}

// TypesCommitment is an auto generated low-level Go binding around an user-defined struct.
type TypesCommitment struct {
	StateRoot        [32]byte
	AccountRoot      [32]byte
	TxHashCommitment [32]byte
	BatchType        uint8
}

// TypesCommitmentInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesCommitmentInclusionProof struct {
	Commitment       TypesCommitment
	PathToCommitment *big.Int
	Siblings         [][32]byte
}

// TypesPDAInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesPDAInclusionProof struct {
	PathToPubkey *big.Int
	PubkeyLeaf   TypesPDALeaf
}

// TypesPDALeaf is an auto generated low-level Go binding around an user-defined struct.
type TypesPDALeaf struct {
	Pubkey []byte
}

// TypesPDAMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type TypesPDAMerkleProof struct {
	Pda      TypesPDAInclusionProof
	Siblings [][32]byte
}

// TypesUserAccount is an auto generated low-level Go binding around an user-defined struct.
type TypesUserAccount struct {
	ID        *big.Int
	TokenType *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	Burn      *big.Int
	LastBurn  *big.Int
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"genesisStateRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[],\"name\":\"SlashAndRollback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"}],\"name\":\"WithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountRegistry\",\"outputs\":[{\"internalType\":\"contractBLSAccountRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"finalisesOn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"depositRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"withdrawn\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositManager\",\"outputs\":[{\"internalType\":\"contractDepositManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batch_id\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHashCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"enumTypes.Usage\",\"name\":\"batchType\",\"type\":\"uint8\"}],\"internalType\":\"structTypes.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"pathToCommitment\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.CommitmentInclusionProof\",\"name\":\"commitmentMP\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountProofs[]\",\"name\":\"accountProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof[]\",\"name\":\"pdaProof\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.BatchValidationProofs\",\"name\":\"batchProofs\",\"type\":\"tuple\"}],\"name\":\"disputeBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subTreeDepth\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_zero_account_mp\",\"type\":\"tuple\"}],\"name\":\"finaliseDepositsAndSubmitBatch\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batch_id\",\"type\":\"uint256\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"finalisesOn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"depositRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"withdrawn\",\"type\":\"bool\"}],\"internalType\":\"structTypes.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestBalanceTreeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractGovernance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"invalidBatchMarker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logger\",\"outputs\":[{\"internalType\":\"contractLogger\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merkleUtils\",\"outputs\":[{\"internalType\":\"contractMerkleTreeUtils\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameRegistry\",\"outputs\":[{\"internalType\":\"contractNameRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfBatchesSubmitted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rollupReddit\",\"outputs\":[{\"internalType\":\"contractIRollupReddit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"txs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"updatedRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumTypes.Usage\",\"name\":\"batchType\",\"type\":\"uint8\"},{\"internalType\":\"uint256[][2]\",\"name\":\"aggregatedSignatures\",\"type\":\"uint256[][2]\"}],\"name\":\"submitBatch\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"contractITokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_Rollup *RollupCaller) ZEROBYTES32(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "ZERO_BYTES32")
	return *ret0, err
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_Rollup *RollupSession) ZEROBYTES32() ([32]byte, error) {
	return _Rollup.Contract.ZEROBYTES32(&_Rollup.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_Rollup *RollupCallerSession) ZEROBYTES32() ([32]byte, error) {
	return _Rollup.Contract.ZEROBYTES32(&_Rollup.CallOpts)
}

// AccountRegistry is a free data retrieval call binding the contract method 0xd089e11a.
//
// Solidity: function accountRegistry() view returns(address)
func (_Rollup *RollupCaller) AccountRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "accountRegistry")
	return *ret0, err
}

// AccountRegistry is a free data retrieval call binding the contract method 0xd089e11a.
//
// Solidity: function accountRegistry() view returns(address)
func (_Rollup *RollupSession) AccountRegistry() (common.Address, error) {
	return _Rollup.Contract.AccountRegistry(&_Rollup.CallOpts)
}

// AccountRegistry is a free data retrieval call binding the contract method 0xd089e11a.
//
// Solidity: function accountRegistry() view returns(address)
func (_Rollup *RollupCallerSession) AccountRegistry() (common.Address, error) {
	return _Rollup.Contract.AccountRegistry(&_Rollup.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32 commitmentRoot, address committer, uint256 finalisesOn, bytes32 depositRoot, bool withdrawn)
func (_Rollup *RollupCaller) Batches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CommitmentRoot [32]byte
	Committer      common.Address
	FinalisesOn    *big.Int
	DepositRoot    [32]byte
	Withdrawn      bool
}, error) {
	ret := new(struct {
		CommitmentRoot [32]byte
		Committer      common.Address
		FinalisesOn    *big.Int
		DepositRoot    [32]byte
		Withdrawn      bool
	})
	out := ret
	err := _Rollup.contract.Call(opts, out, "batches", arg0)
	return *ret, err
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32 commitmentRoot, address committer, uint256 finalisesOn, bytes32 depositRoot, bool withdrawn)
func (_Rollup *RollupSession) Batches(arg0 *big.Int) (struct {
	CommitmentRoot [32]byte
	Committer      common.Address
	FinalisesOn    *big.Int
	DepositRoot    [32]byte
	Withdrawn      bool
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32 commitmentRoot, address committer, uint256 finalisesOn, bytes32 depositRoot, bool withdrawn)
func (_Rollup *RollupCallerSession) Batches(arg0 *big.Int) (struct {
	CommitmentRoot [32]byte
	Committer      common.Address
	FinalisesOn    *big.Int
	DepositRoot    [32]byte
	Withdrawn      bool
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_Rollup *RollupCaller) DepositManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "depositManager")
	return *ret0, err
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_Rollup *RollupSession) DepositManager() (common.Address, error) {
	return _Rollup.Contract.DepositManager(&_Rollup.CallOpts)
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_Rollup *RollupCallerSession) DepositManager() (common.Address, error) {
	return _Rollup.Contract.DepositManager(&_Rollup.CallOpts)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 _batch_id) view returns((bytes32,address,uint256,bytes32,bool) batch)
func (_Rollup *RollupCaller) GetBatch(opts *bind.CallOpts, _batch_id *big.Int) (TypesBatch, error) {
	var (
		ret0 = new(TypesBatch)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "getBatch", _batch_id)
	return *ret0, err
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 _batch_id) view returns((bytes32,address,uint256,bytes32,bool) batch)
func (_Rollup *RollupSession) GetBatch(_batch_id *big.Int) (TypesBatch, error) {
	return _Rollup.Contract.GetBatch(&_Rollup.CallOpts, _batch_id)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 _batch_id) view returns((bytes32,address,uint256,bytes32,bool) batch)
func (_Rollup *RollupCallerSession) GetBatch(_batch_id *big.Int) (TypesBatch, error) {
	return _Rollup.Contract.GetBatch(&_Rollup.CallOpts, _batch_id)
}

// GetLatestBalanceTreeRoot is a free data retrieval call binding the contract method 0xb66f874a.
//
// Solidity: function getLatestBalanceTreeRoot() view returns(bytes32)
func (_Rollup *RollupCaller) GetLatestBalanceTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "getLatestBalanceTreeRoot")
	return *ret0, err
}

// GetLatestBalanceTreeRoot is a free data retrieval call binding the contract method 0xb66f874a.
//
// Solidity: function getLatestBalanceTreeRoot() view returns(bytes32)
func (_Rollup *RollupSession) GetLatestBalanceTreeRoot() ([32]byte, error) {
	return _Rollup.Contract.GetLatestBalanceTreeRoot(&_Rollup.CallOpts)
}

// GetLatestBalanceTreeRoot is a free data retrieval call binding the contract method 0xb66f874a.
//
// Solidity: function getLatestBalanceTreeRoot() view returns(bytes32)
func (_Rollup *RollupCallerSession) GetLatestBalanceTreeRoot() ([32]byte, error) {
	return _Rollup.Contract.GetLatestBalanceTreeRoot(&_Rollup.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Rollup *RollupCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Rollup *RollupSession) Governance() (common.Address, error) {
	return _Rollup.Contract.Governance(&_Rollup.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Rollup *RollupCallerSession) Governance() (common.Address, error) {
	return _Rollup.Contract.Governance(&_Rollup.CallOpts)
}

// InvalidBatchMarker is a free data retrieval call binding the contract method 0x5b097d37.
//
// Solidity: function invalidBatchMarker() view returns(uint256)
func (_Rollup *RollupCaller) InvalidBatchMarker(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "invalidBatchMarker")
	return *ret0, err
}

// InvalidBatchMarker is a free data retrieval call binding the contract method 0x5b097d37.
//
// Solidity: function invalidBatchMarker() view returns(uint256)
func (_Rollup *RollupSession) InvalidBatchMarker() (*big.Int, error) {
	return _Rollup.Contract.InvalidBatchMarker(&_Rollup.CallOpts)
}

// InvalidBatchMarker is a free data retrieval call binding the contract method 0x5b097d37.
//
// Solidity: function invalidBatchMarker() view returns(uint256)
func (_Rollup *RollupCallerSession) InvalidBatchMarker() (*big.Int, error) {
	return _Rollup.Contract.InvalidBatchMarker(&_Rollup.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Rollup *RollupCaller) Logger(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "logger")
	return *ret0, err
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Rollup *RollupSession) Logger() (common.Address, error) {
	return _Rollup.Contract.Logger(&_Rollup.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_Rollup *RollupCallerSession) Logger() (common.Address, error) {
	return _Rollup.Contract.Logger(&_Rollup.CallOpts)
}

// MerkleUtils is a free data retrieval call binding the contract method 0x47b0f08e.
//
// Solidity: function merkleUtils() view returns(address)
func (_Rollup *RollupCaller) MerkleUtils(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "merkleUtils")
	return *ret0, err
}

// MerkleUtils is a free data retrieval call binding the contract method 0x47b0f08e.
//
// Solidity: function merkleUtils() view returns(address)
func (_Rollup *RollupSession) MerkleUtils() (common.Address, error) {
	return _Rollup.Contract.MerkleUtils(&_Rollup.CallOpts)
}

// MerkleUtils is a free data retrieval call binding the contract method 0x47b0f08e.
//
// Solidity: function merkleUtils() view returns(address)
func (_Rollup *RollupCallerSession) MerkleUtils() (common.Address, error) {
	return _Rollup.Contract.MerkleUtils(&_Rollup.CallOpts)
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Rollup *RollupCaller) NameRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "nameRegistry")
	return *ret0, err
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Rollup *RollupSession) NameRegistry() (common.Address, error) {
	return _Rollup.Contract.NameRegistry(&_Rollup.CallOpts)
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Rollup *RollupCallerSession) NameRegistry() (common.Address, error) {
	return _Rollup.Contract.NameRegistry(&_Rollup.CallOpts)
}

// NumOfBatchesSubmitted is a free data retrieval call binding the contract method 0x8267b96c.
//
// Solidity: function numOfBatchesSubmitted() view returns(uint256)
func (_Rollup *RollupCaller) NumOfBatchesSubmitted(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "numOfBatchesSubmitted")
	return *ret0, err
}

// NumOfBatchesSubmitted is a free data retrieval call binding the contract method 0x8267b96c.
//
// Solidity: function numOfBatchesSubmitted() view returns(uint256)
func (_Rollup *RollupSession) NumOfBatchesSubmitted() (*big.Int, error) {
	return _Rollup.Contract.NumOfBatchesSubmitted(&_Rollup.CallOpts)
}

// NumOfBatchesSubmitted is a free data retrieval call binding the contract method 0x8267b96c.
//
// Solidity: function numOfBatchesSubmitted() view returns(uint256)
func (_Rollup *RollupCallerSession) NumOfBatchesSubmitted() (*big.Int, error) {
	return _Rollup.Contract.NumOfBatchesSubmitted(&_Rollup.CallOpts)
}

// RollupReddit is a free data retrieval call binding the contract method 0x49432a62.
//
// Solidity: function rollupReddit() view returns(address)
func (_Rollup *RollupCaller) RollupReddit(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "rollupReddit")
	return *ret0, err
}

// RollupReddit is a free data retrieval call binding the contract method 0x49432a62.
//
// Solidity: function rollupReddit() view returns(address)
func (_Rollup *RollupSession) RollupReddit() (common.Address, error) {
	return _Rollup.Contract.RollupReddit(&_Rollup.CallOpts)
}

// RollupReddit is a free data retrieval call binding the contract method 0x49432a62.
//
// Solidity: function rollupReddit() view returns(address)
func (_Rollup *RollupCallerSession) RollupReddit() (common.Address, error) {
	return _Rollup.Contract.RollupReddit(&_Rollup.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Rollup *RollupCaller) TokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "tokenRegistry")
	return *ret0, err
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Rollup *RollupSession) TokenRegistry() (common.Address, error) {
	return _Rollup.Contract.TokenRegistry(&_Rollup.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Rollup *RollupCallerSession) TokenRegistry() (common.Address, error) {
	return _Rollup.Contract.TokenRegistry(&_Rollup.CallOpts)
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0xdf070983.
//
// Solidity: function SlashAndRollback() returns()
func (_Rollup *RollupTransactor) SlashAndRollback(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "SlashAndRollback")
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0xdf070983.
//
// Solidity: function SlashAndRollback() returns()
func (_Rollup *RollupSession) SlashAndRollback() (*types.Transaction, error) {
	return _Rollup.Contract.SlashAndRollback(&_Rollup.TransactOpts)
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0xdf070983.
//
// Solidity: function SlashAndRollback() returns()
func (_Rollup *RollupTransactorSession) SlashAndRollback() (*types.Transaction, error) {
	return _Rollup.Contract.SlashAndRollback(&_Rollup.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xff34585d.
//
// Solidity: function WithdrawStake(uint256 batch_id) returns()
func (_Rollup *RollupTransactor) WithdrawStake(opts *bind.TransactOpts, batch_id *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "WithdrawStake", batch_id)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xff34585d.
//
// Solidity: function WithdrawStake(uint256 batch_id) returns()
func (_Rollup *RollupSession) WithdrawStake(batch_id *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStake(&_Rollup.TransactOpts, batch_id)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xff34585d.
//
// Solidity: function WithdrawStake(uint256 batch_id) returns()
func (_Rollup *RollupTransactorSession) WithdrawStake(batch_id *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStake(&_Rollup.TransactOpts, batch_id)
}

// DisputeBatch is a paid mutator transaction binding the contract method 0x413255d4.
//
// Solidity: function disputeBatch(uint256 _batch_id, ((bytes32,bytes32,bytes32,uint8),uint256,bytes32[]) commitmentMP, bytes txs, ((((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]),((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]))[],((uint256,(bytes)),bytes32[])[]) batchProofs) returns()
func (_Rollup *RollupTransactor) DisputeBatch(opts *bind.TransactOpts, _batch_id *big.Int, commitmentMP TypesCommitmentInclusionProof, txs []byte, batchProofs TypesBatchValidationProofs) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "disputeBatch", _batch_id, commitmentMP, txs, batchProofs)
}

// DisputeBatch is a paid mutator transaction binding the contract method 0x413255d4.
//
// Solidity: function disputeBatch(uint256 _batch_id, ((bytes32,bytes32,bytes32,uint8),uint256,bytes32[]) commitmentMP, bytes txs, ((((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]),((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]))[],((uint256,(bytes)),bytes32[])[]) batchProofs) returns()
func (_Rollup *RollupSession) DisputeBatch(_batch_id *big.Int, commitmentMP TypesCommitmentInclusionProof, txs []byte, batchProofs TypesBatchValidationProofs) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeBatch(&_Rollup.TransactOpts, _batch_id, commitmentMP, txs, batchProofs)
}

// DisputeBatch is a paid mutator transaction binding the contract method 0x413255d4.
//
// Solidity: function disputeBatch(uint256 _batch_id, ((bytes32,bytes32,bytes32,uint8),uint256,bytes32[]) commitmentMP, bytes txs, ((((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]),((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]))[],((uint256,(bytes)),bytes32[])[]) batchProofs) returns()
func (_Rollup *RollupTransactorSession) DisputeBatch(_batch_id *big.Int, commitmentMP TypesCommitmentInclusionProof, txs []byte, batchProofs TypesBatchValidationProofs) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeBatch(&_Rollup.TransactOpts, _batch_id, commitmentMP, txs, batchProofs)
}

// FinaliseDepositsAndSubmitBatch is a paid mutator transaction binding the contract method 0xcac30223.
//
// Solidity: function finaliseDepositsAndSubmitBatch(uint256 _subTreeDepth, ((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]) _zero_account_mp) payable returns()
func (_Rollup *RollupTransactor) FinaliseDepositsAndSubmitBatch(opts *bind.TransactOpts, _subTreeDepth *big.Int, _zero_account_mp TypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "finaliseDepositsAndSubmitBatch", _subTreeDepth, _zero_account_mp)
}

// FinaliseDepositsAndSubmitBatch is a paid mutator transaction binding the contract method 0xcac30223.
//
// Solidity: function finaliseDepositsAndSubmitBatch(uint256 _subTreeDepth, ((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]) _zero_account_mp) payable returns()
func (_Rollup *RollupSession) FinaliseDepositsAndSubmitBatch(_subTreeDepth *big.Int, _zero_account_mp TypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.FinaliseDepositsAndSubmitBatch(&_Rollup.TransactOpts, _subTreeDepth, _zero_account_mp)
}

// FinaliseDepositsAndSubmitBatch is a paid mutator transaction binding the contract method 0xcac30223.
//
// Solidity: function finaliseDepositsAndSubmitBatch(uint256 _subTreeDepth, ((uint256,(uint256,uint256,uint256,uint256,uint256,uint256)),bytes32[]) _zero_account_mp) payable returns()
func (_Rollup *RollupTransactorSession) FinaliseDepositsAndSubmitBatch(_subTreeDepth *big.Int, _zero_account_mp TypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.FinaliseDepositsAndSubmitBatch(&_Rollup.TransactOpts, _subTreeDepth, _zero_account_mp)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x75e655a6.
//
// Solidity: function submitBatch(bytes[] txs, bytes32[] updatedRoots, uint8 batchType, uint256[][2] aggregatedSignatures) payable returns()
func (_Rollup *RollupTransactor) SubmitBatch(opts *bind.TransactOpts, txs [][]byte, updatedRoots [][32]byte, batchType uint8, aggregatedSignatures [][2]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitBatch", txs, updatedRoots, batchType, aggregatedSignatures)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x75e655a6.
//
// Solidity: function submitBatch(bytes[] txs, bytes32[] updatedRoots, uint8 batchType, uint256[][2] aggregatedSignatures) payable returns()
func (_Rollup *RollupSession) SubmitBatch(txs [][]byte, updatedRoots [][32]byte, batchType uint8, aggregatedSignatures [][2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitBatch(&_Rollup.TransactOpts, txs, updatedRoots, batchType, aggregatedSignatures)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x75e655a6.
//
// Solidity: function submitBatch(bytes[] txs, bytes32[] updatedRoots, uint8 batchType, uint256[][2] aggregatedSignatures) payable returns()
func (_Rollup *RollupTransactorSession) SubmitBatch(txs [][]byte, updatedRoots [][32]byte, batchType uint8, aggregatedSignatures [][2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitBatch(&_Rollup.TransactOpts, txs, updatedRoots, batchType, aggregatedSignatures)
}
