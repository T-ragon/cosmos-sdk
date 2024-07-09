package legacy_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/T-ragon/cosmos-sdk/v3/codec"
	"github.com/T-ragon/cosmos-sdk/v3/codec/legacy"
	"github.com/T-ragon/cosmos-sdk/v3/testutil/testdata"
)

func TestRegisterAminoMsg(t *testing.T) {
	cdc := codec.NewLegacyAmino()

	testCases := map[string]struct {
		msgName  string
		expPanic bool
	}{
		"all good": {
			msgName: "cosmos-sdk/Test",
		},
		"msgName too long": {
			msgName:  strings.Repeat("a", 40),
			expPanic: true,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			fn := func() { legacy.RegisterAminoMsg(cdc, &testdata.TestMsg{}, tc.msgName) }
			if tc.expPanic {
				require.Panics(t, fn)
			} else {
				require.NotPanics(t, fn)
			}
		})
	}
}
