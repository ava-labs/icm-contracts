# Validator Manager Contracts

The contracts in this directory define the Validator Manager used to manage Avalanche L1 validators, as defined in [ACP-77](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets). They comply with [ACP-99](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/99-validatorsetmanager-contract), which specifies the standard minimal functionality that Validator Managers should implement. The contracts in this directory are are related as follows:

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram
class ACP99Manager {
    +initializeValidatorSet()
    +completeValidatorRegistration()
    +completeValidatorRemoval()
    +completeValidatorWeightUpdate()
    -_initiateValidatorRegistration()
    -_initiateValidatorRemoval()
    -_initiateValidatorWeightUpdate()
}
<<Abstract>> ACP99Manager

class ValidatorManager {
    +initializeValidatorSet()
    +completeValidatorRegistration() onlyOwner
    +completeValidatorRemoval() onlyOwner
    +completeValidatorWeightUpdate() onlyOwner
    +initiateValidatorRegistration() onlyOwner
    +initiateValidatorRemoval() onlyOwner
    +initiateValidatorWeightUpdate() onlyOwner
}

class PoAManager {
    +completeValidatorRegistration()
    +completeValidatorRemoval()
    +completeValidatorWeightUpdate()
    +initiateValidatorRegistration() onlyOwner
    +initiateValidatorRemoval() onlyOwner
    +initiateValidatorWeightUpdate() onlyOwner
    +transferValidatorManagerOwnership() onlyOwner
}

class StakingManager {
    +completeValidatorRegistration()
    +initiateValidatorRemoval()
    +completeValidatorRemoval()
    +completeDelegatorRegistration()
    +initiateDelegatorRemoval()
    +completeDelegatorRemoval()
    -_initiateValidatorRegistration()
    -_initiateDelegatorRegistration()
}
<<Abstract>> StakingManager
class ERC20TokenStakingManager {
    +initiateValidatorRegistration()
    +initiateDelegatorRegistration()
}
class NativeTokenStakingManager {
    +initiateValidatorRegistration() payable
    +initiateDelegatorRegistration() payable
}

ACP99Manager <|-- ValidatorManager
ValidatorManager --o PoAManager : owner
ValidatorManager --o StakingManager : owner
StakingManager <|-- ERC20TokenStakingManager
StakingManager <|-- NativeTokenStakingManager
```

## A Note on Nomenclature

The contracts in this directory are only useful to L1s that have been converted from Subnets as described in ACP-77. As such, `l1`/`L1` is generally preferred over `subnet`/`Subnet` in the source code. The one major exception is that `subnetID` should be used to refer to both Subnets that have not been converted, and L1s that have. This is because an L1 must first be initialized as a Subnet by issuing a `CreateSubnetTx` on the P-Chain, the transaction hash of which becomes the `subnetID`. Rather than change the name and/or value of this identifier, it is simpler for both to remain static in perpetuity.

## Deploying

The validator manager system consists of a `ValidatorManager`, and one of `NativeTokenStakingManager`, `ERC20TokenStakingManager`, or `PoAManager`. `ValidatorManager` is `Ownable`, and its owner should be set to the address of the other contract.

All of these are implemented as [upgradeable](https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/3d6a15108b50491ec3c51c03b32802c33e092a0f/contracts/proxy/utils/Initializable.sol#L56) contracts. There are numerous [guides](https://blog.chain.link/upgradable-smart-contracts/) for deploying upgradeable smart contracts, but the general steps are as follows:

1. Deploy the implementation contract
2. Deploy the proxy contract
3. Call the implementation contract's `initialize` function

- Each deployed contract requires different settings. For example, `ValidatorManagerSettings` specifies the churn parameters, while `StakingManagerSettings` specifies the staking and rewards parameters.

4. Initialize the validator set by calling `initializeValidatorSet` on `ValidatorManager`

- When an L1 is first created on the P-Chain, it must be explicitly converted to an L1 via [`ConvertSubnetToL1Tx`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#convertsubnettol1tx). The resulting `SubnetToL1ConversionMessage` ICM [message](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#subnettol1conversionmessage) is provided in the call to `initializeValidatorSet` to specify the starting validator set in the `ValidatorManager`. Regardless of the setup of the overall validator manager system, these initial validators are treated as PoA and are not eligible for staking rewards.

### Proof-of-Authority

PoA validator management is provided by `PoAManager` by providing an `owner` in the call to `initialize`. Only the `owner` may initiate validator set changes, but anybody can complete the validator set change by providing the corresponding ICM message signed by the P-Chain.

> [!NOTE]
> PoA validator management can also be implemented by `ValidatorManager` on its own, by setting the `owner` to the desired admin address. Unlike `PoAManager`, only the admin is able to initiate or complete validator set changes.

### Proof-of-Stake

PoS validator management is provided by the abstract contract `StakingManager`, which has two concrete implementations: `NativeTokenStakingManager` and `ERC20TokenStakingManager`. `StakingManager` supports uptime-based validation rewards, as well as delegation to a chosen validator. The `uptimeBlockchainID` used to initialize the `StakingManager` **must** be validated by the L1 validator set that the contract manages. **There is no way to verify this from within the contract, so take care when setting this value.** This [state transition diagram](./StateTransition.md) illustrates the relationship between validators and delegators. After deploying `StakingManager` and a proxy, call the `initialize` function, which takes a `StakingManagerSettings` as well as any implementation-specific arguments.

> [!NOTE]
> The `weightToValueFactor` fields of `StakingManagerSettings` sets the factor used to convert between the weight that the validator is registered with on the P-Chain, and the value transferred to the contract as stake. This involves integer division, which may result in loss of precision. When selecting `weightToValueFactor`, it's important to make the following considerations:
>
> 1. If `weightToValueFactor` is near the denomination of the asset, then staking amounts on the order of 1 unit of the asset may cause the converted weight to round down to 0. This may impose a larger-than-expected minimum stake amount.
>    - Ex: If USDC (denomination of 6) is used as the staking token and `weightToValueFactor` is 1e9, then any amount less than 1,000 USDC will round down to 0 and therefore be invalid.
> 2. Staked amounts up to `weightValueFactor - 1` may be lost in the contract as dust, as the validator's registered weight is used to calculate the original staked amount.
>    - Ex: `value=1001` and `weightToValueFactor=1e3`. The resulting weight will be `1`. Converting the weight back to a value results in `value=1000`.
> 3. The validator's weight is represented on the P-Chain as a `uint64`. `StakingManager` restricts values such that the calculated weight does not exceed the maximum value for that type.

### Migrating from Proof-of-Authority to Proof-of-Stake

See the [migration guide](./PoAMigration.md) for details.

#### NativeTokenStakingManager

`NativeTokenStakingManager` allows permissionless addition and removal of validators that post the L1's native token as stake. Staking rewards are minted via the Native Minter Precompile, which is configured with a set of addresses with minting privileges. As such, the address that `NativeTokenStakingManager` is deployed to must be added as an admin to the precompile. This can be done by either calling the precompile's `setAdmin` method from an admin address, or setting the address in the Native Minter precompile settings in the chain's genesis (`config.contractNativeMinterConfig.adminAddresses`). There are a couple of methods to get this address: one is to calculate the resulting deployed address based on the deployer's address and account nonce: `keccak256(rlp.encode(address, nonce))`. The second method involves manually placing the `NativeTokenStakingManager` bytecode at a particular address in the genesis, then setting that address as an admin.

```json
{
    "config" : {
        ...
        "contractNativeMinterConfig": {
            "blockTimestamp": 0,
            "adminAddresses": [
                "0xffffffffffffffffffffffffffffffffffffffff"
            ]
        }
    },
    "alloc": {
        "0xffffffffffffffffffffffffffffffffffffffff": {
            "balance": "0x0",
            "code": "<NativeTokenStakingManagerByteCode>",
            "nonce": 1
        }
    }
}
```

#### ERC20TokenStakingManager

`ERC20TokenStakingManager` allows permissionless addition and removal of validators that post the an ERC20 token as stake. The ERC20 is specified in the call to `initialize`, and must implement [`IERC20Mintable`](./interfaces/IERC20Mintable.sol). Care should be taken to enforce that only authorized users are able to `mint` the ERC20 staking token.

## Usage


### Register a Validator

#### PoA

Validator registration is initiated with a call to `PoAManager.initiateValidatorRegistration`. Churn limitations are checked - only a certain (configurable) percentage of the total weight is allowed to be added or removed in a (configurable) period of time. The `ValidatorManager` then constructs a [`RegisterL1ValidatorMessage`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#registerl1validatormessage) ICM message to be sent to the P-Chain. Each validator registration request includes all of the information needed to identify the validator and its stake weight, as well as an `expiry` timestamp before which the `RegisterL1ValidatorMessage` must be delivered to the P-Chain. If the validator is not registered on the P-Chain before the `expiry`, then the validator may be removed from the contract state by calling `completeValidatorRemoval`.

The `RegisterL1ValidatorMessage` is delivered to the P-Chain as the ICM message payload of a `RegisterL1ValidatorTx`. Please see the transaction [specification](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#registerl1validatortx) for validity requirements. The P-Chain then signs a [`L1ValidatorRegistrationMessage`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#l1validatorregistrationmessage) ICM message indicating that the specified validator was successfully registered on the P-Chain.

The `L1ValidatorRegistrationMessage` is delivered by calling `ValidatorManager.completeValidatorRegistration`.

#### PoS

When registering a PoS validator, the same steps as the PoA case apply, with the only difference being that `StakingManager.initiateValidatorRegistration` and `StakingManager.completeValidatorRegistration` must be called instead. 

The sender of the transaction that called `StakingManager.initiateValidatorRegistration` is registered as the validator owner. Only this owner can remove the validator. 

Staking rewards begin accruing once `StakingManager.completeValidatorRegistration` is called.

### Remove a Validator

### PoA

Validator exit is initiated with a call to `PoAManager.initiateValidatorRemoval`. The `ValidatorManager` contructs an [`L1ValidatorWeightMessage`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#l1validatorweightmessage) ICM message with the weight set to `0`. This is delivered to the P-Chain as the payload of a [`SetL1ValidatorWeightTx`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#setl1validatorweighttx). The P-Chain acknowledges the validator exit by signing an `L1ValidatorRegistrationMessage` with `valid=0`, which is delivered by calling `ValidatorManager.completeValidatorRemoval`. The validation is removed from the contract's state.

### PoS

PoS validator removal follows the same flow as the PoA case, except that `StakingManager.initiateValidatorRemoval` and `StakingManager.completeValidatorRemoval` must be called instead.

There are two additional considerations:

- A [`ValidationUptimeMessage`](./UptimeMessageSpec.md) ICM message may optionally be provided in the call to `StakingManager.initiateValidatorRemoval` in order to calculate the staking rewards; otherwise the latest received uptime will be used (see [(PoS only) Submit and Uptime Proof](#pos-only-submit-an-uptime-proof)). This proof may be requested directly from the L1 validators, which will provide it in a `ValidationUptimeMessage` ICM message. If the uptime is not sufficient to earn validation rewards, the call to `initiateValidatorRemoval` will fail. `forceInitiateValidatorRemoval` acts the same as `initiateValidatorRemoval`, but bypasses the uptime-based rewards check. Once `initiateValidatorRemoval` or `forceInitiateValidatorRemoval` is called, staking rewards cease accruing for `StakingManagers`.

- Unlike with PoA, PoS validators are not able to decrease their weight. This can lead to a scenario in which a PoS validator manager with a high proportion of the L1's weight is not able to exit the validator set due to churn restrictions. Additional validators or delegators will need to first be registered to more evenly distribute weight across the L1's validator set.

Once acknowledgement from the P-Chain has been received via a call to  `StakingManager.completeValidatorRemoval`, staking rewards are disbursed and stake is returned.

#### Disable a Validator Directly on the P-Chain

ACP-77 also provides a method to disable a validator without interacting with the L1 directly. The P-Chain transaction [`DisableL1ValidatorTx`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#disablel1validatortx) disables the validator on the P-Chain. The disabled validator's weight will still count towards the L1's total weight.

Disabled L1 validators can re-activate at any time by increasing their balance with an `IncreaseBalanceTx`. Anyone can call `IncreaseBalanceTx` for any validator on the P-Chain. A disabled validator can only be completely and permanently removed from the validator set by a call to `initiateValidatorRemoval`.

### (PoS only) Register a Delegator

`StakingManager` supports delegation to an actively staked validator as a way for users to earn staking rewards without having to validate the chain. Delegators pay a configurable percentage fee on any earned staking rewards to the host validator. A delegator may be registered by calling `initiateDelegatorRegistration` and providing an amount to stake. The delegator will be registered as long as churn restrictions are not violated. The delegator is reflected on the P-Chain by adjusting the validator's registered weight via a [`SetL1ValidatorWeightTx`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#setl1validatorweighttx). The weight change acknowledgement is delivered to the `StakingManager` via an [`L1ValidatorWeightMessage`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#l1validatorweightmessage), which is provided by calling `completeDelegatorRegistration`.

> [!NOTE]
> The P-Chain is only willing to sign an `L1ValidatorWeightMessage` for a validator in the L1's validator set. Once a validator exit has been initiated (via a call to `initiateValidatorRemoval`), the `StakingManager` must assume that the validator has been removed from the L1's validator set on the P-Chain, and therefore that the P-Chain will not sign any further weight updates. The contracts prohibit _initiating_ adding or removing a delegator in between calls to `initiateValidatorRemoval` and `completeValidatorRemoval`. However, if the `L1ValidatorWeightMessage` pertaining to an already initiated delegator action is constructed _before_ the validator is removed on the P-Chain, then the delegator action may be completed. Otherwise, `completeValidatorRemoval` must be called before completing the delegator action.

### (PoS only) Remove a Delegator

Delegator removal may be initiated by calling `initiateDelegatorRemoval`, as long as churn restrictions are not violated. Similar to `initiateValidatorRemoval`, an uptime proof may be provided to be used to determine delegator rewards eligibility. If no proof is provided, the latest known uptime will be used (see [(PoS only) Submit and Uptime Proof](#pos-only-submit-an-uptime-proof)). The validator's weight is updated on the P-Chain by the same mechanism used to register a delegator. The `L1ValidatorWeightMessage` from the P-Chain is delivered to the `StakingManager` in the call to `completeDelegatorRemoval`.

Either the delegator owner or the validator owner may initiate removing a delegator. This is to prevent the validator from being unable to remove itself due to churn limitations if it is has too high a proportion of the Subnet's total weight due to delegator additions. The validator owner may only remove Delegators after the minimum stake duration has elapsed.

### (PoS only) Submit an Uptime Proof

The [rewards calculator](./interfaces/IRewardCalculator.sol) is a function of uptime seconds since the validator's start time. In addition to doing so in the calls to `initiateValidatorRemoval` and `initiateDelegatorRemoval` as described above, uptime proofs may also be supplied by calling `submitUptimeProof`. Unlike `initiateValidatorRemoval` and `initiateDelegatorRemoval`, `submitUptimeProof` may be called by anyone, decreasing the likelihood of a validation or delegation not being able to claim rewards that it deserved based on its actual uptime.

### (PoS only) Collect Staking Rewards

#### Validation Rewards

Validation rewards are distributed in the call to `completeValidatorRemoval` on the `StakingManager`.

#### Delegation Rewards

Delegation rewards are distributed in the call to `completeDelegatorRemoval` on the `StakingManager`.

#### Delegation Fees

Delegation fees owed to validators are _not_ distributed when the validation ends as to bound the amount of gas consumed in the call to `completeValidatorRemoval`. Instead, `claimDelegationFees` on the `StakingManager` may be called after the validation is completed.
