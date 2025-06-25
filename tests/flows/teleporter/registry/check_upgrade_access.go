package registry

import (
	"context"
	"math/big"

	"github.com/ava-labs/icm-contracts/tests/flows"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"

	. "github.com/onsi/gomega"
)

func CheckUpgradeAccess(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	l1Info := network.GetPrimaryNetworkInfo()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy TestMessenger to the L1
	//
	ctx := context.Background()
	teleporterAddress := teleporter.TeleporterMessengerAddress(l1Info)
	_, testMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(l1Info),
		l1Info,
	)

	// Check that owner is the funded address
	owner, err := testMessenger.Owner(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(owner).Should(Equal(fundedAddress))

	// Try to call updateMinTeleporterVersion from a non owner account
	nonOwnerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	nonOwnerAddress := crypto.PubkeyToAddress(nonOwnerKey.PublicKey)

	// Transfer native assets to the non owner account
	fundAmount := big.NewInt(0.1e18) // 0.1avax
	utils.SendNativeTransfer(
		ctx,
		l1Info,
		fundedKey,
		nonOwnerAddress,
		fundAmount,
	)

	// Check that access is not granted to the non owner and has no effect
	nonOwnerOpts, err := bind.NewKeyedTransactorWithChainID(
		nonOwnerKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())
	_, err = testMessenger.PauseTeleporterAddress(nonOwnerOpts, teleporterAddress)
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(ContainSubstring(flows.ErrTxReverted))

	// Check that the teleporter address is not paused, because previous call should have failed
	isPaused, err := testMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())

	// Check that the owner is able to pause the Teleporter address
	ownerOpts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())
	// Try to call pauseTeleporterAddress from the owner account
	tx, err := testMessenger.PauseTeleporterAddress(ownerOpts, teleporterAddress)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, l1Info, tx.Hash())
	pauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessenger.ParseTeleporterAddressPaused)
	Expect(err).Should(BeNil())
	Expect(pauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err = testMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Transfer ownership to the non owner account
	tx, err = testMessenger.TransferOwnership(ownerOpts, nonOwnerAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info, tx.Hash())

	// Try to call unpauseTeleporterAddress from the previous owner account
	_, err = testMessenger.UnpauseTeleporterAddress(ownerOpts, teleporterAddress)
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(ContainSubstring(flows.ErrTxReverted))

	// Make sure the teleporter address is still paused
	isPaused, err = testMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Try to call unpauseTeleporterAddress from the non owner account now
	tx, err = testMessenger.UnpauseTeleporterAddress(nonOwnerOpts, teleporterAddress)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, l1Info, tx.Hash())
	unpauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessenger.ParseTeleporterAddressUnpaused)
	Expect(err).Should(BeNil())
	Expect(unpauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err = testMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())
}
