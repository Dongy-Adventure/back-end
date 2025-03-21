// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/appointment_service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	dto "github.com/Dongy-s-Advanture/back-end/internal/dto"
	model "github.com/Dongy-s-Advanture/back-end/internal/model"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockIAppointmentService is a mock of IAppointmentService interface.
type MockIAppointmentService struct {
	ctrl     *gomock.Controller
	recorder *MockIAppointmentServiceMockRecorder
}

// MockIAppointmentServiceMockRecorder is the mock recorder for MockIAppointmentService.
type MockIAppointmentServiceMockRecorder struct {
	mock *MockIAppointmentService
}

// NewMockIAppointmentService creates a new mock instance.
func NewMockIAppointmentService(ctrl *gomock.Controller) *MockIAppointmentService {
	mock := &MockIAppointmentService{ctrl: ctrl}
	mock.recorder = &MockIAppointmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAppointmentService) EXPECT() *MockIAppointmentServiceMockRecorder {
	return m.recorder
}

// CreateAppointment mocks base method.
func (m *MockIAppointmentService) CreateAppointment(appointment *model.Appointment) (*dto.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppointment", appointment)
	ret0, _ := ret[0].(*dto.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAppointment indicates an expected call of CreateAppointment.
func (mr *MockIAppointmentServiceMockRecorder) CreateAppointment(appointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppointment", reflect.TypeOf((*MockIAppointmentService)(nil).CreateAppointment), appointment)
}

// GetAppointmentByID mocks base method.
func (m *MockIAppointmentService) GetAppointmentByID(appointmentID primitive.ObjectID) (*dto.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppointmentByID", appointmentID)
	ret0, _ := ret[0].(*dto.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppointmentByID indicates an expected call of GetAppointmentByID.
func (mr *MockIAppointmentServiceMockRecorder) GetAppointmentByID(appointmentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppointmentByID", reflect.TypeOf((*MockIAppointmentService)(nil).GetAppointmentByID), appointmentID)
}

// GetAppointmentByOrderID mocks base method.
func (m *MockIAppointmentService) GetAppointmentByOrderID(orderID primitive.ObjectID) (*dto.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppointmentByOrderID", orderID)
	ret0, _ := ret[0].(*dto.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppointmentByOrderID indicates an expected call of GetAppointmentByOrderID.
func (mr *MockIAppointmentServiceMockRecorder) GetAppointmentByOrderID(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppointmentByOrderID", reflect.TypeOf((*MockIAppointmentService)(nil).GetAppointmentByOrderID), orderID)
}

// GetAppointments mocks base method.
func (m *MockIAppointmentService) GetAppointments() ([]dto.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppointments")
	ret0, _ := ret[0].([]dto.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppointments indicates an expected call of GetAppointments.
func (mr *MockIAppointmentServiceMockRecorder) GetAppointments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppointments", reflect.TypeOf((*MockIAppointmentService)(nil).GetAppointments))
}

// UpdateAppointmentDate mocks base method.
func (m *MockIAppointmentService) UpdateAppointmentDate(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppointmentDate", appointmentID, updatedAppointment)
	ret0, _ := ret[0].(*dto.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAppointmentDate indicates an expected call of UpdateAppointmentDate.
func (mr *MockIAppointmentServiceMockRecorder) UpdateAppointmentDate(appointmentID, updatedAppointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppointmentDate", reflect.TypeOf((*MockIAppointmentService)(nil).UpdateAppointmentDate), appointmentID, updatedAppointment)
}

// UpdateAppointmentPlace mocks base method.
func (m *MockIAppointmentService) UpdateAppointmentPlace(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppointmentPlace", appointmentID, updatedAppointment)
	ret0, _ := ret[0].(*dto.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAppointmentPlace indicates an expected call of UpdateAppointmentPlace.
func (mr *MockIAppointmentServiceMockRecorder) UpdateAppointmentPlace(appointmentID, updatedAppointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppointmentPlace", reflect.TypeOf((*MockIAppointmentService)(nil).UpdateAppointmentPlace), appointmentID, updatedAppointment)
}
