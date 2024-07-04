package simulation

import (
	"context"
	"slices"

	"github.com/cosmos/cosmos-sdk/simsx"

	"cosmossdk.io/x/bank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/maps"
)

func MsgSendFactory() simsx.SimMsgFactoryFn[*types.MsgSend] {
	return func(ctx context.Context, testData *simsx.ChainDataSource, reporter simsx.SimulationReporter) ([]simsx.SimAccount, sdk.Msg) {
		from := testData.AnyAccount(reporter, simsx.WithSpendableBalance())
		to := testData.AnyAccount(reporter, simsx.ExcludeAccounts(from))
		coins := from.LiquidBalance().RandSubsetCoins(reporter, simsx.WithSendEnabledCoins())
		return []simsx.SimAccount{from}, types.NewMsgSend(from.AddressBech32, to.AddressBech32, coins)
	}
}

func MsgMultiSendFactory() simsx.SimMsgFactoryFn[*types.MsgMultiSend] {
	return func(ctx context.Context, testData *simsx.ChainDataSource, reporter simsx.SimulationReporter) ([]simsx.SimAccount, sdk.Msg) {
		r := testData.Rand()
		var (
			sending        = make([]types.Input, r.Intn(1)+1)
			receiving      = make([]types.Output, r.Intn(3)+1)
			senderAcc      = make([]simsx.SimAccount, len(sending))
			usedAddrsIdx   = make(map[string]struct{}) // use map to check if address already exists as input
			totalSentCoins sdk.Coins
		)
		for i := range sending {
			// generate random input fields, ignore to address
			from := testData.AnyAccount(reporter, simsx.WithSpendableBalance(), simsx.ExcludeAddresses(maps.Keys(usedAddrsIdx)...))
			if reporter.IsSkipped() {
				return nil, nil
			}
			coins := from.LiquidBalance().RandSubsetCoins(reporter, simsx.WithSendEnabledCoins())
			// set input address in used address map
			usedAddrsIdx[from.AddressBech32] = struct{}{}

			// set signer privkey
			senderAcc[i] = from

			// set next input and accumulate total sent coins
			sending[i] = types.NewInput(from.AddressBech32, coins)
			totalSentCoins = totalSentCoins.Add(coins...)
		}

		for i := range receiving {
			receiver := testData.AnyAccount(reporter)
			if reporter.IsSkipped() {
				return nil, nil
			}

			var outCoins sdk.Coins
			// split total sent coins into random subsets for output
			if i == len(receiving)-1 {
				// last one receives remaining amount
				outCoins = totalSentCoins
			} else {
				// take random subset of remaining coins for output
				// and update remaining coins
				outCoins = r.SubsetCoins(totalSentCoins)
				totalSentCoins = totalSentCoins.Sub(outCoins...)
			}

			receiving[i] = types.NewOutput(receiver.AddressBech32, outCoins)
		}

		// remove any entries that have no coins
		slices.DeleteFunc(receiving, func(o types.Output) bool {
			return o.Coins.Empty()
		})
		return senderAcc, &types.MsgMultiSend{Inputs: sending, Outputs: receiving}
	}
}

// MsgUpdateParamsFactory creates a gov proposal for param updates
func MsgUpdateParamsFactory() simsx.SimMsgFactoryFn[*types.MsgUpdateParams] {
	return func(_ context.Context, testData *simsx.ChainDataSource, reporter simsx.SimulationReporter) ([]simsx.SimAccount, sdk.Msg) {
		params := types.DefaultParams()
		params.DefaultSendEnabled = testData.Rand().Intn(2) == 0
		return nil, &types.MsgUpdateParams{
			Authority: testData.ModuleAccountAddress(reporter, "gov"),
			Params:    params,
		}
	}
}
