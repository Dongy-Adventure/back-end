package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	mock "github.com/Dongy-s-Advanture/back-end/pkg/mock/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	router                 *gin.Engine
	mockOrderService       *mock.MockIOrderService
	mockBuyerService       *mock.MockIBuyerService
	mockSellerService      *mock.MockISellerService
	mockProductService     *mock.MockIProductService
	mockAppointmentService *mock.MockIAppointmentService
	testResponse           *http.Response
	buyerID                primitive.ObjectID
	sellerID               primitive.ObjectID
	product                dto.Product
	textResponse           string
)

func SetUpRouter() {
	if router == nil {
		router = gin.New()
	}
}

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	router = gin.New()
	ctrl := gomock.NewController(nil)
	mockOrderService = mock.NewMockIOrderService(ctrl)
	mockBuyerService = mock.NewMockIBuyerService(ctrl)
	mockSellerService = mock.NewMockISellerService(ctrl)
	mockProductService = mock.NewMockIProductService(ctrl)
	mockAppointmentService = mock.NewMockIAppointmentService(ctrl)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func theResponseStatusShouldBe(statusCode int) error {
	if testResponse == nil {
		return fmt.Errorf("testResponse is nil, make sure the request was made before checking the status")
	}
	if testResponse.StatusCode != statusCode {
		return fmt.Errorf("expected status %d but got %d", statusCode, testResponse.StatusCode)
	}
	return nil
}
func theResponseShouldContain(expectedMessage string) error {
	if !assert.Contains(nil, textResponse, expectedMessage) {
		return fmt.Errorf("expected response to contain '%s' but got '%s'", expectedMessage, textResponse)
	}
	return nil
}
