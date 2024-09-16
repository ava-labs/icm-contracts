// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorRegistrationInput} from "./IValidatorManager.sol";
import {IPoSValidatorManager, PoSValidatorRequirements} from "./IPoSValidatorManager.sol";

interface INativeTokenStakingManager is IPoSValidatorManager {
    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param registrationInput The inputs for a validator registration.
     * @param requirements The requirements for the validator being registered.
     */
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        PoSValidatorRequirements calldata requirements
    ) external payable returns (bytes32 validationID);

    function initializeDelegatorRegistration(bytes32 validationID)
        external
        payable
        returns (bytes32);
}
