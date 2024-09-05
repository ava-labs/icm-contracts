// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IERC20TokenStakingManager} from "./interfaces/IERC20TokenStakingManager.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {ICMInitializable} from "../utilities/ICMInitializable.sol";
import {PoSValidatorManager} from "./PoSValidatorManager.sol";
import {PoSValidatorManagerSettings} from "./interfaces/IPoSValidatorManager.sol";

contract ERC20TokenStakingManager is
    Initializable,
    PoSValidatorManager,
    IERC20TokenStakingManager
{
    using SafeERC20 for IERC20;
    using SafeERC20TransferFrom for IERC20;

    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.ERC20TokenStakingManager
    struct ERC20TokenStakingManagerStorage {
        IERC20 _token;
        uint8 _tokenDecimals;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.ERC20TokenStakingManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Update to correct storage slot
    bytes32 private constant _ERC20_STAKING_MANAGER_STORAGE_LOCATION =
        0x6e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab00;

    // solhint-disable ordering
    function _getERC20StakingManagerStorage()
        private
        pure
        returns (ERC20TokenStakingManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := _ERC20_STAKING_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initialize the ERC20 token staking manager
     * @dev Uses reinitializer(2) on the PoS staking contracts to make sure after migration from PoA, the PoS contracts can reinitialize with its needed values.
     * @param settings Initial settings for the PoS validator manager
     * @param token The ERC20 token to be staked
     */
    function initialize(
        PoSValidatorManagerSettings calldata settings,
        IERC20 token
    ) external reinitializer(2) {
        __ERC20TokenStakingManager_init(settings, token);
    }

    // solhint-disable func-name-mixedcase
    function __ERC20TokenStakingManager_init(
        PoSValidatorManagerSettings calldata settings,
        IERC20 token
    ) internal onlyInitializing {
        __POS_Validator_Manager_init(settings);
        __ERC20TokenStakingManager_init_unchained(token);
    }

    // solhint-disable func-name-mixedcase
    function __ERC20TokenStakingManager_init_unchained(IERC20 token) internal onlyInitializing {
        ERC20TokenStakingManagerStorage storage $ = _getERC20StakingManagerStorage();
        require(address(token) != address(0), "ERC20TokenStakingManager: zero token address");
        $._token = token;
    }

    /**
     * @notice Begins the validator registration process. Locks the configured ERC20 in the contract as the stake.
     * @param stakeAmount The amount to be staked.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The Unix timestamp after which the reigistration is no longer valid on the P-Chain.
     * @param delegationFeeRate The fee charged to delegators by the validator.
     * @param blsPublicKey The BLS public key of the validator.
     */
    function initializeValidatorRegistration(
        uint256 stakeAmount,
        bytes32 nodeID,
        uint64 registrationExpiry,
        uint256 delegationFeeRate,
        bytes memory blsPublicKey
    ) external returns (bytes32 validationID) {
        uint64 weight = _processStake(stakeAmount);
        validationID =
            _initializeValidatorRegistration(nodeID, weight, registrationExpiry, blsPublicKey);
        _setDelegationFeeRate(validationID, delegationFeeRate);
        return validationID;
    }

    /**
     * @notice Begins the delegator registration process. Locks the configured ERC20 in the contract as the delegated stake.
     * @param validationID The ID of the validation period being delegated to.
     * @param delegationAmount The amount to be delegated.
     */
    function initializeDelegatorRegistration(
        bytes32 validationID,
        uint256 delegationAmount
    ) external {
        return _initializeDelegatorRegistration(validationID, _msgSender(), delegationAmount);
    }

    // Must be guarded with reentrancy guard for safe transfer from
    function _lock(uint256 value) internal virtual override returns (uint256) {
        return _getERC20StakingManagerStorage()._token.safeTransferFrom(value);
    }

    function _unlock(uint256 value, address to) internal virtual override {
        _getERC20StakingManagerStorage()._token.safeTransfer(to, value);
    }
}
