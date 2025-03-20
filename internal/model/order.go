package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
)

type Order struct {
	OrderID       primitive.ObjectID `json:"orderID,omitempty" bson:"_id,omitempty"`
	Status        int16              `json:"status" bson:"status"`
	Products      []dto.OrderProduct  `json:"products" bson:"products"`
	AppointmentID primitive.ObjectID `json:"appointmentID" bson:"appointmentID"`
	SellerID      primitive.ObjectID `json:"sellerID" bson:"sellerID"`
	SellerName    string             `json:"sellerName" bson:"sellerName"`
	BuyerID       primitive.ObjectID `json:"buyerID" bson:"buyerID"`
	BuyerName     string             `json:"buyerName" bson:"buyerName"`
	TotalPrice    float64            `json:"totalPrice" bson:"totalPrice"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	Payment     string               `json:"payment" bson:"payment"`
}

