/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/spf13/cobra"
)

var (
	logger        logging.Logger
	teleporterABI *abi.ABI
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "teleporter-cli",
	Short: "A CLI that integrates with the Teleporter protocol",
	Long: `A CLI that integrates with the Teleporter protocol, and allows you
	to debug Teleporter on chain activity. The CLI can help decode Teleporter and Warp events,
	as well as parsing Teleporter messages.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	logLevelArg := rootCmd.PersistentFlags().StringP("log", "l", "", "Log level")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if *logLevelArg == "" {
			*logLevelArg = logging.Info.LowerString()
		}

		logLevel, err := logging.ToLevel(*logLevelArg)
		if err != nil {
			return err
		}
		logger = logging.NewLogger(
			"teleporter-cli",
			logging.NewWrappedCore(
				logLevel,
				os.Stdout,
				logging.JSON.ConsoleEncoder(),
			),
		)
		abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
		if err != nil {
			return err
		}
		teleporterABI = abi
		return nil
	}
}
