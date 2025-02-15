package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID       primitive.ObjectID `json:"orderID,omitempty" bson:"_id,omitempty"`
	Status        int16              `json:"status" bson:"status"`
	Products      []Product          `json:"products" bson:"products"`
	AppointmentID primitive.ObjectID `json:"appointmentID" bson:"appointmentID"`
	SellerID      primitive.ObjectID `json:"sellerID" bson:"sellerID"`
	BuyerID       primitive.ObjectID `json:"buyerID" bson:"buyerID"`
	TotalPrice    float64            `json:"totalPrice" bson:"totalPrice"`
	CreatedAt     time.Time          `json:"createdAt"`
}
