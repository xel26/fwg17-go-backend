package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProductCategoriesRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllProductCategories)
	r.GET("/:id", controllers.DetailProductCategories)
	r.POST("", controllers.CreateProductCategories)
	r.PATCH("/:id", controllers.UpdateProductCategories)
	r.DELETE("/:id", controllers.DeleteProductCategories)
}