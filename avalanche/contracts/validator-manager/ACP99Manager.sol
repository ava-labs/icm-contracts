// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IACP99Manager, PChainOwner} from "./interfaces/IACP99Manager.sol";

/*
 * @title ACP99Manager
 * @notice The ACP99Manager interface represents the functionality for sovereign L1
 * validator management, as specified in ACP-77. 
 * 
 * @dev ACP99Manager defines the private functions specified in ACP-99.
 * The counterpart to this contract is IACP99Manager, which defines the public functions specified in ACP-99.
 * https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/99-validatorsetmanager-contract
 */
abstract contract ACP99Manager is IACP99Manager {
    // solhint-disable ordering

    /**
     * @notice Initiates validator registration by issuing a RegisterL1ValidatorMessage. The validator should
     * not be considered active until completeValidatorRegistration is called.
     *
     * Emits an {InitiatedValidatorRegistration} event on success.
     *
     * @param nodeID The ID of the node to add to the L1.
     * @param blsPublicKey The BLS public key of the validator.
     * @param remainingBalanceOwner The remaining balance owner of the validator.
     * @param disableOwner The disable owner of the validator.
     * @param weight The weight of the node on the L1.
     * @return validationID The ID of the registered validator.
     */
    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) internal virtual returns (bytes32 validationID);

    /**
     * @notice Initiates validator removal by issuing a L1ValidatorWeightMessage with the weight set to zero.
     * The validator should be considered inactive as soon as this function is called.
     *
     * Emits an {InitiatedValidatorRemoval} on success.
     *
     * @param validationID The ID of the validator to remove.
     */
    function _initiateValidatorRemoval(
        bytes32 validationID
    ) internal virtual;

    /**
     * @notice Initiates a validator weight update by issuing an L1ValidatorWeightMessage with a nonzero weight.
     * The validator weight change should not have any effect until completeValidatorWeightUpdate is successfully called.
     *
     * Emits an {InitiatedValidatorWeightUpdate} event on success.
     *
     * @param validationID The ID of the validator to modify.
     * @param weight The new weight of the validator.
     * @return nonce The validator nonce associated with the weight change.
     * @return messageID The ID of the L1ValidatorWeightMessage used to update the validator's weight.
     */
    function _initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 weight
    ) internal virtual returns (uint64 nonce, bytes32 messageID);
}
