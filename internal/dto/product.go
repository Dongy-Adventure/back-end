package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ProductID   primitive.ObjectID `json:"productID,omitempty"`
	ProductName string             `json:"productName"`
	Price       float64            `json:"price,omitempty"`
	Description string             `json:"description,omitempty"`
	ImageURL    string             `json:"imageURL,omitempty"`
	Tag         []string           `json:"tag,omitempty"`
	Color       string             `json:"color,omitempty"`
	SellerID    primitive.ObjectID `json:"sellerID,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
	Amount      int                `json:"amount"`
}
type ProductCreateRequest struct {
	ProductName string             `json:"productName" binding:"required"`
	Price       float64            `json:"price,omitempty" binding:"required,gte=1"`
	Description string             `json:"description,omitempty"`
	ImageURL    string             `json:"imageURL,omitempty"`
	Tag         []string           `json:"tag,omitempty"`
	Color       string             `json:"color,omitempty"`
	SellerID    primitive.ObjectID `json:"sellerID,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
	Amount      int                `json:"amount" binding:"required,gte=0"`
}
type UpdateProductRequest struct {
	ProductName string             `json:"productName" binding:"required"`
	Price       float64            `json:"price,omitempty" binding:"required,gte=1"`
	Description string             `json:"description,omitempty"`
	ImageURL    string             `json:"imageURL,omitempty"`
	Tag         []string           `json:"tag,omitempty"`
	Color       string             `json:"color,omitempty"`
	SellerID    primitive.ObjectID `json:"sellerID,omitempty"`
	Amount      int                `json:"amount" binding:"required,gte=0"`
}
