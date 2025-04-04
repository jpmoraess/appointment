// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jpmoraess/appointment-api/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/jpmoraess/appointment-api/db/sqlc"
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

// CreateTenant mocks base method.
func (m *MockStore) CreateTenant(arg0 context.Context, arg1 db.CreateTenantParams) (db.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", arg0, arg1)
	ret0, _ := ret[0].(db.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant.
func (mr *MockStoreMockRecorder) CreateTenant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockStore)(nil).CreateTenant), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// GetTenant mocks base method.
func (m *MockStore) GetTenant(arg0 context.Context, arg1 int64) (db.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenant", arg0, arg1)
	ret0, _ := ret[0].(db.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTenant indicates an expected call of GetTenant.
func (mr *MockStoreMockRecorder) GetTenant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenant", reflect.TypeOf((*MockStore)(nil).GetTenant), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListTenants mocks base method.
func (m *MockStore) ListTenants(arg0 context.Context, arg1 db.ListTenantsParams) ([]db.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTenants", arg0, arg1)
	ret0, _ := ret[0].([]db.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTenants indicates an expected call of ListTenants.
func (mr *MockStoreMockRecorder) ListTenants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTenants", reflect.TypeOf((*MockStore)(nil).ListTenants), arg0, arg1)
}

// RegisterTx mocks base method.
func (m *MockStore) RegisterTx(arg0 context.Context, arg1 db.RegisterTxParams) (*db.RegisterTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTx", arg0, arg1)
	ret0, _ := ret[0].(*db.RegisterTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTx indicates an expected call of RegisterTx.
func (mr *MockStoreMockRecorder) RegisterTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTx", reflect.TypeOf((*MockStore)(nil).RegisterTx), arg0, arg1)
}
