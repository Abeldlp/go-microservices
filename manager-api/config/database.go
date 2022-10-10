package config

import (
	"github.com/Abeldlp/go-and-compose/models"
	// "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitializeDatabase() {
	// dsn := "user:password@tcp(manager-api-db:3306)/manager-api?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := "host=" + os.Getenv("DATABASE_HOST") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PASSWORD") + " dbname=" + os.Getenv("DATABASE_NAME") + " port=" + os.Getenv("DATABASE_PORT") + " sslmode=disable TimeZone=Europe/Amsterdam"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
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
