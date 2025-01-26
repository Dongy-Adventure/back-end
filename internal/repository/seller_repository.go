package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ISellerRepository interface {
	GetSellerData() ([]dto.Seller, error)
	GetOneSellerData(string) (*dto.Seller, error)
	CreateSellerData(*model.Seller) (*dto.Seller, error)
}

type SellerRepository struct {
	sellerCollection *mongo.Collection
}

func NewSellerRepository(db *mongo.Database, collectionName string) ISellerRepository {
	return SellerRepository{
		sellerCollection: db.Collection(collectionName),
	}
}

func (r SellerRepository) GetOneSellerData(sellerID string) (*dto.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var seller *model.Seller

	err := r.sellerCollection.FindOne(ctx, bson.M{"seller_id": sellerID}).Decode(&seller)
	if err != nil {
		return nil, err
	}
	return converter.SellerModelToDTO(seller)
}
func (r SellerRepository) GetSellerData() ([]dto.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var sellerList []dto.Seller

	dataList, err := r.sellerCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var sellerModel *model.Seller
		if err = dataList.Decode(&sellerModel); err != nil {
			return nil, err
		}
		sellerDTO, sellerErr := converter.SellerModelToDTO(sellerModel)
		if sellerErr != nil {
			return nil, err
		}
		sellerList = append(sellerList, *sellerDTO)
	}

	return sellerList, nil
}

func (r SellerRepository) CreateSellerData(seller *model.Seller) (*dto.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := r.sellerCollection.InsertOne(ctx, seller)
	if err != nil {
		fmt.Println("HERE")
		return nil, err
	}
	var newSeller *model.Seller
	err = r.sellerCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newSeller)

	if err != nil {
		return nil, err
	}

	return converter.SellerModelToDTO(newSeller)
}
