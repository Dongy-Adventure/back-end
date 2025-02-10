package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ReviewID primitive.ObjectID `json:"reviewID,omitempty" bson:"_id"`
	BuyerID  primitive.ObjectID `json:"buyerID" bson:"buyer_id"`
	SellerID primitive.ObjectID `json:"sellerID" bson:"seller_id"`
	Image     string            `json:"image,omitempty" bson:"image,omitempty"`
	Message  string             `json:"message" bson:"message"`
	Score    int                `json:"score" bson:"score"`
	Date     time.Time          `json:"date" bson:"date"`
}
