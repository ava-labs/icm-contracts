package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	avalanchevalidatorsetregistry "github.com/ava-labs/icm-contracts/abi-bindings/go/AvalancheValidatorSetRegistry"
	"github.com/ava-labs/libevm/log"
)

// DebugICMMessage provides detailed analysis of an ICM message for debugging validation issues
func DebugICMMessage(
	icmMessage avalanchevalidatorsetregistry.ICMMessage,
	validatorsBytes []byte,
	expectedNetworkID uint32,
) {
	log.Info("=== ICM MESSAGE DEBUG ANALYSIS ===")

	// 1. Basic message structure
	log.Info("1. Message Structure",
		"networkID", icmMessage.UnsignedMessage.AvalancheNetworkID,
		"sourceBlockchainID", hex.EncodeToString(icmMessage.UnsignedMessage.AvalancheSourceBlockchainID[:]),
		"payloadLength", len(icmMessage.UnsignedMessage.Payload),
	)

	// 2. Network ID validation
	networkIDValid := icmMessage.UnsignedMessage.AvalancheNetworkID == expectedNetworkID
	log.Info("2. Network ID Validation",
		"expected", expectedNetworkID,
		"actual", icmMessage.UnsignedMessage.AvalancheNetworkID,
		"valid", networkIDValid,
	)

	// 3. Try to parse the payload as AddressedCall
	addressedCall, err := payload.ParseAddressedCall(icmMessage.UnsignedMessage.Payload)
	if err != nil {
		log.Error("3. AddressedCall Parsing FAILED", "error", err)
		return
	}

	sourceAddressValid := len(addressedCall.SourceAddress) == 0
	log.Info("3. AddressedCall Structure",
		"sourceAddressLength", len(addressedCall.SourceAddress),
		"sourceAddressValid", sourceAddressValid,
		"innerPayloadLength", len(addressedCall.Payload),
	)

	// 4. Validator set hash verification
	computedHash := sha256.Sum256(validatorsBytes)
	log.Info("4. Validator Set Hash",
		"validatorBytesLength", len(validatorsBytes),
		"computedHash", hex.EncodeToString(computedHash[:]),
	)

	// 5. Signature analysis
	log.Info("5. Signature Analysis",
		"signersLength", len(icmMessage.Signature.Signers),
		"signers", hex.EncodeToString(icmMessage.Signature.Signers),
		"signatureLength", len(icmMessage.Signature.Signature),
		"signature", hex.EncodeToString(icmMessage.Signature.Signature),
	)

	// 6. Try to parse the full signed message
	unsignedMessage, err := warp.ParseUnsignedMessage(icmMessage.UnsignedMessageBytes)
	if err != nil {
		log.Error("6. Signed Message Parsing FAILED", "error", err)
	} else {
		log.Info("6. Signed Message Structure",
			"messageID", unsignedMessage.ID().String(),
			"networkID", unsignedMessage.NetworkID,
			"sourceChainID", unsignedMessage.SourceChainID.String(),
		)
	}

	log.Info("=== END ICM MESSAGE DEBUG ANALYSIS ===")
}

// CompareICMMessageFields helps identify specific field mismatches
func CompareICMMessageFields(
	icmMessage avalanchevalidatorsetregistry.ICMMessage,
	expectedNetworkID uint32,
	expectedSourceBlockchainID [32]byte,
) {
	log.Info("=== ICM MESSAGE FIELD COMPARISON ===")

	// Network ID comparison
	if icmMessage.UnsignedMessage.AvalancheNetworkID != expectedNetworkID {
		log.Error("❌ Network ID Mismatch",
			"expected", expectedNetworkID,
			"actual", icmMessage.UnsignedMessage.AvalancheNetworkID,
		)
	} else {
		log.Info("✅ Network ID Match", "networkID", expectedNetworkID)
	}

	// Blockchain ID comparison
	if icmMessage.UnsignedMessage.AvalancheSourceBlockchainID != expectedSourceBlockchainID {
		log.Error("❌ Source Blockchain ID Mismatch",
			"expected", hex.EncodeToString(expectedSourceBlockchainID[:]),
			"actual", hex.EncodeToString(icmMessage.UnsignedMessage.AvalancheSourceBlockchainID[:]),
		)
	} else {
		log.Info("✅ Source Blockchain ID Match",
			"blockchainID", hex.EncodeToString(expectedSourceBlockchainID[:]),
		)
	}

	log.Info("=== END FIELD COMPARISON ===")
}
