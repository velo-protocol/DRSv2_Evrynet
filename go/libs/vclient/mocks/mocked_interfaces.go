// Code generated by MockGen. DO NOT EDIT.
// Source: ./libs/vclient/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	ethereum "github.com/Evrynetlabs/evrynet-node"
	bind "github.com/Evrynetlabs/evrynet-node/accounts/abi/bind"
	common "github.com/Evrynetlabs/evrynet-node/common"
	types "github.com/Evrynetlabs/evrynet-node/core/types"
	gomock "github.com/golang/mock/gomock"
	vabi "github.com/velo-protocol/DRSv2_Evrynet/go/abi"
	big "math/big"
	reflect "reflect"
)

// MockConnection is a mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// CodeAt mocks base method
func (m *MockConnection) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, contract, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt
func (mr *MockConnectionMockRecorder) CodeAt(ctx, contract, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockConnection)(nil).CodeAt), ctx, contract, blockNumber)
}

// CallContract mocks base method
func (m *MockConnection) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, call, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract
func (mr *MockConnectionMockRecorder) CallContract(ctx, call, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockConnection)(nil).CallContract), ctx, call, blockNumber)
}

// PendingCodeAt mocks base method
func (m *MockConnection) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingCodeAt", ctx, account)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingCodeAt indicates an expected call of PendingCodeAt
func (mr *MockConnectionMockRecorder) PendingCodeAt(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingCodeAt", reflect.TypeOf((*MockConnection)(nil).PendingCodeAt), ctx, account)
}

// PendingNonceAt mocks base method
func (m *MockConnection) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingNonceAt", ctx, account)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingNonceAt indicates an expected call of PendingNonceAt
func (mr *MockConnectionMockRecorder) PendingNonceAt(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingNonceAt", reflect.TypeOf((*MockConnection)(nil).PendingNonceAt), ctx, account)
}

// SuggestGasPrice mocks base method
func (m *MockConnection) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasPrice", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasPrice indicates an expected call of SuggestGasPrice
func (mr *MockConnectionMockRecorder) SuggestGasPrice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasPrice", reflect.TypeOf((*MockConnection)(nil).SuggestGasPrice), ctx)
}

// EstimateGas mocks base method
func (m *MockConnection) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", ctx, call)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas
func (mr *MockConnectionMockRecorder) EstimateGas(ctx, call interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockConnection)(nil).EstimateGas), ctx, call)
}

// SendTransaction mocks base method
func (m *MockConnection) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockConnectionMockRecorder) SendTransaction(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockConnection)(nil).SendTransaction), ctx, tx)
}

// FilterLogs mocks base method
func (m *MockConnection) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterLogs", ctx, query)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterLogs indicates an expected call of FilterLogs
func (mr *MockConnectionMockRecorder) FilterLogs(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterLogs", reflect.TypeOf((*MockConnection)(nil).FilterLogs), ctx, query)
}

// SubscribeFilterLogs mocks base method
func (m *MockConnection) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeFilterLogs", ctx, query, ch)
	ret0, _ := ret[0].(ethereum.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeFilterLogs indicates an expected call of SubscribeFilterLogs
func (mr *MockConnectionMockRecorder) SubscribeFilterLogs(ctx, query, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeFilterLogs", reflect.TypeOf((*MockConnection)(nil).SubscribeFilterLogs), ctx, query, ch)
}

// TransactionReceipt mocks base method
func (m *MockConnection) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, txHash)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt
func (mr *MockConnectionMockRecorder) TransactionReceipt(ctx, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockConnection)(nil).TransactionReceipt), ctx, txHash)
}

// MockDRSContract is a mock of DRSContract interface
type MockDRSContract struct {
	ctrl     *gomock.Controller
	recorder *MockDRSContractMockRecorder
}

// MockDRSContractMockRecorder is the mock recorder for MockDRSContract
type MockDRSContractMockRecorder struct {
	mock *MockDRSContract
}

// NewMockDRSContract creates a new mock instance
func NewMockDRSContract(ctrl *gomock.Controller) *MockDRSContract {
	mock := &MockDRSContract{ctrl: ctrl}
	mock.recorder = &MockDRSContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDRSContract) EXPECT() *MockDRSContractMockRecorder {
	return m.recorder
}

// Setup mocks base method
func (m *MockDRSContract) Setup(opts *bind.TransactOpts, collateralAssetCode, peggedCurrency [32]byte, assetCode string, peggedValue *big.Int) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup", opts, collateralAssetCode, peggedCurrency, assetCode, peggedValue)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Setup indicates an expected call of Setup
func (mr *MockDRSContractMockRecorder) Setup(opts, collateralAssetCode, peggedCurrency, assetCode, peggedValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockDRSContract)(nil).Setup), opts, collateralAssetCode, peggedCurrency, assetCode, peggedValue)
}

// MintFromCollateralAmount mocks base method
func (m *MockDRSContract) MintFromCollateralAmount(opts *bind.TransactOpts, netCollateralAmount *big.Int, assetCode string) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintFromCollateralAmount", opts, netCollateralAmount, assetCode)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintFromCollateralAmount indicates an expected call of MintFromCollateralAmount
func (mr *MockDRSContractMockRecorder) MintFromCollateralAmount(opts, netCollateralAmount, assetCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintFromCollateralAmount", reflect.TypeOf((*MockDRSContract)(nil).MintFromCollateralAmount), opts, netCollateralAmount, assetCode)
}

// MintFromStableCreditAmount mocks base method
func (m *MockDRSContract) MintFromStableCreditAmount(opts *bind.TransactOpts, mintAmount *big.Int, assetCode string) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintFromStableCreditAmount", opts, mintAmount, assetCode)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintFromStableCreditAmount indicates an expected call of MintFromStableCreditAmount
func (mr *MockDRSContractMockRecorder) MintFromStableCreditAmount(opts, mintAmount, assetCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintFromStableCreditAmount", reflect.TypeOf((*MockDRSContract)(nil).MintFromStableCreditAmount), opts, mintAmount, assetCode)
}

// GetExchange mocks base method
func (m *MockDRSContract) GetExchange(opts *bind.CallOpts, assetCode string) (string, [32]byte, *big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchange", opts, assetCode)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([32]byte)
	ret2, _ := ret[2].(*big.Int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetExchange indicates an expected call of GetExchange
func (mr *MockDRSContractMockRecorder) GetExchange(opts, assetCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchange", reflect.TypeOf((*MockDRSContract)(nil).GetExchange), opts, assetCode)
}

// Redeem mocks base method
func (m *MockDRSContract) Redeem(opts *bind.TransactOpts, stableCreditAmount *big.Int, assetCode string) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Redeem", opts, stableCreditAmount, assetCode)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Redeem indicates an expected call of Redeem
func (mr *MockDRSContractMockRecorder) Redeem(opts, stableCreditAmount, assetCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redeem", reflect.TypeOf((*MockDRSContract)(nil).Redeem), opts, stableCreditAmount, assetCode)
}

// CollateralHealthCheck mocks base method
func (m *MockDRSContract) CollateralHealthCheck(opts *bind.CallOpts, assetCode string) (common.Address, [32]byte, *big.Int, *big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollateralHealthCheck", opts, assetCode)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].([32]byte)
	ret2, _ := ret[2].(*big.Int)
	ret3, _ := ret[3].(*big.Int)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// CollateralHealthCheck indicates an expected call of CollateralHealthCheck
func (mr *MockDRSContractMockRecorder) CollateralHealthCheck(opts, assetCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollateralHealthCheck", reflect.TypeOf((*MockDRSContract)(nil).CollateralHealthCheck), opts, assetCode)
}

// Rebalance mocks base method
func (m *MockDRSContract) Rebalance(opts *bind.TransactOpts, assetCode string) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebalance", opts, assetCode)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rebalance indicates an expected call of Rebalance
func (mr *MockDRSContractMockRecorder) Rebalance(opts, assetCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebalance", reflect.TypeOf((*MockDRSContract)(nil).Rebalance), opts, assetCode)
}

// MockHeartContract is a mock of HeartContract interface
type MockHeartContract struct {
	ctrl     *gomock.Controller
	recorder *MockHeartContractMockRecorder
}

// MockHeartContractMockRecorder is the mock recorder for MockHeartContract
type MockHeartContractMockRecorder struct {
	mock *MockHeartContract
}

// NewMockHeartContract creates a new mock instance
func NewMockHeartContract(ctrl *gomock.Controller) *MockHeartContract {
	mock := &MockHeartContract{ctrl: ctrl}
	mock.recorder = &MockHeartContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHeartContract) EXPECT() *MockHeartContractMockRecorder {
	return m.recorder
}

// SetGovernor mocks base method
func (m *MockHeartContract) SetGovernor(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGovernor", opts, address)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetGovernor indicates an expected call of SetGovernor
func (mr *MockHeartContractMockRecorder) SetGovernor(opts, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGovernor", reflect.TypeOf((*MockHeartContract)(nil).SetGovernor), opts, address)
}

// IsGovernor mocks base method
func (m *MockHeartContract) IsGovernor(opts *bind.CallOpts, addr common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsGovernor", opts, addr)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsGovernor indicates an expected call of IsGovernor
func (mr *MockHeartContractMockRecorder) IsGovernor(opts, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsGovernor", reflect.TypeOf((*MockHeartContract)(nil).IsGovernor), opts, addr)
}

// SetTrustedPartner mocks base method
func (m *MockHeartContract) SetTrustedPartner(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTrustedPartner", opts, address)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTrustedPartner indicates an expected call of SetTrustedPartner
func (mr *MockHeartContractMockRecorder) SetTrustedPartner(opts, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrustedPartner", reflect.TypeOf((*MockHeartContract)(nil).SetTrustedPartner), opts, address)
}

// IsTrustedPartner mocks base method
func (m *MockHeartContract) IsTrustedPartner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTrustedPartner", opts, addr)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsTrustedPartner indicates an expected call of IsTrustedPartner
func (mr *MockHeartContractMockRecorder) IsTrustedPartner(opts, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTrustedPartner", reflect.TypeOf((*MockHeartContract)(nil).IsTrustedPartner), opts, addr)
}

// GetStableCreditCount mocks base method
func (m *MockHeartContract) GetStableCreditCount(opts *bind.CallOpts) (uint8, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStableCreditCount", opts)
	ret0, _ := ret[0].(uint8)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStableCreditCount indicates an expected call of GetStableCreditCount
func (mr *MockHeartContractMockRecorder) GetStableCreditCount(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStableCreditCount", reflect.TypeOf((*MockHeartContract)(nil).GetStableCreditCount), opts)
}

// GetRecentStableCredit mocks base method
func (m *MockHeartContract) GetRecentStableCredit(opts *bind.CallOpts) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecentStableCredit", opts)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecentStableCredit indicates an expected call of GetRecentStableCredit
func (mr *MockHeartContractMockRecorder) GetRecentStableCredit(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecentStableCredit", reflect.TypeOf((*MockHeartContract)(nil).GetRecentStableCredit), opts)
}

// GetNextStableCredit mocks base method
func (m *MockHeartContract) GetNextStableCredit(opts *bind.CallOpts, stableCreditId [32]byte) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextStableCredit", opts, stableCreditId)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextStableCredit indicates an expected call of GetNextStableCredit
func (mr *MockHeartContractMockRecorder) GetNextStableCredit(opts, stableCreditId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextStableCredit", reflect.TypeOf((*MockHeartContract)(nil).GetNextStableCredit), opts, stableCreditId)
}

// MockStableCreditContract is a mock of StableCreditContract interface
type MockStableCreditContract struct {
	ctrl     *gomock.Controller
	recorder *MockStableCreditContractMockRecorder
}

// MockStableCreditContractMockRecorder is the mock recorder for MockStableCreditContract
type MockStableCreditContractMockRecorder struct {
	mock *MockStableCreditContract
}

// NewMockStableCreditContract creates a new mock instance
func NewMockStableCreditContract(ctrl *gomock.Controller) *MockStableCreditContract {
	mock := &MockStableCreditContract{ctrl: ctrl}
	mock.recorder = &MockStableCreditContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStableCreditContract) EXPECT() *MockStableCreditContractMockRecorder {
	return m.recorder
}

// StableCreditAssetCode mocks base method
func (m *MockStableCreditContract) StableCreditAssetCode(opts *bind.CallOpts) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StableCreditAssetCode", opts)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StableCreditAssetCode indicates an expected call of StableCreditAssetCode
func (mr *MockStableCreditContractMockRecorder) StableCreditAssetCode(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StableCreditAssetCode", reflect.TypeOf((*MockStableCreditContract)(nil).StableCreditAssetCode), opts)
}

// MockTxHelper is a mock of TxHelper interface
type MockTxHelper struct {
	ctrl     *gomock.Controller
	recorder *MockTxHelperMockRecorder
}

// MockTxHelperMockRecorder is the mock recorder for MockTxHelper
type MockTxHelperMockRecorder struct {
	mock *MockTxHelper
}

// NewMockTxHelper creates a new mock instance
func NewMockTxHelper(ctrl *gomock.Controller) *MockTxHelper {
	mock := &MockTxHelper{ctrl: ctrl}
	mock.recorder = &MockTxHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxHelper) EXPECT() *MockTxHelperMockRecorder {
	return m.recorder
}

// ConfirmTx mocks base method
func (m *MockTxHelper) ConfirmTx(ctx context.Context, tx *types.Transaction, from common.Address) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmTx", ctx, tx, from)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmTx indicates an expected call of ConfirmTx
func (mr *MockTxHelperMockRecorder) ConfirmTx(ctx, tx, from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmTx", reflect.TypeOf((*MockTxHelper)(nil).ConfirmTx), ctx, tx, from)
}

// ExtractSetupEvent mocks base method
func (m *MockTxHelper) ExtractSetupEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemSetup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractSetupEvent", eventName, log)
	ret0, _ := ret[0].(*vabi.DigitalReserveSystemSetup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractSetupEvent indicates an expected call of ExtractSetupEvent
func (mr *MockTxHelperMockRecorder) ExtractSetupEvent(eventName, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractSetupEvent", reflect.TypeOf((*MockTxHelper)(nil).ExtractSetupEvent), eventName, log)
}

// ExtractMintEvent mocks base method
func (m *MockTxHelper) ExtractMintEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemMint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractMintEvent", eventName, log)
	ret0, _ := ret[0].(*vabi.DigitalReserveSystemMint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractMintEvent indicates an expected call of ExtractMintEvent
func (mr *MockTxHelperMockRecorder) ExtractMintEvent(eventName, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractMintEvent", reflect.TypeOf((*MockTxHelper)(nil).ExtractMintEvent), eventName, log)
}

// ExtractRedeemEvent mocks base method
func (m *MockTxHelper) ExtractRedeemEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemRedeem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractRedeemEvent", eventName, log)
	ret0, _ := ret[0].(*vabi.DigitalReserveSystemRedeem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractRedeemEvent indicates an expected call of ExtractRedeemEvent
func (mr *MockTxHelperMockRecorder) ExtractRedeemEvent(eventName, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractRedeemEvent", reflect.TypeOf((*MockTxHelper)(nil).ExtractRedeemEvent), eventName, log)
}

// ExtractRebalanceEvent mocks base method
func (m *MockTxHelper) ExtractRebalanceEvent(eventName string, log *types.Log) (*vabi.DigitalReserveSystemRebalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractRebalanceEvent", eventName, log)
	ret0, _ := ret[0].(*vabi.DigitalReserveSystemRebalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractRebalanceEvent indicates an expected call of ExtractRebalanceEvent
func (mr *MockTxHelperMockRecorder) ExtractRebalanceEvent(eventName, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractRebalanceEvent", reflect.TypeOf((*MockTxHelper)(nil).ExtractRebalanceEvent), eventName, log)
}

// StableCreditAssetCode mocks base method
func (m *MockTxHelper) StableCreditAssetCode(addr common.Address) (*string, *[32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StableCreditAssetCode", addr)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(*[32]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StableCreditAssetCode indicates an expected call of StableCreditAssetCode
func (mr *MockTxHelperMockRecorder) StableCreditAssetCode(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StableCreditAssetCode", reflect.TypeOf((*MockTxHelper)(nil).StableCreditAssetCode), addr)
}
