package simulation_test

import (
	"math/rand"
	"testing"

	"gotest.tools/v3/assert"

	"cosmossdk.io/x/bank/simulation"
	"cosmossdk.io/x/bank/types"

	codectestutil "github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	"github.com/T-ragon/cosmos-sdk/v3/types/address"
	simtypes "github.com/T-ragon/cosmos-sdk/v3/types/simulation"
)

func TestProposalMsgs(t *testing.T) {
	ac := codectestutil.CodecOptions{}.GetAddressCodec()

	// initialize parameters
	s := rand.NewSource(1)
	r := rand.New(s)

	accounts := simtypes.RandomAccounts(r, 3)

	// execute ProposalMsgs function
	weightedProposalMsgs := simulation.ProposalMsgs()
	assert.Assert(t, len(weightedProposalMsgs) == 1)

	w0 := weightedProposalMsgs[0]

	// tests w0 interface:
	assert.Equal(t, simulation.OpWeightMsgUpdateParams, w0.AppParamsKey())
	assert.Equal(t, simulation.DefaultWeightMsgUpdateParams, w0.DefaultWeight())

	msg, err := w0.MsgSimulatorFn()(r, accounts, ac)
	assert.NilError(t, err)
	msgUpdateParams, ok := msg.(*types.MsgUpdateParams)
	assert.Assert(t, ok)

	authority, err := ac.BytesToString(address.Module("gov"))
	assert.NilError(t, err)
	assert.Equal(t, authority, msgUpdateParams.Authority)
	assert.Assert(t, len(msgUpdateParams.Params.SendEnabled) == 0) //nolint:staticcheck // we're testing the old way here
	assert.Equal(t, true, msgUpdateParams.Params.DefaultSendEnabled)
}
