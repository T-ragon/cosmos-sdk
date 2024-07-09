package std

import (
	"cosmossdk.io/core/legacy"
	"cosmossdk.io/core/registry"

	"github.com/T-ragon/cosmos-sdk/v3/codec"
	cryptocodec "github.com/T-ragon/cosmos-sdk/v3/crypto/codec"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	txtypes "github.com/T-ragon/cosmos-sdk/v3/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc legacy.Amino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry registry.InterfaceRegistrar) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
