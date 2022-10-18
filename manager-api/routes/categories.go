package routes

import (
	"github.com/Abeldlp/go-and-compose/controllers"
	"github.com/gin-gonic/gin"
)

var CategoryRoutes *gin.RouterGroup

func AppendCategoryRoutes(r *gin.Engine) {
	category := r.Group("/categories")

	{
		category.GET("/", controllers.GetAllCategories)
		category.GET("/:id", controllers.GetCategoryById)
	}
}
