package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	blswallet "github.com/kilic/bn254/bls"
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
			privKeyBytes, err := hex.DecodeString(privKey)
			if err != nil {
				return err
			}
			pubKey, err := flags.GetString(FlagPubKey)
			if err != nil {
				return err
			}
			pubkeyBytes, err := hex.DecodeString(pubKey)
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

			txHash, err := validateAndTransfer(db, bazooka, fromIndex, toIndex, amount, fee, privKeyBytes, pubkeyBytes)
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

				fmt.Println("Adding new account", "privkey", hex.EncodeToString(secretBytes), "publickey", hex.EncodeToString(publicKeyBytes))

				pubkeyIndex := uint64(i + 2)
				path, err := core.SolidityPathToNodePath(uint64(pubkeyIndex), params.MaxDepth)
				if err != nil {
					return err
				}

				// add accounts to tree
				acc, err := core.NewAccount(pubkeyIndex, publicKeyBytes, path)
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
			txHash, err := validateAndTransfer(db, bazooka, 2, 3, 1, 0, secretBytes, publicKeyBytes)
			if err != nil {
				return err
			}

			fmt.Println("Transaction sent!", "Hash", txHash)

			return nil
		},
	}
	return cmd
}

func dummyCreate2Transfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dummy-create2transfer",
		Short: "Sends a create2transfer transaction",
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

			for i := 0; i < 16; i++ {
				fmt.Println("Sending another tx", i)
				user1, err := wallet.NewWallet()
				if err != nil {
					return err
				}
				secretBytes, publicKeyBytes := user1.Bytes()
				pubkeyIndex := uint64(i + 2)
				path, err := core.SolidityPathToNodePath(uint64(pubkeyIndex), params.MaxDepth)
				if err != nil {
					return err
				}

				// add accounts to tree
				user1Acc, err := core.NewAccount(pubkeyIndex, publicKeyBytes, path)
				if err != nil {
					return err
				}

				err = db.UpdateAccount(*user1Acc)
				if err != nil {
					return err
				}

				// add accounts to state tree
				user1state, err := bazooka.EncodeState(pubkeyIndex, 10, 0, 1)
				if err != nil {
					return err
				}
				newUser := core.NewUserState(pubkeyIndex, core.STATUS_ACTIVE, path, user1state)
				err = db.UpdateState(*newUser)
				if err != nil {
					return err
				}

				user2, err := wallet.NewWallet()
				if err != nil {
					return err
				}

				_, publicKeyBytes2 := user2.Bytes()
				pubkey2, err := core.Pubkey(publicKeyBytes2).ToSol()
				if err != nil {
					return err
				}

				// send a transfer tx between 2
				txData, err := bazooka.EncodeCreate2TransferTxWithPub(int64(newUser.AccountID), pubkey2, 0, 1, 1, core.TX_CREATE_2_TRANSFER)
				if err != nil {
					return err
				}

				tx, err := core.NewPendingTx(newUser.AccountID, 0, core.TX_CREATE_2_TRANSFER, []byte(""), txData)
				if err != nil {
					return err
				}

				if err := signAndBroadcast(tx, secretBytes, publicKeyBytes, bazooka, db); err != nil {
					return err
				}

				fmt.Println("Transaction sent!", "Hash", tx.TxHash)
			}

			return nil
		},
	}
	return cmd
}

func dummyMassMigrate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dummy-massmigrate",
		Short: "Creates 2 accounts and creates a mass migrate",
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
				fmt.Println("Adding new account", "privkey", hex.EncodeToString(secretBytes), "publickey", publicKeyBytes)

				pubkeyIndex := uint64(i + 2)
				path, err := core.SolidityPathToNodePath(uint64(pubkeyIndex), params.MaxDepth)
				if err != nil {
					return err
				}

				// add accounts to tree
				acc, err := core.NewAccount(pubkeyIndex, publicKeyBytes, path)
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

			toSpoke := 10
			secretBytes, publicKeyBytes := users[0].Bytes()
			txData, err := bazooka.EncodeMassMigrationTx(int64(2), int64(toSpoke), int64(0), 0, int64(1), core.TX_MASS_MIGRATIONS)
			if err != nil {
				return err
			}

			tx, err := core.NewPendingTx(2, 0, core.TX_MASS_MIGRATIONS, []byte(""), txData)
			if err != nil {
				return err
			}

			if err = signAndBroadcast(tx, secretBytes, publicKeyBytes, bazooka, db); err != nil {
				return err
			}

			fmt.Println("Transaction sent!", "Hash", tx.TxHash)

			return nil
		},
	}
	return cmd
}

// validateAndTransfer creates and sends a transfer transaction
func validateAndTransfer(db core.DB, bazooka core.Bazooka, fromIndex, toIndex, amount, fee uint64, priv, pub []byte) (txHash string, err error) {
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

	opts := bind.CallOpts{From: config.OperatorAddress}
	txBytes, err := tx.GetSignBytes(bazooka)
	if err != nil {
		return
	}
	err = tx.SignTx(priv, pub, txBytes)
	if err != nil {
		return
	}

	// signature, err := core.BytesToSolSignature(tx.Signature)
	// if err != nil {
	// 	return
	// }

	pubkeyInt, err := core.Pubkey(pub).ToSol()
	if err != nil {
		return
	}

	newWallet, err := wallet.NewWallet()
	if err != nil {
		return
	}
	secret, pubkey := newWallet.Bytes()
	err = tx.SignTx(secret, pubkey, txBytes)
	if err != nil {
		return
	}

	sig, err := blswallet.SignatureKeyFromBytes(tx.Signature)
	if err != nil {
		fmt.Println("error while getting signature", err)
		return
	}

	pubkeyObj, err := blswallet.PublicKeyFromBytes(pubkey)
	if err != nil {
		fmt.Println("error while getting public key", err)
		return
	}
	pubkeyInt, err = core.Pubkey(pubkey).ToSol()
	if err != nil {
		return
	}

	solSignature, err := core.BytesToSolSignature(tx.Signature)
	if err != nil {
		return
	}

	valid, err := newWallet.VerifySignature(common.Keccak256(txBytes).Bytes(), *sig, *pubkeyObj)
	fmt.Println(valid, err)

	err = bazooka.SC.Transfer.VerifySingle(&opts, txBytes, pubkeyInt, solSignature, wallet.DefaultDomain)
	if err != nil {
		fmt.Println("error on validate", err)
		return
	}

	// err = bazooka.SC.Transfer.Validate(&opts, tx.Data, solSignature, pubkeyInt, wallet.DefaultDomain)
	// if err != nil {
	// 	return
	// }

	// if err = signAndBroadcast(tx, priv, pub, bazooka, db); err != nil {
	// 	return
	// }

	return tx.TxHash, nil
}

func signAndBroadcast(tx core.Tx, priv, pub []byte, bazooka core.Bazooka, db core.DB) (err error) {
	txBytes, err := tx.GetSignBytes(bazooka)
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

	fmt.Println("Sending new tx", tx.String())

	err = db.InsertTx(&tx)
	if err != nil {
		return err
	}
	return nil
}
