package types

import (
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/T-ragon/cosmos-sdk/baseapp"
)

// Deprecated.
func ConsensusParamsKeyTable() KeyTable {
	return NewKeyTable(
		NewParamSetPair(
			baseapp.ParamStoreKeyBlockParams, cmtproto.BlockParams{}, baseapp.ValidateBlockParams,
		),
		NewParamSetPair(
			baseapp.ParamStoreKeyEvidenceParams, cmtproto.EvidenceParams{}, baseapp.ValidateEvidenceParams,
		),
		NewParamSetPair(
			baseapp.ParamStoreKeyValidatorParams, cmtproto.ValidatorParams{}, baseapp.ValidateValidatorParams,
		),
	)
}
