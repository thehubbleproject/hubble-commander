package router

import (
	"errors"
	"fmt"

	"github.com/BOPR/config"
	"github.com/BOPR/contracts/rollupcaller"
	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type transferTransaction struct {
	core.Tx
}

func (tx *transferTransaction) processTx(balanceRoot, accountRoot core.ByteArray) (newBalanceRoot core.ByteArray, accounts [][]byte, err error) {
	var PDAProof core.PDAMerkleProof
	var accMPs []core.AccountMerkleProof
	core.DBInstance.FetchPDAProofWithID(tx.Accounts[0], &PDAProof)
	core.DBInstance.FetchMPWithID(tx.Accounts[0], &accMPs[0])
	toAcc, err := core.DBInstance.GetAccountByIndex(tx.Accounts[1])
	if err != nil {
		return
	}
	var toSiblings []core.UserAccount
	dbCopy, _ := core.NewDB()
	mysqlTx := dbCopy.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			mysqlTx.Rollback()
		}
	}()
	dbCopy.Instance = mysqlTx
	updatedAccount, _, err := tx.applyTxWithProof(accMPs[0], tx.Tx)
	if err != nil {
		return
	}
	accMPs[0].Account.Data = updatedAccount
	err = dbCopy.UpdateAccount(accMPs[0].Account)
	if err != nil {
		return
	}
	// TODO add a check to ensure that DB copy of state matches the one returned by ApplyTransferTx
	toSiblings, err = dbCopy.GetSiblings(toAcc.Path)
	if err != nil {
		return
	}
	accMPs[1] = core.NewAccountMerkleProof(toAcc, toSiblings)
	opts := bind.CallOpts{From: config.OperatorAddress}
	fromMP, err := accMPs[0].ToABIVersion()
	if err != nil {
		return
	}
	toMP, err := accMPs[1].ToABIVersion()
	if err != nil {
		return
	}
	typesAccountProofs := rollupcaller.TypesAccountProofs{From: fromMP, To: toMP}
	updatedRoot, updatedFrom, updatedTo, errCode, IsValidTx, err := core.LoadedBazooka.RollupCaller.ProcessTransferTx(&opts,
		balanceRoot,
		accountRoot,
		tx.Signature,
		tx.Data,
		PDAProof.ToABIVersion(),
		typesAccountProofs,
	)
	if err != nil {
		return
	}
	accounts = append(accounts, updatedFrom, updatedTo)

	// b.log.Info("Processed transaction", "IsSuccess", IsValidTx, "newRoot", updatedRoot)

	if !IsValidTx {
		fmt.Println("Invalid transaction", "error_code", errCode)
		return newBalanceRoot, accounts, errors.New("Tx is invalid")
	}
	newBalanceRoot = core.BytesToByteArray(updatedRoot[:])
	return
}

func (transfer *transferTransaction) compressTxs(txs []core.Tx) ([]byte, error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	var data [][]byte
	for _, tx := range txs {
		data = append(data, tx.Data)
	}
	// TOOD remove and update the transfer
	return core.LoadedBazooka.RollupUtils.CompressManyAirdropFromEncoded(&opts, data)
}
func (transfer transferTransaction) applyTxWithProof(accountMP core.AccountMerkleProof, tx core.Tx) (updatedAccount []byte, updatedRoot core.ByteArray, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	accMP, err := accountMP.ToABIVersion()
	if err != nil {
		return nil, core.ByteArray{}, err
	}
	updatedAccountBytes, updatedRoot, err := core.LoadedBazooka.RollupCaller.ApplyAirdropTx(&opts, accMP, tx.Data)
	if err != nil {
		return updatedAccountBytes, updatedRoot, err
	}
	return
}

func (transfer *transferTransaction) applyTxWithoutProof(tx core.Tx, accountID uint64) (updatedAccount []byte, updatedRoot core.ByteArray, err error) {
	opts := bind.CallOpts{From: config.OperatorAddress}
	var accountMP core.AccountMerkleProof
	core.DBInstance.FetchMPWithID(accountID, &accountMP)
	accMP, err := accountMP.ToABIVersion()
	if err != nil {
		return nil, core.ByteArray{}, err
	}
	updatedAccountBytes, updatedRoot, err := core.LoadedBazooka.RollupCaller.ApplyAirdropTx(&opts, accMP, tx.Data)
	if err != nil {
		return updatedAccountBytes, updatedRoot, err
	}
	return
}
