package dto

import (
	"mime/multipart"

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
	ProfilePic  string             `json:"profilePic"`
}

type SellerRegisterRequest struct {
	Username    string                `json:"username" form:"username"`
	Password    string                `json:"password" form:"password"`
	Name        string                `json:"name" form:"name"`
	Surname     string                `json:"surname" form:"surname"`
	Payment     string                `json:"payment" form:"payment"`
	Address     string                `json:"address" form:"address"`
	PhoneNumber string                `json:"phoneNumber" form:"phoneNumber"`
	Score       float32               `json:"score" form:"score"`
	Province    string                `json:"province" form:"province"`
	City        string                `json:"city" form:"city"`
	Zip         string                `json:"zip" form:"zip"`
	ProfilePic  *multipart.FileHeader `json:"profilePic" form:"profilePic"`
}

type SellerWithdrawRequest struct {
	Payment string  `json:"payment"`
	Amount  float64 `json:"amount"`
}
