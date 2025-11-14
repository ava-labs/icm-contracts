// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import "forge-std/Test.sol";
import {
    ICMMessage,
    ICMUnsignedMessage,
    ICMSignature
} from "@avalabs/avalanche-contracts/teleporter/ITeleporterMessenger.sol";
import "../AvalancheValidatorSetRegistry.sol";
import "../EthWarp.sol";
import "../utils/ValidatorSets.sol";
import "../utils/BLST.sol";

contract EthWarpTest is Test {
    AvalancheValidatorSetRegistry registry;
    EthWarp warp;

    uint32 constant NETWORK_ID = 1;
    bytes32 constant AvalancheBlockchainId = bytes32(hex"3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7");

    function setUp() public {
        registry = new AvalancheValidatorSetRegistry(NETWORK_ID, AvalancheBlockchainId);
        warp = new EthWarp(1, AvalancheBlockchainId);
    }

    /**
     * @dev The `getVerifiedWarpBlockHash` method is part of the IWarpExt interface, but it
      * is specifically not implemented for the Ethereum Warp messenger
     */
    function testWarpBlochHashReverts() public {
        vm.expectRevert("This method cannot be called on Ethereum");
        warp.getVerifiedWarpBlockHash(0);
    }

    /**
     * @dev The `getVerifiedWarpMessage` method is part of the IWarpExt interface, but it
     * is specifically not implemented for the Ethereum Warp messenger
     */
    function testGetVerifiedWarpMessageReverts() public {
        vm.expectRevert("This method can't be called on Ethereum, use `getVerifiedICMMessage` instead");
        warp.getVerifiedWarpMessage(0);
    }

    /**
     * @dev An Ethereum warp contract must be able to look up an `AvalancheValidatorSetRegistry`
     * contract from the respective blockchain id in order to validate a message.
     *
     * Here we test that if this look up fails, the function reverts
     */
    function testUnregisteredBlockchain() public {
        // the blockchain is not registered with this contract
        assertFalse(warp.isChainRegistered(AvalancheBlockchainId));
        ICMMessage memory message = icmMessageFixture();
        vm.expectRevert( "Cannot receive a Warp message from a chain whose validator set is unknown");
        warp.getVerifiedICMMessage(message);
    }

    /**
     * @dev The ICM message given to the warp contract will reach out to an `AvalancheValidatorSetRegistry`
     * with no validators registered. This should cause message verification to fail.
     */
    function testWrongValidatorSet() public {
        // register the `AvalancheValidatorSetRegistry` contract
        warp.registerChain(AvalancheBlockchainId, address(registry));
        assertTrue(warp.isChainRegistered(AvalancheBlockchainId));

        ICMMessage memory message = icmMessageFixture();

        // The call should fail when attempting to find the current validator set
        vm.expectRevert();
        warp.getVerifiedICMMessage(message);
    }

    /**
     * @dev Test happy flow that warp can receive a message signed by a quorum of validators
     */
    function testGetVerifiedMessageFromPayload() public {
        // add a validator set to the registry with all equal weights
        registerValidatorSet([uint64(100), uint64(100), uint64(100), uint64(100)]);

        // register this registry with the warp contract
        warp.registerChain(AvalancheBlockchainId, address(registry));
        // verify that we can receive and validate this message
        warp.getVerifiedICMMessage(icmMessageFixture());
    }

    function secretKeysFixture() private pure returns (uint256[4] memory) {
        return [uint256(1337), uint256(1338), uint256(1339), uint256(1340)];
    }

    /**
     * @dev An ICM message used for testing. It is signed by all validators except the second
     */
    function icmMessageFixture() private view returns (ICMMessage memory) {
        AddressedCall memory addressedCall = AddressedCall({
            sourceAddress: hex"deadbeef",
            payload: hex"48656c6c6f2c20576f726c6421"
        });
        ICMUnsignedMessage memory unsignedMessage =  ICMUnsignedMessage ({
            avalancheNetworkID: 1,
            avalancheSourceBlockchainID: AvalancheBlockchainId,
            payload: ICM.serializeAddressedCall(addressedCall)
        });
        bytes memory unsignedMessageBytes = ICM.serializeICMUnsignedMessage(unsignedMessage);
        uint256[4] memory secretKeys = secretKeysFixture();
        uint256[] memory sks = new uint256[] (3);
        sks[0] = secretKeys[0];
        sks[1] = secretKeys[2];
        sks[2] = secretKeys[3];
        bytes memory sig = BLST.createAggregateSignature(sks, unsignedMessageBytes);
        return ICMMessage({
            unsignedMessage: unsignedMessage,
            unsignedMessageBytes: unsignedMessageBytes,
            signature: ICMSignature({
                signers: hex"0d",
                signature: sig
            })
        });
    }

    /**
     * @dev A factory function to register a set of four validators to the registry
     * The weights are passed in as parameter
     */
    function registerValidatorSet(uint64[4] memory weights) private {
        // construct public keys
        bytes[] memory publicKeys = new bytes[](4);
        uint256[4] memory secretKeys = secretKeysFixture();
        publicKeys[0] = BLST.getPublicKeyFromSecret(secretKeys[0]);
        publicKeys[1] = BLST.getPublicKeyFromSecret(secretKeys[1]);
        publicKeys[2] = BLST.getPublicKeyFromSecret(secretKeys[2]);
        publicKeys[3] = BLST.getPublicKeyFromSecret(secretKeys[3]);

        // construct the validator set
        Validator[] memory validators = new Validator[](4);
        for (uint256 i = 0; i < 4; ++i) {
            validators[i] = Validator({
                blsPublicKey: publicKeys[i],
                weight: weights[i]
            });
        }

        bytes memory validatorBytes = ValidatorSets.serializeValidators(validators);

        // construct the validator set state payload
        ValidatorSetStatePayload memory validatorSetState = ValidatorSetStatePayload({
            avalancheBlockchainID: AvalancheBlockchainId,
            pChainHeight: 1,
            pChainTimestamp: 1,
            validatorSetHash: sha256(validatorBytes)
        });

        // construct the addressed call for registering validators
        AddressedCall memory addressedCall = AddressedCall({
            sourceAddress: new bytes(0),
            payload: ValidatorSets.serializeValidatorSetStatePayload(validatorSetState)
        });

        // construct the unsigned ICM message
        ICMUnsignedMessage memory unsignedMessage =  ICMUnsignedMessage ({
            avalancheNetworkID: 1,
            avalancheSourceBlockchainID: AvalancheBlockchainId,
            payload: ICM.serializeAddressedCall(addressedCall)
        });

        bytes memory unsignedMessageBytes = ICM.serializeICMUnsignedMessage(unsignedMessage);

        // sign the ICM message
        uint256[] memory sks = new uint256[] (4);
        sks[0] = secretKeys[0];
        sks[1] = secretKeys[1];
        sks[2] = secretKeys[2];
        sks[3] = secretKeys[3];
        bytes memory sig = BLST.createAggregateSignature(sks, unsignedMessageBytes);

        // construct the final ICM message
        ICMMessage memory icmMessage = ICMMessage({
            unsignedMessage: unsignedMessage,
            unsignedMessageBytes: unsignedMessageBytes,
            signature: ICMSignature({
                signers: hex"0f",
                signature: sig
            })
        });
        uint256 expected_id = registry.nextValidatorSetID();
        uint256 id = registry.registerValidatorSet(icmMessage, validatorBytes);
        assertEq(id, expected_id);
    }
}
