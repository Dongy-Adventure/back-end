package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-playground/validator/v10"
)

type Review struct {
	ReviewID primitive.ObjectID `json:"reviewID,omitempty" bson:"_id"`
	BuyerID  primitive.ObjectID `json:"buyerID" bson:"buyer_id"`
	SellerID primitive.ObjectID `json:"sellerID" bson:"seller_id"`
	Image     string            `json:"image,omitempty" bson:"image,omitempty"`
	Message  string             `json:"message" bson:"message" validate:"max=500"`
	Score    int                `json:"score" bson:"score" validate:"min=0,max=10"`
	Date     time.Time          `json:"date" bson:"date"`
}

var validate = validator.New()

func (r *Review) Validate() error {
	return validate.Struct(r)
}