package routers

import (
	controllers_customer "coffe-shop-be-golang/src/controllers/customer"

	"github.com/gin-gonic/gin"
)

func HistoryOrderRouter(r *gin.RouterGroup){
	r.GET("", controllers_customer.ListAllOrders)
	r.GET("/:id", controllers_customer.DetailOrder)
}