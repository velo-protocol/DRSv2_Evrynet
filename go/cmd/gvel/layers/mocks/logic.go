// Code generated by MockGen. DO NOT EDIT.
// Source: ./layers/logic/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/velo-protocol/DRSv2_Evrynet/go/cmd/gvel/entity"
	reflect "reflect"
)

// MockLogic is a mock of Logic interface
type MockLogic struct {
	ctrl     *gomock.Controller
	recorder *MockLogicMockRecorder
}

// MockLogicMockRecorder is the mock recorder for MockLogic
type MockLogicMockRecorder struct {
	mock *MockLogic
}

// NewMockLogic creates a new mock instance
func NewMockLogic(ctrl *gomock.Controller) *MockLogic {
	mock := &MockLogic{ctrl: ctrl}
	mock.recorder = &MockLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogic) EXPECT() *MockLogicMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockLogic) Init(configFilePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", configFilePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockLogicMockRecorder) Init(configFilePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockLogic)(nil).Init), configFilePath)
}

// CreateAccount mocks base method
func (m *MockLogic) CreateAccount(input *entity.CreateAccountInput) (*entity.CreateAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", input)
	ret0, _ := ret[0].(*entity.CreateAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockLogicMockRecorder) CreateAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockLogic)(nil).CreateAccount), input)
}

// ImportAccount mocks base method
func (m *MockLogic) ImportAccount(input *entity.ImportAccountInput) (*entity.ImportAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportAccount", input)
	ret0, _ := ret[0].(*entity.ImportAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportAccount indicates an expected call of ImportAccount
func (mr *MockLogicMockRecorder) ImportAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportAccount", reflect.TypeOf((*MockLogic)(nil).ImportAccount), input)
}

// ExportAccount mocks base method
func (m *MockLogic) ExportAccount(input *entity.ExportAccountInput) (*entity.ExportAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportAccount", input)
	ret0, _ := ret[0].(*entity.ExportAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportAccount indicates an expected call of ExportAccount
func (mr *MockLogicMockRecorder) ExportAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportAccount", reflect.TypeOf((*MockLogic)(nil).ExportAccount), input)
}

// ListAccount mocks base method
func (m *MockLogic) ListAccount() ([]*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccount")
	ret0, _ := ret[0].([]*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccount indicates an expected call of ListAccount
func (mr *MockLogicMockRecorder) ListAccount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccount", reflect.TypeOf((*MockLogic)(nil).ListAccount))
}

// SetDefaultAccount mocks base method
func (m *MockLogic) SetDefaultAccount(input *entity.SetDefaultAccountInput) (*entity.SetDefaultAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefaultAccount", input)
	ret0, _ := ret[0].(*entity.SetDefaultAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDefaultAccount indicates an expected call of SetDefaultAccount
func (mr *MockLogicMockRecorder) SetDefaultAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultAccount", reflect.TypeOf((*MockLogic)(nil).SetDefaultAccount), input)
}

// SetupCredit mocks base method
func (m *MockLogic) SetupCredit(input *entity.SetupCreditInput) (*entity.SetupCreditOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupCredit", input)
	ret0, _ := ret[0].(*entity.SetupCreditOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupCredit indicates an expected call of SetupCredit
func (mr *MockLogicMockRecorder) SetupCredit(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupCredit", reflect.TypeOf((*MockLogic)(nil).SetupCredit), input)
}

// MintCreditByCollateral mocks base method
func (m *MockLogic) MintCreditByCollateral(input *entity.MintCreditByCollateralInput) (*entity.MintCreditByCollateralOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCreditByCollateral", input)
	ret0, _ := ret[0].(*entity.MintCreditByCollateralOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintCreditByCollateral indicates an expected call of MintCreditByCollateral
func (mr *MockLogicMockRecorder) MintCreditByCollateral(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCreditByCollateral", reflect.TypeOf((*MockLogic)(nil).MintCreditByCollateral), input)
}

// MintCreditByCredit mocks base method
func (m *MockLogic) MintCreditByCredit(input *entity.MintCreditByCreditInput) (*entity.MintCreditByCreditOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCreditByCredit", input)
	ret0, _ := ret[0].(*entity.MintCreditByCreditOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintCreditByCredit indicates an expected call of MintCreditByCredit
func (mr *MockLogicMockRecorder) MintCreditByCredit(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCreditByCredit", reflect.TypeOf((*MockLogic)(nil).MintCreditByCredit), input)
}

// RedeemCredit mocks base method
func (m *MockLogic) RedeemCredit(input *entity.RedeemCreditInput) (*entity.RedeemCreditOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeemCredit", input)
	ret0, _ := ret[0].(*entity.RedeemCreditOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RedeemCredit indicates an expected call of RedeemCredit
func (mr *MockLogicMockRecorder) RedeemCredit(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeemCredit", reflect.TypeOf((*MockLogic)(nil).RedeemCredit), input)
}

// GetCreditExchange mocks base method
func (m *MockLogic) GetCreditExchange(input *entity.GetCreditExchangeInput) (*entity.GetCreditExchangeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreditExchange", input)
	ret0, _ := ret[0].(*entity.GetCreditExchangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreditExchange indicates an expected call of GetCreditExchange
func (mr *MockLogicMockRecorder) GetCreditExchange(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreditExchange", reflect.TypeOf((*MockLogic)(nil).GetCreditExchange), input)
}

// CollateralHealthCheck mocks base method
func (m *MockLogic) CollateralHealthCheck(input *entity.CollateralHealthCheckInput) ([]*entity.CollateralHealthCheckOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollateralHealthCheck", input)
	ret0, _ := ret[0].([]*entity.CollateralHealthCheckOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollateralHealthCheck indicates an expected call of CollateralHealthCheck
func (mr *MockLogicMockRecorder) CollateralHealthCheck(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollateralHealthCheck", reflect.TypeOf((*MockLogic)(nil).CollateralHealthCheck), input)
}

// RebalanceCollateral mocks base method
func (m *MockLogic) RebalanceCollateral(input *entity.RebalanceCollateralInput) ([]*entity.RebalanceCollateralOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RebalanceCollateral", input)
	ret0, _ := ret[0].([]*entity.RebalanceCollateralOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RebalanceCollateral indicates an expected call of RebalanceCollateral
func (mr *MockLogicMockRecorder) RebalanceCollateral(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebalanceCollateral", reflect.TypeOf((*MockLogic)(nil).RebalanceCollateral), input)
}

// SetEnv mocks base method
func (m *MockLogic) SetEnv(input *entity.SetEnvInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEnv", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEnv indicates an expected call of SetEnv
func (mr *MockLogicMockRecorder) SetEnv(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnv", reflect.TypeOf((*MockLogic)(nil).SetEnv), input)
}