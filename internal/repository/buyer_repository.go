package repository

import (
	"context"
	"time"
	"fmt"

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
	UpdateProductInCart(buyerID primitive.ObjectID, product *model.OrderProduct) ([]dto.OrderProduct, error)
	DeleteProductFromCart(buyerID, productID primitive.ObjectID) error
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
	buyer.Cart = []model.OrderProduct{}
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

func (r *BuyerRepository) UpdateProductInCart(buyerID primitive.ObjectID, product *model.OrderProduct) ([]dto.OrderProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var buyer struct {
		Cart []model.OrderProduct `bson:"cart"` 
	}
	err := r.buyerCollection.FindOne(ctx, bson.M{"_id": buyerID}).Decode(&buyer)
	if err != nil {
		return nil, err
	}

	// Check if product is in cart
	found := false
	for i, p := range buyer.Cart {
		if p.ProductID == product.ProductID {
			buyer.Cart[i].Amount = product.Amount
			found = true
			break
		}
	}

	if !found {
		buyer.Cart = append(buyer.Cart, *product) 
	}

	_, err = r.buyerCollection.UpdateOne(
		ctx,
		bson.M{"_id": buyerID},
		bson.M{"$set": bson.M{"cart": buyer.Cart}},
	)
	if err != nil {
		return nil, err
	}

	var updatedCart []dto.OrderProduct
	for _, p := range buyer.Cart {
		productDTO, err := converter.OrderProductModelToDTO(&p)
		if err != nil {
			return nil, err
		}
		updatedCart = append(updatedCart, *productDTO)
	}

	return updatedCart, nil
}


func (r *BuyerRepository) DeleteProductFromCart(buyerID, productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": buyerID}
	update := bson.M{"$pull": bson.M{"cart": bson.M{"productID": productID}}}

	result, err := r.buyerCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("product not found in cart or already removed")
	}

	return nil
}

