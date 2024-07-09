package authz

import (
	"cosmossdk.io/x/authz/client/cli"

	"github.com/T-ragon/cosmos-sdk/client"
	"github.com/T-ragon/cosmos-sdk/testutil"
	clitestutil "github.com/T-ragon/cosmos-sdk/testutil/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	return clitestutil.ExecTestCLICmd(clientCtx, cli.NewCmdGrantAuthorization(), args)
}
