// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../utils/BLSTUtils.sol";

contract BLSTUtilsTest is Test {
    function testFormatUncompressedBLSPublicKey() public pure {
        bytes memory unformatted =
            hex"074a926af04042ab5876a9ba4297a806a34732d7df7a33b2d34d8e77313c34c49f365eba61e923acfbdfa160c3e51d2710b93b3ff2351cc019178b142ecb2eed17b0c99c2b6e9c30bedae8cfc26f8f7cdf75b9f0dbb2ea31e11117a04dbe962f";
        bytes memory formatted =
            hex"00000000000000000000000000000000074a926af04042ab5876a9ba4297a806a34732d7df7a33b2d34d8e77313c34c49f365eba61e923acfbdfa160c3e51d270000000000000000000000000000000010b93b3ff2351cc019178b142ecb2eed17b0c99c2b6e9c30bedae8cfc26f8f7cdf75b9f0dbb2ea31e11117a04dbe962f";

        assertEq(formatted, BLSTUtils.formatUncompressedBLSPublicKey(unformatted));
    }

    function testNoAggregatePublicKeysSuccess() public view {
        // Test data from bls_formatter.go output
        bytes memory pk1 =
            hex"00000000000000000000000000000000105fdda3f4e22ea318d3831af1de2ef487f4b0506dfe61dca9cbd5c042484b3ecad5565855a557c237081456979433ea0000000000000000000000000000000013711c822adbbcb4f5811d48576f14c5ab78a9e5313636bf5a3722ad60ecdb96629f7e7d76873291960f6d08ab312869";

        bytes[] memory publicKeys = new bytes[](1);
        publicKeys[0] = pk1;
        bytes memory aggregated = BLSTUtils.aggregatePublicKeys(publicKeys);
        assertEq(aggregated, pk1);
    }

    function testAggregatePublicKeysSuccess() public view {
        // Test data from bls_formatter.go output
        bytes memory pk1 =
            hex"00000000000000000000000000000000105fdda3f4e22ea318d3831af1de2ef487f4b0506dfe61dca9cbd5c042484b3ecad5565855a557c237081456979433ea0000000000000000000000000000000013711c822adbbcb4f5811d48576f14c5ab78a9e5313636bf5a3722ad60ecdb96629f7e7d76873291960f6d08ab312869";
        bytes memory pk2 =
            hex"000000000000000000000000000000000f872fdc2f552676f3297330d13d888f9921038cce6cc087f1087d315363c79847e43f85d836a796cf6ddf101d13e2d9000000000000000000000000000000000b1737099061d418482275164132e3cea936ebb202661cac3df9c2dfba028fd164a947b0f62b38cae6dba6784ba475d8";
        bytes memory expectedAggregated =
            hex"00000000000000000000000000000000086401cb5b1276627747ca377b2de206fd0b6427bac06fdbd21e4cf5c71b57676367e5cc198c54c8ab0e4d0e5fa05663000000000000000000000000000000001436d02c536eb6e244867b806bc9453b367e923edcf1ada35d743122c5d2503d4114d8c776b64e67d22f3c5c6e4583e7";

        bytes[] memory publicKeys = new bytes[](2);
        publicKeys[0] = pk1;
        publicKeys[1] = pk2;

        bytes memory aggregated = BLSTUtils.aggregatePublicKeys(publicKeys);
        assertEq(aggregated, expectedAggregated);
    }

    function testAddG1Success() public view {
        bytes memory pk1 =
            hex"0984b5daad5db34822a5b73103587ea28209087d3ad92fba376f669692e07771f9945661a2d6fc9114ea098015ef44a0107993eb8136dfbb698c9b36736b093000e7c6d704c131c9534e45bb11069a6ce296ce15878e8cbd39800c52773c8841";
        bytes memory pk2 =
            hex"10ffae60e1c05689695b551f442351306be88aa4f98faa7521a120e12b5ba597bfa310c4ef9c59e31622bf678241ea9a0cebd5284f2326f38f0515652d16471bc699b9654870e49219b091f2717caf07b170e40bb23aad749a99dfd9cc072982";
        bytes memory pk3 =
            hex"0252a35315cec61341b676c42ac45c058f23e76f49194349f8c80d5e1414e6043560ebe4d129f23bff1f3692f4052396175832741f48604400a1b43f76e43f738f196472058158573c1e9d14de62b9137d2e4548507fd48b8f28a5ad9baa1063";
        bytes memory expectedAggregate =
            hex"11a1c66b14aa6d316416583e6490b081cf09dcd59ddd5001c4c9acf236b2c4ecab638ae32003ddce690de7ace0ed32c017363c0c67531f30c01618ab4df0dfe6c13c21ed5efe7943b7110c80e77ee085984d19df331de1514981c559b2659200";

        bytes memory res1 = BLSTUtils.addG1(
            BLSTUtils.formatUncompressedBLSPublicKey(pk1),
            BLSTUtils.formatUncompressedBLSPublicKey(pk2)
        );
        bytes memory res2 = BLSTUtils.addG1(res1, BLSTUtils.formatUncompressedBLSPublicKey(pk3));
        assertEq(res2, BLSTUtils.formatUncompressedBLSPublicKey(expectedAggregate));
    }

    function testVerifyBLSSignatureSuccess() public view {
        bytes memory message = "Hello Ethereum, it's Avalanche.";
        bytes memory pk =
            hex"000000000000000000000000000000000d89fe5ee3754167f889ff991c81e1717b665690955ad302a50e285dbe68aac93d03177fba54622bf5c4926ce90d0ac60000000000000000000000000000000015874901c0288afcdf442b402199d6a2ed38c5ac28d055637e7bafb44dc4a46a56db1c74caf0e339f6ed665dbc0c0c60";
        bytes memory sig =
            hex"1572ad7226ef5d2a52c89d824179e1723a523c7a30c8ff2643aee0b19aca4e3670a28d8749991075d72dd9d64aba703a0ed0d39f4a935cfc5b96c5fbc8985596a9556dec84286c0f945c93ebf33c96bc58ce5cf54a70c3b5221da81d01b6640e058b6a69058a96395acf26fb8013e93dcacbe83f93ed0697e6fd080132062c8e3fc4b63488e00ec5dd60dd2dcda6264112ce2d0e1e024de26bc4d75fb1a7b0fa823b44770570e1539cf500aa5fa2c5abd8e806246d0009d260e224e1f435cc97";

        bool valid = BLSTUtils.verifySignature(pk, sig, message);
        assertTrue(valid);
    }

    function testVerifyBLSAggregateSignatureSuccess() public view {
        bytes memory message = "Hello Ethereum, it's Avalanche.";
        bytes[] memory pks = new bytes[](2);
        pks[0] =
            hex"000000000000000000000000000000000d89fe5ee3754167f889ff991c81e1717b665690955ad302a50e285dbe68aac93d03177fba54622bf5c4926ce90d0ac60000000000000000000000000000000015874901c0288afcdf442b402199d6a2ed38c5ac28d055637e7bafb44dc4a46a56db1c74caf0e339f6ed665dbc0c0c60";
        pks[1] =
            hex"00000000000000000000000000000000001f2946023547012bf42ca45962136ebca74aeea4291a25e9bf46bf37c5fdd37ecad6d5a389bcecf3f3d1f4f73c44db000000000000000000000000000000000385c2753673a29e199f0a0dc02772b5e2dec7a3aa70e012b313195906a27ac95d0a4963d96340df6f4d13cf40f7c78f";
        bytes memory sig =
            hex"0ace4f38b5e22e8cfd27951b8327d5260def6e6370e334c3f74fb9b5ccd209d0d80d1a34bb804f6ac40dd991141d87b5147a0a08008b085fd27341456bb43654fed88ebca4407d0b5ae66a67df8141edc58fced0468b5e34cec33dc2d58f023d1770616fe829ef9af2350c7c9d7b72ee1c05edefc7cd831ac13114e8ae7e1466c7ef2990a8e6387f9def145be2c9be8b0cbc8f7e6ebcb1b099a1785ea2001d70887faaaf86c937a292049134d02c7bc9aeaed7a3354673985a88a6e3aff66ccd";
        bool valid = BLSTUtils.verifyAggregateSignature(pks, sig, message);
        assertTrue(valid);
    }
}
