package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/orderstatus"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/userrole"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/pkg/utils/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderService interface {
	CreateOrder(products []dto.Product, buyerID primitive.ObjectID, sellerID primitive.ObjectID) (*dto.Order, error)
	GetOrdersByUserID(userID primitive.ObjectID, userType userrole.UserType) ([]dto.Order, error)
	GetTotalPrice(products []dto.Product) float64
	DeleteOrderByOrderID(orderID primitive.ObjectID) error
	UpdateOrder(orderID primitive.ObjectID, updatedOrder *model.Order) (*dto.Order, error)
	UpdateOrderStatus(orderID primitive.ObjectID, orderStatus int) (int, error)
}

type OrderService struct {
	orderRepository       repository.IOrderRepository
	appointmentRepository repository.IAppointmentRepository
	sellerRepository      repository.ISellerRepository
}

func NewOrderService(orderRepository repository.IOrderRepository, appointmentRepository repository.IAppointmentRepository, sellerRepository repository.ISellerRepository) IOrderService {
	return OrderService{orderRepository: orderRepository, appointmentRepository: appointmentRepository, sellerRepository: sellerRepository}
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
	orderID := primitive.NewObjectID()
	app, err := s.appointmentRepository.CreateAppointment(&model.Appointment{
		AppointmentID: primitive.NewObjectID(),
		OrderID:       orderID,
		BuyerID:       buyerID,
		SellerID:      sellerID,
		CreatedAt:     time.Now(),
	})

	if err != nil {
		return nil, err
	}
	totalPrice := s.GetTotalPrice(products)
	newOrder, err := s.orderRepository.CreateOrder(&model.Order{
		OrderID:       orderID,
		Status:        orderstatus.WAITFORLOCATION,
		Products:      productsModel,
		AppointmentID: app.AppointmentID,
		BuyerID:       buyerID,
		SellerID:      sellerID,
		TotalPrice:    totalPrice,
		CreatedAt:     time.Now(),
	})

	if err != nil {
		return nil, err
	}

	_, err = s.sellerRepository.UpdateSellerBalance(sellerID, totalPrice)

	if err != nil {
		return nil, err
	}

	_, err = s.sellerRepository.AddTransaction(sellerID, &dto.Transaction{
		Amount:        totalPrice,
		Product:       products,
		Date:          time.Now(),
		PaymentMethod: "PromptPay",
	})
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

func (s OrderService) GetTotalPrice(products []dto.Product) float64 {
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

func (s OrderService) DeleteOrderByOrderID(orderID primitive.ObjectID) error {
	err := s.orderRepository.DeleteOrderByOrderID(orderID)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	return nil
}

func (s OrderService) UpdateOrder(orderID primitive.ObjectID, updatedOrder *model.Order) (*dto.Order, error) {
	updatedOrderFromDB, err := s.orderRepository.UpdateOrder(orderID, updatedOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	return updatedOrderFromDB, nil
}

func (s OrderService) UpdateOrderStatus(orderID primitive.ObjectID, orderStatus int) (int, error) {
	updatedStatus, err := s.orderRepository.UpdateOrderStatus(orderID, orderStatus)
	if err != nil {
		return 0, fmt.Errorf("failed to update order status: %w", err)
	}

	return updatedStatus, nil
}
