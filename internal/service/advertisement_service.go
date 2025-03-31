package service

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAdvertisementService interface {
	GetAdvertisements() ([]dto.Advertisement, error)
	GetAdvertisementByID(advertisementID primitive.ObjectID) (*dto.Advertisement, error)
	GetAdvertisementsBySellerID(sellerID primitive.ObjectID) ([]dto.Advertisement, error)
	GetAdvertisementsByProductID(productID primitive.ObjectID) ([]dto.Advertisement, error)
	CreateAdvertisement(advertisement *model.Advertisement) (*dto.Advertisement, error)
	UpdateAdvertisement(advertisementID primitive.ObjectID, updatedAdvertisement *model.Advertisement) (*dto.Advertisement, error)
	DeleteAdvertisement(advertisementID primitive.ObjectID) error
}

type AdvertisementService struct {
	advertisementRepository repository.IAdvertisementRepository
}

func NewAdvertisementService(r repository.IAdvertisementRepository) IAdvertisementService {
	return AdvertisementService{
		advertisementRepository: r,
	}
}

func (s AdvertisementService) GetAdvertisements() ([]dto.Advertisement, error) {
	advertisements, err := s.advertisementRepository.GetAdvertisements()
	if err != nil {
		return nil, err
	}
	return advertisements, nil
}

func (s AdvertisementService) GetAdvertisementByID(advertisementID primitive.ObjectID) (*dto.Advertisement, error) {
	advertisementDTO, err := s.advertisementRepository.GetAdvertisementByID(advertisementID)
	if err != nil {
		return nil, err
	}
	return advertisementDTO, nil
}

func (s AdvertisementService) GetAdvertisementsBySellerID(sellerID primitive.ObjectID) ([]dto.Advertisement, error) {
	advertisements, err := s.advertisementRepository.GetAdvertisementsBySellerID(sellerID)
	if err != nil {
		return nil, err
	}
	return advertisements, nil
}

func (s AdvertisementService) GetAdvertisementsByProductID(productID primitive.ObjectID) ([]dto.Advertisement, error) {
	advertisements, err := s.advertisementRepository.GetAdvertisementsByProductID(productID)
	if err != nil {
		return nil, err
	}
	return advertisements, nil
}

func (s AdvertisementService) CreateAdvertisement(advertisement *model.Advertisement) (*dto.Advertisement, error) {

	newAdvertisement, err := s.advertisementRepository.CreateAdvertisement(advertisement)

	if err != nil {
		return nil, err
	}

	return newAdvertisement, nil
}


func (s AdvertisementService) UpdateAdvertisement(advertisementID primitive.ObjectID, updatedAdvertisement *model.Advertisement) (*dto.Advertisement, error) {

	updatedAdvertisementDTO, err := s.advertisementRepository.UpdateAdvertisement(advertisementID, updatedAdvertisement)
	if err != nil {
		return nil, err
	}

	return updatedAdvertisementDTO, nil
}

func (s AdvertisementService) DeleteAdvertisement(advertisementID primitive.ObjectID) error {

	err := s.advertisementRepository.DeleteAdvertisement(advertisementID)
	if err != nil {
		return err 
	}

	return nil 
}

