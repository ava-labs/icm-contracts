// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

/**
 * @notice Interface for a contracts that are called to receive native tokens.
 */
interface INativeSendAndCallReceiver {
    /**
     * @notice Called to receive the amount of the native token. Implementations
     * must properly handle the msg.value of the call in order to ensure it doesn't
     * become improperly made inaccessible.
     * @param sourceBlockchainID Blockchain ID that the transfer originated from
     * @param originTokenTransferrerAddress Address of the token transferrer that initiated the Teleporter message
     * @param originSenderAddress Address of the sender that sent the transfer. This value
     * should only be trusted if {originTokenTransferrerAddress} is verified and known.
     * @param payload Arbitrary data provided by the caller
     */
    function receiveTokens(
        bytes32 sourceBlockchainID,
        address originTokenTransferrerAddress,
        address originSenderAddress,
        bytes calldata payload
    ) external payable;
}
