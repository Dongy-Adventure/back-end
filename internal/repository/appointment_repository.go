package repository

import (
	"context"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAppointmentRepository interface {
	GetAppointments() ([]dto.Appointment, error)
	GetAppointmentByID(appointmentID primitive.ObjectID) (*dto.Appointment, error)
	GetAppointmentByOrderID(orderID primitive.ObjectID) (*dto.Appointment, error)
	CreateAppointment(appointment *model.Appointment) (*dto.Appointment, error)
	UpdateAppointmentDate(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error)
	UpdateAppointmentPlace(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error)

}

type AppointmentRepository struct {
	appointmentCollection *mongo.Collection
}

func NewAppointmentRepository(db *mongo.Database, collectionName string) IAppointmentRepository {
	return AppointmentRepository{
		appointmentCollection: db.Collection(collectionName),
	}
}

func (r AppointmentRepository) GetAppointments() ([]dto.Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var appointmentList []dto.Appointment

	dataList, err := r.appointmentCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var appointmentModel *model.Appointment
		if err = dataList.Decode(&appointmentModel); err != nil {
			return nil, err
		}
		appointmentDTO, appointmentErr := converter.AppointmentModelToDTO(appointmentModel)
		if appointmentErr != nil {
			return nil, err
		}
		appointmentList = append(appointmentList, *appointmentDTO)
	}

	return appointmentList, nil
}

func (r AppointmentRepository) GetAppointmentByID(appointmentID primitive.ObjectID) (*dto.Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var appointment *model.Appointment

	err := r.appointmentCollection.FindOne(ctx, bson.M{"_id": appointmentID}).Decode(&appointment)
	if err != nil {
		return nil, err
	}
	return converter.AppointmentModelToDTO(appointment)
}

func (r AppointmentRepository) GetAppointmentByOrderID(orderID primitive.ObjectID) (*dto.Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var appointment *model.Appointment

	err := r.appointmentCollection.FindOne(ctx, bson.M{"order_id": orderID}).Decode(&appointment)
	if err != nil {
		return nil, err
	}
	return converter.AppointmentModelToDTO(appointment)
}

func (r AppointmentRepository) CreateAppointment(appointment *model.Appointment) (*dto.Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	appointment.AppointmentID = primitive.NewObjectID()
	appointment.CreatedAt = time.Now()
	result, err := r.appointmentCollection.InsertOne(ctx, appointment)
	if err != nil {
		return nil, err
	}
	var newAppointment *model.Appointment
	err = r.appointmentCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newAppointment)
	if err != nil {
		return nil, err
	}

	return converter.AppointmentModelToDTO(newAppointment)
}


func (r AppointmentRepository) UpdateAppointmentDate(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Date        time.Time            `json:"date" bson:"date"`
	// TimeSlot    string               `json:"timeSlot" bson:"time_slot"`

	update := bson.M{
		"$set": bson.M{
			"date":    updatedAppointment.Date,
        		"time_slot":   updatedAppointment.TimeSlot,
		},
	}

	filter := bson.M{"_id": appointmentID}
	_, err := r.appointmentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var newUpdatedAppointment *model.Appointment
	err = r.appointmentCollection.FindOne(ctx, filter).Decode(&newUpdatedAppointment)
	if err != nil {
		return nil, err
	}

	return converter.AppointmentModelToDTO(newUpdatedAppointment)
}

func (r AppointmentRepository) UpdateAppointmentPlace(appointmentID primitive.ObjectID, updatedAppointment *model.Appointment) (*dto.Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Address     string               `json:"address" bson:"address"`
	// City        string               `json:"city" bson:"city"`
	// Province    string               `json:"province" bson:"province"`
	// Zip         string               `json:"zip" bson:"zip"`

	update := bson.M{
		"$set": bson.M{
			"address":    updatedAppointment.Address,
        		"city":       updatedAppointment.City,
			"province":   updatedAppointment.Province,
			"zip":        updatedAppointment.Zip,
		},
	}

	filter := bson.M{"_id": appointmentID}
	_, err := r.appointmentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var newUpdatedAppointment *model.Appointment
	err = r.appointmentCollection.FindOne(ctx, filter).Decode(&newUpdatedAppointment)
	if err != nil {
		return nil, err
	}

	return converter.AppointmentModelToDTO(newUpdatedAppointment)
}


