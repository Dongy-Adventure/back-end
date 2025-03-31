package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAdvertisementController interface {
	GetAdvertisements(c *gin.Context)
	GetAdvertisementByID(c *gin.Context)
	GetAdvertisementsBySellerID(c *gin.Context)
	GetAdvertisementsByProductID(c *gin.Context)
	CreateAdvertisement(c *gin.Context)
	UpdateAdvertisement(c *gin.Context)
	DeleteAdvertisement(c *gin.Context)
}

type AdvertisementController struct {
	advertisementService service.IAdvertisementService
}

func NewAdvertisementController(s service.IAdvertisementService) IAdvertisementController {
	return AdvertisementController{
		advertisementService: s,
	}
}

// GetAdvertisements godoc
//
//	@Summary		Get all advertisements
//	@Description	Retrieves all advertisements
//	@Tags			advertisement
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.SuccessResponse{data=[]dto.Advertisement}
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/advertisement/ [get]
func (s AdvertisementController) GetAdvertisements(c *gin.Context) {
	res, err := s.advertisementService.GetAdvertisements()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No advertisements",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get Advertisements success",
		Data:    res,
	})
}

// GetAdvertisementsByID godoc
//
//	@Summary		Get a advertisement by ID
//	@Description	Retrieves a advertisement's data by its ID
//	@Tags			advertisement
//	@Accept			json
//	@Produce		json
//	@Param			advertisement_id	path		string	true	"Advertisement ID"
//	@Success		200			{object}	dto.SuccessResponse{data=dto.Advertisement}
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/advertisement/{advertisement_id} [get]
func (s AdvertisementController) GetAdvertisementByID(c *gin.Context) {
	advertisementIDstr := c.Param("advertisement_id")
	advertisementID, err := primitive.ObjectIDFromHex(advertisementIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid advertisementID format",
			Message: err.Error(),
		})
		return
	}
	res, err := s.advertisementService.GetAdvertisementByID(advertisementID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No advertisement with this advertisementID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get advertisement success",
		Data:    res,
	})
}

// GetAdvertisementsBySellerID godoc
//
//	@Summary		Get advertisements by sellerID
//	@Description	Retrieves each seller's advertisements by seller ID
//	@Tags			advertisement
//	@Accept			json
//	@Produce		json
//	@Param			seller_id	path		string	true	"Seller ID"
//	@Success		200			{object}	dto.SuccessResponse{data=[]dto.Advertisement}
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/advertisement/seller/{seller_id} [get]
func (s AdvertisementController) GetAdvertisementsBySellerID(c *gin.Context) {
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
	res, err := s.advertisementService.GetAdvertisementsBySellerID(sellerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No advertisement with this sellerID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get advertisement success",
		Data:    res,
	})
}

// GetAdvertisementsByProductID godoc
//
//	@Summary		Get advertisements by productID
//	@Description	Retrieves each product's advertisements by product ID
//	@Tags			advertisement
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string	true	"Product ID"
//	@Success		200			{object}	dto.SuccessResponse{data=[]dto.Advertisement}
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/advertisement/product/{product_id} [get]
func (s AdvertisementController) GetAdvertisementsByProductID(c *gin.Context) {
	productIDstr := c.Param("product_id")
	productID, err := primitive.ObjectIDFromHex(productIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid productID format",
			Message: err.Error(),
		})
		return
	}
	res, err := s.advertisementService.GetAdvertisementsByProductID(productID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No advertisement with this productID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get advertisement success",
		Data:    res,
	})
}

// CreateAdvertisement godoc

// CreateAdvertisement godoc
//
//	@Summary		Create a new advertisement
//	@Description	Creates a new advertisement in the database
//	@Tags			advertisement
//	@Accept			json
//	@Produce		json
//	@Param			advertisement	body		dto.AdvertisementCreateRequest	true	"Advertisement to create"
//	@Success		201		{object}	dto.SuccessResponse{data=dto.Advertisement}
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/advertisement/ [post]
func (s AdvertisementController) CreateAdvertisement(c *gin.Context) {
	var newAdvertisement model.Advertisement

	if err := c.BindJSON(&newAdvertisement); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	res, err := s.advertisementService.CreateAdvertisement(&newAdvertisement)

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
		Message: "Advertisement created",
		Data:    res,
	})
}

// UpdateAdvertisement godoc
//
//	@Summary		Update a advertisement by ID
//	@Description	Updates an existing advertisement's data by its ID
//	@Tags			advertisement
//	@Accept			json
//	@Produce		json
//	@Param			advertisement_id	path		string					true	"Advertisement ID"
//	@Param			advertisement		body		dto.AdvertisementUpdateRequest	true	"Advertisement data to update"
//	@Success		200			{object}	dto.SuccessResponse{data=dto.Advertisement}
//	@Failure		400			{object}	dto.ErrorResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/advertisement/{advertisement_id} [put]
func (s AdvertisementController) UpdateAdvertisement(c *gin.Context) {
	advertisementIDstr := c.Param("advertisement_id")
	advertisementID, err := primitive.ObjectIDFromHex(advertisementIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid advertisementID format",
			Message: err.Error(),
		})
		return
	}
	var updatedAdvertisement model.Advertisement
	if err := c.BindJSON(&updatedAdvertisement); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	res, err := s.advertisementService.UpdateAdvertisement(advertisementID, &updatedAdvertisement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to update advertisement data",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Update advertisement success",
		Data:    res,
	})
}

// DeleteAdvertisement godoc
//
//	@Summary		Delete a advertisement by ID
//	@Description	Delete a advertisement's data by its ID
//	@Tags			advertisement
//	@Accept			json
//	@Produce		json
//	@Param			advertisement_id	path		string	true	"Advertisement ID"
//	@Success		200			{object}	dto.SuccessResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/advertisement/{advertisement_id} [delete]
func (s AdvertisementController) DeleteAdvertisement(c *gin.Context) {
	advertisementIDstr := c.Param("advertisement_id")
	advertisementID, err := primitive.ObjectIDFromHex(advertisementIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid advertisementID format",
			Message: err.Error(),
		})
		return
	}
	err = s.advertisementService.DeleteAdvertisement(advertisementID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No advertisement with this advertisementID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Delete advertisement success",
	})
}
