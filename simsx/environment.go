package simsx

import (
	"context"
	"errors"
	"math/rand"
	"slices"
	"time"

	"cosmossdk.io/core/address"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

type contextAwareBalanceSource struct {
	ctx  context.Context
	bank BalanceSource
}

func (s contextAwareBalanceSource) SpendableCoins(accAddress sdk.AccAddress) sdk.Coins {
	return s.bank.SpendableCoins(s.ctx, accAddress)
}

func (s contextAwareBalanceSource) IsSendEnabledDenom(denom string) bool {
	return s.bank.IsSendEnabledDenom(s.ctx, denom)
}

type SimAccount struct {
	simtypes.Account
	r             *rand.Rand
	liquidBalance *SimsAccountBalance
	bank          contextAwareBalanceSource
}

func (a *SimAccount) LiquidBalance() *SimsAccountBalance {
	if a.liquidBalance == nil {
		a.liquidBalance = NewSimsAccountBalance(a, a.r, a.bank.SpendableCoins(a.Address))
	}
	return a.liquidBalance
}

type SimsAccountBalance struct {
	owner *SimAccount
	sdk.Coins
	r *rand.Rand
}

func NewSimsAccountBalance(o *SimAccount, r *rand.Rand, coins sdk.Coins) *SimsAccountBalance {
	return &SimsAccountBalance{Coins: coins, r: r, owner: o}
}

type CoinsFilter interface {
	Accept(c sdk.Coins) bool // returns false to reject
}
type CoinsFilterFn func(c sdk.Coins) bool

func (f CoinsFilterFn) Accept(c sdk.Coins) bool {
	return f(c)
}

func WithSendEnabledCoins() CoinsFilter {
	return statefulCoinsFilterFn(func(s *SimAccount, coins sdk.Coins) bool {
		for _, coin := range coins {
			if !s.bank.IsSendEnabledDenom(coin.Denom) {
				return false
			}
		}
		return true
	})
}

type statefulCoinsFilter struct {
	s  *SimAccount
	do func(s *SimAccount, c sdk.Coins) bool
}

func statefulCoinsFilterFn(f func(s *SimAccount, c sdk.Coins) bool) CoinsFilter {
	return &statefulCoinsFilter{do: f}
}

func (f statefulCoinsFilter) Accept(c sdk.Coins) bool {
	if f.s == nil {
		panic("account not set")
	}
	return f.do(f.s, c)
}

func (f *statefulCoinsFilter) visit(s *SimAccount) {
	f.s = s
}

var _ visitable = &statefulCoinsFilter{}

type visitable interface {
	visit(s *SimAccount)
}

// RandSubsetCoins return random amounts from the current balance. When the coins are empty, skip is called on the reporter.
// The amounts are removed from the liquid balance.
func (b *SimsAccountBalance) RandSubsetCoins(reporter SimulationReporter, filters ...CoinsFilter) sdk.Coins {
	amount := b.randomAmount(5, reporter, b.Coins, filters...)

	b.Coins = b.Coins.Sub(amount...)
	if amount.Empty() {
		reporter.Skip("got empty amounts")
	}
	return amount
}

func (b *SimsAccountBalance) RandSubsetCoin(reporter SimulationReporter, denom string, filters ...CoinsFilter) sdk.Coin {
	ok, coin := b.Find(denom)
	if !ok {
		reporter.Skipf("no such coin: %s", denom)
		return sdk.NewCoin(denom, math.ZeroInt())
	}
	amount := b.randomAmount(5, reporter, sdk.Coins{coin}, filters...)
	b.Coins = b.Coins.Sub(amount...)
	ok, coin = amount.Find(denom)
	if !ok || !coin.IsPositive() {
		reporter.Skip("empty coin")
		return sdk.NewCoin(denom, math.ZeroInt())
	}
	return coin
}

func (b *SimsAccountBalance) randomAmount(retryCount int, reporter SimulationReporter, coins sdk.Coins, filters ...CoinsFilter) sdk.Coins {
	if retryCount < 0 || b.Coins.Empty() {
		reporter.Skip("failed to find matching amount")
		return sdk.Coins{}
	}
	amount := simtypes.RandSubsetCoins(b.r, coins)
	for _, filter := range filters {
		if f, ok := filter.(visitable); ok {
			f.visit(b.owner)
		}
		if !filter.Accept(amount) {
			return b.randomAmount(retryCount-1, reporter, coins, filters...)
		}
	}
	return amount
}

func (b *SimsAccountBalance) RandFees() sdk.Coins {
	amount, err := simtypes.RandomFees(b.r, b.Coins)
	if err != nil {
		panic(err.Error()) // todo: setup a better way to abort execution
	} // todo: revisit the panic
	return amount
}

type SimAccountFilter interface {
	// Accept returns true to accept the account or false to reject
	Accept(a SimAccount) bool
}
type SimAccountFilterFn func(a SimAccount) bool

func (s SimAccountFilterFn) Accept(a SimAccount) bool {
	return s(a)
}

func ExcludeAccounts(others ...SimAccount) SimAccountFilter {
	return SimAccountFilterFn(func(a SimAccount) bool {
		return !slices.ContainsFunc(others, func(o SimAccount) bool {
			return a.Address.Equals(o.Address)
		})
	})
}

func ExcludeAddresses(addrs ...string) SimAccountFilter {
	return SimAccountFilterFn(func(a SimAccount) bool {
		return !slices.Contains(addrs, a.AddressBech32)
	})
}

func WithDenomBalance(denom string) SimAccountFilter {
	return SimAccountFilterFn(func(a SimAccount) bool {
		return a.LiquidBalance().AmountOf(denom).IsPositive()
	})
}
func WithLiquidBalanceGTE(amount ...sdk.Coin) SimAccountFilter {
	return SimAccountFilterFn(func(a SimAccount) bool {
		return a.LiquidBalance().IsAllGTE(amount)
	})
}

// todo: liquid token but sent may not be enabled for all or any
func WithSpendableBalance() SimAccountFilter {
	return SimAccountFilterFn(func(a SimAccount) bool {
		return !a.LiquidBalance().Empty()
	})
}

type ModuleAccountSource interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
}
type SpendableCoinser interface {
	SpendableCoins(addr sdk.AccAddress) sdk.Coins
}

type SpendableCoinserFn func(addr sdk.AccAddress) sdk.Coins

func (b SpendableCoinserFn) SpendableCoins(addr sdk.AccAddress) sdk.Coins {
	return b(addr)
}

type BalanceSource interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	IsSendEnabledDenom(ctx context.Context, denom string) bool
}

func NewBalanceSource(ctx sdk.Context, bk BalanceSource,
) SpendableCoinser {
	return SpendableCoinserFn(func(addr sdk.AccAddress) sdk.Coins {
		return bk.SpendableCoins(ctx, addr)
	})
}

type ChainDataSource struct {
	r                         *rand.Rand
	addressToAccountsPosIndex map[string]int
	accounts                  []SimAccount
	accountSource             ModuleAccountSource
	addressCodec              address.Codec
}

func NewChainDataSource(ctx context.Context, r *rand.Rand, ak ModuleAccountSource, bk BalanceSource, codec address.Codec, oldSimAcc ...simtypes.Account) *ChainDataSource {
	acc := make([]SimAccount, len(oldSimAcc))
	index := make(map[string]int, len(oldSimAcc))
	for i, a := range oldSimAcc {
		acc[i] = SimAccount{
			Account: a,
			r:       r,
			bank:    contextAwareBalanceSource{ctx: ctx, bank: bk},
		}
		index[a.AddressBech32] = i
	}
	return &ChainDataSource{r: r, accountSource: ak, addressCodec: codec, accounts: acc}
}

// no module accounts
func (c *ChainDataSource) AnyAccount(r SimulationReporter, filters ...SimAccountFilter) SimAccount {
	acc := c.randomAccount(r, 5, filters...)
	return acc
}

func (c ChainDataSource) GetAccount(reporter SimulationReporter, addr string) SimAccount {
	pos, ok := c.addressToAccountsPosIndex[addr]
	if !ok {
		reporter.Skipf("no such account: %s", addr)
		return SimAccount{}
	}
	return c.accounts[pos]
}

func (c *ChainDataSource) randomAccount(reporter SimulationReporter, retryCount int, filters ...SimAccountFilter) SimAccount {
	if retryCount < 0 {
		reporter.Skip("failed to find a matching account")
		return SimAccount{
			Account:       simtypes.Account{},
			r:             c.r,
			liquidBalance: &SimsAccountBalance{},
			bank:          c.accounts[0].bank,
		}
	}
	idx := c.r.Intn(len(c.accounts))
	acc := c.accounts[idx]
	for _, filter := range filters {
		if !filter.Accept(acc) {
			return c.randomAccount(reporter, retryCount-1, filters...)
		}
	}
	return acc
}

func (c *ChainDataSource) ModuleAccountAddress(reporter SimulationReporter, moduleName string) string {
	acc := c.accountSource.GetModuleAddress(moduleName)
	if acc == nil {
		reporter.Skipf("unknown module account: %s", moduleName)
		return ""
	}
	res, err := c.addressCodec.BytesToString(acc)
	if err != nil {
		reporter.Skipf("failed to encode module address: %s", err)
		return ""
	}
	return res
}

func (c *ChainDataSource) Rand() *XRand {
	return &XRand{c.r}
}

type XRand struct {
	*rand.Rand
}

func NewXRand(rand *rand.Rand) *XRand {
	return &XRand{Rand: rand}
}

func (r *XRand) StringN(max int) string {
	return simtypes.RandStringOfLength(r.Rand, max)
}

func (r *XRand) SubsetCoins(src sdk.Coins) sdk.Coins {
	return simtypes.RandSubsetCoins(r.Rand, src)
}

func (r *XRand) DecN(max math.LegacyDec) math.LegacyDec {
	return simtypes.RandomDecAmount(r.Rand, max)
}

func (r *XRand) IntInRange(min, max int) int {
	return r.Rand.Intn(max-min) + min
}

// Uint64InRange returns a pseudo-random uint64 number in the range [min, max].
// It panics when min >= max
func (r *XRand) Uint64InRange(min, max uint64) uint64 {
	return uint64(r.Rand.Int63n(int64(max-min)) + int64(min))
}

func (r *XRand) PositiveSDKIntn(max math.Int) (math.Int, error) {
	return simtypes.RandPositiveInt(r.Rand, max)
}

func (r *XRand) PositiveSDKIntInRange(min, max math.Int) (math.Int, error) {
	diff := max.Sub(min)
	if !diff.IsPositive() {
		return math.Int{}, errors.New("min value must not be greater or equal to max")
	}
	result, err := r.PositiveSDKIntn(diff)
	if err != nil {
		return math.Int{}, err
	}
	return result.Add(min), nil
}

// Timestamp returns a timestamp between  Jan 1, 2062 and Jan 1, 2262
func (r *XRand) Timestamp() time.Time {
	return simtypes.RandTimestamp(r.Rand)
}

func (r *XRand) Bool() bool {
	return r.Intn(100) > 50
}
