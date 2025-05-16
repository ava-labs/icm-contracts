// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

/// @notice L1 validator status.
enum ValidatorStatus {
    Unknown,
    PendingAdded,
    Active,
    PendingRemoved,
    Completed,
    Invalidated
}

/**
 * @notice Description of the conversion data used to convert
 * a subnet to an L1 on the P-Chain.
 * This data is the pre-image of a hash that is authenticated by the P-Chain
 * and verified by the Validator Manager.
 */
struct ConversionData {
    bytes32 subnetID;
    bytes32 validatorManagerBlockchainID;
    address validatorManagerAddress;
    InitialValidator[] initialValidators;
}

/// @notice Specifies an initial validator, used in the conversion data.
struct InitialValidator {
    bytes nodeID;
    bytes blsPublicKey;
    uint64 weight;
}

/**
 * @notice Specifies the owner of a validator's remaining balance or disable owner on the P-Chain.
 * P-Chain addresses are also 20-bytes, so we use the address type to represent them.
 */
struct PChainOwner {
    uint32 threshold;
    address[] addresses;
}

/**
 * @notice Contains the active state of a Validator.
 * @param status The validator status.
 * @param nodeID The NodeID of the validator.
 * @param startingWeight The weight of the validator at the time of registration.
 * @param sentNonce The current weight update nonce sent by the manager.
 * @param receivedNonce The highest nonce received from the P-Chain.
 * @param weight The current weight of the validator.
 * @param startTime The start time of the validator.
 * @param endTime The end time of the validator.
 */
struct Validator {
    ValidatorStatus status;
    bytes nodeID;
    uint64 startingWeight;
    uint64 sentNonce;
    uint64 receivedNonce;
    uint64 weight;
    uint64 startTime;
    uint64 endTime;
}

/*
 * @title IACP99Manager
 * @notice The IACP99Manager interface represents the functionality for sovereign L1
 * validator management, as specified in ACP-77.
 *
 * @dev IACP99Manager defines the public functions specified in ACP-99.
 * The counterpart to this interface is ACP99Manager, which defines the private functions specified in ACP-99.
 * https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/99-validatorsetmanager-contract
 */
interface IACP99Manager {
    /**
     * @notice Emitted when an initial validator is registered.
     * @notice The field index is the index of the initial validator in the conversion data.
     * This is used along with the subnetID as the ACP-118 justification in
     * signature requests to P-Chain validators over a L1ValidatorRegistrationMessage
     * when removing the validator
     */
    event RegisteredInitialValidator(
        bytes32 indexed validationID,
        bytes20 indexed nodeID,
        bytes32 indexed subnetID,
        uint64 weight
    );
    /// @notice Emitted when a validator registration to the L1 is initiated.
    event InitiatedValidatorRegistration(
        bytes32 indexed validationID,
        bytes20 indexed nodeID,
        bytes32 registrationMessageID,
        uint64 registrationExpiry,
        uint64 weight
    );
    /// @notice Emitted when a validator registration to the L1 is completed.
    event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight);
    /// @notice Emitted when removal of an L1 validator is initiated.
    event InitiatedValidatorRemoval(
        bytes32 indexed validationID,
        bytes32 validatorWeightMessageID,
        uint64 weight,
        uint64 endTime
    );
    /// @notice Emitted when removal of an L1 validator is completed.
    event CompletedValidatorRemoval(bytes32 indexed validationID);
    /// @notice Emitted when a validator weight update is initiated.
    event InitiatedValidatorWeightUpdate(
        bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight
    );
    /// @notice Emitted when a validator weight update is completed.
    event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight);

    /**
     * @notice Verifies and sets the initial validator set for the chain by consuming a
     * SubnetToL1ConversionMessage from the P-Chain.
     *
     * Emits a {RegisteredInitialValidator} event for each initial validator in {conversionData}.
     *
     * @param conversionData The Subnet conversion message data used to recompute and verify against the ConversionID.
     * @param messageIndex The index that contains the SubnetToL1ConversionMessage ICM message containing the
     * ConversionID to be verified against the provided {conversionData}.
     */
    function initializeValidatorSet(
        ConversionData calldata conversionData,
        uint32 messageIndex
    ) external;

    /**
     * @notice Completes the validator registration process by returning an acknowledgement of the registration of a
     * validationID from the P-Chain. The validator should not be considered active until this method is successfully called.
     *
     * Emits a {CompletedValidatorRegistration} event on success.
     *
     * @param messageIndex The index of the L1ValidatorRegistrationMessage to be received providing the acknowledgement.
     * @return validationID The ID of the registered validator.
     */
    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32 validationID);

    /**
     * @notice Completes validator removal by consuming a RegisterL1ValidatorMessage from the P-Chain acknowledging
     * that the validator has been removed, or that it was not registered on the P-Chain and the expiry time has passed.
     *
     * Emits a {CompletedValidatorRemoval} on success.
     *
     * @param messageIndex The index of the RegisterL1ValidatorMessage.
     * @return validationID The ID of the validator that was removed.
     */
    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32 validationID);

    /**
     * @notice Completes the validator weight update process by consuming an L1ValidatorWeightMessage from the P-Chain
     * acknowledging the weight update. The validator weight change should not have any effect until this method is successfully called.
     *
     * Emits a {CompletedValidatorWeightUpdate} event on success.
     *
     * @param messageIndex The index of the L1ValidatorWeightMessage message to be received providing the acknowledgement.
     * @return validationID The ID of the validator, retreived from the L1ValidatorWeightMessage.
     * @return nonce The nonce of the validator, retreived from the L1ValidatorWeightMessage.
     */
    function completeValidatorWeightUpdate(
        uint32 messageIndex
    ) external returns (bytes32 validationID, uint64 nonce);

    /// @notice Returns the SubnetID of the L1 tied to this manager
    function subnetID() external view returns (bytes32 id);

    /// @notice Returns the validator details for a given validation ID.
    function getValidator(
        bytes32 validationID
    ) external view returns (Validator memory validator);

    /// @notice Returns the total weight of the current L1 validator set.
    function l1TotalWeight() external view returns (uint64 weight);
}
