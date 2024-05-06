package services

import (
	"yt-gorm-scopes/src/application/domain/entity"
	"yt-gorm-scopes/src/application/domain/persistance"
	"yt-gorm-scopes/src/application/model"
)

type orderService struct {
	orderRepository persistance.OrderRepository
}

type OrderService interface {
	CreateOrder(command model.CreateOrderCommand) *entity.Order
	GetOrderById(id int64) *entity.Order
	ShipOrderByCargoCode(cargoCode string) error
}

func (service orderService) GetOrderById(id int64) *entity.Order {
	return service.orderRepository.GetOrderById(id)
}

func (service orderService) CreateOrder(command model.CreateOrderCommand) *entity.Order {
	order := model.MapToOrder(command)
	return service.orderRepository.CreateOrder(order)
}

func (service orderService) ShipOrderByCargoCode(cargoCode string) error {
	return service.orderRepository.ShipOrderByCargoCode(cargoCode)
}

func NewOrderService(orderRepository persistance.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}
