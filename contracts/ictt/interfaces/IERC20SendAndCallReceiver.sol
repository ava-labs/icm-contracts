// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

/**
 * @notice Interface for contracts that are called to receive token transfers.
 */
interface IERC20SendAndCallReceiver {
    /**
     * @notice Called to receive the amount of the given token
     * @param sourceBlockchainID Blockchain ID that the transfer originated from
     * @param originTokenTransferrerAddress Address of the token transferrer that initiated the Teleporter message
     * @param originSenderAddress Address of the sender that sent the transfer. This value
     * should only be trusted if {originTokenTransferrerAddress} is verified and known.
     * @param token Address of the token to be received
     * @param amount Amount of the token to be received
     * @param payload Arbitrary data provided by the caller
     */
    function receiveTokens(
        bytes32 sourceBlockchainID,
        address originTokenTransferrerAddress,
        address originSenderAddress,
        address token,
        uint256 amount,
        bytes calldata payload
    ) external;
}
