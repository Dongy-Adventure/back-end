// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/buyer_service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	dto "github.com/Dongy-s-Advanture/back-end/internal/dto"
	model "github.com/Dongy-s-Advanture/back-end/internal/model"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockIBuyerService is a mock of IBuyerService interface.
type MockIBuyerService struct {
	ctrl     *gomock.Controller
	recorder *MockIBuyerServiceMockRecorder
}

// MockIBuyerServiceMockRecorder is the mock recorder for MockIBuyerService.
type MockIBuyerServiceMockRecorder struct {
	mock *MockIBuyerService
}

// NewMockIBuyerService creates a new mock instance.
func NewMockIBuyerService(ctrl *gomock.Controller) *MockIBuyerService {
	mock := &MockIBuyerService{ctrl: ctrl}
	mock.recorder = &MockIBuyerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBuyerService) EXPECT() *MockIBuyerServiceMockRecorder {
	return m.recorder
}

// CreateBuyerData mocks base method.
func (m *MockIBuyerService) CreateBuyerData(buyer *model.Buyer) (*dto.Buyer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBuyerData", buyer)
	ret0, _ := ret[0].(*dto.Buyer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBuyerData indicates an expected call of CreateBuyerData.
func (mr *MockIBuyerServiceMockRecorder) CreateBuyerData(buyer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBuyerData", reflect.TypeOf((*MockIBuyerService)(nil).CreateBuyerData), buyer)
}

// GetBuyer mocks base method.
func (m *MockIBuyerService) GetBuyer() ([]dto.Buyer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuyer")
	ret0, _ := ret[0].([]dto.Buyer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuyer indicates an expected call of GetBuyer.
func (mr *MockIBuyerServiceMockRecorder) GetBuyer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuyer", reflect.TypeOf((*MockIBuyerService)(nil).GetBuyer))
}

// GetBuyerByID mocks base method.
func (m *MockIBuyerService) GetBuyerByID(buyerID primitive.ObjectID) (*dto.Buyer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuyerByID", buyerID)
	ret0, _ := ret[0].(*dto.Buyer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuyerByID indicates an expected call of GetBuyerByID.
func (mr *MockIBuyerServiceMockRecorder) GetBuyerByID(buyerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuyerByID", reflect.TypeOf((*MockIBuyerService)(nil).GetBuyerByID), buyerID)
}

// UpdateBuyerData mocks base method.
func (m *MockIBuyerService) UpdateBuyerData(buyerID primitive.ObjectID, updatedBuyer *model.Buyer) (*dto.Buyer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBuyerData", buyerID, updatedBuyer)
	ret0, _ := ret[0].(*dto.Buyer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBuyerData indicates an expected call of UpdateBuyerData.
func (mr *MockIBuyerServiceMockRecorder) UpdateBuyerData(buyerID, updatedBuyer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBuyerData", reflect.TypeOf((*MockIBuyerService)(nil).UpdateBuyerData), buyerID, updatedBuyer)
}

// UpdateProductInCart mocks base method.
func (m *MockIBuyerService) UpdateProductInCart(buyerID primitive.ObjectID, product dto.OrderProduct) ([]dto.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductInCart", buyerID, product)
	ret0, _ := ret[0].([]dto.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductInCart indicates an expected call of UpdateProductInCart.
func (mr *MockIBuyerServiceMockRecorder) UpdateProductInCart(buyerID, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductInCart", reflect.TypeOf((*MockIBuyerService)(nil).UpdateProductInCart), buyerID, product)
}
