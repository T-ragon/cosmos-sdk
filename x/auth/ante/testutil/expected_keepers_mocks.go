// Code generated by MockGen. DO NOT EDIT.
// Source: x/auth/ante/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	address "cosmossdk.io/core/address"
	appmodule "cosmossdk.io/core/appmodule"
	types "cosmossdk.io/x/auth/types"
	types0 "cosmossdk.io/x/consensus/types"
	types1 "github.com/T-ragon/cosmos-sdk/v3/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// AddressCodec mocks base method.
func (m *MockAccountKeeper) AddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// AddressCodec indicates an expected call of AddressCodec.
func (mr *MockAccountKeeperMockRecorder) AddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressCodec", reflect.TypeOf((*MockAccountKeeper)(nil).AddressCodec))
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr types1.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetEnvironment mocks base method.
func (m *MockAccountKeeper) GetEnvironment() appmodule.Environment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironment")
	ret0, _ := ret[0].(appmodule.Environment)
	return ret0
}

// GetEnvironment indicates an expected call of GetEnvironment.
func (mr *MockAccountKeeperMockRecorder) GetEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockAccountKeeper)(nil).GetEnvironment))
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types1.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types1.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// GetParams mocks base method.
func (m *MockAccountKeeper) GetParams(ctx context.Context) types.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockAccountKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockAccountKeeper)(nil).GetParams), ctx)
}

// NewAccountWithAddress mocks base method.
func (m *MockAccountKeeper) NewAccountWithAddress(ctx context.Context, addr types1.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccountWithAddress", ctx, addr)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// NewAccountWithAddress indicates an expected call of NewAccountWithAddress.
func (mr *MockAccountKeeperMockRecorder) NewAccountWithAddress(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccountWithAddress", reflect.TypeOf((*MockAccountKeeper)(nil).NewAccountWithAddress), ctx, addr)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx context.Context, acc types1.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}

// MockFeegrantKeeper is a mock of FeegrantKeeper interface.
type MockFeegrantKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockFeegrantKeeperMockRecorder
}

// MockFeegrantKeeperMockRecorder is the mock recorder for MockFeegrantKeeper.
type MockFeegrantKeeperMockRecorder struct {
	mock *MockFeegrantKeeper
}

// NewMockFeegrantKeeper creates a new mock instance.
func NewMockFeegrantKeeper(ctrl *gomock.Controller) *MockFeegrantKeeper {
	mock := &MockFeegrantKeeper{ctrl: ctrl}
	mock.recorder = &MockFeegrantKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeegrantKeeper) EXPECT() *MockFeegrantKeeperMockRecorder {
	return m.recorder
}

// UseGrantedFees mocks base method.
func (m *MockFeegrantKeeper) UseGrantedFees(ctx context.Context, granter, grantee types1.AccAddress, fee types1.Coins, msgs []types1.Msg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseGrantedFees", ctx, granter, grantee, fee, msgs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseGrantedFees indicates an expected call of UseGrantedFees.
func (mr *MockFeegrantKeeperMockRecorder) UseGrantedFees(ctx, granter, grantee, fee, msgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseGrantedFees", reflect.TypeOf((*MockFeegrantKeeper)(nil).UseGrantedFees), ctx, granter, grantee, fee, msgs)
}

// MockConsensusKeeper is a mock of ConsensusKeeper interface.
type MockConsensusKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusKeeperMockRecorder
}

// MockConsensusKeeperMockRecorder is the mock recorder for MockConsensusKeeper.
type MockConsensusKeeperMockRecorder struct {
	mock *MockConsensusKeeper
}

// NewMockConsensusKeeper creates a new mock instance.
func NewMockConsensusKeeper(ctrl *gomock.Controller) *MockConsensusKeeper {
	mock := &MockConsensusKeeper{ctrl: ctrl}
	mock.recorder = &MockConsensusKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsensusKeeper) EXPECT() *MockConsensusKeeperMockRecorder {
	return m.recorder
}

// GetCometInfo mocks base method.
func (m *MockConsensusKeeper) GetCometInfo(ctx context.Context, request *types0.QueryGetCometInfoRequest) (*types0.QueryGetCometInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCometInfo", ctx, request)
	ret0, _ := ret[0].(*types0.QueryGetCometInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCometInfo indicates an expected call of GetCometInfo.
func (mr *MockConsensusKeeperMockRecorder) GetCometInfo(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCometInfo", reflect.TypeOf((*MockConsensusKeeper)(nil).GetCometInfo), ctx, request)
}

// Params mocks base method.
func (m *MockConsensusKeeper) Params(arg0 context.Context, arg1 *types0.QueryParamsRequest) (*types0.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Params", arg0, arg1)
	ret0, _ := ret[0].(*types0.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Params indicates an expected call of Params.
func (mr *MockConsensusKeeperMockRecorder) Params(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Params", reflect.TypeOf((*MockConsensusKeeper)(nil).Params), arg0, arg1)
}
