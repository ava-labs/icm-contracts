# Ethereum ICM Verification

Smart contracts for verifying Avalanche ICM messages on Ethereum using the BLS precompiles from [EIP-2537](https://eips.ethereum.org/EIPS/eip-2537).

> NOTE: These contracts are proof-of-concepts and have not been audited. **Do not use them in production.**

## How it works

- Avalanche validator sets get registered in the `AvalancheValidatorSetRegistry` contract
- Validator set updates are authenticated by signatures from the current set  
- Other contracts call `verifyICMMessage()` to validate ICM messages from Avalanche L1s

## Contracts / Libraries
- `AvalancheValidatorSetRegistry` - manages validator sets and verifies ICM messages
- `BLSTUtils` - BLS12-381 crypto operations using Ethereum precompiles
- `ICMUtils` - ICM message parsing and signature verification
- `ValidatorSetUtils` - validator set data structures and serialization

## Usage Example

```solidity
// Deploy registry for Avalanche Fuji testnet (network ID 5)
AvalancheValidatorSetRegistry registry = new AvalancheValidatorSetRegistry(5);

// Register initial validator set (must be self-signed by that set)
uint256 validatorSetID = registry.registerValidatorSet(icmMessage, validatorBytes);

// Verify messages from that L1
bool isValid = registry.verifyICMMessage(validatorSetID, someICMMessage);

// Update the validator set when changes are made to it.
registry.updateValidatorSet(validatorSetID, validatorSetStateICMMessage, updatedValidatorBytes);

```

## Integration

```solidity
contract MyDApp {
    AvalancheValidatorSetRegistry immutable registry;
    uint256 immutable validatorSetID;
    
    function processAvalancheMessage(ICMMessage memory message) external {
        require(registry.verifyICMMessage(validatorSetID, message), "Invalid ICM message");
        
        // Process the verified message...
        bytes memory data = message.unsignedMessage.payload;
    }
}
```

Since multiple validator sets can be registered per blockchain ID, make sure you're using the correct `validatorSetID` for your intended L1.
