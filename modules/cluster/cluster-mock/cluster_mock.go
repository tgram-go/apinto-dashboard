// Code generated by MockGen. DO NOT EDIT.
// Source: cluster/service.go

// Package mock_cluster is a generated GoMock package.
package mock_cluster

import (
	context "context"
	reflect "reflect"

	v1 "github.com/eolinker/apinto-dashboard/client/v1"
	cluster_dto "github.com/eolinker/apinto-dashboard/modules/cluster/cluster-dto"
	cluster_entry "github.com/eolinker/apinto-dashboard/modules/cluster/cluster-entry"
	cluster_model "github.com/eolinker/apinto-dashboard/modules/cluster/cluster-model"
	gomock "github.com/golang/mock/gomock"
)

// MockIApintoClient is a mock of IApintoClient interface.
type MockIApintoClient struct {
	ctrl     *gomock.Controller
	recorder *MockIApintoClientMockRecorder
}

// MockIApintoClientMockRecorder is the mock recorder for MockIApintoClient.
type MockIApintoClientMockRecorder struct {
	mock *MockIApintoClient
}

// NewMockIApintoClient creates a new mock instance.
func NewMockIApintoClient(ctrl *gomock.Controller) *MockIApintoClient {
	mock := &MockIApintoClient{ctrl: ctrl}
	mock.recorder = &MockIApintoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIApintoClient) EXPECT() *MockIApintoClientMockRecorder {
	return m.recorder
}

// GetClient mocks base method.
func (m *MockIApintoClient) GetClient(ctx context.Context, clusterId int) (v1.IClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient", ctx, clusterId)
	ret0, _ := ret[0].(v1.IClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient.
func (mr *MockIApintoClientMockRecorder) GetClient(ctx, clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockIApintoClient)(nil).GetClient), ctx, clusterId)
}

// SetClient mocks base method.
func (m *MockIApintoClient) SetClient(namespace, clusterId int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", namespace, clusterId)
}

// SetClient indicates an expected call of SetClient.
func (mr *MockIApintoClientMockRecorder) SetClient(namespace, clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockIApintoClient)(nil).SetClient), namespace, clusterId)
}

// MockIClusterCertificateService is a mock of IClusterCertificateService interface.
type MockIClusterCertificateService struct {
	ctrl     *gomock.Controller
	recorder *MockIClusterCertificateServiceMockRecorder
}

// MockIClusterCertificateServiceMockRecorder is the mock recorder for MockIClusterCertificateService.
type MockIClusterCertificateServiceMockRecorder struct {
	mock *MockIClusterCertificateService
}

// NewMockIClusterCertificateService creates a new mock instance.
func NewMockIClusterCertificateService(ctrl *gomock.Controller) *MockIClusterCertificateService {
	mock := &MockIClusterCertificateService{ctrl: ctrl}
	mock.recorder = &MockIClusterCertificateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClusterCertificateService) EXPECT() *MockIClusterCertificateServiceMockRecorder {
	return m.recorder
}

// DeleteById mocks base method.
func (m *MockIClusterCertificateService) DeleteById(ctx context.Context, namespaceId int, clusterName string, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, namespaceId, clusterName, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockIClusterCertificateServiceMockRecorder) DeleteById(ctx, namespaceId, clusterName, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockIClusterCertificateService)(nil).DeleteById), ctx, namespaceId, clusterName, id)
}

// Insert mocks base method.
func (m *MockIClusterCertificateService) Insert(ctx context.Context, operator, namespaceId int, clusterName, key, pem string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, operator, namespaceId, clusterName, key, pem)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockIClusterCertificateServiceMockRecorder) Insert(ctx, operator, namespaceId, clusterName, key, pem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIClusterCertificateService)(nil).Insert), ctx, operator, namespaceId, clusterName, key, pem)
}

// QueryList mocks base method.
func (m *MockIClusterCertificateService) QueryList(ctx context.Context, namespaceId int, clusterName string) ([]*cluster_model.ClusterCertificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryList", ctx, namespaceId, clusterName)
	ret0, _ := ret[0].([]*cluster_model.ClusterCertificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryList indicates an expected call of QueryList.
func (mr *MockIClusterCertificateServiceMockRecorder) QueryList(ctx, namespaceId, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryList", reflect.TypeOf((*MockIClusterCertificateService)(nil).QueryList), ctx, namespaceId, clusterName)
}

// Update mocks base method.
func (m *MockIClusterCertificateService) Update(ctx context.Context, operator, namespaceId, certificateId int, clusterName, key, pem string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, operator, namespaceId, certificateId, clusterName, key, pem)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIClusterCertificateServiceMockRecorder) Update(ctx, operator, namespaceId, certificateId, clusterName, key, pem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIClusterCertificateService)(nil).Update), ctx, operator, namespaceId, certificateId, clusterName, key, pem)
}

// MockIClusterService is a mock of IClusterService interface.
type MockIClusterService struct {
	ctrl     *gomock.Controller
	recorder *MockIClusterServiceMockRecorder
}

// MockIClusterServiceMockRecorder is the mock recorder for MockIClusterService.
type MockIClusterServiceMockRecorder struct {
	mock *MockIClusterService
}

// NewMockIClusterService creates a new mock instance.
func NewMockIClusterService(ctrl *gomock.Controller) *MockIClusterService {
	mock := &MockIClusterService{ctrl: ctrl}
	mock.recorder = &MockIClusterServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClusterService) EXPECT() *MockIClusterServiceMockRecorder {
	return m.recorder
}

// CheckByNamespaceByName mocks base method.
func (m *MockIClusterService) CheckByNamespaceByName(ctx context.Context, namespaceId int, name string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckByNamespaceByName", ctx, namespaceId, name)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckByNamespaceByName indicates an expected call of CheckByNamespaceByName.
func (mr *MockIClusterServiceMockRecorder) CheckByNamespaceByName(ctx, namespaceId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckByNamespaceByName", reflect.TypeOf((*MockIClusterService)(nil).CheckByNamespaceByName), ctx, namespaceId, name)
}

// DeleteByNamespaceIdByName mocks base method.
func (m *MockIClusterService) DeleteByNamespaceIdByName(ctx context.Context, namespaceId, userId int, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByNamespaceIdByName", ctx, namespaceId, userId, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByNamespaceIdByName indicates an expected call of DeleteByNamespaceIdByName.
func (mr *MockIClusterServiceMockRecorder) DeleteByNamespaceIdByName(ctx, namespaceId, userId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByNamespaceIdByName", reflect.TypeOf((*MockIClusterService)(nil).DeleteByNamespaceIdByName), ctx, namespaceId, userId, name)
}

// GetAllCluster mocks base method.
func (m *MockIClusterService) GetAllCluster(ctx context.Context) ([]*cluster_model.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCluster", ctx)
	ret0, _ := ret[0].([]*cluster_model.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCluster indicates an expected call of GetAllCluster.
func (mr *MockIClusterServiceMockRecorder) GetAllCluster(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCluster", reflect.TypeOf((*MockIClusterService)(nil).GetAllCluster), ctx)
}

// GetByClusterId mocks base method.
func (m *MockIClusterService) GetByClusterId(ctx context.Context, clusterId int) (*cluster_model.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByClusterId", ctx, clusterId)
	ret0, _ := ret[0].(*cluster_model.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByClusterId indicates an expected call of GetByClusterId.
func (mr *MockIClusterServiceMockRecorder) GetByClusterId(ctx, clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByClusterId", reflect.TypeOf((*MockIClusterService)(nil).GetByClusterId), ctx, clusterId)
}

// GetByNames mocks base method.
func (m *MockIClusterService) GetByNames(ctx context.Context, namespaceId int, names []string) ([]*cluster_model.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNames", ctx, namespaceId, names)
	ret0, _ := ret[0].([]*cluster_model.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNames indicates an expected call of GetByNames.
func (mr *MockIClusterServiceMockRecorder) GetByNames(ctx, namespaceId, names interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNames", reflect.TypeOf((*MockIClusterService)(nil).GetByNames), ctx, namespaceId, names)
}

// GetByNamespaceByName mocks base method.
func (m *MockIClusterService) GetByNamespaceByName(ctx context.Context, namespaceId int, name string) (*cluster_model.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNamespaceByName", ctx, namespaceId, name)
	ret0, _ := ret[0].(*cluster_model.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNamespaceByName indicates an expected call of GetByNamespaceByName.
func (mr *MockIClusterServiceMockRecorder) GetByNamespaceByName(ctx, namespaceId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNamespaceByName", reflect.TypeOf((*MockIClusterService)(nil).GetByNamespaceByName), ctx, namespaceId, name)
}

// GetByNamespaceId mocks base method.
func (m *MockIClusterService) GetByNamespaceId(ctx context.Context, namespaceId int) ([]*cluster_model.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNamespaceId", ctx, namespaceId)
	ret0, _ := ret[0].([]*cluster_model.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNamespaceId indicates an expected call of GetByNamespaceId.
func (mr *MockIClusterServiceMockRecorder) GetByNamespaceId(ctx, namespaceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNamespaceId", reflect.TypeOf((*MockIClusterService)(nil).GetByNamespaceId), ctx, namespaceId)
}

// Insert mocks base method.
func (m *MockIClusterService) Insert(ctx context.Context, namespaceId, userId int, clusterInput *cluster_dto.ClusterInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, namespaceId, userId, clusterInput)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockIClusterServiceMockRecorder) Insert(ctx, namespaceId, userId, clusterInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIClusterService)(nil).Insert), ctx, namespaceId, userId, clusterInput)
}

// QueryByNamespaceId mocks base method.
func (m *MockIClusterService) QueryByNamespaceId(ctx context.Context, namespaceId int, clusterName string) (*cluster_model.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryByNamespaceId", ctx, namespaceId, clusterName)
	ret0, _ := ret[0].(*cluster_model.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryByNamespaceId indicates an expected call of QueryByNamespaceId.
func (mr *MockIClusterServiceMockRecorder) QueryByNamespaceId(ctx, namespaceId, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryByNamespaceId", reflect.TypeOf((*MockIClusterService)(nil).QueryByNamespaceId), ctx, namespaceId, clusterName)
}

// QueryListByNamespaceId mocks base method.
func (m *MockIClusterService) QueryListByNamespaceId(ctx context.Context, namespaceId int) ([]*cluster_model.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryListByNamespaceId", ctx, namespaceId)
	ret0, _ := ret[0].([]*cluster_model.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryListByNamespaceId indicates an expected call of QueryListByNamespaceId.
func (mr *MockIClusterServiceMockRecorder) QueryListByNamespaceId(ctx, namespaceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryListByNamespaceId", reflect.TypeOf((*MockIClusterService)(nil).QueryListByNamespaceId), ctx, namespaceId)
}

// UpdateAddr mocks base method.
func (m *MockIClusterService) UpdateAddr(ctx context.Context, userId, clusterId int, addr, uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddr", ctx, userId, clusterId, addr, uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddr indicates an expected call of UpdateAddr.
func (mr *MockIClusterServiceMockRecorder) UpdateAddr(ctx, userId, clusterId, addr, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddr", reflect.TypeOf((*MockIClusterService)(nil).UpdateAddr), ctx, userId, clusterId, addr, uuid)
}

// UpdateDesc mocks base method.
func (m *MockIClusterService) UpdateDesc(ctx context.Context, namespaceId, userId int, name, desc string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDesc", ctx, namespaceId, userId, name, desc)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDesc indicates an expected call of UpdateDesc.
func (mr *MockIClusterServiceMockRecorder) UpdateDesc(ctx, namespaceId, userId, name, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDesc", reflect.TypeOf((*MockIClusterService)(nil).UpdateDesc), ctx, namespaceId, userId, name, desc)
}

// MockIClusterConfigService is a mock of IClusterConfigService interface.
type MockIClusterConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockIClusterConfigServiceMockRecorder
}

// MockIClusterConfigServiceMockRecorder is the mock recorder for MockIClusterConfigService.
type MockIClusterConfigServiceMockRecorder struct {
	mock *MockIClusterConfigService
}

// NewMockIClusterConfigService creates a new mock instance.
func NewMockIClusterConfigService(ctrl *gomock.Controller) *MockIClusterConfigService {
	mock := &MockIClusterConfigService{ctrl: ctrl}
	mock.recorder = &MockIClusterConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClusterConfigService) EXPECT() *MockIClusterConfigServiceMockRecorder {
	return m.recorder
}

// CheckInput mocks base method.
func (m *MockIClusterConfigService) CheckInput(configType string, config []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckInput", configType, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckInput indicates an expected call of CheckInput.
func (mr *MockIClusterConfigServiceMockRecorder) CheckInput(configType, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckInput", reflect.TypeOf((*MockIClusterConfigService)(nil).CheckInput), configType, config)
}

// Disable mocks base method.
func (m *MockIClusterConfigService) Disable(ctx context.Context, namespaceId, operator int, clusterName, configType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", ctx, namespaceId, operator, clusterName, configType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockIClusterConfigServiceMockRecorder) Disable(ctx, namespaceId, operator, clusterName, configType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockIClusterConfigService)(nil).Disable), ctx, namespaceId, operator, clusterName, configType)
}

// Edit mocks base method.
func (m *MockIClusterConfigService) Edit(ctx context.Context, namespaceId, operator int, clusterName, configType string, config []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", ctx, namespaceId, operator, clusterName, configType, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit.
func (mr *MockIClusterConfigServiceMockRecorder) Edit(ctx, namespaceId, operator, clusterName, configType, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockIClusterConfigService)(nil).Edit), ctx, namespaceId, operator, clusterName, configType, config)
}

// Enable mocks base method.
func (m *MockIClusterConfigService) Enable(ctx context.Context, namespaceId, operator int, clusterName, configType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", ctx, namespaceId, operator, clusterName, configType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockIClusterConfigServiceMockRecorder) Enable(ctx, namespaceId, operator, clusterName, configType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockIClusterConfigService)(nil).Enable), ctx, namespaceId, operator, clusterName, configType)
}

// FormatOutput mocks base method.
func (m *MockIClusterConfigService) FormatOutput(configType, operator string, config *cluster_entry.ClusterConfig) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatOutput", configType, operator, config)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// FormatOutput indicates an expected call of FormatOutput.
func (mr *MockIClusterConfigServiceMockRecorder) FormatOutput(configType, operator, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatOutput", reflect.TypeOf((*MockIClusterConfigService)(nil).FormatOutput), configType, operator, config)
}

// Get mocks base method.
func (m *MockIClusterConfigService) Get(ctx context.Context, namespaceId int, clusterName, configType string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, namespaceId, clusterName, configType)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIClusterConfigServiceMockRecorder) Get(ctx, namespaceId, clusterName, configType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIClusterConfigService)(nil).Get), ctx, namespaceId, clusterName, configType)
}

// IsConfigTypeExist mocks base method.
func (m *MockIClusterConfigService) IsConfigTypeExist(configType string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConfigTypeExist", configType)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConfigTypeExist indicates an expected call of IsConfigTypeExist.
func (mr *MockIClusterConfigServiceMockRecorder) IsConfigTypeExist(configType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConfigTypeExist", reflect.TypeOf((*MockIClusterConfigService)(nil).IsConfigTypeExist), configType)
}

// OfflineApinto mocks base method.
func (m *MockIClusterConfigService) OfflineApinto(client v1.IClient, name, configType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfflineApinto", client, name, configType)
	ret0, _ := ret[0].(error)
	return ret0
}

// OfflineApinto indicates an expected call of OfflineApinto.
func (mr *MockIClusterConfigServiceMockRecorder) OfflineApinto(client, name, configType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfflineApinto", reflect.TypeOf((*MockIClusterConfigService)(nil).OfflineApinto), client, name, configType)
}

// ResetOnline mocks base method.
func (m *MockIClusterConfigService) ResetOnline(ctx context.Context, namespaceId, clusterId int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetOnline", ctx, namespaceId, clusterId)
}

// ResetOnline indicates an expected call of ResetOnline.
func (mr *MockIClusterConfigServiceMockRecorder) ResetOnline(ctx, namespaceId, clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetOnline", reflect.TypeOf((*MockIClusterConfigService)(nil).ResetOnline), ctx, namespaceId, clusterId)
}

// ToApinto mocks base method.
func (m *MockIClusterConfigService) ToApinto(client v1.IClient, name, configType string, config []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToApinto", client, name, configType, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToApinto indicates an expected call of ToApinto.
func (mr *MockIClusterConfigServiceMockRecorder) ToApinto(client, name, configType, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToApinto", reflect.TypeOf((*MockIClusterConfigService)(nil).ToApinto), client, name, configType, config)
}

// MockIClusterNodeService is a mock of IClusterNodeService interface.
type MockIClusterNodeService struct {
	ctrl     *gomock.Controller
	recorder *MockIClusterNodeServiceMockRecorder
}

// MockIClusterNodeServiceMockRecorder is the mock recorder for MockIClusterNodeService.
type MockIClusterNodeServiceMockRecorder struct {
	mock *MockIClusterNodeService
}

// NewMockIClusterNodeService creates a new mock instance.
func NewMockIClusterNodeService(ctrl *gomock.Controller) *MockIClusterNodeService {
	mock := &MockIClusterNodeService{ctrl: ctrl}
	mock.recorder = &MockIClusterNodeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClusterNodeService) EXPECT() *MockIClusterNodeServiceMockRecorder {
	return m.recorder
}

// GetClusterInfo mocks base method.
func (m *MockIClusterNodeService) GetClusterInfo(addr string) (*v1.ClusterInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInfo", addr)
	ret0, _ := ret[0].(*v1.ClusterInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInfo indicates an expected call of GetClusterInfo.
func (mr *MockIClusterNodeServiceMockRecorder) GetClusterInfo(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInfo", reflect.TypeOf((*MockIClusterNodeService)(nil).GetClusterInfo), addr)
}

// GetNodesByUrl mocks base method.
func (m *MockIClusterNodeService) GetNodesByUrl(addr string) ([]*cluster_model.ClusterNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodesByUrl", addr)
	ret0, _ := ret[0].([]*cluster_model.ClusterNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodesByUrl indicates an expected call of GetNodesByUrl.
func (mr *MockIClusterNodeServiceMockRecorder) GetNodesByUrl(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodesByUrl", reflect.TypeOf((*MockIClusterNodeService)(nil).GetNodesByUrl), addr)
}

// Insert mocks base method.
func (m *MockIClusterNodeService) Insert(ctx context.Context, nodes []*cluster_model.ClusterNode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockIClusterNodeServiceMockRecorder) Insert(ctx, nodes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIClusterNodeService)(nil).Insert), ctx, nodes)
}

// NodeRepeatContrast mocks base method.
func (m *MockIClusterNodeService) NodeRepeatContrast(ctx context.Context, namespaceId, clusterId int, newList []*cluster_model.ClusterNode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeRepeatContrast", ctx, namespaceId, clusterId, newList)
	ret0, _ := ret[0].(error)
	return ret0
}

// NodeRepeatContrast indicates an expected call of NodeRepeatContrast.
func (mr *MockIClusterNodeServiceMockRecorder) NodeRepeatContrast(ctx, namespaceId, clusterId, newList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeRepeatContrast", reflect.TypeOf((*MockIClusterNodeService)(nil).NodeRepeatContrast), ctx, namespaceId, clusterId, newList)
}

// QueryByClusterIds mocks base method.
func (m *MockIClusterNodeService) QueryByClusterIds(ctx context.Context, clusterIds ...int) ([]*cluster_model.ClusterNode, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range clusterIds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryByClusterIds", varargs...)
	ret0, _ := ret[0].([]*cluster_model.ClusterNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryByClusterIds indicates an expected call of QueryByClusterIds.
func (mr *MockIClusterNodeServiceMockRecorder) QueryByClusterIds(ctx interface{}, clusterIds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, clusterIds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryByClusterIds", reflect.TypeOf((*MockIClusterNodeService)(nil).QueryByClusterIds), varargs...)
}

// QueryList mocks base method.
func (m *MockIClusterNodeService) QueryList(ctx context.Context, namespaceId int, clusterName string) ([]*cluster_model.ClusterNode, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryList", ctx, namespaceId, clusterName)
	ret0, _ := ret[0].([]*cluster_model.ClusterNode)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryList indicates an expected call of QueryList.
func (mr *MockIClusterNodeServiceMockRecorder) QueryList(ctx, namespaceId, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryList", reflect.TypeOf((*MockIClusterNodeService)(nil).QueryList), ctx, namespaceId, clusterName)
}

// Reset mocks base method.
func (m *MockIClusterNodeService) Reset(ctx context.Context, namespaceId, userId int, clusterName, clusterAddr, source string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset", ctx, namespaceId, userId, clusterName, clusterAddr, source)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockIClusterNodeServiceMockRecorder) Reset(ctx, namespaceId, userId, clusterName, clusterAddr, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockIClusterNodeService)(nil).Reset), ctx, namespaceId, userId, clusterName, clusterAddr, source)
}

// Update mocks base method.
func (m *MockIClusterNodeService) Update(ctx context.Context, namespaceId int, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, namespaceId, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIClusterNodeServiceMockRecorder) Update(ctx, namespaceId, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIClusterNodeService)(nil).Update), ctx, namespaceId, clusterName)
}
