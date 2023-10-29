// Code generated by MockGen. DO NOT EDIT.
// Source: Food_Shop_Server/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	db "Food_Shop_Server/db/sqlc"
	context "context"
	reflect "reflect"

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

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockStore) DeleteProduct(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), arg0, arg1)
}

// GetProductCount mocks base method.
func (m *MockStore) GetProductCount(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductCount indicates an expected call of GetProductCount.
func (mr *MockStoreMockRecorder) GetProductCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductCount", reflect.TypeOf((*MockStore)(nil).GetProductCount), arg0)
}

// GetProductCountByOwner mocks base method.
func (m *MockStore) GetProductCountByOwner(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductCountByOwner", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductCountByOwner indicates an expected call of GetProductCountByOwner.
func (mr *MockStoreMockRecorder) GetProductCountByOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductCountByOwner", reflect.TypeOf((*MockStore)(nil).GetProductCountByOwner), arg0, arg1)
}

// GetProductList mocks base method.
func (m *MockStore) GetProductList(arg0 context.Context, arg1 db.GetProductListParams) ([]db.GetProductListRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductList", arg0, arg1)
	ret0, _ := ret[0].([]db.GetProductListRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductList indicates an expected call of GetProductList.
func (mr *MockStoreMockRecorder) GetProductList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductList", reflect.TypeOf((*MockStore)(nil).GetProductList), arg0, arg1)
}

// GetProductListByOwner mocks base method.
func (m *MockStore) GetProductListByOwner(arg0 context.Context, arg1 db.GetProductListByOwnerParams) ([]db.GetProductListByOwnerRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductListByOwner", arg0, arg1)
	ret0, _ := ret[0].([]db.GetProductListByOwnerRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductListByOwner indicates an expected call of GetProductListByOwner.
func (mr *MockStoreMockRecorder) GetProductListByOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductListByOwner", reflect.TypeOf((*MockStore)(nil).GetProductListByOwner), arg0, arg1)
}
