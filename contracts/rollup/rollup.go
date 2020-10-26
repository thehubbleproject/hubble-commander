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

// TypesBatch is an auto generated low-level Go binding around an user-defined struct.
type TypesBatch struct {
	CommitmentRoot [32]byte
	Meta           [32]byte
}

// TypesCommitment is an auto generated low-level Go binding around an user-defined struct.
type TypesCommitment struct {
	StateRoot [32]byte
	BodyRoot  [32]byte
}

// TypesCommitmentInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesCommitmentInclusionProof struct {
	Commitment TypesCommitment
	Path       *big.Int
	Witness    [][32]byte
}

// TypesMMCommitmentInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesMMCommitmentInclusionProof struct {
	Commitment TypesMassMigrationCommitment
	Path       *big.Int
	Witness    [][32]byte
}

// TypesMassMigrationBody is an auto generated low-level Go binding around an user-defined struct.
type TypesMassMigrationBody struct {
	AccountRoot  [32]byte
	Signature    [2]*big.Int
	SpokeID      *big.Int
	WithdrawRoot [32]byte
	TokenID      *big.Int
	Amount       *big.Int
	FeeReceiver  *big.Int
	Txs          []byte
}

// TypesMassMigrationCommitment is an auto generated low-level Go binding around an user-defined struct.
type TypesMassMigrationCommitment struct {
	StateRoot [32]byte
	Body      TypesMassMigrationBody
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

// TypesSubtreeVacancyProof is an auto generated low-level Go binding around an user-defined struct.
type TypesSubtreeVacancyProof struct {
	Depth       *big.Int
	PathAtDepth *big.Int
	Witness     [][32]byte
}

// TypesTransferBody is an auto generated low-level Go binding around an user-defined struct.
type TypesTransferBody struct {
	AccountRoot [32]byte
	Signature   [2]*big.Int
	FeeReceiver *big.Int
	Txs         []byte
}

// TypesTransferCommitment is an auto generated low-level Go binding around an user-defined struct.
type TypesTransferCommitment struct {
	StateRoot [32]byte
	Body      TypesTransferBody
}

// TypesTransferCommitmentInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesTransferCommitmentInclusionProof struct {
	Commitment TypesTransferCommitment
	Path       *big.Int
	Witness    [][32]byte
}

// TypesUserState is an auto generated low-level Go binding around an user-defined struct.
type TypesUserState struct {
	PubkeyIndex *big.Int
	TokenType   *big.Int
	Balance     *big.Int
	Nonce       *big.Int
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"genesisStateRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountRegistry\",\"outputs\":[{\"internalType\":\"contractBLSAccountRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"meta\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositManager\",\"outputs\":[{\"internalType\":\"contractDepositManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.MassMigrationBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.MassMigrationCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.MMCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnesses\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProof\",\"name\":\"signatureProof\",\"type\":\"tuple\"}],\"name\":\"disputeSignatureMassMigration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.TransferBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransferCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.TransferCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnesses\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProof\",\"name\":\"signatureProof\",\"type\":\"tuple\"}],\"name\":\"disputeSignatureTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.CommitmentInclusionProof\",\"name\":\"previous\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.MassMigrationBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.MassMigrationCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.MMCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"}],\"name\":\"disputeTransitionMassMigration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.CommitmentInclusionProof\",\"name\":\"previous\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.TransferBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransferCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.TransferCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"}],\"name\":\"disputeTransitionTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"meta\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"govStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractGovernance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"invalidBatchMarker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logger\",\"outputs\":[{\"internalType\":\"contractLogger\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"massMigration\",\"outputs\":[{\"internalType\":\"contractMassMigration\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameRegistry\",\"outputs\":[{\"internalType\":\"contractNameRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfBatchesSubmitted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"slashAndRollback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"stateRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"signatures\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeReceivers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"txss\",\"type\":\"bytes[]\"}],\"name\":\"submitCreate2Transfer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.CommitmentInclusionProof\",\"name\":\"previous\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pathAtDepth\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.SubtreeVacancyProof\",\"name\":\"vacant\",\"type\":\"tuple\"}],\"name\":\"submitDeposits\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"stateRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"signatures\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"meta\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"withdrawRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"txss\",\"type\":\"bytes[]\"}],\"name\":\"submitMassMigration\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"stateRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"signatures\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeReceivers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"txss\",\"type\":\"bytes[]\"}],\"name\":\"submitTransfer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"contractTransfer\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// AppID is a free data retrieval call binding the contract method 0xf9a6cbae.
//
// Solidity: function appID() view returns(bytes32)
func (_Rollup *RollupCaller) AppID(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "appID")
	return *ret0, err
}

// AppID is a free data retrieval call binding the contract method 0xf9a6cbae.
//
// Solidity: function appID() view returns(bytes32)
func (_Rollup *RollupSession) AppID() ([32]byte, error) {
	return _Rollup.Contract.AppID(&_Rollup.CallOpts)
}

// AppID is a free data retrieval call binding the contract method 0xf9a6cbae.
//
// Solidity: function appID() view returns(bytes32)
func (_Rollup *RollupCallerSession) AppID() ([32]byte, error) {
	return _Rollup.Contract.AppID(&_Rollup.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32 commitmentRoot, bytes32 meta)
func (_Rollup *RollupCaller) Batches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CommitmentRoot [32]byte
	Meta           [32]byte
}, error) {
	ret := new(struct {
		CommitmentRoot [32]byte
		Meta           [32]byte
	})
	out := ret
	err := _Rollup.contract.Call(opts, out, "batches", arg0)
	return *ret, err
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32 commitmentRoot, bytes32 meta)
func (_Rollup *RollupSession) Batches(arg0 *big.Int) (struct {
	CommitmentRoot [32]byte
	Meta           [32]byte
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32 commitmentRoot, bytes32 meta)
func (_Rollup *RollupCallerSession) Batches(arg0 *big.Int) (struct {
	CommitmentRoot [32]byte
	Meta           [32]byte
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
// Solidity: function getBatch(uint256 batchID) view returns((bytes32,bytes32) batch)
func (_Rollup *RollupCaller) GetBatch(opts *bind.CallOpts, batchID *big.Int) (TypesBatch, error) {
	var (
		ret0 = new(TypesBatch)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "getBatch", batchID)
	return *ret0, err
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 batchID) view returns((bytes32,bytes32) batch)
func (_Rollup *RollupSession) GetBatch(batchID *big.Int) (TypesBatch, error) {
	return _Rollup.Contract.GetBatch(&_Rollup.CallOpts, batchID)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 batchID) view returns((bytes32,bytes32) batch)
func (_Rollup *RollupCallerSession) GetBatch(batchID *big.Int) (TypesBatch, error) {
	return _Rollup.Contract.GetBatch(&_Rollup.CallOpts, batchID)
}

// GovStakeAmount is a free data retrieval call binding the contract method 0x0c3f72eb.
//
// Solidity: function govStakeAmount() view returns(uint256)
func (_Rollup *RollupCaller) GovStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "govStakeAmount")
	return *ret0, err
}

// GovStakeAmount is a free data retrieval call binding the contract method 0x0c3f72eb.
//
// Solidity: function govStakeAmount() view returns(uint256)
func (_Rollup *RollupSession) GovStakeAmount() (*big.Int, error) {
	return _Rollup.Contract.GovStakeAmount(&_Rollup.CallOpts)
}

// GovStakeAmount is a free data retrieval call binding the contract method 0x0c3f72eb.
//
// Solidity: function govStakeAmount() view returns(uint256)
func (_Rollup *RollupCallerSession) GovStakeAmount() (*big.Int, error) {
	return _Rollup.Contract.GovStakeAmount(&_Rollup.CallOpts)
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

// MassMigration is a free data retrieval call binding the contract method 0x0ed75b9c.
//
// Solidity: function massMigration() view returns(address)
func (_Rollup *RollupCaller) MassMigration(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "massMigration")
	return *ret0, err
}

// MassMigration is a free data retrieval call binding the contract method 0x0ed75b9c.
//
// Solidity: function massMigration() view returns(address)
func (_Rollup *RollupSession) MassMigration() (common.Address, error) {
	return _Rollup.Contract.MassMigration(&_Rollup.CallOpts)
}

// MassMigration is a free data retrieval call binding the contract method 0x0ed75b9c.
//
// Solidity: function massMigration() view returns(address)
func (_Rollup *RollupCallerSession) MassMigration() (common.Address, error) {
	return _Rollup.Contract.MassMigration(&_Rollup.CallOpts)
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

// Transfer is a free data retrieval call binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() view returns(address)
func (_Rollup *RollupCaller) Transfer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "transfer")
	return *ret0, err
}

// Transfer is a free data retrieval call binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() view returns(address)
func (_Rollup *RollupSession) Transfer() (common.Address, error) {
	return _Rollup.Contract.Transfer(&_Rollup.CallOpts)
}

// Transfer is a free data retrieval call binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() view returns(address)
func (_Rollup *RollupCallerSession) Transfer() (common.Address, error) {
	return _Rollup.Contract.Transfer(&_Rollup.CallOpts)
}

// DisputeSignatureMassMigration is a paid mutator transaction binding the contract method 0x2538507d.
//
// Solidity: function disputeSignatureMassMigration(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupTransactor) DisputeSignatureMassMigration(opts *bind.TransactOpts, batchID *big.Int, target TypesMMCommitmentInclusionProof, signatureProof TypesSignatureProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "disputeSignatureMassMigration", batchID, target, signatureProof)
}

// DisputeSignatureMassMigration is a paid mutator transaction binding the contract method 0x2538507d.
//
// Solidity: function disputeSignatureMassMigration(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupSession) DisputeSignatureMassMigration(batchID *big.Int, target TypesMMCommitmentInclusionProof, signatureProof TypesSignatureProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeSignatureMassMigration(&_Rollup.TransactOpts, batchID, target, signatureProof)
}

// DisputeSignatureMassMigration is a paid mutator transaction binding the contract method 0x2538507d.
//
// Solidity: function disputeSignatureMassMigration(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupTransactorSession) DisputeSignatureMassMigration(batchID *big.Int, target TypesMMCommitmentInclusionProof, signatureProof TypesSignatureProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeSignatureMassMigration(&_Rollup.TransactOpts, batchID, target, signatureProof)
}

// DisputeSignatureTransfer is a paid mutator transaction binding the contract method 0xc08caaef.
//
// Solidity: function disputeSignatureTransfer(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupTransactor) DisputeSignatureTransfer(opts *bind.TransactOpts, batchID *big.Int, target TypesTransferCommitmentInclusionProof, signatureProof TypesSignatureProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "disputeSignatureTransfer", batchID, target, signatureProof)
}

// DisputeSignatureTransfer is a paid mutator transaction binding the contract method 0xc08caaef.
//
// Solidity: function disputeSignatureTransfer(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupSession) DisputeSignatureTransfer(batchID *big.Int, target TypesTransferCommitmentInclusionProof, signatureProof TypesSignatureProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeSignatureTransfer(&_Rollup.TransactOpts, batchID, target, signatureProof)
}

// DisputeSignatureTransfer is a paid mutator transaction binding the contract method 0xc08caaef.
//
// Solidity: function disputeSignatureTransfer(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupTransactorSession) DisputeSignatureTransfer(batchID *big.Int, target TypesTransferCommitmentInclusionProof, signatureProof TypesSignatureProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeSignatureTransfer(&_Rollup.TransactOpts, batchID, target, signatureProof)
}

// DisputeTransitionMassMigration is a paid mutator transaction binding the contract method 0xeb84e076.
//
// Solidity: function disputeTransitionMassMigration(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupTransactor) DisputeTransitionMassMigration(opts *bind.TransactOpts, batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesMMCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "disputeTransitionMassMigration", batchID, previous, target, proofs)
}

// DisputeTransitionMassMigration is a paid mutator transaction binding the contract method 0xeb84e076.
//
// Solidity: function disputeTransitionMassMigration(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupSession) DisputeTransitionMassMigration(batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesMMCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeTransitionMassMigration(&_Rollup.TransactOpts, batchID, previous, target, proofs)
}

// DisputeTransitionMassMigration is a paid mutator transaction binding the contract method 0xeb84e076.
//
// Solidity: function disputeTransitionMassMigration(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes32,uint256,uint256,uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupTransactorSession) DisputeTransitionMassMigration(batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesMMCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeTransitionMassMigration(&_Rollup.TransactOpts, batchID, previous, target, proofs)
}

// DisputeTransitionTransfer is a paid mutator transaction binding the contract method 0x4e23e8c3.
//
// Solidity: function disputeTransitionTransfer(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupTransactor) DisputeTransitionTransfer(opts *bind.TransactOpts, batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesTransferCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "disputeTransitionTransfer", batchID, previous, target, proofs)
}

// DisputeTransitionTransfer is a paid mutator transaction binding the contract method 0x4e23e8c3.
//
// Solidity: function disputeTransitionTransfer(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupSession) DisputeTransitionTransfer(batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesTransferCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeTransitionTransfer(&_Rollup.TransactOpts, batchID, previous, target, proofs)
}

// DisputeTransitionTransfer is a paid mutator transaction binding the contract method 0x4e23e8c3.
//
// Solidity: function disputeTransitionTransfer(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupTransactorSession) DisputeTransitionTransfer(batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesTransferCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeTransitionTransfer(&_Rollup.TransactOpts, batchID, previous, target, proofs)
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0x010615fc.
//
// Solidity: function slashAndRollback() returns()
func (_Rollup *RollupTransactor) SlashAndRollback(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "slashAndRollback")
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0x010615fc.
//
// Solidity: function slashAndRollback() returns()
func (_Rollup *RollupSession) SlashAndRollback() (*types.Transaction, error) {
	return _Rollup.Contract.SlashAndRollback(&_Rollup.TransactOpts)
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0x010615fc.
//
// Solidity: function slashAndRollback() returns()
func (_Rollup *RollupTransactorSession) SlashAndRollback() (*types.Transaction, error) {
	return _Rollup.Contract.SlashAndRollback(&_Rollup.TransactOpts)
}

// SubmitCreate2Transfer is a paid mutator transaction binding the contract method 0x9251597f.
//
// Solidity: function submitCreate2Transfer(bytes32[] stateRoots, uint256[2][] signatures, uint256[] feeReceivers, bytes[] txss) payable returns()
func (_Rollup *RollupTransactor) SubmitCreate2Transfer(opts *bind.TransactOpts, stateRoots [][32]byte, signatures [][2]*big.Int, feeReceivers []*big.Int, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitCreate2Transfer", stateRoots, signatures, feeReceivers, txss)
}

// SubmitCreate2Transfer is a paid mutator transaction binding the contract method 0x9251597f.
//
// Solidity: function submitCreate2Transfer(bytes32[] stateRoots, uint256[2][] signatures, uint256[] feeReceivers, bytes[] txss) payable returns()
func (_Rollup *RollupSession) SubmitCreate2Transfer(stateRoots [][32]byte, signatures [][2]*big.Int, feeReceivers []*big.Int, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitCreate2Transfer(&_Rollup.TransactOpts, stateRoots, signatures, feeReceivers, txss)
}

// SubmitCreate2Transfer is a paid mutator transaction binding the contract method 0x9251597f.
//
// Solidity: function submitCreate2Transfer(bytes32[] stateRoots, uint256[2][] signatures, uint256[] feeReceivers, bytes[] txss) payable returns()
func (_Rollup *RollupTransactorSession) SubmitCreate2Transfer(stateRoots [][32]byte, signatures [][2]*big.Int, feeReceivers []*big.Int, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitCreate2Transfer(&_Rollup.TransactOpts, stateRoots, signatures, feeReceivers, txss)
}

// SubmitDeposits is a paid mutator transaction binding the contract method 0x562a2bca.
//
// Solidity: function submitDeposits(((bytes32,bytes32),uint256,bytes32[]) previous, (uint256,uint256,bytes32[]) vacant) payable returns()
func (_Rollup *RollupTransactor) SubmitDeposits(opts *bind.TransactOpts, previous TypesCommitmentInclusionProof, vacant TypesSubtreeVacancyProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitDeposits", previous, vacant)
}

// SubmitDeposits is a paid mutator transaction binding the contract method 0x562a2bca.
//
// Solidity: function submitDeposits(((bytes32,bytes32),uint256,bytes32[]) previous, (uint256,uint256,bytes32[]) vacant) payable returns()
func (_Rollup *RollupSession) SubmitDeposits(previous TypesCommitmentInclusionProof, vacant TypesSubtreeVacancyProof) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitDeposits(&_Rollup.TransactOpts, previous, vacant)
}

// SubmitDeposits is a paid mutator transaction binding the contract method 0x562a2bca.
//
// Solidity: function submitDeposits(((bytes32,bytes32),uint256,bytes32[]) previous, (uint256,uint256,bytes32[]) vacant) payable returns()
func (_Rollup *RollupTransactorSession) SubmitDeposits(previous TypesCommitmentInclusionProof, vacant TypesSubtreeVacancyProof) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitDeposits(&_Rollup.TransactOpts, previous, vacant)
}

// SubmitMassMigration is a paid mutator transaction binding the contract method 0x035695e3.
//
// Solidity: function submitMassMigration(bytes32[] stateRoots, uint256[2][] signatures, uint256[4][] meta, bytes32[] withdrawRoots, bytes[] txss) payable returns()
func (_Rollup *RollupTransactor) SubmitMassMigration(opts *bind.TransactOpts, stateRoots [][32]byte, signatures [][2]*big.Int, meta [][4]*big.Int, withdrawRoots [][32]byte, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitMassMigration", stateRoots, signatures, meta, withdrawRoots, txss)
}

// SubmitMassMigration is a paid mutator transaction binding the contract method 0x035695e3.
//
// Solidity: function submitMassMigration(bytes32[] stateRoots, uint256[2][] signatures, uint256[4][] meta, bytes32[] withdrawRoots, bytes[] txss) payable returns()
func (_Rollup *RollupSession) SubmitMassMigration(stateRoots [][32]byte, signatures [][2]*big.Int, meta [][4]*big.Int, withdrawRoots [][32]byte, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitMassMigration(&_Rollup.TransactOpts, stateRoots, signatures, meta, withdrawRoots, txss)
}

// SubmitMassMigration is a paid mutator transaction binding the contract method 0x035695e3.
//
// Solidity: function submitMassMigration(bytes32[] stateRoots, uint256[2][] signatures, uint256[4][] meta, bytes32[] withdrawRoots, bytes[] txss) payable returns()
func (_Rollup *RollupTransactorSession) SubmitMassMigration(stateRoots [][32]byte, signatures [][2]*big.Int, meta [][4]*big.Int, withdrawRoots [][32]byte, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitMassMigration(&_Rollup.TransactOpts, stateRoots, signatures, meta, withdrawRoots, txss)
}

// SubmitTransfer is a paid mutator transaction binding the contract method 0x5f5b95b8.
//
// Solidity: function submitTransfer(bytes32[] stateRoots, uint256[2][] signatures, uint256[] feeReceivers, bytes[] txss) payable returns()
func (_Rollup *RollupTransactor) SubmitTransfer(opts *bind.TransactOpts, stateRoots [][32]byte, signatures [][2]*big.Int, feeReceivers []*big.Int, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitTransfer", stateRoots, signatures, feeReceivers, txss)
}

// SubmitTransfer is a paid mutator transaction binding the contract method 0x5f5b95b8.
//
// Solidity: function submitTransfer(bytes32[] stateRoots, uint256[2][] signatures, uint256[] feeReceivers, bytes[] txss) payable returns()
func (_Rollup *RollupSession) SubmitTransfer(stateRoots [][32]byte, signatures [][2]*big.Int, feeReceivers []*big.Int, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitTransfer(&_Rollup.TransactOpts, stateRoots, signatures, feeReceivers, txss)
}

// SubmitTransfer is a paid mutator transaction binding the contract method 0x5f5b95b8.
//
// Solidity: function submitTransfer(bytes32[] stateRoots, uint256[2][] signatures, uint256[] feeReceivers, bytes[] txss) payable returns()
func (_Rollup *RollupTransactorSession) SubmitTransfer(stateRoots [][32]byte, signatures [][2]*big.Int, feeReceivers []*big.Int, txss [][]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitTransfer(&_Rollup.TransactOpts, stateRoots, signatures, feeReceivers, txss)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 batchID) returns()
func (_Rollup *RollupTransactor) WithdrawStake(opts *bind.TransactOpts, batchID *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "withdrawStake", batchID)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 batchID) returns()
func (_Rollup *RollupSession) WithdrawStake(batchID *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStake(&_Rollup.TransactOpts, batchID)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 batchID) returns()
func (_Rollup *RollupTransactorSession) WithdrawStake(batchID *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStake(&_Rollup.TransactOpts, batchID)
}
