package simulation_test

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/T-ragon/cosmos-sdk/types"
)

func init() {
	sdk.DefaultPowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}
