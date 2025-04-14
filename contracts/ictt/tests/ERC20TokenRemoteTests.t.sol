// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {ERC20TokenTransferrerTest} from "./ERC20TokenTransferrerTests.t.sol";
import {TokenRemoteTest} from "./TokenRemoteTests.t.sol";
import {IERC20SendAndCallReceiver} from "../interfaces/IERC20SendAndCallReceiver.sol";
import {TokenRemote} from "../TokenRemote/TokenRemote.sol";
import {TokenRemoteSettings} from "../TokenRemote/interfaces/ITokenRemote.sol";
import {ERC20TokenRemoteUpgradeable} from "../TokenRemote/ERC20TokenRemoteUpgradeable.sol";
import {ERC20TokenRemote} from "../TokenRemote/ERC20TokenRemote.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {SendTokensInput} from "../interfaces/ITokenTransferrer.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";

contract ERC20TokenRemoteTest is ERC20TokenTransferrerTest, TokenRemoteTest {
    using SafeERC20 for IERC20;

    string public constant MOCK_TOKEN_NAME = "Test Token";
    string public constant MOCK_TOKEN_SYMBOL = "TST";

    ERC20TokenRemoteUpgradeable public app;

    function setUp() public virtual override {
        TokenRemoteTest.setUp();

        tokenDecimals = 14;
        tokenHomeDecimals = 18;
        app = ERC20TokenRemoteUpgradeable(address(_createNewRemoteInstance()));

        erc20TokenTransferrer = app;
        tokenRemote = app;
        tokenTransferrer = app;

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(this), 10e18);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_HOME_ADDRESS,
            _encodeSingleHopSendMessage(10e18, address(this))
        );
    }

    /**
     * Initialization unit tests
     */
    function testNonUpgradeableInitialization() public {
        app = new ERC20TokenRemote(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
        assertEq(app.getBlockchainID(), DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID);
    }

    function testDisableInitialization() public {
        app = new ERC20TokenRemoteUpgradeable(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));
        app.initialize(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
    }

    function testZeroTeleporterRegistryAddress() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: address(0),
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals,
            "TeleporterRegistryApp: zero Teleporter registry address"
        );
    }

    function testZeroTeleporterManagerAddress() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(0),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals,
            abi.encodeWithSelector(Ownable.OwnableInvalidOwner.selector, address(0))
        );
    }

    function testZeroTokenHomeBlockchainID() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: bytes32(0),
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals,
            _formatErrorMessage("zero token home blockchain ID")
        );
    }

    function testDeployToSameBlockchain() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals,
            _formatErrorMessage("cannot deploy to same blockchain as token home")
        );
    }

    function testZeroTokenHomeAddress() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: address(0),
                tokenHomeDecimals: 18
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            18,
            _formatErrorMessage("zero token home address")
        );
    }

    function testSendWithSeparateFeeAsset() public {
        uint256 amount = 200_000;
        uint256 feeAmount = 100;
        ExampleERC20 separateFeeAsset = new ExampleERC20();
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFeeTokenAddress = address(separateFeeAsset);
        input.primaryFee = feeAmount;

        IERC20(separateFeeAsset).safeIncreaseAllowance(address(tokenTransferrer), feeAmount);
        vm.expectCall(
            address(separateFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom, (address(this), address(tokenTransferrer), feeAmount)
            )
        );
        // Increase the allowance of the token transferrer to transfer the funds from the user
        IERC20(app).safeIncreaseAllowance(address(tokenTransferrer), amount);

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(this), address(0), amount);
        _checkExpectedTeleporterCallsForSend(_createSingleHopTeleporterMessageInput(input, amount));
        vm.expectEmit(true, true, true, true, address(tokenTransferrer));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount);
        _send(input, amount);
    }

    function testDecimals() public view {
        uint8 res = app.decimals();
        assertEq(tokenDecimals, res);
    }

    function _createNewRemoteInstance() internal override returns (TokenRemote) {
        ERC20TokenRemoteUpgradeable instance =
            new ERC20TokenRemoteUpgradeable(ICMInitializable.Allowed);
        instance.initialize(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
        return instance;
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenRemote));
        emit TokensWithdrawn(recipient, amount);
        vm.expectEmit(true, true, true, true, address(tokenRemote));
        emit Transfer(address(0), recipient, amount);
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
        // The transferred tokens will be minted to the contract itself
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(tokenRemote), amount);

        // Then recipient contract is then approved to spend them
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        if (targetHasCode) {
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                IERC20SendAndCallReceiver.receiveTokens,
                (
                    sourceBlockchainID,
                    originInfo.tokenTransferrerAddress,
                    originInfo.senderAddress,
                    address(app),
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
        vm.expectEmit(true, true, true, true, address(app));
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
            vm.expectEmit(true, true, true, true, address(app));
            emit Transfer(address(app), address(DEFAULT_FALLBACK_RECIPIENT_ADDRESS), amount);
        }
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal virtual override {
        // Transfer the fee to the token transferrer if it is greater than 0
        if (feeAmount > 0) {
            IERC20(app).safeIncreaseAllowance(address(tokenTransferrer), feeAmount);
        }

        // Increase the allowance of the token transferrer to transfer the funds from the user
        IERC20(app).safeIncreaseAllowance(address(tokenTransferrer), amount);

        // Account for the burn before sending
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(this), address(0), amount);

        if (feeAmount > 0) {
            vm.expectEmit(true, true, true, true, address(app));
            emit Transfer(address(this), address(tokenTransferrer), feeAmount);
        }
    }

    function _getTotalSupply() internal view override returns (uint256) {
        return app.totalSupply();
    }

    function _setUpMockMint(address, uint256) internal pure override {
        // Don't need to mock the minting of an ERC20TokenRemoteUpgradeable since it is an internal call
        // on the remote contract.
        return;
    }

    function _invalidInitialization(
        TokenRemoteSettings memory settings,
        string memory tokenName,
        string memory tokenSymbol,
        uint8 tokenDecimals_,
        bytes memory expectedErrorMessage
    ) private {
        app = new ERC20TokenRemoteUpgradeable(ICMInitializable.Allowed);
        vm.expectRevert(expectedErrorMessage);
        app.initialize(settings, tokenName, tokenSymbol, tokenDecimals_);
    }
}
