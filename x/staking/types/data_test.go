package types_test

import (
	"fmt"

	codectypes "github.com/T-ragon/cosmos-sdk/v3/codec/types"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/keys/ed25519"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

var (
	pk1      = ed25519.GenPrivKey().PubKey()
	pk1Any   *codectypes.Any
	pk2      = ed25519.GenPrivKey().PubKey()
	pk3      = ed25519.GenPrivKey().PubKey()
	valAddr1 = sdk.ValAddress(pk1.Address())
	valAddr2 = sdk.ValAddress(pk2.Address())
	valAddr3 = sdk.ValAddress(pk3.Address())
)

func init() {
	var err error
	pk1Any, err = codectypes.NewAnyWithValue(pk1)
	if err != nil {
		panic(fmt.Sprintf("Can't pack pk1 %t as Any", pk1))
	}
}
