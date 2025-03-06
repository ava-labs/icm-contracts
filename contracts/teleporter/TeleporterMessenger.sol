// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {
    TeleporterMessageReceipt,
    TeleporterMessageInput,
    TeleporterMessage,
    TeleporterFeeInfo,
    ITeleporterMessenger
} from "./ITeleporterMessenger.sol";
import {ReceiptQueue} from "./ReceiptQueue.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {ITeleporterReceiver} from "./ITeleporterReceiver.sol";
import {ReentrancyGuards} from "@utilities/ReentrancyGuards.sol";

/**
 * @dev Implementation of the {ITeleporterMessenger} interface.
 *
 * This implementation is used to send messages cross chain using the IWarpMessenger precompile,
 * and to receive messages sent from other chains. Teleporter contracts should be deployed through Nick's method
 * of universal deployer, such that the same contract is deployed at the same address on all chains.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract TeleporterMessenger is ITeleporterMessenger, ReentrancyGuards {
    using SafeERC20 for IERC20;
    using ReceiptQueue for ReceiptQueue.TeleporterMessageReceiptQueue;

    /**
     * @notice SentMessageInfo includes the fee information for a given message submitted
     * to be sent, along with the hash of the message itself.
     */
    struct SentMessageInfo {
        bytes32 messageHash;
        TeleporterFeeInfo feeInfo;
    }

    /**
     * @notice Warp precompile used for sending and receiving Warp messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    /**
     * @notice The blockchain ID of the chain the contract is deployed on.
     * @dev Can be initialized by calling initializeBlockchainID, or will be initialized
     * automatically on the first successful call to send or receive a message.
     */
    bytes32 public blockchainID;

    /**
     * @notice A monotonically incremented integer tracking the total number of messages sent by this TeleporterMessenger contract.
     * @dev Used to provide uniqueness when generating message IDs for new messages. The first message sent will use a
     * messageNonce of 1 such that the nonce value can be used to provide replay protection for a given message ID.
     */
    uint256 public messageNonce;

    /**
     * @notice Tracks the outstanding receipts to send back to a given chain in subsequent messages sent to that chain.
     * @dev The key is the blockchain ID of the other chain, and the value is a queue of pending receipts for messages
     * received from that chain.
     */
    mapping(bytes32 sourceBlockchainID => ReceiptQueue.TeleporterMessageReceiptQueue receiptQueue)
        public receiptQueues;

    /**
     * @notice Tracks the message hash and fee information for each message sent that has yet to be acknowledged
     * with a receipt.
     * @dev The key is the message ID, and the value is the info for the uniquely identified message.
     */
    mapping(bytes32 messageID => SentMessageInfo messageInfo) public sentMessageInfo;

    /**
     * @notice  Tracks the hash of messages that have been received but have never succeeded in execution.
     * @dev Enables retrying of failed messages with higher gas limits. Message execution is guaranteed to
     * succeed at most once. The key is the message ID, and the value is the hash of the uniquely
     * identified message whose execution failed.
     */
    mapping(bytes32 messageID => bytes32 messageHash) public receivedFailedMessageHashes;

    /**
     * @dev Tracks the message nonce for each message that has been received. The key is the message ID,
     * and the value is the nonce value that was received as a part of that message.
     * Note: the `messageNonce` values are also used to determine if a given message has been received or not.
     */
    mapping(bytes32 messageID => uint256 messageNonce) internal _receivedMessageNonces;

    /**
     * @dev Tracks the relayer reward address for each message that has been received.
     * The key is the message ID, and the value is the reward address provided by the deliverer of the message.
     */
    mapping(bytes32 messageID => address relayerRewardAddress) internal _relayerRewardAddresses;

    /**
     * @dev Tracks the reward amounts for a given asset able to be redeemed by a given relayer.
     * The first key is the relayer reward address, the second key is the fee token contract address,
     * and the value is the amount of the asset redeemable by the relayer.
     */
    mapping(
        address relayerRewardAddress
            => mapping(address feeTokenContract => uint256 redeemableRewardAmount)
    ) internal _relayerRewardAmounts;

    /**
     * @dev See {ITeleporterMessenger-sendCrossChainMessage}
     *
     * When executed, a relayer may kick off an asynchronous event to have the validators of the
     * chain create an aggregate BLS signature of the message.
     *
     * Emits a {SendCrossChainMessage} event when message successfully gets sent.
     */
    function sendCrossChainMessage(
        TeleporterMessageInput calldata messageInput
    ) external senderNonReentrant returns (bytes32) {
        // Get the outstanding receipts for messages that have been previously received
        // from the destination chain but not yet acknowledged, and attach the receipts
        // to the Teleporter message to be sent.
        return _sendTeleporterMessage(
            messageInput,
            receiptQueues[messageInput.destinationBlockchainID].getOutstandingReceiptsToSend()
        );
    }

    /**
     * @dev Emits a {SendCrossChainMessage} event.
     * Requirements:
     *
     * - `message` must have been previously sent.
     * - `message` encoding must match previously sent message.
     *
     * @inheritdoc ITeleporterMessenger
     */
    function retrySendCrossChainMessage(
        TeleporterMessage calldata message
    ) external senderNonReentrant {
        // Calculate the message ID based on the message nonce.
        // If the blockchain ID has yet to be initialized, no messages have ever been sent by
        // this contract, meaning that the message to be retried will not be found in any event.
        // Thus, don't need to initialize the blockchain ID here.
        bytes32 messageID =
            calculateMessageID(blockchainID, message.destinationBlockchainID, message.messageNonce);

        // Get the previously sent message hash.
        SentMessageInfo memory existingMessageInfo = sentMessageInfo[messageID];
        // If the message hash is zero, the message was never sent.
        require(
            existingMessageInfo.messageHash != bytes32(0), "TeleporterMessenger: message not found"
        );

        // Check that the hash of the provided message matches the one that was originally submitted.
        bytes memory messageBytes = abi.encode(message);
        require(
            keccak256(messageBytes) == existingMessageInfo.messageHash,
            "TeleporterMessenger: invalid message hash"
        );

        // Emit and make state variable changes before external calls when possible,
        // though this function is protected by sender reentrancy guard.
        emit SendCrossChainMessage(
            messageID, message.destinationBlockchainID, message, existingMessageInfo.feeInfo
        );

        // Resubmit the message to the warp precompile now that we know
        // the exact message was already submitted in the past.
        WARP_MESSENGER.sendWarpMessage(messageBytes);
    }

    /**
     * @dev Both `senderNonReentrant` and `receiverNonReentrant` are used here to prevent the
     * external call of `safeTransferFrom` from being reentrant, calling `receiveCrossChainMessage`
     * to attribute rewards for this message, and then still adding fee amount for this message.
     * Emits an {AddFeeAmount} event.
     * Requirements:
     *
     * - `additionalFeeAmount` must be non-zero.
     * - `message` must exist and not have been acknowledged with a receipt yet.
     * - `feeTokenAddress` must match the fee asset contract address used in the original call to `sendCrossChainMessage`.
     * @inheritdoc ITeleporterMessenger
     */
    function addFeeAmount(
        bytes32 messageID,
        address feeTokenAddress,
        uint256 additionalFeeAmount
    ) external senderNonReentrant receiverNonReentrant {
        // The additional fee amount must be non-zero.
        require(additionalFeeAmount > 0, "TeleporterMessenger: zero additional fee amount");

        // Do not allow adding a fee asset with contract address zero.
        require(
            feeTokenAddress != address(0), "TeleporterMessenger: zero fee asset contract address"
        );

        // If a receipt has been received for this message, its hash and fee information
        // will be cleared from state. At this point, you can not add to its fee. This is also the
        // case if the given message never existed.
        require(
            sentMessageInfo[messageID].messageHash != bytes32(0),
            "TeleporterMessenger: message not found"
        );

        // Check that the fee contract address matches the one that was originally used. Only a single
        // fee asset can be used to incentivize the delivery of a given message.
        // We require users to explicitly pass the same fee asset contract address here rather than just using
        // the previously submitted asset type as a defensive measure to avoid having users accidentally confuse
        // which asset they are paying.
        require(
            sentMessageInfo[messageID].feeInfo.feeTokenAddress == feeTokenAddress,
            "TeleporterMessenger: invalid fee asset contract address"
        );

        // Transfer the additional fee amount to this Teleporter instance.
        uint256 adjustedAmount =
            SafeERC20TransferFrom.safeTransferFrom(IERC20(feeTokenAddress), additionalFeeAmount);

        // Store the updated fee amount, and emit it as an event.
        sentMessageInfo[messageID].feeInfo.amount += adjustedAmount;

        emit AddFeeAmount(messageID, sentMessageInfo[messageID].feeInfo);
    }

    /**
     * @dev Emits a {ReceiveCrossChainMessage} event.
     * Re-entrancy is explicitly disallowed between receiving functions. One message is not able to receive another message.
     * Requirements:
     *
     * - `relayerRewardAddress` must not be the zero address.
     * - `messageIndex` must specify a valid warp message in the transaction's storage slots.
     * - Valid warp message provided in storage slots, and sender address matches the address of this contract.
     * - Teleporter message `destinationBlockchainID` must match the `blockchainID` of this contract.
     * - Teleporter message was not previously received.
     * - Transaction was sent by an allowed relayer for corresponding teleporter message.
     *
     * @inheritdoc ITeleporterMessenger
     */
    function receiveCrossChainMessage(
        uint32 messageIndex,
        address relayerRewardAddress
    ) external receiverNonReentrant {
        // Verify and parse the cross chain message included in the transaction access list
        // using the warp message precompile.
        (WarpMessage memory warpMessage, bool success) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(success, "TeleporterMessenger: invalid warp message");

        // Only allow for messages to be received from the same address as this teleporter contract.
        // The contract should be deployed using the universal deployer pattern, such that it knows messages
        // received from the same address on other chains were constructed using the same bytecode of this contract.
        // This allows for trusting the message format and uniqueness as specified by sendCrossChainMessage.
        require(
            warpMessage.originSenderAddress == address(this),
            "TeleporterMessenger: invalid origin sender address"
        );

        // Parse the payload of the message.
        TeleporterMessage memory teleporterMessage =
            abi.decode(warpMessage.payload, (TeleporterMessage));

        // If the blockchain ID has yet to be initialized, do so now.
        bytes32 blockchainID_ = initializeBlockchainID();

        // Require that the message was intended for this blockchain.
        require(
            teleporterMessage.destinationBlockchainID == blockchainID_,
            "TeleporterMessenger: invalid destination chain ID"
        );

        // Calculate the message ID of the message given the source blockchain ID and message nonce.
        bytes32 messageID = calculateMessageID(
            warpMessage.sourceChainID, blockchainID_, teleporterMessage.messageNonce
        );

        // Require that the message has not been received previously.
        require(!_messageReceived(messageID), "TeleporterMessenger: message already received");

        // Check that the caller is allowed to deliver this message.
        require(
            _checkIsAllowedRelayer(msg.sender, teleporterMessage.allowedRelayerAddresses),
            "TeleporterMessenger: unauthorized relayer"
        );

        // Mark the message as received.
        _markMessageReceived(messageID, teleporterMessage.messageNonce);

        // Store the relayer reward address if non-zero.
        if (relayerRewardAddress != address(0)) {
            _relayerRewardAddresses[messageID] = relayerRewardAddress;
        }

        // Process the receipts that were included in the teleporter message by paying the
        // fee for the messages are reward to the given relayers.
        uint256 length = teleporterMessage.receipts.length;
        for (uint256 i; i < length; ++i) {
            _markReceipt(blockchainID_, warpMessage.sourceChainID, teleporterMessage.receipts[i]);
        }

        // Store the receipt of this message delivery.
        receiptQueues[warpMessage.sourceChainID].enqueue(
            TeleporterMessageReceipt({
                receivedMessageNonce: teleporterMessage.messageNonce,
                relayerRewardAddress: relayerRewardAddress
            })
        );

        emit ReceiveCrossChainMessage(
            messageID,
            warpMessage.sourceChainID,
            msg.sender,
            relayerRewardAddress,
            teleporterMessage
        );

        // Execute the message.
        if (teleporterMessage.message.length > 0) {
            _handleInitialMessageExecution(messageID, warpMessage.sourceChainID, teleporterMessage);
        }
    }

    /**
     * @dev A Teleporter message has an associated `requiredGasLimit` that is used to execute the message.
     * If the `requiredGasLimit` is too low, then the message execution will fail. This method allows
     * for retrying the execution of a message with a higher gas limit. Contrary to `receiveCrossChainMessage`,
     * which will only use `requiredGasLimit` in the sub-call to execute the message, this method may
     * use all of the gas available in the transaction.
     *
     * Reverts if the message execution fails again on the specified message.
     * Emits a {MessageExecuted} event if the retry is successful.
     * Requirements:
     *
     * - `message` must have previously failed to execute, and matches the hash of the failed message.
     *
     * @inheritdoc ITeleporterMessenger
     */
    function retryMessageExecution(
        bytes32 sourceBlockchainID,
        TeleporterMessage calldata message
    ) external receiverNonReentrant {
        // Calculate the message ID based on the source blockchainID and message nonce.
        // If the blockchain ID has yet to be initialized, no messages have ever been received by
        // this contract, meaning that the message to be retried will not be found in any event.
        // Thus, don't need to initialize the blockchain ID here.
        bytes32 messageID =
            calculateMessageID(sourceBlockchainID, blockchainID, message.messageNonce);

        // Check that the hash of the payload provided matches the hash of the payload that previously failed to execute.
        bytes32 failedMessageHash = receivedFailedMessageHashes[messageID];
        require(failedMessageHash != bytes32(0), "TeleporterMessenger: message not found");
        require(
            keccak256(abi.encode(message)) == failedMessageHash,
            "TeleporterMessenger: invalid message hash"
        );

        // Check that the target address has fully initialized contract code prior to calling it.
        // If the target address does not have code, the execution automatically fails.
        require(
            message.destinationAddress.code.length > 0,
            "TeleporterMessenger: destination address has no code"
        );

        // Clear the failed message hash from state prior to retrying its execution to redundantly prevent
        // reentrance attacks (on top of the nonReentrant guard).
        emit MessageExecuted(messageID, sourceBlockchainID);
        delete receivedFailedMessageHashes[
            messageID
        ];

        // Re-encode the payload by ABI encoding a call to the {receiveTeleporterMessage} function
        // defined by the {ITeleporterReceiver} interface.
        // If the destination address does not implement {receiveTeleporterMessage}, but does implement
        // a fallback function, then the fallback function will be called instead.
        bytes memory payload = abi.encodeCall(
            ITeleporterReceiver.receiveTeleporterMessage,
            (sourceBlockchainID, message.originSenderAddress, message.message)
        );

        // Reattempt the message execution with all of the gas left available for execution of this transaction.
        // Use all of the gas left because this message has already been successfully received, and it is the
        // caller's responsibility to provide as much gas as is needed. Compared to the initial delivery, where
        // the relayer should still receive their reward even if the message execution takes more gas than expected.
        // Require that the call be successful such that in the failure case this transaction reverts and the
        // message can be retried again if desired.
        bool success = _tryExecuteMessage(message.destinationAddress, gasleft(), payload);
        require(success, "TeleporterMessenger: retry execution failed");
    }

    /**
     * @dev There is no explicit limit to the number of receipts able to be sent by a {sendSpecifiedReceipts} message because
     * this method is intended to be used by relayers themselves to ensure their receipts get returned.
     * There is no fee associated with the empty message, and the same relayer is expected to relay it
     * themselves in order to claim their rewards, so it is their responsibility to ensure that the necessary
     * gas is provided for however many receipts are being retried.
     *
     * These specified receipts are not removed from their corresponding receipt queue because there
     * is no efficient way to remove a specific receipt from an arbitrary position in the queue, and it is
     * harmless for receipts to be sent multiple times within the protocol.
     *
     * Emits {SendCrossChainMessage} event.
     * Requirements:
     * - `messageIDs` must all be valid and have existing receipts.
     *
     * @inheritdoc ITeleporterMessenger
     */
    function sendSpecifiedReceipts(
        bytes32 sourceBlockchainID,
        bytes32[] calldata messageIDs,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external senderNonReentrant returns (bytes32) {
        // If the blockchain ID has yet to be initialized, no messages have ever been received by
        // this contract, meaning that the receipt(s) will not be found in any event.
        // Thus, don't need to initialize the blockchain ID here.
        bytes32 blockchainID_ = blockchainID;

        // Iterate through the specified message IDs and create teleporter receipts to send back.
        TeleporterMessageReceipt[] memory receiptsToSend =
            new TeleporterMessageReceipt[](messageIDs.length);
        uint256 length = messageIDs.length;
        for (uint256 i; i < length; ++i) {
            bytes32 receivedMessageID = messageIDs[i];

            // Get the nonce for this message ID.
            uint256 receivedMessageNonce = _receivedMessageNonces[receivedMessageID];
            require(receivedMessageNonce != 0, "TeleporterMessenger: receipt not found");

            // Check that the message ID was received from the specified source blockchain.
            require(
                receivedMessageID
                    == calculateMessageID(sourceBlockchainID, blockchainID_, receivedMessageNonce),
                "TeleporterMessenger: message ID not from source blockchain"
            );

            // Get the relayer reward address for the message.
            address relayerRewardAddress = _relayerRewardAddresses[receivedMessageID];

            receiptsToSend[i] = TeleporterMessageReceipt({
                receivedMessageNonce: receivedMessageNonce,
                relayerRewardAddress: relayerRewardAddress
            });
        }

        return _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: address(0),
                feeInfo: feeInfo,
                requiredGasLimit: uint256(0),
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: new bytes(0)
            }),
            receiptsToSend
        );
    }

    /**
     * @dev Requirements:
     *
     * - `rewardAmount` must be non-zero.
     *
     * @inheritdoc ITeleporterMessenger
     */
    function redeemRelayerRewards(
        address feeAsset
    ) external {
        uint256 rewardAmount = _relayerRewardAmounts[msg.sender][feeAsset];
        require(rewardAmount > 0, "TeleporterMessenger: no reward to redeem");

        // Zero the reward balance before calling the external ERC20 to transfer the
        // reward to prevent any possible re-entrancy.
        delete _relayerRewardAmounts[msg.sender][feeAsset];

        emit RelayerRewardsRedeemed(msg.sender, feeAsset, rewardAmount);

        // "Fee on transfer" tokens do not require a special case here because
        // the amount credited to the caller does not affect this contract's accounting.
        // The reward is considered paid in full in all cases.
        IERC20(feeAsset).safeTransfer(msg.sender, rewardAmount);
    }

    /**
     * @inheritdoc ITeleporterMessenger
     */
    function getMessageHash(
        bytes32 messageID
    ) external view returns (bytes32) {
        return sentMessageInfo[messageID].messageHash;
    }

    /**
     * @inheritdoc ITeleporterMessenger
     */
    function messageReceived(
        bytes32 messageID
    ) external view returns (bool) {
        return _messageReceived(messageID);
    }

    /**
     * @inheritdoc ITeleporterMessenger
     */
    function getRelayerRewardAddress(
        bytes32 messageID
    ) external view returns (address) {
        require(_messageReceived(messageID), "TeleporterMessenger: message not received");
        return _relayerRewardAddresses[messageID];
    }

    /**
     * @inheritdoc ITeleporterMessenger
     */
    function checkRelayerRewardAmount(
        address relayer,
        address feeAsset
    ) external view returns (uint256) {
        return _relayerRewardAmounts[relayer][feeAsset];
    }

    /**
     * @inheritdoc ITeleporterMessenger
     */
    function getFeeInfo(
        bytes32 messageID
    ) external view returns (address, uint256) {
        TeleporterFeeInfo memory feeInfo = sentMessageInfo[messageID].feeInfo;
        return (feeInfo.feeTokenAddress, feeInfo.amount);
    }

    /**
     * @notice Gets the next message ID to be used for a message sent from the contract instance.
     * @return The next message ID to be used for a message sent from the contract instance.
     */
    function getNextMessageID(
        bytes32 destinationBlockchainID
    ) external view returns (bytes32) {
        bytes32 blockchainID_ = blockchainID;
        require(blockchainID_ != bytes32(0), "TeleporterMessenger: zero blockchain ID");
        uint256 nextMessageNonce = messageNonce + 1;
        return calculateMessageID(blockchainID_, destinationBlockchainID, nextMessageNonce);
    }

    /**
     * @inheritdoc ITeleporterMessenger
     */
    function getReceiptQueueSize(
        bytes32 sourceBlockchainID
    ) external view returns (uint256) {
        return receiptQueues[sourceBlockchainID].size();
    }

    /**
     * @inheritdoc ITeleporterMessenger
     */
    function getReceiptAtIndex(
        bytes32 sourceBlockchainID,
        uint256 index
    ) external view returns (TeleporterMessageReceipt memory) {
        return receiptQueues[sourceBlockchainID].getReceiptAtIndex(index);
    }

    /**
     * @notice If not already set, initializes blockchainID by getting the current
     * blockchain ID value from the Warp precompile.
     * @dev Emits {BlockchainIDInitialized} event.
     * @return The current blockchain ID.
     */
    function initializeBlockchainID() public returns (bytes32) {
        bytes32 blockchainID_ = blockchainID;
        if (blockchainID_ == bytes32(0)) {
            blockchainID_ = WARP_MESSENGER.getBlockchainID();
            require(blockchainID_ != bytes32(0), "TeleporterMessenger: zero blockchain ID");
            blockchainID = blockchainID_;
            emit BlockchainIDInitialized(blockchainID_);
        }
        return blockchainID_;
    }

    /**
     * @notice Calculates the message ID for a message sent from this contract instance with the
     * given source blockchain ID, destination blockchain ID, and message nonce.
     */
    function calculateMessageID(
        bytes32 sourceBlockchainID,
        bytes32 destinationBlockchainID,
        uint256 nonce
    ) public view returns (bytes32) {
        return
            keccak256(abi.encode(address(this), sourceBlockchainID, destinationBlockchainID, nonce));
    }

    /**
     * @dev Checks whether `delivererAddress` is allowed to deliver the message.
     */
    function _checkIsAllowedRelayer(
        address delivererAddress,
        address[] memory allowedRelayers
    ) internal pure returns (bool) {
        // An empty allowed relayers list means anyone is allowed to deliver the message.
        if (allowedRelayers.length == 0) {
            return true;
        }

        // Otherwise, the deliverer address must be included in allowedRelayers.
        uint256 length = allowedRelayers.length;
        for (uint256 i; i < length; ++i) {
            if (allowedRelayers[i] == delivererAddress) {
                return true;
            }
        }
        return false;
    }

    /**
     * @dev Helper function for sending a teleporter message cross chain.
     * Constructs the Teleporter message and sends it through the Warp Messenger precompile,
     * and performs fee transfer if necessary.
     *
     * Emits a {SendCrossChainMessage} event.
     */
    function _sendTeleporterMessage(
        TeleporterMessageInput memory messageInput,
        TeleporterMessageReceipt[] memory receipts
    ) private returns (bytes32) {
        // If the blockchain ID has yet to be initialized, do so now.
        bytes32 blockchainID_ = initializeBlockchainID();

        // Get the message ID to use for this message by incrementing it.
        uint256 messageNonce_ = ++messageNonce;
        bytes32 messageID =
            calculateMessageID(blockchainID_, messageInput.destinationBlockchainID, messageNonce_);

        // Construct and serialize the message.
        TeleporterMessage memory teleporterMessage = TeleporterMessage({
            messageNonce: messageNonce_,
            originSenderAddress: msg.sender,
            destinationBlockchainID: messageInput.destinationBlockchainID,
            destinationAddress: messageInput.destinationAddress,
            requiredGasLimit: messageInput.requiredGasLimit,
            allowedRelayerAddresses: messageInput.allowedRelayerAddresses,
            receipts: receipts,
            message: messageInput.message
        });
        bytes memory teleporterMessageBytes = abi.encode(teleporterMessage);

        // If the fee amount is non-zero, transfer the asset into control of this TeleporterMessenger contract instance.
        // The fee is allowed to be 0 because it's possible for someone to run their own relayer and deliver their own messages,
        // which does not require further incentivization. They still must pay the transaction fee to submit the message, so
        // this is not a DOS vector in terms of being able to submit zero-fee messages.
        uint256 adjustedFeeAmount;
        if (messageInput.feeInfo.amount > 0) {
            // If the fee amount is non-zero, check that the contract address is not address(0)
            require(
                messageInput.feeInfo.feeTokenAddress != address(0),
                "TeleporterMessenger: zero fee asset contract address"
            );

            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(messageInput.feeInfo.feeTokenAddress), messageInput.feeInfo.amount
            );
        }

        // Store the fee asset and amount to be paid to the relayer of this message upon receiving the receipt.
        // Also store the message hash so that it can be retried until a receipt of its delivery is received back.
        TeleporterFeeInfo memory adjustedFeeInfo = TeleporterFeeInfo({
            feeTokenAddress: messageInput.feeInfo.feeTokenAddress,
            amount: adjustedFeeAmount
        });
        sentMessageInfo[messageID] = SentMessageInfo({
            messageHash: keccak256(teleporterMessageBytes),
            feeInfo: adjustedFeeInfo
        });

        emit SendCrossChainMessage(
            messageID, messageInput.destinationBlockchainID, teleporterMessage, adjustedFeeInfo
        );

        // Submit the message to the AWM precompile.
        WARP_MESSENGER.sendWarpMessage(teleporterMessageBytes);

        return messageID;
    }

    /**
     * @dev Marks a message as being received by storing the message nonce associated with the
     * given message ID. The message nonce must not be zero in order to be able to distinguish between
     * received and unreceived messages based on their ID.
     */
    function _markMessageReceived(bytes32 messageID, uint256 messageNonce_) private {
        require(messageNonce_ != 0, "TeleporterMessenger: zero message nonce");
        _receivedMessageNonces[messageID] = messageNonce_;
    }

    /**
     * @dev Records the receival of a receipt for a message previously sent to the `destinationBlockchainID` with the given `messageID`.
     *
     * Returns early if a receipt was already previously received for this message, or if the message never existed. Otherwise, deletes
     * the message information from `sentMessageInfo` and increments the reward redeemable by the specified relayer reward address.
     *
     * Emits a {ReceiptReceived} event if the receipt was successfully received.
     */
    function _markReceipt(
        bytes32 sourceBlockchainID_,
        bytes32 destinationBlockchainID_,
        TeleporterMessageReceipt memory receipt
    ) private {
        bytes32 messageID = calculateMessageID(
            sourceBlockchainID_, destinationBlockchainID_, receipt.receivedMessageNonce
        );

        // Get the information about the sent message to be marked as received.
        SentMessageInfo memory messageInfo = sentMessageInfo[messageID];

        // If the message hash does not exist, it could be the case that the receipt was already
        // received for this message (it's possible for receipts to be sent more than once)
        // or that the other chain sent an invalid receipt. Return early since this is an expected
        // case where there is no fee to be paid for the given message.
        if (messageInfo.messageHash == bytes32(0)) {
            return;
        }

        // Delete the message information from state now that it is known to be received.
        delete sentMessageInfo[messageID];

        // Increment the fee/reward amount owed to the relayer for having delivered
        // the message identified in this receipt.
        _relayerRewardAmounts[receipt.relayerRewardAddress][messageInfo.feeInfo.feeTokenAddress] +=
            messageInfo.feeInfo.amount;

        emit ReceiptReceived(
            messageID, destinationBlockchainID_, receipt.relayerRewardAddress, messageInfo.feeInfo
        );
    }

    /**
     * @dev Attempts to execute the newly received message.
     *
     * Only revert in the event that the message deliverer (relayer) did not provide enough gas to handle the execution
     * (including possibly storing a failed message in state). All execution specific errors (i.e. invalid call data, etc)
     * that are not in the relayer's control are caught and handled properly.
     *
     * Emits a {MessageExecuted} event if the call on destination address is successful.
     * Emits a {MessageExecutionFailed} event if the call on destination address fails with formatted call data.
     * Requirements:
     *
     * - There is enough gas left to cover `message.requiredGasLimit`.
     */
    function _handleInitialMessageExecution(
        bytes32 messageID,
        bytes32 sourceBlockchainID,
        TeleporterMessage memory message
    ) private {
        // Check that the message delivery was provided the required gas amount as specified by the sender.
        // If the required gas amount is provided, the message will be considered received whether or not
        // its execution succeeds, such that the relayer can claim their fee reward. However, if the message
        // execution fails, the message hash will be stored in state such that anyone can try to provide more
        // gas to successfully execute the message.
        require(gasleft() >= message.requiredGasLimit, "TeleporterMessenger: insufficient gas");

        // The destination address must have fully initialized contract code in order for the message
        // to call it. If the destination address does not have code, store the message as a failed
        // execution so that it can be retried in the future should a contract be later deployed to
        // the address.
        if (message.destinationAddress.code.length == 0) {
            _storeFailedMessageExecution(messageID, sourceBlockchainID, message);
            return;
        }

        // Encode the payload by ABI encoding a call to the {receiveTeleporterMessage} function
        // defined by the {ITeleporterReceiver} interface.
        bytes memory payload = abi.encodeCall(
            ITeleporterReceiver.receiveTeleporterMessage,
            (sourceBlockchainID, message.originSenderAddress, message.message)
        );

        // Call the destination address of the message with the formatted call data. Only provide the required
        // gas limit to the sub-call so that the end application cannot consume an arbitrary amount of gas.
        bool success =
            _tryExecuteMessage(message.destinationAddress, message.requiredGasLimit, payload);

        // If the execution failed, store a hash of the message in state such that its
        // execution can be retried again in the future with a higher gas limit (paid by whoever
        // retries). Either way, the message will now be considered received since the relayer
        // provided enough gas to meet the required gas limit.
        if (!success) {
            _storeFailedMessageExecution(messageID, sourceBlockchainID, message);
            return;
        }

        emit MessageExecuted(messageID, sourceBlockchainID);
    }

    function _tryExecuteMessage(
        address target,
        uint256 gasLimit,
        bytes memory payload
    ) private returns (bool) {
        // Call the destination address of the message with the provided payload and amount of gas.
        //
        // Assembly is used for the low-level call to avoid unnecessary expansion of the return data in memory.
        // This prevents possible "return bomb" vectors where the external contract could force the caller
        // to use an arbitrary amount of gas. See Solidity issue here: https://github.com/ethereum/solidity/issues/12306
        bool success;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            success :=
                call(
                    gasLimit, // gas provided to the call
                    target, // call target
                    0, // zero value
                    add(payload, 0x20), // input data - 0x20 needs to be added to an array because the first 32-byte slot contains the array length (0x20 in hex is 32 in decimal).
                    mload(payload), // input data size - mload returns mem[p..(p+32)], which is the first 32-byte slot of the array. In this case, the array length.
                    0, // output
                    0 // output size
                )
        }
        return success;
    }

    /**
     * @dev Stores the hash of a message that has been successfully received but fails to execute properly
     * such that the message execution can be retried by anyone in the future.
     */
    function _storeFailedMessageExecution(
        bytes32 messageID,
        bytes32 sourceBlockchainID,
        TeleporterMessage memory message
    ) private {
        receivedFailedMessageHashes[messageID] = keccak256(abi.encode(message));

        // Emit a failed execution event for anyone monitoring unsuccessful messages to retry.
        emit MessageExecutionFailed(messageID, sourceBlockchainID, message);
    }

    /**
     * @dev Checks if a given message has been received.
     * @return A boolean representing if the given message has been received or not.
     */
    function _messageReceived(
        bytes32 messageID
    ) private view returns (bool) {
        return _receivedMessageNonces[messageID] != 0;
    }
}
