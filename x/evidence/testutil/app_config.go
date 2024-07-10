package testutil

import (
	_ "cosmossdk.io/x/evidence" // import as blank for app wiring

	"github.com/T-ragon/cosmos-sdk/testutil/configurator"
	_ "github.com/T-ragon/cosmos-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/bank"           // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/consensus"      // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/genutil"        // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/params"         // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/slashing"       // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/staking"        // import as blank for app wiring
)

var AppConfig = configurator.NewAppConfig(
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.StakingModule(),
	configurator.SlashingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.ParamsModule(),
	configurator.EvidenceModule(),
	configurator.GenutilModule(),
)
