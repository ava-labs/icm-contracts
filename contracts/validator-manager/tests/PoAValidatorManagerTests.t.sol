// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoAValidatorManager} from "../PoAValidatorManager.sol";
import {
    ValidatorManagerSettings,
    ValidatorRegistrationInput,
    IValidatorManager
} from "../interfaces/IValidatorManager.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";

contract PoAValidatorManagerTest is ValidatorManagerTest {
    PoAValidatorManager public app;

    address public constant DEFAULT_OWNER = address(0x1);

    function setUp() public override {
        ValidatorManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();
        _mockInitializeValidatorSet();
        app.initializeValidatorSet(_defaultConversionData(), 0);
    }

    function testDisableInitialization() public {
        app = new PoAValidatorManager(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));
        app.initialize(
            ValidatorManagerSettings({
                subnetID: DEFAULT_SUBNET_ID,
                churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
            }),
            address(this)
        );
    }

    function testInvalidOwnerRegistration() public {
        vm.prank(vm.addr(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, vm.addr(1)
            )
        );
        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: DEFAULT_NODE_ID,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                registrationExpiry: DEFAULT_EXPIRY,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER
            }),
            DEFAULT_WEIGHT
        );
    }

    // This test applies to all ValidatorManagers, but we test it here to avoid
    // having to source UINT64MAX funds for PoSValidatorManagers.
    function testTotalWeightOverflow() public {
        uint64 weight = type(uint64).max;

        bytes memory nodeID = _newNodeID();
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorManager.InvalidTotalWeight.selector, weight)
        );

        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: nodeID,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                registrationExpiry: DEFAULT_EXPIRY
            }),
            weight
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration(input, weight);
    }

    function _initializeEndValidation(
        bytes32 validationID,
        bool,
        address
    ) internal virtual override {
        return app.initializeEndValidation(validationID);
    }

    function _forceInitializeEndValidation(
        bytes32 validationID,
        bool,
        address
    ) internal virtual override {
        return app.initializeEndValidation(validationID);
    }

    function _setUp() internal override returns (IValidatorManager) {
        app = new PoAValidatorManager(ICMInitializable.Allowed);
        app.initialize(
            ValidatorManagerSettings({
                subnetID: DEFAULT_SUBNET_ID,
                churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
            }),
            address(this)
        );
        validatorManager = app;

        return app;
    }

    // solhint-disable-next-line no-empty-blocks
    function _beforeSend(uint256 amount, address spender) internal virtual override {}
}
