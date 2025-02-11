package service

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type ISellerService interface {
	CreateSellerData(seller *model.Seller) (*dto.Seller, error)
	GetSellerByID(sellerID primitive.ObjectID) (*dto.Seller, error)
	GetSellers() ([]dto.Seller, error)
	UpdateSeller(sellerID primitive.ObjectID, updatedSeller *model.Seller) (*dto.Seller, error)
	AddTransaction(sellerID primitive.ObjectID, transaction *dto.Transaction) (*dto.Transaction, error)
	GetSellerBalanceByID(sellerID primitive.ObjectID) (float64, error)
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

func (s SellerService) GetSellerByID(sellerID primitive.ObjectID) (*dto.Seller, error) {
	sellerDTO, err := s.sellerRepository.GetSellerByID(sellerID)
	if err != nil {
		return nil, err
	}
	return sellerDTO, nil
}

func (s SellerService) GetSellers() ([]dto.Seller, error) {
	sellers, err := s.sellerRepository.GetSellers()
	if err != nil {
		return nil, err
	}
	return sellers, nil
}

func (s SellerService) UpdateSeller(sellerID primitive.ObjectID, updatedSeller *model.Seller) (*dto.Seller, error) {

	if updatedSeller.Password != "" {
		passwordBytes := []byte(updatedSeller.Password)
		hashPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		encryptedPassword := string(hashPasswordBytes)
		updatedSeller.Password = encryptedPassword
	}

	updatedSellerDTO, err := s.sellerRepository.UpdateSeller(sellerID, updatedSeller)
	if err != nil {
		return nil, err
	}

	return updatedSellerDTO, nil
}
func (s SellerService) AddTransaction(sellerID primitive.ObjectID, transaction *dto.Transaction) (*dto.Transaction, error) {
	newTransaction, err := s.sellerRepository.AddTransaction(sellerID, transaction)
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}

func (s SellerService) GetSellerBalanceByID(sellerID primitive.ObjectID) (float64, error) {
	totalBalance, err := s.sellerRepository.GetSellerBalanceByID(sellerID)
	if err != nil {
		return 0, err
	}
	return totalBalance, nil
}