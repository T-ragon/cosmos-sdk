package simsx

import (
	"context"
	"math/rand"

	"cosmossdk.io/core/address"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

type FactoryMethod func(ctx context.Context, testData *ChainDataSource, reporter SimulationReporter) (signer []SimAccount, msg sdk.Msg)

var _ SimMsgFactoryX = SimMsgFactoryFn[sdk.Msg](nil)

type SimMsgFactoryFn[T sdk.Msg] FactoryMethod

// MsgType returns an empty instance of type T, which implements `sdk.Msg`.
func (f SimMsgFactoryFn[T]) MsgType() sdk.Msg {
	var x T
	return x
}

func (f SimMsgFactoryFn[T]) Create() FactoryMethod {
	return FactoryMethod(f)
}

func (f SimMsgFactoryFn[T]) Cast(msg sdk.Msg) T {
	return msg.(T)
}

type SimMsgFactoryX interface {
	MsgType() sdk.Msg
	Create() FactoryMethod
}

// Registry is an abstract entry point to register message factories with weights
type Registry interface {
	Add(weight uint32, f SimMsgFactoryX)
}

type AccountSourceX interface {
	AccountSource
	ModuleAccountSource
}

var _ Registry = &SimsRegistryAdapter[any]{}

// SimsRegistryAdapter is an implementation of the Registry interface that provides adapters to use the new message factories
// with the legacy simulation system
type SimsRegistryAdapter[T any] struct {
	reporter     SimulationReporter
	legacyProps  []T
	ak           AccountSourceX
	bk           BalanceSource
	addressCodec address.Codec
	adapter      func(l *SimsRegistryAdapter[T], weight uint32, example sdk.Msg, f FactoryMethod) T
	txConfig     client.TxConfig
}

// NewSimsMsgRegistryAdapter creates a new instance of SimsRegistryAdapter for WeightedOperation types.
func NewSimsMsgRegistryAdapter(
	reporter SimulationReporter,
	ak AccountSourceX,
	bk BalanceSource,
	txConfig client.TxConfig,
) *SimsRegistryAdapter[simtypes.WeightedOperation] {
	return &SimsRegistryAdapter[simtypes.WeightedOperation]{
		reporter:     reporter,
		ak:           ak,
		bk:           bk,
		txConfig:     txConfig,
		addressCodec: txConfig.SigningContext().AddressCodec(),
		adapter:      LegacyOperationAdapter,
	}
}

// NewSimsProposalRegistryAdapter creates a new instance of SimsRegistryAdapter for WeightedProposalMsg types.
func NewSimsProposalRegistryAdapter(
	reporter SimulationReporter,
	ak AccountSourceX,
	bk BalanceSource,
	addrCodec address.Codec,
) *SimsRegistryAdapter[simtypes.WeightedProposalMsg] {
	return &SimsRegistryAdapter[simtypes.WeightedProposalMsg]{
		reporter:     reporter,
		ak:           ak,
		bk:           bk,
		addressCodec: addrCodec,
		adapter:      LegacyProposalMsgAdapter,
	}
}

// Add adds a new weighted operation to the SimsRegistryAdapter
func (l *SimsRegistryAdapter[T]) Add(weight uint32, f SimMsgFactoryX) {
	if f == nil {
		panic("message factory must not be nil")
	}
	if weight == 0 {
		return
	}
	l.legacyProps = append(l.legacyProps, l.adapter(l, weight, f.MsgType(), f.Create()))
}

// ToLegacy returns the legacy properties of the SimsRegistryAdapter as a slice of type T.
func (l *SimsRegistryAdapter[T]) ToLegacy() []T {
	return l.legacyProps
}

// LegacyOperationAdapter adapter to convert the new msg factory into the weighted operations type
func LegacyOperationAdapter(l *SimsRegistryAdapter[simtypes.WeightedOperation], weight uint32, example sdk.Msg, f FactoryMethod) simtypes.WeightedOperation {
	opAdapter := func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		xCtx, done := context.WithCancel(ctx)
		ctx = sdk.UnwrapSDKContext(xCtx)
		testData := NewChainDataSource(ctx, r, l.ak, l.bk, l.addressCodec, accs...)
		reporter := l.reporter.WithScope(example, SkipHookFn(func(args ...any) { done() }))
		from, msg := runWithFailFast(ctx, testData, reporter, f)
		return DeliverSimsMsg(ctx, reporter, app, r, l.txConfig, l.ak, chainID, msg, from...), nil, reporter.Close()
	}
	return simulation.NewWeightedOperation(int(weight), opAdapter)
}

// LegacyProposalMsgAdapter adapter to convert the new msg factory into the weighted proposal message type
func LegacyProposalMsgAdapter(l *SimsRegistryAdapter[simtypes.WeightedProposalMsg], weight uint32, example sdk.Msg, f FactoryMethod) simtypes.WeightedProposalMsg {
	msgAdapter := func(ctx context.Context, r *rand.Rand, accs []simtypes.Account, cdc address.Codec) (sdk.Msg, error) {
		xCtx, done := context.WithCancel(ctx)
		testData := NewChainDataSource(xCtx, r, l.ak, l.bk, l.addressCodec, accs...)
		reporter := l.reporter.WithScope(example, SkipHookFn(func(args ...any) { done() }))
		_, msg := runWithFailFast(xCtx, testData, reporter, f)
		return msg, nil
	}
	return simulation.NewWeightedProposalMsgX("", int(weight), msgAdapter)
}

type tuple struct {
	signer []SimAccount
	msg    sdk.Msg
}

// runWithFailFast runs the factory method on a separate goroutine to abort early when the context is canceled via reporter skip
func runWithFailFast(ctx context.Context, data *ChainDataSource, reporter SimulationReporter, f FactoryMethod) (signer []SimAccount, msg sdk.Msg) {
	r := make(chan tuple)
	go func() {
		defer recoverPanicForSkipped(reporter, r)
		signer, msg := f(ctx, data, reporter)
		r <- tuple{signer: signer, msg: msg}
	}()
	select {
	case t, ok := <-r:
		if !ok {
			return nil, nil
		}
		return t.signer, t.msg
	case <-ctx.Done():
		reporter.Skip("context closed")
		return nil, nil
	}
}

func recoverPanicForSkipped(reporter SimulationReporter, resultChan chan tuple) {
	if r := recover(); r != nil {
		if !reporter.IsSkipped() {
			panic(r)
		}
		close(resultChan)
	}
}
