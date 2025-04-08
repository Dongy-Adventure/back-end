package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type IPaymentController interface {
	HandlePayment(c *gin.Context)
	OmiseWebhookHandler(c *gin.Context)
	SSEHandler(c *gin.Context)
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
// @Success 200 {object} dto.SuccessResponse{data=string} "Payment successful"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 502 {object} dto.ErrorResponse "Payment failed"
// @Router /payment [post]
func (p PaymentController) HandlePayment(c *gin.Context) {
	var paymentRequest dto.PaymentRequest
	if err := c.BindJSON(&paymentRequest); err != nil {
		fmt.Println("HERE")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	charge, err := p.paymentService.HandlePayment(&paymentRequest)
	if err != nil {
		c.JSON(http.StatusBadGateway, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadGateway,
			Error:   "Payment failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Payment success",
		Data:    charge,
	})
}

// @Summary SSE for Charge Status
// @Description Opens an SSE connection to track charge status updates.
// @Tags Payment
// @Accept json
// @Produce text/event-stream
// @Param charge_id path string true "Charge ID"
// @Success 200 {string} string "SSE connection established for charge status updates"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Router /payment/charge/:charge_id/sse [get]
func (p PaymentController) SSEHandler(c *gin.Context) {
	chargeID := c.Param("charge_id")
	if chargeID == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request",
			Message: "Missing charge_id parameter",
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	clientChan := make(chan string)

	p.paymentService.AddClient(chargeID, clientChan)

	defer func() {
		p.paymentService.RemoveClient(chargeID, clientChan)
		close(clientChan)
		log.Printf("SSE connection closed for charge %s", chargeID)
	}()

	c.Stream(func(w io.Writer) bool {
		select {
		case message := <-clientChan:
			_, err := io.WriteString(w, fmt.Sprintf("data: %s\n\n", message))
			if err != nil {
				log.Printf("Error writing SSE event: %v", err)
				return false
			}

			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			return true

		case <-c.Request.Context().Done():
			log.Printf("Client disconnected from SSE for charge %s", chargeID)
			return false
		}
	})
}

// @Summary Omise Webhook Handler
// @Description Handles webhook events from Omise, such as payment success or failure.
// @Tags Payment
// @Accept json
// @Produce json
// @Param webhookPayload body map[string]interface{} true "Webhook payload from Omise"
// @Success 200 {object} dto.SuccessResponse "Webhook processed successfully"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /webhook/omise [post]
func (p PaymentController) OmiseWebhookHandler(c *gin.Context) {
	var webhookPayload map[string]interface{}
	if err := c.ShouldBindJSON(&webhookPayload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid webhook payload",
			Message: err.Error(),
		})
		return
	}
	data, ok := webhookPayload["data"].(map[string]interface{})
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid webhook data",
			Message: "Missing 'data' field in webhook",
		})
		return
	}
	status, ok := data["status"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid webhook event",
			Message: "Missing 'event' field in webhook",
		})
		return
	}
	switch status {
	case "successful":
		chargeID, ok := data["id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
				Success: false,
				Status:  http.StatusBadRequest,
				Error:   "Invalid webhook data",
				Message: "Missing charge ID in webhook",
			})
			return
		}
		status, ok := data["status"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
				Success: false,
				Status:  http.StatusBadRequest,
				Error:   "Invalid webhook data",
				Message: "Missing status in webhook",
			})
			return
		}

		if err := p.paymentService.UpdatePaymentStatus(chargeID, status); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Status:  http.StatusInternalServerError,
				Error:   "Database error",
				Message: "Failed to update payment status",
			})
			return
		}

		p.paymentService.BroadcastChargeStatus(chargeID, status)

		c.JSON(http.StatusOK, dto.SuccessResponse{
			Success: true,
			Status:  http.StatusOK,
			Message: "Webhook processed successfully",
		})
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Unknown event type",
			Message: fmt.Sprintf("Unknown webhook event type: %s", status),
		})
	}

}
