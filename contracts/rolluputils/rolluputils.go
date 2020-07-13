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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TypesBurnConsent is an auto generated low-level Go binding around an user-defined struct.
type TypesBurnConsent struct {
	FromIndex *big.Int
	Amount    *big.Int
	Cancel    bool
	Signature []byte
}

// TypesBurnExecution is an auto generated low-level Go binding around an user-defined struct.
type TypesBurnExecution struct {
	FromIndex *big.Int
}

// TypesCreateAccount is an auto generated low-level Go binding around an user-defined struct.
type TypesCreateAccount struct {
	ToIndex   *big.Int
	TokenType *big.Int
	Signature []byte
}

// TypesDropTx is an auto generated low-level Go binding around an user-defined struct.
type TypesDropTx struct {
	FromIndex *big.Int
	ToIndex   *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
	Signature []byte
}

// TypesPDALeaf is an auto generated low-level Go binding around an user-defined struct.
type TypesPDALeaf struct {
	Pubkey []byte
}

// TypesTransaction is an auto generated low-level Go binding around an user-defined struct.
type TypesTransaction struct {
	FromIndex *big.Int
	ToIndex   *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
	Signature []byte
}

// TypesUserAccount is an auto generated low-level Go binding around an user-defined struct.
type TypesUserAccount struct {
	ID        *big.Int
	TokenType *big.Int
	Balance   *big.Int
	Nonce     *big.Int
}

// RolluputilsABI is the input ABI used to generate the binding from.
const RolluputilsABI = "[{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"_PDA_Leaf\",\"type\":\"tuple\"}],\"name\":\"PDALeafToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"accountBytes\",\"type\":\"bytes\"}],\"name\":\"AccountFromBytes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"BytesFromAccount\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"name\":\"BytesFromAccountDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"name\":\"getAccountHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"HashFromAccount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"TxFromBytesBurnConsent\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancel\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"TxFromBytesBurnExecution\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressConsent\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancel\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressExecution\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressCreateAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.CreateAccount\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancel\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressConsent\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressExecution\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.CreateAccount\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressCreateAccount\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.CreateAccount\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromCreateAccount\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancel\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromBurnConsent\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromBurnExecution\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"name\":\"BytesFromTxCreateAccountDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancel\",\"type\":\"bool\"}],\"name\":\"BytesFromTxBurnConsentDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"}],\"name\":\"BytesFromTxBurnExecutionDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancel\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromConsent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromExecution\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getTxSignBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BytesFromTxDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"TxFromBytesDeconstructed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"TxFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressTxWithMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.DropTx\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressDrop\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dropBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressDrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.DropTx\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BytesFromTxAirdropDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.DropTx\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromAirdrop\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getDropSignBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressDropNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pub\",\"type\":\"bytes\"}],\"name\":\"calculateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetGenesisLeaves\",\"outputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"leaves\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetGenesisDataBlocks\",\"outputs\":[{\"internalType\":\"bytes[2]\",\"name\":\"dataBlocks\",\"type\":\"bytes[2]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

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

// AccountFromBytes is a free data retrieval call binding the contract method 0x1a636e86.
//
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType)
func (_Rolluputils *RolluputilsCaller) AccountFromBytes(opts *bind.CallOpts, accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
}, error) {
	ret := new(struct {
		ID        *big.Int
		Balance   *big.Int
		Nonce     *big.Int
		TokenType *big.Int
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "AccountFromBytes", accountBytes)
	return *ret, err
}

// AccountFromBytes is a free data retrieval call binding the contract method 0x1a636e86.
//
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType)
func (_Rolluputils *RolluputilsSession) AccountFromBytes(accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
}, error) {
	return _Rolluputils.Contract.AccountFromBytes(&_Rolluputils.CallOpts, accountBytes)
}

// AccountFromBytes is a free data retrieval call binding the contract method 0x1a636e86.
//
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType)
func (_Rolluputils *RolluputilsCallerSession) AccountFromBytes(accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
}, error) {
	return _Rolluputils.Contract.AccountFromBytes(&_Rolluputils.CallOpts, accountBytes)
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(TypesUserAccount account) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromAccount(opts *bind.CallOpts, account TypesUserAccount) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromAccount", account)
	return *ret0, err
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(TypesUserAccount account) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAccount(account TypesUserAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccount(&_Rolluputils.CallOpts, account)
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(TypesUserAccount account) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAccount(account TypesUserAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccount(&_Rolluputils.CallOpts, account)
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0x5bc64a2b.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromAccountDeconstructed(opts *bind.CallOpts, ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromAccountDeconstructed", ID, balance, nonce, tokenType)
	return *ret0, err
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0x5bc64a2b.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAccountDeconstructed(ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccountDeconstructed(&_Rolluputils.CallOpts, ID, balance, nonce, tokenType)
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0x5bc64a2b.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAccountDeconstructed(ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccountDeconstructed(&_Rolluputils.CallOpts, ID, balance, nonce, tokenType)
}

// BytesFromAirdrop is a free data retrieval call binding the contract method 0x29d87dd9.
//
// Solidity: function BytesFromAirdrop(TypesDropTx _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromAirdrop(opts *bind.CallOpts, _tx TypesDropTx) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromAirdrop", _tx)
	return *ret0, err
}

// BytesFromAirdrop is a free data retrieval call binding the contract method 0x29d87dd9.
//
// Solidity: function BytesFromAirdrop(TypesDropTx _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAirdrop(_tx TypesDropTx) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAirdrop(&_Rolluputils.CallOpts, _tx)
}

// BytesFromAirdrop is a free data retrieval call binding the contract method 0x29d87dd9.
//
// Solidity: function BytesFromAirdrop(TypesDropTx _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAirdrop(_tx TypesDropTx) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAirdrop(&_Rolluputils.CallOpts, _tx)
}

// BytesFromBurnConsent is a free data retrieval call binding the contract method 0x1f66ac03.
//
// Solidity: function BytesFromBurnConsent(TypesBurnConsent _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromBurnConsent(opts *bind.CallOpts, _tx TypesBurnConsent) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromBurnConsent", _tx)
	return *ret0, err
}

// BytesFromBurnConsent is a free data retrieval call binding the contract method 0x1f66ac03.
//
// Solidity: function BytesFromBurnConsent(TypesBurnConsent _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromBurnConsent(_tx TypesBurnConsent) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnConsent(&_Rolluputils.CallOpts, _tx)
}

// BytesFromBurnConsent is a free data retrieval call binding the contract method 0x1f66ac03.
//
// Solidity: function BytesFromBurnConsent(TypesBurnConsent _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromBurnConsent(_tx TypesBurnConsent) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnConsent(&_Rolluputils.CallOpts, _tx)
}

// BytesFromBurnExecution is a free data retrieval call binding the contract method 0x74a7ae9b.
//
// Solidity: function BytesFromBurnExecution(TypesBurnExecution _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromBurnExecution(opts *bind.CallOpts, _tx TypesBurnExecution) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromBurnExecution", _tx)
	return *ret0, err
}

// BytesFromBurnExecution is a free data retrieval call binding the contract method 0x74a7ae9b.
//
// Solidity: function BytesFromBurnExecution(TypesBurnExecution _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromBurnExecution(_tx TypesBurnExecution) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnExecution(&_Rolluputils.CallOpts, _tx)
}

// BytesFromBurnExecution is a free data retrieval call binding the contract method 0x74a7ae9b.
//
// Solidity: function BytesFromBurnExecution(TypesBurnExecution _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromBurnExecution(_tx TypesBurnExecution) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnExecution(&_Rolluputils.CallOpts, _tx)
}

// BytesFromCreateAccount is a free data retrieval call binding the contract method 0x32c9b1ab.
//
// Solidity: function BytesFromCreateAccount(TypesCreateAccount _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromCreateAccount(opts *bind.CallOpts, _tx TypesCreateAccount) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromCreateAccount", _tx)
	return *ret0, err
}

// BytesFromCreateAccount is a free data retrieval call binding the contract method 0x32c9b1ab.
//
// Solidity: function BytesFromCreateAccount(TypesCreateAccount _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromCreateAccount(_tx TypesCreateAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromCreateAccount(&_Rolluputils.CallOpts, _tx)
}

// BytesFromCreateAccount is a free data retrieval call binding the contract method 0x32c9b1ab.
//
// Solidity: function BytesFromCreateAccount(TypesCreateAccount _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromCreateAccount(_tx TypesCreateAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromCreateAccount(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTx is a free data retrieval call binding the contract method 0xd0dbc902.
//
// Solidity: function BytesFromTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTx(opts *bind.CallOpts, _tx TypesTransaction) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTx", _tx)
	return *ret0, err
}

// BytesFromTx is a free data retrieval call binding the contract method 0xd0dbc902.
//
// Solidity: function BytesFromTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTx is a free data retrieval call binding the contract method 0xd0dbc902.
//
// Solidity: function BytesFromTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTxAirdropDeconstructed is a free data retrieval call binding the contract method 0x36a36323.
//
// Solidity: function BytesFromTxAirdropDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTxAirdropDeconstructed(opts *bind.CallOpts, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTxAirdropDeconstructed", from, to, tokenType, nonce, txType, amount)
	return *ret0, err
}

// BytesFromTxAirdropDeconstructed is a free data retrieval call binding the contract method 0x36a36323.
//
// Solidity: function BytesFromTxAirdropDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTxAirdropDeconstructed(from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxAirdropDeconstructed(&_Rolluputils.CallOpts, from, to, tokenType, nonce, txType, amount)
}

// BytesFromTxAirdropDeconstructed is a free data retrieval call binding the contract method 0x36a36323.
//
// Solidity: function BytesFromTxAirdropDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTxAirdropDeconstructed(from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxAirdropDeconstructed(&_Rolluputils.CallOpts, from, to, tokenType, nonce, txType, amount)
}

// BytesFromTxBurnConsentDeconstructed is a free data retrieval call binding the contract method 0xe8b6a4db.
//
// Solidity: function BytesFromTxBurnConsentDeconstructed(uint256 from, uint256 amount, bool cancel) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTxBurnConsentDeconstructed(opts *bind.CallOpts, from *big.Int, amount *big.Int, cancel bool) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTxBurnConsentDeconstructed", from, amount, cancel)
	return *ret0, err
}

// BytesFromTxBurnConsentDeconstructed is a free data retrieval call binding the contract method 0xe8b6a4db.
//
// Solidity: function BytesFromTxBurnConsentDeconstructed(uint256 from, uint256 amount, bool cancel) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTxBurnConsentDeconstructed(from *big.Int, amount *big.Int, cancel bool) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxBurnConsentDeconstructed(&_Rolluputils.CallOpts, from, amount, cancel)
}

// BytesFromTxBurnConsentDeconstructed is a free data retrieval call binding the contract method 0xe8b6a4db.
//
// Solidity: function BytesFromTxBurnConsentDeconstructed(uint256 from, uint256 amount, bool cancel) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTxBurnConsentDeconstructed(from *big.Int, amount *big.Int, cancel bool) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxBurnConsentDeconstructed(&_Rolluputils.CallOpts, from, amount, cancel)
}

// BytesFromTxBurnExecutionDeconstructed is a free data retrieval call binding the contract method 0xee928b85.
//
// Solidity: function BytesFromTxBurnExecutionDeconstructed(uint256 from) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTxBurnExecutionDeconstructed(opts *bind.CallOpts, from *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTxBurnExecutionDeconstructed", from)
	return *ret0, err
}

// BytesFromTxBurnExecutionDeconstructed is a free data retrieval call binding the contract method 0xee928b85.
//
// Solidity: function BytesFromTxBurnExecutionDeconstructed(uint256 from) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTxBurnExecutionDeconstructed(from *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxBurnExecutionDeconstructed(&_Rolluputils.CallOpts, from)
}

// BytesFromTxBurnExecutionDeconstructed is a free data retrieval call binding the contract method 0xee928b85.
//
// Solidity: function BytesFromTxBurnExecutionDeconstructed(uint256 from) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTxBurnExecutionDeconstructed(from *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxBurnExecutionDeconstructed(&_Rolluputils.CallOpts, from)
}

// BytesFromTxCreateAccountDeconstructed is a free data retrieval call binding the contract method 0xd65e80bb.
//
// Solidity: function BytesFromTxCreateAccountDeconstructed(uint256 to, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTxCreateAccountDeconstructed(opts *bind.CallOpts, to *big.Int, tokenType *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTxCreateAccountDeconstructed", to, tokenType)
	return *ret0, err
}

// BytesFromTxCreateAccountDeconstructed is a free data retrieval call binding the contract method 0xd65e80bb.
//
// Solidity: function BytesFromTxCreateAccountDeconstructed(uint256 to, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTxCreateAccountDeconstructed(to *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxCreateAccountDeconstructed(&_Rolluputils.CallOpts, to, tokenType)
}

// BytesFromTxCreateAccountDeconstructed is a free data retrieval call binding the contract method 0xd65e80bb.
//
// Solidity: function BytesFromTxCreateAccountDeconstructed(uint256 to, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTxCreateAccountDeconstructed(to *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxCreateAccountDeconstructed(&_Rolluputils.CallOpts, to, tokenType)
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTxDeconstructed(opts *bind.CallOpts, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTxDeconstructed", from, to, tokenType, nonce, txType, amount)
	return *ret0, err
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTxDeconstructed(from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxDeconstructed(&_Rolluputils.CallOpts, from, to, tokenType, nonce, txType, amount)
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTxDeconstructed(from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxDeconstructed(&_Rolluputils.CallOpts, from, to, tokenType, nonce, txType, amount)
}

// CompressConsent is a free data retrieval call binding the contract method 0xa0ad574e.
//
// Solidity: function CompressConsent(TypesBurnConsent _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressConsent(opts *bind.CallOpts, _tx TypesBurnConsent) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressConsent", _tx)
	return *ret0, err
}

// CompressConsent is a free data retrieval call binding the contract method 0xa0ad574e.
//
// Solidity: function CompressConsent(TypesBurnConsent _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressConsent(_tx TypesBurnConsent) ([]byte, error) {
	return _Rolluputils.Contract.CompressConsent(&_Rolluputils.CallOpts, _tx)
}

// CompressConsent is a free data retrieval call binding the contract method 0xa0ad574e.
//
// Solidity: function CompressConsent(TypesBurnConsent _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressConsent(_tx TypesBurnConsent) ([]byte, error) {
	return _Rolluputils.Contract.CompressConsent(&_Rolluputils.CallOpts, _tx)
}

// CompressCreateAccount is a free data retrieval call binding the contract method 0xe12b9991.
//
// Solidity: function CompressCreateAccount(TypesCreateAccount _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressCreateAccount(opts *bind.CallOpts, _tx TypesCreateAccount) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressCreateAccount", _tx)
	return *ret0, err
}

// CompressCreateAccount is a free data retrieval call binding the contract method 0xe12b9991.
//
// Solidity: function CompressCreateAccount(TypesCreateAccount _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressCreateAccount(_tx TypesCreateAccount) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccount(&_Rolluputils.CallOpts, _tx)
}

// CompressCreateAccount is a free data retrieval call binding the contract method 0xe12b9991.
//
// Solidity: function CompressCreateAccount(TypesCreateAccount _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressCreateAccount(_tx TypesCreateAccount) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccount(&_Rolluputils.CallOpts, _tx)
}

// CompressDrop is a free data retrieval call binding the contract method 0x6da01ca1.
//
// Solidity: function CompressDrop(TypesDropTx _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressDrop(opts *bind.CallOpts, _tx TypesDropTx) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressDrop", _tx)
	return *ret0, err
}

// CompressDrop is a free data retrieval call binding the contract method 0x6da01ca1.
//
// Solidity: function CompressDrop(TypesDropTx _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressDrop(_tx TypesDropTx) ([]byte, error) {
	return _Rolluputils.Contract.CompressDrop(&_Rolluputils.CallOpts, _tx)
}

// CompressDrop is a free data retrieval call binding the contract method 0x6da01ca1.
//
// Solidity: function CompressDrop(TypesDropTx _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressDrop(_tx TypesDropTx) ([]byte, error) {
	return _Rolluputils.Contract.CompressDrop(&_Rolluputils.CallOpts, _tx)
}

// CompressDropNoStruct is a free data retrieval call binding the contract method 0x6ebfcd20.
//
// Solidity: function CompressDropNoStruct(uint256 toIndex, uint256 amount, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressDropNoStruct(opts *bind.CallOpts, toIndex *big.Int, amount *big.Int, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressDropNoStruct", toIndex, amount, sig)
	return *ret0, err
}

// CompressDropNoStruct is a free data retrieval call binding the contract method 0x6ebfcd20.
//
// Solidity: function CompressDropNoStruct(uint256 toIndex, uint256 amount, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressDropNoStruct(toIndex *big.Int, amount *big.Int, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressDropNoStruct(&_Rolluputils.CallOpts, toIndex, amount, sig)
}

// CompressDropNoStruct is a free data retrieval call binding the contract method 0x6ebfcd20.
//
// Solidity: function CompressDropNoStruct(uint256 toIndex, uint256 amount, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressDropNoStruct(toIndex *big.Int, amount *big.Int, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressDropNoStruct(&_Rolluputils.CallOpts, toIndex, amount, sig)
}

// CompressExecution is a free data retrieval call binding the contract method 0x324a4335.
//
// Solidity: function CompressExecution(TypesBurnExecution _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressExecution(opts *bind.CallOpts, _tx TypesBurnExecution) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressExecution", _tx)
	return *ret0, err
}

// CompressExecution is a free data retrieval call binding the contract method 0x324a4335.
//
// Solidity: function CompressExecution(TypesBurnExecution _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressExecution(_tx TypesBurnExecution) ([]byte, error) {
	return _Rolluputils.Contract.CompressExecution(&_Rolluputils.CallOpts, _tx)
}

// CompressExecution is a free data retrieval call binding the contract method 0x324a4335.
//
// Solidity: function CompressExecution(TypesBurnExecution _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressExecution(_tx TypesBurnExecution) ([]byte, error) {
	return _Rolluputils.Contract.CompressExecution(&_Rolluputils.CallOpts, _tx)
}

// CompressTx is a free data retrieval call binding the contract method 0x02c36512.
//
// Solidity: function CompressTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressTx(opts *bind.CallOpts, _tx TypesTransaction) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressTx", _tx)
	return *ret0, err
}

// CompressTx is a free data retrieval call binding the contract method 0x02c36512.
//
// Solidity: function CompressTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.CompressTx(&_Rolluputils.CallOpts, _tx)
}

// CompressTx is a free data retrieval call binding the contract method 0x02c36512.
//
// Solidity: function CompressTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.CompressTx(&_Rolluputils.CallOpts, _tx)
}

// CompressTxWithMessage is a free data retrieval call binding the contract method 0x6877401d.
//
// Solidity: function CompressTxWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressTxWithMessage(opts *bind.CallOpts, message []byte, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressTxWithMessage", message, sig)
	return *ret0, err
}

// CompressTxWithMessage is a free data retrieval call binding the contract method 0x6877401d.
//
// Solidity: function CompressTxWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressTxWithMessage(message []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressTxWithMessage(&_Rolluputils.CallOpts, message, sig)
}

// CompressTxWithMessage is a free data retrieval call binding the contract method 0x6877401d.
//
// Solidity: function CompressTxWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressTxWithMessage(message []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressTxWithMessage(&_Rolluputils.CallOpts, message, sig)
}

// DecompressConsent is a free data retrieval call binding the contract method 0x34ff6c03.
//
// Solidity: function DecompressConsent(bytes txBytes) pure returns(TypesBurnConsent)
func (_Rolluputils *RolluputilsCaller) DecompressConsent(opts *bind.CallOpts, txBytes []byte) (TypesBurnConsent, error) {
	var (
		ret0 = new(TypesBurnConsent)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressConsent", txBytes)
	return *ret0, err
}

// DecompressConsent is a free data retrieval call binding the contract method 0x34ff6c03.
//
// Solidity: function DecompressConsent(bytes txBytes) pure returns(TypesBurnConsent)
func (_Rolluputils *RolluputilsSession) DecompressConsent(txBytes []byte) (TypesBurnConsent, error) {
	return _Rolluputils.Contract.DecompressConsent(&_Rolluputils.CallOpts, txBytes)
}

// DecompressConsent is a free data retrieval call binding the contract method 0x34ff6c03.
//
// Solidity: function DecompressConsent(bytes txBytes) pure returns(TypesBurnConsent)
func (_Rolluputils *RolluputilsCallerSession) DecompressConsent(txBytes []byte) (TypesBurnConsent, error) {
	return _Rolluputils.Contract.DecompressConsent(&_Rolluputils.CallOpts, txBytes)
}

// DecompressCreateAccount is a free data retrieval call binding the contract method 0x22002f28.
//
// Solidity: function DecompressCreateAccount(bytes txBytes) pure returns(TypesCreateAccount)
func (_Rolluputils *RolluputilsCaller) DecompressCreateAccount(opts *bind.CallOpts, txBytes []byte) (TypesCreateAccount, error) {
	var (
		ret0 = new(TypesCreateAccount)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressCreateAccount", txBytes)
	return *ret0, err
}

// DecompressCreateAccount is a free data retrieval call binding the contract method 0x22002f28.
//
// Solidity: function DecompressCreateAccount(bytes txBytes) pure returns(TypesCreateAccount)
func (_Rolluputils *RolluputilsSession) DecompressCreateAccount(txBytes []byte) (TypesCreateAccount, error) {
	return _Rolluputils.Contract.DecompressCreateAccount(&_Rolluputils.CallOpts, txBytes)
}

// DecompressCreateAccount is a free data retrieval call binding the contract method 0x22002f28.
//
// Solidity: function DecompressCreateAccount(bytes txBytes) pure returns(TypesCreateAccount)
func (_Rolluputils *RolluputilsCallerSession) DecompressCreateAccount(txBytes []byte) (TypesCreateAccount, error) {
	return _Rolluputils.Contract.DecompressCreateAccount(&_Rolluputils.CallOpts, txBytes)
}

// DecompressDrop is a free data retrieval call binding the contract method 0xc8be93c3.
//
// Solidity: function DecompressDrop(bytes dropBytes) pure returns(TypesDropTx)
func (_Rolluputils *RolluputilsCaller) DecompressDrop(opts *bind.CallOpts, dropBytes []byte) (TypesDropTx, error) {
	var (
		ret0 = new(TypesDropTx)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressDrop", dropBytes)
	return *ret0, err
}

// DecompressDrop is a free data retrieval call binding the contract method 0xc8be93c3.
//
// Solidity: function DecompressDrop(bytes dropBytes) pure returns(TypesDropTx)
func (_Rolluputils *RolluputilsSession) DecompressDrop(dropBytes []byte) (TypesDropTx, error) {
	return _Rolluputils.Contract.DecompressDrop(&_Rolluputils.CallOpts, dropBytes)
}

// DecompressDrop is a free data retrieval call binding the contract method 0xc8be93c3.
//
// Solidity: function DecompressDrop(bytes dropBytes) pure returns(TypesDropTx)
func (_Rolluputils *RolluputilsCallerSession) DecompressDrop(dropBytes []byte) (TypesDropTx, error) {
	return _Rolluputils.Contract.DecompressDrop(&_Rolluputils.CallOpts, dropBytes)
}

// DecompressExecution is a free data retrieval call binding the contract method 0x23f05f72.
//
// Solidity: function DecompressExecution(bytes txBytes) pure returns(TypesBurnExecution)
func (_Rolluputils *RolluputilsCaller) DecompressExecution(opts *bind.CallOpts, txBytes []byte) (TypesBurnExecution, error) {
	var (
		ret0 = new(TypesBurnExecution)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressExecution", txBytes)
	return *ret0, err
}

// DecompressExecution is a free data retrieval call binding the contract method 0x23f05f72.
//
// Solidity: function DecompressExecution(bytes txBytes) pure returns(TypesBurnExecution)
func (_Rolluputils *RolluputilsSession) DecompressExecution(txBytes []byte) (TypesBurnExecution, error) {
	return _Rolluputils.Contract.DecompressExecution(&_Rolluputils.CallOpts, txBytes)
}

// DecompressExecution is a free data retrieval call binding the contract method 0x23f05f72.
//
// Solidity: function DecompressExecution(bytes txBytes) pure returns(TypesBurnExecution)
func (_Rolluputils *RolluputilsCallerSession) DecompressExecution(txBytes []byte) (TypesBurnExecution, error) {
	return _Rolluputils.Contract.DecompressExecution(&_Rolluputils.CallOpts, txBytes)
}

// DecompressTx is a free data retrieval call binding the contract method 0xeedeb9d9.
//
// Solidity: function DecompressTx(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 nonce, bytes sig)
func (_Rolluputils *RolluputilsCaller) DecompressTx(opts *bind.CallOpts, txBytes []byte) (struct {
	From  *big.Int
	To    *big.Int
	Nonce *big.Int
	Sig   []byte
}, error) {
	ret := new(struct {
		From  *big.Int
		To    *big.Int
		Nonce *big.Int
		Sig   []byte
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "DecompressTx", txBytes)
	return *ret, err
}

// DecompressTx is a free data retrieval call binding the contract method 0xeedeb9d9.
//
// Solidity: function DecompressTx(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 nonce, bytes sig)
func (_Rolluputils *RolluputilsSession) DecompressTx(txBytes []byte) (struct {
	From  *big.Int
	To    *big.Int
	Nonce *big.Int
	Sig   []byte
}, error) {
	return _Rolluputils.Contract.DecompressTx(&_Rolluputils.CallOpts, txBytes)
}

// DecompressTx is a free data retrieval call binding the contract method 0xeedeb9d9.
//
// Solidity: function DecompressTx(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 nonce, bytes sig)
func (_Rolluputils *RolluputilsCallerSession) DecompressTx(txBytes []byte) (struct {
	From  *big.Int
	To    *big.Int
	Nonce *big.Int
	Sig   []byte
}, error) {
	return _Rolluputils.Contract.DecompressTx(&_Rolluputils.CallOpts, txBytes)
}

// GetGenesisDataBlocks is a free data retrieval call binding the contract method 0x5e31c831.
//
// Solidity: function GetGenesisDataBlocks() pure returns(bytes[2] dataBlocks)
func (_Rolluputils *RolluputilsCaller) GetGenesisDataBlocks(opts *bind.CallOpts) ([2][]byte, error) {
	var (
		ret0 = new([2][]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "GetGenesisDataBlocks")
	return *ret0, err
}

// GetGenesisDataBlocks is a free data retrieval call binding the contract method 0x5e31c831.
//
// Solidity: function GetGenesisDataBlocks() pure returns(bytes[2] dataBlocks)
func (_Rolluputils *RolluputilsSession) GetGenesisDataBlocks() ([2][]byte, error) {
	return _Rolluputils.Contract.GetGenesisDataBlocks(&_Rolluputils.CallOpts)
}

// GetGenesisDataBlocks is a free data retrieval call binding the contract method 0x5e31c831.
//
// Solidity: function GetGenesisDataBlocks() pure returns(bytes[2] dataBlocks)
func (_Rolluputils *RolluputilsCallerSession) GetGenesisDataBlocks() ([2][]byte, error) {
	return _Rolluputils.Contract.GetGenesisDataBlocks(&_Rolluputils.CallOpts)
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

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(TypesUserAccount account) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromAccount(opts *bind.CallOpts, account TypesUserAccount) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromAccount", account)
	return *ret0, err
}

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(TypesUserAccount account) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromAccount(account TypesUserAccount) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromAccount(&_Rolluputils.CallOpts, account)
}

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(TypesUserAccount account) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromAccount(account TypesUserAccount) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromAccount(&_Rolluputils.CallOpts, account)
}

// HashFromConsent is a free data retrieval call binding the contract method 0xbe8d1339.
//
// Solidity: function HashFromConsent(TypesBurnConsent _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromConsent(opts *bind.CallOpts, _tx TypesBurnConsent) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromConsent", _tx)
	return *ret0, err
}

// HashFromConsent is a free data retrieval call binding the contract method 0xbe8d1339.
//
// Solidity: function HashFromConsent(TypesBurnConsent _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromConsent(_tx TypesBurnConsent) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromConsent(&_Rolluputils.CallOpts, _tx)
}

// HashFromConsent is a free data retrieval call binding the contract method 0xbe8d1339.
//
// Solidity: function HashFromConsent(TypesBurnConsent _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromConsent(_tx TypesBurnConsent) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromConsent(&_Rolluputils.CallOpts, _tx)
}

// HashFromExecution is a free data retrieval call binding the contract method 0x8c3d042c.
//
// Solidity: function HashFromExecution(TypesBurnExecution _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromExecution(opts *bind.CallOpts, _tx TypesBurnExecution) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromExecution", _tx)
	return *ret0, err
}

// HashFromExecution is a free data retrieval call binding the contract method 0x8c3d042c.
//
// Solidity: function HashFromExecution(TypesBurnExecution _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromExecution(_tx TypesBurnExecution) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromExecution(&_Rolluputils.CallOpts, _tx)
}

// HashFromExecution is a free data retrieval call binding the contract method 0x8c3d042c.
//
// Solidity: function HashFromExecution(TypesBurnExecution _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromExecution(_tx TypesBurnExecution) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromExecution(&_Rolluputils.CallOpts, _tx)
}

// HashFromTx is a free data retrieval call binding the contract method 0xb90cbf51.
//
// Solidity: function HashFromTx(TypesTransaction _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromTx(opts *bind.CallOpts, _tx TypesTransaction) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromTx", _tx)
	return *ret0, err
}

// HashFromTx is a free data retrieval call binding the contract method 0xb90cbf51.
//
// Solidity: function HashFromTx(TypesTransaction _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromTx(_tx TypesTransaction) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromTx(&_Rolluputils.CallOpts, _tx)
}

// HashFromTx is a free data retrieval call binding the contract method 0xb90cbf51.
//
// Solidity: function HashFromTx(TypesTransaction _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromTx(_tx TypesTransaction) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromTx(&_Rolluputils.CallOpts, _tx)
}

// PDALeafToHash is a free data retrieval call binding the contract method 0xc2ddab33.
//
// Solidity: function PDALeafToHash(TypesPDALeaf _PDA_Leaf) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) PDALeafToHash(opts *bind.CallOpts, _PDA_Leaf TypesPDALeaf) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "PDALeafToHash", _PDA_Leaf)
	return *ret0, err
}

// PDALeafToHash is a free data retrieval call binding the contract method 0xc2ddab33.
//
// Solidity: function PDALeafToHash(TypesPDALeaf _PDA_Leaf) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) PDALeafToHash(_PDA_Leaf TypesPDALeaf) ([32]byte, error) {
	return _Rolluputils.Contract.PDALeafToHash(&_Rolluputils.CallOpts, _PDA_Leaf)
}

// PDALeafToHash is a free data retrieval call binding the contract method 0xc2ddab33.
//
// Solidity: function PDALeafToHash(TypesPDALeaf _PDA_Leaf) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) PDALeafToHash(_PDA_Leaf TypesPDALeaf) ([32]byte, error) {
	return _Rolluputils.Contract.PDALeafToHash(&_Rolluputils.CallOpts, _PDA_Leaf)
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns(TypesTransaction)
func (_Rolluputils *RolluputilsCaller) TxFromBytes(opts *bind.CallOpts, txBytes []byte) (TypesTransaction, error) {
	var (
		ret0 = new(TypesTransaction)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "TxFromBytes", txBytes)
	return *ret0, err
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns(TypesTransaction)
func (_Rolluputils *RolluputilsSession) TxFromBytes(txBytes []byte) (TypesTransaction, error) {
	return _Rolluputils.Contract.TxFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns(TypesTransaction)
func (_Rolluputils *RolluputilsCallerSession) TxFromBytes(txBytes []byte) (TypesTransaction, error) {
	return _Rolluputils.Contract.TxFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytesBurnConsent is a free data retrieval call binding the contract method 0xa8c42f1e.
//
// Solidity: function TxFromBytesBurnConsent(bytes txBytes) pure returns(TypesBurnConsent)
func (_Rolluputils *RolluputilsCaller) TxFromBytesBurnConsent(opts *bind.CallOpts, txBytes []byte) (TypesBurnConsent, error) {
	var (
		ret0 = new(TypesBurnConsent)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "TxFromBytesBurnConsent", txBytes)
	return *ret0, err
}

// TxFromBytesBurnConsent is a free data retrieval call binding the contract method 0xa8c42f1e.
//
// Solidity: function TxFromBytesBurnConsent(bytes txBytes) pure returns(TypesBurnConsent)
func (_Rolluputils *RolluputilsSession) TxFromBytesBurnConsent(txBytes []byte) (TypesBurnConsent, error) {
	return _Rolluputils.Contract.TxFromBytesBurnConsent(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytesBurnConsent is a free data retrieval call binding the contract method 0xa8c42f1e.
//
// Solidity: function TxFromBytesBurnConsent(bytes txBytes) pure returns(TypesBurnConsent)
func (_Rolluputils *RolluputilsCallerSession) TxFromBytesBurnConsent(txBytes []byte) (TypesBurnConsent, error) {
	return _Rolluputils.Contract.TxFromBytesBurnConsent(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytesBurnExecution is a free data retrieval call binding the contract method 0x0c83d4bc.
//
// Solidity: function TxFromBytesBurnExecution(bytes txBytes) pure returns(TypesBurnExecution)
func (_Rolluputils *RolluputilsCaller) TxFromBytesBurnExecution(opts *bind.CallOpts, txBytes []byte) (TypesBurnExecution, error) {
	var (
		ret0 = new(TypesBurnExecution)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "TxFromBytesBurnExecution", txBytes)
	return *ret0, err
}

// TxFromBytesBurnExecution is a free data retrieval call binding the contract method 0x0c83d4bc.
//
// Solidity: function TxFromBytesBurnExecution(bytes txBytes) pure returns(TypesBurnExecution)
func (_Rolluputils *RolluputilsSession) TxFromBytesBurnExecution(txBytes []byte) (TypesBurnExecution, error) {
	return _Rolluputils.Contract.TxFromBytesBurnExecution(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytesBurnExecution is a free data retrieval call binding the contract method 0x0c83d4bc.
//
// Solidity: function TxFromBytesBurnExecution(bytes txBytes) pure returns(TypesBurnExecution)
func (_Rolluputils *RolluputilsCallerSession) TxFromBytesBurnExecution(txBytes []byte) (TypesBurnExecution, error) {
	return _Rolluputils.Contract.TxFromBytesBurnExecution(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytesDeconstructed is a free data retrieval call binding the contract method 0x2013a0cf.
//
// Solidity: function TxFromBytesDeconstructed(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount)
func (_Rolluputils *RolluputilsCaller) TxFromBytesDeconstructed(opts *bind.CallOpts, txBytes []byte) (struct {
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
}, error) {
	ret := new(struct {
		From      *big.Int
		To        *big.Int
		TokenType *big.Int
		Nonce     *big.Int
		TxType    *big.Int
		Amount    *big.Int
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "TxFromBytesDeconstructed", txBytes)
	return *ret, err
}

// TxFromBytesDeconstructed is a free data retrieval call binding the contract method 0x2013a0cf.
//
// Solidity: function TxFromBytesDeconstructed(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount)
func (_Rolluputils *RolluputilsSession) TxFromBytesDeconstructed(txBytes []byte) (struct {
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
}, error) {
	return _Rolluputils.Contract.TxFromBytesDeconstructed(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytesDeconstructed is a free data retrieval call binding the contract method 0x2013a0cf.
//
// Solidity: function TxFromBytesDeconstructed(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount)
func (_Rolluputils *RolluputilsCallerSession) TxFromBytesDeconstructed(txBytes []byte) (struct {
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
}, error) {
	return _Rolluputils.Contract.TxFromBytesDeconstructed(&_Rolluputils.CallOpts, txBytes)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) pure returns(address addr)
func (_Rolluputils *RolluputilsCaller) CalculateAddress(opts *bind.CallOpts, pub []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "calculateAddress", pub)
	return *ret0, err
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) pure returns(address addr)
func (_Rolluputils *RolluputilsSession) CalculateAddress(pub []byte) (common.Address, error) {
	return _Rolluputils.Contract.CalculateAddress(&_Rolluputils.CallOpts, pub)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) pure returns(address addr)
func (_Rolluputils *RolluputilsCallerSession) CalculateAddress(pub []byte) (common.Address, error) {
	return _Rolluputils.Contract.CalculateAddress(&_Rolluputils.CallOpts, pub)
}

// GetAccountHash is a free data retrieval call binding the contract method 0x61a8c3c2.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) GetAccountHash(opts *bind.CallOpts, id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getAccountHash", id, balance, nonce, tokenType)
	return *ret0, err
}

// GetAccountHash is a free data retrieval call binding the contract method 0x61a8c3c2.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) GetAccountHash(id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetAccountHash(&_Rolluputils.CallOpts, id, balance, nonce, tokenType)
}

// GetAccountHash is a free data retrieval call binding the contract method 0x61a8c3c2.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) GetAccountHash(id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetAccountHash(&_Rolluputils.CallOpts, id, balance, nonce, tokenType)
}

// GetDropSignBytes is a free data retrieval call binding the contract method 0x4a5e70d5.
//
// Solidity: function getDropSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) GetDropSignBytes(opts *bind.CallOpts, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getDropSignBytes", fromIndex, toIndex, tokenType, txType, nonce, amount)
	return *ret0, err
}

// GetDropSignBytes is a free data retrieval call binding the contract method 0x4a5e70d5.
//
// Solidity: function getDropSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) GetDropSignBytes(fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetDropSignBytes(&_Rolluputils.CallOpts, fromIndex, toIndex, tokenType, txType, nonce, amount)
}

// GetDropSignBytes is a free data retrieval call binding the contract method 0x4a5e70d5.
//
// Solidity: function getDropSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) GetDropSignBytes(fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetDropSignBytes(&_Rolluputils.CallOpts, fromIndex, toIndex, tokenType, txType, nonce, amount)
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) GetTxSignBytes(opts *bind.CallOpts, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getTxSignBytes", fromIndex, toIndex, tokenType, txType, nonce, amount)
	return *ret0, err
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) GetTxSignBytes(fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes(&_Rolluputils.CallOpts, fromIndex, toIndex, tokenType, txType, nonce, amount)
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) GetTxSignBytes(fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes(&_Rolluputils.CallOpts, fromIndex, toIndex, tokenType, txType, nonce, amount)
}
