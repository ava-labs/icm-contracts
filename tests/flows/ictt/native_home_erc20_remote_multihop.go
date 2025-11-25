package ictt

import (
	"context"
	"math/big"

	nativetokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/NativeTokenHome"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"

	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys ERC20TokenRemote to L1 A and L1 B
 * Transfers C-Chain native tokens to L1 A
 * Transfer tokens from L1 A to L1 B through multi-hop
 * Brige back tokens from L1 B to L1 A through multi-hop
 */
func NativeTokenHomeERC20TokenRemoteMultiHop(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, l1BInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a NativeTokenHome on the primary network
	wavaxAddress, wavax := utils.DeployWrappedNativeToken(
		ctx,
		fundedKey,
		cChainInfo,
		"AVAX",
	)
	nativeTokenHomeAddress, nativeTokenHome := utils.DeployNativeTokenHome(
		ctx,
		teleporter,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Token representation on L1 B will have same name, symbol, and decimals
	tokenName, err := wavax.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := wavax.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := wavax.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenRemote on L1 A
	erc20TokenRemoteAddressA, erc20TokenRemoteA := utils.DeployERC20TokenRemote(
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

	// Deploy an ERC20TokenRemote on L1 B
	erc20TokenRemoteAddressB, erc20TokenRemoteB := utils.DeployERC20TokenRemote(
		ctx,
		teleporter,
		fundedKey,
		l1BInfo,
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

	// Register both ERC20Destinations on the NativeTokenHome
	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		nativeTokenHomeAddress,
		l1AInfo,
		erc20TokenRemoteAddressA,
		fundedKey,
		aggregator,
	)

	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		nativeTokenHomeAddress,
		l1BInfo,
		erc20TokenRemoteAddressB,
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
		DestinationTokenTransferrerAddress: erc20TokenRemoteAddressA,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             wavaxAddress,
		PrimaryFee:                         big.NewInt(1e10),
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
		utils.BigIntSub(amount, input.PrimaryFee),
		fundedKey,
	)

	// Relay the message to L1 B and check for message delivery
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
		erc20TokenRemoteA,
		receipt,
		recipientAddress,
		transferredAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenRemoteA.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(transferredAmount))

	// Send tokens from L1 A to recipient on L1 B through a multi-hop
	secondaryFeeAmount := new(big.Int).Div(transferredAmount, big.NewInt(4))
	utils.SendERC20TokenMultiHopAndVerify(
		ctx,
		teleporter,
		fundedKey,
		recipientKey,
		recipientAddress,
		l1AInfo,
		erc20TokenRemoteA,
		erc20TokenRemoteAddressA,
		l1BInfo,
		erc20TokenRemoteB,
		erc20TokenRemoteAddressB,
		cChainInfo,
		transferredAmount,
		secondaryFeeAmount,
		aggregator,
	)
}
