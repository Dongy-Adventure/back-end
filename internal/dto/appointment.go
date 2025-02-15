package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	AppointmentID primitive.ObjectID `json:"appointmentID,omitempty"`
	OrderID  primitive.ObjectID 	   `json:"orderID"`
	BuyerID  primitive.ObjectID 	   `json:"buyerID"`
	SellerID primitive.ObjectID      `json:"sellerID"`
	Address     string               `json:"address"`
	City        string               `json:"city"`
	Province    string               `json:"province"`
	Zip         string               `json:"zip"`
	Date    time.Time                `json:"date"`
	TimeSlot    string               `json:"timeSlot"`
	CreatedAt    time.Time           `json:"createdAt"`
}

type AppointmentCreateRequest struct {
	OrderID  primitive.ObjectID 	   `json:"orderID"`
	BuyerID  primitive.ObjectID  `json:"buyerID"`
	SellerID primitive.ObjectID  `json:"sellerID"`
}

type AppointmentPlaceRequest struct {
	Address     string               `json:"address"`
	City        string               `json:"city"`
	Province    string               `json:"province"`
	Zip         string               `json:"zip"`
}

type AppointmentDateRequest struct {
	Date        string              `json:"date"`
	TimeSlot    string               `json:"timeSlot"`
}







