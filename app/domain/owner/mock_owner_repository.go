// Code generated by MockGen. DO NOT EDIT.
// Source: owner_repository.go
//
// Generated by this command:
//
//	mockgen -package owner -source owner_repository.go -destination mock_owner_repository.go
//
// Package owner is a generated GoMock package.
package owner

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockOwnerRepository is a mock of OwnerRepository interface.
type MockOwnerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOwnerRepositoryMockRecorder
}

// MockOwnerRepositoryMockRecorder is the mock recorder for MockOwnerRepository.
type MockOwnerRepositoryMockRecorder struct {
	mock *MockOwnerRepository
}

// NewMockOwnerRepository creates a new mock instance.
func NewMockOwnerRepository(ctrl *gomock.Controller) *MockOwnerRepository {
	mock := &MockOwnerRepository{ctrl: ctrl}
	mock.recorder = &MockOwnerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOwnerRepository) EXPECT() *MockOwnerRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockOwnerRepository) FindByID(ctx context.Context, id string) (*Owner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*Owner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockOwnerRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockOwnerRepository)(nil).FindByID), ctx, id)
}

// Save mocks base method.
func (m *MockOwnerRepository) Save(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOwnerRepositoryMockRecorder) Save(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOwnerRepository)(nil).Save), ctx)
}