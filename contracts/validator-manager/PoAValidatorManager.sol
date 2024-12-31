// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManager} from "./ValidatorManager.sol";
import {
    ValidatorManagerSettings,
    ValidatorRegistrationInput
} from "./interfaces/IValidatorManager.sol";
import {IPoAValidatorManager} from "./interfaces/IPoAValidatorManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

/**
 * @dev Implementation of the {IPoAValidatorManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 * @custom:oz-upgrades-unsafe-allow constructor external-library-linking
 */
contract PoAValidatorManager is IPoAValidatorManager, ValidatorManager, OwnableUpgradeable {
    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    function initialize(
        ValidatorManagerSettings calldata settings,
        address initialOwner
    ) external initializer {
        __PoAValidatorManager_init(settings, initialOwner);
    }

    // solhint-disable func-name-mixedcase, ordering
    function __PoAValidatorManager_init(
        ValidatorManagerSettings calldata settings,
        address initialOwner
    ) internal onlyInitializing {
        __ValidatorManager_init(settings);
        __Ownable_init(initialOwner);
    }

    // solhint-disable-next-line no-empty-blocks
    function __PoAValidatorManager_init_unchained() internal onlyInitializing {}

    // solhint-enable func-name-mixedcase

    /**
     * @notice See {IPoAValidatorManager-initializeValidatorRegistration}.
     */
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint64 weight
    ) external onlyOwner returns (bytes32 validationID) {
        return _initializeValidatorRegistration(registrationInput, weight);
    }

    // solhint-enable ordering
    /**
     * @notice See {IPoAValidatorManager-initializeEndValidation}.
     */
    function initializeEndValidation(bytes32 validationID) external override onlyOwner {
        _initializeEndValidation(validationID);
    }

    /**
     * @notice See {IValidatorManager-completeEndValidation}.
     */
    function completeEndValidation(uint32 messageIndex) external {
        _completeEndValidation(messageIndex);
    }
}
