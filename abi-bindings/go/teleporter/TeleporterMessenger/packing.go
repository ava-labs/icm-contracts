// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/pkg/errors"
)

var teleporterMessageType abi.Type

func init() {
	// Create an ABI binding for TeleporterMessage, defined in ITeleporterMessenger.sol
	// abigen does not support ABI bindings for standalone structs, only methods and events,
	// so we must manually keep this up-to-date with the struct defined in the contract.
	var err error
	teleporterMessageType, err = abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "messageNonce", Type: "uint256"},
		{Name: "originSenderAddress", Type: "address"},
		{Name: "destinationBlockchainID", Type: "bytes32"},
		{Name: "destinationAddress", Type: "address"},
		{Name: "requiredGasLimit", Type: "uint256"},
		{Name: "allowedRelayerAddresses", Type: "address[]"},
		{Name: "receipts", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "receivedMessageNonce", Type: "uint256"},
			{Name: "relayerRewardAddress", Type: "address"},
		}},
		{Name: "message", Type: "bytes"},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create TeleporterMessage ABI type: %v", err))
	}
}

func (m *TeleporterMessage) Pack() ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "teleporterMessage",
			Type: teleporterMessageType,
		},
	}
	return args.Pack(m)
}

func (m *TeleporterMessage) Unpack(b []byte) error {
	args := abi.Arguments{
		{
			Name: "teleporterMessage",
			Type: teleporterMessageType,
		},
	}
	unpacked, err := args.Unpack(b)
	if err != nil {
		return fmt.Errorf("failed to unpack to teleporter message with err: %v", err)
	}
	return args.Copy(&m, unpacked)
}

func PackSendCrossChainMessage(input TeleporterMessageInput) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("sendCrossChainMessage", input)
}

func PackRetryMessageExecution(sourceBlockchainID ids.ID, message TeleporterMessage) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("retryMessageExecution", sourceBlockchainID, message)
}

// PackReceiveCrossChainMessage packs a ReceiveCrossChainMessageInput to form
// a call to the receiveCrossChainMessage function
func PackReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// PackReceiveInterChainMessage packs a ReceiveInterChainMessageInput to form
// a call to the receiveInterChainMessage function
func PackReceiveInterChainMessage(messagePayload *ICMMessage, relayerRewardAddress common.Address) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("receiveInterChainMessage", messagePayload, relayerRewardAddress)
}

// PackCalculateMessageID packs input to form a call to the calculateMessageID function
func PackCalculateMessageID(
	sourceBlockchainID [32]byte,
	destinationBlockchainID [32]byte,
	nonce *big.Int,
) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("calculateMessageID", sourceBlockchainID, destinationBlockchainID, nonce)
}

func PackCalculateMessageIDOutput(messageID [32]byte) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.PackOutput("calculateMessageID", messageID)
}

// PackMessageReceived packs a MessageReceivedInput to form a call to the messageReceived function
func PackMessageReceived(messageID [32]byte) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}
	return abi.Pack("messageReceived", messageID)
}

// UnpackMessageReceivedResult attempts to unpack result bytes to a bool indicating whether the message was received
func UnpackMessageReceivedResult(result []byte) (bool, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return false, errors.Wrap(err, "failed to get abi")
	}

	var success bool
	err = abi.UnpackIntoInterface(&success, "messageReceived", result)
	return success, err
}

func PackMessageReceivedOutput(success bool) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.PackOutput("messageReceived", success)
}

// UnpackEvent unpacks the event data and topics into the provided interface
func UnpackEvent(out interface{}, event string, topics []common.Hash, data []byte) error {
	teleporterABI, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get abi: %v", err)
	}
	if len(data) > 0 {
		if err := teleporterABI.UnpackIntoInterface(out, event, data); err != nil {
			return err
		}
	}

	var indexed abi.Arguments
	for _, arg := range teleporterABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, topics[1:])
}
