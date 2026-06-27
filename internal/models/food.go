package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Food struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Price       float64   `gorm:"type:decimal(10,2)" json:"price"`
	ImageURL    string    `json:"image_url"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (f *Food) BeforeCreate(tx *gorm.DB) error {
	f.ID = uuid.New()
	return nil
}
