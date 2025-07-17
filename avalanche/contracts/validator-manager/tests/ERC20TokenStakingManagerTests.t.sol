// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {StakingManagerTest} from "./StakingManagerTests.t.sol";
import {ERC20TokenStakingManager} from "../ERC20TokenStakingManager.sol";
import {StakingManager, StakingManagerSettings} from "../StakingManager.sol";
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {IERC20Mintable} from "../interfaces/IERC20Mintable.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {IACP99Manager, PChainOwner, ConversionData} from "../interfaces/IACP99Manager.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";

contract ERC20TokenStakingManagerTest is StakingManagerTest {
    using SafeERC20 for IERC20Mintable;

    ERC20TokenStakingManager public app;
    IERC20Mintable public token;

    function setUp() public override {
        ValidatorManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();

        ConversionData memory conversion = _defaultConversionData();
        bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
        _mockInitializeValidatorSet(conversionID);
        validatorManager.initializeValidatorSet(conversion, 0);
    }

    //
    // Initialization unit tests
    // The pattern in these tests requires that only non-admin validator manager functions are called,
    // as each test re-deploys the ERC20TokenStakingManager contract.
    //
    function testDisableInitialization() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        app.initialize(defaultPoSSettings, token);
    }

    function testZeroTokenAddress() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(StakingManager.ZeroAddress.selector));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        app.initialize(defaultPoSSettings, IERC20Mintable(address(0)));
    }

    function testZeroRewardCalculatorAddress() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(StakingManager.ZeroAddress.selector));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        app.initialize(defaultPoSSettings, token);
    }

    function testZeroManagerddress() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(StakingManager.ZeroAddress.selector));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        app.initialize(defaultPoSSettings, token);
    }

    function testZeroMinimumDelegationFee() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(StakingManager.InvalidDelegationFee.selector, 0));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        defaultPoSSettings.minimumDelegationFeeBips = 0;
        app.initialize(defaultPoSSettings, token);
    }

    function testMaxMinimumDelegationFee() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        uint16 minimumDelegationFeeBips = app.MAXIMUM_DELEGATION_FEE_BIPS() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidDelegationFee.selector, minimumDelegationFeeBips
            )
        );

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        defaultPoSSettings.minimumDelegationFeeBips = minimumDelegationFeeBips;
        app.initialize(defaultPoSSettings, token);
    }

    function testInvalidStakeAmountRange() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidStakeAmount.selector, DEFAULT_MAXIMUM_STAKE_AMOUNT
            )
        );

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        defaultPoSSettings.minimumStakeAmount = DEFAULT_MAXIMUM_STAKE_AMOUNT;
        defaultPoSSettings.maximumStakeAmount = DEFAULT_MINIMUM_STAKE_AMOUNT;
        app.initialize(defaultPoSSettings, token);
    }

    function testZeroMaxStakeMultiplier() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(StakingManager.InvalidStakeMultiplier.selector, 0));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        defaultPoSSettings.maximumStakeMultiplier = 0;
        app.initialize(defaultPoSSettings, token);
    }

    function testMinStakeDurationTooLow() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        uint64 minimumStakeDuration = DEFAULT_CHURN_PERIOD - 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidMinStakeDuration.selector, minimumStakeDuration
            )
        );

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        defaultPoSSettings.minimumStakeDuration = minimumStakeDuration;
        app.initialize(defaultPoSSettings, token);
    }

    function testMaxStakeMultiplierOverLimit() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        uint8 maximumStakeMultiplier = app.MAXIMUM_STAKE_MULTIPLIER_LIMIT() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidStakeMultiplier.selector, maximumStakeMultiplier
            )
        );

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        defaultPoSSettings.maximumStakeMultiplier = maximumStakeMultiplier;
        app.initialize(defaultPoSSettings, token);
    }

    function testZeroWeightToValueFactor() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(StakingManager.ZeroWeightToValueFactor.selector));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        defaultPoSSettings.weightToValueFactor = 0;
        app.initialize(defaultPoSSettings, token);
    }

    function testInvalidValidatorManager() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        ValidatorManager invalidManager = ValidatorManager(address(token)); // The contract type is arbitrary

        vm.expectRevert();

        StakingManagerSettings memory settings = _defaultPoSSettings();
        settings.manager = invalidManager;
        app.initialize(settings, token);
    }

    function testUnsetValidatorManager() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert();

        app.initialize(_defaultPoSSettings(), token); // settings.manager is not set
    }

    function testInvalidValidatorMinStakeDuration() public {
        uint256 stakeAmount = _weightToValue(DEFAULT_WEIGHT);
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidMinStakeDuration.selector, DEFAULT_MINIMUM_STAKE_DURATION - 1
            )
        );
        app.initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION - 1,
            stakeAmount: stakeAmount,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
    }

    function testERC20TokenStakingManagerStorageSlot() public view {
        assertEq(
            _erc7201StorageSlot("ERC20TokenStakingManager"),
            app.ERC20_STAKING_MANAGER_STORAGE_LOCATION()
        );
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight,
        address rewardRecipient
    ) internal virtual override returns (bytes32) {
        return app.initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            stakeAmount: _weightToValue(weight),
            rewardRecipient: rewardRecipient
        });
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount,
        address rewardRecipient
    ) internal virtual override returns (bytes32) {
        return app.initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration,
            stakeAmount: stakeAmount,
            rewardRecipient: rewardRecipient
        });
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            stakeAmount: _weightToValue(weight),
            rewardRecipient: address(this)
        });
    }

    function _initiateDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight,
        address rewardRecipient
    ) internal virtual override returns (bytes32) {
        uint256 value = _weightToValue(weight);
        vm.startPrank(delegatorAddress);
        bytes32 delegationID =
            app.initiateDelegatorRegistration(validationID, value, rewardRecipient);
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

    function _setUp() internal override returns (IACP99Manager) {
        // Construct the object under test
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        token = new ExampleERC20();
        rewardCalculator = new ExampleRewardCalculator(DEFAULT_REWARD_RATE);
        validatorManager = new ValidatorManager(ICMInitializable.Allowed);

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        defaultPoSSettings.manager = validatorManager;

        validatorManager.initialize(_defaultSettings(address(app)));
        app.initialize(defaultPoSSettings, token);

        stakingManager = app;

        return validatorManager;
    }

    function _getStakeAssetBalance(
        address account
    ) internal view override returns (uint256) {
        return token.balanceOf(account);
    }
}
