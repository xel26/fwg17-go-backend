package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func OrdersRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllOrders)
	r.GET("/:id", controllers.DetailOrders)
	r.POST("", controllers.CreateOrders)
	r.PATCH("/:id", controllers.UpdateOrders)
	r.DELETE("/:id", controllers.DeleteOrders)
}