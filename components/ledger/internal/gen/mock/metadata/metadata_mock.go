// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/ledger/internal/domain/metadata (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen --destination=../../gen/mock/metadata/metadata_mock.go --package=mock . Repository
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	http "github.com/LerianStudio/midaz/common/net/http"
	metadata "github.com/LerianStudio/midaz/components/ledger/internal/domain/metadata"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 context.Context, arg1 string, arg2 *metadata.Metadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockRepository) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), arg0, arg1, arg2)
}

// FindByEntity mocks base method.
func (m *MockRepository) FindByEntity(arg0 context.Context, arg1, arg2 string) (*metadata.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEntity", arg0, arg1, arg2)
	ret0, _ := ret[0].(*metadata.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEntity indicates an expected call of FindByEntity.
func (mr *MockRepositoryMockRecorder) FindByEntity(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEntity", reflect.TypeOf((*MockRepository)(nil).FindByEntity), arg0, arg1, arg2)
}

// FindList mocks base method.
func (m *MockRepository) FindList(arg0 context.Context, arg1 string, arg2 http.QueryHeader) ([]*metadata.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*metadata.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.z
func (mr *MockRepositoryMockRecorder) FindList(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockRepository)(nil).FindList), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockRepository) Update(arg0 context.Context, arg1, arg2 string, arg3 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), arg0, arg1, arg2, arg3)
}
