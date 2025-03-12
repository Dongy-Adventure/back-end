package model

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Buyer struct {
	BuyerID     primitive.ObjectID `json:"buyerID,omitempty" bson:"_id,omitempty"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password" copier:"-"`
	Name        string             `json:"name" bson:"name"`
	Surname     string             `json:"surname" bson:"surname"`
	Payment     string             `json:"payment" bson:"payment"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Address     string             `json:"address" bson:"addresss"`
	City        string             `json:"city" bson:"city"`
	Province    string             `json:"province" bson:"province"`
	Zip         string             `json:"zip" bson:"zip"`
	Cart        []dto.Product      `json:"cart" bson:"cart"`
}
