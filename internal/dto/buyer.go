package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Buyer struct {
	BuyerID     primitive.ObjectID `json:"buyerID"`
	Username    string             `json:"username"`
	Password    string             `json:"password"`
	Name        string             `json:"name"`
	Surname     string             `json:"surname"`
	Payment     string             `json:"payment"`
	PhoneNumber string             `json:"phoneNumber"`
	Address     string             `json:"address"`
	City        string             `json:"city"`
	Province    string             `json:"province"`
	Zip         string             `json:"zip"`
	Cart        []Product          `json:"cart"`
	Transaction []Transaction      `json:"transaction"`
}

type BuyerRegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Payment     string `json:"payment"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
}

type UpdateCartRequest struct {
	Product Product `json:"product"`
}

type BuyerPaymentRequest struct {
	BuyerID       string    `json:"buyerID" binding:"required"`
	PaymentMethod string    `json:"paymentMethod" binding:"required"`
	Amount        float64   `json:"amount" binding:"required"`
	CreatedAt     time.Time `json:"createdAt" binding:"required"`
}
