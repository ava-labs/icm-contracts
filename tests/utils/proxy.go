package utils

import (
	"context"
	"crypto/ecdsa"

	proxyadmin "github.com/ava-labs/icm-contracts/abi-bindings/go/ProxyAdmin"
	transparentupgradeableproxy "github.com/ava-labs/icm-contracts/abi-bindings/go/TransparentUpgradeableProxy"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

func DeployTransparentUpgradeableProxy(
	ctx context.Context,
	l1 interfaces.L1TestInfo,
	senderKey *ecdsa.PrivateKey,
	implAddress common.Address,
) (common.Address, *proxyadmin.ProxyAdmin) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		l1.EVMChainID,
	)
	Expect(err).Should((BeNil()))

	senderAddress := crypto.PubkeyToAddress(senderKey.PublicKey)
	proxyAddress, tx, proxy, err := transparentupgradeableproxy.DeployTransparentUpgradeableProxy(
		opts,
		l1.RPCClient,
		implAddress,
		senderAddress,
		[]byte{},
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, l1, tx.Hash())
	proxyAdminEvent, err := GetEventFromLogs(receipt.Logs, proxy.ParseAdminChanged)
	Expect(err).Should(BeNil())

	proxyAdmin, err := proxyadmin.NewProxyAdmin(proxyAdminEvent.NewAdmin, l1.RPCClient)
	Expect(err).Should(BeNil())

	return proxyAddress, proxyAdmin
}
