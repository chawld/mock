// Code generated by MockGen. DO NOT EDIT.
// Source: intf.go

// Package mock_stacked is a generated GoMock package.
package mock_stacked

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	stacked "github.com/golang/mock/sample/stacked"
)

// MockStackedIntf is a mock of StackedIntf interface.
type MockStackedIntf struct {
	frame    gomock.StackFrame
	ctrl     *gomock.Controller
	recorder *MockStackedIntfMockRecorder
}

// MockStackedIntfMockRecorder is the mock recorder for MockStackedIntf.
type MockStackedIntfMockRecorder struct {
	mock *MockStackedIntf
}

// MockStackedIntfOption is a function that customizes the mock.
type MockStackedIntfOption func(m *MockStackedIntf)

// NewMockStackedIntf creates a new mock instance.
func NewMockStackedIntf(ctrl *gomock.Controller, opt ...MockStackedIntfOption) *MockStackedIntf {
	mock := &MockStackedIntf{ctrl: ctrl}
	mock.recorder = &MockStackedIntfMockRecorder{mock}
	for _, o := range opt {
		o(mock)
	}
	return mock
}

// WithBackingStackedIntf returns an Option that adds the given backing implementation to the mock.
func WithBackingStackedIntf(next stacked.StackedIntf) MockStackedIntfOption {
	return func(m *MockStackedIntf) {
		m.frame = gomock.StackFrame{Next: next, Level: 0}
		for level := 1; ; level++ {
			b, ok := m.frame.Next.(*MockStackedIntf)
			if !ok {
				break
			}
			b.frame.Level = level
			m = b
		}
	}
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStackedIntf) EXPECT() *MockStackedIntfMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockStackedIntf) Get() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.CallStacked(m, &m.frame, "Get")
	ret0, _ := ret[0].(int)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockStackedIntfMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStackedIntf)(nil).Get))
}

// Inc mocks base method.
func (m *MockStackedIntf) Inc() {
	m.ctrl.T.Helper()
	m.ctrl.CallStacked(m, &m.frame, "Inc")
}

// Inc indicates an expected call of Inc.
func (mr *MockStackedIntfMockRecorder) Inc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inc", reflect.TypeOf((*MockStackedIntf)(nil).Inc))
}
