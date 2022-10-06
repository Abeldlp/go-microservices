package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName       string          `json:"product_name"`
	Price             int32           `json:"price"`
	ProductCategoryID int             `json:"product_category_id"`
	ProductCategory   ProductCategory `json:"product_category"`
}
