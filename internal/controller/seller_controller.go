package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type ISellerController interface {
	CreateSeller(c *gin.Context)
	GetSellerByID(c *gin.Context)
	GetSellers(c *gin.Context)
	UpdateSeller(c *gin.Context)
}

type SellerController struct {
	sellerService service.ISellerService
}

func NewSellerController(s service.ISellerService) ISellerController {
	return SellerController{
		sellerService: s,
	}
}

// CreateSeller godoc
// @Summary Create a new seller
// @Description Creates a new seller in the database
// @Tags seller
// @Accept json
// @Produce json
// @Param seller body model.Seller true "Seller to create"
// @Success 201 {object} dto.SuccessResponse{data=dto.Seller}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /seller/ [post]
func (s SellerController) CreateSeller(c *gin.Context) {
	var newSeller model.Seller

	if err := c.BindJSON(&newSeller); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	res, err := s.sellerService.CreateSellerData(&newSeller)

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
		Message: "Seller created",
		Data:    res,
	})
}

// GetSellerByID godoc
// @Summary Get a seller by ID
// @Description Retrieves a seller's data by their ID
// @Tags seller
// @Accept json
// @Produce json
// @Param seller_id path string true "Seller ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.Seller}
// @Failure 500 {object} dto.ErrorResponse
// @Router /seller/{seller_id} [get]
func (s SellerController) GetSellerByID(c *gin.Context) {
	sellerID := c.Param("seller_id")
	res, err := s.sellerService.GetSellerByID(sellerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No seller with this sellerID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get seller success",
		Data:    res,
	})
}

// GetSellers godoc
// @Summary Get all sellers
// @Description Retrieves all sellers
// @Tags seller
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.Seller}
// @Failure 500 {object} dto.ErrorResponse
// @Router /seller/ [get]
func (s SellerController) GetSellers(c *gin.Context) {
	res, err := s.sellerService.GetSellers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No sellers",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get Sellers success",
		Data:    res,
	})
}

// UpdateSeller godoc
// @Summary Update a seller by ID
// @Description Updates an existing seller's data by their ID
// @Tags seller
// @Accept json
// @Produce json
// @Param id path string true "Seller ID"
// @Param seller body model.Seller true "Seller data to update"
// @Success 200 {object} dto.SuccessResponse{data=dto.Seller}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /seller/{id} [put]
func (s SellerController) UpdateSeller(c *gin.Context) {
	sellerID := c.Param("id")

	var updatedSeller model.Seller
	if err := c.BindJSON(&updatedSeller); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	res, err := s.sellerService.UpdateSellerData(sellerID, &updatedSeller)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to update seller data",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Update seller success",
		Data:    res,
	})
}
