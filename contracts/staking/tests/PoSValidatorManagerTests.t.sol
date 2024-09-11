// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {PoSValidatorManager} from "../PoSValidatorManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";

abstract contract PoSValidatorManagerTest is ValidatorManagerTest {
    uint64 public constant DEFAULT_UPTIME = uint64(100);
    uint64 public constant DEFAULT_DELEGATOR_WEIGHT = uint64(1e5);
    uint64 public constant DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP = uint64(2000);
    uint64 public constant DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP = uint64(3000);
    uint64 public constant DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP = uint64(4000);
    address public constant DEFAULT_DELEGATOR_ADDRESS =
        address(0x1234123412341234123412341234123412341234);
    uint64 public constant DEFAULT_REWARD_RATE = uint64(10);

    PoSValidatorManager public posValidatorManager;
    ExampleRewardCalculator public rewardCalculator;

    event DelegatorAdded(
        bytes32 indexed delegationID,
        bytes32 indexed validationID,
        address indexed delegatorAddress,
        uint64 nonce,
        uint64 validatorWeight,
        uint64 delegatorWeight,
        bytes32 setWeightMessageID
    );

    event DelegatorRegistered(
        bytes32 indexed delegationID,
        bytes32 indexed validationID,
        uint64 indexed nonce,
        uint256 startTime
    );

    event DelegatorRemovalInitialized(
        bytes32 indexed delegationID,
        bytes32 indexed validationID,
        uint64 indexed nonce,
        uint64 validatorWeight,
        uint256 endTime,
        bytes32 setWeightMessageID
    );

    event DelegationEnded(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint64 indexed nonce
    );

    function testInitializeEndValidation() public override {
        _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            includeUptime: true
        });
    }

    function testInvalidUptimeWarpMessage() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _mockGetVerifiedWarpMessage(new bytes(0), false);
        vm.expectRevert(_formatErrorMessage("invalid warp message"));
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeChainID() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _mockGetVerifiedWarpMessage(new bytes(0), true);
        _mockGetBlockchainID();
        vm.expectRevert(_formatErrorMessage("invalid source chain ID"));
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeSenderAddress() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _mockGetBlockchainID();
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

        vm.expectRevert(_formatErrorMessage("invalid origin sender address"));
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeValidationID() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _mockGetBlockchainID();
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

        vm.expectRevert(_formatErrorMessage("invalid uptime validation ID"));
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInitializeDelegatorRegistration() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
    }

    function testResendDelegatorRegistration() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        bytes32 delegationID = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(
            validationID, 1, DEFAULT_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
        );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendDelegatorRegistration(delegationID);
    }

    function testCompleteDelegatorRegistration() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        bytes32 delegationID = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1
        );
    }

    function testCompleteDelegatorRegistrationWrongNonce() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        // Initialize two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
                + DEFAULT_WEIGHT,
            expectedNonce: 2
        });

        // Complete registration of delegator2 with delegator1's nonce
        // Note that registering delegator1 with delegator2's nonce is valid
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSubnetValidatorWeightUpdateMessage(
            validationID, 1, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _mockGetVerifiedWarpMessage(setValidatorWeightPayload, true);

        vm.warp(DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP);
        vm.expectRevert("PoSValidatorManager: nonce does not match");
        posValidatorManager.completeDelegatorRegistration(0, delegationID2);
    }

    function testCompleteDelegatorRegistrationImplicitNonce() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        // Initialize two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        address delegator2 = address(0x5678567856785678567856785678567856785678);
        _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
                + DEFAULT_WEIGHT,
            expectedNonce: 2
        });
        // Mark delegator1 as registered by delivering the weight update from nonce 2 (delegator 2's nonce)
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID1,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            2
        );
    }

    function testInitializeEndDelegation() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        bytes32 delegationID = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1
        );
        _setUpInitializeEndDelegation({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });
    }

    function testResendEndDelegation() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        bytes32 delegationID = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1
        );
        _setUpInitializeEndDelegation({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 2, DEFAULT_WEIGHT);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendEndDelegation(delegationID);
    }

    function testResendEndValidation() public override {
        bytes32 validationID = _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            includeUptime: true
        });
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        validatorManager.resendEndValidatorMessage(validationID);
    }

    function testCompleteEndDelegation() public {
        uint256 registrationDuration = 1000 * 60 * 60 * 24; // 1 day

        uint256 registrationExpiry = vm.getBlockTimestamp() + registrationDuration;

        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: uint64(registrationExpiry),
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: uint64(registrationExpiry)
        });
        bytes32 delegationID = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        _setUpCompleteDelegatorRegistration({
            validationID: validationID,
            delegationID: delegationID,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        _setUpInitializeEndDelegation({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: uint64(registrationExpiry),
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });

        address delegator = DEFAULT_DELEGATOR_ADDRESS;
        uint256 balanceBefore = _getStakeAssetBalance(delegator);
        uint256 expectedReward = rewardCalculator.calculateReward(
            DEFAULT_DELEGATOR_WEIGHT,
            uint64(registrationExpiry) - DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            0,
            0
        );

        _setUpCompleteEndDelegation({
            validationID: validationID,
            delegationID: delegationID,
            delegator: delegator,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });

        uint256 balanceAfter = _getStakeAssetBalance(delegator);
        vm.assertEq(balanceAfter, balanceAfter);
        vm.assertEq(expectedReward, expectedReward);
        uint256 balanceChange = balanceAfter - balanceBefore;
        require(
            balanceChange == DEFAULT_DELEGATOR_WEIGHT + expectedReward,
            "delegator should have received their stake back and been rewarded"
        );
    }

    function testCompleteEndDelegationWrongNonce() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        // Register two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID1,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1
        );
        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
                + DEFAULT_WEIGHT,
            expectedNonce: 2
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID2,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            2
        );

        // Initialize end delegation for both delegators
        _setUpInitializeEndDelegation({
            validationID: validationID,
            delegatorAddress: delegator1,
            delegationID: delegationID1,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 3
        });
        _setUpInitializeEndDelegation({
            validationID: validationID,
            delegatorAddress: delegator2,
            delegationID: delegationID2,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4
        });

        // Complete ending delegator2 with delegator1's nonce
        // Note that ending delegator1 with delegator2's nonce is valid
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSubnetValidatorWeightUpdateMessage(
            validationID, 3, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _mockGetVerifiedWarpMessage(setValidatorWeightPayload, true);

        vm.expectRevert("PoSValidatorManager: nonce does not match");
        posValidatorManager.completeEndDelegation(0, delegationID2);
    }

    function testCompleteEndDelegationImplicitNonce() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        // Register two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID1,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1
        );
        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
                + DEFAULT_WEIGHT,
            expectedNonce: 2
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID2,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            2
        );

        // Initialize end delegation for both delegators
        _setUpInitializeEndDelegation({
            validationID: validationID,
            delegatorAddress: delegator1,
            delegationID: delegationID1,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 3
        });
        _setUpInitializeEndDelegation({
            validationID: validationID,
            delegatorAddress: delegator2,
            delegationID: delegationID2,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4
        });

        // Complete delegation1 by delivering the weight update from nonce 4 (delegator2's nonce)
        _setUpCompleteEndDelegation({
            validationID: validationID,
            delegationID: delegationID1,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedReward: rewardCalculator.calculateReward(
                DEFAULT_DELEGATOR_WEIGHT,
                DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP
                    - DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
                0,
                0
            ),
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4
        });
    }

    function testCompleteEndValidation() public override {
        uint256 registrationDuration = 1000 * 60 * 60 * 24; // 1 day

        uint256 registrationExpiry = vm.getBlockTimestamp() + registrationDuration;

        bytes32 validationID = _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: uint64(registrationExpiry),
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: uint64(registrationExpiry),
            includeUptime: true
        });

        uint256 balanceBefore = _getStakeAssetBalance(address(this));

        uint256 expectedReward = DEFAULT_WEIGHT * DEFAULT_REWARD_RATE / 365 - 1;

        _expectStakeUnlock(address(this), DEFAULT_WEIGHT);
        _expectRewardIssuance(address(this), expectedReward);

        _testCompleteEndValidation(validationID);

        uint256 balanceChange = _getStakeAssetBalance(address(this)) - balanceBefore;
        require(
            balanceChange == DEFAULT_WEIGHT + expectedReward,
            "validator should have received their stake back and been rewarded"
        );
    }

    function testValueToWeight() public view {
        uint64 w1 = posValidatorManager.valueToWeight(1e12);
        uint64 w2 = posValidatorManager.valueToWeight(1e18);
        uint64 w3 = posValidatorManager.valueToWeight(1e27);

        assertEq(w1, 1);
        assertEq(w2, 1e6);
        assertEq(w3, 1e15);
    }

    function testWeightToValue() public view {
        uint256 v1 = posValidatorManager.weightToValue(1);
        uint256 v2 = posValidatorManager.weightToValue(1e6);
        uint256 v3 = posValidatorManager.weightToValue(1e15);

        assertEq(v1, 1e12);
        assertEq(v2, 1e18);
        assertEq(v3, 1e27);
    }

    function _initializeEndValidation(bytes32 validationID) internal virtual override {
        return posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight
    ) internal virtual returns (bytes32);

    //
    // Delegation setup utilities
    //
    function _setUpInitializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight,
        uint64 registrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal returns (bytes32) {
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, expectedNonce, expectedValidatorWeight);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.warp(registrationTimestamp);

        _beforeSend(weight, delegatorAddress);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorAdded({
            delegationID: keccak256(abi.encodePacked(validationID, delegatorAddress, expectedNonce)),
            validationID: validationID,
            delegatorAddress: delegatorAddress,
            nonce: expectedNonce,
            validatorWeight: expectedValidatorWeight,
            delegatorWeight: weight,
            setWeightMessageID: bytes32(0)
        });

        return _initializeDelegatorRegistration(validationID, delegatorAddress, weight);
    }

    function _setUpCompleteDelegatorRegistration(
        bytes32 validationID,
        bytes32 delegationID,
        uint64 completeRegistrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal returns (bytes32) {
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSubnetValidatorWeightUpdateMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
        _mockGetVerifiedWarpMessage(setValidatorWeightPayload, true);

        vm.warp(completeRegistrationTimestamp);
        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRegistered({
            delegationID: delegationID,
            validationID: validationID,
            nonce: expectedNonce,
            startTime: completeRegistrationTimestamp
        });
        posValidatorManager.completeDelegatorRegistration(0, delegationID);
        return delegationID;
    }

    function _setUpInitializeEndDelegation(
        bytes32 validationID,
        address delegatorAddress,
        bytes32 delegationID,
        uint64 startDelegationTimestamp,
        uint64 endDelegationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal returns (bytes32) {
        vm.warp(endDelegationTimestamp);
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, expectedNonce, expectedValidatorWeight);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, endDelegationTimestamp - startDelegationTimestamp
        );
        _mockGetVerifiedWarpMessage(uptimeMsg, true);
        _mockGetBlockchainID(P_CHAIN_BLOCKCHAIN_ID);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRemovalInitialized({
            delegationID: delegationID,
            validationID: validationID,
            nonce: expectedNonce,
            validatorWeight: expectedValidatorWeight,
            endTime: endDelegationTimestamp,
            setWeightMessageID: bytes32(0)
        });
        vm.prank(delegatorAddress);
        posValidatorManager.initializeEndDelegation(delegationID, true, 0);
        return delegationID;
    }

    function _setUpCompleteEndDelegation(
        bytes32 validationID,
        bytes32 delegationID,
        address delegator,
        uint256 delegatorWeight,
        uint256 expectedReward,
        uint64 validatorWeight,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal returns (bytes32) {
        bytes memory weightUpdateMessage = ValidatorMessages.packSubnetValidatorWeightUpdateMessage(
            validationID, expectedNonce, validatorWeight
        );
        _mockGetVerifiedWarpMessage(weightUpdateMessage, true);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegationEnded(delegationID, validationID, expectedNonce);
        _expectStakeUnlock(delegator, delegatorWeight);
        _expectRewardIssuance(delegator, expectedReward);
        posValidatorManager.completeEndDelegation(0, delegationID);
        assertEq(posValidatorManager.getWeight(validationID), expectedValidatorWeight);
        return delegationID;
    }

    function _getStakeAssetBalance(address account) internal virtual returns (uint256);
    function _expectStakeUnlock(address account, uint256 amount) internal virtual;
    function _expectRewardIssuance(address account, uint256 amount) internal virtual;

    function _formatErrorMessage(bytes memory errorMessage) internal pure returns (bytes memory) {
        return abi.encodePacked("PoSValidatorManager: ", errorMessage);
    }
}
