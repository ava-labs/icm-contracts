package ethereum_icm_verification

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/log"
	. "github.com/onsi/gomega"
)

func EthereumICMVerification(
	avalancheNetwork *network.LocalAvalancheNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	// 1. Deploy an AvalancheValidatorSetRegistry contract on the Ethereum network
	// 2. Register the Avalanche L1 network's validator set in the registry via an ICM message from the P-Chain, self-signed by the validator set of the Avalanche network
	// 3. Deploy an adapted Warp contract to Ethereum that can verify ICM messages against validator set signatures
	// 4. Deploy a Teleporter contract to Ethereum using the adapted Warp contract
	// 5. Send a transaction on the Avalanche L1 network to send an ICM message to the Ethereum network
	// 6. Get an aggregate signature of the ICM message using the ICM signature aggregator
	// 7. Broadcast the signed ICM message to the Ethereum network, and confirm successful inclusion and verification
	// 8. Add a new validator to the Avalanche L1 network, representing 40% of the new total weight
	// 9. Update the validator set registered on the Ethereum registry via a validator set update message signed by sufficient weight of the previous validator set
	// 10. Send and verify another ICM message from the Avalanche L1 network to the Ethereum network
	ctx := context.Background()
	_, ethereumFundedKey := ethereumNetwork.GetFundedAccountInfo()

	// Start a signature aggregator for the Avalanche network to be able to generate signed ICM messages.
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	Expect(len(primaryNetworkInfo.NodeURIs)).Should(BeNumerically(">", 0), "No NodeURIs found in primaryNetworkInfo")
	pChainInfo := utils.GetPChainInfo(primaryNetworkInfo)
	l1Infos := avalancheNetwork.GetL1Infos()
	l1Info := l1Infos[0]
	Expect(len(l1Infos)).Should(BeNumerically(">", 0), "No L1s found in Avalanche network")

	// Deploy an AvalancheValidatorSetRegistry contract on the Ethereum network
	validatorSetRegistryContractAddress, validatorSetRegistry := utils.DeployAvalancheValidatorSetRegistry(
		ctx,
		ethereumFundedKey,
		ethereumNetwork.ChainID,
		ethereumNetwork.Client,
		avalancheNetwork.GetNetworkID(),
		l1Info.BlockchainID,
	)
	log.Info("Deployed AvalancheValidatorSetRegistry contract on the Ethereum network", "address", validatorSetRegistryContractAddress)

	// Deploy a test cross chain messenger to the Avalanche subnet
	fundedAvaAddress, fundedAvaKey := avalancheNetwork.GetFundedAccountInfo()
	_, testMessengerL1 := utils.DeployTestMessenger(
		ctx,
		fundedAvaKey,
		fundedAvaAddress,
		teleporter.TeleporterRegistryAddress(l1Info),
		l1Info,
	)

	signatureAggregator := utils.NewSignatureAggregator(
		primaryNetworkInfo.NodeURIs[0],
		[]ids.ID{
			l1Info.SubnetID,
		},
	)
	defer signatureAggregator.Shutdown()

	// Register the Avalanche L1 network's validator set in the registry via an ICM message from the P-Chain, self-signed by the validator set of the Avalanche network
	pChainClient := platformvm.NewClient(primaryNetworkInfo.NodeURIs[0])
	utils.RegisterAvalancheValidatorSet(
		ctx,
		ethereumFundedKey,
		ethereumNetwork.ChainID,
		ethereumNetwork.Client,
		validatorSetRegistry,
		avalancheNetwork.GetNetworkID(),
		pChainInfo.BlockchainID,
		l1Info.SubnetID,
		l1Info.BlockchainID,
		pChainClient,
		signatureAggregator,
	)
	log.Info("registered warp with source id", "id", l1Info.EVMChainID)
	// Register the EthWarp contract and register the above ValidatorSetRegistry with it
	ethWarpAddress, ethWarpContract := utils.DeployEthWarp(
		ctx,
		ethereumFundedKey,
		ethereumNetwork.ChainID,
		ethereumNetwork.Client,
		l1Info.BlockchainID,
	)
	log.Info("Deployed EthWarp contract on the Ethereum network", "address", ethWarpAddress)

	utils.RegisterValSetRegistryWithEthWarp(
		ctx,
		ethereumFundedKey,
		ethereumNetwork.ChainID,
		ethereumNetwork.Client,
		ethWarpContract,
		l1Info.BlockchainID,
		validatorSetRegistryContractAddress,
	)

	// Deploy Teleporter contract to Ethereum using the adapted Warp contract
	ethTeleporterAddress, ethTeleporter := utils.DeployTeleporterMessengerToEthereum(
		ctx,
		ethereumFundedKey,
		ethereumNetwork.ChainID,
		ethereumNetwork.Client,
		ethWarpAddress,
	)
	log.Info("Deployed Teleporter to Ethereum network", "address", ethTeleporterAddress)

	// Send and verify an ICM message from the Avalanche L1 network to the Ethereum network.
	teleporter.SendExampleInterChainMessageAndVerify(
		ctx,
		ethereumNetwork.ChainID,
		ethereumNetwork.Client,
		avalancheNetwork.GetNetworkID(),
		pChainInfo.BlockchainID,
		l1Info,
		testMessengerL1,
		ethTeleporterAddress,
		ethTeleporter,
		fundedAvaKey,
		ethereumFundedKey,
		"Hello Ethereum",
		signatureAggregator,
		false,
	)

	log.Info("Message successfully sent to Ethereum")

	// TODO:
	// 1. Add a new validator to the L1, and update the validator set on the Ethereum network.
	// 2. Send and verify another message, using the updated validator set.
}
