package staking

import (
	"context"
	"log"
	"math/big"

	"crypto/ecdsa"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	emcoin "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/EmCoin"
	slotauctionmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/SlotAuctionManager"
	validatormanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/ValidatorManager"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	acp99manager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/ACP99Manager"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ethereum/go-ethereum/common"
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
func ValidatorSlotAuction(network *localnetwork.LocalNetwork) {
	// Get the L1s info
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	_, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	nodes, initialValidationIDs := network.ConvertSubnet(
		ctx,
		l1AInfo,
		utils.SlotAuctionManager,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		fundedKey,
		false,
	)
	
	validatorManagerProxy, slotAuctionManagerAddress := network.GetValidatorManager(l1AInfo.SubnetID)
	slotAuctionAddress := slotAuctionManagerAddress.Address

	validatorManager, err := validatormanager.NewValidatorManager(validatorManagerProxy.Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())

	vmAddr, _ := validatorManager.Owner(&bind.CallOpts{})

	log.Println("Validator manager owner:", vmAddr)
	log.Println("Slot auction owner:", slotAuctionAddress)

	slotAuctionManager, err := slotauctionmanager.NewSlotAuctionManager(
		slotAuctionAddress,
		l1AInfo.RPCClient,
		initialValidationIDs,
	)
	Expect(err).Should(BeNil())

	opts, _ := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)

	emCoinAddress, err := slotAuctionManager.TOKENCONTRACT(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	emCoin, err := emcoin.NewEmCoin(emCoinAddress, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			l1AInfo.SubnetID,
		},
	)

	defer signatureAggregator.Shutdown()
	node := nodes[0]

	// buys coin and approves it to be sent to slot auction
	transactionInfo, err := emCoin.BuyCoin(opts, big.NewInt(1000))
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1AInfo, transactionInfo.Hash())

	transactionInfo, err = emCoin.Approve(opts, slotAuctionAddress, big.NewInt(10))
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1AInfo, transactionInfo.Hash())

	val, err := slotAuctionManager.ValidatorCount(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	log.Println("Number of validators", val)

	//Remove validator

	validatorManager.

	log.Println("Initializing initial validator removal")
	






	transactionInfo, err = slotAuctionManager.BecomeValidator(
		opts,
		node.NodeID[:],
		node.NodePoP.PublicKey[:],
		slotauctionmanager.PChainOwner{},
		slotauctionmanager.PChainOwner{},
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1AInfo, transactionInfo.Hash())

	val, err = slotAuctionManager.ValidatorCount(&bind.CallOpts{})
	log.Println("Number of validators:", val)
	Expect(err).Should(BeNil())

	//signatureAggregator.CreateSignedMessage()

}

func InitiateAndCompleteEndInitialPoPValidation(
	ctx context.Context,
	signatureAggregator *utils.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	slotAuctionManager *slotauctionmanager.SlotAuctionManager, //replace this with auctionManager
	stakingManagerAddress common.Address,
	validatorManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
	pchainWallet pwallet.Wallet,
	networkID uint32,
	validatorManager *validatormanager.ValidatorManager,
) {
	log.Println("Initializing initial validator removal")
	utils.WaitMinStakeDuration(ctx, l1Info, fundedKey)
	receipt := ForceInitiateEndPoPValidation(
		ctx,
		fundedKey,
		l1Info,
		slotAuctionManager,
		validationID,
	)
	acp99Manager, err := acp99manager.NewACP99Manager(validatorManagerAddress, l1Info.RPCClient)
	Expect(err).Should(BeNil())
	validatorRemovalEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		acp99Manager.ParseInitiatedValidatorRemoval,
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
	receipt = CompleteEndPoSValidation(
		ctx,
		fundedKey,
		l1Info,
		stakingManagerAddress,
		registrationSignedMessage,
	)

	// Check that the validator is has been delisted from the staking contract
	validationEndedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		acp99Manager.ParseCompletedValidatorRemoval,
	)
	Expect(err).Should(BeNil())
	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

func ForceInitiateEndPoPValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	slotAuctionManager *slotauctionmanager.SlotAuctionManager, 
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := slotAuctionManager.InitiateRemoveInitialValidator(
		opts,
		validationID,
	)
	Expect(err).Should(BeNil())
	return utils.WaitForTransactionSuccess(ctx, l1, tx.Hash())
}