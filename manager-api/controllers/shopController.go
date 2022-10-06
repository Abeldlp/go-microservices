package controllers

import (
	"net/http"

	"github.com/Abeldlp/go-and-compose/config"
	"github.com/Abeldlp/go-and-compose/models"
	"github.com/gin-gonic/gin"
)

func GetAllShops(c *gin.Context) {
	var shops []models.Shop
	config.DB.Find(&shops)

	c.JSON(http.StatusOK, shops)
}
