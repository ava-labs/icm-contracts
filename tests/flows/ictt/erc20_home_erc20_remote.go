package ictt

import (
	"context"
	"math/big"

	erc20tokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	erc20tokenremote "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenRemote/ERC20TokenRemote"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

/**
 * Deploy an ERC20TokenHome on the primary network
 * Deploys ERC20TokenRemote to L1 A
 * Transfers C-Chain example ERC20 tokens to L1 A
 * Transfer tokens from L1 A to C-Chain
 */
func ERC20TokenHomeERC20TokenRemote(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on the primary network as the token to be transferred
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

	// Token representation on L1 A will have same name, symbol, and decimals
	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenRemote to L1 A
	erc20TokenRemoteAddress, erc20TokenRemote := utils.DeployERC20TokenRemote(
		ctx,
		teleporter,
		fundedKey,
		l1AInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		exampleERC20Decimals,
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
		erc20TokenHomeAddress,
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
	input := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:            l1AInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenRemoteAddress,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             exampleERC20Address,
		PrimaryFee:                         big.NewInt(1e18),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
	}
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))

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

	// Transfer back to home chain
	// Fund recipient with gas tokens on L1 A
	utils.SendNativeTransfer(
		ctx,
		l1AInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputB := erc20tokenremote.SendTokensInput{
		DestinationBlockchainID:            cChainInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenHomeAddress,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             erc20TokenRemoteAddress,
		PrimaryFee:                         big.NewInt(1e10),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
	}

	receipt, transferredAmount = utils.SendERC20TokenRemote(
		ctx,
		l1AInfo,
		erc20TokenRemote,
		erc20TokenRemoteAddress,
		inputB,
		utils.BigIntSub(transferredAmount, inputB.PrimaryFee),
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

	utils.CheckERC20TokenHomeWithdrawal(
		ctx,
		erc20TokenHomeAddress,
		exampleERC20,
		receipt,
		recipientAddress,
		transferredAmount,
	)

	// Check that the recipient received the tokens
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(transferredAmount))
}
