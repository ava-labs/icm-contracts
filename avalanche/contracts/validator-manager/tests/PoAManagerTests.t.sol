// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {PoAManager} from "../PoAManager.sol";
import {ValidatorManager, ValidatorManagerSettings} from "../ValidatorManager.sol";
import {IValidatorManagerExternalOwnable} from "../interfaces/IValidatorManagerExternalOwnable.sol";
import {IACP99Manager, PChainOwner, ConversionData} from "../interfaces/IACP99Manager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

contract PoAManagerTest is ValidatorManagerTest {
    PoAManager public poaManager;

    function setUp() public override {
        ValidatorManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();

        ConversionData memory conversion = _defaultConversionData();
        bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
        _mockInitializeValidatorSet(conversionID);
        validatorManager.initializeValidatorSet(conversion, 0);
    }

    function testOnlyOwnerCanInitiateValidatorRegistration() public {
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

    function testAnyoneCanCompleteValidatorRegistration() public {
        bytes32 validationID = _setUpInitiateValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_BLS_PUBLIC_KEY,
            address(this)
        );
        bytes memory l1ValidatorRegistrationMessage =
            ValidatorMessages.packL1ValidatorRegistrationMessage(validationID, true);

        _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit CompletedValidatorRegistration(validationID, DEFAULT_WEIGHT);

        vm.prank(vm.addr(1));
        _completeValidatorRegistration(0);
    }

    function testOnlyOwnerCanInitiateValidatorRemoval() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.prank(vm.addr(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, vm.addr(1)
            )
        );
        _initiateValidatorRemoval(validationID, false);
    }

    function testAnyoneCanCompleteValidatorRemoval() public {
        bytes32 validationID = _registerDefaultValidator();
        _initiateValidatorRemoval(validationID, false);

        bytes memory l1ValidatorRegistrationMessage =
            ValidatorMessages.packL1ValidatorRegistrationMessage(validationID, false);

        _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit CompletedValidatorRemoval(validationID);

        vm.prank(vm.addr(1));
        _completeValidatorRemoval(0);
    }

    function testOnlyOwnerCanInitiateValidatorWeightUpdate() public {
        vm.prank(vm.addr(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, vm.addr(1)
            )
        );
        poaManager.initiateValidatorWeightUpdate(bytes32(0), 0);
    }

    function testAnyoneCanCompleteValidatorWeightUpdate() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 newWeight = DEFAULT_WEIGHT + 1;

        poaManager.initiateValidatorWeightUpdate(validationID, newWeight);

        bytes memory l1ValidatorWeightUpdateMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 0, newWeight);

        _mockGetPChainWarpMessage(l1ValidatorWeightUpdateMessage, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit CompletedValidatorWeightUpdate(validationID, 0, newWeight);

        vm.prank(vm.addr(1));
        poaManager.completeValidatorWeightUpdate(0);
    }

    function testPoAMangerOnlyOwnerCanTransferOwnership() public {
        address newOwner = vm.addr(1);
        poaManager.transferValidatorManagerOwnership(newOwner);

        assertEq(validatorManager.owner(), newOwner, "Ownership should be transferred");
    }

    function testPoAMangerFailedTransferOwnership() public {
        address newOwner = vm.addr(1);
        vm.prank(newOwner);
        vm.expectRevert(
            abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, newOwner)
        );
        poaManager.transferValidatorManagerOwnership(newOwner);
    }

    function testPoAManagerIsOwnerOfValidatorManager() public view {
        assertEq(validatorManager.owner(), address(poaManager), "PoAManager should be owner");
    }

    function _setUp() internal override returns (IACP99Manager) {
        validatorManager = new ValidatorManager(ICMInitializable.Allowed);
        poaManager = new PoAManager(
            address(this), IValidatorManagerExternalOwnable(address(validatorManager))
        );

        // Construct ValidatorManagerSettings with the correct fields
        ValidatorManagerSettings memory settings = _defaultSettings(address(poaManager));

        validatorManager.initialize(settings);
        return validatorManager;
    }

    // Implement required abstract functions
    // solhint-disable-next-line no-empty-blocks
    function _beforeSend(uint256 amount, address spender) internal virtual override {}

    function _beforeRegisterValidator(
        bytes32 validationID,
        address rewardRecipient
    ) internal virtual override 
    // solhint-disable-next-line no-empty-blocks
    {}

    // Override helper functions to call PoAManager instead of ValidatorManager
    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight,
        address /*rewardRecipient*/
    ) internal virtual override returns (bytes32) {
        return poaManager.initiateValidatorRegistration(
            nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, weight
        );
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return poaManager.initiateValidatorRegistration(
            nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, weight
        );
    }

    function _completeValidatorRegistration(
        uint32 messageIndex
    ) internal virtual override returns (bytes32) {
        return poaManager.completeValidatorRegistration(messageIndex);
    }

    function _initiateValidatorRemoval(
        bytes32 validationID,
        bool /*includeUptime*/
    ) internal virtual override {
        poaManager.initiateValidatorRemoval(validationID);
    }

    function _forceInitiateValidatorRemoval(
        bytes32 validationID,
        bool /*includeUptime*/
    ) internal virtual override {
        poaManager.initiateValidatorRemoval(validationID);
    }

    function _completeValidatorRemoval(
        uint32 messageIndex
    ) internal virtual override returns (bytes32) {
        return poaManager.completeValidatorRemoval(messageIndex);
    }
}
