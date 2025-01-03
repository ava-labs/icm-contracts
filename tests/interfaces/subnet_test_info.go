package interfaces

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/ethclient"
)

// Tracks information about a test L1 used for executing tests against.
type L1TestInfo struct {
	SubnetID                     ids.ID
	BlockchainID                 ids.ID
	NodeURIs                     []string
	WSClient                     ethclient.Client
	RPCClient                    ethclient.Client
	EVMChainID                   *big.Int
	RequirePrimaryNetworkSigners bool
}
