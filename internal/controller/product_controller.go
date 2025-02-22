package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductController interface {
	CreateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	UpdateProduct(c *gin.Context)
}

type ProductController struct {
	productService service.IProductService
}

func NewProductController(s service.IProductService) IProductController {
	return ProductController{
		productService: s,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Creates a new product in the database
// @Tags product
// @Accept json
// @Produce json
// @Param product body dto.ProductRegisterRequest true "Product to create"
// @Success 201 {object} dto.SuccessResponse{data=dto.Product}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /product/ [post]
func (s ProductController) CreateProduct(c *gin.Context) {
	var newProduct dto.ProductPost

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	product := model.Product{
		// map fields from newProduct to product
		ProductName: newProduct.ProductName,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		// add other fields as necessary
	}
	res, err := s.productService.CreateProduct(&product)

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
		Message: "Product created",
		Data:    res,
	})
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Retrieves a product's data by their ID
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.Product}
// @Failure 500 {object} dto.ErrorResponse
// @Router /product/{product_id} [get]
func (s ProductController) GetProductByID(c *gin.Context) {
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
	res, err := s.productService.GetProductByID(productID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No product with this productID",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get product success",
		Data:    res,
	})
}

// GetProductBySellerID godoc
// @Summary Get a product by ID
// @Description Retrieves a product's data by their ID
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Success 200 {object} dto.SuccessResponse{data=dto.Product}
// @Failure 500 {object} dto.ErrorResponse
// @Router /product/{product_id} [get]
func (s ProductController) GetProductBySellerID(c *gin.Context) {
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
	res, err := s.productService.GetProductBySellerID(sellerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No product with this productID",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get product success",
		Data:    res,
	})
}

// GetProducts godoc
// @Summary Get all products
// @Description Retrieves all products
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=[]dto.Product}
// @Failure 500 {object} dto.ErrorResponse
// @Router /product/ [get]
func (s ProductController) GetProducts(c *gin.Context) {
	res, err := s.productService.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No products found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get products success",
		Data:    res,
	})
}

// UpdateProduct godoc
// @Summary Update a product by ID
// @Description Updates an existing product's data by their ID
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Param product body model.Product true "Product data to update"
// @Success 200 {object} dto.SuccessResponse{data=dto.Product}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /product/{product_id} [put]
func (s ProductController) UpdateProduct(c *gin.Context) {
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
	var updatedProduct model.Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	res, err := s.productService.UpdateProduct(productID, &updatedProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to update product data",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Update product success",
		Data:    res,
	})
}
