package testutil

import (
	"cosmossdk.io/x/auth/tx"

	"github.com/T-ragon/cosmos-sdk/v3/client"
	"github.com/T-ragon/cosmos-sdk/v3/codec"
	"github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	"github.com/T-ragon/cosmos-sdk/v3/codec/types"
	"github.com/T-ragon/cosmos-sdk/v3/std"
	"github.com/T-ragon/cosmos-sdk/v3/types/module"
)

// TestEncodingConfig defines an encoding configuration that is used for testing
// purposes. Note, MakeTestEncodingConfig takes a series of AppModule types
// which should only contain the relevant module being tested and any potential
// dependencies.
type TestEncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeTestEncodingConfig(codecOpt testutil.CodecOptions, modules ...module.AppModule) TestEncodingConfig {
	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := codecOpt.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	signingCtx := cdc.InterfaceRegistry().SigningContext()

	encCfg := TestEncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             cdc,
		TxConfig:          tx.NewTxConfig(cdc, signingCtx.AddressCodec(), signingCtx.ValidatorAddressCodec(), tx.DefaultSignModes),
		Amino:             aminoCodec,
	}

	mb := module.NewManager(modules...)
	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
	mb.RegisterLegacyAminoCodec(encCfg.Amino)
	mb.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}

func MakeTestTxConfig(codecOpt testutil.CodecOptions) client.TxConfig {
	interfaceRegistry := codecOpt.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	signingCtx := interfaceRegistry.SigningContext()
	return tx.NewTxConfig(cdc, signingCtx.AddressCodec(), signingCtx.ValidatorAddressCodec(), tx.DefaultSignModes)
}

type TestBuilderTxConfig struct {
	client.TxConfig
	TxBuilder *TestTxBuilder
}

func MakeBuilderTestTxConfig(codecOpt testutil.CodecOptions) TestBuilderTxConfig {
	return TestBuilderTxConfig{
		TxConfig: MakeTestTxConfig(codecOpt),
	}
}

func (cfg TestBuilderTxConfig) NewTxBuilder() client.TxBuilder {
	if cfg.TxBuilder == nil {
		cfg.TxBuilder = &TestTxBuilder{
			TxBuilder: cfg.TxConfig.NewTxBuilder(),
		}
	}
	return cfg.TxBuilder
}

type TestTxBuilder struct {
	client.TxBuilder
	ExtOptions []*types.Any
}

func (b *TestTxBuilder) SetExtensionOptions(extOpts ...*types.Any) {
	b.ExtOptions = extOpts
}
