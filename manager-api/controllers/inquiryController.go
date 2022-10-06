package controllers

import (
	"github.com/Abeldlp/go-and-compose/config"
	"github.com/Abeldlp/go-and-compose/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllInquiries(c *gin.Context) {
	var inquiries []models.Inquiry
	config.DB.Order("created_at desc").Preload("Shop").Find(&inquiries)
	c.JSON(http.StatusOK, inquiries)
}

func GetInquiryById(c *gin.Context) {
	id := c.Param("id")
	var inquiry models.Inquiry
	config.DB.First(&inquiry, id)

	c.JSON(http.StatusOK, inquiry)
}

func CreateInquiry(c *gin.Context) {
	var inquiry models.Inquiry
	c.BindJSON(&inquiry)

	config.DB.Create(&models.Inquiry{
		ShopID:   inquiry.ShopID,
		Message:  inquiry.Message,
		Assignee: inquiry.Assignee,
		Deadline: inquiry.Deadline,
	})

	c.JSON(http.StatusOK, gin.H{
		"message":  inquiry.Message,
		"assignee": inquiry.Assignee,
	})
}

func UpdateInquiry(c *gin.Context) {
	id := c.Param("id")
	var inquiry models.Inquiry
	c.BindJSON(&inquiry)

	config.DB.First(&inquiry, id)

	config.DB.Model(&inquiry).Updates(models.Inquiry{
		ShopID:   inquiry.ShopID,
		Message:  inquiry.Message,
		Deadline: inquiry.Deadline,
		Assignee: inquiry.Assignee,
	})

	c.JSON(http.StatusOK, gin.H{
		"inquiry": id,
		"message": "updated succesfully",
	})
}

func DeleteInquiry(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Inquiry{}, id)

	c.JSON(http.StatusOK, gin.H{
		"deleted_inquiry": id,
	})
}
