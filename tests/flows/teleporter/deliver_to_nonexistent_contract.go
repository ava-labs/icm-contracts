package teleporter

import (
	"context"
	"math/big"

	testmessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/tests/TestMessenger"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/log"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

func DeliverToNonExistentContract(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	deployerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	deployerAddress := crypto.PubkeyToAddress(deployerKey.PublicKey)

	//
	// Fund the deployer address on L1 B
	//
	ctx := context.Background()
	log.Info("Funding address on L1 B", "address", deployerAddress.Hex())

	fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
	fundDeployerTx := utils.CreateNativeTransferTransaction(
		ctx, l1BInfo, fundedKey, deployerAddress, fundAmount,
	)
	utils.SendTransactionAndWaitForSuccess(ctx, l1BInfo, fundDeployerTx)

	//
	// Deploy ExampleMessenger to L1 A, but not to L1 B
	// Send a message that should fail to be executed on L1 B
	//
	log.Info("Deploying ExampleMessenger to L1 A")
	_, L1AExampleMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(l1AInfo),
		l1AInfo,
	)

	// Derive the eventual address of the destination contract on L1 B
	nonce, err := l1BInfo.RPCClient.NonceAt(ctx, deployerAddress, nil)
	Expect(err).Should(BeNil())
	destinationContractAddress := crypto.CreateAddress(deployerAddress, nonce)

	//
	// Call the example messenger contract on L1 A
	//
	log.Info("Calling ExampleMessenger on L1 A")
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := L1AExampleMessenger.SendMessage(
		optsA,
		l1BInfo.BlockchainID,
		destinationContractAddress,
		common.BigToAddress(common.Big0),
		big.NewInt(0),
		testmessenger.SendMessageRequiredGas,
		message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, l1AInfo, tx.Hash())

	sendEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(l1AInfo).ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(sendEvent.DestinationBlockchainID[:]).Should(Equal(l1BInfo.BlockchainID[:]))

	teleporterMessageID := sendEvent.MessageID

	//
	// Relay the message to the destination
	//

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	log.Info("Relaying the message to the destination")
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
		utils.GetEventFromLogs(receipt.Logs, teleporter.TeleporterMessenger(l1AInfo).ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	//
	// Check that the message was successfully relayed
	//
	log.Info("Checking the message was successfully relayed")
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(
		&bind.CallOpts{},
		teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Check that the message was not successfully executed
	//
	log.Info("Checking the message was not successfully executed")
	executionFailedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(l1AInfo).ParseMessageExecutionFailed,
	)
	Expect(err).Should(BeNil())
	Expect(executionFailedEvent.MessageID).Should(Equal(receiveEvent.MessageID))

	//
	// Deploy the contract on L1 B
	//
	log.Info("Deploying the contract on L1 B")
	exampleMessengerContractB, L1BExampleMessenger := utils.DeployTestMessenger(
		ctx,
		deployerKey,
		deployerAddress,
		teleporter.TeleporterRegistryAddress(l1BInfo),
		l1BInfo,
	)

	// Confirm that it was deployed at the expected address
	Expect(exampleMessengerContractB).Should(Equal(destinationContractAddress))

	//
	// Call retryMessageExecution on L1 B
	//
	log.Info("Calling retryMessageExecution on L1 B")
	receipt = utils.RetryMessageExecutionAndWaitForAcceptance(
		ctx,
		l1AInfo.BlockchainID,
		teleporter.TeleporterMessenger(l1BInfo),
		l1BInfo,
		receiveEvent.Message,
		fundedKey,
	)
	log.Info("Checking the message was successfully executed")
	messageExecutedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(l1AInfo).ParseMessageExecuted,
	)
	Expect(err).Should(BeNil())
	Expect(messageExecutedEvent.MessageID).Should(Equal(receiveEvent.MessageID))

	//
	// Verify we received the expected string
	//
	log.Info("Verifying we received the expected string")
	_, currMessage, err := L1BExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, l1AInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))
}
