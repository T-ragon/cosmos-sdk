package testutil

import (
	"github.com/T-ragon/cosmos-sdk/testutil/configurator"
	_ "github.com/T-ragon/cosmos-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/bank"           // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/consensus"      // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/distribution"   // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/genutil"        // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/mint"           // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/params"         // import as blank for app wiring
	_ "github.com/T-ragon/cosmos-sdk/x/staking"        // import as blank for app wiring
)

var AppConfig = configurator.NewAppConfig(
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.StakingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.ParamsModule(),
	configurator.GenutilModule(),
	configurator.DistributionModule(),
	configurator.MintModule(),
)
