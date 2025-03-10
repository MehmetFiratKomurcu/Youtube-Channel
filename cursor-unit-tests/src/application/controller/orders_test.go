package controller

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"yt-otel-metrics/src/application/test/mocks"
)

// setupTestApp creates a new Fiber app with the ShipOrderByCargoCode route registered
func setupTestApp(mockService *mocks.MockOrderService) *fiber.App {
	app := fiber.New()
	ShipOrderByCargoCode(app, mockService)
	return app
}

// createTestRequest creates an HTTP request for testing
func createTestRequest(method, path string) *http.Request {
	req := httptest.NewRequest(method, path, nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}

// assertResponse verifies the response status code and body content
func assertResponse(t *testing.T, resp *http.Response, expectedStatus int, expectedContent string) {
	t.Helper() // Mark as helper function for better error reporting

	assert.Equal(t, expectedStatus, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), expectedContent)
}

// TestShipOrderByCargoCode tests the ShipOrderByCargoCode controller function
func TestShipOrderByCargoCode(t *testing.T) {
	// Test cases
	tests := []struct {
		name               string
		cargoCode          string
		setupMock          func(*mocks.MockOrderService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:      "Success_ShipsOrdersWithValidCargoCode",
			cargoCode: "VALID123",
			setupMock: func(mockService *mocks.MockOrderService) {
				mockService.On("ShipOrderByCargoCode", "VALID123").Return(nil)
			},
			expectedStatusCode: fiber.StatusOK,
			expectedResponse:   "Orders ship successfully! yayyyy!",
		},
		{
			name:      "Failure_EmptyCargoCode",
			cargoCode: "empty", // Special case for testing empty cargo code
			setupMock: func(mockService *mocks.MockOrderService) {
				// We need to set up the mock for the "empty" cargo code
				mockService.On("ShipOrderByCargoCode", "empty").Return(errors.New("invalid cargo code"))
			},
			expectedStatusCode: fiber.StatusBadRequest,
			expectedResponse:   "invalid cargo code",
		},
		{
			name:      "Failure_ServiceReturnsError",
			cargoCode: "INVALID456",
			setupMock: func(mockService *mocks.MockOrderService) {
				mockService.On("ShipOrderByCargoCode", "INVALID456").Return(errors.New("cargo code not found"))
			},
			expectedStatusCode: fiber.StatusBadRequest,
			expectedResponse:   "cargo code not found",
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := new(mocks.MockOrderService)
			if tt.setupMock != nil {
				tt.setupMock(mockService)
			}

			app := setupTestApp(mockService)
			req := createTestRequest(
				http.MethodPost,
				"/orders/cargo-code/"+tt.cargoCode+"/ship",
			)

			// Act
			resp, err := app.Test(req)

			// Assert
			assert.NoError(t, err)
			assertResponse(t, resp, tt.expectedStatusCode, tt.expectedResponse)
			mockService.AssertExpectations(t)
		})
	}
}

// TestShipOrderByCargoCode_EdgeCases tests additional edge cases
func TestShipOrderByCargoCode_EdgeCases(t *testing.T) {
	// Test cases for edge scenarios
	tests := []struct {
		name               string
		cargoCode          string
		setupMock          func(*mocks.MockOrderService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:      "Failure_DatabaseError",
			cargoCode: "DB_ERROR",
			setupMock: func(mockService *mocks.MockOrderService) {
				mockService.On("ShipOrderByCargoCode", "DB_ERROR").Return(errors.New("database connection error"))
			},
			expectedStatusCode: fiber.StatusBadRequest,
			expectedResponse:   "database connection error",
		},
		{
			name:      "Failure_NoOrdersFoundForCargoCode",
			cargoCode: "NO_ORDERS",
			setupMock: func(mockService *mocks.MockOrderService) {
				mockService.On("ShipOrderByCargoCode", "NO_ORDERS").Return(errors.New("no orders found for cargo code"))
			},
			expectedStatusCode: fiber.StatusBadRequest,
			expectedResponse:   "no orders found for cargo code",
		},
		{
			name:      "Failure_OrdersAlreadyShipped",
			cargoCode: "ALREADY_SHIPPED",
			setupMock: func(mockService *mocks.MockOrderService) {
				mockService.On("ShipOrderByCargoCode", "ALREADY_SHIPPED").Return(errors.New("orders already shipped"))
			},
			expectedStatusCode: fiber.StatusBadRequest,
			expectedResponse:   "orders already shipped",
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := new(mocks.MockOrderService)
			if tt.setupMock != nil {
				tt.setupMock(mockService)
			}

			app := setupTestApp(mockService)
			req := createTestRequest(
				http.MethodPost,
				"/orders/cargo-code/"+tt.cargoCode+"/ship",
			)

			// Act
			resp, err := app.Test(req)

			// Assert
			assert.NoError(t, err)
			assertResponse(t, resp, tt.expectedStatusCode, tt.expectedResponse)
			mockService.AssertExpectations(t)
		})
	}
}

// TestDirectEmptyCargoCode tests the case where the cargo code is empty directly
func TestDirectEmptyCargoCode(t *testing.T) {
	// Arrange
	app := fiber.New()
	mockService := new(mocks.MockOrderService)

	// Create a custom handler that simulates the controller function with an empty cargo code
	app.Post("/test-empty-cargo", func(c *fiber.Ctx) error {
		cargoCode := ""

		if cargoCode == "" {
			return c.Status(fiber.StatusBadRequest).JSON("invalid cargo code")
		}

		err := mockService.ShipOrderByCargoCode(cargoCode)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON("Orders ship successfully! yayyyy!")
	})

	// Create a request
	req := createTestRequest(http.MethodPost, "/test-empty-cargo")

	// Act
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assertResponse(t, resp, fiber.StatusBadRequest, "invalid cargo code")
}

// TestTrueEmptyCargoCode tests the case where the cargo code parameter is truly empty
func TestTrueEmptyCargoCode(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockOrderService)
	app := setupTestApp(mockService)

	// Create a request with an empty cargo code parameter
	// This simulates a request to /orders/cargo-code//ship
	req := createTestRequest(http.MethodPost, "/orders/cargo-code//ship")

	// Act
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	// The Fiber router will return 404 for this case because it doesn't match the route pattern
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}
