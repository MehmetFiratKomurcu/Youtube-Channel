package entity

import "time"

type Order struct {
	Id             int64           `gorm:"column:id;primaryKey"`
	ShipmentNumber int64           `gorm:"column:shipment_number"`
	CargoId        int             `gorm:"column:cargo_id"`
	IsShipped      bool            `gorm:"column:is_shipped"`
	CreatedAt      time.Time       `gorm:"column:created_at"`
	OrderLineItems []OrderLineItem `gorm:"referenceKey:OrderId"`
}
