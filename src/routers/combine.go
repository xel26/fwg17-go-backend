package routers

import (
	"github.com/gin-gonic/gin"
	routers "github.com/xel26/fwg17-go-backend/src/routers/admin"
)

func Combine(r *gin.Engine){
	routers.CombineAdmin(r.Group("/admin"))
}