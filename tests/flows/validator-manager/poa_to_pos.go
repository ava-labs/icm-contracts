package staking

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	nativetokenstakingmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	poamanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/PoAManager"
	validatormanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/ValidatorManager"
	istakingmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/interfaces/IStakingManager"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

/*
 * Register a PoA validator manager on a L1 with a proxy. The steps are as follows:
 * - Generate random address to be the owner address
 * - Fund native assets to the owner address
 * - Deploy the PoAValidatorManager contract
 * - Deploy a TransparentUpgradeableProxy contract that points to the PoAValidatorManager
 * - Call initialize on the PoAValidatorManager through the proxy
 * - Initialize and complete PoA validator registration
 *
 * Migrates the proxy to a PoS validator manager. The steps are as follows:
 * - Deploy the StakingManager contract
 * - Upgrade the TransparentUpgradeableProxy to point to the StakingManager
 * - Call initialize on the StakingManager through the proxy
 * - Check that previous validator is still active
 * - Initialize and complete PoS validator registration
 * - Attempt to delist previous PoA validator with wrong owner and check that it fails
 * - Delist the previous PoA validator properly
 * - Delist the PoS validator
 */
func PoAMigrationToPoS(network *localnetwork.LocalNetwork) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	// Generate random address to be the owner address
	ownerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	ownerAddress := crypto.PubkeyToAddress(ownerKey.PublicKey)

	// Transfer native assets to the owner account
	ctx := context.Background()
	fundAmount := big.NewInt(1e18) // 10avax
	fundAmount.Mul(fundAmount, big.NewInt(10))
	utils.SendNativeTransfer(
		ctx,
		l1AInfo,
		fundedKey,
		ownerAddress,
		fundAmount,
	)

	// Deploy PoAManager
	nodes, initialValidationIDs := network.ConvertSubnet(
		ctx,
		l1AInfo,
		utils.PoAValidatorManager,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		ownerKey,
		false,
	)
	validatorManagerAddr, poaManagerAddr := network.GetValidatorManager(l1AInfo.SubnetID)
	poaManager, err := poamanager.NewPoAManager(poaManagerAddr.Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())
	validatorManager, err := validatormanager.NewValidatorManager(validatorManagerAddr.Address, l1AInfo.RPCClient)
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
	utils.InitiateAndCompleteEndInitialPoAValidation(
		ctx,
		signatureAggregator,
		ownerKey,
		l1AInfo,
		pChainInfo,
		poaManager,
		poaManagerAddr.Address,
		validatorManagerAddr.Address,
		initialValidationIDs[0],
		0,
		nodes[0].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	// Try to call with invalid owner
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())

	_, err = poaManager.InitiateValidatorRegistration(
		opts,
		nodes[0].NodeID[:],
		nodes[0].NodePoP.PublicKey[:],
		poamanager.PChainOwner{},
		poamanager.PChainOwner{},
		nodes[0].Weight,
	)
	Expect(err).ShouldNot(BeNil())

	//
	// Re-register the validator as a SoV validator
	//
	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	registrationInitiatedEvent := utils.InitiateAndCompletePoAValidatorRegistration(
		ctx,
		signatureAggregator,
		ownerKey,
		l1AInfo,
		pChainInfo,
		poaManager,
		poaManagerAddr.Address,
		validatorManagerAddr.Address,
		expiry,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
	poaValidationID := registrationInitiatedEvent.ValidationID

	poaValidator, err := validatorManager.GetValidator(&bind.CallOpts{}, poaValidationID)
	Expect(err).Should(BeNil())
	poaNodeID := poaValidator.NodeID

	/*
	 ******************
	 * Migrate PoAValidatorManager to StakingManager
	 ******************
	 */

	// Deploy StakingManager contract
	stakingManagerAddress, _ := utils.DeployAndInitializeValidatorManagerSpecialization(
		ctx,
		ownerKey,
		l1AInfo,
		validatorManagerAddr.Address,
		utils.NativeTokenStakingManager,
		false,
	)

	utils.AddNativeMinterAdmin(ctx, l1AInfo, fundedKey, stakingManagerAddress)

	nativeStakingManager, err := nativetokenstakingmanager.NewNativeTokenStakingManager(
		stakingManagerAddress,
		l1AInfo.RPCClient,
	)
	Expect(err).Should(BeNil())

	// Transfer ownership from PoA -> the new staking manager
	opts, err = bind.NewKeyedTransactorWithChainID(ownerKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := poaManager.TransferValidatorManagerOwnership(opts, stakingManagerAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(context.Background(), l1AInfo, tx.Hash())

	// Check that previous validator is still registered
	validationID, err := validatorManager.GetNodeValidationID(&bind.CallOpts{}, poaNodeID)
	Expect(err).Should(BeNil())
	Expect(validationID[:]).Should(Equal(poaValidationID[:]))

	validator, err := validatorManager.GetValidator(&bind.CallOpts{}, validationID)
	Expect(err).Should(BeNil())
	Expect(validator.EndTime).Should(Equal(uint64(0)))

	//
	// Remove the PoA validator and re-register as a PoS validator
	//
	posStakingManager, err := istakingmanager.NewIStakingManager(stakingManagerAddress, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())
	utils.InitiateAndCompleteEndPoSValidation(
		ctx,
		signatureAggregator,
		ownerKey,
		l1AInfo,
		pChainInfo,
		posStakingManager,
		stakingManagerAddress,
		validatorManagerAddr.Address,
		poaValidationID,
		registrationInitiatedEvent.RegistrationExpiry,
		nodes[0],
		1,
		false,
		time.Time{},
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	posRegistrationInitiatedEvent := utils.InitiateAndCompleteNativeValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		nativeStakingManager,
		stakingManagerAddress,
		validatorManagerAddr.Address,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
	validatorStartTime := time.Now()
	posValidationID := posRegistrationInitiatedEvent.ValidationID

	// Delist the PoS validator
	utils.InitiateAndCompleteEndPoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		posStakingManager,
		stakingManagerAddress,
		validatorManagerAddr.Address,
		posValidationID,
		posRegistrationInitiatedEvent.RegistrationExpiry,
		nodes[0],
		1,
		true,
		validatorStartTime,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
}
