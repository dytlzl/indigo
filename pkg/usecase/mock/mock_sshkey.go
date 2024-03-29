// Code generated by MockGen. DO NOT EDIT.
// Source: sshkey.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	domain "github.com/dytlzl/indigo/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockSSHKeyRepository is a mock of SSHKeyRepository interface.
type MockSSHKeyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSSHKeyRepositoryMockRecorder
}

// MockSSHKeyRepositoryMockRecorder is the mock recorder for MockSSHKeyRepository.
type MockSSHKeyRepositoryMockRecorder struct {
	mock *MockSSHKeyRepository
}

// NewMockSSHKeyRepository creates a new mock instance.
func NewMockSSHKeyRepository(ctrl *gomock.Controller) *MockSSHKeyRepository {
	mock := &MockSSHKeyRepository{ctrl: ctrl}
	mock.recorder = &MockSSHKeyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSHKeyRepository) EXPECT() *MockSSHKeyRepositoryMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockSSHKeyRepository) List(ctx context.Context) ([]domain.SSHKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]domain.SSHKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSSHKeyRepositoryMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSSHKeyRepository)(nil).List), ctx)
}

// MockSSHKeyUseCase is a mock of SSHKeyUseCase interface.
type MockSSHKeyUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSSHKeyUseCaseMockRecorder
}

// MockSSHKeyUseCaseMockRecorder is the mock recorder for MockSSHKeyUseCase.
type MockSSHKeyUseCaseMockRecorder struct {
	mock *MockSSHKeyUseCase
}

// NewMockSSHKeyUseCase creates a new mock instance.
func NewMockSSHKeyUseCase(ctrl *gomock.Controller) *MockSSHKeyUseCase {
	mock := &MockSSHKeyUseCase{ctrl: ctrl}
	mock.recorder = &MockSSHKeyUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSHKeyUseCase) EXPECT() *MockSSHKeyUseCaseMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockSSHKeyUseCase) List(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockSSHKeyUseCaseMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSSHKeyUseCase)(nil).List), ctx)
}
