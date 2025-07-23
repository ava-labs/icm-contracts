package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	avalanchevalidatorsetregistry "github.com/ava-labs/icm-contracts/abi-bindings/go/AvalancheValidatorSetRegistry"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	. "github.com/onsi/gomega"
)

func DeployAvalancheValidatorSetRegistry(
	ctx context.Context,
	fundedKey *ecdsa.PrivateKey,
	chainID *big.Int,
	client ethclient.Client,
	avalancheNetworkID uint32,
) (common.Address, *avalanchevalidatorsetregistry.AvalancheValidatorSetRegistry) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, chainID)
	Expect(err).Should(BeNil())

	address, tx, validatorSetRegistry, err := avalanchevalidatorsetregistry.DeployAvalancheValidatorSetRegistry(opts, client, avalancheNetworkID)
	Expect(err).Should(BeNil())

	WaitForTransactionSuccessWithClient(ctx, client, tx.Hash())

	return address, validatorSetRegistry
}
