// Code generated by MockGen. DO NOT EDIT.
// Source: internal/database/database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
	database "gitlab.ozon.dev/kunata928/telegramBot/internal/database"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// AddExpense mocks base method.
func (m *MockDB) AddExpense(ctx context.Context, userID int64, newExpense *database.Expense) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddExpense", ctx, userID, newExpense)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddExpense indicates an expected call of AddExpense.
func (mr *MockDBMockRecorder) AddExpense(ctx, userID, newExpense interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExpense", reflect.TypeOf((*MockDB)(nil).AddExpense), ctx, userID, newExpense)
}

// AddRate mocks base method.
func (m *MockDB) AddRate(ctx context.Context, rates *database.Rates, date time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRate", ctx, rates, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRate indicates an expected call of AddRate.
func (mr *MockDBMockRecorder) AddRate(ctx, rates, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRate", reflect.TypeOf((*MockDB)(nil).AddRate), ctx, rates, date)
}

// GetClientCurrency mocks base method.
func (m *MockDB) GetClientCurrency(ctx context.Context, userID int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientCurrency", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientCurrency indicates an expected call of GetClientCurrency.
func (mr *MockDBMockRecorder) GetClientCurrency(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientCurrency", reflect.TypeOf((*MockDB)(nil).GetClientCurrency), ctx, userID)
}

// GetClientExpenses mocks base method.
func (m *MockDB) GetClientExpenses(ctx context.Context, userID int64, fromDate time.Time) ([]*database.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientExpenses", ctx, userID, fromDate)
	ret0, _ := ret[0].([]*database.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientExpenses indicates an expected call of GetClientExpenses.
func (mr *MockDBMockRecorder) GetClientExpenses(ctx, userID, fromDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientExpenses", reflect.TypeOf((*MockDB)(nil).GetClientExpenses), ctx, userID, fromDate)
}

// GetRate mocks base method.
func (m *MockDB) GetRate(ctx context.Context, name string, date time.Time) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRate", ctx, name, date)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRate indicates an expected call of GetRate.
func (mr *MockDBMockRecorder) GetRate(ctx, name, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRate", reflect.TypeOf((*MockDB)(nil).GetRate), ctx, name, date)
}

// InitClient mocks base method.
func (m *MockDB) InitClient(ctx context.Context, userID int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitClient", ctx, userID)
}

// InitClient indicates an expected call of InitClient.
func (mr *MockDBMockRecorder) InitClient(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitClient", reflect.TypeOf((*MockDB)(nil).InitClient), ctx, userID)
}

// RefreshClientCurrency mocks base method.
func (m *MockDB) RefreshClientCurrency(ctx context.Context, userID int64, currency string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshClientCurrency", ctx, userID, currency)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshClientCurrency indicates an expected call of RefreshClientCurrency.
func (mr *MockDBMockRecorder) RefreshClientCurrency(ctx, userID, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshClientCurrency", reflect.TypeOf((*MockDB)(nil).RefreshClientCurrency), ctx, userID, currency)
}

// RefreshClientLimit mocks base method.
func (m *MockDB) RefreshClientLimit(ctx context.Context, userID int64, amount decimal.Decimal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshClientLimit", ctx, userID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshClientLimit indicates an expected call of RefreshClientLimit.
func (mr *MockDBMockRecorder) RefreshClientLimit(ctx, userID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshClientLimit", reflect.TypeOf((*MockDB)(nil).RefreshClientLimit), ctx, userID, amount)
}