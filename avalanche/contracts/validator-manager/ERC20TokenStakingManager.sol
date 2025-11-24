// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {StakingManager} from "./StakingManager.sol";
import {StakingManagerSettings} from "./interfaces/IStakingManager.sol";
import {PChainOwner} from "./ACP99Manager.sol";
import {IERC20TokenStakingManager} from "./interfaces/IERC20TokenStakingManager.sol";
import {IERC20Mintable} from "./interfaces/IERC20Mintable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

/**
 * @dev Implementation of the {IERC20TokenStakingManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract ERC20TokenStakingManager is Initializable, StakingManager, IERC20TokenStakingManager {
    using SafeERC20 for IERC20Mintable;
    using SafeERC20TransferFrom for IERC20Mintable;

    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.ERC20TokenStakingManager
    struct ERC20TokenStakingManagerStorage {
        IERC20Mintable _token;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.ERC20TokenStakingManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant ERC20_STAKING_MANAGER_STORAGE_LOCATION =
        0x6e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab00;

    // solhint-disable ordering
    function _getERC20StakingManagerStorage()
        private
        pure
        returns (ERC20TokenStakingManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := ERC20_STAKING_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(
        ICMInitializable init
    ) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initialize the ERC20 token staking manager
     * @param settings Initial settings for the PoS validator manager
     * @param token The ERC20 token to be staked
     */
    function initialize(
        StakingManagerSettings calldata settings,
        IERC20Mintable token
    ) external initializer {
        __ERC20TokenStakingManager_init(settings, token);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ERC20TokenStakingManager_init(
        StakingManagerSettings calldata settings,
        IERC20Mintable token
    ) internal onlyInitializing {
        __StakingManager_init(settings);
        __ERC20TokenStakingManager_init_unchained(token);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ERC20TokenStakingManager_init_unchained(
        IERC20Mintable token
    ) internal onlyInitializing {
        ERC20TokenStakingManagerStorage storage $ = _getERC20StakingManagerStorage();
        if (address(token) == address(0)) {
            revert ZeroAddress();
        }
        $._token = token;
    }

    /**
     * @notice See {IERC20TokenStakingManager-initiateValidatorRegistration}
     */
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount,
        address rewardRecipient
    ) external nonReentrant returns (bytes32) {
        return _initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration,
            stakeAmount: stakeAmount,
            rewardRecipient: rewardRecipient
        });
    }

    /**
     * @notice See {IERC20TokenStakingManager-initiateDelegatorRegistration}
     */
    function initiateDelegatorRegistration(
        bytes32 validationID,
        uint256 delegationAmount,
        address rewardRecipient
    ) external nonReentrant returns (bytes32) {
        return _initiateDelegatorRegistration(
            validationID, _msgSender(), delegationAmount, rewardRecipient
        );
    }

    /**
     * @notice Returns the ERC20 token being staked
     */
    function erc20() external view returns (IERC20Mintable) {
        return _getERC20StakingManagerStorage()._token;
    }

    /**
     * @notice See {StakingManager-_lock}
     * Note: Must be guarded with reentrancy guard for safe transfer from.
     */
    function _lock(
        uint256 value
    ) internal virtual override returns (uint256) {
        return _getERC20StakingManagerStorage()._token.safeTransferFrom(value);
    }

    /**
     * @notice See {StakingManager-_unlock}
     * Note: Must be guarded with reentrancy guard for safe transfer.
     */
    function _unlock(address to, uint256 value) internal virtual override {
        _getERC20StakingManagerStorage()._token.safeTransfer(to, value);
    }

    /**
     * @notice See {StakingManager-_reward}
     */
    function _reward(address account, uint256 amount) internal virtual override {
        ERC20TokenStakingManagerStorage storage $ = _getERC20StakingManagerStorage();
        $._token.mint(account, amount);
    }
}
