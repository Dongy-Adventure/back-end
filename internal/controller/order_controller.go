package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/userrole"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderController interface {
	CreateOrder(c *gin.Context)
	GetOrdersByUserID(c *gin.Context)
}
type OrderController struct {
	orderService service.IOrderService
}

func NewOrderController(orderService service.IOrderService) IOrderController {
	return OrderController{orderService: orderService}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Creates a new order in the database
// @Tags order
// @Accept json
// @Produce json
// @Param buyer body dto.OrderCreateRequest true "Order to create"
// @Success 201 {object} dto.SuccessResponse{data=dto.Order}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order/ [post]
func (o OrderController) CreateOrder(c *gin.Context) {
	var req dto.OrderCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	newOrder, err := o.orderService.CreateOrder(req.Products, req.BuyerID, req.SellerID)
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
		Message: "Order created",
		Data:    newOrder,
	})
}

// GetOrdersByUserID godoc
// @Summary Get orders by userID and userType
// @Description Get all orders by userID and userType
// @Tags order
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param user_type path int true "User Type"
// @Success 200 {object} dto.SuccessResponse{data=[]dto.Order}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order/ [get]
func (o OrderController) GetOrdersByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userTypeStr := c.Param("user_type")

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid userID format",
			Message: err.Error(),
		})
		return
	}
	var userType userrole.UserType
	switch userTypeStr {
	case "0":
		userType = userrole.UserRole.BUYER

	case "1":
		userType = userrole.UserRole.SELLER
	default:
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Invalid userType",
			Message: "UserType must be either 0 or 1",
		})
		return
	}
	orders, err := o.orderService.GetOrdersByUserID(userID, userType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Failed to get orders",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Get orders success",
		Data:    orders,
	})
}
