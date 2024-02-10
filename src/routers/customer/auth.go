package routers

import (
	controllers_customer "coffe-shop-be-golang/src/controllers/customer"
	"coffe-shop-be-golang/src/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup){
	authMiddleware, _ := middleware.Auth()
	
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", controllers_customer.Register)
	r.POST("/forgot-password", controllers_customer.ForgotPassword)
}