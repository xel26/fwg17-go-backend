package routers

import (
	"coffe-shop-be-golang/src/controllers"
	"coffe-shop-be-golang/src/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup){
	authMiddleware, _ := middleware.Auth()
	
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", controllers.Register)
	r.POST("/forgot-password", controllers.ForgotPassword)
}