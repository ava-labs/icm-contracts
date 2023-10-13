// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "@openzeppelin/contracts/utils/Address.sol";

/**
 * @dev Registry entry that represents a mapping between protocolAddress and version.
 */
struct ProtocolRegistryEntry {
    uint256 version;
    address protocolAddress;
}

/**
 * @dev Implementation of an abstract `WarpProtocolRegistry` contract.
 *
 * This implementation is a contract that can be used as a base contract for protocols that are
 * built on top of Warp. It allows the protocol to be upgraded through a Warp out-of-band message.
 */
abstract contract WarpProtocolRegistry {
    // Address that the out-of-band Warp message sets as the "source" address.
    // The address is obviously not owned by any EOA or smart contract account, so it
    // cannot possibly be the source address of any other Warp message emitted by the VM.
    address public constant VALIDATORS_SOURCE_ADDRESS = address(0);

    WarpMessenger public constant WARP_MESSENGER =
        WarpMessenger(0x0200000000000000000000000000000000000005);

    bytes32 internal immutable _blockchainID;

    // The latest protocol version. 0 means no protocol version has been added, and isn't a valid version.
    uint256 internal _latestVersion;

    mapping(uint256 => address) internal _versionToAddress;
    mapping(address => uint256) internal _addressToVersion;

    /**
     * @dev Emitted when a new protocol version is added to the registry.
     */
    event AddProtocolVersion(
        uint256 indexed version,
        address indexed protocolAddress
    );

    // Errors
    error InvalidWarpMessage();
    error InvalidSourceChainID();
    error InvalidOriginSenderAddress();
    error InvalidDestinationChainID();
    error InvalidDestinationAddress();
    error InvalidProtocolVersion();
    error InvalidProtocolAddress();

    /**
     * @dev Initializes the contract by setting a `chainID` and `latestVersion`.
     */
    constructor(ProtocolRegistryEntry[] memory initialEntries) {
        _latestVersion = 0;
        _blockchainID = WARP_MESSENGER.getBlockchainID();

        for (uint256 i = 0; i < initialEntries.length; i++) {
            _addToRegistry(initialEntries[i]);
        }
    }

    /**
     * @dev Gets and verifies a warp out-of-band message, and adds the new protocol version
     * address to the registry.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Requirements:
     *
     * - a valid warp out-of-band message must be provided.
     * - the version must be the increment of the latest version.
     * - the protocol address must be a contract address.
     */
    function addProtocolVersion(uint32 messageIndex) external virtual {
        _addProtocolVersion(messageIndex);
    }

    /**
     * @dev Gets the address of a protocol version.
     * Requirements:
     *
     * - the version must be a valid version.
     */
    function getAddressFromVersion(
        uint256 version
    ) external view virtual returns (address) {
        return _getAddressFromVersion(version);
    }

    /**
     * @dev Gets the version of the given `protocolAddress`.
     * If `protocolAddress` is not a valid protocol address, returns 0, which is an invalid version.
     */
    function getVersionFromAddress(
        address protocolAddress
    ) external view virtual returns (uint256) {
        return _addressToVersion[protocolAddress];
    }

    /**
     * @dev Gets the latest protocol version.
     */
    function getLatestVersion() external view virtual returns (uint256) {
        return _latestVersion;
    }

    /**
     * @dev Gets and verifies for a warp out-of-band message, and adds the new protocol version
     * address to the registry.
     */
    function _addProtocolVersion(uint32 messageIndex) internal virtual {
        // Get and validate for a warp out-of-band message.
        (WarpMessage memory message, bool valid) = WARP_MESSENGER
            .getVerifiedWarpMessage(messageIndex);
        if (!valid) {
            revert InvalidWarpMessage();
        }
        if (message.sourceChainID != _blockchainID) {
            revert InvalidSourceChainID();
        }

        // Check that the message is sent through a warp out of band message.
        if (message.originSenderAddress != VALIDATORS_SOURCE_ADDRESS) {
            revert InvalidOriginSenderAddress();
        }
        if (message.destinationChainID != _blockchainID) {
            revert InvalidDestinationChainID();
        }
        if (message.destinationAddress != address(this)) {
            revert InvalidDestinationAddress();
        }

        ProtocolRegistryEntry memory entry = abi.decode(
            message.payload,
            (ProtocolRegistryEntry)
        );

        _addToRegistry(entry);
    }

    /**
     * @dev Adds the new protocol version address to the registry.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Note: `protocolAddress` doesn't have to be a contract address, this is primarily
     * to support the case we want to register a new protocol address meant for a security patch
     * before the contract is deployed, to prevent the vulnerabilitiy from being exposed before registry update.
     * Requirements:
     *
     * - `version` must be the increment of the latest version.
     */
    function _addToRegistry(
        ProtocolRegistryEntry memory entry
    ) internal virtual {
        // Check that the version is the increment of the latest version.
        if (entry.version != _latestVersion + 1) {
            revert InvalidProtocolVersion();
        }

        if (entry.protocolAddress == address(0)) {
            revert InvalidProtocolAddress();
        }

        _latestVersion++;
        _versionToAddress[entry.version] = entry.protocolAddress;
        _addressToVersion[entry.protocolAddress] = entry.version;
        emit AddProtocolVersion(entry.version, entry.protocolAddress);
    }

    /**
     * @dev Gets the address of a protocol version.
     * Requirements:
     *
     * - `version` must be a valid version, i.e. greater than 0 and not greater than the latest version.
     */
    function _getAddressFromVersion(
        uint256 version
    ) internal view virtual returns (address) {
        // Check that the version provided is a valid version.
        if (version == 0) {
            revert InvalidProtocolVersion();
        }

        if (version > _latestVersion) {
            revert InvalidProtocolVersion();
        }
        return _versionToAddress[version];
    }
}
