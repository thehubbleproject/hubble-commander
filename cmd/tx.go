package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
	"github.com/BOPR/wallet"
	"github.com/spf13/cobra"
)

var (
	ErrInvalidAmount = errors.New("Invalid amount")
	ErrStateInActive = errors.New("User state inactive")
)

//  sendTransferTx generated init command to initialise the config file
func sendTransferTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Transfers assets between 2 accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()

			toIndex, err := flags.GetUint64(FlagToID)
			if err != nil {
				return err
			}
			fromIndex, err := flags.GetUint64(FlagFromID)
			if err != nil {
				return err
			}
			privKey, err := flags.GetString(FlagPrivKey)
			if err != nil {
				return err
			}
			pubKey, err := flags.GetString(FlagPubKey)
			if err != nil {
				return err
			}
			amount, err := flags.GetUint64(FlagAmount)
			if err != nil {
				return err
			}
			fee, err := flags.GetUint64(FlagFee)
			if err != nil {
				return err
			}

			db, err := core.NewDB()
			if err != nil {
				return err
			}
			defer db.Close()

			bazooka, err := core.NewPreLoadedBazooka()
			if err != nil {
				return err
			}

			txHash, err := validateAndTransfer(db, bazooka, fromIndex, toIndex, amount, fee, privKey, pubKey)
			if err != nil {
				return err
			}

			fmt.Println("Transaction submitted successfully", "hash", txHash)
			return nil
		},
	}
	cmd.Flags().StringP(FlagToID, "", "", "--to=<to-account>")
	cmd.Flags().StringP(FlagFromID, "", "", "--from=<from-account>")
	cmd.Flags().StringP(FlagPubKey, "", "", "--pubkey=<pubkey>")
	cmd.Flags().StringP(FlagPrivKey, "", "", "--privkey=<privkey>")
	cmd.Flags().StringP(FlagAmount, "", "", "--amount=<amount>")
	err := cmd.MarkFlagRequired(FlagToID)
	common.PanicIfError(err)
	err = cmd.MarkFlagRequired(FlagFromID)
	common.PanicIfError(err)
	err = cmd.MarkFlagRequired(FlagPubKey)
	common.PanicIfError(err)
	err = cmd.MarkFlagRequired(FlagPrivKey)
	common.PanicIfError(err)
	err = cmd.MarkFlagRequired(FlagAmount)
	common.PanicIfError(err)
	return cmd
}

func dummyTransfer() *cobra.Command {
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
				publicKey, err := core.NewPubkeyFromBytes(publicKeyBytes)
				if err != nil {
					return err
				}
				pubkeyStr, err := publicKey.String()
				if err != nil {
					return err
				}
				fmt.Println("Adding new account", "privkey", hex.EncodeToString(secretBytes), "publickey", publicKey)

				pubkeyIndex := uint64(i + 2)
				path, err := core.SolidityPathToNodePath(uint64(pubkeyIndex), params.MaxDepth)
				if err != nil {
					return err
				}

				// add accounts to tree
				acc, err := core.NewAccount(pubkeyIndex, pubkeyStr, path)
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
			txHash, err := validateAndTransfer(db, bazooka, 2, 3, 1, 0, hex.EncodeToString(secretBytes), hex.EncodeToString(publicKeyBytes))
			if err != nil {
				return err
			}

			fmt.Println("Transaction sent!", "Hash", txHash)

			return nil
		},
	}
	return cmd
}

// validateAndTransfer creates and sends a transfer transaction
func validateAndTransfer(db core.DB, bazooka core.Bazooka, fromIndex, toIndex, amount, fee uint64, priv, pub string) (txHash string, err error) {
	from, err := db.GetStateByIndex(fromIndex)
	if err != nil {
		return
	}

	if !from.IsActive() {
		return "", ErrStateInActive
	}

	to, err := db.GetStateByIndex(toIndex)
	if err != nil {
		return
	}

	if !to.IsActive() {
		return "", ErrStateInActive
	}

	_, bal, nonce, _, err := bazooka.DecodeState(from.Data)
	if err != nil {
		return
	}

	if bal.Int64() <= int64(amount+fee) {
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
	err = tx.SignTx(priv, pub, common.Keccak256(tx.GetSignBytes()))
	if err != nil {
		return
	}
	err = tx.AssignHash()
	if err != nil {
		return
	}

	fmt.Println("Sending new tx", tx.String())

	err = db.InsertTx(&tx)
	if err != nil {
		return
	}

	return tx.TxHash, nil
}
