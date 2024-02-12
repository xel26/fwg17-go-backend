package routers

import (
	"coffe-shop-be-golang/src/controllers"
	controllers_customer "coffe-shop-be-golang/src/controllers/customer"
	"coffe-shop-be-golang/src/middleware"

	"github.com/gin-gonic/gin"
)

func CombineCustomer(r *gin.RouterGroup){
	AuthRouter(r.Group("/"))
	ProductsRouter(r.Group("/products"))
	r.GET("/testimonial", controllers.ListAllTestimonial)

	authMiddleware, _ := middleware.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	
	ProfileRouter(r.Group("/profile"))
	HistoryOrderRouter(r.Group("/history-order"))
	r.GET("/order-products", controllers_customer.ListOrderProducts)
	r.POST("/checkout", controllers_customer.Checkout)
	r.GET("/data-size", controllers_customer.GetPriceSize)
	r.GET("/data-variant", controllers_customer.GetPriceVariant)
}