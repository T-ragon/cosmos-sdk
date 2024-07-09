package exported

import (
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

// GenesisBalance defines a genesis balance interface that allows for account
// address and balance retrieval.
type GenesisBalance interface {
	GetAddress() string
	GetCoins() sdk.Coins
}
