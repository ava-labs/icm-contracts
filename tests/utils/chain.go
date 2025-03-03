package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	goLog "log"
	"math/big"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	nativeMinter "github.com/ava-labs/icm-contracts/abi-bindings/go/INativeMinter"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	gasUtils "github.com/ava-labs/icm-contracts/utils/gas-utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/eth/tracers"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetEvmInterfaces "github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/precompile/contracts/nativeminter"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	CChainPathSpecifier = "C"
)

var NativeTransferGas uint64 = 21_000

var WarpEnabledChainConfig = tmpnet.FlagsMap{
	"log-level":         "debug",
	"warp-api-enabled":  true,
	"local-txs-enabled": true,
	"eth-apis": []string{
		"eth",
		"eth-filter",
		"net",
		"admin",
		"web3",
		"internal-eth",
		"internal-blockchain",
		"internal-transaction",
		"internal-debug",
		"internal-account",
		"internal-personal",
		"debug",
		"debug-tracer",
		"debug-file-tracer",
		"debug-handler",
	},
}

type Node struct {
	NodeID  ids.NodeID
	NodePoP *signer.ProofOfPossession
	Weight  uint64
}

//
// URL utils
//

func HttpToWebsocketURI(uri string, blockchainID string) string {
	return fmt.Sprintf("ws://%s/ext/bc/%s/ws", strings.TrimPrefix(uri, "http://"), blockchainID)
}

func HttpToRPCURI(uri string, blockchainID string) string {
	return fmt.Sprintf("http://%s/ext/bc/%s/rpc", strings.TrimPrefix(uri, "http://"), blockchainID)
}

// Get the host and port from a URI. The URI should be in the format http://host:port or https://host:port
func GetURIHostAndPort(uri string) (string, uint32, error) {
	// At a minimum uri should have http:// of 7 characters
	Expect(len(uri)).Should(BeNumerically(">", 7))
	if uri[:7] == "http://" {
		uri = uri[7:]
	} else if uri[:8] == "https://" {
		uri = uri[8:]
	} else {
		return "", 0, fmt.Errorf("invalid uri: %s", uri)
	}

	// Split the uri into host and port
	hostAndPort := strings.Split(uri, ":")
	Expect(len(hostAndPort)).Should(Equal(2))

	// Parse the port
	port, err := strconv.ParseUint(hostAndPort[1], 10, 32)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse port: %w", err)
	}

	return hostAndPort[0], uint32(port), nil
}

//
// Transaction utils
//

func CreateNativeTransferTransaction(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Transaction {
	fromAddress := crypto.PubkeyToAddress(fromKey.PublicKey)
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, l1Info, fromAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   l1Info.EVMChainID,
		Nonce:     nonce,
		To:        &recipient,
		Gas:       NativeTransferGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     amount,
	})

	return SignTransaction(tx, fromKey, l1Info.EVMChainID)
}

func SendNativeTransfer(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Receipt {
	tx := CreateNativeTransferTransaction(ctx, l1Info, fromKey, recipient, amount)
	return SendTransactionAndWaitForSuccess(ctx, l1Info, tx)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals success.
func sendAndWaitForTransaction(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	tx *types.Transaction,
	success bool,
) *types.Receipt {
	err := l1Info.RPCClient.SendTransaction(ctx, tx)
	Expect(err).Should(BeNil())

	return waitForTransaction(ctx, l1Info, tx.Hash(), success)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals false.
func SendTransactionAndWaitForFailure(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	tx *types.Transaction,
) *types.Receipt {
	return sendAndWaitForTransaction(ctx, l1Info, tx, false)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals true.
func SendTransactionAndWaitForSuccess(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	tx *types.Transaction,
) *types.Receipt {
	return sendAndWaitForTransaction(ctx, l1Info, tx, true)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals true.
func WaitForTransactionSuccess(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	txHash common.Hash,
) *types.Receipt {
	return waitForTransaction(ctx, l1Info, txHash, true)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals false.
func WaitForTransactionFailure(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	txHash common.Hash,
) *types.Receipt {
	return waitForTransaction(ctx, l1Info, txHash, false)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals success.
func waitForTransaction(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	txHash common.Hash,
	success bool,
) *types.Receipt {
	cctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	receipt, err := WaitMined(cctx, l1Info.RPCClient, txHash)
	Expect(err).Should(BeNil())

	if success {
		if receipt.Status == types.ReceiptStatusFailed {
			TraceTransactionAndExit(ctx, l1Info.RPCClient, receipt.TxHash)
		}
	} else {
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusFailed))
	}
	return receipt
}

// Polls for a transaction receipt of the given txHash on each queryTicker tick until
// either a transaction receipt returned, or the context is cancelled or expired.
func waitForTransactionReceipt(
	cctx context.Context,
	rpcClient ethclient.Client,
	txHash common.Hash,
) (*types.Receipt, error) {
	queryTicker := time.NewTicker(200 * time.Millisecond)
	defer queryTicker.Stop()
	for {
		receipt, err := rpcClient.TransactionReceipt(cctx, txHash)
		if err == nil {
			return receipt, nil
		}

		if errors.Is(err, subnetEvmInterfaces.NotFound) {
			log.Debug("Transaction not yet mined")
		} else {
			log.Error("Receipt retrieval failed", "err", err)
			return nil, err
		}

		// Wait for the next round.
		select {
		case <-cctx.Done():
			return nil, cctx.Err()
		case <-queryTicker.C:
		}
	}
}

// Signs a transaction using the provided key for the specified chainID
func SignTransaction(tx *types.Transaction, key *ecdsa.PrivateKey, chainID *big.Int) *types.Transaction {
	txSigner := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(tx, txSigner, key)
	Expect(err).Should(BeNil())

	return signedTx
}

// Returns the gasFeeCap, gasTipCap, and nonce the be used when constructing a transaction from fundedAddress
func CalculateTxParams(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	fundedAddress common.Address,
) (*big.Int, *big.Int, uint64) {
	baseFee, err := l1Info.RPCClient.EstimateBaseFee(ctx)
	Expect(err).Should(BeNil())

	gasTipCap, err := l1Info.RPCClient.SuggestGasTipCap(ctx)
	Expect(err).Should(BeNil())

	nonce, err := l1Info.RPCClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	gasFeeCap := baseFee.Mul(baseFee, big.NewInt(gasUtils.BaseFeeFactor))
	gasFeeCap.Add(gasFeeCap, big.NewInt(gasUtils.MaxPriorityFeePerGas))

	return gasFeeCap, gasTipCap, nonce
}

// Gomega will print the transaction trace and exit
func TraceTransactionAndExit(ctx context.Context, rpcClient ethclient.Client, txHash common.Hash) {
	Expect(TraceTransaction(ctx, rpcClient, txHash)).Should(Equal(""))
}

func TraceTransaction(ctx context.Context, rpcClient ethclient.Client, txHash common.Hash) string {
	var result interface{}
	ct := "callTracer"
	err := rpcClient.Client().Call(&result, "debug_traceTransaction", txHash.String(), tracers.TraceConfig{Tracer: &ct})
	Expect(err).Should(BeNil())

	jsonStr, err := json.Marshal(result)
	Expect(err).Should(BeNil())

	return string(jsonStr)
}

//
// Block utils
//

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
// Takes a tx hash instead of the full tx in the subnet-evm version of this function.
// Copied and modified from https://github.com/ava-labs/subnet-evm/blob/v0.6.0-fuji/accounts/abi/bind/util.go#L42
func WaitMined(ctx context.Context, rpcClient ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	now := time.Now()
	receipt, err := waitForTransactionReceipt(ctx, rpcClient, txHash)
	if err != nil {
		return nil, err
	}
	since := time.Since(now)
	goLog.Println("Transaction mined", "txHash", txHash.Hex(), "duration", since)

	// Check that the block height endpoint returns a block height as high as the block number that the transaction was
	// included in. This is to workaround the issue where multiple nodes behind a public RPC endpoint see
	// transactions/blocks at different points in time. Ideally, all nodes in the network should have seen this block
	// and transaction before returning from WaitMined. The block height endpoint of public RPC endpoints is
	// configured to return the lowest value currently returned by any node behind the load balancer, so waiting for
	// it to be at least as high as the block height specified in the receipt should provide a relatively strong
	// indication that the transaction has been seen widely throughout the network.
	err = waitForBlockHeight(ctx, rpcClient, receipt.BlockNumber.Uint64())
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

// Polls for the eth_blockNumber endpoint for the latest blockheight on each queryTicker tick until
// either the returned height is greater than or equal to the expectedBlockNumber, or the context
// is cancelled or expired.
func waitForBlockHeight(
	cctx context.Context,
	rpcClient ethclient.Client,
	expectedBlockNumber uint64,
) error {
	queryTicker := time.NewTicker(2 * time.Second)
	defer queryTicker.Stop()
	for {
		currentBlockNumber, err := rpcClient.BlockNumber(cctx)
		if err != nil {
			return err
		}

		if currentBlockNumber >= expectedBlockNumber {
			return nil
		} else {
			log.Info("Waiting for block height where transaction was included",
				"blockNumber", expectedBlockNumber)
		}

		// Wait for the next round.
		select {
		case <-cctx.Done():
			return cctx.Err()
		case <-queryTicker.C:
		}
	}
}

//
// Log utils
//

// Returns the first log in 'logs' that is successfully parsed by 'parser'
// Errors and prints a trace of the transaction if no log is found.
func GetEventFromLogsOrTrace[T any](
	ctx context.Context,
	receipt *types.Receipt,
	l1Info interfaces.L1TestInfo,
	parser func(log types.Log) (T, error),
) T {
	log, err := GetEventFromLogs(receipt.Logs, parser)
	if err != nil {
		TraceTransactionAndExit(ctx, l1Info.RPCClient, receipt.TxHash)
	}
	return log
}

// Returns the first log in 'logs' that is successfully parsed by 'parser'
func GetEventFromLogs[T any](logs []*types.Log, parser func(log types.Log) (T, error)) (T, error) {
	for _, log := range logs {
		event, err := parser(*log)
		if err == nil {
			return event, nil
		}
	}
	return *new(T), fmt.Errorf("failed to find %T event in receipt logs", *new(T))
}

//
// Account utils
//

func PrivateKeyToAddress(k *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(k.PublicKey)
}

// Throws a Gomega error if there is a mismatch
func CheckBalance(ctx context.Context, addr common.Address, expectedBalance *big.Int, rpcClient ethclient.Client) {
	bal, err := rpcClient.BalanceAt(ctx, addr, nil)
	Expect(err).Should(BeNil())
	ExpectBigEqual(bal, expectedBalance)
}

//
// Big int utils
//

func ExpectBigEqual(v1 *big.Int, v2 *big.Int) {
	// Compare strings, so gomega will print the numbers if they differ
	Expect(v1.String()).Should(Equal(v2.String()))
}

func BigIntSub(v1 *big.Int, v2 *big.Int) *big.Int {
	return big.NewInt(0).Sub(v1, v2)
}

func BigIntMul(v1 *big.Int, v2 *big.Int) *big.Int {
	return big.NewInt(0).Mul(v1, v2)
}

//
// Network utils
//

func GetPChainInfo(cChainInfo interfaces.L1TestInfo) interfaces.L1TestInfo {
	pChainBlockchainID, err := info.NewClient(cChainInfo.NodeURIs[0]).GetBlockchainID(context.Background(), "P")
	Expect(err).Should(BeNil())
	return interfaces.L1TestInfo{
		BlockchainID: pChainBlockchainID,
		SubnetID:     ids.Empty,
	}
}

type ChainConfigMap map[string]string

// Sets the chain config in customChainConfigs for the specified L!
func (m ChainConfigMap) Add(l1 interfaces.L1TestInfo, chainConfig string) {
	if l1.SubnetID == constants.PrimaryNetworkID {
		m[CChainPathSpecifier] = chainConfig
	} else {
		m[l1.BlockchainID.String()] = chainConfig
	}
}

func GetChainConfigWithOffChainMessages(offChainMessages []avalancheWarp.UnsignedMessage) string {
	// Convert messages to hex
	hexOffChainMessages := []string{}
	for _, message := range offChainMessages {
		hexOffChainMessages = append(hexOffChainMessages, hexutil.Encode(message.Bytes()))
	}

	chainConfig := WarpEnabledChainConfig
	chainConfig["warp-off-chain-messages"] = hexOffChainMessages

	// Marshal the map to JSON
	offChainMessageJson, err := tmpnet.DefaultJSONMarshal(chainConfig)
	Expect(err).Should(BeNil())

	return string(offChainMessageJson)
}

// read in the template file, make the substitutions declared at the beginning
// of the function, write out the instantiation to a temp file, and then return
// the path to that temp file.
func InstantiateGenesisTemplate(
	templateFileName string,
	chainID uint64,
	teleporterContractAddress common.Address,
	teleporterDeployedBytecode string,
	teleporterDeployerAddress common.Address,
	requirePrimaryNetworkSigners bool,
) string {
	if teleporterContractAddress.Big().Cmp(big.NewInt(0)) == 0 {
		teleporterContractAddress = common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff")
	}
	if teleporterDeployerAddress.Big().Cmp(big.NewInt(0)) == 0 {
		teleporterDeployerAddress = common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
	if teleporterDeployedBytecode == "" {
		teleporterDeployedBytecode = "0x00"
	}

	substitutions := []struct {
		Target string
		Value  string
	}{
		{
			"<EVM_CHAIN_ID>",
			strconv.FormatUint(chainID, 10),
		},
		{
			"<TELEPORTER_MESSENGER_CONTRACT_ADDRESS>",
			teleporterContractAddress.Hex(),
		},
		{
			"<TELEPORTER_MESSENGER_DEPLOYED_BYTECODE>",
			teleporterDeployedBytecode,
		},
		{
			"<TELEPORTER_MESSENGER_DEPLOYER_ADDRESS>",
			teleporterDeployerAddress.Hex(),
		},
		{
			"<REQUIRE_PRIMARY_NETWORK_SIGNERS>",
			strconv.FormatBool(requirePrimaryNetworkSigners),
		},
	}

	templateFileBytes, err := os.ReadFile(templateFileName)
	Expect(err).Should(BeNil())

	l1GenesisFile, err := os.CreateTemp(os.TempDir(), "")
	Expect(err).Should(BeNil())

	defer l1GenesisFile.Close()

	var replaced string = string(templateFileBytes[:])
	for _, s := range substitutions {
		replaced = strings.ReplaceAll(replaced, s.Target, s.Value)
	}

	l1GenesisFile.WriteString(replaced)

	return l1GenesisFile.Name()
}

// Native minter utils
//

// Funded key must have admin access to set new admin.
func AddNativeMinterAdmin(
	ctx context.Context,
	l1 interfaces.L1TestInfo,
	fundedKey *ecdsa.PrivateKey,
	address common.Address,
) {
	nativeMinterPrecompile, err := nativeMinter.NewINativeMinter(nativeminter.ContractAddress, l1.RPCClient)
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := nativeMinterPrecompile.SetAdmin(opts, address)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, l1, tx.Hash())
}

// Blocks until all validators specified in nodeURIs have reached the specified block height
func WaitForAllValidatorsToAcceptBlock(ctx context.Context, nodeURIs []string, blockchainID ids.ID, height uint64) {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for i, uri := range nodeURIs {
		chainAWSURI := HttpToWebsocketURI(uri, blockchainID.String())
		log.Debug("Creating ethclient for blockchain", "blockchainID", blockchainID.String(), "wsURI", chainAWSURI)
		client, err := ethclient.Dial(chainAWSURI)
		Expect(err).Should(BeNil())
		defer client.Close()

		// Loop until each node has advanced to >= the height of the block that emitted the warp log
		for {
			block, err := client.BlockByNumber(cctx, nil)
			Expect(err).Should(BeNil())
			if block.NumberU64() >= height {
				log.Debug("Client accepted the block containing SendWarpMessage", "client", i, "height", block.NumberU64())
				break
			}
		}
	}
}

func ExtractWarpMessageFromLog(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.L1TestInfo,
) *avalancheWarp.UnsignedMessage {
	log.Info("Fetching relevant warp logs from the newly produced block")
	logs, err := source.RPCClient.FilterLogs(ctx, subnetEvmInterfaces.FilterQuery{
		BlockHash: &sourceReceipt.BlockHash,
		Addresses: []common.Address{warp.Module.Address},
	})
	Expect(err).Should(BeNil())
	Expect(len(logs)).Should(Equal(1))

	// Check for relevant warp log from subscription and ensure that it matches
	// the log extracted from the last block.
	txLog := logs[0]
	log.Info("Parsing logData as unsigned warp message")
	unsignedMsg, err := warp.UnpackSendWarpEventDataToMessage(txLog.Data)
	Expect(err).Should(BeNil())
	return unsignedMsg
}

func ConstructSignedWarpMessage(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	justification []byte,
	signatureAggregator *SignatureAggregator,
) *avalancheWarp.Message {
	unsignedMsg := ExtractWarpMessageFromLog(ctx, sourceReceipt, source)

	// Loop over each client on source chain to ensure they all have time to accept the block.
	// Note: if we did not confirm this here, the next stage could be racy since it assumes every node
	// has accepted the block.
	WaitForAllValidatorsToAcceptBlock(ctx, source.NodeURIs, source.BlockchainID, sourceReceipt.BlockNumber.Uint64())

	// Get the aggregate signature for the Warp message
	log.Info("Fetching aggregate signature from the source chain validators")
	return GetSignedMessage(source, destination, unsignedMsg, justification, signatureAggregator)
}

func GetSignedMessage(
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	unsignedWarpMessage *avalancheWarp.UnsignedMessage,
	justification []byte,
	signatureAggregator *SignatureAggregator,
) *avalancheWarp.Message {
	signingSubnetID := source.SubnetID
	if source.SubnetID == constants.PrimaryNetworkID && !destination.RequirePrimaryNetworkSigners {
		signingSubnetID = destination.SubnetID
	}

	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedWarpMessage,
		justification,
		signingSubnetID,
		warp.WarpDefaultQuorumNumerator,
	)
	Expect(err).Should(BeNil())

	return signedWarpMessage
}

func SetupProposerVM(ctx context.Context, fundedKey *ecdsa.PrivateKey, network *tmpnet.Network, SubnetID ids.ID) {
	l1Details := network.Subnets[slices.IndexFunc(
		network.Subnets,
		func(s *tmpnet.Subnet) bool { return s.SubnetID == SubnetID },
	)]

	chainID := l1Details.Chains[0].ChainID

	node, err := network.GetNode(l1Details.ValidatorIDs[0])
	Expect(err).Should(BeNil())
	nodeURI, _, err := node.GetLocalURI(context.Background())
	Expect(err).Should(BeNil())
	uri := HttpToWebsocketURI(nodeURI, chainID.String())

	client, err := ethclient.Dial(uri)
	Expect(err).Should(BeNil())
	chainIDInt, err := client.ChainID(ctx)
	Expect(err).Should(BeNil())

	err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(ctx, chainIDInt, fundedKey, client)
	Expect(err).Should(BeNil())
}
