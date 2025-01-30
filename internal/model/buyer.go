package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Buyer struct {
	BuyerID  primitive.ObjectID `json:"buyerID,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password" copier:"-"`
	Name     string             `json:"name" bson:"name"`
	Surname  string             `json:"surname" bson:"surname"`
}
