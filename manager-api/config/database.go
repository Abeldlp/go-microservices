package config

import (
	models "github.com/Abeldlp/go-and-compose/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	dsn := "user:password@tcp(manager-api-db:3306)/manager-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Inquiry{})

	DB = db
}
