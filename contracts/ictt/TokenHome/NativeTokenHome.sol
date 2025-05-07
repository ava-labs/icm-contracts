// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {NativeTokenHomeUpgradeable} from "./NativeTokenHomeUpgradeable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @title NativeTokenHome
 * @notice A non-upgradeable version of {NativeTokenHomeUpgradeable} that calls the parent upgradeable contract's initialize function.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract NativeTokenHome is NativeTokenHomeUpgradeable {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        uint256 minTeleporterVersion,
        address wrappedTokenAddress
    ) NativeTokenHomeUpgradeable(ICMInitializable.Allowed) {
        initialize(
            teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress
        );
    }
}
