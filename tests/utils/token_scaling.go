package utils

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	erc20tokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	nativetokenhome "github.com/ava-labs/icm-contracts/abi-bindings/go/ictt/TokenHome/NativeTokenHome"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"

	. "github.com/onsi/gomega"
)

// ApplyTokenScaling applies token scaling to the given amount of home tokens.
// Token scaling is applied when sending tokens from the home to the TokenRemote instances.
func ApplyTokenScaling(
	tokenMultiplier *big.Int,
	multiplyOnRemote bool,
	homeTokenAmount *big.Int,
) *big.Int {
	return scaleTokens(tokenMultiplier, multiplyOnRemote, homeTokenAmount, true)
}

// RemoveTokenScaling removes token scaling from the given amount of remote tokens.
// Token scaling is removed when sending tokens from the remote back to the TokenHome instance.
func RemoveTokenScaling(
	tokenMultiplier *big.Int,
	multiplyOnRemote bool,
	remoteTokenAmount *big.Int,
) *big.Int {
	return scaleTokens(tokenMultiplier, multiplyOnRemote, remoteTokenAmount, false)
}

func scaleTokens(
	tokenMultiplier *big.Int,
	multiplyOnRemote bool,
	amount *big.Int,
	isSendToRemote bool,
) *big.Int {
	// Multiply when multiplyOnRemote and isSendToRemote are
	// both true or both false.
	if multiplyOnRemote == isSendToRemote {
		return big.NewInt(0).Mul(amount, tokenMultiplier)
	}

	return big.NewInt(0).Div(amount, tokenMultiplier)
}

// GetScaledAmountFromERC20TokenHome returns the scaled amount of remote tokens that
// will be sent to the remote token transferrer for an amount of home tokens.
func GetScaledAmountFromERC20TokenHome(
	erc20TokenHome *erc20tokenhome.ERC20TokenHome,
	remoteBlockchainID ids.ID,
	remoteAddress common.Address,
	homeTokenAmount *big.Int,
) *big.Int {
	remoteSettings, err := erc20TokenHome.GetRemoteTokenTransferrerSettings(
		&bind.CallOpts{},
		remoteBlockchainID,
		remoteAddress,
	)
	Expect(err).Should(BeNil())

	return ApplyTokenScaling(
		remoteSettings.TokenMultiplier,
		remoteSettings.MultiplyOnRemote,
		homeTokenAmount,
	)
}

// GetScaledAmountFromNativeTokenHome returns the scaled amount of tokens that will be sent to
// the remote token transferrer for corresponding amount of home tokens.
func GetScaledAmountFromNativeTokenHome(
	nativeTokenHome *nativetokenhome.NativeTokenHome,
	remoteBlockchainID ids.ID,
	remoteAddress common.Address,
	amount *big.Int,
) *big.Int {
	remoteSettings, err := nativeTokenHome.GetRemoteTokenTransferrerSettings(
		&bind.CallOpts{},
		remoteBlockchainID,
		remoteAddress,
	)
	Expect(err).Should(BeNil())

	return ApplyTokenScaling(
		remoteSettings.TokenMultiplier,
		remoteSettings.MultiplyOnRemote,
		amount,
	)
}

func calculateCollateralNeeded(
	initialReserveImbalance *big.Int,
	tokenMultiplier *big.Int,
	multiplyOnRemote bool,
) *big.Int {
	collateralNeeded := RemoveTokenScaling(tokenMultiplier, multiplyOnRemote, initialReserveImbalance)

	remainder := big.NewInt(0).Mod(collateralNeeded, tokenMultiplier)
	if multiplyOnRemote && (remainder.Cmp(big.NewInt(0)) != 0) {
		collateralNeeded.Add(collateralNeeded, big.NewInt(1))
	}
	return collateralNeeded
}
