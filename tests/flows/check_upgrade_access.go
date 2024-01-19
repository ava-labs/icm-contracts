package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func CheckUpgradeAccess(network interfaces.Network) {
	subnetInfo := network.GetPrimaryNetworkInfo()
	_, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()
	teleporterAddress := network.GetTeleporterContractAddress()
	_, exampleMessenger := utils.DeployExampleCrossChainMessenger(
		ctx, fundedKey, subnetInfo,
	)

	// Try to call updateMinTeleporterVersion from a non owner account
	nonOwnerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	nonOwnerAddress := crypto.PubkeyToAddress(nonOwnerKey.PublicKey)

	// Transfer native assets to the non owner account
	fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1)) // 10eth
	fundNonOwnerTx := utils.CreateNativeTransferTransaction(
		ctx,
		subnetInfo,
		fundedKey,
		nonOwnerAddress,
		fundAmount)
	utils.SendTransactionAndWaitForSuccess(ctx, subnetInfo, fundNonOwnerTx)

	nonOwnerOpts, err := bind.NewKeyedTransactorWithChainID(
		nonOwnerKey, subnetInfo.EVMChainID)
	Expect(err).Should(BeNil())
	ownerOpts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetInfo.EVMChainID)
	Expect(err).Should(BeNil())

	_, err = exampleMessenger.PauseTeleporterAddress(nonOwnerOpts, teleporterAddress)
	Expect(err).ShouldNot(BeNil())

	// Check that the teleporter address is not paused, because previous call should have failed
	isPaused, err := exampleMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())

	// Try to call pauseTeleporterAddress from the owner account
	tx, err := exampleMessenger.PauseTeleporterAddress(ownerOpts, teleporterAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetInfo, tx)
	isPaused, err = exampleMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Transfer ownership to the non owner account
	tx, err = exampleMessenger.TransferOwnership(ownerOpts, nonOwnerAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetInfo, tx)

	// Try to call unpauseTeleporterAddress from the previous owner account
	_, err = exampleMessenger.UnpauseTeleporterAddress(ownerOpts, teleporterAddress)
	Expect(err).ShouldNot(BeNil())

	// Make sure the teleporter address is still paused
	isPaused, err = exampleMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Try to call unpauseTeleporterAddress from the non owner account now
	tx, err = exampleMessenger.UnpauseTeleporterAddress(nonOwnerOpts, teleporterAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetInfo, tx)
	isPaused, err = exampleMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())
}
