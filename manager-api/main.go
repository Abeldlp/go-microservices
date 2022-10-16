package main

import (
	"github.com/Abeldlp/go-and-compose/config"
	"github.com/Abeldlp/go-and-compose/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.InitializeDatabase()
	r := config.BuildRouter()

	//Routes
	routes.AppendInquiryRoutes(r)
	routes.AppendShopRoutes(r)

	r.Run()
}
