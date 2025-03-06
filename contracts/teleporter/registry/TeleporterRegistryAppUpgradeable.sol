// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TeleporterRegistry} from "./TeleporterRegistry.sol";
import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";
import {ITeleporterMessenger, TeleporterMessageInput} from "@teleporter/ITeleporterMessenger.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

/**
 * @dev TeleporterRegistryAppUpgradeable provides upgrade utility for applications built on top
 * of the Teleporter protocol by integrating with the {TeleporterRegistry}.
 *
 * This contract is intended to be inherited by other contracts that wish to use the
 * upgrade mechanism. It provides an interface that restricts access to only Teleporter
 * versions that are greater than or equal to `minTeleporterVersion`.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract TeleporterRegistryAppUpgradeable is
    Initializable,
    ContextUpgradeable,
    ITeleporterReceiver,
    ReentrancyGuardUpgradeable
{
    using SafeERC20 for IERC20;

    // solhint-disable private-vars-leading-underscore
    /**
     * @dev Namespace storage slots following the ERC-7201 standard to prevent
     * storage collisions between upgradeable contracts.
     * @custom:storage-location erc7201:teleporter.storage.TeleporterRegistryApp
     */
    struct TeleporterRegistryAppStorage {
        // The Teleporter registry contract manages different Teleporter contract versions.
        TeleporterRegistry _teleporterRegistry;
        /**
         * @dev A mapping that keeps track of paused Teleporter addresses.
         */
        mapping(address teleporterAddress => bool paused) _pausedTeleporterAddresses;
        /**
         * @dev The minimum required Teleporter version that the contract is allowed
         * to receive messages from. Should only be updated by `_setMinTeleporterVersion`
         */
        uint256 _minTeleporterVersion;
    }
    // solhint-enable private-vars-leading-underscore

    /**
     * @dev Storage slot computed based off ERC-7201 formula
     * keccak256(abi.encode(uint256(keccak256("teleporter.storage.TeleporterRegistryApp")) - 1)) & ~bytes32(uint256(0xff));
     */
    bytes32 public constant TELEPORTER_REGISTRY_APP_STORAGE_LOCATION =
        0xde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00;

    /**
     * @dev Emitted when `minTeleporterVersion` is updated.
     */
    event MinTeleporterVersionUpdated(
        uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion
    );

    /**
     * @dev Emitted when a Teleporter address is paused.
     */
    event TeleporterAddressPaused(address indexed teleporterAddress);

    /**
     * @dev Emitted when a Teleporter address is unpaused.
     */
    event TeleporterAddressUnpaused(address indexed teleporterAddress);

    function _getTeleporterRegistryAppStorage()
        private
        pure
        returns (TeleporterRegistryAppStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := TELEPORTER_REGISTRY_APP_STORAGE_LOCATION
        }
    }

    /**
     * @dev Initializes the {TeleporterRegistryApp} contract by getting `teleporterRegistry`
     * instance and setting `_minTeleporterVersion`.
     * @param teleporterRegistryAddress The address of the Teleporter registry contract.
     * The Teleporter registry contract should be deployed, and have at least one Teleporter version.
     * @param minTeleporterVersion The minimum Teleporter version allowed for delivering messages.
     * The minimum version should be less than or equal to the latest Teleporter version, and greater than 0.
     */
    // solhint-disable ordering
    // solhint-disable-next-line func-name-mixedcase
    function __TeleporterRegistryApp_init(
        address teleporterRegistryAddress,
        uint256 minTeleporterVersion
    ) internal onlyInitializing {
        __ReentrancyGuard_init();
        __Context_init();
        __TeleporterRegistryApp_init_unchained(teleporterRegistryAddress, minTeleporterVersion);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __TeleporterRegistryApp_init_unchained(
        address teleporterRegistryAddress,
        uint256 minTeleporterVersion
    ) internal onlyInitializing {
        require(
            teleporterRegistryAddress != address(0),
            "TeleporterRegistryApp: zero Teleporter registry address"
        );

        TeleporterRegistryAppStorage storage $ = _getTeleporterRegistryAppStorage();
        TeleporterRegistry registry = TeleporterRegistry(teleporterRegistryAddress);
        require(registry.latestVersion() > 0, "TeleporterRegistryApp: invalid Teleporter registry");
        $._teleporterRegistry = registry;
        _setMinTeleporterVersion(minTeleporterVersion);
    }
    // solhint-enable ordering

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}
     * `nonReentrant` is a reentrancy guard that protects against multiple versions of the
     * TeleporterMessengerContract delivering a message in the same call. Any internal calls
     * will not be able to call functions also marked with `nonReentrant`.
     *
     * Requirements:
     *
     * - `_msgSender()` must be a Teleporter version greater than or equal to `minTeleporterVersion`.
     */
    function receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes calldata message
    ) external nonReentrant {
        TeleporterRegistryAppStorage storage $ = _getTeleporterRegistryAppStorage();
        // Checks that `_msgSender()` matches a Teleporter version greater than or equal to `minTeleporterVersion`.
        require(
            $._teleporterRegistry.getVersionFromAddress(_msgSender()) >= $._minTeleporterVersion,
            "TeleporterRegistryApp: invalid Teleporter sender"
        );

        // Check against the paused Teleporter addresses.
        require(
            !_isTeleporterAddressPaused($, _msgSender()),
            "TeleporterRegistryApp: Teleporter address paused"
        );

        _receiveTeleporterMessage(sourceBlockchainID, originSenderAddress, message);
    }

    /**
     * @dev Checks if a Teleporter address is paused.
     */
    function isTeleporterAddressPaused(
        address teleporterAddress
    ) external view virtual returns (bool) {
        TeleporterRegistryAppStorage storage $ = _getTeleporterRegistryAppStorage();
        return _isTeleporterAddressPaused($, teleporterAddress);
    }

    /**
     * @dev Updates the minimum Teleporter version allowed for delivering Teleporer messages
     * to this contract.
     *
     * To prevent anyone from being able to call this function, which would disallow messages
     * from old Teleporter versions from being received, this function should be safeguarded with access
     * controls. This is done by overriding the implementation of {_checkTeleporterRegistryAppAccess}.
     */
    function updateMinTeleporterVersion(
        uint256 version
    ) public virtual {
        _checkTeleporterRegistryAppAccess();
        _setMinTeleporterVersion(version);
    }

    /**
     * @dev Pauses a Teleporter address from interacting with this contract.
     * After pausing a Teleporter address, it will no longer be able to deliver messages
     * to this contract, and this contract will not send messages through that Teleporter address.
     * The address does not need to be registered with the Teleporter registry.
     * Emits a {TeleporterAddressPaused} event if successfully paused.
     * Requirements:
     *
     * - `_msgSender()` must have Teleporter upgrade access.
     * - `teleporterAddress` is not the zero address.
     * - `teleporterAddress` is not already paused.
     */
    function pauseTeleporterAddress(
        address teleporterAddress
    ) public virtual {
        TeleporterRegistryAppStorage storage $ = _getTeleporterRegistryAppStorage();
        _checkTeleporterRegistryAppAccess();
        require(teleporterAddress != address(0), "TeleporterRegistryApp: zero Teleporter address");
        require(
            !_isTeleporterAddressPaused($, teleporterAddress),
            "TeleporterRegistryApp: address already paused"
        );
        $._pausedTeleporterAddresses[teleporterAddress] = true;
        emit TeleporterAddressPaused(teleporterAddress);
    }

    /**
     * @dev Unpauses a Teleporter address from interacting with this contract.
     * After unpausing a Teleporter address, it will again be able to deliver messages
     * to this contract, and this contract can send messages through that Teleporter address.
     * The address does not need to be registered with the Teleporter registry.
     * Emits a {TeleporterAddressUnpaused} event if successfully unpaused.
     * Requirements:
     *
     * - `_msgSender()` must have Teleporter upgrade access.
     * - `teleporterAddress` is not the zero address.
     * - `teleporterAddress` is already paused.
     */
    function unpauseTeleporterAddress(
        address teleporterAddress
    ) public virtual {
        TeleporterRegistryAppStorage storage $ = _getTeleporterRegistryAppStorage();
        _checkTeleporterRegistryAppAccess();
        require(teleporterAddress != address(0), "TeleporterRegistryApp: zero Teleporter address");
        require(
            _isTeleporterAddressPaused($, teleporterAddress),
            "TeleporterRegistryApp: address not paused"
        );
        $._pausedTeleporterAddresses[teleporterAddress] = false;
        emit TeleporterAddressUnpaused(teleporterAddress);
    }

    /**
     * @dev Public getter for `_minTeleporterVersion`.
     */
    function getMinTeleporterVersion() public view returns (uint256) {
        TeleporterRegistryAppStorage storage $ = _getTeleporterRegistryAppStorage();
        return $._minTeleporterVersion;
    }

    /**
     * @dev Sets the minimum Teleporter version allowed for delivering Teleporter messages.
     * Emits a {MinTeleporterVersionUpdated} event if the minimum Teleporter version was updated.
     * Requirements:
     *
     * - `version` must be less than or equal to the latest Teleporter version.
     * - `version` must be greater than the current minimum Teleporter version.
     *
     */
    function _setMinTeleporterVersion(
        uint256 version
    ) internal virtual {
        TeleporterRegistryAppStorage storage $ = _getTeleporterRegistryAppStorage();
        uint256 latestTeleporterVersion = $._teleporterRegistry.latestVersion();
        uint256 oldMinTeleporterVersion = $._minTeleporterVersion;

        require(
            version <= latestTeleporterVersion, "TeleporterRegistryApp: invalid Teleporter version"
        );
        require(
            version > oldMinTeleporterVersion,
            "TeleporterRegistryApp: not greater than current minimum version"
        );

        $._minTeleporterVersion = version;
        emit MinTeleporterVersionUpdated(oldMinTeleporterVersion, version);
    }

    /**
     * @dev Receives Teleporter messages and handles accordingly.
     * This function should be overridden by contracts that inherit from this contract.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal virtual;

    /**
     * @dev Checks that the caller has access to update the minimum Teleporter version
     * allowed for delivering Teleporter messages to this contract.
     *
     * This function should be overridden by contracts that inherit from this contract.
     */
    function _checkTeleporterRegistryAppAccess() internal virtual;

    /**
     * @dev Sends a cross chain message using the TeleporterMessenger contract.
     *
     * The fee amount should be transferred to this contract prior to calling this function.
     * The fee amount is then allocated from this contract's token balance to
     * Teleporter's allowance to pay for sending the message.
     *
     * @return `messageID` The unique identifier for the Teleporter message.
     */
    function _sendTeleporterMessage(
        TeleporterMessageInput memory messageInput
    ) internal virtual returns (bytes32) {
        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();
        // For non-zero fee amounts increase the Teleporter contract's allowance.
        if (messageInput.feeInfo.amount > 0) {
            require(
                messageInput.feeInfo.feeTokenAddress != address(0),
                "TeleporterRegistryApp: zero fee token address"
            );
            IERC20(messageInput.feeInfo.feeTokenAddress).safeIncreaseAllowance(
                address(teleporterMessenger), messageInput.feeInfo.amount
            );
        }
        return teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    /**
     * @dev Returns the Teleporter messenger used to send Teleporter messages,
     * and checks that the Teleporter messenger is not paused.
     *
     * By default returns the latest Teleporter messenger, but can be overriden to
     * return a Teleporter messenger of a specific version.
     */
    function _getTeleporterMessenger() internal view virtual returns (ITeleporterMessenger) {
        TeleporterRegistryAppStorage storage $ = _getTeleporterRegistryAppStorage();
        ITeleporterMessenger teleporter = $._teleporterRegistry.getLatestTeleporter();
        require(
            !_isTeleporterAddressPaused($, address(teleporter)),
            "TeleporterRegistryApp: Teleporter sending paused"
        );

        return teleporter;
    }

    function _isTeleporterAddressPaused(
        TeleporterRegistryAppStorage storage $,
        address teleporterAddress
    ) internal view returns (bool) {
        return $._pausedTeleporterAddresses[teleporterAddress];
    }
}
