package service

import (
	"log"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type IPaymentService interface {
	HandlePayment(paymentRequest *dto.PaymentRequest) (*omise.Charge, error)
	AddClient(chargeID string, clientChan chan string)
	RemoveClient(chargeID string, client chan string)
	BroadcastChargeStatus(chargeID, status string)
	UpdatePaymentStatus(chargeID, status string) error
}
type PaymentService struct {
	client     *omise.Client
	sseClients map[string][]chan string
}

func NewPaymentService(client *omise.Client) IPaymentService {
	return PaymentService{client: client, sseClients: make(map[string][]chan string)}
}

func (s PaymentService) HandlePayment(paymentRequest *dto.PaymentRequest) (*omise.Charge, error) {
	token := paymentRequest.Token
	charge := &omise.Charge{}
	if err := s.client.Do(charge, &operations.CreateCharge{
		Amount:   paymentRequest.Amount,
		Currency: "thb",
		Card:     token,
	}); err != nil {
		return nil, err
	}

	return charge, nil
}

func (s PaymentService) BroadcastChargeStatus(chargeID, status string) {
	if clients, ok := s.sseClients[chargeID]; ok {
		for _, client := range clients {
			select {
			case client <- status:
				log.Printf("Sent SSE update for charge %s: %s", chargeID, status)
			default:
				log.Printf("Client for charge %s not ready, closing channel", chargeID)
				close(client)
				s.RemoveClient(chargeID, client)
			}
		}
	}
}

func (s PaymentService) AddClient(chargeID string, clientChan chan string) {
	s.sseClients[chargeID] = append(s.sseClients[chargeID], clientChan)
}
func (s PaymentService) RemoveClient(chargeID string, client chan string) {
	if clients, ok := s.sseClients[chargeID]; ok {
		for i, c := range clients {
			if c == client {
				s.sseClients[chargeID] = append(clients[:i], clients[i+1:]...)
				break
			}
		}
		if len(s.sseClients[chargeID]) == 0 {
			delete(s.sseClients, chargeID)
		}
	}
}

func (s PaymentService) UpdatePaymentStatus(chargeID, status string) error {
	log.Printf("Simulating database update: Charge %s status updated to %s", chargeID, status)
	// ... (database logic) ...
	return nil
}
