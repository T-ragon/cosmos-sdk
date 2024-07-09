package genutil

import (
	"context"

	"cosmossdk.io/core/genesis"

	"github.com/T-ragon/cosmos-sdk/v3/client"
	"github.com/T-ragon/cosmos-sdk/v3/types/module"
	"github.com/T-ragon/cosmos-sdk/v3/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx context.Context, stakingKeeper types.StakingKeeper,
	deliverTx genesis.TxHandler, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) (validatorUpdates []module.ValidatorUpdate, err error) {
	if len(genesisState.GenTxs) > 0 {
		moduleValidatorUpdates, err := DeliverGenTxs(ctx, genesisState.GenTxs, stakingKeeper, deliverTx, txEncodingConfig)
		if err != nil {
			return nil, err
		}
		return moduleValidatorUpdates, nil
	}
	return
}
