package distribution

import (
	"github.com/T-ragon/cosmos-sdk/telemetry"
	sdk "github.com/T-ragon/cosmos-sdk/types"
	"github.com/T-ragon/cosmos-sdk/x/distribution/keeper"
	"github.com/T-ragon/cosmos-sdk/x/distribution/types"
)

// BeginBlocker sets the proposer for determining distribution during endblock
// and distribute rewards for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, telemetry.Now(), telemetry.MetricKeyBeginBlocker)

	// determine the total power signing the block
	var previousTotalPower int64
	for _, voteInfo := range ctx.VoteInfos() {
		previousTotalPower += voteInfo.Validator.Power
	}

	// TODO this is Tendermint-dependent
	// ref https://github.com/T-ragon/cosmos-sdk/issues/3095
	if ctx.BlockHeight() > 1 {
		if err := k.AllocateTokens(ctx, previousTotalPower, ctx.VoteInfos()); err != nil {
			return err
		}
	}

	// record the proposer for when we payout on the next block
	consAddr := sdk.ConsAddress(ctx.BlockHeader().ProposerAddress)
	return k.SetPreviousProposerConsAddr(ctx, consAddr)
}
