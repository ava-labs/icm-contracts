package teleporter

import (
	"context"
	"math/big"

	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

// Tests basic one-way send from L1 A to L1 B and vice versa
func BasicSendReceive(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	teleporterContractAddress := teleporter.TeleporterMessengerAddress(l1AInfo)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	// Send a transaction to L1 A to issue an ICM Message from the Teleporter contract to L1 B
	ctx := context.Background()

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Clear the receipt queue from L1 B -> L1 A to have a clean slate for the test flow.
	// This is only done if the test non-external networks because external networks may have
	// an arbitrarily high number of receipts to be cleared from a given queue from unrelated messages.
	teleporter.ClearReceiptQueue(ctx, fundedKey, l1BInfo, l1AInfo, aggregator)

	feeAmount := big.NewInt(1)
	feeTokenAddress, feeToken := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		l1AInfo,
	)
	utils.ERC20Approve(
		ctx,
		feeToken,
		teleporterContractAddress,
		big.NewInt(0).Mul(big.NewInt(1e18),
			big.NewInt(10)),
		l1AInfo,
		fundedKey,
	)

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: l1BInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: feeTokenAddress,
			Amount:          feeAmount,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(l1AInfo),
		l1AInfo,
		l1BInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)
	expectedReceiptID := teleporterMessageID

	// Relay the message to the destination
	deliveryReceipt := teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		l1AInfo,
		l1BInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)
	receiveEvent, err := utils.GetEventFromLogs(
		deliveryReceipt.Logs,
		teleporter.TeleporterMessenger(l1BInfo).ParseReceiveCrossChainMessage,
	)
	Expect(err).Should(BeNil())

	// Check Teleporter message received on the destination
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Send a transaction to L1 B to issue an ICM Message from the Teleporter contract to L1 A
	sendCrossChainMessageInput.DestinationBlockchainID = l1AInfo.BlockchainID
	sendCrossChainMessageInput.FeeInfo.Amount = big.NewInt(0)
	receipt, teleporterMessageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(l1BInfo),
		l1BInfo,
		l1AInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	// Relay the message to the destination
	deliveryReceipt = teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		l1BInfo,
		l1AInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	Expect(utils.CheckReceiptReceived(
		deliveryReceipt,
		expectedReceiptID,
		teleporter.TeleporterMessenger(l1AInfo))).Should(BeTrue())

	// Check Teleporter message received on the destination
	delivered, err = teleporter.TeleporterMessenger(l1AInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// If the reward address of the message from A->B is the funded address, which is able to send
	// transactions on L1 A, then redeem the rewards.
	if receiveEvent.RewardRedeemer == fundedAddress {
		utils.RedeemRelayerRewardsAndConfirm(
			ctx, teleporter.TeleporterMessenger(l1AInfo), l1AInfo, feeToken, feeTokenAddress, fundedKey, feeAmount,
		)
	}
}
