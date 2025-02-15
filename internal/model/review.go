package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ReviewID primitive.ObjectID `json:"reviewID,omitempty" bson:"_id"`
	BuyerID  primitive.ObjectID `json:"buyerID" bson:"buyer_id"`
	BuyerName string `json:"buyerName" bson:"buyerName"`
	SellerID primitive.ObjectID `json:"sellerID" bson:"seller_id"`
	SellerName string `json:"sellerName" bson:"sellerName"`
	Image     string            `json:"image,omitempty" bson:"image,omitempty"`
	Message  string             `json:"message" bson:"message" validate:"max=500"`
	Score    int                `json:"score" bson:"score" validate:"min=0,max=10"`
	Date     time.Time          `json:"date" bson:"date"`
}

var validate = validator.New()

func (r *Review) Validate() error {
	return validate.Struct(r)
}