package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
)



type Transaction struct {
	Type    dto.TransactionType 	   `json:"type" bson:"type"`
	Amount  float64   		   `json:"amount" bson:"amount" binding:"gte=0"`
	OrderID  primitive.ObjectID `json:"orderID,omitempty" bson:"_id,omitempty"`
	Payment  string             `json:"payment" bson:"payment"`
	Date    time.Time            `json:"data" bson:"date"`
}