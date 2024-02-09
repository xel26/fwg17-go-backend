package routers

import (
	controllers_customer "coffe-shop-be-golang/src/controllers/customer"
	"coffe-shop-be-golang/src/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(r *gin.RouterGroup){
	authMiddleware, _ := middleware.Auth()
	r.Use(authMiddleware.MiddlewareFunc())

	r.GET("", controllers_customer.GetProfile)
	r.PATCH("", controllers_customer.UpdateProfile)
}