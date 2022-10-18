package routes

import (
	"github.com/Abeldlp/go-and-compose/controllers"
	"github.com/gin-gonic/gin"
)

var ShopRoutes *gin.RouterGroup

func AppendShopRoutes(r *gin.Engine) {
	shop := r.Group("/shops")

	{
		shop.GET("/", controllers.GetAllShops)
		shop.GET("/:id", controllers.GetShopById)
		shop.POST("/", controllers.CreateShop)
		shop.DELETE("/:id", controllers.DeleteShop)
	}
}
