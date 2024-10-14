package services

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"yt-otel-tracing/src/application/domain/entity"
	"yt-otel-tracing/src/application/domain/persistance"
	"yt-otel-tracing/src/application/model"
)

type orderService struct {
	orderRepository persistance.OrderRepository
	tracer          trace.Tracer
}

type OrderService interface {
	CreateOrder(command model.CreateOrderCommand) *entity.Order
	GetOrderById(ctx context.Context, id int64) *entity.Order
	ShipOrderByCargoCode(cargoCode string) error
}

func (service orderService) GetOrderById(ctx context.Context, id int64) *entity.Order {
	ctx, span := service.tracer.Start(ctx, "GetOrderById Service")
	defer span.End()

	return service.orderRepository.GetOrderById(ctx, id)
}

func (service orderService) CreateOrder(command model.CreateOrderCommand) *entity.Order {
	order := model.MapToOrder(command)
	return service.orderRepository.CreateOrder(order)
}

func (service orderService) ShipOrderByCargoCode(cargoCode string) error {
	return service.orderRepository.ShipOrderByCargoCode(cargoCode)
}

func NewOrderService(orderRepository persistance.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository, tracer: otel.Tracer("Order Service")}
}
