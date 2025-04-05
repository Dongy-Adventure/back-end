package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductController interface {
	CreateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	GetProductsBySellerID(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type ProductController struct {
	productService service.IProductService
	s3Service      service.IS3Service
}

func NewProductController(s service.IProductService, s3 service.IS3Service) IProductController {
	return ProductController{
		productService: s,
		s3Service:      s3,
	}
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Creates a new product in the database
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			product	body		dto.ProductCreateRequest	true	"Product to create"
//	@Success		201		{object}	dto.SuccessResponse{data=dto.Product}
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/product/ [post]
func (s ProductController) CreateProduct(c *gin.Context) {
	var newProduct dto.ProductCreateRequest

	if err := c.ShouldBind(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	var imageURL string
	if newProduct.Image != nil {
		fileUrl, err := s.s3Service.UploadFile(newProduct.Image, "products")
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Status:  http.StatusInternalServerError,
				Error:   "Failed to upload image to S3",
				Message: err.Error(),
			})
			return
		}
		imageURL = fileUrl
	}

	var newProductData model.Product
	if err := copier.Copy(&newProductData, &newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to copy seller data",
			Message: err.Error(),
		})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(newProduct.SellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid Seller ID",
			Message: err.Error(),
		})
		return
	}

	newProductData.Image = imageURL
	newProductData.SellerID = sellerID

	res, err := s.productService.CreateProduct(&newProductData)

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
//
//	@Summary		Get a product by ID
//	@Description	Retrieves a product's data by their ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string	true	"Product ID"
//	@Success		200			{object}	dto.SuccessResponse{data=dto.Product}
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/product/{product_id} [get]
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

// GetProductsBySellerID godoc
//
//	@Summary		Get products by sellerID
//	@Description	Retrieves each seller's products-on-display by seller ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			seller_id	path		string	true	"Seller ID"
//	@Success		200			{object}	dto.SuccessResponse{data=[]dto.Product}
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/product/seller/{seller_id} [get]
func (s ProductController) GetProductsBySellerID(c *gin.Context) {
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
	res, err := s.productService.GetProductsBySellerID(sellerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No product with this sellerID",
			Message: err.Error()})
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
//
//	@Summary		Get all products
//	@Description	Retrieves all products
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.SuccessResponse{data=[]dto.Product}
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/product/ [get]
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
//
//	@Summary		Update a product by ID
//	@Description	Updates an existing product's data by their ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string			true	"Product ID"
//	@Param			updatedProduct		body		dto.UpdateProductRequest	true	"Product data to update"
//	@Success		200			{object}	dto.SuccessResponse{data=dto.Product}
//	@Failure		400			{object}	dto.ErrorResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/product/{product_id} [put]
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
	var updatedProduct dto.UpdateProductRequest
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	sellerID, err := primitive.ObjectIDFromHex(updatedProduct.SellerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid productID format",
			Message: err.Error(),
		})
		return
	}
	res, err := s.productService.UpdateProduct(productID, &model.Product{
		ProductID:   productID,
		ProductName: updatedProduct.ProductName,
		Price:       updatedProduct.Price,
		Description: updatedProduct.Description,
		Image:       updatedProduct.Image,
		Tag:         updatedProduct.Tag,
		Color:       updatedProduct.Color,
		SellerID:    sellerID,
		Amount:      updatedProduct.Amount,
		CreatedAt:   updatedProduct.CreatedAt,
	})
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

// DeleteProduct godoc
//
//	@Summary		Delete a product by ID
//	@Description	Delete a product's data by its ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string	true	"Product ID"
//	@Success		200			{object}	dto.SuccessResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/product/{product_id} [delete]
func (s ProductController) DeleteProduct(c *gin.Context) {
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
	err = s.productService.DeleteProduct(productID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "No product with this productID",
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Delete product success",
	})
}
