package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type IPaymentController interface {
	HandlePayment(c *gin.Context)
}
type PaymentController struct {
	paymentService service.IPaymentService
}

func NewPaymentController(paymentService service.IPaymentService) IPaymentController {
	return PaymentController{paymentService: paymentService}
}

// @Summary Process payment
// @Description Handles payment processing through Omise.
// @Tags Payment
// @Accept json
// @Produce json
// @Param paymentRequest body dto.PaymentRequest true "Payment request payload"
// @Success 200 {object} dto.PaymentResponse "Payment successful"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 502 {object} dto.ErrorResponse "Payment failed"
// @Router /payment [post]
func (p PaymentController) HandlePayment(c *gin.Context) {
	var paymentRequest dto.BuyerPaymentRequest
	if err := c.BindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	err := p.paymentService.HandlePayment(&paymentRequest)
	if err != nil {
		c.JSON(http.StatusBadGateway, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadGateway,
			Error:   "Payment failed",
			Message: err.Error(),
		})
		return
	}
}
