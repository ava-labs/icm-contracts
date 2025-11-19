package utils

import (
	"encoding/hex"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	avalanchevalidatorsetregistry "github.com/ava-labs/icm-contracts/abi-bindings/go/AvalancheValidatorSetRegistry"
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	"github.com/ava-labs/libevm/log"
	. "github.com/onsi/gomega"
)

func PrepareICMMessageForEthereum(
	avalancheNetworkID uint32,
	avalanchePChainBlockchainID ids.ID,
	avalancheSubnetID ids.ID,
	avalancheL1BlockchainID ids.ID,
	signatureAggregator *SignatureAggregator,
	unsignedMessage *warp.UnsignedMessage,
) teleportermessenger.ICMMessage {

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
	icmMessage := teleportermessenger.ICMMessage{
		UnsignedMessage: teleportermessenger.ICMUnsignedMessage{
			AvalancheNetworkID:          avalancheNetworkID,
			AvalancheSourceBlockchainID: avalancheL1BlockchainID,
			Payload:                     unsignedMessage.Payload,
		},
		UnsignedMessageBytes: unsignedMessage.Bytes(),
		Signature: teleportermessenger.ICMSignature{
			Signers:   bitSetSignature.Signers,
			Signature: blsSignatureBytes,
		},
	}

	// Debug logging before sending transaction
	log.Debug("DEBUG: ICM Message Details",
		"avalancheNetworkID", avalancheNetworkID,
		"avalancheSourceBlockchainID", hex.EncodeToString(avalanchePChainBlockchainID[:]),
		"avalancheL1BlockchainID", hex.EncodeToString(avalancheL1BlockchainID[:]),
		"payloadLength", len(unsignedMessage.Payload),
		"payload", hex.EncodeToString(unsignedMessage.Payload),
		"unsignedMessageLength", len(unsignedMessage.Bytes()),
		"signersLength", len(bitSetSignature.Signers),
		"signers", hex.EncodeToString(bitSetSignature.Signers),
		"signatureLength", len(bitSetSignature.Signature[:]),
		"signature", hex.EncodeToString(bitSetSignature.Signature[:]),
	)
	return icmMessage
}

func into(msg teleportermessenger.ICMMessage) avalanchevalidatorsetregistry.ICMMessage {
	return avalanchevalidatorsetregistry.ICMMessage{
		UnsignedMessage: avalanchevalidatorsetregistry.ICMUnsignedMessage{
			AvalancheNetworkID:          msg.UnsignedMessage.AvalancheNetworkID,
			AvalancheSourceBlockchainID: msg.UnsignedMessage.AvalancheSourceBlockchainID,
			Payload:                     msg.UnsignedMessage.Payload,
		},
		UnsignedMessageBytes: msg.UnsignedMessageBytes,
		Signature: avalanchevalidatorsetregistry.ICMSignature{
			Signers:   msg.Signature.Signers,
			Signature: msg.Signature.Signature,
		},
	}
}
