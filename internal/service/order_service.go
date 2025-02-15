package service

import (
	"errors"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/orderstatus"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/userrole"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderService interface {
	CreateOrder(products []dto.Product, buyerID primitive.ObjectID, sellerID primitive.ObjectID) (*dto.Order, error)
	GetOrdersByUserID(userID primitive.ObjectID, userType userrole.UserType) ([]dto.Order, error)
	getTotalPrice(products []dto.Product) float64
}

type OrderService struct {
	orderRepository repository.IOrderRepository
}

func NewOrderService(r repository.IOrderRepository) IOrderService {
	return OrderService{orderRepository: r}
}

func (s OrderService) CreateOrder(products []dto.Product, buyerID primitive.ObjectID, sellerID primitive.ObjectID) (*dto.Order, error) {
	var productsModel []model.Product
	if len(products) <= 0 {
		return nil, errors.New("no product")
	}
	for i := 0; i < len(products); i++ {
		product, err := converter.ProductDTOToModel(&products[i])
		if err != nil {
			return nil, err
		}
		productsModel = append(productsModel, *product)
	}

	return s.orderRepository.CreateOrder(&model.Order{
		OrderID:       primitive.NewObjectID(),
		Status:        orderstatus.PENDING,
		Products:      productsModel,
		AppointmentID: primitive.NilObjectID,
		BuyerID:       buyerID,
		SellerID:      sellerID,
		TotalPrice:    s.getTotalPrice(products),
		CreatedAt:     time.Now(),
	})
}

func (s OrderService) getTotalPrice(products []dto.Product) float64 {
	var totalPrice float64
	for i := 0; i < len(products); i++ {
		totalPrice += products[i].Price
	}
	return totalPrice
}

func (s OrderService) GetOrdersByUserID(userID primitive.ObjectID, userType userrole.UserType) ([]dto.Order, error) {
	orders, err := s.orderRepository.GetOrdersByUserID(userID, userType)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
