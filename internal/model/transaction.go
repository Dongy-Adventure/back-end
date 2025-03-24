package model

import (
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
)

type Transaction struct {
	Amount        float64       `json:"amount" bson:"amount" binding:"min=0"`
	Product       []dto.Product `json:"product" bson:"product"`
	Date          time.Time     `json:"data" bson:"date"`
	PaymentMethod string        `json:"paymentMethod" bson:"paymentMethod"`
}
