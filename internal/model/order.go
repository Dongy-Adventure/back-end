package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID       primitive.ObjectID `json:"orderID,omitempty" bson:"_id,omitempty"`
	Status        int16              `json:"status" bson:"status"`
	Products      []OrderProduct  `json:"products" bson:"products"`
	AppointmentID primitive.ObjectID `json:"appointmentID" bson:"appointmentID"`
	SellerID      primitive.ObjectID `json:"sellerID" bson:"sellerID"`
	SellerName    string             `json:"sellerName" bson:"sellerName"`
	BuyerID       primitive.ObjectID `json:"buyerID" bson:"buyerID"`
	BuyerName     string             `json:"buyerName" bson:"buyerName"`
	TotalPrice    float64            `json:"totalPrice" bson:"totalPrice"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	Payment     string               `json:"payment" bson:"payment"`
}

type OrderProduct struct {
	ProductID primitive.ObjectID    `json:"productID" bson:"_id"`
	Amount  int                     `json:"amount" bson:"amount" binding:"required,gte=0"`
}
 

