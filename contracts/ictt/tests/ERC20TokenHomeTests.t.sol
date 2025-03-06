// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ERC20TokenTransferrerTest} from "./ERC20TokenTransferrerTests.t.sol";
import {TokenHomeTest} from "./TokenHomeTests.t.sol";
import {IERC20SendAndCallReceiver} from "../interfaces/IERC20SendAndCallReceiver.sol";
import {SendTokensInput} from "../interfaces/ITokenTransferrer.sol";
import {ERC20TokenHomeUpgradeable} from "../TokenHome/ERC20TokenHomeUpgradeable.sol";
import {ERC20TokenHome} from "../TokenHome/ERC20TokenHome.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TokenScalingUtils} from "@utilities/TokenScalingUtils.sol";
import {RemoteTokenTransferrerSettings} from "../TokenHome/interfaces/ITokenHome.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";

contract ERC20TokenHomeTest is ERC20TokenTransferrerTest, TokenHomeTest {
    using SafeERC20 for IERC20;

    ERC20TokenHomeUpgradeable public app;
    IERC20 public mockERC20;

    function setUp() public override {
        TokenHomeTest.setUp();

        mockERC20 = new ExampleERC20();
        tokenHomeDecimals = 6;
        app = new ERC20TokenHomeUpgradeable(ICMInitializable.Allowed);
        app.initialize(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            1,
            address(mockERC20),
            tokenHomeDecimals
        );
        erc20TokenTransferrer = app;
        tokenHome = app;
        tokenTransferrer = app;

        transferredToken = mockERC20;
    }

    /**
     * Initialization unit tests
     */
    function testNonUpgradeableInitialization() public {
        app = new ERC20TokenHome(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            1,
            address(mockERC20),
            tokenHomeDecimals
        );
        assertEq(app.getBlockchainID(), DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID);
    }

    function testDisableInitialization() public {
        app = new ERC20TokenHomeUpgradeable(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));
        app.initialize(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            1,
            address(mockERC20),
            tokenHomeDecimals
        );
    }

    function testZeroTeleporterRegistryAddress() public {
        _invalidInitialization(
            address(0),
            address(this),
            address(mockERC20),
            tokenHomeDecimals,
            "TeleporterRegistryApp: zero Teleporter registry address"
        );
    }

    function testZeroTeleporterManagerAddress() public {
        _invalidInitialization(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            address(mockERC20),
            tokenHomeDecimals,
            abi.encodeWithSelector(Ownable.OwnableInvalidOwner.selector, address(0))
        );
    }

    function testZeroFeeTokenAddress() public {
        _invalidInitialization(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(0),
            tokenHomeDecimals,
            _formatErrorMessage("zero token address")
        );
    }

    function testTokenDecimalsTooHigh() public {
        _invalidInitialization(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(mockERC20),
            uint8(TokenScalingUtils.MAX_TOKEN_DECIMALS) + 1,
            _formatErrorMessage("token decimals too high")
        );
    }

    function testReceiveZeroHomeTokenAmount() public {
        // Set up a registered remote that will scale down the received amount
        // to zero home tokens.
        uint256 tokenMultiplier = 100_000;
        _setUpRegisteredRemote(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            0,
            tokenMultiplier,
            true
        );

        // Send over home token to the remote
        // and check for expected calls for scaled amount of tokens sent.
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = 1;
        _setUpExpectedDeposit(amount, input.primaryFee);

        uint256 scaledAmount = tokenMultiplier * amount;
        _checkExpectedTeleporterCallsForSend(
            _createSingleHopTeleporterMessageInput(input, scaledAmount)
        );
        vm.expectEmit(true, true, true, true, address(tokenTransferrer));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, scaledAmount);
        _send(input, amount);

        // Receive an amount from remote less than `scaledAmount`
        // which will be scaled down to zero home tokens.
        vm.expectRevert(_formatErrorMessage("zero token amount"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeSingleHopSendMessage(scaledAmount - 1, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testRegisterDestinationRoundUpCollateralNeeded() public {
        _setUpRegisteredRemote(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 11, 10, true
        );
        RemoteTokenTransferrerSettings memory settings = tokenHome.getRemoteTokenTransferrerSettings(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS
        );
        assertEq(settings.collateralNeeded, 2);
    }

    function testSendScaledUpAmount() public {
        uint256 amount = 100;
        uint256 feeAmount = 1;

        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        // Raw amount sent over wire should be multipled by 1e2.
        uint256 tokenMultiplier = 1e2;
        _setUpRegisteredRemote(
            input.destinationBlockchainID,
            input.destinationTokenTransferrerAddress,
            0,
            tokenMultiplier,
            true
        );
        _setUpExpectedDeposit(amount, input.primaryFee);
        TeleporterMessageInput memory expectedMessage = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationTokenTransferrerAddress,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: address(transferredToken),
                amount: input.primaryFee
            }),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(amount * tokenMultiplier, DEFAULT_RECIPIENT_ADDRESS)
        });
        _checkExpectedTeleporterCallsForSend(expectedMessage);
        vm.expectEmit(true, true, true, true, address(tokenTransferrer));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount * tokenMultiplier);
        _send(input, amount);
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenHome));
        emit TokensWithdrawn(recipient, amount);
        vm.expectCall(
            address(mockERC20), abi.encodeCall(IERC20.transfer, (address(recipient), amount))
        );
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Transfer(address(app), recipient, amount);
    }

    function _setUpExpectedSendAndCall(
        bytes32 sourceBlockchainID,
        OriginSenderInfo memory originInfo,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal override {
        // The recipient contract will be approved to spend the tokens from the token home contract.
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        if (targetHasCode) {
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                IERC20SendAndCallReceiver.receiveTokens,
                (
                    sourceBlockchainID,
                    originInfo.tokenTransferrerAddress,
                    originInfo.senderAddress,
                    address(mockERC20),
                    amount,
                    payload
                )
            );
            if (expectSuccess) {
                vm.mockCall(recipient, expectedCalldata, new bytes(0));
            } else {
                vm.mockCallRevert(recipient, expectedCalldata, new bytes(0));
            }
            vm.expectCall(recipient, 0, uint64(gasLimit), expectedCalldata);
        } else {
            vm.etch(recipient, new bytes(0));
        }

        // Then recipient contract approval is reset
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, 0);

        if (targetHasCode && expectSuccess) {
            // The call should have succeeded.
            vm.expectEmit(true, true, true, true, address(app));
            emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        } else {
            // The call should have failed.
            vm.expectEmit(true, true, true, true, address(app));
            emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

            // Then the amount should be sent to the fallback recipient.
            vm.expectEmit(true, true, true, true, address(mockERC20));
            emit Transfer(address(app), address(DEFAULT_FALLBACK_RECIPIENT_ADDRESS), amount);
        }
    }

    function _addCollateral(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount
    ) internal override {
        app.addCollateral(remoteBlockchainID, remoteTokenTransferrerAddress, amount);
    }

    function _setUpDeposit(
        uint256 amount
    ) internal virtual override {
        // Increase the allowance of the token transferrer to transfer the funds from the user
        transferredToken.safeIncreaseAllowance(address(tokenTransferrer), amount);
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert("SafeERC20TransferFrom: balance not increased");
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal virtual override {
        // Transfer the fee to the token transferrer if it is greater than 0
        if (feeAmount > 0) {
            transferredToken.safeIncreaseAllowance(address(tokenTransferrer), feeAmount);
            vm.expectCall(
                address(transferredToken),
                abi.encodeCall(
                    IERC20.transferFrom, (address(this), address(tokenTransferrer), feeAmount)
                )
            );
        }
        // Increase the allowance of the token transferrer to transfer the funds from the user
        transferredToken.safeIncreaseAllowance(address(tokenTransferrer), amount);

        // Check that transferFrom is called to deposit the funds sent from the user to the token transferrer
        // This is the case because for the {ERC20TokenHomeUpgradeable) is not the fee token itself
        vm.expectCall(
            address(transferredToken),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenTransferrer), amount))
        );
        vm.expectEmit(true, true, true, true, address(transferredToken));
        emit Transfer(address(this), address(tokenTransferrer), amount);
    }

    function _invalidInitialization(
        address teleporterRegistryAddress,
        address teleporterManagerAddress,
        address feeTokenAddress,
        uint8 tokenDecimals,
        bytes memory expectedErrorMessage
    ) private {
        app = new ERC20TokenHomeUpgradeable(ICMInitializable.Allowed);
        vm.expectRevert(expectedErrorMessage);
        app.initialize(
            teleporterRegistryAddress, teleporterManagerAddress, 1, feeTokenAddress, tokenDecimals
        );
    }
}
