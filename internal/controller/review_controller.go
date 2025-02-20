package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IReviewController interface {
	GetReviews(c *gin.Context)
	GetReviewByID(c *gin.Context)
	GetReviewsBySellerID(c *gin.Context)
	GetReviewsByBuyerID(c *gin.Context)
	CreateReview(c *gin.Context)
	UpdateReview(c *gin.Context)
}

type ReviewController struct {
	reviewService service.IReviewService
}

func NewReviewController(s service.IReviewService) IReviewController {
	return ReviewController{
		reviewService: s,
	}
}

// GetReviews godoc
// @Summary Get all reviews
// @Description Retrieves all reviews
// @Tags review
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.Review}
// @Failure 500 {object} dto.ErrorResponse
// @Router /review/ [get]
func (s ReviewController) GetReviews(c *gin.Context) {
	res, err := s.reviewService.GetReviews()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No reviews",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get Reviews success",
		Data:    res,
	})
}

// GetReviewsByID godoc
// @Summary Get a review by ID
// @Description Retrieves a review's data by its ID
// @Tags review
// @Accept json
// @Produce json
// @Param review_id path string true "Review ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.Review}
// @Failure 500 {object} dto.ErrorResponse
// @Router /review/{review_id} [get]
func (s ReviewController) GetReviewByID(c *gin.Context) {
	reviewIDstr := c.Param("review_id")
	reviewID, err := primitive.ObjectIDFromHex(reviewIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid reviewID format",
			Message: err.Error(),
		})
		return
	}
	res, err := s.reviewService.GetReviewByID(reviewID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No review with this reviewID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get review success",
		Data:    res,
	})
}

// GetReviewsBySellerID godoc
// @Summary Get reviews by sellerID
// @Description Retrieves each seller's reviews by seller ID
// @Tags review
// @Accept json
// @Produce json
// @Param seller_id path string true "Seller ID"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.Review}
// @Failure 500 {object} dto.ErrorResponse
// @Router /review/seller/{seller_id} [get]
func (s ReviewController) GetReviewsBySellerID(c *gin.Context) {
	sellerIDstr := c.Param("seller_id")
	sellerID, err := primitive.ObjectIDFromHex(sellerIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid sellerID format",
			Message: err.Error(),
		})
		return
	}
	res, err := s.reviewService.GetReviewsBySellerID(sellerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No review with this sellerID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get review success",
		Data:    res,
	})
}

// GetReviewsByBuyerID godoc
// @Summary Get reviews by buyerID
// @Description Retrieves each buyer's reviews by buyer ID
// @Tags review
// @Accept json
// @Produce json
// @Param buyer_id path string true "Buyer ID"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.Review}
// @Failure 500 {object} dto.ErrorResponse
// @Router /review/buyer/{buyer_id} [get]
func (s ReviewController) GetReviewsByBuyerID(c *gin.Context) {
	buyerIDstr := c.Param("buyer_id")
	buyerID, err := primitive.ObjectIDFromHex(buyerIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid buyerID format",
			Message: err.Error(),
		})
		return
	}
	res, err := s.reviewService.GetReviewsByBuyerID(buyerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No review with this buyerID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get review success",
		Data:    res,
	})
}

// CreateReview godoc

// CreateReview godoc
// @Summary Create a new review
// @Description Creates a new review in the database
// @Tags review
// @Accept json
// @Produce json
// @Param review body dto.ReviewCreateRequest true "Review to create"
// @Success 201 {object} dto.SuccessResponse{data=dto.Review}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /review/ [post]
func (s ReviewController) CreateReview(c *gin.Context) {
	var newReview model.Review

	if err := c.BindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	if err := newReview.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Validation failed",
			Message: err.Error(),
		})
		return
	}

	res, err := s.reviewService.CreateReview(&newReview)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to insert to database",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Review created",
		Data:    res,
	})
}

// UpdateReview godoc
// @Summary Update a review by ID
// @Description Updates an existing review's data by its ID
// @Tags review
// @Accept json
// @Produce json
// @Param review_id path string true "Review ID"
// @Param review body dto.ReviewUpdateRequest true "Review data to update"
// @Success 200 {object} dto.SuccessResponse{data=dto.Review}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /review/{review_id} [put]
func (s ReviewController) UpdateReview(c *gin.Context) {
	reviewIDstr := c.Param("review_id")
	reviewID, err := primitive.ObjectIDFromHex(reviewIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid reviewID format",
			Message: err.Error(),
		})
		return
	}
	var updatedReview model.Review
	if err := c.BindJSON(&updatedReview); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	if err := updatedReview.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Validation failed",
			Message: err.Error(),
		})
		return
	}

	res, err := s.reviewService.UpdateReview(reviewID, &updatedReview)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to update review data",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Update review success",
		Data:    res,
	})
}

