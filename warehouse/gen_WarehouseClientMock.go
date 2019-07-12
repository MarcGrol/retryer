// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package warehouse is a generated GoMock package.
package warehouse

import (
	context "context"
	httpclient "github.com/MarcGrol/forwardhttp/httpclient"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWarehouser is a mock of Warehouser interface
type MockWarehouser struct {
	ctrl     *gomock.Controller
	recorder *MockWarehouserMockRecorder
}

// MockWarehouserMockRecorder is the mock recorder for MockWarehouser
type MockWarehouserMockRecorder struct {
	mock *MockWarehouser
}

// NewMockWarehouser creates a new mock instance
func NewMockWarehouser(ctrl *gomock.Controller) *MockWarehouser {
	mock := &MockWarehouser{ctrl: ctrl}
	mock.recorder = &MockWarehouserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWarehouser) EXPECT() *MockWarehouserMockRecorder {
	return m.recorder
}

// Put mocks base method
func (m *MockWarehouser) Put(c context.Context, req httpclient.Request, resp *httpclient.Response, err error, stats Stats) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", c, req, resp, err, stats)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockWarehouserMockRecorder) Put(c, req, resp, err, stats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockWarehouser)(nil).Put), c, req, resp, err, stats)
}
