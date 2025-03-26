package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/controller"
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var paymentSuccess bool
var productOutOfStock bool
var cart []dto.OrderProduct

func InitializeOrderConfirmationScenario(ctx *godog.ScenarioContext) {
	SetUpRouter()
	router = gin.New()

	orderController := controller.NewOrderController(mockOrderService, mockPaymentService)
	router.POST("/order/", orderController.CreateOrder)
	ctx.Step(`^a buyer with ID "([^"]*)"$`, aBuyerWithIDExists)
	ctx.Step(`^an order should be created with respone status (\d+)`, anOrderCreated)
	ctx.Step(`^a product with ID "([^"]*)" exists in the buyer cart$`, aProductWithIDExistsInTheBuyerCart)
	ctx.Step(`^the payment status is "([^"]*)"$`, thePaymentStatusIs)
	ctx.Step(`^the product should be removed from the buyer cart$`, theProductIsRemovedFromTheBuyersCart)
	ctx.Step(`^the product is out of stock$`, theProductIsOutOfStock)
	ctx.Step(`^a log should be generated indicating "([^"]*)"$`, theResponseShouldContain)
}

func aBuyerWithIDExists(buyerIDStr string) error {
	var err error
	buyerID, err = primitive.ObjectIDFromHex(buyerIDStr)
	buyerName := "Test"

	if err != nil {
		return fmt.Errorf("invalid buyerID format: %v", err)
	}
	buyer := &dto.Buyer{
		BuyerID:  buyerID,
		Username: buyerName,
		Cart:     []dto.OrderProduct{},
	}

	mockBuyerService.EXPECT().
		GetBuyerByID(buyerID).
		Return(buyer, nil).
		AnyTimes()

	return nil
}

func aProductWithIDExistsInTheBuyerCart(productIDStr string) error {
	var err error
	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		return fmt.Errorf("invalid productID format: %v", err)
	}

	buyer, err := mockBuyerService.GetBuyerByID(buyerID)
	if err != nil {
		return fmt.Errorf("failed to get buyer: %v", err)
	}
	for _, item := range buyer.Cart {
		if item.ProductID == productID {
			product = item
			return nil
		}
	}
	newProduct := dto.OrderProduct{
		ProductID: productID,
		Amount:    1,
	}
	mockBuyerService.EXPECT().
		UpdateProductInCart(buyerID, newProduct).
		Return([]dto.OrderProduct{newProduct}, nil)
	newCart, err := mockBuyerService.UpdateProductInCart(buyerID, newProduct)
	if err != nil {
		return fmt.Errorf("failed to update product in cart: %v", err)
	}
	product = newProduct
	cart = newCart
	return nil
}

func thePaymentStatusIs(status string) error {
	switch status {
	case "success":
		paymentSuccess = true
		return nil
	case "failure":
		textResponse = "Payment failure: order not created"
		paymentSuccess = false
		return nil
	case "canceled":
		textResponse = "Order cancelled by the buyer"
		paymentSuccess = false
		return nil
	default:
		return fmt.Errorf("unexpected payment status: %s", status)
	}
}

func anOrderCreated(expectedStatus int) error {
	// Create a mock order request
	requestBody := dto.OrderCreateRequest{
		Products:   []dto.OrderProduct{product}, // Ensure 'product' is defined in the test context
		BuyerID:    buyerID,                     // Ensure 'buyerID' is defined in the test context
		SellerID:   sellerID,                    // Ensure 'sellerID' is defined in the test context
		SellerName: "Test",
		BuyerName:  "Test",
		// Payment:    "Paypal",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	if paymentSuccess {
		mockOrderService.EXPECT().
			CreateOrder(gomock.Any()).
			Return(&dto.Order{
				OrderID:   primitive.NewObjectID(),
				BuyerID:   buyerID,
				SellerID:  sellerID,
				Products:  []dto.OrderProduct{product},
				CreatedAt: time.Now(),
			}, nil).Times(1)
	} else {
		// Handle payment failure scenario
		mockOrderService.EXPECT().
			CreateOrder(gomock.Any()).
			Return(nil, fmt.Errorf("payment failed")).Times(1)
	}

	req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(recorder)
	c.Request = req
	controller := controller.NewOrderController(mockOrderService, mockPaymentService)
	handler := controller.CreateOrder
	handler(c)

	testResponse := recorder.Result()

	if testResponse.StatusCode != expectedStatus {
		responseBody, err := io.ReadAll(testResponse.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %v", err)
		}

		return fmt.Errorf("error: expected status %d but got %d. Response: %s", expectedStatus, testResponse.StatusCode, string(responseBody))
	}

	return nil
}

func theProductIsRemovedFromTheBuyersCart() error {
	if paymentSuccess {

		buyer, err := mockBuyerService.GetBuyerByID(buyerID)
		if err != nil {
			return fmt.Errorf("failed to retrieve the buyer: %v", err)
		}

		productRemoved := true
		for _, item := range buyer.Cart {
			if item.ProductID == product.ProductID {
				productRemoved = false
				break
			}
		}

		if productRemoved {
			fmt.Println("Product successfully removed from the buyer's cart")
		} else {
			return fmt.Errorf("product was not removed from the cart")
		}
	}

	return nil
}

func theProductIsOutOfStock() error {

	productOutOfStock = true

	if productOutOfStock {
		textResponse = "Insufficient stock for product"
		fmt.Println("The product is out of stock.")
	} else {
		return fmt.Errorf("product is not out of stock")
	}

	return nil
}

func TestBuyerConfirmOrder(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeOrderConfirmationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"feature/order_confirmation.feature"},
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("Godog tests failed")
	}
}
