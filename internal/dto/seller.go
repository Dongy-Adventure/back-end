package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Seller struct {
	SellerID primitive.ObjectID `json:"sellerID"`
	Username string             `json:"username"`
	Name     string             `json:"name"`
	Surname  string             `json:"surname"`
	Payment  string             `json:"payment"`
}
