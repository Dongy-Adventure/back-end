package dto

import (
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
	Cart        []OrderProduct     `json:"cart"`
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
