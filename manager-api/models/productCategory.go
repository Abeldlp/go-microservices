package models

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	CategoryName string `json:"category_name"`
}
