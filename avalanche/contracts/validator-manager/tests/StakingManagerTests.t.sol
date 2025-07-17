// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {StakingManager} from "../StakingManager.sol";
import {DelegatorStatus, StakingManagerSettings} from "../interfaces/IStakingManager.sol";
import {IValidatorManager, ValidatorManager, ValidatorStatus} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";
import {PChainOwner} from "../ACP99Manager.sol";

abstract contract StakingManagerTest is ValidatorManagerTest {
    uint64 public constant DEFAULT_UPTIME = uint64(100);
    uint64 public constant DEFAULT_DELEGATOR_WEIGHT = uint64(1e5);
    uint64 public constant DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP =
        DEFAULT_REGISTRATION_TIMESTAMP + 1000;
    uint64 public constant DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP =
        DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1000;
    uint64 public constant DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP =
        DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION;
    address public constant DEFAULT_DELEGATOR_ADDRESS =
        address(0x1234123412341234123412341234123412341234);
    address public constant DEFAULT_VALIDATOR_OWNER_ADDRESS =
        address(0x2345234523452345234523452345234523452345);
    uint64 public constant DEFAULT_REWARD_RATE = uint64(10);
    uint64 public constant DEFAULT_MINIMUM_STAKE_DURATION = 24 hours;
    uint16 public constant DEFAULT_MINIMUM_DELEGATION_FEE_BIPS = 100;
    uint16 public constant DEFAULT_DELEGATION_FEE_BIPS = 150;
    uint8 public constant DEFAULT_MAXIMUM_STAKE_MULTIPLIER = 4;
    uint256 public constant DEFAULT_WEIGHT_TO_VALUE_FACTOR = 1e12;
    uint256 public constant SECONDS_IN_YEAR = 31536000;

    StakingManager public stakingManager;
    IRewardCalculator public rewardCalculator;

    event InitiatedDelegatorRegistration(
        bytes32 indexed delegationID,
        bytes32 indexed validationID,
        address indexed delegatorAddress,
        uint64 nonce,
        uint64 validatorWeight,
        uint64 delegatorWeight,
        bytes32 setWeightMessageID,
        address rewardRecipient
    );

    event InitiatedStakingValidatorRegistration(
        bytes32 indexed validationID,
        address indexed owner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        address rewardRecipient
    );

    event CompletedDelegatorRegistration(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime
    );

    event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID);

    event CompletedDelegatorRemoval(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees
    );

    event UptimeUpdated(bytes32 indexed validationID, uint64 uptime);

    event ValidatorRewardClaimed(
        bytes32 indexed validationID, address indexed recipient, uint256 amount
    );

    event ValidatorRewardRecipientChanged(
        bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient
    );

    event DelegatorRewardClaimed(
        bytes32 indexed delegationID, address indexed recipient, uint256 amount
    );

    event DelegatorRewardRecipientChanged(
        bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient
    );

    function testDelegationFeeBipsTooLow() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidDelegationFee.selector,
                DEFAULT_MINIMUM_DELEGATION_FEE_BIPS - 1
            )
        );
        _initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            delegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS - 1,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            stakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
            rewardRecipient: address(this)
        });
    }

    function testDelegationFeeBipsTooHigh() public {
        uint16 delegationFeeBips = stakingManager.MAXIMUM_DELEGATION_FEE_BIPS() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(StakingManager.InvalidDelegationFee.selector, delegationFeeBips)
        );

        _initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            stakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
            rewardRecipient: address(this)
        });
    }

    function testInvalidMinStakeDuration() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidMinStakeDuration.selector, DEFAULT_MINIMUM_STAKE_DURATION - 1
            )
        );
        _initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION - 1,
            stakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
            rewardRecipient: address(this)
        });
    }

    function testStakeAmountTooLow() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidStakeAmount.selector, DEFAULT_MINIMUM_STAKE_AMOUNT - 1
            )
        );
        _initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            stakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT - 1,
            rewardRecipient: address(this)
        });
    }

    function testStakeAmountTooHigh() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidStakeAmount.selector, DEFAULT_MAXIMUM_STAKE_AMOUNT + 1
            )
        );
        _initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            stakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT + 1,
            rewardRecipient: address(this)
        });
    }

    function testInvalidInitializeEndTime() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.MinStakeDurationNotPassed.selector, block.timestamp
            )
        );
        stakingManager.initiateValidatorRemoval(validationID, false, 0);
    }

    function testInvalidUptimeWarpMessage() public {
        bytes32 validationID = _registerDefaultValidator();

        _mockGetUptimeWarpMessage(new bytes(0), false);
        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        vm.expectRevert(IValidatorManager.InvalidWarpMessage.selector);
        stakingManager.initiateValidatorRemoval(validationID, true, 0);
    }

    function testInvalidUptimeSenderAddress() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    originSenderAddress: address(this),
                    payload: new bytes(0)
                }),
                true
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        vm.expectRevert(
            abi.encodeWithSelector(
                IValidatorManager.InvalidWarpOriginSenderAddress.selector, address(this)
            )
        );
        stakingManager.initiateValidatorRemoval(validationID, true, 0);
    }

    function testInvalidUptimeValidationID() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    originSenderAddress: address(0),
                    payload: ValidatorMessages.packValidationUptimeMessage(bytes32(0), 0)
                }),
                true
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.UnexpectedValidationID.selector, bytes32(0), validationID
            )
        );
        stakingManager.initiateValidatorRemoval(validationID, true, 0);
    }

    function testInvalidRewardRecipient() public {
        vm.expectRevert(
            abi.encodeWithSelector(StakingManager.InvalidRewardRecipient.selector, address(0))
        );
        _initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            stakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
            rewardRecipient: address(0)
        });
    }

    function testInvalidDelegatorRewardRecipient() public {
        bytes32 validationID = _registerDefaultValidator();
        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(StakingManager.InvalidRewardRecipient.selector, address(0))
        );
        _initiateDelegatorRegistration(
            validationID, DEFAULT_DELEGATOR_ADDRESS, DEFAULT_DELEGATOR_WEIGHT, address(0)
        );
    }

    function testInitiateDelegatorRegistration() public {
        bytes32 validationID = _registerDefaultValidator();

        _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
    }

    function testInitiateDelegatorRegistrationInactiveValidator() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        address validatorOwner = address(this);

        _completeValidatorRemovalWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: validatorOwner
        });

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                IValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.Completed
            )
        );
        _initiateDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_ADDRESS
        );
    }

    function testResendDelegatorRegistration() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes32 delegationID = _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, 1, DEFAULT_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
        );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        stakingManager.resendUpdateDelegator(delegationID);
    }

    function testCompleteDelegatorRegistration() public {
        bytes32 validationID = _registerDefaultValidator();

        _registerDefaultDelegator(validationID);
    }

    function testCompleteDelegatorRegistrationWrongNonce() public {
        bytes32 validationID = _registerDefaultValidator();

        // Initialize two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
                + DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        // Complete registration of delegator2 with delegator1's nonce
        // Note that registering delegator1 with delegator2's nonce is valid
        uint64 nonce = 1;
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, nonce, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.warp(DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP);
        vm.expectRevert(abi.encodeWithSelector(IValidatorManager.InvalidNonce.selector, nonce));
        stakingManager.completeDelegatorRegistration(delegationID2, 0);
    }

    function testCompleteDelegatorRegistrationImplicitNonce() public {
        bytes32 validationID = _registerDefaultValidator();

        // Initialize two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
        address delegator2 = address(0x5678567856785678567856785678567856785678);
        _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
                + DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
        // Mark delegator1 as registered by delivering the weight update from nonce 2 (delegator 2's nonce)
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, 2, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );

        _setUpCompleteDelegatorRegistration(
            delegationID1,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            setValidatorWeightPayload
        );
    }

    function testInitiateeEndValidationNotOwner() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.prank(address(1));
        vm.expectRevert(
            abi.encodeWithSelector(StakingManager.UnauthorizedOwner.selector, address(1))
        );
        stakingManager.initiateValidatorRemoval(validationID, false, 0);
    }

    function testInitiateDelegatorRemoval() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });
    }

    function testInitiateDelegatorRemovalByValidator() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: address(this),
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });
    }

    function testInitiateDelegatorRemovalByValidatorMinStakeDurationNotPassed() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        uint64 invalidEndTime = DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1 hours;
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.MinStakeDurationNotPassed.selector, invalidEndTime
            )
        );
        _initiateDelegatorRemoval({
            sender: address(this),
            delegationID: delegationID,
            endDelegationTimestamp: invalidEndTime,
            includeUptime: false,
            force: false
        });
    }

    function testInitiateDelegatorRemovalMinStakeDurationNotPassed() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        uint64 invalidEndTime =
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION - 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.MinStakeDurationNotPassed.selector, invalidEndTime
            )
        );
        _initiateDelegatorRemoval({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: invalidEndTime,
            includeUptime: false,
            force: false
        });
    }

    function testCompleteDelegatorRemovalChurnPeriodSecondsNotPassed() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 delegatorRegistrationTime =
            DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION + 1;
        bytes32 delegationID = _registerDelegator({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: delegatorRegistrationTime - 1,
            completeRegistrationTimestamp: delegatorRegistrationTime,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        address validatorOwner = address(this);

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: delegatorRegistrationTime + 1,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: validatorOwner
        });

        uint64 invalidEndTime = delegatorRegistrationTime + DEFAULT_CHURN_PERIOD - 1;

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.MinStakeDurationNotPassed.selector, invalidEndTime
            )
        );

        // Initialize end delegation will also call _completeDelegatorRemoval because the validator is copmleted.
        _initiateDelegatorRemoval({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: invalidEndTime,
            includeUptime: false,
            force: false
        });
    }

    function testInitiateDelegatorRemovalInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.DelegatorIneligibleForRewards.selector, delegationID
            )
        );
        vm.warp(DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP);
        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        stakingManager.initiateDelegatorRemoval(delegationID, false, 0);
    }

    function testForceInitiateDelegatorRemoval() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: true
        });
    }

    function testForceInitiateDelegatorRemovalInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: false,
            force: true
        });
    }

    function testResendEndDelegation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 2, DEFAULT_WEIGHT);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        stakingManager.resendUpdateDelegator(delegationID);
    }

    function testResendEndValidation() public override {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        validatorManager.resendValidatorRemovalMessage(validationID);
    }

    function testCompleteDelegatorRemoval() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _completeDefaultDelegator(validationID, delegationID);
    }

    function testClaimDelegationFeesInvalidValidatorStatus() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _completeDefaultDelegator(validationID, delegationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                IValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.Active
            )
        );

        stakingManager.claimDelegationFees(validationID);
    }

    function testClaimDelegationFeesInvalidSender() public {
        bytes32 validationID = _registerDefaultValidator();
        _registerDefaultDelegator(validationID);

        _endDefaultValidatorWithChecks(validationID, 2);

        vm.expectRevert(
            abi.encodeWithSelector(StakingManager.UnauthorizedOwner.selector, address(123))
        );

        vm.prank(address(123));
        stakingManager.claimDelegationFees(validationID);
    }

    function testClaimDelegationFees() public {
        bytes32 validationID = _registerDefaultValidator();
        address rewardRecipient = address(42);
        bytes32 delegationID = _registerDefaultDelegatorWithRecipient(validationID, rewardRecipient);

        _endDefaultValidatorWithChecks(validationID, 2);

        // Validator is Completed, so this will also complete the delegation.
        _initiateDelegatorRemoval({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            includeUptime: true,
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedDelegationFee = expectedTotalReward * DEFAULT_DELEGATION_FEE_BIPS / 10000;

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidatorRewardClaimed(validationID, address(this), expectedDelegationFee);
        _expectRewardIssuance(address(this), expectedDelegationFee);
        stakingManager.claimDelegationFees(validationID);
    }

    function testClaimDelegationFeesWithRewardRecipient() public {
        bytes32 validationID = _registerDefaultValidator();
        address rewardRecipient = address(42);
        bytes32 delegationID = _registerDefaultDelegatorWithRecipient(validationID, rewardRecipient);

        _endDefaultValidatorWithChecks(validationID, 2);

        // Validator is Completed, so this will also complete the delegation.
        _initiateDelegatorRemoval({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            includeUptime: true,
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedDelegationFee = expectedTotalReward * DEFAULT_DELEGATION_FEE_BIPS / 10000;

        // Change the reward recipient to a different address
        address newRewardRecipient = address(43);
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidatorRewardRecipientChanged(validationID, newRewardRecipient, address(this));
        stakingManager.changeValidatorRewardRecipient(validationID, newRewardRecipient);

        // Claim delegation fees should now go to the new reward recipient
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidatorRewardClaimed(validationID, newRewardRecipient, expectedDelegationFee);
        _expectRewardIssuance(newRewardRecipient, expectedDelegationFee);
        stakingManager.claimDelegationFees(validationID);
    }

    function testCompleteDelegatorRemovalWithNonDelegatorRewardRecipient() public {
        bytes32 validationID = _registerDefaultValidator();
        address rewardRecipient = address(42);
        bytes32 delegationID = _registerDefaultDelegatorWithRecipient(validationID, rewardRecipient);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        _completeDelegatorRemovalWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });
    }

    function testChangeDelegatorRewardRecipientWithNullAddress() public {
        bytes32 validationID = _registerDefaultValidator();
        address rewardRecipient = address(42);
        bytes32 delegationID = _registerDefaultDelegatorWithRecipient(validationID, rewardRecipient);
        address newRewardRecipient = address(0);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidRewardRecipient.selector, newRewardRecipient
            )
        );

        stakingManager.changeDelegatorRewardRecipient(delegationID, newRewardRecipient);
    }

    function testChangeDelegatorRewardRecipientByNonDelegator() public {
        bytes32 validationID = _registerDefaultValidator();
        address rewardRecipient = address(42);
        bytes32 delegationID = _registerDefaultDelegatorWithRecipient(validationID, rewardRecipient);
        address badActor = address(43);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        vm.prank(badActor);

        vm.expectRevert(abi.encodeWithSelector(StakingManager.UnauthorizedOwner.selector, badActor));

        stakingManager.changeDelegatorRewardRecipient(delegationID, badActor);
    }

    function testChangeDelegatorRewardRecipientBackToSelf() public {
        bytes32 validationID = _registerDefaultValidator();
        address rewardRecipient = address(42);
        bytes32 delegationID = _registerDefaultDelegatorWithRecipient(validationID, rewardRecipient);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit DelegatorRewardRecipientChanged(
            delegationID, DEFAULT_DELEGATOR_ADDRESS, rewardRecipient
        );

        stakingManager.changeDelegatorRewardRecipient(delegationID, DEFAULT_DELEGATOR_ADDRESS);

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        _completeDelegatorRemovalWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });
    }

    function testChangeDelegatorRewardRecipient() public {
        bytes32 validationID = _registerDefaultValidator();
        address rewardRecipient = address(42);
        bytes32 delegationID = _registerDefaultDelegatorWithRecipient(validationID, rewardRecipient);
        address newRewardRecipient = address(43);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit DelegatorRewardRecipientChanged(delegationID, newRewardRecipient, rewardRecipient);
        stakingManager.changeDelegatorRewardRecipient(delegationID, newRewardRecipient);

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        _completeDelegatorRemovalWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });
    }

    // Delegator registration is not allowed when Validator is pending removed.
    function testInitiateDelegatorRegistrationValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                IValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.PendingRemoved
            )
        );
        _initiateDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_ADDRESS
        );
    }

    // Complete delegator registration may be called when validator is pending removed.
    function testCompleteRegisterDelegatorValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes32 delegationID = _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 2, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _setUpCompleteDelegatorRegistrationWithChecks(
            validationID,
            delegationID,
            DEFAULT_COMPLETION_TIMESTAMP + 1,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1
        );
    }

    // Delegator cannot initiate end delegation when validator is pending removed.
    function testInitiateDelegatorRemovalValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                IValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.PendingRemoved
            )
        );
        _initiateDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_ADDRESS
        );
    }

    // Delegator may complete end delegation while validator is pending removed.
    function testCompleteDelegatorRemovalValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        uint64 validationEndTime = DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1;
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 3, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, validationEndTime - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: validationEndTime,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: validationEndTime - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        address delegator = DEFAULT_DELEGATOR_ADDRESS;

        _completeDelegatorRemovalWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: delegator,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: 0,
            expectedNonce: 2
        });
    }

    function testInitiateDelegatorRegistrationValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        _endDefaultValidatorWithChecks(validationID, 1);

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP + 1);
        vm.expectRevert(
            abi.encodeWithSelector(
                IValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.Completed
            )
        );
        _initiateDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_ADDRESS
        );
    }

    function testCompleteDelegatorRegistrationValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initiateDefaultDelegatorRegistration(validationID);

        _endDefaultValidatorWithChecks(validationID, 2);

        // completeDelegatorRegistration should fall through to _completeDelegatorRemoval and refund the stake
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit CompletedDelegatorRemoval(delegationID, validationID, 0, 0);

        uint256 balanceBefore = _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS);

        _expectStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, _weightToValue(DEFAULT_DELEGATOR_WEIGHT));

        // warp to right after validator ended
        vm.warp(DEFAULT_COMPLETION_TIMESTAMP + 1);
        stakingManager.completeDelegatorRegistration(delegationID, 0);

        assertEq(
            _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS),
            balanceBefore + _weightToValue(DEFAULT_DELEGATOR_WEIGHT)
        );
    }

    function testInitiateDelegatorRemovalValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _endDefaultValidatorWithChecks(validationID, 2);

        uint64 delegationEndTime = DEFAULT_COMPLETION_TIMESTAMP + 1;

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees = expectedTotalReward * DEFAULT_DELEGATION_FEE_BIPS / 10000;
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        // completeDelegatorRegistration should fall through to _completeDelegatorRemoval and refund the stake
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit CompletedDelegatorRemoval(
            delegationID, validationID, expectedDelegatorReward, expectedValidatorFees
        );

        uint256 balanceBefore = _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS);

        _expectStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, _weightToValue(DEFAULT_DELEGATOR_WEIGHT));
        _expectRewardIssuance(DEFAULT_DELEGATOR_ADDRESS, expectedDelegatorReward);

        // warp to right after validator ended
        vm.warp(delegationEndTime);
        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        stakingManager.initiateDelegatorRemoval(delegationID, false, 0);

        assertEq(
            _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS),
            balanceBefore + _weightToValue(DEFAULT_DELEGATOR_WEIGHT) + expectedDelegatorReward
        );
    }

    function testCompleteDelegatorRemovalValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        _endDefaultValidatorWithChecks(validationID, 3);

        uint256 expectedTotalReward = _defaultDelegatorExpectedTotalReward();

        uint256 expectedValidatorFees = (expectedTotalReward * DEFAULT_DELEGATION_FEE_BIPS) / 10000;
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit CompletedDelegatorRemoval(
            delegationID, validationID, expectedDelegatorReward, expectedValidatorFees
        );
        uint256 balanceBefore = _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS);

        _expectStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, _weightToValue(DEFAULT_DELEGATOR_WEIGHT));
        _expectRewardIssuance(DEFAULT_DELEGATOR_ADDRESS, expectedDelegatorReward);

        stakingManager.completeDelegatorRemoval(delegationID, 0);

        assertEq(
            _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS),
            balanceBefore + _weightToValue(DEFAULT_DELEGATOR_WEIGHT) + expectedDelegatorReward
        );
    }

    function testCompleteDelegatorRemovalWrongNonce() public {
        bytes32 validationID = _registerDefaultValidator();
        // Register two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _registerDelegator({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _registerDelegator({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT * 2 + DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        // Initialize end delegation for both delegators
        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: delegator1,
            delegationID: delegationID1,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 3,
            includeUptime: true,
            force: false
        });
        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: delegator2,
            delegationID: delegationID2,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4,
            includeUptime: true,
            force: false
        });

        // Complete ending delegator2 with delegator1's nonce
        // Note that ending delegator1 with delegator2's nonce is valid
        uint64 nonce = 3;
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, nonce, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.expectRevert(abi.encodeWithSelector(IValidatorManager.InvalidNonce.selector, nonce));
        stakingManager.completeDelegatorRemoval(delegationID2, 0);
    }

    function testCompleteDelegatorRemovalImplicitNonce() public {
        bytes32 validationID = _registerDefaultValidator();

        // Register two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _registerDelegator({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _registerDelegator({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT * 2 + DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        // Initialize end delegation for both delegators
        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: delegator1,
            delegationID: delegationID1,
            startDelegationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 3,
            includeUptime: true,
            force: false
        });
        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: delegator2,
            delegationID: delegationID2,
            startDelegationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4,
            includeUptime: true,
            force: false
        });

        uint256 expectedTotalReward = _defaultDelegatorExpectedTotalReward();

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        address delegator = DEFAULT_DELEGATOR_ADDRESS;

        // Complete delegation1 by delivering the weight update from nonce 4 (delegator2's nonce)
        _completeDelegatorRemovalWithChecks({
            validationID: validationID,
            delegationID: delegationID1,
            delegator: delegator,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4
        });
    }

    function testCompleteValidatorRemoval() public virtual override {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        address validatorOwner = address(this);

        _completeValidatorRemovalWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: validatorOwner
        });
    }

    function testReplayValidatorRegistration() public virtual override {
        uint64 initialTimestamp = uint64(block.timestamp);
        bytes32 validationID = _registerDefaultValidator();

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        address validatorOwner = address(this);

        _completeValidatorRemovalWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: validatorOwner
        });

        // Set the timestamp to be the same as when we registered the initial validator so that the
        // expiries will be the same, leading to the same validation ID.
        vm.warp(initialTimestamp);
        _beforeSend(_weightToValue(DEFAULT_WEIGHT), address(this));

        vm.expectRevert(
            abi.encodeWithSelector(IValidatorManager.InvalidValidatorStatus.selector, 4)
        );

        _initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            weight: DEFAULT_WEIGHT
        });
    }

    function testCompleteValidatorRemovalWithNonValidatorRewardRecipient() public virtual {
        address rewardRecipient = address(42);
        bytes32 validationID = _registerDefaultValidatorWithRecipient(rewardRecipient);
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        address validatorOwner = address(this);

        _completeValidatorRemovalWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: rewardRecipient
        });
    }

    function testChangeValidatorRewardRecipient() public virtual {
        address rewardRecipient = address(42);
        bytes32 validationID = _registerDefaultValidatorWithRecipient(rewardRecipient);
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address newRecipient = address(43);

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidatorRewardRecipientChanged(validationID, newRecipient, rewardRecipient);
        stakingManager.changeValidatorRewardRecipient(validationID, newRecipient);

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        _completeValidatorRemovalWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: newRecipient
        });
    }

    function testChangeValidatorRewardRecipientBackToSelf() public {
        address rewardRecipient = address(42);
        bytes32 validationID = _registerDefaultValidatorWithRecipient(rewardRecipient);
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address newRecipient = address(this);

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidatorRewardRecipientChanged(validationID, newRecipient, rewardRecipient);
        stakingManager.changeValidatorRewardRecipient(validationID, newRecipient);

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        address validatorOwner = address(this);

        _completeValidatorRemovalWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: validatorOwner
        });
    }

    function testChangeValidatorRewardRecipientWithNullAddress() public virtual {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address newRecipient = address(0);

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        vm.expectRevert(
            abi.encodeWithSelector(StakingManager.InvalidRewardRecipient.selector, newRecipient)
        );

        stakingManager.changeValidatorRewardRecipient(validationID, newRecipient);
    }

    function testChangeValidatorRewardRecipientByNonValidator() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address badActor = address(43);

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        vm.prank(badActor);

        vm.expectRevert(abi.encodeWithSelector(StakingManager.UnauthorizedOwner.selector, badActor));

        stakingManager.changeValidatorRewardRecipient(validationID, badActor);
    }

    function testInitiateValidatorRemoval() public virtual override {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });
    }

    function testInitiateValidatorRemovalUseStoredUptime() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        // Submit an uptime proof via submitUptime
        uint64 uptimePercentage1 = 80;
        uint64 uptime1 = (
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage1
        ) / 100;
        bytes memory uptimeMsg1 =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime1);
        _mockGetUptimeWarpMessage(uptimeMsg1, true);

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit UptimeUpdated(validationID, uptime1);
        stakingManager.submitUptimeProof(validationID, 0);

        // Submit a second uptime proof via initiateValidatorRemoval. This one is not sufficient for rewards
        // Submit an uptime proof via submitUptime
        uint64 uptimePercentage2 = 79;
        uint64 uptime2 = (
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage2
        ) / 100;
        bytes memory uptimeMsg2 =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime2);
        _mockGetUptimeWarpMessage(uptimeMsg2, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit InitiatedValidatorRemoval(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _initiateValidatorRemoval(validationID, true);
    }

    function testInitiateValidatorRemovalWithoutNewUptime() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        // Submit an uptime proof via submitUptime
        uint64 uptimePercentage1 = 80;
        uint64 uptime1 = (
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage1
        ) / 100;
        bytes memory uptimeMsg1 =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime1);
        _mockGetUptimeWarpMessage(uptimeMsg1, true);

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit UptimeUpdated(validationID, uptime1);
        stakingManager.submitUptimeProof(validationID, 0);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit InitiatedValidatorRemoval(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _initiateValidatorRemoval(validationID, false);
    }

    function testInitiateValidatorRemovalInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 79;

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            ((DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage)
                / 100
        );
        _mockGetUptimeWarpMessage(uptimeMsg, true);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.ValidatorIneligibleForRewards.selector, validationID
            )
        );

        _initiateValidatorRemoval(validationID, true);
    }

    function testSubmitUptimeProofPoaValidator() public {
        bytes32 defaultInitialValidationID = sha256(abi.encodePacked(DEFAULT_SUBNET_ID, uint32(1)));

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.ValidatorNotPoS.selector, defaultInitialValidationID
            )
        );
        stakingManager.submitUptimeProof(defaultInitialValidationID, 0);
    }

    function testSubmitUptimeProofInactiveValidator() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                IValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.PendingRemoved
            )
        );
        stakingManager.submitUptimeProof(validationID, 0);
    }

    function testEndValidationPoAValidator() public {
        bytes32 validationID = sha256(abi.encodePacked(DEFAULT_SUBNET_ID, uint32(1)));

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit InitiatedValidatorRemoval(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _initiateValidatorRemoval(validationID, false);

        uint256 balanceBefore = _getStakeAssetBalance(address(this));

        bytes memory l1ValidatorRegistrationMessage =
            ValidatorMessages.packL1ValidatorRegistrationMessage(validationID, false);
        _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);

        stakingManager.completeValidatorRemoval(0);

        assertEq(_getStakeAssetBalance(address(this)), balanceBefore);
    }

    function testDelegationToPoAValidator() public {
        bytes32 defaultInitialValidationID = sha256(abi.encodePacked(DEFAULT_SUBNET_ID, uint32(1)));

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.ValidatorNotPoS.selector, defaultInitialValidationID
            )
        );

        _initiateDelegatorRegistration(
            defaultInitialValidationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            address(0)
        );
    }

    function testDelegationOverWeightLimit() public {
        bytes32 validationID = _registerDefaultValidator();

        uint64 delegatorWeight = DEFAULT_WEIGHT * DEFAULT_MAXIMUM_STAKE_MULTIPLIER + 1;

        _beforeSend(_weightToValue(delegatorWeight), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.MaxWeightExceeded.selector, delegatorWeight + DEFAULT_WEIGHT
            )
        );

        _initiateDelegatorRegistration(
            validationID, DEFAULT_DELEGATOR_ADDRESS, delegatorWeight, DEFAULT_DELEGATOR_ADDRESS
        );
    }

    function testCompleteDelegatorRegistrationAlreadyRegistered() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidDelegatorStatus.selector, DelegatorStatus.Active
            )
        );

        stakingManager.completeDelegatorRegistration(delegationID, 0);
    }

    function testCompleteDelegatorRegistrationWrongValidationID() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initiateDefaultDelegatorRegistration(validationID);

        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            delegationID, 1, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        // The invalid validationID has sent no weight updates, so its nonce should be 0
        vm.expectRevert(abi.encodeWithSelector(IValidatorManager.InvalidNonce.selector, 1));

        vm.warp(DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP);
        stakingManager.completeDelegatorRegistration(delegationID, 0);
    }

    function testCompleteDelegatorRemovalWrongValidationID() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(delegationID, 2, DEFAULT_WEIGHT);
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.expectRevert(abi.encodeWithSelector(IValidatorManager.InvalidNonce.selector, 2));

        stakingManager.completeDelegatorRemoval(delegationID, 0);
    }

    function testInitiateDelegatorRemovalNotRegistered() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initiateDefaultDelegatorRegistration(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidDelegatorStatus.selector, DelegatorStatus.PendingAdded
            )
        );

        stakingManager.initiateDelegatorRemoval(delegationID, true, 0);
    }

    function testInitiateDelegatorRemovalWrongSender() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(StakingManager.UnauthorizedOwner.selector, address(123))
        );

        vm.prank(address(123));
        stakingManager.initiateDelegatorRemoval(delegationID, true, 0);
    }

    function testCompleteDelegatorRegistrationForCompletedDelegatorRegistrationWhileValidatorPendingRemoved(
    ) public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initiateDefaultDelegatorRegistration(validationID);

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 2, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _setUpCompleteDelegatorRegistrationWithChecks(
            validationID, delegationID, DEFAULT_COMPLETION_TIMESTAMP + 1, 0, 2
        );

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        _completeValidatorRemovalWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: address(this)
        });

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP + 1 + DEFAULT_MINIMUM_STAKE_DURATION);
        _expectStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, _weightToValue(DEFAULT_DELEGATOR_WEIGHT));
        stakingManager.initiateDelegatorRemoval(delegationID, true, 0);
    }

    function testCompleteDelegatorRemovalWhileActive() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidDelegatorStatus.selector, DelegatorStatus.Active
            )
        );

        stakingManager.completeDelegatorRemoval(delegationID, 0);
    }

    function testCompleteDelegatorRegistrationValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initiateDefaultDelegatorRegistration(validationID);

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 2, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );

        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _setUpCompleteDelegatorRegistrationWithChecks(
            validationID, delegationID, DEFAULT_COMPLETION_TIMESTAMP + 1, 0, 2
        );

        vm.expectRevert(
            abi.encodeWithSelector(
                IValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.PendingRemoved
            )
        );

        stakingManager.initiateDelegatorRemoval(delegationID, false, 0);
    }

    function testResendEndDelegationWhileActive() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidDelegatorStatus.selector, DelegatorStatus.Active
            )
        );

        stakingManager.resendUpdateDelegator(delegationID);
    }

    function testForceInitiateValidatorRemoval() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: true
        });
    }

    function testForceInitiateValidatorRemovalInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 79;

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            ((DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage)
                / 100
        );
        _mockGetUptimeWarpMessage(uptimeMsg, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit InitiatedValidatorRemoval(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _forceInitiateValidatorRemoval(validationID, true);
    }

    function testValueToWeightTruncated() public {
        // default weightToValueFactor is 1e12
        vm.expectRevert(abi.encodeWithSelector(StakingManager.InvalidStakeAmount.selector, 1e11));
        stakingManager.valueToWeight(1e11);
    }

    function testValueToWeightExceedsUInt64Max() public {
        // default weightToValueFactor is 1e12
        vm.expectRevert(abi.encodeWithSelector(StakingManager.InvalidStakeAmount.selector, 1e40));
        stakingManager.valueToWeight(1e40);
    }

    function testValueToWeight() public view {
        uint64 w1 = stakingManager.valueToWeight(1e12);
        uint64 w2 = stakingManager.valueToWeight(1e18);
        uint64 w3 = stakingManager.valueToWeight(1e27);

        assertEq(w1, 1);
        assertEq(w2, 1e6);
        assertEq(w3, 1e15);
    }

    function testWeightToValue() public view {
        uint256 v1 = stakingManager.weightToValue(1);
        uint256 v2 = stakingManager.weightToValue(1e6);
        uint256 v3 = stakingManager.weightToValue(1e15);

        assertEq(v1, 1e12);
        assertEq(v2, 1e18);
        assertEq(v3, 1e27);
    }

    function testStakingManagerStorageSlot() public view {
        assertEq(
            _erc7201StorageSlot("StakingManager"), stakingManager.STAKING_MANAGER_STORAGE_LOCATION()
        );
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
    ) internal virtual returns (bytes32);

    function _completeValidatorRegistration(
        uint32 messageIndex
    ) internal virtual override returns (bytes32) {
        return stakingManager.completeValidatorRegistration(messageIndex);
    }

    function _initiateValidatorRemoval(
        bytes32 validationID,
        bool includeUptime
    ) internal virtual override {
        stakingManager.initiateValidatorRemoval(validationID, includeUptime, 0);
    }

    function _forceInitiateValidatorRemoval(
        bytes32 validationID,
        bool includeUptime
    ) internal virtual override {
        stakingManager.forceInitiateValidatorRemoval(validationID, includeUptime, 0);
    }

    function _completeValidatorRemoval(
        uint32 messageIndex
    ) internal virtual override returns (bytes32) {
        return stakingManager.completeValidatorRemoval(messageIndex);
    }

    function _initiateDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight,
        address rewardRecipient
    ) internal virtual returns (bytes32);

    //
    // Staking specific setup utilities
    //
    function _beforeRegisterValidator(
        bytes32 validationID,
        address rewardRecipient
    ) internal virtual override {
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit InitiatedStakingValidatorRegistration(
            validationID,
            address(this),
            DEFAULT_DELEGATION_FEE_BIPS,
            DEFAULT_MINIMUM_STAKE_DURATION,
            rewardRecipient
        );
    }

    function _setUpInitiateDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight,
        uint64 registrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce,
        address rewardRecipient
    ) internal returns (bytes32) {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.warp(registrationTimestamp);

        _beforeSend(_weightToValue(weight), delegatorAddress);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit InitiatedValidatorWeightUpdate(
            validationID, expectedNonce, bytes32(0), expectedValidatorWeight
        );

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit InitiatedDelegatorRegistration({
            delegationID: keccak256(abi.encodePacked(validationID, expectedNonce)),
            validationID: validationID,
            delegatorAddress: delegatorAddress,
            nonce: expectedNonce,
            validatorWeight: expectedValidatorWeight,
            delegatorWeight: weight,
            setWeightMessageID: bytes32(0),
            rewardRecipient: rewardRecipient
        });

        return
            _initiateDelegatorRegistration(validationID, delegatorAddress, weight, rewardRecipient);
    }

    function _setUpCompleteDelegatorRegistration(
        bytes32 delegationID,
        uint64 completeRegistrationTimestamp,
        bytes memory setValidatorWeightPayload
    ) internal {
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.warp(completeRegistrationTimestamp);
        stakingManager.completeDelegatorRegistration(delegationID, 0);
    }

    function _setUpCompleteDelegatorRegistrationWithChecks(
        bytes32 validationID,
        bytes32 delegationID,
        uint64 completeRegistrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit CompletedValidatorWeightUpdate(validationID, expectedNonce, expectedValidatorWeight);
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit CompletedDelegatorRegistration({
            delegationID: delegationID,
            validationID: validationID,
            startTime: completeRegistrationTimestamp
        });
        _setUpCompleteDelegatorRegistration(
            delegationID, completeRegistrationTimestamp, setValidatorWeightPayload
        );
    }

    function _registerDefaultDelegator(
        bytes32 validationID
    ) internal returns (bytes32 delegationID) {
        return _registerDelegator({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
    }

    function _registerDefaultDelegatorWithRecipient(
        bytes32 validationID,
        address rewardRecipient
    ) internal returns (bytes32 delegationID) {
        return _registerDelegator({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: rewardRecipient
        });
    }

    function _initiateDefaultDelegatorRegistration(
        bytes32 validationID
    ) internal returns (bytes32) {
        return _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
    }

    function _completeDefaultDelegatorRegistration(
        bytes32 validationID,
        bytes32 delegationID
    ) internal {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, 1, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _setUpCompleteDelegatorRegistration({
            delegationID: delegationID,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            setValidatorWeightPayload: setValidatorWeightPayload
        });
    }

    function _completeDefaultDelegator(bytes32 validationID, bytes32 delegationID) internal {
        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        uint256 expectedTotalReward = _defaultDelegatorExpectedTotalReward();
        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        _completeDelegatorRemovalWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });
    }

    function _registerDelegator(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight,
        uint64 initRegistrationTimestamp,
        uint64 completeRegistrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce,
        address rewardRecipient
    ) internal returns (bytes32) {
        bytes32 delegationID = _setUpInitiateDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegatorAddress,
            weight: weight,
            registrationTimestamp: initRegistrationTimestamp,
            expectedValidatorWeight: expectedValidatorWeight,
            expectedNonce: expectedNonce,
            rewardRecipient: rewardRecipient
        });
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );

        _setUpCompleteDelegatorRegistration(
            delegationID, completeRegistrationTimestamp, setValidatorWeightPayload
        );
        return delegationID;
    }

    function _initiateDelegatorRemovalValidatorActiveWithChecks(
        bytes32 validationID,
        address sender,
        bytes32 delegationID,
        uint64 startDelegationTimestamp,
        uint64 endDelegationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce,
        bool includeUptime,
        bool force
    ) internal {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, endDelegationTimestamp - startDelegationTimestamp
        );
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit InitiatedValidatorWeightUpdate({
            validationID: validationID,
            nonce: expectedNonce,
            weight: expectedValidatorWeight,
            weightUpdateMessageID: bytes32(0)
        });

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit InitiatedDelegatorRemoval({delegationID: delegationID, validationID: validationID});

        _initiateDelegatorRemovalValidatorActive({
            sender: sender,
            delegationID: delegationID,
            endDelegationTimestamp: endDelegationTimestamp,
            includeUptime: includeUptime,
            force: force,
            setValidatorWeightPayload: setValidatorWeightPayload,
            uptimePayload: uptimeMsg
        });
    }

    function _initiateDelegatorRemovalValidatorActive(
        address sender,
        bytes32 delegationID,
        uint64 endDelegationTimestamp,
        bool includeUptime,
        bool force,
        bytes memory setValidatorWeightPayload,
        bytes memory uptimePayload
    ) internal {
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        if (includeUptime) {
            _mockGetUptimeWarpMessage(uptimePayload, true);
        }
        _initiateDelegatorRemoval({
            sender: sender,
            delegationID: delegationID,
            endDelegationTimestamp: endDelegationTimestamp,
            includeUptime: includeUptime,
            force: force
        });
    }

    function _initiateDelegatorRemoval(
        address sender,
        bytes32 delegationID,
        uint64 endDelegationTimestamp,
        bool includeUptime,
        bool force
    ) internal {
        vm.warp(endDelegationTimestamp);
        vm.prank(sender);
        if (force) {
            stakingManager.forceInitiateDelegatorRemoval(delegationID, includeUptime, 0);
        } else {
            stakingManager.initiateDelegatorRemoval(delegationID, includeUptime, 0);
        }
    }

    function _endDefaultValidatorWithChecks(bytes32 validationID, uint64 expectedNonce) internal {
        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: expectedNonce,
            rewardRecipient: address(this)
        });
    }

    function _endDefaultValidator(bytes32 validationID, uint64 expectedNonce) internal {
        address validatorOwner = address(this);
        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: expectedNonce,
            rewardRecipient: validatorOwner
        });
    }

    function _endValidationWithChecks(
        bytes32 validationID,
        address validatorOwner,
        uint64 completeRegistrationTimestamp,
        uint64 completionTimestamp,
        uint64 validatorWeight,
        uint64 expectedNonce,
        address rewardRecipient
    ) internal {
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, expectedNonce, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, completionTimestamp - completeRegistrationTimestamp
        );
        _initiateValidatorRemoval({
            validationID: validationID,
            completionTimestamp: completionTimestamp,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(validatorWeight),
            validatorStartTime: completeRegistrationTimestamp,
            stakingStartTime: completeRegistrationTimestamp,
            stakingEndTime: completionTimestamp,
            uptimeSeconds: completionTimestamp - completeRegistrationTimestamp
        });

        _completeValidatorRemovalWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: validatorWeight,
            rewardRecipient: rewardRecipient
        });
    }

    function _completeValidatorRemovalWithChecks(
        bytes32 validationID,
        address validatorOwner,
        uint256 expectedReward,
        uint64 validatorWeight,
        address rewardRecipient
    ) internal {
        bytes memory l1ValidatorRegistrationMessage =
            ValidatorMessages.packL1ValidatorRegistrationMessage(validationID, false);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit CompletedValidatorRemoval(validationID);
        uint256 balanceBefore = _getStakeAssetBalance(validatorOwner);
        uint256 rewardRecipientBalanceBefore = _getStakeAssetBalance(rewardRecipient);

        _expectStakeUnlock(validatorOwner, _weightToValue(validatorWeight));

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidatorRewardClaimed(validationID, rewardRecipient, expectedReward);
        _expectRewardIssuance(rewardRecipient, expectedReward);

        _completeValidatorRemoval(l1ValidatorRegistrationMessage);

        if (rewardRecipient == validatorOwner) {
            assertEq(
                _getStakeAssetBalance(validatorOwner),
                balanceBefore + _weightToValue(validatorWeight) + expectedReward
            );
        } else {
            assertEq(
                _getStakeAssetBalance(validatorOwner),
                balanceBefore + _weightToValue(validatorWeight)
            );

            assertEq(
                _getStakeAssetBalance(rewardRecipient),
                rewardRecipientBalanceBefore + expectedReward
            );
        }
    }

    function _completeValidatorRemoval(
        bytes memory l1ValidatorRegistrationMessage
    ) internal {
        _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);
        stakingManager.completeValidatorRemoval(0);
    }

    function _completeDelegatorRemovalWithChecks(
        bytes32 validationID,
        bytes32 delegationID,
        address delegator,
        uint64 delegatorWeight,
        uint256 expectedValidatorFees,
        uint256 expectedDelegatorReward,
        uint64 validatorWeight,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal {
        (address rewardRecipient,) = stakingManager.getDelegatorRewardInfo(delegationID);
        bytes memory weightUpdateMessage = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, validatorWeight
        );
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit CompletedValidatorWeightUpdate(validationID, expectedNonce, validatorWeight);

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit DelegatorRewardClaimed(delegationID, rewardRecipient, expectedDelegatorReward);

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit CompletedDelegatorRemoval(
            delegationID, validationID, expectedDelegatorReward, expectedValidatorFees
        );
        uint256 balanceBefore = _getStakeAssetBalance(delegator);
        uint256 rewardRecipientBalanceBefore = _getStakeAssetBalance(rewardRecipient);

        _expectStakeUnlock(delegator, _weightToValue(delegatorWeight));
        _expectRewardIssuance(rewardRecipient, expectedDelegatorReward);

        _completeDelegatorRemoval(delegationID, weightUpdateMessage);

        assertEq(validatorManager.getValidator(validationID).weight, expectedValidatorWeight);

        if (rewardRecipient == delegator) {
            assertEq(
                _getStakeAssetBalance(delegator),
                balanceBefore + _weightToValue(delegatorWeight) + expectedDelegatorReward
            );
        } else {
            assertEq(
                _getStakeAssetBalance(delegator), balanceBefore + _weightToValue(delegatorWeight)
            );

            assertEq(
                _getStakeAssetBalance(rewardRecipient),
                rewardRecipientBalanceBefore + expectedDelegatorReward
            );
        }
    }

    function _completeDelegatorRemoval(
        bytes32 delegationID,
        bytes memory weightUpdateMessage
    ) internal {
        _mockGetPChainWarpMessage(weightUpdateMessage, true);
        stakingManager.completeDelegatorRemoval(delegationID, 0);
    }

    function _initiateAndCompleteDelegatorRemovalWithChecks(
        bytes32 validationID,
        bytes32 delegationID
    ) internal {
        _initiateDelegatorRemovalValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: 0
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;
        address delegator = DEFAULT_DELEGATOR_ADDRESS;

        _completeDelegatorRemovalWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: delegator,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });
    }

    function _getStakeAssetBalance(
        address account
    ) internal virtual returns (uint256);
    function _expectStakeUnlock(address account, uint256 amount) internal virtual;
    function _expectRewardIssuance(address account, uint256 amount) internal virtual;

    function _defaultDelegatorExpectedTotalReward() internal view returns (uint256) {
        return rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });
    }

    function _defaultPoSSettings() internal pure returns (StakingManagerSettings memory) {
        return StakingManagerSettings({
            manager: ValidatorManager(address(0)),
            minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
            maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
            minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            minimumDelegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS,
            maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
            weightToValueFactor: DEFAULT_WEIGHT_TO_VALUE_FACTOR,
            rewardCalculator: IRewardCalculator(address(0)),
            uptimeBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID
        });
    }

    function _calculateValidatorFeesFromDelegator(
        uint256 totalReward,
        uint64 delegationFeeBips
    ) internal pure returns (uint256) {
        return totalReward * delegationFeeBips / 10000;
    }
}
