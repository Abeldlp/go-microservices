package controllers

import (
	"net/http"

	"github.com/Abeldlp/go-and-compose/config"
	"github.com/Abeldlp/go-and-compose/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var categories []models.ProductCategory
	config.DB.Find(&categories)

	c.JSON(http.StatusOK, categories)
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	var category models.ProductCategory

	config.DB.First(&category, id)

	c.JSON(http.StatusOK, category)
}

func CreateCategory(c *gin.Context) {
	var category models.ProductCategory
	c.BindJSON(&category)

	config.DB.Create(&models.ProductCategory{
		CategoryName: category.CategoryName,
	})
}
