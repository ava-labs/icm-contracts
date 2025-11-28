package registry

import (
	"context"

	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

func PauseTeleporter(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy TestMessenger to L1s A and B
	//
	ctx := context.Background()
	teleporterAddress := teleporter.TeleporterMessengerAddress(l1AInfo)
	_, testMessengerA := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(l1AInfo),
		l1AInfo,
	)
	testMessengerAddressB, testMessengerB := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(l1BInfo),
		l1BInfo,
	)

	// Pause Teleporter on L1 B
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey,
		l1BInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())
	tx, err := testMessengerB.PauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, l1BInfo, tx.Hash())
	pauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessengerB.ParseTeleporterAddressPaused)
	Expect(err).Should(BeNil())
	Expect(pauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err := testMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Send a message from L1 A to L1 B, which should fail
	teleporter.SendExampleCrossChainMessageAndVerify(
		ctx,
		l1AInfo,
		testMessengerA,
		l1BInfo,
		testMessengerAddressB,
		testMessengerB,
		fundedKey,
		"message_1",
		aggregator,
		false,
	)

	// Unpause Teleporter on L1 B
	tx, err = testMessengerB.UnpauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt = utils.WaitForTransactionSuccess(ctx, l1BInfo, tx.Hash())
	unpauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessengerB.ParseTeleporterAddressUnpaused)
	Expect(err).Should(BeNil())
	Expect(unpauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err = testMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())

	// Send a message from L1 A to L1 B again, which should now succeed
	teleporter.SendExampleCrossChainMessageAndVerify(
		ctx,
		l1AInfo,
		testMessengerA,
		l1BInfo,
		testMessengerAddressB,
		testMessengerB,
		fundedKey,
		"message_2",
		aggregator,
		true,
	)
}
