// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// Event is a Teleporter log event
type Event uint8

const (
	Unknown Event = iota
	SendCrossChainMessage
	ReceiveCrossChainMessage
	AddFeeAmount
	MessageExecutionFailed
	MessageExecuted
	RelayerRewardsRedeemed
	ReceiptReceived

	sendCrossChainMessageStr    = "SendCrossChainMessage"
	receiveCrossChainMessageStr = "ReceiveCrossChainMessage"
	addFeeAmountStr             = "AddFeeAmount"
	messageExecutionFailedStr   = "MessageExecutionFailed"
	messageExecutedStr          = "MessageExecuted"
	relayerRewardsRedeemedStr   = "RelayerRewardsRedeemed"
	receiptReceivedStr          = "ReceiptReceived"
	unknownStr                  = "Unknown"
)

// String returns the string representation of an Event
func (e Event) String() string {
	switch e {
	case SendCrossChainMessage:
		return sendCrossChainMessageStr
	case ReceiveCrossChainMessage:
		return receiveCrossChainMessageStr
	case AddFeeAmount:
		return addFeeAmountStr
	case MessageExecutionFailed:
		return messageExecutionFailedStr
	case MessageExecuted:
		return messageExecutedStr
	case RelayerRewardsRedeemed:
		return relayerRewardsRedeemedStr
	case ReceiptReceived:
		return receiptReceivedStr
	default:
		return unknownStr
	}
}

// ToEvent converts a string to an Event
func ToEvent(e string) (Event, error) {
	switch strings.ToLower(e) {
	case strings.ToLower(sendCrossChainMessageStr):
		return SendCrossChainMessage, nil
	case strings.ToLower(receiveCrossChainMessageStr):
		return ReceiveCrossChainMessage, nil
	case strings.ToLower(addFeeAmountStr):
		return AddFeeAmount, nil
	case strings.ToLower(messageExecutionFailedStr):
		return MessageExecutionFailed, nil
	case strings.ToLower(messageExecutedStr):
		return MessageExecuted, nil
	case strings.ToLower(relayerRewardsRedeemedStr):
		return RelayerRewardsRedeemed, nil
	case strings.ToLower(receiptReceivedStr):
		return ReceiptReceived, nil
	default:
		return Unknown, fmt.Errorf("unknown event %s", e)
	}
}

// FilterTeleporterEvents parses the topics and data of a Teleporter log into the corresponding Teleporter event
func FilterTeleporterEvents(topics []common.Hash, data []byte, event string) (fmt.Stringer, error) {
	e, err := ToEvent(event)
	if err != nil {
		return nil, err
	}
	var out fmt.Stringer
	switch e {
	case SendCrossChainMessage:
		out = new(TeleporterMessengerSendCrossChainMessage)
	case ReceiveCrossChainMessage:
		out = new(TeleporterMessengerReceiveCrossChainMessage)
	case AddFeeAmount:
		out = new(TeleporterMessengerAddFeeAmount)
	case MessageExecutionFailed:
		out = new(TeleporterMessengerMessageExecutionFailed)
	case MessageExecuted:
		out = new(TeleporterMessengerMessageExecuted)
	case RelayerRewardsRedeemed:
		out = new(TeleporterMessengerRelayerRewardsRedeemed)
	case ReceiptReceived:
		out = new(TeleporterMessengerReceiptReceived)
	default:
		return nil, fmt.Errorf("unknown event %s", e.String())
	}
	if err := UnpackEvent(out, e.String(), topics, data); err != nil {
		return nil, err
	}
	return out, nil
}

func (t TeleporterMessengerSendCrossChainMessage) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerSendCrossChainMessage{
		MessageID:               common.Hash(t.MessageID),
		DestinationBlockchainID: ids.ID(t.DestinationBlockchainID),
		Message:                 toReadableTeleporterMessage(t.Message),
		FeeInfo:                 t.FeeInfo,
		Raw:                     t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerSendCrossChainMessage struct {
	MessageID               common.Hash
	DestinationBlockchainID ids.ID
	Message                 ReadableTeleporterMessage
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log
}

func (t TeleporterMessengerReceiveCrossChainMessage) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerReceiveCrossChainMessage{
		MessageID:          common.Hash(t.MessageID),
		SourceBlockchainID: ids.ID(t.SourceBlockchainID),
		Deliverer:          t.Deliverer,
		RewardRedeemer:     t.RewardRedeemer,
		Message:            toReadableTeleporterMessage(t.Message),
		Raw:                t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerReceiveCrossChainMessage struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            ReadableTeleporterMessage
	Raw                types.Log
}

func (t TeleporterMessengerAddFeeAmount) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerAddFeeAmount{
		MessageID:      common.Hash(t.MessageID),
		UpdatedFeeInfo: t.UpdatedFeeInfo,
		Raw:            t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerAddFeeAmount struct {
	MessageID      common.Hash
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log
}

func (t TeleporterMessengerMessageExecutionFailed) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerMessageExecutionFailed{
		MessageID:          common.Hash(t.MessageID),
		SourceBlockchainID: ids.ID(t.SourceBlockchainID),
		Message:            toReadableTeleporterMessage(t.Message),
		Raw:                t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerMessageExecutionFailed struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Message            ReadableTeleporterMessage
	Raw                types.Log
}

func (t TeleporterMessengerMessageExecuted) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerMessageExecuted{
		MessageID:          common.Hash(t.MessageID),
		SourceBlockchainID: ids.ID(t.SourceBlockchainID),
		Raw:                t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerMessageExecuted struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Raw                types.Log
}

func (t TeleporterMessengerRelayerRewardsRedeemed) String() string {
	outJson, _ := json.MarshalIndent(t, "", "  ")

	return string(outJson)
}

func (t TeleporterMessengerReceiptReceived) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerReceiptReceived{
		MessageID:               common.Hash(t.MessageID),
		DestinationBlockchainID: ids.ID(t.DestinationBlockchainID),
		RelayerRewardAddress:    t.RelayerRewardAddress,
		FeeInfo:                 t.FeeInfo,
		Raw:                     t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerReceiptReceived struct {
	MessageID               common.Hash
	DestinationBlockchainID ids.ID
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log
}

func toReadableTeleporterMessage(t TeleporterMessage) ReadableTeleporterMessage {
	return ReadableTeleporterMessage{
		MessageNonce:            t.MessageNonce,
		OriginSenderAddress:     t.OriginSenderAddress,
		DestinationBlockchainID: ids.ID(t.DestinationBlockchainID),
		DestinationAddress:      t.DestinationAddress,
		RequiredGasLimit:        t.RequiredGasLimit,
		AllowedRelayerAddresses: t.AllowedRelayerAddresses,
		Receipts:                t.Receipts,
		Message:                 t.Message,
	}
}

func (t TeleporterMessage) String() string {
	outJson, _ := json.MarshalIndent(toReadableTeleporterMessage(t), "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessage struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	DestinationBlockchainID ids.ID
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}
