# Upgrading TeleporterMessenger

This document outlines the high-level steps necessary to upgrade `TeleporterMessenger` as an Avalanche L1 operator, a relayer operator, and a dApp admin. As a reference, the [Teleporter Registry test suite](../../../tests/flows/teleporter_registry.go) implements the steps described in a test environment.

## Register a New TeleporterMessenger Version

Once a new `TeleporterMessenger` contract instance is [deployed](../../../utils/contract-deployment/README.md), the `addProtocolVersion` method of [TeleporterRegistry.sol](./TeleporterRegistry.sol) must be called. This method is only callable if the associated ICM message was sent via an off-chain ICM message, which is provided by a validator's chain config.

The steps to do so are as follows:

1. Construct the unsigned ICM message bytes:

    a. Pack the ICM payload to be parsed by `addProtocolVersion`. This is a tuple of `(ProtocolRegistryEntry, address)`, where the `ProtocolRegistryEntry` specifies the new `TeleporterMessenger` contract address and the version, and the `address` specifies the `TeleporterRegistry` contract address that the new `TeleporterMessenger` version will be registered with.

    b. Pack the ICM payload as an [AddressedCall](https://github.com/ava-labs/avalanchego/blob/v1.11.3/vms/platformvm/warp/payload/addressed_call.go#L15), with the source address set to the zero address. This is how `addProtocolVersion` identifies this message as an off-chain ICM message.
    
    c. Pack the `AddressedCall` message into an [unsigned ICM message](https://github.com/ava-labs/avalanchego/blob/v1.11.3/vms/platformvm/warp/unsigned_message.go#L14), specifying the blockchain ID that `TeleporterRegistry` and the new `TeleporterMessenger` are deployed to.

2. Populate the "warp-off-chain-messages" field of each validator node's chain config with the hex-encoded unsigned ICM message bytes. 

3. Restart the node to mark the off-chain ICM message as eligible for signing by the validator.

To actually register the new `TeleporterMessenger` version with the registry, the validators must be queried for their signature of the message, the signatures aggregated, and a signed ICM message created to be included in the transaction that calls `addProtocolVersion`. As an example, [ICM Relayer](https://github.com/ava-labs/icm-services) provides this functionality. The following steps illustrate how to use ICM Relayer to register the new `TeleporterMessenger` version.

1. Construct a "manual-warp-messages" entry in the ICM Relayer configuration file, using the following values:

    a. "unsigned-message-bytes": the hex-encoded unsigned ICM message bytes derived above.

    b. "source-blockchain-id": the blockchain ID that that `TeleporterRegistry` and the new `TeleporterMessenger` are deployed to.

    c. "destination-blockchain-id": the blockchain ID that that `TeleporterRegistry` and the new `TeleporterMessenger` are deployed to.

    d. "source-address": the zero address, "0x0000000000000000000000000000000000000000". This is the "source" address for off-chain ICM messages.

2. Add the `TeleporterRegistry` as a supported message type by adding the following entry to the list of "message-contracts":

    a. Set the key to the zero address, "0x0000000000000000000000000000000000000000".
    
    b. Set the value as the following, populating the <TELEPORTER_REGISTRY_ADDRESS>:
    ```json
    {
        "message-format": "off-chain-registry",
        "settings": {
            "teleporter-registry-address": "<TELEPORTER_REGISTRY_ADDRESS>"
        }
    }
    ```

3. Restart the relayer. On startup, it will query the validator nodes for their BLS signatures on the off-chain ICM message, construct an aggregate signature and signed ICM message, and use it to call `addProtocolVersion` in the registry.
