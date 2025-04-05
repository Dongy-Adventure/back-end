package dto

import (
	"mime/multipart"

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
	ProfilePic  string             `json:"profilePic"`
}

type BuyerRegisterRequest struct {
	Username    string                `json:"username" form:"username"`
	Password    string                `json:"password" form:"password"`
	Name        string                `json:"name" form:"name"`
	Surname     string                `json:"surname" form:"surname"`
	Payment     string                `json:"payment" form:"payment"`
	Address     string                `json:"address" form:"address"`
	PhoneNumber string                `json:"phoneNumber" form:"phoneNumber"`
	Province    string                `json:"province" form:"province"`
	City        string                `json:"city" form:"city"`
	Zip         string                `json:"zip" form:"zip"`
	ProfilePic  *multipart.FileHeader `json:"profilePic" form:"profilePic"`
}

type UpdateCartRequest struct {
	Product Product `json:"product"`
}
