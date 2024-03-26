// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iotexproject/iotex-address/address (interfaces: Address)

// Package mock_iotex_address is a generated GoMock package.
package mock_iotex_address

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAddress is a mock of Address interface.
type MockAddress struct {
	ctrl     *gomock.Controller
	recorder *MockAddressMockRecorder
}

// MockAddressMockRecorder is the mock recorder for MockAddress.
type MockAddressMockRecorder struct {
	mock *MockAddress
}

// NewMockAddress creates a new mock instance.
func NewMockAddress(ctrl *gomock.Controller) *MockAddress {
	mock := &MockAddress{ctrl: ctrl}
	mock.recorder = &MockAddressMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddress) EXPECT() *MockAddressMockRecorder {
	return m.recorder
}

// Bytes mocks base method.
func (m *MockAddress) Bytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes.
func (mr *MockAddressMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockAddress)(nil).Bytes))
}

// Hex mocks base method.
func (m *MockAddress) Hex() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hex")
	ret0, _ := ret[0].(string)
	return ret0
}

// Hex indicates an expected call of Hex.
func (mr *MockAddressMockRecorder) Hex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hex", reflect.TypeOf((*MockAddress)(nil).Hex))
}

// String mocks base method.
func (m *MockAddress) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockAddressMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockAddress)(nil).String))
}