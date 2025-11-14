package ethereum_icm_verification_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	ethereumICMVerification "github.com/ava-labs/icm-contracts/tests/flows/ethereum-icm-verification"
	"github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile       = "./avalanche/out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisTemplateFile      = "./tests/utils/warp-genesis-template.json"
	ethereumICMVerificationLabel = "ethereum-icm-verification"
)

var (
	LocalAvalancheNetworkInstance *network.LocalAvalancheNetwork
	LocalEthereumNetworkInstance  *network.LocalEthereumNetwork
	TeleporterInfo                utils.TeleporterTestInfo
	e2eFlags                      *e2e.FlagVars
)

func TestEthereumICMVerification(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	e2eFlags = e2e.RegisterFlags()
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Ethereum ICM Verification e2e test")
}

var _ = ginkgo.BeforeSuite(func() {
	// Configure logging for tests
	utils.ConfigureDefaultLoggingForTests()

	// Generate the Teleporter deployment values
	teleporterDeployerTransaction,
		teleporterDeployedBytecode,
		teleporterDeployerAddress,
		teleporterContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		teleporterByteCodeFile,
		false,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
	)
	Expect(err).Should(BeNil())

	// Create the local network instances
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	LocalAvalancheNetworkInstance = network.NewLocalAvalancheNetwork(
		ctx,
		"ethereum-icm-verification-test-local-network",
		warpGenesisTemplateFile,
		[]network.L1Spec{
			{
				Name:                       "L1",
				EVMChainID:                 12345,
				TeleporterContractAddress:  teleporterContractAddress,
				TeleporterDeployedBytecode: teleporterDeployedBytecode,
				TeleporterDeployerAddress:  teleporterDeployerAddress,
				NodeCount:                  2,
			},
		},
		2,
		2,
		e2eFlags,
	)
	TeleporterInfo = utils.NewTeleporterTestInfo(LocalAvalancheNetworkInstance.GetAllL1Infos())
	log.Info("Started local Avalanche network", "networkID", LocalAvalancheNetworkInstance.NetworkID)

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the L1 chains.
	_, fundedKey := LocalAvalancheNetworkInstance.GetFundedAccountInfo()
	TeleporterInfo.DeployTeleporterMessenger(
		ctx,
		LocalAvalancheNetworkInstance.GetPrimaryNetworkInfo(),
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)
	for _, l1 := range LocalAvalancheNetworkInstance.GetAllL1Infos() {
		TeleporterInfo.SetTeleporter(teleporterContractAddress, l1)
		TeleporterInfo.Initialize(l1, fundedKey, common.HexToAddress("0x0200000000000000000000000000000000000005"))
		TeleporterInfo.InitializeBlockchainID(l1, fundedKey)
		TeleporterInfo.DeployTeleporterRegistry(l1, fundedKey)
	}

	LocalEthereumNetworkInstance = network.NewLocalEthereumNetwork(ctx)
	log.Info("Started local Ethereum network", "chainID", LocalEthereumNetworkInstance.ChainID)

	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	LocalAvalancheNetworkInstance.TearDownNetwork()
	LocalAvalancheNetworkInstance = nil
	LocalEthereumNetworkInstance.TearDownNetwork()
	LocalEthereumNetworkInstance = nil
})

var _ = ginkgo.Describe("[Ethereum ICM verification integration tests]", func() {
	// Ethereum ICM verification tests
	ginkgo.It("Deploy and test Ethereum ICM verification",
		ginkgo.Label(ethereumICMVerificationLabel),
		func() {
			ethereumICMVerification.EthereumICMVerification(LocalAvalancheNetworkInstance, LocalEthereumNetworkInstance, TeleporterInfo)
		})
})
