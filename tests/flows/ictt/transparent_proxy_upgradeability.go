package ictt

import (
	"context"
	"math/big"

	erc20tokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	erc20tokenhomeupgradeable "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/ERC20TokenHomeUpgradeable"
	erc20tokenremote "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenRemote/ERC20TokenRemote"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

/**
 * Deploy an upgradeable ERC20TokenHome on the primary network
 * Deploys a transparent upgradeable proxy that uses the ERC20TokenHome logic contract
 * Deploys a proxy admin contract to manage the upgradeable proxy
 * Deploy an ERC20TokenRemote to L1 A
 * Transfers example erc20 tokens from the primary network to L1 A
 * Deploy a new ERC20TokenHome logic contract on the primary network
 * Upgrade the transparent upgradeable proxy to use the new logic contract
 * Transfer tokens from L1 A back to the primary network
 * Check that the transfer was successful, and expected balances are correct
 */

func TransparentUpgradeableProxy(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
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

	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey,
		cChainInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())
	implAddress, tx, _, err := erc20tokenhomeupgradeable.DeployERC20TokenHomeUpgradeable(
		opts,
		cChainInfo.RPCClient,
		uint8(1),
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Deploy a TransparentUpgradeableProxy contract on primary network for the ERC20TokenHome logic contract
	erc20TokenHomeAddress, proxyAdmin := utils.DeployTransparentUpgradeableProxy(
		ctx,
		cChainInfo,
		fundedKey,
		implAddress,
	)
	erc20TokenHome, err := erc20tokenhome.NewERC20TokenHome(erc20TokenHomeAddress, cChainInfo.RPCClient)
	Expect(err).Should(BeNil())

	tx, err = erc20TokenHome.Initialize(
		opts,
		teleporter.TeleporterRegistryAddress(cChainInfo),
		fundedAddress,
		big.NewInt(1),
		exampleERC20Address,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Deploy the ERC20TokenRemote contract on L1 A
	erc20TokenRemoteAddress, erc20TokenRemote := utils.DeployERC20TokenRemote(
		ctx,
		teleporter,
		fundedKey,
		l1AInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		tokenDecimals,
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

	// Send a transfer from primary network to L1 A
	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

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

	// Check that the transfer was successful, and expected balances are correct
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

	// Deploy a new ERC20TokenHome logic contract on primary network
	newLogic, tx, _, err := erc20tokenhomeupgradeable.DeployERC20TokenHomeUpgradeable(
		opts,
		cChainInfo.RPCClient,
		uint8(1),
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Upgrade the TransparentUpgradeableProxy contract to use the new logic contract
	tx, err = proxyAdmin.UpgradeAndCall(opts, erc20TokenHomeAddress, newLogic, []byte{})
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Send a transfer from L1 A back to primary network
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

	// Check that the transfer was successful, and expected balances are correct
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
