// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
)

const (
	ReceiveCrossChainMessageBaseGasCost uint64 = 250_000
	MarkMessageReceiptGasCost           uint64 = 2500
	DecodeMessageGasCostPerByte         uint64 = 35

	BaseFeeFactor        = 2
	MaxPriorityFeePerGas = 2500000000 // 2.5 gwei
)

// CalculateReceiveMessageGasLimit calculates the estimated gas amount used by a single call
// to receiveCrossChainMessage for the given message and validator bit vector. The result amount
// depends on the following:
// - Required gas limit for the message execution
// - The size of the Warp message
// - The size of the Teleporter message included in the Warp message
// - The number of Teleporter receipts
// - Base gas cost for {receiveCrossChainMessage} call
// - The number of validator signatures included in the aggregate signature
func CalculateReceiveMessageGasLimit(
	numSigners int,
	executionRequiredGasLimit *big.Int,
	numPredicateChunks int,
	teleporterMessageSize int,
	teleporterReceiptsCount int,
) (uint64, error) {
	if !executionRequiredGasLimit.IsUint64() {
		return 0, errors.New("required gas limit too high")
	}

	gasAmounts := []uint64{
		executionRequiredGasLimit.Uint64(),
		// The variable gas on message bytes is accounted for both when used in predicate verification
		// and also when used in `getVerifiedWarpMessage`
		uint64(numPredicateChunks) * warp.GasCostPerWarpMessageChunk * 2,
		// Take into the variable gas cost for decoding the Teleporter message
		// and marking the receipts as received.
		uint64(teleporterMessageSize) * DecodeMessageGasCostPerByte,
		uint64(teleporterReceiptsCount) * MarkMessageReceiptGasCost,
		ReceiveCrossChainMessageBaseGasCost,
		uint64(numSigners) * warp.GasCostPerWarpSigner,
		warp.GasCostPerSignatureVerification,
	}

	res := gasAmounts[0]
	var err error
	for i := 1; i < len(gasAmounts); i++ {
		res, err = math.Add64(res, gasAmounts[i])
		if err != nil {
			return 0, err
		}
	}

	return res, nil
}
