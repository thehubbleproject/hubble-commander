package core

import (
	"encoding/hex"
	"fmt"
	"math"
	"testing"

	"github.com/BOPR/aggregator"
	"github.com/BOPR/core"

	"github.com/BOPR/config"
	"github.com/stretchr/testify/require"
)

// func TestPubkeyHashCreation(t *testing.T) {
// 	bz, err := core.ABIEncodePubkey("90718dcbc2477c86294742fb72bf098ba85ff671b88c8d79b2e09ce19bdbd88fd87047aaebc775b168372752aa8bc4e5be1ca5d39284fed00722f341927888c3")
// 	if err != nil {
// 		panic(err)
// 	}
// 	hash := common.Keccak256(bz)
// 	fmt.Println(hash.String())
// }

func TestTxProcessing(t *testing.T) {
	err := config.ParseAndInitGlobalConfig("../.")
	require.Equal(t, err, nil, "Error while parsing and init config")
	db, err := core.NewDB()
	require.Equal(t, err, nil, "Error while creating database")
	core.DBInstance = db
	bazooka, err := core.NewPreLoadedBazooka()
	require.Equal(t, err, nil, "Error while creating bazooka")
	core.LoadedBazooka = bazooka
	agg := aggregator.NewAggregator()
	genesisAccounts, err := core.LoadedBazooka.GetGenesisAccounts()
	require.Equal(t, err, nil, "error loading genesis accounts")
	zeroAccount := genesisAccounts[0]
	diff := int(math.Exp2(float64(17)))
	var allAccounts []core.UserAccount
	var allPDALeaf []core.PDA

	// fill the tree with zero leaves
	for diff > 0 {
		newAcc := core.EmptyAccount()
		newAcc.Data = zeroAccount.Data
		newAcc.Hash = core.ZERO_VALUE_LEAF.String()
		newPDA := core.NewEmptyPDA()
		newPDA.Hash = core.ZERO_VALUE_LEAF.String()
		allPDALeaf = append(allPDALeaf, *newPDA)
		allAccounts = append(allAccounts, newAcc)
		diff--
	}

	// newParams := core.Params{StakeAmount: 10, MaxDepth: 17, MaxDepositSubTreeHeight: 1}
	// core.DBInstance.UpdateStakeAmount(newParams.StakeAmount)
	// core.DBInstance.UpdateMaxDepth(newParams.MaxDepth)
	// core.DBInstance.UpdateDepositSubTreeHeight(newParams.MaxDepositSubTreeHeight)
	// core.DBInstance.UpdateFinalisationTimePerBatch(40320)
	// // load accounts
	// err = core.DBInstance.InitBalancesTree(17, allAccounts)
	// require.Equal(t, err, nil, "error initing balances tree")

	// err = core.DBInstance.InitPDATree(17, allPDALeaf)
	// require.Equal(t, err, nil, "error initing balances tree")
	// account, err := core.DBInstance.GetAccountByIndex(2)
	// aliceAccountBytes, err := core.LoadedBazooka.EncodeAccount(2, 10, 0, 1, 2, 1)
	// require.Equal(t, err, nil, "error encoding alice account")

	// fmt.Println("account 2", hex.EncodeToString(aliceAccountBytes))
	// id, balance, _, _, _, _, err := core.LoadedBazooka.DecodeAccount(aliceAccountBytes)
	// if err != nil {
	// 	panic("error while decoding alice account")
	// }
	// fmt.Println(balance, id)
	// start := time.Now()
	// account.Data = aliceAccountBytes
	// db.UpdateAccount(account)
	// elapsed := time.Since(start)
	// log.Printf("Binomial took %s", elapsed)

	var txs []core.Tx
	// txBytes, err := bazooka.EncodeCreateAccountTx(11, 11, 1)
	// if err != nil {
	// 	panic(err)
	// }

	// txCore := core.NewPendingTx(0, uint64(11), core.TX_CREATE_ACCOUNT, "0x1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)
	account, err := core.DBInstance.GetAccountByIndex(2)
	fmt.Println("account data", hex.EncodeToString(account.Data))
	id, balance, nonce, token, burn, lastBurn, err := core.LoadedBazooka.DecodeAccount(account.Data)
	if err != nil {
		panic("error while decoding alice account")
	}
	fmt.Println("data", id, "balance", balance, "nonce", nonce, "token", token, burn, lastBurn)
	txBytes, err := bazooka.EncodeBurnConsentTx(2, 1, 1, core.TX_BURN_CONSENT)
	if err != nil {
		panic(err)
	}
	txCore := core.NewPendingTx(2, uint64(2), core.TX_BURN_CONSENT, "1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)
	// account, err := core.DBInstance.GetAccountByIndex(4)
	// fmt.Println("account data", hex.EncodeToString(account.Data))
	// id, balance, nonce, token, burn, lastBurn, err := core.LoadedBazooka.DecodeAccount(account.Data)
	// if err != nil {
	// 	panic("error while decoding alice account")
	// }
	// fmt.Println("data", id, "balance", balance, "nonce", nonce, "token", token, burn, lastBurn)
	// txBytes, err := bazooka.EncodeBurnExecTx(4, core.TX_BURN_EXEC)
	// if err != nil {
	// 	panic(err)
	// }
	// txCore := core.NewPendingTx(4, uint64(2), core.TX_BURN_EXEC, "0x1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)
	txs = append(txs, txCore)
	err = agg.ProcessTx(txs)
	require.Equal(t, err, nil, "error processing tx")
}
