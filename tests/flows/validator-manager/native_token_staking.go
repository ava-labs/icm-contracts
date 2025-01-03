package staking

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	nativetokenstakingmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	iposvalidatormanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/interfaces/IPoSValidatorManager"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

/*
 * Registers a native token staking validator on a L1. The steps are as follows:
 * - Deploy the NativeTokenStakingManager
 * - Initiate validator registration
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is registered in the staking contract
 *
 * Delists the validator from the L1. The steps are as follows:
 * - Initiate validator delisting
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is delisted from the staking contract
 */
func NativeTokenStakingManager(network *localnetwork.LocalNetwork) {
	// Get the L1s info
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			l1AInfo.SubnetID,
		},
	)
	ctx := context.Background()

	nodes, initialValidationIDs, _ := network.ConvertSubnet(
		ctx,
		l1AInfo,
		utils.NativeTokenStakingManager,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		fundedKey,
		false,
	)
	stakingManagerAddress := network.GetValidatorManager(l1AInfo.SubnetID)
	nativeStakingManager, err := nativetokenstakingmanager.NewNativeTokenStakingManager(
		stakingManagerAddress,
		l1AInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	utils.AddNativeMinterAdmin(ctx, l1AInfo, fundedKey, stakingManagerAddress)

	//
	// Delist one initial validator
	//
	posStakingManager, err := iposvalidatormanager.NewIPoSValidatorManager(stakingManagerAddress, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())
	utils.InitializeAndCompleteEndInitialPoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		posStakingManager,
		stakingManagerAddress,
		initialValidationIDs[0],
		0,
		nodes[0].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	//
	// Register the validator as PoS
	//
	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	validationID := utils.InitializeAndCompleteNativeValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		nativeStakingManager,
		stakingManagerAddress,
		expiry,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
	validatorStartTime := time.Now()

	//
	// Register a delegator
	//
	var delegationID ids.ID
	{
		log.Println("Registering delegator")
		delegatorStake, err := nativeStakingManager.WeightToValue(
			&bind.CallOpts{},
			nodes[0].Weight,
		)
		Expect(err).Should(BeNil())
		delegatorStake.Div(delegatorStake, big.NewInt(10))
		delegatorWeight, err := nativeStakingManager.ValueToWeight(
			&bind.CallOpts{},
			delegatorStake,
		)
		Expect(err).Should(BeNil())
		newValidatorWeight := nodes[0].Weight + delegatorWeight

		nonce := uint64(1)

		receipt := utils.InitializeNativeDelegatorRegistration(
			ctx,
			fundedKey,
			l1AInfo,
			validationID,
			delegatorStake,
			stakingManagerAddress,
			nativeStakingManager,
		)
		initRegistrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			nativeStakingManager.ParseDelegatorAdded,
		)
		Expect(err).Should(BeNil())
		delegationID = initRegistrationEvent.DelegationID

		aggregator := network.GetSignatureAggregator()
		defer aggregator.Shutdown()

		// Gather subnet-evm Warp signatures for the L1ValidatorWeightMessage & relay to the P-Chain
		signedWarpMessage := utils.ConstructSignedWarpMessage(
			context.Background(),
			receipt,
			l1AInfo,
			pChainInfo,
			nil,
			aggregator,
		)

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, l1AInfo, fundedKey, 5)

		// Construct a L1ValidatorWeightMessage Warp message from the P-Chain
		registrationSignedMessage := utils.ConstructL1ValidatorWeightMessage(
			validationID,
			nonce,
			newValidatorWeight,
			l1AInfo,
			pChainInfo,
			signatureAggregator,
			network.GetNetworkID(),
		)

		// Deliver the Warp message to the L1
		receipt = utils.CompleteDelegatorRegistration(
			ctx,
			fundedKey,
			delegationID,
			l1AInfo,
			stakingManagerAddress,
			registrationSignedMessage,
		)
		// Check that the validator is registered in the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			nativeStakingManager.ParseDelegatorRegistered,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	}
	//
	// Delist the delegator
	//
	{
		log.Println("Delisting delegator")
		nonce := uint64(2)
		receipt := utils.InitializeEndDelegation(
			ctx,
			fundedKey,
			l1AInfo,
			stakingManagerAddress,
			delegationID,
		)
		delegatorRemovalEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			nativeStakingManager.ParseDelegatorRemovalInitialized,
		)
		Expect(err).Should(BeNil())
		Expect(delegatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(delegatorRemovalEvent.DelegationID[:]).Should(Equal(delegationID[:]))

		aggregator := network.GetSignatureAggregator()
		defer aggregator.Shutdown()

		// Gather subnet-evm Warp signatures for the SetL1ValidatorWeightMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := utils.ConstructSignedWarpMessage(
			context.Background(),
			receipt,
			l1AInfo,
			pChainInfo,
			nil,
			aggregator,
		)
		Expect(err).Should(BeNil())

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, l1AInfo, fundedKey, 5)

		// Construct a L1ValidatorWeightMessage Warp message from the P-Chain
		signedMessage := utils.ConstructL1ValidatorWeightMessage(
			validationID,
			nonce,
			nodes[0].Weight,
			l1AInfo,
			pChainInfo,
			signatureAggregator,
			network.GetNetworkID(),
		)

		// Deliver the Warp message to the L1
		receipt = utils.CompleteEndDelegation(
			ctx,
			fundedKey,
			delegationID,
			l1AInfo,
			stakingManagerAddress,
			signedMessage,
		)

		// Check that the delegator has been delisted from the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			nativeStakingManager.ParseDelegationEnded,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	}

	//
	// Delist the validator
	//
	utils.InitializeAndCompleteEndPoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		posStakingManager,
		stakingManagerAddress,
		validationID,
		expiry,
		nodes[0],
		1,
		true,
		validatorStartTime,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
}
