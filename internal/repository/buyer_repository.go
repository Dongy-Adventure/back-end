package repository

import (
	"context"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBuyerRepository interface {
	GetBuyerByID() ([]dto.Buyer, error)
	GetBuyer(string) (*dto.Buyer, error)
	CreateBuyerData(*model.Buyer) (*dto.Buyer, error)
	UpdateBuyerData(string, *model.Buyer) (*dto.Buyer, error)
}

type BuyerRepository struct {
	buyerCollection *mongo.Collection
}

func NewBuyerRepository(db *mongo.Database, collectionName string) IBuyerRepository {
	return BuyerRepository{
		buyerCollection: db.Collection(collectionName),
	}
}

func (r BuyerRepository) GetBuyerByID(buyerID string) (*dto.Buyer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var buyer *model.Buyer

	err := r.buyerCollection.FindOne(ctx, bson.M{"buyer_id": buyerID}).Decode(&buyer)
	if err != nil {
		return nil, err
	}
	return converter.BuyerModelToDTO(buyer)
}
func (r BuyerRepository) GetBuyer() ([]dto.Buyer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var buyerList []dto.Buyer

	dataList, err := r.buyerCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var buyerModel *model.Buyer
		if err = dataList.Decode(&buyerModel); err != nil {
			return nil, err
		}
		buyerDTO, buyerErr := converter.BuyerModelToDTO(buyerModel)
		if buyerErr != nil {
			return nil, err
		}
		buyerList = append(buyerList, *buyerDTO)
	}

	return buyerList, nil
}

func (r BuyerRepository) CreateBuyerData(buyer *model.Buyer) (*dto.Buyer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := r.buyerCollection.InsertOne(ctx, buyer)
	if err != nil {
		return nil, err
	}
	var newBuyer *model.Buyer
	err = r.buyerCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newBuyer)

	if err != nil {
		return nil, err
	}

	return converter.BuyerModelToDTO(newBuyer)
}

func (r BuyerRepository) UpdateBuyerData(buyerID string, updatedBuyer *model.Buyer) (*dto.Buyer,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
		    "buyer_id": updatedBuyer.BuyerID,
		    "username": updatedBuyer.Username,
		    "password": updatedBuyer.Password,
		    "name":     updatedBuyer.Name,
		    "surname":  updatedBuyer.Surname,
		},
	 }

	filter := bson.M{"buyer_id": buyerID}
	_,error := r.buyerCollection.UpdateOne(ctx,filter,update)
	if err != nil {
		return nil, err
	}

	var newUpdatedBuyer *model.Buyer
	err = r.buyerCollection.FindOne(ctx, filter).Decode(&newUpdatedBuyer)

	if err != nil {
		return nil, err
	}

	return converter.BuyerModelToDTO(newUpdatedBuyer)
}


