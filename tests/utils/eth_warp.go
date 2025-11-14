package utils

import (
    "context"
    "crypto/ecdsa"
    "encoding/hex"
    "math/big"

    "github.com/ava-labs/avalanchego/ids"
    ethwarp "github.com/ava-labs/icm-contracts/abi-bindings/go/EthWarp"
    "github.com/ava-labs/libevm/common"
    "github.com/ava-labs/libevm/log"
    "github.com/ava-labs/subnet-evm/accounts/abi/bind"
    "github.com/ava-labs/subnet-evm/ethclient"
    . "github.com/onsi/gomega"
)

func DeployEthWarp(
    ctx context.Context,
    fundedKey *ecdsa.PrivateKey,
    chainID *big.Int,
    client ethclient.Client,
    warpChainID ids.ID,
) (common.Address, *ethwarp.EthWarp) {
    opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, chainID)
    Expect(err).Should(BeNil())
    log.Info("Deploying WARP with dest chain", "id", warpChainID.Hex())
    address, tx, ethWarp, err := ethwarp.DeployEthWarp(opts, client, chainID, warpChainID)
    Expect(err).Should(BeNil())

    WaitForTransactionSuccessWithClient(ctx, client, tx.Hash())

    return address, ethWarp
}

func RegisterValSetRegistryWithEthWarp(
    ctx context.Context,
    fundedKey *ecdsa.PrivateKey,
    chainID *big.Int,
    client ethclient.Client,
    ethWarp *ethwarp.EthWarp,
    avalancheL1BlockchainID ids.ID,
    validatorSetRegistryAddress common.Address,
) {
    opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, chainID)
    Expect(err).Should(BeNil())
    tx, err := ethWarp.RegisterChain(opts, avalancheL1BlockchainID, validatorSetRegistryAddress)
    if err != nil {
        log.Error("Failed to register ValidatorSetRegistry to EthWarp contract", "error", err)
        log.Info("Registering a ValidatorSetRegistry to an EthWarp contract failed", "error", err)
    } else if tx != nil {
        log.Info("EthWarp contract registered validator set", "Blockchain ID", avalancheL1BlockchainID, "contract", validatorSetRegistryAddress)
        log.Info("Register chain to EthWarp transaction", "tx", tx.Hash(), "txData", hex.EncodeToString(tx.Data()))
    } else {
        log.Error("Transaction is nil but no error returned")
    }
    Expect(err).Should(BeNil())

    // Wait for the transaction to be mined and get the receipt.
    txReceipt := WaitForTransactionSuccessWithClient(ctx, client, tx.Hash())
    log.Info("DEBUG: Transaction receipt", "txReceipt", txReceipt, "gasUsed", txReceipt.GasUsed)
}
