// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/utils/units"
	teleporterFlows "github.com/ava-labs/icm-contracts/tests/flows/teleporter"
	registryFlows "github.com/ava-labs/icm-contracts/tests/flows/teleporter/registry"
	"github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile  = "./out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"

	teleporterMessengerLabel = "TeleporterMessenger"
	upgradabilityLabel       = "upgradability"
	utilsLabel               = "utils"
)

var (
	LocalNetworkInstance *network.LocalNetwork
	TeleporterInfo       utils.TeleporterTestInfo
	e2eFlags             *e2e.FlagVars
)

func TestTeleporter(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	e2eFlags = e2e.RegisterFlags()
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
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

	// Create the local network instance
	ctx, cancel := context.WithTimeout(context.Background(), 240*2*time.Second)
	defer cancel()

	LocalNetworkInstance = network.NewLocalNetwork(
		ctx,
		"teleporter-test-local-network",
		warpGenesisTemplateFile,
		[]network.L1Spec{
			{
				Name:                         "A",
				EVMChainID:                   12345,
				TeleporterContractAddress:    teleporterContractAddress,
				TeleporterDeployedBytecode:   teleporterDeployedBytecode,
				TeleporterDeployerAddress:    teleporterDeployerAddress,
				NodeCount:                    5,
				RequirePrimaryNetworkSigners: true,
			},
			{
				Name:                         "B",
				EVMChainID:                   54321,
				TeleporterContractAddress:    teleporterContractAddress,
				TeleporterDeployedBytecode:   teleporterDeployedBytecode,
				TeleporterDeployerAddress:    teleporterDeployerAddress,
				NodeCount:                    5,
				RequirePrimaryNetworkSigners: true,
			},
		},
		2,
		2,
		e2eFlags,
	)
	TeleporterInfo = utils.NewTeleporterTestInfo(LocalNetworkInstance.GetAllL1Infos())
	log.Info("Started local network")

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the l1 chains.
	_, fundedKey := LocalNetworkInstance.GetFundedAccountInfo()
	TeleporterInfo.DeployTeleporterMessenger(
		ctx,
		LocalNetworkInstance.GetPrimaryNetworkInfo(),
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)

	for _, l1 := range LocalNetworkInstance.GetAllL1Infos() {
		TeleporterInfo.SetTeleporter(teleporterContractAddress, l1)
		TeleporterInfo.InitializeBlockchainID(l1, fundedKey)
		TeleporterInfo.DeployTeleporterRegistry(l1, fundedKey)
	}

	balance := 100 * units.Avax
	for _, subnet := range LocalNetworkInstance.GetL1Infos() {
		// Choose weights such that we can test validator churn
		LocalNetworkInstance.ConvertSubnet(
			ctx,
			subnet,
			utils.PoAValidatorManager,
			[]uint64{units.Schmeckle, units.Schmeckle, units.Schmeckle, units.Schmeckle, units.Schmeckle},
			[]uint64{balance, balance, balance, balance, balance},
			fundedKey,
			false,
		)
	}

	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
	LocalNetworkInstance = nil
})

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Teleporter tests
	ginkgo.It("Send a message from L1 A to L1 B, and one from B to A",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.BasicSendReceive(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Deliver to the wrong chain",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.DeliverToWrongChain(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Deliver to non-existent contract",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.DeliverToNonExistentContract(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Retry successful execution",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.RetrySuccessfulExecution(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Unallowed relayer",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.UnallowedRelayer(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Relay message twice",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.RelayMessageTwice(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Add additional fee amount",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.AddFeeAmount(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Send specific receipts",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.SendSpecificReceipts(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Insufficient gas",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.InsufficientGas(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Resubmit altered message",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.ResubmitAlteredMessage(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Calculate Teleporter message IDs",
		ginkgo.Label(utilsLabel),
		func() {
			teleporterFlows.CalculateMessageID(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Relayer modifies message",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.RelayerModifiesMessage(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Validator churn",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.ValidatorChurn(LocalNetworkInstance, TeleporterInfo)
		})

	// Teleporter Registry tests
	ginkgo.It("Teleporter registry",
		ginkgo.Label(upgradabilityLabel),
		func() {
			registryFlows.TeleporterRegistry(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Check upgrade access",
		ginkgo.Label(upgradabilityLabel),
		func() {
			registryFlows.CheckUpgradeAccess(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Pause and Unpause Teleporter",
		ginkgo.Label(upgradabilityLabel),
		func() {
			registryFlows.PauseTeleporter(LocalNetworkInstance, TeleporterInfo)
		})
})
