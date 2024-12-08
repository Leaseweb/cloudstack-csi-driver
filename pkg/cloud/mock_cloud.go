// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloud.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_cloud.go -package=cloud -source ./cloud.go
//

// Package cloud is a generated GoMock package.
package cloud

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCloud is a mock of Cloud interface.
type MockCloud struct {
	ctrl     *gomock.Controller
	recorder *MockCloudMockRecorder
	isgomock struct{}
}

// MockCloudMockRecorder is the mock recorder for MockCloud.
type MockCloudMockRecorder struct {
	mock *MockCloud
}

// NewMockCloud creates a new mock instance.
func NewMockCloud(ctrl *gomock.Controller) *MockCloud {
	mock := &MockCloud{ctrl: ctrl}
	mock.recorder = &MockCloudMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloud) EXPECT() *MockCloudMockRecorder {
	return m.recorder
}

// AttachVolume mocks base method.
func (m *MockCloud) AttachVolume(ctx context.Context, volumeID, vmID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachVolume", ctx, volumeID, vmID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachVolume indicates an expected call of AttachVolume.
func (mr *MockCloudMockRecorder) AttachVolume(ctx, volumeID, vmID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolume", reflect.TypeOf((*MockCloud)(nil).AttachVolume), ctx, volumeID, vmID)
}

// CreateVolume mocks base method.
func (m *MockCloud) CreateVolume(ctx context.Context, diskOfferingID, zoneID, name string, sizeInGB int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", ctx, diskOfferingID, zoneID, name, sizeInGB)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockCloudMockRecorder) CreateVolume(ctx, diskOfferingID, zoneID, name, sizeInGB any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockCloud)(nil).CreateVolume), ctx, diskOfferingID, zoneID, name, sizeInGB)
}

// DeleteVolume mocks base method.
func (m *MockCloud) DeleteVolume(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockCloudMockRecorder) DeleteVolume(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockCloud)(nil).DeleteVolume), ctx, id)
}

// DetachVolume mocks base method.
func (m *MockCloud) DetachVolume(ctx context.Context, volumeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachVolume", ctx, volumeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachVolume indicates an expected call of DetachVolume.
func (mr *MockCloudMockRecorder) DetachVolume(ctx, volumeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolume", reflect.TypeOf((*MockCloud)(nil).DetachVolume), ctx, volumeID)
}

// ExpandVolume mocks base method.
func (m *MockCloud) ExpandVolume(ctx context.Context, volumeID string, newSizeInGB int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpandVolume", ctx, volumeID, newSizeInGB)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExpandVolume indicates an expected call of ExpandVolume.
func (mr *MockCloudMockRecorder) ExpandVolume(ctx, volumeID, newSizeInGB any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandVolume", reflect.TypeOf((*MockCloud)(nil).ExpandVolume), ctx, volumeID, newSizeInGB)
}

// GetNodeInfo mocks base method.
func (m *MockCloud) GetNodeInfo(ctx context.Context, vmName string) (*VM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeInfo", ctx, vmName)
	ret0, _ := ret[0].(*VM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeInfo indicates an expected call of GetNodeInfo.
func (mr *MockCloudMockRecorder) GetNodeInfo(ctx, vmName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeInfo", reflect.TypeOf((*MockCloud)(nil).GetNodeInfo), ctx, vmName)
}

// GetVMByID mocks base method.
func (m *MockCloud) GetVMByID(ctx context.Context, vmID string) (*VM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVMByID", ctx, vmID)
	ret0, _ := ret[0].(*VM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVMByID indicates an expected call of GetVMByID.
func (mr *MockCloudMockRecorder) GetVMByID(ctx, vmID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVMByID", reflect.TypeOf((*MockCloud)(nil).GetVMByID), ctx, vmID)
}

// GetVolumeByID mocks base method.
func (m *MockCloud) GetVolumeByID(ctx context.Context, volumeID string) (*Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeByID", ctx, volumeID)
	ret0, _ := ret[0].(*Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeByID indicates an expected call of GetVolumeByID.
func (mr *MockCloudMockRecorder) GetVolumeByID(ctx, volumeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeByID", reflect.TypeOf((*MockCloud)(nil).GetVolumeByID), ctx, volumeID)
}

// GetVolumeByName mocks base method.
func (m *MockCloud) GetVolumeByName(ctx context.Context, name string) (*Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeByName", ctx, name)
	ret0, _ := ret[0].(*Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeByName indicates an expected call of GetVolumeByName.
func (mr *MockCloudMockRecorder) GetVolumeByName(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeByName", reflect.TypeOf((*MockCloud)(nil).GetVolumeByName), ctx, name)
}

// ListZonesID mocks base method.
func (m *MockCloud) ListZonesID(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListZonesID", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListZonesID indicates an expected call of ListZonesID.
func (mr *MockCloudMockRecorder) ListZonesID(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZonesID", reflect.TypeOf((*MockCloud)(nil).ListZonesID), ctx)
}