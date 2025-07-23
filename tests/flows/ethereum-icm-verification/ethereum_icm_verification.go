package ethereum_icm_verification

import (
	"context"

	"github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/log"
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
	validatorSetRegistryContractAddress, _ := utils.DeployAvalancheValidatorSetRegistry(
		ctx,
		ethereumFundedKey,
		ethereumNetwork.ChainID,
		ethereumNetwork.Client,
		avalancheNetwork.NetworkID,
	)

	log.Info("Deployed AvalancheValidatorSetRegistry contract on the Ethereum network", "address", validatorSetRegistryContractAddress)
}
