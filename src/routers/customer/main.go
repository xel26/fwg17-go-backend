package routers

import (
	"coffe-shop-be-golang/src/controllers"
	controllers_customer "coffe-shop-be-golang/src/controllers/customer"
	"coffe-shop-be-golang/src/middleware"

	"github.com/gin-gonic/gin"
)

func CombineCustomer(r *gin.RouterGroup){
	AuthRouter(r.Group("/"))

	authMiddleware, _ := middleware.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	
	ProductsRouter(r.Group("/products"))
	ProfileRouter(r.Group("/profile"))
	r.GET("/history-order", controllers.ListAllOrders)
	r.POST("/checkout", controllers_customer.Checkout)
}