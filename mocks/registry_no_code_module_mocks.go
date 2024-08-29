// Code generated by MockGen. DO NOT EDIT.
// Source: registry_no_code_module.go
//
// Generated by this command:
//
//	mockgen -source=registry_no_code_module.go -destination=mocks/registry_no_code_module_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/hashicorp/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockRegistryNoCodeModules is a mock of RegistryNoCodeModules interface.
type MockRegistryNoCodeModules struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryNoCodeModulesMockRecorder
}

// MockRegistryNoCodeModulesMockRecorder is the mock recorder for MockRegistryNoCodeModules.
type MockRegistryNoCodeModulesMockRecorder struct {
	mock *MockRegistryNoCodeModules
}

// NewMockRegistryNoCodeModules creates a new mock instance.
func NewMockRegistryNoCodeModules(ctrl *gomock.Controller) *MockRegistryNoCodeModules {
	mock := &MockRegistryNoCodeModules{ctrl: ctrl}
	mock.recorder = &MockRegistryNoCodeModulesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistryNoCodeModules) EXPECT() *MockRegistryNoCodeModulesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRegistryNoCodeModules) Create(ctx context.Context, organization string, options tfe.RegistryNoCodeModuleCreateOptions) (*tfe.RegistryNoCodeModule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.RegistryNoCodeModule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRegistryNoCodeModulesMockRecorder) Create(ctx, organization, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRegistryNoCodeModules)(nil).Create), ctx, organization, options)
}

// CreateWorkspace mocks base method.
func (m *MockRegistryNoCodeModules) CreateWorkspace(ctx context.Context, noCodeModuleID string, options *tfe.RegistryNoCodeModuleCreateWorkspaceOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkspace", ctx, noCodeModuleID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkspace indicates an expected call of CreateWorkspace.
func (mr *MockRegistryNoCodeModulesMockRecorder) CreateWorkspace(ctx, noCodeModuleID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkspace", reflect.TypeOf((*MockRegistryNoCodeModules)(nil).CreateWorkspace), ctx, noCodeModuleID, options)
}

// Delete mocks base method.
func (m *MockRegistryNoCodeModules) Delete(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRegistryNoCodeModulesMockRecorder) Delete(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRegistryNoCodeModules)(nil).Delete), ctx, ID)
}

// Read mocks base method.
func (m *MockRegistryNoCodeModules) Read(ctx context.Context, noCodeModuleID string, options *tfe.RegistryNoCodeModuleReadOptions) (*tfe.RegistryNoCodeModule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, noCodeModuleID, options)
	ret0, _ := ret[0].(*tfe.RegistryNoCodeModule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRegistryNoCodeModulesMockRecorder) Read(ctx, noCodeModuleID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRegistryNoCodeModules)(nil).Read), ctx, noCodeModuleID, options)
}

// Update mocks base method.
func (m *MockRegistryNoCodeModules) Update(ctx context.Context, noCodeModuleID string, options tfe.RegistryNoCodeModuleUpdateOptions) (*tfe.RegistryNoCodeModule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, noCodeModuleID, options)
	ret0, _ := ret[0].(*tfe.RegistryNoCodeModule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRegistryNoCodeModulesMockRecorder) Update(ctx, noCodeModuleID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRegistryNoCodeModules)(nil).Update), ctx, noCodeModuleID, options)
}

// UpgradeWorkspace mocks base method.
func (m *MockRegistryNoCodeModules) UpgradeWorkspace(ctx context.Context, noCodeModuleID, workspaceID string, options *tfe.RegistryNoCodeModuleUpgradeWorkspaceOptions) (*tfe.WorkspaceUpgrade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeWorkspace", ctx, noCodeModuleID, workspaceID, options)
	ret0, _ := ret[0].(*tfe.WorkspaceUpgrade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeWorkspace indicates an expected call of UpgradeWorkspace.
func (mr *MockRegistryNoCodeModulesMockRecorder) UpgradeWorkspace(ctx, noCodeModuleID, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeWorkspace", reflect.TypeOf((*MockRegistryNoCodeModules)(nil).UpgradeWorkspace), ctx, noCodeModuleID, workspaceID, options)
}
