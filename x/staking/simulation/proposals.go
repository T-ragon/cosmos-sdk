package simulation

import (
	"math/rand"
	"time"

	coreaddress "cosmossdk.io/core/address"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/x/staking/types"

	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	"github.com/T-ragon/cosmos-sdk/v3/types/address"
	simtypes "github.com/T-ragon/cosmos-sdk/v3/types/simulation"
	"github.com/T-ragon/cosmos-sdk/v3/x/simulation"
)

// Simulation operation weights constants
const (
	DefaultWeightMsgUpdateParams int = 100

	OpWeightMsgUpdateParams = "op_weight_msg_update_params"
)

// ProposalMsgs defines the module weighted proposals' contents
func ProposalMsgs() []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			OpWeightMsgUpdateParams,
			DefaultWeightMsgUpdateParams,
			SimulateMsgUpdateParams,
		),
	}
}

// SimulateMsgUpdateParams returns a random MsgUpdateParams
func SimulateMsgUpdateParams(r *rand.Rand, _ []simtypes.Account, addressCodec coreaddress.Codec) (sdk.Msg, error) {
	// use the default gov module account address as authority
	var authority sdk.AccAddress = address.Module("gov")

	params := types.DefaultParams()
	params.BondDenom = simtypes.RandStringOfLength(r, 10)
	params.HistoricalEntries = uint32(simtypes.RandIntBetween(r, 0, 1000))
	params.MaxEntries = uint32(simtypes.RandIntBetween(r, 1, 1000))
	params.MaxValidators = uint32(simtypes.RandIntBetween(r, 1, 1000))
	params.UnbondingTime = time.Duration(simtypes.RandTimestamp(r).UnixNano())
	params.MinCommissionRate = simtypes.RandomDecAmount(r, sdkmath.LegacyNewDec(1))

	addr, err := addressCodec.BytesToString(authority)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateParams{
		Authority: addr,
		Params:    params,
	}, nil
}
