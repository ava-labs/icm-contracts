// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

struct FieldPoint2 {
    bytes32[2] u;
    bytes32[2] u_I;
}

/**
 * @title BLST
 * @notice Utility library for BLS12-381 operations.
 */
library BLST {
    uint256 constant BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH = 96;

    address constant BLS12381_G1_ADD_PRECOMPILE =
        address(0x000000000000000000000000000000000000000b);
    address constant BLS12381_G1_MSM_PRECOMPILE =
        address(0x000000000000000000000000000000000000000C);
    address constant BLS12381_G2_ADD_PRECOMPILE =
        address(0x000000000000000000000000000000000000000d);
    address constant BLS12381_G2_MSM_PRECOMPILE =
        address(0x000000000000000000000000000000000000000E);
    address constant BLS12381_PAIRING_CHECK_PRECOMPILE =
        address(0x000000000000000000000000000000000000000F);
    address constant BLS12381_MAP_FP2_G2_PRECOMPILE =
        address(0x0000000000000000000000000000000000000011);

    bytes constant BLS_G1_GENERATOR =
        hex"0000000000000000000000000000000017f1d3a73197d7942695638c4fa9ac0fc3688c4f9774b905a14e3a3f171bac586c55e83ff97a1aeffb3af00adb22c6bb0000000000000000000000000000000008b3f481e3aaa0f1a09e30ed741d8ae4fcf5e095d5d00af600db18cb2c04b3edd03cc744a2888ae40caa232946c5e7e1";
    bytes constant BLS_G1_NEG_GENERATOR =
        hex"0000000000000000000000000000000017f1d3a73197d7942695638c4fa9ac0fc3688c4f9774b905a14e3a3f171bac586c55e83ff97a1aeffb3af00adb22c6bb00000000000000000000000000000000114d1d6855d545a8aa7d76c8cf2e21f267816aef1db507c96655b9d5caac42364e6f38ba0ecb751bad54dcd6b939c2ca";
    bytes constant BLS_G2_GENERATOR =
        hex"00000000000000000000000000000000024aa2b2f08f0a91260805272dc51051c6e47ad4fa403b02b4510b647ae3d1770bac0326a805bbefd48056c8c121bdb80000000000000000000000000000000013e02b6052719f607dacd3a088274f65596bd0d09920b61ab5da61bbdc7f5049334cf11213945d57e5ac7d055d042b7e000000000000000000000000000000000ce5d527727d6e118cc9cdc6da2e351aadfd9baa8cbdd3a76d429a695160d12c923ac9cc3baca289e193548608b82801000000000000000000000000000000000606c4a02ea734cc32acd2b02bc28b99cb3e287e85a763af267492ab572e99ab3f370d275cec1da1aaa9075ff05f79be";
    bytes constant BLS12381G2_SIG_DST = "BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_";
    bytes constant BLS12381G2_POP_DST = "BLS_POP_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_";

    /**
     * @notice Formats a 96-byte uncompressed BLS public key into the 128-byte format expected by the
     * BLS12381_G1_ADD_PRECOMPILE.
     * @param publicKey The 96-byte uncompressed BLS public key, as produced by the BLST library's
     * P1Affine.Serialize() function.
     * @return The 128-byte serialized public key. Its X and Y coordinates are left-padded to be 64 bytes each, for a
     * total of 128 bytes.
     */
    function formatUncompressedBLSPublicKey(
        bytes memory publicKey
    ) internal pure returns (bytes memory) {
        require(publicKey.length == 96, "Invalid input public key length");
        bytes memory res = new bytes(128);

        // Copy the X coordinate.
        for (uint256 i = 0; i < 48; ++i) {
            res[16 + i] = publicKey[i];
        }

        // Copy the Y coordinate.
        for (uint256 i = 0; i < 48; ++i) {
            res[80 + i] = publicKey[i + 48];
        }

        return res;
    }

    /**
     * @notice Turns the 128-byte format expected by the BLS12381_G1_ADD_PRECOMPILE into
     * a 96-byte uncompressed BLS public key
     * @param publicKey The the 128-byte padded format.
     * @return The 96-byte uncompressed BLS public key.
     */
    function getUncompressedBlsPublicKey(
        bytes memory publicKey
    ) internal pure returns (bytes memory) {
        require(publicKey.length == 128, "Invalid input public key length");
        bytes memory res = new bytes(96);

        // Copy the X coordinate.
        for (uint256 i = 0; i < 48; ++i) {
            res[i] = publicKey[16 + i];
        }

        // Copy the Y coordinate.
        for (uint256 i = 0; i < 48; ++i) {
            res[48 + i] = publicKey[i + 80];
        }

        return res;
    }

    /**
     * @notice Given a secret key, get the public counterpart
     * @dev According to the EIP-2537 spec, this method can fail as it may not land on the curve or
     * is not in the correct subgroup
     * @param secretKey is a 256 bit scalar.
     * @return The 128-byte serialized public key. Its X and Y coordinates are left-padded to be 64 bytes each,
     * for a total of 128 bytes.
     */
    function getPublicKeyFromSecret(
        uint256 secretKey
    ) internal view returns (bytes memory) {
        bytes memory input = new bytes(160);
        for (uint i = 0; i < 128; i++) {
            input[i] = BLS_G1_GENERATOR[i];
        }
        bytes32 sk = bytes32(secretKey);
        for (uint j = 0; j < 32; j++) {
            input[128 + j] = sk[j];
        }

        (bool success, bytes memory result) =
            BLS12381_G1_MSM_PRECOMPILE.staticcall(abi.encodePacked(input));
        require(success, "Failed to perform scalar multiplication over G1");
        return result;
    }

    /**
     * @notice Formats a 192-byte uncompressed BLS signature into the 256-byte format expected by the
     * BLS12381_PAIRING_CHECK_PRECOMPILE.
     * @param signature The 192-byte uncompressed BLS signature. Must have the format [x.c1, x.c0, y.c1, y.c0].
     * @return The 256-byte formatted signature. Has the format
     * [16 pad + x.c0, 16 pad + x.c1, 16 pad +  y.c0, 16 pad + y.c1].
     */
    function formatUncompressedBLSTSignature(
        bytes memory signature
    ) internal pure returns (bytes memory) {
        require(signature.length == 192, "Invalid input signature length");
        bytes memory res = new bytes(256);

        // Copy X0
        for (uint256 i = 0; i < 48; ++i) {
            res[16 + i] = signature[i + 48];
        }

        // Copy X1
        for (uint256 i = 0; i < 48; ++i) {
            res[80 + i] = signature[i];
        }

        // Copy Y0
        for (uint256 i = 0; i < 48; ++i) {
            res[144 + i] = signature[i + 144];
        }

        // Copy Y1
        for (uint256 i = 0; i < 48; ++i) {
            res[208 + i] = signature[i + 96];
        }

        return res;
    }

    /**
     * @notice Formats the the 256-byte format expected by the  BLS12381_PAIRING_CHECK_PRECOMPILE into the
     * 192-byte uncompressed BLS signature.
     * @param signature The 256-byte BLS signature. Must have the format
     * [16 pad + x.c0, 16 pad + x.c1, 16 pad +  y.c0, 16 pad + y.c1].
     * @return The 192-byte uncompressed signature. Has the format [x.c1, x.c0, y.c1, y.c0]
     */
    function getUncompressedBLSTSignature(
        bytes memory signature
    ) internal pure returns (bytes memory) {
        require(signature.length == 256, "Invalid input signature length");
        bytes memory res = new bytes(192);

        // COPY X0
        for (uint256 i = 0; i < 48; i++) {
            res[i + 48] = signature[16 + i];
        }

        // COPY X1
        for (uint256 i = 0; i < 48; i++) {
            res[i] = signature[80 + i];
        }

        // COPY Y0
        for (uint256 i = 0; i < 48; i++) {
            res[i + 144] = signature[144 + i];
        }

        // COPY Y1
        for (uint256 i = 0; i < 48; i++) {
            res[i + 96] = signature[208 + i];
        }

        return res;
    }

    /**
     * @notice Aggregates a list of public keys.
     * @param publicKeys The public keys to aggregate. Each public key must be in uncompressed, and its X and Y coordinates must be
     * left-padded to be 64 bytes each, for a total of 128 bytes.
     * @return The aggregated public key.
     */
    function aggregatePublicKeys(
        bytes[] memory publicKeys
    ) internal view returns (bytes memory) {
        // Use the BLS public key aggregation precompile to aggregate the public keys.
        require(publicKeys.length > 0, "Missing public keys");
        bytes memory aggregatedPublicKey = publicKeys[0];
        for (uint256 i = 1; i < publicKeys.length; i++) {
            aggregatedPublicKey = addG1(aggregatedPublicKey, publicKeys[i]);
        }
        return aggregatedPublicKey;
    }

    /**
     * @notice Create a signature over a message with a secret key.
     * @dev According to the EIP-2537 spec, G2 MSM method can fail as it may not land on the curve or
     * is not in the correct subgroup
     * @param secretKey is a 256 bit scalar.
     * @param message the message to be signed.
     * @return The 192-byte uncompressed BLS signature
     */
    function createSignature(
        uint256 secretKey,
        bytes memory message
    ) internal view returns (bytes memory) {
        return getUncompressedBLSTSignature(_createSignatureRaw(secretKey, message));
    }

    /**
     * @notice Create an aggregate signature over a message with a list of secret keys.
     * @dev According to the EIP-2537 spec, this method can fail as it may not land on the curve or
     * is not in the correct subgroup
     * @param secretKeys a vector 256 bit scalars.
     * @param message the message to be signed.
     * @return The 192-byte uncompressed aggregated BLS signature
     */
    function createAggregateSignature(
        uint256[] memory secretKeys,
        bytes memory message
    ) internal view returns (bytes memory) {
        bytes memory sig = new bytes(256);
        for (uint256 i = 0; i < secretKeys.length; i++) {
            sig = addG2(sig, _createSignatureRaw(secretKeys[i], message));
        }
        return getUncompressedBLSTSignature(sig);
    }

    function _createSignatureRaw(
        uint256 secretKey,
        bytes memory message
    ) private view returns (bytes memory) {
        bytes memory messageG2 = hashToG2(message, BLS12381G2_SIG_DST);
        bytes memory input = new bytes(288);
        for (uint i = 0; i < 256; i++) {
            input[i] = messageG2[i];
        }
        bytes32 sk = bytes32(secretKey);
        for (uint j = 0; j < 32;  j++) {
            input[256 + j] = sk[j];
        }

        (bool success, bytes memory result) =
            BLS12381_G2_MSM_PRECOMPILE.staticcall(abi.encodePacked(input));
        require(success, "Failed to perform scalar multiplication over G2");
        return result;
    }

    /**
     * @notice Verifies a BLS12-381 signature of the given message using the given public key.
     * @param publicKey The public key to verify the signature against. Must be in uncompressed, and its X and Y coordinates must be
     * left-padded to be 64 bytes each, for a total of 128 bytes.
     * @param signature The signature to verify. Must be in uncompressed form, 192 bytes long, and have the format [x.c1, x.c0, y.c1, y.c0].
     * This is the format used by the BLST library for the P2Affine.Serialize() function.
     * @param message The message to verify the signature against.
     * @return True if the signature is valid, false otherwise.
     */
    function verifySignature(
        bytes memory publicKey,
        bytes memory signature,
        bytes memory message
    ) internal view returns (bool) {
        return _verifySignature(publicKey, signature, message, BLS12381G2_SIG_DST);
    }

    /**
     * @notice Verifies a BLS12-381 aggregate signature of the given message using the given public keys.
     * @param publicKeys The public keys to verify the signature against. Each public key must be in uncompressed, and its X and Y coordinates must be
     * left-padded to be 64 bytes each, for a total of 128 bytes.
     * @param signature The signature to verify. Must be in uncompressed form, 192 bytes long, and have the format [x.c1, x.c0, y.c1, y.c0].
     * This is the format used by the BLST library for the P2Affine.Serialize() function.
     * @param message The message to verify the signature against.
     * @return True if the signature is valid for the public key resulting from aggregating the given public keys, false otherwise.
     */
    function verifyAggregateSignature(
        bytes[] memory publicKeys,
        bytes memory signature,
        bytes memory message
    ) internal view returns (bool) {
        bytes memory aggregatePublicKey = aggregatePublicKeys(publicKeys);
        return verifySignature(aggregatePublicKey, signature, message);
    }

    /**
     * @notice Verifies a BLS12-381 proof of possession of the given public key.
     * @param publicKey The public key to verify the proof of possession against. Must be in uncompressed, and its X and Y coordinates must be
     * left-padded to be 64 bytes each, for a total of 128 bytes.
     * @param signature The signature to verify. Must be in uncompressed form, 192 bytes long, and have the format [x.c1, x.c0, y.c1, y.c0].
     * This is the format used by the BLST library for the P2Affine.Serialize() function.
     * @param message The message to verify the proof of possession against.
     * @return True if the proof of possession is valid, false otherwise.
     */
    function verifyProofOfPossession(
        bytes memory publicKey,
        bytes memory signature,
        bytes memory message
    ) internal view returns (bool) {
        return _verifySignature(publicKey, signature, message, BLS12381G2_POP_DST);
    }

    /**
     * @notice Hashes a message to the G2 curve
     * @dev Original source: https://github.com/ethyla/bls12-381-hash-to-curve/blob/main/src/HashToCurve.sol
     */
    function hashToG2(
        bytes memory message,
        bytes memory dst
    ) internal view returns (bytes memory) {
        FieldPoint2[2] memory u = hashToFieldFp2(message, dst);
        bytes memory q0 = _mapFpToG2(u[0]);
        bytes memory q1 = _mapFpToG2(u[1]);
        return addG2(q0, q1);
    }

    /**
     * @notice Computes a field point from a message
     * @dev Follows https://datatracker.ietf.org/doc/html/rfc9380#section-5.2
     * @param message Arbitrarylength byte string to be hashed
     * @param dst The domain separation tag
     * @return Two field points
     */
    function hashToFieldFp2(
        bytes memory message,
        bytes memory dst
    ) internal view returns (FieldPoint2[2] memory) {
        // 1. len_in_bytes = count * m * L
        // so always 2 * 2 * 64 = 256
        uint16 lenInBytes = 256;
        // 2. uniform_bytes = expand_message(msg, DST, len_in_bytes)
        bytes32[] memory pseudoRandomBytes = _expandMsgXmd(message, dst, lenInBytes);
        FieldPoint2[2] memory u;
        // No loop here saves 800 gas hardcoding offset an additional 300
        // 3. for i in (0, ..., count - 1):
        // 4.   for j in (0, ..., m - 1):
        // 5.     elm_offset = L * (j + i * m)
        // 6.     tv = substr(uniform_bytes, elm_offset, HTF_L)
        // uint8 HTF_L = 64;
        // bytes memory tv = new bytes(64);
        // 7.     e_j = OS2IP(tv) mod p
        // 8.   u_i = (e_0, ..., e_(m - 1))
        // tv = bytes.concat(pseudo_random_bytes[0], pseudo_random_bytes[1]);
        u[0].u = _modfield(pseudoRandomBytes[0], pseudoRandomBytes[1]);
        u[0].u_I = _modfield(pseudoRandomBytes[2], pseudoRandomBytes[3]);
        u[1].u = _modfield(pseudoRandomBytes[4], pseudoRandomBytes[5]);
        u[1].u_I = _modfield(pseudoRandomBytes[6], pseudoRandomBytes[7]);
        // 9. return (u_0, ..., u_(count - 1))
        return u;
    }

    /**
     * @notice Adds two G1 points
     * @param p0 The first G1 point
     * @param p1 The second G1 point
     * @return The sum of the two G1 points
     */
    function addG1(bytes memory p0, bytes memory p1) internal view returns (bytes memory) {
        require(p0.length == 128 && p1.length == 128, "Invalid G1 point length");
        (bool success, bytes memory result) =
            BLS12381_G1_ADD_PRECOMPILE.staticcall(abi.encodePacked(p0, p1));
        require(success, "Failed to add G1 points");
        return result;
    }

    /**
     * @notice Adds two G2 points
     * @param q0 The first G2 point
     * @param q1 The second G2 point
     * @return The sum of the two G2 points
     */
    function addG2(bytes memory q0, bytes memory q1) internal view returns (bytes memory) {
        require(q0.length == 256 && q1.length == 256, "Invalid G2 point length");
        bytes memory addG2input = abi.encodePacked(q0, q1);
        (bool success, bytes memory output) = BLS12381_G2_ADD_PRECOMPILE.staticcall(addG2input);
        require(success, "Failed to add G2 points");
        return output;
    }

    function _verifySignature(
        bytes memory publicKey,
        bytes memory signature,
        bytes memory message,
        bytes memory dst
    ) private view returns (bool) {
        // Check the input lengths
        require(publicKey.length == 128, "Invalid public key length");
        require(signature.length == 192, "Invalid signature length");

        // Hash the message to the G2 curve
        bytes memory messageG2 = hashToG2(message, dst);

        bytes memory pairingCheckInput = abi.encodePacked(
            publicKey, messageG2, BLS_G1_NEG_GENERATOR, formatUncompressedBLSTSignature(signature)
        );
        (bool success, bytes memory output) =
            BLS12381_PAIRING_CHECK_PRECOMPILE.staticcall(pairingCheckInput);
        require(success, "Failed to perform pairing check");
        require(output.length == 32, "Invalid pairing check output length");
        return uint256(bytes32(output)) == 1;
    }

    /**
     * @dev Maps a field point to a G2 point
     * @param fp2 The field point to map
     * @return The G2 point
     */
    function _mapFpToG2(
        FieldPoint2 memory fp2
    ) private view returns (bytes memory) {
        bytes memory mapFp2ToG2input = abi.encodePacked(fp2.u[0], fp2.u[1], fp2.u_I[0], fp2.u_I[1]);
        (bool success, bytes memory output) =
            BLS12381_MAP_FP2_G2_PRECOMPILE.staticcall(mapFp2ToG2input);
        require(success, "Failed to map Fp2 to G2");
        return output;
    }

    /**
     * @notice Computes a field point from a message
     * @dev Follows https://datatracker.ietf.org/doc/html/rfc9380#section-5.3
     * @dev bytes32[] because len_in_bytes is always a multiple of 32 in our case even 128
     * @param message Arbitrarylength byte string to be hashed
     * @param dst The domain separation tag of at most 255 bytes
     * @param lenInBytes The length of the requested output in bytes
     * @return A field point
     */
    function _expandMsgXmd(
        bytes memory message,
        bytes memory dst,
        uint16 lenInBytes
    ) private pure returns (bytes32[] memory) {
        // 1.  ell = ceil(len_in_bytes / b_in_bytes)
        // b_in_bytes seems to be 32 for sha256
        // ceil the division
        uint256 ell = (lenInBytes - 1) / 32 + 1;

        // 2.  ABORT if ell > 255 or len_in_bytes > 65535 or len(DST) > 255
        require(ell <= 255, "len_in_bytes too large for sha256");
        // Not really needed because of parameter type
        // require(lenInBytes <= 65535, "len_in_bytes too large");
        // no length normalizing via hashing
        require(dst.length <= 255, "dst too long");

        bytes memory dstPrime = bytes.concat(dst, bytes1(uint8(dst.length)));

        // 4.  Z_pad = I2OSP(0, s_in_bytes)
        // this should be sha256 blocksize so 64 bytes
        // bytes
        //     memory zPad = hex"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";
        // spliting like below saves 250 gas
        bytes32 zPad1 = hex"0000000000000000000000000000000000000000000000000000000000000000";
        bytes32 zPad2 = hex"0000000000000000000000000000000000000000000000000000000000000000";

        // 5.  l_i_b_str = I2OSP(len_in_bytes, 2)
        // length in bytes string
        bytes2 libStr = bytes2(lenInBytes);

        // 6.  msg_prime = Z_pad || msg || l_i_b_str || I2OSP(0, 1) || DST_prime
        bytes memory msgPrime = bytes.concat(zPad1, zPad2, message, libStr, hex"00", dstPrime);

        bytes32 b_0;
        bytes32[] memory b = new bytes32[](ell);

        // 7.  b_0 = H(msg_prime)
        b_0 = sha256(msgPrime);

        // 8.  b_1 = H(b_0 || I2OSP(1, 1) || DST_prime)
        b[0] = sha256(bytes.concat(b_0, hex"01", dstPrime));

        // 9.  for i in (2, ..., ell):
        for (uint8 i = 2; i <= ell; i++) {
            // 10.    b_i = H(strxor(b_0, b_(i - 1)) || I2OSP(i, 1) || DST_prime)
            bytes memory tmp = abi.encodePacked(b_0 ^ b[i - 2], i, dstPrime);
            b[i - 1] = sha256(tmp);
        }
        // 11. uniform_bytes = b_1 || ... || b_ell
        // 12. return substr(uniform_bytes, 0, len_in_bytes)
        // Here we don't need the uniform_bytes because b is already properly formed
        return b;
    }

    /**
     * @dev Computes the mod against the bls12-381 field modulus
     */
    function _modfield(bytes32 _b1, bytes32 _b2) private view returns (bytes32[2] memory r) {
        assembly {
            let bl := 0x40
            let ml := 0x40

            let freemem := mload(0x40) // Free memory pointer is always stored at 0x40

            // arg[0] = base.length @ +0
            mstore(freemem, bl)
            // arg[1] = exp.length @ +0x20
            mstore(add(freemem, 0x20), 0x20)
            // arg[2] = mod.length @ +0x40
            mstore(add(freemem, 0x40), ml)

            // arg[3] = base.bits @ + 0x60
            // places the first 32 bytes of _b1 and the last 32 bytes of _b2
            mstore(add(freemem, 0x60), _b1)
            mstore(add(freemem, 0x80), _b2)

            // arg[4] = exp.bits @ +0x60+base.length
            // exponent always 1
            mstore(add(freemem, 0xa0), 1)

            // arg[5] = mod.bits @ +96+base.length+exp.length
            // this field_modulus as hex 4002409555221667393417789825735904156556882819939007885332058136124031650490837864442687629129015664037894272559787
            // we add the 0 prefix so that the result will be exactly 64 bytes
            // saves 300 gas per call instead of sending it along every time
            // places the first 32 bytes and the last 32 bytes of the field modulus
            mstore(
                add(freemem, 0xc0),
                0x000000000000000000000000000000001a0111ea397fe69a4b1ba7b6434bacd7
            )
            mstore(
                add(freemem, 0xe0),
                0x64774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab
            )

            // Invoke contract 0x5, put return value right after mod.length, @ 0x60
            let success :=
                staticcall(
                    sub(gas(), 1350), // gas
                    0x5, // mpdexp precompile
                    freemem, //input offset
                    0x100, // input size  = 0x60+base.length+exp.length+mod.length
                    freemem, // output offset
                    ml // output size
                )
            switch success
            case 0 { invalid() } //fail where we haven't enough gas to make the call

            // point to mod length, result was placed immediately after
            r := freemem
            //adjust freemem pointer
            mstore(0x40, add(freemem, ml))
        }
    }
}
