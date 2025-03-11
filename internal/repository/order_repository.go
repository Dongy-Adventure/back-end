package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/userrole"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IOrderRepository interface {
	CreateOrder(order *model.Order) (*dto.Order, error)
	GetOrdersByUserID(userID primitive.ObjectID, userType userrole.UserType) ([]dto.Order, error)
	DeleteOrderByOrderID(orderID primitive.ObjectID) error
	UpdateOrder(orderID primitive.ObjectID, updatedOrder *model.Order) (*dto.Order, error)
	UpdateOrderStatus(orderID primitive.ObjectID, orderStatus int) (int, error)
}

type OrderRepository struct {
	orderCollection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database, collectionName string) IOrderRepository {
	return OrderRepository{
		orderCollection: db.Collection(collectionName),
	}
}
func (r OrderRepository) CreateOrder(order *model.Order) (*dto.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := r.orderCollection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	var newOrder *model.Order
	err = r.orderCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newOrder)
	if err != nil {
		return nil, err
	}
	return converter.OrderModelToDTO(newOrder)
}

func (r OrderRepository) GetOrdersByUserID(userID primitive.ObjectID, userType userrole.UserType) ([]dto.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	var orders []dto.Order
	var filter bson.M

	defer cancel()
	switch userType {
	case userrole.UserRole.BUYER:
		filter = bson.M{"buyerID": userID}
	case userrole.UserRole.SELLER:
		filter = bson.M{"sellerID": userID}
	default:
		return nil, fmt.Errorf("invalid user type")

	}
	dataList, err := r.orderCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var nextOrder *model.Order
		if err = dataList.Decode(&nextOrder); err != nil {
			return nil, err
		}
		order, err := converter.OrderModelToDTO(nextOrder)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}
	return orders, nil
}

func (r OrderRepository) DeleteOrderByOrderID(orderID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	filter := bson.M{"_id": orderID}
	deleteResult, err := r.orderCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("no order found")
	}

	return nil
}

func (r OrderRepository) UpdateOrder(orderID primitive.ObjectID, updatedOrder *model.Order) (*dto.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	filter := bson.M{"_id": orderID}
	update := bson.M{
		"$set": updatedOrder,
	}
	result := r.orderCollection.FindOneAndUpdate(ctx, filter, update)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no order found with the given ID")
		}
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	var updatedOrderFromDB model.Order
	err := r.orderCollection.FindOne(ctx, filter).Decode(&updatedOrderFromDB)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated order: %w", err)
	}

	return converter.OrderModelToDTO(&updatedOrderFromDB)
}

func (r OrderRepository) UpdateOrderStatus(orderID primitive.ObjectID, orderStatus int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"_id": orderID}

	update := bson.M{
		"$set": bson.M{
			"status": orderStatus,
		},
	}

	result := r.orderCollection.FindOneAndUpdate(ctx, filter, update)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, fmt.Errorf("no order found with the given ID")
		}
		return 0, fmt.Errorf("failed to update order status: %w", err)
	}

	return orderStatus, nil
}
