// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TokenTransferrerTest} from "./TokenTransferrerTests.t.sol";
import {TokenRemote, IWarpMessenger} from "../TokenRemote/TokenRemote.sol";
import {TeleporterRegistry} from "@teleporter/registry/TeleporterRegistry.sol";
import {SendTokensInput, SendAndCallInput} from "../interfaces/ITokenTransferrer.sol";
import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";
import {TokenScalingUtils} from "@utilities/TokenScalingUtils.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    TransferrerMessageType,
    TransferrerMessage,
    RegisterRemoteMessage
} from "../interfaces/ITokenTransferrer.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

abstract contract TokenRemoteTest is TokenTransferrerTest {
    using SafeERC20 for IERC20;

    TokenRemote public tokenRemote;
    uint8 public tokenDecimals = 18;

    function setUp() public virtual {
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(ITeleporterMessenger.sendCrossChainMessage.selector),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.getBlockchainID, ()),
            abi.encode(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID)
        );
        vm.expectCall(WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getBlockchainID, ()));

        _initMockTeleporterRegistry();

        vm.expectCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeCall(TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion, ())
        );
    }

    function testZeroDestinationBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = bytes32(0);
        vm.expectRevert(_formatErrorMessage("zero destination blockchain ID"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testZeroDestinationBlockchainWithSendAndCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = bytes32(0);
        vm.expectRevert(_formatErrorMessage("zero destination blockchain ID"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testZeroDestinationTokenTransferrerAddress() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationTokenTransferrerAddress = address(0);
        vm.expectRevert(_formatErrorMessage("zero destination token transferrer address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testZeroDestinationTokenTransferrerAddressWithSendAndCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationTokenTransferrerAddress = address(0);
        vm.expectRevert(_formatErrorMessage("zero destination token transferrer address"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testInvalidSendingBackToHomeBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationTokenTransferrerAddress = address(this);
        vm.expectRevert(_formatErrorMessage("invalid destination token transferrer address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testInvalidSendAndCallingBackToHomeBlockchain() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationTokenTransferrerAddress = address(this);
        vm.expectRevert(_formatErrorMessage("invalid destination token transferrer address"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroSecondaryFeeToHomeBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.secondaryFee = 1;
        vm.expectRevert(_formatErrorMessage("non-zero secondary fee"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroSecondaryFeeToHomeBlockchainCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.secondaryFee = 1;
        vm.expectRevert(_formatErrorMessage("non-zero secondary fee"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroMultiHopFallbackForSingleHop() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("non-zero multi-hop fallback"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroMultiHopFallbackForSingleHopCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("non-zero multi-hop fallback"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendingToSameInstance() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenRemote.getBlockchainID();
        input.destinationTokenTransferrerAddress = address(tokenRemote);
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("invalid destination token transferrer address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendingToOtherInstanceOnSameChain() public {
        _sendMultiHopSendSuccess(tokenRemote.getBlockchainID(), 1e18, 999, 555);
    }

    function testSendAndCallingToSameInstance() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = tokenRemote.getBlockchainID();
        input.destinationTokenTransferrerAddress = address(tokenRemote);
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("invalid destination token transferrer address"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendAndCallingToOtherInstanceOnSameChain() public {
        _sendMultiHopCallSuccess(tokenRemote.getBlockchainID(), 1e18, 999, 555);
    }

    function testSendZeroAmountAfterRemoveScaling() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.primaryFee = 0;
        input.secondaryFee = 0;
        uint256 amount = 1;

        if (
            TokenScalingUtils.removeTokenScale(
                tokenRemote.getTokenMultiplier(), tokenRemote.getMultiplyOnRemote(), amount
            ) != 0
        ) {
            return;
        }

        _setUpExpectedDeposit(amount, input.primaryFee);
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
        _sendAndCall(input, amount);
    }

    function testSendToSameBlockchainDifferentDestination() public {
        uint256 amount = 2e15;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenRemote.getBlockchainID();
        // Set the desintation token transferrer address to an address different than the token remote contract.
        input.destinationTokenTransferrerAddress = address(0x55);

        _sendSingleHopSendSuccess(amount, input.primaryFee);
    }

    function testSendMultiHopInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendMultiHopInput();
        uint256 amount = 1;
        input.secondaryFee = 2;
        _setUpExpectedDeposit(amount, input.primaryFee);
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
        _send(input, amount);
    }

    function testSendMultiHopZeroMultiHopFallback() public {
        uint256 amount = 200_000;
        SendTokensInput memory input = _createDefaultSendMultiHopInput();
        input.multiHopFallback = address(0);
        vm.expectRevert(_formatErrorMessage("zero multi-hop fallback"));
        _send(input, amount);
    }

    function testSendAndCallMultiHopZeroMultiHopFallback() public {
        uint256 amount = 200_000;
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = OTHER_BLOCKCHAIN_ID;
        input.multiHopFallback = address(0);
        vm.expectRevert(_formatErrorMessage("zero multi-hop fallback"));
        _sendAndCall(input, amount);
    }

    function testSendMultiHopSendSuccess() public {
        uint256 amount = 4e15;
        uint256 primaryFee = 5;
        uint256 secondaryFee = 2;
        _sendMultiHopSendSuccess(OTHER_BLOCKCHAIN_ID, amount, primaryFee, secondaryFee);
    }

    function testSendMultiHopCallSuccess() public {
        uint256 amount = 4e18;
        uint256 primaryFee = 5;
        uint256 secondaryFee = 1;
        _sendMultiHopCallSuccess(OTHER_BLOCKCHAIN_ID, amount, primaryFee, secondaryFee);
    }

    function testReceiveInvalidSourceBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("invalid source blockchain ID"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_HOME_ADDRESS, new bytes(0)
        );
    }

    function testReceiveInvalidOriginSenderAddress() public {
        vm.expectRevert(_formatErrorMessage("invalid origin sender address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, address(0), new bytes(0)
        );
    }

    function testReceiveInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, DEFAULT_TOKEN_HOME_ADDRESS, bytes("test")
        );
    }

    function testReceiveWithdrawSuccess() public {
        uint256 amount = 200;
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, amount);
        uint256 initialSupply = _getTotalSupply();
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_HOME_ADDRESS,
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );
        assertEq(_getTotalSupply(), initialSupply + amount);
    }

    function testReceiveSendAndCallSuccess() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        bytes32 sourceBlockchainID = DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID;
        OriginSenderInfo memory originInfo;
        originInfo.senderAddress = address(this);
        _setUpExpectedSendAndCall({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: true
        });

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        uint256 initialSupply = _getTotalSupply();

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, DEFAULT_TOKEN_HOME_ADDRESS, message
        );

        assertEq(_getTotalSupply(), initialSupply + amount);
    }

    function testReceiveSendAndCallFailure() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        OriginSenderInfo memory originInfo;
        originInfo.tokenTransferrerAddress = address(tokenTransferrer);
        originInfo.senderAddress = address(this);
        _setUpExpectedSendAndCall({
            sourceBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            originInfo: originInfo,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: false
        });

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, DEFAULT_TOKEN_HOME_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailureNoCode() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        bytes32 sourceBlockchainID = DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID;
        OriginSenderInfo memory originInfo;
        originInfo.tokenTransferrerAddress = address(tokenTransferrer);
        originInfo.senderAddress = address(this);
        _setUpExpectedSendAndCall({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: false,
            expectSuccess: false
        });

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, DEFAULT_TOKEN_HOME_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailureInsufficientGas() public {
        uint256 amount = 200;
        bytes memory payload = hex"DEADBEEF";
        uint256 gasLimit = 5_000_000;
        OriginSenderInfo memory originInfo;
        originInfo.tokenTransferrerAddress = address(tokenTransferrer);
        originInfo.senderAddress = address(this);

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: gasLimit,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        _setUpMockMint(address(tokenRemote), amount);
        vm.expectRevert("CallUtils: insufficient gas");
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage{gas: gasLimit - 1}(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, DEFAULT_TOKEN_HOME_ADDRESS, message
        );
    }

    function testReceiveInvalidMessageType() public {
        vm.expectRevert(_formatErrorMessage("invalid message type"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_HOME_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: 1,
                destinationBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
                destinationTokenTransferrerAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 1_000,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
            })
        );
    }

    function testRegisterWithHomeSuccess() public {
        // Create a new instance that has not yet received any messages.
        tokenRemote = _createNewRemoteInstance();
        tokenTransferrer = tokenRemote;

        // Deploy a separate fee asset for registering with the token home.
        ExampleERC20 separateFeeAsset = new ExampleERC20();
        uint256 feeAmount = 13;
        TeleporterFeeInfo memory feeInfo =
            TeleporterFeeInfo({feeTokenAddress: address(separateFeeAsset), amount: feeAmount});

        IERC20(separateFeeAsset).safeIncreaseAllowance(address(tokenTransferrer), feeAmount);
        vm.expectCall(
            address(separateFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom, (address(this), address(tokenTransferrer), feeAmount)
            )
        );

        TransferrerMessage memory expectedTokenTransfer = TransferrerMessage({
            messageType: TransferrerMessageType.REGISTER_REMOTE,
            payload: abi.encode(
                RegisterRemoteMessage({
                    initialReserveImbalance: tokenRemote.getInitialReserveImbalance(),
                    remoteTokenDecimals: tokenDecimals,
                    homeTokenDecimals: tokenHomeDecimals
                })
            )
        });
        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: tokenRemote.getTokenHomeBlockchainID(),
            destinationAddress: tokenRemote.getTokenHomeAddress(),
            feeInfo: feeInfo,
            requiredGasLimit: tokenRemote.REGISTER_REMOTE_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: abi.encode(expectedTokenTransfer)
        });
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        tokenRemote.registerWithHome(feeInfo);
    }

    function testRegisterWithHomeAlreadyRegistered() public {
        // Mock receiving a message from the home so that the remote knows
        // that it is registered already.
        uint256 amount = 1;
        _setUpMockMint(DEFAULT_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            tokenRemote.getTokenHomeBlockchainID(),
            tokenRemote.getTokenHomeAddress(),
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );

        vm.expectRevert(_formatErrorMessage("already registered"));
        tokenRemote.registerWithHome(TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}));
    }

    function testCalculateNumWords() public view {
        assertEq(tokenRemote.calculateNumWords(0), 0);
        assertEq(tokenRemote.calculateNumWords(1), 1);
        assertEq(tokenRemote.calculateNumWords(32), 1);
        assertEq(tokenRemote.calculateNumWords(33), 2);
        assertEq(tokenRemote.calculateNumWords(64), 2);
        assertEq(tokenRemote.calculateNumWords(65), 3);
    }

    function _sendMultiHopSendSuccess(
        bytes32 destinationBlockchainID,
        uint256 amount,
        uint256 primaryFee,
        uint256 secondaryFee
    ) internal {
        uint256 transferAmount = amount;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = destinationBlockchainID;
        input.primaryFee = primaryFee;
        input.secondaryFee = secondaryFee;
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;

        _setUpExpectedDeposit(amount, input.primaryFee);
        _checkExpectedTeleporterCallsForSend(
            _createMultiHopSendTeleporterMessageInput(input, transferAmount)
        );
        vm.expectEmit(true, true, true, true, address(tokenTransferrer));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, transferAmount);
        _send(input, amount);
    }

    function _sendMultiHopCallSuccess(
        bytes32 destinationBlockchainID,
        uint256 amount,
        uint256 primaryFee,
        uint256 secondaryFee
    ) internal {
        uint256 transferAmount = amount;
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = destinationBlockchainID;
        input.primaryFee = primaryFee;
        input.secondaryFee = secondaryFee;
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;

        _setUpExpectedDeposit(amount, input.primaryFee);

        // Only tokens destinations scale tokens, so isReceive is always false here.
        address originSenderAddress = address(this);

        _checkExpectedTeleporterCallsForSend(
            _createMultiHopCallTeleporterMessageInput(originSenderAddress, input, transferAmount)
        );
        vm.expectEmit(true, true, true, true, address(tokenTransferrer));
        emit TokensAndCallSent(_MOCK_MESSAGE_ID, originSenderAddress, input, transferAmount);
        _sendAndCall(input, amount);
    }

    function _setUpMockMint(address recipient, uint256 amount) internal virtual;

    function _setUpExpectedSendAndCall(
        bytes32 sourceBlockchainID,
        OriginSenderInfo memory originInfo,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal virtual;

    // Remotes don't need to register supported remotes because they
    // only send messages to their configured token home.
    function _setUpRegisteredRemote(bytes32, address, uint256) internal virtual override {
        return;
    }

    function _createNewRemoteInstance() internal virtual returns (TokenRemote);

    function _getTotalSupply() internal view virtual returns (uint256);

    function _createMultiHopSendTeleporterMessageInput(
        SendTokensInput memory input,
        uint256 transferAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenRemote.getTokenHomeBlockchainID(),
            destinationAddress: tokenRemote.getTokenHomeAddress(),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(tokenRemote), amount: input.primaryFee}),
            requiredGasLimit: tokenRemote.MULTI_HOP_SEND_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: _encodeMultiHopSendMessage({
                amount: transferAmount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationTokenTransferrerAddress: input.destinationTokenTransferrerAddress,
                recipient: input.recipient,
                secondaryFee: input.secondaryFee,
                secondaryGasLimit: input.requiredGasLimit,
                multiHopFallback: input.multiHopFallback
            })
        });
    }

    function _createMultiHopCallTeleporterMessageInput(
        address originSenderAddress,
        SendAndCallInput memory input,
        uint256 transferAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenRemote.getTokenHomeBlockchainID(),
            destinationAddress: tokenRemote.getTokenHomeAddress(),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(tokenRemote), amount: input.primaryFee}),
            requiredGasLimit: tokenRemote.MULTI_HOP_CALL_REQUIRED_GAS()
                + (
                    tokenRemote.calculateNumWords(input.recipientPayload.length)
                        * tokenRemote.MULTI_HOP_CALL_GAS_PER_WORD()
                ),
            allowedRelayerAddresses: new address[](0),
            message: _encodeMultiHopCallMessage({
                originSenderAddress: originSenderAddress,
                amount: transferAmount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationTokenTransferrerAddress: input.destinationTokenTransferrerAddress,
                recipientContract: input.recipientContract,
                recipientPayload: input.recipientPayload,
                recipientGasLimit: input.recipientGasLimit,
                fallbackRecipient: input.fallbackRecipient,
                multiHopFallback: input.multiHopFallback,
                secondaryRequiredGasLimit: input.requiredGasLimit,
                secondaryFee: input.secondaryFee
            })
        });
    }

    function _createDefaultSendTokensInput()
        internal
        view
        override
        returns (SendTokensInput memory)
    {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            destinationTokenTransferrerAddress: DEFAULT_TOKEN_HOME_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFeeTokenAddress: address(tokenRemote),
            primaryFee: 0,
            secondaryFee: 0,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            multiHopFallback: address(0)
        });
    }

    function _createDefaultSendMultiHopInput() internal view returns (SendTokensInput memory) {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = OTHER_BLOCKCHAIN_ID;
        input.multiHopFallback = DEFAULT_FALLBACK_RECIPIENT_ADDRESS;
        return input;
    }

    function _createDefaultSendAndCallInput()
        internal
        view
        override
        returns (SendAndCallInput memory)
    {
        return SendAndCallInput({
            destinationBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            destinationTokenTransferrerAddress: DEFAULT_TOKEN_HOME_ADDRESS,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: new bytes(16),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
            multiHopFallback: address(0),
            primaryFeeTokenAddress: address(tokenRemote),
            primaryFee: 0,
            secondaryFee: 0
        });
    }

    function _getDefaultMessageSourceBlockchainID() internal pure override returns (bytes32) {
        return DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID;
    }

    function _formatErrorMessage(
        string memory message
    ) internal pure override returns (bytes memory) {
        return bytes(string.concat("TokenRemote: ", message));
    }
}
