package signing_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	authsign "cosmossdk.io/x/auth/signing"

	codectestutil "github.com/T-ragon/cosmos-sdk/v3/codec/testutil"
	"github.com/T-ragon/cosmos-sdk/v3/testutil/testdata"
	moduletestutil "github.com/T-ragon/cosmos-sdk/v3/types/module/testutil"
	"github.com/T-ragon/cosmos-sdk/v3/types/tx/signing"
)

func TestGetSignBytesAdapterNoPublicKey(t *testing.T) {
	encodingConfig := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{})
	txConfig := encodingConfig.TxConfig
	_, _, addr := testdata.KeyTestPubAddr()
	signerData := authsign.SignerData{
		Address:       addr.String(),
		ChainID:       "test-chain",
		AccountNumber: 11,
		Sequence:      15,
	}
	w := txConfig.NewTxBuilder()
	_, err := authsign.GetSignBytesAdapter(
		context.Background(),
		txConfig.SignModeHandler(),
		signing.SignMode_SIGN_MODE_DIRECT,
		signerData,
		w.GetTx())
	require.NoError(t, err)
}
