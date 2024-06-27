// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {INativeTokenTransferer} from "../../interfaces/INativeTokenTransferer.sol";
import {ITokenHome} from "./ITokenHome.sol";

/**
 * @notice Interface for a native token "home" contract that locks the native token
 * on its chain to be transferred to supported remote bridge contracts on other chains.
 *
 * @custom:security-contact https://github.com/ava-labs/avalanche-interchain-token-transfer/blob/main/SECURITY.md
 */
interface INativeTokenHome is INativeTokenTransferer, ITokenHome {
    /**
     * @notice Adds collateral to the home bridge contract for the specified TokenRemote instance. If more value is provided
     * than the amount of collateral needed, the excess amount is returned to the caller.
     * @param remoteBlockchainID The blockchain ID of the remote bridge contract to add collateral for.
     * @param remoteBridgeAddress The address of the remote bridge contract to add collateral for on
     * the {remoteBlockchainID}.
     */
    function addCollateral(
        bytes32 remoteBlockchainID,
        address remoteBridgeAddress
    ) external payable;
}
