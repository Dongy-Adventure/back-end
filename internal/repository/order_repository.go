package repository

import (
	"context"
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
