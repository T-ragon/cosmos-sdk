package simulation

import (
	"math/rand"

	coreaddress "cosmossdk.io/core/address"
	"cosmossdk.io/x/bank/types"

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
func SimulateMsgUpdateParams(r *rand.Rand, _ []simtypes.Account, ac coreaddress.Codec) (sdk.Msg, error) {
	// use the default gov module account address as authority
	authority, err := ac.BytesToString(address.Module("gov"))
	if err != nil {
		return nil, err
	}

	params := types.DefaultParams()
	params.DefaultSendEnabled = r.Intn(2) == 0

	return &types.MsgUpdateParams{
		Authority: authority,
		Params:    params,
	}, nil
}
