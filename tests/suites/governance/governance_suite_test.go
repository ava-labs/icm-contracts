package governance_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	governanceFlows "github.com/ava-labs/icm-contracts/tests/flows/governance"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/libevm/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"
	validatorSetSigLabel    = "ValidatorSetSig"
)

var (
	LocalNetworkInstance *localnetwork.LocalNetwork
	e2eFlags             *e2e.FlagVars
)

func TestGovernance(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	e2eFlags = e2e.RegisterFlags()
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Governance e2e test")
}

// Define the before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
	// Create the local network instance
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	LocalNetworkInstance = localnetwork.NewLocalNetwork(
		ctx,
		"governance-test-local-network",
		warpGenesisTemplateFile,
		[]localnetwork.L1Spec{
			{
				Name:       "A",
				EVMChainID: 12345,
				NodeCount:  2,
			},
			{
				Name:       "B",
				EVMChainID: 54321,
				NodeCount:  2,
			},
		},
		2,
		2,
		e2eFlags,
	)
	log.Info("Started local network")
})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
})

var _ = ginkgo.Describe("[Governance integration tests]", func() {
	// Governance tests
	ginkgo.It("Deliver ValidatorSetSig signed message",
		ginkgo.Label(validatorSetSigLabel),
		func() {
			governanceFlows.ValidatorSetSig(LocalNetworkInstance)
		})
})
