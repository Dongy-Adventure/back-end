// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/seller_service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	dto "github.com/Dongy-s-Advanture/back-end/internal/dto"
	model "github.com/Dongy-s-Advanture/back-end/internal/model"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockISellerService is a mock of ISellerService interface.
type MockISellerService struct {
	ctrl     *gomock.Controller
	recorder *MockISellerServiceMockRecorder
}

// MockISellerServiceMockRecorder is the mock recorder for MockISellerService.
type MockISellerServiceMockRecorder struct {
	mock *MockISellerService
}

// NewMockISellerService creates a new mock instance.
func NewMockISellerService(ctrl *gomock.Controller) *MockISellerService {
	mock := &MockISellerService{ctrl: ctrl}
	mock.recorder = &MockISellerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISellerService) EXPECT() *MockISellerServiceMockRecorder {
	return m.recorder
}

// // AddTransaction mocks base method.
// func (m *MockISellerService) AddTransaction(sellerID primitive.ObjectID, transaction *dto.Transaction) (*dto.Transaction, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "AddTransaction", sellerID, transaction)
// 	ret0, _ := ret[0].(*dto.Transaction)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// // AddTransaction indicates an expected call of AddTransaction.
// func (mr *MockISellerServiceMockRecorder) AddTransaction(sellerID, transaction interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockISellerService)(nil).AddTransaction), sellerID, transaction)
// }

// CreateSellerData mocks base method.
func (m *MockISellerService) CreateSellerData(seller *model.Seller) (*dto.Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSellerData", seller)
	ret0, _ := ret[0].(*dto.Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSellerData indicates an expected call of CreateSellerData.
func (mr *MockISellerServiceMockRecorder) CreateSellerData(seller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSellerData", reflect.TypeOf((*MockISellerService)(nil).CreateSellerData), seller)
}

// GetSellerBalanceByID mocks base method.
func (m *MockISellerService) GetSellerBalanceByID(sellerID primitive.ObjectID) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSellerBalanceByID", sellerID)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSellerBalanceByID indicates an expected call of GetSellerBalanceByID.
func (mr *MockISellerServiceMockRecorder) GetSellerBalanceByID(sellerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSellerBalanceByID", reflect.TypeOf((*MockISellerService)(nil).GetSellerBalanceByID), sellerID)
}

// GetSellerByID mocks base method.
func (m *MockISellerService) GetSellerByID(sellerID primitive.ObjectID) (*dto.Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSellerByID", sellerID)
	ret0, _ := ret[0].(*dto.Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSellerByID indicates an expected call of GetSellerByID.
func (mr *MockISellerServiceMockRecorder) GetSellerByID(sellerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSellerByID", reflect.TypeOf((*MockISellerService)(nil).GetSellerByID), sellerID)
}

// GetSellers mocks base method.
func (m *MockISellerService) GetSellers() ([]dto.Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSellers")
	ret0, _ := ret[0].([]dto.Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSellers indicates an expected call of GetSellers.
func (mr *MockISellerServiceMockRecorder) GetSellers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSellers", reflect.TypeOf((*MockISellerService)(nil).GetSellers))
}

// UpdateSeller mocks base method.
func (m *MockISellerService) UpdateSeller(sellerID primitive.ObjectID, updatedSeller *model.Seller) (*dto.Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSeller", sellerID, updatedSeller)
	ret0, _ := ret[0].(*dto.Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSeller indicates an expected call of UpdateSeller.
func (mr *MockISellerServiceMockRecorder) UpdateSeller(sellerID, updatedSeller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSeller", reflect.TypeOf((*MockISellerService)(nil).UpdateSeller), sellerID, updatedSeller)
}
