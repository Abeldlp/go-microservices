package main

import (
	"github.com/Abeldlp/go-and-compose/config"
	"github.com/Abeldlp/go-and-compose/routes"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	env := os.Getenv("PROD_ENV")

	if env == "" {
		godotenv.Load(".env")
	}

	config.InitializeDatabase()
	r := config.BuildRouter()

	//Routes
	routes.AppendInquiryRoutes(r)
	routes.AppendShopRoutes(r)
	routes.AppendCategoryRoutes(r)

	r.Run()
}
