package main

import (
	"github.com/Abeldlp/go-and-compose/config"
	"github.com/Abeldlp/go-and-compose/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitializeDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	inquiry := r.Group("/inquiries")

	{
		inquiry.GET("/", controllers.GetAllInquiries)
		inquiry.GET("/:id", controllers.GetInquiryById)
		inquiry.POST("/", controllers.CreateInquiry)
		inquiry.DELETE("/:id", controllers.DeleteInquiry)
	}

	r.Run()
}
