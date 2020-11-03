package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
	"github.com/BOPR/wallet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ErrInvalidAmount = errors.New("Invalid amount")
	ErrStateInActive = errors.New("User state inactive")
)

//  SendTransferTx generated init command to initialise the config file
func SendTransferTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Transfers assets between 2 accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			toIndex := viper.GetUint64(FlagToID)
			fromIndex := viper.GetUint64(FlagFromID)
			privKey := viper.GetString(FlagPrivKey)
			pubKey := viper.GetString(FlagPubKey)
			amount := viper.GetUint64(FlagAmount)

			db, err := core.NewDB()
			if err != nil {
				return err
			}
			defer db.Close()

			bazooka, err := core.NewPreLoadedBazooka()
			if err != nil {
				return err
			}
			err, txHash := ValidateAndTransfer(db, bazooka, fromIndex, toIndex, amount, privKey, pubKey)
			if err != nil {
				return err
			}

			fmt.Println("Transaction submitted successfully", "hash", txHash)
			return nil
		},
	}
	cmd.Flags().StringP(FlagToID, "", "", "--to=<to-account>")
	cmd.Flags().StringP(FlagFromID, "", "", "--from=<from-account>")
	cmd.Flags().StringP(FlagTokenID, "", "", "--token=<token-id>")
	cmd.Flags().StringP(FlagPubKey, "", "", "--pubkey=<pubkey>")
	cmd.Flags().StringP(FlagPrivKey, "", "", "--privkey=<privkey>")
	cmd.Flags().StringP(FlagAmount, "", "", "--amount=<amount>")
	cmd.MarkFlagRequired(FlagTokenID)
	return cmd
}

func DummyTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dummy-transfer",
		Short: "Creates 2 accounts and creates a transfer between them",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := core.NewDB()
			if err != nil {
				return err
			}
			defer db.Close()

			bazooka, err := core.NewPreLoadedBazooka()
			if err != nil {
				return err
			}
			params, err := db.GetParams()
			if err != nil {
				return err
			}

			// create 2 accounts
			var users []wallet.Wallet
			for i := 0; i < 2; i++ {
				user, err := wallet.NewWallet()
				if err != nil {
					return err
				}
				users = append(users, user)
				secretBytes, publicKeyBytes := user.Bytes()
				publicKey := hex.EncodeToString(publicKeyBytes)
				fmt.Println("Adding new account", "privkey", hex.EncodeToString(secretBytes), "publickey", publicKey)

				pubkeyIndex := uint64(i + 2)
				path, err := core.SolidityPathToNodePath(uint64(pubkeyIndex), params.MaxDepth)
				if err != nil {
					return err
				}

				// add accounts to tree
				acc, err := core.NewAccount(pubkeyIndex, publicKey, path)
				if err != nil {
					return err
				}
				err = db.UpdateAccount(*acc)
				if err != nil {
					return err
				}
				// add accounts to state tree
				userState, err := bazooka.EncodeState(pubkeyIndex, 10, 0, 1)
				if err != nil {
					return err
				}
				newUser := core.NewUserState(pubkeyIndex, core.STATUS_ACTIVE, path, userState)
				err = db.UpdateState(*newUser)
				if err != nil {
					return err
				}
			}

			secretBytes, publicKeyBytes := users[0].Bytes()
			// send a transfer tx between 2
			err, txHash := ValidateAndTransfer(db, bazooka, 2, 3, 1, hex.EncodeToString(secretBytes), hex.EncodeToString(publicKeyBytes))
			if err != nil {
				return err
			}
			fmt.Println("Transaction sent!", "Hash", txHash)
			return nil
		},
	}
	return cmd
}

// ValidateAndTransfer creates and sends a transfer transaction
func ValidateAndTransfer(db core.DB, bazooka core.Bazooka, fromIndex, toIndex, amount uint64, priv, pub string) (err error, txHash string) {
	from, err := db.GetStateByIndex(fromIndex)
	if err != nil {
		return
	}

	if !from.IsActive() {
		return ErrStateInActive, ""
	}

	to, err := db.GetStateByIndex(toIndex)
	if err != nil {
		return
	}

	if !to.IsActive() {
		return ErrStateInActive, ""
	}

	_, bal, nonce, token, err := bazooka.DecodeState(from.Data)
	if err != nil {
		return
	}

	if bal.Int64() <= int64(amount) {
		return ErrInvalidAmount, ""
	}

	txData, err := bazooka.EncodeTransferTx(int64(fromIndex), int64(toIndex), token.Int64(), nonce.Int64(), int64(amount), core.TX_TRANSFER_TYPE)
	if err != nil {
		return
	}

	tx := core.NewPendingTx(toIndex, fromIndex, core.TX_TRANSFER_TYPE, []byte(""), txData)
	tx.SignTx(priv, pub, common.Keccak256(tx.GetSignBytes()))
	tx.AssignHash()

	err = db.InsertTx(&tx)
	if err != nil {
		return
	}

	return nil, tx.TxHash
}
