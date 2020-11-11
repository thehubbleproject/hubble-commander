package main

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/wallet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	WithConfigPathFlag = "config-path"
	ConfigFileName     = "config"
)

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "hubble",
		Short: "Optimistic Rollup Daemon (server)",
	}

	// add new persistent flag for heimdall-config
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
	rootCmd.AddCommand(resetCmd())
	rootCmd.AddCommand(addGenesisAcccountsCmd())
	rootCmd.AddCommand(sendTransferTx())
	rootCmd.AddCommand(dummyTransfer())
	rootCmd.AddCommand(createDatabase())
	rootCmd.AddCommand(createUsers())
	rootCmd.AddCommand(migrationCmd)

	executor := Executor{rootCmd, os.Exit}
	if err = executor.Command.Execute(); err != nil {
		fmt.Println("Error while executing command", err)
		return
	}
}

// resetCmd resets all the collections
func resetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "reset database",
		Run: func(cmd *cobra.Command, args []string) {
			err := config.ParseAndInitGlobalConfig("")
			common.PanicIfError(err)
			// TODO fix this command for mysql database
			// create new DB instance
			// dbInstance, err := db.NewDB()
			// defer dbInstance.Close()
			// common.PanicIfError(err)
			// fmt.Println("Resetting database", "db", common.DATABASE)
			// err = dbInstance.MgoSession.DropDatabase(common.DATABASE)
			// common.PanicIfError(err)
		},
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
			var users []User
			for i := 0; i < 2; i++ {
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
	return cmd
}

func addGenesisAcccountsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add-gen-accounts",
		Short: "Adds the accounts present in genesis account to the contract",
		Run: func(cmd *cobra.Command, args []string) {
			viperObj := viper.New()
			dir, err := os.Getwd()
			common.PanicIfError(err)

			viperObj.SetConfigName(ConfigFileName) // name of config file (without extension)
			viperObj.AddConfigPath(dir)
			err = viperObj.ReadInConfig()
			common.PanicIfError(err)

			var cfg config.Configuration
			if err = viperObj.UnmarshalExact(&cfg); err != nil {
				common.PanicIfError(err)
			}
			// init global config
			config.GlobalCfg = cfg
		},
	}
}

// createDatabase creates the database
func createDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-database",
		Short: "Create a new database",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.ParseAndInitGlobalConfig(""); err != nil {
				return err
			}
			splitStrings := strings.Split(config.GlobalCfg.FormattedDBURL(), "/")
			connectionString := []string{splitStrings[0], "/"}
			dbNew, err := sql.Open("mysql", strings.Join(connectionString, ""))
			if err != nil {
				return err
			}
			defer dbNew.Close()
			_, err = dbNew.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", "hubble"))
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP(FlagDatabaseName, "", "", "--dbname=<database-name>")
	// cmd.MarkFlagRequired(FlagDatabaseName)
	return cmd
}
