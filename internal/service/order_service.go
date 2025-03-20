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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderService interface {
	CreateOrder(products []dto.OrderProduct, buyerID primitive.ObjectID, sellerID primitive.ObjectID, sellerName string, buyerName string,payment string) (*dto.Order, error)
	GetOrdersByUserID(userID primitive.ObjectID, userType userrole.UserType) ([]dto.Order, error)
	GetTotalPrice(products []dto.OrderProduct) (float64,error)
	DeleteOrderByOrderID(orderID primitive.ObjectID) error
	UpdateOrder(orderID primitive.ObjectID, updatedOrder *model.Order) (*dto.Order, error)
	UpdateOrderStatus(orderID primitive.ObjectID, orderStatus int) (int, error)
}

type OrderService struct {
	orderRepository       repository.IOrderRepository
	appointmentRepository repository.IAppointmentRepository
	sellerRepository      repository.ISellerRepository
	productRepository 	  repository.IProductRepository
}

func NewOrderService(r repository.IOrderRepository, a repository.IAppointmentRepository, sr repository.ISellerRepository, p repository.IProductRepository) IOrderService {
	return OrderService{orderRepository: r, appointmentRepository: a, sellerRepository: sr, productRepository: p}
}

func (s OrderService) CreateOrder(products []dto.OrderProduct, buyerID primitive.ObjectID, sellerID primitive.ObjectID, sellerName string, buyerName string, payment string) (*dto.Order, error) {
	if len(products) == 0 {
		return nil, errors.New("no product")
	}

	var productsModel []model.OrderProduct
	for _, product := range products {
		stockProduct, err := s.productRepository.GetProductByID(product.ProductID)
			if err != nil {
				return nil, err
			}

		if stockProduct.Amount < product.Amount {
			return nil, fmt.Errorf("not enough stock for product %s", stockProduct.ProductName)
		}

		productsModel = append(productsModel, model.OrderProduct{
			ProductID: product.ProductID,
			Amount:    product.Amount,
		})
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
	

	// Get total price 
	totalPrice,err := s.GetTotalPrice(products)
	if err != nil {
		return nil, err
	}

	// Add transaction and update (+deposit) seller balance
	err = s.sellerRepository.DepositSellerBalance(sellerID, orderID, payment, totalPrice)
	if err != nil {
		return nil, err
	}

	// Deduct product amount
	for _, product := range products {
		err = s.productRepository.UpdateProductAmount(product.ProductID, product.Amount)
		if err != nil {
			return nil, err
		}
	}


	return s.orderRepository.CreateOrder(&model.Order{
		OrderID:       orderID,
		Status:        orderstatus.WAITFORLOCATION,
		Products:      productsModel,
		AppointmentID: app.AppointmentID,
		BuyerID:       buyerID,
		BuyerName:     buyerName,
		SellerID:      sellerID,
		SellerName:    sellerName,
		TotalPrice:    totalPrice,
		Payment:       payment,
		CreatedAt:     time.Now(),
	})
}

func (s OrderService) GetTotalPrice(products []dto.OrderProduct) (float64,error) {
	var totalPrice float64
	for _, product := range products {
		prod, err := s.productRepository.GetProductByID(product.ProductID)
		if err != nil {
			return 0, err
		}
		totalPrice += prod.Price * float64(product.Amount)
	}
	return totalPrice, nil
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
