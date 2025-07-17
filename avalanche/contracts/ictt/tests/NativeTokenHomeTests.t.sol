// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TokenHomeTest} from "./TokenHomeTests.t.sol";
import {NativeTokenTransferrerTest} from "./NativeTokenTransferrerTests.t.sol";
import {NativeTokenHomeUpgradeable} from "../TokenHome/NativeTokenHomeUpgradeable.sol";
import {NativeTokenHome} from "../TokenHome/NativeTokenHome.sol";
import {IWrappedNativeToken} from "../interfaces/IWrappedNativeToken.sol";
import {INativeSendAndCallReceiver} from "../interfaces/INativeSendAndCallReceiver.sol";
import {WrappedNativeToken} from "../WrappedNativeToken.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";

contract NativeTokenHomeTest is NativeTokenTransferrerTest, TokenHomeTest {
    using SafeERC20 for IERC20;

    NativeTokenHomeUpgradeable public app;
    IWrappedNativeToken public wavax;

    receive() external payable {
        require(msg.sender == address(app), "NativeTokenHomeTest: invalid receive payable sender");
    }

    function setUp() public override {
        TokenHomeTest.setUp();

        wavax = new WrappedNativeToken("AVAX");
        app = new NativeTokenHomeUpgradeable(ICMInitializable.Allowed);
        app.initialize(
            MOCK_TELEPORTER_REGISTRY_ADDRESS, MOCK_TELEPORTER_MESSENGER_ADDRESS, 1, address(wavax)
        );
        tokenHome = app;
        nativeTokenTransferrer = app;
        tokenTransferrer = app;
        transferredToken = IERC20(wavax);
        tokenHomeDecimals = 18;
    }

    /**
     * Initialization unit tests
     */
    function testNonUpgradeableInitialization() public {
        app =
            new NativeTokenHome(MOCK_TELEPORTER_REGISTRY_ADDRESS, address(this), 1, address(wavax));
        assertEq(app.getBlockchainID(), DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID);
    }

    function testDisableInitialization() public {
        app = new NativeTokenHomeUpgradeable(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));
        app.initialize(MOCK_TELEPORTER_REGISTRY_ADDRESS, address(this), 1, address(wavax));
    }

    function testZeroTeleporterRegistryAddress() public {
        _invalidInitialization(
            address(0),
            address(this),
            address(wavax),
            "TeleporterRegistryApp: zero Teleporter registry address"
        );
    }

    function testZeroTeleporterManagerAddress() public {
        _invalidInitialization(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            address(wavax),
            abi.encodeWithSelector(Ownable.OwnableInvalidOwner.selector, address(0))
        );
    }

    function testZeroFeeTokenAddress() public {
        _invalidInitialization(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(0),
            _formatErrorMessage("zero token address")
        );
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenHome));
        emit TokensWithdrawn(recipient, amount);
        vm.expectCall(address(wavax), abi.encodeCall(IWrappedNativeToken.withdraw, (amount)));
        vm.expectEmit(true, true, true, true, address(wavax));
        emit Withdrawal(address(app), amount);
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
        if (targetHasCode) {
            // Non-zero code length
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                INativeSendAndCallReceiver.receiveTokens,
                (
                    sourceBlockchainID,
                    originInfo.tokenTransferrerAddress,
                    originInfo.senderAddress,
                    payload
                )
            );
            if (expectSuccess) {
                vm.mockCall(recipient, amount, expectedCalldata, new bytes(0));
            } else {
                vm.mockCallRevert(recipient, amount, expectedCalldata, new bytes(0));
            }
            vm.expectCall(recipient, amount, uint64(gasLimit), expectedCalldata);
        } else {
            // No code at target
            vm.etch(recipient, new bytes(0));
        }

        if (targetHasCode && expectSuccess) {
            vm.expectEmit(true, true, true, true, address(app));
            emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        } else {
            vm.expectEmit(true, true, true, true, address(app));
            emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        }
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal override {
        wavax.deposit{value: feeAmount}();
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

        vm.expectCall(address(transferredToken), abi.encodeCall(IWrappedNativeToken.deposit, ()));
        vm.expectEmit(true, true, true, true, address(transferredToken));
        emit Deposit(address(nativeTokenTransferrer), amount);
    }

    function _setUpDeposit(
        uint256 amount
    ) internal virtual override 
    // solhint-disable-next-line no-empty-blocks
    {}

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert("SafeWrappedNativeTokenDeposit: balance not increased");
    }

    function _addCollateral(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferrerAddress,
        uint256 amount
    ) internal override {
        app.addCollateral{value: amount}(remoteBlockchainID, remoteTokenTransferrerAddress);
    }

    function _invalidInitialization(
        address teleporterRegistryAddress,
        address teleporterManagerAddress,
        address wrappedTokenAddress,
        bytes memory expectedErrorMessage
    ) private {
        app = new NativeTokenHomeUpgradeable(ICMInitializable.Allowed);
        vm.expectRevert(expectedErrorMessage);
        app.initialize(teleporterRegistryAddress, teleporterManagerAddress, 1, wrappedTokenAddress);
    }
}
