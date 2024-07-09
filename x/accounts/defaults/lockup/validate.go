package lockup

import (
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	sdkerrors "github.com/T-ragon/cosmos-sdk/v3/types/errors"
)

func validateAmount(amount sdk.Coins) error {
	if !amount.IsValid() {
		return sdkerrors.ErrInvalidCoins.Wrap(amount.String())
	}

	if amount.IsZero() {
		return sdkerrors.ErrInvalidCoins.Wrap(amount.String())
	}

	return nil
}
