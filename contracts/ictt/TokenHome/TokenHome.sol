// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterRegistryOwnableAppUpgradeable} from
    "@teleporter/registry/TeleporterRegistryOwnableAppUpgradeable.sol";
import {ITokenHome, RemoteTokenTransferrerSettings} from "./interfaces/ITokenHome.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    TransferrerMessageType,
    TransferrerMessage,
    SingleHopSendMessage,
    SingleHopCallMessage,
    MultiHopSendMessage,
    MultiHopCallMessage,
    RegisterRemoteMessage
} from "../interfaces/ITokenTransferrer.sol";
import {SendReentrancyGuardUpgradeable} from "@utilities/SendReentrancyGuardUpgradeable.sol";
import {TokenScalingUtils} from "@utilities/TokenScalingUtils.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/ERC20.sol";

/**
 * @title TokenHome
 * @dev Abstract contract for a token transferrer home that sends its specified token to {TokenRemote} instances.
 *
 * This contract also handles multi-hop transfers, where tokens sent from a {TokenRemote}
 * instance are forwarded to another {TokenRemote} instance.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract TokenHome is
    ITokenHome,
    TeleporterRegistryOwnableAppUpgradeable,
    SendReentrancyGuardUpgradeable
{
    // solhint-disable private-vars-leading-underscore
    /**
     * @dev Namespace storage slots following the ERC-7201 standard to prevent
     * storage collisions between upgradeable contracts.
     *
     * @custom:storage-location erc7201:avalanche-ictt.storage.TokenHome
     */
    struct TokenHomeStorage {
        /// @notice The blockchain ID of the chain this contract is deployed on.
        bytes32 _blockchainID;
        /**
         * @notice The token address this home contract transfers to TokenRemote instances.
         * For multi-hop transfers, this {tokenAddress} is always used to pay for the secondary message fees.
         * If the token is an ERC20 token, the contract address is directly passed in.
         * If the token is a native asset, the contract address is the wrapped token contract.
         */
        address _tokenAddress;
        uint8 _tokenDecimals;
        /**
         * @notice Tracks the settings for each remote token transferrer instance. TokenRemote instances
         * must register with their {TokenHome} contracts via Teleporter message to be able to
         * receive tokens from this contract.
         */
        mapping(
            bytes32 remoteBlockchainID
                => mapping(
                    address remoteTokenTransferrerAddress
                        => RemoteTokenTransferrerSettings remoteSettings
                )
        ) _registeredRemotes;
        /**
         * @notice Tracks the balances of tokens sent to TokenRemote instances.
         * Balances are represented in the remote token's denomination,
         * and token transferrers are not allowed to unwrap more than has been sent to them.
         * @dev (remoteBlockchainID, remoteTokenTransferrerAddress) -> balance
         */
        mapping(
            bytes32 remoteBlockchainID
                => mapping(address remoteTokenTransferrerAddress => uint256 balance)
        ) _transferredBalances;
    }
    // solhint-enable private-vars-leading-underscore

    /**
     * @dev Storage slot computed based off ERC-7201 formula
     * keccak256(abi.encode(uint256(keccak256("avalanche-ictt.storage.TokenHome")) - 1)) & ~bytes32(uint256(0xff));
     */
    bytes32 public constant TOKEN_HOME_STORAGE_LOCATION =
        0x9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e600;

    // solhint-disable ordering
    function _getTokenHomeStorage() private pure returns (TokenHomeStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := TOKEN_HOME_STORAGE_LOCATION
        }
    }

    /**
     * @notice Initializes this home token transferrer instance to send tokens to TokenRemote instances on other chains.
     * @param teleporterRegistryAddress The address of the TeleporterRegistry contract.
     * @param teleporterManager Address that manages this contract's integration with the
     * Teleporter registry and Teleporter versions.
     * @param minTeleporterVersion Minimum Teleporter version supported by this contract.
     * @param tokenAddress The token contract address to be transferredd by the home instance.
     * @param tokenDecimals The number of decimals for the token being transferred.
     */
    // solhint-disable-next-line func-name-mixedcase
    function __TokenHome_init(
        address teleporterRegistryAddress,
        address teleporterManager,
        uint256 minTeleporterVersion,
        address tokenAddress,
        uint8 tokenDecimals
    ) internal virtual onlyInitializing {
        __TeleporterRegistryOwnableApp_init(
            teleporterRegistryAddress, teleporterManager, minTeleporterVersion
        );
        __SendReentrancyGuard_init();
        __TokenHome_init_unchained(tokenAddress, tokenDecimals);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __TokenHome_init_unchained(
        address tokenAddress,
        uint8 tokenDecimals
    ) internal onlyInitializing {
        require(tokenAddress != address(0), "TokenHome: zero token address");
        require(
            tokenDecimals <= TokenScalingUtils.MAX_TOKEN_DECIMALS,
            "TokenHome: token decimals too high"
        );
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        $._blockchainID =
            IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
        $._tokenAddress = tokenAddress;
        $._tokenDecimals = tokenDecimals;
    }
    // solhint-enable ordering

    function getRemoteTokenTransferrerSettings(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress
    ) public view returns (RemoteTokenTransferrerSettings memory) {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        return $._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress];
    }

    function getTransferredBalance(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress
    ) public view returns (uint256) {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        return $._transferredBalances[remoteBlockchainID][remoteTokenTransferrerAddress];
    }

    function getTokenAddress() public view returns (address) {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        return $._tokenAddress;
    }

    function getBlockchainID() public view returns (bytes32) {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        return $._blockchainID;
    }

    function _registerRemote(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        RegisterRemoteMessage memory message
    ) internal {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        require(remoteBlockchainID != bytes32(0), "TokenHome: zero remote blockchain ID");
        require(
            remoteBlockchainID != $._blockchainID, "TokenHome: cannot register remote on same chain"
        );
        require(
            remoteTokenTransferrerAddress != address(0),
            "TokenHome: zero remote token transferrer address"
        );
        require(
            !$._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress].registered,
            "TokenHome: remote already registered"
        );
        require(
            message.remoteTokenDecimals <= TokenScalingUtils.MAX_TOKEN_DECIMALS,
            "TokenHome: remote token decimals too high"
        );
        require(
            message.homeTokenDecimals == $._tokenDecimals, "TokenHome: invalid home token decimals"
        );

        (uint256 tokenMultiplier, bool multiplyOnRemote) = TokenScalingUtils
            .deriveTokenMultiplierValues($._tokenDecimals, message.remoteTokenDecimals);

        // Calculate the collateral needed in home token denomination.
        uint256 collateralNeeded = TokenScalingUtils.removeTokenScale(
            tokenMultiplier, multiplyOnRemote, message.initialReserveImbalance
        );

        // Round up the collateral needed by 1 in the case that {multiplyOnRemote} is true and
        // {initialReserveImbalance} is not divisible by the {tokenMultiplier} to
        // ensure that the full amount is accounted for.
        if (multiplyOnRemote && message.initialReserveImbalance % tokenMultiplier != 0) {
            collateralNeeded += 1;
        }

        $._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress] =
        RemoteTokenTransferrerSettings({
            registered: true,
            collateralNeeded: collateralNeeded,
            tokenMultiplier: tokenMultiplier,
            multiplyOnRemote: multiplyOnRemote
        });

        emit RemoteRegistered(
            remoteBlockchainID,
            remoteTokenTransferrerAddress,
            collateralNeeded,
            message.remoteTokenDecimals
        );
    }

    /**
     * @notice Sends tokens to the specified remote token transferrer instance.
     *
     * @dev Increases the balance sent to the remote token transferrer instance,
     * and uses Teleporter to send a cross chain message. The amount passed is assumed to
     * be already scaled to the local denomination for this token home.
     * Requirements:
     *
     * - The TokenRemote instance specified by {input.destinationBlockchainID} and {input.destinationTokenTransferrerAddress} must
     *   be registered with this contract.
     * - {input.recipient} cannot be the zero address
     * - {amount} must be greater than 0
     */
    function _send(SendTokensInput memory input, uint256 amount) internal sendNonReentrant {
        _validateSendTokensInput(input);
        // Require that a single hop transfer does not have a multi-hop fallback recipient.
        require(input.multiHopFallback == address(0), "TokenHome: non-zero multi-hop fallback");

        (uint256 adjustedAmount, uint256 feeAmount) = _prepareSend({
            remoteBlockchainID: input.destinationBlockchainID,
            remoteTokenTransferrerAddress: input.destinationTokenTransferrerAddress,
            amount: amount,
            primaryFeeTokenAddress: input.primaryFeeTokenAddress,
            feeAmount: input.primaryFee
        });

        TransferrerMessage memory message = TransferrerMessage({
            messageType: TransferrerMessageType.SINGLE_HOP_SEND,
            payload: abi.encode(
                SingleHopSendMessage({recipient: input.recipient, amount: adjustedAmount})
            )
        });

        // Send message to the TokenRemote instance
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationTokenTransferrerAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensSent(messageID, _msgSender(), input, adjustedAmount);
    }

    /**
     * @notice Routes tokens from a multi-hop message to the specified remote token transferrer instance.
     *
     * @dev Increases the balance sent to the remote token transferrer instance,
     * and uses Teleporter to send a cross chain message. The amount passed is assumed to
     * be already scaled to the local denomination for this token home.
     * Requirements:
     *
     * - The TokenRemote instance specified by {input.destinationBlockchainID} and {input.destinationTokenTransferrerAddress} must
     *   be registered with this contract.
     * - {input.recipient} cannot be the zero address
     * - {amount} must be greater than 0
     */
    function _routeMultiHop(
        SendTokensInput memory input,
        uint256 amount
    ) internal sendNonReentrant {
        _validateSendTokensInput(input);

        uint256 adjustedAmount = _prepareMultiHopRouting(
            input.destinationBlockchainID,
            input.destinationTokenTransferrerAddress,
            amount,
            input.primaryFee
        );

        if (adjustedAmount == 0) {
            // If the adjusted amount is zero for any reason (i.e. unregistered remote,
            // being scaled down to zero, etc.), send the tokens to the multi-hop fallback.
            _withdraw(input.multiHopFallback, amount);
            return;
        }

        TransferrerMessage memory message = TransferrerMessage({
            messageType: TransferrerMessageType.SINGLE_HOP_SEND,
            payload: abi.encode(
                SingleHopSendMessage({recipient: input.recipient, amount: adjustedAmount})
            )
        });

        // Send message to the TokenRemote instance.
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationTokenTransferrerAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: input.primaryFee
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensRouted(messageID, input, adjustedAmount);
    }

    function _sendAndCall(
        bytes32 sourceBlockchainID,
        address originTokenTransferrerAddress,
        address originSenderAddress,
        SendAndCallInput memory input,
        uint256 amount
    ) internal sendNonReentrant {
        _validateSendAndCallInput(input);

        // Require that a single hop transfer does not have a multi-hop fallback recipient.
        require(input.multiHopFallback == address(0), "TokenHome: non-zero multi-hop fallback");

        (uint256 adjustedAmount, uint256 feeAmount) = _prepareSend({
            remoteBlockchainID: input.destinationBlockchainID,
            remoteTokenTransferrerAddress: input.destinationTokenTransferrerAddress,
            amount: amount,
            primaryFeeTokenAddress: input.primaryFeeTokenAddress,
            feeAmount: input.primaryFee
        });

        TransferrerMessage memory message = TransferrerMessage({
            messageType: TransferrerMessageType.SINGLE_HOP_CALL,
            payload: abi.encode(
                SingleHopCallMessage({
                    sourceBlockchainID: sourceBlockchainID,
                    originTokenTransferrerAddress: originTokenTransferrerAddress,
                    originSenderAddress: originSenderAddress,
                    recipientContract: input.recipientContract,
                    amount: adjustedAmount,
                    recipientPayload: input.recipientPayload,
                    recipientGasLimit: input.recipientGasLimit,
                    fallbackRecipient: input.fallbackRecipient
                })
            )
        });

        // Send message to the TokenRemote instance.
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationTokenTransferrerAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensAndCallSent(messageID, originSenderAddress, input, adjustedAmount);
    }

    function _routeMultiHopSendAndCall(
        bytes32 sourceBlockchainID,
        address originTokenTransferrerAddress,
        address originSenderAddress,
        SendAndCallInput memory input,
        uint256 amount
    ) internal sendNonReentrant {
        _validateSendAndCallInput(input);
        uint256 adjustedAmount = _prepareMultiHopRouting(
            input.destinationBlockchainID,
            input.destinationTokenTransferrerAddress,
            amount,
            input.primaryFee
        );

        if (adjustedAmount == 0) {
            // If the adjusted amount is zero for any reason (i.e. unregistered remote,
            // being scaled down to zero, etc.), send the tokens to the multi-hop fallback recipient.
            _withdraw(input.multiHopFallback, amount);
            return;
        }

        TransferrerMessage memory message = TransferrerMessage({
            messageType: TransferrerMessageType.SINGLE_HOP_CALL,
            payload: abi.encode(
                SingleHopCallMessage({
                    sourceBlockchainID: sourceBlockchainID,
                    originTokenTransferrerAddress: originTokenTransferrerAddress,
                    originSenderAddress: originSenderAddress,
                    recipientContract: input.recipientContract,
                    amount: adjustedAmount,
                    recipientPayload: input.recipientPayload,
                    recipientGasLimit: input.recipientGasLimit,
                    fallbackRecipient: input.fallbackRecipient
                })
            )
        });

        // Send message to the TokenRemote instance.
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationTokenTransferrerAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: input.primaryFee
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensAndCallRouted(messageID, input, adjustedAmount);
    }

    /**
     * @dev See {INativeTokenHome-addCollateral}
     */
    function _addCollateral(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount
    ) internal sendNonReentrant {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        RemoteTokenTransferrerSettings memory remoteSettings =
            $._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress];
        require(remoteSettings.registered, "TokenHome: remote not registered");
        require(remoteSettings.collateralNeeded > 0, "TokenHome: zero collateral needed");

        // Deposit the full amount, and withdraw back to the sender if there is excess.
        amount = _deposit(amount);

        // Calculate the remaining collateral needed, any excess amount, and adjust
        // {amount} to represent the amount of tokens added as collateral.
        uint256 remainingCollateralNeeded;
        uint256 excessAmount;
        if (amount >= remoteSettings.collateralNeeded) {
            remainingCollateralNeeded = 0;
            excessAmount = amount - remoteSettings.collateralNeeded;
            amount = remoteSettings.collateralNeeded;
        } else {
            remainingCollateralNeeded = remoteSettings.collateralNeeded - amount;
        }

        // Update the remaining collateral needed.
        $._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress].collateralNeeded =
            remainingCollateralNeeded;
        emit CollateralAdded(
            remoteBlockchainID, remoteTokenTransferrerAddress, amount, remainingCollateralNeeded
        );

        // If there is excess amount, send it back to the sender.
        if (excessAmount > 0) {
            _withdraw(_msgSender(), excessAmount);
        }
    }

    /**
     * @dev See {ITeleporterUpgradeable-_receiveTeleporterMessage}
     *
     * Handles the processing of Teleporter messages sent to this contract.
     * Supported message types include registering a TokenRemote instance, single-hop sends,
     * single-hop calls, multi-hop sends, and multi-hop calls.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        TransferrerMessage memory transferrerMessage = abi.decode(message, (TransferrerMessage));
        if (transferrerMessage.messageType == TransferrerMessageType.SINGLE_HOP_SEND) {
            SingleHopSendMessage memory payload =
                abi.decode(transferrerMessage.payload, (SingleHopSendMessage));

            uint256 homeAmount =
                _processSingleHopTransfer(sourceBlockchainID, originSenderAddress, payload.amount);

            // Send the tokens to the recipient.
            _withdraw(payload.recipient, homeAmount);
            return;
        } else if (transferrerMessage.messageType == TransferrerMessageType.SINGLE_HOP_CALL) {
            SingleHopCallMessage memory payload =
                abi.decode(transferrerMessage.payload, (SingleHopCallMessage));

            uint256 homeAmount =
                _processSingleHopTransfer(sourceBlockchainID, originSenderAddress, payload.amount);

            // Verify that the payload's source blockchain ID and origin token transferrer address matches the source blockchain ID
            // and origin sender address passed from Teleporter.
            require(
                payload.sourceBlockchainID == sourceBlockchainID,
                "TokenHome: mismatched source blockchain ID"
            );
            require(
                payload.originTokenTransferrerAddress == originSenderAddress,
                "TokenHome: mismatched origin sender address"
            );

            _handleSendAndCall(payload, homeAmount);
            return;
        } else if (transferrerMessage.messageType == TransferrerMessageType.MULTI_HOP_SEND) {
            MultiHopSendMessage memory payload =
                abi.decode(transferrerMessage.payload, (MultiHopSendMessage));

            (uint256 homeAmount, uint256 fee) = _processMultiHopTransfer(
                sourceBlockchainID, originSenderAddress, payload.amount, payload.secondaryFee
            );

            // For a multi-hop send, the fee token address has to be {tokenAddress},
            // because the fee is taken from the amount that has already been deposited.
            // For ERC20 tokens, the token address of the contract is directly passed.
            // For native assets, the contract address is the wrapped token contract.
            _routeMultiHop(
                SendTokensInput({
                    destinationBlockchainID: payload.destinationBlockchainID,
                    destinationTokenTransferrerAddress: payload.destinationTokenTransferrerAddress,
                    recipient: payload.recipient,
                    primaryFeeTokenAddress: $._tokenAddress,
                    primaryFee: fee,
                    secondaryFee: 0,
                    requiredGasLimit: payload.secondaryGasLimit,
                    multiHopFallback: payload.multiHopFallback
                }),
                homeAmount
            );
            return;
        } else if (transferrerMessage.messageType == TransferrerMessageType.MULTI_HOP_CALL) {
            MultiHopCallMessage memory payload =
                abi.decode(transferrerMessage.payload, (MultiHopCallMessage));

            (uint256 homeAmount, uint256 fee) = _processMultiHopTransfer(
                sourceBlockchainID, originSenderAddress, payload.amount, payload.secondaryFee
            );

            // For a multi-hop send, the fee token address has to be {tokenAddress},
            // because the fee is taken from the amount that has already been deposited.
            // For ERC20 tokens, the token address of the contract is directly passed.
            // For native assets, the contract address is the wrapped token contract.
            _routeMultiHopSendAndCall({
                sourceBlockchainID: sourceBlockchainID,
                originTokenTransferrerAddress: originSenderAddress,
                originSenderAddress: payload.originSenderAddress,
                input: SendAndCallInput({
                    destinationBlockchainID: payload.destinationBlockchainID,
                    destinationTokenTransferrerAddress: payload.destinationTokenTransferrerAddress,
                    recipientContract: payload.recipientContract,
                    recipientPayload: payload.recipientPayload,
                    requiredGasLimit: payload.secondaryRequiredGasLimit,
                    recipientGasLimit: payload.recipientGasLimit,
                    multiHopFallback: payload.multiHopFallback,
                    fallbackRecipient: payload.fallbackRecipient,
                    primaryFeeTokenAddress: $._tokenAddress,
                    primaryFee: fee,
                    secondaryFee: 0
                }),
                amount: homeAmount
            });
            return;
        } else if (transferrerMessage.messageType == TransferrerMessageType.REGISTER_REMOTE) {
            RegisterRemoteMessage memory payload =
                abi.decode(transferrerMessage.payload, (RegisterRemoteMessage));
            _registerRemote(sourceBlockchainID, originSenderAddress, payload);
        }
    }

    /**
     * @notice Deposits tokens from the sender to this contract,
     * and returns the adjusted amount of tokens deposited.
     * @param amount The initial amount sent to this contract.
     * @return The actual amount deposited to this contract.
     */
    function _deposit(
        uint256 amount
    ) internal virtual returns (uint256);

    /**
     * @notice Withdraws tokens to the recipient address.
     * @param recipient The address to withdraw tokens to
     * @param amount The amount of tokens to withdraw
     */
    function _withdraw(address recipient, uint256 amount) internal virtual;

    /**
     * @notice Processes a send and call message by calling the recipient contract.
     * @param message The send and call message include recipient calldata
     * @param amount The amount of tokens to be sent to the recipient. This amount is assumed to be
     * already scaled to the local denomination for this token home.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual;

    /**
     * @notice Processes a received single hop transfer from a TokenRemote instance.
     * Validates that the message is sent from a registered TokenRemote instance,
     * and is already collateralized.
     * @param remoteBlockchainID The blockchain ID of the TokenRemote instance.
     * @param remoteTokenTransferrerAddress The address of the TokenRemote instance.
     * @param amount The amount of tokens sent back from remote, denominated by the
     * remote's token scale amount.
     */
    function _processSingleHopTransfer(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount
    ) private returns (uint256) {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        RemoteTokenTransferrerSettings memory remoteSettings =
            $._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress];

        return _processReceivedTransfer(
            remoteSettings, remoteBlockchainID, remoteTokenTransferrerAddress, amount
        );
    }

    /**
     * @notice Processes a received multi-hop transfer from a TokenRemote instance.
     * Validates that the message is sent from a registered TokenRemote instance,
     * and is already collateralized.
     * @param remoteBlockchainID The blockchain ID of the TokenRemote instance.
     * @param remoteTokenTransferrerAddress The address of the TokenRemote instance.
     * @param amount The amount of tokens sent back from remote, denominated by the
     * remote's token scale amount.
     * @param secondaryFee The Teleporter fee for the second hop of the mutihop transfer
     */
    function _processMultiHopTransfer(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount,
        uint256 secondaryFee
    ) private returns (uint256, uint256) {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        RemoteTokenTransferrerSettings memory remoteSettings =
            $._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress];

        uint256 transferAmount = _processReceivedTransfer(
            remoteSettings, remoteBlockchainID, remoteTokenTransferrerAddress, amount
        );

        uint256 fee = TokenScalingUtils.removeTokenScale(
            remoteSettings.tokenMultiplier, remoteSettings.multiplyOnRemote, secondaryFee
        );

        return (transferAmount, fee);
    }

    /**
     * @notice Processes a received transfer from a TokenRemote instance.
     * Deducts the balance transferred to the given TokenRemote instance.
     * Removes the token scaling of the remote, checks the associated home token
     * amount is greater than zero, and returns the home token amount.
     * @param remoteSettings The token transferrer settings for the TokenRemote instance we received the transfer from.
     * @param remoteBlockchainID The blockchain ID of the TokenRemote instance.
     * @param remoteTokenTransferrerAddress The address of the TokenRemote instance.
     * @param amount The amount of tokens sent back from remote, denominated by the
     * remote's token scale amount.
     */
    function _processReceivedTransfer(
        RemoteTokenTransferrerSettings memory remoteSettings,
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount
    ) private returns (uint256) {
        // Require that the remote is registered and has no collateral needed.
        require(remoteSettings.registered, "TokenHome: remote not registered");
        require(remoteSettings.collateralNeeded == 0, "TokenHome: remote not collateralized");

        // Deduct the balance transferred to the given TokenRemote instance prior to scaling the amount.
        _deductSenderBalance(remoteBlockchainID, remoteTokenTransferrerAddress, amount);

        // Remove the token scaling of the remote and get home token amount.
        uint256 homeAmount = TokenScalingUtils.removeTokenScale(
            remoteSettings.tokenMultiplier, remoteSettings.multiplyOnRemote, amount
        );

        // Require that the home token amount is greater than zero after removed scaling.
        require(homeAmount > 0, "TokenHome: zero token amount");
        return homeAmount;
    }

    /**
     * @notice Prepares a multi-hop send by checking the TokenRemote instance settings
     * and adjusting the amount to be sent.
     * @return The scaled amount to be sent to the TokenRemote instance. Zero can be returned if the
     * TokenRemote instance is not registered, needs collateral, or the scaled amount is zero.
     */
    function _prepareMultiHopRouting(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount,
        uint256 fee
    ) private returns (uint256) {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        RemoteTokenTransferrerSettings memory remoteSettings =
            $._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress];
        if (!remoteSettings.registered || remoteSettings.collateralNeeded > 0) {
            return 0;
        }

        // Subtract fee amount from amount prior to scaling.
        require(amount > fee, "TokenHome: insufficient amount to cover fees");
        amount -= fee;

        // Scale the amount based on the token multiplier for the given TokenRemote instance.
        uint256 scaledAmount = TokenScalingUtils.applyTokenScale(
            remoteSettings.tokenMultiplier, remoteSettings.multiplyOnRemote, amount
        );
        if (scaledAmount == 0) {
            return 0;
        }

        // Increase the balance of the TokenRemote instance by the scaled amount.
        $._transferredBalances[remoteBlockchainID][remoteTokenTransferrerAddress] += scaledAmount;

        return scaledAmount;
    }

    /**
     * @dev Prepares tokens to be sent to another chain by handling the
     * locking of the token amount in this contract and updating the accounting
     * balances.
     */
    function _prepareSend(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount,
        address primaryFeeTokenAddress,
        uint256 feeAmount
    ) private returns (uint256, uint256) {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        RemoteTokenTransferrerSettings memory remoteSettings =
            $._registeredRemotes[remoteBlockchainID][remoteTokenTransferrerAddress];
        require(remoteSettings.registered, "TokenHome: remote not registered");
        require(remoteSettings.collateralNeeded == 0, "TokenHome: collateral needed for remote");

        // Deposit the funds sent from the user to the token transferrer,
        // and set to adjusted amount after deposit.
        amount = _deposit(amount);

        if (feeAmount > 0) {
            feeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(primaryFeeTokenAddress), _msgSender(), feeAmount
            );
        }

        // Scale the amount based on the token multiplier for the given TokenRemote instance.
        uint256 scaledAmount = TokenScalingUtils.applyTokenScale(
            remoteSettings.tokenMultiplier, remoteSettings.multiplyOnRemote, amount
        );
        require(scaledAmount > 0, "TokenHome: zero scaled amount");

        // Increase the balance of the TokenRemote instance by the scaled amount.
        $._transferredBalances[remoteBlockchainID][remoteTokenTransferrerAddress] += scaledAmount;

        return (scaledAmount, feeAmount);
    }

    function _deductSenderBalance(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount
    ) private {
        TokenHomeStorage storage $ = _getTokenHomeStorage();
        uint256 senderBalance =
            $._transferredBalances[remoteBlockchainID][remoteTokenTransferrerAddress];
        require(senderBalance >= amount, "TokenHome: insufficient token transfer balance");
        $._transferredBalances[remoteBlockchainID][remoteTokenTransferrerAddress] =
            senderBalance - amount;
    }

    function _validateSendAndCallInput(
        SendAndCallInput memory input
    ) private pure {
        require(input.recipientContract != address(0), "TokenHome: zero recipient contract address");
        require(input.requiredGasLimit > 0, "TokenHome: zero required gas limit");
        require(input.recipientGasLimit > 0, "TokenHome: zero recipient gas limit");
        require(
            input.recipientGasLimit < input.requiredGasLimit,
            "TokenHome: invalid recipient gas limit"
        );
        require(input.fallbackRecipient != address(0), "TokenHome: zero fallback recipient address");
        require(input.secondaryFee == 0, "TokenHome: non-zero secondary fee");
    }

    function _validateSendTokensInput(
        SendTokensInput memory input
    ) private pure {
        require(input.recipient != address(0), "TokenHome: zero recipient address");
        require(input.requiredGasLimit > 0, "TokenHome: zero required gas limit");
        require(input.secondaryFee == 0, "TokenHome: non-zero secondary fee");
    }
}
