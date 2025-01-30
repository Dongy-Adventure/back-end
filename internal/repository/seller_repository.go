package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ISellerRepository interface {
	GetSellers() ([]dto.Seller, error)
	GetSellerByID(sellerID primitive.ObjectID) (*dto.Seller, error)
	CreateSellerData(seller *model.Seller) (*dto.Seller, error)
	GetSellerByUsername(req *dto.LoginRequest) (*model.Seller, error)
	UpdateSeller(sellerID primitive.ObjectID, updatedSeller *model.Seller) (*dto.Seller, error)
}

type SellerRepository struct {
	sellerCollection *mongo.Collection
}

func NewSellerRepository(db *mongo.Database, collectionName string) ISellerRepository {
	return SellerRepository{
		sellerCollection: db.Collection(collectionName),
	}
}

func (r SellerRepository) GetSellerByID(sellerID primitive.ObjectID) (*dto.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var seller *model.Seller

	err := r.sellerCollection.FindOne(ctx, bson.M{"_id": sellerID}).Decode(&seller)
	if err != nil {
		return nil, err
	}
	return converter.SellerModelToDTO(seller)
}

func (r SellerRepository) GetSellers() ([]dto.Seller, error) {
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
	seller.SellerID = primitive.NewObjectID()
	result, err := r.sellerCollection.InsertOne(ctx, seller)
	if err != nil {
		return nil, err
	}
	var newSeller *model.Seller
	err = r.sellerCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newSeller)

	if err != nil {
		return nil, err
	}

	return converter.SellerModelToDTO(newSeller)
}

func (r SellerRepository) GetSellerByUsername(req *dto.LoginRequest) (*model.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var seller *model.Seller

	err := r.sellerCollection.FindOne(ctx, bson.M{"username": req.Username}).Decode(&seller)
	if err != nil {
		return nil, err
	}
	return seller, nil
}

func (r SellerRepository) UpdateSeller(sellerID primitive.ObjectID, updatedSeller *model.Seller) (*dto.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Println(updatedSeller.Username, "HERE")
	update := bson.M{
		"$set": bson.M{
			"username": updatedSeller.Username,
			"password": updatedSeller.Password,
			"name":     updatedSeller.Name,
			"surname":  updatedSeller.Surname,
			"payment":  updatedSeller.Payment,
		},
	}

	filter := bson.M{"_id": sellerID}
	_, err := r.sellerCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var newUpdatedSeller *model.Seller
	err = r.sellerCollection.FindOne(ctx, filter).Decode(&newUpdatedSeller)
	if err != nil {
		return nil, err
	}

	return converter.SellerModelToDTO(newUpdatedSeller)
}
