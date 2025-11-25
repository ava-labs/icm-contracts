package network

import (
	"crypto/ecdsa"

	"github.com/ava-labs/libevm/common"
)

type LocalNetwork interface {
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)
	TearDownNetwork()
}
