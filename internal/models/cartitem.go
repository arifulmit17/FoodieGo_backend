package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CartID uuid.UUID
	Cart   Cart

	FoodID uuid.UUID
	Food   Food

	Quantity int `gorm:"default:1"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *CartItem) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}
