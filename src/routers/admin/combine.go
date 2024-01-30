package routers

import "github.com/gin-gonic/gin"

func CombineAdmin(r *gin.RouterGroup){
	UserRouter(r.Group("/users"))
}