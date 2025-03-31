package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advertisement struct {
	AdvertisementID primitive.ObjectID `json:"advertisementID,omitempty" bson:"_id"`
	SellerID        primitive.ObjectID `json:"sellerID" bson:"seller_id"`
	ProductID       primitive.ObjectID `json:"productID" bson:"product_id"`
	ImageURL        string             `json:"imageURL,omitempty" bson:"imageURL"`
	Amount          float64            `json:"amount" bson:"amount" binding:"gte=0"`
	Payment         string             `json:"payment" bson:"payment"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
}
