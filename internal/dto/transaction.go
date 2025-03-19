package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionType string

const (
    Credit TransactionType = "credit" // Money added
    Debit  TransactionType = "debit"  // Money withdrawn
)

type Transaction struct {
	Type    TransactionType 	   `json:"type"`
	Amount  float64   		   `json:"amount"`
	OrderID  primitive.ObjectID `json:"orderID,omitempty"`
	Payment  string             `json:"payment"`
	Date    time.Time            `json:"data"`
}



