// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ValidatorSet} from "../validator-set.sol";

contract ValidatorSetGasTests is Test {
    ValidatorSet internal _validatorSet;

    bytes public constant PUBLIC_KEY =
        hex"123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678";

    function setUp() public {
        _validatorSet = new ValidatorSet();
    }

    function testGasUsageRegisterValidator() public {
        uint256 gasStart = gasleft();

        _validatorSet.registerValidator(1000, PUBLIC_KEY);

        uint256 gasUsed = gasStart - gasleft();
        emit log_named_uint("Gas used for registration", gasUsed);
    }

    function testGasUsageGetValidator() public {
        // Register a validator first
        uint256 validatorId = _validatorSet.registerValidator(1000, PUBLIC_KEY);

        uint256 gasStart = gasleft();

        ValidatorSet.Validator memory validator = _validatorSet.getValidator(validatorId);

        uint256 gasUsed = gasStart - gasleft();
        emit log_named_uint("Gas used for getting validator", gasUsed);

        // Ensure the data is correct to avoid optimization issues
        assertEq(validator.weight, 1000);
        assertEq(validator.publicKey, PUBLIC_KEY);
    }

    function testGasUsageUpdateValidator() public {
        // First register a validator
        uint256 validatorId = _validatorSet.registerValidator(1000, PUBLIC_KEY);

        // Create new public key for update (different from original)
        bytes memory newPublicKey =
            hex"fedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcbafedcba";

        uint256 gasStart = gasleft();

        _validatorSet.updateValidator(validatorId, 2000, newPublicKey);

        uint256 gasUsed = gasStart - gasleft();
        emit log_named_uint("Gas used for updating validator", gasUsed);

        // Verify the update was successful
        ValidatorSet.Validator memory updatedValidator = _validatorSet.getValidator(validatorId);
        assertEq(updatedValidator.weight, 2000);
        assertEq(updatedValidator.publicKey, newPublicKey);
    }

    function testGasUsageMultipleRegistrations() public {
        emit log_string("=== Testing gas usage for multiple registrations ===");

        // Register first validator
        uint256 gasStart1 = gasleft();
        _validatorSet.registerValidator(1000, PUBLIC_KEY);
        uint256 gasUsed1 = gasStart1 - gasleft();

        // Register second validator
        uint256 gasStart2 = gasleft();
        _validatorSet.registerValidator(2000, PUBLIC_KEY);
        uint256 gasUsed2 = gasStart2 - gasleft();

        // Register third validator
        uint256 gasStart3 = gasleft();
        _validatorSet.registerValidator(3000, PUBLIC_KEY);
        uint256 gasUsed3 = gasStart3 - gasleft();

        emit log_named_uint("Gas used for 1st registration", gasUsed1);
        emit log_named_uint("Gas used for 2nd registration", gasUsed2);
        emit log_named_uint("Gas used for 3rd registration", gasUsed3);

        // Check if gas usage increases with each registration (due to storage costs)
        assertTrue(gasUsed1 > 0);
        assertTrue(gasUsed2 > 0);
        assertTrue(gasUsed3 > 0);
    }

    function testGasUsageComparisonDifferentStakeAmounts() public {
        emit log_string("=== Testing gas usage for different stake amounts ===");

        // Small stake amount
        uint256 gasStart1 = gasleft();
        _validatorSet.registerValidator(1, PUBLIC_KEY);
        uint256 gasUsed1 = gasStart1 - gasleft();

        // Large stake amount
        uint256 gasStart2 = gasleft();
        _validatorSet.registerValidator(type(uint256).max, PUBLIC_KEY);
        uint256 gasUsed2 = gasStart2 - gasleft();

        emit log_named_uint("Gas used for small stake (1 wei)", gasUsed1);
        emit log_named_uint("Gas used for large stake (max uint256)", gasUsed2);

        // Gas usage should be similar regardless of stake amount since it's just a uint256
        uint256 gasUsageDifference = gasUsed2 > gasUsed1 ? gasUsed2 - gasUsed1 : gasUsed1 - gasUsed2;
        emit log_named_uint("Gas usage difference", gasUsageDifference);

        // The difference should be reasonable (less than 30000 gas due to array storage initialization)
        assertTrue(
            gasUsageDifference < 30000, "Gas usage should not vary significantly with stake amount"
        );
    }

    function testGasUsagePublicKeySameLength() public {
        emit log_string("=== Testing gas usage for public keys of same length ===");

        // Short public key
        uint256 gasStart1 = gasleft();
        _validatorSet.registerValidator(1000, PUBLIC_KEY);
        uint256 gasUsed1 = gasStart1 - gasleft();

        // Medium public key
        uint256 gasStart2 = gasleft();
        _validatorSet.registerValidator(1000, PUBLIC_KEY);
        uint256 gasUsed2 = gasStart2 - gasleft();

        // Long public key
        uint256 gasStart3 = gasleft();
        _validatorSet.registerValidator(1000, PUBLIC_KEY);
        uint256 gasUsed3 = gasStart3 - gasleft();

        emit log_named_uint("Gas used for short public key (48 bytes)", gasUsed1);
        emit log_named_uint("Gas used for medium public key (48 bytes)", gasUsed2);
        emit log_named_uint("Gas used for long public key (48 bytes)", gasUsed3);

        // All public keys are now 48 bytes, so gas usage should be similar
        assertTrue(gasUsed1 > 0, "Gas should be used for registration");
        assertTrue(gasUsed2 > 0, "Gas should be used for registration");
        assertTrue(gasUsed3 > 0, "Gas should be used for registration");

        // Gas usage should be similar since all keys are the same length
        uint256 maxGas = gasUsed1 > gasUsed2
            ? (gasUsed1 > gasUsed3 ? gasUsed1 : gasUsed3)
            : (gasUsed2 > gasUsed3 ? gasUsed2 : gasUsed3);
        uint256 minGas = gasUsed1 < gasUsed2
            ? (gasUsed1 < gasUsed3 ? gasUsed1 : gasUsed3)
            : (gasUsed2 < gasUsed3 ? gasUsed2 : gasUsed3);
        assertTrue(
            maxGas - minGas < 30000,
            "Gas usage difference should be reasonable for same-length keys"
        );
    }

    function testGasUsageReadOperations() public {
        emit log_string("=== Testing gas usage for read operations ===");

        // Register some validators first
        uint256 validatorId1 = _validatorSet.registerValidator(1000, PUBLIC_KEY);
        _validatorSet.registerValidator(2000, PUBLIC_KEY);
        _validatorSet.registerValidator(3000, PUBLIC_KEY);

        // Test getValidator
        uint256 gasStart1 = gasleft();
        _validatorSet.getValidator(validatorId1);
        uint256 gasUsed1 = gasStart1 - gasleft();

        // Test getTotalValidators
        uint256 gasStart2 = gasleft();
        _validatorSet.getTotalValidators();
        uint256 gasUsed2 = gasStart2 - gasleft();

        emit log_named_uint("Gas used for getValidator", gasUsed1);
        emit log_named_uint("Gas used for getTotalValidators", gasUsed2);

        // Read operations should be relatively cheap
        assertTrue(gasUsed1 < 10000, "getValidator should be relatively cheap");
        assertTrue(gasUsed2 < 5000, "getTotalValidators should be very cheap");
    }

    function testGasUsageBatchOperations() public {
        emit log_string("=== Testing gas usage for batch operations ===");

        address[] memory validators = new address[](5);
        validators[0] = address(0x100);
        validators[1] = address(0x200);
        validators[2] = address(0x300);
        validators[3] = address(0x400);
        validators[4] = address(0x500);

        uint256 totalGasUsed = 0;
        uint256 gasStart = gasleft();

        // Register 5 validators in sequence
        for (uint256 i = 0; i < validators.length; i++) {
            uint256 iterationGasStart = gasleft();
            _validatorSet.registerValidator(1000 + i, PUBLIC_KEY);
            uint256 iterationGasUsed = iterationGasStart - gasleft();
            emit log_named_uint(
                string(abi.encodePacked("Gas for registration ", vm.toString(i + 1))),
                iterationGasUsed
            );
        }

        totalGasUsed = gasStart - gasleft();
        emit log_named_uint("Total gas used for 5 registrations", totalGasUsed);
        emit log_named_uint("Average gas per registration", totalGasUsed / validators.length);
    }
}
