package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Buyer struct {
	BuyerID  primitive.ObjectID `json:"buyerID"`
	Username string             `json:"username"`
	Name     string             `json:"name"`
	Surname  string             `json:"surname"`
}
