package service

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IReviewService interface {
	GetReviews() ([]dto.Review, error)
	GetReviewByID(reviewID primitive.ObjectID) (*dto.Review, error)
	GetReviewsBySellerID(sellerID primitive.ObjectID) ([]dto.Review, error)
	GetReviewsByBuyerID(buyerID primitive.ObjectID) ([]dto.Review, error)
	CreateReview(review *model.Review) (*dto.Review, error)
	UpdateReview(reviewID primitive.ObjectID, updatedReview *model.Review) (*dto.Review, error)
}

type ReviewService struct {
	reviewRepository repository.IReviewRepository
}

func NewReviewService(r repository.IReviewRepository) IReviewService {
	return ReviewService{
		reviewRepository: r,
	}
}

func (s ReviewService) GetReviews() ([]dto.Review, error) {
	reviews, err := s.reviewRepository.GetReviews()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s ReviewService) GetReviewByID(reviewID primitive.ObjectID) (*dto.Review, error) {
	reviewDTO, err := s.reviewRepository.GetReviewByID(reviewID)
	if err != nil {
		return nil, err
	}
	return reviewDTO, nil
}

func (s ReviewService) GetReviewsBySellerID(sellerID primitive.ObjectID) ([]dto.Review, error) {
	reviews, err := s.reviewRepository.GetReviewsBySellerID(sellerID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s ReviewService) GetReviewsByBuyerID(buyerID primitive.ObjectID) ([]dto.Review, error) {
	reviews, err := s.reviewRepository.GetReviewsByBuyerID(buyerID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s ReviewService) CreateReview(review *model.Review) (*dto.Review, error) {

	newReview, err := s.reviewRepository.CreateReview(review)

	if err != nil {
		return nil, err
	}

	return newReview, nil
}


func (s ReviewService) UpdateReview(reviewID primitive.ObjectID, updatedReview *model.Review) (*dto.Review, error) {

	updatedReviewDTO, err := s.reviewRepository.UpdateReview(reviewID, updatedReview)
	if err != nil {
		return nil, err
	}

	return updatedReviewDTO, nil
}

