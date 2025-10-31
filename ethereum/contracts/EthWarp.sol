// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.25;

import {
    WarpMessage,
    WarpBlockHash,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";
import {IWarpExt} from "@avalabs/avalanche-contracts/teleporter/IWarpExt.sol";
import {ICM, ICMMessage} from "./utils/ICM.sol";
import {IVerifyICMMessage} from "./interfaces/IVerifyWarpMessage.sol";


contract EthWarp is IWarpExt {

    /**
     * @notice The chain ID of the Ethereum network the contract is deployed on.
     * @dev The chain ID for Ethereum is a uint which we reinterpret as bytes32
     * to remain consistent with the existing interface
     */
    bytes32 public blockchainID;

    /**
     * @notice A mapping of avalanche chain IDs to contract addresses that know how
     * to validate received Warp message.
     */
    mapping(bytes32 avaBlockchainId => address verifyWarpMessage) internal _registeredChains;

    constructor (uint blockChainId) {
        blockChainId = blockChainId;
    }

    function getVerifiedMessageFromPayload(
        bytes calldata payload
    ) external view returns (WarpMessage memory warpMessage) {
        ICMMessage memory icmMessage = abi.decode(payload, (ICMMessage));
        require(
            isChainRegistered(icmMessage.unsignedMessage.avalancheSourceBlockchainID),
            "Cannot receive a Warp message from a chain whose validator set is unknown"
        );
        bool isValid = IVerifyICMMessage(_registeredChains[warpMessage.sourceChainID])
            .verifyICMMessage(icmMessage);
        require(isValid, "Received an invalid ICM message");
        warpMessage = ICM.handleMessage(icmMessage.unsignedMessage);
        return warpMessage;
    }

    function sendWarpMessage(bytes calldata payload) external returns (bytes32 messageID) {
        revert("Sending Warp messages from Ethereum is not currently supported");
    }

    function getVerifiedWarpMessage(uint32 index) external view returns (WarpMessage calldata message, bool valid) {
        revert("This method cannot be called on Ethereum, use `getMessageFromPayload` instead");
    }

    function getVerifiedWarpBlockHash(
        uint32 index
    ) external view returns (WarpBlockHash calldata warpBlockHash, bool valid) {
        revert("This method cannot be called on Ethereum");
    }

    function getBlockchainID() external view returns (bytes32 chainID) {
        return blockchainID;
    }

    /**
     * @notice Check if a source chain is registered with this Warp contract. If it is not,
     * this contract will be unable to verify a quorum of signatures is present on the
     * received message.
     * @return registered A boolean indicating presence of the given key
     */
    function isChainRegistered(bytes32 avaBlockchainId) public view returns (bool) {
        return _registeredChains[avaBlockchainId] != address(0);
    }

    /**
     * @notice Register a contract implementing `IVerifyWarpMessage` here to validate messages
     * originating from `avaBlockchainId`.
     */
    function registerChain(bytes32 avaBlockchainId, address verifyWarpMessage) external {
        require(!isChainRegistered(avaBlockchainId), "This chain is already registered");
        require(verifyWarpMessage != address(0), "Provided address does not exist");
        _registeredChains[avaBlockchainId] = verifyWarpMessage;
    }
}