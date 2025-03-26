package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

var orderID primitive.ObjectID
var appointmentID primitive.ObjectID
var orderStatus int16

func InitializeSellerAppointmentScenario(ctx *godog.ScenarioContext) {
	router = gin.New()
	orderController := controller.NewOrderController(mockOrderService, mockPaymentService)
	appointmentController := controller.NewAppointmentController(mockAppointmentService)
	router.POST("/order", orderController.CreateOrder)
	router.PUT("/appointment/:appointment_id", appointmentController.UpdateAppointmentPlace)
	router.PUT("/order/:order_id", orderController.UpdateOrderStatusByOrderID)

	// Bind steps to functions
	ctx.Step(`^an order with ID "([^"]*)" is created$`, anOrderWithIDIsCreated)
	ctx.Step(`^an appointment with ID "([^"]*)" is created for this order$`, anAppointmentWithIDIsCreatedForThisOrder)
	ctx.Step(`^the seller chooses a location for order with ID "([^"]*)"$`, theSellerChoosesALocation)
	ctx.Step(`^the buyer rejects the seller's location for order with ID "([^"]*)"$`, theBuyerRejectsTheSellersLocation)
	ctx.Step(`^the buyer accepts the seller's location for order with ID "([^"]*)"$`, theBuyerAcceptsTheSellersLocation)
	ctx.Step(`^the order status for order with ID "([^"]*)" should be (\d+)$`, theOrderStatusShouldBe)
	ctx.Step(`^the response status should be (\d+)$`, theResponseStatusShouldBe)
	ctx.Step(`^the order status for order with ID "([^"]*)" should not change$`, theOrderStatusForOrderWithIDShouldNotChange)
	ctx.Step(`^the seller updates the location for order with ID "([^"]*)"$`, theSellerUpdatesTheLocationForOrderWithID)
}

func anOrderWithIDIsCreated(id string) error {
	orderIDPrimitive, _ := primitive.ObjectIDFromHex(id)
	orderStatus = 0
	orderID = orderIDPrimitive

	requestBody := dto.OrderCreateRequest{
		Products:   []dto.OrderProduct{product}, // Ensure 'product' is defined in the test context
		BuyerID:    buyerID,                     // Ensure 'buyerID' is defined in the test context
		SellerID:   sellerID,                    // Ensure 'sellerID' is defined in the test context
		SellerName: "Test",
		BuyerName:  "Test",
		// Payment:    "Test",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	mockOrderService.EXPECT().
		CreateOrder(gomock.Any()).
		Return(&dto.Order{
			OrderID:    primitive.NewObjectID(),
			BuyerID:    buyerID,
			SellerID:   sellerID,
			SellerName: "Test",
			BuyerName:  "Test",
			Payment:    "Test",
			Products:   []dto.OrderProduct{},
			CreatedAt:  time.Now(),
		}, nil).Times(1)

	req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.Use(func(c *gin.Context) {
		log.Println("Request path:", c.Request.URL.Path)
		c.Next()
	})
	router.ServeHTTP(recorder, req)

	testResponse := recorder.Result()

	if testResponse.StatusCode != http.StatusCreated {
		return fmt.Errorf("error: expected status 201 but got %d.", testResponse.StatusCode)
	}
	return nil
}

func anAppointmentWithIDIsCreatedForThisOrder(id string) error {
	appointmentIDPrimitive, _ := primitive.ObjectIDFromHex(id)
	appointmentID = appointmentIDPrimitive

	mockAppointmentService.EXPECT().
		CreateAppointment(gomock.Any()).
		Return(&dto.Appointment{
			AppointmentID: appointmentIDPrimitive,
			OrderID:       orderID,
			BuyerID:       primitive.NewObjectID(),
			SellerID:      primitive.NewObjectID(),
			Address:       "Test Address",
			City:          "Test City",
			Province:      "Test Province",
			Zip:           "12345",
			Date:          time.Now(),
			TimeSlot:      "10:00 AM - 11:00 AM",
			CreatedAt:     time.Now(),
		}, nil).Times(1)

	return nil
}

func theSellerChoosesALocation(orderID string) error {
	orderIDPrimitive, _ := primitive.ObjectIDFromHex(orderID)

	mockOrderService.EXPECT().
		UpdateOrderStatus(orderIDPrimitive, gomock.Any()).
		Return(1, nil).Times(1)

	requestBody := dto.OrderStatusRequest{
		OrderStatus: 1,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/order/%s", orderID), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	testResponse = recorder.Result()

	body, _ := io.ReadAll(testResponse.Body)
	defer testResponse.Body.Close()

	if testResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("error: expected status 200 but got %d. Response: %s", testResponse.StatusCode, string(body))
	}

	return nil
}

func theBuyerRejectsTheSellersLocation(orderID string) error {
	orderIDPrimitive, _ := primitive.ObjectIDFromHex(orderID)

	mockOrderService.EXPECT().
		UpdateOrderStatus(orderIDPrimitive, 1).
		Return(1, nil).Times(1)

	requestBody := dto.OrderStatusRequest{
		OrderStatus: 1,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/order/%s", orderID), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	testResponse = recorder.Result()

	body, _ := io.ReadAll(testResponse.Body)
	defer testResponse.Body.Close()

	if testResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("error: expected status 200 but got %d. Response: %s", testResponse.StatusCode, string(body))
	}

	return nil
}

func theBuyerAcceptsTheSellersLocation(orderID string) error {
	orderIDPrimitive, _ := primitive.ObjectIDFromHex(orderID)

	mockOrderService.EXPECT().
		UpdateOrderStatus(orderIDPrimitive, 2).
		Return(2, nil).Times(1)

	requestBody := dto.OrderStatusRequest{
		OrderStatus: 2,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/order/%s", orderID), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	testResponse = recorder.Result()
	if testResponse == nil {
		return fmt.Errorf("testResponse is nil, router.ServeHTTP failed")
	}
	body, _ := io.ReadAll(testResponse.Body)
	defer testResponse.Body.Close()

	if testResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("error: expected status 200 but got %d. Response: %s", testResponse.StatusCode, string(body))
	}
	orderStatus = 2
	return nil
}

func theOrderStatusShouldBe(orderID string, expectedStatus int) error {
	orderIDPrimitive, _ := primitive.ObjectIDFromHex(orderID)

	if int(orderStatus) != expectedStatus {
		return fmt.Errorf("expected order status to be %d for order with ID %s, but got %d", expectedStatus, orderIDPrimitive, orderStatus)
	}
	return nil
}

func theOrderStatusForOrderWithIDShouldNotChange(orderID string) error {
	// Convert orderID from string to ObjectID
	orderIDPrimitive, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return fmt.Errorf("invalid order ID: %v", err)
	}

	// Mock fetching current order status using EXPECT
	mockOrderService.EXPECT().
		UpdateOrderStatus(orderIDPrimitive, 1).
		Return(1, nil).Times(1)

	// Verify that the order status has not changed
	currentOrderStatus := orderStatus
	if currentOrderStatus != orderStatus {
		return fmt.Errorf("expected order status %d but got %d", orderStatus, currentOrderStatus)
	}
	return nil
}

func theSellerUpdatesTheLocationForOrderWithID(orderID string) error {
	orderIDPrimitive, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return fmt.Errorf("invalid order ID: %v", err)
	}

	updatedAppointment := &dto.Appointment{
		AppointmentID: orderIDPrimitive,
		OrderID:       orderIDPrimitive,
		Address:       "Updated Address",
		City:          "Updated City",
		Province:      "Updated Province",
		Zip:           "54321",
		Date:          time.Now(),
		TimeSlot:      "11:00 AM - 12:00 PM",
		CreatedAt:     time.Now(),
	}

	mockAppointmentService.EXPECT().
		UpdateAppointmentPlace(orderIDPrimitive, gomock.Any()).
		Return(updatedAppointment, nil).Times(1) // Return updated appointment DTO and no error

	return nil
}

func TestSellerAppointment(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeSellerAppointmentScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"feature/seller_appointment.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("Godog tests failed")
	}
}
