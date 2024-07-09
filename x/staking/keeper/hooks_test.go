package keeper_test

import (
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

func (s *KeeperTestSuite) TestHookAfterConsensusPubKeyUpdate() {
	stKeeper := s.stakingKeeper
	ctx := s.ctx
	require := s.Require()

	rotationFee := sdk.NewInt64Coin("stake", 1000000)
	err := stKeeper.Hooks().AfterConsensusPubKeyUpdate(ctx, PKs[0], PKs[1], rotationFee)
	require.NoError(err)
}
