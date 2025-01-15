// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorRegistrationInput} from "../ValidatorManager.sol";

/**
 * @notice Interface for Proof of Authority Validator Manager contracts
 */
interface IPoAValidatorManager {
    /**
     * @notice Begins the validator registration process, and sets the {weight} of the validator.
     * @param registrationInput The inputs for a validator registration.
     * @param weight The weight of the validator being registered.
     */
    function initiateValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint64 weight
    ) external returns (bytes32 validationID);

    /**
     * @notice Begins the process of ending an active validation period. The validation period must have been previously
     * started by a successful call to {completeValidatorRegistration} with the given validationID.
     * @param validationID The ID of the validation period being ended.
     */
    function initiateValidatorRemoval(bytes32 validationID) external;
}
