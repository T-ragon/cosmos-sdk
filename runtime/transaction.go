package runtime

import (
	"context"

	"cosmossdk.io/core/transaction"

	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

var _ transaction.Service = TransactionService{}

type TransactionService struct{}

// ExecMode implements transaction.Service.
func (t TransactionService) ExecMode(ctx context.Context) transaction.ExecMode {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return transaction.ExecMode(sdkCtx.ExecMode())
}
