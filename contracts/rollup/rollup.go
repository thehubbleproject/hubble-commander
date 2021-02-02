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
	PubkeyID *big.Int
	TokenID  *big.Int
	Balance  *big.Int
	Nonce    *big.Int
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"inputs\":[{\"internalType\":\"contractChooser\",\"name\":\"_chooser\",\"type\":\"address\"},{\"internalType\":\"contractIDepositManager\",\"name\":\"_depositManager\",\"type\":\"address\"},{\"internalType\":\"contractBLSAccountRegistry\",\"name\":\"_accountRegistry\",\"type\":\"address\"},{\"internalType\":\"contractTransfer\",\"name\":\"_transfer\",\"type\":\"address\"},{\"internalType\":\"contractMassMigration\",\"name\":\"_massMigration\",\"type\":\"address\"},{\"internalType\":\"contractCreate2Transfer\",\"name\":\"_create2Transfer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"genesisStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksToFinalise\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minGasLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTxsPerCommit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"depositSubTreeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pathToSubTree\",\"type\":\"uint256\"}],\"name\":\"DepositsFinalised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTypes.Usage\",\"name\":\"batchType\",\"type\":\"uint8\"}],\"name\":\"NewBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nDeleted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"name\":\"RollbackStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"StakeWithdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountRegistry\",\"outputs\":[{\"internalType\":\"contractBLSAccountRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"meta\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chooser\",\"outputs\":[{\"internalType\":\"contractChooser\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"create2Transfer\",\"outputs\":[{\"internalType\":\"contractCreate2Transfer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositManager\",\"outputs\":[{\"internalType\":\"contractIDepositManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.TransferBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransferCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.TransferCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeysSender\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnessesSender\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeysReceiver\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnessesReceiver\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProofWithReceiver\",\"name\":\"signatureProof\",\"type\":\"tuple\"}],\"name\":\"disputeSignatureCreate2Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.MassMigrationBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.MassMigrationCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.MMCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnesses\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProof\",\"name\":\"signatureProof\",\"type\":\"tuple\"}],\"name\":\"disputeSignatureMassMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.TransferBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransferCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.TransferCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"stateWitnesses\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"pubkeyWitnesses\",\"type\":\"bytes32[][]\"}],\"internalType\":\"structTypes.SignatureProof\",\"name\":\"signatureProof\",\"type\":\"tuple\"}],\"name\":\"disputeSignatureTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.CommitmentInclusionProof\",\"name\":\"previous\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.TransferBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransferCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.TransferCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"}],\"name\":\"disputeTransitionCreate2Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.CommitmentInclusionProof\",\"name\":\"previous\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"spokeID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.MassMigrationBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.MassMigrationCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.MMCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"}],\"name\":\"disputeTransitionMassMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.CommitmentInclusionProof\",\"name\":\"previous\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accountRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"feeReceiver\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.TransferBody\",\"name\":\"body\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.TransferCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.TransferCommitmentInclusionProof\",\"name\":\"target\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubkeyID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.StateMerkleProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"}],\"name\":\"disputeTransitionTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"meta\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invalidBatchMarker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keepRollingBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massMigration\",\"outputs\":[{\"internalType\":\"contractMassMigration\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBatchID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paramBlocksToFinalise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paramMaxTxsPerCommit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paramMinGasLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paramStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"stateRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"signatures\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeReceivers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"txss\",\"type\":\"bytes[]\"}],\"name\":\"submitCreate2Transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.CommitmentInclusionProof\",\"name\":\"previous\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pathAtDepth\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"witness\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.SubtreeVacancyProof\",\"name\":\"vacant\",\"type\":\"tuple\"}],\"name\":\"submitDeposits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"stateRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"signatures\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"meta\",\"type\":\"uint256[4][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"withdrawRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"txss\",\"type\":\"bytes[]\"}],\"name\":\"submitMassMigration\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"stateRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"signatures\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeReceivers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"txss\",\"type\":\"bytes[]\"}],\"name\":\"submitTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"contractTransfer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// Chooser is a free data retrieval call binding the contract method 0x98d17621.
//
// Solidity: function chooser() view returns(address)
func (_Rollup *RollupCaller) Chooser(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "chooser")
	return *ret0, err
}

// Chooser is a free data retrieval call binding the contract method 0x98d17621.
//
// Solidity: function chooser() view returns(address)
func (_Rollup *RollupSession) Chooser() (common.Address, error) {
	return _Rollup.Contract.Chooser(&_Rollup.CallOpts)
}

// Chooser is a free data retrieval call binding the contract method 0x98d17621.
//
// Solidity: function chooser() view returns(address)
func (_Rollup *RollupCallerSession) Chooser() (common.Address, error) {
	return _Rollup.Contract.Chooser(&_Rollup.CallOpts)
}

// Create2Transfer is a free data retrieval call binding the contract method 0x7ae8c568.
//
// Solidity: function create2Transfer() view returns(address)
func (_Rollup *RollupCaller) Create2Transfer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "create2Transfer")
	return *ret0, err
}

// Create2Transfer is a free data retrieval call binding the contract method 0x7ae8c568.
//
// Solidity: function create2Transfer() view returns(address)
func (_Rollup *RollupSession) Create2Transfer() (common.Address, error) {
	return _Rollup.Contract.Create2Transfer(&_Rollup.CallOpts)
}

// Create2Transfer is a free data retrieval call binding the contract method 0x7ae8c568.
//
// Solidity: function create2Transfer() view returns(address)
func (_Rollup *RollupCallerSession) Create2Transfer() (common.Address, error) {
	return _Rollup.Contract.Create2Transfer(&_Rollup.CallOpts)
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

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_Rollup *RollupCaller) Deposits(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "deposits", arg0)
	return *ret0, err
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_Rollup *RollupSession) Deposits(arg0 *big.Int) ([32]byte, error) {
	return _Rollup.Contract.Deposits(&_Rollup.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_Rollup *RollupCallerSession) Deposits(arg0 *big.Int) ([32]byte, error) {
	return _Rollup.Contract.Deposits(&_Rollup.CallOpts, arg0)
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

// NextBatchID is a free data retrieval call binding the contract method 0xa72598ce.
//
// Solidity: function nextBatchID() view returns(uint256)
func (_Rollup *RollupCaller) NextBatchID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "nextBatchID")
	return *ret0, err
}

// NextBatchID is a free data retrieval call binding the contract method 0xa72598ce.
//
// Solidity: function nextBatchID() view returns(uint256)
func (_Rollup *RollupSession) NextBatchID() (*big.Int, error) {
	return _Rollup.Contract.NextBatchID(&_Rollup.CallOpts)
}

// NextBatchID is a free data retrieval call binding the contract method 0xa72598ce.
//
// Solidity: function nextBatchID() view returns(uint256)
func (_Rollup *RollupCallerSession) NextBatchID() (*big.Int, error) {
	return _Rollup.Contract.NextBatchID(&_Rollup.CallOpts)
}

// ParamBlocksToFinalise is a free data retrieval call binding the contract method 0x31c2b7db.
//
// Solidity: function paramBlocksToFinalise() view returns(uint256)
func (_Rollup *RollupCaller) ParamBlocksToFinalise(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "paramBlocksToFinalise")
	return *ret0, err
}

// ParamBlocksToFinalise is a free data retrieval call binding the contract method 0x31c2b7db.
//
// Solidity: function paramBlocksToFinalise() view returns(uint256)
func (_Rollup *RollupSession) ParamBlocksToFinalise() (*big.Int, error) {
	return _Rollup.Contract.ParamBlocksToFinalise(&_Rollup.CallOpts)
}

// ParamBlocksToFinalise is a free data retrieval call binding the contract method 0x31c2b7db.
//
// Solidity: function paramBlocksToFinalise() view returns(uint256)
func (_Rollup *RollupCallerSession) ParamBlocksToFinalise() (*big.Int, error) {
	return _Rollup.Contract.ParamBlocksToFinalise(&_Rollup.CallOpts)
}

// ParamMaxTxsPerCommit is a free data retrieval call binding the contract method 0xe42f7554.
//
// Solidity: function paramMaxTxsPerCommit() view returns(uint256)
func (_Rollup *RollupCaller) ParamMaxTxsPerCommit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "paramMaxTxsPerCommit")
	return *ret0, err
}

// ParamMaxTxsPerCommit is a free data retrieval call binding the contract method 0xe42f7554.
//
// Solidity: function paramMaxTxsPerCommit() view returns(uint256)
func (_Rollup *RollupSession) ParamMaxTxsPerCommit() (*big.Int, error) {
	return _Rollup.Contract.ParamMaxTxsPerCommit(&_Rollup.CallOpts)
}

// ParamMaxTxsPerCommit is a free data retrieval call binding the contract method 0xe42f7554.
//
// Solidity: function paramMaxTxsPerCommit() view returns(uint256)
func (_Rollup *RollupCallerSession) ParamMaxTxsPerCommit() (*big.Int, error) {
	return _Rollup.Contract.ParamMaxTxsPerCommit(&_Rollup.CallOpts)
}

// ParamMinGasLeft is a free data retrieval call binding the contract method 0xca858d9d.
//
// Solidity: function paramMinGasLeft() view returns(uint256)
func (_Rollup *RollupCaller) ParamMinGasLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "paramMinGasLeft")
	return *ret0, err
}

// ParamMinGasLeft is a free data retrieval call binding the contract method 0xca858d9d.
//
// Solidity: function paramMinGasLeft() view returns(uint256)
func (_Rollup *RollupSession) ParamMinGasLeft() (*big.Int, error) {
	return _Rollup.Contract.ParamMinGasLeft(&_Rollup.CallOpts)
}

// ParamMinGasLeft is a free data retrieval call binding the contract method 0xca858d9d.
//
// Solidity: function paramMinGasLeft() view returns(uint256)
func (_Rollup *RollupCallerSession) ParamMinGasLeft() (*big.Int, error) {
	return _Rollup.Contract.ParamMinGasLeft(&_Rollup.CallOpts)
}

// ParamStakeAmount is a free data retrieval call binding the contract method 0xd1b243e2.
//
// Solidity: function paramStakeAmount() view returns(uint256)
func (_Rollup *RollupCaller) ParamStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "paramStakeAmount")
	return *ret0, err
}

// ParamStakeAmount is a free data retrieval call binding the contract method 0xd1b243e2.
//
// Solidity: function paramStakeAmount() view returns(uint256)
func (_Rollup *RollupSession) ParamStakeAmount() (*big.Int, error) {
	return _Rollup.Contract.ParamStakeAmount(&_Rollup.CallOpts)
}

// ParamStakeAmount is a free data retrieval call binding the contract method 0xd1b243e2.
//
// Solidity: function paramStakeAmount() view returns(uint256)
func (_Rollup *RollupCallerSession) ParamStakeAmount() (*big.Int, error) {
	return _Rollup.Contract.ParamStakeAmount(&_Rollup.CallOpts)
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

// WithdrawalBitmap is a free data retrieval call binding the contract method 0xfe5f9604.
//
// Solidity: function withdrawalBitmap(uint256 ) view returns(uint256)
func (_Rollup *RollupCaller) WithdrawalBitmap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "withdrawalBitmap", arg0)
	return *ret0, err
}

// WithdrawalBitmap is a free data retrieval call binding the contract method 0xfe5f9604.
//
// Solidity: function withdrawalBitmap(uint256 ) view returns(uint256)
func (_Rollup *RollupSession) WithdrawalBitmap(arg0 *big.Int) (*big.Int, error) {
	return _Rollup.Contract.WithdrawalBitmap(&_Rollup.CallOpts, arg0)
}

// WithdrawalBitmap is a free data retrieval call binding the contract method 0xfe5f9604.
//
// Solidity: function withdrawalBitmap(uint256 ) view returns(uint256)
func (_Rollup *RollupCallerSession) WithdrawalBitmap(arg0 *big.Int) (*big.Int, error) {
	return _Rollup.Contract.WithdrawalBitmap(&_Rollup.CallOpts, arg0)
}

// DisputeSignatureCreate2Transfer is a paid mutator transaction binding the contract method 0x50980f40.
//
// Solidity: function disputeSignatureCreate2Transfer(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupTransactor) DisputeSignatureCreate2Transfer(opts *bind.TransactOpts, batchID *big.Int, target TypesTransferCommitmentInclusionProof, signatureProof TypesSignatureProofWithReceiver) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "disputeSignatureCreate2Transfer", batchID, target, signatureProof)
}

// DisputeSignatureCreate2Transfer is a paid mutator transaction binding the contract method 0x50980f40.
//
// Solidity: function disputeSignatureCreate2Transfer(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupSession) DisputeSignatureCreate2Transfer(batchID *big.Int, target TypesTransferCommitmentInclusionProof, signatureProof TypesSignatureProofWithReceiver) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeSignatureCreate2Transfer(&_Rollup.TransactOpts, batchID, target, signatureProof)
}

// DisputeSignatureCreate2Transfer is a paid mutator transaction binding the contract method 0x50980f40.
//
// Solidity: function disputeSignatureCreate2Transfer(uint256 batchID, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256)[],bytes32[][],uint256[4][],bytes32[][],uint256[4][],bytes32[][]) signatureProof) returns()
func (_Rollup *RollupTransactorSession) DisputeSignatureCreate2Transfer(batchID *big.Int, target TypesTransferCommitmentInclusionProof, signatureProof TypesSignatureProofWithReceiver) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeSignatureCreate2Transfer(&_Rollup.TransactOpts, batchID, target, signatureProof)
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

// DisputeTransitionCreate2Transfer is a paid mutator transaction binding the contract method 0x5f6e91d5.
//
// Solidity: function disputeTransitionCreate2Transfer(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupTransactor) DisputeTransitionCreate2Transfer(opts *bind.TransactOpts, batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesTransferCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "disputeTransitionCreate2Transfer", batchID, previous, target, proofs)
}

// DisputeTransitionCreate2Transfer is a paid mutator transaction binding the contract method 0x5f6e91d5.
//
// Solidity: function disputeTransitionCreate2Transfer(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupSession) DisputeTransitionCreate2Transfer(batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesTransferCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeTransitionCreate2Transfer(&_Rollup.TransactOpts, batchID, previous, target, proofs)
}

// DisputeTransitionCreate2Transfer is a paid mutator transaction binding the contract method 0x5f6e91d5.
//
// Solidity: function disputeTransitionCreate2Transfer(uint256 batchID, ((bytes32,bytes32),uint256,bytes32[]) previous, ((bytes32,(bytes32,uint256[2],uint256,bytes)),uint256,bytes32[]) target, ((uint256,uint256,uint256,uint256),bytes32[])[] proofs) returns()
func (_Rollup *RollupTransactorSession) DisputeTransitionCreate2Transfer(batchID *big.Int, previous TypesCommitmentInclusionProof, target TypesTransferCommitmentInclusionProof, proofs []TypesStateMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeTransitionCreate2Transfer(&_Rollup.TransactOpts, batchID, previous, target, proofs)
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

// KeepRollingBack is a paid mutator transaction binding the contract method 0xac96f0cd.
//
// Solidity: function keepRollingBack() returns()
func (_Rollup *RollupTransactor) KeepRollingBack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "keepRollingBack")
}

// KeepRollingBack is a paid mutator transaction binding the contract method 0xac96f0cd.
//
// Solidity: function keepRollingBack() returns()
func (_Rollup *RollupSession) KeepRollingBack() (*types.Transaction, error) {
	return _Rollup.Contract.KeepRollingBack(&_Rollup.TransactOpts)
}

// KeepRollingBack is a paid mutator transaction binding the contract method 0xac96f0cd.
//
// Solidity: function keepRollingBack() returns()
func (_Rollup *RollupTransactorSession) KeepRollingBack() (*types.Transaction, error) {
	return _Rollup.Contract.KeepRollingBack(&_Rollup.TransactOpts)
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

// RollupDepositsFinalisedIterator is returned from FilterDepositsFinalised and is used to iterate over the raw logs and unpacked data for DepositsFinalised events raised by the Rollup contract.
type RollupDepositsFinalisedIterator struct {
	Event *RollupDepositsFinalised // Event containing the contract specifics and raw log

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
func (it *RollupDepositsFinalisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupDepositsFinalised)
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
		it.Event = new(RollupDepositsFinalised)
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
func (it *RollupDepositsFinalisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupDepositsFinalisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupDepositsFinalised represents a DepositsFinalised event raised by the Rollup contract.
type RollupDepositsFinalised struct {
	DepositSubTreeRoot [32]byte
	PathToSubTree      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositsFinalised is a free log retrieval operation binding the contract event 0xd1d49921e88d39febcc645126f95d5eb4fab4d9e436d650983b7221eb7bf5c28.
//
// Solidity: event DepositsFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree)
func (_Rollup *RollupFilterer) FilterDepositsFinalised(opts *bind.FilterOpts) (*RollupDepositsFinalisedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "DepositsFinalised")
	if err != nil {
		return nil, err
	}
	return &RollupDepositsFinalisedIterator{contract: _Rollup.contract, event: "DepositsFinalised", logs: logs, sub: sub}, nil
}

// WatchDepositsFinalised is a free log subscription operation binding the contract event 0xd1d49921e88d39febcc645126f95d5eb4fab4d9e436d650983b7221eb7bf5c28.
//
// Solidity: event DepositsFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree)
func (_Rollup *RollupFilterer) WatchDepositsFinalised(opts *bind.WatchOpts, sink chan<- *RollupDepositsFinalised) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "DepositsFinalised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupDepositsFinalised)
				if err := _Rollup.contract.UnpackLog(event, "DepositsFinalised", log); err != nil {
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

// ParseDepositsFinalised is a log parse operation binding the contract event 0xd1d49921e88d39febcc645126f95d5eb4fab4d9e436d650983b7221eb7bf5c28.
//
// Solidity: event DepositsFinalised(bytes32 depositSubTreeRoot, uint256 pathToSubTree)
func (_Rollup *RollupFilterer) ParseDepositsFinalised(log types.Log) (*RollupDepositsFinalised, error) {
	event := new(RollupDepositsFinalised)
	if err := _Rollup.contract.UnpackLog(event, "DepositsFinalised", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupNewBatchIterator is returned from FilterNewBatch and is used to iterate over the raw logs and unpacked data for NewBatch events raised by the Rollup contract.
type RollupNewBatchIterator struct {
	Event *RollupNewBatch // Event containing the contract specifics and raw log

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
func (it *RollupNewBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNewBatch)
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
		it.Event = new(RollupNewBatch)
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
func (it *RollupNewBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNewBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNewBatch represents a NewBatch event raised by the Rollup contract.
type RollupNewBatch struct {
	Committer common.Address
	Index     *big.Int
	BatchType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBatch is a free log retrieval operation binding the contract event 0xb34aa33e0e9ecb485e5d2fe496d4135cbedda94674af8795691fe0877da3916f.
//
// Solidity: event NewBatch(address committer, uint256 index, uint8 batchType)
func (_Rollup *RollupFilterer) FilterNewBatch(opts *bind.FilterOpts) (*RollupNewBatchIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return &RollupNewBatchIterator{contract: _Rollup.contract, event: "NewBatch", logs: logs, sub: sub}, nil
}

// WatchNewBatch is a free log subscription operation binding the contract event 0xb34aa33e0e9ecb485e5d2fe496d4135cbedda94674af8795691fe0877da3916f.
//
// Solidity: event NewBatch(address committer, uint256 index, uint8 batchType)
func (_Rollup *RollupFilterer) WatchNewBatch(opts *bind.WatchOpts, sink chan<- *RollupNewBatch) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNewBatch)
				if err := _Rollup.contract.UnpackLog(event, "NewBatch", log); err != nil {
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

// ParseNewBatch is a log parse operation binding the contract event 0xb34aa33e0e9ecb485e5d2fe496d4135cbedda94674af8795691fe0877da3916f.
//
// Solidity: event NewBatch(address committer, uint256 index, uint8 batchType)
func (_Rollup *RollupFilterer) ParseNewBatch(log types.Log) (*RollupNewBatch, error) {
	event := new(RollupNewBatch)
	if err := _Rollup.contract.UnpackLog(event, "NewBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupRollbackStatusIterator is returned from FilterRollbackStatus and is used to iterate over the raw logs and unpacked data for RollbackStatus events raised by the Rollup contract.
type RollupRollbackStatusIterator struct {
	Event *RollupRollbackStatus // Event containing the contract specifics and raw log

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
func (it *RollupRollbackStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRollbackStatus)
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
		it.Event = new(RollupRollbackStatus)
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
func (it *RollupRollbackStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRollbackStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRollbackStatus represents a RollbackStatus event raised by the Rollup contract.
type RollupRollbackStatus struct {
	StartID   *big.Int
	NDeleted  *big.Int
	Completed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRollbackStatus is a free log retrieval operation binding the contract event 0x595cb200f38fcdbf73429ffce272e53b2e923c231b2059ce20d1f4909cff1bec.
//
// Solidity: event RollbackStatus(uint256 startID, uint256 nDeleted, bool completed)
func (_Rollup *RollupFilterer) FilterRollbackStatus(opts *bind.FilterOpts) (*RollupRollbackStatusIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RollbackStatus")
	if err != nil {
		return nil, err
	}
	return &RollupRollbackStatusIterator{contract: _Rollup.contract, event: "RollbackStatus", logs: logs, sub: sub}, nil
}

// WatchRollbackStatus is a free log subscription operation binding the contract event 0x595cb200f38fcdbf73429ffce272e53b2e923c231b2059ce20d1f4909cff1bec.
//
// Solidity: event RollbackStatus(uint256 startID, uint256 nDeleted, bool completed)
func (_Rollup *RollupFilterer) WatchRollbackStatus(opts *bind.WatchOpts, sink chan<- *RollupRollbackStatus) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RollbackStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRollbackStatus)
				if err := _Rollup.contract.UnpackLog(event, "RollbackStatus", log); err != nil {
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

// ParseRollbackStatus is a log parse operation binding the contract event 0x595cb200f38fcdbf73429ffce272e53b2e923c231b2059ce20d1f4909cff1bec.
//
// Solidity: event RollbackStatus(uint256 startID, uint256 nDeleted, bool completed)
func (_Rollup *RollupFilterer) ParseRollbackStatus(log types.Log) (*RollupRollbackStatus, error) {
	event := new(RollupRollbackStatus)
	if err := _Rollup.contract.UnpackLog(event, "RollbackStatus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupStakeWithdrawIterator is returned from FilterStakeWithdraw and is used to iterate over the raw logs and unpacked data for StakeWithdraw events raised by the Rollup contract.
type RollupStakeWithdrawIterator struct {
	Event *RollupStakeWithdraw // Event containing the contract specifics and raw log

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
func (it *RollupStakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupStakeWithdraw)
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
		it.Event = new(RollupStakeWithdraw)
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
func (it *RollupStakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupStakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupStakeWithdraw represents a StakeWithdraw event raised by the Rollup contract.
type RollupStakeWithdraw struct {
	Committed common.Address
	BatchID   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdraw is a free log retrieval operation binding the contract event 0x1248d48e2de900a1010c7fce73506969ecec243600bfc08b641b158f26d857cd.
//
// Solidity: event StakeWithdraw(address committed, uint256 batchID)
func (_Rollup *RollupFilterer) FilterStakeWithdraw(opts *bind.FilterOpts) (*RollupStakeWithdrawIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "StakeWithdraw")
	if err != nil {
		return nil, err
	}
	return &RollupStakeWithdrawIterator{contract: _Rollup.contract, event: "StakeWithdraw", logs: logs, sub: sub}, nil
}

// WatchStakeWithdraw is a free log subscription operation binding the contract event 0x1248d48e2de900a1010c7fce73506969ecec243600bfc08b641b158f26d857cd.
//
// Solidity: event StakeWithdraw(address committed, uint256 batchID)
func (_Rollup *RollupFilterer) WatchStakeWithdraw(opts *bind.WatchOpts, sink chan<- *RollupStakeWithdraw) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "StakeWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupStakeWithdraw)
				if err := _Rollup.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
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

// ParseStakeWithdraw is a log parse operation binding the contract event 0x1248d48e2de900a1010c7fce73506969ecec243600bfc08b641b158f26d857cd.
//
// Solidity: event StakeWithdraw(address committed, uint256 batchID)
func (_Rollup *RollupFilterer) ParseStakeWithdraw(log types.Log) (*RollupStakeWithdraw, error) {
	event := new(RollupStakeWithdraw)
	if err := _Rollup.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
