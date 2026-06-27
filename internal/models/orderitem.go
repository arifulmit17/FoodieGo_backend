package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrderID uuid.UUID
	Order   Order

	FoodID uuid.UUID
	Food   Food

	Quantity int

	Price float64 `gorm:"type:decimal(10,2)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *OrderItem) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New()
	return nil
}
