package routers

import (
	admin "coffe-shop-be-golang/src/routers/admin"
	customer "coffe-shop-be-golang/src/routers/customer"

	"github.com/gin-gonic/gin"
)

func Combine(r *gin.Engine){
	admin.CombineAdmin(r.Group("/admin"))
	customer.CombineCustomer(r.Group("/"))
}