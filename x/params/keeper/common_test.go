package keeper_test

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/T-ragon/cosmos-sdk/codec"
	sdktestutil "github.com/T-ragon/cosmos-sdk/testutil"
	sdk "github.com/T-ragon/cosmos-sdk/types"
	moduletestutil "github.com/T-ragon/cosmos-sdk/types/module/testutil"
	"github.com/T-ragon/cosmos-sdk/x/params"
	paramskeeper "github.com/T-ragon/cosmos-sdk/x/params/keeper"
)

func testComponents() (*codec.LegacyAmino, sdk.Context, storetypes.StoreKey, storetypes.StoreKey, paramskeeper.Keeper) {
	encodingConfig := moduletestutil.MakeTestEncodingConfig(params.AppModuleBasic{})
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
	cdc.RegisterConcrete(s{}, "test/s", nil)
	cdc.RegisterConcrete(invalid{}, "test/invalid", nil)
	return cdc
}
