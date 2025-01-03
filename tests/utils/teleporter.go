package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	validatorsetsig "github.com/ava-labs/icm-contracts/abi-bindings/go/governance/ValidatorSetSig"
	exampleerc20 "github.com/ava-labs/icm-contracts/abi-bindings/go/mocks/ExampleERC20"
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	teleporterregistry "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/registry/TeleporterRegistry"
	testmessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/tests/TestMessenger"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	deploymentUtils "github.com/ava-labs/icm-contracts/utils/deployment-utils"
	gasUtils "github.com/ava-labs/icm-contracts/utils/gas-utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

var (
	DefaultTeleporterTransactionGas   uint64 = 300_000
	DefaultTeleporterTransactionValue        = common.Big0
)

type ChainTeleporterInfo struct {
	TeleporterRegistry        *teleporterregistry.TeleporterRegistry
	TeleporterRegistryAddress common.Address

	TeleporterMessenger        *teleportermessenger.TeleporterMessenger
	TeleporterMessengerAddress common.Address
}

type TeleporterTestInfo map[ids.ID]*ChainTeleporterInfo

func NewTeleporterTestInfo(l1s []interfaces.L1TestInfo) TeleporterTestInfo {
	t := make(TeleporterTestInfo)
	for _, l1 := range l1s {
		t[l1.BlockchainID] = &ChainTeleporterInfo{}
	}
	return t
}

func (t TeleporterTestInfo) TeleporterMessenger(
	l1 interfaces.L1TestInfo,
) *teleportermessenger.TeleporterMessenger {
	return t[l1.BlockchainID].TeleporterMessenger
}

func (t TeleporterTestInfo) TeleporterMessengerAddress(l1 interfaces.L1TestInfo) common.Address {
	return t[l1.BlockchainID].TeleporterMessengerAddress
}

func (t TeleporterTestInfo) TeleporterRegistry(
	l1 interfaces.L1TestInfo,
) *teleporterregistry.TeleporterRegistry {
	return t[l1.BlockchainID].TeleporterRegistry
}

func (t TeleporterTestInfo) TeleporterRegistryAddress(l1 interfaces.L1TestInfo) common.Address {
	return t[l1.BlockchainID].TeleporterRegistryAddress
}

func (t TeleporterTestInfo) SetTeleporter(address common.Address, l1 interfaces.L1TestInfo) {
	teleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		address, l1.RPCClient,
	)
	Expect(err).Should(BeNil())
	info := t[l1.BlockchainID]
	info.TeleporterMessengerAddress = address
	info.TeleporterMessenger = teleporterMessenger
}

func (t TeleporterTestInfo) InitializeBlockchainID(l1 interfaces.L1TestInfo, fundedKey *ecdsa.PrivateKey) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := t.TeleporterMessenger(l1).InitializeBlockchainID(opts)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(context.Background(), l1, tx.Hash())
}

func (t TeleporterTestInfo) DeployTeleporterRegistry(l1 interfaces.L1TestInfo, deployerKey *ecdsa.PrivateKey) {
	ctx := context.Background()
	entries := []teleporterregistry.ProtocolRegistryEntry{
		{
			Version:         big.NewInt(1),
			ProtocolAddress: t.TeleporterMessengerAddress(l1),
		},
	}
	opts, err := bind.NewKeyedTransactorWithChainID(deployerKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	teleporterRegistryAddress, tx, teleporterRegistry, err := teleporterregistry.DeployTeleporterRegistry(
		opts, l1.RPCClient, entries,
	)
	Expect(err).Should(BeNil())
	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	info := t[l1.BlockchainID]
	info.TeleporterRegistryAddress = teleporterRegistryAddress
	info.TeleporterRegistry = teleporterRegistry
}

func (t TeleporterTestInfo) DeployTeleporterMessenger(
	ctx context.Context,
	l1 interfaces.L1TestInfo,
	transactionBytes []byte,
	deployerAddress common.Address,
	contractAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
) {
	// Fund the deployer address
	fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(11)) // 11 AVAX
	fundDeployerTx := CreateNativeTransferTransaction(
		ctx, l1, fundedKey, deployerAddress, fundAmount,
	)
	SendTransactionAndWaitForSuccess(ctx, l1, fundDeployerTx)

	log.Info("Finished funding Teleporter deployer", "blockchainID", l1.BlockchainID.Hex())

	// Deploy Teleporter contract
	rpcClient, err := rpc.DialContext(
		ctx,
		HttpToRPCURI(l1.NodeURIs[0], l1.BlockchainID.String()),
	)
	Expect(err).Should(BeNil())
	defer rpcClient.Close()

	txHash := common.Hash{}
	err = rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, l1, txHash)

	teleporterCode, err := l1.RPCClient.CodeAt(ctx, contractAddress, nil)
	Expect(err).Should(BeNil())
	Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
}

func (t TeleporterTestInfo) RelayTeleporterMessage(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	expectSuccess bool,
	fundedKey *ecdsa.PrivateKey,
	justification []byte,
	signatureAggregator *SignatureAggregator,
) *types.Receipt {
	// Fetch the Teleporter message from the logs
	sendEvent, err := GetEventFromLogs(sourceReceipt.Logs, t.TeleporterMessenger(source).ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())

	signedWarpMessage := ConstructSignedWarpMessage(
		ctx,
		sourceReceipt,
		source,
		destination,
		justification,
		signatureAggregator,
	)

	// Construct the transaction to send the Warp message to the destination chain
	signedTx := CreateReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessage,
		sendEvent.Message.RequiredGasLimit,
		t.TeleporterMessengerAddress(source),
		fundedKey,
		destination,
	)

	log.Info("Sending transaction to destination chain")
	if !expectSuccess {
		return SendTransactionAndWaitForFailure(ctx, destination, signedTx)
	}

	receipt := SendTransactionAndWaitForSuccess(ctx, destination, signedTx)

	// Check the transaction logs for the ReceiveCrossChainMessage event emitted by the Teleporter contract
	receiveEvent, err := GetEventFromLogs(
		receipt.Logs,
		t.TeleporterMessenger(destination).ParseReceiveCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(receiveEvent.SourceBlockchainID[:]).Should(Equal(source.BlockchainID[:]))
	return receipt
}

func (t TeleporterTestInfo) SendExampleCrossChainMessageAndVerify(
	ctx context.Context,
	source interfaces.L1TestInfo,
	sourceExampleMessenger *testmessenger.TestMessenger,
	destination interfaces.L1TestInfo,
	destExampleMessengerAddress common.Address,
	destExampleMessenger *testmessenger.TestMessenger,
	senderKey *ecdsa.PrivateKey,
	message string,
	signatureAggregator *SignatureAggregator,
	expectSuccess bool,
) {
	// Call the example messenger contract on L1 A
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := sourceExampleMessenger.SendMessage(
		optsA,
		destination.BlockchainID,
		destExampleMessengerAddress,
		common.BigToAddress(common.Big0),
		big.NewInt(0),
		testmessenger.SendMessageRequiredGas,
		message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := WaitForTransactionSuccess(ctx, source, tx.Hash())

	sourceTeleporterMessenger := t.TeleporterMessenger(source)
	destTeleporterMessenger := t.TeleporterMessenger(destination)

	event, err := GetEventFromLogs(receipt.Logs, sourceTeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destination.BlockchainID[:]))

	teleporterMessageID := event.MessageID

	//
	// Relay the message to the destination
	//
	receipt = t.RelayTeleporterMessage(ctx, receipt, source, destination, true, senderKey, nil, signatureAggregator)

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := destTeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	if expectSuccess {
		// Check that message execution was successful
		messageExecutedEvent, err := GetEventFromLogs(
			receipt.Logs,
			destTeleporterMessenger.ParseMessageExecuted,
		)
		Expect(err).Should(BeNil())
		Expect(messageExecutedEvent.MessageID[:]).Should(Equal(teleporterMessageID[:]))
	} else {
		// Check that message execution failed
		messageExecutionFailedEvent, err := GetEventFromLogs(
			receipt.Logs,
			destTeleporterMessenger.ParseMessageExecutionFailed,
		)
		Expect(err).Should(BeNil())
		Expect(messageExecutionFailedEvent.MessageID[:]).Should(Equal(teleporterMessageID[:]))
	}

	//
	// Verify we received the expected string
	//
	_, currMessage, err := destExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, source.BlockchainID)
	Expect(err).Should(BeNil())
	if expectSuccess {
		Expect(currMessage).Should(Equal(message))
	} else {
		Expect(currMessage).ShouldNot(Equal(message))
	}
}

func (t TeleporterTestInfo) AddProtocolVersionAndWaitForAcceptance(
	ctx context.Context,
	l1 interfaces.L1TestInfo,
	newTeleporterAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	signatureAggregator *SignatureAggregator,
) {
	signedWarpMsg := GetSignedMessage(l1, l1, unsignedMessage, nil, signatureAggregator)
	log.Info("Got signed warp message", "messageID", signedWarpMsg.ID())

	// Construct tx to add protocol version and send to destination chain
	signedTx := CreateAddProtocolVersionTransaction(
		ctx,
		signedWarpMsg,
		t.TeleporterRegistryAddress(l1),
		senderKey,
		l1,
	)

	curLatestVersion := t.GetLatestTeleporterVersion(l1)
	expectedLatestVersion := big.NewInt(curLatestVersion.Int64() + 1)

	// Wait for tx to be accepted, and verify events emitted
	receipt := SendTransactionAndWaitForSuccess(ctx, l1, signedTx)
	teleporterRegistry := t.TeleporterRegistry(l1)
	addProtocolVersionEvent, err := GetEventFromLogs(receipt.Logs, teleporterRegistry.ParseAddProtocolVersion)
	Expect(err).Should(BeNil())
	Expect(addProtocolVersionEvent.Version.Cmp(expectedLatestVersion)).Should(Equal(0))
	Expect(addProtocolVersionEvent.ProtocolAddress).Should(Equal(newTeleporterAddress))

	versionUpdatedEvent, err := GetEventFromLogs(receipt.Logs, teleporterRegistry.ParseLatestVersionUpdated)
	Expect(err).Should(BeNil())
	Expect(versionUpdatedEvent.OldVersion.Cmp(curLatestVersion)).Should(Equal(0))
	Expect(versionUpdatedEvent.NewVersion.Cmp(expectedLatestVersion)).Should(Equal(0))
}

func (t TeleporterTestInfo) GetLatestTeleporterVersion(l1 interfaces.L1TestInfo) *big.Int {
	version, err := t.TeleporterRegistry(l1).LatestVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	return version
}

func (t TeleporterTestInfo) ClearReceiptQueue(
	ctx context.Context,
	fundedKey *ecdsa.PrivateKey,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	signatureAggregator *SignatureAggregator,
) {
	sourceTeleporterMessenger := t.TeleporterMessenger(source)
	outstandReceiptCount := GetOutstandingReceiptCount(
		t.TeleporterMessenger(source),
		destination.BlockchainID,
	)
	for outstandReceiptCount.Cmp(big.NewInt(0)) != 0 {
		log.Info("Emptying receipt queue", "remainingReceipts", outstandReceiptCount.String())
		// Send message from L1 B to L1 A to trigger the "regular" method of delivering receipts.
		// The next message from B->A will contain the same receipts that were manually sent in the above steps,
		// but they should not be processed again on L1 A.
		sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
			DestinationBlockchainID: destination.BlockchainID,
			DestinationAddress:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
			RequiredGasLimit:        big.NewInt(1),
			FeeInfo: teleportermessenger.TeleporterFeeInfo{
				FeeTokenAddress: common.Address{},
				Amount:          big.NewInt(0),
			},
			AllowedRelayerAddresses: []common.Address{},
			Message:                 []byte{1, 2, 3, 4},
		}

		// This message will also have the same receipts as the previous message
		receipt, _ := SendCrossChainMessageAndWaitForAcceptance(
			ctx, sourceTeleporterMessenger, source, destination, sendCrossChainMessageInput, fundedKey)

		// Relay message
		t.RelayTeleporterMessage(ctx, receipt, source, destination, true, fundedKey, nil, signatureAggregator)

		outstandReceiptCount = GetOutstandingReceiptCount(sourceTeleporterMessenger, destination.BlockchainID)
	}
	log.Info("Receipt queue emptied")
}

// Deploys a new version of Teleporter and returns its address
// Does NOT modify the global Teleporter contract address to provide greater testing flexibility.
func (t TeleporterTestInfo) DeployNewTeleporterVersion(
	ctx context.Context,
	l1 interfaces.L1TestInfo,
	fundedKey *ecdsa.PrivateKey,
	teleporterByteCodeFile string,
) common.Address {
	contractCreationGasPrice := (&big.Int{}).Add(deploymentUtils.GetDefaultContractCreationGasPrice(), big.NewInt(1))
	teleporterDeployerTransaction,
		_,
		teleporterDeployerAddress,
		teleporterContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		teleporterByteCodeFile,
		false,
		contractCreationGasPrice,
	)
	Expect(err).Should(BeNil())

	t.DeployTeleporterMessenger(
		ctx,
		l1,
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)

	return teleporterContractAddress
}

//
// Deployment utils
//

func DeployTestMessenger(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	teleporterManager common.Address,
	registryAddress common.Address,
	l1 interfaces.L1TestInfo,
) (common.Address, *testmessenger.TestMessenger) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		l1.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, exampleMessenger, err := testmessenger.DeployTestMessenger(
		opts,
		l1.RPCClient,
		registryAddress,
		teleporterManager,
		big.NewInt(1),
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	return address, exampleMessenger
}

//
// Parsing utils
//

func ParseTeleporterMessage(unsignedMessage avalancheWarp.UnsignedMessage) *teleportermessenger.TeleporterMessage {
	addressedPayload, err := payload.ParseAddressedCall(unsignedMessage.Payload)
	Expect(err).Should(BeNil())

	teleporterMessage := teleportermessenger.TeleporterMessage{}
	err = teleporterMessage.Unpack(addressedPayload.Payload)
	Expect(err).Should(BeNil())

	return &teleporterMessage
}

//
// Function call utils
//

func SendAddFeeAmountAndWaitForAcceptance(
	ctx context.Context,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	messageID ids.ID,
	amount *big.Int,
	feeContractAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		source.EVMChainID,
	)
	Expect(err).Should(BeNil())

	tx, err := transactor.AddFeeAmount(opts, messageID, feeContractAddress, amount)
	Expect(err).Should(BeNil())

	receipt := WaitForTransactionSuccess(ctx, source, tx.Hash())

	addFeeAmountEvent, err := GetEventFromLogs(receipt.Logs, transactor.ParseAddFeeAmount)
	Expect(err).Should(BeNil())
	Expect(addFeeAmountEvent.MessageID[:]).Should(Equal(messageID[:]))

	log.Info("Send AddFeeAmount transaction on source chain",
		"messageID", messageID,
		"sourceChainID", source.BlockchainID,
		"destinationBlockchainID", destination.BlockchainID,
	)

	return receipt
}

func RetryMessageExecutionAndWaitForAcceptance(
	ctx context.Context,
	sourceBlockchainID ids.ID,
	destinationTeleporterMessenger *teleportermessenger.TeleporterMessenger,
	destinationL1 interfaces.L1TestInfo,
	message teleportermessenger.TeleporterMessage,
	senderKey *ecdsa.PrivateKey,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, destinationL1.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := destinationTeleporterMessenger.RetryMessageExecution(opts, sourceBlockchainID, message)
	Expect(err).Should(BeNil())

	return WaitForTransactionSuccess(ctx, destinationL1, tx.Hash())
}

func RedeemRelayerRewardsAndConfirm(
	ctx context.Context,
	teleporterMessenger *teleportermessenger.TeleporterMessenger,
	l1 interfaces.L1TestInfo,
	feeToken *exampleerc20.ExampleERC20,
	feeTokenAddress common.Address,
	redeemerKey *ecdsa.PrivateKey,
	expectedAmount *big.Int,
) *types.Receipt {
	redeemerAddress := crypto.PubkeyToAddress(redeemerKey.PublicKey)

	// Check the ERC20 balance before redemption
	balanceBeforeRedemption, err := feeToken.BalanceOf(
		&bind.CallOpts{}, redeemerAddress,
	)
	Expect(err).Should(BeNil())

	// Redeem the rewards
	txOpts, err := bind.NewKeyedTransactorWithChainID(
		redeemerKey, l1.EVMChainID,
	)
	Expect(err).Should(BeNil())
	tx, err := teleporterMessenger.RedeemRelayerRewards(
		txOpts, feeTokenAddress,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, l1, tx.Hash())

	// Check that the ERC20 balance was incremented
	balanceAfterRedemption, err := feeToken.BalanceOf(
		&bind.CallOpts{}, redeemerAddress,
	)
	Expect(err).Should(BeNil())
	Expect(balanceAfterRedemption).Should(
		Equal(
			big.NewInt(0).Add(
				balanceBeforeRedemption, expectedAmount,
			),
		),
	)

	// Check that the redeemable rewards amount is now zero.
	updatedRewardAmount, err := teleporterMessenger.CheckRelayerRewardAmount(
		&bind.CallOpts{},
		redeemerAddress,
		feeTokenAddress,
	)
	Expect(err).Should(BeNil())
	Expect(updatedRewardAmount.Cmp(big.NewInt(0))).Should(Equal(0))

	return receipt
}

func SendSpecifiedReceiptsAndWaitForAcceptance(
	ctx context.Context,
	sourceTeleporterMessenger *teleportermessenger.TeleporterMessenger,
	source interfaces.L1TestInfo,
	destinationBlockchainID ids.ID,
	messageIDs [][32]byte,
	feeInfo teleportermessenger.TeleporterFeeInfo,
	allowedRelayerAddresses []common.Address,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := sourceTeleporterMessenger.SendSpecifiedReceipts(
		opts, destinationBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
	Expect(err).Should(BeNil())

	receipt := WaitForTransactionSuccess(ctx, source, tx.Hash())

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, sourceTeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationBlockchainID[:]))

	log.Info("Sending SendSpecifiedReceipts transaction",
		"destinationBlockchainID", destinationBlockchainID,
		"txHash", tx.Hash())

	return receipt, event.MessageID
}

func SendCrossChainMessageAndWaitForAcceptance(
	ctx context.Context,
	sourceTeleporterMessenger *teleportermessenger.TeleporterMessenger,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	input teleportermessenger.TeleporterMessageInput,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	// Send a transaction to the Teleporter contract
	tx, err := sourceTeleporterMessenger.SendCrossChainMessage(opts, input)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be accepted
	receipt := WaitForTransactionSuccess(ctx, source, tx.Hash())

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, sourceTeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())

	log.Info("Sending SendCrossChainMessage transaction on source chain",
		"sourceChainID", source.BlockchainID,
		"destinationBlockchainID", destination.BlockchainID,
		"txHash", tx.Hash())

	return receipt, event.MessageID
}

// Returns true if the transaction receipt contains a ReceiptReceived log with the specified messageID
func CheckReceiptReceived(
	receipt *types.Receipt,
	messageID [32]byte,
	transactor *teleportermessenger.TeleporterMessenger,
) bool {
	for _, log := range receipt.Logs {
		event, err := transactor.ParseReceiptReceived(*log)
		if err == nil && bytes.Equal(event.MessageID[:], messageID[:]) {
			return true
		}
	}
	return false
}

func GetOutstandingReceiptCount(
	teleporterMessenger *teleportermessenger.TeleporterMessenger,
	destinationBlockchainID ids.ID,
) *big.Int {
	size, err := teleporterMessenger.GetReceiptQueueSize(&bind.CallOpts{}, destinationBlockchainID)
	Expect(err).Should(BeNil())
	return size
}

//
// Transaction utils
//

// Constructs a transaction to call sendCrossChainMessage
// Returns the signed transaction.
func CreateSendCrossChainMessageTransaction(
	ctx context.Context,
	source interfaces.L1TestInfo,
	input teleportermessenger.TeleporterMessageInput,
	senderKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	data, err := teleportermessenger.PackSendCrossChainMessage(input)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, source, PrivateKeyToAddress(senderKey))

	// Send a transaction to the Teleporter contract
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   source.EVMChainID,
		Nonce:     nonce,
		To:        &teleporterContractAddress,
		Gas:       DefaultTeleporterTransactionGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     DefaultTeleporterTransactionValue,
		Data:      data,
	})

	return SignTransaction(tx, senderKey, source.EVMChainID)
}

// Constructs a transaction to call receiveCrossChainMessage
// Returns the signed transaction.
func CreateReceiveCrossChainMessageTransaction(
	ctx context.Context,
	signedMessage *avalancheWarp.Message,
	requiredGasLimit *big.Int,
	teleporterContractAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
) *types.Transaction {
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing receiveCrossChainMessage transaction for the destination chain")
	numSigners, err := signedMessage.Signature.NumSigners()
	Expect(err).Should(BeNil())

	teleporterMessage := ParseTeleporterMessage(signedMessage.UnsignedMessage)
	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(
		numSigners,
		requiredGasLimit,
		len(signedMessage.Bytes()),
		len(signedMessage.Payload),
		len(teleporterMessage.Receipts),
	)
	Expect(err).Should(BeNil())

	callData, err := teleportermessenger.PackReceiveCrossChainMessage(0, PrivateKeyToAddress(senderKey))
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, l1Info, PrivateKeyToAddress(senderKey))

	destinationTx := predicateutils.NewPredicateTx(
		l1Info.EVMChainID,
		nonce,
		&teleporterContractAddress,
		gasLimit,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessage.Bytes(),
	)

	return SignTransaction(destinationTx, senderKey, l1Info.EVMChainID)
}

// Constructs a transaction to call addProtocolVersion
// Returns the signed transaction.
func CreateAddProtocolVersionTransaction(
	ctx context.Context,
	signedMessage *avalancheWarp.Message,
	teleporterRegistryAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
) *types.Transaction {
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing addProtocolVersion transaction for the destination chain")

	callData, err := teleporterregistry.PackAddProtocolVersion(0)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, l1Info, PrivateKeyToAddress(senderKey))

	destinationTx := predicateutils.NewPredicateTx(
		l1Info.EVMChainID,
		nonce,
		&teleporterRegistryAddress,
		500_000,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessage.Bytes(),
	)

	return SignTransaction(destinationTx, senderKey, l1Info.EVMChainID)
}

func CreateExecuteCallPredicateTransaction(
	ctx context.Context,
	signedMessage *avalancheWarp.Message,
	validatorSetSigAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
) *types.Transaction {
	log.Info("Constructing executeCall transaction for the destination chain")

	callData, err := validatorsetsig.PackExecuteCall(0)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, l1Info, PrivateKeyToAddress(senderKey))

	destinationTx := predicateutils.NewPredicateTx(
		l1Info.EVMChainID,
		nonce,
		&validatorSetSigAddress,
		500_000,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessage.Bytes(),
	)
	return SignTransaction(destinationTx, senderKey, l1Info.EVMChainID)
}

//
// Off-chain message utils
//

// Creates an Warp message that registers a Teleporter protocol version with TeleporterRegistry.
// Returns the Warp message, as well as the chain config adding the message to the list of approved
// off-chain Warp messages
func InitOffChainMessageChainConfig(
	networkID uint32,
	l1 interfaces.L1TestInfo,
	registryAddress common.Address,
	teleporterAddress common.Address,
	version uint64,
) (*avalancheWarp.UnsignedMessage, string) {
	unsignedMessage := CreateOffChainRegistryMessage(
		networkID,
		l1,
		registryAddress,
		teleporterregistry.ProtocolRegistryEntry{
			Version:         big.NewInt(int64(version)),
			ProtocolAddress: teleporterAddress,
		},
	)
	log.Info("Adding off-chain message to Warp chain config",
		"messageID", unsignedMessage.ID(),
		"blockchainID", l1.BlockchainID.String(),
	)

	return unsignedMessage, GetChainConfigWithOffChainMessages([]avalancheWarp.UnsignedMessage{*unsignedMessage})
}

// Creates an off-chain Warp message that registers a Teleporter protocol version with TeleporterRegistry
func CreateOffChainRegistryMessage(
	networkID uint32,
	l1 interfaces.L1TestInfo,
	registryAddress common.Address,
	entry teleporterregistry.ProtocolRegistryEntry,
) *avalancheWarp.UnsignedMessage {
	sourceAddress := []byte{}
	payloadBytes, err := teleporterregistry.PackTeleporterRegistryWarpPayload(entry, registryAddress)
	Expect(err).Should(BeNil())

	addressedPayload, err := payload.NewAddressedCall(sourceAddress, payloadBytes)
	Expect(err).Should(BeNil())

	unsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		l1.BlockchainID,
		addressedPayload.Bytes(),
	)
	Expect(err).Should(BeNil())

	return unsignedMessage
}
