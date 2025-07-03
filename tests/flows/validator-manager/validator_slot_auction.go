package staking

import (
	"context"
	"log"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	exampleerc20 "github.com/ava-labs/icm-contracts/abi-bindings/go/mocks/ExampleERC20"
	slotauctionmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/SlotAuctionManager"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"

	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
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
	_, fundedKey := network.GetFundedAccountInfo() //funded address
	ctx := context.Background()
	pChainInfo := utils.GetPChainInfo(cChainInfo) //pchainInfo

	nodes, initialValidationIDs := network.ConvertSubnet(
		ctx,
		l1AInfo,
		utils.SlotAuctionManager,
		[]uint64{units.Schmeckle, units.Schmeckle, units.Schmeckle, units.Schmeckle, units.Schmeckle, 2000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		fundedKey,
		false,
	)

	validatorManagerProxy, slotAuctionManagerProxy := network.GetValidatorManager(l1AInfo.SubnetID)
	slotAuctionAddress := slotAuctionManagerProxy.Address

	slotAuctionManager, err := slotauctionmanager.NewSlotAuctionManager(
		slotAuctionAddress,
		l1AInfo.RPCClient,
	)
	Expect(err).Should(BeNil())

	exampleERC20Address, err := slotAuctionManager.TOKENCONTRACT(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	exampleERC20, err := exampleerc20.NewExampleERC20(exampleERC20Address, l1AInfo.RPCClient)
	Expect(err).Should(BeNil())

	// signatureAggregator := network.GetSignatureAggregator()
	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			l1AInfo.SubnetID,
		},
	)
	defer signatureAggregator.Shutdown()

	log.Println("Creating extra accounts with balances")

	Account1PrivKey, _ := utils.CreateAndFundNewAddress(ctx, l1AInfo, exampleERC20, fundedKey, big.NewInt(1000))
	Account2PrivKey, _ := utils.CreateAndFundNewAddress(ctx, l1AInfo, exampleERC20, fundedKey, big.NewInt(500))

	utils.InitiateAndCompleteEndInitialProofOfPurchaseValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		slotAuctionManager,
		slotAuctionAddress,
		validatorManagerProxy.Address,
		initialValidationIDs[0],
		0,
		nodes[0].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	utils.InitiateAndCompleteEndInitialProofOfPurchaseValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		slotAuctionManager,
		slotAuctionAddress,
		validatorManagerProxy.Address,
		initialValidationIDs[1],
		1,
		nodes[1].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	utils.InitiateAndCompleteEndInitialProofOfPurchaseValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		slotAuctionManager,
		slotAuctionAddress,
		validatorManagerProxy.Address,
		initialValidationIDs[2],
		2,
		nodes[2].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	utils.InitiateAuction(
		ctx,
		l1AInfo,
		fundedKey,
		2,
		units.Schmeckle,
		big.NewInt(0),
		big.NewInt(31556926),
		slotAuctionManager,
	)

	utils.PlaceBidOnAuction(
		ctx,
		Account1PrivKey,
		l1AInfo,
		big.NewInt(50),
		exampleERC20,
		nodes[0],
		slotAuctionManager,
		slotAuctionAddress,
		validatorManagerProxy.Address,
	)

	utils.PlaceBidOnAuction(
		ctx,
		Account2PrivKey,
		l1AInfo,
		big.NewInt(10),
		exampleERC20,
		nodes[1],
		slotAuctionManager,
		slotAuctionAddress,
		validatorManagerProxy.Address,
	)

	utils.PlaceBidOnAuction(
		ctx,
		Account2PrivKey,
		l1AInfo,
		big.NewInt(20),
		exampleERC20,
		nodes[2],
		slotAuctionManager,
		slotAuctionAddress,
		validatorManagerProxy.Address,
	)

	utils.EndAuction(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		slotAuctionManager,
		slotAuctionAddress,
		validatorManagerProxy.Address,
		exampleERC20,
		network.GetPChainWallet(),
		network.GetNetworkID(),
		nodes,
	)

	// OwnerOpts, _ := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)

}
