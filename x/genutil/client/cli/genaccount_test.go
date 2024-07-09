package cli_test

import (
	"context"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	corectx "cosmossdk.io/core/context"
	"cosmossdk.io/log"
	"cosmossdk.io/x/auth"

	"github.com/T-ragon/cosmos-sdk/v3/client"
	addresscodec "github.com/T-ragon/cosmos-sdk/v3/codec/address"
	codectestutil "github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/hd"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/keyring"
	"github.com/T-ragon/cosmos-sdk/v3/testutil/testdata"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	moduletestutil "github.com/T-ragon/cosmos-sdk/v3/types/module/testutil"
	genutilcli "github.com/T-ragon/cosmos-sdk/v3/x/genutil/client/cli"
	genutiltest "github.com/T-ragon/cosmos-sdk/v3/x/genutil/client/testutil"
)

func TestAddGenesisAccountCmd(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	ac := codectestutil.CodecOptions{}.GetAddressCodec()
	addr1Str, err := ac.BytesToString(addr1)
	require.NoError(t, err)

	tests := []struct {
		name        string
		addr        string
		denom       string
		withKeyring bool
		expectErr   bool
	}{
		{
			name:        "invalid address",
			addr:        "",
			denom:       "1000atom",
			withKeyring: false,
			expectErr:   true,
		},
		{
			name:        "valid address",
			addr:        addr1Str,
			denom:       "1000atom",
			withKeyring: false,
			expectErr:   false,
		},
		{
			name:        "multiple denoms",
			addr:        addr1Str,
			denom:       "1000atom, 2000stake",
			withKeyring: false,
			expectErr:   false,
		},
		{
			name:        "with keyring",
			addr:        "set",
			denom:       "1000atom",
			withKeyring: true,
			expectErr:   false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			home := t.TempDir()
			logger := log.NewNopLogger()
			viper := viper.New()

			appCodec := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, auth.AppModule{}).Codec
			err = genutiltest.ExecInitCmd(testMbm, home, appCodec)
			require.NoError(t, err)

			err := writeAndTrackDefaultConfig(viper, home)
			require.NoError(t, err)
			clientCtx := client.Context{}.WithCodec(appCodec).WithHomeDir(home).WithAddressCodec(ac)

			if tc.withKeyring {
				path := hd.CreateHDPath(118, 0, 0).String()
				kr, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendMemory, home, nil, appCodec)
				require.NoError(t, err)
				_, _, err = kr.NewMnemonic(tc.addr, keyring.English, path, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
				require.NoError(t, err)
				clientCtx = clientCtx.WithKeyring(kr)
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
			ctx = context.WithValue(ctx, corectx.ViperContextKey, viper)
			ctx = context.WithValue(ctx, corectx.LoggerContextKey, logger)

			cmd := genutilcli.AddGenesisAccountCmd(addresscodec.NewBech32Codec("cosmos"))
			cmd.SetArgs([]string{
				tc.addr,
				tc.denom,
			})

			if tc.expectErr {
				require.Error(t, cmd.ExecuteContext(ctx))
			} else {
				require.NoError(t, cmd.ExecuteContext(ctx))
			}
		})
	}
}
