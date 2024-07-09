package authz

import (
	"cosmossdk.io/x/authz/client/cli"

	"github.com/T-ragon/cosmos-sdk/v3/client"
	"github.com/T-ragon/cosmos-sdk/v3/testutil"
	clitestutil "github.com/T-ragon/cosmos-sdk/v3/testutil/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	return clitestutil.ExecTestCLICmd(clientCtx, cli.NewCmdGrantAuthorization(), args)
}
