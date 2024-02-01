package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func TagsRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllTags)
	r.GET("/:id", controllers.DetailTags)
	r.POST("", controllers.CreateTags)
	r.PATCH("/:id", controllers.UpdateTags)
	r.DELETE("/:id", controllers.DeleteTags)
}