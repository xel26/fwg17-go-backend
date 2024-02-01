package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func MessageRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllMessage)
	r.GET("/:id", controllers.DetailMessage)
	r.POST("", controllers.CreateMessage)
	r.PATCH("/:id", controllers.UpdateMessage)
	r.DELETE("/:id", controllers.DeleteMessage)
}