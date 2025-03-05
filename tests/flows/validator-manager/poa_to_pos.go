package staking

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	ownableupgradeable "github.com/ava-labs/icm-contracts/abi-bindings/go/OwnableUpgradeable"
	nativetokenstakingmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	poavalidatormanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/PoAValidatorManagerV1"
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

	// Deploy PoAValidatorManager
	nodes, initialValidationIDs := network.ConvertSubnetPoAV1(
		ctx,
		l1AInfo,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		ownerKey,
	)
	validatorManagerProxy, proxyAdmin := network.GetValidatorManager(l1AInfo.SubnetID)
	poaValidatorManager, err := poavalidatormanager.NewPoAValidatorManager(validatorManagerProxy.Address, l1AInfo.RPCClient)
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
	utils.InitializeAndCompleteEndInitialPoAValidationV1(
		ctx,
		signatureAggregator,
		ownerKey,
		fundedKey,
		l1AInfo,
		pChainInfo,
		poaValidatorManager,
		validatorManagerProxy.Address,
		initialValidationIDs[0],
		0,
		nodes[0].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	// Try to call with invalid owner
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())

	_, err = poaValidatorManager.InitializeValidatorRegistration(
		opts,
		poavalidatormanager.ValidatorRegistrationInput{
			NodeID:                nodes[0].NodeID[:],
			BlsPublicKey:          nodes[0].NodePoP.PublicKey[:],
			RegistrationExpiry:    uint64(time.Now().Add(24 * time.Hour).Unix()),
			RemainingBalanceOwner: poavalidatormanager.PChainOwner{},
			DisableOwner:          poavalidatormanager.PChainOwner{},
		},
		nodes[0].Weight,
	)
	Expect(err).ShouldNot(BeNil())

	//
	// Re-register the validator as a SoV validator
	//
	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	poaValidationID := utils.InitializeAndCompletePoAValidatorRegistrationV1(
		ctx,
		signatureAggregator,
		ownerKey,
		fundedKey,
		l1AInfo,
		pChainInfo,
		poaValidatorManager,
		validatorManagerProxy.Address,
		expiry,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	poaValidator, err := poaValidatorManager.GetValidator(&bind.CallOpts{}, poaValidationID)
	Expect(err).Should(BeNil())
	poaNodeID := poaValidator.NodeID

	/*
	 ******************
	 * Migrate PoAValidatorManager to StakingManager
	 ******************
	 */
	// Deploy ValidatorManager contract
	validatorManagerAddress, _ := utils.DeployValidatorManager(ctx, ownerKey, l1AInfo, false)
	opts, err = bind.NewKeyedTransactorWithChainID(ownerKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := proxyAdmin.UpgradeAndCall(opts, validatorManagerProxy.Address, validatorManagerAddress, []byte{})
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1AInfo, tx.Hash())

	// Deploy StakingManager contract
	stakingManagerAddress, _ := utils.DeployAndInitializeValidatorManagerSpecialization(
		ctx,
		ownerKey,
		l1AInfo,
		validatorManagerProxy.Address,
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
	ownable, err := ownableupgradeable.NewOwnableUpgradeable(validatorManagerProxy.Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())

	opts, err = bind.NewKeyedTransactorWithChainID(ownerKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = ownable.TransferOwnership(opts, stakingManagerAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(context.Background(), l1AInfo, tx.Hash())

	validatorManager, err := validatormanager.NewValidatorManager(validatorManagerProxy.Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())

	// Check that previous validator is still registered
	validationID, err := validatorManager.RegisteredValidators(&bind.CallOpts{}, poaNodeID)
	Expect(err).Should(BeNil())
	Expect(validationID[:]).Should(Equal(poaValidationID[:]))

	// Check the subnetID was migrated
	subnetID, err := validatorManager.SubnetID(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(subnetID[:]).Should(Equal(l1AInfo.SubnetID[:]))

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
		validatorManagerProxy.Address,
		poaValidationID,
		expiry,
		nodes[0],
		1,
		false,
		time.Time{},
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	expiry2 := uint64(time.Now().Add(24 * time.Hour).Unix())
	posValidationID := utils.InitiateAndCompleteNativeValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		nativeStakingManager,
		stakingManagerAddress,
		validatorManagerProxy.Address,
		expiry2,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
	validatorStartTime := time.Now()

	// Delist the PoS validator
	utils.InitiateAndCompleteEndPoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		posStakingManager,
		stakingManagerAddress,
		validatorManagerProxy.Address,
		posValidationID,
		expiry2,
		nodes[0],
		1,
		true,
		validatorStartTime,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
}
