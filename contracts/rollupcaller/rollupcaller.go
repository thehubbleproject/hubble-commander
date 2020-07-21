// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollupcaller

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
	_ = abi.U256
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

// RollupcallerABI is the input ABI used to generate the binding from.
const RollupcallerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"airdrop\",\"outputs\":[{\"internalType\":\"contractIReddit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnConsent\",\"outputs\":[{\"internalType\":\"contractIReddit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnExecution\",\"outputs\":[{\"internalType\":\"contractIReddit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractIReddit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameRegistry\",\"outputs\":[{\"internalType\":\"contractNameRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"contractIReddit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"createPublickeys\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_merkle_proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"ApplyCreateAccountTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_balanceRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof\",\"name\":\"_to_pda_proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"to_account_proof\",\"type\":\"tuple\"}],\"name\":\"processCreateAccountTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"createdAccountBytes\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.ErrorCode\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_merkle_proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"ApplyAirdropTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_balanceRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof\",\"name\":\"_from_pda_proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountProofs\",\"name\":\"accountProofs\",\"type\":\"tuple\"}],\"name\":\"processAirdropTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.ErrorCode\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_merkle_proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"ApplyTransferTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_balanceRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof\",\"name\":\"_from_pda_proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountProofs\",\"name\":\"accountProofs\",\"type\":\"tuple\"}],\"name\":\"processTransferTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.ErrorCode\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_merkle_proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"ApplyBurnConsentTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"updatedAccount\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_balanceRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof\",\"name\":\"_from_pda_proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_fromAccountProof\",\"type\":\"tuple\"}],\"name\":\"processBurnConsentTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.ErrorCode\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_merkle_proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"ApplyBurnExecutionTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"updatedAccount\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_balanceRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_fromAccountProof\",\"type\":\"tuple\"}],\"name\":\"processBurnExecutionTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"enumTypes.ErrorCode\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Rollupcaller is an auto generated Go binding around an Ethereum contract.
type Rollupcaller struct {
	RollupcallerCaller     // Read-only binding to the contract
	RollupcallerTransactor // Write-only binding to the contract
	RollupcallerFilterer   // Log filterer for contract events
}

// RollupcallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupcallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupcallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupcallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupcallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupcallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupcallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupcallerSession struct {
	Contract     *Rollupcaller     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupcallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupcallerCallerSession struct {
	Contract *RollupcallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RollupcallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupcallerTransactorSession struct {
	Contract     *RollupcallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RollupcallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupcallerRaw struct {
	Contract *Rollupcaller // Generic contract binding to access the raw methods on
}

// RollupcallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupcallerCallerRaw struct {
	Contract *RollupcallerCaller // Generic read-only contract binding to access the raw methods on
}

// RollupcallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupcallerTransactorRaw struct {
	Contract *RollupcallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupcaller creates a new instance of Rollupcaller, bound to a specific deployed contract.
func NewRollupcaller(address common.Address, backend bind.ContractBackend) (*Rollupcaller, error) {
	contract, err := bindRollupcaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollupcaller{RollupcallerCaller: RollupcallerCaller{contract: contract}, RollupcallerTransactor: RollupcallerTransactor{contract: contract}, RollupcallerFilterer: RollupcallerFilterer{contract: contract}}, nil
}

// NewRollupcallerCaller creates a new read-only instance of Rollupcaller, bound to a specific deployed contract.
func NewRollupcallerCaller(address common.Address, caller bind.ContractCaller) (*RollupcallerCaller, error) {
	contract, err := bindRollupcaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupcallerCaller{contract: contract}, nil
}

// NewRollupcallerTransactor creates a new write-only instance of Rollupcaller, bound to a specific deployed contract.
func NewRollupcallerTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupcallerTransactor, error) {
	contract, err := bindRollupcaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupcallerTransactor{contract: contract}, nil
}

// NewRollupcallerFilterer creates a new log filterer instance of Rollupcaller, bound to a specific deployed contract.
func NewRollupcallerFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupcallerFilterer, error) {
	contract, err := bindRollupcaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupcallerFilterer{contract: contract}, nil
}

// bindRollupcaller binds a generic wrapper to an already deployed contract.
func bindRollupcaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupcallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollupcaller *RollupcallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollupcaller.Contract.RollupcallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollupcaller *RollupcallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollupcaller.Contract.RollupcallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollupcaller *RollupcallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollupcaller.Contract.RollupcallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollupcaller *RollupcallerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollupcaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollupcaller *RollupcallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollupcaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollupcaller *RollupcallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollupcaller.Contract.contract.Transact(opts, method, params...)
}

// ApplyAirdropTx is a free data retrieval call binding the contract method 0xcef3f773.
//
// Solidity: function ApplyAirdropTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCaller) ApplyAirdropTx(opts *bind.CallOpts, _merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	var (
		ret0 = new([]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Rollupcaller.contract.Call(opts, out, "ApplyAirdropTx", _merkle_proof, txBytes)
	return *ret0, *ret1, err
}

// ApplyAirdropTx is a free data retrieval call binding the contract method 0xcef3f773.
//
// Solidity: function ApplyAirdropTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerSession) ApplyAirdropTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	return _Rollupcaller.Contract.ApplyAirdropTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyAirdropTx is a free data retrieval call binding the contract method 0xcef3f773.
//
// Solidity: function ApplyAirdropTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCallerSession) ApplyAirdropTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	return _Rollupcaller.Contract.ApplyAirdropTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyBurnConsentTx is a free data retrieval call binding the contract method 0xea2849bc.
//
// Solidity: function ApplyBurnConsentTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCaller) ApplyBurnConsentTx(opts *bind.CallOpts, _merkle_proof TypesAccountMerkleProof, txBytes []byte) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	ret := new(struct {
		UpdatedAccount []byte
		NewRoot        [32]byte
	})
	out := ret
	err := _Rollupcaller.contract.Call(opts, out, "ApplyBurnConsentTx", _merkle_proof, txBytes)
	return *ret, err
}

// ApplyBurnConsentTx is a free data retrieval call binding the contract method 0xea2849bc.
//
// Solidity: function ApplyBurnConsentTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Rollupcaller *RollupcallerSession) ApplyBurnConsentTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	return _Rollupcaller.Contract.ApplyBurnConsentTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyBurnConsentTx is a free data retrieval call binding the contract method 0xea2849bc.
//
// Solidity: function ApplyBurnConsentTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCallerSession) ApplyBurnConsentTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	return _Rollupcaller.Contract.ApplyBurnConsentTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyBurnExecutionTx is a free data retrieval call binding the contract method 0x952c7118.
//
// Solidity: function ApplyBurnExecutionTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCaller) ApplyBurnExecutionTx(opts *bind.CallOpts, _merkle_proof TypesAccountMerkleProof, txBytes []byte) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	ret := new(struct {
		UpdatedAccount []byte
		NewRoot        [32]byte
	})
	out := ret
	err := _Rollupcaller.contract.Call(opts, out, "ApplyBurnExecutionTx", _merkle_proof, txBytes)
	return *ret, err
}

// ApplyBurnExecutionTx is a free data retrieval call binding the contract method 0x952c7118.
//
// Solidity: function ApplyBurnExecutionTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Rollupcaller *RollupcallerSession) ApplyBurnExecutionTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	return _Rollupcaller.Contract.ApplyBurnExecutionTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyBurnExecutionTx is a free data retrieval call binding the contract method 0x952c7118.
//
// Solidity: function ApplyBurnExecutionTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCallerSession) ApplyBurnExecutionTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	return _Rollupcaller.Contract.ApplyBurnExecutionTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyCreateAccountTx is a free data retrieval call binding the contract method 0xfa41dbd3.
//
// Solidity: function ApplyCreateAccountTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCaller) ApplyCreateAccountTx(opts *bind.CallOpts, _merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	var (
		ret0 = new([]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Rollupcaller.contract.Call(opts, out, "ApplyCreateAccountTx", _merkle_proof, txBytes)
	return *ret0, *ret1, err
}

// ApplyCreateAccountTx is a free data retrieval call binding the contract method 0xfa41dbd3.
//
// Solidity: function ApplyCreateAccountTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerSession) ApplyCreateAccountTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	return _Rollupcaller.Contract.ApplyCreateAccountTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyCreateAccountTx is a free data retrieval call binding the contract method 0xfa41dbd3.
//
// Solidity: function ApplyCreateAccountTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCallerSession) ApplyCreateAccountTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	return _Rollupcaller.Contract.ApplyCreateAccountTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyTransferTx is a free data retrieval call binding the contract method 0x66e468c6.
//
// Solidity: function ApplyTransferTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCaller) ApplyTransferTx(opts *bind.CallOpts, _merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	var (
		ret0 = new([]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Rollupcaller.contract.Call(opts, out, "ApplyTransferTx", _merkle_proof, txBytes)
	return *ret0, *ret1, err
}

// ApplyTransferTx is a free data retrieval call binding the contract method 0x66e468c6.
//
// Solidity: function ApplyTransferTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerSession) ApplyTransferTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	return _Rollupcaller.Contract.ApplyTransferTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// ApplyTransferTx is a free data retrieval call binding the contract method 0x66e468c6.
//
// Solidity: function ApplyTransferTx(TypesAccountMerkleProof _merkle_proof, bytes txBytes) view returns(bytes, bytes32 newRoot)
func (_Rollupcaller *RollupcallerCallerSession) ApplyTransferTx(_merkle_proof TypesAccountMerkleProof, txBytes []byte) ([]byte, [32]byte, error) {
	return _Rollupcaller.Contract.ApplyTransferTx(&_Rollupcaller.CallOpts, _merkle_proof, txBytes)
}

// Airdrop is a free data retrieval call binding the contract method 0x3884d635.
//
// Solidity: function airdrop() view returns(address)
func (_Rollupcaller *RollupcallerCaller) Airdrop(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollupcaller.contract.Call(opts, out, "airdrop")
	return *ret0, err
}

// Airdrop is a free data retrieval call binding the contract method 0x3884d635.
//
// Solidity: function airdrop() view returns(address)
func (_Rollupcaller *RollupcallerSession) Airdrop() (common.Address, error) {
	return _Rollupcaller.Contract.Airdrop(&_Rollupcaller.CallOpts)
}

// Airdrop is a free data retrieval call binding the contract method 0x3884d635.
//
// Solidity: function airdrop() view returns(address)
func (_Rollupcaller *RollupcallerCallerSession) Airdrop() (common.Address, error) {
	return _Rollupcaller.Contract.Airdrop(&_Rollupcaller.CallOpts)
}

// BurnConsent is a free data retrieval call binding the contract method 0x2f23041e.
//
// Solidity: function burnConsent() view returns(address)
func (_Rollupcaller *RollupcallerCaller) BurnConsent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollupcaller.contract.Call(opts, out, "burnConsent")
	return *ret0, err
}

// BurnConsent is a free data retrieval call binding the contract method 0x2f23041e.
//
// Solidity: function burnConsent() view returns(address)
func (_Rollupcaller *RollupcallerSession) BurnConsent() (common.Address, error) {
	return _Rollupcaller.Contract.BurnConsent(&_Rollupcaller.CallOpts)
}

// BurnConsent is a free data retrieval call binding the contract method 0x2f23041e.
//
// Solidity: function burnConsent() view returns(address)
func (_Rollupcaller *RollupcallerCallerSession) BurnConsent() (common.Address, error) {
	return _Rollupcaller.Contract.BurnConsent(&_Rollupcaller.CallOpts)
}

// BurnExecution is a free data retrieval call binding the contract method 0xf4fc967e.
//
// Solidity: function burnExecution() view returns(address)
func (_Rollupcaller *RollupcallerCaller) BurnExecution(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollupcaller.contract.Call(opts, out, "burnExecution")
	return *ret0, err
}

// BurnExecution is a free data retrieval call binding the contract method 0xf4fc967e.
//
// Solidity: function burnExecution() view returns(address)
func (_Rollupcaller *RollupcallerSession) BurnExecution() (common.Address, error) {
	return _Rollupcaller.Contract.BurnExecution(&_Rollupcaller.CallOpts)
}

// BurnExecution is a free data retrieval call binding the contract method 0xf4fc967e.
//
// Solidity: function burnExecution() view returns(address)
func (_Rollupcaller *RollupcallerCallerSession) BurnExecution() (common.Address, error) {
	return _Rollupcaller.Contract.BurnExecution(&_Rollupcaller.CallOpts)
}

// CreateAccount is a free data retrieval call binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() view returns(address)
func (_Rollupcaller *RollupcallerCaller) CreateAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollupcaller.contract.Call(opts, out, "createAccount")
	return *ret0, err
}

// CreateAccount is a free data retrieval call binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() view returns(address)
func (_Rollupcaller *RollupcallerSession) CreateAccount() (common.Address, error) {
	return _Rollupcaller.Contract.CreateAccount(&_Rollupcaller.CallOpts)
}

// CreateAccount is a free data retrieval call binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() view returns(address)
func (_Rollupcaller *RollupcallerCallerSession) CreateAccount() (common.Address, error) {
	return _Rollupcaller.Contract.CreateAccount(&_Rollupcaller.CallOpts)
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Rollupcaller *RollupcallerCaller) NameRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollupcaller.contract.Call(opts, out, "nameRegistry")
	return *ret0, err
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Rollupcaller *RollupcallerSession) NameRegistry() (common.Address, error) {
	return _Rollupcaller.Contract.NameRegistry(&_Rollupcaller.CallOpts)
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Rollupcaller *RollupcallerCallerSession) NameRegistry() (common.Address, error) {
	return _Rollupcaller.Contract.NameRegistry(&_Rollupcaller.CallOpts)
}

// ProcessAirdropTx is a free data retrieval call binding the contract method 0x7810dc60.
//
// Solidity: function processAirdropTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerCaller) ProcessAirdropTx(opts *bind.CallOpts, _balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, uint8, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([]byte)
		ret2 = new([]byte)
		ret3 = new(uint8)
		ret4 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Rollupcaller.contract.Call(opts, out, "processAirdropTx", _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, accountProofs)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ProcessAirdropTx is a free data retrieval call binding the contract method 0x7810dc60.
//
// Solidity: function processAirdropTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerSession) ProcessAirdropTx(_balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessAirdropTx(&_Rollupcaller.CallOpts, _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, accountProofs)
}

// ProcessAirdropTx is a free data retrieval call binding the contract method 0x7810dc60.
//
// Solidity: function processAirdropTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerCallerSession) ProcessAirdropTx(_balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessAirdropTx(&_Rollupcaller.CallOpts, _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, accountProofs)
}

// ProcessBurnConsentTx is a free data retrieval call binding the contract method 0x8fec478c.
//
// Solidity: function processBurnConsentTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountMerkleProof _fromAccountProof) view returns(bytes32, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerCaller) ProcessBurnConsentTx(opts *bind.CallOpts, _balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, _fromAccountProof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([]byte)
		ret2 = new(uint8)
		ret3 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Rollupcaller.contract.Call(opts, out, "processBurnConsentTx", _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, _fromAccountProof)
	return *ret0, *ret1, *ret2, *ret3, err
}

// ProcessBurnConsentTx is a free data retrieval call binding the contract method 0x8fec478c.
//
// Solidity: function processBurnConsentTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountMerkleProof _fromAccountProof) view returns(bytes32, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerSession) ProcessBurnConsentTx(_balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, _fromAccountProof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessBurnConsentTx(&_Rollupcaller.CallOpts, _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, _fromAccountProof)
}

// ProcessBurnConsentTx is a free data retrieval call binding the contract method 0x8fec478c.
//
// Solidity: function processBurnConsentTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountMerkleProof _fromAccountProof) view returns(bytes32, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerCallerSession) ProcessBurnConsentTx(_balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, _fromAccountProof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessBurnConsentTx(&_Rollupcaller.CallOpts, _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, _fromAccountProof)
}

// ProcessBurnExecutionTx is a free data retrieval call binding the contract method 0x3c697977.
//
// Solidity: function processBurnExecutionTx(bytes32 _balanceRoot, bytes txBytes, TypesAccountMerkleProof _fromAccountProof) view returns(bytes32, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerCaller) ProcessBurnExecutionTx(opts *bind.CallOpts, _balanceRoot [32]byte, txBytes []byte, _fromAccountProof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([]byte)
		ret2 = new(uint8)
		ret3 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Rollupcaller.contract.Call(opts, out, "processBurnExecutionTx", _balanceRoot, txBytes, _fromAccountProof)
	return *ret0, *ret1, *ret2, *ret3, err
}

// ProcessBurnExecutionTx is a free data retrieval call binding the contract method 0x3c697977.
//
// Solidity: function processBurnExecutionTx(bytes32 _balanceRoot, bytes txBytes, TypesAccountMerkleProof _fromAccountProof) view returns(bytes32, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerSession) ProcessBurnExecutionTx(_balanceRoot [32]byte, txBytes []byte, _fromAccountProof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessBurnExecutionTx(&_Rollupcaller.CallOpts, _balanceRoot, txBytes, _fromAccountProof)
}

// ProcessBurnExecutionTx is a free data retrieval call binding the contract method 0x3c697977.
//
// Solidity: function processBurnExecutionTx(bytes32 _balanceRoot, bytes txBytes, TypesAccountMerkleProof _fromAccountProof) view returns(bytes32, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerCallerSession) ProcessBurnExecutionTx(_balanceRoot [32]byte, txBytes []byte, _fromAccountProof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessBurnExecutionTx(&_Rollupcaller.CallOpts, _balanceRoot, txBytes, _fromAccountProof)
}

// ProcessCreateAccountTx is a free data retrieval call binding the contract method 0xa25ce890.
//
// Solidity: function processCreateAccountTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _to_pda_proof, TypesAccountMerkleProof to_account_proof) view returns(bytes32 newRoot, bytes createdAccountBytes, uint8, bool)
func (_Rollupcaller *RollupcallerCaller) ProcessCreateAccountTx(opts *bind.CallOpts, _balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _to_pda_proof TypesPDAMerkleProof, to_account_proof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([]byte)
		ret2 = new(uint8)
		ret3 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Rollupcaller.contract.Call(opts, out, "processCreateAccountTx", _balanceRoot, _accountsRoot, sig, txBytes, _to_pda_proof, to_account_proof)
	return *ret0, *ret1, *ret2, *ret3, err
}

// ProcessCreateAccountTx is a free data retrieval call binding the contract method 0xa25ce890.
//
// Solidity: function processCreateAccountTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _to_pda_proof, TypesAccountMerkleProof to_account_proof) view returns(bytes32 newRoot, bytes createdAccountBytes, uint8, bool)
func (_Rollupcaller *RollupcallerSession) ProcessCreateAccountTx(_balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _to_pda_proof TypesPDAMerkleProof, to_account_proof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessCreateAccountTx(&_Rollupcaller.CallOpts, _balanceRoot, _accountsRoot, sig, txBytes, _to_pda_proof, to_account_proof)
}

// ProcessCreateAccountTx is a free data retrieval call binding the contract method 0xa25ce890.
//
// Solidity: function processCreateAccountTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _to_pda_proof, TypesAccountMerkleProof to_account_proof) view returns(bytes32 newRoot, bytes createdAccountBytes, uint8, bool)
func (_Rollupcaller *RollupcallerCallerSession) ProcessCreateAccountTx(_balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _to_pda_proof TypesPDAMerkleProof, to_account_proof TypesAccountMerkleProof) ([32]byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessCreateAccountTx(&_Rollupcaller.CallOpts, _balanceRoot, _accountsRoot, sig, txBytes, _to_pda_proof, to_account_proof)
}

// ProcessTransferTx is a free data retrieval call binding the contract method 0xb6626b7d.
//
// Solidity: function processTransferTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerCaller) ProcessTransferTx(opts *bind.CallOpts, _balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, uint8, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([]byte)
		ret2 = new([]byte)
		ret3 = new(uint8)
		ret4 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Rollupcaller.contract.Call(opts, out, "processTransferTx", _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, accountProofs)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ProcessTransferTx is a free data retrieval call binding the contract method 0xb6626b7d.
//
// Solidity: function processTransferTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerSession) ProcessTransferTx(_balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessTransferTx(&_Rollupcaller.CallOpts, _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, accountProofs)
}

// ProcessTransferTx is a free data retrieval call binding the contract method 0xb6626b7d.
//
// Solidity: function processTransferTx(bytes32 _balanceRoot, bytes32 _accountsRoot, bytes sig, bytes txBytes, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint8, bool)
func (_Rollupcaller *RollupcallerCallerSession) ProcessTransferTx(_balanceRoot [32]byte, _accountsRoot [32]byte, sig []byte, txBytes []byte, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, uint8, bool, error) {
	return _Rollupcaller.Contract.ProcessTransferTx(&_Rollupcaller.CallOpts, _balanceRoot, _accountsRoot, sig, txBytes, _from_pda_proof, accountProofs)
}

// Transfer is a free data retrieval call binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() view returns(address)
func (_Rollupcaller *RollupcallerCaller) Transfer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollupcaller.contract.Call(opts, out, "transfer")
	return *ret0, err
}

// Transfer is a free data retrieval call binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() view returns(address)
func (_Rollupcaller *RollupcallerSession) Transfer() (common.Address, error) {
	return _Rollupcaller.Contract.Transfer(&_Rollupcaller.CallOpts)
}

// Transfer is a free data retrieval call binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() view returns(address)
func (_Rollupcaller *RollupcallerCallerSession) Transfer() (common.Address, error) {
	return _Rollupcaller.Contract.Transfer(&_Rollupcaller.CallOpts)
}

// CreatePublickeys is a paid mutator transaction binding the contract method 0x807b55dc.
//
// Solidity: function createPublickeys(bytes[] publicKeys) returns(uint256[])
func (_Rollupcaller *RollupcallerTransactor) CreatePublickeys(opts *bind.TransactOpts, publicKeys [][]byte) (*types.Transaction, error) {
	return _Rollupcaller.contract.Transact(opts, "createPublickeys", publicKeys)
}

// CreatePublickeys is a paid mutator transaction binding the contract method 0x807b55dc.
//
// Solidity: function createPublickeys(bytes[] publicKeys) returns(uint256[])
func (_Rollupcaller *RollupcallerSession) CreatePublickeys(publicKeys [][]byte) (*types.Transaction, error) {
	return _Rollupcaller.Contract.CreatePublickeys(&_Rollupcaller.TransactOpts, publicKeys)
}

// CreatePublickeys is a paid mutator transaction binding the contract method 0x807b55dc.
//
// Solidity: function createPublickeys(bytes[] publicKeys) returns(uint256[])
func (_Rollupcaller *RollupcallerTransactorSession) CreatePublickeys(publicKeys [][]byte) (*types.Transaction, error) {
	return _Rollupcaller.Contract.CreatePublickeys(&_Rollupcaller.TransactOpts, publicKeys)
}
