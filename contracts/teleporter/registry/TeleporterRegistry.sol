// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";
import {
    IWarpMessenger,
    WarpMessage
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

// Registry entry that represents a mapping between protocolAddress and Teleporter version.
struct ProtocolRegistryEntry {
    uint256 version;
    address protocolAddress;
}

/**
 * @dev TeleporterRegistry contract provides an upgrade mechanism for {ITeleporterMessenger} contracts
 * through Warp off-chain messages
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract TeleporterRegistry {
    /**
     * @notice Address that the off-chain Warp message sets as the "source" address.
     * @dev The address is not owned by any EOA or smart contract account, so it
     * cannot possibly be the source address of any other Warp message emitted by the VM.
     */
    address public constant VALIDATORS_SOURCE_ADDRESS = address(0);

    /**
     * @notice Warp precompile used for sending and receiving Warp messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    /**
     * @notice The blockchain ID of the chain the contract is deployed on.
     */
    bytes32 public immutable blockchainID;

    /**
     * @notice The maximum version increment allowed when adding a new protocol version.
     */
    uint256 public constant MAX_VERSION_INCREMENT = 500;

    /**
     * @notice The latest protocol version.
     * @dev 0 means no protocol version has been added, and isn't a valid version.
     */
    uint256 public latestVersion;

    /**
     * @notice Mappings that keep track of the protocol version and corresponding contract address.
     */
    mapping(uint256 version => address protocolAddress) private _versionToAddress;
    mapping(address protocolAddress => uint256 version) private _addressToVersion;

    /**
     * @notice Emitted when a new protocol version is added to the registry.
     */
    event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress);

    /**
     * @notice Emitted when the latest version is updated.
     */
    event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion);

    /**
     * @dev Initializes the contract by setting `blockchainID` and `latestVersion`.
     * Also adds the initial protocol versions to the registry.
     */
    constructor(
        ProtocolRegistryEntry[] memory initialEntries
    ) {
        blockchainID = WARP_MESSENGER.getBlockchainID();

        uint256 length = initialEntries.length;
        for (uint256 i; i < length; ++i) {
            _addToRegistry(initialEntries[i]);
        }
    }

    /**
     * @dev Receives a Warp off-chain message containing a new protocol version and address to be registered,
     * and adds the new values to the respective mappings.
     * If a version is greater than the current latest version, it will be set as the latest version.
     * If a version is less than the current latest version, it is added to the registry, but
     * doesn't change the latest version.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Emits a {LatestVersionUpdated} event when a new protocol version greater than the current latest version is added.
     * Requirements:
     *
     * - a valid Warp off-chain message must be provided.
     * - source chain ID must be the same as the blockchain ID of the registry.
     * - origin sender address must be the same as the `VALIDATORS_SOURCE_ADDRESS`.
     * - destination address must be the same as the address of this registry.
     */
    function addProtocolVersion(
        uint32 messageIndex
    ) external {
        // Get the verified Warp message, and check that it was sent to this registry via a Warp off-chain message.
        (WarpMessage memory message, bool success) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(success, "TeleporterRegistry: invalid warp message");
        require(
            message.sourceChainID == blockchainID, "TeleporterRegistry: invalid source chain ID"
        );
        // Check that the message is sent through a Warp off-chain message.
        require(
            message.originSenderAddress == VALIDATORS_SOURCE_ADDRESS,
            "TeleporterRegistry: invalid origin sender address"
        );

        (ProtocolRegistryEntry memory entry, address destinationAddress) =
            abi.decode(message.payload, (ProtocolRegistryEntry, address));

        // Check that the message is sent to the registry.
        require(
            destinationAddress == address(this), "TeleporterRegistry: invalid destination address"
        );

        _addToRegistry(entry);
    }

    /**
     * @dev Gets the latest {ITeleporterMessenger} contract.
     */
    function getLatestTeleporter() external view returns (ITeleporterMessenger) {
        return ITeleporterMessenger(getAddressFromVersion(latestVersion));
    }

    /**
     * @dev Gets the {ITeleporterMessenger} contract of the given `version`.
     * Requirements:
     *
     * - `version` must be a valid registered version.
     */
    function getTeleporterFromVersion(
        uint256 version
    ) external view returns (ITeleporterMessenger) {
        return ITeleporterMessenger(getAddressFromVersion(version));
    }

    /**
     * @dev Gets the address of a protocol version.
     * Requirements:
     *
     * - `version` must be a valid registered version.
     */
    function getAddressFromVersion(
        uint256 version
    ) public view returns (address) {
        require(version != 0, "TeleporterRegistry: zero version");
        address protocolAddress = _versionToAddress[version];
        require(protocolAddress != address(0), "TeleporterRegistry: version not found");
        return protocolAddress;
    }

    /**
     * @dev Gets the version of the given `protocolAddress`.
     * Requirements:
     *
     * - `protocolAddress` must be a valid registered protocol address.
     */
    function getVersionFromAddress(
        address protocolAddress
    ) public view returns (uint256) {
        require(protocolAddress != address(0), "TeleporterRegistry: zero protocol address");
        uint256 version = _addressToVersion[protocolAddress];
        require(version != 0, "TeleporterRegistry: protocol address not found");
        return version;
    }

    /**
     * @dev Adds the new protocol version address to the registry.
     * Updates latest version if the version is greater than the current latest version.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Emits a {LatestVersionUpdated} event when a new protocol version greater than the current latest version is added.
     * Note: `entry.protocolAddress` doesn't have to be a contract address, allowing a new protocol address to be registered before the contract is deployed.
     * Requirements:
     *
     * - `entry.version` is not zero
     * - `entry.version` is not already registered
     * - `entry.protocolAddress` is not zero address
     */
    function _addToRegistry(
        ProtocolRegistryEntry memory entry
    ) private {
        require(entry.version != 0, "TeleporterRegistry: zero version");
        // Check that the version has not previously been registered.
        require(
            _versionToAddress[entry.version] == address(0),
            "TeleporterRegistry: version already exists"
        );
        require(entry.protocolAddress != address(0), "TeleporterRegistry: zero protocol address");

        uint256 latestVersion_ = latestVersion;
        require(
            entry.version <= latestVersion_ + MAX_VERSION_INCREMENT,
            "TeleporterRegistry: version increment too high"
        );

        _versionToAddress[entry.version] = entry.protocolAddress;

        // Since a protocol address can be registered multiple times,
        // only update the version if the new version is greater than the current version.
        if (entry.version > _addressToVersion[entry.protocolAddress]) {
            _addressToVersion[entry.protocolAddress] = entry.version;
        }
        emit AddProtocolVersion(entry.version, entry.protocolAddress);

        // Set latest version if the version is greater than the current latest version.
        if (entry.version > latestVersion_) {
            latestVersion = entry.version;
            emit LatestVersionUpdated(latestVersion_, entry.version);
        }
    }
}
