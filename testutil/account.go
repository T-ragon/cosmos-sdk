package testutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/T-ragon/cosmos-sdk/v3/crypto/hd"
	"github.com/T-ragon/cosmos-sdk/v3/crypto/keyring"
	"github.com/T-ragon/cosmos-sdk/v3/types"
)

type TestAccount struct {
	Name    string
	Address types.AccAddress
}

func CreateKeyringAccounts(t *testing.T, kr keyring.Keyring, num int) []TestAccount {
	t.Helper()
	accounts := make([]TestAccount, num)
	for i := range accounts {
		record, _, err := kr.NewMnemonic(
			fmt.Sprintf("key-%d", i),
			keyring.English,
			types.FullFundraiserPath,
			keyring.DefaultBIP39Passphrase,
			hd.Secp256k1)
		assert.NoError(t, err)

		addr, err := record.GetAddress()
		assert.NoError(t, err)

		accounts[i] = TestAccount{Name: record.Name, Address: addr}
	}

	return accounts
}
