package ictt

import (
	"context"
	"math/big"

	nativetokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/NativeTokenHome"
	nativetokenremote "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenRemote/NativeTokenRemote"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys a NativeTokenRemote to L1 A
 * Transfers C-Chain native tokens to L1 A
 * Transfer back tokens from L1 A to C-Chain
 */
func NativeTokenHomeNativeDestination(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example WAVAX on the primary network
	cChainWAVAXAddress, wavax := utils.DeployWrappedNativeToken(
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
		cChainWAVAXAddress,
	)

	// Deploy a NativeTokenRemote to L1 A
	nativeTokenRemoteAddress, nativeTokenRemote := utils.DeployNativeTokenRemote(
		ctx,
		teleporter,
		l1AInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHomeAddress,
		18,
		initialReserveImbalance,
		burnedFeesReportingRewardPercentage,
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Register the NativeTokenRemote on the NativeTokenHome
	collateralAmount := utils.RegisterTokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		nativeTokenHomeAddress,
		l1AInfo,
		nativeTokenRemoteAddress,
		initialReserveImbalance,
		big.NewInt(1),
		multiplyOnRemote,
		fundedKey,
		aggregator,
	)

	utils.AddCollateralToNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		l1AInfo.BlockchainID,
		nativeTokenRemoteAddress,
		collateralAmount,
		fundedKey,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on L1 A that fully collateralize token transferrer with leftover tokens.
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))
	{
		input := nativetokenhome.SendTokensInput{
			DestinationBlockchainID:            l1AInfo.BlockchainID,
			DestinationTokenTransferrerAddress: nativeTokenRemoteAddress,
			Recipient:                          recipientAddress,
			PrimaryFeeTokenAddress:             cChainWAVAXAddress,
			PrimaryFee:                         big.NewInt(1e18),
			SecondaryFee:                       big.NewInt(0),
			RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
		}

		// Send initialReserveImbalance tokens to fully collateralize token transferrer and mint the remainder.
		receipt, _ := utils.SendNativeTokenHome(
			ctx,
			cChainInfo,
			nativeTokenHome,
			nativeTokenHomeAddress,
			wavax,
			input,
			amount,
			fundedKey,
		)

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

		utils.CheckBalance(
			ctx,
			recipientAddress,
			amount,
			l1AInfo.RPCClient,
		)
	}

	// Send tokens on L1 A back for native tokens on C-Chain
	{
		input_A := nativetokenremote.SendTokensInput{
			DestinationBlockchainID:            cChainInfo.BlockchainID,
			DestinationTokenTransferrerAddress: nativeTokenHomeAddress,
			Recipient:                          recipientAddress,
			PrimaryFeeTokenAddress:             nativeTokenRemoteAddress,
			PrimaryFee:                         big.NewInt(1e18),
			SecondaryFee:                       big.NewInt(0),
			RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
		}

		// Send half of the tokens back to C-Chain
		amount := big.NewInt(0).Div(amount, big.NewInt(2))
		receipt, transferredAmount := utils.SendNativeTokenRemote(
			ctx,
			l1AInfo,
			nativeTokenRemote,
			nativeTokenRemoteAddress,
			input_A,
			amount,
			recipientKey,
		)

		receipt = teleporter.RelayTeleporterMessage(
			ctx,
			receipt,
			l1AInfo,
			cChainInfo,
			true,
			fundedKey,
			nil,
			aggregator,
		)

		// Check that the recipient received the tokens
		homeAmount := transferredAmount
		utils.CheckNativeTokenHomeWithdrawal(
			ctx,
			nativeTokenHomeAddress,
			wavax,
			receipt,
			homeAmount,
		)

		utils.CheckBalance(ctx, recipientAddress, homeAmount, cChainInfo.RPCClient)
	}
}
