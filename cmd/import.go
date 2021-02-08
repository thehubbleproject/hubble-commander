package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BOPR/config"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/spf13/cobra"
)

const (
	defaultKeystorePath = ".keystore/"
)

func importCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "import",
		Short: "Import keystore for operator",
		RunE: func(cmd *cobra.Command, args []string) error {
			keyfile := args[0]
			passphrase, err := getPassPhrase("Unlock the keystore with the password", false)
			if err != nil {
				return err
			}

			keystoreBytes, err := ioutil.ReadFile(keyfile)
			if err != nil {
				return err
			}

			ks := keystore.NewKeyStore(defaultKeystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

			newPassphrase, err := getPassPhrase("Please give a new password", true)
			if err != nil {
				return err
			}
			acct, err := ks.Import(keystoreBytes, passphrase, newPassphrase)
			if err != nil {
				return err
			}
			fmt.Printf("Operator Address: {%x}\n", acct.Address)
			cfg, err := config.ParseConfig()
			cfg.KeystorePassphrase = newPassphrase
			config.WriteConfigFile("./config.toml", &cfg)

			return nil
		},
	}
}

func getPassPhrase(text string, confirmation bool) (string, error) {
	if text != "" {
		fmt.Println(text)
	}
	password, err := prompt.Stdin.PromptPassword("Password: ")
	if err != nil {
		return "", fmt.Errorf("Failed to read password: %v", err)
	}
	if confirmation {
		confirm, err := prompt.Stdin.PromptPassword("Repeat password: ")
		if err != nil {
			return "", fmt.Errorf("Failed to read password confirmation: %v", err)
		}
		if password != confirm {
			return "", fmt.Errorf("Passwords do not match")
		}
	}
	return password, nil
}
