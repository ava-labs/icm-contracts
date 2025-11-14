// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import "forge-std/Test.sol";
import "../AvalancheValidatorSetRegistry.sol";
import "../utils/ValidatorSets.sol";
import "../utils/ICM.sol";
import "../utils/BLST.sol";

contract AvalancheValidatorSetRegistryTest is Test {
    AvalancheValidatorSetRegistry registry;

    uint32 constant NETWORK_ID = 1;

    function setUp() public {
        registry = new AvalancheValidatorSetRegistry(NETWORK_ID, bytes32(hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7"));
    }

    function testGetAvalancheNetworkID() public view {
        assertEq(registry.getAvalancheNetworkID(), NETWORK_ID);
    }

    function testRegisterValidatorSet() public {
        bytes memory validatorBytes =
            hex"000000000004055b1e5892c401dc04a72699da740f768f25f16fbac3debaed3caa7f502242b9b8b5890c6795d30bd6011451aef8b547133e7251b1d022c842645433528f94d9e118d2b3a7bd2ec8e8feb971e7d2fe7f5ba81f056147a61f4b4e5e747529ac1d0000000000000064064bb2e2313f21f36907e4c010a7bfee0216a3f0000bf2b1c83adf6aa6675bada4b68dd023f4f4c9285be2c75ff745ac0b790412ba453f7d1a0c49c9c7e1fd7a791110499ea762a96d57b69e67d2e25fb5756a37e26f4c4a69b27939799b92f9000000000000006413d331edb3c1f4113b5148e25bf32658f962d43dea4115fe40fb0c38ac3acd58537cc147218e34e86fd5366ed1ef0d00173010db4bbb9e64cb676d39f4bae156b7b8bc76102d6e41e7bc46c6b6a369d21726f9e7be4b929ce4dcf84fa1a1a56600000000000000641831d57287d398005355881a59ee05384c9467c7974650cb6a38fd1346d3ec58484ab0ad97678bf21521c6e6824e08180d16fb7a9257aac66442415c3f51f35b691f5ce4389ce22c0701b2a1ac6d05f2907e040eb7a03b5aef68bbd90cfbea770000000000000064";

        bytes memory signedValidtorSetStateMessage =
            hex"0000000000013d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a70000006400000000000100000000000000560000000000043d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a700000000000000640000000068712555234da730fe88131b61174151e4a00b7f47b024818550869e9d11a75e2cf4fa4900000000000000010b0525ddc62d270271ab07307867f19b7b70b83a11231b463980fc359d2622a99e2c160f29461422e0455d5279480d77fc08f8d71020f434478e25aa2a1a7c1e468add5a0e113c9544d846335965215f404f59b6e198f6c02659747803fe9502f40bb6e436071f743d6453ecfc97e30d840119bb1003fc3e71df8fa5e822687a9cf7021e40fa75c7756bfe5079c2e8896e1454653389759028106f34e2c1810e08caa6e2c88edbb9bdecb8914c2201c9ae4c82c3c06236d0df0a886068fcfa3f81";
        bytes32 expectedBlockchainID =
            bytes32(hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7");
        uint32 expectedPChainHeight = 100;
        uint32 expectedPChainTimestamp = 1752245589;
        bytes[] memory expectedPublicKeys = new bytes[](4);
        expectedPublicKeys[0] =
            hex"055b1e5892c401dc04a72699da740f768f25f16fbac3debaed3caa7f502242b9b8b5890c6795d30bd6011451aef8b547133e7251b1d022c842645433528f94d9e118d2b3a7bd2ec8e8feb971e7d2fe7f5ba81f056147a61f4b4e5e747529ac1d";
        expectedPublicKeys[1] =
            hex"064bb2e2313f21f36907e4c010a7bfee0216a3f0000bf2b1c83adf6aa6675bada4b68dd023f4f4c9285be2c75ff745ac0b790412ba453f7d1a0c49c9c7e1fd7a791110499ea762a96d57b69e67d2e25fb5756a37e26f4c4a69b27939799b92f9";
        expectedPublicKeys[2] =
            hex"13d331edb3c1f4113b5148e25bf32658f962d43dea4115fe40fb0c38ac3acd58537cc147218e34e86fd5366ed1ef0d00173010db4bbb9e64cb676d39f4bae156b7b8bc76102d6e41e7bc46c6b6a369d21726f9e7be4b929ce4dcf84fa1a1a566";
        expectedPublicKeys[3] =
            hex"1831d57287d398005355881a59ee05384c9467c7974650cb6a38fd1346d3ec58484ab0ad97678bf21521c6e6824e08180d16fb7a9257aac66442415c3f51f35b691f5ce4389ce22c0701b2a1ac6d05f2907e040eb7a03b5aef68bbd90cfbea77";
        uint64 expectedWeight = 100;
        uint64 expectedTotalWeight = 400;

        // Parse the ICM message.
        ICMMessage memory message = ICM.parseICMMessage(signedValidtorSetStateMessage);

        // Get the current validator set ID.
        uint256 expectedValidatorSetID = registry.nextValidatorSetID();

        // Register the validator set.
        uint256 validatorSetID = registry.registerValidatorSet(message, validatorBytes);

        // Check that the validator set ID is correct.
        assertEq(validatorSetID, expectedValidatorSetID);

        // Check that the validator set is correct.
        ValidatorSet memory validatorSet = registry.getValidatorSet(validatorSetID);
        assertEq(validatorSet.avalancheBlockchainID, expectedBlockchainID);
        assertEq(validatorSet.pChainHeight, expectedPChainHeight);
        assertEq(validatorSet.pChainTimestamp, expectedPChainTimestamp);
        assertEq(validatorSet.validators.length, 4);
        for (uint256 i = 0; i < 4; i++) {
            assertEq(
                validatorSet.validators[i].blsPublicKey,
                BLST.formatUncompressedBLSPublicKey(expectedPublicKeys[i])
            );
            assertEq(validatorSet.validators[i].weight, expectedWeight);
        }
        assertEq(validatorSet.totalWeight, expectedTotalWeight);
    }

    function testIncorrectNetworkId() public {
        // create a new registry with a network id different than the one the message is intended for
        registry = new AvalancheValidatorSetRegistry(NETWORK_ID + 1, bytes32(hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7"));
        bytes memory validatorBytes =
            hex"000000000004055b1e5892c401dc04a72699da740f768f25f16fbac3debaed3caa7f502242b9b8b5890c6795d30bd6011451aef8b547133e7251b1d022c842645433528f94d9e118d2b3a7bd2ec8e8feb971e7d2fe7f5ba81f056147a61f4b4e5e747529ac1d0000000000000064064bb2e2313f21f36907e4c010a7bfee0216a3f0000bf2b1c83adf6aa6675bada4b68dd023f4f4c9285be2c75ff745ac0b790412ba453f7d1a0c49c9c7e1fd7a791110499ea762a96d57b69e67d2e25fb5756a37e26f4c4a69b27939799b92f9000000000000006413d331edb3c1f4113b5148e25bf32658f962d43dea4115fe40fb0c38ac3acd58537cc147218e34e86fd5366ed1ef0d00173010db4bbb9e64cb676d39f4bae156b7b8bc76102d6e41e7bc46c6b6a369d21726f9e7be4b929ce4dcf84fa1a1a56600000000000000641831d57287d398005355881a59ee05384c9467c7974650cb6a38fd1346d3ec58484ab0ad97678bf21521c6e6824e08180d16fb7a9257aac66442415c3f51f35b691f5ce4389ce22c0701b2a1ac6d05f2907e040eb7a03b5aef68bbd90cfbea770000000000000064";
        bytes memory signedValidtorSetStateMessage =
            hex"0000000000013d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a70000006400000000000100000000000000560000000000043d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a700000000000000640000000068712555234da730fe88131b61174151e4a00b7f47b024818550869e9d11a75e2cf4fa4900000000000000010b0525ddc62d270271ab07307867f19b7b70b83a11231b463980fc359d2622a99e2c160f29461422e0455d5279480d77fc08f8d71020f434478e25aa2a1a7c1e468add5a0e113c9544d846335965215f404f59b6e198f6c02659747803fe9502f40bb6e436071f743d6453ecfc97e30d840119bb1003fc3e71df8fa5e822687a9cf7021e40fa75c7756bfe5079c2e8896e1454653389759028106f34e2c1810e08caa6e2c88edbb9bdecb8914c2201c9ae4c82c3c06236d0df0a886068fcfa3f81";
        // Parse the ICM message.
        ICMMessage memory message = ICM.parseICMMessage(signedValidtorSetStateMessage);
        vm.expectRevert();
        registry.registerValidatorSet(message, validatorBytes);
    }
}
