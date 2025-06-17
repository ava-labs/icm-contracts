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

func createTestTeleporterMessage(messageNonce *big.Int) TeleporterMessage {
	m := TeleporterMessage{
		MessageNonce:            messageNonce,
		OriginSenderAddress:     common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		DestinationBlockchainID: ids.ID{1, 2, 3, 4},
		DestinationAddress:      common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		RequiredGasLimit:        big.NewInt(2),
		AllowedRelayerAddresses: []common.Address{
			common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		},
		Receipts: []TeleporterMessageReceipt{
			{
				ReceivedMessageNonce: big.NewInt(1),
				RelayerRewardAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
			},
		},
		Message: []byte{1, 2, 3, 4},
	}
	return m
}

func TestUnpackEvent(t *testing.T) {
	mockBlockchainID := ids.ID{1, 2, 3, 4}
	mockMessageNonce := big.NewInt(5)
	mockMessageID := common.Hash{9, 10, 11, 12}
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
			out      interface{}
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
				out: new(TeleporterMessengerSendCrossChainMessage),
				expected: &TeleporterMessengerSendCrossChainMessage{
					DestinationBlockchainID: mockBlockchainID,
					MessageID:               mockMessageID,
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
				out: new(TeleporterMessengerReceiveCrossChainMessage),
				expected: &TeleporterMessengerReceiveCrossChainMessage{
					SourceBlockchainID: mockBlockchainID,
					MessageID:          mockMessageID,
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
				out: new(TeleporterMessengerMessageExecuted),
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

			err = UnpackEvent(test.out, test.event.String(), topics, data)
			require.NoError(t, err)

			require.Equal(t, test.expected, test.out)
		})
	}
}
