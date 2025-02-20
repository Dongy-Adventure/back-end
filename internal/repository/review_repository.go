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

type IReviewRepository interface {
	GetReviews() ([]dto.Review, error)
	GetReviewByID(reviewID primitive.ObjectID) (*dto.Review, error)
	GetReviewsBySellerID(sellerID primitive.ObjectID) ([]dto.Review, error)
	GetReviewsByBuyerID(buyerID primitive.ObjectID) ([]dto.Review, error)
	CreateReview(review *model.Review) (*dto.Review, error)
	UpdateReview(reviewID primitive.ObjectID, updatedReview *model.Review) (*dto.Review, error)
}

type ReviewRepository struct {
	reviewCollection *mongo.Collection
	sellerRepo 	  ISellerRepository
}

func NewReviewRepository(db *mongo.Database, reviewcollectionName string, sellerRepo ISellerRepository) IReviewRepository {
	return ReviewRepository{
		reviewCollection: db.Collection(reviewcollectionName),
		sellerRepo:       sellerRepo,
	}
}

func (r ReviewRepository) GetReviews() ([]dto.Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var reviewList []dto.Review

	dataList, err := r.reviewCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var reviewModel *model.Review
		if err = dataList.Decode(&reviewModel); err != nil {
			return nil, err
		}
		reviewDTO, reviewErr := converter.ReviewModelToDTO(reviewModel)
		if reviewErr != nil {
			return nil, err
		}
		reviewList = append(reviewList, *reviewDTO)
	}

	return reviewList, nil
}

func (r ReviewRepository) GetReviewByID(reviewID primitive.ObjectID) (*dto.Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var review *model.Review

	err := r.reviewCollection.FindOne(ctx, bson.M{"_id": reviewID}).Decode(&review)
	if err != nil {
		return nil, err
	}
	return converter.ReviewModelToDTO(review)
}

func (r ReviewRepository) GetReviewsBySellerID(sellerID primitive.ObjectID) ([]dto.Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var reviewList []dto.Review

	dataList, err := r.reviewCollection.Find(ctx, bson.M{"seller_id": sellerID})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var reviewModel *model.Review
		if err = dataList.Decode(&reviewModel); err != nil {
			return nil, err
		}
		reviewDTO, reviewErr := converter.ReviewModelToDTO(reviewModel)
		if reviewErr != nil {
			return nil, err
		}
		reviewList = append(reviewList, *reviewDTO)
	}

	return reviewList, nil
}

func (r ReviewRepository) GetReviewsByBuyerID(buyerID primitive.ObjectID) ([]dto.Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var reviewList []dto.Review

	dataList, err := r.reviewCollection.Find(ctx, bson.M{"buyer_id": buyerID})
	if err != nil {
		return nil, err
	}
	defer dataList.Close(ctx)
	for dataList.Next(ctx) {
		var reviewModel *model.Review
		if err = dataList.Decode(&reviewModel); err != nil {
			return nil, err
		}
		reviewDTO, reviewErr := converter.ReviewModelToDTO(reviewModel)
		if reviewErr != nil {
			return nil, err
		}
		reviewList = append(reviewList, *reviewDTO)
	}

	return reviewList, nil
}



func (r ReviewRepository) CreateReview(review *model.Review) (*dto.Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	review.ReviewID = primitive.NewObjectID()
	review.Date = time.Now()
	result, err := r.reviewCollection.InsertOne(ctx, review)
	if err != nil {
		return nil, err
	}
	var newReview *model.Review
	err = r.reviewCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newReview)
	if err != nil {
		return nil, err
	}

	//update seller's score
	err = r.sellerRepo.UpdateSellerScore(newReview.SellerID)
	if err != nil {
		return nil, err
	}

	return converter.ReviewModelToDTO(newReview)
}


func (r ReviewRepository) UpdateReview(reviewID primitive.ObjectID, updatedReview *model.Review) (*dto.Review, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"message": updatedReview.Message,
        		"score":   updatedReview.Score,
        		"image":   updatedReview.Image,
        		"date":    time.Now(),
		},
	}

	filter := bson.M{"_id": reviewID}
	_, err := r.reviewCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var newUpdatedReview *model.Review
	err = r.reviewCollection.FindOne(ctx, filter).Decode(&newUpdatedReview)
	if err != nil {
		return nil, err
	}

	//update seller's score
	err = r.sellerRepo.UpdateSellerScore(newUpdatedReview.SellerID)
	if err != nil {
		return nil, err
	}

	return converter.ReviewModelToDTO(newUpdatedReview)
}

