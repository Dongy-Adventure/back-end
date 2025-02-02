package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Buyer struct {
	BuyerID     primitive.ObjectID `json:"buyerID"`
	Username    string             `json:"username"`
	Name        string             `json:"name"`
	Surname     string             `json:"surname"`
	Address     string             `json:"address"`
	PhoneNumber string             `json:"phoneNumber"`
}

type BuyerRegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Payment     string `json:"payment"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}
