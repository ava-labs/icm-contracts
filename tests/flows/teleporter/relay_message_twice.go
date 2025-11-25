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

func RelayMessageTwice(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Send a transaction to L1 A to issue an ICM Message from the Teleporter contract to L1 B
	//
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

	log.Info(
		"Sending Teleporter transaction on source chain",
		"destinationBlockchainID", l1BInfo.BlockchainID,
	)
	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(l1AInfo),
		l1AInfo,
		l1BInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	//
	// Relay the message to the destination
	//
	teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		l1AInfo,
		l1BInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	//
	// Check Teleporter message received on the destination
	//
	log.Info("Checking the message was received on the destination")
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Attempt to send the same message again, should fail
	//
	log.Info("Relaying the same Teleporter message again on the destination")
	teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		l1AInfo,
		l1BInfo,
		false,
		fundedKey,
		nil,
		aggregator,
	)
}
