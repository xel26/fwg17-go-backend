package routers

import (
	controllers_customer "coffe-shop-be-golang/src/controllers/customer"

	"github.com/gin-gonic/gin"
)

func IntRandRouter(r *gin.RouterGroup){
	r.GET("", controllers_customer.FindOneIntRandom)
	r.DELETE("", controllers_customer.DeleteIntRandom)
}