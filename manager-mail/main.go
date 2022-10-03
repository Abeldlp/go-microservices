package main

import (
	"net/http"

	"github.com/Abeldlp/manager-mail/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitializeDatabase()

	r := gin.Default()
	r.GET("/mail", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is the mail microservice",
		})
	})
	r.Run()
}
