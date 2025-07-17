// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TokenHome} from "./TokenHome.sol";
import {INativeTokenHome} from "./interfaces/INativeTokenHome.sol";
import {INativeSendAndCallReceiver} from "../interfaces/INativeSendAndCallReceiver.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    SingleHopCallMessage
} from "../interfaces/ITokenTransferrer.sol";
import {IWrappedNativeToken} from "../interfaces/IWrappedNativeToken.sol";
import {CallUtils} from "@utilities/CallUtils.sol";
import {SafeWrappedNativeTokenDeposit} from "@utilities/SafeWrappedNativeTokenDeposit.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @title NativeTokenHomeUpgradeable
 * @notice An {INativeTokenHome} implementation that locks the native token of this chain to be transferred to
 * TokenRemote instances on other chains.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract NativeTokenHomeUpgradeable is INativeTokenHome, TokenHome {
    using Address for address payable;

    // solhint-disable private-vars-leading-underscore
    /**
     * @dev Namespace storage slots following the ERC-7201 standard to prevent
     * storage collisions between upgradeable contracts.
     *
     * @custom:storage-location erc7201:avalanche-ictt.storage.NativeTokenHome
     */
    struct NativeTokenHomeStorage {
        /**
         * @notice The wrapped native token contract that represents the native tokens on this chain.
         */
        IWrappedNativeToken _wrappedToken;
    }
    // solhint-enable private-vars-leading-underscore

    /**
     * @dev Storage slot computed based off ERC-7201 formula
     * keccak256(abi.encode(uint256(keccak256("avalanche-ictt.storage.NativeTokenHome")) - 1)) & ~bytes32(uint256(0xff));
     */
    bytes32 public constant NATIVE_TOKEN_HOME_STORAGE_LOCATION =
        0x3b5030f10c94fcbdaa3022348ff0b82dbd4c0c71339e41ff59d0bdc92179d600;

    // solhint-disable ordering
    function _getNativeTokenHomeStorage() private pure returns (NativeTokenHomeStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := NATIVE_TOKEN_HOME_STORAGE_LOCATION
        }
    }

    constructor(
        ICMInitializable init
    ) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initializes this token TokenHome instance to send native tokens to TokenRemote instances on other chains.
     * Always uses a {tokenDecimals} value of 18 since it is the denomination of the native token of EVM instances.
     * @param teleporterRegistryAddress The current blockchain ID's Teleporter registry
     * address. See here for details: https://github.com/ava-labs/icm-contracts/tree/main/contracts/teleporter/registry
     * @param teleporterManager Address that manages this contract's integration with the
     * Teleporter registry and Teleporter versions.
     * @param minTeleporterVersion Minimum Teleporter version supported by this contract.
     * @param wrappedTokenAddress The wrapped native token contract address of the native asset
     * to be transferred to TokenRemote instances.
     */
    function initialize(
        address teleporterRegistryAddress,
        address teleporterManager,
        uint256 minTeleporterVersion,
        address wrappedTokenAddress
    ) public initializer {
        __NativeTokenHome_init(
            teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __NativeTokenHome_init(
        address teleporterRegistryAddress,
        address teleporterManager,
        uint256 minTeleporterVersion,
        address wrappedTokenAddress
    ) internal onlyInitializing {
        __TokenHome_init(
            teleporterRegistryAddress,
            teleporterManager,
            minTeleporterVersion,
            wrappedTokenAddress,
            18
        );
        __NativeTokenHome_init_unchained(wrappedTokenAddress);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __NativeTokenHome_init_unchained(
        address wrappedTokenAddress
    ) internal onlyInitializing {
        _getNativeTokenHomeStorage()._wrappedToken = IWrappedNativeToken(wrappedTokenAddress);
    }
    // solhint-enable ordering

    /**
     * @notice Receives native tokens transferred to this contract.
     * @dev This function is called when the token transferrer is withdrawing native tokens to
     * transfer to the recipient. The caller must be the wrapped native token contract.
     */
    receive() external payable {
        // The caller here is expected to be {tokenAddress} directly, and not through a meta-transaction,
        // so we check for `msg.sender` instead of `_msgSender()`.
        require(
            msg.sender == address(_getNativeTokenHomeStorage()._wrappedToken),
            "NativeTokenHome: invalid receive payable sender"
        );
    }

    /**
     * @dev See {INativeTokenTransferrer-send}
     */
    function send(
        SendTokensInput calldata input
    ) external payable {
        _send(input, msg.value);
    }

    /**
     * @dev See {INativeTokenTransferrer-sendAndCall}
     */
    function sendAndCall(
        SendAndCallInput calldata input
    ) external payable {
        _sendAndCall({
            sourceBlockchainID: getBlockchainID(),
            originTokenTransferrerAddress: address(this),
            originSenderAddress: _msgSender(),
            input: input,
            amount: msg.value
        });
    }

    /**
     * @dev See {INativeTokenHome-addCollateral}
     */
    function addCollateral(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress
    ) external payable {
        _addCollateral(remoteBlockchainID, remoteTokenTransferrerAddress, msg.value);
    }

    /**
     * @dev See {TokenHome-_deposit}
     * Deposits the native tokens sent to this contract
     */
    function _deposit(
        uint256 amount
    ) internal virtual override returns (uint256) {
        NativeTokenHomeStorage storage $ = _getNativeTokenHomeStorage();
        return SafeWrappedNativeTokenDeposit.safeDeposit($._wrappedToken, amount);
    }

    /**
     * @dev See {TokenHome-_withdraw}
     * Withdraws the wrapped tokens for native tokens,
     * and sends them to the recipient.
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        NativeTokenHomeStorage storage $ = _getNativeTokenHomeStorage();
        emit TokensWithdrawn(recipient, amount);
        $._wrappedToken.withdraw(amount);
        payable(recipient).sendValue(amount);
    }

    /**
     * @dev See {TokenHome-_handleSendAndCall}
     *
     * Send the native tokens to the recipient contract as a part of the call to
     * {INativeSendAndCallReceiver-receiveTokens} on the recipient contract.
     * If the call fails or doesn't spend all of the tokens, the remaining amount is
     * sent to the fallback recipient.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual override {
        NativeTokenHomeStorage storage $ = _getNativeTokenHomeStorage();
        // Withdraw the native token from the wrapped native token contract.
        $._wrappedToken.withdraw(amount);

        // Encode the call to {INativeSendAndCallReceiver-receiveTokens}
        bytes memory payload = abi.encodeCall(
            INativeSendAndCallReceiver.receiveTokens,
            (
                message.sourceBlockchainID,
                message.originTokenTransferrerAddress,
                message.originSenderAddress,
                message.recipientPayload
            )
        );

        // Call the recipient contract with the given payload, gas amount, and value.
        bool success = CallUtils._callWithExactGasAndValue(
            message.recipientGasLimit, amount, message.recipientContract, payload
        );

        // If the call failed, send the funds to the fallback recipient.
        if (success) {
            emit CallSucceeded(message.recipientContract, amount);
        } else {
            emit CallFailed(message.recipientContract, amount);
            payable(message.fallbackRecipient).sendValue(amount);
        }
    }
}
