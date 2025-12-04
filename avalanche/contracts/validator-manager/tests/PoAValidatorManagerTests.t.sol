// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";
import {IACP99Manager, PChainOwner, ConversionData} from "../interfaces/IACP99Manager.sol";
import {IValidatorManager, ValidatorManager} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";

contract PoAValidatorManagerTest is ValidatorManagerTest {
    address public constant DEFAULT_OWNER = address(0x1);

    ValidatorManager public app;

    function setUp() public override {
        ValidatorManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();

        ConversionData memory conversion = _defaultConversionData();
        bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
        _mockInitializeValidatorSet(conversionID);
        validatorManager.initializeValidatorSet(conversion, 0);
    }

    function testDisableInitialization() public {
        app = new ValidatorManager(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));
        app.initialize(_defaultSettings(address(this)));
    }

    function testInvalidOwnerRegistration() public {
        vm.prank(vm.addr(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, vm.addr(1)
            )
        );
        _initiateValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            weight: DEFAULT_WEIGHT
        });
    }

    // This test applies to all ValidatorManagers, but we test it here to avoid
    // having to source UINT64MAX funds for StakingManagers.
    function testTotalWeightOverflow() public {
        uint64 weight = type(uint64).max;

        bytes memory nodeID = _newNodeID();
        vm.expectRevert(
            abi.encodeWithSelector(IValidatorManager.InvalidTotalWeight.selector, weight)
        );

        _initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER,
            weight: weight
        });
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            weight: weight
        });
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight,
        address
    ) internal virtual override returns (bytes32) {
        return app.initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            weight: weight
        });
    }

    function _completeValidatorRegistration(
        uint32 messageIndex
    ) internal virtual override returns (bytes32) {
        return app.completeValidatorRegistration(messageIndex);
    }

    function _initiateValidatorRemoval(bytes32 validationID, bool) internal virtual override {
        return app.initiateValidatorRemoval(validationID);
    }

    function _forceInitiateValidatorRemoval(bytes32 validationID, bool) internal virtual override {
        return app.initiateValidatorRemoval(validationID);
    }

    function _completeValidatorRemoval(
        uint32 messageIndex
    ) internal virtual override returns (bytes32) {
        return app.completeValidatorRemoval(messageIndex);
    }

    function _setUp() internal override returns (IACP99Manager) {
        validatorManager = new ValidatorManager(ICMInitializable.Allowed);
        app = validatorManager;

        validatorManager.initialize(_defaultSettings(address(this)));

        return validatorManager;
    }

    // solhint-disable-next-line no-empty-blocks
    function _beforeSend(uint256 amount, address spender) internal virtual override {}

    function _beforeRegisterValidator(
        bytes32 validationID,
        address rewardRecipient
    ) internal virtual override 
    // solhint-disable-next-line no-empty-blocks
    {}
}
