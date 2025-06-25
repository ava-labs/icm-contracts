package staking

import (
	// "context"
	"log"
	// "math/big"

	// "github.com/ava-labs/avalanchego/ids"
	// "github.com/ava-labs/avalanchego/utils/units"
	// exampleerc20 "github.com/ava-labs/icm-contracts/abi-bindings/go/mocks/ExampleERC20"

	// //emcoin "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/EmCoin"
	// slotauctionmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/SlotAuctionManager"

	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	// "github.com/ava-labs/icm-contracts/tests/utils"
	// "github.com/ava-labs/subnet-evm/accounts/abi/bind"
	// . "github.com/onsi/gomega"
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
	foo := network.NetworkID + 5
	log.Println(foo)
	// Get the L1s info
	// cChainInfo := network.GetPrimaryNetworkInfo()
	// l1AInfo, _ := network.GetTwoL1s()
	// _, fundedKey := network.GetFundedAccountInfo()
	// ctx := context.Background()
	// pChainInfo := utils.GetPChainInfo(cChainInfo)
	// var stakeAmount big.Int
	// stakeAmount.SetInt64(10)

	// nodes, _ := network.ConvertSubnet( // _ was initial validation ids
	// 	ctx,
	// 	l1AInfo,
	// 	utils.SlotAuctionManager,
	// 	[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
	// 	fundedKey,
	// 	false,
	// )

	// log.Println("Emre:", len(nodes))

	// _, slotAuctionManagerProxy := network.GetValidatorManager(l1AInfo.SubnetID) // _ was validator manager proxy
	// slotAuctionAddress := slotAuctionManagerProxy.Address

	// // _, err := validatormanager.NewValidatorManager(validatorManagerProxy.Address, l1AInfo.RPCClient)
	// // Expect(err).Should(BeNil())

	// slotAuctionManager, err := slotauctionmanager.NewSlotAuctionManager(
	// 	slotAuctionAddress,
	// 	l1AInfo.RPCClient,
	// )
	// Expect(err).Should(BeNil())

	// exampleERC20Address, err := slotAuctionManager.TOKENCONTRACT(&bind.CallOpts{})
	// Expect(err).Should(BeNil())

	// exampleERC20, err := exampleerc20.NewExampleERC20(exampleERC20Address, l1AInfo.RPCClient)
	// Expect(err).Should(BeNil())

	// signatureAggregator := utils.NewSignatureAggregator(
	// 	cChainInfo.NodeURIs[0],
	// 	[]ids.ID{
	// 		l1AInfo.SubnetID,
	// 	},
	// )

	// defer signatureAggregator.Shutdown()
	// opts, _ := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)

	// rewardAddress := opts.From

	// balance, _ := exampleERC20.BalanceOf(&bind.CallOpts{}, rewardAddress)
	// log.Println("Balance before: ", balance)

	// log.Println("Initiating and completing end Initial Proof of Purchase")
	// //Remove initial Validator
	// utils.InitiateAndCompleteEndInitialProofOfPurchaseValidation(
	// 	ctx,
	// 	signatureAggregator,
	// 	fundedKey,
	// 	l1AInfo,
	// 	pChainInfo,
	// 	slotAuctionManager,
	// 	slotAuctionAddress,
	// 	validatorManagerProxy.Address,
	// 	initialValidationIDs[0],
	// 	0,
	// 	nodes[0].Weight,
	// 	network.GetPChainWallet(),
	// 	network.GetNetworkID(),
	// )
	// log.Println("Starting addition of validator")

	// registrationInitiatedEvent := utils.InitiateAndCompleteERC20AuctionValidatorRegistration(
	// 	ctx,
	// 	signatureAggregator,
	// 	fundedKey,
	// 	l1AInfo,
	// 	&stakeAmount,
	// 	pChainInfo,
	// 	slotAuctionManager,
	// 	slotAuctionAddress,
	// 	validatorManagerProxy.Address,
	// 	opts.From,
	// 	exampleERC20,
	// 	nodes[0],
	// 	network.GetPChainWallet(),
	// 	network.GetNetworkID(),
	// )
	// balance, _ = exampleERC20.BalanceOf(&bind.CallOpts{}, rewardAddress)

	// nodes[0].NodeID = registrationInitiatedEvent.NodeID
	// nodes[0].Weight = registrationInitiatedEvent.Weight

	// log.Println("Balance after validator registration: ", balance)

	// utils.InitiateAndCompleteEndProofOfPurchaseValidation(
	// 	ctx,
	// 	signatureAggregator,
	// 	fundedKey,
	// 	l1AInfo,
	// 	pChainInfo,
	// 	slotAuctionManager,
	// 	slotAuctionAddress,
	// 	validatorManagerProxy.Address,
	// 	registrationInitiatedEvent.ValidationID,
	// 	registrationInitiatedEvent.RegistrationExpiry,
	// 	nodes[0],
	// 	network.GetPChainWallet(),
	// 	network.GetNetworkID(),
	// )

	// balance, _ = exampleERC20.BalanceOf(&bind.CallOpts{}, rewardAddress)

	// log.Println("Balance after validator removal: ", balance)

	// log.Println("Emre: AAAAAAAAAAAAAAAAAAAAA")

	// Expect(err).Should(BeNil())
	// utils.WaitForTransactionSuccess(ctx, l1AInfo, transactionInfo.Hash())
	// transactionInfo.Data()
	// log.Println(emCoin.BalanceOf(&bind.CallOpts{}, opts.From))

	// tempValID, _ := slotAuctionManager.TemporaryID(&bind.CallOpts{})

	// transactionInfo, err = slotAuctionManager.RemoveValidator(opts, tempValID)
	// Expect(err).Should(BeNil())
	// utils.WaitForTransactionSuccess(ctx, l1AInfo, transactionInfo.Hash())

	// log.Println(emCoin.BalanceOf(&bind.CallOpts{}, opts.From))

}
