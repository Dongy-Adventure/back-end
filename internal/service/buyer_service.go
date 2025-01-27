package service

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type IBuyerService interface {
	CreateBuyerData(*model.Buyer) (*dto.Buyer, error)
}

type BuyerService struct {
	buyerRepository repository.IBuyerRepository
}

func NewBuyerService(r repository.IBuyerRepository) IBuyerService {
	return BuyerService{
		buyerRepository: r,
	}
}

func (s BuyerService) CreateBuyerData(buyer *model.Buyer) (*dto.Buyer, error) {

	passwordBytes := []byte(buyer.Password)
	hashPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	encryptPassword := string(hashPasswordBytes)

	buyer.Password = encryptPassword

	newBuyer, err := s.buyerRepository.CreateBuyerData(buyer)

	if err != nil {
		return nil, err
	}

	return newBuyer, nil
}
