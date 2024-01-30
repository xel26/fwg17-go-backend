package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xel26/fwg17-go-backend/src/controllers"
)

func UserRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllUsers)
	r.GET("/:id", controllers.DetailUser)
	r.POST("", controllers.CreateUser)
	r.PATCH("/:id", controllers.UpdateUser)
	r.DELETE("/:id", controllers.DeleteUser)
}