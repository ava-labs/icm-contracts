// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoSValidatorManagerTest} from "./PoSValidatorManagerTests.t.sol";
import {ERC20TokenStakingManager} from "../ERC20TokenStakingManager.sol";
import {PoSValidatorManager, PoSValidatorManagerSettings} from "../PoSValidatorManager.sol";
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
import {ValidatorRegistrationInput} from "../interfaces/IValidatorManager.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {IERC20Mintable} from "../interfaces/IERC20Mintable.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {ACP99Manager} from "../ACP99Manager.sol";

contract ERC20TokenStakingManagerTest is PoSValidatorManagerTest {
    using SafeERC20 for IERC20Mintable;

    ERC20TokenStakingManager public app;
    IERC20Mintable public token;

    function setUp() public override {
        ValidatorManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();
        _mockInitializeValidatorSet();
        app.initializeValidatorSet(_defaultConversionData(), 0);
    }

    function testDisableInitialization() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));
        app.initialize(_defaultPoSSettings(), token);
    }

    function testZeroTokenAddress() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(
                ERC20TokenStakingManager.InvalidTokenAddress.selector, address(0)
            )
        );
        app.initialize(_defaultPoSSettings(), IERC20Mintable(address(0)));
    }

    function testZeroMinimumDelegationFee() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.InvalidDelegationFee.selector, 0)
        );

        PoSValidatorManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.minimumDelegationFeeBips = 0;
        app.initialize(defaultPoSSettings, token);
    }

    function testMaxMinimumDelegationFee() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        uint16 minimumDelegationFeeBips = app.MAXIMUM_DELEGATION_FEE_BIPS() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegationFee.selector, minimumDelegationFeeBips
            )
        );

        PoSValidatorManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.minimumDelegationFeeBips = minimumDelegationFeeBips;
        app.initialize(defaultPoSSettings, token);
    }

    function testInvalidStakeAmountRange() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidStakeAmount.selector, DEFAULT_MAXIMUM_STAKE_AMOUNT
            )
        );

        PoSValidatorManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.minimumStakeAmount = DEFAULT_MAXIMUM_STAKE_AMOUNT;
        defaultPoSSettings.maximumStakeAmount = DEFAULT_MINIMUM_STAKE_AMOUNT;
        app.initialize(defaultPoSSettings, token);
    }

    function testZeroMaxStakeMultiplier() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.InvalidStakeMultiplier.selector, 0)
        );

        PoSValidatorManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.maximumStakeMultiplier = 0;
        app.initialize(defaultPoSSettings, token);
    }

    function testMinStakeDurationTooLow() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        uint64 minimumStakeDuration = DEFAULT_CHURN_PERIOD - 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidMinStakeDuration.selector, minimumStakeDuration
            )
        );

        PoSValidatorManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.minimumStakeDuration = minimumStakeDuration;
        app.initialize(defaultPoSSettings, token);
    }

    function testMaxStakeMultiplierOverLimit() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        uint8 maximumStakeMultiplier = app.MAXIMUM_STAKE_MULTIPLIER_LIMIT() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidStakeMultiplier.selector, maximumStakeMultiplier
            )
        );

        PoSValidatorManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.maximumStakeMultiplier = maximumStakeMultiplier;
        app.initialize(defaultPoSSettings, token);
    }

    function testZeroWeightToValueFactor() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.ZeroWeightToValueFactor.selector)
        );

        PoSValidatorManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.weightToValueFactor = 0;
        app.initialize(defaultPoSSettings, token);
    }

    function testInvalidValidatorMinStakeDuration() public {
        ValidatorRegistrationInput memory input = ValidatorRegistrationInput({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationExpiry: DEFAULT_EXPIRY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });
        uint256 stakeAmount = _weightToValue(DEFAULT_WEIGHT);
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidMinStakeDuration.selector,
                DEFAULT_MINIMUM_STAKE_DURATION - 1
            )
        );
        app.initializeValidatorRegistration(
            input, DEFAULT_DELEGATION_FEE_BIPS, DEFAULT_MINIMUM_STAKE_DURATION - 1, stakeAmount
        );
    }

    function testERC20TokenStakingManagerStorageSlot() public view {
        assertEq(
            _erc7201StorageSlot("ERC20TokenStakingManager"),
            app.ERC20_STAKING_MANAGER_STORAGE_LOCATION()
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory registrationInput,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration(
            registrationInput, delegationFeeBips, minStakeDuration, stakeAmount
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration(
            input,
            DEFAULT_DELEGATION_FEE_BIPS,
            DEFAULT_MINIMUM_STAKE_DURATION,
            _weightToValue(weight)
        );
    }

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        uint256 value = _weightToValue(weight);
        vm.startPrank(delegatorAddress);
        bytes32 delegationID = app.initializeDelegatorRegistration(validationID, value);
        vm.stopPrank();
        return delegationID;
    }

    function _beforeSend(uint256 amount, address spender) internal override {
        token.safeIncreaseAllowance(spender, amount);
        token.safeTransfer(spender, amount);

        // ERC20 tokens need to be pre-approved
        vm.startPrank(spender);
        token.safeIncreaseAllowance(address(app), amount);
        vm.stopPrank();
    }

    function _expectStakeUnlock(address account, uint256 amount) internal override {
        vm.expectCall(address(token), abi.encodeCall(IERC20.transfer, (account, amount)));
    }

    function _expectRewardIssuance(address account, uint256 amount) internal override {
        vm.expectCall(address(token), abi.encodeCall(IERC20Mintable.mint, (account, amount)));
    }

    function _setUp() internal override returns (ACP99Manager) {
        // Construct the object under test
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        token = new ExampleERC20();
        rewardCalculator = new ExampleRewardCalculator(DEFAULT_REWARD_RATE);

        PoSValidatorManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        app.initialize(defaultPoSSettings, token);

        validatorManager = app;
        posValidatorManager = app;

        return app;
    }

    function _getStakeAssetBalance(address account) internal view override returns (uint256) {
        return token.balanceOf(account);
    }
}
