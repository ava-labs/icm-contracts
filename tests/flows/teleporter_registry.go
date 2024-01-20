package flows

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	runner_sdk "github.com/ava-labs/avalanche-network-runner/client"
	"github.com/ava-labs/avalanchego/utils/constants"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleporterregistry "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/upgrades/TeleporterRegistry"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile = "./contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json"
	// TODO: when avalanche-network-runner depedency is updated, use constants.DefaultLocalNetworkID
	defaultLocalNetworkID uint32 = 1337
)

func TeleporterRegistry(network interfaces.LocalNetwork) {
	// Deploy dapp on both chains that use Teleporter Registry
	// Deploy version 2 of Teleporter to both chains
	// Construct AddProtocolVersion txs for both chains
	// Send tx for one of the two chains, verify events emitted
	// Send a message from chain with old version to chain with new version, verify message is received
	// Update minTeleporterVersion on chain with new version
	// Send same message from chain with old version, verify message is not received
	// Send a message from chain with new version to chain with old version, verify message is not received
	// Send tx with new protocol version to other chain, verify events emitted
	// Retry the previously failed message execution, verify message is now able to be delivered to dapp

	subnetAInfo, subnetBInfo := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example cross chain messenger to both chains
	exampleMessengerContractA, exampleMessengerA := utils.DeployExampleCrossChainMessenger(ctx, fundedKey, subnetAInfo)
	exampleMessengerContractB, exampleMessengerB := utils.DeployExampleCrossChainMessenger(
		ctx, fundedKey, subnetBInfo,
	)

	// Deploy the new version of Teleporter to both chains
	newTeleporterAddress := deployNewTeleporterVersion(ctx, network, fundedKey)

	// Create chain config file with off chain message for each chain
	offchainMessageA, warpEnabledChainConfigA := initChainConfig(subnetAInfo, newTeleporterAddress)
	offchainMessageB, warpEnabledChainConfigB := initChainConfig(subnetBInfo, newTeleporterAddress)

	// Create chain config with off chain messages
	chainConfigs := make(map[string]string)
	setChainConfig(chainConfigs, subnetAInfo, warpEnabledChainConfigA)
	setChainConfig(chainConfigs, subnetBInfo, warpEnabledChainConfigB)

	// Restart nodes with new chain config
	nodeNames := network.GetSubnetNodeNames()
	network.RestartNodes(ctx, nodeNames, runner_sdk.WithChainConfigs(chainConfigs))

	// Call addProtocolVersion on subnetB to register the new Teleporter version
	utils.AddProtocolVersionAndWaitForAcceptance(
		ctx,
		network,
		subnetBInfo,
		subnetAInfo,
		newTeleporterAddress,
		fundedKey,
		offchainMessageB)

	// Send a message using old Teleporter version to example messenger using new Teleporter version.
	// Message should be received successfully since we haven't updated mininum Teleporter version yet.
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		subnetAInfo,
		exampleMessengerA,
		subnetBInfo,
		exampleMessengerContractB,
		exampleMessengerB,
		fundedKey,
		true)

	// Update minimum Teleporter version on destination chain
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetBInfo.EVMChainID)
	Expect(err).Should(BeNil())

	latestVersion, err := subnetBInfo.TeleporterRegistry.LatestVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	minTeleporterVersion, err := exampleMessengerB.GetMinTeleporterVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tx, err := exampleMessengerB.UpdateMinTeleporterVersion(opts, latestVersion)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, subnetBInfo, tx)

	// Verify that minTeleporterVersion updated
	minTeleporterVersionUpdatedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		exampleMessengerB.ParseMinTeleporterVersionUpdated)
	Expect(err).Should(BeNil())
	Expect(minTeleporterVersionUpdatedEvent.OldMinTeleporterVersion.Cmp(minTeleporterVersion)).Should(Equal(0))
	Expect(minTeleporterVersionUpdatedEvent.NewMinTeleporterVersion.Cmp(latestVersion)).Should(Equal(0))

	// Send a message using old Teleporter version to example messenger with updated minimum Teleporter version.
	// Message should fail since we updated minimum Teleporter version.
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		subnetAInfo,
		exampleMessengerA,
		subnetBInfo,
		exampleMessengerContractB,
		exampleMessengerB,
		fundedKey,
		false)

	// Update the subnets to use new Teleporter messengers
	network.SetTeleporterContractAddress(newTeleporterAddress)
	subnetAInfo, subnetBInfo = utils.GetTwoSubnets(network)
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		subnetBInfo,
		exampleMessengerB,
		subnetAInfo,
		exampleMessengerContractA,
		exampleMessengerA,
		fundedKey,
		false)

	// Call addProtocolVersion on subnetA to register the new Teleporter version
	utils.AddProtocolVersionAndWaitForAcceptance(
		ctx,
		network,
		subnetAInfo,
		subnetBInfo,
		newTeleporterAddress,
		fundedKey,
		offchainMessageA)

	// Send a message from A->B, which previously failed, but now using the new Teleporter version.
	// Teleporter versions should match, so message should be received successfully.
	utils.SendExampleCrossChainMessageAndVerify(ctx,
		network,
		subnetBInfo,
		exampleMessengerB,
		subnetAInfo,
		exampleMessengerContractA,
		exampleMessengerA,
		fundedKey,
		true)
}

func initChainConfig(
	subnet interfaces.SubnetTestInfo,
	teleporterAddresses common.Address,
) (*avalancheWarp.UnsignedMessage, string) {
	warpApiEnabled := true
	unsignedMessage := createOffChainRegistryMessage(subnet, teleporterregistry.ProtocolRegistryEntry{
		Version:         big.NewInt(2),
		ProtocolAddress: teleporterAddresses,
	})
	offChainMessage := hexutil.Encode(unsignedMessage.Bytes())
	log.Info("Adding off-chain message to Warp chain config",
		"messageID", unsignedMessage.ID(),
		"blockchainID", subnet.BlockchainID.String())

	return unsignedMessage, fmt.Sprintf(`{
    "warp-api-enabled": %t, 
    "warp-off-chain-messages": ["%s"],
	"log-level": "debug",
    "eth-apis":["eth","eth-filter","net","admin","web3",
                "internal-eth","internal-blockchain","internal-transaction",
                "internal-debug","internal-account","internal-personal",
                "debug","debug-tracer","debug-file-tracer","debug-handler"]
	}`, warpApiEnabled, offChainMessage)
}

func createOffChainRegistryMessage(
	subnet interfaces.SubnetTestInfo,
	entry teleporterregistry.ProtocolRegistryEntry,
) *avalancheWarp.UnsignedMessage {
	sourceAddress := []byte{}
	payloadBytes, err := teleporterregistry.PackTeleporterRegistryWarpPayload(entry, subnet.TeleporterRegistryAddress)
	Expect(err).Should(BeNil())

	addressedPayload, err := payload.NewAddressedCall(sourceAddress, payloadBytes)
	Expect(err).Should(BeNil())

	unsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		defaultLocalNetworkID,
		subnet.BlockchainID,
		addressedPayload.Bytes())
	Expect(err).Should(BeNil())

	return unsignedMessage
}

func deployNewTeleporterVersion(
	ctx context.Context,
	network interfaces.LocalNetwork,
	fundedKey *ecdsa.PrivateKey,
) common.Address {
	deploymentUtils.AlterNicksTransaction()
	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err :=
		deploymentUtils.ConstructKeylessTransaction(teleporterByteCodeFile, false)
	Expect(err).Should(BeNil())

	network.DeployTeleporterContracts(
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
		false)
	return teleporterContractAddress
}

func setChainConfig(customChainConfigs map[string]string, subnet interfaces.SubnetTestInfo, chainConfig string) {
	if subnet.SubnetID == constants.PrimaryNetworkID {
		customChainConfigs[utils.CChainPathSpecifier] = chainConfig
	} else {
		customChainConfigs[subnet.BlockchainID.String()] = chainConfig
	}
}
