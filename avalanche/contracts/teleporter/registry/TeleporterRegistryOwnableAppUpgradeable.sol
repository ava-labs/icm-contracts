// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TeleporterRegistryAppUpgradeable} from "./TeleporterRegistryAppUpgradeable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

/**
 * @dev Contract that inherits {TeleporterRegistryAppUpgradeable} and allows
 * only owners of the contract to update the minimum Teleporter version or
 * pause and unpause specific Teleporter versions.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract TeleporterRegistryOwnableAppUpgradeable is
    TeleporterRegistryAppUpgradeable,
    OwnableUpgradeable
{
    // solhint-disable-next-line func-name-mixedcase
    function __TeleporterRegistryOwnableApp_init(
        address teleporterRegistryAddress,
        address initialOwner,
        uint256 minTeleporterVersion
    ) internal onlyInitializing {
        __TeleporterRegistryApp_init(teleporterRegistryAddress, minTeleporterVersion);
        __Ownable_init(initialOwner);
    }

    // solhint-disable-next-line func-name-mixedcase, no-empty-blocks, func-name-mixedcase
    function _TeleporterRegistryOwnableApp_init_unchained() internal onlyInitializing {}

    /**
     * @dev See {TeleporterRegistryAppUpgradeable-_checkTeleporterRegistryAppAccess}
     *
     * Checks that the caller is the owner of the contract for updating {minTeleporterVersion},
     * and pausing/unpausing specific Teleporter version interactions.
     */
    function _checkTeleporterRegistryAppAccess() internal view virtual override {
        _checkOwner();
    }
}
