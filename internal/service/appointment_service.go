package service

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAppointmentService interface {
	GetAppointments() ([]dto.Appointment, error)
	GetAppointmentByID(appointmentID primitive.ObjectID) (*dto.Appointment, error)
	GetAppointmentByOrderID(orderID primitive.ObjectID) (*dto.Appointment, error)
	CreateAppointment(appointment *model.Appointment) (*dto.Appointment, error)
	UpdateAppointmentDate(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error)
	UpdateAppointmentPlace(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error)
}

type AppointmentService struct {
	appointmentRepository repository.IAppointmentRepository
}

func NewAppointmentService(r repository.IAppointmentRepository) IAppointmentService {
	return AppointmentService{
		appointmentRepository: r,
	}
}

func (s AppointmentService) GetAppointments() ([]dto.Appointment, error) {
	appointments, err := s.appointmentRepository.GetAppointments()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (s AppointmentService) GetAppointmentByID(appointmentID primitive.ObjectID) (*dto.Appointment, error) {
	appointmentDTO, err := s.appointmentRepository.GetAppointmentByID(appointmentID)
	if err != nil {
		return nil, err
	}
	return appointmentDTO, nil
}

func (s AppointmentService) GetAppointmentByOrderID(orderID primitive.ObjectID) (*dto.Appointment, error) {
	appointmentDTO, err := s.appointmentRepository.GetAppointmentByOrderID(orderID)
	if err != nil {
		return nil, err
	}
	return appointmentDTO, nil
}

func (s AppointmentService) CreateAppointment(appointment *model.Appointment) (*dto.Appointment, error) {

	newAppointment, err := s.appointmentRepository.CreateAppointment(appointment)

	if err != nil {
		return nil, err
	}

	return newAppointment, nil
}


func (s AppointmentService) UpdateAppointmentDate(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error) {

	updatedAppointmentDTO, err := s.appointmentRepository.UpdateAppointmentDate(appointmentID, updatedAppointment)
	if err != nil {
		return nil, err
	}

	return updatedAppointmentDTO, nil
}

func (s AppointmentService) UpdateAppointmentPlace(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error) {

	updatedAppointmentDTO, err := s.appointmentRepository.UpdateAppointmentPlace(appointmentID, updatedAppointment)
	if err != nil {
		return nil, err
	}

	return updatedAppointmentDTO, nil
}


