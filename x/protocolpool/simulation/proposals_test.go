package simulation_test

import (
	"math/rand"
	"testing"

	"gotest.tools/v3/assert"

	"cosmossdk.io/x/protocolpool/simulation"
	pooltypes "cosmossdk.io/x/protocolpool/types"

	codectestutil "github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	"github.com/T-ragon/cosmos-sdk/v3/types/address"
	simtypes "github.com/T-ragon/cosmos-sdk/v3/types/simulation"
)

func TestProposalMsgs(t *testing.T) {
	// initialize parameters
	s := rand.NewSource(1)
	r := rand.New(s)
	addressCodec := codectestutil.CodecOptions{}.GetAddressCodec()
	accounts := simtypes.RandomAccounts(r, 3)

	// execute ProposalMsgs function
	weightedProposalMsgs := simulation.ProposalMsgs()
	assert.Assert(t, len(weightedProposalMsgs) == 1)

	w0 := weightedProposalMsgs[0]

	// tests w0 interface:
	assert.Equal(t, simulation.OpWeightMsgCommunityPoolSpend, w0.AppParamsKey())
	assert.Equal(t, simulation.DefaultWeightMsgCommunityPoolSpend, w0.DefaultWeight())

	msg, err := w0.MsgSimulatorFn()(r, accounts, addressCodec)
	assert.NilError(t, err)
	msgCommunityPoolSpend, ok := msg.(*pooltypes.MsgCommunityPoolSpend)
	assert.Assert(t, ok)

	coins, err := sdk.ParseCoinsNormalized("100stake,2testtoken")
	assert.NilError(t, err)

	authAddr, err := addressCodec.BytesToString(address.Module("gov"))
	assert.NilError(t, err)
	assert.Equal(t, authAddr, msgCommunityPoolSpend.Authority)
	assert.Assert(t, msgCommunityPoolSpend.Amount.Equal(coins))
}
