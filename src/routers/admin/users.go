package routers

import (
	"coffe-shop-be-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllUsers)
	r.GET("/:id", controllers.DetailUser)
	r.POST("", controllers.CreateUser)
	r.PATCH("/:id", controllers.UpdateUser)
	r.DELETE("/:id", controllers.DeleteUser)
}