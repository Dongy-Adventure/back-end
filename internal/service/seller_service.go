package service

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type ISellerService interface {
	CreateSellerData(*model.Seller) (*dto.Seller, error)
	GetSellerByID(string) (*dto.Seller, error)
	UpdateSellerData(string, *model.Seller) (*dto.Seller, error)
}

type SellerService struct {
	sellerRepository repository.ISellerRepository
}

func NewSellerService(r repository.ISellerRepository) ISellerService {
	return SellerService{
		sellerRepository: r,
	}
}

func (s SellerService) CreateSellerData(seller *model.Seller) (*dto.Seller, error) {

	passwordBytes := []byte(seller.Password)
	hashPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	encryptPassword := string(hashPasswordBytes)

	seller.Password = encryptPassword

	newSeller, err := s.sellerRepository.CreateSellerData(seller)

	if err != nil {
		return nil, err
	}

	return newSeller, nil
}

func (s SellerService) GetSellerByID(sellerID string) (*dto.Seller, error) {
	selletDTO, err := s.sellerRepository.GetSellerByID(sellerID)
	if err != nil {
		return nil, err
	}
	return selletDTO, nil
}

func (s SellerService) UpdateSellerData(sellerID string, updatedSeller *model.Seller) (*dto.Seller, error) {
	
	if updatedSeller.Password != "" {
		passwordBytes := []byte(updatedSeller.Password)
		hashPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		encryptedPassword := string(hashPasswordBytes)
		updatedSeller.Password = encryptedPassword
	}

	updatedSellerDTO, err := s.sellerRepository.UpdateSellerData(sellerID, updatedSeller)
	if err != nil {
		return nil, err
	}

	return updatedSellerDTO, nil
}
