// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	"github.com/ava-labs/libevm/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	topicArgs []string
	data      []byte
)

var eventCmd = &cobra.Command{
	Use:   "event --topics topic1,topic2 [--data data]",
	Short: "Parses a Teleporter log's topics and data",
	Long: `Given the topics and data of a Teleporter log, parses the log into
the corresponding Teleporter event. Topics are represented by a hash,
and data is the hex encoding of the bytes.`,
	Args: cobra.NoArgs,
	Run:  eventRun,
}

func eventRun(cmd *cobra.Command, args []string) {
	var topics []common.Hash
	for _, topic := range topicArgs {
		topics = append(topics, common.HexToHash(topic))
	}

	event, err := teleporterABI.EventByID(topics[0])
	cobra.CheckErr(err)

	out, err := teleportermessenger.FilterTeleporterEvents(topics, data, event.Name)
	cobra.CheckErr(err)
	logger.Info("Parsed Teleporter event", zap.String("name", event.Name), zap.String("event", out.String()))
	cmd.Println("Event command ran successfully for", event.Name)
}

func init() {
	rootCmd.AddCommand(eventCmd)
	eventCmd.PersistentFlags().StringSliceVar(&topicArgs, "topics", []string{}, "Topic hashes of the event")
	eventCmd.Flags().BytesHexVar(&data, "data", []byte{}, "Hex encoded data of the event")

	err := eventCmd.MarkPersistentFlagRequired("topics")
	cobra.CheckErr(err)
}
