// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterregistry

import (
	"fmt"

	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	"github.com/pkg/errors"
)

var protocolRegistryEntryType abi.Type
var addressType abi.Type

func init() {
	// Create an ABI binding for ProtocolRegistryEntry, defined in TeleporterRegistry.sol
	// abigen does not support ABI bindings for standalone structs, only methods and events,
	// so we must manually keep this up-to-date with the struct defined in the contract.
	var err error
	protocolRegistryEntryType, err = abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "version", Type: "uint256"},
		{Name: "protocolAddress", Type: "address"},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create ProtocolRegistryEntry ABI type: %v", err))
	}

	addressType, err = abi.NewType("address", "", nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create address ABI type: %v", err))
	}
}

// ProtocolRegistryEntry is currently only packed together with the destinationAddress
// in the TeleporterRegistryWarpPayload struct. We still implement the ABIPacker interface for exhaustiveness testing.

func (p *ProtocolRegistryEntry) Pack() ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "protocolRegistryEntry",
			Type: protocolRegistryEntryType,
		},
	}
	return args.Pack(p)
}

func (p *ProtocolRegistryEntry) Unpack(b []byte) error {
	args := abi.Arguments{
		{
			Name: "protocolRegistryEntry",
			Type: protocolRegistryEntryType,
		},
	}
	unpacked, err := args.Unpack(b)
	if err != nil {
		return fmt.Errorf("failed to unpack to Teleporter registry entry with err: %v", err)
	}

	return args.Copy(&p, unpacked)
}

func PackTeleporterRegistryWarpPayload(entry ProtocolRegistryEntry, destinationAddress common.Address) ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "protocolRegistryEntry",
			Type: protocolRegistryEntryType,
		},
		{
			Name: "destinationAddress",
			Type: addressType,
		},
	}
	return args.Pack(entry, destinationAddress)
}

func UnpackTeleporterRegistryWarpPayload(entryBytes []byte) (ProtocolRegistryEntry, common.Address, error) {
	args := abi.Arguments{
		{
			Name: "protocolRegistryEntry",
			Type: protocolRegistryEntryType,
		},
		{
			Name: "destinationAddress",
			Type: addressType,
		},
	}
	unpacked, err := args.Unpack(entryBytes)
	fmt.Println("unpacked: ", unpacked)
	if err != nil {
		return ProtocolRegistryEntry{}, common.Address{},
			fmt.Errorf("failed to unpack to Teleporter registry entry with err: %v", err)
	}
	type teleporterRegistryWarpPayload struct {
		ProtocolRegistryEntry ProtocolRegistryEntry `json:"protocolRegistryEntry"`
		DestinationAddress    common.Address        `json:"destinationAddress"`
	}
	var payload teleporterRegistryWarpPayload
	err = args.Copy(&payload, unpacked)
	if err != nil {
		return ProtocolRegistryEntry{}, common.Address{}, err
	}

	return payload.ProtocolRegistryEntry, payload.DestinationAddress, nil
}

// PackAddProtocolVersion packs input to form a call to the addProtocolVersion function
func PackAddProtocolVersion(messageIndex uint32) ([]byte, error) {
	abi, err := TeleporterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("addProtocolVersion", messageIndex)
}
