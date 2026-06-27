package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	UserID uuid.UUID
	User   User

	Items []CartItem `gorm:"foreignKey:CartID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}
