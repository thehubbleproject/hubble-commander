package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	db "github.com/BOPR/db"
	"github.com/BOPR/wallet"
	"github.com/spf13/cobra"
)

var (
	ErrInvalidAmount = errors.New("Invalid amount")
)

// sendTransferTx generated init command to initialise the config file
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
			wallet, err := wallet.SecretToWallet(privKeyBytes, pubkeyBytes)
			if err != nil {
				return err
			}
			secretBytes, pubkeyBytes := wallet.Bytes()

			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}

			DBI, err := db.NewDB(cfg)
			if err != nil {
				return err
			}
			defer DBI.Close()

			bazooka, err := bazooka.NewPreLoadedBazooka(cfg)
			if err != nil {
				return err
			}

			txHash, err := validateAndTransfer(&DBI, &bazooka, fromIndex, toIndex, amount, fee, secretBytes, pubkeyBytes)
			if err != nil {
				return err
			}

			fmt.Println("Transaction submitted successfully", "hash", txHash)
			return nil
		},
	}
	cmd.Flags().Uint64P(FlagToID, "", 0, "--to=<to-account>")
	cmd.Flags().Uint64P(FlagFee, "", 0, "--fee=<fee>")
	cmd.Flags().Uint64P(FlagFromID, "", 0, "--from=<from-account>")
	cmd.Flags().StringP(FlagPubKey, "", "", "--pubkey=<pubkey>")
	cmd.Flags().StringP(FlagPrivKey, "", "", "--privkey=<privkey>")
	cmd.Flags().Uint64P(FlagAmount, "", 0, "--amount=<amount>")
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

// sendCreate2TransferTx generated init command to initialise the config file
func sendCreate2TransferTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "c2t",
		Short: "Transfers assets between sender to non existent receiver",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}
			DBI, err := db.NewDB(cfg)
			if err != nil {
				return err
			}
			defer DBI.Close()
			bazooka, err := bazooka.NewPreLoadedBazooka(cfg)
			if err != nil {
				return err
			}
			flags := cmd.Flags()
			fromIndex, err := flags.GetUint64(FlagFromID)
			if err != nil {
				return err
			}

			// fetching to pubkey
			toPubkeyStr, err := flags.GetString(FlagToPubkey)
			if err != nil {
				return err
			}
			toPubHexBz, err := hex.DecodeString(toPubkeyStr)
			if err != nil {
				return err
			}
			toPubkey, err := core.Pubkey(toPubHexBz).ToSol()
			if err != nil {
				return err
			}
			// fetching from pubkey
			pubKey, err := flags.GetString(FlagPubKey)
			if err != nil {
				return err
			}
			pubkeyBytes, err := hex.DecodeString(pubKey)
			if err != nil {
				return err
			}

			// fetching priv key
			privKey, err := flags.GetString(FlagPrivKey)
			if err != nil {
				return err
			}
			privKeyBytes, err := hex.DecodeString(privKey)
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

			wallet, err := wallet.SecretToWallet(privKeyBytes, pubkeyBytes)
			if err != nil {
				return err
			}
			secretBytes, pubkeyBytes := wallet.Bytes()

			from, err := DBI.GetStateByIndex(fromIndex)
			if err != nil {
				return err
			}
			_, _, nonce, _, err := bazooka.DecodeState(from.Data)
			if err != nil {
				return err
			}
			txData, err := bazooka.EncodeCreate2TransferTxWithPub(int64(fromIndex), toPubkey, int64(fee), nonce.Int64(), int64(amount), core.TX_CREATE_2_TRANSFER)
			if err != nil {
				return err
			}

			tx, err := core.NewPendingTx(fromIndex, 0, core.TX_CREATE_2_TRANSFER, []byte(""), txData)
			if err != nil {
				return err
			}

			if err = signAndBroadcast(&bazooka, &DBI, tx, secretBytes, pubkeyBytes); err != nil {
				return err
			}

			fmt.Println("Transaction submitted successfully", "hash", tx.TxHash)

			return nil
		},
	}
	cmd.Flags().StringP(FlagToPubkey, "", "", "--to-pub=<to-pubkey>")
	cmd.Flags().Uint64P(FlagFee, "", 0, "--fee=<fee>")
	cmd.Flags().Uint64P(FlagFromID, "", 0, "--from=<from-account>")
	cmd.Flags().StringP(FlagPubKey, "", "", "--pubkey=<pubkey>")
	cmd.Flags().StringP(FlagPrivKey, "", "", "--privkey=<privkey>")
	cmd.Flags().Uint64P(FlagAmount, "", 0, "--amount=<amount>")
	err := cmd.MarkFlagRequired(FlagToPubkey)
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
			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}
			DBI, err := db.NewDB(cfg)
			if err != nil {
				return err
			}
			defer DBI.Close()

			bazooka, err := bazooka.NewPreLoadedBazooka(cfg)
			if err != nil {
				return err
			}
			params, err := DBI.GetParams()
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
				nodeType, err := DBI.FindNodeType(path)
				if err != nil {
					panic(err)
				}
				// add accounts to tree
				acc, err := core.NewAccount(pubkeyIndex, publicKeyBytes, path, nodeType)
				if err != nil {
					return err
				}
				err = DBI.UpdateAccount(*acc)
				if err != nil {
					return err
				}
				// add accounts to state tree
				userState, err := bazooka.EncodeState(pubkeyIndex, 10, 0, 1)
				if err != nil {
					return err
				}
				newUser := core.NewUserState(pubkeyIndex, path, userState)
				err = DBI.UpdateState(*newUser)
				if err != nil {
					return err
				}
			}

			secretBytes, publicKeyBytes := users[0].Bytes()

			// send a transfer tx between 2
			txHash, err := validateAndTransfer(&DBI, &bazooka, 2, 3, 1, 0, secretBytes, publicKeyBytes)
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
			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}
			DBI, err := db.NewDB(cfg)
			if err != nil {
				return err
			}
			defer DBI.Close()
			bazooka, err := bazooka.NewPreLoadedBazooka(cfg)
			if err != nil {
				return err
			}
			params, err := DBI.GetParams()
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

				nodeType, err := DBI.FindNodeType(path)
				if err != nil {
					panic(err)
				}

				// add accounts to tree
				user1Acc, err := core.NewAccount(pubkeyIndex, publicKeyBytes, path, nodeType)
				if err != nil {
					return err
				}

				err = DBI.UpdateAccount(*user1Acc)
				if err != nil {
					return err
				}

				// add accounts to state tree
				user1state, err := bazooka.EncodeState(pubkeyIndex, 10, 0, 1)
				if err != nil {
					return err
				}
				newUser := core.NewUserState(pubkeyIndex, path, user1state)
				err = DBI.UpdateState(*newUser)
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

				if err := signAndBroadcast(&bazooka, &DBI, tx, secretBytes, publicKeyBytes); err != nil {
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
			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}
			DBI, err := db.NewDB(cfg)
			if err != nil {
				return err
			}
			defer DBI.Close()

			bazooka, err := bazooka.NewPreLoadedBazooka(cfg)
			if err != nil {
				return err
			}
			params, err := DBI.GetParams()
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

				nodeType, err := DBI.FindNodeType(path)
				if err != nil {
					panic(err)
				}
				// add accounts to tree
				acc, err := core.NewAccount(pubkeyIndex, publicKeyBytes, path, nodeType)
				if err != nil {
					return err
				}
				err = DBI.UpdateAccount(*acc)
				if err != nil {
					return err
				}
				// add accounts to state tree
				userState, err := bazooka.EncodeState(pubkeyIndex, 10, 0, 1)
				if err != nil {
					return err
				}
				newUser := core.NewUserState(pubkeyIndex, path, userState)
				err = DBI.UpdateState(*newUser)
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

			if err = signAndBroadcast(&bazooka, &DBI, tx, secretBytes, publicKeyBytes); err != nil {
				return err
			}

			fmt.Println("Transaction sent!", "Hash", tx.TxHash)

			return nil
		},
	}
	return cmd
}

// validateAndTransfer creates and sends a transfer transaction
func validateAndTransfer(DBI *db.DB, bazooka *bazooka.Bazooka, fromIndex, toIndex, amount, fee uint64, priv, pub []byte) (txHash string, err error) {
	from, err := DBI.GetStateByIndex(fromIndex)
	if err != nil {
		return
	}

	_, err = DBI.GetStateByIndex(toIndex)
	if err != nil {
		return
	}

	_, bal, nonce, _, err := bazooka.DecodeState(from.Data)
	if err != nil {
		return
	}
	fmt.Println("nonce here", nonce)
	// nonce = big.NewInt(0)

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

	if err = signAndBroadcast(bazooka, DBI, tx, priv, pub); err != nil {
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

	fmt.Println("Sending new tx", tx.String())
	err = DBI.InsertTx(&tx)
	if err != nil {
		return err
	}
	return nil
}
