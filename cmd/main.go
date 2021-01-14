package main

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BOPR/config"
	"github.com/BOPR/wallet"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	WithConfigPathFlag = "config-path"
)

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}

func main() {
	myFigure := figure.NewColorFigure("Hubble", "", "red", true)
	myFigure.Print()
	rootCmd := &cobra.Command{
		Use:   "hubble",
		Short: "Optimistic Rollup Daemon (server)",
	}

	rootCmd.PersistentFlags().String(
		WithConfigPathFlag,
		"",
		"Config file path (default ./config.toml)",
	)

	// bind with-heimdall-config config with root cmd
	err := viper.BindPFlag(
		WithConfigPathFlag,
		rootCmd.Flags().Lookup(WithConfigPathFlag),
	)
	if err != nil {
		fmt.Println("BindPFlag Error", err)
		return
	}

	rootCmd.AddCommand(initCmd())
	rootCmd.AddCommand(startCmd())
	rootCmd.AddCommand(startRestServerCmd())
	rootCmd.AddCommand(sendTransferTx())
	rootCmd.AddCommand(dummyTransfer())
	rootCmd.AddCommand(dummyCreate2Transfer())
	rootCmd.AddCommand(dummyMassMigrate())
	rootCmd.AddCommand(createDatabase())
	rootCmd.AddCommand(createUsers())
	rootCmd.AddCommand(viewState())
	rootCmd.AddCommand(migrationCmd)

	if err := viper.BindPFlag(WithConfigPathFlag, rootCmd.Flags().Lookup(WithConfigPathFlag)); err != nil {
		fmt.Println("Error binding flags with viper", "Error", err)
		return
	}

	executor := Executor{rootCmd, os.Exit}
	if err = executor.Command.Execute(); err != nil {
		fmt.Println("Error while executing command", err)
		return
	}
}

type UserList struct {
	Users []User `json:"users"`
}

type User struct {
	PublicKey string `json:"pubkey"`
	PrivKey   string `json:"privkey"`
}

// createUsers creates the database
func createUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-users",
		Short: "Create users to be used in simulations",
		RunE: func(cmd *cobra.Command, args []string) error {
			userCount, err := cmd.Flags().GetInt(FlagDatabaseName)
			if err != nil {
				return err
			}
			var users []User
			for i := 0; i < userCount; i++ {
				newWallet, err := wallet.NewWallet()
				if err != nil {
					return err
				}
				secretBytes, publicKeyBytes := newWallet.Bytes()
				publicKey := hex.EncodeToString(publicKeyBytes)
				secretKey := hex.EncodeToString(secretBytes)
				newUser := User{PublicKey: publicKey, PrivKey: secretKey}
				users = append(users, newUser)
			}
			bz, err := json.MarshalIndent(UserList{Users: users}, "", " ")
			if err != nil {
				return err
			}
			return ioutil.WriteFile("users.json", bz, 0644)
		},
	}

	cmd.Flags().Int(FlagUserCount, 0, "--count=<user-count>")
	err := cmd.MarkFlagRequired(FlagUserCount)
	if err != nil {
		panic(err)
	}
	return cmd
}

// createDatabase creates the database
func createDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-database",
		Short: "Create a new database",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbName, err := cmd.Flags().GetString(FlagDatabaseName)
			if err != nil {
				return err
			}
			cfg, err := config.ParseConfig()
			if err != nil {
				return err
			}
			splitStrings := strings.Split(cfg.FormattedDBURL(), "/")
			connectionString := []string{splitStrings[0], "/"}
			dbNew, err := sql.Open("mysql", strings.Join(connectionString, ""))
			if err != nil {
				return err
			}
			defer dbNew.Close()
			_, err = dbNew.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName))
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP(FlagDatabaseName, "", "", "--dbname=<database-name>")
	err := cmd.MarkFlagRequired(FlagDatabaseName)
	if err != nil {
		panic(err)
	}
	return cmd
}
