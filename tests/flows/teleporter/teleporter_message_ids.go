// // (c) 2024, Ava Labs, Inc. All rights reserved.
// // See the file LICENSE for licensing terms.

package teleporter

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	teleporterutils "github.com/ava-labs/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

// Tests Teleporter message ID calculation
func CalculateMessageID(network *localnetwork.LocalAvalancheNetwork, teleporter utils.TeleporterTestInfo) {
	l1Info := network.GetPrimaryNetworkInfo()
	teleporterContractAddress := teleporter.TeleporterMessengerAddress(l1Info)

	sourceBlockchainID := common.HexToHash("0xabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd")
	destinationBlockchainID := common.HexToHash("0x1234567812345678123456781234567812345678123456781234567812345678")
	nonce := big.NewInt(42)

	expectedMessageID, err := teleporter.TeleporterMessenger(l1Info).CalculateMessageID(
		&bind.CallOpts{},
		sourceBlockchainID,
		destinationBlockchainID,
		nonce,
	)
	Expect(err).Should(BeNil())

	calculatedMessageID, err := teleporterutils.CalculateMessageID(
		teleporterContractAddress,
		ids.ID(sourceBlockchainID),
		ids.ID(destinationBlockchainID),
		nonce,
	)
	Expect(err).Should(BeNil())
	Expect(ids.ID(expectedMessageID)).Should(Equal(calculatedMessageID))
}
