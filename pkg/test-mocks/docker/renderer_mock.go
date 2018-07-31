// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicatedhq/ship/pkg/lifecycle/render/docker (interfaces: Renderer)

// Package docker is a generated GoMock package.
package docker

import (
	context "context"
	reflect "reflect"

	log "github.com/go-kit/kit/log"
	gomock "github.com/golang/mock/gomock"
	libyaml "github.com/replicatedhq/libyaml"
	api "github.com/replicatedhq/ship/pkg/api"
)

// MockRenderer is a mock of Renderer interface
type MockRenderer struct {
	ctrl     *gomock.Controller
	recorder *MockRendererMockRecorder
}

// MockRendererMockRecorder is the mock recorder for MockRenderer
type MockRendererMockRecorder struct {
	mock *MockRenderer
}

// NewMockRenderer creates a new mock instance
func NewMockRenderer(ctrl *gomock.Controller) *MockRenderer {
	mock := &MockRenderer{ctrl: ctrl}
	mock.recorder = &MockRendererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRenderer) EXPECT() *MockRendererMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockRenderer) Execute(arg0 api.DockerAsset, arg1 api.ReleaseMetadata, arg2 func(chan interface{}, log.Logger) error, arg3 string, arg4 map[string]interface{}, arg5 []libyaml.ConfigGroup) func(context.Context) error {
	ret := m.ctrl.Call(m, "Execute", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(func(context.Context) error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockRendererMockRecorder) Execute(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRenderer)(nil).Execute), arg0, arg1, arg2, arg3, arg4, arg5)
}
