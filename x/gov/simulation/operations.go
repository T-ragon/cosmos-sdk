package simulation

import (
	"github.com/cosmos/cosmos-sdk/simsx"
	"math"
	"math/rand"
	"sync/atomic"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/x/gov/keeper"
	"cosmossdk.io/x/gov/types"
	v1 "cosmossdk.io/x/gov/types/v1"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

const unsetProposalID = 100000000000000

// Governance message types and routes
var (
	TypeMsgSubmitProposal = sdk.MsgTypeURL(&v1.MsgSubmitProposal{})
	TypeMsgVote           = sdk.MsgTypeURL(&v1.MsgVote{})
)

// Simulation operation weights constants
const (
	DefaultWeightTextProposal = 5
)

// SharedState shared state between message invocations
type SharedState struct {
	minProposalID atomic.Uint64
}

// NewSharedState constructor
func NewSharedState() *SharedState {
	r := &SharedState{}
	r.setMinProposalID(unsetProposalID)
	return r
}

func (s *SharedState) getMinProposalID() uint64 {
	return s.minProposalID.Load()
}

func (s *SharedState) setMinProposalID(id uint64) {
	s.minProposalID.Store(id)
}

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams,
	txGen client.TxConfig,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	wMsgs []simtypes.WeightedProposalMsg,
	wContents []simtypes.WeightedProposalContent, //nolint:staticcheck // used for legacy testing
) simulation.WeightedOperations {
	// generate the weighted operations for the proposal msgs
	var wProposalOps simulation.WeightedOperations
	for _, wMsg := range wMsgs {
		wMsg := wMsg // pin variable
		var weight int
		appParams.GetOrGenerate(wMsg.AppParamsKey(), &weight, nil,
			func(_ *rand.Rand) { weight = wMsg.DefaultWeight() },
		)

		wProposalOps = append(
			wProposalOps,
			simulation.NewWeightedOperation(
				weight,
				SimulateMsgSubmitProposal(txGen, ak, bk, k, wMsg.MsgSimulatorFn()),
			),
		)
	}

	// generate the weighted operations for the proposal contents
	var wLegacyProposalOps simulation.WeightedOperations
	for _, wContent := range wContents {
		wContent := wContent // pin variable
		var weight int
		appParams.GetOrGenerate(wContent.AppParamsKey(), &weight, nil,
			func(_ *rand.Rand) { weight = wContent.DefaultWeight() },
		)

		wLegacyProposalOps = append(
			wLegacyProposalOps,
			simulation.NewWeightedOperation(
				weight,
				SimulateMsgSubmitLegacyProposal(txGen, ak, bk, k, wContent.ContentSimulatorFn()),
			),
		)
	}

	state := NewSharedState()

	reg := simsx.NewSimsMsgRegistryAdapter(simsx.NewBasicSimulationReporter(), ak, bk.(simsx.BalanceSource), txGen)
	weights := simsx.ParamWeightSource(appParams)
	reg.Add(weights.Get("msg_submit_proposal", 100), MsgDepositFactory(k, state))
	reg.Add(weights.Get("msg_deposit", 100), MsgDepositFactory(k, state))
	reg.Add(weights.Get("msg_vote", 67), MsgVoteFactory(k, state))
	reg.Add(weights.Get("msg_weighted_vote", 33), MsgWeightedVoteFactory(k, state))
	reg.Add(weights.Get("cancel_proposal", 5), MsgCancelProposalFactory(k, state))

	return append(reg.ToLegacy(), append(wProposalOps, wLegacyProposalOps...)...)
}

// SimulateMsgSubmitProposal simulates creating a msg Submit Proposal
// voting on the proposal, and subsequently slashing the proposal. It is implemented using
// future operations.
func SimulateMsgSubmitProposal(
	txGen client.TxConfig,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	msgSim simtypes.MsgSimulatorFnX,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msgs := []sdk.Msg{}
		proposalMsg, err := msgSim(ctx, r, accs, ak.AddressCodec())
		if err != nil {
			return simtypes.OperationMsg{}, nil, err
		}
		if proposalMsg != nil {
			msgs = append(msgs, proposalMsg)
		}

		return simulateMsgSubmitProposal(txGen, ak, bk, k, msgs)(r, app, ctx, accs, chainID)
	}
}

// SimulateMsgSubmitLegacyProposal simulates creating a msg Submit Proposal
// voting on the proposal, and subsequently slashing the proposal. It is implemented using
// future operations.
func SimulateMsgSubmitLegacyProposal(
	txGen client.TxConfig,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	contentSim simtypes.ContentSimulatorFn, //nolint:staticcheck // used for legacy testing
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// 1) submit proposal now
		content := contentSim(r, ctx, accs)
		if content == nil {
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgSubmitProposal, "content is nil"), nil, nil
		}

		govacc, err := ak.AddressCodec().BytesToString(k.GetGovernanceAccount(ctx).GetAddress())
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgSubmitProposal, "error getting governance account address"), nil, err
		}
		contentMsg, err := v1.NewLegacyContent(content, govacc)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgSubmitProposal, "error converting legacy content into proposal message"), nil, err
		}

		return simulateMsgSubmitProposal(txGen, ak, bk, k, []sdk.Msg{contentMsg})(r, app, ctx, accs, chainID)
	}
}

func simulateMsgSubmitProposal(
	txGen client.TxConfig,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	proposalMsgs []sdk.Msg,
) simtypes.Operation {
	// The states are:
	// column 1: All validators vote
	// column 2: 90% vote
	// column 3: 75% vote
	// column 4: 40% vote
	// column 5: 15% vote
	// column 6: no one votes
	// All columns sum to 100 for simplicity, values chosen by @valardragon semi-arbitrarily,
	// feel free to change.
	numVotesTransitionMatrix, _ := simulation.CreateTransitionMatrix([][]int{
		{20, 10, 0, 0, 0, 0},
		{55, 50, 20, 10, 0, 0},
		{25, 25, 30, 25, 30, 15},
		{0, 15, 30, 25, 30, 30},
		{0, 0, 20, 30, 30, 30},
		{0, 0, 0, 10, 10, 25},
	})

	statePercentageArray := []float64{1, .9, .75, .4, .15, 0}
	curNumVotesState := 1

	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accs []simtypes.Account,
		chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		expedited := r.Intn(2) == 0
		deposit, skip, err := randomDeposit(r, ctx, ak, bk, k, simAccount.Address, true, expedited)
		switch {
		case skip:
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgSubmitProposal, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgSubmitProposal, "unable to generate deposit"), nil, err
		}

		proposalType := v1.ProposalType_PROPOSAL_TYPE_STANDARD
		if expedited {
			proposalType = v1.ProposalType_PROPOSAL_TYPE_EXPEDITED
		}

		accAddr, err := ak.AddressCodec().BytesToString(simAccount.Address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgSubmitProposal, "error getting simAccount address"), nil, err
		}
		msg, err := v1.NewMsgSubmitProposal(
			proposalMsgs,
			deposit,
			accAddr,
			simtypes.RandStringOfLength(r, 100),
			simtypes.RandStringOfLength(r, 100),
			simtypes.RandStringOfLength(r, 100),
			proposalType,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "unable to generate a submit proposal msg"), nil, err
		}

		account := ak.GetAccount(ctx, simAccount.Address)
		tx, err := simtestutil.GenSignedMockTx(
			r,
			txGen,
			[]sdk.Msg{msg},
			sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)},
			simtestutil.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "unable to deliver tx"), nil, err
		}

		opMsg := simtypes.NewOperationMsg(msg, true, "")

		// get the submitted proposal ID
		proposalID, err := k.ProposalID.Peek(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "unable to generate proposalID"), nil, err
		}

		// 2) Schedule operations for votes
		// 2.1) first pick a number of people to vote.
		curNumVotesState = numVotesTransitionMatrix.NextState(r, curNumVotesState)
		numVotes := int(math.Ceil(float64(len(accs)) * statePercentageArray[curNumVotesState]))

		// 2.2) select who votes and when
		whoVotes := r.Perm(len(accs))

		// didntVote := whoVotes[numVotes:]
		whoVotes = whoVotes[:numVotes]
		params, _ := k.Params.Get(ctx)
		votingPeriod := params.VotingPeriod

		fops := make([]simtypes.FutureOperation, numVotes+1)
		for i := 0; i < numVotes; i++ {
			whenVote := ctx.HeaderInfo().Time.Add(time.Duration(r.Int63n(int64(votingPeriod.Seconds()))) * time.Second)
			fops[i] = simtypes.FutureOperation{
				BlockTime: whenVote,
				Op:        operationSimulateMsgVote(txGen, ak, bk, k, accs[whoVotes[i]], int64(proposalID), nil),
			}
		}

		return opMsg, fops, nil
	}
}

func operationSimulateMsgVote(
	txGen client.TxConfig,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	simAccount simtypes.Account,
	proposalIDInt int64,
	s *SharedState,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		if simAccount.Equals(simtypes.Account{}) {
			simAccount, _ = simtypes.RandomAcc(r, accs)
		}

		var proposalID uint64

		switch {
		case proposalIDInt < 0:
			var ok bool
			proposalID, ok = randomProposalID(r, k, ctx, v1.StatusVotingPeriod, s)
			if !ok {
				return simtypes.NoOpMsg(types.ModuleName, TypeMsgVote, "unable to generate proposalID"), nil, nil
			}
		default:
			proposalID = uint64(proposalIDInt)
		}

		option := randomVotingOption(r)
		addr, err := ak.AddressCodec().BytesToString(simAccount.Address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgVote, "unable to get simAccount address"), nil, err
		}
		msg := v1.NewMsgVote(addr, proposalID, option, "")

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// Pick a random deposit with a random denomination with a
// deposit amount between (0, min(balance, minDepositAmount))
// This is to simulate multiple users depositing to get the
// proposal above the minimum deposit amount
func randomDeposit(
	r *rand.Rand,
	ctx sdk.Context,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	addr sdk.AccAddress,
	useMinAmount bool,
	expedited bool,
) (deposit sdk.Coins, skip bool, err error) {
	account := ak.GetAccount(ctx, addr)
	spendable := bk.SpendableCoins(ctx, account.GetAddress())

	if spendable.Empty() {
		return nil, true, nil // skip
	}

	params, _ := k.Params.Get(ctx)
	minDeposit := params.MinDeposit
	if expedited {
		minDeposit = params.ExpeditedMinDeposit
	}
	denomIndex := r.Intn(len(minDeposit))
	denom := minDeposit[denomIndex].Denom

	spendableBalance := spendable.AmountOf(denom)
	if spendableBalance.IsZero() {
		return nil, true, nil
	}

	minDepositAmount := minDeposit[denomIndex].Amount

	minDepositRatio, err := sdkmath.LegacyNewDecFromStr(params.GetMinDepositRatio())
	if err != nil {
		return nil, false, err
	}

	threshold := minDepositAmount.ToLegacyDec().Mul(minDepositRatio).TruncateInt()

	minAmount := sdkmath.ZeroInt()
	if useMinAmount {
		minDepositPercent, err := sdkmath.LegacyNewDecFromStr(params.MinInitialDepositRatio)
		if err != nil {
			return nil, false, err
		}

		minAmount = sdkmath.LegacyNewDecFromInt(minDepositAmount).Mul(minDepositPercent).TruncateInt()
	}

	amount, err := simtypes.RandPositiveInt(r, minDepositAmount.Sub(minAmount))
	if err != nil {
		return nil, false, err
	}
	amount = amount.Add(minAmount)

	if amount.GT(spendableBalance) || amount.LT(threshold) {
		return nil, true, nil
	}

	return sdk.Coins{sdk.NewCoin(denom, amount)}, false, nil
}

// Pick a random proposal ID between the initial proposal ID
// (defined in gov GenesisState) and the latest proposal ID
// that matches a given Status.
// It does not provide a default ID.
func randomProposalID(r *rand.Rand, k *keeper.Keeper, ctx sdk.Context, status v1.ProposalStatus, s *SharedState) (proposalID uint64, found bool) {
	proposalID, _ = k.ProposalID.Peek(ctx)
	if initialProposalID := s.getMinProposalID(); initialProposalID == unsetProposalID {
		s.setMinProposalID(proposalID)
	} else if initialProposalID < proposalID {
		proposalID = uint64(simtypes.RandIntBetween(r, int(initialProposalID), int(proposalID)))
	}
	proposal, err := k.Proposals.Get(ctx, proposalID)
	if err != nil || proposal.Status != status {
		return proposalID, false
	}

	return proposalID, true
}

// Pick a random voting option
func randomVotingOption(r *rand.Rand) v1.VoteOption {
	switch r.Intn(4) {
	case 0:
		return v1.OptionYes
	case 1:
		return v1.OptionAbstain
	case 2:
		return v1.OptionNo
	case 3:
		return v1.OptionNoWithVeto
	default:
		panic("invalid vote option")
	}
}
