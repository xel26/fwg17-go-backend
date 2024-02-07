package routers

import (
	"coffe-shop-be-golang/src/middleware"

	"github.com/gin-gonic/gin"
)

func CombineCustomer(r *gin.RouterGroup){
	authMiddleware, _ := middleware.Auth()
	r.Use(authMiddleware.MiddlewareFunc())

	AuthRouter(r.Group("/"))
	ProductsRouter(r.Group("/products"))
	ProfileRouter(r.Group("/profile"))
}