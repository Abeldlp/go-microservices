package config

import (
	"fmt"
	"github.com/Abeldlp/go-and-compose/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitializeDatabase() {
	DATABASE_HOST := os.Getenv("DATABASE_HOST")
	DATABASE_USER := os.Getenv("DATABASE_USER")
	DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")
	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	DATABASE_PORT := os.Getenv("DATABASE_PORT")

	dsn := "host=" + DATABASE_HOST + " user=" + DATABASE_USER + " password=" + DATABASE_PASSWORD + " dbname=" + DATABASE_NAME + " port=" + DATABASE_PORT + " sslmode=disable TimeZone=Europe/Amsterdam"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("DATABASE_HOST")
		fmt.Println(DATABASE_HOST)
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&models.Shop{},
		&models.Inquiry{},
		&models.ProductCategory{},
		&models.Product{},
	)

	DB = db
}
