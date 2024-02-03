package routers

import "github.com/gin-gonic/gin"

func CombineAdmin(r *gin.RouterGroup){
	UserRouter(r.Group("/users"))
	ProductsRouter(r.Group("/products"))
	CategoriesRouter(r.Group("/categories"))
	ForgotPasswordRouter(r.Group("/forgot-password"))
	MessageRouter(r.Group("/message"))
	OrderDetailsRouter(r.Group("/order-details"))
	OrdersRouter(r.Group("/orders"))
	ProductCategoriesRouter(r.Group("/product-categories"))
	ProductRatingsRouter(r.Group("/product-ratings"))
	ProductVariantsRouter(r.Group("/product-variants"))
	PromoRouter(r.Group("/promo"))
	SizesRouter(r.Group("/sizes"))
	TagsRouter(r.Group("/tags"))
	TestimonialRouter(r.Group("/testimonial"))
	VariantsRouter(r.Group("/variants"))
}