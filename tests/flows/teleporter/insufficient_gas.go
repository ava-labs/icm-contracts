package teleporter

import (
	"context"
	"math/big"

	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

func InsufficientGas(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Deploy TestMessenger to L1s A
	_, l1ATestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(l1AInfo),
		l1AInfo,
	)
	// Deploy TestMessenger to L1s B
	testMessengerContractB, l1BTestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(l1BInfo),
		l1BInfo,
	)

	// Send message from L1A to L1B with 0 execution gas, which should fail to execute
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(
		fundedKey,
		l1AInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())
	tx, err := l1ATestMessenger.SendMessage(
		optsA, l1BInfo.BlockchainID, testMessengerContractB, fundedAddress, big.NewInt(0), big.NewInt(0), message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, l1AInfo, tx.Hash())

	event, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(l1AInfo).ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(l1BInfo.BlockchainID[:]))

	messageID := event.MessageID

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Relay message from L1 A to L1 B
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

	// Check Teleporter message received on the destination
	delivered, err :=
		teleporter.TeleporterMessenger(l1BInfo).MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check message execution failed event
	failedMessageExecutionEvent, err := utils.GetEventFromLogs(
		receipt.Logs, teleporter.TeleporterMessenger(l1BInfo).ParseMessageExecutionFailed,
	)
	Expect(err).Should(BeNil())
	Expect(failedMessageExecutionEvent.MessageID[:]).Should(Equal(messageID[:]))
	Expect(failedMessageExecutionEvent.SourceBlockchainID[:]).Should(Equal(l1AInfo.BlockchainID[:]))

	// Retry message execution. This will execute the message with as much gas as needed
	// (up to the transaction gas limit), rather than using the required gas specified in the message itself.
	receipt = utils.RetryMessageExecutionAndWaitForAcceptance(
		ctx,
		l1AInfo.BlockchainID,
		teleporter.TeleporterMessenger(l1BInfo),
		l1BInfo,
		failedMessageExecutionEvent.Message,
		fundedKey,
	)
	executedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(l1BInfo).ParseMessageExecuted,
	)
	Expect(err).Should(BeNil())
	Expect(executedEvent.MessageID[:]).Should(Equal(messageID[:]))
	Expect(executedEvent.SourceBlockchainID[:]).Should(Equal(l1AInfo.BlockchainID[:]))

	//
	// Verify we received the expected string
	//
	_, currMessage, err := l1BTestMessenger.GetCurrentMessage(&bind.CallOpts{}, l1AInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))
}
