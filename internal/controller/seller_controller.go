package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISellerController interface {
	CreateSeller(c *gin.Context)
	GetSellerByID(c *gin.Context)
	GetSellers(c *gin.Context)
	UpdateSeller(c *gin.Context)
	AddTransaction(c *gin.Context)
	GetSellerBalanceByID(c *gin.Context)
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
//	@Summary		Create a new seller
//	@Description	Creates a new seller in the database
//	@Tags			seller
//	@Accept			json
//	@Produce		json
//	@Param			seller	body		dto.SellerRegisterRequest	true	"Seller to create"
//	@Success		201		{object}	dto.SuccessResponse{data=dto.Seller}
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/seller/ [post]
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
//	@Summary		Get a seller by ID
//	@Description	Retrieves a seller's data by their ID
//	@Tags			seller
//	@Accept			json
//	@Produce		json
//	@Param			seller_id	path		string	true	"Seller ID"
//	@Success		200			{object}	dto.SuccessResponse{data=dto.Seller}
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/seller/{seller_id} [get]
func (s SellerController) GetSellerByID(c *gin.Context) {
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
//	@Summary		Get all sellers
//	@Description	Retrieves all sellers
//	@Tags			seller
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.SuccessResponse{data=[]dto.Seller}
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/seller/ [get]
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
//	@Summary		Update a seller by ID
//	@Description	Updates an existing seller's data by their ID
//	@Tags			seller
//	@Accept			json
//	@Produce		json
//	@Param			seller_id	path		string			true	"Seller ID"
//	@Param			seller		body		model.Seller	true	"Seller data to update"
//	@Success		200			{object}	dto.SuccessResponse{data=dto.Seller}
//	@Failure		400			{object}	dto.ErrorResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/seller/{seller_id} [put]
func (s SellerController) UpdateSeller(c *gin.Context) {
	sellerIDstr := c.Param("seller_id")
	userID, exists := c.Get("userID")
	if userID != sellerIDstr || !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Error:   "ID not match or not exists",
			Message: "param ID doesn't match with callerID"})
		return
	}
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

	res, err := s.sellerService.UpdateSeller(sellerID, &updatedSeller)
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

// AddTransaction godoc
//	@Summary		Add a transaction by sellerID
//	@Description	Append transaction to seller transactions
//	@Tags			seller
//	@Accept			json
//	@Produce		json
//	@Param			seller_id	path		string			true	"Seller ID"
//	@Param			transaction	body		dto.Transaction	true	"Transaction to append"
//	@Success		201			{object}	dto.SuccessResponse{data=dto.Transaction}
//	@Failure		400			{object}	dto.ErrorResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/seller/{seller_id}/transaction [post]
func (s SellerController) AddTransaction(c *gin.Context) {
	sellerIDstr := c.Param("seller_id")
	userID, exists := c.Get("userID")
	if userID != sellerIDstr || !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Error:   "ID not match or not exists",
			Message: "param ID doesn't match with callerID"})
		return
	}
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

	var newTransaction dto.Transaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	res, err := s.sellerService.AddTransaction(sellerID, &newTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to add transaction",
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Add transaction success",
		Data:    res,
	})
}

// GetSellerBalanceByID godoc
//	@Summary		Get a seller's total balance by ID
//	@Description	Retrieves a seller's total balance by their ID
//	@Tags			seller
//	@Accept			json
//	@Produce		json
//	@Param			seller_id	path		string	true	"Seller ID"
//	@Success		200			{object}	dto.SuccessResponse{data=float64}
//	@Failure		400			{object}	dto.ErrorResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/seller/{seller_id}/balance [get]
func (s SellerController) GetSellerBalanceByID(c *gin.Context) {
	sellerIDstr := c.Param("seller_id")
	userID, exists := c.Get("userID")
	if userID != sellerIDstr || !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Error:   "ID not match or not exists",
			Message: "param ID doesn't match with callerID"})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(sellerIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid sellerID format",
			Message: err.Error(),
		})
		return
	}

	balance, err := s.sellerService.GetSellerBalanceByID(sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to retrieve seller balance",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get seller balance success",
		Data:    balance,
	})
}
