package entity

import "time"

type Cargo struct {
	Id          int32     `gorm:"column:id;primaryKey"`
	Code        string    `gorm:"column:code"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
