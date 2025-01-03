package teleporter

import (
	"context"
	"math/big"

	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func AddFeeAmount(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	teleporterContractAddress := teleporter.TeleporterMessengerAddress(l1AInfo)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Use mock token as the fee token
	mockTokenAddress, mockToken := utils.DeployExampleERC20(
		context.Background(),
		fundedKey,
		l1AInfo,
	)
	utils.ERC20Approve(
		ctx,
		mockToken,
		teleporterContractAddress,
		big.NewInt(1e18),
		l1AInfo,
		fundedKey,
	)

	initFeeAmount := big.NewInt(1)

	// Send a transaction to L1 A that sends a message to L1 B.
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: l1BInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          initFeeAmount,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	sendCrossChainMsgReceipt, messageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(l1AInfo),
		l1AInfo,
		l1BInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	// Add a fee amount to the message.
	additionalFeeAmount := big.NewInt(2)
	utils.SendAddFeeAmountAndWaitForAcceptance(
		ctx,
		l1AInfo,
		l1BInfo,
		messageID,
		additionalFeeAmount,
		mockTokenAddress,
		fundedKey,
		teleporter.TeleporterMessenger(l1AInfo),
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()
	// Relay message from L1 A to L1 B
	deliveryReceipt := teleporter.RelayTeleporterMessage(
		ctx,
		sendCrossChainMsgReceipt,
		l1AInfo,
		l1BInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)
	receiveEvent, err := utils.GetEventFromLogs(
		deliveryReceipt.Logs,
		teleporter.TeleporterMessenger(l1BInfo).ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	// Check Teleporter message received on the destination (L1 B)
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the initial relayer reward amount on L1 A.
	initialRewardAmount, err := teleporter.TeleporterMessenger(l1AInfo).CheckRelayerRewardAmount(
		&bind.CallOpts{},
		receiveEvent.RewardRedeemer,
		mockTokenAddress)
	Expect(err).Should(BeNil())

	// Send a message from L1 B back to L1 A that includes the specific receipt for the message.
	sendSpecificReceiptsReceipt, sendSpecificReceiptsMessageID := utils.SendSpecifiedReceiptsAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(l1BInfo),
		l1BInfo,
		l1AInfo.BlockchainID,
		[][32]byte{receiveEvent.MessageID},
		teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          big.NewInt(0),
		},
		[]common.Address{},
		fundedKey,
	)

	// Relay message containing the specific receipt from L1 B to L1 A
	teleporter.RelayTeleporterMessage(
		ctx,
		sendSpecificReceiptsReceipt,
		l1BInfo,
		l1AInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	// Check message delivered
	delivered, err = teleporter.TeleporterMessenger(l1AInfo).MessageReceived(
		&bind.CallOpts{},
		sendSpecificReceiptsMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the updated relayer reward amount
	expectedIncrease := new(big.Int).Add(initFeeAmount, additionalFeeAmount)
	newRewardAmount, err := teleporter.TeleporterMessenger(l1AInfo).CheckRelayerRewardAmount(
		&bind.CallOpts{},
		receiveEvent.RewardRedeemer,
		mockTokenAddress)
	Expect(err).Should(BeNil())
	Expect(newRewardAmount).Should(Equal(new(big.Int).Add(initialRewardAmount, expectedIncrease)))

	// If the funded address is the one able to redeem the rewards, do so and check the reward amount is reset.
	if fundedAddress == receiveEvent.RewardRedeemer {
		utils.RedeemRelayerRewardsAndConfirm(
			ctx,
			teleporter.TeleporterMessenger(l1AInfo),
			l1AInfo,
			mockToken,
			mockTokenAddress,
			fundedKey,
			newRewardAmount,
		)
	}
}
