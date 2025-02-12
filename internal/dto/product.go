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
	SellerID    primitive.ObjectID `json:"SellerID,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
}
type ProductPost struct {
	ProductName string             `json:"productName"`
	Price       float64            `json:"price,omitempty"`
	Description string             `json:"description,omitempty"`
	ImageURL    string             `json:"imageURL,omitempty"`
	Tag         []string           `json:"tag,omitempty"`
	Color       string             `json:"color,omitempty"`
	SellerID    primitive.ObjectID `json:"SellerID,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
}
type UpdateProductRequest struct {
	ProductName string   `json:"productName"`
	Price       float64  `json:"price,omitempty"`
	Description string   `json:"description,omitempty"`
	ImageURL    string   `json:"imageURL,omitempty"`
	Tag         []string `json:"tag,omitempty"`
	Color       string   `json:"color,omitempty"`
}
