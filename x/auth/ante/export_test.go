package ante

import (
	cryptotypes "github.com/T-ragon/cosmos-sdk/crypto/types"
	sdk "github.com/T-ragon/cosmos-sdk/types"
)

var SimSecp256k1PubkeyInternal = simSecp256k1Pubkey

func SetSVDPubKey(svd SigVerificationDecorator, ctx sdk.Context, acc sdk.AccountI, txPubKey cryptotypes.PubKey) error {
	return svd.setPubKey(ctx, acc, txPubKey)
}
