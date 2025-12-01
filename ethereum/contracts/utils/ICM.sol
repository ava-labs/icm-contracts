// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {WarpMessage} from "@subnet-evm/IWarpMessenger.sol";
import {
    ICMMessage,
    ICMUnsignedMessage,
    ICMSignature
} from "@avalabs/avalanche-contracts/teleporter/ITeleporterMessenger.sol";
import {Validator, ValidatorSet} from "../utils/ValidatorSets.sol";
import {BLST} from "./BLST.sol";
import {ByteSlicer} from "./ByteSlicer.sol";

struct AddressedCall {
    bytes sourceAddress;
    bytes payload;
}

/**
 * @title ICM
 * @notice Utility library for Interchain Messaging (ICM) operations
 * @dev This library provides helper functions for working with ICM signatures and validation
 */
library ICM {
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
        uint256 resultSize = data.length * 8 - leadingZeros;

        // Create result array with the exact size needed
        bool[] memory result = new bool[](resultSize);

        // Fill the result array directly, starting from the first significant bit
        uint256 resultIndex = 0;

        // Process all bytes, skipping leading zeros in the first byte
        for (uint256 j = leadingZeros; j < 8; ++j) {
            result[resultIndex] = (uint8(data[0]) & (1 << (7 - j))) != 0;
            resultIndex++;
        }
        for (uint256 i = 1; i < data.length; ++i) {
            for (uint256 j = 0; j < 8; ++j) {
                result[resultIndex] = (uint8(data[i]) & (1 << (7 - j))) != 0;
                resultIndex++;
            }
        }

        return result;
    }

    function parseICMMessage(
        bytes memory data
    ) internal pure returns (ICMMessage memory) {
        ICMUnsignedMessage memory unsignedMessage = parseICMUnsignedMessage(data);
        uint32 payloadLength = uint32(unsignedMessage.payload.length);
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
            unsignedMessage: unsignedMessage,
            unsignedMessageBytes: unsignedMessageBytes,
            signature: ICMSignature({signers: bitSet, signature: signature})
        });
    }

    function parseICMUnsignedMessage(
        bytes memory data
    ) internal pure returns (ICMUnsignedMessage memory) {
        // Validate the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the avalancheNetworkID
        uint32 avalancheNetworkID = uint32(bytes4(ByteSlicer.slice(data, 2, 4)));

        // Parse the avalancheSourceBlockchainID
        bytes32 avalancheSourceBlockchainID = abi.decode(ByteSlicer.slice(data, 6, 32), (bytes32));

        // Parse the payload
        uint32 payloadLength = uint32(bytes4(ByteSlicer.slice(data, 38, 4)));
        bytes memory payload = ByteSlicer.slice(data, 42, payloadLength);
        return ICMUnsignedMessage({
            avalancheNetworkID: avalancheNetworkID,
            avalancheSourceBlockchainID: avalancheSourceBlockchainID,
            payload: payload
        });
    }

    /**
     * @notice Serialize an ICM unsigned message to bytes.
     */
    function serializeICMUnsignedMessage(
        ICMUnsignedMessage memory message
    ) internal pure returns (bytes memory) {
        bytes memory serialized = new bytes(
            // the codec
            2
            // the avalancheNetworkID
            + 4
            // the avalancheSourceBlockchainID
            + 32
            // the payload length
            + 4
            // the message payload
            + message.payload.length
        );
        // the codec
        serialized[0] = 0x00;
        serialized[1] = 0x00;
        // encode the avalancheNetworkID
        bytes4 avalancheNetworkID = bytes4(message.avalancheNetworkID);
        for (uint256 i = 0; i < 4; i++) {
            serialized[2 + i] = avalancheNetworkID[i];
        }
        // encode the avalancheSourceBlockchainID
        for (uint256 i = 0; i < 32; i++) {
            serialized[6 + i] = message.avalancheSourceBlockchainID[i];
        }
        // encode the payload length
        bytes4 payloadLength = bytes4(uint32(message.payload.length));
        for (uint256 i = 0; i < 4; i++) {
            serialized[38 + i] = payloadLength[i];
        }
        // encode the payload
        for (uint256 i = 0; i < message.payload.length; i++) {
            serialized[42 + i] = message.payload[i];
        }

        return serialized;
    }

    /**
     * @notice Serialize an ICM message to bytes
     */
    function serializeICMMessage(
        ICMMessage memory message
    ) internal pure returns (bytes memory) {

        bytes memory data = new bytes(
            // unsigned message length
            message.unsignedMessageBytes.length
            + 4
            // to encode the length of the serialize bitset
            + 4
            // the length of the serialized bitset
            + message.signature.signers.length
            // the signature
            + message.signature.signature.length
        );
        uint256 cursor = 0;
        // add the unsigned message bytes
        (data, cursor) = ByteSlicer.extendFromSlice(data, message.unsignedMessageBytes, cursor);

        (data, cursor) =  ByteSlicer.extendFromSlice(data, new bytes(4), cursor);

        // add the length of the signers (as a bitset)
        bytes4 bitsetLength = bytes4(uint32(message.signature.signers.length));
        for (uint256 i = 0; i < 4; i++) {
            data[cursor + i] = bitsetLength[i];
        }
        cursor += 4;

        // add the signers bitset
        (data, cursor) = ByteSlicer.extendFromSlice(data, message.signature.signers, cursor);
        // add the signature bytes
        (data, cursor) = ByteSlicer.extendFromSlice(data, message.signature.signature, cursor);

        return data;
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

    function serializeAddressedCall(
        AddressedCall memory addressedCall
    ) internal pure returns (bytes memory) {
        bytes memory serialized = new bytes(
            // codec bytes
            2
            // payload type ID
            + 4
            // length of the source address
            + 4
            // the source address
            + addressedCall.sourceAddress.length
            // length of the payload
            + 4
            // the payload
            + addressedCall.payload.length
        );
        // encode the codec
        serialized[0] = 0x00;
        serialized[1] = 0x00;
        // encode the payload type ID
        serialized[2] = 0x00;
        serialized[3] = 0x00;
        serialized[4] = 0x00;
        serialized[5] = 0x01;
        // encode the source address length
        bytes4 sourceAddrLength = bytes4(uint32(addressedCall.sourceAddress.length));
        for (uint256 i = 0; i < 4; i++) {
            serialized[6 + i] = sourceAddrLength[i];
        }
        // encode the the source address
        for (uint256 i = 0; i < addressedCall.sourceAddress.length; i++) {
            serialized[10 + i] = addressedCall.sourceAddress[i];
        }
        // encode the payload length
        bytes4 payloadLength = bytes4(uint32(addressedCall.payload.length));
        for (uint256 i = 0; i < 4; i++) {
            serialized[10 + addressedCall.sourceAddress.length + i] = payloadLength[i];
        }
        // encode the payload
        for (uint256 i = 0; i < addressedCall.payload.length; i++) {
            serialized[14 + addressedCall.sourceAddress.length + i] = addressedCall.payload[i];
        }

        return serialized;
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
                    aggregatePublicKey = BLST.addG1(aggregatePublicKey, validator.blsPublicKey);
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
        bytes32 avalancheBlockChainID,
        ValidatorSet memory validatorSet
    ) internal view  {
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


         if (
            avalancheBlockChainID != validatorSet.avalancheBlockchainID
         ) {
             revert("Invalid avalanche source blockchain ID");
         }

        bool[] memory signers = bytesToBoolArray(message.signature.signers);
        (bytes memory aggregatePublicKey, uint64 aggregateWeight) =
            filterValidators(signers, validatorSet.validators);

        if (!verifyWeight(aggregateWeight, validatorSet.totalWeight)) {
            revert("Insufficient weight");
        }

        bool result = BLST.verifySignature(
            aggregatePublicKey, message.signature.signature, message.unsignedMessageBytes
        );
        if (!result) {
            revert("Invalid signature");
        }
    }

    /**
     * @notice Convert the unsigned part of an ICM message to Warp message. This should
     * be called after an ICM message has been verified.
     * @dev This function should only be called when communicating with an external
     * chain. If used during internal interop, the teleport contracts will reject the
     * `originSenderAddress`.
     *
     * This function loosely reproduces the logic of the handleMessage method used in the
     * Warp precompile.
     * @param message The unsigned part of an ICM message
     * @return A Warp message
     */
    function handleMessage(ICMUnsignedMessage memory message ) internal pure returns (WarpMessage memory) {
        AddressedCall memory addressedCall = parseAddressedCall(message.payload);
        address sourceAddress;
        bytes memory sourceAddressBytes = addressedCall.sourceAddress;
        assembly {
            sourceAddress := mload(add(sourceAddressBytes, 20))
        }
        return WarpMessage({
            sourceChainID: message.avalancheSourceBlockchainID,
            originSenderAddress: sourceAddress,
            payload: addressedCall.payload
        });
    }
}
