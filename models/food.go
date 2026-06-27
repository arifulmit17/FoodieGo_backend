package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model

	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}
