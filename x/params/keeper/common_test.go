package keeper_test

import (
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/params"
	paramskeeper "cosmossdk.io/x/params/keeper"

	"github.com/T-ragon/cosmos-sdk/v3/codec"
	codectestutil "github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	sdktestutil "github.com/T-ragon/cosmos-sdk/v3/testutil"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	moduletestutil "github.com/T-ragon/cosmos-sdk/v3/types/module/testutil"
)

func testComponents() (*codec.LegacyAmino, sdk.Context, storetypes.StoreKey, storetypes.StoreKey, paramskeeper.Keeper) {
	encodingConfig := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, params.AppModule{})
	cdc := encodingConfig.Codec

	legacyAmino := createTestCodec()
	mkey := storetypes.NewKVStoreKey("test")
	tkey := storetypes.NewTransientStoreKey("transient_test")
	ctx := sdktestutil.DefaultContext(mkey, tkey)
	keeper := paramskeeper.NewKeeper(cdc, legacyAmino, mkey, tkey)

	return legacyAmino, ctx, mkey, tkey, keeper
}

type invalid struct{}

type s struct {
	I int
}

func createTestCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cdc.RegisterConcrete(s{}, "test/s")
	cdc.RegisterConcrete(invalid{}, "test/invalid")
	return cdc
}
