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

/**
 * Deploys an ERC20TokenHome contract on the C-Chain
 * Deploys an ERC20TokenRemote contract on L1 A
 * Check sending to  unregistered remote fails
 * Register the ERC20TokenRemote to home contract
 * Check sending to non-collateralized remote fails
 * Collateralize the remote
 * Check sending to collateralized remote succeeds and withdraws with correct scale.
 */
func RegistrationAndCollateralCheck(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
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

	// Create an ERC20TokenHome for transferring the ERC20 token
	erc20TokenHomeAddress, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		teleporter,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		erc20TokenHomeDecimals,
	)

	// Deploy a NativeTokenRemote to L1 A
	nativeTokenRemoteAddressA, _ := utils.DeployNativeTokenRemote(
		ctx,
		teleporter,
		l1AInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		erc20TokenHomeDecimals,
		initialReserveImbalance,
		burnedFeesReportingRewardPercentage,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
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

	initialBalance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())

	// Send the tokens and expect for failure since TokenRemote instance is not registered.
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, cChainInfo.EVMChainID)
	Expect(err).Should(BeNil())
	_, err = erc20TokenHome.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(ContainSubstring(ErrRemoteNotRegistered))

	// Check the balance of the ERC20TokenHome to ensure it was not changed
	balance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(balance, initialBalance)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Register the NativeTokenRemote to the ERC20TokenHome
	collateralNeeded := utils.RegisterTokenRemoteOnHome(
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

	// Try sending again and expect failure since remote is not collateralized
	_, err = erc20TokenHome.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(ContainSubstring(ErrNonZeroCollateralNeeded))

	// Check the balance of the ERC20TokenHome to ensure it was not changed
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(balance, initialBalance)

	// Add collateral to the ERC20TokenHome
	utils.AddCollateralToERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		l1AInfo.BlockchainID,
		nativeTokenRemoteAddressA,
		collateralNeeded,
		fundedKey,
	)

	// Check the balance of the ERC20TokenHome to ensure it was increased by the collateral amount.
	// Also set the new initial balance before sending tokens
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(balance, initialBalance.Add(initialBalance, collateralNeeded))

	// Send the tokens and expect success now that collateral is added
	utils.ERC20DecimalsApprove(
		ctx,
		exampleERC20,
		erc20TokenHomeAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		cChainInfo,
		fundedKey,
	)
	tx, err := erc20TokenHome.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())
	event, err := utils.GetEventFromLogs(receipt.Logs, erc20TokenHome.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(fundedKey.PublicKey)))

	// Compute the scaled amount
	scaledAmount := utils.GetScaledAmountFromERC20TokenHome(
		erc20TokenHome,
		input.DestinationBlockchainID,
		input.DestinationTokenTransferrerAddress,
		amount,
	)
	utils.ExpectBigEqual(event.Amount, scaledAmount)

	// Check the balance of the ERC20TokenHome increased by the transferred amount
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(balance, big.NewInt(0).Add(initialBalance, amount))

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
	utils.CheckBalance(ctx, recipientAddress, scaledAmount, l1AInfo.RPCClient)
}
