// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {
    IACP99Manager,
    PChainOwner,
    ConversionData,
    InitialValidator
} from "../interfaces/IACP99Manager.sol";
import {
    IValidatorManager,
    ValidatorManager,
    ValidatorStatus,
    ValidatorManagerSettings
} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {NativeTokenSlotAuctionManager} from "../NativeTokenSlotAuctionManager.sol";
import {SlotAuctionManager} from "../SlotAuctionManager.sol";
import {
    ISlotAuctionManager, ValidatorBid, ValidatorInfo, SlotAuctionManagerSettings, AuctionSettings
} from "../interfaces/ISlotAuctionManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";
import {SlotAuctionManagerTest} from "./SlotAuctionManagerTests.t.sol";

contract NativeTokenSlotAuctionManagerTest is SlotAuctionManagerTest {
    NativeTokenSlotAuctionManager public app;

    function setUp() public override {
        SlotAuctionManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();

        ConversionData memory conversion = _defaultConversionData();
        bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
        _mockInitializeValidatorSet(conversionID);
        validatorManager.initializeValidatorSet(conversion, 0);
    }

    function _setUp() internal override {
        // Construct the object under test
        validatorManager = new ValidatorManager(ICMInitializable.Allowed);
        app = new NativeTokenSlotAuctionManager(
            ICMInitializable.Allowed
        );

        app.initialize(
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
        //Native tokes dont need to be pre approved
    }

    function _placeBid(
        uint256 bid,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) internal override {
        app.placeBid{value: bid}(nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
    }
}
