package ictt

import (
	"context"
	"math/big"

	nativetokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/NativeTokenHome"
	erc20tokenremote "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenRemote/ERC20TokenRemote"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys ERC20TokenRemote to L1 A
 * Transfers C-Chain native tokens to L1 A
 * Transfer back tokens from L1 A to C-Chain
 */
func NativeTokenHomeERC20TokenRemote(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example WAVAX on the primary network
	wavaxAddress, wavax := utils.DeployWrappedNativeToken(
		ctx,
		fundedKey,
		cChainInfo,
		"AVAX",
	)

	// Create a NativeTokenHome for transferring the native token
	nativeTokenHomeAddress, nativeTokenHome := utils.DeployNativeTokenHome(
		ctx,
		teleporter,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Token representation on L1 A will have same name, symbol, and decimals
	tokenName, err := wavax.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := wavax.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := wavax.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenRemote to L1 A
	erc20TokenRemoteAddress, erc20TokenRemote := utils.DeployERC20TokenRemote(
		ctx,
		teleporter,
		fundedKey,
		l1AInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHomeAddress,
		18,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		nativeTokenHomeAddress,
		l1AInfo,
		erc20TokenRemoteAddress,
		fundedKey,
		aggregator,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on L1 A
	input := nativetokenhome.SendTokensInput{
		DestinationBlockchainID:            l1AInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenRemoteAddress,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             wavaxAddress,
		PrimaryFee:                         big.NewInt(1e18),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
	}

	// Send the tokens and verify expected events
	amount := big.NewInt(2e18)
	receipt, transferredAmount := utils.SendNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		wavax,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to L1 A and check for message delivery
	receipt = teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		cChainInfo,
		l1AInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	utils.CheckERC20TokenRemoteWithdrawal(
		ctx,
		erc20TokenRemote,
		receipt,
		recipientAddress,
		transferredAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenRemote.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(transferredAmount))

	// Fund recipient with gas tokens on L1 A
	utils.SendNativeTransfer(
		ctx,
		l1AInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputA := erc20tokenremote.SendTokensInput{
		DestinationBlockchainID:            cChainInfo.BlockchainID,
		DestinationTokenTransferrerAddress: nativeTokenHomeAddress,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             erc20TokenRemoteAddress,
		PrimaryFee:                         big.NewInt(1e10),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
	}

	// Send tokens on L1 A back for native tokens on C-Chain
	receipt, transferredAmount = utils.SendERC20TokenRemote(
		ctx,
		l1AInfo,
		erc20TokenRemote,
		erc20TokenRemoteAddress,
		inputA,
		utils.BigIntSub(transferredAmount, inputA.PrimaryFee),
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
	utils.CheckNativeTokenHomeWithdrawal(
		ctx,
		nativeTokenHomeAddress,
		wavax,
		receipt,
		transferredAmount,
	)

	utils.CheckBalance(ctx, recipientAddress, transferredAmount, cChainInfo.RPCClient)
}
