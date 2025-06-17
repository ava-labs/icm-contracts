// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

func TestEventString(t *testing.T) {
	var (
		tests = []struct {
			event Event
			str   string
		}{
			{Unknown, unknownStr},
			{SendCrossChainMessage, sendCrossChainMessageStr},
			{ReceiveCrossChainMessage, receiveCrossChainMessageStr},
			{AddFeeAmount, addFeeAmountStr},
			{MessageExecutionFailed, messageExecutionFailedStr},
			{MessageExecuted, messageExecutedStr},
			{RelayerRewardsRedeemed, relayerRewardsRedeemedStr},
		}
	)

	for _, test := range tests {
		t.Run(test.event.String(), func(t *testing.T) {
			require.Equal(t, test.event.String(), test.str)
		})
	}
}

func TestToEvent(t *testing.T) {
	var (
		tests = []struct {
			str     string
			event   Event
			isError bool
		}{
			{unknownStr, Unknown, true},
			{sendCrossChainMessageStr, SendCrossChainMessage, false},
			{receiveCrossChainMessageStr, ReceiveCrossChainMessage, false},
			{addFeeAmountStr, AddFeeAmount, false},
			{messageExecutionFailedStr, MessageExecutionFailed, false},
			{messageExecutedStr, MessageExecuted, false},
			{relayerRewardsRedeemedStr, RelayerRewardsRedeemed, false},
		}
	)

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			event, err := ToEvent(test.str)
			if test.isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, test.event, event)
		})
	}
}

func TestFilterTeleporterEvents(t *testing.T) {
	mockBlockchainID := ids.ID{1, 2, 3, 4}
	mockMessageNonce := big.NewInt(8)
	mockMessageID := ids.ID{9, 10, 11, 12}
	message := createTestTeleporterMessage(mockMessageNonce)
	feeInfo := TeleporterFeeInfo{
		FeeTokenAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		Amount:          big.NewInt(1),
	}
	deliverer := common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")

	teleporterABI, err := TeleporterMessengerMetaData.GetAbi()
	require.NoError(t, err)

	var (
		tests = []struct {
			event    Event
			args     []interface{}
			expected interface{}
		}{
			{
				event: SendCrossChainMessage,
				args: []interface{}{
					mockMessageID,
					mockBlockchainID,
					message,
					feeInfo,
				},
				expected: &TeleporterMessengerSendCrossChainMessage{
					MessageID:               mockMessageID,
					DestinationBlockchainID: mockBlockchainID,
					Message:                 message,
					FeeInfo:                 feeInfo,
				},
			},
			{
				event: ReceiveCrossChainMessage,
				args: []interface{}{
					mockMessageID,
					mockBlockchainID,
					deliverer,
					deliverer,
					message,
				},
				expected: &TeleporterMessengerReceiveCrossChainMessage{
					MessageID:          mockMessageID,
					SourceBlockchainID: mockBlockchainID,
					Deliverer:          deliverer,
					RewardRedeemer:     deliverer,
					Message:            message,
				},
			},
			{
				event: MessageExecuted,
				args: []interface{}{
					mockMessageID,
					mockBlockchainID,
				},
				expected: &TeleporterMessengerMessageExecuted{
					MessageID:          mockMessageID,
					SourceBlockchainID: mockBlockchainID,
				},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.event.String(), func(t *testing.T) {
			topics, data, err := teleporterABI.PackEvent(test.event.String(), test.args...)
			require.NoError(t, err)

			out, err := FilterTeleporterEvents(topics, data, test.event.String())
			require.NoError(t, err)

			require.Equal(t, test.expected, out)
		})
	}
}
