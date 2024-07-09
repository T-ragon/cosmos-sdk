package simulation

import (
	"math/rand"

	coreaddress "cosmossdk.io/core/address"
	pooltypes "cosmossdk.io/x/protocolpool/types"

	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	"github.com/T-ragon/cosmos-sdk/v3/types/address"
	simtypes "github.com/T-ragon/cosmos-sdk/v3/types/simulation"
	"github.com/T-ragon/cosmos-sdk/v3/x/simulation"
)

const (
	OpWeightMsgCommunityPoolSpend = "op_weight_msg_community_pool_spend"

	DefaultWeightMsgCommunityPoolSpend int = 50
)

func ProposalMsgs() []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			OpWeightMsgCommunityPoolSpend,
			DefaultWeightMsgCommunityPoolSpend,
			SimulateMsgCommunityPoolSpend,
		),
	}
}

func SimulateMsgCommunityPoolSpend(r *rand.Rand, _ []simtypes.Account, cdc coreaddress.Codec) (sdk.Msg, error) {
	// use the default gov module account address as authority
	var authority sdk.AccAddress = address.Module("gov")

	accs := simtypes.RandomAccounts(r, 5)
	acc, _ := simtypes.RandomAcc(r, accs)

	coins, err := sdk.ParseCoinsNormalized("100stake,2testtoken")
	if err != nil {
		return nil, err
	}

	authorityAddr, err := cdc.BytesToString(authority)
	if err != nil {
		return nil, err
	}
	recipentAddr, err := cdc.BytesToString(acc.Address)
	if err != nil {
		return nil, err
	}
	return &pooltypes.MsgCommunityPoolSpend{
		Authority: authorityAddr,
		Recipient: recipentAddr,
		Amount:    coins,
	}, nil
}
