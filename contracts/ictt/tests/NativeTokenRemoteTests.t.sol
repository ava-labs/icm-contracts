// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TokenRemoteTest} from "./TokenRemoteTests.t.sol";
import {NativeTokenTransferrerTest} from "./NativeTokenTransferrerTests.t.sol";
import {INativeSendAndCallReceiver} from "../interfaces/INativeSendAndCallReceiver.sol";
import {TokenRemote} from "../TokenRemote/TokenRemote.sol";
import {NativeTokenRemoteUpgradeable} from "../TokenRemote/NativeTokenRemoteUpgradeable.sol";
import {NativeTokenRemote} from "../TokenRemote/NativeTokenRemote.sol";
import {TokenRemoteSettings} from "../TokenRemote/interfaces/ITokenRemote.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/INativeMinter.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {SendTokensInput} from "../interfaces/ITokenTransferrer.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";

contract NativeTokenRemoteTest is NativeTokenTransferrerTest, TokenRemoteTest {
    using SafeERC20 for IERC20;

    address public constant TEST_ACCOUNT = 0xd4E96eF8eee8678dBFf4d535E033Ed1a4F7605b7;
    string public constant DEFAULT_SYMBOL = "XYZ";
    NativeTokenRemoteUpgradeable public app;

    event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned);

    function setUp() public override {
        TokenRemoteTest.setUp();

        tokenHomeDecimals = 6;
        app = NativeTokenRemoteUpgradeable(payable(address(_createNewRemoteInstance())));
        tokenRemote = app;
        nativeTokenTransferrer = app;
        tokenTransferrer = app;
        assertEq(app.totalNativeAssetSupply(), _DEFAULT_INITIAL_RESERVE_IMBALANCE);
        _collateralizeTokenTransferrer();
    }

    /**
     * Initialization unit tests
     */
    function testNonUpgradeableInitialization() public {
        app = new NativeTokenRemote({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            nativeAssetSymbol: DEFAULT_SYMBOL,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        });
        assertEq(app.getBlockchainID(), DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID);
    }

    function testDisableInitialization() public {
        app = new NativeTokenRemoteUpgradeable(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));
        app.initialize({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            nativeAssetSymbol: DEFAULT_SYMBOL,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        });
    }

    function testZeroInitialReserveImbalance() public {
        _invalidInitialization({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            nativeAssetSymbol: DEFAULT_SYMBOL,
            initialReserveImbalance: 0,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE,
            expectedErrorMessage: "NativeTokenRemote: zero initial reserve imbalance"
        });
    }

    function testInvalidBurnedRewardPercentage() public {
        _invalidInitialization({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            nativeAssetSymbol: DEFAULT_SYMBOL,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            burnedFeesReportingRewardPercentage: 100,
            expectedErrorMessage: "NativeTokenRemote: invalid percentage"
        });
    }

    function testZeroTokenHomeBlockchainID() public {
        _invalidInitialization({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: bytes32(0),
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            nativeAssetSymbol: DEFAULT_SYMBOL,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE,
            expectedErrorMessage: _formatErrorMessage("zero token home blockchain ID")
        });
    }

    function testDeployToSameBlockchain() public {
        _invalidInitialization({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            nativeAssetSymbol: DEFAULT_SYMBOL,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE,
            expectedErrorMessage: _formatErrorMessage("cannot deploy to same blockchain as token home")
        });
    }

    function testSendBeforeCollateralized() public {
        // Need a new instance since the default set up pre-collateralizes the contract.
        app = NativeTokenRemoteUpgradeable(payable(address(_createNewRemoteInstance())));
        tokenRemote = app;
        nativeTokenTransferrer = app;
        tokenTransferrer = app;

        vm.expectRevert("NativeTokenRemote: contract undercollateralized");
        // solhint-disable-next-line check-send-result
        app.send{value: 1e17}(_createDefaultSendTokensInput());

        // Now mark the contract as collateralized and confirm sending is enabled.
        _collateralizeTokenTransferrer();
        _sendSingleHopSendSuccess(1e17, 0);
    }

    function testSendAndCallBeforeCollateralized() public {
        // Need a new instance since the default set up pre-collateralizes the contract.
        app = NativeTokenRemoteUpgradeable(payable(address(_createNewRemoteInstance())));
        tokenRemote = app;
        nativeTokenTransferrer = app;
        tokenTransferrer = app;

        vm.expectRevert("NativeTokenRemote: contract undercollateralized");
        app.sendAndCall{value: 1e15}(_createDefaultSendAndCallInput());

        // Now mark the contract as collateralized and confirm sending is enabled.
        _collateralizeTokenTransferrer();
        _sendSingleHopSendSuccess(1e15, 0);
    }

    function testSendWithSeparateFeeAsset() public {
        uint256 amount = 2e15;
        uint256 feeAmount = 100;
        ExampleERC20 separateFeeAsset = new ExampleERC20();
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFeeTokenAddress = address(separateFeeAsset);
        input.primaryFee = feeAmount;

        IERC20(separateFeeAsset).safeIncreaseAllowance(address(app), feeAmount);
        vm.expectCall(
            address(separateFeeAsset),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(app), feeAmount))
        );

        _checkExpectedTeleporterCallsForSend(_createSingleHopTeleporterMessageInput(input, amount));
        vm.expectEmit(true, true, true, true, address(app));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount);
        _send(input, amount);
    }

    function testTotalNativeAssetSupply() public {
        uint256 initialTotalMinted = app.getTotalMinted();
        uint256 initialExpectedBalance = _DEFAULT_INITIAL_RESERVE_IMBALANCE + initialTotalMinted;
        assertEq(app.totalNativeAssetSupply(), initialExpectedBalance);

        // Mock tokens being burned as tx fees.
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), initialExpectedBalance - 1);
        assertEq(app.totalNativeAssetSupply(), 1);

        // Reset the burned tx fee amount.
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), 0);
        assertEq(app.totalNativeAssetSupply(), initialExpectedBalance);

        // Mock tokens being transferred out by crediting them to the native token remote contract
        vm.deal(app.BURNED_FOR_TRANSFER_ADDRESS(), initialExpectedBalance - 1);
        assertEq(app.totalNativeAssetSupply(), 1);

        // Depositing native tokens into the contract to be wrapped native tokens shouldn't affect the supply
        // of the native asset, but should be reflected in the total supply of the ERC20 representation.
        app.deposit{value: 2}();
        assertEq(app.totalNativeAssetSupply(), 1);
        assertEq(app.totalSupply(), 2);
    }

    function testTransferToHome() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = _DEFAULT_TRANSFER_AMOUNT;
        vm.expectEmit(true, true, true, true, address(app));
        emit TokensSent({
            teleporterMessageID: _MOCK_MESSAGE_ID,
            sender: address(this),
            input: input,
            amount: amount
        });

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationTokenTransferrerAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(app), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(amount, input.recipient)
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );
        // solhint-disable-next-line check-send-result
        app.send{value: amount}(input);
    }

    function testReceiveSendAndCallFailureInsufficientValue() public {
        uint256 amount = 200;
        bytes memory payload = hex"DEADBEEF";
        OriginSenderInfo memory originInfo;
        originInfo.tokenTransferrerAddress = address(app);
        originInfo.senderAddress = address(this);
        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        _setUpMockMint(address(app), amount);
        vm.deal(address(app), amount - 1);
        vm.expectRevert("CallUtils: insufficient value");
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenRemote.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, DEFAULT_TOKEN_HOME_ADDRESS, message
        );
    }

    function testReportBurnFeesNoNewAmount() public {
        vm.expectRevert("NativeTokenRemote: burn address balance not greater than last report");
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testReportBurnFeesScaledToZero() public {
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), 1);
        vm.expectRevert("NativeTokenRemote: zero scaled amount to report burn");
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testReportBurnFeesSuccess() public {
        uint256 initialBurnedTxFeeAmount = 1e19;
        uint256 expectedReward = initialBurnedTxFeeAmount / 100; // 1% of 1e17
        uint256 expectedReportedAmount = initialBurnedTxFeeAmount - expectedReward;
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), initialBurnedTxFeeAmount);

        _setUpMockMint(address(app), expectedReward);
        TeleporterMessageInput memory expectedMessageInput = _createSingleHopTeleporterMessageInput(
            SendTokensInput({
                destinationBlockchainID: app.getTokenHomeBlockchainID(),
                destinationTokenTransferrerAddress: app.getTokenHomeAddress(),
                recipient: app.HOME_CHAIN_BURN_ADDRESS(),
                primaryFeeTokenAddress: address(app),
                primaryFee: expectedReward,
                secondaryFee: 0,
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
                multiHopFallback: address(0)
            }),
            expectedReportedAmount
        );
        _checkExpectedTeleporterCallsForSend(expectedMessageInput);
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);

        // Calling it again should revert since no additional amount as been burned.
        vm.expectRevert("NativeTokenRemote: burn address balance not greater than last report");
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);

        // Mock more transaction fees being burned.
        uint256 additionalBurnTxFeeAmount = 5 * 1e15 + 3;
        expectedReward = additionalBurnTxFeeAmount / 100; // 1%, rounded down due to integer division.
        expectedReportedAmount = additionalBurnTxFeeAmount - expectedReward;
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), initialBurnedTxFeeAmount + additionalBurnTxFeeAmount);

        _setUpMockMint(address(app), expectedReward);
        expectedMessageInput = _createSingleHopTeleporterMessageInput(
            SendTokensInput({
                destinationBlockchainID: app.getTokenHomeBlockchainID(),
                destinationTokenTransferrerAddress: app.getTokenHomeAddress(),
                recipient: app.HOME_CHAIN_BURN_ADDRESS(),
                primaryFeeTokenAddress: address(app),
                primaryFee: expectedReward,
                secondaryFee: 0,
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
                multiHopFallback: address(0)
            }),
            expectedReportedAmount
        );
        _checkExpectedTeleporterCallsForSend(expectedMessageInput);
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testReportBurnFeesNoRewardSuccess() public {
        // Create a new TokenRemote instance with no rewards for reporting burned fees.
        app = new NativeTokenRemoteUpgradeable(ICMInitializable.Allowed);
        app.initialize({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            nativeAssetSymbol: DEFAULT_SYMBOL,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            burnedFeesReportingRewardPercentage: 0
        });
        tokenRemote = app;
        nativeTokenTransferrer = app;
        tokenTransferrer = app;

        uint256 burnedTxFeeAmount = 1e15;
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), burnedTxFeeAmount);
        TeleporterMessageInput memory expectedMessageInput = _createSingleHopTeleporterMessageInput(
            SendTokensInput({
                destinationBlockchainID: app.getTokenHomeBlockchainID(),
                destinationTokenTransferrerAddress: app.getTokenHomeAddress(),
                recipient: app.HOME_CHAIN_BURN_ADDRESS(),
                primaryFeeTokenAddress: address(app),
                primaryFee: 0,
                secondaryFee: 0,
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
                multiHopFallback: address(0)
            }),
            burnedTxFeeAmount
        );
        _checkExpectedTeleporterCallsForSend(expectedMessageInput);
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testFallback() public {
        (bool success,) = address(app).call{value: 1}("1234567812345678");
        assertTrue(success);
        assertEq(app.balanceOf(address(this)), 1);
    }

    function testDepositWithdrawWrappedNativeToken() public {
        uint256 depositAmount = 500;
        uint256 withdrawAmount = 100;
        vm.deal(TEST_ACCOUNT, depositAmount);
        vm.startPrank(TEST_ACCOUNT);
        app.deposit{value: depositAmount}();
        assertEq(app.balanceOf(TEST_ACCOUNT), depositAmount);
        app.withdraw(withdrawAmount);
        assertEq(app.balanceOf(TEST_ACCOUNT), depositAmount - withdrawAmount);
        assertEq(TEST_ACCOUNT.balance, withdrawAmount);
    }

    function _createNewRemoteInstance() internal override returns (TokenRemote) {
        NativeTokenRemoteUpgradeable instance =
            new NativeTokenRemoteUpgradeable(ICMInitializable.Allowed);
        instance.initialize({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            nativeAssetSymbol: DEFAULT_SYMBOL,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        });
        return instance;
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenRemote));
        emit TokensWithdrawn(recipient, amount);
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, amount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, amount))
        );
        vm.deal(recipient, amount);
    }

    function _setUpMockMint(address recipient, uint256 amount) internal override {
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, amount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, amount))
        );
        vm.deal(recipient, amount);
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
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (address(app), amount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (address(app), amount))
        );

        // Mock the native minter precompile crediting native balance to the contract.
        vm.deal(address(app), amount);

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

    function _setUpExpectedDeposit(uint256, uint256 feeAmount) internal override {
        app.deposit{value: feeAmount}();
        // Transfer the fee to the token transferrer if it is greater than 0
        if (feeAmount > 0) {
            IERC20(app).safeIncreaseAllowance(address(tokenTransferrer), feeAmount);
        }

        if (feeAmount > 0) {
            vm.expectEmit(true, true, true, true, address(app));
            emit Transfer(address(this), address(tokenTransferrer), feeAmount);
        }
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
    }

    function _getTotalSupply() internal view override returns (uint256) {
        return app.totalNativeAssetSupply();
    }

    // The native token remote contract is considered collateralized once it has received
    // a message from its configured home to mint tokens. Until then, the home contract is
    // still assumed to have insufficient collateral.
    function _collateralizeTokenTransferrer() private {
        assertFalse(app.getIsCollateralized());
        uint256 amount = 10e18;
        _setUpMockMint(DEFAULT_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            app.getTokenHomeBlockchainID(),
            app.getTokenHomeAddress(),
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );
        assertTrue(app.getIsCollateralized());
    }

    function _invalidInitialization(
        TokenRemoteSettings memory settings,
        string memory nativeAssetSymbol,
        uint256 initialReserveImbalance,
        uint256 burnedFeesReportingRewardPercentage,
        bytes memory expectedErrorMessage
    ) private {
        app = new NativeTokenRemoteUpgradeable(ICMInitializable.Allowed);
        vm.expectRevert(expectedErrorMessage);
        app.initialize(
            settings,
            nativeAssetSymbol,
            initialReserveImbalance,
            burnedFeesReportingRewardPercentage
        );
    }
}
