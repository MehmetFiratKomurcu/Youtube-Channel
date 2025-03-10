package mocks

import (
	"context"
	"yt-otel-metrics/src/application/domain/entity"
	"yt-otel-metrics/src/application/model"

	"github.com/stretchr/testify/mock"
)

// MockOrderService is a mock implementation of the OrderService interface
type MockOrderService struct {
	mock.Mock
}

// CreateOrder mocks the CreateOrder method
func (m *MockOrderService) CreateOrder(command model.CreateOrderCommand) *entity.Order {
	args := m.Called(command)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entity.Order)
}

// GetOrderById mocks the GetOrderById method
func (m *MockOrderService) GetOrderById(ctx context.Context, id int64) *entity.Order {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entity.Order)
}

// ShipOrderByCargoCode mocks the ShipOrderByCargoCode method
func (m *MockOrderService) ShipOrderByCargoCode(cargoCode string) error {
	args := m.Called(cargoCode)
	return args.Error(0)
}
