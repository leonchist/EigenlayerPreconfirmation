package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/model"
	object "gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
)

// MockTaskRepositoryInterface is a mock of TaskRepositoryInterface interface
type MockTaskRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryInterfaceMockRecorder
}

// MockTaskRepositoryInterfaceMockRecorder is the mock recorder for MockTaskRepositoryInterface
type MockTaskRepositoryInterfaceMockRecorder struct {
	mock *MockTaskRepositoryInterface
}

// NewMockTaskRepositoryInterface creates a new mock instance
func NewMockTaskRepositoryInterface(ctrl *gomock.Controller) *MockTaskRepositoryInterface {
	mock := &MockTaskRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskRepositoryInterface) EXPECT() *MockTaskRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockTaskRepositoryInterface) GetAll(ctx context.Context) ([]model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockTaskRepositoryInterfaceMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTaskRepositoryInterface)(nil).GetAll), ctx)
}

// Create mocks base method
func (m *MockTaskRepositoryInterface) Create(ctx context.Context, req object.TaskObjRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockTaskRepositoryInterfaceMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskRepositoryInterface)(nil).Create), ctx, req)
}
