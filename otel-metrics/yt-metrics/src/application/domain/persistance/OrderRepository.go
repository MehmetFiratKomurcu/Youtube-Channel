package persistance

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"yt-otel-metrics/src/application/domain/entity"
)

type orderRepository struct {
	db     *gorm.DB
	tracer trace.Tracer
}

type OrderRepository interface {
	GetOrderById(ctx context.Context, id int64) *entity.Order
	CreateOrder(order entity.Order) *entity.Order
	ShipOrderByCargoCode(cargoCode string) error
}

func (repo orderRepository) CreateOrder(order entity.Order) *entity.Order {
	result := repo.db.Create(&order)
	if result.Error != nil {
		panic(result.Error)
	}
	return &order
}

func (repo orderRepository) GetOrderById(ctx context.Context, id int64) *entity.Order {
	ctx, span := repo.tracer.Start(ctx, "GetOrderById Repository")
	defer span.End()

	span.AddEvent("getting Order data from database")

	var order entity.Order
	result := repo.db.WithContext(ctx).Preload("OrderLineItems").First(&order, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(result.Error)
	}

	span.AddEvent("got Order data from database")

	return &order
}

func (repo orderRepository) ShipOrderByCargoCode(cargoCode string) error {
	result := repo.db.Model(&entity.Order{}).Scopes(GetCargoCodeById(cargoCode), NonCancelledOrders).Update("is_shipped", true)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetCargoCodeById(code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if code != "" {
			var cargo entity.Cargo
			if db.Session(&gorm.Session{}).Model(&cargo).First(&cargo, "code = ?", code).Error == nil {
				return db.Where("cargo_id = ?", cargo.Id)
			}
		}

		_ = db.AddError(errors.New("invalid Cargo Code"))
		return db
	}
}

func NonCancelledOrders(db *gorm.DB) *gorm.DB {
	return db.Where("is_cancelled = ?", false)
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db, tracer: otel.Tracer("OrderRepository")}
}
