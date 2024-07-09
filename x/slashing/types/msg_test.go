package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/T-ragon/cosmos-sdk/v3/codec"
	addresscodec "github.com/T-ragon/cosmos-sdk/v3/codec/address"
	"github.com/T-ragon/cosmos-sdk/v3/codec/types"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

func TestMsgUnjailGetSignBytes(t *testing.T) {
	addr := sdk.AccAddress("abcd")
	valAddrStr, err := addresscodec.NewBech32Codec("cosmosvaloper").BytesToString(addr)
	require.NoError(t, err)
	msg := NewMsgUnjail(valAddrStr)
	pc := codec.NewProtoCodec(types.NewInterfaceRegistry())
	bytes, err := pc.MarshalAminoJSON(msg)
	require.NoError(t, err)
	require.Equal(
		t,
		`{"type":"cosmos-sdk/MsgUnjail","value":{"address":"cosmosvaloper1v93xxeqhg9nn6"}}`,
		string(bytes),
	)
}
