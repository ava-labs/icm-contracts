# Governance Contracts

## Overview

This contract provides an alternative to traditional multi-signature contracts where instead of requiring signatures from `K` of `N` pre-specified signers, an aggregate signature is required from a quorum of the current validators for a given blockchain.

The contract leverages off-chain [Avalanche ICM Messages](https://docs.avax.network/build/cross-chain/awm/overview), which are manually approved for signing by a chain's validators. It requires these messages to have the source address set to the zero address to enforce this, since on-chain ICM messages cannot have zero source address.

Note:
1. The blockchain validating the message may or may not be the same chain where the target contract and the `ValidatorSetSig` contract are deployed.
2. [Off-Chain ICM messages](https://github.com/ava-labs/subnet-evm/issues/729) are ICM messages that validators can include in their config to indicate that they are willing to sign them even though they are not a result of on-chain activity.

#### Creating a valid Off-chain ICM Message for interaction with the ValidatorSetSig Contract

1. ABI-encode a `ValidatorSetSigMessage` as defined in `ValidatorSetSig.sol`
2. OPTIONAL: call `validateMessage` view method of the intended contract to confirm that the message has correct fields, including the nonce.
3. Pack that as the payload of an [AddressedCall](https://github.com/ava-labs/avalanchego/blob/0c4efd743e1d737f4e8970d0e0ebf229ea44406c/vms/platformvm/warp/payload/addressed_call.go#L15) ICM Message format. Note that the `SourceAddress` field has to be set to the zero address.
4. Pack the `AddressedCall` as the payload of the [UnsignedWarpMessage](https://github.com/ava-labs/avalanchego/blob/f17ea6a7ab4036c41b693e47b94d8f0c81cb69ec/vms/platformvm/warp/unsigned_message.go#L14).
