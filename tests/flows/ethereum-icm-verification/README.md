# Ethereum External ICM Verification E2E Tests

1. Creates an Avalanche L1 on the Fuji testnet with 3 local validators, running a custom version of AvalancheGo.
1. Deploys an `AvalancheValidatorSetRegistry` contract on the Sepolia testnet.
1. Registers the Avalanche L1's validator set in the registry via an ICM message from the P-Chain, self-signed by the validator set of the L1.
1. Sends a transaction on the L1 to send an ICM message to Sepolia.
1. Gets an aggregate signature of the ICM message using the ICM signature aggregator.
1. Broadcasts the signed ICM message to Sepolia, and confirms successful inclusion and verification.
1. Adds a new validator to the L1, representing 40% of the new total weight.
1. Updates the validator set reigstered on the Sepolia registry via a validator set update message signed by sufficient weight of the previous validator set.
1. Sends and verifies another ICM message from the L1 to Sepolia.
