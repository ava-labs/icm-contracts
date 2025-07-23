package teleporter

import (
	"context"
	"math/big"

	testmessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/tests/TestMessenger"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

func RetrySuccessfulExecution(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy TestMessenger to L1s A and B
	//
	ctx := context.Background()

	_, l1ATestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(l1AInfo),
		l1AInfo,
	)
	testMessengerContractAddressB, l1BTestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(l1BInfo),
		l1BInfo,
	)

	//
	// Call the test messenger contract on L1 A
	//
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := l1ATestMessenger.SendMessage(
		optsA,
		l1BInfo.BlockchainID,
		testMessengerContractAddressB,
		fundedAddress,
		big.NewInt(0),
		testmessenger.SendMessageRequiredGas,
		message,
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

	teleporterMessageID := event.MessageID

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	//
	// Relay the message to the destination
	//
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
	receiveEvent, err :=
		utils.GetEventFromLogs(receipt.Logs, teleporter.TeleporterMessenger(l1BInfo).ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	deliveredTeleporterMessage := receiveEvent.Message

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Verify we received the expected string
	//
	_, currMessage, err := l1BTestMessenger.GetCurrentMessage(&bind.CallOpts{}, l1AInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))

	//
	// Attempt to retry message execution, which should fail
	//
	optsB, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1BInfo.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err = teleporter.TeleporterMessenger(l1BInfo).RetryMessageExecution(
		optsB,
		l1AInfo.BlockchainID,
		deliveredTeleporterMessage,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(tx).Should(BeNil())
}
