package service

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type IPaymentService interface {
	HandlePayment(paymentRequest *dto.BuyerPaymentRequest) error
}
type PaymentService struct {
	client *omise.Client
}

func NewPaymentService(client *omise.Client) IPaymentService {
	return PaymentService{client: client}
}
func (s PaymentService) HandlePayment(paymentRequest *dto.BuyerPaymentRequest) error {
	// TODO: get token from client side
	mockToken := "tokn_test_4yupqgt7k9t9rdhn2m9" // This is a mock Omise token, replace it as needed for your tests.

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   int64(paymentRequest.Amount),
		Currency: "thb",
		Card:     mockToken,
	}
	if e := s.client.Do(charge, createCharge); e != nil {
		return e
	}
	return nil
}
