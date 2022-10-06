package models

import (
	"gorm.io/gorm"
)

type Inquiry struct {
	gorm.Model
	Message  string `json:"message"`
	Deadline string `json:"deadline"`
	Assignee string `json:"assignee"`
	ShopID   uint   `json:"shop_id"`
	Shop     Shop   `json:"shop"`
}
