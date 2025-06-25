// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterregistry

import (
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

func TestPackUnpackProtocolRegistryEntry(t *testing.T) {
	entry := ProtocolRegistryEntry{
		Version:         big.NewInt(1),
		ProtocolAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
	}
	e, err := entry.Pack()
	require.NoError(t, err)

	d := ProtocolRegistryEntry{}
	err = d.Unpack(e)
	require.NoError(t, err)
	require.Equal(t, entry, d)

	destinationAddress := common.HexToAddress("0x0123456789abcdef0123456789abcdef01234568")

	b, err := PackTeleporterRegistryWarpPayload(entry, destinationAddress)
	require.NoError(t, err)

	// Confirm that packing of the TeleporterRegistryWarpPayload is prefixed with the ProtocolRegistryEntry packed bytes.
	require.Equal(t, e, b[:len(e)])

	unpackedEntry, unpackedDestinationAddress, err := UnpackTeleporterRegistryWarpPayload(b)
	require.NoError(t, err)

	require.Equal(t, entry.Version, unpackedEntry.Version)
	require.Equal(t, entry.ProtocolAddress, unpackedEntry.ProtocolAddress)
	require.Equal(t, destinationAddress, unpackedDestinationAddress)
}
