package testutil

import (
	_ "cosmossdk.io/x/accounts"       // import as blank for app wiring
	_ "cosmossdk.io/x/auth"           // import as blank for app wiring
	_ "cosmossdk.io/x/auth/tx/config" // import as blank for app wiring
	_ "cosmossdk.io/x/authz"          // import as blank for app wiring
	_ "cosmossdk.io/x/bank"           // import as blank for app wiring
	_ "cosmossdk.io/x/consensus"      // import as blank for app wiring
	_ "cosmossdk.io/x/group/module"   // import as blank for app wiring
	_ "cosmossdk.io/x/mint"           // import as blank for app wiring
	_ "cosmossdk.io/x/staking"        // import as blank for app wiring

	"github.com/T-ragon/cosmos-sdk/v3/testutil/configurator"
	_ "github.com/T-ragon/cosmos-sdk/v3/x/genutil" // import as blank for app wiring
)

var AppConfig = configurator.NewAppConfig(
	configurator.AccountsModule(),
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.StakingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.GenutilModule(),
	configurator.GroupModule(),
)
