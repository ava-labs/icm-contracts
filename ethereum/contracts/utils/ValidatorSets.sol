// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ByteSlicer} from "./ByteSlicer.sol";
import {BLST} from "./BLST.sol";
import {ByteComparator} from "./ByteComparator.sol";

struct Validator {
    bytes blsPublicKey;
    uint64 weight;
}

struct ValidatorSet {
    bytes32 avalancheBlockchainID;
    Validator[] validators;
    uint64 totalWeight;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
}

struct ValidatorSetStatePayload {
    bytes32 avalancheBlockchainID;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
    bytes32 validatorSetHash;
}

library ValidatorSets {
    /**
     * @notice Parses the validators from a serialized validator set
     * @param data The serialized validator set. The serialized format is:
     * - 2 bytes: codec ID (0x0000)
     * - 4 bytes: number of validators
     * - 96 bytes per validator: unformatted BLS public key
     * - 8 bytes per validator: weight
     * @return validators The parsed validators
     * @return totalWeight The total weight of all validators
     */
    function parseValidators(
        bytes memory data
    ) internal pure returns (Validator[] memory, uint64) {
        // Check the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the number of validators
        uint32 numValidators = uint32(bytes4(ByteSlicer.slice(data, 2, 4)));

        // Parse the validators
        Validator[] memory validators = new Validator[](numValidators);
        uint64 totalWeight = 0;
        uint256 offset = 6;
        bytes memory previousPublicKey = new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);
        for (uint32 i = 0; i < numValidators; i++) {
            bytes memory unformattedPublicKey = ByteSlicer.slice(data, offset, 96);
            require(
                ByteComparator.compare(unformattedPublicKey, previousPublicKey) > 0,
                "BLS public key must be greater than the latest public key"
            );
            uint64 weight = uint64(bytes8(ByteSlicer.slice(data, offset + 96, 8)));
            require(weight > 0, "Validator weight must be greater than 0");
            validators[i] = Validator({
                blsPublicKey: BLST.formatUncompressedBLSPublicKey(unformattedPublicKey),
                weight: weight
            });
            totalWeight += weight;
            offset += 104;
        }
        return (validators, totalWeight);
    }

    function serializeValidators(
        Validator[] memory validators
    ) public pure returns (bytes memory) {
        bytes memory serialized = new bytes(
            2 + 4 + validators.length * (96 + 8)
        );
        // the codec
        serialized[0] = 0x00;
        serialized[1] = 0x00;
        //encode the number of validators
        bytes4 numValidators = bytes4(uint32(validators.length));
        for (uint256 i = 0; i < 4; i++) {
            serialized[2 + i] = numValidators[i];
        }
        // encode the validators
        uint256 offset = 6;
        for (uint256 i = 0; i < validators.length; i++) {
            // encode the 96-bytes uncompressed BLS public key
            bytes memory uncompressedBlsPublicKey = BLST.getUncompressedBlsPublicKey(validators[i].blsPublicKey);
            for(uint256 j = 0; j < 96; j++ ) {
                serialized[j + offset] = uncompressedBlsPublicKey[j];
            }
            offset += 96;
            // encode the validator weight
            bytes8 weight = bytes8(validators[i].weight);
            for(uint256 j = 0; j < 8; j++ ) {
                serialized[j + offset] = weight[j];
            }
            offset += 8;
        }
        return serialized;
    }

    function parseValidatorSetStatePayload(
        bytes memory data
    ) internal pure returns (ValidatorSetStatePayload memory) {
        // Check the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the payload type ID, and confirm it is 4 for ValidatorSetState
        uint32 payloadTypeID = uint32(bytes4(ByteSlicer.slice(data, 2, 4)));
        require(payloadTypeID == 4, "Invalid ValidatorSetState payload type ID");

        // Parse the avalancheBlockchainID
        bytes32 avalancheBlockchainID = abi.decode(ByteSlicer.slice(data, 6, 32), (bytes32));

        // Parse the pChainHeight
        uint64 pChainHeight = uint64(bytes8(ByteSlicer.slice(data, 38, 8)));

        // Parse the pChainTimestamp
        uint64 pChainTimestamp = uint64(bytes8(ByteSlicer.slice(data, 46, 8)));

        // Parse the validatorSetHash
        bytes32 validatorSetHash = abi.decode(ByteSlicer.slice(data, 54, 32), (bytes32));

        return ValidatorSetStatePayload({
            avalancheBlockchainID: avalancheBlockchainID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            validatorSetHash: validatorSetHash
        });
    }

    function serializeValidatorSetStatePayload(
        ValidatorSetStatePayload memory payload
    ) public pure returns (bytes memory) {
        bytes memory serialized = new bytes(
            // codec
            2
            // payload type ID
            + 4
            // Avalanche block chain ID
            + 32
            // the P-chain height
            + 8
            // the P-chain timestamp
            + 8
            // hash of validator set
            + 32
        );
        // encode the codec
        serialized[0] = 0x00;
        serialized[1] = 0x00;
        // the payload type ID
        serialized[2] = 0x00;
        serialized[3] = 0x00;
        serialized[4] = 0x00;
        serialized[5] = 0x04;
        // encode avalancheBlockchainID
        for (uint256 i = 0; i < 32; i++) {
            serialized[6 + i] = payload.avalancheBlockchainID[i];
        }
        // encode the P-chain block height
        bytes8 pChainHeight = bytes8(payload.pChainHeight);
        for (uint256 i = 0; i < 8; i++) {
            serialized[38 + i] = pChainHeight[i];
        }
        // encode the P-chain timestamp
        bytes8 pChainTimestamp = bytes8(payload.pChainTimestamp);
        for (uint256 i = 0; i < 8; i++) {
            serialized[46 + i] = pChainTimestamp[i];
        }
        // encode the validator set hash
        for (uint256 i = 0; i < 32; i++) {
            serialized[54 + i] = payload.validatorSetHash[i];
        }
        return serialized;
    }
}
