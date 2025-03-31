package repository

import (
	"context"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/pkg/utils/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAdvertisementRepository interface {
	GetAdvertisements() ([]dto.Advertisement, error)
	GetAdvertisementByID(advertisementID primitive.ObjectID) (*dto.Advertisement, error)
	GetAdvertisementsBySellerID(sellerID primitive.ObjectID) ([]dto.Advertisement, error)
	GetAdvertisementsByProductID(productID primitive.ObjectID) ([]dto.Advertisement, error)
	CreateAdvertisement(advertisement *model.Advertisement) (*dto.Advertisement, error)
	UpdateAdvertisement(advertisementID primitive.ObjectID, updatedAdvertisement *model.Advertisement) (*dto.Advertisement, error)
	DeleteAdvertisement(advertisementID primitive.ObjectID) error
}

type AdvertisementRepository struct {
	advertisementCollection *mongo.Collection
}

func NewAdvertisementRepository(db *mongo.Database, advertisementcollectionName string) IAdvertisementRepository {
	return AdvertisementRepository{
		advertisementCollection: db.Collection(advertisementcollectionName),
	}
}

func (r AdvertisementRepository) GetAdvertisements() ([]dto.Advertisement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var advertisementList []dto.Advertisement

	dataList, err := r.advertisementCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var advertisementModel *model.Advertisement
		if err = dataList.Decode(&advertisementModel); err != nil {
			return nil, err
		}
		advertisementDTO, advertisementErr := converter.AdvertisementModelToDTO(advertisementModel)
		if advertisementErr != nil {
			return nil, err
		}
		advertisementList = append(advertisementList, *advertisementDTO)
	}

	return advertisementList, nil
}

func (r AdvertisementRepository) GetAdvertisementByID(advertisementID primitive.ObjectID) (*dto.Advertisement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var advertisement *model.Advertisement

	err := r.advertisementCollection.FindOne(ctx, bson.M{"_id": advertisementID}).Decode(&advertisement)
	if err != nil {
		return nil, err
	}
	return converter.AdvertisementModelToDTO(advertisement)
}

func (r AdvertisementRepository) GetAdvertisementsBySellerID(sellerID primitive.ObjectID) ([]dto.Advertisement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var advertisementList []dto.Advertisement

	dataList, err := r.advertisementCollection.Find(ctx, bson.M{"seller_id": sellerID})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var advertisementModel *model.Advertisement
		if err = dataList.Decode(&advertisementModel); err != nil {
			return nil, err
		}
		advertisementDTO, advertisementErr := converter.AdvertisementModelToDTO(advertisementModel)
		if advertisementErr != nil {
			return nil, err
		}
		advertisementList = append(advertisementList, *advertisementDTO)
	}

	return advertisementList, nil
}

func (r AdvertisementRepository) GetAdvertisementsByProductID(productID primitive.ObjectID) ([]dto.Advertisement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var advertisementList []dto.Advertisement

	dataList, err := r.advertisementCollection.Find(ctx, bson.M{"product_id": productID})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var advertisementModel *model.Advertisement
		if err = dataList.Decode(&advertisementModel); err != nil {
			return nil, err
		}
		advertisementDTO, advertisementErr := converter.AdvertisementModelToDTO(advertisementModel)
		if advertisementErr != nil {
			return nil, err
		}
		advertisementList = append(advertisementList, *advertisementDTO)
	}

	return advertisementList, nil
}

func (r AdvertisementRepository) CreateAdvertisement(advertisement *model.Advertisement) (*dto.Advertisement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	advertisement.AdvertisementID = primitive.NewObjectID()
	advertisement.CreatedAt = time.Now()
	result, err := r.advertisementCollection.InsertOne(ctx, advertisement)
	if err != nil {
		return nil, err
	}
	var newAdvertisement *model.Advertisement
	err = r.advertisementCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newAdvertisement)
	if err != nil {
		return nil, err
	}
	return converter.AdvertisementModelToDTO(newAdvertisement)
}

func (r AdvertisementRepository) UpdateAdvertisement(advertisementID primitive.ObjectID, updatedAdvertisement *model.Advertisement) (*dto.Advertisement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"productID": updatedAdvertisement.ProductID,
			"imageURL":  updatedAdvertisement.ImageURL,
			"amount":    updatedAdvertisement.Amount,
			"payment":   updatedAdvertisement.Payment,
			"createdAt": time.Now(),
		},
	}

	filter := bson.M{"_id": advertisementID}
	_, err := r.advertisementCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var newUpdatedAdvertisement *model.Advertisement
	err = r.advertisementCollection.FindOne(ctx, filter).Decode(&newUpdatedAdvertisement)
	if err != nil {
		return nil, err
	}

	return converter.AdvertisementModelToDTO(newUpdatedAdvertisement)
}

func (r AdvertisementRepository) DeleteAdvertisement(advertisementID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := r.advertisementCollection.DeleteOne(ctx, bson.M{"_id": advertisementID})
	return err
}
