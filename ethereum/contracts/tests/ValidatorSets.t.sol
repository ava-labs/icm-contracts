// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import "forge-std/Test.sol";
import "../utils/ValidatorSets.sol";
import "../utils/ICM.sol";
import "../utils/BLST.sol";

contract ValidatorSetTest is Test {
    function testParseValidatorsRoundtrip() public view {
        Validator[] memory validators = new Validator[](3);
        validators[0] = Validator({
            blsPublicKey: BLST.getPublicKeyFromSecret(1),
            weight: 1
        });
        validators[1] = Validator({
            blsPublicKey: BLST.getPublicKeyFromSecret(2),
            weight: 2
        });
        validators[2] = Validator({
            blsPublicKey: BLST.getPublicKeyFromSecret(3),
            weight: 3
        });
        bytes memory serialized = ValidatorSets.serializeValidators(validators);
        (Validator[] memory deserialized, uint64 totalWeight) = ValidatorSets.parseValidators(serialized);
        for (uint256 i = 0; i < 3; i++) {
            assertEq(validators[i].blsPublicKey, deserialized[i].blsPublicKey);
            assertEq(validators[i].weight, deserialized[i].weight);
        }
        assertEq(deserialized.length, 3);
        assertEq(totalWeight, 6);
    }

    function testPayloadRoundtrip() public pure {
        ValidatorSetStatePayload memory payload = ValidatorSetStatePayload({
            avalancheBlockchainID: hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7",
            pChainHeight: 22,
            pChainTimestamp: 32,
            validatorSetHash: hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7"
        });
        bytes memory serialized = ValidatorSets.serializeValidatorSetStatePayload(payload);
        ValidatorSetStatePayload memory deserialized = ValidatorSets.parseValidatorSetStatePayload(serialized);
        assertEq(payload.avalancheBlockchainID, deserialized.avalancheBlockchainID);
        assertEq(payload.pChainHeight, deserialized.pChainHeight);
        assertEq(payload.pChainTimestamp, deserialized.pChainTimestamp);
        assertEq(payload.validatorSetHash, deserialized.validatorSetHash);
    }
}