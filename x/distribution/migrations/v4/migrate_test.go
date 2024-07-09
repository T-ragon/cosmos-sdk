package v4_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/core/log"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/distribution"
	v4 "cosmossdk.io/x/distribution/migrations/v4"

	codectestutil "github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/keys/secp256k1"
	"github.com/T-ragon/cosmos-sdk/v3/runtime"
	"github.com/T-ragon/cosmos-sdk/v3/testutil"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	moduletestutil "github.com/T-ragon/cosmos-sdk/v3/types/module/testutil"
)

func TestMigration(t *testing.T) {
	cdc := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, distribution.AppModule{}).Codec
	storeKey := storetypes.NewKVStoreKey("distribution")
	storeService := runtime.NewKVStoreService(storeKey)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	env := runtime.NewEnvironment(storeService, log.NewNopLogger())

	addr1 := secp256k1.GenPrivKey().PubKey().Address()
	consAddr1 := sdk.ConsAddress(addr1)

	// Set and check the previous proposer
	err := v4.SetPreviousProposerConsAddr(ctx, storeService, cdc, consAddr1)
	require.NoError(t, err)

	gotAddr, err := v4.GetPreviousProposerConsAddr(ctx, storeService, cdc)
	require.NoError(t, err)
	require.Equal(t, consAddr1, gotAddr)

	err = v4.MigrateStore(ctx, env, cdc)
	require.NoError(t, err)

	// Check that the previous proposer has been removed
	_, err = v4.GetPreviousProposerConsAddr(ctx, storeService, cdc)
	require.ErrorContains(t, err, "previous proposer not set")
}
