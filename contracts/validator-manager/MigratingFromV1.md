# Migrating From V1 Validator Manager Contracts

The V1 Validator Manager contracts are implemented as a single deployed contract consisting of multiple contracts related through inheritance. The V2 Validator Manager, on the other hand, may be consist of multiple deployed contracts that interact via external function calls.

> Note: This guide only applies to validator manager contracts at or above [validator-manager-v1.0.0](https://github.com/ava-labs/icm-contracts/releases/tag/validator-manager-v1.0.0). Contracts at commits before then must use other methods to convert to V2.

## Migrating Proof-of-Authority

In the V2 contracts, Proof-of-Authority validator management is implemented fully in `ValidatorManager`, so there is no need to deploy a separate contract (contrast with the Proof-of-Stake case, which requires a standalone `StakingManager` deployment). Additionally, both the V1 `PoAValidatorManager` and the V2 `ValidatorManager` are `Ownable`, meaning that the V1 `PoAValidatorManager` owner will persist after upgrading to a V2 `ValidatorManager`. Migrating from a V1 to V2 therefore only requires migrating the state variables that track validators. The helper method `migrateFromV1` is provided for this purpose. The general steps are as follows:

1. Upgrade the proxy contract's implementation to a newly deployed V2 `ValidatorManager` using standard methods
2. For each validator (active or expired), call `migrateFromV1` with its `validationID`. The `receivedNonce` argument can be set to 0, since V1 `PoAValidatorManager`s do not support weight changes.

Some notes on this process:
- `migrateFromV1` may only be called once per `validationID`. Any active validators will not be able to be removed until they are migrated.
- It is up to the contract owner to track all of the `validationID`s to be migrated. This is because Solidity mappings are not iterable, so there is no getter to retreive all `validationID`s as a batch.

## Migrating Proof-of-Stake

**Migrating from a V1 `PoSValidatorManager` deployment to a V2 Proof-of-Stake Validator Manager is not supported.**

V2 Proof-of-Stake Validator Managers consist of two deployed contracts: an instance of a concrete implementation of `StakingManager`, and an instance of a `ValidatorManager`. Because `StakingManager` is a standalone contract that does not share state with `ValidatorManager`, migrating from a V1 `PoSValidatorManager` to a V2 `StakingManager` would require replicating state from one contract to another. 