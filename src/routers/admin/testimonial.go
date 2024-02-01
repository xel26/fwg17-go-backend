package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func TestimonialRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllTestimonial)
	r.GET("/:id", controllers.DetailTestimonial)
	r.POST("", controllers.CreateTestimonial)
	r.PATCH("/:id", controllers.UpdateTestimonial)
	r.DELETE("/:id", controllers.DeleteTestimonial)
}