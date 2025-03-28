# TeleporterMessenger Contracts Upgradability

## Overview

The `TeleporterMessenger` contract is non-upgradable, once a version of the contract is deployed it cannot be changed. This is with the intention of preventing any changes to the deployed contract that could potentially introduce bugs or vulnerabilities.

However, there could still be new versions of `TeleporterMessenger` contracts needed to be deployed in the future. `TeleporterRegistry` provides applications that use a `TeleporterMessenger` instance a minimal step process to integrate with new versions of `TeleporterMessenger`.

The `TeleporterRegistry` maintains a mapping of `TeleporterMessenger` contract versions to their addresses. When a new `TeleporterMessenger` version is deployed, its address can be added to the `TeleporterRegistry`. The `TeleporterRegistry` can only be updated through an ICM off-chain message that meets the following requirements:

- `sourceChainAddress` must match `VALIDATORS_SOURCE_ADDRESS = address(0)`
  - The zero address can only be set as the source chain address by an ICM off-chain message, and cannot be set by an on-chain ICM message.
- `sourceBlockchainID` must match the blockchain ID that the registry is deployed on
- `destinationBlockchainID` must match the blockchain ID that the registry is deployed on
- `destinationAddress` must match the address of the registry

In the `TeleporterRegistry` contract, the `latestVersion` state variable returns the highest version number that has been registered in the registry. The `getLatestTeleporter` function returns the `ITeleporterMessenger` that is registered with the corresponding version.

## Design

- `TeleporterRegistry` is deployed on each blockchain that needs to keep track of `TeleporterMessenger` contract versions.
- The registry's contract address on each blockchain does not need to be the same, and does not require a Nick's method transaction for deployment.
- Each registry's mapping of version to contract address is independent of registries on other blockchains, and chains can decide on their own registry mapping entries.
- Each blockchain should only have one canonical `TeleporterRegistry` contract.
- `TeleporterRegistry` contract can be initialized through a list of initial registry entries, which are `TeleporterMessenger` contract versions and their addresses.
- The registry keeps track of a mapping of `TeleporterMessenger` contract versions to their addresses, and vice versa, a mapping of `TeleporterMessenger` contract addresses to their versions.
- Version zero is an invalid version, and is used to indicate that a `TeleporterMessenger` contract has not been registered yet.
- Once a version number is registered in the registry, it cannot be changed, but a previous registered protocol address can be added to the registry with a new version. This is especially important in the case of a rollback to a previous `TeleporterMessenger` version, in which case the previous `TeleporterMessenger` contract address would need to be registered with a new version to the registry.

## Integrating `TeleporterRegistryApp` into a dApp

<div align="center">
  <img src="./upgrade-uml.png?raw=true" alt="Upgrade UML diagram"/>
</div>

[TeleporterRegistryApp](./TeleporterRegistryApp.sol) is an abstract contract that helps integrate the `TeleporterRegistry` into ICM contracts. To support upgradeable contracts, there is also a corresponding `TeleporterRegistryAppUpgradeable` contract that is upgrade compatible. By inheriting from `TeleporterRegistryApp`, dApps get:

- Ability to send ICM messages through the latest version of the `TeleporterMessenger` contract registered in the Teleporter registry. (The dApp can also override this to use a specific version of the `TeleporterMessenger` contract.)
- `minTeleporterVersion` management that allows the dApp to specify the minimum `TeleporterMessenger` version that can send messages to the dApp.
- Access controlled utility to update the `minTeleporterVersion`
- Access controlled utility to pause/unpause interaction with specific `TeleporterMessenger` addresses.

To integrate `TeleporterRegistryApp` with a dApp, pass in the Teleporter registry address inside the constructor. For upgradeable contracts `TeleporterRegistryAppUpgradeable` can be inherited, and the derived contract's `initializer` function should call either `__TeleporterRegistryApp_init` or `__TeleporterRegistryApp_init_unchained` An example dApp looks like:

```solidity
// An example dApp that integrates with the Teleporter registry
// to send/receive ICM messages.
contract ExampleApp is
    TeleporterRegistryApp
{
    ...
    // Constructor passes in the Teleporter registry address
    // to the TeleporterRegistryApp contract.
    constructor(
        address teleporterRegistryAddress,
        uint256 minTeleporterVersion
    ) TeleporterRegistryApp(teleporterRegistryAddress, minTeleporterVersion) {
        currentBlockchainID = IWarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();
    }
    ...
    // Handles receiving ICM messages,
    // and also checks that the sender is a valid TeleporterMessenger contract.
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        // implementation
    }

    // Implements the access control checks for the dApp's interaction with TeleporterMessenger versions.
    function _checkTeleporterRegistryAppAccess() internal view virtual override {
        //implementation
    }

}
```

### Checking TeleporterRegistryApp access

To prevent anyone from calling the dApp's `updateMinTeleporterVersion`, which would disallow messages from old `TeleporterMessenger` versions from being received, this function should be safeguarded with access controls. All contracts deriving from `TeleporterRegistryApp` will need to implement `TeleporterRegistryApp._checkTeleporterRegistryAppAccess`. For example, [TeleporterRegistryOwnableApp](./TeleporterRegistryOwnableApp.sol) is an abstract contract that inherits `TeleporterRegistryApp`, and implements `_checkTeleporterRegistryAppAccess` to check whether the caller is the owner. There is also a corresponding `TeleporterRegistryOwnableAppUpgradeable` contract that is upgrade compatible.

```solidity
    function _checkTeleporterRegistryAppAccess() internal view virtual override {
        _checkOwner();
    }
```

Another example would be a dApp that has different roles and priveleges. `_checkTeleporterRegistryAppAccess` can be implemented to check whether the caller is a specific role.

```solidity
    function _checkTeleporterRegistryAppAccess() internal view virtual override {
        require(
            hasRole(TELEPORTER_REGISTRY_APP_ADMIN, _msgSender()),
            "TeleporterRegistryApp: caller does not have access"
        );
    }
```

### Sending with specific TeleporterMessenger version

For sending messages with the Teleporter registry, dApps should use `TeleporterRegistryApp._getTeleporterMessenger`. This function by default extends `TeleporterRegistry.getLatestTeleporter`, using the latest version, and adds an extra check on whether the latest `TeleporterMessenger` address is paused. If the dApp wants to send a message through a specific `TeleporterMessenger` version, it can override `_getTeleporterMessenger()` to use the specific `TeleporterMessenger` version with `TeleporterRegistry.getTeleporterFromVersion`.

The `TeleporterRegistryApp._sendTeleporterMessage` function makes sending ICM messages easier. The function uses `_getTeleporterMessenger` to get the sending `TeleporterMessenger` version, pays for `TeleporterMessenger` fees from the dApp's balance, and sends the cross chain message.

Using latest version:

```solidity
        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();
```

Using specific version:

```solidity
        // Override _getTeleporterMessenger to use specific version.
        function _getTeleporterMessenger() internal view override returns (ITeleporterMessenger) {
            ITeleporterMessenger teleporter = teleporterRegistry
                .getTeleporterFromVersion($VERSION);
            require(
                !pausedTeleporterAddresses[address(teleporter)],
                "TeleporterRegistryApp: Teleporter sending version paused"
            );

            return teleporter;
        }

        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();
```

### Receiving from specific TeleporterMessenger versions

`TeleporterRegistryApp` also provides an initial implementation of [ITeleporterReceiver.receiveTeleporterMessage](../ITeleporterReceiver.sol) that ensures `_msgSender` is a `TeleporterMessenger` contract with a version greater than or equal to `minTeleporterVersion`. This supports the case where a dApp wants to use a new version of the `TeleporterMessenger` contract, but still wants to be able to receive messages from the old `TeleporterMessenger` contract.The dApp can override `_receiveTeleporterMessage` to implement its own logic for receiving messages from `TeleporterMessenger` contracts.

## Managing a TeleporterRegistryApp dApp

dApps that implement `TeleporterRegistryApp` automatically use the latest `TeleporterMessenger` version registered with the `TeleporterRegistry`. Interaction with underlying `TeleporterMessenger` versions can be managed by setting the minimum `TeleporterMessenger` version, and pausing and unpausing specific versions.

The following sections include example `cast send` commands for issuing transactions that call contract functions. See the [Foundry Book](https://book.getfoundry.sh/reference/cast/cast-send) for details on how to issue transactions using common wallet options.

### Managing the Minimum TeleporterMessenger version

The `TeleporterRegistryApp` contract constructor saves the Teleporter registry in a state variable used by the inheriting dApp contract, and initializes a `minTeleporterVersion` to the highest `TeleporterMessenger` version registered in `TeleporterRegistry`. `minTeleporterVersion` is used to allow dApp's to specify the `TeleporterMessenger` versions allowed to interact with it.

#### Updating `minTeleporterVersion`

The `TeleporterRegistryApp.updateMinTeleporterVersion` function updates the `minTeleporterVersion` used to check which `TeleporterMessenger` versions can be used for sending and receiving messages. **Once the `minTeleporterVersion` is increased, any undelivered messages sent by other chains using older versions of `TeleporterMessenger` will never be able to be received**. The `updateMinTeleporterVersion` function can only be called with a version greater than the current `minTeleporterVersion` and less than `latestVersion` in the Teleporter registry.

> Example: Update the minimum TeleporterMessenger version to 2
>
> ```bash
> cast send <DAPP_ADDRESS> "updateMinTeleporterVersion(uint256)" 2
> ```

### Pausing TeleporterMessenger version interactions

dApps that inherit from `TeleporterRegistryApp` can pause `TeleporterMessenger` interactions by calling `TeleporterRegistryApp.pauseTeleporterAddress`. This function prevents the dApp contract from interacting with the paused `TeleporterMessenger` address when sending or receiving ICM messages.

`pauseTeleporterAddress` can only be called by addresses that passes the dApp's `TeleporterRegistryApp._checkTeleporterRegistryAppAccess` check.

The `TeleporterMessenger` address corresponding to a `TeleporterMessenger` version can be fetched from the registry with `TeleporterRegistry.getAddressFromVersion`

> Example: Pause TeleporterMessenger version 3
>
> ```bash
> versionThreeAddress=$(cast call <REGISTRY_ADDRESS> "getAddressFromVersion(uint256)(address)" 3)
> cast send <DAPP_ADDRESS> "pauseTeleporterAddress(address)" $versionThreeAddress
> ```

#### Pause all TeleporterMessenger interactions

To pause all `TeleporterMessenger` interactions, `TeleporterRegistryApp.pauseTeleporterAddress` must be called for every `TeleporterMessenger` version from the `minTeleporterVersion` to the latest `TeleporterMessenger` version registered in `TeleporterRegistry`. Note that there may be gaps in `TeleporterMessenger` versions registered with `TeleporterRegistry`, but they will always be in increasing order. The latest `TeleporterMessenger` version can be obtained by inspecting the public variable `TeleporterRegistry.latestVersion`. The `minTeleporterVersion` can be obtained by calling `TeleporterRegistryApp.getMinTeleporterVersion`.

> Example: Pause all registered TeleporterMessenger versions
>
> ```bash
> # Fetch the minimum TeleporterMessenger version
> minVersion=$(cast call <DAPP_ADDRESS> "getMinTeleporterVersion()(uint256)")
>
> # Fetch the latest registered version
> latestVersion=$(cast call <REGISTRY_ADDRESS> "latestVersion()(uint256)")
>
> # Pause all registered versions
> for ((version=minVersion; version<=latestVersion; version++))
> do
>  # Fetch the version address if it's registered
>  versionAddress=$(cast call <REGISTRY_ADDRESS> "getAddressFromVersion(uint256)(address)" $version)
>
>  if [ $? -eq 0 ]; then
>    # If cast call is successful, proceed to cast send
>    cast send <DAPP_ADDRESS> "pauseTeleporterAddress(address)" $versionAddress
>  else
>    # If cast call fails, print an error message and skip to the next iteration
>    echo "Version $version not registered. Skipping."
>  fi
> done
> ```

#### Unpausing TeleporterMessenger version interactions

As with pausing, dApps can unpause `TeleporterMessenger` interactions by calling `TeleporterRegistryApp.unpauseTeleporterAddress`. This unpause function allows receiving `TeleporterMessenger` message from the unpaused `TeleporterMessenger` address, and also enables the sending of messages through the unpaused `TeleporterMessenger` address in `_getTeleporterMessenger()`. Unpausing is also only allowed by addresses passing the dApp's `_checkTeleporterRegistryAppAccess` check.

Note that receiving `TeleporterMessenger` messages is still governed by the `minTeleporterVersion` check, so even if a `TeleporterMessenger` address is unpaused, the dApp will not receive messages from the unpaused `TeleporterMessenger` address if the `TeleporterMessenger` version is less than `minTeleporterVersion`.

> Example: Unpause TeleporterMessenger version 3
>
> ```bash
> versionThreeAddress=$(cast call <REGISTRY_ADDRESS> "getAddressFromVersion(uint256)(address)" 3)
> cast send <DAPP_ADDRESS> "unpauseTeleporterAddress(address)" $versionThreeAddress
> ```
