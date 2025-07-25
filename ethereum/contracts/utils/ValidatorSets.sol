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
}
