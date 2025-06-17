package teleporter

import (
	"context"
	"math/big"

	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/log"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

func ResubmitAlteredMessage(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	// Send a transaction to L1 A to issue an ICM Message from the Teleporter contract to L1 B
	ctx := context.Background()

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: l1BInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	receipt, messageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, teleporter.TeleporterMessenger(l1AInfo), l1AInfo, l1BInfo, sendCrossChainMessageInput, fundedKey)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Relay the message to the destination
	receipt = teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		l1AInfo,
		l1BInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	log.Info("Checking the message was received on the destination")
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Get the Teleporter message from receive event
	event, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(l1BInfo).ParseReceiveCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(event.MessageID[:]).Should(Equal(messageID[:]))
	teleporterMessage := event.Message

	// Alter the message
	alteredMessage := make([]byte, len(teleporterMessage.Message))
	copy(alteredMessage, teleporterMessage.Message)
	alteredMessage[0] = ^alteredMessage[0]
	Expect(alteredMessage[:]).ShouldNot(Equal(teleporterMessage.Message[:]))
	teleporterMessage.Message = alteredMessage

	// Resubmit the altered message
	log.Info("Submitting the altered Teleporter message on the source chain")
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := teleporter.TeleporterMessenger(l1AInfo).RetrySendCrossChainMessage(opts, teleporterMessage)
	Expect(err).ShouldNot(BeNil())

	// We expect the tx to be nil because the ICM message failed verification, which happens in the predicate
	// In that case, the block is never built, and the transaction is never mined
	Expect(tx).Should(BeNil())
}
