package persistance

import (
	"errors"
	"gorm.io/gorm"
	"yt-gorm/src/application/domain/entity"
)

type orderRepository struct {
	db *gorm.DB
}

type OrderRepository interface {
	GetOrderById(id int64) *entity.Order
	CreateOrder(order entity.Order) *entity.Order
}

func (repo orderRepository) CreateOrder(order entity.Order) *entity.Order {
	result := repo.db.Create(&order)
	if result.Error != nil {
		panic(result.Error)
	}
	return &order
}

func (repo orderRepository) GetOrderById(id int64) *entity.Order {
	var order entity.Order
	result := repo.db.Preload("OrderLineItems").First(&order, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(result.Error)
	}

	return &order
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
