// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {ICMMessage, ICMUnsignedMessage, ICMSignature, ICM} from "../utils/ICM.sol";
import {Validator, ValidatorSet} from "../utils/ValidatorSets.sol";
import {BLST} from "../utils/BLST.sol";

contract ICMUtilsTest is Test {
    function testParseICMMessage() public pure {
        bytes memory messageBytes =
            hex"0000000000013d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a70000000d48656c6c6f2c20576f726c64210000000000000001010cb7f52fa291c273810bcf0dd1d3dc41e0449e6da4ff2b90243ce33b6ff4733cbd3f8446878109e369db1ae9a67622430d3d6ec6a9b1abb9e6cd86df330080a546d2b8599950618eb76efbdfe13ef74fc14e4fa9919bd440dd2bd02bf84421af089653aef1a209f9bbf2837562b02adc6e766753e0247417225a4de4ad095ad50d619b719e234702df7f87e8cd54873310c55d7516d586441d7dd978e199f08383013aa6e3b2f9a48ac3366e9a5df7a003786b93b937d9ccf463a6f10b342306";
        ICMMessage memory message = ICM.parseICMMessage(messageBytes);
        assertEq(message.unsignedMessage.avalancheNetworkID, 1);
        assertEq(
            message.unsignedMessage.avalancheSourceBlockchainID,
            bytes32(hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7")
        );
        assertEq(message.unsignedMessage.payload, hex"48656c6c6f2c20576f726c6421");
        assertEq(
            message.unsignedMessageBytes,
            hex"0000000000013d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a70000000d48656c6c6f2c20576f726c6421"
        );
        assertEq(message.signature.signers, hex"01");
        assertEq(
            message.signature.signature,
            hex"0cb7f52fa291c273810bcf0dd1d3dc41e0449e6da4ff2b90243ce33b6ff4733cbd3f8446878109e369db1ae9a67622430d3d6ec6a9b1abb9e6cd86df330080a546d2b8599950618eb76efbdfe13ef74fc14e4fa9919bd440dd2bd02bf84421af089653aef1a209f9bbf2837562b02adc6e766753e0247417225a4de4ad095ad50d619b719e234702df7f87e8cd54873310c55d7516d586441d7dd978e199f08383013aa6e3b2f9a48ac3366e9a5df7a003786b93b937d9ccf463a6f10b342306"
        );
    }

    function testBytesToBoolArray() public pure {
        bytes memory data = hex"01";
        bool[] memory result = ICM.bytesToBoolArray(data);
        assertEq(result.length, 1);
        assertEq(result[0], true);

        data = hex"80";
        result = ICM.bytesToBoolArray(data);
        assertEq(result.length, 8);
        assertEq(result[0], true);
        assertEq(result[1], false);
        assertEq(result[2], false);
        assertEq(result[3], false);
        assertEq(result[4], false);
        assertEq(result[5], false);
        assertEq(result[6], false);
        assertEq(result[7], false);
    }

    function testVerifyWeight() public pure {
        assertEq(ICM.verifyWeight(100, 100), true);
        assertEq(ICM.verifyWeight(68, 100), true);
        assertEq(ICM.verifyWeight(67, 100), true);
        assertEq(ICM.verifyWeight(3936137349139582, 5874831864387435), true);
        assertEq(ICM.verifyWeight(3936137349139581, 5874831864387435), false);
        assertEq(ICM.verifyWeight(67, 100), true);
        assertEq(ICM.verifyWeight(67000000000, 100000000000), true);
        assertEq(ICM.verifyWeight(66666666669, 100000000000), false);
        assertEq(ICM.verifyWeight(66, 100), false);
        assertEq(ICM.verifyWeight(0, 100), false);
    }

    function testVerifyICMMessageSingleSignerSuccess() public view {
        bytes memory unformattedPublicKey =
            hex"074a926af04042ab5876a9ba4297a806a34732d7df7a33b2d34d8e77313c34c49f365eba61e923acfbdfa160c3e51d2710b93b3ff2351cc019178b142ecb2eed17b0c99c2b6e9c30bedae8cfc26f8f7cdf75b9f0dbb2ea31e11117a04dbe962f";
        bytes memory formattedPublicKey = BLST.formatUncompressedBLSPublicKey(unformattedPublicKey);

        Validator[] memory validators = new Validator[](1);
        validators[0] = Validator({blsPublicKey: formattedPublicKey, weight: 100});

        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: bytes32(
                hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7"
            ),
            validators: validators,
            totalWeight: 100,
            pChainHeight: 1,
            pChainTimestamp: 1
        });
        ICMMessage memory message = ICM.parseICMMessage(
            hex"0000000000013d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a70000000d48656c6c6f2c20576f726c64210000000000000001010cb7f52fa291c273810bcf0dd1d3dc41e0449e6da4ff2b90243ce33b6ff4733cbd3f8446878109e369db1ae9a67622430d3d6ec6a9b1abb9e6cd86df330080a546d2b8599950618eb76efbdfe13ef74fc14e4fa9919bd440dd2bd02bf84421af089653aef1a209f9bbf2837562b02adc6e766753e0247417225a4de4ad095ad50d619b719e234702df7f87e8cd54873310c55d7516d586441d7dd978e199f08383013aa6e3b2f9a48ac3366e9a5df7a003786b93b937d9ccf463a6f10b342306"
        );
        assertEq(ICM.verifyICMMessage(message, 1, validatorSet), true);
    }

    function testVerifyICMMessageMultiSignerSuccess() public view {
        bytes[] memory unformattedPublicKeys = new bytes[](4);
        unformattedPublicKeys[0] =
            hex"0984b5daad5db34822a5b73103587ea28209087d3ad92fba376f669692e07771f9945661a2d6fc9114ea098015ef44a0107993eb8136dfbb698c9b36736b093000e7c6d704c131c9534e45bb11069a6ce296ce15878e8cbd39800c52773c8841";
        unformattedPublicKeys[1] =
            hex"10ffae60e1c05689695b551f442351306be88aa4f98faa7521a120e12b5ba597bfa310c4ef9c59e31622bf678241ea9a0cebd5284f2326f38f0515652d16471bc699b9654870e49219b091f2717caf07b170e40bb23aad749a99dfd9cc072982";
        unformattedPublicKeys[2] =
            hex"0a8c197bfda5978a8fb1ee76b36b544029594a32191916f46568a18b31594e53aebea9cb0a0c218d9f2a7ddf63d7cc68196c71b6c4c816e3d3adc8189e9e89d677160272eebc84070185da00ac4b97ed3a0b3b18aab35381239882c8a3250072";
        unformattedPublicKeys[3] =
            hex"0252a35315cec61341b676c42ac45c058f23e76f49194349f8c80d5e1414e6043560ebe4d129f23bff1f3692f4052396175832741f48604400a1b43f76e43f738f196472058158573c1e9d14de62b9137d2e4548507fd48b8f28a5ad9baa1063";

        Validator[] memory validators = new Validator[](4);
        for (uint256 i = 0; i < 4; ++i) {
            validators[i] = Validator({
                blsPublicKey: BLST.formatUncompressedBLSPublicKey(unformattedPublicKeys[i]),
                weight: 100
            });
        }

        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: bytes32(
                hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7"
            ),
            validators: validators,
            totalWeight: 400,
            pChainHeight: 1,
            pChainTimestamp: 1
        });

        ICMMessage memory message = ICM.parseICMMessage(
            hex"0000000000013d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a70000000d48656c6c6f2c20576f726c642100000000000000010b0a09516a15348d6d981453e57d7cfc8c9f4d965b96773cefa8edf77c763d549507a2ba420e4a717011b3d9c1f4f1eb3f158970399057617f161b6e9dcfd82682291c21bbe1d7600334c1fe70efa58b03f2eeacea089bd9ec940d64f98d6b177e086ba348a850a39b4ee758645b0c273857133d448f310be2d3d7ad16c2c377d4d9156a6f18e47437e405ed7f042cec7605c18a3fac0ae6569fa482a3210d6a50b913d9e607494f239f178d11ea9904bd3598c87970a8b69d955dc6fe0178e8b7"
        );
        assertEq(ICM.verifyICMMessage(message, 1, validatorSet), true);
    }
}
