package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seller struct {
	SellerID    primitive.ObjectID `json:"sellerID"`
	Username    string             `json:"username"`
	Name        string             `json:"name"`
	Surname     string             `json:"surname"`
	Payment     string             `json:"payment"`
	Address     string             `json:"address"`
	PhoneNumber string             `json:"phoneNumber"`
	Score       float32            `json:"score"`
	Transaction []Transaction      `json:"transaction"`
}

type SellerRegisterRequest struct {
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	Payment     string  `json:"payment"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phoneNumber"`
	Score       float32 `json:"score"`
}

type Transaction struct {
	Amount  float64   `json:"amount" bson:"amount" validate:"min=0"`
	Product []string  `json:"product" bson:"product"`
	Date    time.Time `json:"data" bson:"date"`
}
