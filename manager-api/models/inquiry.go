package models

import (
	// "github.com/Abeldlp/go-and-compose/config"
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

// func (i *Inquiry) Get() {
// 	config.DB.Order("created_at desc").Preload("Shop").Find(&i)
// }
