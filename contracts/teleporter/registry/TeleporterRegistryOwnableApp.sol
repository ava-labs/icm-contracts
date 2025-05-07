// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TeleporterRegistryApp} from "./TeleporterRegistryApp.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";

/**
 * @dev Contract that inherits {TeleporterRegistryApp} and allows
 * only owners of the contract to update the minimum Teleporter version or
 * pause and unpause specific Teleporter versions.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract TeleporterRegistryOwnableApp is TeleporterRegistryApp, Ownable {
    constructor(
        address teleporterRegistryAddress,
        address initialOwner,
        uint256 minTeleporterVersion
    )
        TeleporterRegistryApp(teleporterRegistryAddress, minTeleporterVersion)
        Ownable(initialOwner)
    {}

    /**
     * @dev See {TeleporterRegistryApp-_checkTeleporterRegistryAppAccess}
     *
     * Checks that the caller is the owner of the contract for updating {minTeleporterVersion},
     * and pausing/unpausing specific Teleporter version interactions.
     */
    function _checkTeleporterRegistryAppAccess() internal view virtual override {
        _checkOwner();
    }
}
