// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

/**
 * @dev Validator status
 */
enum ValidatorStatus {
    Unknown,
    PendingAdded,
    Active,
    PendingRemoved,
    Completed,
    Invalidated
}

/**
 * @dev Specifies the owner of a validator's remaining balance or disable owner on the P-Chain.
 * P-Chain addresses are also 20-bytes, so we use the address type to represent them.
 */
struct PChainOwner {
    uint32 threshold;
    address[] addresses;
}

/**
 * @dev Contains the active state of a Validator
 */
struct Validator {
    ValidatorStatus status;
    bytes nodeID;
    uint64 startingWeight;
    uint64 messageNonce;
    uint64 weight;
    uint64 startedAt;
    uint64 endedAt;
}

/**
 * @dev Describes the current churn period
 */
struct ValidatorChurnPeriod {
    uint256 startedAt;
    uint64 initialWeight;
    uint64 totalWeight;
    uint64 churnAmount;
}

/**
 * @notice Validator Manager settings, used to initialize the Validator Manager
 * @notice The subnetID is the ID of the L1 that the Validator Manager is managing
 * @notice The churnPeriodSeconds is the duration of the churn period in seconds
 * @notice The maximumChurnPercentage is the maximum percentage of the total weight that can be added or removed in a single churn period
 */
struct ValidatorManagerSettings {
    bytes32 subnetID;
    uint64 churnPeriodSeconds;
    uint8 maximumChurnPercentage;
}

/**
 * @dev Description of the conversion data used to convert
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

/**
 * @dev Specifies an initial validator, used in the conversion data.
 */
struct InitialValidator {
    bytes nodeID;
    bytes blsPublicKey;
    uint64 weight;
}

/**
 * @dev Specifies a validator to register.
 */
struct ValidatorRegistrationInput {
    bytes nodeID;
    bytes blsPublicKey;
    uint64 registrationExpiry;
    PChainOwner remainingBalanceOwner;
    PChainOwner disableOwner;
}

/**
 * @notice Interface for Validator Manager contracts that implement Subnet-only Validator management.
 */
interface IValidatorManager {
    /**
     * @notice Emitted when a new validation period is created by locking stake in the manager contract.
     * Note: This event does not mean that the validation period has been successfully registered on the P-Chain,
     * and rewards for this validation period will not begin accruing until the {ValidationPeriodRegistered} event is
     * emitted.
     * @param validationID The ID of the validation period being created.
     * @param nodeID The node ID of the validator being registered.
     * @param registerValidationMessageID The ID of the ICM message that will be sent to the P-Chain to register the
     * validation period.
     * @param weight The weight of the validator being registered.
     * @param registrationExpiry The Unix timestamp after which the reigistration is no longer valid on the P-Chain.
     */
    event ValidationPeriodCreated(
        bytes32 indexed validationID,
        bytes indexed nodeID,
        bytes32 indexed registerValidationMessageID,
        uint64 weight,
        uint64 registrationExpiry
    );

    event InitialValidatorCreated(
        bytes32 indexed validationID, bytes indexed nodeID, uint64 weight
    );

    /**
     * @notice Emitted when the staking manager learns that the validation period has been successfully registered
     * on the P-Chain. Rewards for this validation period will begin accruing when this event is emitted.
     * @param validationID The ID of the validation period being registered.
     * @param weight The weight of the validator being registered.
     * @param timestamp The time at which the validation period was registered with the contract.
     */
    event ValidationPeriodRegistered(
        bytes32 indexed validationID, uint64 weight, uint256 timestamp
    );

    /**
     * @notice Emitted when the process of ending a registered validation period is started by calling
     * {initializeEndValidation}.
     * Note: The stake for this validation period remains locked until a {ValidationPeriodRemoved} event is emitted.
     * @param validationID The ID of the validation period being removed.
     * @param setWeightMessageID The ID of the ICM message that updates the validator's weight on the P-Chain.
     * @param weight The weight of the validator being removed.
     * @param endTime The time at which the removal was initiated.
     */
    event ValidatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        uint64 weight,
        uint256 endTime
    );

    /**
     * @notice Emitted when the stake for a validation period is unlocked and returned to the staker.
     * This is done by calling {completeEndValidation}, which provides proof from the P-Chain that the
     * validation period is not active and will never be active in the future.
     * @param validationID The ID of the validation period being removed.
     */
    event ValidationPeriodEnded(bytes32 indexed validationID, ValidatorStatus indexed status);

    /**
     * @notice Event emitted when validator weight is updated.
     * @param validationID The ID of the validation period being updated
     * @param nonce The message nonce used to update the validator weight
     * @param weight The updated validator weight that is sent to the P-Chain
     * @param setWeightMessageID The ID of the ICM message that updates the validator's weight on the P-Chain
     */
    event ValidatorWeightUpdate(
        bytes32 indexed validationID,
        uint64 indexed nonce,
        uint64 weight,
        bytes32 setWeightMessageID
    );

    /**
     * @notice Verifies and sets the initial validator set for the chain through a P-Chain SubnetToL1ConversionMessage.
     * @param conversionData The subnet conversion message data used to recompute and verify against the conversionID.
     * @param messsageIndex The index that contains the SubnetToL1ConversionMessage ICM message containing the conversionID to be verified against the provided {ConversionData}
     */
    function initializeValidatorSet(
        ConversionData calldata conversionData,
        uint32 messsageIndex
    ) external;

    /**
     * @notice Resubmits a validator registration message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being registered.
     */
    function resendRegisterValidatorMessage(bytes32 validationID) external;

    /**
     * @notice Completes the validator registration process by returning an acknowledgement of the registration of a
     * validationID from the P-Chain.
     * @param messageIndex The index of the ICM message to be received providing the acknowledgement.
     */
    function completeValidatorRegistration(uint32 messageIndex) external;

    /**
     * @notice Resubmits a validator end message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being ended.
     */
    function resendEndValidatorMessage(bytes32 validationID) external;

    /**
     * @notice Completes the process of ending a validation period by receiving an acknowledgement from the P-Chain
     * that the validation ID is not active and will never be active in the future. Returns the the stake associated
     * with the validation.
     * Note: This function can be used for successful validation periods that have been explicitly ended by calling
     * {initializeEndValidation} or for validation periods that never began on the P-Chain due to the {registrationExpiry} being reached.
     * @param messageIndex The index of the ICM message to be received providing the proof the validation is not active
     * and never will be active on the P-Chain.
     */
    function completeEndValidation(uint32 messageIndex) external;
}
