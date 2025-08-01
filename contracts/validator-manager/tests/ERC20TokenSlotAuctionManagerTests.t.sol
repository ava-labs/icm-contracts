// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner, ConversionData} from "../interfaces/IACP99Manager.sol";
import {IValidatorManager, ValidatorManager} from "../ValidatorManager.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {ERC20TokenSlotAuctionManager} from "../ERC20TokenSlotAuctionManager.sol";
import {SlotAuctionManagerSettings} from "../interfaces/ISlotAuctionManager.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC20Mintable} from "../interfaces/IERC20Mintable.sol";
import {SlotAuctionManagerTest} from "./SlotAuctionManagerTests.t.sol";

contract ERC20TokenSlotAuctionManagerTest is SlotAuctionManagerTest {
    using SafeERC20 for IERC20Mintable;

    ERC20TokenSlotAuctionManager public app;
    IERC20Mintable public token;

    function setUp() public override {
        SlotAuctionManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();

        ConversionData memory conversion = _defaultConversionData();
        bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
        _mockInitializeValidatorSet(conversionID);
        validatorManager.initializeValidatorSet(conversion, 0);
    }

    function testInsufficientBalance() public {
        _startAuctionWithCheck();
        _beforeBid(DEFAULT_MINIMUM_BID - 1, address(this));
        vm.expectRevert(); // Not sure what the name of the error is but somewhere in safe transfer should revert
        _placeBid(
            DEFAULT_MINIMUM_BID,
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
    }

    function _setUp() internal override {
        // Construct the object under test
        token = new ExampleERC20();
        validatorManager = new ValidatorManager(ICMInitializable.Allowed);
        app = new ERC20TokenSlotAuctionManager(ICMInitializable.Allowed);
        app.initialize(
            token,
            SlotAuctionManagerSettings(
                DEFAULT_AUCTION_OWNER_ADRESS,
                IValidatorManager(address(validatorManager)),
                DEFAULT_AUCTION_SETTINGS
            )
        );

        validatorManager.initialize(_defaultSettings(address(app)));

        slotauctionmanager = app;
    }

    function _beforeBid(uint256 amount, address spender) internal override {
        token.safeIncreaseAllowance(spender, amount);
        token.safeTransfer(spender, amount);

        // ERC20 tokens need to be pre-approved
        vm.startPrank(spender);
        token.safeIncreaseAllowance(address(app), amount);
        vm.stopPrank();
    }

    function _placeBid(
        uint256 bid,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) internal override {
        app.placeBid(bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
    }
}
