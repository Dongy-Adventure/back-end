package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/paymenttype"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/pkg/utils/converter"
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
	UpdateSellerScore(sellerID primitive.ObjectID) error
	GetSellerBalanceByID(sellerID primitive.ObjectID) (float64, error)
	DepositSellerBalance(sellerID primitive.ObjectID, orderID primitive.ObjectID, payment string, amount float64) error
	WithdrawSellerBalance(sellerID primitive.ObjectID, payment string, amount float64) error
}

type SellerRepository struct {
	sellerCollection *mongo.Collection
	reviewCollection *mongo.Collection
}

func NewSellerRepository(db *mongo.Database, sellercollectionName string, reviewcollectionName string) ISellerRepository {
	return SellerRepository{
		sellerCollection: db.Collection(sellercollectionName),
		reviewCollection: db.Collection(reviewcollectionName),
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

	// var existingSeller model.Seller
	// err := r.sellerCollection.FindOne(ctx, bson.M{"username": seller.Username}).Decode(&existingSeller)

	// if err == nil {
	// 	return nil, fmt.Errorf("this username is already exists")
	// }
	seller.SellerID = primitive.NewObjectID()
	seller.Transaction = []model.Transaction{}
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
	data, err := bson.Marshal(updatedSeller)
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

	filter := bson.M{"_id": sellerID}
	_, err = r.sellerCollection.UpdateOne(ctx, filter, bson.M{"$set": update})

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

func (r SellerRepository) UpdateSellerScore(sellerID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	pipeline := []bson.M{
		{"$match": bson.M{"seller_id": sellerID}},
		{"$group": bson.M{
			"_id":      "$seller_id",
			"avgScore": bson.M{"$avg": "$score"},
		}},
	}

	var result struct {
		AvgScore float64 `bson:"avgScore"`
	}

	cursor, err := r.reviewCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return err
		}
	} else {
		result.AvgScore = 0
	}

	_, err = r.sellerCollection.UpdateOne(
		ctx,
		bson.M{"_id": sellerID},
		bson.M{"$set": bson.M{"score": result.AvgScore}},
	)

	if err != nil {
		return err
	}

	return nil
}

func (r SellerRepository) GetSellerBalanceByID(sellerID primitive.ObjectID) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var seller model.Seller

	err := r.sellerCollection.FindOne(ctx, bson.M{"_id": sellerID}).Decode(&seller)
	if err != nil {
		return 0, err
	}

	totalBalance := seller.Balance

	return totalBalance, nil
}

func (r SellerRepository) DepositSellerBalance(sellerID primitive.ObjectID, orderID primitive.ObjectID, payment string, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	transaction := model.Transaction{
		Type:    paymenttype.CREDIT,
		Amount:  amount,
		OrderID: orderID,
		Payment: payment,
		Date:    time.Now(),
	}

	// Update Seller Balance & Add Transaction
	update := bson.M{
		"$inc":  bson.M{"balance": amount},          // Increase balance
		"$push": bson.M{"transaction": transaction}, // Add transaction record
	}

	_, err := r.sellerCollection.UpdateOne(ctx, bson.M{"_id": sellerID}, update)
	return err
}

func (r SellerRepository) WithdrawSellerBalance(sellerID primitive.ObjectID, payment string, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var seller model.Seller
	err := r.sellerCollection.FindOne(ctx, bson.M{"_id": sellerID}).Decode(&seller)
	if err != nil {
		return err
	}

	if seller.Balance < amount {
		return errors.New("insufficient balance")
	}

	transaction := model.Transaction{
		Type:    paymenttype.DEBIT,
		Amount:  -amount,
		Payment: payment,
		Date:    time.Now(),
	}

	// Update balance & add transaction
	update := bson.M{
		"$inc":  bson.M{"balance": -amount},         // Decrease balance
		"$push": bson.M{"transaction": transaction}, // Log transaction
	}

	_, err = r.sellerCollection.UpdateOne(ctx, bson.M{"_id": sellerID}, update)
	return err
}
