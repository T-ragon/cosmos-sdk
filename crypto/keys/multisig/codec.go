package multisig

import (
	"github.com/cometbft/cometbft/crypto/sr25519"

	"github.com/T-ragon/cosmos-sdk/v3/codec"
	bls12_381 "github.com/T-ragon/cosmos-sdk/v3/crypto/keys/bls12_381"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/keys/ed25519"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/keys/secp256k1"
	cryptotypes "github.com/T-ragon/cosmos-sdk/v3/crypto/types"
)

// TODO: Figure out API for others to either add their own pubkey types, or
// to make verify / marshal accept a AminoCdc.
const (
	// PubKeyAminoRoute defines the amino route for a multisig threshold public key
	PubKeyAminoRoute = "tendermint/PubKeyMultisigThreshold"
)

// AminoCdc is being deprecated in the SDK. But even if you need to
// use Amino for some reason, please use `codec/legacy.Cdc` instead.
var AminoCdc = codec.NewLegacyAmino()

func init() {
	AminoCdc.RegisterInterface((*cryptotypes.PubKey)(nil), nil)
	AminoCdc.RegisterConcrete(ed25519.PubKey{},
		ed25519.PubKeyName)
	AminoCdc.RegisterConcrete(sr25519.PubKey{},
		sr25519.PubKeyName)
	AminoCdc.RegisterConcrete(&secp256k1.PubKey{},
		secp256k1.PubKeyName)
	AminoCdc.RegisterConcrete(&bls12_381.PubKey{},
		bls12_381.PubKeyName)
	AminoCdc.RegisterConcrete(&LegacyAminoPubKey{},
		PubKeyAminoRoute)
}
