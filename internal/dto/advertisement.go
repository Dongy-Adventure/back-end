package dto

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advertisement struct {
	AdvertisementID primitive.ObjectID `json:"advertisementID,omitempty"`
	SellerID        primitive.ObjectID `json:"sellerID"`
	ProductID       primitive.ObjectID `json:"productID"`
	ImageURL        string             `json:"imageURL,omitempty"`
	Amount          float64            `json:"amount"`
	Payment         string             `json:"payment"`
	CreatedAt       time.Time          `json:"createdAt"`
}

type AdvertisementCreateRequest struct {
	SellerID  primitive.ObjectID    `json:"sellerID" form:"sellerID"`
	ProductID primitive.ObjectID    `json:"productID" form:"productID"`
	ImageURL  *multipart.FileHeader `json:"imageURL,omitempty" form:"imageURL" swaggerignore:"true"`
	Amount    float64               `json:"amount" form:"amount"`
	Payment   string                `json:"payment" form:"payment"`
}

type AdvertisementUpdateRequest struct {
	ProductID primitive.ObjectID    `json:"productID,omitempty" form:"productID"`
	ImageURL  *multipart.FileHeader `json:"imageURL,omitempty" form:"imageURL" swaggerignore:"true"`
	Amount    float64               `json:"amount" form:"amount"`
	Payment   string                `json:"payment" form:"payment"`
}
