package routes

import (
	"github.com/Abeldlp/go-and-compose/controllers"
	"github.com/gin-gonic/gin"
)

var InquiryRoutes *gin.RouterGroup

func AppendInquiryRoutes(r *gin.Engine) {

	inquiry := r.Group("/inquiries")

	{
		inquiry.GET("/", controllers.GetAllInquiries)
		inquiry.GET("/:id", controllers.GetInquiryById)
		inquiry.POST("/", controllers.CreateInquiry)
		inquiry.DELETE("/:id", controllers.DeleteInquiry)
	}
}
