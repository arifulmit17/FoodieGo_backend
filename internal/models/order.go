package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	UserID uuid.UUID
	User   User

	Address string

	TotalPrice float64 `gorm:"type:decimal(10,2)"`

	Status string `gorm:"default:Pending"`

	Items []OrderItem `gorm:"foreignKey:OrderID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New()
	return nil
}
