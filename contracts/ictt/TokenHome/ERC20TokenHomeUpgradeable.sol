// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TokenHome} from "./TokenHome.sol";
import {IERC20TokenHome} from "./interfaces/IERC20TokenHome.sol";
import {IERC20SendAndCallReceiver} from "../interfaces/IERC20SendAndCallReceiver.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    SingleHopCallMessage
} from "../interfaces/ITokenTransferrer.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/ERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {CallUtils} from "@utilities/CallUtils.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @title ERC20TokenHomeUpgradeable
 * @notice An {IERC20TokenHome} implementation that locks a specified ERC20 token to be sent to
 * TokenRemote instances on other chains.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract ERC20TokenHomeUpgradeable is IERC20TokenHome, TokenHome {
    using SafeERC20 for IERC20;

    // solhint-disable private-vars-leading-underscore
    /**
     * @dev Namespace storage slots following the ERC-7201 standard to prevent
     * storage collisions between upgradeable contracts.
     *
     * @custom:storage-location erc7201:avalanche-ictt.storage.ERC20TokenHome
     */
    struct ERC20TokenHomeStorage {
        /// @notice The ERC20 token this home contract transfers to TokenRemote instances.
        IERC20 _token;
    }
    // solhint-enable private-vars-leading-underscore

    /**
     * @dev Storage slot computed based off ERC-7201 formula
     * keccak256(abi.encode(uint256(keccak256("avalanche-ictt.storage.ERC20TokenHome")) - 1)) & ~bytes32(uint256(0xff));
     */
    bytes32 public constant ERC20_TOKEN_HOME_STORAGE_LOCATION =
        0x914a9547f6c3ddce1d5efbd9e687708f0d1d408ce129e8e1a88bce4f40e29500;

    // solhint-disable ordering
    function _getERC20TokenHomeStorage() private pure returns (ERC20TokenHomeStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := ERC20_TOKEN_HOME_STORAGE_LOCATION
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
     * @notice Initializes the token TokenHome instance to send ERC20 tokens to TokenRemote instances on other chains.
     * @param teleporterRegistryAddress The current blockchain ID's Teleporter registry
     * address. See here for details: https://github.com/ava-labs/icm-contracts/tree/main/contracts/teleporter/registry
     * @param teleporterManager Address that manages this contract's integration with the
     * Teleporter registry and Teleporter versions.
     * @param minTeleporterVersion Minimum Teleporter version supported by this contract.
     * @param tokenAddress The ERC20 token contract address to be transferred by the home.
     * @param tokenDecimals The number of decimals for the ERC20 token
     */
    function initialize(
        address teleporterRegistryAddress,
        address teleporterManager,
        uint256 minTeleporterVersion,
        address tokenAddress,
        uint8 tokenDecimals
    ) public initializer {
        __ERC20TokenHome_init(
            teleporterRegistryAddress,
            teleporterManager,
            minTeleporterVersion,
            tokenAddress,
            tokenDecimals
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ERC20TokenHome_init(
        address teleporterRegistryAddress,
        address teleporterManager,
        uint256 minTeleporterVersion,
        address tokenAddress,
        uint8 tokenDecimals
    ) internal onlyInitializing {
        __TokenHome_init(
            teleporterRegistryAddress,
            teleporterManager,
            minTeleporterVersion,
            tokenAddress,
            tokenDecimals
        );
        __ERC20TokenHome_init_unchained(tokenAddress);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ERC20TokenHome_init_unchained(
        address tokenAddress
    ) internal onlyInitializing {
        _getERC20TokenHomeStorage()._token = IERC20(tokenAddress);
    }
    // solhint-enable ordering

    /**
     * @dev See {IERC20TokenTransferrer-send}
     */
    function send(SendTokensInput calldata input, uint256 amount) external {
        _send(input, amount);
    }

    /**
     * @dev See {IERC20TokenTransferrer-sendAndCall}
     */
    function sendAndCall(SendAndCallInput calldata input, uint256 amount) external {
        _sendAndCall({
            sourceBlockchainID: getBlockchainID(),
            originTokenTransferrerAddress: address(this),
            originSenderAddress: _msgSender(),
            input: input,
            amount: amount
        });
    }

    /**
     * @dev See {IERC20TokenHome-addCollateral}
     */
    function addCollateral(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount
    ) external {
        _addCollateral(remoteBlockchainID, remoteTokenTransferrerAddress, amount);
    }

    /**
     * @dev See {TokenHome-_deposit}
     */
    function _deposit(
        uint256 amount
    ) internal virtual override returns (uint256) {
        ERC20TokenHomeStorage storage $ = _getERC20TokenHomeStorage();
        return SafeERC20TransferFrom.safeTransferFrom($._token, _msgSender(), amount);
    }

    /**
     * @dev See {TokenHome-_withdraw}
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        ERC20TokenHomeStorage storage $ = _getERC20TokenHomeStorage();
        emit TokensWithdrawn(recipient, amount);
        $._token.safeTransfer(recipient, amount);
    }

    /**
     * @dev See {TokenHome-_handleSendAndCall}
     *
     * Approves the recipient contract to spend the amount of tokens from this contract,
     * and calls {IERC20SendAndCallReceiver-receiveTokens} on the recipient contract.
     * If the call fails or doesn't spend all of the tokens, the remaining amount is
     * sent to the fallback recipient.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual override {
        ERC20TokenHomeStorage storage $ = _getERC20TokenHomeStorage();
        IERC20 token = $._token;
        // Approve the recipient contract to spend the amount from the collateral.
        SafeERC20.safeIncreaseAllowance($._token, message.recipientContract, amount);

        // Encode the call to {IERC20SendAndCallReceiver-receiveTokens}
        bytes memory payload = abi.encodeCall(
            IERC20SendAndCallReceiver.receiveTokens,
            (
                message.sourceBlockchainID,
                message.originTokenTransferrerAddress,
                message.originSenderAddress,
                address(token),
                amount,
                message.recipientPayload
            )
        );

        // Call the recipient contract with the given payload and gas amount.
        bool success = CallUtils._callWithExactGas(
            message.recipientGasLimit, message.recipientContract, payload
        );

        uint256 remainingAllowance = token.allowance(address(this), message.recipientContract);

        // Reset the recipient contract allowance to 0.
        SafeERC20.forceApprove(token, message.recipientContract, 0);

        if (success) {
            emit CallSucceeded(message.recipientContract, amount);
        } else {
            emit CallFailed(message.recipientContract, amount);
        }

        // Transfer any remaining allowance to the fallback recipient. This will be the
        // full amount if the call failed.
        if (remainingAllowance > 0) {
            token.safeTransfer(message.fallbackRecipient, remainingAllowance);
        }
    }
}
