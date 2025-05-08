// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    StakingManagerSettings
} from "./interfaces/IStakingManager.sol";
import {
    StakingManager
} from "./StakingManager.sol";
import {
    Delegator,
    DelegatorStatus,
    IStakingManager,
    PoSValidatorInfo,
    StakingManagerSettings
} from "./interfaces/IStakingManager.sol";
import { Math } from "@openzeppelin/contracts@5.0.2/utils/math/Math.sol";
import {INative721TokenStakingManager} from "./interfaces/INative721TokenStakingManager.sol";
import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {IERC721Receiver} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721Receiver.sol";

import {Validator, ValidatorStatus, PChainOwner} from "./ACP99Manager.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
/**
 * @dev Implementation of the {INative721TokenStakingManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract Native721TokenStakingManager is
    Initializable,
    StakingManager,
    OwnableUpgradeable,
    INative721TokenStakingManager,
    IERC721Receiver
{
    using Address for address payable;

    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.Native721TokenStakingManager
    struct Native721TokenStakingManagerStorage {
        IERC721 _token;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.Native721TokenStakingManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant ERC721_STAKING_MANAGER_STORAGE_LOCATION =
        0xf2d79c30881febd0da8597832b5b1bf1f4d4b2209b19059420303eb8fcab8a00;
    
    uint8 public constant UPTIME_REWARDS_THRESHOLD_PERCENTAGE = 80;
    uint64 public constant REWARD_CLAIM_DELAY = 7 days;

    error InvalidNFTAmount(uint256 nftAmount);
    error InvalidTokenAddress(address tokenAddress);
    error InvalidInputLengths(uint256 inputLength1, uint256 inputLength2);
    error TooEarly(uint256 actualTime, uint256 expectedTime);
    error TooLate(uint256 actualTime, uint256 expectedTime);


    // solhint-disable ordering
    function _getERC721StakingManagerStorage()
        private
        pure
        returns (Native721TokenStakingManagerStorage storage $)
    {
        assembly {
            $.slot := ERC721_STAKING_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initialize the ERC721 token staking manager
     * @dev Uses reinitializer(2) on the PoS staking contracts to make sure after migration from PoA, the PoS contracts can reinitialize with its needed values.
     * @param settings Initial settings for the PoS validator manager
     * @param stakingToken The ERC721 token to be staked
     */
    function initialize(
        StakingManagerSettings calldata settings,
        IERC721 stakingToken
    ) external reinitializer(3) {
        __Ownable_init(settings.admin);
        __StakingManager_init(settings);

        Native721TokenStakingManagerStorage storage $ = _getERC721StakingManagerStorage();

        if (address(stakingToken) == address(0)) {
            revert InvalidTokenAddress(address(stakingToken));
        }

        $._token = stakingToken;
    }

    function onERC721Received(
        address,
        address,
        uint256,
        bytes memory
    ) public virtual returns (bytes4) {
        return this.onERC721Received.selector;
    }

    /**
     * @notice See {INative721TokenStakingManager-initiateValidatorRegistration}.
     */
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        uint64 registrationExpiry,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256[] memory tokenIDs
    ) external payable nonReentrant returns (bytes32) {
        return _initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            registrationExpiry: registrationExpiry,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration,
            stakeAmount: msg.value,
            tokenIDs: tokenIDs
        });
    }

    /**
     * @notice See {INative721TokenStakingManager-initiateDelegatorRegistration}.
     */
    function initiateDelegatorRegistration(bytes32 validationID)
        external
        payable
        nonReentrant
        returns (bytes32)
    {
        return _initiateDelegatorRegistration(validationID, _msgSender(), msg.value);
    }

    /**
    * @notice See {INative721TokenStakingManager-registerNFTDelegation}.
    *
    */
    function registerNFTDelegation(
        bytes32 validationID,
        uint256[] memory tokenIDs
    ) external nonReentrant returns (bytes32) {
        _lockNFTs(tokenIDs);
        return _registerNFTDelegation(validationID, _msgSender(), tokenIDs);
    }

    /**
    * @notice See {INative721TokenStakingManager-initializeEndNFTDelegation}.
    *
    */
    function initiateNFTDelegatorRemoval(
        bytes32 delegationID
    ) external nonReentrant {
        _initiateNFTDelegatorRemoval(delegationID);
    }

    /**
    * @notice See {INative721TokenStakingManager-completeEndNFTDelegation}.
    *
    */
    function completeNFTDelegatorRemoval(
        bytes32 delegationID
    ) external nonReentrant {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID];

        // Ensure the delegator is pending removed. Since anybody can call this function once
        // end delegation has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        if (delegator.status != DelegatorStatus.PendingRemoved) {
            revert InvalidDelegatorStatus(delegator.status);
        } 

        if(block.timestamp < delegator.endTime + $._unlockDuration) {
            revert UnlockDurationNotPassed(uint64(block.timestamp));
        }

        _unlockNFTs(delegator.owner, _completeNFTDelegatorRemoval(delegationID));
    }

    /**
    * @notice See {INative721TokenStakingManager-registerNFTRedelegation}.
    */
    function registerNFTRedelegation(
        bytes32 delegationID,
        bytes32 nextValidationID
    ) external nonReentrant {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID];

        _initiateNFTDelegatorRemoval(delegationID);
        uint256[] memory tokenIDs = _completeNFTDelegatorRemoval(delegationID);
        _registerNFTDelegation(nextValidationID, delegator.owner, tokenIDs);
    }

    /**
     * @notice unlocks the validator stake, to be called after removal and passing of unlock duration
     * @param validationID The unique identifier of the validator to unlock.
     */ 
    function unlockValidator(bytes32 validationID) external override nonReentrant {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        _unlockValidator(validationID);
        _unlockNFTs($._posValidatorInfo[validationID].owner, $._posValidatorInfo[validationID].tokenIDs);
    }

    /**
     * @notice See {INative721TokenStakingManager-submitUptimeProofs}.
     */
    function submitUptimeProofs(bytes32[] memory validationIDs, uint32[] memory messageIndexes) external {
        if(validationIDs.length != messageIndexes.length){
            revert InvalidInputLengths(validationIDs.length, messageIndexes.length);
        }
        for (uint256 i = 0; i < validationIDs.length; i++) {
            _updateUptime(validationIDs[i], messageIndexes[i]);
        }
    }

    /**
     * @notice Calculates the rewards for the caller in a given epoch for the specified tokens.
     * @dev This function determines the available rewards based on the user's weight in the staking system.
     *      It supports both primary and NFT-based reward pools.
     * @param primary A boolean indicating whether to retrieve rewards from the primary pool (true) or the NFT pool (false).
     * @param epoch The staking epoch for which to retrieve rewards.
     * @param tokens An array of token addresses for which to check the rewards.
     * @param account The account for which the rewards are being queried.
     * @return rewards An array of reward amounts corresponding to the provided token addresses.
     *
     * Requirements:
     * - The caller must have participated in staking or NFT delegation during the given epoch.
     * - The function calculates rewards based on the caller’s recorded weight and subtracts any already withdrawn rewards.
    */
    function getRewards(
        bool primary,
        uint64 epoch,
        address[] memory tokens,
        address account
    ) public view returns (uint256[] memory) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        uint256[] memory rewards = new uint256[](tokens.length);

        if(primary && $._totalRewardWeight[epoch] == 0){ return rewards; }
        if(!primary && $._totalRewardWeightNFT[epoch] == 0){ return rewards; }

        for(uint256 i = 0; i < tokens.length; i++){
            if(primary){
                rewards[i] = (($._rewardPools[epoch][tokens[i]] * $._accountRewardWeight[epoch][account])
                    / $._totalRewardWeight[epoch]) - $._rewardWithdrawn[epoch][account][tokens[i]];
            } else {
                rewards[i] = (($._rewardPoolsNFT[epoch][tokens[i]] * $._accountRewardWeightNFT[epoch][account])
                    / $._totalRewardWeightNFT[epoch]) - $._rewardWithdrawnNFT[epoch][account][tokens[i]];
            }
        }
        return rewards;
    }

    /**
     * @notice See {INative721TokenStakingManager-claimRewards}.
     */
    function claimRewards(
        bool primary,
        uint64 epoch,
        address[] memory tokens,
        address recipient
    ) external nonReentrant {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        uint64 claimStart = (epoch + 1) * $._epochDuration + REWARD_CLAIM_DELAY - $._epochOffset;
        if(block.timestamp < claimStart){
            revert TooEarly(block.timestamp, claimStart);
        }
        
        address sender = _msgSender();
        uint256[] memory rewards = getRewards(primary, epoch, tokens, sender);
        for(uint256 i = 0; i < tokens.length; i++){
            if(primary){
                $._rewardWithdrawn[epoch][sender][tokens[i]] += rewards[i];
            } else {
                $._rewardWithdrawnNFT[epoch][sender][tokens[i]] += rewards[i];
            }
            emit RewardClaimed(primary, epoch, sender, tokens[i], rewards[i]);
            IERC20(tokens[i]).transfer(recipient, rewards[i]);
        }
    }

    /**
     * @notice See {INative721TokenStakingManager-registerRewards}.
     */
    function registerRewards(
        bool primary,
        uint64 epoch,
        address token,
        uint256 amount
    ) external onlyOwner nonReentrant {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        if(primary){
            $._rewardPools[epoch][token] += amount;
        } else {
            $._rewardPoolsNFT[epoch][token] += amount;
        }
        IERC20(token).transferFrom(_msgSender(), address(this), amount);
        emit RewardRegistered(primary, epoch, token, amount);
    }

    /**
     * @notice See {INative721TokenStakingManager-cancelRewards}.
     */
    function cancelRewards(
        bool primary,
        uint64 epoch,
        address token
    ) external onlyOwner nonReentrant {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        uint64 claimStart = (epoch + 1) * $._epochDuration + REWARD_CLAIM_DELAY;
        if(block.timestamp >= claimStart){
            revert TooLate(block.timestamp, claimStart);
        }

        if(primary){
            IERC20(token).transfer(_msgSender(), $._rewardPools[epoch][token]);
            $._rewardPools[epoch][token] = 0;
        } else {
            IERC20(token).transfer(_msgSender(), $._rewardPoolsNFT[epoch][token]);
            $._rewardPoolsNFT[epoch][token] = 0;
        }
        emit RewardCancelled(primary, epoch, token);
    }

    /**
     * @notice See {INative721TokenStakingManager-erc721}.
     */
    function erc721() external view returns (IERC721) {
        return _getERC721StakingManagerStorage()._token;
    }

    /**
     * @notice See {StakingManager-_lock}
     */
    function _lock(uint256 value) internal virtual override returns (uint256) {
        return value;
    }

    /**
     * @notice See {StakingManager-_unlock}
     * Note: Must be guarded with reentrancy guard for safe transfer.
     */
    function _unlock(address to, uint256 value) internal virtual override {
        payable(to).sendValue(value);
    }

    /**
     * @notice Locks a list of ERC-721 tokens by transferring them to the contract.
     * @dev Transfers each token in `tokenIDs` from the caller to the contract.
     * This function is used to stake NFTs as part of the staking mechanism.
     * @param tokenIDs The array of token IDs to be locked.
     * @return The number of tokens successfully locked.
     */
    function _lockNFTs(uint256[] memory tokenIDs) internal returns (uint256) {
        for (uint256 i = 0; i < tokenIDs.length; i++) {
            _getERC721StakingManagerStorage()._token.safeTransferFrom(_msgSender(), address(this), tokenIDs[i]);
        }
        return tokenIDs.length;
    }

    /**
     * @notice Unlocks a list of ERC-721 tokens by transferring them back to the specified address.
     * @dev Transfers each token in `tokenIDs` from the contract to the recipient.
     * This function is used when unstaking NFTs from the staking mechanism.
     * @param to The address that will receive the unlocked NFTs.
     * @param tokenIDs The array of token IDs to be unlocked and transferred.
     */
    function _unlockNFTs(address to, uint256[] memory tokenIDs) internal virtual {
        for (uint256 i = 0; i < tokenIDs.length; i++) {
            _getERC721StakingManagerStorage()._token.transferFrom(address(this), to, tokenIDs[i]);
        }
    }

    /**
     * @notice See {StakingManager-_reward}
     * @dev Distributes ERC20 rewards to stakers
     */
    function _reward(address account, uint256 amount) internal virtual override {
    }

    function _getEpoch() internal view virtual returns (uint64) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        return uint64((block.timestamp + $._epochOffset) / $._epochDuration);
    }

    /**
     * @notice Initiates validator registration. Extends the functionality of {ACP99Manager-_initiateValidatorRegistration}
     * by locking stake and setting staking and delegation parameters.
     * @param delegationFeeBips The delegation fee in basis points.
     * @param minStakeDuration The minimum stake duration in seconds.
     * @param stakeAmount The amount of stake to lock.
     */
    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        uint64 registrationExpiry,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount,
        uint256[] memory tokenIDs
    ) internal virtual returns (bytes32) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        // Validate and save the validator requirements
        if (
            delegationFeeBips < $._minimumDelegationFeeBips
                || delegationFeeBips > MAXIMUM_DELEGATION_FEE_BIPS
        ) {
            revert InvalidDelegationFee(delegationFeeBips);
        }

        if (minStakeDuration < $._minimumStakeDuration) {
            revert InvalidMinStakeDuration(minStakeDuration);
        }

        // Ensure the weight is within the valid range.
        if (stakeAmount < $._minimumStakeAmount || stakeAmount > $._maximumStakeAmount) {
            revert InvalidStakeAmount(stakeAmount);
        }

        if (tokenIDs.length < 1 || tokenIDs.length > $._maximumNFTAmount) {
            revert InvalidNFTAmount(tokenIDs.length);
        }

        // Lock the stake in the contract.
        uint64 weight = valueToWeight(_lock(stakeAmount));
        _lockNFTs(tokenIDs);

        bytes32 validationID = $._manager.initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            registrationExpiry: registrationExpiry,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            weight: weight
        });

        address owner = _msgSender();

        $._posValidatorInfo[validationID].owner = owner;
        $._posValidatorInfo[validationID].delegationFeeBips = delegationFeeBips;
        $._posValidatorInfo[validationID].minStakeDuration = minStakeDuration;
        $._posValidatorInfo[validationID].uptimeSeconds = 0;
        $._posValidatorInfo[validationID].tokenIDs = tokenIDs;
        $._posValidatorInfo[validationID].totalTokens = tokenIDs.length;

        return validationID;
    }

    /**
     * @notice Registers an NFT-based delegation to a PoS validator.
     * @dev This function records the delegation details, ensures the validator is active, 
     *      updates the staking records, and emits relevant events.
     * @param validationID The identifier of the PoS validator to delegate to.
     * @param delegatorAddress The address of the user delegating NFTs.
     * @param tokenIDs The array of NFT token IDs being delegated.
     * @return delegationID A unique identifier for this delegation.
     *
     * Requirements:
     * - The validator must be a valid PoS validator.
     * - The validator must be in an active status.
     * - The function generates a unique delegation ID based on the validation ID and nonce.
     *
     * Emits:
     * - `InitiatedDelegatorRegistration` upon starting the delegation process.
     * - `CompletedDelegatorRegistration` once the delegation is successfully recorded.
     * - `DelegatedNFTs` containing the delegated NFT token IDs.
     */
    function _registerNFTDelegation(
        bytes32 validationID,
        address delegatorAddress,
        uint256[] memory tokenIDs
    ) internal returns (bytes32) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        uint64 weight = uint64(tokenIDs.length * 1e6);

        // Ensure the validation period is active
        Validator memory validator = $._manager.getValidator(validationID);
        // Check that the validation ID is a PoS validator
        if (!_isPoSValidator(validationID)) {
            revert ValidatorNotPoS(validationID);
        }
        if (validator.status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus(validator.status);
        }

        if ($._posValidatorInfo[validationID].totalTokens + tokenIDs.length > $._maximumNFTAmount) {
            revert InvalidNFTAmount(uint64($._posValidatorInfo[validationID].totalTokens + tokenIDs.length));
        }

        uint64 nonce = ++$._posValidatorInfo[validationID].tokenNonce;
        
        // Update the delegation status
        bytes32 delegationID = keccak256(abi.encodePacked(validationID, nonce, "ERC721"));
        $._delegatorStakes[delegationID].owner = delegatorAddress;
        $._delegatorStakes[delegationID].validationID = validationID;
        $._delegatorStakes[delegationID].weight = weight;
        $._delegatorStakes[delegationID].status = DelegatorStatus.Active;
        $._delegatorStakes[delegationID].startTime = uint64(block.timestamp);
        $._lockedNFTs[delegationID] = tokenIDs;

        $._posValidatorInfo[validationID].totalTokens += tokenIDs.length;

        emit InitiatedDelegatorRegistration({
            delegationID: delegationID,
            validationID: validationID,
            delegatorAddress: delegatorAddress,
            nonce: nonce,
            validatorWeight: validator.weight,
            delegatorWeight: weight,
            setWeightMessageID: 0
        });

        emit CompletedDelegatorRegistration({
            delegationID: delegationID,
            validationID: validationID,
            startTime: uint64(block.timestamp)
        });

        emit DelegatedNFTs(delegationID, tokenIDs);

        return delegationID;
    }

   /**
    * @notice Initiates the process of ending an NFT delegation for a given delegation ID.
    * @dev This function ensures that the delegation is active and validates that the caller is authorized to end it.
    *      If the validator status is valid, the delegation status is updated to `PendingRemoved`. If the validator
    *      is complete, then removal is completed directly. Status is updated to `Completed` and initate
    *      `InitiatedDelegatorRemoval` is not emitted. 
    * @param delegationID The unique identifier of the NFT delegation to be ended.
    *
    * Reverts if:
    * - The delegation is not active (`InvalidDelegatorStatus`).
    * - The caller is not authorized to end the delegation (`UnauthorizedOwner`).
    * - The minimum stake duration has not passed for the validator or the delegator (`MinStakeDurationNotPassed`).
    * - The validator is not in a valid state to end the delegation (`InvalidValidatorStatus`).
    */
    function _initiateNFTDelegatorRemoval(
        bytes32 delegationID
    ) internal {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;

        Validator memory validator = $._manager.getValidator(validationID);

        // Ensure the delegator is active
        if (delegator.status != DelegatorStatus.Active) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        if (delegator.owner != _msgSender()) {
            revert UnauthorizedOwner(_msgSender());
        }

        if (validator.status == ValidatorStatus.Active || validator.status == ValidatorStatus.Completed) {
            // Check that minimum stake duration has passed.
            if (validator.status != ValidatorStatus.Completed && block.timestamp < delegator.startTime + $._minimumStakeDuration) {
                revert MinStakeDurationNotPassed(uint64(block.timestamp));
            }

            $._delegatorStakes[delegationID].status = DelegatorStatus.PendingRemoved;
            $._delegatorStakes[delegationID].endTime = uint64(block.timestamp);
            if (validator.status == ValidatorStatus.Completed) {
                $._delegatorStakes[delegationID].endTime = validator.endTime; 
            }
            emit InitiatedDelegatorRemoval(delegationID, validationID);
        } else {
            revert InvalidValidatorStatus(validator.status);
        }
    }

    /**
    * @notice Completes the process of ending an NFT delegation and returns the associated token IDs.
    * @dev This function removes the delegation from the validator and account, retrieves the associated NFTs,
    *      and clears the delegation data from storage. It emits a `DelegationEnded` event upon completion.
    * @param delegationID The unique identifier of the NFT delegation to be completed.
    * @return tokenIDs An array of token IDs associated with the completed delegation.
    *
    * Emits:
    * - `DelegationEnded` when the delegation is successfully completed and removed from storage.
    */
    function _completeNFTDelegatorRemoval(
        bytes32 delegationID
    ) internal returns (uint256[] memory tokenIDs) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;

        tokenIDs = $._lockedNFTs[delegationID];

        $._posValidatorInfo[validationID].totalTokens -= tokenIDs.length;
        $._delegatorStakes[delegationID].status = DelegatorStatus.Removed;

        emit CompletedDelegatorRemoval(delegationID, validationID, 0, 0);

        return tokenIDs;
    }

    /**
    * @notice Updates the uptime of a validator based on a verified ValidationUptimeMessage received via Warp.
    * @dev This function extracts the uptime from a Warp message, validates its authenticity, and updates the
    *      stored uptime for the specified validator if the provided uptime is greater than the currently stored uptime.
    *      It also updates the validator's epoch information and balance trackers for both standard and NFT delegations.
    * @param validationID The unique identifier of the validator whose uptime is being updated.
    * @param messageIndex The index of the Warp message in the Warp messenger to validate and process.
    * @return The updated uptime for the specified validator, or the current uptime if no update is performed.
    *
    * Reverts if:
    * - The Warp message is invalid.
    * - The source chain ID in the Warp message does not match the expected uptime blockchain ID.
    * - The origin sender address in the Warp message is not the zero address.
    * - The `validationID` in the Warp message payload does not match the provided `validationID`.
    *
    * Emits:
    * - `UptimeUpdated` event when the uptime is successfully updated for a validator.
    */
    function _updateUptime(bytes32 validationID, uint32 messageIndex) internal override returns (uint64) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        if ($._uptimeKeeper != _msgSender()) {
            revert OwnableUnauthorizedAccount(_msgSender());
        }
        
        uint64 uptime = _validateUptime(validationID, messageIndex);
        uint64 epoch = _getEpoch() - 1;
        uint64 dur = $._epochDuration;

        PoSValidatorInfo storage validatorInfo = $._posValidatorInfo[validationID];

        if(validatorInfo.uptimeSeconds >= uptime ){
            return validatorInfo.uptimeSeconds;
        }

        uint256 validationUptime = uptime - validatorInfo.uptimeSeconds;
        if (validationUptime * 100 / dur >= UPTIME_REWARDS_THRESHOLD_PERCENTAGE){
            validationUptime = dur;
        }

        // Calculate validator weights
        uint256 valWeight = $._manager.getValidator(validationID).startingWeight * validationUptime / dur;
        uint256 valWeightNFT = (validatorInfo.tokenIDs.length * 1e6) * validationUptime / dur;

        // Update reward weights for validator owner
        $._accountRewardWeight[epoch][validatorInfo.owner] += valWeight;
        $._accountRewardWeightNFT[epoch][validatorInfo.owner] += valWeightNFT;
        $._totalRewardWeight[epoch] += valWeight;
        $._totalRewardWeightNFT[epoch] += valWeightNFT;

        validatorInfo.uptimeSeconds = uptime;

        $._validationUptimes[epoch][validationID] += validationUptime;
        
        emit UptimeUpdated(validationID, uptime, epoch);
        return uptime;
    }

    /**
    * @notice Calculated and updates the reward weight of the given delegations for the previous epoch after successfully 
    *         submitting the uptime of the respective validators
    * @param delegationIDs An array of delegation IDs associated with the staking process.
    */
    function resolveRewards(bytes32[] memory delegationIDs) external {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        if ($._uptimeKeeper != _msgSender()) {
            revert OwnableUnauthorizedAccount(_msgSender());
        }
        
        uint64 epoch = _getEpoch() - 1;
        uint64 dur = $._epochDuration;

        for (uint256 i = 0; i < delegationIDs.length; i++) {
            Delegator memory delegator = $._delegatorStakes[delegationIDs[i]];
            PoSValidatorInfo storage validatorInfo = $._posValidatorInfo[delegator.validationID];

            uint64 epochStart = epoch * dur;
            uint64 epochEnd = epochStart + dur;

            uint64 delegationUptime;
            {
                uint64 delegationStart = uint64(Math.max(delegator.startTime, epochStart));
                uint64 delegationEnd = delegator.endTime != 0 ? delegator.endTime : epochEnd;
                if (epochStart > delegationEnd){ continue; }
                delegationUptime = uint64(Math.min(delegationEnd - delegationStart, $._validationUptimes[epoch][delegator.validationID]));

                if (delegationUptime * 100 / dur >= UPTIME_REWARDS_THRESHOLD_PERCENTAGE){
                    delegationUptime = dur;
                }
            }
            uint256 delWeight = (delegator.weight * delegationUptime) / dur;
            uint256 feeWeight = (delWeight * validatorInfo.delegationFeeBips) / BIPS_CONVERSION_FACTOR;

            if ($._lockedNFTs[delegationIDs[i]].length == 0){
                $._accountRewardWeight[epoch][validatorInfo.owner] += feeWeight;
                $._accountRewardWeight[epoch][delegator.owner] += delWeight - feeWeight;
                $._totalRewardWeight[epoch] += delWeight;
            } else {
                $._accountRewardWeightNFT[epoch][validatorInfo.owner] += feeWeight;
                $._accountRewardWeightNFT[epoch][delegator.owner] += delWeight - feeWeight;
                $._totalRewardWeightNFT[epoch] += delWeight;
            }
            emit RewardResolved(delegationIDs[i], epoch);
        }        
    }

    /**
     * @notice Validates the uptime proof for a given validator.
     * @dev This function checks whether the provided validation ID corresponds to a PoS validator,
     *      verifies the associated warp message, and extracts the uptime value.
     * @param validationID The unique identifier of the validator whose uptime is being validated.
     * @param messageIndex The index of the warp message to retrieve and validate.
     * @return uptime The validated uptime value extracted from the warp message.
     *
     * Requirements:
     * - The validation ID must belong to a PoS validator.
     * - The referenced warp message must be valid and verified.
     * - The source chain ID of the warp message must match the expected uptime blockchain ID.
     * - The origin sender address of the warp message must be the zero address, ensuring direct validator signing.
     * - The extracted validation ID from the warp message must match the provided `validationID`.
     *
     * Reverts:
     * - `ValidatorNotPoS` if the validator is not a PoS validator.
     * - `InvalidWarpMessage` if the warp message is not valid.
     * - `InvalidWarpSourceChainID` if the message originates from an incorrect blockchain.
     * - `InvalidWarpOriginSenderAddress` if the sender is not the zero address.
     * - `UnexpectedValidationID` if the extracted validation ID does not match the provided one.
     */
    function _validateUptime(bytes32 validationID, uint32 messageIndex) internal view returns (uint64) {
        if (!_isPoSValidator(validationID)) {
            revert ValidatorNotPoS(validationID);
        }

        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        if (!valid) {
            revert InvalidWarpMessage();
        }

        StakingManagerStorage storage $ = _getStakingManagerStorage();
        // The uptime proof must be from the specifed uptime blockchain
        if (warpMessage.sourceChainID != $._uptimeBlockchainID) {
            revert InvalidWarpSourceChainID(warpMessage.sourceChainID);
        }

        // The sender is required to be the zero address so that we know the validator node
        // signed the proof directly, rather than as an arbitrary on-chain message
        if (warpMessage.originSenderAddress != address(0)) {
            revert InvalidWarpOriginSenderAddress(warpMessage.originSenderAddress);
        }

        (bytes32 uptimeValidationID, uint64 uptime) =
            ValidatorMessages.unpackValidationUptimeMessage(warpMessage.payload);
        if (validationID != uptimeValidationID) {
            revert UnexpectedValidationID(uptimeValidationID, validationID);
        }

        return uptime;
    }
}