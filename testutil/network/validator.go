package network

import (
	"context"
	"net/http"

	"github.com/cometbft/cometbft/node"
	cmtclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"cosmossdk.io/log"

	"github.com/T-ragon/cosmos-sdk/v3/client"
	cryptotypes "github.com/T-ragon/cosmos-sdk/v3/crypto/types"
	"github.com/T-ragon/cosmos-sdk/v3/server/api"
	srvconfig "github.com/T-ragon/cosmos-sdk/v3/server/config"
	servertypes "github.com/T-ragon/cosmos-sdk/v3/server/types"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

// Validator defines an in-process CometBFT validator node. Through this object,
// a client can make RPC and API calls and interact with any client command
// or handler.
type Validator struct {
	AppConfig  *srvconfig.Config
	clientCtx  client.Context
	viper      *viper.Viper
	logger     log.Logger
	dir        string
	nodeID     string
	pubKey     cryptotypes.PubKey
	moniker    string
	aPIAddress string
	rPCAddress string
	p2PAddress string
	address    sdk.AccAddress
	valAddress sdk.ValAddress
	rPCClient  cmtclient.Client

	app      servertypes.Application
	tmNode   *node.Node
	api      *api.Server
	grpc     *grpc.Server
	grpcWeb  *http.Server
	errGroup *errgroup.Group
	cancelFn context.CancelFunc
}

var _ ValidatorI = &Validator{}

func (v *Validator) GetViper() *viper.Viper {
	return v.viper
}

func (v *Validator) GetLogger() log.Logger {
	return v.logger
}

func (v *Validator) GetClientCtx() client.Context {
	return v.clientCtx
}

func (v *Validator) GetAppConfig() *srvconfig.Config {
	return v.AppConfig
}

func (v *Validator) GetAddress() sdk.AccAddress {
	return v.address
}

func (v *Validator) GetValAddress() sdk.ValAddress {
	return v.valAddress
}

func (v *Validator) GetAPIAddress() string {
	return v.aPIAddress
}

func (v *Validator) GetRPCAddress() string {
	return v.rPCAddress
}

func (v *Validator) GetPubKey() cryptotypes.PubKey {
	return v.pubKey
}

func (v *Validator) GetMoniker() string {
	return v.moniker
}
