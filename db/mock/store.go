// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dorasaicu12/simplebank/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/dorasaicu12/simplebank/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockStore) AddAccountBalance(arg0 context.Context, arg1 db.AddAccountBalanceParams) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockStoreMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockStore)(nil).AddAccountBalance), arg0, arg1)
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateEntrie mocks base method.
func (m *MockStore) CreateEntrie(arg0 context.Context, arg1 db.CreateEntrieParams) (db.Entries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntrie", arg0, arg1)
	ret0, _ := ret[0].(db.Entries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntrie indicates an expected call of CreateEntrie.
func (mr *MockStoreMockRecorder) CreateEntrie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntrie", reflect.TypeOf((*MockStore)(nil).CreateEntrie), arg0, arg1)
}

// CreateTransfers mocks base method.
func (m *MockStore) CreateTransfers(arg0 context.Context, arg1 db.CreateTransfersParams) (db.Transfers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfers", arg0, arg1)
	ret0, _ := ret[0].(db.Transfers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfers indicates an expected call of CreateTransfers.
func (mr *MockStoreMockRecorder) CreateTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfers", reflect.TypeOf((*MockStore)(nil).CreateTransfers), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// DeleteEntrie mocks base method.
func (m *MockStore) DeleteEntrie(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntrie", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntrie indicates an expected call of DeleteEntrie.
func (mr *MockStoreMockRecorder) DeleteEntrie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntrie", reflect.TypeOf((*MockStore)(nil).DeleteEntrie), arg0, arg1)
}

// DeleteTransfers mocks base method.
func (m *MockStore) DeleteTransfers(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransfers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransfers indicates an expected call of DeleteTransfers.
func (mr *MockStoreMockRecorder) DeleteTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransfers", reflect.TypeOf((*MockStore)(nil).DeleteTransfers), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int64) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetAccountForUpdate mocks base method.
func (m *MockStore) GetAccountForUpdate(arg0 context.Context, arg1 int64) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockStoreMockRecorder) GetAccountForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountForUpdate), arg0, arg1)
}

// GetEntrie mocks base method.
func (m *MockStore) GetEntrie(arg0 context.Context, arg1 int64) (db.Entries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntrie", arg0, arg1)
	ret0, _ := ret[0].(db.Entries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntrie indicates an expected call of GetEntrie.
func (mr *MockStoreMockRecorder) GetEntrie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntrie", reflect.TypeOf((*MockStore)(nil).GetEntrie), arg0, arg1)
}

// GetListAccount mocks base method.
func (m *MockStore) GetListAccount(arg0 context.Context, arg1 db.GetListAccountParams) ([]db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListAccount", arg0, arg1)
	ret0, _ := ret[0].([]db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListAccount indicates an expected call of GetListAccount.
func (mr *MockStoreMockRecorder) GetListAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListAccount", reflect.TypeOf((*MockStore)(nil).GetListAccount), arg0, arg1)
}

// GetListEntrie mocks base method.
func (m *MockStore) GetListEntrie(arg0 context.Context, arg1 db.GetListEntrieParams) ([]db.Entries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListEntrie", arg0, arg1)
	ret0, _ := ret[0].([]db.Entries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListEntrie indicates an expected call of GetListEntrie.
func (mr *MockStoreMockRecorder) GetListEntrie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListEntrie", reflect.TypeOf((*MockStore)(nil).GetListEntrie), arg0, arg1)
}

// GetListTransfers mocks base method.
func (m *MockStore) GetListTransfers(arg0 context.Context, arg1 db.GetListTransfersParams) ([]db.Transfers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]db.Transfers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListTransfers indicates an expected call of GetListTransfers.
func (mr *MockStoreMockRecorder) GetListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListTransfers", reflect.TypeOf((*MockStore)(nil).GetListTransfers), arg0, arg1)
}

// GetTransfers mocks base method.
func (m *MockStore) GetTransfers(arg0 context.Context, arg1 int64) (db.Transfers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfers", arg0, arg1)
	ret0, _ := ret[0].(db.Transfers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfers indicates an expected call of GetTransfers.
func (mr *MockStoreMockRecorder) GetTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfers", reflect.TypeOf((*MockStore)(nil).GetTransfers), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(arg0 context.Context, arg1 db.TransferTxParams) (db.TransferResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransferResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdateEntrie mocks base method.
func (m *MockStore) UpdateEntrie(arg0 context.Context, arg1 db.UpdateEntrieParams) (db.Entries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntrie", arg0, arg1)
	ret0, _ := ret[0].(db.Entries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntrie indicates an expected call of UpdateEntrie.
func (mr *MockStoreMockRecorder) UpdateEntrie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntrie", reflect.TypeOf((*MockStore)(nil).UpdateEntrie), arg0, arg1)
}

// UpdateTransfers mocks base method.
func (m *MockStore) UpdateTransfers(arg0 context.Context, arg1 db.UpdateTransfersParams) (db.Transfers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransfers", arg0, arg1)
	ret0, _ := ret[0].(db.Transfers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransfers indicates an expected call of UpdateTransfers.
func (mr *MockStoreMockRecorder) UpdateTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransfers", reflect.TypeOf((*MockStore)(nil).UpdateTransfers), arg0, arg1)
}
