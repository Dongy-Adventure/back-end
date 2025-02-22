package repository

import (
	"context"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	GetProductByID(ctx context.Context, productID primitive.ObjectID) (*dto.Product, error)
	GetProducts(ctx context.Context) ([]dto.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) (*dto.Product, error)
	UpdateProduct(ctx context.Context, productID primitive.ObjectID, updatedProduct *model.Product) (*dto.Product, error)
	DeleteProduct(ctx context.Context, productID primitive.ObjectID) error
}

type ProductRepository struct {
	productCollection *mongo.Collection
}

func NewProductRepository(db *mongo.Database, collectionName string) IProductRepository {
	return &ProductRepository{
		productCollection: db.Collection(collectionName),
	}
}
func (r *ProductRepository) GetProductSellerByID(ctx context.Context, sellerID primitive.ObjectID) (*dto.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var product *model.Product

	err := r.productCollection.FindOne(ctx, bson.M{"sellerID": sellerID}).Decode(&sellerID)
	if err != nil {
		return nil, err
	}
	return converter.ProductModelToDTO(product)
}

func (r *ProductRepository) GetProductByID(ctx context.Context, productID primitive.ObjectID) (*dto.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var product *model.Product

	err := r.productCollection.FindOne(ctx, bson.M{"_id": productID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return converter.ProductModelToDTO(product)
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]dto.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var productList []dto.Product

	dataList, err := r.productCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)

	for dataList.Next(ctx) {
		var productModel *model.Product
		if err = dataList.Decode(&productModel); err != nil {
			return nil, err
		}
		productDTO, productErr := converter.ProductModelToDTO(productModel)
		if productErr != nil {
			return nil, err
		}
		productList = append(productList, *productDTO)
	}

	return productList, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *model.Product) (*dto.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	product.ProductID = primitive.NewObjectID()
	result, err := r.productCollection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	var newProduct *model.Product
	err = r.productCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newProduct)

	if err != nil {
		return nil, err
	}

	return converter.ProductModelToDTO(newProduct)
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, productID primitive.ObjectID, updatedProduct *model.Product) (*dto.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        updatedProduct.ProductName,
			"description": updatedProduct.Description,
			"price":       updatedProduct.Price,
			"tag":         updatedProduct.Tag,
			"imageURL":    updatedProduct.ImageURL,
			"color":       updatedProduct.Color,
		},
	}

	filter := bson.M{"_id": productID}
	_, err := r.productCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var newUpdatedProduct *model.Product
	err = r.productCollection.FindOne(ctx, filter).Decode(&newUpdatedProduct)

	if err != nil {
		return nil, err
	}

	return converter.ProductModelToDTO(newUpdatedProduct)
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := r.productCollection.DeleteOne(ctx, bson.M{"_id": productID})
	return err
}
