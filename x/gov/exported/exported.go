package exported

import (
	sdk "github.com/T-ragon/cosmos-sdk/types"
	paramtypes "github.com/T-ragon/cosmos-sdk/x/params/types"
)

type (
	ParamSet = paramtypes.ParamSet

	// Subspace defines an interface that implements the legacy x/params Subspace
	// type.
	//
	// NOTE: This is used solely for migration of x/params managed parameters.
	ParamSubspace interface {
		Get(ctx sdk.Context, key []byte, ptr interface{})
	}
)
