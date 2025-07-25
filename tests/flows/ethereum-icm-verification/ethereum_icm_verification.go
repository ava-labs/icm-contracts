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

func EthereumICMVerification(avalancheNetwork *network.LocalAvalancheNetwork, ethereumNetwork *network.LocalEthereumNetwork) {
	// 1. Deploy an AvalancheValidatorSetRegistry contract on the Ethereum network
	// 2. Register the Avalanche L1 network's validator set in the registry via an ICM message from the P-Chain, self-signed by the validator set of the Avalanche network
	// 3. Send a transaction on the Avalanche L1 network to send an ICM message to the Ethereum network
	// 4. Get an aggregate signature of the ICM message using the ICM signature aggregator
	// 5. Broadcast the signed ICM message to the Ethereum network, and confirm successful inclusion and verification
	// 6. Add a new validator to the Avalanche L1 network, representing 40% of the new total weight
	// 7. Update the validator set reigstered on the Ethereum registry via a validator set update message signed by sufficient weight of the previous validator set
	// 8. Send and verify another ICM message from the Avalanche L1 network to the Ethereum network
	ctx := context.Background()
	_, ethereumFundedKey := ethereumNetwork.GetFundedAccountInfo()

	// Deploy an AvalancheValidatorSetRegistry contract on the Ethereum network
	validatorSetRegistryContractAddress, validatorSetRegistry := utils.DeployAvalancheValidatorSetRegistry(
		ctx,
		ethereumFundedKey,
		ethereumNetwork.ChainID,
		ethereumNetwork.Client,
		avalancheNetwork.GetNetworkID(),
	)
	log.Info("Deployed AvalancheValidatorSetRegistry contract on the Ethereum network", "address", validatorSetRegistryContractAddress)

	// Start a signature aggregator for the Avalanche network to be able to generate signed ICM messages.
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	Expect(len(primaryNetworkInfo.NodeURIs)).Should(BeNumerically(">", 0), "No NodeURIs found in primaryNetworkInfo")
	pChainInfo := utils.GetPChainInfo(primaryNetworkInfo)
	l1Infos := avalancheNetwork.GetL1Infos()
	l1Info := l1Infos[0]
	Expect(len(l1Infos)).Should(BeNumerically(">", 0), "No L1s found in Avalanche network")

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
}
