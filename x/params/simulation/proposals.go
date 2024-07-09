package simulation

import (
	simtypes "github.com/T-ragon/cosmos-sdk/v3/types/simulation"
	"github.com/T-ragon/cosmos-sdk/v3/x/simulation"
)

const (
	// OpWeightSubmitParamChangeProposal app params key for param change proposal
	OpWeightSubmitParamChangeProposal = "op_weight_submit_param_change_proposal"
	DefaultWeightParamChangeProposal  = 5
)

// ProposalContents defines the module weighted proposals' contents
//
//nolint:staticcheck // used for legacy testing
func ProposalContents(paramChanges []simtypes.LegacyParamChange) []simtypes.WeightedProposalContent {
	return []simtypes.WeightedProposalContent{
		simulation.NewWeightedProposalContent(
			OpWeightSubmitParamChangeProposal,
			DefaultWeightParamChangeProposal,
			SimulateParamChangeProposalContent(paramChanges),
		),
	}
}
