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
			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}
			wallet, err := wallet.SecretToWallet(privKeyBytes, pubkeyBytes, cfg.AppID)
			if err != nil {
				return err
			}
			secretBytes, pubkeyBytes := wallet.Bytes()

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

			wallet, err := wallet.SecretToWallet(privKeyBytes, pubkeyBytes, cfg.AppID)
			if err != nil {
				return err
			}
			secretBytes, pubkeyBytes := wallet.Bytes()

			from, err := DBI.GetStateByIndex(fromIndex)
			if err != nil {
				return err
			}
			_, _, nonce, token, err := bazooka.DecodeState(from.Data)
			if err != nil {
				return err
			}
			txData, err := bazooka.EncodeCreate2TransferTxWithPub(int64(fromIndex), toPubkey, int64(fee), nonce.Int64(), int64(amount), core.TX_CREATE_2_TRANSFER)
			if err != nil {
				return err
			}

			tx, err := core.NewPendingTx(txData, nil, fromIndex, nonce.Uint64(), 0, token.Uint64(), core.TX_CREATE_2_TRANSFER)
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

	_, bal, _, token, err := bazooka.DecodeState(from.Data)
	if err != nil {
		return
	}

	pendingNonce, err := DBI.GetPendingNonce(fromIndex)
	if err != nil {
		return
	}

	if bal.Int64() <= int64(amount+fee) {
		return "", ErrInvalidAmount
	}

	txData, err := bazooka.EncodeTransferTx(int64(fromIndex), int64(toIndex), int64(fee), int64(pendingNonce+1), int64(amount), core.TX_TRANSFER_TYPE)
	if err != nil {
		return
	}

	tx, err := core.NewPendingTx(txData, nil, fromIndex, pendingNonce+1, 0, token.Uint64(), core.TX_TRANSFER_TYPE)
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

	err = tx.SignTx(priv, pub, txBytes, b.Cfg.AppID)
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
