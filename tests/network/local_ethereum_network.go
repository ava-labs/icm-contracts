package network

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/log"
	"github.com/ava-labs/subnet-evm/ethclient"
	. "github.com/onsi/gomega"
)

var _ LocalNetwork = &LocalEthereumNetwork{}

type LocalEthereumNetwork struct {
	BaseURL         string
	Client          ethclient.Client
	ChainID         *big.Int
	globalFundedKey *secp256k1.PrivateKey
}

const (
	localEthereumNetworkBaseURL   = "http://127.0.0.1:5050"
	localEthereumNetworkFundedKey = "764A4A322753120B4667A20B6309CB5EC754A22BDBCBD62398BE8F803B255337"
)

func NewLocalEthereumNetwork(ctx context.Context) *LocalEthereumNetwork {
	// The local Ethereum network must already be running and accessible at a known endpoint.
	// We test that it is accessible here.
	client, err := ethclient.Dial(localEthereumNetworkBaseURL)
	Expect(err).Should(BeNil())

	// Get the chain ID of the local Ethereum network
	chainID, err := client.ChainID(ctx)
	Expect(err).Should(BeNil())

	fundedKeyBytes, err := hex.DecodeString(localEthereumNetworkFundedKey)
	Expect(err).Should(BeNil())
	globalFundedKey, err := secp256k1.ToPrivateKey(fundedKeyBytes)
	Expect(err).Should(BeNil())

	return &LocalEthereumNetwork{
		BaseURL:         localEthereumNetworkBaseURL,
		Client:          client,
		ChainID:         chainID,
		globalFundedKey: globalFundedKey,
	}
}

func (n *LocalEthereumNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	ecdsaKey := n.globalFundedKey.ToECDSA()
	fundedAddress := crypto.PubkeyToAddress(ecdsaKey.PublicKey)
	return fundedAddress, ecdsaKey
}

func (n *LocalEthereumNetwork) TearDownNetwork() {
	log.Info("Tearing down local Ethereum network")
	Expect(n).ShouldNot(BeNil())
	Expect(n.Client).ShouldNot(BeNil())
}
