// Code generated by MockGen. DO NOT EDIT.
// Source: D:\go\src\github.com\dedensmkn4\movie-backend\internal\logging\repository\init.go

// Package mock_repository is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/dedensmkn4/movie-backend/internal/logging/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockLoggingRepository is a mock of LoggingRepository interface.
type MockLoggingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLoggingRepositoryMockRecorder
}

// MockLoggingRepositoryMockRecorder is the mock recorder for MockLoggingRepository.
type MockLoggingRepositoryMockRecorder struct {
	mock *MockLoggingRepository
}

// NewMockLoggingRepository creates a new mock instance.
func NewMockLoggingRepository(ctrl *gomock.Controller) *MockLoggingRepository {
	mock := &MockLoggingRepository{ctrl: ctrl}
	mock.recorder = &MockLoggingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggingRepository) EXPECT() *MockLoggingRepositoryMockRecorder {
	return m.recorder
}

// InsertLoggingSearchOmdb mocks base method.
func (m *MockLoggingRepository) InsertLoggingSearchOmdb(ent *domain.LoggingOmdbSearch) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertLoggingSearchOmdb", ent)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertLoggingSearchOmdb indicates an expected call of InsertLoggingSearchOmdb.
func (mr *MockLoggingRepositoryMockRecorder) InsertLoggingSearchOmdb(ent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertLoggingSearchOmdb", reflect.TypeOf((*MockLoggingRepository)(nil).InsertLoggingSearchOmdb), ent)
}