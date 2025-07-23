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

func DeliverToWrongChain(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, L1CInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Get the expected teleporter message ID for L1 C
	//
	expectedAtoCMessageID, err := teleporter.TeleporterMessenger(l1AInfo).GetNextMessageID(
		&bind.CallOpts{},
		L1CInfo.BlockchainID,
	)
	Expect(err).Should(BeNil())

	//
	// Submit a message to be sent from L1A to L1B
	//
	ctx := context.Background()
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: l1BInfo.BlockchainID, // Message intended for L1B
		DestinationAddress:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	log.Info(
		"Sending Teleporter transaction on source chain",
		"destinationBlockchainID", l1BInfo.BlockchainID,
	)

	receipt, _ := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(l1AInfo),
		l1AInfo,
		l1BInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		l1AInfo,
		L1CInfo,
		false,
		fundedKey,
		nil,
		aggregator,
	)

	//
	// Check that the message was not received on the L1 C
	//
	delivered, err := teleporter.TeleporterMessenger(L1CInfo).MessageReceived(
		&bind.CallOpts{}, expectedAtoCMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}
