package dto

import (
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
	Province    string             `json:"province"`
	City        string             `json:"city"`
	Zip         string             `json:"zip"`
	Transaction []Transaction      `json:"transaction"`
	Balance     float64            `json:"balance"`
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
	Province    string  `json:"province"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
}

type SellerWithdrawRequest struct {
	Payment     string             `json:"payment"`
	Amount  float64   		      `json:"amount"`
}
