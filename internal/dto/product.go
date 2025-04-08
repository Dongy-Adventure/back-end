package dto

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ProductID   primitive.ObjectID `json:"productID,omitempty"`
	ProductName string             `json:"productName"`
	Price       float64            `json:"price,omitempty"`
	Description string             `json:"description,omitempty"`
	Image       string             `json:"image,omitempty"`
	Tag         []string           `json:"tag,omitempty"`
	Color       string             `json:"color,omitempty"`
	SellerID    primitive.ObjectID `json:"sellerID,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
	Amount      int                `json:"amount"`
}
type ProductCreateRequest struct {
	ProductName string                `json:"productName" binding:"required" form:"productName"`
	Price       float64               `json:"price,omitempty" binding:"required,gte=1" form:"price"`
	Description string                `json:"description,omitempty" form:"description"`
	Tag         []string              `json:"tag,omitempty" form:"tag[]"`
	Color       string                `json:"color,omitempty" form:"color"`
	SellerID    string                `json:"sellerID,omitempty" form:"sellerID"`
	CreatedAt   time.Time             `json:"createdAt,omitempty" form:"createdAt"`
	Amount      int                   `json:"amount" binding:"required,gte=0" form:"amount"`
	Image       *multipart.FileHeader `json:"image,omitempty" form:"image" swaggerignore:"true"`
}

type UpdateProductRequest struct {
	ProductName string    `json:"productName" binding:"required"`
	Price       float64   `json:"price,omitempty" binding:"required,gte=1"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	Tag         []string  `json:"tag,omitempty"`
	Color       string    `json:"color,omitempty"`
	SellerID    string    `json:"sellerID,omitempty"`
	Amount      int       `json:"amount" binding:"required,gte=0"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
