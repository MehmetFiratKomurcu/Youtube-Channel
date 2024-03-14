package services

import (
	"yt-gorm/src/application/domain/entity"
	"yt-gorm/src/application/domain/persistance"
	"yt-gorm/src/application/model"
)

type orderService struct {
	orderRepository persistance.OrderRepository
}

type OrderService interface {
	CreateOrder(command model.CreateOrderCommand) *entity.Order
	GetOrderById(id int64) *entity.Order
}

func (service orderService) GetOrderById(id int64) *entity.Order {
	return service.orderRepository.GetOrderById(id)
}

func (service orderService) CreateOrder(command model.CreateOrderCommand) *entity.Order {
	order := model.MapToOrder(command)
	return service.orderRepository.CreateOrder(order)
}

func NewOrderService(orderRepository persistance.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}
