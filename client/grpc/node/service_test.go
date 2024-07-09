package node

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/T-ragon/cosmos-sdk/client"
	"github.com/T-ragon/cosmos-sdk/server/config"
	sdk "github.com/T-ragon/cosmos-sdk/types"
)

func TestServiceServer_Config(t *testing.T) {
	defaultCfg := config.DefaultConfig()
	svr := NewQueryServer(client.Context{}, *defaultCfg)
	ctx := sdk.Context{}.WithMinGasPrices(sdk.NewDecCoins(sdk.NewInt64DecCoin("stake", 15)))

	resp, err := svr.Config(ctx, &ConfigRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, ctx.MinGasPrices().String(), resp.MinimumGasPrice)
	require.Equal(t, defaultCfg.HaltHeight, resp.HaltHeight)
}
