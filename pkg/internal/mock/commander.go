// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/commander/commander.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Commander is a mock of Commander interface
type Commander struct {
	ctrl     *gomock.Controller
	recorder *CommanderMockRecorder
}

// CommanderMockRecorder is the mock recorder for Commander
type CommanderMockRecorder struct {
	mock *Commander
}

// NewCommander creates a new mock instance
func NewCommander(ctrl *gomock.Controller) *Commander {
	mock := &Commander{ctrl: ctrl}
	mock.recorder = &CommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Commander) EXPECT() *CommanderMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *Commander) Execute(cmd string, args []string) ([]string, error) {
	ret := m.ctrl.Call(m, "Execute", cmd, args)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *CommanderMockRecorder) Execute(cmd, args interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*Commander)(nil).Execute), cmd, args)
}
