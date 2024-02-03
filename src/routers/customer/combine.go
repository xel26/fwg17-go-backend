package routers

import "github.com/gin-gonic/gin"

func CombineCustomer(r *gin.RouterGroup){
	AuthRouter(r.Group("/"))
}