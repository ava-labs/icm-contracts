package teleporter

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	gasUtils "github.com/ava-labs/icm-contracts/utils/gas-utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/log"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	. "github.com/onsi/gomega"
)

// Disallow this test from being run on anything but a local network, since it requires special behavior by the relayer
func RelayerModifiesMessage(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	// Send a transaction to L1 A to issue a Warp Message from the Teleporter contract to L1 B
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

	// Relay the message to the destination
	// Relayer modifies the message in flight
	relayAlteredMessage(
		ctx,
		teleporter,
		receipt,
		l1AInfo,
		l1BInfo,
		network,
	)

	// Check Teleporter message was not received on the destination
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}

func relayAlteredMessage(
	ctx context.Context,
	teleporter utils.TeleporterTestInfo,
	sourceReceipt *types.Receipt,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	network *localnetwork.LocalAvalancheNetwork,
) {
	// Fetch the Teleporter message from the logs
	sendEvent, err := utils.GetEventFromLogs(
		sourceReceipt.Logs,
		teleporter.TeleporterMessenger(source).ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	signedWarpMessage := utils.ConstructSignedWarpMessage(
		ctx,
		sourceReceipt,
		source,
		destination,
		nil,
		aggregator,
	)

	// Construct the transaction to send the Warp message to the destination chain
	_, fundedKey := network.GetFundedAccountInfo()
	signedTx := createAlteredReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessage,
		&sendEvent.Message,
		sendEvent.Message.RequiredGasLimit,
		teleporter.TeleporterMessengerAddress(source),
		fundedKey,
		destination,
	)

	log.Info("Sending transaction to destination chain")
	utils.SendTransactionAndWaitForFailure(ctx, destination, signedTx)
}

func createAlteredReceiveCrossChainMessageTransaction(
	ctx context.Context,
	signedMessage *avalancheWarp.Message,
	teleporterMessage *teleportermessenger.TeleporterMessage,
	requiredGasLimit *big.Int,
	teleporterContractAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
) *types.Transaction {
	fundedAddress := crypto.PubkeyToAddress(fundedKey.PublicKey)
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing transaction for the destination chain")

	numSigners, err := signedMessage.Signature.NumSigners()
	Expect(err).Should(BeNil())

	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(
		numSigners,
		requiredGasLimit,
		len(signedMessage.Bytes()),
		len(signedMessage.Payload),
		len(teleporterMessage.Receipts),
	)
	Expect(err).Should(BeNil())

	callData, err := teleportermessenger.PackReceiveCrossChainMessage(0, fundedAddress)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := utils.CalculateTxParams(ctx, l1Info, fundedAddress)

	alterTeleporterMessage(signedMessage)

	destinationTx := predicateutils.NewPredicateTx(
		l1Info.EVMChainID,
		nonce,
		&teleporterContractAddress,
		gasLimit,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessage.Bytes(),
	)

	return utils.SignTransaction(destinationTx, fundedKey, l1Info.EVMChainID)
}

func alterTeleporterMessage(signedMessage *avalancheWarp.Message) {
	warpMsgPayload, err := warpPayload.ParseAddressedCall(signedMessage.UnsignedMessage.Payload)
	Expect(err).Should(BeNil())

	teleporterMessage := teleportermessenger.TeleporterMessage{}
	err = teleporterMessage.Unpack(warpMsgPayload.Payload)
	Expect(err).Should(BeNil())
	// Alter the message
	teleporterMessage.Message[0] = ^teleporterMessage.Message[0]

	// Pack the teleporter message
	teleporterMessageBytes, err := teleporterMessage.Pack()
	Expect(err).Should(BeNil())

	payload, err := warpPayload.NewAddressedCall(warpMsgPayload.SourceAddress, teleporterMessageBytes)
	Expect(err).Should(BeNil())

	signedMessage.UnsignedMessage.Payload = payload.Bytes()

	signedMessage.Initialize()
}
