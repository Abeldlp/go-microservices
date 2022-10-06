package main

import (
	"github.com/Abeldlp/go-and-compose/config"
	"github.com/Abeldlp/go-and-compose/routes"
)

func main() {
	config.InitializeDatabase()
	r := config.BuildRouter()

	//Routes
	routes.AppendInquiryRoutes(r)

	r.Run()
}
