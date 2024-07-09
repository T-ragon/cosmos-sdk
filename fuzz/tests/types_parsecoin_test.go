//go:build gofuzz || go1.18

package tests

import (
	"testing"

	"github.com/T-ragon/cosmos-sdk/v3/types"
)

func FuzzTypesParseCoin(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = types.ParseCoinNormalized(string(data))
	})
}
