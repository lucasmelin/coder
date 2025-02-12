// Code generated by MockGen. DO NOT EDIT.
// Source: .. (interfaces: Execer)
//
// Generated by this command:
//
//	mockgen -destination ./execmock.go -package execmock .. Execer
//

// Package execmock is a generated GoMock package.
package execmock

import (
	context "context"
	exec "os/exec"
	reflect "reflect"

	pty "github.com/coder/coder/v2/pty"
	gomock "go.uber.org/mock/gomock"
)

// MockExecer is a mock of Execer interface.
type MockExecer struct {
	ctrl     *gomock.Controller
	recorder *MockExecerMockRecorder
	isgomock struct{}
}

// MockExecerMockRecorder is the mock recorder for MockExecer.
type MockExecerMockRecorder struct {
	mock *MockExecer
}

// NewMockExecer creates a new mock instance.
func NewMockExecer(ctrl *gomock.Controller) *MockExecer {
	mock := &MockExecer{ctrl: ctrl}
	mock.recorder = &MockExecerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecer) EXPECT() *MockExecerMockRecorder {
	return m.recorder
}

// CommandContext mocks base method.
func (m *MockExecer) CommandContext(ctx context.Context, cmd string, args ...string) *exec.Cmd {
	m.ctrl.T.Helper()
	varargs := []any{ctx, cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommandContext", varargs...)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// CommandContext indicates an expected call of CommandContext.
func (mr *MockExecerMockRecorder) CommandContext(ctx, cmd any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandContext", reflect.TypeOf((*MockExecer)(nil).CommandContext), varargs...)
}

// PTYCommandContext mocks base method.
func (m *MockExecer) PTYCommandContext(ctx context.Context, cmd string, args ...string) *pty.Cmd {
	m.ctrl.T.Helper()
	varargs := []any{ctx, cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PTYCommandContext", varargs...)
	ret0, _ := ret[0].(*pty.Cmd)
	return ret0
}

// PTYCommandContext indicates an expected call of PTYCommandContext.
func (mr *MockExecerMockRecorder) PTYCommandContext(ctx, cmd any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PTYCommandContext", reflect.TypeOf((*MockExecer)(nil).PTYCommandContext), varargs...)
}
