// Code generated by MockGen. DO NOT EDIT.
// Source: pinger.go
//
// Generated by this command:
//
//	mockgen -source pinger.go -destination mocks/pinger.go -package healthcheckmocks
//

// Package healthcheckmocks is a generated GoMock package.
package healthcheckmocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPinger is a mock of Pinger interface.
type MockPinger struct {
	ctrl     *gomock.Controller
	recorder *MockPingerMockRecorder
}

// MockPingerMockRecorder is the mock recorder for MockPinger.
type MockPingerMockRecorder struct {
	mock *MockPinger
}

// NewMockPinger creates a new mock instance.
func NewMockPinger(ctrl *gomock.Controller) *MockPinger {
	mock := &MockPinger{ctrl: ctrl}
	mock.recorder = &MockPingerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPinger) EXPECT() *MockPingerMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockPinger) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPingerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPinger)(nil).Name))
}

// Ping mocks base method.
func (m *MockPinger) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockPingerMockRecorder) Ping(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockPinger)(nil).Ping), ctx)
}
