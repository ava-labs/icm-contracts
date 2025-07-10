// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";

import {IACP99Manager, PChainOwner, ConversionData} from "../interfaces/IACP99Manager.sol";

import {IValidatorManager, ValidatorManager, ValidatorStatus} from "../ValidatorManager.sol";

import {ValidatorMessages} from "../ValidatorMessages.sol";

import {ICMInitializable} from "../../utilities/ICMInitializable.sol";

import {SlotAuctionManager} from "../SlotAuctionManager.sol";

import {ExampleERC20} from "@mocks/ExampleERC20.sol";

import {IERC20Mintable} from "../interfaces/IERC20Mintable.sol";

import { WarpMessage, IWarpMessenger } from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";

abstract contract SlotAuctionManagerTest is ValidatorManagerTest {
    SlotAuctionManager public app;
    IERC20Mintable public token;
    
    uint16 public constant DEFAULT_VALIDATOR_SLOTS = uint16(2);
    uint256 public constant DEFAULT_MINIMUM_AUCTION_DURATION = uint256(10);
    uint256 public constant DEFAULT_MINIMUM_BID = uint256(10); 

    function setUp() public override {
        ValidatorManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();

        ConversionData memory conversion = _defaultConversionData();
        bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
        _mockInitializeValidatorSet(conversionID);
        validatorManager.initializeValidatorSet(conversion, 0);
    }
    
    function testStartAuction() public {
        vm.expectRevert(abi.encodeWithSelector(SlotAuctionManager.AuctionNotInProgress.selector));
        PChainOwner memory PCO;
        app.placeBid(100, DEFAULT_NODE_ID, DEFAULT_BLS_PUBLIC_KEY, PCO, PCO);
    }

    function _setUp() internal override returns (IACP99Manager) {
        // Construct the object under test
        token = new ExampleERC20();
        validatorManager = new ValidatorManager(ICMInitializable.Allowed);
        app = new SlotAuctionManager(address(token), address(validatorManager));

        validatorManager.initialize(_defaultSettings(address(app)));

        return validatorManager;
    }
    
}