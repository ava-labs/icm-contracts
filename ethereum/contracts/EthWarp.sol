// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.25;

import {
    WarpMessage,
    WarpBlockHash
} from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";
import {IWarpExt} from "@avalabs/avalanche-contracts/teleporter/IWarpExt.sol";
import {ICMMessage} from "@avalabs/avalanche-contracts/teleporter/ITeleporterMessenger.sol";
import {ICM} from "./utils/ICM.sol";
import {IVerifyICMMessage} from "./interfaces/IVerifyWarpMessage.sol";


contract EthWarp is IWarpExt {

    /**
     * @notice The chain ID of the Ethereum network the contract is deployed on.
     * @dev The chain ID for Ethereum is a uint which we reinterpret as bytes32
     * to remain consistent with the existing interface
     */
    bytes32 public blockchainID;

    /**
     * @notice The chain ID of that the Warp message originated from. Used
     * to ensure that message IDs are computed computed consistently on source and
     * target teleporter contracts.
     */
    bytes32 public warpSourceChainID;

    /**
     * @notice A mapping of avalanche chain IDs to contract addresses that know how
     * to validate received Warp message.
     */
    mapping(bytes32 avaBlockchainId => address verifyWarpMessage) internal _registeredChains;

    constructor (uint256 blockChainId, bytes32 _warpSourceChainID) {
        blockchainID = bytes32(uint256(blockChainId));
        warpSourceChainID = _warpSourceChainID;
    }

    function getVerifiedICMMessage(
        ICMMessage calldata icmMessage
    ) external view returns (WarpMessage memory warpMessage) {
        require(
            isChainRegistered(icmMessage.unsignedMessage.avalancheSourceBlockchainID),
            "Cannot receive a Warp message from a chain whose validator set is unknown"
        );
        bool isValid = IVerifyICMMessage(_registeredChains[icmMessage.unsignedMessage.avalancheSourceBlockchainID])
            .verifyICMMessage(icmMessage);
        require(isValid, "Received an invalid ICM message");
        warpMessage = ICM.handleMessage(icmMessage.unsignedMessage, warpSourceChainID);
        return warpMessage;
    }

    function sendWarpMessage(bytes calldata payload) external returns (bytes32) {
        revert("Sending Warp messages from Ethereum is not currently supported");
    }

    function getVerifiedWarpMessage(uint32 index) external pure returns (WarpMessage calldata, bool) {
        revert("This method can't be called on Ethereum, use `getVerifiedICMMessage` instead");
    }

    function getVerifiedWarpBlockHash(
        uint32 index
    ) external pure returns (WarpBlockHash calldata, bool) {
        revert("This method cannot be called on Ethereum");
    }

    function getBlockchainID() external view returns (bytes32) {
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