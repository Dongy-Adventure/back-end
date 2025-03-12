package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID       primitive.ObjectID `json:"orderID,omitempty"`
	Status        int16              `json:"status"`
	Products      []Product          `json:"products"`
	AppointmentID primitive.ObjectID `json:"appointmentID"`
	SellerID      primitive.ObjectID `json:"sellerID"`
	SellerName    string             `json:"sellerName"`
	BuyerID       primitive.ObjectID `json:"buyerID"`
	BuyerName     string             `json:"buyerName"`
	TotalPrice    float64            `json:"totalPrice"`
	CreatedAt     time.Time          `json:"createdAt"`
}
type OrderCreateRequest struct {
	Products []Product          `json:"products"`
	BuyerID  primitive.ObjectID `json:"buyerID"`
	SellerID primitive.ObjectID `json:"sellerID"`
}

type OrderStatusRequest struct {
	OrderStatus int `json:"orderStatus" binding:"required,gte=0,lte=3"`
}
