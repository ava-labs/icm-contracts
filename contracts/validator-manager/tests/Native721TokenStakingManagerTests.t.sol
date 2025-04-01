// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {StakingManagerTest} from "./StakingManagerTests.t.sol";
import {Native721TokenStakingManager} from "../Native721TokenStakingManager.sol";
import {StakingManager, StakingManagerSettings} from "../StakingManager.sol";
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";
import {ACP99Manager, PChainOwner, ConversionData} from "../ACP99Manager.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {ExampleERC721} from "@mocks/ExampleERC721.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {IERC721Receiver} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721Receiver.sol";
import {console} from "forge-std/console.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
contract Native721TokenStakingManagerTest is StakingManagerTest, IERC721Receiver {
    Native721TokenStakingManager public app;

    ExampleERC721 public stakingToken;
    IERC20 public rewardToken;

    uint128 public constant REWARD_PER_EPOCH = 100e18;
    uint128 public constant REWARD_CLAIM_DELAY = 7 days;

    uint256 testTokenID = 0;

    function setUp() public override {
        ValidatorManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();

        ConversionData memory conversion = _defaultConversionData();
        bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
        _mockInitializeValidatorSet(conversionID);
        validatorManager.initializeValidatorSet(conversion, 0);
    }

    function onERC721Received(
        address,
        address,
        uint256,
        bytes memory
    ) public virtual returns (bytes4) {
        return IERC721Receiver.onERC721Received.selector;
    }

    //
    // Initialization unit tests
    // The pattern in these tests requires that only non-admin validator manager functions are called,
    // as each test re-deploys the Native721TokenStakingManager contract.
    //
    function testDisableInitialization() public {
        app = new Native721TokenStakingManager(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        app.initialize(defaultPoSSettings, stakingToken);
    }

    function testInvalidTokenAddress() public {
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(Native721TokenStakingManager.InvalidTokenAddress.selector, address(0)));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        app.initialize(defaultPoSSettings, IERC721(address(0)));
    }

    function testZeroMinimumDelegationFee() public {
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(StakingManager.InvalidDelegationFee.selector, 0));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.minimumDelegationFeeBips = 0;
        app.initialize(defaultPoSSettings, stakingToken);
    }

    function testMaxMinimumDelegationFee() public {
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        uint16 minimumDelegationFeeBips = app.MAXIMUM_DELEGATION_FEE_BIPS() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidDelegationFee.selector, minimumDelegationFeeBips
            )
        );

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.minimumDelegationFeeBips = minimumDelegationFeeBips;
        app.initialize(defaultPoSSettings, stakingToken);
    }

    function testInvalidStakeAmountRange() public {
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidStakeAmount.selector, DEFAULT_MAXIMUM_STAKE_AMOUNT
            )
        );

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.minimumStakeAmount = DEFAULT_MAXIMUM_STAKE_AMOUNT;
        defaultPoSSettings.maximumStakeAmount = DEFAULT_MINIMUM_STAKE_AMOUNT;
        app.initialize(defaultPoSSettings, stakingToken);
    }

    function testZeroWeightToValueFactor() public {
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(abi.encodeWithSelector(StakingManager.ZeroWeightToValueFactor.selector));

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.weightToValueFactor = 0;
        app.initialize(defaultPoSSettings, stakingToken);
    }

    function testMinStakeDurationTooLow() public {
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        uint64 minStakeDuration = DEFAULT_CHURN_PERIOD - 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidMinStakeDuration.selector, minStakeDuration
            )
        );

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;
        defaultPoSSettings.minimumStakeDuration = minStakeDuration;
        app.initialize(defaultPoSSettings, stakingToken);
    }

    function testInvalidValidatorManager() public {
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        Native721TokenStakingManager invalidManager =
            new Native721TokenStakingManager(ICMInitializable.Allowed); // the contract type is arbitrary

        vm.expectRevert();

        StakingManagerSettings memory settings = _defaultPoSSettings();
        settings.manager = ValidatorManager(address(invalidManager));
        app.initialize(settings, stakingToken);
    }

    function testUnsetValidatorManager() public {
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert();

        app.initialize(_defaultPoSSettings(), stakingToken); // settings.manager is not set
    }

    function testNFTDelegationOverWeightLimit() public {
        bytes32 validationID = _registerDefaultValidator();

        uint256 tokenID = 6;
        uint256[] memory tokens = new uint256[](DEFAULT_MAXIMUM_NFT_AMOUNT);
        for(uint256 i = 0; i<DEFAULT_MAXIMUM_NFT_AMOUNT; i++){
            stakingToken.mint(DEFAULT_DELEGATOR_ADDRESS, ++tokenID);
            tokens[i] = tokenID;
        }

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        stakingToken.setApprovalForAll(address(app), true);

        vm.expectRevert(
            abi.encodeWithSelector(
                Native721TokenStakingManager.InvalidNFTAmount.selector, DEFAULT_MAXIMUM_NFT_AMOUNT + 1
            )
        );

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        app.registerNFTDelegation(validationID, tokens);
    }

    function testSubmitUptimeNonOwner() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);
        bytes memory uptimeMessage =
            ValidatorMessages.packValidationUptimeMessage(validationID, 0);

        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, DEFAULT_DELEGATOR_ADDRESS
            )
        );

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        app.submitUptimeProof(validationID, 0);
    }

     function testSubmitUptimes() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes32 nextValidationID = _registerValidator({
            nodeID: _newNodeID(),
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);
        
        bytes memory uptimeMessage0 =
            ValidatorMessages.packValidationUptimeMessage(validationID, DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    originSenderAddress: address(0),
                    payload: uptimeMessage0
                }),
                true
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );

        bytes memory uptimeMessage1 =
            ValidatorMessages.packValidationUptimeMessage(nextValidationID, DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(1)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    originSenderAddress: address(0),
                    payload: uptimeMessage1
                }),
                true
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 1)
        );

        bytes32[] memory validationIDs = new bytes32[](2);
        validationIDs[0] = validationID; 
        validationIDs[1] = nextValidationID;

        uint32[] memory messageIndexes = new uint32[](2);
        messageIndexes[0] = 0; 
        messageIndexes[1] = 1;

        vm.prank(DEFAULT_UPTIME_KEEPER);
        app.submitUptimeProofs(validationIDs, messageIndexes);
    }

    function testSubmitUptimesInvalidInput() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32[] memory validationIDs = new bytes32[](1);
        validationIDs[0] = validationID; 

        uint32[] memory messageIndexes = new uint32[](2);
        messageIndexes[0] = 0; 
        messageIndexes[1] = 1; 

        vm.expectRevert(
            abi.encodeWithSelector(
                Native721TokenStakingManager.InvalidInputLengths.selector, 1, 2
            )
        ); 
        vm.prank(DEFAULT_UPTIME_KEEPER);
        app.submitUptimeProofs(validationIDs, messageIndexes);
    }

    function testRewardRegistrationNonOwner() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, DEFAULT_DELEGATOR_ADDRESS
            )
        );

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        app.registerRewards(true, 0, address(rewardToken), REWARD_PER_EPOCH);
    }

    function testRewardCancellationTooLate() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                Native721TokenStakingManager.TooLate.selector, 2 * DEFAULT_EPOCH_DURATION, 604800
            )
        );

        vm.warp(2 * DEFAULT_EPOCH_DURATION);
        app.cancelRewards(true, 0, address(rewardToken));    
    }

     function testRewardCancellationNonOwner() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, DEFAULT_DELEGATOR_ADDRESS
            )
        );

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        app.cancelRewards(true, 0, address(rewardToken));
    }

    function testCancelRewards() public {
        app.cancelRewards(true, 0, address(rewardToken));    
        app.cancelRewards(false, 0, address(rewardToken));    
    }

    function testDelegatorNFTRemoval() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerNFTDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);

        vm.warp(DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP);

        _initiateNFTDelegatorRemoval({
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID
        });

        vm.warp(block.timestamp + DEFAULT_UNLOCK_DURATION);
        _completeNFTDelegatorRemoval(DEFAULT_DELEGATOR_ADDRESS, delegationID);

        _expectNFTStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, 1);
    }

    function testDelegationRewards() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: address(this)
        });

        // Validator is Completed, so this will also complete the delegation.
        _initiateDelegatorRemoval({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            includeUptime: false,
            force: false,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);
        _submitUptime(validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP);

        bytes32[] memory delegationIDs = new bytes32[](1);
        delegationIDs[0] = delegationID;
        _resolveRewards(delegationIDs);

        (uint256 validatorReward, uint256 delegatorReward) = _calculateExpectedRewards(
            DEFAULT_WEIGHT, DEFAULT_DELEGATOR_WEIGHT, DEFAULT_DELEGATION_FEE_BIPS);

        _claimReward(true, address(this), validatorReward);
        _claimReward(true, DEFAULT_DELEGATOR_ADDRESS, delegatorReward);
    }

     function testRewardsTooEarly() public {
        bytes32 validationID = _registerDefaultValidator();

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: address(this)
        });

        vm.warp(DEFAULT_EPOCH_DURATION);
        _submitUptime(validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP);


        address[] memory tokens = new address[](1);
        tokens[0] = address(rewardToken);
        
        vm.expectRevert(
            abi.encodeWithSelector(
                Native721TokenStakingManager.TooEarly.selector, block.timestamp, DEFAULT_EPOCH_DURATION + REWARD_CLAIM_DELAY
            )
        );
        app.claimRewards(true, 0, tokens, address(this));
    }    

    function testDelegationRewardsForSameValidatorAndDelegator() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes32 delegationID = _registerDelegator({
            validationID: validationID,
            delegatorAddress: address(this),
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: address(this)
        });

        // Validator is Completed, so this will also complete the delegation.
        _initiateDelegatorRemoval({
            sender: address(this),
            delegationID: delegationID,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            includeUptime: false,
            force: false,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);
        _submitUptime(validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP);

        bytes32[] memory delegationIDs = new bytes32[](1);
        delegationIDs[0] = delegationID;
        _resolveRewards(delegationIDs);

        (uint256 validatorReward, uint256 delegatorReward) = _calculateExpectedRewards(
            DEFAULT_WEIGHT, DEFAULT_DELEGATOR_WEIGHT, DEFAULT_DELEGATION_FEE_BIPS);

        _claimReward(true, address(this), validatorReward + delegatorReward);
    }

    function testNFTDelegationRewards() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerNFTDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: address(this)
        });

        // Validator is Completed, so this will also complete the delegation.
        _initiateNFTDelegatorRemoval({
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID
        });

        _expectNFTStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, 1);


        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);
        _submitUptime(validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP);

        bytes32[] memory delegationIDs = new bytes32[](1);
        delegationIDs[0] = delegationID;
        _resolveRewards(delegationIDs);

        (uint256 validatorReward, uint256 delegatorReward) = _calculateExpectedRewards(
            1e6, 1e6, DEFAULT_DELEGATION_FEE_BIPS);

        _claimReward(false, address(this), validatorReward );
        _claimReward(false, DEFAULT_DELEGATOR_ADDRESS, delegatorReward);
    }

    function testDoubleDelegationRewards() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        bytes32 newDelegationID = _registerDelegator({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: 2 * DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 2
        });

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 3,
            rewardRecipient: address(this)
        });

        // Validator is Completed, so this will also complete the delegation.
        _initiateDelegatorRemoval({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            includeUptime: false,
            force: false,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        _initiateDelegatorRemoval({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: newDelegationID,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            includeUptime: false,
            force: false,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);
        _submitUptime(validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP);

        bytes32[] memory delegationIDs = new bytes32[](2);
        delegationIDs[0] = delegationID;
        delegationIDs[1] = newDelegationID;
        _resolveRewards(delegationIDs);

        (uint256 validatorReward, uint256 delegatorReward) = _calculateExpectedRewards(
            DEFAULT_WEIGHT, DEFAULT_DELEGATOR_WEIGHT * 2, DEFAULT_DELEGATION_FEE_BIPS);

        _claimReward(true, address(this), validatorReward);
        _claimReward(true, DEFAULT_DELEGATOR_ADDRESS, delegatorReward);
    }

    function testDefaultAndNFTDelegationRewards() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);
        bytes32 nftDelegationID = _registerNFTDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: address(this)
        });

        // Validator is Completed, so this will also complete the delegation.
        _initiateNFTDelegatorRemoval({
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: nftDelegationID
        });

        _expectNFTStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, 1);

        // Validator is Completed, so this will also complete the delegation.
        _initiateDelegatorRemoval({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION + 1,
            includeUptime: true,
            force: true,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });

        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION);
        _submitUptime(validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP);

        bytes32[] memory delegationIDs = new bytes32[](2);
        delegationIDs[0] = delegationID;
        delegationIDs[1] = nftDelegationID;
        _resolveRewards(delegationIDs);

        (uint256 validatorReward, uint256 delegatorReward) = _calculateExpectedRewards(
            DEFAULT_WEIGHT, DEFAULT_DELEGATOR_WEIGHT, DEFAULT_DELEGATION_FEE_BIPS);

        _claimReward(true, address(this), validatorReward);
        _claimReward(true, DEFAULT_DELEGATOR_ADDRESS, delegatorReward);

        (validatorReward, delegatorReward) = _calculateExpectedRewards(
            1e6, 1e6, DEFAULT_DELEGATION_FEE_BIPS);
        
        _claimReward(false, address(this), validatorReward);
        _claimReward(false, DEFAULT_DELEGATOR_ADDRESS, delegatorReward);
    }

    function testNFTRedelegation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerNFTDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);

        address rewardRecipient = address(42);

        bytes32 nextValidationID = _registerValidator({
            nodeID: _newNodeID(),
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        vm.warp(block.timestamp + DEFAULT_MINIMUM_STAKE_DURATION + 1);

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        app.registerNFTRedelegation(delegationID, nextValidationID);
    }

    function testEndDelegationNFTBeforeUnlock() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerNFTDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);

        vm.warp(block.timestamp + DEFAULT_MINIMUM_STAKE_DURATION + 1);

        _initiateNFTDelegatorRemoval({
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID
        });

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.UnlockDurationNotPassed.selector, block.timestamp
            )
        );
        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        app.completeNFTDelegatorRemoval(delegationID); 
    }

    function testRevertDoubleCompletion() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerNFTDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: address(this)
        });

        vm.warp(block.timestamp + DEFAULT_MINIMUM_STAKE_DURATION + 1);

        // completes the delegation as validation already ended
        _initiateNFTDelegatorRemoval({
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID
        });

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.InvalidDelegatorStatus.selector, 4
            )
        );
        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        app.completeNFTDelegatorRemoval(delegationID); 
    }

    function testEndNFTDelegationRevertBeforeMinStakeDuration() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerNFTDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.MinStakeDurationNotPassed.selector, block.timestamp
            )
        );
        _initiateNFTDelegatorRemoval({
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID
        });
    }

    function testValidationRegistrationWithoutNFT() public {
         vm.expectRevert(
            abi.encodeWithSelector(
                Native721TokenStakingManager.InvalidNFTAmount.selector, 0
            )
        );
        uint256[] memory tokens = new uint256[](0);
        app.initiateValidatorRegistration{value: DEFAULT_MINIMUM_STAKE_AMOUNT}({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationExpiry: DEFAULT_EXPIRY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            tokenIDs: tokens
        });
    }

    function testRevertRemovalDelgationNFTForNonOwner() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerNFTDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EPOCH_DURATION,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 1,
            rewardRecipient: address(this)
        });

        vm.expectRevert(
            abi.encodeWithSelector(
                StakingManager.UnauthorizedOwner.selector, address(42)
            )
        );
        _initiateNFTDelegatorRemoval({
            delegatorAddress: address(42),
            delegationID: delegationID
        });
    }
 
    // Helpers
    function _calculateExpectedRewards(
        uint256 validatorStake,
        uint256 delegatorStake,
        uint256 delegationFeeBips
    ) internal pure returns (uint256 validatorReward, uint256 delegatorReward) {
        uint256 feeWeight = delegatorStake * delegationFeeBips / 10000;
        delegatorReward = (REWARD_PER_EPOCH * (delegatorStake - feeWeight)) / (delegatorStake + validatorStake);
        validatorReward = (REWARD_PER_EPOCH * (validatorStake + feeWeight)) / (delegatorStake + validatorStake);
    }

    function _registerNFTDelegation(
        bytes32 validationID,
        address delegatorAddress
    ) internal virtual returns (bytes32) {
        uint256[] memory tokens = new uint256[](1);
        tokens[0] = ++testTokenID;

        _beforeSendNFT(tokens[0], delegatorAddress);

        vm.prank(delegatorAddress);
        return app.registerNFTDelegation(validationID, tokens);
    }

    function _initiateNFTDelegatorRemoval(
        address delegatorAddress,
        bytes32 delegationID
    ) internal virtual returns (bytes32) {
        vm.prank(delegatorAddress);
        app.initiateNFTDelegatorRemoval(delegationID);
    }

    function _completeNFTDelegatorRemoval(
        address delegatorAddress,
        bytes32 delegationID
    ) internal virtual returns (bytes32) {
        vm.warp(block.timestamp + DEFAULT_UNLOCK_DURATION);
        vm.prank(delegatorAddress);
        app.completeNFTDelegatorRemoval(delegationID);
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        uint64 registrationExpiry,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) internal virtual override returns (bytes32) {
        uint256[] memory tokens = new uint256[](1);
        tokens[0] = ++testTokenID;
        return app.initiateValidatorRegistration{value: stakeAmount}({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            registrationExpiry: registrationExpiry,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration,
            tokenIDs: tokens
        });
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        uint64 registrationExpiry,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        uint256[] memory tokens = new uint256[](1);
        tokens[0] = ++testTokenID;
        return app.initiateValidatorRegistration{value: _weightToValue(weight)}({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            registrationExpiry: registrationExpiry,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            tokenIDs: tokens
        });
    }

    function _initiateDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        uint256 value = _weightToValue(weight);
        vm.prank(delegatorAddress);
        vm.deal(delegatorAddress, value);
        return app.initiateDelegatorRegistration{value: value}(validationID);
    }

    // solhint-disable no-empty-blocks
    function _beforeSend(uint256 amount, address spender) internal override {
        // Native tokens no need pre approve
    }

    function _beforeSendNFT(uint256 tokenId, address spender) internal {
        stakingToken.transferFrom(address(this), spender, tokenId);

        vm.prank(spender);
        stakingToken.approve(address(app), tokenId);
    }
    // solhint-enable no-empty-blocks

    function _expectStakeUnlock(address account, uint256 amount) internal override {
        // empty calldata implies the receive function will be called
        vm.expectCall(account, amount, "");
    }

    function _expectNFTStakeUnlock(address account, uint256 amount) internal view {
        assertEq(stakingToken.balanceOf(account), amount);
    }

    function _expectRewardIssuance(address account, uint256 amount) internal override {
    }

    function _claimReward(bool primary, address account, uint256 expectedAmount) internal returns (uint256) {
        uint256 balanceBefore = rewardToken.balanceOf(account);
        
        address[] memory tokens = new address[](1);
        tokens[0] = address(rewardToken);
        
        vm.prank(account);
        vm.warp(block.timestamp + REWARD_CLAIM_DELAY);
        app.claimRewards(primary, 0, tokens, account);
        
        assertApproxEqRel(expectedAmount, rewardToken.balanceOf(account) - balanceBefore, 0.1e18);
    }

    function _submitUptime(bytes32 validationID, uint64 uptime) internal {
        bytes memory uptimeMessage =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime);
        _mockGetUptimeWarpMessage(uptimeMessage, true);

        vm.prank(DEFAULT_UPTIME_KEEPER);
        app.submitUptimeProof(validationID, 0);
    }

    function _resolveRewards(bytes32[] memory delegationIDs) internal {
        vm.prank(DEFAULT_UPTIME_KEEPER);
        app.resolveRewards(delegationIDs);
    }

    function _setUp() internal override returns (ACP99Manager) {
        // Construct the object under test
        app = new Native721TokenStakingManager(ICMInitializable.Allowed);
        validatorManager = new ValidatorManager(ICMInitializable.Allowed);

        rewardToken = new ExampleERC20();
        stakingToken = new ExampleERC721();
        rewardCalculator = new ExampleRewardCalculator(DEFAULT_REWARD_RATE);

        stakingToken.setApprovalForAll(address(app), true);

        StakingManagerSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.manager = validatorManager;

        validatorManager.initialize(_defaultSettings(address(app)));
        app.initialize(defaultPoSSettings, stakingToken);

        rewardToken.approve(address(app), REWARD_PER_EPOCH * 2);

        app.registerRewards(true, 0, address(rewardToken), REWARD_PER_EPOCH);
        app.registerRewards(false, 0, address(rewardToken), REWARD_PER_EPOCH);

        stakingManager = app;

        return validatorManager;
    }

    function _getStakeAssetBalance(address account) internal view override returns (uint256) {
        return account.balance;
    }
}