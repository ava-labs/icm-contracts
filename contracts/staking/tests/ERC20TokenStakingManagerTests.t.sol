// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoSValidatorManagerTest} from "./PoSValidatorManagerTests.t.sol";
import {ERC20TokenStakingManager} from "../ERC20TokenStakingManager.sol";
import {ValidatorManagerSettings} from "../interfaces/IValidatorManager.sol";
import {PoSValidatorManagerSettings} from "../interfaces/IPoSValidatorManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
contract ERC20TokenStakingManagerTest is PoSValidatorManagerTest {
    using SafeERC20 for IERC20;

    ERC20TokenStakingManager public app;
    IERC20 public token;

    function setUp() public virtual {
        // Construct the object under test
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        token = new ExampleERC20();
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFeeRate: DEFAULT_MINIMUM_DELEGATION_FEE_RATE,
                rewardCalculator: IRewardCalculator(address(0))
            }),
            token
        );
        validatorManager = app;
        posValidatorManager = app;
    }

    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        uint256 delegationFeeRate,
        bytes memory signature,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration(
            app.weightToValue(weight), nodeID, registrationExpiry, delegationFeeRate, signature
        );
    }

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegator,
        uint64 weight
    ) internal virtual override {
        uint256 value = app.weightToValue(weight);
        vm.startPrank(delegator);
        app.initializeDelegatorRegistration(validationID, value);
        vm.stopPrank();
    }

    function _beforeSend(uint64 weight, address spender) internal override {
        uint256 value = app.weightToValue(weight);
        token.safeIncreaseAllowance(spender, value);
        token.safeTransfer(spender, value);

        // ERC20 tokens need to be pre-approved
        vm.startPrank(spender);
        token.safeIncreaseAllowance(address(app), value);
        vm.stopPrank();
    }
}
