package staking

import (
	"context"
	"log"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	exampleerc20 "github.com/ava-labs/icm-contracts/abi-bindings/go/mocks/ExampleERC20"
	erc20tokenslotauctionmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/ERC20TokenSlotAuctionManager"
	islotauctionmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/interfaces/ISlotAuctionManager"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"

	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	. "github.com/onsi/gomega"
)

/*
 * Registers an ERC20 auction validator on a L1. The steps are as follows:
 * - Deploy the SlotAuctionManager
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
 * - Verify that the validator is delisted from the auction contract
*/
func ERC20TokenSlotAuctionManager(network *localnetwork.LocalNetwork) {
	// Get the L1s info
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	_, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	balance := 100 * units.Avax
	nodes, initialValidationIDs := network.ConvertSubnet(
		ctx,
		l1AInfo,
		utils.ERC20TokenSlotAuctionManager,
		[]uint64{units.Schmeckle, units.Schmeckle, units.Schmeckle,
			units.Schmeckle, units.Schmeckle, 2000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		[]uint64{balance, balance, balance, balance, balance, balance},
		fundedKey,
		false,
	)

	validatorManagerProxy, slotAuctionManagerProxy := network.GetValidatorManager(l1AInfo.SubnetID)
	erc20TokenSlotAuctionAddress := slotAuctionManagerProxy.Address

	erc20TokenSlotAuctionManager, err := erc20tokenslotauctionmanager.NewERC20TokenSlotAuctionManager(
		erc20TokenSlotAuctionAddress,
		l1AInfo.RPCClient,
	)

	Expect(err).Should(BeNil())

	exampleERC20Address, err := erc20TokenSlotAuctionManager.Erc20(&bind.CallOpts{})
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
	account1PrivKey, _ := utils.CreateAndFundNewAddressWithERC20(ctx, l1AInfo, fundedKey, exampleERC20, big.NewInt(1000))
	account2PrivKey, _ := utils.CreateAndFundNewAddressWithERC20(ctx, l1AInfo, fundedKey, exampleERC20, big.NewInt(500))

	iSlotAuctionManager, err := islotauctionmanager.NewISlotAuctionManager(
		slotAuctionManagerProxy.Address,
		l1AInfo.RPCClient,
	)
	Expect(err).Should(BeNil())

	utils.InitiateAndCompleteEndInitialAuctionValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		iSlotAuctionManager,
		erc20TokenSlotAuctionAddress,
		validatorManagerProxy.Address,
		initialValidationIDs[0],
		0,
		nodes[0].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	utils.InitiateAndCompleteEndInitialAuctionValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		iSlotAuctionManager,
		erc20TokenSlotAuctionAddress,
		validatorManagerProxy.Address,
		initialValidationIDs[1],
		1,
		nodes[1].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	utils.InitiateAndCompleteEndInitialAuctionValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		iSlotAuctionManager,
		erc20TokenSlotAuctionAddress,
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
		iSlotAuctionManager,
	)

	utils.PlaceBidOnERC20TokenAuction(
		ctx,
		account1PrivKey,
		l1AInfo,
		big.NewInt(50),
		exampleERC20,
		nodes[0],
		erc20TokenSlotAuctionManager,
		erc20TokenSlotAuctionAddress,
	)
	utils.PlaceBidOnERC20TokenAuction(
		ctx,
		account2PrivKey,
		l1AInfo,
		big.NewInt(10),
		exampleERC20,
		nodes[1],
		erc20TokenSlotAuctionManager,
		erc20TokenSlotAuctionAddress,
	)
	utils.PlaceBidOnERC20TokenAuction(
		ctx,
		account2PrivKey,
		l1AInfo,
		big.NewInt(20),
		exampleERC20,
		nodes[2],
		erc20TokenSlotAuctionManager,
		erc20TokenSlotAuctionAddress,
	)

	utils.EndAuctionAndCompleteValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		iSlotAuctionManager,
		erc20TokenSlotAuctionAddress,
		validatorManagerProxy.Address,
		network.GetPChainWallet(),
		network.GetNetworkID(),
		nodes,
	)
}
