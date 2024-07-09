package types

import (
	"context"

	"cosmossdk.io/x/auth/types"

	sdk "github.com/T-ragon/cosmos-sdk/types"
)

// BankKeeper defines the expected interface contract the vesting module requires
// for creating vesting accounts with funds.
type BankKeeper interface {
	IsSendEnabledCoins(ctx context.Context, coins ...sdk.Coin) error
	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
	BlockedAddr(addr sdk.AccAddress) bool
}

type AccountsModKeeper interface {
	types.AccountsModKeeper
}
