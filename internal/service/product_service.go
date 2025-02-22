package service

import (
	"context"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductService interface {
	GetProductByID(productID primitive.ObjectID) (*dto.Product, error)
	GetProducts() ([]dto.Product, error)
	CreateProduct(product *model.Product) (*dto.Product, error)
	UpdateProduct(productID primitive.ObjectID, updatedProduct *model.Product) (*dto.Product, error)
	GetProductBySellerID(sellerID primitive.ObjectID) (*dto.Product, error)
}

type ProductService struct {
	productRepository repository.IProductRepository
}

func NewProductService(r repository.IProductRepository) IProductService {
	return ProductService{
		productRepository: r,
	}
}

func (s ProductService) CreateProduct(product *model.Product) (*dto.Product, error) {
	// You may not need to hash passwords for products, so you can remove that part.
	newProduct, err := s.productRepository.CreateProduct(context.Background(), product)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (s ProductService) GetProductByID(productID primitive.ObjectID) (*dto.Product, error) {
	productDTO, err := s.productRepository.GetProductByID(context.Background(), productID)
	if err != nil {
		return nil, err
	}
	return productDTO, nil
}

func (s ProductService) GetProductBySellerID(sellerID primitive.ObjectID) (*dto.Product, error) {
	productDTO, err := s.productRepository.GetProductByID(context.Background(), sellerID)
	if err != nil {
		return nil, err
	}
	return productDTO, nil
}

func (s ProductService) GetProducts() ([]dto.Product, error) {
	products, err := s.productRepository.GetProducts(context.Background())
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s ProductService) UpdateProduct(productID primitive.ObjectID, updatedProduct *model.Product) (*dto.Product, error) {
	updatedProductDTO, err := s.productRepository.UpdateProduct(context.Background(), productID, updatedProduct)
	if err != nil {
		return nil, err
	}
	return updatedProductDTO, nil
}
