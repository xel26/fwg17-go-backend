package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriesRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllCategories)
	r.GET("/:id", controllers.DetailCategories)
	r.POST("", controllers.CreateCategories)
	r.PATCH("/:id", controllers.UpdateCategories)
	r.DELETE("/:id", controllers.DeleteCategories)
}