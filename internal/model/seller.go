package model

import (
	"time"

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
	Address     string             `json:"address" bson:"addresss"`
	Score       float64            `json:"score" bson:"score"`
	Transaction []Transaction      `json:"transaction" bson:"transaction"`
}

type Transaction struct {
	Amount  float64   `json:"amount" bson:"amount" validate:"min=0"`
	Product []string  `json:"product" bson:"product"`
	Date    time.Time `json:"data" bson:"date"`
}
