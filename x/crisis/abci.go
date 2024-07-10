package crisis

import (
	"context"

	"github.com/T-ragon/cosmos-sdk/telemetry"
	sdk "github.com/T-ragon/cosmos-sdk/types"
	"github.com/T-ragon/cosmos-sdk/x/crisis/keeper"
	"github.com/T-ragon/cosmos-sdk/x/crisis/types"
)

// check all registered invariants
func EndBlocker(ctx context.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, telemetry.Now(), telemetry.MetricKeyEndBlocker)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.InvCheckPeriod() == 0 || sdkCtx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(sdkCtx)
}
