// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IAvalancheValidatorSetRegistry} from "./interfaces/IAvalancheValidatorSetRegistry.sol";
import {
    Validator,
    ValidatorSet,
    ValidatorSetStatePayload,
    ValidatorSets
} from "./utils/ValidatorSets.sol";
import {
    WarpMessage,
    WarpBlockHash
} from "@subnet-evm/IWarpMessenger.sol";
import {ICMMessage} from "@avalabs/avalanche-contracts/teleporter/ITeleporterMessenger.sol";
import {ICM, AddressedCall} from "./utils/ICM.sol";
import {IWarpExt} from "@avalabs/avalanche-contracts/teleporter/IWarpExt.sol";

/**
 * @title AvalancheValidatorSetRegistry
 * @notice Registry for managing Avalanche validator sets
 * @dev This contract allows registration and updates of validator sets for Avalanche blockchains.
 * Updates are authenticated through signed ICM messages from the current validator set.
 */
contract AvalancheValidatorSetRegistry is
    IAvalancheValidatorSetRegistry,
    IWarpExt
{
    uint32 public immutable avalancheNetworkID;
    uint32 public nextValidatorSetID = 0;
    /**
     * @notice The chain ID of the Ethereum network the contract is deployed on.
     * @dev The chain ID for Ethereum is a uint which we reinterpret as bytes32
     * to remain consistent with the existing interface
     */
    uint256 public ethereumBlockchainID;


    // Mapping of validator set IDs to their complete validator set data.
    mapping(uint256 => ValidatorSet) private _validatorSets;

    constructor(uint32 _avalancheNetworkID, uint256 _ethereumBlockchainID) {
        avalancheNetworkID = _avalancheNetworkID;
        ethereumBlockchainID = _ethereumBlockchainID;
    }

    /**
     * @notice Registers a new validator set
     * @dev A validator set can be registered by anyone. The correctness should be verified
     * with the actual validator set on the Avalanche P-Chain.
     * @param message The ICM message containing the validator set to register. The message must be signed by validator set
     * @param validatorBytes The serialized validator set to register.
     * @return The ID of the registered validator set
     */
    function registerValidatorSet(
        ICMMessage calldata message,
        bytes memory validatorBytes
    ) external override returns (uint256) {
        // Parse and validate the validator set data
        (
            ValidatorSetStatePayload memory validatorSetStatePayload,
            Validator[] memory validators,
            uint64 totalWeight
        ) = _parseAndValidateValidatorSetData(message, validatorBytes);

        // Construct the validator set and confirm the ICM was self-signed by it.
        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: validatorSetStatePayload.avalancheBlockchainID,
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: validatorSetStatePayload.pChainHeight,
            pChainTimestamp: validatorSetStatePayload.pChainTimestamp
        });

        ICM.verifyICMMessage(message, avalancheNetworkID,validatorSet);

        // Store the validator set.
        uint256 validatorSetID = nextValidatorSetID++;
        _validatorSets[validatorSetID] = validatorSet;

        emit ValidatorSetRegistered(validatorSetID, validatorSet.avalancheBlockchainID);
        return validatorSetID;
    }

    /**
     * @notice Updates a validator set
     * @dev Updates are authenticated by a signed ICM message from the current validator set
     * @param validatorSetID The ID of the validator set to update
     * @param message The ICM message containing the update
     */
    function updateValidatorSet(
        uint256 validatorSetID,
        ICMMessage calldata message,
        bytes memory validatorBytes
    ) external override {
        require(validatorSetID < nextValidatorSetID, "Validator set does not exist");

        ValidatorSet storage currentValidatorSet = _validatorSets[validatorSetID];

        // Verify the ICM message using the current validator set
        ICM.verifyICMMessage(message, avalancheNetworkID,currentValidatorSet);

        // Parse and validate the validator set data
        (
            ValidatorSetStatePayload memory validatorSetStatePayload,
            Validator[] memory validators,
            uint64 totalWeight
        ) = _parseAndValidateValidatorSetData(message, validatorBytes);

        // Check that blockchain ID matches the current validator set.
        require(
            validatorSetStatePayload.avalancheBlockchainID
                == currentValidatorSet.avalancheBlockchainID,
            "Blockchain ID mismatch"
        );

        // Check that the pChain height is greater than the current validator set.
        require(
            validatorSetStatePayload.pChainHeight > currentValidatorSet.pChainHeight,
            "P-Chain height must be greater than the current validator set"
        );

        // Check that the pChain timestamp is greater than the current validator set.
        require(
            validatorSetStatePayload.pChainTimestamp > currentValidatorSet.pChainTimestamp,
            "P-Chain timestamp must be greater than the current validator set"
        );

        // Update the validator set
        _validatorSets[validatorSetID] = ValidatorSet({
            avalancheBlockchainID: validatorSetStatePayload.avalancheBlockchainID,
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: validatorSetStatePayload.pChainHeight,
            pChainTimestamp: validatorSetStatePayload.pChainTimestamp
        });

        emit ValidatorSetUpdated(validatorSetID, validatorSetStatePayload.avalancheBlockchainID);
    }


    /**
     * @notice Gets a validator set by its ID
     * @param validatorSetID The ID of the validator set to get
     * @return The validator set
     */
    function getValidatorSet(
        uint256 validatorSetID
    ) external view override returns (ValidatorSet memory) {
        return _getValidatorSet(validatorSetID);
    }

    /**
     * @notice Gets the Avalanche network ID
     * @return The Avalanche network ID
     */
    function getAvalancheNetworkID() external view returns (uint32) {
        return avalancheNetworkID;
    }

    /**
     * @notice Verifies an ICM message against a validator set or reverts
     * @dev This function validates that the message is properly signed by a sufficient quorum of validators
     * from the validator set identified by validatorSetID. The verification includes checking the network ID,
     * blockchain ID, and cryptographic signature verification.
     * @param validatorSetID The ID of the validator set to verify the message against
     * @param message The ICM message to verify
     */
    function verifyICMMessage(
        uint256 validatorSetID,
        ICMMessage calldata message
    ) public view override {
        ValidatorSet memory validatorSet = _getValidatorSet(validatorSetID);
        ICM.verifyICMMessage(message, avalancheNetworkID, validatorSet);
    }

    /**
     * * @notice Verifies an ICM message against a validator set or reverts.
     * @param validatorSetID The ID of the validator set to verify the message against
     * @param icmMessage The ICM message to verify
     * @return The derived message that TeleporterMessenger expects
     */
    function getVerifiedICMMessage(
        uint256 validatorSetID,
        ICMMessage calldata icmMessage
    ) external view returns (WarpMessage memory) {
        verifyICMMessage(validatorSetID, icmMessage);
        return ICM.handleMessage(icmMessage.unsignedMessage);
    }

    function sendWarpMessage(bytes calldata payload) external returns (bytes32) {
        revert("Sending Warp messages from Ethereum is not currently supported");
    }

    function getVerifiedWarpMessage(uint32 index) external pure returns (WarpMessage calldata, bool) {
        revert("This method can't be called on Ethereum, use `getVerifiedICMMessage` instead");
    }

    function getVerifiedWarpBlockHash(
        uint32 index
    ) external pure returns (WarpBlockHash calldata, bool) {
        revert("This method cannot be called on Ethereum");
    }

    function getBlockchainID() external view returns (bytes32) {
        return bytes32(ethereumBlockchainID);
    }

    function _getValidatorSet(
        uint256 validatorSetID
    ) private view returns (ValidatorSet memory) {
        require(validatorSetID < nextValidatorSetID, "Validator set does not exist");
        return _validatorSets[validatorSetID];
    }

    /**
     * @notice Parses and validates validator set data from an ICM message
     * @dev This is a helper function that consolidates common validation logic
     * @param message The ICM message containing the validator set data
     * @param validatorBytes The serialized validator set
     * @return validatorSetStatePayload The parsed validator set state payload
     * @return validators The parsed validators array
     * @return totalWeight The total weight of all validators
     */
    function _parseAndValidateValidatorSetData(
        ICMMessage calldata message,
        bytes memory validatorBytes
    )
    private
    pure
    returns (
        ValidatorSetStatePayload memory validatorSetStatePayload,
        Validator[] memory validators,
        uint64 totalWeight
    )
    {
        // Parse the addressed call and validate that the source address is empty.
        AddressedCall memory addressedCall = ICM.parseAddressedCall(message.unsignedMessage.payload);
        require(addressedCall.sourceAddress.length == 0, "Source address must be empty");

        // Parse the validator set state payload.
        validatorSetStatePayload =
            ValidatorSets.parseValidatorSetStatePayload(addressedCall.payload);

        // Check that the validator set hash matches the hash of the serialized validator set.
        require(
            validatorSetStatePayload.validatorSetHash == sha256(validatorBytes),
            "Validator set hash mismatch"
        );

        // Parse the validators.
        (validators, totalWeight) = ValidatorSets.parseValidators(validatorBytes);
        require(validators.length > 0, "Validator set cannot be empty");
        require(totalWeight > 0, "Total weight must be greater than 0");
    }
}
