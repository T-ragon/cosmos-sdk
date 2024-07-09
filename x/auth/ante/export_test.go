package ante

import (
	cryptotypes "github.com/T-ragon/cosmos-sdk/v3/crypto/types"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

var SimSecp256k1PubkeyInternal = simSecp256k1Pubkey

func SetSVDPubKey(svd SigVerificationDecorator, ctx sdk.Context, acc sdk.AccountI, txPubKey cryptotypes.PubKey) error {
	return svd.setPubKey(ctx, acc, txPubKey)
}
