package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID       primitive.ObjectID `json:"orderID,omitempty"`
	Status        int16              `json:"status"`
	Products      []OrderProduct     `json:"products"`
	AppointmentID primitive.ObjectID `json:"appointmentID"`
	SellerID      primitive.ObjectID `json:"sellerID"`
	SellerName    string             `json:"sellerName"`
	BuyerID       primitive.ObjectID `json:"buyerID"`
	BuyerName     string             `json:"buyerName"`
	TotalPrice    float64            `json:"totalPrice"`
	CreatedAt     time.Time          `json:"createdAt"`
	Payment       string             `json:"payment"`
}
type OrderCreateRequest struct {
	Products    []OrderProduct     `json:"products"`
	BuyerID     primitive.ObjectID `json:"buyerID"`
	SellerID    primitive.ObjectID `json:"sellerID"`
	BuyerName   string             `json:"buyerName"`
	SellerName  string             `json:"sellerName"`
	Payment		string			   `json:"payment"`
	CreatedAt	time.Time          `json:"createdAt"`				
}

type OrderStatusRequest struct {
	OrderStatus int `json:"orderStatus" binding:"required,gte=0,lte=3"`
}

type OrderProduct struct {
	ProductID primitive.ObjectID `json:"productID"`
	Amount    int                `json:"amount"`
}
type PaymentRequest struct {
	BuyerID       string    `json:"buyerID" binding:"required"`
	PaymentMethod string    `json:"paymentMethod" binding:"required"`
	Amount        int64     `json:"amount" binding:"required"`
	Address       string    `json:"address,omitempty"`
	City          string    `json:"city,omitempty"`
	Province      string    `json:"province,omitempty"`
	Zip           string    `json:"zip,omitempty"`
	Token         string    `json:"token" binding:"required"`
	CreatedAt     time.Time `json:"createdAt" binding:"required"`
}
