package types

import (
	corelegacy "cosmossdk.io/core/legacy"
	"cosmossdk.io/core/registry"
	coretransaction "cosmossdk.io/core/transaction"

	"github.com/T-ragon/cosmos-sdk/v3/codec/legacy"
	"github.com/T-ragon/cosmos-sdk/v3/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc corelegacy.Amino) {
	cdc.RegisterConcrete(Params{}, "cosmos-sdk/x/mint/Params")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/mint/MsgUpdateParams")
}

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registrar registry.InterfaceRegistrar) {
	registrar.RegisterImplementations(
		(*coretransaction.Msg)(nil),
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
