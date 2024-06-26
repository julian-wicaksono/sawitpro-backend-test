// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

// GetUserByPhoneNumber mocks base method.
func (m *MockRepositoryInterface) GetUserByPhoneNumber(ctx context.Context, input UserData) (UserId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhoneNumber", ctx, input)
	ret0, _ := ret[0].(UserId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhoneNumber indicates an expected call of GetUserByPhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByPhoneNumber(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByPhoneNumber), ctx, input)
}

// GetUserDataById mocks base method.
func (m *MockRepositoryInterface) GetUserDataById(ctx context.Context, input UserId) (UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDataById", ctx, input)
	ret0, _ := ret[0].(UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDataById indicates an expected call of GetUserDataById.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserDataById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDataById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserDataById), ctx, input)
}

// GetUserDataByPhoneNumber mocks base method.
func (m *MockRepositoryInterface) GetUserDataByPhoneNumber(ctx context.Context, input UserData) (UserData, UserId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDataByPhoneNumber", ctx, input)
	ret0, _ := ret[0].(UserData)
	ret1, _ := ret[1].(UserId)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserDataByPhoneNumber indicates an expected call of GetUserDataByPhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserDataByPhoneNumber(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDataByPhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserDataByPhoneNumber), ctx, input)
}

// InsertUser mocks base method.
func (m *MockRepositoryInterface) InsertUser(ctx context.Context, input UserData) (UserId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", ctx, input)
	ret0, _ := ret[0].(UserId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockRepositoryInterfaceMockRecorder) InsertUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockRepositoryInterface)(nil).InsertUser), ctx, input)
}

// UpdateFullName mocks base method.
func (m *MockRepositoryInterface) UpdateFullName(ctx context.Context, userData UserData, userId UserId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFullName", ctx, userData, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFullName indicates an expected call of UpdateFullName.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateFullName(ctx, userData, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFullName", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateFullName), ctx, userData, userId)
}

// UpdateLoginCount mocks base method.
func (m *MockRepositoryInterface) UpdateLoginCount(ctx context.Context, input UserId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLoginCount", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLoginCount indicates an expected call of UpdateLoginCount.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateLoginCount(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLoginCount", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateLoginCount), ctx, input)
}

// UpdatePhoneNumber mocks base method.
func (m *MockRepositoryInterface) UpdatePhoneNumber(ctx context.Context, userData UserData, userId UserId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoneNumber", ctx, userData, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhoneNumber indicates an expected call of UpdatePhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) UpdatePhoneNumber(ctx, userData, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdatePhoneNumber), ctx, userData, userId)
}

// UpdateUserData mocks base method.
func (m *MockRepositoryInterface) UpdateUserData(ctx context.Context, userData UserData, userId UserId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserData", ctx, userData, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserData indicates an expected call of UpdateUserData.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateUserData(ctx, userData, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserData", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateUserData), ctx, userData, userId)
}
