// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner} from "./IACP99Manager.sol";

/**
 * @dev Proof-of-authority manager interface meant to be used as the owner of an
 * {IValidatorManager} contract instance that only allows modifications to the validator set
 * to be initiated by a fixed address. Once initiated, modifications are meant to be able to be
 * completed by any address to allow for improved UX flows.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
interface IPoAManager {
    /**
     * @notice Initiates validator registration. Only callable by the contract owner.
     * @param nodeID The node ID of the validator.
     * @param blsPublicKey The BLS public key of the validator.
     * @param remainingBalanceOwner The P-chain owner for remaining balance.
     * @param disableOwner The P-chain owner for disabling.
     * @param weight The validator's weight.
     * @return validationID The unique identifier for the validator registration.
     */
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) external returns (bytes32 validationID);

    /**
     * @notice Initiates validator removal. Only callable by the contract owner.
     * @param validationID The unique identifier of the validator to remove.
     */
    function initiateValidatorRemoval(
        bytes32 validationID
    ) external;

    /**
     * @notice Initiates a validator weight update. Only callable by the contract owner.
     * @param validationID The unique identifier of the validator.
     * @param newWeight The new weight for the validator.
     * @return nonce The validator nonce associated with the weight change.
     * @return messageID The ICM message ID of the L1ValidatorWeightMessage that executes the weight change on the P-Chain.
     */
    function initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 newWeight
    ) external returns (uint64 nonce, bytes32 messageID);

    /**
     * @notice Completes validator registration. Callable by anyone.
     * @param messageIndex The index of the registration message.
     * @return validationID The unique identifier for the completed registration.
     */
    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32 validationID);

    /**
     * @notice Completes validator removal. Callable by anyone.
     * @param messageIndex The index of the removal message.
     * @return validationID The unique identifier for the completed removal.
     */
    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32 validationID);

    /**
     * @notice Completes a validator weight update. Callable by anyone.
     * @param messageIndex The index of the weight update message.
     * @return validationID The unique identifier for the validator whose weight was updated.
     * @return nonce The validator nonce associated with the executed weight change.
     */
    function completeValidatorWeightUpdate(
        uint32 messageIndex
    ) external returns (bytes32 validationID, uint64 nonce);

    /**
     * @notice Transfers ownership of the validator manager contract.
     * @param newOwner The address to transfer ownership to.
     */
    function transferValidatorManagerOwnership(
        address newOwner
    ) external;
}
