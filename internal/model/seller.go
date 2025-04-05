package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seller struct {
	SellerID    primitive.ObjectID `json:"sellerID,omitempty" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password" copier:"-"`
	Name        string             `json:"name" bson:"name"`
	Surname     string             `json:"surname" bson:"surname"`
	Payment     string             `json:"payment" bson:"payment"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Address     string             `json:"address" bson:"address"`
	City        string             `json:"city" bson:"city"`
	Province    string             `json:"province" bson:"province"`
	Zip         string             `json:"zip" bson:"zip"`
	Score       float64            `json:"score" bson:"score,omitempty"`
	Transaction []Transaction      `json:"transaction" bson:"transaction"`
	Balance     float64            `json:"balance" bson:"balance"`
	ProfilePic  string             `json:"profilePic" bson:"profilePic"`
}
