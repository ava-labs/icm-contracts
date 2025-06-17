package ictt

import (
	"context"
	"math/big"

	erc20tokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

/*
*
  - Deploy a ERC20TokenHome on the primary network
  - Deploys NativeTokenRemote toA and L1 B
  - Transfers C-Chain example ERC20 tokens to L1 A as L1 A's native token
  - Transfers C-Chain example ERC20 tokens to L1 B as L1 B's native token
    to collateralize the token transferrer on L1 B
  - Transfer tokens from L1 A to L1 B through multi-hop
  - Transfer back tokens from L1 B to L1 A through multi-hop
*/
func ERC20TokenHomeNativeTokenRemoteMultiHop(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, l1BInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on L1 A as the token to be transferred
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20Decimals(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHomeDecimals,
	)

	exampleERC20Decimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	erc20TokenHomeAddress, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		teleporter,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		exampleERC20Decimals,
	)

	// Deploy a NativeTokenRemote to L1 A
	nativeTokenRemoteAddressA, nativeTokenRemoteA := utils.DeployNativeTokenRemote(
		ctx,
		teleporter,
		l1AInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		exampleERC20Decimals,
		initialReserveImbalance,
		burnedFeesReportingRewardPercentage,
	)

	// Deploy a NativeTokenRemote to L1 B
	nativeTokenRemoteAddressB, nativeTokenRemoteB := utils.DeployNativeTokenRemote(
		ctx,
		teleporter,
		l1BInfo,
		"SUBB",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		exampleERC20Decimals,
		initialReserveImbalance,
		burnedFeesReportingRewardPercentage,
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Register both NativeTokenDestinations on the ERC20TokenHome
	collateralAmountA := utils.RegisterTokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		erc20TokenHomeAddress,
		l1AInfo,
		nativeTokenRemoteAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnRemote,
		fundedKey,
		aggregator,
	)

	collateralAmountB := utils.RegisterTokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		erc20TokenHomeAddress,
		l1BInfo,
		nativeTokenRemoteAddressB,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnRemote,
		fundedKey,
		aggregator,
	)

	// Add collateral for both NativeTokenDestinations
	utils.AddCollateralToERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		l1AInfo.BlockchainID,
		nativeTokenRemoteAddressA,
		collateralAmountA,
		fundedKey,
	)

	utils.AddCollateralToERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		l1BInfo.BlockchainID,
		nativeTokenRemoteAddressB,
		collateralAmountB,
		fundedKey,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// These are set during the initial transferring, and used in the multi-hop transfers
	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))

	// Send tokens from C-Chain to L1 A
	inputA := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:            l1AInfo.BlockchainID,
		DestinationTokenTransferrerAddress: nativeTokenRemoteAddressA,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             exampleERC20Address,
		PrimaryFee:                         big.NewInt(1e18),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
	}

	receipt, transferredAmountA := utils.SendERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		inputA,
		amount,
		fundedKey,
	)

	// Relay the message to L1 A and check for a native token mint withdrawal
	teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		cChainInfo,
		l1AInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	// Verify the recipient received the tokens
	utils.CheckBalance(ctx, recipientAddress, transferredAmountA, l1AInfo.RPCClient)

	// Send tokens from C-Chain to L1 B
	inputB := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:            l1BInfo.BlockchainID,
		DestinationTokenTransferrerAddress: nativeTokenRemoteAddressB,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             exampleERC20Address,
		PrimaryFee:                         big.NewInt(1e18),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
	}

	receipt, transferredAmountB := utils.SendERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		inputB,
		amount,
		fundedKey,
	)

	// Relay the message to L1 B and check for a native token mint withdrawal
	teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		cChainInfo,
		l1BInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	// Verify the recipient received the tokens
	utils.CheckBalance(ctx, recipientAddress, transferredAmountB, l1BInfo.RPCClient)

	// Multi-hop transfer to L1 B
	// Send half of the received amount to account for gas expenses
	amountToSend := new(big.Int).Div(transferredAmountA, big.NewInt(2))

	utils.SendNativeMultiHopAndVerify(
		ctx,
		teleporter,
		fundedKey,
		recipientAddress,
		l1AInfo,
		nativeTokenRemoteA,
		nativeTokenRemoteAddressA,
		l1BInfo,
		nativeTokenRemoteB,
		nativeTokenRemoteAddressB,
		cChainInfo,
		amountToSend,
		big.NewInt(0),
		aggregator,
	)

	// Multi-hop transfer back to L1 A
	secondaryFeeAmount := new(big.Int).Div(amountToSend, big.NewInt(4))
	utils.SendNativeMultiHopAndVerify(
		ctx,
		teleporter,
		fundedKey,
		recipientAddress,
		l1BInfo,
		nativeTokenRemoteB,
		nativeTokenRemoteAddressB,
		l1AInfo,
		nativeTokenRemoteA,
		nativeTokenRemoteAddressA,
		cChainInfo,
		amountToSend,
		secondaryFeeAmount,
		aggregator,
	)
}
