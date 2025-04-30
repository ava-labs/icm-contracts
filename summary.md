# Validator Management System Documentation

## Overview
The ValidatorManager, Native721TokenStakingManager, and StakingManager contracts form a system for managing validator staking using ERC721 tokens (NFTs) and native tokens. These contracts are part of the BEAM ecosystem and handle validator registration, delegation, and reward distribution.

## Core Contracts

### ValidatorManager
The main contract responsible for managing validators and their lifecycle.

Key Functions:
```solidity
function initialize(ValidatorManagerSettings calldata settings)
function initializeValidatorSet(ConversionData calldata conversionData, uint32 messageIndex)
function registerValidator(
    bytes memory nodeID,
    bytes memory blsPublicKey,
    uint64 registrationExpiry,
    PChainOwner memory remainingBalanceOwner,
    PChainOwner memory disableOwner
) external returns (bytes32)
function completeValidatorRemoval(uint32 messageIndex) external returns (bytes32)
function getValidator(bytes32 validationID) external view returns (Validator memory)
```

### StakingManager
Abstract base contract providing core staking functionality.

Key Functions:
```solidity
function initiateValidatorRegistration(
    bytes memory nodeID,
    bytes memory blsPublicKey,
    uint64 registrationExpiry,
    PChainOwner memory remainingBalanceOwner,
    PChainOwner memory disableOwner,
    uint16 delegationFeeBips,
    uint64 minStakeDuration
) external payable returns (bytes32)
function initiateDelegatorRegistration(bytes32 validationID) external payable returns (bytes32)
function completeValidatorRemoval(uint32 messageIndex) external returns (bytes32)
function getRewards(
    bool primary,
    uint64 epoch,
    address[] memory tokens,
    address account
) external view returns (uint256[] memory)
```

### Native721TokenStakingManager
Extension of StakingManager for ERC721 token staking.

Key Functions:
```solidity
function initialize(StakingManagerSettings calldata settings, IERC721 stakingToken)
function initiateValidatorRegistration(
    bytes memory nodeID,
    bytes memory blsPublicKey,
    uint64 registrationExpiry,
    PChainOwner memory remainingBalanceOwner,
    PChainOwner memory disableOwner,
    uint16 delegationFeeBips,
    uint64 minStakeDuration,
    uint256[] memory tokenIDs
) external payable returns (bytes32)
function registerNFTDelegation(bytes32 validationID, uint256[] memory tokenIDs) external returns (bytes32)
function claimRewards(
    bool primary,
    uint64 epoch,
    address[] memory tokens,
    address recipient
) external
function registerRewards(
    bool primary,
    uint64 epoch,
    address token,
    uint256 amount
) external
```

## Detailed Functionality

### Validator Management
- **Validator Registration**
  - Requires nodeID, BLS public key, and registration expiry
  - Supports both native token and NFT staking
  - Minimum stake requirements must be met
  - Delegation fee must be within allowed range (0-10000 bips)
  - Validators can set minimum stake duration

- **Validator Removal**
  - Can be initiated by validator or admin
  - Requires unlock duration to pass before stake can be withdrawn
  - Automatically handles NFT unlocking for NFT-based validators

### Delegation System
- **NFT Delegation**
  - Delegators can stake NFTs to support validators
  - Each NFT contributes to the validator's weight
  - Maximum NFT amount per validator is enforced
  - Delegation fee is automatically distributed

- **Delegation Management**
  - Delegators can remove their stake after unlock period
  - Supports redelegation to different validators
  - Tracks delegation start and end times
  - Handles delegation fee distribution

### Reward System
- **Reward Pools**
  - Dual reward system: Primary and NFT-based pools
  - Primary pool for native token stakers
  - NFT pool for NFT-based stakers
  - Each pool has independent reward distribution

- **Reward Distribution**
  - Epoch-based distribution system
  - Rewards are distributed based on:
    - Stake weight
    - Uptime performance (80% threshold)
    - Delegation fee percentage
  - Formula: `reward = (pool_amount * stake_weight * uptime_factor) / total_weight`

- **Reward Management**
  - Rewards can be registered by contract owner
  - 7-day claim delay after epoch end
  - Rewards can be cancelled before claim period starts
  - Supports multiple reward tokens per epoch

- **Reward Calculation**
  ```solidity
  // For primary rewards
  rewards[i] = (reward_pool[epoch][token] * account_weight[epoch][account]) 
               / total_weight[epoch] 
               - withdrawn_rewards[epoch][account][token]

  // For NFT rewards
  rewards[i] = (reward_pool_nft[epoch][token] * account_weight_nft[epoch][account]) 
               / total_weight_nft[epoch] 
               - withdrawn_rewards_nft[epoch][account][token]
  ```

- **Uptime Requirements**
  - Validators and delegators receive full epoch rewards if they maintain 80% uptime during the epoch
  - If uptime is below 80%, rewards are prorated based on actual uptime
  - Uptime is calculated as: `uptime = (actual_uptime_seconds * 100) / epoch_duration`
  - When uptime ≥ 80%, the full epoch duration is used for reward calculation
  - When uptime < 80%, the actual uptime is used for reward calculation
  - Uptime proofs can be submitted by authorized keepers

## Events
- DelegatedNFTs
- RewardRegistered
- RewardCancelled
- RewardClaimed
- Validator registration and removal events
- Delegation events

## User Flows

### Validator Management
1. **Validator Registration**
   ```
   initiateValidatorRegistration (with stake + NFTs) 
   --> registerValidator (relayer) 
   --> validator becomes active
   ```

2. **Validator Removal**
   ```
   initiateValidatorRemoval 
   --> completeValidatorRemoval (relayer) 
   --21days--> unlockValidator
   ```

3. **Expired Validator Removal**
   ```
   initiateValidatorRegistration 
   --> completeValidatorRemoval (relayer or anyone)
   ```

### Delegation Management
1. **Native Token Delegation**
   ```
   initiateDelegatorRegistration (with native tokens) 
   --> registerDelegator (relayer) 
   --> delegation becomes active
   ```

2. **NFT Delegation**
   ```
   registerNFTDelegation (with NFTs) 
   --> delegation becomes active
   ```

3. **Native Token Delegator Removal**
   ```
   initiateDelegatorRemoval 
   --> completeDelegatorRemoval (relayer) 
   --21days--> unlockDelegator (native tokens)
   ```

4. **NFT Delegator Removal**
   ```
   initiateNFTDelegatorRemoval 
   --21days--> completeNFTDelegatorRemoval
   ```

5. **Native Token Redelegation**
   ```
   initiateDelegatorRemoval 
   --> completeDelegatorRemoval (relayer) 
   --> initiateDelegatorRegistration (with new validator)
   ```

6. **NFT Redelegation**
   ```
   initiateNFTDelegatorRemoval 
   --> completeNFTDelegatorRemoval (relayer) 
   --> registerNFTDelegation (with new validator)
   ```

### Reward Management
1. **Reward Registration**
   ```
   registerRewards (owner) 
   --> rewards available after epoch end + 7 days
   ```

2. **Reward Claiming**
   ```
   getRewards (check available rewards) 
   --> claimRewards (after 7-day delay)
   ```

3. **Reward Cancellation**
   ```
   cancelRewards (owner, before claim period starts)
   ```

### Uptime Management
1. **Uptime Proof Submission**
   ```solidity
   function submitUptimeProofs(
       bytes32[] memory validationIDs,
       uint32[] memory messageIndexes
   ) external
   ```
   - Can only be called by the uptime keeper
   - Submits uptime proofs for multiple validators in a single transaction
   - Each validator's uptime is validated against Warp messages
   - Updates the validator's uptime record for the current epoch
   - If uptime ≥ 80%, full epoch duration is used for reward calculation
   - If uptime < 80%, actual uptime is used for reward calculation

2. **Reward Resolution**
   ```solidity
   function resolveRewards(bytes32[] memory delegationIDs) external
   ```
   - Can only be called by the uptime keeper
   - Calculates and updates reward weights for delegations
   - Processes both native token and NFT delegations
   - For each delegation:
     - Calculates delegation uptime based on validator uptime
     - Computes reward weights considering:
       - Delegation duration within epoch
       - Validator's delegation fee
       - Uptime performance
     - Updates reward weights for both validator and delegator
   - Reward distribution:
     - Validator receives fee portion of rewards
     - Delegator receives remaining rewards
     - Weights are stored for later reward claiming

3. **Reward Calculation Flow**
   ```
   submitUptimeProofs (update validator uptime)
   --> resolveRewards (calculate delegation weights)
   --> getRewards (check available rewards)
   --> claimRewards (after 7-day delay)
   ```

## Unlock Periods and Functions

### Unlock Periods
- All unlock periods are 21 days by default but configurable through contract settings
- Unlock periods start after the completion of removal operations
- During the unlock period:
  - Stakes remain locked
  - No rewards are earned
  - No operations can be performed on the locked assets

### Unlock Functions
1. **Native Token Unlocking**
   ```solidity
   function unlockValidator(bytes32 validationID)
   function unlockDelegator(bytes32 delegationID)
   ```
   - Unlocks native tokens after the unlock period
   - Can only be called after the unlock period has passed
   - Transfers native tokens back to the owner

2. **NFT Unlocking**
   - NFT unlocking is handled automatically within `completeNFTDelegatorRemoval`
   - No separate unlock function is needed
   - NFTs are transferred back to the owner as part of the completion process

### Unlock Process
1. **Validator Unlock**
   ```
   completeValidatorRemoval 
   --21days--> unlockValidator 
   --> unlocks both native tokens and NFTs
   ```

2. **Native Token Delegator Unlock**
   ```
   completeDelegatorRemoval 
   --21days--> unlockDelegator 
   --> unlocks native tokens
   ```

3. **NFT Delegator Unlock**
   ```
   completeNFTDelegatorRemoval 
   --> automatically unlocks and transfers NFTs back to owner
   ```

Note: The unlock period is enforced to ensure security and prevent certain types of attacks.