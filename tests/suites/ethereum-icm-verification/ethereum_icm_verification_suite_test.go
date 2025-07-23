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
	"github.com/ava-labs/libevm/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	warpGenesisTemplateFile      = "./tests/utils/warp-genesis-template.json"
	ethereumICMVerificationLabel = "ethereum-icm-verification"
)

var (
	LocalAvalancheNetworkInstance *network.LocalAvalancheNetwork
	LocalEthereumNetworkInstance  *network.LocalEthereumNetwork
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

	// Create the local network instances
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	LocalAvalancheNetworkInstance = network.NewLocalAvalancheNetwork(
		ctx,
		"ethereum-icm-verification-test-local-network",
		warpGenesisTemplateFile,
		[]network.L1Spec{
			{
				Name:       "L1",
				EVMChainID: 12345,
				NodeCount:  2,
			},
		},
		2,
		2,
		e2eFlags,
	)
	log.Info("Started local Avalanche network")

	LocalEthereumNetworkInstance = network.NewLocalEthereumNetwork(ctx)
	log.Info("Started local Ethereum network")

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
			ethereumICMVerification.EthereumICMVerification(LocalAvalancheNetworkInstance, LocalEthereumNetworkInstance)
		})
})
