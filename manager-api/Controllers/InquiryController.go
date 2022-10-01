package controllers

import (
	models "github.com/Abeldlp/go-and-compose/Models"
	"github.com/Abeldlp/go-and-compose/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllInquiries(c *gin.Context) {
	var inquiries []models.Inquiry
	config.DB.Order("created_at desc").Find(&inquiries)
	c.JSON(http.StatusOK, inquiries)
}

func CreateInquiry(c *gin.Context) {
	var inquiry models.Inquiry
	c.BindJSON(&inquiry)

	config.DB.Create(&models.Inquiry{Message: inquiry.Message, Assignee: inquiry.Assignee})

	c.JSON(http.StatusOK, gin.H{
		"message":  inquiry.Message,
		"assignee": inquiry.Assignee,
	})
}