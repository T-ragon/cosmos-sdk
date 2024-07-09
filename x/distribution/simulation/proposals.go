package simulation

import (
	"math/rand"

	coreaddress "cosmossdk.io/core/address"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/x/distribution/types"

	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	"github.com/T-ragon/cosmos-sdk/v3/types/address"
	simtypes "github.com/T-ragon/cosmos-sdk/v3/types/simulation"
	"github.com/T-ragon/cosmos-sdk/v3/x/simulation"
)

// Simulation operation weights constants
const (
	DefaultWeightMsgUpdateParams int = 50

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
func SimulateMsgUpdateParams(r *rand.Rand, _ []simtypes.Account, cdc coreaddress.Codec) (sdk.Msg, error) {
	// use the default gov module account address as authority
	var authority sdk.AccAddress = address.Module("gov")

	params := types.DefaultParams()
	params.CommunityTax = simtypes.RandomDecAmount(r, sdkmath.LegacyNewDec(1))
	params.WithdrawAddrEnabled = r.Intn(2) == 0

	authorityAddr, err := cdc.BytesToString(authority)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateParams{
		Authority: authorityAddr,
		Params:    params,
	}, nil
}
