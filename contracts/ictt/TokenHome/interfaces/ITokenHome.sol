// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    ITokenTransferrer,
    SendTokensInput,
    SendAndCallInput
} from "../../interfaces/ITokenTransferrer.sol";

/**
 * @notice Each {ITokenRemote} instance registers with the home contract, and provides settings for transferring
 * to the remote token transfer contract.
 * @param registered Whether the remote token transferrer is registered
 * @param collateralNeeded The amount of tokens that must be first added as collateral,
 * through {addCollateral} calls, before tokens can be transferred to the remote token transferrer.
 * @param tokenMultiplier The scaling factor for the amount of tokens to be transferred to the remote.
 * @param multiplyOnRemote Whether the {tokenMultiplier} should be applied when transferring tokens to
 * the remote (multiplyOnRemote=true), or when transferring tokens back to the home (multiplyOnRemote=false).
 */
struct RemoteTokenTransferrerSettings {
    bool registered;
    uint256 collateralNeeded;
    uint256 tokenMultiplier;
    bool multiplyOnRemote;
}
/**
 * @dev Interface for a "home" token transferrer contract that locks a specific token
 * on its chain to be transferred to supported "remote" token transferrer contracts on other chains.
 */

interface ITokenHome is ITokenTransferrer {
    /**
     * @dev Emitted when tokens are added as collateral for a given {ITokenRemote} instance.
     * The event emits a {remaining} value of 0 when the {ITokenRemote} instance is fully collateralized.
     */
    event CollateralAdded(
        bytes32 indexed remoteBlockchainID,
        address indexed remoteTokenTransferrerAddress,
        uint256 amount,
        uint256 remaining
    );

    /**
     * @notice Emitted when a new {ITokenRemote} instance is registered with the token transferrer.
     */
    event RemoteRegistered(
        bytes32 indexed remoteBlockchainID,
        address indexed remoteTokenTransferrerAddress,
        uint256 initialCollateralNeeded,
        uint8 tokenDecimals
    );

    /**
     * @notice Emitted when tokens are routed from a multi-hop send message to another chain.
     */
    event TokensRouted(bytes32 indexed teleporterMessageID, SendTokensInput input, uint256 amount);

    /**
     * @notice Emitted when tokens are routed from a mulit-hop send message,
     * with calldata for a contract recipient, to another chain.
     */
    event TokensAndCallRouted(
        bytes32 indexed teleporterMessageID, SendAndCallInput input, uint256 amount
    );
}
