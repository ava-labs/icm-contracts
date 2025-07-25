package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	avalanchevalidatorsetregistry "github.com/ava-labs/icm-contracts/abi-bindings/go/AvalancheValidatorSetRegistry"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/log"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	. "github.com/onsi/gomega"
)

func DeployAvalancheValidatorSetRegistry(
	ctx context.Context,
	fundedKey *ecdsa.PrivateKey,
	chainID *big.Int,
	client ethclient.Client,
	avalancheNetworkID uint32,
) (common.Address, *avalanchevalidatorsetregistry.AvalancheValidatorSetRegistry) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, chainID)
	Expect(err).Should(BeNil())

	address, tx, validatorSetRegistry, err := avalanchevalidatorsetregistry.DeployAvalancheValidatorSetRegistry(opts, client, avalancheNetworkID)
	Expect(err).Should(BeNil())

	WaitForTransactionSuccessWithClient(ctx, client, tx.Hash())

	return address, validatorSetRegistry
}

func RegisterAvalancheValidatorSet(
	ctx context.Context,
	fundedKey *ecdsa.PrivateKey,
	chainID *big.Int,
	client ethclient.Client,
	validatorSetRegistry *avalanchevalidatorsetregistry.AvalancheValidatorSetRegistry,
	avalancheNetworkID uint32,
	avalanchePChainBlockchainID ids.ID,
	avalancheSubnetID ids.ID,
	avalancheL1BlockchainID ids.ID,
	pChainClient *platformvm.Client,
	signatureAggregator *SignatureAggregator,
) *big.Int {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, chainID)
	Expect(err).Should(BeNil())

	// Get the validator set information to create the serialized pre-image.
	pChainHeight, err := pChainClient.GetHeight(ctx)
	Expect(err).Should(BeNil())
	pChainBlockBytes, err := pChainClient.GetBlockByHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())

	// Parse the block
	pChainBlock, err := block.Parse(block.Codec, pChainBlockBytes)
	Expect(err).Should(BeNil())

	// Confirm that the block is a Banff block
	banffBlock, ok := pChainBlock.(block.BanffBlock)
	Expect(ok).Should(BeTrue())

	// Get the validators from the block height
	rawValidators, err := pChainClient.GetValidatorsAt(ctx, avalancheSubnetID, api.Height(pChainHeight))
	Expect(err).Should(BeNil())

	canonicalValidatorSet, err := warp.FlattenValidatorSet(rawValidators)
	Expect(err).Should(BeNil())

	log.Info("Canonical validator set", "numValidators", len(canonicalValidatorSet.Validators), "totalWeight", canonicalValidatorSet.TotalWeight)

	// Create the unsigned message including the validator set hash in the payload.
	validators := make([]*message.Validator, len(canonicalValidatorSet.Validators))
	for i, validator := range canonicalValidatorSet.Validators {
		validators[i] = &message.Validator{
			UncompressedPublicKeyBytes: [96]byte(validator.PublicKey.Serialize()),
			Weight:                     validator.Weight,
		}
	}
	validatorsBytes, err := message.Codec.Marshal(message.CodecVersion, validators)
	Expect(err).Should(BeNil())

	validatorSetHash := sha256.Sum256(validatorsBytes)
	validaterSetStatePayload, err := message.NewValidatorSetState(
		avalancheL1BlockchainID,
		pChainHeight,
		uint64(banffBlock.Timestamp().Unix()),
		validatorSetHash,
	)
	Expect(err).Should(BeNil())

	addressCallPayload, err := payload.NewAddressedCall([]byte{}, validaterSetStatePayload.Bytes())
	Expect(err).Should(BeNil())

	// Create the unsigned message
	unsignedMessage, err := warp.NewUnsignedMessage(avalancheNetworkID, avalanchePChainBlockchainID, addressCallPayload.Bytes())
	Expect(err).Should(BeNil())

	// Generate the signature of the unsigned message using the signature aggregator.
	log.Info("Generating signature of unsigned message", "networkID", avalancheNetworkID, "pChainBlockchainID", avalanchePChainBlockchainID, "subnetID", avalancheSubnetID, "weightThreshold", 70)
	signedMessage, err := signatureAggregator.CreateSignedMessage(unsignedMessage, nil, avalancheSubnetID, 70)
	Expect(err).Should(BeNil())
	bitSetSignature, ok := signedMessage.Signature.(*warp.BitSetSignature)
	Expect(ok).Should(BeTrue())

	// Parse the BLS signature
	blsSignature, err := bls.SignatureFromBytes(bitSetSignature.Signature[:])
	Expect(err).Should(BeNil())

	// Serialize the BLS signature
	blsSignatureBytes := blsSignature.Serialize()

	// Create the ICM message struct
	icmMessage := avalanchevalidatorsetregistry.ICMMessage{
		UnsignedMessage: avalanchevalidatorsetregistry.ICMUnsignedMessage{
			AvalancheNetworkID:          avalancheNetworkID,
			AvalancheSourceBlockchainID: avalanchePChainBlockchainID,
			Payload:                     addressCallPayload.Bytes(),
		},
		UnsignedMessageBytes: unsignedMessage.Bytes(),
		Signature: avalanchevalidatorsetregistry.ICMSignature{
			Signers:   bitSetSignature.Signers,
			Signature: blsSignatureBytes,
		},
	}

	// Debug logging before sending transaction
	log.Info("DEBUG: ICM Message Details",
		"avalancheNetworkID", avalancheNetworkID,
		"avalancheSourceBlockchainID", hex.EncodeToString(avalanchePChainBlockchainID[:]),
		"avalancheL1BlockchainID", hex.EncodeToString(avalancheL1BlockchainID[:]),
		"payloadLength", len(addressCallPayload.Bytes()),
		"payload", hex.EncodeToString(addressCallPayload.Bytes()),
		"unsignedMessageLength", len(unsignedMessage.Bytes()),
		"unsignedMessage", hex.EncodeToString(unsignedMessage.Bytes()),
		"signersLength", len(bitSetSignature.Signers),
		"signers", hex.EncodeToString(bitSetSignature.Signers),
		"signatureLength", len(bitSetSignature.Signature[:]),
		"signature", hex.EncodeToString(bitSetSignature.Signature[:]),
		"validatorsBytesLength", len(validatorsBytes),
		"validatorsBytes", hex.EncodeToString(validatorsBytes),
	)

	// Try to encode the transaction data to see what would be sent
	registryABI, abiErr := abi.JSON(strings.NewReader(avalanchevalidatorsetregistry.AvalancheValidatorSetRegistryABI))
	if abiErr == nil {
		txData, encodeErr := registryABI.Pack("registerValidatorSet", icmMessage, validatorsBytes)
		if encodeErr == nil {
			log.Info("DEBUG: Raw transaction calldata",
				"data", hex.EncodeToString(txData),
				"dataLength", len(txData),
			)
		} else {
			log.Error("Failed to encode transaction data", "error", encodeErr)
		}
	} else {
		log.Error("Failed to parse contract ABI", "error", abiErr)
	}

	// Test the verification before sending the transaction
	registryNetworkID, err := validatorSetRegistry.GetAvalancheNetworkID(nil)
	if err != nil {
		log.Error("Failed to get registry network ID", "error", err)
	} else {
		log.Info("DEBUG: Registry network ID", "registryNetworkID", registryNetworkID, "expectedNetworkID", avalancheNetworkID)
		if registryNetworkID != avalancheNetworkID {
			log.Error("Network ID mismatch", "registryNetworkID", registryNetworkID, "expectedNetworkID", avalancheNetworkID)
		}
	}

	// Parse the addressed call payload to debug its contents
	log.Info("DEBUG: AddressedCall payload details",
		"payloadBytes", hex.EncodeToString(addressCallPayload.Bytes()),
		"sourceAddressLength", len([]byte{}),
		"validatorSetStatePayloadBytes", hex.EncodeToString(validaterSetStatePayload.Bytes()),
	)

	// Additional debug: check the validator set hash that was computed vs what's expected
	computedHash := sha256.Sum256(validatorsBytes)
	log.Info("DEBUG: Validator set hash validation",
		"expectedHash", hex.EncodeToString(validatorSetHash[:]),
		"computedHash", hex.EncodeToString(computedHash[:]),
		"hashMatch", bytes.Equal(validatorSetHash[:], computedHash[:]),
	)

	// Debug the unsigned message structure
	log.Info("DEBUG: Unsigned message structure breakdown",
		"networkID", avalancheNetworkID,
		"sourceBlockchainID", hex.EncodeToString(avalanchePChainBlockchainID[:]),
		"targetBlockchainID", hex.EncodeToString(avalancheL1BlockchainID[:]),
		"isSameChain", bytes.Equal(avalanchePChainBlockchainID[:], avalancheL1BlockchainID[:]),
		"payloadSize", len(addressCallPayload.Bytes()),
	)

	// Comprehensive ICM message debugging
	DebugICMMessage(icmMessage, validatorsBytes, avalancheNetworkID)
	CompareICMMessageFields(icmMessage, avalancheNetworkID, avalanchePChainBlockchainID)

	// Try a static call first to simulate the transaction and see detailed error
	log.Info("DEBUG: Attempting static call simulation...")
	callOpts := &bind.CallOpts{
		From:    opts.From,
		Context: ctx,
	}

	// We can't directly simulate registerValidatorSet because it's not a view function
	// But we can at least validate the network ID and other accessible info
	currentValidatorSetID, err := validatorSetRegistry.NextValidatorSetID(callOpts)
	if err != nil {
		log.Error("Failed to get next validator set ID", "error", err)
	} else {
		log.Info("DEBUG: Next validator set ID", "nextID", currentValidatorSetID)
	}

	// Register the validator set in the registry.
	tx, err := validatorSetRegistry.RegisterValidatorSet(opts, icmMessage, validatorsBytes)
	if err != nil {
		log.Error("Failed to register validator set", "error", err)
		log.Info("Register validator set transaction failed", "error", err)

		// Additional debugging: examine the specific error message
		errorStr := err.Error()
		log.Info("DEBUG: Detailed error analysis", "errorString", errorStr)
		if strings.Contains(errorStr, "Invalid ICM message") {
			log.Error("DIAGNOSIS: ICM message validation failed - check networkID, sourceBlockchainID, signature verification, or weight quorum")
		}
		if strings.Contains(errorStr, "execution reverted") {
			log.Error("DIAGNOSIS: Transaction reverted - likely a require() statement failed in the contract")
		}
	} else if tx != nil {
		log.Info("Register validator set transaction", "tx", tx.Hash(), "txData", hex.EncodeToString(tx.Data()))
	} else {
		log.Error("Transaction is nil but no error returned")
	}
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined and get the receipt.
	txReceipt := WaitForTransactionSuccessWithClient(ctx, client, tx.Hash())
	log.Info("DEBUG: Transaction receipt", "txReceipt", txReceipt, "gasUsed", txReceipt.GasUsed)

	// Get the event from the logs, and confirm the validator set was registered.
	validatorSetRegisteredEvent, err := GetEventFromLogs(txReceipt.Logs, validatorSetRegistry.ParseValidatorSetRegistered)
	Expect(err).Should(BeNil())
	Expect(validatorSetRegisteredEvent.AvalancheBlockchainID).Should(BeEquivalentTo(avalancheL1BlockchainID))

	// Return the validator set ID.
	return validatorSetRegisteredEvent.ValidatorSetID
}
