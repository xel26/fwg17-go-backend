package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProductVariantsRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllProductVariants)
	r.GET("/:id", controllers.DetailProductVariant)
	r.POST("", controllers.CreateProductVariants)
	r.PATCH("/:id", controllers.UpdateProductVariants)
	r.DELETE("/:id", controllers.DeleteProductVariants)
}