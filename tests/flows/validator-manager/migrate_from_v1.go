package staking

import (
	"context"
	"crypto/ecdsa"
	"log"
	goLog "log"
	"math/big"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	warpMessage "github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	ownableupgradeable "github.com/ava-labs/icm-contracts/abi-bindings/go/OwnableUpgradeable"
	proxyadmin "github.com/ava-labs/icm-contracts/abi-bindings/go/ProxyAdmin"
	nativetokenstakingmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	poavalidatormanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/PoAValidatorManagerV1"
	validatormanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/ValidatorManager"
	istakingmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/interfaces/IStakingManager"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	"github.com/ava-labs/icm-contracts/tests/network"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

/*
 * Register a PoA validator manager on a L1 with a proxy. The steps are as follows:
 * - Generate random address to be the owner address
 * - Fund native assets to the owner address
 * - Deploy the V1 PoAValidatorManager contract
 * - Deploy a TransparentUpgradeableProxy contract that points to the PoAValidatorManager
 * - Call initialize on the V1 PoAValidatorManager through the proxy
 * - Initialize and complete PoA validator registration
 *
 * Migrates the V1 PoAValidatorManager to a V2 ValidatorManager. The steps are as follows:
 * - Deploy the ValidatorManager contract
 * - Upgrade the TransparentUpgradeableProxy to point to the ValidatorManager
 * - Call initialize on the ValidatorManager through the proxy
 * - Migrate the registered validator to the V2 contract
 * - Check that previous validator is still active
 * - Attempt to delist previous PoA validator with wrong owner and check that it fails
 * - Delist the previous PoA validator properly
 * - Initialize and complete a new validator registration
 * - Delist the new validator
 */
func MigrateFromV1PoA(network *localnetwork.LocalNetwork) {
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
	validatorManagerProxy, nodes, initialValidationIDs := convertSubnetPoAV1(
		network,
		ctx,
		l1AInfo,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		ownerKey,
	)

	poaValidatorManager, err := poavalidatormanager.NewPoAValidatorManager(
		validatorManagerProxy.Address,
		l1AInfo.RPCClient,
	)
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
	initializeAndCompleteEndInitialPoAValidationV1(
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
	poaValidationID := initializeAndCompletePoAValidatorRegistrationV1(
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

	tx, err := validatorManagerProxy.ProxyAdmin.UpgradeAndCall(
		opts,
		validatorManagerProxy.Address,
		validatorManagerAddress,
		[]byte{},
	)
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

	// Check that previous validator is still registered
	// Migrate already registered validators to the new ValidatorManager version
	validatorManager, err := validatormanager.NewValidatorManager(validatorManagerProxy.Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())

	validationID, err := validatorManager.RegisteredValidators(&bind.CallOpts{}, poaNodeID)
	Expect(err).Should(BeNil())
	Expect(validationID[:]).Should(Equal(poaValidationID[:]))

	// Check the subnetID was migrated
	subnetID, err := validatorManager.SubnetID(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(subnetID[:]).Should(Equal(l1AInfo.SubnetID[:]))

	//
	// Attempt to remove the PoA validator before migrating it
	//
	posStakingManager, err := istakingmanager.NewIStakingManager(stakingManagerAddress, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())

	opts, err = bind.NewKeyedTransactorWithChainID(ownerKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = posStakingManager.ForceInitiateValidatorRemoval(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err.Error()).Should(ContainSubstring("reverted"))

	//
	// Migrate the validator
	//

	opts, err = bind.NewKeyedTransactorWithChainID(ownerKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = validatorManager.MigrateFromV1(opts, poaValidationID, 0)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1AInfo, tx.Hash())

	//
	// Remove the PoA validator and re-register as a PoS validator
	//

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

func convertSubnetPoAV1(
	n *localnetwork.LocalNetwork,
	ctx context.Context,
	l1 interfaces.L1TestInfo,
	weights []uint64,
	senderKey *ecdsa.PrivateKey,
) (proxy network.ProxyAddress,
	nodes []utils.Node,
	validationIDs []ids.ID,
) {
	goLog.Println("Converting l1", l1.SubnetID)
	cChainInfo := n.GetPrimaryNetworkInfo()
	pClient := platformvm.NewClient(cChainInfo.NodeURIs[0])
	currentValidators, err := pClient.GetCurrentValidators(ctx, l1.SubnetID, nil)
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())

	// Reset the global binary data for better test isolation
	poavalidatormanager.PoAValidatorManagerBin = poavalidatormanager.PoAValidatorManagerMetaData.Bin

	vdrManagerAddress, tx, _, err := poavalidatormanager.DeployPoAValidatorManager(
		opts,
		l1.RPCClient,
		0, // ICMInitializable.Allowed
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1, tx.Hash())

	var vdrManagerProxyAdmin *proxyadmin.ProxyAdmin
	// Overwrite the manager address with the proxy address
	vdrManagerAddress, vdrManagerProxyAdmin = utils.DeployTransparentUpgradeableProxy(
		ctx,
		l1,
		senderKey,
		vdrManagerAddress,
	)

	proxy = network.ProxyAddress{
		Address:    vdrManagerAddress,
		ProxyAdmin: vdrManagerProxyAdmin,
	}

	poaValidatorManager, err := poavalidatormanager.NewPoAValidatorManager(vdrManagerAddress, l1.RPCClient)
	Expect(err).Should(BeNil())
	tx, err = poaValidatorManager.Initialize(
		opts,
		poavalidatormanager.ValidatorManagerSettings{
			L1ID:                   l1.SubnetID,
			ChurnPeriodSeconds:     uint64(0),
			MaximumChurnPercentage: uint8(20),
		},
		utils.PrivateKeyToAddress(senderKey),
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1, tx.Hash())

	tmpnetNodes := n.GetExtraNodes(len(weights))
	sort.Slice(tmpnetNodes, func(i, j int) bool {
		return string(tmpnetNodes[i].NodeID.Bytes()) < string(tmpnetNodes[j].NodeID.Bytes())
	})

	// Construct the converted l1 info
	destAddr, err := address.ParseToID(utils.DefaultPChainAddress)
	Expect(err).Should(BeNil())
	vdrs := make([]*txs.ConvertSubnetToL1Validator, len(tmpnetNodes))
	for i, node := range tmpnetNodes {
		signer, err := node.GetProofOfPossession()
		Expect(err).Should(BeNil())
		nodes = append(nodes, utils.Node{
			NodeID:  node.NodeID,
			NodePoP: signer,
			Weight:  weights[i],
		})
		vdrs[i] = &txs.ConvertSubnetToL1Validator{
			NodeID:  node.NodeID.Bytes(),
			Weight:  weights[i],
			Balance: units.Avax * 100,
			Signer:  *signer,
			RemainingBalanceOwner: warpMessage.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
			DeactivationOwner: warpMessage.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
		}
	}
	pChainWallet := n.GetPChainWallet()
	_, err = pChainWallet.IssueConvertSubnetToL1Tx(
		l1.SubnetID,
		l1.BlockchainID,
		vdrManagerAddress[:],
		vdrs,
	)
	Expect(err).Should(BeNil())

	l1 = n.AddSubnetValidators(tmpnetNodes, l1, true)

	utils.PChainProposerVMWorkaround(pChainWallet)
	utils.AdvanceProposerVM(ctx, l1, senderKey, 5)

	aggregator := n.GetSignatureAggregator()
	defer aggregator.Shutdown()
	validationIDs = initializeValidatorSetV1(
		ctx,
		senderKey,
		l1,
		utils.GetPChainInfo(cChainInfo),
		vdrManagerAddress,
		n.GetNetworkID(),
		aggregator,
		nodes,
	)

	// Remove the bootstrap nodes as l1 validators
	for _, vdr := range currentValidators {
		_, err := pChainWallet.IssueRemoveSubnetValidatorTx(vdr.NodeID, l1.SubnetID)
		Expect(err).Should(BeNil())
		for _, node := range n.Network.Nodes {
			if node.NodeID == vdr.NodeID {
				port := network.GetTmpnetNodePort(node)
				node.Flags[config.HTTPPortKey] = port
				goLog.Println("Restarting bootstrap node", node.NodeID)
				n.Network.RestartNode(ctx, logging.NoLog{}, node)
			}
		}
	}
	utils.PChainProposerVMWorkaround(pChainWallet)
	utils.AdvanceProposerVM(ctx, l1, senderKey, 5)

	return
}

func initializeValidatorSetV1(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	validatorManagerAddress common.Address,
	networkID uint32,
	signatureAggregator *utils.SignatureAggregator,
	nodes []utils.Node,
) []ids.ID {
	log.Println("Initializing validator set", "l1", l1Info.SubnetID)
	initialValidators := make([]warpMessage.SubnetToL1ConversionValidatorData, len(nodes))
	initialValidatorsABI := make([]poavalidatormanager.InitialValidator, len(nodes))
	for i, node := range nodes {
		initialValidators[i] = warpMessage.SubnetToL1ConversionValidatorData{
			NodeID:       node.NodeID.Bytes(),
			BLSPublicKey: node.NodePoP.PublicKey,
			Weight:       nodes[i].Weight,
		}
		initialValidatorsABI[i] = poavalidatormanager.InitialValidator{
			NodeID:       node.NodeID.Bytes(),
			BlsPublicKey: node.NodePoP.PublicKey[:],
			Weight:       nodes[i].Weight,
		}
	}

	l1ConversionData := warpMessage.SubnetToL1ConversionData{
		SubnetID:       l1Info.SubnetID,
		ManagerChainID: l1Info.BlockchainID,
		ManagerAddress: validatorManagerAddress[:],
		Validators:     initialValidators,
	}
	l1ConversionDataABI := poavalidatormanager.ConversionData{
		L1ID:                         l1Info.SubnetID,
		ValidatorManagerBlockchainID: l1Info.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators:            initialValidatorsABI,
	}
	l1ConversionID, err := warpMessage.SubnetToL1ConversionID(l1ConversionData)
	Expect(err).Should(BeNil())
	l1ConversionSignedMessage := utils.ConstructL1ConversionMessage(
		l1ConversionID,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the L1
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", l1ConversionDataABI, uint32(0))
	Expect(err).Should(BeNil())
	receipt := utils.CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1Info,
		validatorManagerAddress,
		l1ConversionSignedMessage.Bytes(),
	)

	manager, err := poavalidatormanager.NewPoAValidatorManager(validatorManagerAddress, l1Info.RPCClient)
	Expect(err).Should(BeNil())
	initialValidatorCreatedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		manager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	var validationIDs []ids.ID
	for i := range nodes {
		validationIDs = append(validationIDs, l1Info.SubnetID.Append(uint32(i)))
	}

	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(nodes[0].Weight))

	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(validationIDs[0]))

	return validationIDs
}

func initializeAndCompleteEndInitialPoAValidationV1(
	ctx context.Context,
	signatureAggregator *utils.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) {
	log.Println("Initializing initial validator removal")
	utils.WaitMinStakeDuration(ctx, l1Info, fundedKey)
	opts, err := bind.NewKeyedTransactorWithChainID(ownerKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := validatorManager.InitializeEndValidation(
		opts,
		validationID,
	)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, l1Info, tx.Hash())

	validatorRemovalEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight).Should(Equal(weight))

	// Gather subnet-evm Warp signatures for the SetL1ValidatorWeightMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	unsignedMessage := utils.ExtractWarpMessageFromLog(ctx, receipt, l1Info)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		l1Info.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	pchainWallet.IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
	utils.PChainProposerVMWorkaround(pchainWallet)
	utils.AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a L1ValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing initial validator removal")
	registrationSignedMessage := utils.ConstructL1ValidatorRegistrationMessageForInitialValidator(
		validationID,
		index,
		false,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the L1
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndValidation", uint32(0))
	Expect(err).Should(BeNil())
	receipt = utils.CallWarpReceiver(
		ctx,
		callData,
		ownerKey,
		l1Info,
		validatorManagerAddress,
		registrationSignedMessage.Bytes(),
	)

	// Check that the validator is has been delisted from the staking contract
	validationEndedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidationPeriodEnded,
	)
	Expect(err).Should(BeNil())
	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

func initializeAndCompletePoAValidatorRegistrationV1(
	ctx context.Context,
	signatureAggregator *utils.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	expiry uint64,
	node utils.Node,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) ids.ID {
	// Initiate validator registration
	// receipt, validationID := InitializePoAValidatorRegistration(
	// 	ctx,
	// 	ownerKey,
	// 	l1Info,
	// 	node,
	// 	expiry,
	// 	validatorManager,
	// )

	opts, err := bind.NewKeyedTransactorWithChainID(ownerKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := validatorManager.InitializeValidatorRegistration(
		opts,
		poavalidatormanager.ValidatorRegistrationInput{
			NodeID:             node.NodeID[:],
			RegistrationExpiry: expiry,
			BlsPublicKey:       node.NodePoP.PublicKey[:],
		},
		node.Weight,
	)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, l1Info, tx.Hash())
	registrationInitiatedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	validationID := ids.ID(registrationInitiatedEvent.ValidationID)

	// Gather subnet-evm Warp signatures for the RegisterL1ValidatorMessage & relay to the P-Chain
	signedWarpMessage := utils.ConstructSignedWarpMessage(ctx, receipt, l1Info, pChainInfo, nil, signatureAggregator)

	_, err = pchainWallet.IssueRegisterL1ValidatorTx(
		100*units.Avax,
		node.NodePoP.ProofOfPossession,
		signedWarpMessage.Bytes(),
	)
	Expect(err).Should(BeNil())
	utils.PChainProposerVMWorkaround(pchainWallet)
	utils.AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a L1ValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing validator registration")
	registrationSignedMessage := utils.ConstructL1ValidatorRegistrationMessage(
		validationID,
		expiry,
		node,
		true,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the L1
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	receipt = utils.CallWarpReceiver(
		ctx,
		callData,
		ownerKey,
		l1Info,
		validatorManagerAddress,
		registrationSignedMessage.Bytes(),
	)

	// Check that the validator is registered in the staking contract
	registrationEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidationPeriodRegistered,
	)
	Expect(err).Should(BeNil())
	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))

	return validationID
}
