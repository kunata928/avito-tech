// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/messages/messages.go

// Package mock_messages is a generated GoMock package.
package mock_messages

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
	rate "gitlab.ozon.dev/kunata928/telegramBot/internal/clients/rate"
	database "gitlab.ozon.dev/kunata928/telegramBot/internal/database"
)

// MockMessageSender is a mock of MessageSender interface.
type MockMessageSender struct {
	ctrl     *gomock.Controller
	recorder *MockMessageSenderMockRecorder
}

// MockMessageSenderMockRecorder is the mock recorder for MockMessageSender.
type MockMessageSenderMockRecorder struct {
	mock *MockMessageSender
}

// NewMockMessageSender creates a new mock instance.
func NewMockMessageSender(ctrl *gomock.Controller) *MockMessageSender {
	mock := &MockMessageSender{ctrl: ctrl}
	mock.recorder = &MockMessageSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageSender) EXPECT() *MockMessageSenderMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockMessageSender) SendMessage(text string, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", text, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockMessageSenderMockRecorder) SendMessage(text, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockMessageSender)(nil).SendMessage), text, userID)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddExpense mocks base method.
func (m *MockStorage) AddExpense(ctx context.Context, userID int64, expense *database.Expense) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddExpense", ctx, userID, expense)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddExpense indicates an expected call of AddExpense.
func (mr *MockStorageMockRecorder) AddExpense(ctx, userID, expense interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExpense", reflect.TypeOf((*MockStorage)(nil).AddExpense), ctx, userID, expense)
}

// GetClientCurrency mocks base method.
func (m *MockStorage) GetClientCurrency(ctx context.Context, userID int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientCurrency", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientCurrency indicates an expected call of GetClientCurrency.
func (mr *MockStorageMockRecorder) GetClientCurrency(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientCurrency", reflect.TypeOf((*MockStorage)(nil).GetClientCurrency), ctx, userID)
}

// GetClientExpenses mocks base method.
func (m *MockStorage) GetClientExpenses(ctx context.Context, userID int64, fromDate time.Time) ([]*database.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientExpenses", ctx, userID, fromDate)
	ret0, _ := ret[0].([]*database.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientExpenses indicates an expected call of GetClientExpenses.
func (mr *MockStorageMockRecorder) GetClientExpenses(ctx, userID, fromDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientExpenses", reflect.TypeOf((*MockStorage)(nil).GetClientExpenses), ctx, userID, fromDate)
}

// GetRate mocks base method.
func (m *MockStorage) GetRate(ctx context.Context, name string, date time.Time) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRate", ctx, name, date)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRate indicates an expected call of GetRate.
func (mr *MockStorageMockRecorder) GetRate(ctx, name, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRate", reflect.TypeOf((*MockStorage)(nil).GetRate), ctx, name, date)
}

// InitClient mocks base method.
func (m *MockStorage) InitClient(ctx context.Context, userID int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitClient", ctx, userID)
}

// InitClient indicates an expected call of InitClient.
func (mr *MockStorageMockRecorder) InitClient(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitClient", reflect.TypeOf((*MockStorage)(nil).InitClient), ctx, userID)
}

// RefreshClientCurrency mocks base method.
func (m *MockStorage) RefreshClientCurrency(ctx context.Context, userID int64, currency string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshClientCurrency", ctx, userID, currency)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshClientCurrency indicates an expected call of RefreshClientCurrency.
func (mr *MockStorageMockRecorder) RefreshClientCurrency(ctx, userID, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshClientCurrency", reflect.TypeOf((*MockStorage)(nil).RefreshClientCurrency), ctx, userID, currency)
}

// RefreshClientLimit mocks base method.
func (m *MockStorage) RefreshClientLimit(ctx context.Context, userID int64, amount decimal.Decimal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshClientLimit", ctx, userID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshClientLimit indicates an expected call of RefreshClientLimit.
func (mr *MockStorageMockRecorder) RefreshClientLimit(ctx, userID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshClientLimit", reflect.TypeOf((*MockStorage)(nil).RefreshClientLimit), ctx, userID, amount)
}

// SetRate mocks base method.
func (m *MockStorage) SetRate(ctx context.Context, rates *database.Rates, date time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRate", ctx, rates, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRate indicates an expected call of SetRate.
func (mr *MockStorageMockRecorder) SetRate(ctx, rates, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRate", reflect.TypeOf((*MockStorage)(nil).SetRate), ctx, rates, date)
}

// MockRateClient is a mock of RateClient interface.
type MockRateClient struct {
	ctrl     *gomock.Controller
	recorder *MockRateClientMockRecorder
}

// MockRateClientMockRecorder is the mock recorder for MockRateClient.
type MockRateClientMockRecorder struct {
	mock *MockRateClient
}

// NewMockRateClient creates a new mock instance.
func NewMockRateClient(ctrl *gomock.Controller) *MockRateClient {
	mock := &MockRateClient{ctrl: ctrl}
	mock.recorder = &MockRateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateClient) EXPECT() *MockRateClientMockRecorder {
	return m.recorder
}

// GetRateDate mocks base method.
func (m *MockRateClient) GetRateDate(ctx context.Context, date time.Time) (*rate.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRateDate", ctx, date)
	ret0, _ := ret[0].(*rate.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRateDate indicates an expected call of GetRateDate.
func (mr *MockRateClientMockRecorder) GetRateDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateDate", reflect.TypeOf((*MockRateClient)(nil).GetRateDate), ctx, date)
}

// MockLRU is a mock of LRU interface.
type MockLRU struct {
	ctrl     *gomock.Controller
	recorder *MockLRUMockRecorder
}

// MockLRUMockRecorder is the mock recorder for MockLRU.
type MockLRUMockRecorder struct {
	mock *MockLRU
}

// NewMockLRU creates a new mock instance.
func NewMockLRU(ctrl *gomock.Controller) *MockLRU {
	mock := &MockLRU{ctrl: ctrl}
	mock.recorder = &MockLRUMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLRU) EXPECT() *MockLRUMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockLRU) Add(key string, value interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", key, value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockLRUMockRecorder) Add(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockLRU)(nil).Add), key, value)
}

// Get mocks base method.
func (m *MockLRU) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockLRUMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLRU)(nil).Get), key)
}

// Len mocks base method.
func (m *MockLRU) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockLRUMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockLRU)(nil).Len))
}

// Remove mocks base method.
func (m *MockLRU) Remove(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockLRUMockRecorder) Remove(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockLRU)(nil).Remove), key)
}
