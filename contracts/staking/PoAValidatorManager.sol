// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IPoAValidatorManager} from "./interfaces/IPoAValidatorManager.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {ICMInitializable} from "../utilities/ICMInitializable.sol";
import {ValidatorManagerSettings} from "./interfaces/IValidatorManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";

contract PoAStakingManager is IPoAValidatorManager, ValidatorManager, OwnableUpgradeable {
    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    function initialize(
        ValidatorManagerSettings calldata settings,
        address initialOwner
    ) external initializer {
        __PoAStakingManager_init(settings, initialOwner);
    }

    // solhint-disable func-name-mixedcase, ordering
    function __PoAStakingManager_init(
        ValidatorManagerSettings calldata settings,
        address initialOwner
    ) internal onlyInitializing {
        __ValidatorManager_init(settings);
        __Ownable_init(initialOwner);
        __PoAStakingManager_init_unchained();
    }

    // solhint-disable-next-line no-empty-blocks
    function __PoAStakingManager_init_unchained() internal onlyInitializing {}

    // solhint-enable func-name-mixedcase

    function initializeValidatorRegistration(
        uint64 weight,
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature
    ) external override onlyOwner returns (bytes32 validationID) {
        return _initializeValidatorRegistration(nodeID, weight, registrationExpiry, signature);
    }

    // solhint-enable ordering
    function initializeEndValidation(bytes32 validationID) external override {
        _initializeEndValidation(validationID, 0);
    }
}