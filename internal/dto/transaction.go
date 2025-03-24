package dto

import (
	"time"

type Transaction struct {
	Type    int16 	  		   `json:"type"`
	Amount  float64   		   `json:"amount"`
	OrderID  primitive.ObjectID `json:"orderID,omitempty"`
	Payment  string             `json:"payment"`
	Date    time.Time            `json:"date"`
}