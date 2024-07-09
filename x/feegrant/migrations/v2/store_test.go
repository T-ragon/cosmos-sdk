package v2_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/core/header"
	"cosmossdk.io/core/log"
	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/feegrant"
	v2 "cosmossdk.io/x/feegrant/migrations/v2"
	"cosmossdk.io/x/feegrant/module"

	addresscodec "github.com/T-ragon/cosmos-sdk/v3/codec/address"
	codectestutil "github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/keys/ed25519"
	"github.com/T-ragon/cosmos-sdk/v3/runtime"
	"github.com/T-ragon/cosmos-sdk/v3/testutil"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	moduletestutil "github.com/T-ragon/cosmos-sdk/v3/types/module/testutil"
)

func TestMigration(t *testing.T) {
	encodingConfig := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, module.AppModule{})
	cdc := encodingConfig.Codec
	ac := addresscodec.NewBech32Codec("cosmos")

	feegrantKey := storetypes.NewKVStoreKey(v2.ModuleName)
	ctx := testutil.DefaultContext(feegrantKey, storetypes.NewTransientStoreKey("transient_test"))
	granter1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	grantee1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	granter2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	grantee2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())

	spendLimit := sdk.NewCoins(sdk.NewCoin("stake", sdkmath.NewInt(1000)))
	now := ctx.HeaderInfo().Time
	oneDay := now.AddDate(0, 0, 1)
	twoDays := now.AddDate(0, 0, 2)

	grants := []struct {
		granter    sdk.AccAddress
		grantee    sdk.AccAddress
		spendLimit sdk.Coins
		expiration *time.Time
	}{
		{
			granter:    granter1,
			grantee:    grantee1,
			spendLimit: spendLimit,
			expiration: &twoDays,
		},
		{
			granter:    granter2,
			grantee:    grantee2,
			spendLimit: spendLimit,
			expiration: &oneDay,
		},
		{
			granter:    granter1,
			grantee:    grantee2,
			spendLimit: spendLimit,
		},
		{
			granter:    granter2,
			grantee:    grantee1,
			expiration: &oneDay,
		},
	}

	store := ctx.KVStore(feegrantKey)
	for _, grant := range grants {
		granterStr, err := ac.BytesToString(grant.granter)
		require.NoError(t, err)
		granteeStr, err := ac.BytesToString(grant.grantee)
		require.NoError(t, err)
		newGrant, err := feegrant.NewGrant(granterStr, granteeStr, &feegrant.BasicAllowance{
			SpendLimit: grant.spendLimit,
			Expiration: grant.expiration,
		})
		require.NoError(t, err)

		bz, err := cdc.Marshal(&newGrant)
		require.NoError(t, err)

		store.Set(v2.FeeAllowanceKey(grant.granter, grant.grantee), bz)
	}

	ctx = ctx.WithHeaderInfo(header.Info{Time: now.Add(30 * time.Hour)})
	require.NoError(t, v2.MigrateStore(ctx, runtime.NewEnvironment(runtime.NewKVStoreService(feegrantKey), log.NewNopLogger()), cdc))
	store = ctx.KVStore(feegrantKey)

	require.NotNil(t, store.Get(v2.FeeAllowanceKey(granter1, grantee1)))
	require.Nil(t, store.Get(v2.FeeAllowanceKey(granter2, grantee2)))
	require.NotNil(t, store.Get(v2.FeeAllowanceKey(granter1, grantee2)))
	require.Nil(t, store.Get(v2.FeeAllowanceKey(granter2, grantee1)))
}
