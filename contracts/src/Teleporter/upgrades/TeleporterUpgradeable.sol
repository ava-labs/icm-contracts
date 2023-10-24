// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./TeleporterRegistry.sol";

/**
 * @dev TeleporterUpgradeable provides upgrade utility for applications built on top
 * of the Teleporter protocol by integrating with the {TeleporterRegistry}.
 *
 * This contract is intended to be inherited by other contracts that wish to use the
 * upgrade mechanism. It provides a modifier that restricts access to only Teleporter
 * versions that are greater than or equal to `_minTeleporterVersion`.
 */
abstract contract TeleporterUpgradeable {
    TeleporterRegistry public immutable teleporterRegistry;
    uint256 public minTeleporterVersion;

    event MinTeleporterVersionUpdated(
        uint256 oldMinTeleporterVersion,
        uint256 newMinTeleporterVersion
    );

    /**
     * @dev Throws if called by a `msg.sender` that is not an allowed Teleporter version.
     * Checks that `msg.sender` matches a Teleporter version greater than or equal to `_minTeleporterVersion`.
     */
    modifier onlyAllowedTeleporter() {
        require(
            teleporterRegistry.getVersionFromAddress(msg.sender) >=
                minTeleporterVersion,
            "TeleporterUpgradeable: invalid teleporter sender"
        );
        _;
    }

    /**
     * @dev Initializes the {TeleporterUpgradeable} contract by getting `teleporterRegistry`
     * instance and setting `_minTeleporterVersion`.
     */
    constructor(address teleporterRegistryAddress) {
        require(
            teleporterRegistryAddress != address(0),
            "TeleporterUpgradeable: zero teleporter registry address"
        );

        teleporterRegistry = TeleporterRegistry(teleporterRegistryAddress);
        minTeleporterVersion = teleporterRegistry.getLatestVersion();
    }

    /**
     * @dev Updates `minTeleporterVersion` to the latest version.
     */
    function updateMinTeleporterVersion() external {
        uint256 prevVersion = minTeleporterVersion;
        minTeleporterVersion = teleporterRegistry.getLatestVersion();
        emit MinTeleporterVersionUpdated(prevVersion, minTeleporterVersion);
    }
}
