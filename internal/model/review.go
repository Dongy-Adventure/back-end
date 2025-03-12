package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ReviewID   primitive.ObjectID `json:"reviewID,omitempty" bson:"_id"`
	BuyerID    primitive.ObjectID `json:"buyerID" bson:"buyer_id"`
	BuyerName  string             `json:"buyerName" bson:"buyerName"`
	SellerID   primitive.ObjectID `json:"sellerID" bson:"seller_id"`
	SellerName string             `json:"sellerName" bson:"sellerName"`
	Image      string             `json:"image,omitempty" bson:"image,omitempty"`
	Message    string             `json:"message" bson:"message" binding:"max=500"`
	Score      int                `json:"score" bson:"score" binding:"gte=0,lte=10"`
	Date       time.Time          `json:"date" bson:"date"`
}
