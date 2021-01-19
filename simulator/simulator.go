package simulator

import (
	"errors"
	"fmt"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/db"
	"github.com/BOPR/wallet"
)

var (
	ErrInvalidAmount = errors.New("Invalid amount")
	ErrStateInActive = errors.New("User state inactive")
)

func Run(n int64) error {
	cfg, err := config.ParseConfig()
	if err != nil {
		return err
	}
	bazooka, err := bazooka.NewPreLoadedBazooka(cfg)
	if err != nil {
		return err
	}
	DBI, err := db.NewDB(cfg)
	if err != nil {
		return err
	}
	defer DBI.Close()
	params, err := DBI.GetParams()
	if err != nil {
		return err
	}
	err = run(bazooka, DBI, params, n)
	if err != nil {
		return err
	}
	return nil
}

func run(bazooka bazooka.Bazooka, DBI db.DB, params core.Params, n int64) error {
	wallets, err := makeWallets(bazooka, DBI, params)
	if err != nil {
		return err
	}
	secretBytes, publicKeyBytes := wallets[0].Bytes()
	f := func(int64) {
		_, err := transfer(&DBI, &bazooka, 2, 3, 1, 0, secretBytes, publicKeyBytes)
		if err != nil {
			panic(err)
		}
	}
	perSecond := benchmarkPerSecond(n, f)
	fmt.Println("Transactions per second", perSecond)
	return nil
}

func makeWallets(bazooka bazooka.Bazooka, DBI db.DB, params core.Params) ([]wallet.Wallet, error) {
	var wallets []wallet.Wallet
	for i := 0; i < 2; i++ {
		user, err := wallet.NewWallet()
		if err != nil {
			return nil, err
		}

		wallets = append(wallets, user)

		_, publicKeyBytes := user.Bytes()

		pubkeyIndex := uint64(i + 2)
		path, err := core.SolidityPathToNodePath(pubkeyIndex, params.MaxDepth)
		if err != nil {
			return nil, err
		}

		// add accounts to tree
		acc, err := core.NewAccount(pubkeyIndex, publicKeyBytes, path)
		if err != nil {
			return nil, err
		}
		err = DBI.UpdateAccount(*acc)
		if err != nil {
			return nil, err
		}
		// add accounts to state tree
		userState, err := bazooka.EncodeState(pubkeyIndex, 10, 0, 1)
		if err != nil {
			return nil, err
		}
		newUser := core.NewUserState(pubkeyIndex, core.STATUS_ACTIVE, path, userState)
		err = DBI.UpdateState(*newUser)
		if err != nil {
			return nil, err
		}
	}
	return wallets, nil
}

func transfer(DBI *db.DB, bazooka *bazooka.Bazooka, fromIndex, toIndex, amount, fee uint64, priv, pub []byte) (txHash string, err error) {
	from, err := DBI.GetStateByIndex(fromIndex)
	if err != nil {
		return
	}

	if !from.IsActive() {
		return "", ErrStateInActive
	}

	to, err := DBI.GetStateByIndex(toIndex)
	if err != nil {
		return
	}

	if !to.IsActive() {
		return "", ErrStateInActive
	}

	_, balance, nonce, _, err := bazooka.DecodeState(from.Data)
	if err != nil {
		return
	}

	if balance.Int64() <= int64(amount+fee) {
		return "", ErrInvalidAmount
	}

	txData, err := bazooka.EncodeTransferTx(int64(fromIndex), int64(toIndex), int64(fee), nonce.Int64(), int64(amount), core.TX_TRANSFER_TYPE)
	if err != nil {
		return
	}

	tx, err := core.NewPendingTx(fromIndex, toIndex, core.TX_TRANSFER_TYPE, []byte(""), txData)
	if err != nil {
		return
	}

	err = signAndBroadcast(bazooka, DBI, tx, priv, pub)
	if err != nil {
		return
	}

	return tx.TxHash, nil
}

func signAndBroadcast(b *bazooka.Bazooka, DBI *db.DB, tx core.Tx, priv, pub []byte) (err error) {
	txBytes, err := bazooka.GetSignBytes(*b, &tx)
	if err != nil {
		return
	}

	err = tx.SignTx(priv, pub, txBytes)
	if err != nil {
		return
	}
	err = tx.AssignHash()
	if err != nil {
		return
	}

	err = DBI.InsertTx(&tx)
	if err != nil {
		return err
	}
	return nil
}
