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

func getShopById(c *gin.Context) {
	id := c.Param("id")
	var shop models.Shop

	config.DB.First(&shop, id)

	c.JSON(http.StatusOK, shop)
}

func CreateShop(c *gin.Context) {
	var shop models.Shop

	c.BindJSON(&shop)

	config.DB.Create(&models.Shop{
		Name:               shop.Name,
		Address:            shop.Address,
		ClosedOnWeekends:   shop.ClosedOnWeekends,
		HasIndoorSkatepark: shop.HasIndoorSkatepark,
	})
}

func DeleteShop(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Inquiry{}, id)

	c.JSON(http.StatusOK, gin.H{
		"deleted_inquiry": id,
	})
}
