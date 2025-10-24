// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"errors"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/subnet-evm/precompile/precompileconfig"
)

const (
	MarkMessageReceiptGasCost   uint64 = 2500
	DecodeMessageGasCostPerByte uint64 = 35
	TeleporterOverheadGasCost   uint64 = 1000_000

	BaseFeeFactor        = 2
	MaxPriorityFeePerGas = 2500000000 // 2.5 gwei
)

var _ precompileconfig.Rules = &UpgradeRules{}

type UpgradeRules struct {
	UpgradeConfig upgrade.Config
}

func (u *UpgradeRules) IsGraniteActivated() bool {
	return u.UpgradeConfig.IsGraniteActivated(time.Now())
}

// CalculateReceiveMessageGasLimit calculates the estimated gas amount used by a single call
// to Teleporter receiveCrossChainMessage for the given message and validator bit vector. The result amount
// depends on the following:
// - Required gas limit for the message execution
// - The size of the Warp message
// - The size of the Teleporter message included in the Warp message
// - The number of Teleporter receipts
// - Base gas cost for {receiveCrossChainMessage} call
// - The number of validator signatures included in the aggregate signature
// TODO: Benchmark to confirm that gas limits estimates are accurate.
// specifically confirm that numTeleporterMessageBytes and TeleporterOverheadGasCost are correct.
func CalculateReceiveMessageGasLimit(
	rules precompileconfig.Rules,
	numSigners int,
	executionRequiredGasLimit *big.Int,
	numPredicateChunks int,
	numTeleporterMessageBytes int,
	teleporterReceiptsCount int,
) (uint64, error) {
	if !executionRequiredGasLimit.IsUint64() {
		return 0, errors.New("required gas limit too high")
	}

	gasConfig := warp.CurrentGasConfig(rules)

	gasAmounts := []uint64{
		executionRequiredGasLimit.Uint64(),
		// The variable gas on message bytes is accounted for both when used in predicate verification
		// and also when used in `getVerifiedWarpMessage`
		uint64(numPredicateChunks) * gasConfig.PerWarpMessageChunk * 2,
		// Take into the variable gas cost for decoding the Teleporter message
		// and marking the receipts as received.
		uint64(numTeleporterMessageBytes) * DecodeMessageGasCostPerByte,
		uint64(teleporterReceiptsCount) * MarkMessageReceiptGasCost,
		uint64(numSigners) * gasConfig.PerWarpSigner,
		gasConfig.VerifyPredicateBase,
		gasConfig.GetVerifiedWarpMessageBase,
		TeleporterOverheadGasCost,
	}

	res := gasAmounts[0]
	var err error
	for i := 1; i < len(gasAmounts); i++ {
		res, err = math.Add(res, gasAmounts[i])
		if err != nil {
			return 0, err
		}
	}

	return res, nil
}
