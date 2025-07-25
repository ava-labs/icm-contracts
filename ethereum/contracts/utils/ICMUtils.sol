// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Validator, ValidatorSet} from "../utils/ValidatorSetUtils.sol";
import {BLSTUtils} from "./BLSTUtils.sol";
import {ByteSlicer} from "./ByteSlicer.sol";

struct ICMSignature {
    bytes signers;
    bytes signature;
}

struct ICMUnsignedMessage {
    uint32 avalancheNetworkID;
    bytes32 avalancheSourceBlockchainID;
    bytes payload;
}

struct ICMMessage {
    ICMUnsignedMessage unsignedMessage;
    bytes unsignedMessageBytes;
    ICMSignature signature;
}

struct AddressedCall {
    bytes sourceAddress;
    bytes payload;
}

/**
 * @title ICMUtils
 * @notice Utility library for Interchain Messaging (ICM) operations
 * @dev This library provides helper functions for working with ICM signatures and validation
 */
library ICMUtils {
    uint256 constant QUORUM_NUM = 67;
    uint256 constant QUORUM_DEN = 100;

    /**
     * @dev Converts a bytes array to a bool[] array
     * @param data The bytes array to convert
     * @return A bool[] array where each element is true if the corresponding bit is 1, false otherwise
     */
    function bytesToBoolArray(
        bytes memory data
    ) internal pure returns (bool[] memory) {
        require(data.length > 0 && data[0] != 0x00, "Invalid bit set");

        // Find the position of the most significant bit in the first byte
        // This determines how many leading zeros we have
        uint8 firstByte = uint8(data[0]);
        uint256 leadingZeros = 0;
        uint8 mask = 0x80; // Start with 10000000

        while ((firstByte & mask) == 0) {
            leadingZeros++;
            mask >>= 1;
        }

        // Calculate the final result size (total bits minus leading zeros)
        uint256 totalBits = data.length * 8;
        uint256 resultSize = totalBits - leadingZeros;

        // Create result array with the exact size needed
        bool[] memory result = new bool[](resultSize);

        // Fill the result array directly, starting from the first significant bit
        uint256 resultIndex = 0;

        // Process all bytes, skipping leading zeros in the first byte
        for (uint256 i = 0; i < data.length; ++i) {
            uint256 startJ = (i == 0) ? leadingZeros : 0;
            for (uint256 j = startJ; j < 8; ++j) {
                result[resultIndex] = (uint8(data[i]) & (1 << (7 - j))) != 0;
                resultIndex++;
            }
        }

        return result;
    }

    function parseICMMessage(
        bytes memory data
    ) internal pure returns (ICMMessage memory) {
        // Validate the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the avalancheNetworkID
        uint32 avalancheNetworkID = uint32(bytes4(ByteSlicer.slice(data, 2, 4)));

        // Parse the avalancheSourceBlockchainID
        bytes32 avalancheSourceBlockchainID = abi.decode(ByteSlicer.slice(data, 6, 32), (bytes32));

        // Parse the payload
        uint32 payloadLength = uint32(bytes4(ByteSlicer.slice(data, 38, 4)));
        bytes memory payload = ByteSlicer.slice(data, 42, payloadLength);

        // Parse the unsigned message bytes
        bytes memory unsignedMessageBytes = ByteSlicer.slice(data, 0, 42 + payloadLength);

        // Check that the signature type ID is 0
        uint32 signatureType = uint32(bytes4(ByteSlicer.slice(data, 42 + payloadLength, 4)));
        require(signatureType == 0, "Invalid signature type ID");

        // Parse the bit set length
        uint32 bitSetLength = uint32(bytes4(ByteSlicer.slice(data, 46 + payloadLength, 4)));

        // Parse the bit set
        bytes memory bitSet = ByteSlicer.slice(data, 50 + payloadLength, bitSetLength);

        // Check that there are exactly 96 bytes remaining, and parse them as the signature
        require(data.length == 50 + payloadLength + bitSetLength + 192, "Invalid signature length");
        bytes memory signature = ByteSlicer.slice(data, 50 + payloadLength + bitSetLength, 192);

        return ICMMessage({
            unsignedMessage: ICMUnsignedMessage({
                avalancheNetworkID: avalancheNetworkID,
                avalancheSourceBlockchainID: avalancheSourceBlockchainID,
                payload: payload
            }),
            unsignedMessageBytes: unsignedMessageBytes,
            signature: ICMSignature({signers: bitSet, signature: signature})
        });
    }

    function parseAddressedCall(
        bytes memory data
    ) internal pure returns (AddressedCall memory) {
        // Validate the codec ID is 0.
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the payload type ID, and confirm it is 1 for AddressedCall
        uint32 payloadTypeID = uint32(bytes4(ByteSlicer.slice(data, 2, 4)));
        require(payloadTypeID == 1, "Invalid payload type ID");

        // Parse the source address length
        uint32 sourceAddressLength = uint32(bytes4(ByteSlicer.slice(data, 6, 4)));

        // Parse the source address
        bytes memory sourceAddress = ByteSlicer.slice(data, 10, sourceAddressLength);

        // Parse the payload length
        uint32 payloadLength = uint32(bytes4(ByteSlicer.slice(data, 10 + sourceAddressLength, 4)));

        // Parse the payload
        bytes memory payload = ByteSlicer.slice(data, 14 + sourceAddressLength, payloadLength);

        return AddressedCall({sourceAddress: sourceAddress, payload: payload});
    }

    function filterValidators(
        bool[] memory signers,
        Validator[] memory validators
    ) internal view returns (bytes memory, uint64) {
        require(
            signers.length <= validators.length,
            "Signers length must be less than or equal to validators length"
        );
        bytes memory aggregatePublicKey;
        uint64 aggregateWeight = 0;
        for (uint256 i = 0; i < validators.length; ++i) {
            if (i < signers.length && signers[signers.length - i - 1]) {
                // Cache the validator to avoid repeated array access
                Validator memory validator = validators[i];

                if (aggregateWeight > 0) {
                    aggregatePublicKey = BLSTUtils.addG1(aggregatePublicKey, validator.blsPublicKey);
                    aggregateWeight += validator.weight;
                } else {
                    aggregatePublicKey = validator.blsPublicKey;
                    aggregateWeight = validator.weight;
                }
            }
        }
        return (aggregatePublicKey, aggregateWeight);
    }

    // Verifies that quorumNum * totalWeight <= quorumDen * signatureWeight
    function verifyWeight(
        uint64 signatureWeight,
        uint64 totalWeight
    ) internal pure returns (bool) {
        uint256 scaledTotalWeight = QUORUM_NUM * uint256(totalWeight);
        uint256 scaledSignatureWeight = QUORUM_DEN * uint256(signatureWeight);
        return scaledTotalWeight <= scaledSignatureWeight;
    }

    function verifyICMMessage(
        ICMMessage memory message,
        uint32 avalancheNetworkID,
        ValidatorSet memory validatorSet
    ) internal view returns (bool) {
        if (message.unsignedMessage.avalancheNetworkID != avalancheNetworkID) {
            revert("Invalid avalanche network ID");
        }

        // TODO: Do we need to check the avalanche source blockchain ID?
        // It's expected to be different in cases where the message is from the primary network.
        // if (
        //     message.unsignedMessage.avalancheSourceBlockchainID
        //         != validatorSet.avalancheBlockchainID
        // ) {
        //     revert("Invalid avalanche source blockchain ID");
        // }

        bool[] memory signers = bytesToBoolArray(message.signature.signers);
        (bytes memory aggregatePublicKey, uint64 aggregateWeight) =
            filterValidators(signers, validatorSet.validators);

        if (!verifyWeight(aggregateWeight, validatorSet.totalWeight)) {
            revert("Invalid weight");
        }

        bool result = BLSTUtils.verifySignature(
            aggregatePublicKey, message.signature.signature, message.unsignedMessageBytes
        );
        if (!result) {
            revert("Invalid signature");
        }
        return result;
    }
}
