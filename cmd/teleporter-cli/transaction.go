// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"encoding/json"

	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/eth/tracers"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/spf13/cobra"
)

const (
	ICMPrecompileAddressHex = "0x0200000000000000000000000000000000000005"
)

var (
	debug             bool
	rpcEndpoint       string
	teleporterAddress common.Address
	client            ethclient.Client
)

var transactionCmd = &cobra.Command{
	Use:   "transaction --rpc RPC_URL --teleporter-address CONTRACT_ADDRESS TRANSACTION_HASH",
	Short: "Parses relevant Teleporter logs from a transaction",
	Long: `Given a transaction this command looks through the transaction's receipt
for TeleporterMessenger and ICM log events. When corresponding log events are found,
the command parses to log event fields to a more human readable format. Optionally pass -d 
or --debug for extra transaction output. This may require enabling debug enpoints on your RPC node`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(args[0])
		if debug {
			printTransaction(cmd, txHash)
			traceTransaction(cmd, txHash)
		}
		checkReceipt(cmd, txHash)
		cmd.Println("Transaction command ran successfully")
	},
}

func checkReceipt(cmd *cobra.Command, txHash common.Hash) {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	cobra.CheckErr(err)

	ICMPrecompileAddress := common.HexToAddress(ICMPrecompileAddressHex)
	for _, log := range receipt.Logs {
		switch log.Address {
		case teleporterAddress:
			printTeleporterLogs(cmd, log)
		case ICMPrecompileAddress:
			printICMLogs(cmd, log)
		}
	}
}

func printTeleporterLogs(cmd *cobra.Command, log *types.Log) {
	logJson, err := json.MarshalIndent(log, "", "  ")
	cobra.CheckErr(err)

	cmd.Println("Teleporter Log:\n" + string(logJson) + "\n")

	event, err := teleporterABI.EventByID(log.Topics[0])
	cobra.CheckErr(err)

	out, err := teleportermessenger.FilterTeleporterEvents(log.Topics, log.Data, event.Name)
	cobra.CheckErr(err)

	cmd.Println(event.Name + " Log:")
	cmd.Println(out.String() + "\n")
}

func printICMLogs(cmd *cobra.Command, log *types.Log) {
	logJson, err := json.MarshalIndent(log, "", "  ")
	cobra.CheckErr(err)

	cmd.Println("ICM Log:\n" + string(logJson) + "\n")

	unsignedMsg, err := warp.UnpackSendWarpEventDataToMessage(log.Data)
	cobra.CheckErr(err)
	cmd.Println("ICM Message ID: " + unsignedMsg.ID().Hex())

	icmPayload, err := warpPayload.ParseAddressedCall(unsignedMsg.Payload)
	cobra.CheckErr(err)

	icmPayloadJson, err := json.MarshalIndent(icmPayload, "", "  ")
	cobra.CheckErr(err)
	cmd.Println("ICM Payload:")
	cmd.Println(string(icmPayloadJson))

	teleporterMessage := teleportermessenger.TeleporterMessage{}
	err = teleporterMessage.Unpack(icmPayload.Payload)
	cobra.CheckErr(err)

	cmd.Println("Teleporter Message:")
	cmd.Println(teleporterMessage.String())
}

func traceTransaction(cmd *cobra.Command, txHash common.Hash) {
	var result interface{}
	ct := "callTracer"
	err := client.Client().Call(&result, "debug_traceTransaction", txHash.String(), tracers.TraceConfig{Tracer: &ct})
	if err != nil {
		cmd.PrintErr("Error calling debug_traceTransaction: " + err.Error())
		return
	}
	json, err := json.MarshalIndent(result, "", "  ")
	cobra.CheckErr(err)

	cmd.Println("Transaction Trace:\n" + string(json) + "\n")
}

func printTransaction(cmd *cobra.Command, txHash common.Hash) {
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	cobra.CheckErr(err)
	json, err := json.MarshalIndent(tx, "", "  ")
	cobra.CheckErr(err)

	cmd.Println("Transaction:\n" + string(json) + "\n")
}

func init() {
	rootCmd.AddCommand(transactionCmd)
	transactionCmd.PersistentFlags().StringVar(&rpcEndpoint, "rpc", "", "RPC endpoint to connect to the node")
	address := transactionCmd.PersistentFlags().StringP("teleporter-address", "t", "", "Teleporter contract address")
	err := transactionCmd.MarkPersistentFlagRequired("rpc")
	cobra.CheckErr(err)
	err = transactionCmd.MarkPersistentFlagRequired("teleporter-address")
	cobra.CheckErr(err)
	transactionCmd.Flags().BoolVarP(&debug, "debug", "d", false, "default: false.")
	transactionCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return transactionPreRunE(cmd, args, address)
	}
}

func transactionPreRunE(cmd *cobra.Command, args []string, address *string) error {
	// Run the persistent pre-run function of the root command if it exists.
	if err := callPersistentPreRunE(cmd, args); err != nil {
		return err
	}
	teleporterAddress = common.HexToAddress(*address)
	c, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return err
	}

	client = c
	return nil
}
