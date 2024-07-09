//go:build sims

package simapp

import (
	"os"
	"testing"

	"github.com/T-ragon/cosmos-sdk/v3/testutils/sims"

	"cosmossdk.io/core/log"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/T-ragon/cosmos-sdk/v3/baseapp"
	"github.com/T-ragon/cosmos-sdk/v3/client/flags"
	"github.com/T-ragon/cosmos-sdk/v3/server"
	simtestutil "github.com/T-ragon/cosmos-sdk/v3/testutil/sims"
	simtypes "github.com/T-ragon/cosmos-sdk/v3/types/simulation"
	"github.com/T-ragon/cosmos-sdk/v3/x/simulation"
	simcli "github.com/T-ragon/cosmos-sdk/v3/x/simulation/client/cli"
)

var FlagEnableBenchStreamingValue bool

// Get flags every time the simulator is run
func init() {
	flag.BoolVar(&FlagEnableBenchStreamingValue, "EnableStreaming", false, "Enable streaming service")
}

// Profile with:
// /usr/local/go/bin/go test -benchmem -run=^$ cosmossdk.io/simapp -bench ^BenchmarkFullAppSimulation$ -Commit=true -cpuprofile cpu.out
func BenchmarkFullAppSimulation(b *testing.B) {
	b.ReportAllocs()

	config := simcli.NewConfigFromFlags()
	config.ChainID = sims.SimAppChainID

	db, dir, logger, skip, err := simtestutil.SetupSimulation(config, "goleveldb-app-sim", "Simulation", simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	if err != nil {
		b.Fatalf("simulation setup failed: %s", err.Error())
	}

	if skip {
		b.Skip("skipping benchmark application simulation")
	}

	defer func() {
		require.NoError(b, db.Close())
		require.NoError(b, os.RemoveAll(dir))
	}()

	appOptions := viper.New()
	appOptions.SetDefault(flags.FlagHome, DefaultNodeHome)
	appOptions.SetDefault(server.FlagInvCheckPeriod, simcli.FlagPeriodValue)

	app := NewSimApp(logger, db, nil, true, appOptions, interBlockCacheOpt(), baseapp.SetChainID(sims.SimAppChainID))

	// run randomized simulation
	simParams, simErr := simulation.SimulateFromSeedX(
		b,
		log.NewNopLogger(),
		os.Stdout,
		app.BaseApp,
		simtestutil.AppStateFn(app.AppCodec(), app.AuthKeeper.AddressCodec(), app.StakingKeeper.ValidatorAddressCodec(), app.SimulationManager(), app.DefaultGenesis()),
		simtypes.RandomAccounts,
		simtestutil.SimulationOperations(app, app.AppCodec(), config, app.txConfig),
		BlockedAddresses(),
		config,
		app.AppCodec(),
		app.txConfig.SigningContext().AddressCodec(),
		&simulation.DummyLogWriter{},
	)

	// export state and simParams before the simulation error is checked
	if err = simtestutil.CheckExportSimulation(app, config, simParams); err != nil {
		b.Fatal(err)
	}

	if simErr != nil {
		b.Fatal(simErr)
	}

	if config.Commit {
		simtestutil.PrintStats(db)
	}
}
