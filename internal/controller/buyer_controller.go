package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IBuyerController interface {
	CreateBuyer(c *gin.Context)
	GetBuyers(c *gin.Context)
	GetBuyerByID(c *gin.Context)
	UpdateBuyer(c *gin.Context)
	UpdateProductInCart(c *gin.Context)
	DeleteProductFromCart(c *gin.Context)
}

type BuyerController struct {
	buyerService service.IBuyerService
}

func NewBuyerController(s service.IBuyerService) IBuyerController {
	return BuyerController{
		buyerService: s,
	}
}

// CreateBuyer godoc
// @Summary Create a new buyer
// @Description Creates a new buyer in the database
// @Tags buyer
// @Accept json
// @Produce json
// @Param buyer body dto.BuyerRegisterRequest true "Buyer to create"
// @Success 201 {object} dto.SuccessResponse{data=dto.Buyer}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /buyer/ [post]
func (s BuyerController) CreateBuyer(c *gin.Context) {
	var newBuyer model.Buyer

	if err := c.BindJSON(&newBuyer); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	res, err := s.buyerService.CreateBuyerData(&newBuyer)

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
		Message: "Buyer created",
		Data:    res,
	})
}

// GetBuyerByID godoc
// @Summary Get a buyer by ID
// @Description Retrieves a buyer's data by their ID
// @Tags buyer
// @Accept json
// @Produce json
// @Param buyer_id path string true "Buyer ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.Buyer}
// @Failure 500 {object} dto.ErrorResponse
// @Router /buyer/{buyer_id} [get]
func (s BuyerController) GetBuyerByID(c *gin.Context) {
	buyerIDstr := c.Param("buyer_id")
	// userID, exists := c.Get("userID")
	// if userID != buyerIDstr || !exists {
	// 	c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
	// 		Success: false,
	// 		Status:  http.StatusUnauthorized,
	// 		Error:   "ID not match or not exists",
	// 		Message: "param ID doesn't match with callerID"})
	// 	return
	// }
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
	res, err := s.buyerService.GetBuyerByID(buyerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No buyer with this buyerID",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get buyer success",
		Data:    res,
	})
}

// GetBuyer godoc
// @Summary Get all buyers
// @Description Retrieves all buyers
// @Tags buyer
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.Buyer}
// @Failure 500 {object} dto.ErrorResponse
// @Router /buyer/ [get]
func (s BuyerController) GetBuyers(c *gin.Context) {
	res, err := s.buyerService.GetBuyer()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No buyers",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get buyers success",
		Data:    res,
	})
}

// UpdateBuyer godoc
// @Summary Update a buyer by ID
// @Description Updates an existing buyer's data by their ID
// @Tags buyer
// @Accept json
// @Produce json
// @Param id path string true "Buyer ID"
// @Param buyer body model.Buyer true "Buyer data to update"
// @Success 200 {object} dto.SuccessResponse{data=dto.Buyer}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /buyer/{id} [put]
func (s BuyerController) UpdateBuyer(c *gin.Context) {
	buyerIDstr := c.Param("buyer_id")
	userID, exists := c.Get("userID")
	if userID != buyerIDstr || !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Error:   "ID not match or not exists",
			Message: "param ID doesn't match with callerID"})
		return
	}
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
	var updatedBuyer model.Buyer
	if err := c.BindJSON(&updatedBuyer); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	res, err := s.buyerService.UpdateBuyerData(buyerID, &updatedBuyer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to update buyer data",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Update buyer success",
		Data:    res,
	})
}

// UpdateProductInCart godoc
// @Summary Update a product in buyer's cart
// @Description Adds the product with specified amount if not in the cart, set amount if already in the cart
// @Tags buyer
// @Accept json
// @Produce json
// @Param buyer_id path string true "Buyer ID"
// @Param orderProduct body dto.OrderProduct true "Product to update in cart"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.OrderProduct}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /buyer/{buyer_id}/cart [post]
func (s BuyerController) UpdateProductInCart(c *gin.Context) {
	buyerIDStr := c.Param("buyer_id")
	userID, exists := c.Get("userID")

	if userID != buyerIDStr || !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Error:   "Unauthorized",
			Message: "You are not allowed to modify this cart",
		})
		return
	}

	buyerID, err := primitive.ObjectIDFromHex(buyerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid buyerID format",
			Message: err.Error(),
		})
		return
	}

	var request dto.OrderProduct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}

	updatedCart, err := s.buyerService.UpdateProductInCart(buyerID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to update cart",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Cart updated successfully",
		Data:    updatedCart,
	})
}

// DeleteProductFromCart godoc
// @Summary Delete a product from the buyer's cart
// @Description Deletes the product with the specified productID from the cart
// @Tags buyer
// @Accept json
// @Produce json
// @Param buyer_id path string true "Buyer ID"
// @Param product_id path string true "Product ID to delete from cart"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /buyer/{buyer_id}/cart/{product_id} [delete]
func (s BuyerController) DeleteProductFromCart(c *gin.Context) {
	buyerIDStr := c.Param("buyer_id")
	productIDStr := c.Param("product_id")
	userID, exists := c.Get("userID")

	if userID != buyerIDStr || !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Error:   "Unauthorized",
			Message: "You are not allowed to modify this cart",
		})
		return
	}

	buyerID, err := primitive.ObjectIDFromHex(buyerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid buyerID format",
			Message: err.Error(),
		})
		return
	}

	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid productID format",
			Message: err.Error(),
		})
		return
	}

	err = s.buyerService.DeleteProductFromCart(buyerID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to delete product from cart",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Product deleted from cart successfully",
	})
}
