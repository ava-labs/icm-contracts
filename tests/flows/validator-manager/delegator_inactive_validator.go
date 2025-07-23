package staking

import (
	"context"
	"log"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	exampleerc20 "github.com/ava-labs/icm-contracts/abi-bindings/go/mocks/ExampleERC20"
	erc20tokenstakingmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/ERC20TokenStakingManager"
	istakingmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/interfaces/IStakingManager"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

/*
 * Tests that a delegator can recover its funds from an inactive validator.
 * The steps are as follows:
 * - Deploy an ERC20TokenStakingManager
 * - Initiate and complete validator registration
 * - Initiate and complete delegator registration
 * - Disable the validator by issuing a DisableL1ValidatorTx on the P-Chain
 * - Initiate and complete validator removal
 */
func RemoveDelegatorInactiveValidator(network *localnetwork.LocalAvalancheNetwork) {
	// Get the L1s info
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	ctx := context.Background()

	nodes, initialValidationIDs := network.ConvertSubnet(
		ctx,
		l1AInfo,
		utils.ERC20TokenStakingManager,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		fundedKey,
		false,
	)
	validatorManagerProxy, stakingManagerProxy := network.GetValidatorManager(l1AInfo.SubnetID)
	erc20StakingManager, err := erc20tokenstakingmanager.NewERC20TokenStakingManager(
		stakingManagerProxy.Address,
		l1AInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	erc20Address, err := erc20StakingManager.Erc20(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	erc20, err := exampleerc20.NewExampleERC20(erc20Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			l1AInfo.SubnetID,
		},
	)
	defer signatureAggregator.Shutdown()

	//
	// Delist one initial validator
	//
	posStakingManager, err := istakingmanager.NewIStakingManager(stakingManagerProxy.Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())
	utils.InitiateAndCompleteEndInitialPoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		posStakingManager,
		stakingManagerProxy.Address,
		validatorManagerProxy.Address,
		initialValidationIDs[0],
		0,
		nodes[0].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	//
	// Register the validator as PoS
	//
	registrationInitiatedEvent := utils.InitiateAndCompleteERC20ValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		erc20StakingManager,
		stakingManagerProxy.Address,
		validatorManagerProxy.Address,
		erc20,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
	validationID := ids.ID(registrationInitiatedEvent.ValidationID)

	//
	// Register a delegator
	//
	var delegationID ids.ID
	delegatorStake, err := erc20StakingManager.WeightToValue(
		&bind.CallOpts{},
		nodes[0].Weight,
	)
	Expect(err).Should(BeNil())
	delegatorStake.Div(delegatorStake, big.NewInt(10))
	delegatorWeight, err := erc20StakingManager.ValueToWeight(
		&bind.CallOpts{},
		delegatorStake,
	)
	Expect(err).Should(BeNil())

	// Get the delegator's staking token balance
	delegatorBalance, err := erc20.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	{
		log.Println("Registering delegator")
		newValidatorWeight := nodes[0].Weight + delegatorWeight

		nonce := uint64(1)

		receipt := utils.InitiateERC20DelegatorRegistration(
			ctx,
			fundedKey,
			l1AInfo,
			validationID,
			delegatorStake,
			erc20,
			stakingManagerProxy.Address,
			erc20StakingManager,
		)
		initRegistrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			erc20StakingManager.ParseInitiatedDelegatorRegistration,
		)
		Expect(err).Should(BeNil())
		delegationID = initRegistrationEvent.DelegationID

		// Gather subnet-evm Warp signatures for the L1ValidatorWeightMessage & relay to the P-Chain
		signedWarpMessage := utils.ConstructSignedWarpMessage(
			context.Background(),
			receipt,
			l1AInfo,
			pChainInfo,
			nil,
			signatureAggregator,
		)

		// Issue a tx to update the validator's weight on the P-Chain
		_, err = network.GetPChainWallet().IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
		Expect(err).Should(BeNil())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, l1AInfo, fundedKey, 5)

		// Construct an L1ValidatorWeightMessage Warp message from the P-Chain
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
			stakingManagerProxy.Address,
			registrationSignedMessage,
		)
		// Check that the validator is registered in the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			erc20StakingManager.ParseCompletedDelegatorRegistration,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	}

	//
	// Disable the validator on the P-Chain
	//
	log.Println("Disabling the validator on the P-Chain")
	_, err = network.GetPChainWallet(validationID).IssueDisableL1ValidatorTx(
		validationID,
	)
	Expect(err).Should(BeNil())
	utils.PChainProposerVMWorkaround(network.GetPChainWallet())
	utils.AdvanceProposerVM(ctx, l1AInfo, fundedKey, 5)

	//
	// Delist the delegator
	//
	{
		log.Println("Delisting delegator")
		nonce := uint64(2)
		receipt := utils.InitiateDelegatorRemoval(
			ctx,
			fundedKey,
			l1AInfo,
			stakingManagerProxy.Address,
			delegationID,
		)
		delegatorRemovalEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			erc20StakingManager.ParseInitiatedDelegatorRemoval,
		)
		Expect(err).Should(BeNil())
		Expect(delegatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(delegatorRemovalEvent.DelegationID[:]).Should(Equal(delegationID[:]))

		// Gather subnet-evm Warp signatures for the SetL1ValidatorWeightMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := utils.ConstructSignedWarpMessage(
			context.Background(),
			receipt,
			l1AInfo,
			pChainInfo,
			nil,
			signatureAggregator,
		)
		Expect(err).Should(BeNil())

		// Issue a tx to update the validator's weight on the P-Chain
		_, err = network.GetPChainWallet().IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
		Expect(err).Should(BeNil())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, l1AInfo, fundedKey, 5)

		// Construct an L1ValidatorWeightMessage Warp message from the P-Chain
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
		receipt = utils.CompleteDelegatorRemoval(
			ctx,
			fundedKey,
			delegationID,
			l1AInfo,
			stakingManagerProxy.Address,
			signedMessage,
		)

		// Check that the delegator has been delisted from the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			erc20StakingManager.ParseCompletedDelegatorRemoval,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))

		// Check that the delegator's stake was returned within an error
		// margin of the weight to value factor
		delegatorBalanceAfter, err := erc20.BalanceOf(&bind.CallOpts{}, fundedAddress)
		Expect(err).Should(BeNil())
		diff := new(big.Int).Sub(delegatorBalance, delegatorBalanceAfter)
		diff.Abs(diff)
		Expect(diff.Cmp(big.NewInt(int64(utils.DefaultWeightToValueFactor)))).Should(Equal(-1))
	}
}
