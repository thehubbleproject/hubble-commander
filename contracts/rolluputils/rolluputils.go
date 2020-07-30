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

// TxBurnConsent is an auto generated low-level Go binding around an user-defined struct.
type TxBurnConsent struct {
	FromIndex *big.Int
	Amount    *big.Int
	Signature []byte
}

// TxBurnExecution is an auto generated low-level Go binding around an user-defined struct.
type TxBurnExecution struct {
	FromIndex *big.Int
}

// TxCreateAccount is an auto generated low-level Go binding around an user-defined struct.
type TxCreateAccount struct {
	AccountID *big.Int
	StateID   *big.Int
	TokenType *big.Int
}

// TxTransfer is an auto generated low-level Go binding around an user-defined struct.
type TxTransfer struct {
	FromIndex *big.Int
	ToIndex   *big.Int
	Amount    *big.Int
	Signature []byte
}

// TypesBurnConsent is an auto generated low-level Go binding around an user-defined struct.
type TypesBurnConsent struct {
	TxType    *big.Int
	FromIndex *big.Int
	Amount    *big.Int
	Nonce     *big.Int
	Signature []byte
}

// TypesBurnExecution is an auto generated low-level Go binding around an user-defined struct.
type TypesBurnExecution struct {
	TxType    *big.Int
	FromIndex *big.Int
}

// TypesCreateAccount is an auto generated low-level Go binding around an user-defined struct.
type TypesCreateAccount struct {
	TxType    *big.Int
	AccountID *big.Int
	StateID   *big.Int
	TokenType *big.Int
}

// TypesDropTx is an auto generated low-level Go binding around an user-defined struct.
type TypesDropTx struct {
	TxType    *big.Int
	FromIndex *big.Int
	ToIndex   *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	Amount    *big.Int
	Signature []byte
}

// TypesPDALeaf is an auto generated low-level Go binding around an user-defined struct.
type TypesPDALeaf struct {
	Pubkey []byte
}

// TypesTransaction is an auto generated low-level Go binding around an user-defined struct.
type TypesTransaction struct {
	TxType    *big.Int
	FromIndex *big.Int
	ToIndex   *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	Amount    *big.Int
	Signature []byte
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

// RolluputilsABI is the input ABI used to generate the binding from.
const RolluputilsABI = "[{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"_PDA_Leaf\",\"type\":\"tuple\"}],\"name\":\"PDALeafToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"accountBytes\",\"type\":\"bytes\"}],\"name\":\"AccountFromBytes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"BytesFromAccount\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"name\":\"BytesFromAccountDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"name\":\"getAccountHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurn\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"HashFromAccount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pub\",\"type\":\"bytes\"}],\"name\":\"calculateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetGenesisLeaves\",\"outputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"leaves\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetGenesisDataBlocks\",\"outputs\":[{\"internalType\":\"bytes[2]\",\"name\":\"dataBlocks\",\"type\":\"bytes[2]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.CreateAccount\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromCreateAccount\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"name\":\"BytesFromCreateAccountNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"CreateAccountFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.CreateAccount\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.CreateAccount\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressCreateAccount\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accountID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"name\":\"CompressCreateAccountNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"CompressCreateAccountWithMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressCreateAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accountID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"CompressCreateAccountFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"txBytes\",\"type\":\"bytes[]\"}],\"name\":\"CompressManyCreateAccountFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"DecompressManyCreateAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accountID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"internalType\":\"structTx.CreateAccount[]\",\"name\":\"structTxs\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.DropTx\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromAirdrop\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BytesFromAirdropNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"AirdropFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.DropTx\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"AirdropFromBytesNoStruct\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AirdropSignBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.DropTx\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressAirdrop\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressAirdropNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressAirdropTxWithMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressAirdrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressAirdropFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"txBytes\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"CompressManyAirdropFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"DecompressManyAirdrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTx.Transfer[]\",\"name\":\"structTxs\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BytesFromTxDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"TxFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"TxFromBytesDeconstructed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getTxSignBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressTxWithMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"DecompressTransfers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTx.Transfer[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressTransferFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"txBytes\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"CompressManyTransferFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"DecompressManyTransfer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTx.Transfer[]\",\"name\":\"structTxs\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromBurnConsent\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BytesFromBurnConsentNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"BurnConsentFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BurnConsentSignBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressBurnConsent\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressBurnConsentNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressBurnConsentWithMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressBurnConsent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BurnConsent\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromBurnConsent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CompressBurnConsentFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"txBytes\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"CompressManyBurnConsentFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"DecompressManyBurnConsent\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTx.BurnConsent[]\",\"name\":\"structTxs\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromBurnExecution\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"name\":\"BytesFromBurnExecutionNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"BurnExecutionFromBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"name\":\"BurnExecutionSignBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressBurnExecution\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"CompressBurnExecutionWithMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"name\":\"CompressBurnExecutionNoStruct\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"DecompressBurnExecution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BurnExecution\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromBurnExecution\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"CompressBurnExecutionFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"txBytes\",\"type\":\"bytes[]\"}],\"name\":\"CompressManyBurnExecutionFromEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"DecompressManyBurnExecution\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"}],\"internalType\":\"structTx.BurnExecution[]\",\"name\":\"structTxs\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetYearMonth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"yearMonth\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn)
func (_Rolluputils *RolluputilsCaller) AccountFromBytes(opts *bind.CallOpts, accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
	Burn      *big.Int
	LastBurn  *big.Int
}, error) {
	ret := new(struct {
		ID        *big.Int
		Balance   *big.Int
		Nonce     *big.Int
		TokenType *big.Int
		Burn      *big.Int
		LastBurn  *big.Int
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "AccountFromBytes", accountBytes)
	return *ret, err
}

// AccountFromBytes is a free data retrieval call binding the contract method 0x1a636e86.
//
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn)
func (_Rolluputils *RolluputilsSession) AccountFromBytes(accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
	Burn      *big.Int
	LastBurn  *big.Int
}, error) {
	return _Rolluputils.Contract.AccountFromBytes(&_Rolluputils.CallOpts, accountBytes)
}

// AccountFromBytes is a free data retrieval call binding the contract method 0x1a636e86.
//
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn)
func (_Rolluputils *RolluputilsCallerSession) AccountFromBytes(accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
	Burn      *big.Int
	LastBurn  *big.Int
}, error) {
	return _Rolluputils.Contract.AccountFromBytes(&_Rolluputils.CallOpts, accountBytes)
}

// AirdropFromBytes is a free data retrieval call binding the contract method 0xbfecfbf3.
//
// Solidity: function AirdropFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes))
func (_Rolluputils *RolluputilsCaller) AirdropFromBytes(opts *bind.CallOpts, txBytes []byte) (TypesDropTx, error) {
	var (
		ret0 = new(TypesDropTx)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "AirdropFromBytes", txBytes)
	return *ret0, err
}

// AirdropFromBytes is a free data retrieval call binding the contract method 0xbfecfbf3.
//
// Solidity: function AirdropFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes))
func (_Rolluputils *RolluputilsSession) AirdropFromBytes(txBytes []byte) (TypesDropTx, error) {
	return _Rolluputils.Contract.AirdropFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// AirdropFromBytes is a free data retrieval call binding the contract method 0xbfecfbf3.
//
// Solidity: function AirdropFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes))
func (_Rolluputils *RolluputilsCallerSession) AirdropFromBytes(txBytes []byte) (TypesDropTx, error) {
	return _Rolluputils.Contract.AirdropFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// AirdropFromBytesNoStruct is a free data retrieval call binding the contract method 0xb34dc92b.
//
// Solidity: function AirdropFromBytesNoStruct(bytes txBytes) pure returns(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount)
func (_Rolluputils *RolluputilsCaller) AirdropFromBytesNoStruct(opts *bind.CallOpts, txBytes []byte) (struct {
	TxType    *big.Int
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	Amount    *big.Int
}, error) {
	ret := new(struct {
		TxType    *big.Int
		From      *big.Int
		To        *big.Int
		TokenType *big.Int
		Nonce     *big.Int
		Amount    *big.Int
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "AirdropFromBytesNoStruct", txBytes)
	return *ret, err
}

// AirdropFromBytesNoStruct is a free data retrieval call binding the contract method 0xb34dc92b.
//
// Solidity: function AirdropFromBytesNoStruct(bytes txBytes) pure returns(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount)
func (_Rolluputils *RolluputilsSession) AirdropFromBytesNoStruct(txBytes []byte) (struct {
	TxType    *big.Int
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	Amount    *big.Int
}, error) {
	return _Rolluputils.Contract.AirdropFromBytesNoStruct(&_Rolluputils.CallOpts, txBytes)
}

// AirdropFromBytesNoStruct is a free data retrieval call binding the contract method 0xb34dc92b.
//
// Solidity: function AirdropFromBytesNoStruct(bytes txBytes) pure returns(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount)
func (_Rolluputils *RolluputilsCallerSession) AirdropFromBytesNoStruct(txBytes []byte) (struct {
	TxType    *big.Int
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	Amount    *big.Int
}, error) {
	return _Rolluputils.Contract.AirdropFromBytesNoStruct(&_Rolluputils.CallOpts, txBytes)
}

// AirdropSignBytes is a free data retrieval call binding the contract method 0x73f43f31.
//
// Solidity: function AirdropSignBytes(uint256 txType, uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) AirdropSignBytes(opts *bind.CallOpts, txType *big.Int, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "AirdropSignBytes", txType, fromIndex, toIndex, tokenType, nonce, amount)
	return *ret0, err
}

// AirdropSignBytes is a free data retrieval call binding the contract method 0x73f43f31.
//
// Solidity: function AirdropSignBytes(uint256 txType, uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) AirdropSignBytes(txType *big.Int, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.AirdropSignBytes(&_Rolluputils.CallOpts, txType, fromIndex, toIndex, tokenType, nonce, amount)
}

// AirdropSignBytes is a free data retrieval call binding the contract method 0x73f43f31.
//
// Solidity: function AirdropSignBytes(uint256 txType, uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) AirdropSignBytes(txType *big.Int, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.AirdropSignBytes(&_Rolluputils.CallOpts, txType, fromIndex, toIndex, tokenType, nonce, amount)
}

// BurnConsentFromBytes is a free data retrieval call binding the contract method 0x6811105b.
//
// Solidity: function BurnConsentFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,bytes))
func (_Rolluputils *RolluputilsCaller) BurnConsentFromBytes(opts *bind.CallOpts, txBytes []byte) (TypesBurnConsent, error) {
	var (
		ret0 = new(TypesBurnConsent)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BurnConsentFromBytes", txBytes)
	return *ret0, err
}

// BurnConsentFromBytes is a free data retrieval call binding the contract method 0x6811105b.
//
// Solidity: function BurnConsentFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,bytes))
func (_Rolluputils *RolluputilsSession) BurnConsentFromBytes(txBytes []byte) (TypesBurnConsent, error) {
	return _Rolluputils.Contract.BurnConsentFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// BurnConsentFromBytes is a free data retrieval call binding the contract method 0x6811105b.
//
// Solidity: function BurnConsentFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,bytes))
func (_Rolluputils *RolluputilsCallerSession) BurnConsentFromBytes(txBytes []byte) (TypesBurnConsent, error) {
	return _Rolluputils.Contract.BurnConsentFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// BurnConsentSignBytes is a free data retrieval call binding the contract method 0xb58a2f79.
//
// Solidity: function BurnConsentSignBytes(uint256 txType, uint256 fromIndex, uint256 amount, uint256 nonce) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) BurnConsentSignBytes(opts *bind.CallOpts, txType *big.Int, fromIndex *big.Int, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BurnConsentSignBytes", txType, fromIndex, amount, nonce)
	return *ret0, err
}

// BurnConsentSignBytes is a free data retrieval call binding the contract method 0xb58a2f79.
//
// Solidity: function BurnConsentSignBytes(uint256 txType, uint256 fromIndex, uint256 amount, uint256 nonce) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) BurnConsentSignBytes(txType *big.Int, fromIndex *big.Int, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.BurnConsentSignBytes(&_Rolluputils.CallOpts, txType, fromIndex, amount, nonce)
}

// BurnConsentSignBytes is a free data retrieval call binding the contract method 0xb58a2f79.
//
// Solidity: function BurnConsentSignBytes(uint256 txType, uint256 fromIndex, uint256 amount, uint256 nonce) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) BurnConsentSignBytes(txType *big.Int, fromIndex *big.Int, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.BurnConsentSignBytes(&_Rolluputils.CallOpts, txType, fromIndex, amount, nonce)
}

// BurnExecutionFromBytes is a free data retrieval call binding the contract method 0x7b1849fa.
//
// Solidity: function BurnExecutionFromBytes(bytes txBytes) pure returns((uint256,uint256))
func (_Rolluputils *RolluputilsCaller) BurnExecutionFromBytes(opts *bind.CallOpts, txBytes []byte) (TypesBurnExecution, error) {
	var (
		ret0 = new(TypesBurnExecution)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BurnExecutionFromBytes", txBytes)
	return *ret0, err
}

// BurnExecutionFromBytes is a free data retrieval call binding the contract method 0x7b1849fa.
//
// Solidity: function BurnExecutionFromBytes(bytes txBytes) pure returns((uint256,uint256))
func (_Rolluputils *RolluputilsSession) BurnExecutionFromBytes(txBytes []byte) (TypesBurnExecution, error) {
	return _Rolluputils.Contract.BurnExecutionFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// BurnExecutionFromBytes is a free data retrieval call binding the contract method 0x7b1849fa.
//
// Solidity: function BurnExecutionFromBytes(bytes txBytes) pure returns((uint256,uint256))
func (_Rolluputils *RolluputilsCallerSession) BurnExecutionFromBytes(txBytes []byte) (TypesBurnExecution, error) {
	return _Rolluputils.Contract.BurnExecutionFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// BurnExecutionSignBytes is a free data retrieval call binding the contract method 0xe79c804b.
//
// Solidity: function BurnExecutionSignBytes(uint256 txType, uint256 fromIndex) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) BurnExecutionSignBytes(opts *bind.CallOpts, txType *big.Int, fromIndex *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BurnExecutionSignBytes", txType, fromIndex)
	return *ret0, err
}

// BurnExecutionSignBytes is a free data retrieval call binding the contract method 0xe79c804b.
//
// Solidity: function BurnExecutionSignBytes(uint256 txType, uint256 fromIndex) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) BurnExecutionSignBytes(txType *big.Int, fromIndex *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.BurnExecutionSignBytes(&_Rolluputils.CallOpts, txType, fromIndex)
}

// BurnExecutionSignBytes is a free data retrieval call binding the contract method 0xe79c804b.
//
// Solidity: function BurnExecutionSignBytes(uint256 txType, uint256 fromIndex) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) BurnExecutionSignBytes(txType *big.Int, fromIndex *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.BurnExecutionSignBytes(&_Rolluputils.CallOpts, txType, fromIndex)
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x04af2a83.
//
// Solidity: function BytesFromAccount((uint256,uint256,uint256,uint256,uint256,uint256) account) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromAccount(opts *bind.CallOpts, account TypesUserAccount) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromAccount", account)
	return *ret0, err
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x04af2a83.
//
// Solidity: function BytesFromAccount((uint256,uint256,uint256,uint256,uint256,uint256) account) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAccount(account TypesUserAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccount(&_Rolluputils.CallOpts, account)
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x04af2a83.
//
// Solidity: function BytesFromAccount((uint256,uint256,uint256,uint256,uint256,uint256) account) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAccount(account TypesUserAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccount(&_Rolluputils.CallOpts, account)
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0xa9ef7f3a.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromAccountDeconstructed(opts *bind.CallOpts, ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int, burn *big.Int, lastBurn *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromAccountDeconstructed", ID, balance, nonce, tokenType, burn, lastBurn)
	return *ret0, err
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0xa9ef7f3a.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAccountDeconstructed(ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int, burn *big.Int, lastBurn *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccountDeconstructed(&_Rolluputils.CallOpts, ID, balance, nonce, tokenType, burn, lastBurn)
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0xa9ef7f3a.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAccountDeconstructed(ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int, burn *big.Int, lastBurn *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccountDeconstructed(&_Rolluputils.CallOpts, ID, balance, nonce, tokenType, burn, lastBurn)
}

// BytesFromAirdrop is a free data retrieval call binding the contract method 0x29d87dd9.
//
// Solidity: function BytesFromAirdrop((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
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
// Solidity: function BytesFromAirdrop((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAirdrop(_tx TypesDropTx) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAirdrop(&_Rolluputils.CallOpts, _tx)
}

// BytesFromAirdrop is a free data retrieval call binding the contract method 0x29d87dd9.
//
// Solidity: function BytesFromAirdrop((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAirdrop(_tx TypesDropTx) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAirdrop(&_Rolluputils.CallOpts, _tx)
}

// BytesFromAirdropNoStruct is a free data retrieval call binding the contract method 0x7368b4a0.
//
// Solidity: function BytesFromAirdropNoStruct(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromAirdropNoStruct(opts *bind.CallOpts, txType *big.Int, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromAirdropNoStruct", txType, from, to, tokenType, nonce, amount)
	return *ret0, err
}

// BytesFromAirdropNoStruct is a free data retrieval call binding the contract method 0x7368b4a0.
//
// Solidity: function BytesFromAirdropNoStruct(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAirdropNoStruct(txType *big.Int, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAirdropNoStruct(&_Rolluputils.CallOpts, txType, from, to, tokenType, nonce, amount)
}

// BytesFromAirdropNoStruct is a free data retrieval call binding the contract method 0x7368b4a0.
//
// Solidity: function BytesFromAirdropNoStruct(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAirdropNoStruct(txType *big.Int, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAirdropNoStruct(&_Rolluputils.CallOpts, txType, from, to, tokenType, nonce, amount)
}

// BytesFromBurnConsent is a free data retrieval call binding the contract method 0x9b4bf464.
//
// Solidity: function BytesFromBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromBurnConsent(opts *bind.CallOpts, _tx TypesBurnConsent) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromBurnConsent", _tx)
	return *ret0, err
}

// BytesFromBurnConsent is a free data retrieval call binding the contract method 0x9b4bf464.
//
// Solidity: function BytesFromBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromBurnConsent(_tx TypesBurnConsent) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnConsent(&_Rolluputils.CallOpts, _tx)
}

// BytesFromBurnConsent is a free data retrieval call binding the contract method 0x9b4bf464.
//
// Solidity: function BytesFromBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromBurnConsent(_tx TypesBurnConsent) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnConsent(&_Rolluputils.CallOpts, _tx)
}

// BytesFromBurnConsentNoStruct is a free data retrieval call binding the contract method 0x722ce51a.
//
// Solidity: function BytesFromBurnConsentNoStruct(uint256 txType, uint256 fromIndex, uint256 amount, uint256 nonce) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromBurnConsentNoStruct(opts *bind.CallOpts, txType *big.Int, fromIndex *big.Int, amount *big.Int, nonce *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromBurnConsentNoStruct", txType, fromIndex, amount, nonce)
	return *ret0, err
}

// BytesFromBurnConsentNoStruct is a free data retrieval call binding the contract method 0x722ce51a.
//
// Solidity: function BytesFromBurnConsentNoStruct(uint256 txType, uint256 fromIndex, uint256 amount, uint256 nonce) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromBurnConsentNoStruct(txType *big.Int, fromIndex *big.Int, amount *big.Int, nonce *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnConsentNoStruct(&_Rolluputils.CallOpts, txType, fromIndex, amount, nonce)
}

// BytesFromBurnConsentNoStruct is a free data retrieval call binding the contract method 0x722ce51a.
//
// Solidity: function BytesFromBurnConsentNoStruct(uint256 txType, uint256 fromIndex, uint256 amount, uint256 nonce) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromBurnConsentNoStruct(txType *big.Int, fromIndex *big.Int, amount *big.Int, nonce *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnConsentNoStruct(&_Rolluputils.CallOpts, txType, fromIndex, amount, nonce)
}

// BytesFromBurnExecution is a free data retrieval call binding the contract method 0xd7199f92.
//
// Solidity: function BytesFromBurnExecution((uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromBurnExecution(opts *bind.CallOpts, _tx TypesBurnExecution) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromBurnExecution", _tx)
	return *ret0, err
}

// BytesFromBurnExecution is a free data retrieval call binding the contract method 0xd7199f92.
//
// Solidity: function BytesFromBurnExecution((uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromBurnExecution(_tx TypesBurnExecution) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnExecution(&_Rolluputils.CallOpts, _tx)
}

// BytesFromBurnExecution is a free data retrieval call binding the contract method 0xd7199f92.
//
// Solidity: function BytesFromBurnExecution((uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromBurnExecution(_tx TypesBurnExecution) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnExecution(&_Rolluputils.CallOpts, _tx)
}

// BytesFromBurnExecutionNoStruct is a free data retrieval call binding the contract method 0x583d218f.
//
// Solidity: function BytesFromBurnExecutionNoStruct(uint256 txType, uint256 fromIndex) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromBurnExecutionNoStruct(opts *bind.CallOpts, txType *big.Int, fromIndex *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromBurnExecutionNoStruct", txType, fromIndex)
	return *ret0, err
}

// BytesFromBurnExecutionNoStruct is a free data retrieval call binding the contract method 0x583d218f.
//
// Solidity: function BytesFromBurnExecutionNoStruct(uint256 txType, uint256 fromIndex) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromBurnExecutionNoStruct(txType *big.Int, fromIndex *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnExecutionNoStruct(&_Rolluputils.CallOpts, txType, fromIndex)
}

// BytesFromBurnExecutionNoStruct is a free data retrieval call binding the contract method 0x583d218f.
//
// Solidity: function BytesFromBurnExecutionNoStruct(uint256 txType, uint256 fromIndex) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromBurnExecutionNoStruct(txType *big.Int, fromIndex *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromBurnExecutionNoStruct(&_Rolluputils.CallOpts, txType, fromIndex)
}

// BytesFromCreateAccount is a free data retrieval call binding the contract method 0x89fd704b.
//
// Solidity: function BytesFromCreateAccount((uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromCreateAccount(opts *bind.CallOpts, _tx TypesCreateAccount) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromCreateAccount", _tx)
	return *ret0, err
}

// BytesFromCreateAccount is a free data retrieval call binding the contract method 0x89fd704b.
//
// Solidity: function BytesFromCreateAccount((uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromCreateAccount(_tx TypesCreateAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromCreateAccount(&_Rolluputils.CallOpts, _tx)
}

// BytesFromCreateAccount is a free data retrieval call binding the contract method 0x89fd704b.
//
// Solidity: function BytesFromCreateAccount((uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromCreateAccount(_tx TypesCreateAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromCreateAccount(&_Rolluputils.CallOpts, _tx)
}

// BytesFromCreateAccountNoStruct is a free data retrieval call binding the contract method 0x0a720c49.
//
// Solidity: function BytesFromCreateAccountNoStruct(uint256 txType, uint256 accountID, uint256 stateID, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromCreateAccountNoStruct(opts *bind.CallOpts, txType *big.Int, accountID *big.Int, stateID *big.Int, tokenType *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromCreateAccountNoStruct", txType, accountID, stateID, tokenType)
	return *ret0, err
}

// BytesFromCreateAccountNoStruct is a free data retrieval call binding the contract method 0x0a720c49.
//
// Solidity: function BytesFromCreateAccountNoStruct(uint256 txType, uint256 accountID, uint256 stateID, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromCreateAccountNoStruct(txType *big.Int, accountID *big.Int, stateID *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromCreateAccountNoStruct(&_Rolluputils.CallOpts, txType, accountID, stateID, tokenType)
}

// BytesFromCreateAccountNoStruct is a free data retrieval call binding the contract method 0x0a720c49.
//
// Solidity: function BytesFromCreateAccountNoStruct(uint256 txType, uint256 accountID, uint256 stateID, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromCreateAccountNoStruct(txType *big.Int, accountID *big.Int, stateID *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromCreateAccountNoStruct(&_Rolluputils.CallOpts, txType, accountID, stateID, tokenType)
}

// BytesFromTx is a free data retrieval call binding the contract method 0xd0dbc902.
//
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
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
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTx is a free data retrieval call binding the contract method 0xd0dbc902.
//
// Solidity: function BytesFromTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTxDeconstructed(opts *bind.CallOpts, txType *big.Int, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTxDeconstructed", txType, from, to, tokenType, nonce, amount)
	return *ret0, err
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTxDeconstructed(txType *big.Int, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxDeconstructed(&_Rolluputils.CallOpts, txType, from, to, tokenType, nonce, amount)
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTxDeconstructed(txType *big.Int, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxDeconstructed(&_Rolluputils.CallOpts, txType, from, to, tokenType, nonce, amount)
}

// CompressAirdrop is a free data retrieval call binding the contract method 0x296d2f03.
//
// Solidity: function CompressAirdrop((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressAirdrop(opts *bind.CallOpts, _tx TypesDropTx) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressAirdrop", _tx)
	return *ret0, err
}

// CompressAirdrop is a free data retrieval call binding the contract method 0x296d2f03.
//
// Solidity: function CompressAirdrop((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressAirdrop(_tx TypesDropTx) ([]byte, error) {
	return _Rolluputils.Contract.CompressAirdrop(&_Rolluputils.CallOpts, _tx)
}

// CompressAirdrop is a free data retrieval call binding the contract method 0x296d2f03.
//
// Solidity: function CompressAirdrop((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressAirdrop(_tx TypesDropTx) ([]byte, error) {
	return _Rolluputils.Contract.CompressAirdrop(&_Rolluputils.CallOpts, _tx)
}

// CompressAirdropFromEncoded is a free data retrieval call binding the contract method 0x175284f7.
//
// Solidity: function CompressAirdropFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressAirdropFromEncoded(opts *bind.CallOpts, txBytes []byte, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressAirdropFromEncoded", txBytes, sig)
	return *ret0, err
}

// CompressAirdropFromEncoded is a free data retrieval call binding the contract method 0x175284f7.
//
// Solidity: function CompressAirdropFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressAirdropFromEncoded(txBytes []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressAirdropFromEncoded(&_Rolluputils.CallOpts, txBytes, sig)
}

// CompressAirdropFromEncoded is a free data retrieval call binding the contract method 0x175284f7.
//
// Solidity: function CompressAirdropFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressAirdropFromEncoded(txBytes []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressAirdropFromEncoded(&_Rolluputils.CallOpts, txBytes, sig)
}

// CompressAirdropNoStruct is a free data retrieval call binding the contract method 0x1345dc05.
//
// Solidity: function CompressAirdropNoStruct(uint256 toIndex, uint256 amount, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressAirdropNoStruct(opts *bind.CallOpts, toIndex *big.Int, amount *big.Int, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressAirdropNoStruct", toIndex, amount, sig)
	return *ret0, err
}

// CompressAirdropNoStruct is a free data retrieval call binding the contract method 0x1345dc05.
//
// Solidity: function CompressAirdropNoStruct(uint256 toIndex, uint256 amount, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressAirdropNoStruct(toIndex *big.Int, amount *big.Int, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressAirdropNoStruct(&_Rolluputils.CallOpts, toIndex, amount, sig)
}

// CompressAirdropNoStruct is a free data retrieval call binding the contract method 0x1345dc05.
//
// Solidity: function CompressAirdropNoStruct(uint256 toIndex, uint256 amount, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressAirdropNoStruct(toIndex *big.Int, amount *big.Int, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressAirdropNoStruct(&_Rolluputils.CallOpts, toIndex, amount, sig)
}

// CompressAirdropTxWithMessage is a free data retrieval call binding the contract method 0x79f8cb6f.
//
// Solidity: function CompressAirdropTxWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressAirdropTxWithMessage(opts *bind.CallOpts, message []byte, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressAirdropTxWithMessage", message, sig)
	return *ret0, err
}

// CompressAirdropTxWithMessage is a free data retrieval call binding the contract method 0x79f8cb6f.
//
// Solidity: function CompressAirdropTxWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressAirdropTxWithMessage(message []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressAirdropTxWithMessage(&_Rolluputils.CallOpts, message, sig)
}

// CompressAirdropTxWithMessage is a free data retrieval call binding the contract method 0x79f8cb6f.
//
// Solidity: function CompressAirdropTxWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressAirdropTxWithMessage(message []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressAirdropTxWithMessage(&_Rolluputils.CallOpts, message, sig)
}

// CompressBurnConsent is a free data retrieval call binding the contract method 0x625c7972.
//
// Solidity: function CompressBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressBurnConsent(opts *bind.CallOpts, _tx TypesBurnConsent) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressBurnConsent", _tx)
	return *ret0, err
}

// CompressBurnConsent is a free data retrieval call binding the contract method 0x625c7972.
//
// Solidity: function CompressBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressBurnConsent(_tx TypesBurnConsent) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnConsent(&_Rolluputils.CallOpts, _tx)
}

// CompressBurnConsent is a free data retrieval call binding the contract method 0x625c7972.
//
// Solidity: function CompressBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressBurnConsent(_tx TypesBurnConsent) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnConsent(&_Rolluputils.CallOpts, _tx)
}

// CompressBurnConsentFromEncoded is a free data retrieval call binding the contract method 0xce1f2349.
//
// Solidity: function CompressBurnConsentFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressBurnConsentFromEncoded(opts *bind.CallOpts, txBytes []byte, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressBurnConsentFromEncoded", txBytes, sig)
	return *ret0, err
}

// CompressBurnConsentFromEncoded is a free data retrieval call binding the contract method 0xce1f2349.
//
// Solidity: function CompressBurnConsentFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressBurnConsentFromEncoded(txBytes []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnConsentFromEncoded(&_Rolluputils.CallOpts, txBytes, sig)
}

// CompressBurnConsentFromEncoded is a free data retrieval call binding the contract method 0xce1f2349.
//
// Solidity: function CompressBurnConsentFromEncoded(bytes txBytes, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressBurnConsentFromEncoded(txBytes []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnConsentFromEncoded(&_Rolluputils.CallOpts, txBytes, sig)
}

// CompressBurnConsentNoStruct is a free data retrieval call binding the contract method 0x70db63d0.
//
// Solidity: function CompressBurnConsentNoStruct(uint256 fromIndex, uint256 amount, uint256 nonce, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressBurnConsentNoStruct(opts *bind.CallOpts, fromIndex *big.Int, amount *big.Int, nonce *big.Int, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressBurnConsentNoStruct", fromIndex, amount, nonce, sig)
	return *ret0, err
}

// CompressBurnConsentNoStruct is a free data retrieval call binding the contract method 0x70db63d0.
//
// Solidity: function CompressBurnConsentNoStruct(uint256 fromIndex, uint256 amount, uint256 nonce, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressBurnConsentNoStruct(fromIndex *big.Int, amount *big.Int, nonce *big.Int, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnConsentNoStruct(&_Rolluputils.CallOpts, fromIndex, amount, nonce, sig)
}

// CompressBurnConsentNoStruct is a free data retrieval call binding the contract method 0x70db63d0.
//
// Solidity: function CompressBurnConsentNoStruct(uint256 fromIndex, uint256 amount, uint256 nonce, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressBurnConsentNoStruct(fromIndex *big.Int, amount *big.Int, nonce *big.Int, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnConsentNoStruct(&_Rolluputils.CallOpts, fromIndex, amount, nonce, sig)
}

// CompressBurnConsentWithMessage is a free data retrieval call binding the contract method 0x0ef721cd.
//
// Solidity: function CompressBurnConsentWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressBurnConsentWithMessage(opts *bind.CallOpts, message []byte, sig []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressBurnConsentWithMessage", message, sig)
	return *ret0, err
}

// CompressBurnConsentWithMessage is a free data retrieval call binding the contract method 0x0ef721cd.
//
// Solidity: function CompressBurnConsentWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressBurnConsentWithMessage(message []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnConsentWithMessage(&_Rolluputils.CallOpts, message, sig)
}

// CompressBurnConsentWithMessage is a free data retrieval call binding the contract method 0x0ef721cd.
//
// Solidity: function CompressBurnConsentWithMessage(bytes message, bytes sig) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressBurnConsentWithMessage(message []byte, sig []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnConsentWithMessage(&_Rolluputils.CallOpts, message, sig)
}

// CompressBurnExecution is a free data retrieval call binding the contract method 0x89903a66.
//
// Solidity: function CompressBurnExecution((uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressBurnExecution(opts *bind.CallOpts, _tx TypesBurnExecution) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressBurnExecution", _tx)
	return *ret0, err
}

// CompressBurnExecution is a free data retrieval call binding the contract method 0x89903a66.
//
// Solidity: function CompressBurnExecution((uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressBurnExecution(_tx TypesBurnExecution) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnExecution(&_Rolluputils.CallOpts, _tx)
}

// CompressBurnExecution is a free data retrieval call binding the contract method 0x89903a66.
//
// Solidity: function CompressBurnExecution((uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressBurnExecution(_tx TypesBurnExecution) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnExecution(&_Rolluputils.CallOpts, _tx)
}

// CompressBurnExecutionFromEncoded is a free data retrieval call binding the contract method 0xe3c996b7.
//
// Solidity: function CompressBurnExecutionFromEncoded(bytes txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressBurnExecutionFromEncoded(opts *bind.CallOpts, txBytes []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressBurnExecutionFromEncoded", txBytes)
	return *ret0, err
}

// CompressBurnExecutionFromEncoded is a free data retrieval call binding the contract method 0xe3c996b7.
//
// Solidity: function CompressBurnExecutionFromEncoded(bytes txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressBurnExecutionFromEncoded(txBytes []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnExecutionFromEncoded(&_Rolluputils.CallOpts, txBytes)
}

// CompressBurnExecutionFromEncoded is a free data retrieval call binding the contract method 0xe3c996b7.
//
// Solidity: function CompressBurnExecutionFromEncoded(bytes txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressBurnExecutionFromEncoded(txBytes []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnExecutionFromEncoded(&_Rolluputils.CallOpts, txBytes)
}

// CompressBurnExecutionNoStruct is a free data retrieval call binding the contract method 0xe2b75030.
//
// Solidity: function CompressBurnExecutionNoStruct(uint256 fromIndex) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressBurnExecutionNoStruct(opts *bind.CallOpts, fromIndex *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressBurnExecutionNoStruct", fromIndex)
	return *ret0, err
}

// CompressBurnExecutionNoStruct is a free data retrieval call binding the contract method 0xe2b75030.
//
// Solidity: function CompressBurnExecutionNoStruct(uint256 fromIndex) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressBurnExecutionNoStruct(fromIndex *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnExecutionNoStruct(&_Rolluputils.CallOpts, fromIndex)
}

// CompressBurnExecutionNoStruct is a free data retrieval call binding the contract method 0xe2b75030.
//
// Solidity: function CompressBurnExecutionNoStruct(uint256 fromIndex) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressBurnExecutionNoStruct(fromIndex *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnExecutionNoStruct(&_Rolluputils.CallOpts, fromIndex)
}

// CompressBurnExecutionWithMessage is a free data retrieval call binding the contract method 0x574108dc.
//
// Solidity: function CompressBurnExecutionWithMessage(bytes message) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressBurnExecutionWithMessage(opts *bind.CallOpts, message []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressBurnExecutionWithMessage", message)
	return *ret0, err
}

// CompressBurnExecutionWithMessage is a free data retrieval call binding the contract method 0x574108dc.
//
// Solidity: function CompressBurnExecutionWithMessage(bytes message) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressBurnExecutionWithMessage(message []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnExecutionWithMessage(&_Rolluputils.CallOpts, message)
}

// CompressBurnExecutionWithMessage is a free data retrieval call binding the contract method 0x574108dc.
//
// Solidity: function CompressBurnExecutionWithMessage(bytes message) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressBurnExecutionWithMessage(message []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressBurnExecutionWithMessage(&_Rolluputils.CallOpts, message)
}

// CompressCreateAccount is a free data retrieval call binding the contract method 0x0a40b1f0.
//
// Solidity: function CompressCreateAccount((uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressCreateAccount(opts *bind.CallOpts, _tx TypesCreateAccount) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressCreateAccount", _tx)
	return *ret0, err
}

// CompressCreateAccount is a free data retrieval call binding the contract method 0x0a40b1f0.
//
// Solidity: function CompressCreateAccount((uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressCreateAccount(_tx TypesCreateAccount) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccount(&_Rolluputils.CallOpts, _tx)
}

// CompressCreateAccount is a free data retrieval call binding the contract method 0x0a40b1f0.
//
// Solidity: function CompressCreateAccount((uint256,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressCreateAccount(_tx TypesCreateAccount) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccount(&_Rolluputils.CallOpts, _tx)
}

// CompressCreateAccountFromEncoded is a free data retrieval call binding the contract method 0x87bb25b8.
//
// Solidity: function CompressCreateAccountFromEncoded(bytes txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressCreateAccountFromEncoded(opts *bind.CallOpts, txBytes []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressCreateAccountFromEncoded", txBytes)
	return *ret0, err
}

// CompressCreateAccountFromEncoded is a free data retrieval call binding the contract method 0x87bb25b8.
//
// Solidity: function CompressCreateAccountFromEncoded(bytes txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressCreateAccountFromEncoded(txBytes []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccountFromEncoded(&_Rolluputils.CallOpts, txBytes)
}

// CompressCreateAccountFromEncoded is a free data retrieval call binding the contract method 0x87bb25b8.
//
// Solidity: function CompressCreateAccountFromEncoded(bytes txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressCreateAccountFromEncoded(txBytes []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccountFromEncoded(&_Rolluputils.CallOpts, txBytes)
}

// CompressCreateAccountNoStruct is a free data retrieval call binding the contract method 0xc7d4234d.
//
// Solidity: function CompressCreateAccountNoStruct(uint256 accountID, uint256 stateID, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressCreateAccountNoStruct(opts *bind.CallOpts, accountID *big.Int, stateID *big.Int, tokenType *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressCreateAccountNoStruct", accountID, stateID, tokenType)
	return *ret0, err
}

// CompressCreateAccountNoStruct is a free data retrieval call binding the contract method 0xc7d4234d.
//
// Solidity: function CompressCreateAccountNoStruct(uint256 accountID, uint256 stateID, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressCreateAccountNoStruct(accountID *big.Int, stateID *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccountNoStruct(&_Rolluputils.CallOpts, accountID, stateID, tokenType)
}

// CompressCreateAccountNoStruct is a free data retrieval call binding the contract method 0xc7d4234d.
//
// Solidity: function CompressCreateAccountNoStruct(uint256 accountID, uint256 stateID, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressCreateAccountNoStruct(accountID *big.Int, stateID *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccountNoStruct(&_Rolluputils.CallOpts, accountID, stateID, tokenType)
}

// CompressCreateAccountWithMessage is a free data retrieval call binding the contract method 0x0a220a6b.
//
// Solidity: function CompressCreateAccountWithMessage(bytes message) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressCreateAccountWithMessage(opts *bind.CallOpts, message []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressCreateAccountWithMessage", message)
	return *ret0, err
}

// CompressCreateAccountWithMessage is a free data retrieval call binding the contract method 0x0a220a6b.
//
// Solidity: function CompressCreateAccountWithMessage(bytes message) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressCreateAccountWithMessage(message []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccountWithMessage(&_Rolluputils.CallOpts, message)
}

// CompressCreateAccountWithMessage is a free data retrieval call binding the contract method 0x0a220a6b.
//
// Solidity: function CompressCreateAccountWithMessage(bytes message) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressCreateAccountWithMessage(message []byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressCreateAccountWithMessage(&_Rolluputils.CallOpts, message)
}

// CompressManyAirdropFromEncoded is a free data retrieval call binding the contract method 0xe1b72cc2.
//
// Solidity: function CompressManyAirdropFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressManyAirdropFromEncoded(opts *bind.CallOpts, txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressManyAirdropFromEncoded", txBytes, sigs)
	return *ret0, err
}

// CompressManyAirdropFromEncoded is a free data retrieval call binding the contract method 0xe1b72cc2.
//
// Solidity: function CompressManyAirdropFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressManyAirdropFromEncoded(txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyAirdropFromEncoded(&_Rolluputils.CallOpts, txBytes, sigs)
}

// CompressManyAirdropFromEncoded is a free data retrieval call binding the contract method 0xe1b72cc2.
//
// Solidity: function CompressManyAirdropFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressManyAirdropFromEncoded(txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyAirdropFromEncoded(&_Rolluputils.CallOpts, txBytes, sigs)
}

// CompressManyBurnConsentFromEncoded is a free data retrieval call binding the contract method 0xb88efada.
//
// Solidity: function CompressManyBurnConsentFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressManyBurnConsentFromEncoded(opts *bind.CallOpts, txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressManyBurnConsentFromEncoded", txBytes, sigs)
	return *ret0, err
}

// CompressManyBurnConsentFromEncoded is a free data retrieval call binding the contract method 0xb88efada.
//
// Solidity: function CompressManyBurnConsentFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressManyBurnConsentFromEncoded(txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyBurnConsentFromEncoded(&_Rolluputils.CallOpts, txBytes, sigs)
}

// CompressManyBurnConsentFromEncoded is a free data retrieval call binding the contract method 0xb88efada.
//
// Solidity: function CompressManyBurnConsentFromEncoded(bytes[] txBytes, bytes[] sigs) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressManyBurnConsentFromEncoded(txBytes [][]byte, sigs [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyBurnConsentFromEncoded(&_Rolluputils.CallOpts, txBytes, sigs)
}

// CompressManyBurnExecutionFromEncoded is a free data retrieval call binding the contract method 0x56f29b49.
//
// Solidity: function CompressManyBurnExecutionFromEncoded(bytes[] txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressManyBurnExecutionFromEncoded(opts *bind.CallOpts, txBytes [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressManyBurnExecutionFromEncoded", txBytes)
	return *ret0, err
}

// CompressManyBurnExecutionFromEncoded is a free data retrieval call binding the contract method 0x56f29b49.
//
// Solidity: function CompressManyBurnExecutionFromEncoded(bytes[] txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressManyBurnExecutionFromEncoded(txBytes [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyBurnExecutionFromEncoded(&_Rolluputils.CallOpts, txBytes)
}

// CompressManyBurnExecutionFromEncoded is a free data retrieval call binding the contract method 0x56f29b49.
//
// Solidity: function CompressManyBurnExecutionFromEncoded(bytes[] txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressManyBurnExecutionFromEncoded(txBytes [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyBurnExecutionFromEncoded(&_Rolluputils.CallOpts, txBytes)
}

// CompressManyCreateAccountFromEncoded is a free data retrieval call binding the contract method 0xdc3498b5.
//
// Solidity: function CompressManyCreateAccountFromEncoded(bytes[] txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressManyCreateAccountFromEncoded(opts *bind.CallOpts, txBytes [][]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressManyCreateAccountFromEncoded", txBytes)
	return *ret0, err
}

// CompressManyCreateAccountFromEncoded is a free data retrieval call binding the contract method 0xdc3498b5.
//
// Solidity: function CompressManyCreateAccountFromEncoded(bytes[] txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressManyCreateAccountFromEncoded(txBytes [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyCreateAccountFromEncoded(&_Rolluputils.CallOpts, txBytes)
}

// CompressManyCreateAccountFromEncoded is a free data retrieval call binding the contract method 0xdc3498b5.
//
// Solidity: function CompressManyCreateAccountFromEncoded(bytes[] txBytes) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressManyCreateAccountFromEncoded(txBytes [][]byte) ([]byte, error) {
	return _Rolluputils.Contract.CompressManyCreateAccountFromEncoded(&_Rolluputils.CallOpts, txBytes)
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

// CompressTx is a free data retrieval call binding the contract method 0x02c36512.
//
// Solidity: function CompressTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
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
// Solidity: function CompressTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.CompressTx(&_Rolluputils.CallOpts, _tx)
}

// CompressTx is a free data retrieval call binding the contract method 0x02c36512.
//
// Solidity: function CompressTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes)
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

// CreateAccountFromBytes is a free data retrieval call binding the contract method 0xb3c2ec13.
//
// Solidity: function CreateAccountFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsCaller) CreateAccountFromBytes(opts *bind.CallOpts, txBytes []byte) (TypesCreateAccount, error) {
	var (
		ret0 = new(TypesCreateAccount)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CreateAccountFromBytes", txBytes)
	return *ret0, err
}

// CreateAccountFromBytes is a free data retrieval call binding the contract method 0xb3c2ec13.
//
// Solidity: function CreateAccountFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsSession) CreateAccountFromBytes(txBytes []byte) (TypesCreateAccount, error) {
	return _Rolluputils.Contract.CreateAccountFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// CreateAccountFromBytes is a free data retrieval call binding the contract method 0xb3c2ec13.
//
// Solidity: function CreateAccountFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256))
func (_Rolluputils *RolluputilsCallerSession) CreateAccountFromBytes(txBytes []byte) (TypesCreateAccount, error) {
	return _Rolluputils.Contract.CreateAccountFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// DecompressAirdrop is a free data retrieval call binding the contract method 0x006de34f.
//
// Solidity: function DecompressAirdrop(bytes txBytes) pure returns(uint256 toIndex, uint256 amount, bytes signature)
func (_Rolluputils *RolluputilsCaller) DecompressAirdrop(opts *bind.CallOpts, txBytes []byte) (struct {
	ToIndex   *big.Int
	Amount    *big.Int
	Signature []byte
}, error) {
	ret := new(struct {
		ToIndex   *big.Int
		Amount    *big.Int
		Signature []byte
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "DecompressAirdrop", txBytes)
	return *ret, err
}

// DecompressAirdrop is a free data retrieval call binding the contract method 0x006de34f.
//
// Solidity: function DecompressAirdrop(bytes txBytes) pure returns(uint256 toIndex, uint256 amount, bytes signature)
func (_Rolluputils *RolluputilsSession) DecompressAirdrop(txBytes []byte) (struct {
	ToIndex   *big.Int
	Amount    *big.Int
	Signature []byte
}, error) {
	return _Rolluputils.Contract.DecompressAirdrop(&_Rolluputils.CallOpts, txBytes)
}

// DecompressAirdrop is a free data retrieval call binding the contract method 0x006de34f.
//
// Solidity: function DecompressAirdrop(bytes txBytes) pure returns(uint256 toIndex, uint256 amount, bytes signature)
func (_Rolluputils *RolluputilsCallerSession) DecompressAirdrop(txBytes []byte) (struct {
	ToIndex   *big.Int
	Amount    *big.Int
	Signature []byte
}, error) {
	return _Rolluputils.Contract.DecompressAirdrop(&_Rolluputils.CallOpts, txBytes)
}

// DecompressBurnConsent is a free data retrieval call binding the contract method 0x0cda9938.
//
// Solidity: function DecompressBurnConsent(bytes txBytes) pure returns(uint256 fromIndex, uint256 amount, uint256 nonce, bytes signature)
func (_Rolluputils *RolluputilsCaller) DecompressBurnConsent(opts *bind.CallOpts, txBytes []byte) (struct {
	FromIndex *big.Int
	Amount    *big.Int
	Nonce     *big.Int
	Signature []byte
}, error) {
	ret := new(struct {
		FromIndex *big.Int
		Amount    *big.Int
		Nonce     *big.Int
		Signature []byte
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "DecompressBurnConsent", txBytes)
	return *ret, err
}

// DecompressBurnConsent is a free data retrieval call binding the contract method 0x0cda9938.
//
// Solidity: function DecompressBurnConsent(bytes txBytes) pure returns(uint256 fromIndex, uint256 amount, uint256 nonce, bytes signature)
func (_Rolluputils *RolluputilsSession) DecompressBurnConsent(txBytes []byte) (struct {
	FromIndex *big.Int
	Amount    *big.Int
	Nonce     *big.Int
	Signature []byte
}, error) {
	return _Rolluputils.Contract.DecompressBurnConsent(&_Rolluputils.CallOpts, txBytes)
}

// DecompressBurnConsent is a free data retrieval call binding the contract method 0x0cda9938.
//
// Solidity: function DecompressBurnConsent(bytes txBytes) pure returns(uint256 fromIndex, uint256 amount, uint256 nonce, bytes signature)
func (_Rolluputils *RolluputilsCallerSession) DecompressBurnConsent(txBytes []byte) (struct {
	FromIndex *big.Int
	Amount    *big.Int
	Nonce     *big.Int
	Signature []byte
}, error) {
	return _Rolluputils.Contract.DecompressBurnConsent(&_Rolluputils.CallOpts, txBytes)
}

// DecompressBurnExecution is a free data retrieval call binding the contract method 0x659175bb.
//
// Solidity: function DecompressBurnExecution(bytes txBytes) pure returns(uint256 fromIndex)
func (_Rolluputils *RolluputilsCaller) DecompressBurnExecution(opts *bind.CallOpts, txBytes []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressBurnExecution", txBytes)
	return *ret0, err
}

// DecompressBurnExecution is a free data retrieval call binding the contract method 0x659175bb.
//
// Solidity: function DecompressBurnExecution(bytes txBytes) pure returns(uint256 fromIndex)
func (_Rolluputils *RolluputilsSession) DecompressBurnExecution(txBytes []byte) (*big.Int, error) {
	return _Rolluputils.Contract.DecompressBurnExecution(&_Rolluputils.CallOpts, txBytes)
}

// DecompressBurnExecution is a free data retrieval call binding the contract method 0x659175bb.
//
// Solidity: function DecompressBurnExecution(bytes txBytes) pure returns(uint256 fromIndex)
func (_Rolluputils *RolluputilsCallerSession) DecompressBurnExecution(txBytes []byte) (*big.Int, error) {
	return _Rolluputils.Contract.DecompressBurnExecution(&_Rolluputils.CallOpts, txBytes)
}

// DecompressCreateAccount is a free data retrieval call binding the contract method 0x22002f28.
//
// Solidity: function DecompressCreateAccount(bytes txBytes) pure returns(uint256 accountID, uint256 stateID, uint256 tokenType)
func (_Rolluputils *RolluputilsCaller) DecompressCreateAccount(opts *bind.CallOpts, txBytes []byte) (struct {
	AccountID *big.Int
	StateID   *big.Int
	TokenType *big.Int
}, error) {
	ret := new(struct {
		AccountID *big.Int
		StateID   *big.Int
		TokenType *big.Int
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "DecompressCreateAccount", txBytes)
	return *ret, err
}

// DecompressCreateAccount is a free data retrieval call binding the contract method 0x22002f28.
//
// Solidity: function DecompressCreateAccount(bytes txBytes) pure returns(uint256 accountID, uint256 stateID, uint256 tokenType)
func (_Rolluputils *RolluputilsSession) DecompressCreateAccount(txBytes []byte) (struct {
	AccountID *big.Int
	StateID   *big.Int
	TokenType *big.Int
}, error) {
	return _Rolluputils.Contract.DecompressCreateAccount(&_Rolluputils.CallOpts, txBytes)
}

// DecompressCreateAccount is a free data retrieval call binding the contract method 0x22002f28.
//
// Solidity: function DecompressCreateAccount(bytes txBytes) pure returns(uint256 accountID, uint256 stateID, uint256 tokenType)
func (_Rolluputils *RolluputilsCallerSession) DecompressCreateAccount(txBytes []byte) (struct {
	AccountID *big.Int
	StateID   *big.Int
	TokenType *big.Int
}, error) {
	return _Rolluputils.Contract.DecompressCreateAccount(&_Rolluputils.CallOpts, txBytes)
}

// DecompressManyAirdrop is a free data retrieval call binding the contract method 0x988ce327.
//
// Solidity: function DecompressManyAirdrop(bytes txs) pure returns((uint256,uint256,uint256,bytes)[] structTxs)
func (_Rolluputils *RolluputilsCaller) DecompressManyAirdrop(opts *bind.CallOpts, txs []byte) ([]TxTransfer, error) {
	var (
		ret0 = new([]TxTransfer)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressManyAirdrop", txs)
	return *ret0, err
}

// DecompressManyAirdrop is a free data retrieval call binding the contract method 0x988ce327.
//
// Solidity: function DecompressManyAirdrop(bytes txs) pure returns((uint256,uint256,uint256,bytes)[] structTxs)
func (_Rolluputils *RolluputilsSession) DecompressManyAirdrop(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressManyAirdrop(&_Rolluputils.CallOpts, txs)
}

// DecompressManyAirdrop is a free data retrieval call binding the contract method 0x988ce327.
//
// Solidity: function DecompressManyAirdrop(bytes txs) pure returns((uint256,uint256,uint256,bytes)[] structTxs)
func (_Rolluputils *RolluputilsCallerSession) DecompressManyAirdrop(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressManyAirdrop(&_Rolluputils.CallOpts, txs)
}

// DecompressManyBurnConsent is a free data retrieval call binding the contract method 0x62108cc2.
//
// Solidity: function DecompressManyBurnConsent(bytes txs) pure returns((uint256,uint256,bytes)[] structTxs)
func (_Rolluputils *RolluputilsCaller) DecompressManyBurnConsent(opts *bind.CallOpts, txs []byte) ([]TxBurnConsent, error) {
	var (
		ret0 = new([]TxBurnConsent)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressManyBurnConsent", txs)
	return *ret0, err
}

// DecompressManyBurnConsent is a free data retrieval call binding the contract method 0x62108cc2.
//
// Solidity: function DecompressManyBurnConsent(bytes txs) pure returns((uint256,uint256,bytes)[] structTxs)
func (_Rolluputils *RolluputilsSession) DecompressManyBurnConsent(txs []byte) ([]TxBurnConsent, error) {
	return _Rolluputils.Contract.DecompressManyBurnConsent(&_Rolluputils.CallOpts, txs)
}

// DecompressManyBurnConsent is a free data retrieval call binding the contract method 0x62108cc2.
//
// Solidity: function DecompressManyBurnConsent(bytes txs) pure returns((uint256,uint256,bytes)[] structTxs)
func (_Rolluputils *RolluputilsCallerSession) DecompressManyBurnConsent(txs []byte) ([]TxBurnConsent, error) {
	return _Rolluputils.Contract.DecompressManyBurnConsent(&_Rolluputils.CallOpts, txs)
}

// DecompressManyBurnExecution is a free data retrieval call binding the contract method 0xa00d5166.
//
// Solidity: function DecompressManyBurnExecution(bytes txs) pure returns((uint256)[] structTxs)
func (_Rolluputils *RolluputilsCaller) DecompressManyBurnExecution(opts *bind.CallOpts, txs []byte) ([]TxBurnExecution, error) {
	var (
		ret0 = new([]TxBurnExecution)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressManyBurnExecution", txs)
	return *ret0, err
}

// DecompressManyBurnExecution is a free data retrieval call binding the contract method 0xa00d5166.
//
// Solidity: function DecompressManyBurnExecution(bytes txs) pure returns((uint256)[] structTxs)
func (_Rolluputils *RolluputilsSession) DecompressManyBurnExecution(txs []byte) ([]TxBurnExecution, error) {
	return _Rolluputils.Contract.DecompressManyBurnExecution(&_Rolluputils.CallOpts, txs)
}

// DecompressManyBurnExecution is a free data retrieval call binding the contract method 0xa00d5166.
//
// Solidity: function DecompressManyBurnExecution(bytes txs) pure returns((uint256)[] structTxs)
func (_Rolluputils *RolluputilsCallerSession) DecompressManyBurnExecution(txs []byte) ([]TxBurnExecution, error) {
	return _Rolluputils.Contract.DecompressManyBurnExecution(&_Rolluputils.CallOpts, txs)
}

// DecompressManyCreateAccount is a free data retrieval call binding the contract method 0x1392f5d4.
//
// Solidity: function DecompressManyCreateAccount(bytes txs) pure returns((uint256,uint256,uint256)[] structTxs)
func (_Rolluputils *RolluputilsCaller) DecompressManyCreateAccount(opts *bind.CallOpts, txs []byte) ([]TxCreateAccount, error) {
	var (
		ret0 = new([]TxCreateAccount)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "DecompressManyCreateAccount", txs)
	return *ret0, err
}

// DecompressManyCreateAccount is a free data retrieval call binding the contract method 0x1392f5d4.
//
// Solidity: function DecompressManyCreateAccount(bytes txs) pure returns((uint256,uint256,uint256)[] structTxs)
func (_Rolluputils *RolluputilsSession) DecompressManyCreateAccount(txs []byte) ([]TxCreateAccount, error) {
	return _Rolluputils.Contract.DecompressManyCreateAccount(&_Rolluputils.CallOpts, txs)
}

// DecompressManyCreateAccount is a free data retrieval call binding the contract method 0x1392f5d4.
//
// Solidity: function DecompressManyCreateAccount(bytes txs) pure returns((uint256,uint256,uint256)[] structTxs)
func (_Rolluputils *RolluputilsCallerSession) DecompressManyCreateAccount(txs []byte) ([]TxCreateAccount, error) {
	return _Rolluputils.Contract.DecompressManyCreateAccount(&_Rolluputils.CallOpts, txs)
}

// DecompressManyTransfer is a free data retrieval call binding the contract method 0xd1b4bf51.
//
// Solidity: function DecompressManyTransfer(bytes txs) pure returns((uint256,uint256,uint256,bytes)[] structTxs)
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
// Solidity: function DecompressManyTransfer(bytes txs) pure returns((uint256,uint256,uint256,bytes)[] structTxs)
func (_Rolluputils *RolluputilsSession) DecompressManyTransfer(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressManyTransfer(&_Rolluputils.CallOpts, txs)
}

// DecompressManyTransfer is a free data retrieval call binding the contract method 0xd1b4bf51.
//
// Solidity: function DecompressManyTransfer(bytes txs) pure returns((uint256,uint256,uint256,bytes)[] structTxs)
func (_Rolluputils *RolluputilsCallerSession) DecompressManyTransfer(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressManyTransfer(&_Rolluputils.CallOpts, txs)
}

// DecompressTransfers is a free data retrieval call binding the contract method 0x67347755.
//
// Solidity: function DecompressTransfers(bytes txs) pure returns((uint256,uint256,uint256,bytes)[])
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
// Solidity: function DecompressTransfers(bytes txs) pure returns((uint256,uint256,uint256,bytes)[])
func (_Rolluputils *RolluputilsSession) DecompressTransfers(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressTransfers(&_Rolluputils.CallOpts, txs)
}

// DecompressTransfers is a free data retrieval call binding the contract method 0x67347755.
//
// Solidity: function DecompressTransfers(bytes txs) pure returns((uint256,uint256,uint256,bytes)[])
func (_Rolluputils *RolluputilsCallerSession) DecompressTransfers(txs []byte) ([]TxTransfer, error) {
	return _Rolluputils.Contract.DecompressTransfers(&_Rolluputils.CallOpts, txs)
}

// DecompressTx is a free data retrieval call binding the contract method 0xeedeb9d9.
//
// Solidity: function DecompressTx(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 amount, bytes sig)
func (_Rolluputils *RolluputilsCaller) DecompressTx(opts *bind.CallOpts, txBytes []byte) (struct {
	From   *big.Int
	To     *big.Int
	Amount *big.Int
	Sig    []byte
}, error) {
	ret := new(struct {
		From   *big.Int
		To     *big.Int
		Amount *big.Int
		Sig    []byte
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "DecompressTx", txBytes)
	return *ret, err
}

// DecompressTx is a free data retrieval call binding the contract method 0xeedeb9d9.
//
// Solidity: function DecompressTx(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 amount, bytes sig)
func (_Rolluputils *RolluputilsSession) DecompressTx(txBytes []byte) (struct {
	From   *big.Int
	To     *big.Int
	Amount *big.Int
	Sig    []byte
}, error) {
	return _Rolluputils.Contract.DecompressTx(&_Rolluputils.CallOpts, txBytes)
}

// DecompressTx is a free data retrieval call binding the contract method 0xeedeb9d9.
//
// Solidity: function DecompressTx(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 amount, bytes sig)
func (_Rolluputils *RolluputilsCallerSession) DecompressTx(txBytes []byte) (struct {
	From   *big.Int
	To     *big.Int
	Amount *big.Int
	Sig    []byte
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

// GetYearMonth is a free data retrieval call binding the contract method 0xda296257.
//
// Solidity: function GetYearMonth() view returns(uint256 yearMonth)
func (_Rolluputils *RolluputilsCaller) GetYearMonth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "GetYearMonth")
	return *ret0, err
}

// GetYearMonth is a free data retrieval call binding the contract method 0xda296257.
//
// Solidity: function GetYearMonth() view returns(uint256 yearMonth)
func (_Rolluputils *RolluputilsSession) GetYearMonth() (*big.Int, error) {
	return _Rolluputils.Contract.GetYearMonth(&_Rolluputils.CallOpts)
}

// GetYearMonth is a free data retrieval call binding the contract method 0xda296257.
//
// Solidity: function GetYearMonth() view returns(uint256 yearMonth)
func (_Rolluputils *RolluputilsCallerSession) GetYearMonth() (*big.Int, error) {
	return _Rolluputils.Contract.GetYearMonth(&_Rolluputils.CallOpts)
}

// HashFromAccount is a free data retrieval call binding the contract method 0x2953145f.
//
// Solidity: function HashFromAccount((uint256,uint256,uint256,uint256,uint256,uint256) account) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromAccount(opts *bind.CallOpts, account TypesUserAccount) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromAccount", account)
	return *ret0, err
}

// HashFromAccount is a free data retrieval call binding the contract method 0x2953145f.
//
// Solidity: function HashFromAccount((uint256,uint256,uint256,uint256,uint256,uint256) account) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromAccount(account TypesUserAccount) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromAccount(&_Rolluputils.CallOpts, account)
}

// HashFromAccount is a free data retrieval call binding the contract method 0x2953145f.
//
// Solidity: function HashFromAccount((uint256,uint256,uint256,uint256,uint256,uint256) account) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromAccount(account TypesUserAccount) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromAccount(&_Rolluputils.CallOpts, account)
}

// HashFromBurnConsent is a free data retrieval call binding the contract method 0xe38a529d.
//
// Solidity: function HashFromBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromBurnConsent(opts *bind.CallOpts, _tx TypesBurnConsent) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromBurnConsent", _tx)
	return *ret0, err
}

// HashFromBurnConsent is a free data retrieval call binding the contract method 0xe38a529d.
//
// Solidity: function HashFromBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromBurnConsent(_tx TypesBurnConsent) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromBurnConsent(&_Rolluputils.CallOpts, _tx)
}

// HashFromBurnConsent is a free data retrieval call binding the contract method 0xe38a529d.
//
// Solidity: function HashFromBurnConsent((uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromBurnConsent(_tx TypesBurnConsent) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromBurnConsent(&_Rolluputils.CallOpts, _tx)
}

// HashFromBurnExecution is a free data retrieval call binding the contract method 0x88f6837b.
//
// Solidity: function HashFromBurnExecution((uint256,uint256) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromBurnExecution(opts *bind.CallOpts, _tx TypesBurnExecution) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromBurnExecution", _tx)
	return *ret0, err
}

// HashFromBurnExecution is a free data retrieval call binding the contract method 0x88f6837b.
//
// Solidity: function HashFromBurnExecution((uint256,uint256) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromBurnExecution(_tx TypesBurnExecution) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromBurnExecution(&_Rolluputils.CallOpts, _tx)
}

// HashFromBurnExecution is a free data retrieval call binding the contract method 0x88f6837b.
//
// Solidity: function HashFromBurnExecution((uint256,uint256) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromBurnExecution(_tx TypesBurnExecution) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromBurnExecution(&_Rolluputils.CallOpts, _tx)
}

// HashFromTx is a free data retrieval call binding the contract method 0xb90cbf51.
//
// Solidity: function HashFromTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes32)
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
// Solidity: function HashFromTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromTx(_tx TypesTransaction) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromTx(&_Rolluputils.CallOpts, _tx)
}

// HashFromTx is a free data retrieval call binding the contract method 0xb90cbf51.
//
// Solidity: function HashFromTx((uint256,uint256,uint256,uint256,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromTx(_tx TypesTransaction) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromTx(&_Rolluputils.CallOpts, _tx)
}

// PDALeafToHash is a free data retrieval call binding the contract method 0xc2ddab33.
//
// Solidity: function PDALeafToHash((bytes) _PDA_Leaf) pure returns(bytes32)
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
// Solidity: function PDALeafToHash((bytes) _PDA_Leaf) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) PDALeafToHash(_PDA_Leaf TypesPDALeaf) ([32]byte, error) {
	return _Rolluputils.Contract.PDALeafToHash(&_Rolluputils.CallOpts, _PDA_Leaf)
}

// PDALeafToHash is a free data retrieval call binding the contract method 0xc2ddab33.
//
// Solidity: function PDALeafToHash((bytes) _PDA_Leaf) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) PDALeafToHash(_PDA_Leaf TypesPDALeaf) ([32]byte, error) {
	return _Rolluputils.Contract.PDALeafToHash(&_Rolluputils.CallOpts, _PDA_Leaf)
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes))
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
// Solidity: function TxFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes))
func (_Rolluputils *RolluputilsSession) TxFromBytes(txBytes []byte) (TypesTransaction, error) {
	return _Rolluputils.Contract.TxFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes))
func (_Rolluputils *RolluputilsCallerSession) TxFromBytes(txBytes []byte) (TypesTransaction, error) {
	return _Rolluputils.Contract.TxFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytesDeconstructed is a free data retrieval call binding the contract method 0x2013a0cf.
//
// Solidity: function TxFromBytesDeconstructed(bytes txBytes) pure returns(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount)
func (_Rolluputils *RolluputilsCaller) TxFromBytesDeconstructed(opts *bind.CallOpts, txBytes []byte) (struct {
	TxType    *big.Int
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	Amount    *big.Int
}, error) {
	ret := new(struct {
		TxType    *big.Int
		From      *big.Int
		To        *big.Int
		TokenType *big.Int
		Nonce     *big.Int
		Amount    *big.Int
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "TxFromBytesDeconstructed", txBytes)
	return *ret, err
}

// TxFromBytesDeconstructed is a free data retrieval call binding the contract method 0x2013a0cf.
//
// Solidity: function TxFromBytesDeconstructed(bytes txBytes) pure returns(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount)
func (_Rolluputils *RolluputilsSession) TxFromBytesDeconstructed(txBytes []byte) (struct {
	TxType    *big.Int
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	Amount    *big.Int
}, error) {
	return _Rolluputils.Contract.TxFromBytesDeconstructed(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytesDeconstructed is a free data retrieval call binding the contract method 0x2013a0cf.
//
// Solidity: function TxFromBytesDeconstructed(bytes txBytes) pure returns(uint256 txType, uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 amount)
func (_Rolluputils *RolluputilsCallerSession) TxFromBytesDeconstructed(txBytes []byte) (struct {
	TxType    *big.Int
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
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

// GetAccountHash is a free data retrieval call binding the contract method 0xdc84864c.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) GetAccountHash(opts *bind.CallOpts, id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int, burn *big.Int, lastBurn *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getAccountHash", id, balance, nonce, tokenType, burn, lastBurn)
	return *ret0, err
}

// GetAccountHash is a free data retrieval call binding the contract method 0xdc84864c.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) GetAccountHash(id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int, burn *big.Int, lastBurn *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetAccountHash(&_Rolluputils.CallOpts, id, balance, nonce, tokenType, burn, lastBurn)
}

// GetAccountHash is a free data retrieval call binding the contract method 0xdc84864c.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType, uint256 burn, uint256 lastBurn) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) GetAccountHash(id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int, burn *big.Int, lastBurn *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetAccountHash(&_Rolluputils.CallOpts, id, balance, nonce, tokenType, burn, lastBurn)
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 txType, uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) GetTxSignBytes(opts *bind.CallOpts, txType *big.Int, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getTxSignBytes", txType, fromIndex, toIndex, tokenType, nonce, amount)
	return *ret0, err
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 txType, uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) GetTxSignBytes(txType *big.Int, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes(&_Rolluputils.CallOpts, txType, fromIndex, toIndex, tokenType, nonce, amount)
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 txType, uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) GetTxSignBytes(txType *big.Int, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes(&_Rolluputils.CallOpts, txType, fromIndex, toIndex, tokenType, nonce, amount)
}
