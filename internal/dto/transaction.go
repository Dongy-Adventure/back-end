package dto

import (
	"time"
)

type Transaction struct {
	Amount        float64   `json:"amount" binding:"min=0"`
	Product       []Product `json:"product"`
	Date          time.Time `json:"data"`
	PaymentMethod string    `json:"paymentMethod"`
}
