package teleporter

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	poamanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/PoAManager"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	newNodeCount       = 2
	sleepPeriodSeconds = 5
)

func ValidatorChurn(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	l1AInfo, l1BInfo := network.GetTwoL1s()
	teleporterContractAddress := teleporter.TeleporterMessengerAddress(l1AInfo)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	//
	// Send a Teleporter message on L1 A
	//
	log.Info("Sending Teleporter message on source chain", "destinationBlockchainID", l1BInfo.BlockchainID)
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

	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(l1AInfo),
		l1AInfo,
		l1BInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	sendEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(l1AInfo).ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	sentTeleporterMessage := sendEvent.Message

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Construct the signed warp message
	signedWarpMessage := utils.ConstructSignedWarpMessage(
		ctx,
		receipt,
		l1AInfo,
		l1BInfo,
		nil,
		aggregator,
	)

	//
	// Modify the validator set on L1 A
	//

	// Add new nodes to the validator set
	addValidatorsCtx, cancel := context.WithTimeout(ctx, (90+sleepPeriodSeconds)*newNodeCount*time.Second)
	defer cancel()
	newNodes := network.GetExtraNodes(newNodeCount)
	validatorManagerProxy, poaManagerProxy := network.GetValidatorManager(l1AInfo.SubnetID)
	poaManager, err := poamanager.NewPoAManager(poaManagerProxy.Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())
	pChainInfo := utils.GetPChainInfo(network.GetPrimaryNetworkInfo())
	Expect(err).Should(BeNil())

	l1AInfo = network.AddSubnetValidators(newNodes, l1AInfo, true)

	for i := 0; i < newNodeCount; i++ {
		expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
		pop, err := newNodes[i].GetProofOfPossession()
		Expect(err).Should(BeNil())
		node := utils.Node{
			NodeID:  newNodes[i].NodeID,
			NodePoP: pop,
			Weight:  units.Schmeckle,
		}
		utils.InitiateAndCompletePoAValidatorRegistration(
			addValidatorsCtx,
			aggregator,
			fundedKey,
			l1AInfo,
			pChainInfo,
			poaManager,
			poaManagerProxy.Address,
			validatorManagerProxy.Address,
			expiry,
			node,
			network.GetPChainWallet(),
			network.GetNetworkID(),
		)
		// Sleep to ensure the validator manager uses a new churn tracking period
		time.Sleep(sleepPeriodSeconds * time.Second)
	}

	// Refresh the L1 info
	l1AInfo, l1BInfo = network.GetTwoL1s()

	// Trigger the proposer VM to update its height so that the inner VM can see the new validator set
	// We have to update all L1s, not just the ones directly involved in this test to ensure that the
	// proposer VM is updated on all L1s.
	for _, l1Info := range network.GetL1Infos() {
		err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(
			ctx, l1Info.EVMChainID, fundedKey, l1Info.WSClient,
		)
		Expect(err).Should(BeNil())
	}

	//
	// Attempt to deliver the warp message signed by the old validator set. This should fail.
	//
	// Construct the transaction to send the Warp message to the destination chain
	signedTx := utils.CreateReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessage,
		sendEvent.Message.RequiredGasLimit,
		teleporterContractAddress,
		fundedKey,
		l1BInfo,
	)

	log.Info("Sending transaction to destination chain")
	utils.SendTransactionAndWaitForFailure(ctx, l1BInfo, signedTx)

	// Verify the message was not delivered
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())

	//
	// Retry sending the message, and attempt to relay again. This should succeed.
	//
	log.Info("Retrying message sending on source chain")
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := teleporter.TeleporterMessenger(l1AInfo).RetrySendCrossChainMessage(
		optsA, sentTeleporterMessage,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt = utils.WaitForTransactionSuccess(ctx, l1AInfo, tx.Hash())

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

	// Verify the message was delivered
	delivered, err = teleporter.TeleporterMessenger(l1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// The test cases now do not require any specific nodes to be validators, so leave the validator set as is.
	// If this changes in the future, this test will need to perform cleanup by removing the nodes that were added
	// and re-adding the nodes that were removed.
}
