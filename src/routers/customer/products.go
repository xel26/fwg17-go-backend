package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProductsRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllProducts)
	r.GET("/:id", controllers.DetailProducts)
}