package authz

import (
	"context"

	"cosmossdk.io/core/address"

	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	AddressCodec() address.Codec
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	IsSendEnabledCoins(ctx context.Context, coins ...sdk.Coin) error
}
