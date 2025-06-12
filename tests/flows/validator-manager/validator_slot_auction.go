package staking

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	slotauctionManager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/SlotAuctionManager"
	slotauctionmanager "github.com/ava-labs/icm-contracts/abi-bindings/go/validator-manager/SlotAuctionManager"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

/*
 * Registers a native token staking validator on a L1. The steps are as follows:
 * - Deploy the NativeTokenStakingManager
 * - Initiate validator registration
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is registered in the staking contract
 *
 * Delists the validator from the L1. The steps are as follows:
 * - Initiate validator delisting
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is delisted from the staking contract
 */
func ValidatorSlotAuction(network *localnetwork.LocalNetwork) {
	// Get the L1s info
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	_, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	nodes, _ := network.ConvertSubnet(
		ctx,
		l1AInfo,
		utils.SlotAuctionManager,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		fundedKey,
		false,
	)

	validatorManagerProxy, _ := network.GetValidatorManager(l1AInfo.SubnetID)

	slotauctionManager, err := slotauctionManager.NewSlotAuctionManager(
		validatorManagerProxy.Address,
		l1AInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	
	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			l1AInfo.SubnetID,
		},
	)
	defer signatureAggregator.Shutdown()

	node := nodes[0]

	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)
	slotauctionManager.BecomeValidator(
		opts, 
		node.NodeID[:],
		node.NodePoP.PublicKey[:],
		slotauctionmanager.PChainOwner{},
		slotauctionmanager.PChainOwner{},
	)
}
