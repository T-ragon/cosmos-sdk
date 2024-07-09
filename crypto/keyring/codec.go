package keyring

import (
	"github.com/T-ragon/cosmos-sdk/v3/codec"
	"github.com/T-ragon/cosmos-sdk/v3/codec/legacy"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/hd"
)

func init() {
	RegisterLegacyAminoCodec(legacy.Cdc)
}

// RegisterLegacyAminoCodec registers concrete types and interfaces on the given codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*LegacyInfo)(nil), nil)
	cdc.RegisterConcrete(hd.BIP44Params{}, "crypto/keys/hd/BIP44Params")
	cdc.RegisterConcrete(legacyLocalInfo{}, "crypto/keys/localInfo")
	cdc.RegisterConcrete(legacyLedgerInfo{}, "crypto/keys/ledgerInfo")
	cdc.RegisterConcrete(legacyOfflineInfo{}, "crypto/keys/offlineInfo")
	cdc.RegisterConcrete(LegacyMultiInfo{}, "crypto/keys/multiInfo")
}
