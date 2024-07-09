//go:build e2e
// +build e2e

package group

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/simapp"

	"github.com/T-ragon/cosmos-sdk/testutil/network"
)

func TestE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig(simapp.NewTestNetworkFixture)
	cfg.NumValidators = 1
	suite.Run(t, NewE2ETestSuite(cfg))
}
