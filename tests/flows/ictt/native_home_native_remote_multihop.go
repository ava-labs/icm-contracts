package ictt

import (
	"context"
	"math/big"

	nativetokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/NativeTokenHome"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

/*
*
  - Deploy a NativeTokenHome on the primary network
  - Deploys NativeTokenRemote to L1 A and L1 B
  - Transfers native tokens from the C-Chain to L1 A as L1 A's native token
  - Transfers native tokens from the C-Chain to L1 B as L1 B's native token
    to collateralize the L1 B token transferrer
  - Transfer tokens from L1 A to L1 B through multi-hop
  - Transfer back tokens from L1 B to L1 A through multi-hop
*/
func NativeTokenHomeNativeTokenRemoteMultiHop(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, l1BInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// decimalsShift is always 0 for native to native
	decimalsShift := uint8(0)

	// Deploy an example WAVAX on the primary network
	wavaxAddress, wavax := utils.DeployWrappedNativeToken(
		ctx,
		fundedKey,
		cChainInfo,
		"AVAX",
	)

	// Create a NativeTokenHome on the primary network
	nativeTokenHomeAddress, nativeTokenHome := utils.DeployNativeTokenHome(
		ctx,
		teleporter,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Deploy a NativeTokenRemote to L1 A
	nativeTokenRemoteAddressA, nativeTokenRemoteA := utils.DeployNativeTokenRemote(
		ctx,
		teleporter,
		l1AInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHomeAddress,
		utils.NativeTokenDecimals,
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
		nativeTokenHomeAddress,
		utils.NativeTokenDecimals,
		initialReserveImbalance,
		burnedFeesReportingRewardPercentage,
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Register both NativeTokenDestinations on the NativeTokenHome
	collateralAmountA := utils.RegisterTokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		nativeTokenHomeAddress,
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
		nativeTokenHomeAddress,
		l1BInfo,
		nativeTokenRemoteAddressB,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnRemote,
		fundedKey,
		aggregator,
	)

	// Add collateral for both NativeTokenDestinations
	utils.AddCollateralToNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		l1AInfo.BlockchainID,
		nativeTokenRemoteAddressA,
		collateralAmountA,
		fundedKey,
	)

	utils.AddCollateralToNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		l1BInfo.BlockchainID,
		nativeTokenRemoteAddressB,
		collateralAmountB,
		fundedKey,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))

	// Send tokens from C-Chain to L1 A
	inputA := nativetokenhome.SendTokensInput{
		DestinationBlockchainID:            l1AInfo.BlockchainID,
		DestinationTokenTransferrerAddress: nativeTokenRemoteAddressA,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             wavaxAddress,
		PrimaryFee:                         big.NewInt(1e18),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
	}

	receipt, transferredAmountA := utils.SendNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		wavax,
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
	inputB := nativetokenhome.SendTokensInput{
		DestinationBlockchainID:            l1BInfo.BlockchainID,
		DestinationTokenTransferrerAddress: nativeTokenRemoteAddressB,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             wavaxAddress,
		PrimaryFee:                         big.NewInt(1e18),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
	}
	receipt, transferredAmountB := utils.SendNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		wavax,
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
	amountToSendA := new(big.Int).Div(transferredAmountA, big.NewInt(2))

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
		amountToSendA,
		big.NewInt(0),
		aggregator,
	)

	// Again, send half of the received amount to account for gas expenses
	amountToSendB := new(big.Int).Div(amountToSendA, big.NewInt(2))
	secondaryFeeAmount := new(big.Int).Div(amountToSendB, big.NewInt(4))

	// Multi-hop transfer back to L1 A
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
		amountToSendB,
		secondaryFeeAmount,
		aggregator,
	)
}
