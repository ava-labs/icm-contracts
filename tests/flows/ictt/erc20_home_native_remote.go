package ictt

import (
	"context"
	"math/big"

	erc20tokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	nativetokenremote "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenRemote/NativeTokenRemote"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

var (
	decimalsShift           = uint8(1)
	tokenMultiplier         = utils.GetTokenMultiplier(decimalsShift)
	initialReserveImbalance = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e6))

	// These two should be changed together
	multiplyOnRemote       = true
	erc20TokenHomeDecimals = utils.NativeTokenDecimals - decimalsShift

	burnedFeesReportingRewardPercentage = big.NewInt(1)
)

/**
 * Deploy a ERC20Token on the primary network
 * Deploys NativeTokenRemote to L1 A and L1 B
 * Transfers C-Chain example ERC20 tokens to L1 A as L1 A's native token
 * Transfer back tokens from L1 A to C-Chain
 */
func ERC20TokenHomeNativeTokenRemote(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
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

	// Create an ERC20TokenHome for transferring the ERC20 token
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

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	collateralAmount := utils.RegisterTokenRemoteOnHome(
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

	utils.AddCollateralToERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		l1AInfo.BlockchainID,
		nativeTokenRemoteAddressA,
		collateralAmount,
		fundedKey,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	recipientKey.ECDH()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to L1 A
	input := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:            l1AInfo.BlockchainID,
		DestinationTokenTransferrerAddress: nativeTokenRemoteAddressA,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             exampleERC20Address,
		PrimaryFee:                         big.NewInt(1e18),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
	}

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))
	receipt, transferredAmount := utils.SendERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		input,
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
	utils.CheckBalance(ctx, recipientAddress, transferredAmount, l1AInfo.RPCClient)

	// Send back to the home chain and check that ERC20TokenHome received the tokens
	input_A := nativetokenremote.SendTokensInput{
		DestinationBlockchainID:            cChainInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenHomeAddress,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             nativeTokenRemoteAddressA,
		PrimaryFee:                         big.NewInt(1e10),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
	}
	// Send half of the received amount to account for gas expenses
	amountToSendA := new(big.Int).Div(transferredAmount, big.NewInt(2))
	receipt, transferredAmount = utils.SendNativeTokenRemote(
		ctx,
		l1AInfo,
		nativeTokenRemoteA,
		nativeTokenRemoteAddressA,
		input_A,
		amountToSendA,
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
	scaledAmount := utils.RemoveTokenScaling(tokenMultiplier, multiplyOnRemote, transferredAmount)
	utils.CheckERC20TokenHomeWithdrawal(
		ctx,
		erc20TokenHomeAddress,
		exampleERC20,
		receipt,
		recipientAddress,
		scaledAmount,
	)

	// Check that the recipient received the tokens
	balance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(scaledAmount))
}
