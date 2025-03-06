// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {TeleporterRegistry} from "@teleporter/registry/TeleporterRegistry.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {
    ITokenTransferrer,
    SendTokensInput,
    SendAndCallInput,
    TransferrerMessageType,
    TransferrerMessage,
    SingleHopSendMessage,
    SingleHopCallMessage,
    MultiHopSendMessage,
    MultiHopCallMessage
} from "../interfaces/ITokenTransferrer.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {IERC20Errors} from "@openzeppelin/contracts@5.0.2/interfaces/draft-IERC6093.sol";

abstract contract TokenTransferrerTest is Test {
    // convenience struct to reduce stack usage
    struct OriginSenderInfo {
        address tokenTransferrerAddress;
        address senderAddress;
    }

    bytes32 public constant DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant OTHER_BLOCKCHAIN_ID =
        bytes32(hex"9876987698769876987698769876987698769876987698769876987698769876");
    address public constant DEFAULT_TOKEN_REMOTE_ADDRESS =
        0xd878229c9c3575F224784DE610911B5607a3ad15;
    address public constant DEFAULT_TOKEN_HOME_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_SENDER_ADDRESS = 0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5;
    address public constant DEFAULT_RECIPIENT_ADDRESS = 0xABCDabcdABcDabcDaBCDAbcdABcdAbCdABcDABCd;
    address public constant DEFAULT_RECIPIENT_CONTRACT_ADDRESS =
        0xa83114A443dA1CecEFC50368531cACE9F37fCCcb;
    uint256 public constant DEFAULT_REQUIRED_GAS_LIMIT = 200_000;
    uint256 public constant DEFAULT_RECIPIENT_GAS_LIMIT = 100_000;
    address public constant DEFAULT_MULTIHOP_FALLBACK_ADDRESS =
        0x043448b3FE2F24522D9CeB32AD8623c0b6b53E26;
    address public constant DEFAULT_FALLBACK_RECIPIENT_ADDRESS =
        0xe69Ea1BAF997002602c0A3D451b2b5c9B7F8E6A1;
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    address public constant NATIVE_MINTER_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000001);
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");

    uint256 internal constant _DEFAULT_FEE_AMOUNT = 123456;
    uint256 internal constant _DEFAULT_TRANSFER_AMOUNT = 1e18;
    uint256 internal constant _DEFAULT_INITIAL_RESERVE_IMBALANCE = 1e18;
    uint256 internal constant _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE = 1;

    uint8 public tokenHomeDecimals;

    ITokenTransferrer public tokenTransferrer;

    event TokensSent(
        bytes32 indexed teleporterMessageID,
        address indexed sender,
        SendTokensInput input,
        uint256 amount
    );
    event TokensRouted(bytes32 indexed teleporterMessageID, SendTokensInput input, uint256 amount);
    event TokensAndCallSent(
        bytes32 indexed teleporterMessageID,
        address indexed sender,
        SendAndCallInput input,
        uint256 amount
    );
    event TokensAndCallRouted(
        bytes32 indexed teleporterMessageID, SendAndCallInput input, uint256 amount
    );

    event TokensWithdrawn(address indexed recipient, uint256 amount);
    event CallSucceeded(address indexed recipientContract, uint256 amount);
    event CallFailed(address indexed recipientContract, uint256 amount);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    function testSendZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        vm.expectRevert(_formatErrorMessage("zero recipient address"));
        _send(input, 0);
    }

    function testSendAndCallZeroRecipientContract() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.recipientContract = address(0);
        vm.expectRevert(_formatErrorMessage("zero recipient contract address"));
        _sendAndCall(input, 0);
    }

    function testSendZeroRequiredGasLimit() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.requiredGasLimit = 0;
        vm.expectRevert(_formatErrorMessage("zero required gas limit"));
        _send(input, 0);
    }

    function testSendAndCallZeroRequiredGasLimit() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.requiredGasLimit = 0;
        vm.expectRevert(_formatErrorMessage("zero required gas limit"));
        _sendAndCall(input, 0);
    }

    function testSendAndCallZeroRecipientGasLimit() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.recipientGasLimit = 0;
        vm.expectRevert(_formatErrorMessage("zero recipient gas limit"));
        _sendAndCall(input, 0);
    }

    function testSendAndCallInvalidRecipientGasLimit() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.recipientGasLimit = input.requiredGasLimit + 1;
        vm.expectRevert(_formatErrorMessage("invalid recipient gas limit"));
        _sendAndCall(input, 0);
    }

    function testSendAndCallZeroFallbackRecipient() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.fallbackRecipient = address(0);
        vm.expectRevert(_formatErrorMessage("zero fallback recipient address"));
        _sendAndCall(input, 0);
    }

    function testSendWithNoFeeAllowance() public {
        uint256 amount = 2e15;
        uint256 primaryFee = 100;

        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = primaryFee;

        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        _setUpExpectedDeposit(amount, 0);
        vm.expectRevert(
            abi.encodeWithSelector(
                IERC20Errors.ERC20InsufficientAllowance.selector, tokenTransferrer, 0, primaryFee
            )
        );
        _send(input, amount);
    }

    function testSendAndCallWithNoFeeAllowance() public {
        uint256 amount = 2e15;
        uint256 primaryFee = 100;

        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.primaryFee = primaryFee;

        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        _setUpExpectedDeposit(amount, 0);
        vm.expectRevert(
            abi.encodeWithSelector(
                IERC20Errors.ERC20InsufficientAllowance.selector, tokenTransferrer, 0, primaryFee
            )
        );
        _sendAndCall(input, amount);
    }

    function testSendWithFees() public {
        uint256 amount = 2e15;
        uint256 primaryFee = 100;
        _sendSingleHopSendSuccess(amount, primaryFee);
    }

    function testSendNoFees() public {
        uint256 amount = 2e15;
        uint256 primaryFee = 0;
        _sendSingleHopSendSuccess(amount, primaryFee);
    }

    function testSendAndCallWithFees() public {
        uint256 amount = 1e17;
        uint256 primaryFee = 10;
        _sendSingleHopCallSuccess(amount, primaryFee);
    }

    function _initMockTeleporterRegistry() internal {
        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion.selector
            ),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getVersionFromAddress.selector),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getAddressFromVersion.selector, (1)),
            abi.encode(MOCK_TELEPORTER_MESSENGER_ADDRESS)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getLatestTeleporter.selector),
            abi.encode(ITeleporterMessenger(MOCK_TELEPORTER_MESSENGER_ADDRESS))
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getVersionFromAddress.selector),
            abi.encode(1)
        );
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual;

    function _sendAndCall(SendAndCallInput memory input, uint256 amount) internal virtual;

    function _sendSingleHopSendSuccess(uint256 amount, uint256 feeAmount) internal {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        _setUpExpectedDeposit(amount, input.primaryFee);
        _checkExpectedTeleporterCallsForSend(_createSingleHopTeleporterMessageInput(input, amount));
        vm.expectEmit(true, true, true, true, address(tokenTransferrer));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount);
        _send(input, amount);
    }

    function _sendSingleHopCallSuccess(uint256 amount, uint256 feeAmount) internal {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.primaryFee = feeAmount;

        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        _setUpExpectedDeposit(amount, input.primaryFee);
        OriginSenderInfo memory originInfo;
        originInfo.tokenTransferrerAddress = address(tokenTransferrer);
        originInfo.senderAddress = address(this);
        _checkExpectedTeleporterCallsForSend(
            _createSingleHopCallTeleporterMessageInput(
                _getDefaultMessageSourceBlockchainID(), originInfo, input, amount
            )
        );
        vm.expectEmit(true, true, true, true, address(tokenTransferrer));
        emit TokensAndCallSent(_MOCK_MESSAGE_ID, address(this), input, amount);
        _sendAndCall(input, amount);
    }

    function _setUpRegisteredRemote(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 initialReserveImbalance
    ) internal virtual;

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal virtual;

    function _setUpExpectedZeroAmountRevert() internal virtual;

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal virtual;

    function _checkExpectedTeleporterCallsForSend(
        TeleporterMessageInput memory expectedMessageInput
    ) internal {
        if (expectedMessageInput.feeInfo.amount > 0) {
            vm.expectCall(
                expectedMessageInput.feeInfo.feeTokenAddress,
                abi.encodeCall(
                    IERC20.allowance,
                    (address(tokenTransferrer), address(MOCK_TELEPORTER_MESSENGER_ADDRESS))
                )
            );
        }
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );
    }

    // This function is overridden by TeleporterTokenDestinationTests
    function _scaleTokens(uint256 amount, bool) internal virtual returns (uint256) {
        return amount;
    }

    function _createDefaultSendTokensInput()
        internal
        view
        virtual
        returns (SendTokensInput memory);

    function _createDefaultSendAndCallInput()
        internal
        view
        virtual
        returns (SendAndCallInput memory);

    function _createSingleHopTeleporterMessageInput(
        SendTokensInput memory input,
        uint256 transferAmount
    ) internal pure returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationTokenTransferrerAddress,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: address(input.primaryFeeTokenAddress),
                amount: input.primaryFee
            }),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(transferAmount, input.recipient)
        });
    }

    function _createSingleHopCallTeleporterMessageInput(
        bytes32 sourceBlockchainID,
        OriginSenderInfo memory originInfo,
        SendAndCallInput memory input,
        uint256 transferAmount
    ) internal pure returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationTokenTransferrerAddress,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: address(input.primaryFeeTokenAddress),
                amount: input.primaryFee
            }),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopCallMessage({
                sourceBlockchainID: sourceBlockchainID,
                originInfo: originInfo,
                amount: transferAmount,
                recipientContract: input.recipientContract,
                recipientPayload: input.recipientPayload,
                recipientGasLimit: input.recipientGasLimit,
                fallbackRecipient: input.fallbackRecipient
            })
        });
    }

    function _getDefaultMessageSourceBlockchainID() internal pure virtual returns (bytes32);

    function _formatErrorMessage(
        string memory message
    ) internal pure virtual returns (bytes memory);

    function _encodeSingleHopSendMessage(
        uint256 amount,
        address recipient
    ) internal pure returns (bytes memory) {
        return abi.encode(
            TransferrerMessage({
                messageType: TransferrerMessageType.SINGLE_HOP_SEND,
                payload: abi.encode(SingleHopSendMessage({recipient: recipient, amount: amount}))
            })
        );
    }

    function _encodeSingleHopCallMessage(
        bytes32 sourceBlockchainID,
        OriginSenderInfo memory originInfo,
        uint256 amount,
        address recipientContract,
        bytes memory recipientPayload,
        uint256 recipientGasLimit,
        address fallbackRecipient
    ) internal pure returns (bytes memory) {
        return abi.encode(
            TransferrerMessage({
                messageType: TransferrerMessageType.SINGLE_HOP_CALL,
                payload: abi.encode(
                    SingleHopCallMessage({
                        sourceBlockchainID: sourceBlockchainID,
                        originTokenTransferrerAddress: originInfo.tokenTransferrerAddress,
                        originSenderAddress: originInfo.senderAddress,
                        recipientContract: recipientContract,
                        amount: amount,
                        recipientPayload: recipientPayload,
                        recipientGasLimit: recipientGasLimit,
                        fallbackRecipient: fallbackRecipient
                    })
                )
            })
        );
    }

    function _encodeMultiHopSendMessage(
        uint256 amount,
        bytes32 destinationBlockchainID,
        address destinationTokenTransferrerAddress,
        address recipient,
        uint256 secondaryFee,
        uint256 secondaryGasLimit,
        address multiHopFallback
    ) internal pure returns (bytes memory) {
        return abi.encode(
            TransferrerMessage({
                messageType: TransferrerMessageType.MULTI_HOP_SEND,
                payload: abi.encode(
                    MultiHopSendMessage({
                        destinationBlockchainID: destinationBlockchainID,
                        destinationTokenTransferrerAddress: destinationTokenTransferrerAddress,
                        recipient: recipient,
                        amount: amount,
                        secondaryFee: secondaryFee,
                        secondaryGasLimit: secondaryGasLimit,
                        multiHopFallback: multiHopFallback
                    })
                )
            })
        );
    }

    function _encodeMultiHopCallMessage(
        address originSenderAddress,
        uint256 amount,
        bytes32 destinationBlockchainID,
        address destinationTokenTransferrerAddress,
        address recipientContract,
        bytes memory recipientPayload,
        uint256 recipientGasLimit,
        address fallbackRecipient,
        address multiHopFallback,
        uint256 secondaryRequiredGasLimit,
        uint256 secondaryFee
    ) internal pure returns (bytes memory) {
        return abi.encode(
            TransferrerMessage({
                messageType: TransferrerMessageType.MULTI_HOP_CALL,
                payload: abi.encode(
                    MultiHopCallMessage({
                        originSenderAddress: originSenderAddress,
                        destinationBlockchainID: destinationBlockchainID,
                        destinationTokenTransferrerAddress: destinationTokenTransferrerAddress,
                        recipientContract: recipientContract,
                        amount: amount,
                        recipientPayload: recipientPayload,
                        recipientGasLimit: recipientGasLimit,
                        fallbackRecipient: fallbackRecipient,
                        multiHopFallback: multiHopFallback,
                        secondaryRequiredGasLimit: secondaryRequiredGasLimit,
                        secondaryFee: secondaryFee
                    })
                )
            })
        );
    }
}
