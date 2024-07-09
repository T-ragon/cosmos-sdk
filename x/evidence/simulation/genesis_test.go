package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	"cosmossdk.io/x/evidence/simulation"
	"cosmossdk.io/x/evidence/types"

	"github.com/T-ragon/cosmos-sdk/v3/codec"
	"github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	codectypes "github.com/T-ragon/cosmos-sdk/v3/codec/types"
	"github.com/T-ragon/cosmos-sdk/v3/types/module"
	simtypes "github.com/T-ragon/cosmos-sdk/v3/types/simulation"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	cdcOpts := testutil.CodecOptions{}
	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:      make(simtypes.AppParams),
		Cdc:            cdc,
		AddressCodec:   cdcOpts.GetAddressCodec(),
		ValidatorCodec: cdcOpts.GetValidatorCodec(),
		Rand:           r,
		NumBonded:      3,
		Accounts:       simtypes.RandomAccounts(r, 3),
		InitialStake:   math.NewInt(1000),
		GenState:       make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var evidenceGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &evidenceGenesis)

	require.Len(t, evidenceGenesis.Evidence, 0)
}
