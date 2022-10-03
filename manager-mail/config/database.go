package config

import (
	models "github.com/Abeldlp/manager-mail/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	dns := "user:password@tcp(manager-mail-db:3306)/manager-mail?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.Mail{})

	DB = db
}
