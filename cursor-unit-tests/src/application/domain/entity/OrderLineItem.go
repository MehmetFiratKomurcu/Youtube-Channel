package entity

type OrderLineItem struct {
	Id        int64 `gorm:"column:id;primaryKey"`
	ProductId int64 `gorm:"column:product_id"`
	SellerId  int64 `gorm:"column:seller_id"`
	OrderId   int64 `gorm:"column:order_id"`
}
