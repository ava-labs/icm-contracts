package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	exampleerc20 "github.com/ava-labs/icm-contracts/abi-bindings/go/mocks/ExampleERC20"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/log"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

var (
	ExpectedExampleERC20DeployerBalance = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e10))
)

func DeployExampleERC20(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	source interfaces.L1TestInfo,
) (common.Address, *exampleerc20.ExampleERC20) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	// Deploy Mock ERC20 contract
	address, tx, token, err := exampleerc20.DeployExampleERC20(opts, source.RPCClient)
	Expect(err).Should(BeNil())
	log.Info("Deployed Mock ERC20 contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, source, tx.Hash())

	// Check that the deployer has the expected initial balance
	senderAddress := crypto.PubkeyToAddress(senderKey.PublicKey)
	balance, err := token.BalanceOf(&bind.CallOpts{}, senderAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(ExpectedExampleERC20DeployerBalance))

	return address, token
}

func ERC20Approve(
	ctx context.Context,
	token *exampleerc20.ExampleERC20,
	spender common.Address,
	amount *big.Int,
	source interfaces.L1TestInfo,
	senderKey *ecdsa.PrivateKey,
) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := token.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())
	log.Info("Approved ERC20", "spender", spender.Hex(), "txHash", tx.Hash().Hex())

	WaitForTransactionSuccess(ctx, source, tx.Hash())
}
