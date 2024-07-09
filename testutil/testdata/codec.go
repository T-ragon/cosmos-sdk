package testdata

import (
	"github.com/tendermint/go-amino"

	"github.com/T-ragon/cosmos-sdk/v3/codec/types"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	"github.com/T-ragon/cosmos-sdk/v3/types/msgservice"
	"github.com/T-ragon/cosmos-sdk/v3/types/tx"
)

func NewTestInterfaceRegistry() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	RegisterInterfaces(registry)
	return registry
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &TestMsg{})

	registry.RegisterInterface("Animal", (*Animal)(nil))
	registry.RegisterImplementations(
		(*Animal)(nil),
		&Dog{},
		&Cat{},
	)
	registry.RegisterImplementations(
		(*HasAnimalI)(nil),
		&HasAnimal{},
	)
	registry.RegisterImplementations(
		(*HasHasAnimalI)(nil),
		&HasHasAnimal{},
	)
	registry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&Cat{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func NewTestAmino() *amino.Codec {
	cdc := amino.NewCodec()
	cdc.RegisterInterface((*Animal)(nil), nil)
	cdc.RegisterConcrete(&Dog{}, "test/Dog", nil)
	cdc.RegisterConcrete(&Cat{}, "test/Cat", nil)

	cdc.RegisterInterface((*HasAnimalI)(nil), nil)
	cdc.RegisterConcrete(&HasAnimal{}, "test/HasAnimal", nil)

	cdc.RegisterInterface((*HasHasAnimalI)(nil), nil)
	cdc.RegisterConcrete(&HasHasAnimal{}, "test/HasHasAnimal", nil)

	return cdc
}
