// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeable} from "./TeleporterUpgradeable.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @dev Contract that implements the {TeleporterUpgradeable} interface and allows
 * only owners of the contract to update the minimum Teleporter version.
 */
abstract contract TeleporterOwnerUpgradeable is TeleporterUpgradeable, Ownable {
    constructor(
        address teleporterRegistryAddress
    ) TeleporterUpgradeable(teleporterRegistryAddress) {}

    /**
     * @dev See {TeleporterUpgradeable-updateMinTeleporterVersion}
     *
     * Updates the minimum Teleporter version allowed for receiving on this contract
     * to the latest version registered in the {TeleporterRegistry}.
     * Restricted to only owners of the contract.
     * Emits a {MinTeleporterVersionUpdated} event if the minimum Teleporter version
     * was updated.
     */
    function updateMinTeleporterVersion() public override onlyOwner {
        uint256 curMinTeleporterVersion = minTeleporterVersion;
        uint256 newLatestVersion = teleporterRegistry.latestVersion();
        if (newLatestVersion != curMinTeleporterVersion) {
            minTeleporterVersion = newLatestVersion;
            emit MinTeleporterVersionUpdated(
                curMinTeleporterVersion,
                newLatestVersion
            );
        }
    }
}
