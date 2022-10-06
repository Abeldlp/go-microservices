package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name               string    `json:"name"`
	Address            string    `json:"address"`
	ClosedOnWeekends   bool      `json:"closed_on_weekends"`
	HasIndoorSkatepark bool      `json:"has_indoor_skatepark"`
	Inquiries          []Inquiry `json:"inquiries"  gorm:"foreignKey:ShopID"`
}
