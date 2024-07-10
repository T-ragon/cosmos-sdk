package authz

import (
	"github.com/T-ragon/cosmos-sdk/client"
	addresscodec "github.com/T-ragon/cosmos-sdk/codec/address"
	"github.com/T-ragon/cosmos-sdk/testutil"
	clitestutil "github.com/T-ragon/cosmos-sdk/testutil/cli"
	"github.com/T-ragon/cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization(addresscodec.NewBech32Codec("cosmos"))
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
