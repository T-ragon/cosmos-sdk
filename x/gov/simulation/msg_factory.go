package simulation

import (
	"context"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/x/gov/keeper"
	v1 "cosmossdk.io/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/simsx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
)

func MsgDepositFactory(k *keeper.Keeper, sharedState *SharedState) simsx.SimMsgFactoryFn[*v1.MsgDeposit] {
	return func(ctx context.Context, testData *simsx.ChainDataSource, reporter simsx.SimulationReporter) ([]simsx.SimAccount, sdk.Msg) {
		r := testData.Rand()
		proposalID, ok := randomProposalIDX(r, k, ctx, v1.StatusDepositPeriod, sharedState)
		if !ok {
			reporter.Skip("no proposal in deposit state")
			return nil, nil
		}
		proposal, err := k.Proposals.Get(ctx, proposalID)
		if err != nil {
			reporter.Skip(err.Error())
			return nil, nil
		}
		// calculate deposit amount
		deposit := randDeposit(ctx, proposal, k, r, reporter)
		if reporter.IsSkipped() {
			return nil, nil
		}
		from := testData.AnyAccount(reporter, simsx.WithLiquidBalanceGTE(deposit))
		return []simsx.SimAccount{from}, v1.NewMsgDeposit(from.AddressBech32, proposalID, sdk.NewCoins(deposit))
	}
}

func MsgVoteFactory(k *keeper.Keeper, sharedState *SharedState) simsx.SimMsgFactoryFn[*v1.MsgVote] {
	return func(ctx context.Context, testData *simsx.ChainDataSource, reporter simsx.SimulationReporter) ([]simsx.SimAccount, sdk.Msg) {
		r := testData.Rand()
		proposalID, ok := randomProposalIDX(r, k, ctx, v1.StatusVotingPeriod, sharedState)
		if !ok {
			reporter.Skip("no proposal in deposit state")
			return nil, nil
		}
		from := testData.AnyAccount(reporter, simsx.WithSpendableBalance())
		msg := v1.NewMsgVote(from.AddressBech32, proposalID, randomVotingOption(r.Rand), "")
		return []simsx.SimAccount{from}, msg
	}
}
func MsgWeightedVoteFactory(k *keeper.Keeper, sharedState *SharedState) simsx.SimMsgFactoryFn[*v1.MsgVoteWeighted] {
	return func(ctx context.Context, testData *simsx.ChainDataSource, reporter simsx.SimulationReporter) ([]simsx.SimAccount, sdk.Msg) {
		r := testData.Rand()
		proposalID, ok := randomProposalIDX(r, k, ctx, v1.StatusVotingPeriod, sharedState)
		if !ok {
			reporter.Skip("no proposal in deposit state")
			return nil, nil
		}
		from := testData.AnyAccount(reporter, simsx.WithSpendableBalance())
		msg := v1.NewMsgVoteWeighted(from.AddressBech32, proposalID, randomWeightedVotingOptions(r.Rand), "")
		return []simsx.SimAccount{from}, msg
	}
}
func MsgCancelProposalFactory(k *keeper.Keeper, sharedState *SharedState) simsx.SimMsgFactoryFn[*v1.MsgCancelProposal] {
	return func(ctx context.Context, testData *simsx.ChainDataSource, reporter simsx.SimulationReporter) ([]simsx.SimAccount, sdk.Msg) {
		r := testData.Rand()
		status := simsx.OneOf(r, []v1.ProposalStatus{v1.StatusDepositPeriod, v1.StatusVotingPeriod})
		proposalID, ok := randomProposalIDX(r, k, ctx, status, sharedState)
		if !ok {
			reporter.Skip("no proposal in deposit state")
			return nil, nil
		}
		proposal, err := k.Proposals.Get(ctx, proposalID)
		if err != nil {
			reporter.Skip(err.Error())
			return nil, nil
		}

		from := testData.GetAccount(reporter, proposal.Proposer)
		if from.LiquidBalance().Empty() {
			reporter.Skip("proposer is broke")
			return nil, nil
		}
		msg := v1.NewMsgCancelProposal(proposalID, from.AddressBech32)
		return []simsx.SimAccount{from}, msg
	}
}
func randDeposit(ctx context.Context, proposal v1.Proposal, k *keeper.Keeper, r *simsx.XRand, reporter simsx.SimulationReporter) sdk.Coin {
	params, err := k.Params.Get(ctx)
	if err != nil {
		reporter.Skipf("gov params: %s", err)
		return sdk.Coin{}
	}
	minDeposits := params.MinDeposit
	if proposal.ProposalType == v1.ProposalType_PROPOSAL_TYPE_EXPEDITED {
		minDeposits = params.ExpeditedMinDeposit
	}
	minDeposit := simsx.OneOf(r, minDeposits)
	minDepositRatio, err := sdkmath.LegacyNewDecFromStr(params.GetMinDepositRatio())
	if err != nil {
		reporter.Skip(err.Error())
		return sdk.Coin{}
	}

	threshold := minDeposit.Amount.ToLegacyDec().Mul(minDepositRatio).TruncateInt()
	depositAmount, err := r.PositiveSDKIntInRange(threshold, minDeposit.Amount)
	if err != nil {
		reporter.Skipf("deposit amount: %s", err)
		return sdk.Coin{}
	}
	return sdk.Coin{Denom: minDeposit.Denom, Amount: depositAmount}
}

// Pick a random proposal ID between the initial proposal ID
// (defined in gov GenesisState) and the latest proposal ID
// that matches a given Status.
// It does not provide a default ID.
func randomProposalIDX(r *simsx.XRand, k *keeper.Keeper, ctx context.Context, status v1.ProposalStatus, s *SharedState) (proposalID uint64, found bool) {
	proposalID, _ = k.ProposalID.Peek(ctx)
	if initialProposalID := s.getMinProposalID(); initialProposalID == unsetProposalID {
		s.setMinProposalID(proposalID)
	} else if initialProposalID < proposalID {
		proposalID = r.Uint64InRange(initialProposalID, proposalID)
	}
	proposal, err := k.Proposals.Get(ctx, proposalID)
	if err != nil || proposal.Status != status {

		// todo (Alex): use recursion to find something with n tries
		return proposalID, false
	}

	return proposalID, true
}

// Pick a random weighted voting options
func randomWeightedVotingOptions(r *rand.Rand) v1.WeightedVoteOptions {
	w1 := r.Intn(100 + 1)
	w2 := r.Intn(100 - w1 + 1)
	w3 := r.Intn(100 - w1 - w2 + 1)
	w4 := 100 - w1 - w2 - w3
	weightedVoteOptions := v1.WeightedVoteOptions{}
	if w1 > 0 {
		weightedVoteOptions = append(weightedVoteOptions, &v1.WeightedVoteOption{
			Option: v1.OptionYes,
			Weight: sdkmath.LegacyNewDecWithPrec(int64(w1), 2).String(),
		})
	}
	if w2 > 0 {
		weightedVoteOptions = append(weightedVoteOptions, &v1.WeightedVoteOption{
			Option: v1.OptionAbstain,
			Weight: sdkmath.LegacyNewDecWithPrec(int64(w2), 2).String(),
		})
	}
	if w3 > 0 {
		weightedVoteOptions = append(weightedVoteOptions, &v1.WeightedVoteOption{
			Option: v1.OptionNo,
			Weight: sdkmath.LegacyNewDecWithPrec(int64(w3), 2).String(),
		})
	}
	if w4 > 0 {
		weightedVoteOptions = append(weightedVoteOptions, &v1.WeightedVoteOption{
			Option: v1.OptionNoWithVeto,
			Weight: sdkmath.LegacyNewDecWithPrec(int64(w4), 2).String(),
		})
	}
	return weightedVoteOptions
}
