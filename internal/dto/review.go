package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ReviewID primitive.ObjectID  `json:"reviewID"`
	BuyerID  primitive.ObjectID `json:"buyerID"`
	SellerID primitive.ObjectID `json:"sellerID"`
	Image     string            `json:"image,omitempty"`
	Message  string             `json:"message"`
	Score    int                `json:"score"`
	Date     time.Time          `json:"date"`
}

type ReviewCreateRequest struct {
	BuyerID  primitive.ObjectID `json:"buyerID"`
	SellerID primitive.ObjectID `json:"sellerID"`
	Image     string            `json:"image,omitempty"`
	Message  string             `json:"message"`
	Score    int                `json:"score"`
}

type ReviewUpdateRequest struct {
	Image     string            `json:"image,omitempty"`
	Message  string             `json:"message"`
	Score    int                `json:"score"`
}




