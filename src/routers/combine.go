package routers

import (
	routers "coffe-shop-be-golang/src/routers/admin"

	"github.com/gin-gonic/gin"
)

func Combine(r *gin.Engine){
	routers.CombineAdmin(r.Group("/admin"))
}