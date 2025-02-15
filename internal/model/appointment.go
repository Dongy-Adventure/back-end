package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	AppointmentID primitive.ObjectID `json:"appointmentID,omitempty" bson:"_id"`
	OrderID  primitive.ObjectID 	   `json:"orderID" bson:"order_id"`
	BuyerID  primitive.ObjectID 	   `json:"buyerID" bson:"buyer_id"`
	SellerID primitive.ObjectID      `json:"sellerID" bson:"seller_id"`
	Address     string               `json:"address" bson:"address"`
	City        string               `json:"city" bson:"city"`
	Province    string               `json:"province" bson:"province"`
	Zip         string               `json:"zip" bson:"zip"`
	Date        time.Time            `json:"date" bson:"date"`
	TimeSlot    string               `json:"timeSlot" bson:"time_slot"`
	CreatedAt   time.Time            `json:"createdAt" bson:"created_at"`
}

