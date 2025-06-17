// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi"
)

var (
	uint256Ty abi.Type
	bytes32Ty abi.Type
	addressTy abi.Type
)

func init() {
	uint256Ty, _ = abi.NewType("uint256", "uint256", nil)
	bytes32Ty, _ = abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ = abi.NewType("address", "address", nil)
}

func CalculateMessageID(
	teleporterMessengerAddress common.Address,
	sourceBlockchainID ids.ID,
	destinationBlockchainID ids.ID,
	nonce *big.Int,
) (ids.ID, error) {
	arguments := abi.Arguments{
		{Type: addressTy},
		{Type: bytes32Ty},
		{Type: bytes32Ty},
		{Type: uint256Ty},
	}

	bytes, err := arguments.Pack(
		teleporterMessengerAddress,
		sourceBlockchainID,
		destinationBlockchainID,
		nonce,
	)
	if err != nil {
		return ids.ID{}, err
	}

	return ids.ID(crypto.Keccak256Hash(bytes)), nil
}
