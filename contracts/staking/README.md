# Validator Manager Contract

`ValidatorManager.sol` defines the abstract staking contract used to manage subnet-only validators, as defined in [ACP-77](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets).

_Disclaimer:_ The contracts in this directory are still under active development, are unaudited, and should not be used in production.

## Usage

The common case for adding or removing a validator from a Subnet's validator set follows these general steps:

1. Inform the Validator Manager contract of the change. The Validator Manager contract will construct a Warp message attesting to the change in the validator set.
2. Deliver the Warp message containing the validator set change to the P-Chain. The P-Chain will construct a Warp message acknowledging the change to the validator set.
3. Deliver the Warp message containing the P-Chain acknowledgement back to the Validator Manager contract on the Subnet. The Validator Manager contract finalizes the validator set change.

### Initializing a Validator Set

If the Subnet has no validators registered on the P-Chain, then it will not be able to sign a Warp message to add a validator (step 1 above). In this case, Subnet recovery must be performed on the P-Chain, which allows an initial validator set to be specified, along with each validator's expiry. These validators are *not* registered with the Subnet's Validator Manager, and are only used to bootstrap subnet consensus and Warp message signing. The following steps describe how to register validators with the Subnet's Validator Manager, after which the initial validator set specified during Subnet recovery may safely exit the validator set or expire.

### Registering a Validator

Validator registration is initiated with a call to `initializeValidatorRegistration` on the concrete Validator Manager contract. The Validator Manager contract constructs a [`RegisterSubnetValidatorMessage`](#registersubnetvalidatormessage) Warp message to be sent to the P-Chain. Each validator registration request includes all of the information needed to identify the validator and its stake weight, as well as an `expiry` timestamp before which the `RegisterSubnetValidatorMessage` must be delivered to the P-Chain.

The `RegisterSubnetValidatorMessage` is delivered to the P-Chain as the Warp message payload of a `RegisterSubnetValidatorTx`. Please see the transaction [specification](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#step-2-issue-a-registersubnetvalidatortx-on-the-p-chain) for validity requirements. The successful `RegisterSubnetValidatorTx` results in a `SubnetValidatorRegistrationMessage` Warp message indicating that the specified validation was successfully registered on the P-Chain.

The `SubnetValidatorRegistrationMessage` is delivered to the Validator Manager contract via a call to `completeValidatorRegistration`. For PoS Validator Managers, staking rewards begin accruing at this time.

#### Handling a Missed Expiry

In the case of a missed expiry, the `RegisterSubnetValidatorTx` will result in a `SubnetValidatorRegistrationMessage` Warp message with `valid=0`. This serves as proof that the corresponding validation has not started and may never start. The `SubnetValidatorRegistrationMessage` can be provided to the Validator Manager contract by calling `completeEndValidation` with `setWeightMessageType=false`.

### Exiting the Validator Set

Validator exit is initiated with a call to `initializeEndValidation` on the Validator Manager contract. For PoS Validator Managers, a `ValidationUptimeMessage` Warp message may optionally be provided in order to calculate the staking rewards; otherwise the validator's uptime will be set to `0`. This proof may be requested from the P-Chain, which will provide it in a [`ValidationUptimeMessage`](#validationuptimemessage) Warp message. Once `initializeEndValidation` is called, staking rewards cease accruing for PoS Validator Managers. The Validator Manager contract constructs a [`SetSubnetValidatorWeightMessage`](#setsubnetvalidatorweightmessage) Warp message, setting the weight to `0`.

The `SetSubnetValidatorWeightMessage` is delivered to the P-Chain as the payload of a `SetValidatorWeightTx`. The P-Chain acknowledges validator exit by signing either a `SetSubnetValidatorWeightMessage` with the `weight=0`, or a `SubnetValidatorRegistrationMessage` with `valid=0`. `completeEndValidation` is then called on the Validator Manager contract, with `setWeightMessageType` set according to the Warp message type being delivered. The validation is removed from the contract's state, and for PoS Validator Managers, staking rewards are disbursed and stake is returned.

#### Exit the Validator Set Directly on the P-Chain

ACP-77 also provides a method for validators to exit a Subnet's validator set without interacting with the Subnet directly. The P-Chain transaction [`ExitValidatorSetTx`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#exitvalidatorsettx) removes the validator from the Subnet validator set tracked on the P-Chain. The P-Chain will will sign a `SetSubnetValidatorWeightMessage` or a `SubnetValidatorRegistrationMessage`, that may then be provided to the Staking Manager contract via a call to `completeEndValidation`, as described above. Note however that without the uptime proof provided in the initial call to `initializeEndValidation`, the Staking Manager will not credit any staking rewards to the validator.

## Warp Message Format Reference

### `RegisterSubnetValidatorMessage`

Description: Used to register a Subnet validator on the P-Chain

Signed by: Subnet

Specification:

```
+--------------+----------+-----------+
|      codecID :   uint16 |   2 bytes |
+--------------+----------+-----------+
|       typeID :   uint32 |   4 bytes |
+--------------+----------+-----------+
|     subnetID : [32]byte |  32 bytes |
+--------------+----------+-----------+
|       nodeID : [32]byte |  32 bytes |
+--------------+----------+-----------+
|       weight :   uint64 |   8 bytes |
+--------------+----------+-----------+
| blsPublicKey : [48]byte |  48 bytes |
+--------------+----------+-----------+
|       expiry :   uint64 |   8 bytes |
+--------------+----------+-----------+
                          | 134 bytes |
                          +-----------+

```

### `SubnetValidatorRegistrationMessage`

Description: Used to confirm Subnet validator registration on the P-Chain

Signed by: P-Chain

Specification:

```
+--------------+----------+----------+
|      codecID :   uint16 |  2 bytes |
+--------------+----------+----------+
|       typeID :   uint32 |  4 bytes |
+--------------+----------+----------+
| validationID : [32]byte | 32 bytes |
+--------------+----------+----------+
|        valid :     bool |  1 byte  |
+--------------+----------+----------+
                          | 39 bytes |
                          +----------+
```

### `ValidationUptimeMessage`

Description: Used to provide validator uptime to the Subnet to calculate staking rewards

Signed by: P-Chain

Specification:

```
+--------------+----------+----------+
|      codecID :   uint16 |  2 bytes |
+--------------+----------+----------+
|       typeID :   uint32 |  4 bytes |
+--------------+----------+----------+
| validationID : [32]byte | 32 bytes |
+--------------+----------+----------+
|       uptime :   uint64 |  8 bytes |
+--------------+----------+----------+
                          | 46 bytes |
                          +----------+
```

### `SetSubnetValidatorWeightMessage`

Description: Used to set a validator's stake weight on another chain

Signed by: Subnet, P-Chain

Specification:

```
+--------------+----------+----------+
|      codecID :   uint16 |  2 bytes |
+--------------+----------+----------+
|       typeID :   uint32 |  4 bytes |
+--------------+----------+----------+
| validationID : [32]byte | 32 bytes |
+--------------+----------+----------+
|        nonce :   uint64 |  8 bytes |
+--------------+----------+----------+
|       weight :   uint64 |  8 bytes |
+--------------+----------+----------+
                          | 54 bytes |
                          +----------+
```

## Types of Staking Managers

### Proof of Authority

The Proof of Authority (PoA) staking manager is ownable, and only allows the owner to modify changes to the validator set of the Subnet. Validators are given a weight by the owner, but do not accrue staking rewards.

### Proof of Stake

The Proof of Stake (PoS) staking manager allows any validator to register and exit the Subnet validator set. Validators are given a weight based on the amount of stake they provide, and accrue staking rewards based on their uptime. The staking rewards can be either native or erc20 tokens.