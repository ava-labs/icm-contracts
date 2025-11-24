// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {NativeTokenRemoteUpgradeable} from "./NativeTokenRemoteUpgradeable.sol";
import {TokenRemoteSettings} from "./interfaces/ITokenRemote.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @title NativeTokenRemote
 * @notice A non-upgradeable version of {NativeTokenRemoteUpgradeable} that calls the parent upgradeable contract's initialize function.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract NativeTokenRemote is NativeTokenRemoteUpgradeable {
    constructor(
        TokenRemoteSettings memory settings,
        string memory nativeAssetSymbol,
        uint256 initialReserveImbalance,
        uint256 burnedFeesReportingRewardPercentage
    ) NativeTokenRemoteUpgradeable(ICMInitializable.Allowed) {
        initialize(
            settings,
            nativeAssetSymbol,
            initialReserveImbalance,
            burnedFeesReportingRewardPercentage
        );
    }
}
