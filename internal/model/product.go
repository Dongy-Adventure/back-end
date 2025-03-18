package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ProductID   primitive.ObjectID `json:"productID,omitempty" bson:"_id,omitempty"`
	ProductName string             `json:"productName" bson:"productName" binding:"required"`
	Price       float64            `json:"price,omitempty" bson:"price" binding:"required,gte=1"`
	Description string             `json:"description,omitempty" bson:"description"`
	ImageURL    string             `json:"imageURL,omitempty" bson:"imageURL"`
	Tag         []string           `json:"tag,omitempty" bson:"tag"`
	Color       string             `json:"color,omitempty" bson:"color"`
	SellerID    primitive.ObjectID `json:"sellerID,omitempty" bson:"sellerID"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	Amount      int                `json:"amount" bson:"amount" binding:"required,gte=0"`
}
