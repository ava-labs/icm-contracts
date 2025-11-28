// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

var teleporterMessengerAddress = common.HexToAddress("0xfeabb3b3f4eeae6b5769507a5e6b808704e5c626")

func TestCalculateMessageID(t *testing.T) {
	testCases := []struct {
		name                       string
		teleporterMessengerAddress common.Address
		sourceBlockchainID         string
		destinationBlockchainID    string
		nonce                      int64
		expectedID                 string
		expectedError              bool
	}{
		{
			name:                       "success1",
			teleporterMessengerAddress: teleporterMessengerAddress,
			sourceBlockchainID:         "2D8RG4UpSXbPbvPCAWppNJyqTG2i2CAXSkTgmTBBvs7GKNZjsY",
			destinationBlockchainID:    "yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp",
			nonce:                      1,
			expectedID:                 "jYP9LS4Hsz6GykPxVrucf6dUZRfXHLRpjpzLxyccjBoJr1dom",
			expectedError:              false,
		},
		{
			name:                       "success2",
			teleporterMessengerAddress: teleporterMessengerAddress,
			sourceBlockchainID:         "2D8RG4UpSXbPbvPCAWppNJyqTG2i2CAXSkTgmTBBvs7GKNZjsY",
			destinationBlockchainID:    "yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp",
			nonce:                      2,
			expectedID:                 "2gzmsJRZ3H2NShgkGNUnf2889a1bbVEZwfZngaRUWNcC7RWr5w",
			expectedError:              false,
		},
		{
			name:                       "success3",
			teleporterMessengerAddress: teleporterMessengerAddress,
			sourceBlockchainID:         "2D8RG4UpSXbPbvPCAWppNJyqTG2i2CAXSkTgmTBBvs7GKNZjsY",
			destinationBlockchainID:    "yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp",
			nonce:                      3,
			expectedID:                 "MCV4YGY9FfdsVUT5mo66zHANm9NNBTeYRqrpASQNCdEP7gKbq",
			expectedError:              false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			sourceID, err := ids.FromString(test.sourceBlockchainID)
			require.NoError(t, err)
			destinationID, err := ids.FromString(test.destinationBlockchainID)
			require.NoError(t, err)
			id, err := CalculateMessageID(
				sourceID,
				destinationID,
				big.NewInt(test.nonce),
			)
			if (err != nil) != test.expectedError {
				t.Fatalf("expected error to be %v but got %v", test.expectedError, err)
			}
			if id.String() != test.expectedID {
				t.Fatalf("expected id to be %v but got %v", test.expectedID, id)
			}
		})
	}
}
