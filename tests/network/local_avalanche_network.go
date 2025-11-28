package network

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/fs"
	goLog "log"
	"os"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	warpMessage "github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"
	ownableupgradeable "github.com/ava-labs/icm-contracts/abi-bindings/go/OwnableUpgradeable"
	proxyadmin "github.com/ava-labs/icm-contracts/abi-bindings/go/ProxyAdmin"
	validatormanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/ValidatorManager"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetEvmTestUtils "github.com/ava-labs/subnet-evm/tests/utils"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/log"
	. "github.com/onsi/gomega"
)

var _ LocalNetwork = &LocalAvalancheNetwork{}

type ProxyAddress struct {
	common.Address
	*proxyadmin.ProxyAdmin
}

// Implements Network, pointing to the network setup in local_network_setup.go
type LocalAvalancheNetwork struct {
	*tmpnet.Network

	extraNodes                      []*tmpnet.Node // to add as more L1 validators in the tests
	primaryNetworkValidators        []*tmpnet.Node
	globalFundedKey                 *secp256k1.PrivateKey
	validatorManagers               map[ids.ID]ProxyAddress
	validatorManagerSpecializations map[ids.ID]ProxyAddress
	logger                          logging.Logger
	deployedL1Specs                 map[string]L1Spec
	graniteEpochDuration            time.Duration
}

const (
	fundedKeyStr         = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
	timeout              = 120 * time.Second
	defaultEpochDuration = 4 * time.Second
)

type L1Spec struct {
	Name       string
	EVMChainID uint64
	NodeCount  int

	// Optional fields
	TeleporterContractAddress    common.Address
	TeleporterDeployedBytecode   string
	TeleporterDeployerAddress    common.Address
	RequirePrimaryNetworkSigners bool
}

func newTmpnetNetwork(
	name string,
	globalFundedKey *secp256k1.PrivateKey,
	warpGenesisTemplateFile string,
	numPrimaryNetworkValidators int,
	l1Specs []L1Spec,
	flagVars *e2e.FlagVars,
) *tmpnet.Network {
	var l1s []*tmpnet.Subnet

	bootstrapNodes := subnetEvmTestUtils.NewTmpnetNodes(numPrimaryNetworkValidators)
	for i, l1Spec := range l1Specs {
		// Create a single bootstrap node. This will be removed from the L1 validator set after it is converted,
		// but will remain a primary network validator
		initialL1Bootstrapper := bootstrapNodes[i] // One bootstrap node per L1

		l1 := subnetEvmTestUtils.NewTmpnetSubnet(
			l1Spec.Name,
			utils.InstantiateGenesisTemplate(
				warpGenesisTemplateFile,
				l1Spec.EVMChainID,
				l1Spec.TeleporterContractAddress,
				l1Spec.TeleporterDeployedBytecode,
				l1Spec.TeleporterDeployerAddress,
				l1Spec.RequirePrimaryNetworkSigners,
			),
			utils.WarpEnabledChainConfig,
			initialL1Bootstrapper,
		)
		l1.OwningKey = globalFundedKey
		l1s = append(l1s, l1)
	}

	// Create new network
	network := subnetEvmTestUtils.NewTmpnetNetwork(
		name,
		bootstrapNodes,
		tmpnet.FlagsMap{},
		l1s...,
	)

	Expect(network).ShouldNot(BeNil())

	// Specify only a subset of the nodes to be bootstrapped
	keysToFund := []*secp256k1.PrivateKey{
		genesis.VMRQKey,
		genesis.EWOQKey,
		tmpnet.HardhatKey,
	}
	keysToFund = append(keysToFund, network.PreFundedKeys...)
	genesis, err := tmpnet.NewTestGenesis(88888, bootstrapNodes, keysToFund)
	Expect(err).Should(BeNil())
	network.Genesis = genesis
	network.PreFundedKeys = keysToFund

	runtimeCfg, err := flagVars.NodeRuntimeConfig()
	Expect(err).Should(BeNil())
	runtimeCfg.Process.ReuseDynamicPorts = true
	network.DefaultRuntimeConfig = *runtimeCfg

	return network
}

func NewLocalAvalancheNetwork(
	ctx context.Context,
	name string,
	warpGenesisTemplateFile string,
	l1Specs []L1Spec,
	numPrimaryNetworkValidators int,
	extraNodeCount int, // for use by tests, e.g. to add new L1 validators
	flagVars *e2e.FlagVars,
) *LocalAvalancheNetwork {
	// There must be at least one primary network validator per L1
	Expect(numPrimaryNetworkValidators).Should(BeNumerically(">=", len(l1Specs)))

	// Create extra nodes to be used to add more validators later
	extraNodes := subnetEvmTestUtils.NewTmpnetNodes(extraNodeCount)

	for _, l1Spec := range l1Specs {
		initialVdrNodes := subnetEvmTestUtils.NewTmpnetNodes(l1Spec.NodeCount)
		extraNodes = append(extraNodes, initialVdrNodes...)
	}

	fundedKey, err := hex.DecodeString(fundedKeyStr)
	Expect(err).Should(BeNil())
	globalFundedKey, err := secp256k1.ToPrivateKey(fundedKey)
	Expect(err).Should(BeNil())

	globalFundedECDSAKey := globalFundedKey.ToECDSA()
	Expect(err).Should(BeNil())

	deployedL1Specs := make(map[string]L1Spec)
	for _, l1Spec := range l1Specs {
		deployedL1Specs[l1Spec.Name] = l1Spec
	}

	isReuseNetwork := flagVars != nil && flagVars.NetworkDir() != ""

	var network *tmpnet.Network
	// All nodes are specified as bootstrap validators
	var primaryNetworkValidators []*tmpnet.Node
	if isReuseNetwork {
		// Load existing network and restart nodes
		network, err = tmpnet.ReadNetwork(context.Background(), logging.NoLog{}, flagVars.NetworkDir())
		Expect(err).Should(BeNil())
		Expect(network).ShouldNot(BeNil())

		extraNodes = make([]*tmpnet.Node, 0)
		for _, node := range network.Nodes {
			err := node.Restart(ctx)
			Expect(err).Should(BeNil(), "Failed to restart node %s: %v", node.NodeID, err)

			if node.Flags[config.PartialSyncPrimaryNetworkKey] == "true" {
				extraNodes = append(extraNodes, node)
			} else {
				primaryNetworkValidators = append(primaryNetworkValidators, node)
			}
		}

		err := tmpnet.WaitForHealthyNodes(ctx, logging.NoLog{}, network.Nodes)
		Expect(err).Should(BeNil())
	} else {
		network = newTmpnetNetwork(
			name,
			globalFundedKey,
			warpGenesisTemplateFile,
			numPrimaryNetworkValidators,
			l1Specs,
			flagVars,
		)

		primaryNetworkValidators = append(primaryNetworkValidators, network.Nodes...)
	}

	upgrades := upgrade.Default
	graniteEpochDuration := defaultEpochDuration
	if envDuration := os.Getenv("GRANITE_EPOCH_DURATION"); envDuration != "" {
		if parsed, err := time.ParseDuration(envDuration); err == nil {
			graniteEpochDuration = parsed
			goLog.Printf("Using Granite epoch duration from environment: %v", graniteEpochDuration)
		} else {
			goLog.Printf("Invalid GRANITE_EPOCH_DURATION '%s', using default: %v", envDuration, graniteEpochDuration)
		}
	}
	upgrades.GraniteEpochDuration = graniteEpochDuration

	upgradeJSON, err := json.Marshal(upgrades)
	Expect(err).Should(BeNil())

	upgradeBase64 := base64.StdEncoding.EncodeToString(upgradeJSON)

	defaultFlags := tmpnet.FlagsMap{
		config.UpgradeFileContentKey: upgradeBase64,
	}
	defaultFlags.SetDefaults(tmpnet.DefaultE2EFlags())
	network.DefaultFlags = defaultFlags

	tc := e2e.NewTestContext()
	env := e2e.NewTestEnvironment(tc, flagVars, network)
	Expect(env).ShouldNot(BeNil())

	ctx, cancelBootstrap := context.WithCancel(ctx)
	defer cancelBootstrap()

	logger := logging.NewLogger("tmpnet")
	goLog.Println("Network bootstrapped")

	// Issue transactions to activate the proposerVM fork on the chains
	if !isReuseNetwork {
		for _, l1 := range network.Subnets {
			utils.SetupProposerVM(ctx, globalFundedECDSAKey, network, l1.SubnetID)
		}
	}

	localNetwork := &LocalAvalancheNetwork{
		Network:                         network,
		extraNodes:                      extraNodes,
		globalFundedKey:                 globalFundedKey,
		primaryNetworkValidators:        primaryNetworkValidators,
		validatorManagers:               make(map[ids.ID]ProxyAddress),
		validatorManagerSpecializations: make(map[ids.ID]ProxyAddress),
		logger:                          logger,
		deployedL1Specs:                 deployedL1Specs,
		graniteEpochDuration:            upgrades.GraniteEpochDuration,
	}

	return localNetwork
}

func (n *LocalAvalancheNetwork) ConvertSubnet(
	ctx context.Context,
	l1 interfaces.L1TestInfo,
	managerType utils.ValidatorManagerConcreteType,
	weights []uint64,
	balances []uint64,
	senderKey *ecdsa.PrivateKey,
	proxy bool,
) ([]utils.Node, []ids.ID) {
	Expect(len(weights)).Should(Equal(len(balances)))
	goLog.Println("Converting l1", l1.SubnetID)
	cChainInfo := n.GetPrimaryNetworkInfo()
	pClient := platformvm.NewClient(cChainInfo.NodeURIs[0])
	currentValidators, err := pClient.GetCurrentValidators(ctx, l1.SubnetID, nil)
	Expect(err).Should(BeNil())

	vdrManagerAddress, vdrManagerProxyAdmin := utils.DeployValidatorManager(
		ctx,
		senderKey,
		l1,
		proxy,
	)

	validatorManager, err := validatormanager.NewValidatorManager(vdrManagerAddress, l1.RPCClient)
	Expect(err).Should(BeNil())

	sender := utils.PrivateKeyToAddress(senderKey)

	utils.InitializeValidatorManager(
		ctx,
		senderKey,
		l1,
		validatorManager,
		sender,
	)

	n.validatorManagers[l1.SubnetID] = ProxyAddress{
		Address:    vdrManagerAddress,
		ProxyAdmin: vdrManagerProxyAdmin,
	}

	specializationAddress, specializationProxyAdmin := utils.DeployAndInitializeValidatorManagerSpecialization(
		ctx,
		senderKey,
		l1,
		vdrManagerAddress,
		managerType,
		proxy,
	)

	ownable, err := ownableupgradeable.NewOwnableUpgradeable(vdrManagerAddress, l1.RPCClient)
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := ownable.TransferOwnership(opts, specializationAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(context.Background(), l1, tx.Hash())

	n.validatorManagerSpecializations[l1.SubnetID] = ProxyAddress{
		Address:    specializationAddress,
		ProxyAdmin: specializationProxyAdmin,
	}

	tmpnetNodes := n.GetExtraNodes(len(weights))
	sort.Slice(tmpnetNodes, func(i, j int) bool {
		return string(tmpnetNodes[i].NodeID.Bytes()) < string(tmpnetNodes[j].NodeID.Bytes())
	})

	var nodes []utils.Node
	// Construct the converted l1 info
	destAddr, err := address.ParseToID(utils.DefaultPChainAddress)
	Expect(err).Should(BeNil())
	vdrs := make([]*txs.ConvertSubnetToL1Validator, len(tmpnetNodes))
	for i, node := range tmpnetNodes {
		signer, err := node.GetProofOfPossession()
		Expect(err).Should(BeNil())
		nodes = append(nodes, utils.Node{
			NodeID:  node.NodeID,
			NodePoP: signer,
			Weight:  weights[i],
		})
		vdrs[i] = &txs.ConvertSubnetToL1Validator{
			NodeID:  node.NodeID.Bytes(),
			Weight:  weights[i],
			Balance: balances[i],
			Signer:  *signer,
			RemainingBalanceOwner: warpMessage.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
			DeactivationOwner: warpMessage.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
		}
	}
	pChainWallet := n.GetPChainWallet()
	_, err = pChainWallet.IssueConvertSubnetToL1Tx(
		l1.SubnetID,
		l1.BlockchainID,
		vdrManagerAddress[:],
		vdrs,
	)
	Expect(err).Should(BeNil())

	l1 = n.AddSubnetValidators(tmpnetNodes, l1, true)

	// Wait for P-Chain to finalize and propagate transactions
	utils.AdvanceProposerVM(ctx, l1, senderKey, 5)
	goLog.Println("Waiting for Granite epoch to complete for ", n.graniteEpochDuration)
	time.Sleep(n.graniteEpochDuration)

	aggregator := n.GetSignatureAggregator()
	defer aggregator.Shutdown()

	validationIDs := utils.InitializeValidatorSet(
		ctx,
		senderKey,
		l1,
		utils.GetPChainInfo(cChainInfo),
		vdrManagerAddress,
		n.GetNetworkID(),
		aggregator,
		nodes,
	)

	// Remove the bootstrap nodes as l1 validators
	for _, vdr := range currentValidators {
		_, err := pChainWallet.IssueRemoveSubnetValidatorTx(vdr.NodeID, l1.SubnetID)
		Expect(err).Should(BeNil())
		for _, node := range n.Network.Nodes {
			if node.NodeID == vdr.NodeID {
				Expect(n.Network.DefaultRuntimeConfig).ShouldNot(BeNil())
				Expect(n.Network.DefaultRuntimeConfig.Process.ReuseDynamicPorts).Should(BeTrue())
				node.RuntimeConfig = &n.Network.DefaultRuntimeConfig
				goLog.Println("Restarting bootstrap node", node.NodeID)
				err = node.Restart(ctx)
				Expect(err).Should(BeNil())
			}
		}
	}
	utils.PChainProposerVMWorkaround(pChainWallet)
	utils.IssueTxsToAdvanceChain(ctx, l1.EVMChainID, senderKey, l1.RPCClient, 5)

	return nodes, validationIDs
}

func (n *LocalAvalancheNetwork) AddSubnetValidators(
	nodes []*tmpnet.Node,
	l1 interfaces.L1TestInfo,
	partialSync bool,
) interfaces.L1TestInfo {
	// Modify each node's config to track the l1
	for _, node := range nodes {
		goLog.Printf("Adding node %s @ %s to l1 %s", node.NodeID, node.URI, l1.SubnetID)
		existingTrackedSubnets := node.Flags[config.TrackSubnetsKey]
		if existingTrackedSubnets == l1.SubnetID.String() {
			goLog.Printf("Node %s @ %s already tracking l1 %s\n", node.NodeID, node.URI, l1.SubnetID)
			continue
		}
		node.Flags[config.TrackSubnetsKey] = l1.SubnetID.String()

		if partialSync {
			node.Flags[config.PartialSyncPrimaryNetworkKey] = "true"
		}

		// Add the node to the network
		n.Network.Nodes = append(n.Network.Nodes, node)
	}
	err := n.Network.StartNodes(context.Background(), n.logger, nodes...)
	Expect(err).Should(BeNil())

	// Update the tmpnet Subnet struct
	for _, tmpnetSubnet := range n.Network.Subnets {
		if tmpnetSubnet.SubnetID == l1.SubnetID {
			for _, tmpnetNode := range nodes {
				tmpnetSubnet.ValidatorIDs = append(tmpnetSubnet.ValidatorIDs, tmpnetNode.NodeID)
			}
		}
	}

	// Refresh the l1 info after restarting the nodes
	return n.GetL1Info(l1.SubnetID)
}

func (n *LocalAvalancheNetwork) GetValidatorManager(subnetID ids.ID) (ProxyAddress, ProxyAddress) {
	return n.validatorManagers[subnetID], n.validatorManagerSpecializations[subnetID]
}

func (n *LocalAvalancheNetwork) GetSignatureAggregator() *utils.SignatureAggregator {
	var subnetIDs []ids.ID
	for _, l1 := range n.GetL1Infos() {
		subnetIDs = append(subnetIDs, l1.SubnetID)
	}
	return utils.NewSignatureAggregator(
		n.GetPrimaryNetworkInfo().NodeURIs[0],
		subnetIDs,
	)
}

func (n *LocalAvalancheNetwork) GetExtraNodes(count int) []*tmpnet.Node {
	Expect(len(n.extraNodes) >= count).Should(
		BeTrue(),
		"not enough extra nodes to use",
	)
	nodes := n.extraNodes[0:count]
	n.extraNodes = n.extraNodes[count:]
	return nodes
}

func (n *LocalAvalancheNetwork) GetPrimaryNetworkValidators() []*tmpnet.Node {
	return n.primaryNetworkValidators
}

func (n *LocalAvalancheNetwork) GetPrimaryNetworkInfo() interfaces.L1TestInfo {
	var nodeURIs []string
	for _, node := range n.primaryNetworkValidators {
		nodeURIs = append(nodeURIs, node.URI)
	}
	infoClient := info.NewClient(nodeURIs[0])
	cChainBlockchainID, err := infoClient.GetBlockchainID(context.Background(), "C")
	Expect(err).Should(BeNil())

	wsClient, err := ethclient.Dial(utils.HttpToWebsocketURI(nodeURIs[0], cChainBlockchainID.String()))
	Expect(err).Should(BeNil())

	rpcClient, err := ethclient.Dial(utils.HttpToRPCURI(nodeURIs[0], cChainBlockchainID.String()))
	Expect(err).Should(BeNil())

	evmChainID, err := rpcClient.ChainID(context.Background())
	Expect(err).Should(BeNil())
	return interfaces.L1TestInfo{
		SubnetID:                     ids.Empty,
		BlockchainID:                 cChainBlockchainID,
		NodeURIs:                     nodeURIs,
		WSClient:                     wsClient,
		RPCClient:                    rpcClient,
		EVMChainID:                   evmChainID,
		RequirePrimaryNetworkSigners: false,
	}
}

func (n *LocalAvalancheNetwork) GetL1Info(subnetID ids.ID) interfaces.L1TestInfo {
	for _, l1 := range n.Network.Subnets {
		if l1.SubnetID == subnetID {
			var nodeURIs []string
			for _, nodeID := range l1.ValidatorIDs {
				node, err := n.Network.GetNode(nodeID)
				Expect(err).Should(BeNil())

				nodeURIs = append(nodeURIs, node.URI)
			}
			blockchainID := l1.Chains[0].ChainID
			wsClient, err := ethclient.Dial(utils.HttpToWebsocketURI(nodeURIs[0], blockchainID.String()))
			Expect(err).Should(BeNil())

			rpcClient, err := ethclient.Dial(utils.HttpToRPCURI(nodeURIs[0], blockchainID.String()))
			Expect(err).Should(BeNil())
			evmChainID, err := rpcClient.ChainID(context.Background())
			Expect(err).Should(BeNil())
			spec, ok := n.deployedL1Specs[l1.Name]
			Expect(ok).Should(BeTrue())
			return interfaces.L1TestInfo{
				SubnetID:                     subnetID,
				BlockchainID:                 blockchainID,
				NodeURIs:                     nodeURIs,
				WSClient:                     wsClient,
				RPCClient:                    rpcClient,
				EVMChainID:                   evmChainID,
				RequirePrimaryNetworkSigners: spec.RequirePrimaryNetworkSigners,
			}
		}
	}
	return interfaces.L1TestInfo{}
}

// Returns all l1 info sorted in lexicographic order of L1Name.
func (n *LocalAvalancheNetwork) GetL1Infos() []interfaces.L1TestInfo {
	l1s := make([]interfaces.L1TestInfo, len(n.Network.Subnets))
	for i, l1 := range n.Network.Subnets {
		var nodeURIs []string
		for _, nodeID := range l1.ValidatorIDs {
			node, err := n.Network.GetNode(nodeID)
			Expect(err).Should(BeNil())

			nodeURIs = append(nodeURIs, node.URI)
		}
		blockchainID := l1.Chains[0].ChainID
		wsClient, err := ethclient.Dial(utils.HttpToWebsocketURI(nodeURIs[0], blockchainID.String()))
		Expect(err).Should(BeNil())

		rpcClient, err := ethclient.Dial(utils.HttpToRPCURI(nodeURIs[0], blockchainID.String()))
		Expect(err).Should(BeNil())
		evmChainID, err := rpcClient.ChainID(context.Background())
		Expect(err).Should(BeNil())
		spec, ok := n.deployedL1Specs[l1.Name]
		Expect(ok).Should(BeTrue())
		l1s[i] = interfaces.L1TestInfo{
			SubnetID:                     l1.SubnetID,
			BlockchainID:                 blockchainID,
			NodeURIs:                     nodeURIs,
			WSClient:                     wsClient,
			RPCClient:                    rpcClient,
			EVMChainID:                   evmChainID,
			RequirePrimaryNetworkSigners: spec.RequirePrimaryNetworkSigners,
		}
	}
	return l1s
}

// Returns L1 info for all L1s, including the primary network
func (n *LocalAvalancheNetwork) GetAllL1Infos() []interfaces.L1TestInfo {
	l1s := n.GetL1Infos()
	return append(l1s, n.GetPrimaryNetworkInfo())
}

func (n *LocalAvalancheNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	ecdsaKey := n.globalFundedKey.ToECDSA()
	fundedAddress := crypto.PubkeyToAddress(ecdsaKey.PublicKey)
	return fundedAddress, ecdsaKey
}

func (n *LocalAvalancheNetwork) TearDownNetwork() {
	log.Info("Tearing down local Avalanche network")
	Expect(n).ShouldNot(BeNil())
	Expect(n.Network).ShouldNot(BeNil())
	Expect(n.Network.Stop(context.Background())).Should(BeNil())
}

func (n *LocalAvalancheNetwork) SetChainConfigs(chainConfigs map[string]string) {
	for chainIDStr, chainConfig := range chainConfigs {
		var cfg tmpnet.ConfigMap
		err := json.Unmarshal([]byte(chainConfig), &cfg)
		if err != nil {
			log.Error(
				"failed to unmarshal chain config",
				"error", err,
				"chainConfig", chainConfig,
			)
		}
		if chainIDStr == utils.CChainPathSpecifier {
			n.Network.PrimarySubnetConfig = cfg
			n.Network.PrimaryChainConfigs[utils.CChainPathSpecifier] = cfg
			continue
		}

		for _, l1 := range n.Network.Subnets {
			for _, chain := range l1.Chains {
				if chain.ChainID.String() == chainIDStr {
					chain.Config = chainConfig
				}
			}
		}
	}
	err := n.Network.Write()
	if err != nil {
		log.Error("failed to write network", "error", err)
	}

	for _, l1 := range n.Network.Subnets {
		err := l1.Write(n.Network.GetSubnetDir())
		if err != nil {
			log.Error("failed to write L1s", "error", err)
		}
	}

	// Restart the network to apply the new chain configs
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60*len(n.Network.Nodes))*time.Second)
	defer cancel()
	err = n.Network.Restart(ctx)
	Expect(err).Should(BeNil())
}

func (n *LocalAvalancheNetwork) GetNetworkID() uint32 {
	return n.Network.Genesis.NetworkID
}

func (n *LocalAvalancheNetwork) Dir() string {
	return n.Network.Dir
}

func (n *LocalAvalancheNetwork) GetPChainWallet(validationIDs ...ids.ID) pwallet.Wallet {
	// Create the P-Chain wallet to issue transactions
	kc := secp256k1fx.NewKeychain(n.globalFundedKey)
	var subnetIDs []ids.ID
	for _, l1 := range n.GetL1Infos() {
		subnetIDs = append(subnetIDs, l1.SubnetID)
	}
	wallet, err := primary.MakeWallet(
		context.Background(),
		n.GetPrimaryNetworkInfo().NodeURIs[0],
		kc,
		kc,
		primary.WalletConfig{
			SubnetIDs:     subnetIDs,
			ValidationIDs: validationIDs,
		})
	Expect(err).Should(BeNil())
	return wallet.P()
}

func (n *LocalAvalancheNetwork) GetTwoL1s() (
	interfaces.L1TestInfo,
	interfaces.L1TestInfo,
) {
	l1s := n.GetL1Infos()
	Expect(len(l1s)).Should(BeNumerically(">=", 2))
	return l1s[0], l1s[1]
}

func (n *LocalAvalancheNetwork) SaveValidatorAddress(
	fileName string,
) {
	validatorAddresses := make(map[string]map[string]string)
	for _, subnet := range n.GetL1Infos() {
		validator, validatorSpec := n.GetValidatorManager(subnet.SubnetID)
		validatorAddresses[subnet.BlockchainID.Hex()] = make(map[string]string)
		validatorAddresses[subnet.BlockchainID.Hex()]["validator"] = validator.Address.Hex()
		validatorAddresses[subnet.BlockchainID.Hex()]["spec"] = validatorSpec.Address.Hex()
	}

	jsonData, err := json.Marshal(validatorAddresses)
	Expect(err).Should(BeNil())
	err = os.WriteFile(fileName, jsonData, fs.ModePerm)
	Expect(err).Should(BeNil())
}

func (n *LocalAvalancheNetwork) SetValidatorAddressFromFile(fileName string) {
	validatorAddresses := make(map[string]map[string]string)
	data, err := os.ReadFile(fileName)
	Expect(err).Should(BeNil())
	err = json.Unmarshal(data, &validatorAddresses)
	Expect(err).Should(BeNil())

	// Set the validator manager for each L1
	for _, subnet := range n.GetL1Infos() {
		validatorAddress := common.HexToAddress(validatorAddresses[subnet.BlockchainID.Hex()]["validator"])
		validatorSpecAddress := common.HexToAddress(validatorAddresses[subnet.BlockchainID.Hex()]["spec"])

		n.validatorManagers[subnet.SubnetID] = ProxyAddress{
			Address: validatorAddress,
		}
		n.validatorManagerSpecializations[subnet.SubnetID] = ProxyAddress{
			Address: validatorSpecAddress,
		}
	}
}
