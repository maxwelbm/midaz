// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/transaction/internal/adapters/redis (interfaces: RedisRepository)
//
// Generated by this command:
//
//	mockgen --destination=redis.mock.go --package=redis . RedisRepository
//

// Package redis is a generated GoMock package.
package redis

import (
	context "context"
	gomock "go.uber.org/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockRedisRepository is a mock of RedisRepository interface.
type MockRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisRepositoryMockRecorder
}

// MockRedisRepositoryMockRecorder is the mock recorder for MockRedisRepository.
type MockRedisRepositoryMockRecorder struct {
	mock *MockRedisRepository
}

// NewMockRedisRepository creates a new mock instance.
func NewMockRedisRepository(ctrl *gomock.Controller) *MockRedisRepository {
	mock := &MockRedisRepository{ctrl: ctrl}
	mock.recorder = &MockRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisRepository) EXPECT() *MockRedisRepositoryMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockRedisRepository) Del(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockRedisRepositoryMockRecorder) Del(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisRepository)(nil).Del), arg0, arg1)
}

// Get mocks base method.
func (m *MockRedisRepository) Get(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisRepositoryMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisRepository)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockRedisRepository) Set(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisRepositoryMockRecorder) Set(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisRepository)(nil).Set), arg0, arg1, arg2, arg3)
}
