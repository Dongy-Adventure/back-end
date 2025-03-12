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

type IBuyerRepository interface {
	GetBuyer() ([]dto.Buyer, error)
	GetBuyerByID(buyerID primitive.ObjectID) (*dto.Buyer, error)
	CreateBuyerData(buyer *model.Buyer) (*dto.Buyer, error)
	GetBuyerByUsername(req *dto.LoginRequest) (*model.Buyer, error)
	UpdateBuyerData(buyerID primitive.ObjectID, updatedBuyer *model.Buyer) (*dto.Buyer, error)
	UpdateProductInCart(buyerID primitive.ObjectID, product dto.Product) ([]dto.Product, error)
}

type BuyerRepository struct {
	buyerCollection *mongo.Collection
}

func NewBuyerRepository(db *mongo.Database, collectionName string) IBuyerRepository {
	return &BuyerRepository{
		buyerCollection: db.Collection(collectionName),
	}
}

func (r *BuyerRepository) GetBuyerByID(buyerID primitive.ObjectID) (*dto.Buyer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var buyer *model.Buyer

	err := r.buyerCollection.FindOne(ctx, bson.M{"_id": buyerID}).Decode(&buyer)
	if err != nil {
		return nil, err
	}
	return converter.BuyerModelToDTO(buyer)
}

func (r *BuyerRepository) GetBuyer() ([]dto.Buyer, error) {
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

func (r *BuyerRepository) GetBuyerByUsername(req *dto.LoginRequest) (*model.Buyer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var buyer *model.Buyer

	err := r.buyerCollection.FindOne(ctx, bson.M{"username": req.Username}).Decode(&buyer)
	if err != nil {
		return nil, err
	}
	return buyer, nil
}

func (r *BuyerRepository) CreateBuyerData(buyer *model.Buyer) (*dto.Buyer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	buyer.BuyerID = primitive.NewObjectID()
	buyer.Cart = []dto.Product{}
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

func (r *BuyerRepository) UpdateBuyerData(buyerID primitive.ObjectID, updatedBuyer *model.Buyer) (*dto.Buyer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	data, err := bson.Marshal(updatedBuyer)
	if err != nil {
		return nil, err
	}
	var update bson.M
	err = bson.Unmarshal(data, &update)
	if err != nil {
		return nil, err
	}
	for key, value := range update {
		if value == "" || value == nil || key == "_id" {
			delete(update, key)
		}
	}

	filter := bson.M{"_id": buyerID}
	_, err = r.buyerCollection.UpdateOne(ctx, filter, bson.M{"$set": update})
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

func (r *BuyerRepository) UpdateProductInCart(buyerID primitive.ObjectID, product dto.Product) ([]dto.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var buyer struct {
		Cart []dto.Product `bson:"cart"`
	}
	err := r.buyerCollection.FindOne(ctx, bson.M{"_id": buyerID}).Decode(&buyer)
	if err != nil {
		return nil, err
	}

	// Check if product is in cart
	newCart := buyer.Cart
	found := false
	for _, pr := range newCart {
		if pr.ProductID == product.ProductID {
			found = true
		}
	}

	// If not found, add the product
	if !found {
		newCart = append(newCart, product)
	}

	_, err = r.buyerCollection.UpdateOne(
		ctx,
		bson.M{"_id": buyerID},
		bson.M{"$set": bson.M{"cart": newCart}},
	)
	if err != nil {
		return nil, err
	}

	return newCart, nil
}
