package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(r *gin.RouterGroup){
	r.GET("/:id", controllers.DetailUser)
	r.PATCH("/:id", controllers.UpdateUser)
}