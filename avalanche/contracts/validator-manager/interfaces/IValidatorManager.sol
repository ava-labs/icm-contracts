// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner, ValidatorStatus, IACP99Manager} from "./IACP99Manager.sol";

/**
 * @dev Validator Manager interface that provides additional functionality on top of {IACP99Manager}
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
interface IValidatorManager is IACP99Manager {
    error InvalidValidatorManagerAddress(address validatorManagerAddress);
    error InvalidWarpOriginSenderAddress(address senderAddress);
    error InvalidValidatorManagerBlockchainID(bytes32 blockchainID);
    error InvalidWarpSourceChainID(bytes32 sourceChainID);
    error InvalidInitializationStatus();
    error InvalidMaximumChurnPercentage(uint8 maximumChurnPercentage);
    error InvalidChurnPeriodLength(uint64 churnPeriodLength);
    error InvalidBLSKeyLength(uint256 length);
    error InvalidNodeID(bytes nodeID);
    error InvalidConversionID(bytes32 encodedConversionID, bytes32 expectedConversionID);
    error InvalidTotalWeight(uint64 weight);
    error InvalidValidationID(bytes32 validationID);
    error InvalidValidatorStatus(ValidatorStatus status);
    error InvalidNonce(uint64 nonce);
    error InvalidWarpMessage();
    error MaxChurnRateExceeded(uint64 churnAmount);
    error NodeAlreadyRegistered(bytes nodeID);
    error UnexpectedRegistrationStatus(bool validRegistration);
    error InvalidPChainOwnerThreshold(uint256 threshold, uint256 addressesLength);
    error InvalidPChainOwnerAddresses();
    error ZeroAddress();

    /**
     * @notice Migrates a validator from the V1 contract to the V2 contract.
     * @param validationID The ID of the validation period to migrate.
     * @param receivedNonce The latest nonce received from the P-Chain.
     */
    function migrateFromV1(bytes32 validationID, uint32 receivedNonce) external;

    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) external returns (bytes32);

    /**
     * @notice Resubmits a validator registration message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being registered.
     */
    function resendRegisterValidatorMessage(
        bytes32 validationID
    ) external;

    function initiateValidatorRemoval(
        bytes32 validationID
    ) external;

    /**
     * @notice Resubmits a validator removal message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being ended.
     */
    function resendValidatorRemovalMessage(
        bytes32 validationID
    ) external;

    function initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 newWeight
    ) external returns (uint64, bytes32);

    /**
     * @notice Returns a validation ID registered to the given nodeID
     * @param nodeID ID of the node associated with the validation ID
     */
    function getNodeValidationID(
        bytes calldata nodeID
    ) external view returns (bytes32);

    function getChurnPeriodSeconds() external view returns (uint64);
}
