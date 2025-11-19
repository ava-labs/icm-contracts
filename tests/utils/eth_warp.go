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
) (common.Address, *ethwarp.EthWarp) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, chainID)
	Expect(err).Should(BeNil())
	log.Info("Deploying WARP with dest chain")
	address, tx, ethWarp, err := ethwarp.DeployEthWarp(opts, client, chainID)
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
    Expect(err).Should(BeNil())

    log.Info("EthWarp contract registered validator set", "Blockchain ID", avalancheL1BlockchainID, "contract", validatorSetRegistryAddress)
    log.Info("Register chain to EthWarp transaction", "tx", tx.Hash(), "txData", hex.EncodeToString(tx.Data()))

	// Wait for the transaction to be mined and get the receipt.
	txReceipt := WaitForTransactionSuccessWithClient(ctx, client, tx.Hash())
	log.Debug("DEBUG: Transaction receipt", "txReceipt", txReceipt, "gasUsed", txReceipt.GasUsed)
}
